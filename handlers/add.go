package handlers

import (
	"microservices/data"
	"net/http"
)

// @Summary		Create Product
// @Description	Add a new product to the data store
// @Tags			products
// @Success		200	{object}	data.Product	"The newly created product"
// @Router			/ [post]
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(prod)

	prod.ToJSON(rw)
}
