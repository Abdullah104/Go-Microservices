package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"microservices/data"
)

// @Summary		Update Product
// @Description	Update a product
// @Tags			products
// @Success		200	{object}	data.Product	"Product after update"
// @Failure		404	"Product not found"
// @Failure		500	"Product not found"
// @Router			/products/:id [put]
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

	rw.Header().Set("Content-Type", "application/json")

	prod.ToJSON(rw)
}
