# \InvoiceApi

All URIs are relative to *http://localhost:8080/backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoice**](InvoiceApi.md#GetInvoice) | **Get** /rest/platform/invoice/{token} | Get a single invoice
[**GetInvoices**](InvoiceApi.md#GetInvoices) | **Get** /rest/platform/invoice | Get a set of invoices



## GetInvoice

> PlatformInvoice GetInvoice(ctx, token).Execute()

Get a single invoice



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoiceApi.GetInvoice(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: PlatformInvoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformInvoice**](PlatformInvoice.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> PlatformInvoiceItemWrapper GetInvoices(ctx).NumberLike(numberLike).Page(page).Size(size).Sort(sort).Execute()

Get a set of invoices



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    numberLike := "numberLike_example" // string | The name of the filter (optional)
    page := int32(56) // int32 | The number of the page. Possible values: 0 to unlimited (optional)
    size := int32(56) // int32 | The number of items inside a page. Possible values: 1 to 100 (optional)
    sort := "sort_example" // string | The order direction for sorted results. Possible values: ASC or DESC (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoiceApi.GetInvoices(context.Background()).NumberLike(numberLike).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoices`: PlatformInvoiceItemWrapper
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **numberLike** | **string** | The name of the filter | 
 **page** | **int32** | The number of the page. Possible values: 0 to unlimited | 
 **size** | **int32** | The number of items inside a page. Possible values: 1 to 100 | 
 **sort** | **string** | The order direction for sorted results. Possible values: ASC or DESC | 

### Return type

[**PlatformInvoiceItemWrapper**](PlatformInvoiceItemWrapper.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

