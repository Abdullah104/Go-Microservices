package handlers

import (
	"net/http"

	"microservices/data"
)

// @Summary		Delete Product
// @Description	deletes a product from the data store
// @Tags			products
// @Param			id	path	int	true	"Product ID"
// @Success		204
// @Failure		404	{object}	GenericError
// @Failure		500	{object}	GenericError
// @Router			/products/{id} [delete]
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := GetProductID(r)

	p.l.Println("[DEBUG] Deleting record id", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		rw.WriteHeader(http.StatusNotFound)
		rw.Header().Set("Content-Type", "application/json")

		data.ToJSON(&GenericError{Message: err.Error()}, rw)

		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Header().Set("Content-Type", "application/json")

		data.ToJSON(&GenericError{Message: err.Error()}, rw)

		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
