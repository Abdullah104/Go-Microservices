package handlers

import (
	"microservices/data"
	"net/http"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(prod)

	prod.ToJSON(rw)
}
