# \ProductsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsGet**](ProductsAPI.md#ProductsGet) | **Get** /products | List Products
[**ProductsIdDelete**](ProductsAPI.md#ProductsIdDelete) | **Delete** /products/{id} | Delete Product
[**ProductsIdGet**](ProductsAPI.md#ProductsIdGet) | **Get** /products/:id | List Single Product
[**ProductsIdPut**](ProductsAPI.md#ProductsIdPut) | **Put** /products/:id | Update Product
[**ProductsPost**](ProductsAPI.md#ProductsPost) | **Post** /products | Create Product



## ProductsGet

> []DataProduct ProductsGet(ctx).Execute()

List Products



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Abdullah104/Go-Microservices/products_client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsGet`: []DataProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductsGetRequest struct via the builder pattern


### Return type

[**[]DataProduct**](DataProduct.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsIdDelete

> ProductsIdDelete(ctx, id).Execute()

Delete Product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Abdullah104/Go-Microservices/products_client"
)

func main() {
	id := int32(56) // int32 | Product ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductsAPI.ProductsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsIdGet

> DataProduct ProductsIdGet(ctx).Execute()

List Single Product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Abdullah104/Go-Microservices/products_client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsIdGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsIdGet`: DataProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsIdGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductsIdGetRequest struct via the builder pattern


### Return type

[**DataProduct**](DataProduct.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsIdPut

> DataProduct ProductsIdPut(ctx).Execute()

Update Product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Abdullah104/Go-Microservices/products_client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsIdPut(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsIdPut`: DataProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsIdPut`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductsIdPutRequest struct via the builder pattern


### Return type

[**DataProduct**](DataProduct.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsPost

> DataProduct ProductsPost(ctx).Execute()

Create Product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Abdullah104/Go-Microservices/products_client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsPost`: DataProduct
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductsPostRequest struct via the builder pattern


### Return type

[**DataProduct**](DataProduct.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

