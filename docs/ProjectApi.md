# \ProjectApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /project | Create a project 
[**CreateProjectVersion**](ProjectApi.md#CreateProjectVersion) | **Post** /project/{project_id}/{version} | Create a new project version 
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /project/{project_id} | Delete a project 
[**DeleteProjectVersion**](ProjectApi.md#DeleteProjectVersion) | **Delete** /project/{project_id}/{version} | Delete a project version 
[**GetProject**](ProjectApi.md#GetProject) | **Get** /project/{project_id} | Get a project&#39;s metadata 
[**GetProjectDeploymentStatus**](ProjectApi.md#GetProjectDeploymentStatus) | **Post** /project/{project_id}/deploy/status | Get deployment status 
[**GetProjectDeploymentURL**](ProjectApi.md#GetProjectDeploymentURL) | **Post** /project/{project_id}/deploy/url | Get the deployment URL 
[**GetProjectStats**](ProjectApi.md#GetProjectStats) | **Get** /project/{project_id}/stats | Get a project&#39;s stats 
[**ListProjects**](ProjectApi.md#ListProjects) | **Get** /project/list | List projects 
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Post** /project/{project_id} | Update a project 
[**UpdateProjectDocumentation**](ProjectApi.md#UpdateProjectDocumentation) | **Post** /project/{project_id}/{version}/documentation | Update the project&#39;s documentation 



## CreateProject

> Project CreateProject(ctx).ProjectCreateRequest(projectCreateRequest).Execute()

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
    projectCreateRequest := *openapiclient.NewProjectCreateRequest() // ProjectCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.CreateProject(context.Background()).ProjectCreateRequest(projectCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectCreateRequest** | [**ProjectCreateRequest**](ProjectCreateRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectVersion

> Project CreateProjectVersion(ctx, projectId, version).Execute()

Create a new project version 



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
    projectId := "L3cPm0cJlf5MeDXzGLQJJ4vlzndJCC" // string | The ID of the project. Created and returned when a project is created.
    version := "2" // string | The version of the project.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.CreateProjectVersion(context.Background(), projectId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.CreateProjectVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectVersion`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.CreateProjectVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 
**version** | **string** | The version of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Project**](Project.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProject(ctx, projectId).Execute()

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
    projectId := "L3cPm0cJlf5MeDXzGLQJJ4vlzndJCC" // string | The ID of the project. Created and returned when a project is created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.DeleteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeleteProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


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


## DeleteProjectVersion

> Project DeleteProjectVersion(ctx, projectId, version).Execute()

Delete a project version 



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
    projectId := "L3cPm0cJlf5MeDXzGLQJJ4vlzndJCC" // string | The ID of the project. Created and returned when a project is created.
    version := "2" // string | The version of the project.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.DeleteProjectVersion(context.Background(), projectId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeleteProjectVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProjectVersion`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeleteProjectVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 
**version** | **string** | The version of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Project**](Project.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> Project GetProject(ctx, projectId).Execute()

Get a project's metadata 



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
    projectId := "L3cPm0cJlf5MeDXzGLQJJ4vlzndJCC" // string | The ID of the project. Created and returned when a project is created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.GetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDeploymentStatus

> interface{} GetProjectDeploymentStatus(ctx, projectId).Execute()

Get deployment status 



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
    projectId := "L3cPm0cJlf5MeDXzGLQJJ4vlzndJCC" // string | The ID of the project. Created and returned when a project is created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.GetProjectDeploymentStatus(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProjectDeploymentStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectDeploymentStatus`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProjectDeploymentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDeploymentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDeploymentURL

> ProjectDeploymentURL GetProjectDeploymentURL(ctx, projectId).InlineObject(inlineObject).Execute()

Get the deployment URL 



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
    projectId := "L3cPm0cJlf5MeDXzGLQJJ4vlzndJCC" // string | The ID of the project. Created and returned when a project is created.
    inlineObject := *openapiclient.NewInlineObject() // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.GetProjectDeploymentURL(context.Background(), projectId).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProjectDeploymentURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectDeploymentURL`: ProjectDeploymentURL
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProjectDeploymentURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDeploymentURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**ProjectDeploymentURL**](ProjectDeploymentURL.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectStats

> []StatItem GetProjectStats(ctx, projectId).Execute()

Get a project's stats 



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
    projectId := "L3cPm0cJlf5MeDXzGLQJJ4vlzndJCC" // string | The ID of the project. Created and returned when a project is created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.GetProjectStats(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProjectStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectStats`: []StatItem
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProjectStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StatItem**](StatItem.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjects

> []Project ListProjects(ctx).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.ListProjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ListProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjects`: []Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ListProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectsRequest struct via the builder pattern


### Return type

[**[]Project**](Project.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> Project UpdateProject(ctx, projectId).ProjectCreateRequest(projectCreateRequest).Execute()

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
    projectId := "L3cPm0cJlf5MeDXzGLQJJ4vlzndJCC" // string | The ID of the project. Created and returned when a project is created.
    projectCreateRequest := *openapiclient.NewProjectCreateRequest() // ProjectCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.UpdateProject(context.Background(), projectId).ProjectCreateRequest(projectCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectCreateRequest** | [**ProjectCreateRequest**](ProjectCreateRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectDocumentation

> Project UpdateProjectDocumentation(ctx, projectId, version).Execute()

Update the project's documentation 



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
    projectId := "L3cPm0cJlf5MeDXzGLQJJ4vlzndJCC" // string | The ID of the project. Created and returned when a project is created.
    version := "2" // string | The version of the project.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.UpdateProjectDocumentation(context.Background(), projectId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.UpdateProjectDocumentation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProjectDocumentation`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.UpdateProjectDocumentation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 
**version** | **string** | The version of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectDocumentationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Project**](Project.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

