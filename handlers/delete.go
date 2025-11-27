package handlers

import (
	"net/http"

	"microservices/data"
)

// DeleteProduct deletes a product
//
//	@Description	deletes a product from the data store
//	@Tags			products
//
//	@Param			id	path	int	true	"Product ID"
//
//	@Success		201
//	@Router			/{id} [delete]
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] Deleting record id", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)

		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)

		return
	}
}
