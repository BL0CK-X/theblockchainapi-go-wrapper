# \CCProjectApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCCProject**](CCProjectApi.md#CreateCCProject) | **Post** /checkout/v1/project | Create a project 
[**DeleteCCProject**](CCProjectApi.md#DeleteCCProject) | **Delete** /checkout/v1/project/{project_id} | Delete a project 
[**GetCCProject**](CCProjectApi.md#GetCCProject) | **Get** /checkout/v1/project/{project_id} | Get a project 
[**ListCCProjects**](CCProjectApi.md#ListCCProjects) | **Get** /checkout/v1/projects | List projects 
[**UpdateCCProject**](CCProjectApi.md#UpdateCCProject) | **Put** /checkout/v1/project/{project_id} | Update a project 



## CreateCCProject

> CCProject CreateCCProject(ctx).CCProjectCreateRequest(cCProjectCreateRequest).Execute()

Create a project 



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
    cCProjectCreateRequest := *openapiclient.NewCCProjectCreateRequest("Name_example") // CCProjectCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CCProjectApi.CreateCCProject(context.Background()).CCProjectCreateRequest(cCProjectCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCProjectApi.CreateCCProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCCProject`: CCProject
    fmt.Fprintf(os.Stdout, "Response from `CCProjectApi.CreateCCProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCCProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cCProjectCreateRequest** | [**CCProjectCreateRequest**](CCProjectCreateRequest.md) |  | 

### Return type

[**CCProject**](CCProject.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCCProject

> DeleteCCProject(ctx, projectId).Execute()

Delete a project 



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
    projectId := "project_WDQskRIfHQxj53N5mk5K" // string | The ID of the project. Created and returned when a project is created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CCProjectApi.DeleteCCProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCProjectApi.DeleteCCProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCCProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCCProject

> CCProject GetCCProject(ctx, projectId).Execute()

Get a project 



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
    projectId := "project_WDQskRIfHQxj53N5mk5K" // string | The ID of the project. Created and returned when a project is created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CCProjectApi.GetCCProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCProjectApi.GetCCProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCCProject`: CCProject
    fmt.Fprintf(os.Stdout, "Response from `CCProjectApi.GetCCProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCCProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CCProject**](CCProject.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCCProjects

> []CCProject ListCCProjects(ctx).Execute()

List projects 



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CCProjectApi.ListCCProjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCProjectApi.ListCCProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCCProjects`: []CCProject
    fmt.Fprintf(os.Stdout, "Response from `CCProjectApi.ListCCProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCCProjectsRequest struct via the builder pattern


### Return type

[**[]CCProject**](CCProject.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCCProject

> CCProject UpdateCCProject(ctx, projectId).CCProject(cCProject).Execute()

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
    projectId := "project_WDQskRIfHQxj53N5mk5K" // string | The ID of the project. Created and returned when a project is created.
    cCProject := *openapiclient.NewCCProject("ProjectId_example", "Name_example") // CCProject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CCProjectApi.UpdateCCProject(context.Background(), projectId).CCProject(cCProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCProjectApi.UpdateCCProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCCProject`: CCProject
    fmt.Fprintf(os.Stdout, "Response from `CCProjectApi.UpdateCCProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCCProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cCProject** | [**CCProject**](CCProject.md) |  | 

### Return type

[**CCProject**](CCProject.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

