# \CompanyApi

All URIs are relative to *http://localhost:8080/backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompany**](CompanyApi.md#CreateCompany) | **Post** /rest/platform/company | Create a new Company
[**GetCompany**](CompanyApi.md#GetCompany) | **Get** /rest/platform/company | Get your Company
[**UpdateCompany**](CompanyApi.md#UpdateCompany) | **Put** /rest/platform/company | Update Company data
[**UpdateCompanyLegal**](CompanyApi.md#UpdateCompanyLegal) | **Put** /rest/platform/company/legal | Update legal Company info
[**UpdateCompanyTopic**](CompanyApi.md#UpdateCompanyTopic) | **Put** /rest/platform/company/topic | Update selected topics for your Company



## CreateCompany

> PlatformCompany CreateCompany(ctx).Company(company).Legal(legal).Admin(admin).Execute()

Create a new Company



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
    company := TODO // PlatformCompanyCreate | 
    legal := TODO // PlatformCompanyCreateLegal | 
    admin := TODO // PlatformCompanyCreateAdmin | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompanyApi.CreateCompany(context.Background()).Company(company).Legal(legal).Admin(admin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.CreateCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompany`: PlatformCompany
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.CreateCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **company** | [**PlatformCompanyCreate**](PlatformCompanyCreate.md) |  | 
 **legal** | [**PlatformCompanyCreateLegal**](PlatformCompanyCreateLegal.md) |  | 
 **admin** | [**PlatformCompanyCreateAdmin**](PlatformCompanyCreateAdmin.md) |  | 

### Return type

[**PlatformCompany**](PlatformCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> PlatformCompany GetCompany(ctx).Execute()

Get your Company



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
    resp, r, err := api_client.CompanyApi.GetCompany(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.GetCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompany`: PlatformCompany
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.GetCompany`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


### Return type

[**PlatformCompany**](PlatformCompany.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompany

> PlatformCompany UpdateCompany(ctx).Company(company).Logo(logo).Execute()

Update Company data



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
    company := TODO // PlatformCompanyUpdate | 
    logo := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompanyApi.UpdateCompany(context.Background()).Company(company).Logo(logo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.UpdateCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompany`: PlatformCompany
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.UpdateCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **company** | [**PlatformCompanyUpdate**](PlatformCompanyUpdate.md) |  | 
 **logo** | ***os.File** |  | 

### Return type

[**PlatformCompany**](PlatformCompany.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanyLegal

> PlatformCompany UpdateCompanyLegal(ctx).Legal(legal).Execute()

Update legal Company info



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
    legal := TODO // PlatformCompanyUpdateLegal | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompanyApi.UpdateCompanyLegal(context.Background()).Legal(legal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.UpdateCompanyLegal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompanyLegal`: PlatformCompany
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.UpdateCompanyLegal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyLegalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legal** | [**PlatformCompanyUpdateLegal**](PlatformCompanyUpdateLegal.md) |  | 

### Return type

[**PlatformCompany**](PlatformCompany.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanyTopic

> PlatformCompany UpdateCompanyTopic(ctx).Sdg(sdg).Execute()

Update selected topics for your Company



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
    sdg := TODO // PlatformCompanyUpdateTopic | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompanyApi.UpdateCompanyTopic(context.Background()).Sdg(sdg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.UpdateCompanyTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompanyTopic`: PlatformCompany
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.UpdateCompanyTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sdg** | [**PlatformCompanyUpdateTopic**](PlatformCompanyUpdateTopic.md) |  | 

### Return type

[**PlatformCompany**](PlatformCompany.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

