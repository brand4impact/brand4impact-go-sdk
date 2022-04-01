# \CampaignApi

All URIs are relative to *http://localhost:8080/backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignApi.md#CreateCampaign) | **Post** /rest/platform/campaign | Create a new campaign
[**GetCampaign**](CampaignApi.md#GetCampaign) | **Get** /rest/platform/campaign/{token} | Get a single campaign
[**GetCampaigns**](CampaignApi.md#GetCampaigns) | **Get** /rest/platform/campaign | Get a set of campaigns
[**SwitchCampaign**](CampaignApi.md#SwitchCampaign) | **Put** /rest/platform/campaign/{token}/switch | Switch campaign status
[**UpdateCampaign**](CampaignApi.md#UpdateCampaign) | **Put** /rest/platform/campaign/{token} | Update a campaign



## CreateCampaign

> PlatformCampaign CreateCampaign(ctx).Campaign(campaign).PrimaryImage(primaryImage).SecondaryImage(secondaryImage).Execute()

Create a new campaign



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
    campaign := TODO // PlatformCampaignCreate | 
    primaryImage := os.NewFile(1234, "some_file") // *os.File | 
    secondaryImage := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignApi.CreateCampaign(context.Background()).Campaign(campaign).PrimaryImage(primaryImage).SecondaryImage(secondaryImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.CreateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaign`: PlatformCampaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaign** | [**PlatformCampaignCreate**](PlatformCampaignCreate.md) |  | 
 **primaryImage** | ***os.File** |  | 
 **secondaryImage** | ***os.File** |  | 

### Return type

[**PlatformCampaign**](PlatformCampaign.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> PlatformCampaign GetCampaign(ctx, token).Execute()

Get a single campaign



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
    token := "token_example" // string | The unique token for the campaign.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignApi.GetCampaign(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.GetCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaign`: PlatformCampaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique token for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformCampaign**](PlatformCampaign.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> PlatformCampaignItemWrapper GetCampaigns(ctx).TitleFilter(titleFilter).Page(page).Size(size).Sort(sort).Execute()

Get a set of campaigns



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
    titleFilter := "titleFilter_example" // string | The name of the filter (optional)
    page := int32(56) // int32 | The number of the page. Possible values: 0 to unlimited (optional)
    size := int32(56) // int32 | The number of items inside a page. Possible values: 1 to 100 (optional)
    sort := "sort_example" // string | The order direction for sorted results. Possible values: ASC or DESC (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignApi.GetCampaigns(context.Background()).TitleFilter(titleFilter).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.GetCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaigns`: PlatformCampaignItemWrapper
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **titleFilter** | **string** | The name of the filter | 
 **page** | **int32** | The number of the page. Possible values: 0 to unlimited | 
 **size** | **int32** | The number of items inside a page. Possible values: 1 to 100 | 
 **sort** | **string** | The order direction for sorted results. Possible values: ASC or DESC | 

### Return type

[**PlatformCampaignItemWrapper**](PlatformCampaignItemWrapper.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchCampaign

> PlatformCampaign SwitchCampaign(ctx, token).Execute()

Switch campaign status



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
    resp, r, err := api_client.CampaignApi.SwitchCampaign(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.SwitchCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwitchCampaign`: PlatformCampaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.SwitchCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformCampaign**](PlatformCampaign.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> PlatformCampaign UpdateCampaign(ctx, token).Campaign(campaign).PrimaryImage(primaryImage).SecondaryImage(secondaryImage).Execute()

Update a campaign



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
    campaign := TODO // PlatformCampaignUpdate | 
    primaryImage := os.NewFile(1234, "some_file") // *os.File | 
    secondaryImage := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignApi.UpdateCampaign(context.Background(), token).Campaign(campaign).PrimaryImage(primaryImage).SecondaryImage(secondaryImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignApi.UpdateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaign`: PlatformCampaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignApi.UpdateCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaign** | [**PlatformCampaignUpdate**](PlatformCampaignUpdate.md) |  | 
 **primaryImage** | ***os.File** |  | 
 **secondaryImage** | ***os.File** |  | 

### Return type

[**PlatformCampaign**](PlatformCampaign.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

