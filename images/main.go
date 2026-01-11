package main

import (
	"context"
	"microservices/images/files"
	"microservices/images/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")
var logLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")
var basePath = env.String("BASE_PATH", false, "./imagestore", "Base path to save images")

func main() {
	env.Parse()

	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString(*logLevel),
		},
	)

	// create a logger for the server from the default logger
	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	// create the storage class, use local storage
	// max filesize 5MB
	stor, err := files.NewLocal(*basePath, 1024*1000*5)
	if err != nil {
		l.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

	// Create the handlers.
	fh := handlers.NewFiles(stor, l)

	// create a new serve mux and register handlers
	sm := mux.NewRouter()

	path := "/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}"

	// filename regex: {filename:[a-zA-Z]+\\.[a-z]{3}}
	// problem with file server is that it is dumb
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc(path, fh.ServeHTTP)

	// get files
	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.HandleFunc(path, http.StripPrefix("/images/", http.FileServer(http.Dir(*basePath))).ServeHTTP)

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,
		Handler:      sm,
		ErrorLog:     sl,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		error := s.ListenAndServe()

		if error != nil && error != http.ErrServerClosed {
			sl.Fatal(error)
		}
	}()

	// Trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	sl.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(tc); err != nil {
		sl.Printf("Server shutdown error: %v\n", err)
	}

	sl.Println("Server stopped")
}
