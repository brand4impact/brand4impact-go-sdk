# \SecurityApi

All URIs are relative to *http://localhost:8080/backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJWTToken**](SecurityApi.md#GetJWTToken) | **Get** /security/jwt | get JWT token



## GetJWTToken

> SecurityUser GetJWTToken(ctx).Execute()

get JWT token

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
    resp, r, err := api_client.SecurityApi.GetJWTToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetJWTToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJWTToken`: SecurityUser
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetJWTToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJWTTokenRequest struct via the builder pattern


### Return type

[**SecurityUser**](SecurityUser.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

