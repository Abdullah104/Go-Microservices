//	@Title			of Product API
//	@Description	Documentation for Product API

//	@Schemes	http
//	@BasePath	/
//	@Version	1.0.0

//	@Accept		json
//	@Produce	json

package handlers

import (
	"context"
	"fmt"
	"log"
	"microservices/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products is an HTTP handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new Products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GenericError is a generic error message returned by the server
type GenericError struct {
	Message string `json:"message"`
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)

			return
		}

		// Validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product")
			http.Error(rw, fmt.Sprintf("Error reading product %s", err), http.StatusBadRequest)

			return
		}

		// Add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(rw, req)
	})
}

// getProductId returns the product's id from the url
// Panics if cannot convert the id into an integer
// This should never happen as the router ensures that this is a valid number
func GetProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
