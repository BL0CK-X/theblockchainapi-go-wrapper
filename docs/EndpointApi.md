# \EndpointApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEndpoint**](EndpointApi.md#DeleteEndpoint) | **Post** /endpoint/delete | Delete an endpoint 
[**GetEndpoint**](EndpointApi.md#GetEndpoint) | **Post** /endpoint/metadata | Get an endpoint&#39;s metadata 
[**ListEndpoints**](EndpointApi.md#ListEndpoints) | **Get** /endpoint/list | List all endpoints 
[**SetEndpoint**](EndpointApi.md#SetEndpoint) | **Post** /endpoint | Create / update an endpoint 



## DeleteEndpoint

> DeleteEndpoint(ctx).EndpointReference(endpointReference).Execute()

Delete an endpoint 



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
    endpointReference := *openapiclient.NewEndpointReference("ProjectId_example", "Version_example", "Path_example") // EndpointReference |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndpointApi.DeleteEndpoint(context.Background()).EndpointReference(endpointReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndpointApi.DeleteEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointReference** | [**EndpointReference**](EndpointReference.md) |  | 

### Return type

 (empty response body)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpoint

> Endpoint GetEndpoint(ctx).EndpointReference(endpointReference).Execute()

Get an endpoint's metadata 



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
    endpointReference := *openapiclient.NewEndpointReference("ProjectId_example", "Version_example", "Path_example") // EndpointReference |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndpointApi.GetEndpoint(context.Background()).EndpointReference(endpointReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndpointApi.GetEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEndpoint`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `EndpointApi.GetEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointReference** | [**EndpointReference**](EndpointReference.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpoints

> []Endpoint ListEndpoints(ctx).Execute()

List all endpoints 



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
    resp, r, err := apiClient.EndpointApi.ListEndpoints(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndpointApi.ListEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpoints`: []Endpoint
    fmt.Fprintf(os.Stdout, "Response from `EndpointApi.ListEndpoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsRequest struct via the builder pattern


### Return type

[**[]Endpoint**](Endpoint.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetEndpoint

> Endpoint SetEndpoint(ctx).Endpoint(endpoint).Execute()

Create / update an endpoint 



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
    endpoint := *openapiclient.NewEndpoint("ProjectId_example", "Version_example", "Path_example", "OperationId_example", "ReadableName_example", float32(123), []openapiclient.ParameterSpecification{*openapiclient.NewParameterSpecification()}, map[string]interface{}(123), []openapiclient.ParameterSpecification{*openapiclient.NewParameterSpecification()}, map[string]interface{}(123)) // Endpoint |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndpointApi.SetEndpoint(context.Background()).Endpoint(endpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndpointApi.SetEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetEndpoint`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `EndpointApi.SetEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpoint** | [**Endpoint**](Endpoint.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

