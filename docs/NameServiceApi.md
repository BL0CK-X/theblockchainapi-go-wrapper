# \NameServiceApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockchainIdentifierFromName**](NameServiceApi.md#GetBlockchainIdentifierFromName) | **Post** /{blockchain}/{network}/name_service/name_to_blockchain_identifier | Get the identifier
[**GetNameForBlockchainIdentifier**](NameServiceApi.md#GetNameForBlockchainIdentifier) | **Post** /{blockchain}/{network}/name_service/blockchain_identifier_to_name | Get the name



## GetBlockchainIdentifierFromName

> InputBlockchainIdentifier GetBlockchainIdentifierFromName(ctx, blockchain, network).InputName(inputName).Execute()

Get the identifier



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
    inputName := *openapiclient.NewInputName() // InputName |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NameServiceApi.GetBlockchainIdentifierFromName(context.Background(), blockchain, network).InputName(inputName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NameServiceApi.GetBlockchainIdentifierFromName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockchainIdentifierFromName`: InputBlockchainIdentifier
    fmt.Fprintf(os.Stdout, "Response from `NameServiceApi.GetBlockchainIdentifierFromName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 
**network** | **string** | The network of the blockchain you selected  - Solana: &#x60;devnet&#x60;, &#x60;mainnet-beta&#x60; - Ethereum: &#x60;ropsten&#x60;, &#x60;mainnet&#x60;  Defaults when not provided (not applicable to path parameters): - Solana: &#x60;devnet&#x60; - Ethereum: &#x60;ropsten&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainIdentifierFromNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputName** | [**InputName**](InputName.md) |  | 

### Return type

[**InputBlockchainIdentifier**](InputBlockchainIdentifier.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNameForBlockchainIdentifier

> InputName GetNameForBlockchainIdentifier(ctx, blockchain, network, blockchainIdentifier).InputBlockchainIdentifier(inputBlockchainIdentifier).Execute()

Get the name



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
    blockchainIdentifier := "blockchainIdentifier_example" // string | The identifier of the token (e.g., `mint_address` on `Solana` or `token_address` on `Ethereum`) 
    inputBlockchainIdentifier := *openapiclient.NewInputBlockchainIdentifier() // InputBlockchainIdentifier |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NameServiceApi.GetNameForBlockchainIdentifier(context.Background(), blockchain, network, blockchainIdentifier).InputBlockchainIdentifier(inputBlockchainIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NameServiceApi.GetNameForBlockchainIdentifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNameForBlockchainIdentifier`: InputName
    fmt.Fprintf(os.Stdout, "Response from `NameServiceApi.GetNameForBlockchainIdentifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 
**network** | **string** | The network of the blockchain you selected  - Solana: &#x60;devnet&#x60;, &#x60;mainnet-beta&#x60; - Ethereum: &#x60;ropsten&#x60;, &#x60;mainnet&#x60;  Defaults when not provided (not applicable to path parameters): - Solana: &#x60;devnet&#x60; - Ethereum: &#x60;ropsten&#x60; | 
**blockchainIdentifier** | **string** | The identifier of the token (e.g., &#x60;mint_address&#x60; on &#x60;Solana&#x60; or &#x60;token_address&#x60; on &#x60;Ethereum&#x60;)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNameForBlockchainIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **inputBlockchainIdentifier** | [**InputBlockchainIdentifier**](InputBlockchainIdentifier.md) |  | 

### Return type

[**InputName**](InputName.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

