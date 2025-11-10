package handlers

import (
	"log"
	"microservices/data"
	"net/http"
	"regexp"
	"strconv"
)

// Products is an HTTP handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new Products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw)

		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)

		return
	}

	if r.Method == http.MethodPut {
		// expect the ID in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")

			http.Error(rw, "Invalid URI", http.StatusBadRequest)

			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group", g)

			http.Error(rw, "Invalid URI", http.StatusBadRequest)

			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("Invalid URI cannot convert to number", idString)

			http.Error(rw, "Invalid URI", http.StatusBadRequest)

			return
		}

		p.updateProduct(id, rw, r)

		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(rw http.ResponseWriter) {
	// fetch the products from the data store
	lp := data.GetProducts()

	// serialize the list into JSON
	error := lp.ToJSON(rw)

	if error != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)

		return
	}

	data.AddProduct(prod)

	prod.ToJSON(rw)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)

		return
	}

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
