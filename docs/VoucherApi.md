# \VoucherApi

All URIs are relative to *http://localhost:8080/backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadVoucherQr**](VoucherApi.md#DownloadVoucherQr) | **Get** /rest/platform/voucher/{token}/qr | Download the voucher QR image
[**DownloadVouchersFile**](VoucherApi.md#DownloadVouchersFile) | **Get** /rest/platform/voucher/campaign/{campaignToken}/file | Download a set of voucher&#39;s links file
[**DownloadVouchersQr**](VoucherApi.md#DownloadVouchersQr) | **Get** /rest/platform/voucher/campaign/{campaignToken}/qr | Download a set of voucher QR images
[**GetVoucher**](VoucherApi.md#GetVoucher) | **Get** /rest/platform/voucher/{token} | Get a single voucher
[**GetVouchers**](VoucherApi.md#GetVouchers) | **Get** /rest/platform/voucher/campaign/{campaignToken} | Get a set of vouchers



## DownloadVoucherQr

> string DownloadVoucherQr(ctx, token).Execute()

Download the voucher QR image



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
    resp, r, err := api_client.VoucherApi.DownloadVoucherQr(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.DownloadVoucherQr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadVoucherQr`: string
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.DownloadVoucherQr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadVoucherQrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadVouchersFile

> string DownloadVouchersFile(ctx, campaignToken).Status(status).Min(min).Max(max).Execute()

Download a set of voucher's links file



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
    campaignToken := "campaignToken_example" // string | The unique token for the campaign.
    status := "status_example" // string | The status of the voucher. Possible values: Active or Inactive (optional)
    min := int32(56) // int32 | Min value of the voucher (optional)
    max := int32(56) // int32 | Max value of the voucher (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VoucherApi.DownloadVouchersFile(context.Background(), campaignToken).Status(status).Min(min).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.DownloadVouchersFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadVouchersFile`: string
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.DownloadVouchersFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignToken** | **string** | The unique token for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadVouchersFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | The status of the voucher. Possible values: Active or Inactive | 
 **min** | **int32** | Min value of the voucher | 
 **max** | **int32** | Max value of the voucher | 

### Return type

**string**

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadVouchersQr

> string DownloadVouchersQr(ctx, campaignToken).Status(status).Min(min).Max(max).Execute()

Download a set of voucher QR images



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
    campaignToken := "campaignToken_example" // string | The unique token for the campaign.
    status := "status_example" // string | The status of the voucher. Possible values: Active or Inactive (optional)
    min := int32(56) // int32 | Min value of the voucher (optional)
    max := int32(56) // int32 | Max value of the voucher (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VoucherApi.DownloadVouchersQr(context.Background(), campaignToken).Status(status).Min(min).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.DownloadVouchersQr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadVouchersQr`: string
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.DownloadVouchersQr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignToken** | **string** | The unique token for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadVouchersQrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | The status of the voucher. Possible values: Active or Inactive | 
 **min** | **int32** | Min value of the voucher | 
 **max** | **int32** | Max value of the voucher | 

### Return type

**string**

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoucher

> PlatformVoucher GetVoucher(ctx, token).Execute()

Get a single voucher



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
    resp, r, err := api_client.VoucherApi.GetVoucher(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.GetVoucher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVoucher`: PlatformVoucher
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.GetVoucher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformVoucher**](PlatformVoucher.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVouchers

> PlatformVoucherItemWrapper GetVouchers(ctx, campaignToken).Status(status).Min(min).Max(max).Page(page).Size(size).Execute()

Get a set of vouchers



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
    campaignToken := "campaignToken_example" // string | The unique token for the campaign.
    status := "status_example" // string | The status of the voucher. Possible values: Active or Inactive (optional)
    min := int32(56) // int32 | Min value of the voucher (optional)
    max := int32(56) // int32 | Max value of the voucher (optional)
    page := int32(56) // int32 | The number of the page. Possible values: 0 to unlimited (optional)
    size := int32(56) // int32 | The number of items inside a page. Possible values: 1 to 100 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VoucherApi.GetVouchers(context.Background(), campaignToken).Status(status).Min(min).Max(max).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.GetVouchers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVouchers`: PlatformVoucherItemWrapper
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.GetVouchers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignToken** | **string** | The unique token for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVouchersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | The status of the voucher. Possible values: Active or Inactive | 
 **min** | **int32** | Min value of the voucher | 
 **max** | **int32** | Max value of the voucher | 
 **page** | **int32** | The number of the page. Possible values: 0 to unlimited | 
 **size** | **int32** | The number of items inside a page. Possible values: 1 to 100 | 

### Return type

[**PlatformVoucherItemWrapper**](PlatformVoucherItemWrapper.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

