package handlers

import (
	"log"
	"microservices/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw)

		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter) {
	lp := data.GetProducts()

	error := lp.ToJSON(rw)

	if error != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
