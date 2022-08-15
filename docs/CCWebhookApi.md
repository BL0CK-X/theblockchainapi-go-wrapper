# \CCWebhookApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCCWebhook**](CCWebhookApi.md#GetCCWebhook) | **Post** /webhook/v1/project/{project_id}/webhook/{webhook_identifier} | Get a webhook 
[**ListCCProjectWebhooks**](CCWebhookApi.md#ListCCProjectWebhooks) | **Post** /webhook/v1/project/{project_id}/webhooks | List project&#39;s webhooks 
[**ValidateCCWebhook**](CCWebhookApi.md#ValidateCCWebhook) | **Post** /webhook/v1/project/{project_id}/webhook/{webhook_identifier}/validate | Validate a webhook 



## GetCCWebhook

> CCWebhook GetCCWebhook(ctx, projectId, webhookIdentifier).Execute()

Get a webhook 



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
    webhookIdentifier := "webhook/payment.received/payment_EyfDL6CsvFinpUnXgBXaUuqq3hCXa0" // string | The ID of the webhook. Created and returned when a webhook is sent.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CCWebhookApi.GetCCWebhook(context.Background(), projectId, webhookIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCWebhookApi.GetCCWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCCWebhook`: CCWebhook
    fmt.Fprintf(os.Stdout, "Response from `CCWebhookApi.GetCCWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 
**webhookIdentifier** | **string** | The ID of the webhook. Created and returned when a webhook is sent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCCWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CCWebhook**](CCWebhook.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCCProjectWebhooks

> []CCWebhook ListCCProjectWebhooks(ctx, projectId).Execute()

List project's webhooks 



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
    resp, r, err := apiClient.CCWebhookApi.ListCCProjectWebhooks(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCWebhookApi.ListCCProjectWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCCProjectWebhooks`: []CCWebhook
    fmt.Fprintf(os.Stdout, "Response from `CCWebhookApi.ListCCProjectWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCCProjectWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CCWebhook**](CCWebhook.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCCWebhook

> ValidateCCWebhook(ctx, projectId, webhookIdentifier).CCWebhookValidateRequest(cCWebhookValidateRequest).Execute()

Validate a webhook 



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
    webhookIdentifier := "webhook/payment.received/payment_EyfDL6CsvFinpUnXgBXaUuqq3hCXa0" // string | The ID of the webhook. Created and returned when a webhook is sent.
    cCWebhookValidateRequest := *openapiclient.NewCCWebhookValidateRequest("webhook/payment.received/payment_2nFvUIGAsczBHSgiIz5JMenEiiJZet", map[string]interface{}([68,112]), float32(1651018365)) // CCWebhookValidateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CCWebhookApi.ValidateCCWebhook(context.Background(), projectId, webhookIdentifier).CCWebhookValidateRequest(cCWebhookValidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCWebhookApi.ValidateCCWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 
**webhookIdentifier** | **string** | The ID of the webhook. Created and returned when a webhook is sent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateCCWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cCWebhookValidateRequest** | [**CCWebhookValidateRequest**](CCWebhookValidateRequest.md) |  | 

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

