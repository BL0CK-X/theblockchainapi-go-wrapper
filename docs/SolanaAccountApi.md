# \SolanaAccountApi

All URIs are relative to *https://api.theblockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolanaGetAccount**](SolanaAccountApi.md#SolanaGetAccount) | **Get** /solana/account/{network}/{public_key} | Get the details of an account on Solana
[**SolanaGetAccountIsCandyMachine**](SolanaAccountApi.md#SolanaGetAccountIsCandyMachine) | **Get** /solana/account/{network}/{public_key}/is_candy_machine | Get if account is candy machine
[**SolanaGetAccountIsNFT**](SolanaAccountApi.md#SolanaGetAccountIsNFT) | **Get** /solana/account/{network}/{public_key}/is_nft | Get if account is NFT



## SolanaGetAccount

> Account SolanaGetAccount(ctx, network, publicKey).Execute()

Get the details of an account on Solana



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
    network := "mainnet-beta" // string | The network ID (devnet, mainnet-beta)
    publicKey := "EEr5yQpNXf7Bru6Rt5podx56HGW9CEehXqgRGh2wa71w" // string | The public key of the account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaAccountApi.SolanaGetAccount(context.Background(), network, publicKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaAccountApi.SolanaGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `SolanaAccountApi.SolanaGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**publicKey** | **string** | The public key of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Account**](Account.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetAccountIsCandyMachine

> AccountIsCandyMachine SolanaGetAccountIsCandyMachine(ctx, network, publicKey).Execute()

Get if account is candy machine



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
    network := "mainnet-beta" // string | The network ID (devnet, mainnet-beta)
    publicKey := "EEr5yQpNXf7Bru6Rt5podx56HGW9CEehXqgRGh2wa71w" // string | The public key of the account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaAccountApi.SolanaGetAccountIsCandyMachine(context.Background(), network, publicKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaAccountApi.SolanaGetAccountIsCandyMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetAccountIsCandyMachine`: AccountIsCandyMachine
    fmt.Fprintf(os.Stdout, "Response from `SolanaAccountApi.SolanaGetAccountIsCandyMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**publicKey** | **string** | The public key of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetAccountIsCandyMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountIsCandyMachine**](AccountIsCandyMachine.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetAccountIsNFT

> AccountIsNFT SolanaGetAccountIsNFT(ctx, network, publicKey).Execute()

Get if account is NFT



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
    network := "mainnet-beta" // string | The network ID (devnet, mainnet-beta)
    publicKey := "EEr5yQpNXf7Bru6Rt5podx56HGW9CEehXqgRGh2wa71w" // string | The public key of the account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaAccountApi.SolanaGetAccountIsNFT(context.Background(), network, publicKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaAccountApi.SolanaGetAccountIsNFT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetAccountIsNFT`: AccountIsNFT
    fmt.Fprintf(os.Stdout, "Response from `SolanaAccountApi.SolanaGetAccountIsNFT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**publicKey** | **string** | The public key of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetAccountIsNFTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountIsNFT**](AccountIsNFT.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

