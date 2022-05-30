# \TokenApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokenMetadata**](TokenApi.md#GetTokenMetadata) | **Get** /{blockchain}/{network}/token/{token_blockchain_identifier} | Get a token&#39;s metadata
[**ListAllTokens**](TokenApi.md#ListAllTokens) | **Get** /{blockchain}/{network}/all_tokens | List all tokens



## GetTokenMetadata

> TokenMetadataResponse GetTokenMetadata(ctx, blockchain, network, tokenBlockchainIdentifier).Execute()

Get a token's metadata



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
    blockchain := "blockchain_example" // string | The blockchain you want to use 
    network := "ropsten" // string | The network of the blockchain you selected  - Solana: `devnet`, `mainnet-beta` - Ethereum: `ropsten`, `mainnet`  Defaults when not provided (not applicable to path parameters): - Solana: `devnet` - Ethereum: `ropsten`
    tokenBlockchainIdentifier := "tokenBlockchainIdentifier_example" // string | The identifier of the token (e.g., `mint_address` on `Solana` or `token_address` on `Ethereum`) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.GetTokenMetadata(context.Background(), blockchain, network, tokenBlockchainIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.GetTokenMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenMetadata`: TokenMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.GetTokenMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 
**network** | **string** | The network of the blockchain you selected  - Solana: &#x60;devnet&#x60;, &#x60;mainnet-beta&#x60; - Ethereum: &#x60;ropsten&#x60;, &#x60;mainnet&#x60;  Defaults when not provided (not applicable to path parameters): - Solana: &#x60;devnet&#x60; - Ethereum: &#x60;ropsten&#x60; | 
**tokenBlockchainIdentifier** | **string** | The identifier of the token (e.g., &#x60;mint_address&#x60; on &#x60;Solana&#x60; or &#x60;token_address&#x60; on &#x60;Ethereum&#x60;)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TokenMetadataResponse**](TokenMetadataResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllTokens

> []TokenMetadataResponse ListAllTokens(ctx, blockchain, network).Execute()

List all tokens



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
    blockchain := "blockchain_example" // string | The blockchain you want to use 
    network := "ropsten" // string | The network of the blockchain you selected  - Solana: `devnet`, `mainnet-beta` - Ethereum: `ropsten`, `mainnet`  Defaults when not provided (not applicable to path parameters): - Solana: `devnet` - Ethereum: `ropsten`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.ListAllTokens(context.Background(), blockchain, network).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.ListAllTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllTokens`: []TokenMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.ListAllTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 
**network** | **string** | The network of the blockchain you selected  - Solana: &#x60;devnet&#x60;, &#x60;mainnet-beta&#x60; - Ethereum: &#x60;ropsten&#x60;, &#x60;mainnet&#x60;  Defaults when not provided (not applicable to path parameters): - Solana: &#x60;devnet&#x60; - Ethereum: &#x60;ropsten&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TokenMetadataResponse**](TokenMetadataResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

