package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"microservices/data"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)

		return
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)

		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)

		return
	}

	prod.ToJSON(rw)
}
