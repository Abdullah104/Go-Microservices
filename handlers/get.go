package handlers

import (
	"microservices/data"
	"net/http"
)

// @Summary		List Products
// @Description	returns the products from the data store
// @Tags			products
// @Success		200	{object}	data.Products	"a list of products in the response"
// @Failure		500	"Unable to marshal JSON"
// @Router			/products [get]
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")

	// fetch the products from the data store
	lp := data.GetProducts()
	rw.Header().Set("Content-Type", "application/json")

	// Serialize the list to JSON
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// @Summary		List Single Product
// @Description	returns a product with the given ID from the data store
// @Tags			products
// @Produce		json
// @Success		200	{object}	data.Product	"A Product in the response"
// @Failure		404	"Product not found"
// @Router			/products/:id [get]
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET product")

	id := GetProductID(r)

	// Fetch the product from the data store
	prod, _, err := data.FindProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)

		return
	}
	rw.Header().Set("Content-Type", "application/json")

	data.ToJSON(prod, rw)
}
