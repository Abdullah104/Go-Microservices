package handlers

import (
	"microservices/data"
	"net/http"
)

// GetProducts returns a list of products
//
//	@Description	returns the products from the data store
//	@Tags			products
//	@Success		200	{object}	data.Products	"a list of products in the response"
//	@Router			/ [get]
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")

	// fetch the products from the data store
	lp := data.GetProducts()

	// Serialize the list to JSON
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
