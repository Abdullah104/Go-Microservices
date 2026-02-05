package main

import (
	"microservices/currency/protos/currency"
	"microservices/currency/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	currency.RegisterCurrencyServer(gs, cs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
