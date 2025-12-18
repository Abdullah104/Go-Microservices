package main

import (
	"fmt"
	pc "products_client"
	"testing"
)

func TestOurClient(t *testing.T) {
	configuration := pc.NewConfiguration()
	configuration.Scheme = "http"
	configuration.Host = "localhost:9090"

	client := pc.NewAPIClient(configuration)

	service := client.ProductsAPI

	p, r, e := service.ProductsGetExecute(pc.ApiProductsGetRequest{})

	fmt.Println(p)
	fmt.Println(r)
	fmt.Println(e)
}
