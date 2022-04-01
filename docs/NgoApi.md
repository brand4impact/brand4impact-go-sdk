# \NgoApi

All URIs are relative to *http://localhost:8080/backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNgo**](NgoApi.md#CreateNgo) | **Post** /rest/platform/ngo | Create a new NGO
[**GetNgo**](NgoApi.md#GetNgo) | **Get** /rest/platform/ngo | Get your NGO
[**UpdateNgo**](NgoApi.md#UpdateNgo) | **Put** /rest/platform/ngo | Update NGO data
[**UpdateNgoBank**](NgoApi.md#UpdateNgoBank) | **Put** /rest/platform/ngo/bank | Update bank NGO info
[**UpdateNgoLegal**](NgoApi.md#UpdateNgoLegal) | **Put** /rest/platform/ngo/legal | Update legal NGO info
[**UploadNgoFiles**](NgoApi.md#UploadNgoFiles) | **Put** /rest/platform/ngo/upload | Upload NGO files



## CreateNgo

> PlatformNgo CreateNgo(ctx).Ngo(ngo).Legal(legal).Admin(admin).Execute()

Create a new NGO



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
    ngo := TODO // PlatformNgoCreate | 
    legal := TODO // PlatformNgoCreateLegal | 
    admin := TODO // PlatformNgoCreateAdmin | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NgoApi.CreateNgo(context.Background()).Ngo(ngo).Legal(legal).Admin(admin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgoApi.CreateNgo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNgo`: PlatformNgo
    fmt.Fprintf(os.Stdout, "Response from `NgoApi.CreateNgo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNgoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ngo** | [**PlatformNgoCreate**](PlatformNgoCreate.md) |  | 
 **legal** | [**PlatformNgoCreateLegal**](PlatformNgoCreateLegal.md) |  | 
 **admin** | [**PlatformNgoCreateAdmin**](PlatformNgoCreateAdmin.md) |  | 

### Return type

[**PlatformNgo**](PlatformNgo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNgo

> PlatformNgo GetNgo(ctx).Execute()

Get your NGO



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NgoApi.GetNgo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgoApi.GetNgo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNgo`: PlatformNgo
    fmt.Fprintf(os.Stdout, "Response from `NgoApi.GetNgo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNgoRequest struct via the builder pattern


### Return type

[**PlatformNgo**](PlatformNgo.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNgo

> PlatformNgo UpdateNgo(ctx).Ngo(ngo).Logo(logo).Execute()

Update NGO data



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
    ngo := TODO // PlatformNgoUpdate | 
    logo := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NgoApi.UpdateNgo(context.Background()).Ngo(ngo).Logo(logo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgoApi.UpdateNgo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNgo`: PlatformNgo
    fmt.Fprintf(os.Stdout, "Response from `NgoApi.UpdateNgo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNgoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ngo** | [**PlatformNgoUpdate**](PlatformNgoUpdate.md) |  | 
 **logo** | ***os.File** |  | 

### Return type

[**PlatformNgo**](PlatformNgo.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNgoBank

> PlatformNgo UpdateNgoBank(ctx).Bank(bank).Execute()

Update bank NGO info



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
    bank := TODO // PlatformNgoUpdateBank | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NgoApi.UpdateNgoBank(context.Background()).Bank(bank).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgoApi.UpdateNgoBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNgoBank`: PlatformNgo
    fmt.Fprintf(os.Stdout, "Response from `NgoApi.UpdateNgoBank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNgoBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bank** | [**PlatformNgoUpdateBank**](PlatformNgoUpdateBank.md) |  | 

### Return type

[**PlatformNgo**](PlatformNgo.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNgoLegal

> PlatformNgo UpdateNgoLegal(ctx).Legal(legal).Execute()

Update legal NGO info



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
    legal := TODO // PlatformNgoUpdateLegal | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NgoApi.UpdateNgoLegal(context.Background()).Legal(legal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgoApi.UpdateNgoLegal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNgoLegal`: PlatformNgo
    fmt.Fprintf(os.Stdout, "Response from `NgoApi.UpdateNgoLegal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNgoLegalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legal** | [**PlatformNgoUpdateLegal**](PlatformNgoUpdateLegal.md) |  | 

### Return type

[**PlatformNgo**](PlatformNgo.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadNgoFiles

> PlatformNgo UploadNgoFiles(ctx).File(file).Execute()

Upload NGO files



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
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NgoApi.UploadNgoFiles(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NgoApi.UploadNgoFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadNgoFiles`: PlatformNgo
    fmt.Fprintf(os.Stdout, "Response from `NgoApi.UploadNgoFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadNgoFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**PlatformNgo**](PlatformNgo.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

