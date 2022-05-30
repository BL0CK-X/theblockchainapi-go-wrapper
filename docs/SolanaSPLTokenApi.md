# \SolanaSPLTokenApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolanaGetSPLToken**](SolanaSPLTokenApi.md#SolanaGetSPLToken) | **Get** /solana/spl-token/{network}/{public_key} | Get SPL token metadata



## SolanaGetSPLToken

> GetSPLTokenResponse SolanaGetSPLToken(ctx, publicKey, network).Execute()

Get SPL token metadata



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
    publicKey := "MangoCzJ36AjZyKwVj3VnYU4GTonjfVEnJmvvWaxLac" // string | The public key of the token
    network := "mainnet-beta" // string | The network ID (devnet, mainnet-beta)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaSPLTokenApi.SolanaGetSPLToken(context.Background(), publicKey, network).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaSPLTokenApi.SolanaGetSPLToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetSPLToken`: GetSPLTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaSPLTokenApi.SolanaGetSPLToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicKey** | **string** | The public key of the token | 
**network** | **string** | The network ID (devnet, mainnet-beta) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetSPLTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSPLTokenResponse**](GetSPLTokenResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

