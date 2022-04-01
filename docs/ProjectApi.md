# \ProjectApi

All URIs are relative to *http://localhost:8080/backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /rest/platform/project | Create a new project
[**GetProject**](ProjectApi.md#GetProject) | **Get** /rest/platform/project/{token} | Get a single project
[**GetProjects**](ProjectApi.md#GetProjects) | **Get** /rest/platform/project | Get a set of projects
[**SwitchProject**](ProjectApi.md#SwitchProject) | **Put** /rest/platform/project/{token}/switch | Switch project status
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Put** /rest/platform/project/{token} | Update a project



## CreateProject

> PlatformProject CreateProject(ctx).Project(project).PrimaryImage(primaryImage).SecondaryImage(secondaryImage).Execute()

Create a new project



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
    project := TODO // PlatformProjectCreate | 
    primaryImage := os.NewFile(1234, "some_file") // *os.File | 
    secondaryImage := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.CreateProject(context.Background()).Project(project).PrimaryImage(primaryImage).SecondaryImage(secondaryImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: PlatformProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**PlatformProjectCreate**](PlatformProjectCreate.md) |  | 
 **primaryImage** | ***os.File** |  | 
 **secondaryImage** | ***os.File** |  | 

### Return type

[**PlatformProject**](PlatformProject.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> PlatformProject GetProject(ctx, token).Execute()

Get a single project



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
    resp, r, err := api_client.ProjectApi.GetProject(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: PlatformProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique token for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformProject**](PlatformProject.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> PlatformProjectItemWrapper GetProjects(ctx).DescriptionLike(descriptionLike).Country(country).TopicToken(topicToken).Page(page).Size(size).Sort(sort).Execute()

Get a set of projects



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
    descriptionLike := "descriptionLike_example" // string |  (optional)
    country := "country_example" // string | The country (optional)
    topicToken := "topicToken_example" // string | The topicToken (optional)
    page := int32(56) // int32 | The number of the page. Possible values: 0 to unlimited (optional)
    size := int32(56) // int32 | The number of items inside a page. Possible values: 1 to 100 (optional)
    sort := "sort_example" // string | The order direction for sorted results. Possible values: ASC or DESC (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.GetProjects(context.Background()).DescriptionLike(descriptionLike).Country(country).TopicToken(topicToken).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjects`: PlatformProjectItemWrapper
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **descriptionLike** | **string** |  | 
 **country** | **string** | The country | 
 **topicToken** | **string** | The topicToken | 
 **page** | **int32** | The number of the page. Possible values: 0 to unlimited | 
 **size** | **int32** | The number of items inside a page. Possible values: 1 to 100 | 
 **sort** | **string** | The order direction for sorted results. Possible values: ASC or DESC | 

### Return type

[**PlatformProjectItemWrapper**](PlatformProjectItemWrapper.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchProject

> PlatformProject SwitchProject(ctx, token).Execute()

Switch project status



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
    resp, r, err := api_client.ProjectApi.SwitchProject(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.SwitchProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwitchProject`: PlatformProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.SwitchProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformProject**](PlatformProject.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> PlatformProject UpdateProject(ctx, token).Project(project).PrimaryImage(primaryImage).SecondaryImage(secondaryImage).Execute()

Update a project



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
    project := TODO // PlatformProjectUpdate | 
    primaryImage := os.NewFile(1234, "some_file") // *os.File | 
    secondaryImage := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.UpdateProject(context.Background(), token).Project(project).PrimaryImage(primaryImage).SecondaryImage(secondaryImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: PlatformProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**PlatformProjectUpdate**](PlatformProjectUpdate.md) |  | 
 **primaryImage** | ***os.File** |  | 
 **secondaryImage** | ***os.File** |  | 

### Return type

[**PlatformProject**](PlatformProject.md)

### Authorization

[BearerAuthentication](../README.md#BearerAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

