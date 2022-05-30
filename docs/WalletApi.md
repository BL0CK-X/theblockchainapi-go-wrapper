# \WalletApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DerivePrivateKey**](WalletApi.md#DerivePrivateKey) | **Post** /{blockchain}/wallet/private_key | Derive private key
[**DeriveWalletIdentifier**](WalletApi.md#DeriveWalletIdentifier) | **Post** /{blockchain}/wallet/identifier | Derive wallet identifier
[**GeneratePrivateKey**](WalletApi.md#GeneratePrivateKey) | **Post** /{blockchain}/wallet/generate/private_key | Generate private key
[**GenerateSeedPhrase**](WalletApi.md#GenerateSeedPhrase) | **Post** /{blockchain}/wallet/generate/secret_recovery_phrase | Generate seed phrase
[**GetAirdrop**](WalletApi.md#GetAirdrop) | **Post** /{blockchain}/wallet/airdrop | Get an airdrop
[**GetBalance**](WalletApi.md#GetBalance) | **Post** /{blockchain}/wallet/balance | Get wallet&#39;s balance of X
[**Transfer**](WalletApi.md#Transfer) | **Post** /{blockchain}/wallet/transfer | Transfer crypto, a token, or an NFT to another wallet



## DerivePrivateKey

> GeneralGeneratePrivateKeyResponse DerivePrivateKey(ctx, blockchain).SupplyWalletRequest(supplyWalletRequest).Execute()

Derive private key



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
    supplyWalletRequest := *openapiclient.NewSupplyWalletRequest(openapiclient.GeneralWallet{GeneralB58PrivateKey: openapiclient.NewGeneralB58PrivateKey("B58PrivateKey_example")}) // SupplyWalletRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletApi.DerivePrivateKey(context.Background(), blockchain).SupplyWalletRequest(supplyWalletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.DerivePrivateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DerivePrivateKey`: GeneralGeneratePrivateKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.DerivePrivateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDerivePrivateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supplyWalletRequest** | [**SupplyWalletRequest**](SupplyWalletRequest.md) |  | 

### Return type

[**GeneralGeneratePrivateKeyResponse**](GeneralGeneratePrivateKeyResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeriveWalletIdentifier

> WalletIdentifiers DeriveWalletIdentifier(ctx, blockchain).SupplyWalletRequest(supplyWalletRequest).Execute()

Derive wallet identifier



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
    supplyWalletRequest := *openapiclient.NewSupplyWalletRequest(openapiclient.GeneralWallet{GeneralB58PrivateKey: openapiclient.NewGeneralB58PrivateKey("B58PrivateKey_example")}) // SupplyWalletRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletApi.DeriveWalletIdentifier(context.Background(), blockchain).SupplyWalletRequest(supplyWalletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.DeriveWalletIdentifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeriveWalletIdentifier`: WalletIdentifiers
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.DeriveWalletIdentifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeriveWalletIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supplyWalletRequest** | [**SupplyWalletRequest**](SupplyWalletRequest.md) |  | 

### Return type

[**WalletIdentifiers**](WalletIdentifiers.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneratePrivateKey

> GeneralGeneratePrivateKeyResponse GeneratePrivateKey(ctx, blockchain).Execute()

Generate private key



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletApi.GeneratePrivateKey(context.Background(), blockchain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GeneratePrivateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneratePrivateKey`: GeneralGeneratePrivateKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GeneratePrivateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePrivateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GeneralGeneratePrivateKeyResponse**](GeneralGeneratePrivateKeyResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSeedPhrase

> GeneralSecretPhrase GenerateSeedPhrase(ctx, blockchain).GenerateSeedPhraseRequest(generateSeedPhraseRequest).Execute()

Generate seed phrase



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
    generateSeedPhraseRequest := *openapiclient.NewGenerateSeedPhraseRequest() // GenerateSeedPhraseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletApi.GenerateSeedPhrase(context.Background(), blockchain).GenerateSeedPhraseRequest(generateSeedPhraseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GenerateSeedPhrase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSeedPhrase`: GeneralSecretPhrase
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GenerateSeedPhrase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSeedPhraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateSeedPhraseRequest** | [**GenerateSeedPhraseRequest**](GenerateSeedPhraseRequest.md) |  | 

### Return type

[**GeneralSecretPhrase**](GeneralSecretPhrase.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAirdrop

> GeneralTransactionMadeResponse GetAirdrop(ctx, blockchain).GeneralAirdropRequest(generalAirdropRequest).Execute()

Get an airdrop



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
    generalAirdropRequest := *openapiclient.NewGeneralAirdropRequest("RecipientBlockchainIdentifier_example") // GeneralAirdropRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletApi.GetAirdrop(context.Background(), blockchain).GeneralAirdropRequest(generalAirdropRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetAirdrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAirdrop`: GeneralTransactionMadeResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetAirdrop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAirdropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generalAirdropRequest** | [**GeneralAirdropRequest**](GeneralAirdropRequest.md) |  | 

### Return type

[**GeneralTransactionMadeResponse**](GeneralTransactionMadeResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalance

> GeneralBalanceResponse GetBalance(ctx, blockchain).GeneralBalanceRequest(generalBalanceRequest).Execute()

Get wallet's balance of X



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
    generalBalanceRequest := *openapiclient.NewGeneralBalanceRequest() // GeneralBalanceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletApi.GetBalance(context.Background(), blockchain).GeneralBalanceRequest(generalBalanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalance`: GeneralBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generalBalanceRequest** | [**GeneralBalanceRequest**](GeneralBalanceRequest.md) |  | 

### Return type

[**GeneralBalanceResponse**](GeneralBalanceResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Transfer

> GeneralTransactionResponse Transfer(ctx, blockchain).GeneralTransferRequest(generalTransferRequest).Execute()

Transfer crypto, a token, or an NFT to another wallet



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
    generalTransferRequest := *openapiclient.NewGeneralTransferRequest("RecipientBlockchainIdentifier_example") // GeneralTransferRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletApi.Transfer(context.Background(), blockchain).GeneralTransferRequest(generalTransferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.Transfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Transfer`: GeneralTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.Transfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generalTransferRequest** | [**GeneralTransferRequest**](GeneralTransferRequest.md) |  | 

### Return type

[**GeneralTransactionResponse**](GeneralTransactionResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

