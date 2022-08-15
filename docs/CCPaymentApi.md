# \CCPaymentApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCCPayment**](CCPaymentApi.md#GetCCPayment) | **Get** /checkout/v1/project/{project_id}/payment/{payment_identifier} | Get a payment 
[**ListAllCCPayments**](CCPaymentApi.md#ListAllCCPayments) | **Get** /checkout/v1/payments | List all payments  
[**ListCCProjectPayments**](CCPaymentApi.md#ListCCProjectPayments) | **Get** /checkout/v1/project/{project_id}/payments | List a project&#39;s payments 



## GetCCPayment

> CCPayment GetCCPayment(ctx, projectId, paymentIdentifier).Execute()

Get a payment 



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
    paymentIdentifier := "["payment_97jA1szpVjrSI5EvOb5zxNQFq5CfrC"]" // string | You can supply either `payment_id` or `payment_validation_code`.  The `payment_id` is only visible by you and uniquely identifies a payment.  The `payment_validation_code` is shown exclusively to the customer and the owner of the project. It can be used to redeem a payment. You can verify a payment by pulling the payment with the `payment_validation_code`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CCPaymentApi.GetCCPayment(context.Background(), projectId, paymentIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCPaymentApi.GetCCPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCCPayment`: CCPayment
    fmt.Fprintf(os.Stdout, "Response from `CCPaymentApi.GetCCPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 
**paymentIdentifier** | **string** | You can supply either &#x60;payment_id&#x60; or &#x60;payment_validation_code&#x60;.  The &#x60;payment_id&#x60; is only visible by you and uniquely identifies a payment.  The &#x60;payment_validation_code&#x60; is shown exclusively to the customer and the owner of the project. It can be used to redeem a payment. You can verify a payment by pulling the payment with the &#x60;payment_validation_code&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCCPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CCPayment**](CCPayment.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllCCPayments

> []CCPayment ListAllCCPayments(ctx).Execute()

List all payments  



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
    resp, r, err := apiClient.CCPaymentApi.ListAllCCPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCPaymentApi.ListAllCCPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllCCPayments`: []CCPayment
    fmt.Fprintf(os.Stdout, "Response from `CCPaymentApi.ListAllCCPayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllCCPaymentsRequest struct via the builder pattern


### Return type

[**[]CCPayment**](CCPayment.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCCProjectPayments

> []CCPayment ListCCProjectPayments(ctx, projectId).UNKNOWN_PARAMETER_NAME(UNKNOWN_PARAMETER_NAME).Execute()

List a project's payments 



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
    UNKNOWN_PARAMETER_NAME := TODO //  | Filter payments by a user's wallet identifier (i.e., Solana public key)  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CCPaymentApi.ListCCProjectPayments(context.Background(), projectId).UNKNOWN_PARAMETER_NAME(UNKNOWN_PARAMETER_NAME).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CCPaymentApi.ListCCProjectPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCCProjectPayments`: []CCPayment
    fmt.Fprintf(os.Stdout, "Response from `CCPaymentApi.ListCCProjectPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. Created and returned when a project is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCCProjectPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **UNKNOWN_PARAMETER_NAME** | [****](.md) | Filter payments by a user&#39;s wallet identifier (i.e., Solana public key)  | 

### Return type

[**[]CCPayment**](CCPayment.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

