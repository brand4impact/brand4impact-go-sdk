# \TopicApi

All URIs are relative to *http://localhost:8080/backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTopic**](TopicApi.md#GetTopic) | **Get** /rest/platform/topic/{token} | Get a single topic
[**GetTopics**](TopicApi.md#GetTopics) | **Get** /rest/platform/topic | Get a set of topics



## GetTopic

> PlatformTopic GetTopic(ctx, token).Execute()

Get a single topic



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
    token := "token_example" // string | The unique token for the topic.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopicApi.GetTopic(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicApi.GetTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopic`: PlatformTopic
    fmt.Fprintf(os.Stdout, "Response from `TopicApi.GetTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique token for the topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformTopic**](PlatformTopic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopics

> PlatformTopicItemWrapper GetTopics(ctx).TitleLike(titleLike).Page(page).Size(size).Sort(sort).Execute()

Get a set of topics



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
    titleLike := "titleLike_example" // string | The name of the filter (optional)
    page := int32(56) // int32 | The number of the page. Possible values: 0 to unlimited (optional)
    size := int32(56) // int32 | The number of items inside a page. Possible values: 1 to 100 (optional)
    sort := "sort_example" // string | The order direction for sorted results. Possible values: ASC or DESC (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopicApi.GetTopics(context.Background()).TitleLike(titleLike).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicApi.GetTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopics`: PlatformTopicItemWrapper
    fmt.Fprintf(os.Stdout, "Response from `TopicApi.GetTopics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **titleLike** | **string** | The name of the filter | 
 **page** | **int32** | The number of the page. Possible values: 0 to unlimited | 
 **size** | **int32** | The number of items inside a page. Possible values: 1 to 100 | 
 **sort** | **string** | The order direction for sorted results. Possible values: ASC or DESC | 

### Return type

[**PlatformTopicItemWrapper**](PlatformTopicItemWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

