package main

import (
	"context"
	"log"
	"microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "microservices/swagger"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

// @title			Products API
// @version		1.0
// @description	A simple products management API
// @host			localhost:9090
// @BasePath		/
func main() {
	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Create the handlers
	ph := handlers.NewProducts(l)

	// Create a new serve mux and register the handlers
	sm := mux.NewRouter()

	idPath := "/products/{id:[0-9]+}"

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc(idPath, ph.GetProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.Use(ph.MiddlewareProductValidation)
	putRouter.HandleFunc(idPath, ph.UpdateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.Use(ph.MiddlewareProductValidation)
	postRouter.HandleFunc("/products", ph.AddProduct)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(idPath, ph.DeleteProduct)

	opts := middleware.RedocOpts{SpecURL: "/swagger/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)

	directory := http.Dir("./swagger")
	fileServer := http.FileServer(directory)
	getRouter.PathPrefix("/swagger").Handler(http.StripPrefix("/swagger", fileServer))

	// Create a new server
	s := &http.Server{
		Addr:         *bindAddress,      // Configure the bind address
		Handler:      sm,                // Set the default handler
		ErrorLog:     l,                 // Set the logger for the server
		ReadTimeout:  1 * time.Second,   // Max time to read request from the client
		WriteTimeout: 1 * time.Second,   // Max time to write response to the client
		IdleTimeout:  120 * time.Second, // Max time for connections using TCP Keep-Alive
	}

	// Start the server
	go func() {
		error := s.ListenAndServe()

		if error != nil && error != http.ErrServerClosed {
			l.Fatal(error)
		}
	}()

	// Trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // Important: release resources

	if err := s.Shutdown(tc); err != nil {
		l.Printf("Server shutdown error: %v\n", err)
	}

	l.Println("Server stopped")
}
