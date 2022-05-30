# \SolanaWalletApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolanaDeriveAssociatedTokenAccountAddress**](SolanaWalletApi.md#SolanaDeriveAssociatedTokenAccountAddress) | **Get** /solana/wallet/{public_key}/associated_token_account/{mint_address} | Derive an associated token account address
[**SolanaDerivePrivateKey**](SolanaWalletApi.md#SolanaDerivePrivateKey) | **Post** /solana/wallet/private_key | Derive private key
[**SolanaDerivePublicKey**](SolanaWalletApi.md#SolanaDerivePublicKey) | **Post** /solana/wallet/public_key | Derive public key
[**SolanaGeneratePrivateKey**](SolanaWalletApi.md#SolanaGeneratePrivateKey) | **Post** /solana/wallet/generate/private_key | Generate private key
[**SolanaGenerateSecretRecoveryPhrase**](SolanaWalletApi.md#SolanaGenerateSecretRecoveryPhrase) | **Post** /solana/wallet/generate/secret_recovery_phrase | Generate secret phrase
[**SolanaGetAirdrop**](SolanaWalletApi.md#SolanaGetAirdrop) | **Post** /solana/wallet/airdrop | Get an airdrop on devnet
[**SolanaGetBalance**](SolanaWalletApi.md#SolanaGetBalance) | **Post** /solana/wallet/balance | Get wallet&#39;s balance in SOL or any SPL
[**SolanaGetNFTsBelongingToWallet**](SolanaWalletApi.md#SolanaGetNFTsBelongingToWallet) | **Get** /solana/wallet/{network}/{public_key}/nfts | Get address&#39;s NFTs
[**SolanaGetTokensBelongingToWallet**](SolanaWalletApi.md#SolanaGetTokensBelongingToWallet) | **Get** /solana/wallet/{network}/{public_key}/tokens | Get address&#39;s tokens and respective balances
[**SolanaGetWalletTransactions**](SolanaWalletApi.md#SolanaGetWalletTransactions) | **Get** /solana/wallet/{network}/{public_key}/transactions | Get address&#39;s associated transaction signatures
[**SolanaTransfer**](SolanaWalletApi.md#SolanaTransfer) | **Post** /solana/wallet/transfer | Transfer SOL, a token, or an NFT to another address



## SolanaDeriveAssociatedTokenAccountAddress

> ATAResponse SolanaDeriveAssociatedTokenAccountAddress(ctx, publicKey, mintAddress).Execute()

Derive an associated token account address



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
    publicKey := "31LKs39pjT5oj6fWjC3F76dHWf9489asCthmgj8wu7pj" // string | The public key of the wallet
    mintAddress := "7EWJNaNYfPYMaqzdAAa4ps5kpqW95B7VHsjhW1kr18sj" // string | The mint address of the token (either SPL or NFT)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaDeriveAssociatedTokenAccountAddress(context.Background(), publicKey, mintAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaDeriveAssociatedTokenAccountAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaDeriveAssociatedTokenAccountAddress`: ATAResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaDeriveAssociatedTokenAccountAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicKey** | **string** | The public key of the wallet | 
**mintAddress** | **string** | The mint address of the token (either SPL or NFT) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaDeriveAssociatedTokenAccountAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ATAResponse**](ATAResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaDerivePrivateKey

> GeneratePrivateKey SolanaDerivePrivateKey(ctx).GetPublicKeyRequest(getPublicKeyRequest).Execute()

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
    getPublicKeyRequest := *openapiclient.NewGetPublicKeyRequest(openapiclient.Wallet{B58PrivateKey: openapiclient.NewB58PrivateKey("B58PrivateKey_example")}) // GetPublicKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaDerivePrivateKey(context.Background()).GetPublicKeyRequest(getPublicKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaDerivePrivateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaDerivePrivateKey`: GeneratePrivateKey
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaDerivePrivateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaDerivePrivateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPublicKeyRequest** | [**GetPublicKeyRequest**](GetPublicKeyRequest.md) |  | 

### Return type

[**GeneratePrivateKey**](GeneratePrivateKey.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaDerivePublicKey

> PublicKey SolanaDerivePublicKey(ctx).GetPublicKeyRequest(getPublicKeyRequest).Execute()

Derive public key



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
    getPublicKeyRequest := *openapiclient.NewGetPublicKeyRequest(openapiclient.Wallet{B58PrivateKey: openapiclient.NewB58PrivateKey("B58PrivateKey_example")}) // GetPublicKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaDerivePublicKey(context.Background()).GetPublicKeyRequest(getPublicKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaDerivePublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaDerivePublicKey`: PublicKey
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaDerivePublicKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaDerivePublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPublicKeyRequest** | [**GetPublicKeyRequest**](GetPublicKeyRequest.md) |  | 

### Return type

[**PublicKey**](PublicKey.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGeneratePrivateKey

> GeneratePrivateKey SolanaGeneratePrivateKey(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaGeneratePrivateKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaGeneratePrivateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGeneratePrivateKey`: GeneratePrivateKey
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaGeneratePrivateKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGeneratePrivateKeyRequest struct via the builder pattern


### Return type

[**GeneratePrivateKey**](GeneratePrivateKey.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGenerateSecretRecoveryPhrase

> SecretPhrase SolanaGenerateSecretRecoveryPhrase(ctx).Execute()

Generate secret phrase



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
    resp, r, err := apiClient.SolanaWalletApi.SolanaGenerateSecretRecoveryPhrase(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaGenerateSecretRecoveryPhrase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGenerateSecretRecoveryPhrase`: SecretPhrase
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaGenerateSecretRecoveryPhrase`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGenerateSecretRecoveryPhraseRequest struct via the builder pattern


### Return type

[**SecretPhrase**](SecretPhrase.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetAirdrop

> TransferResponse SolanaGetAirdrop(ctx).AirdropRequest(airdropRequest).Execute()

Get an airdrop on devnet



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
    airdropRequest := *openapiclient.NewAirdropRequest("RecipientAddress_example") // AirdropRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaGetAirdrop(context.Background()).AirdropRequest(airdropRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaGetAirdrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetAirdrop`: TransferResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaGetAirdrop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetAirdropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **airdropRequest** | [**AirdropRequest**](AirdropRequest.md) |  | 

### Return type

[**TransferResponse**](TransferResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetBalance

> BalanceResponse SolanaGetBalance(ctx).BalanceRequest(balanceRequest).Execute()

Get wallet's balance in SOL or any SPL



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
    balanceRequest := *openapiclient.NewBalanceRequest("PublicKey_example") // BalanceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaGetBalance(context.Background()).BalanceRequest(balanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaGetBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetBalance`: BalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaGetBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **balanceRequest** | [**BalanceRequest**](BalanceRequest.md) |  | 

### Return type

[**BalanceResponse**](BalanceResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetNFTsBelongingToWallet

> ListNFTsResponse SolanaGetNFTsBelongingToWallet(ctx, network, publicKey).Execute()

Get address's NFTs



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
    publicKey := "HE3ZYk4aWfKD7R9EmFQbxjj75JdgHuDztNAsseKVan82" // string | The public key of the account whose list of owned NFTs you want to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaGetNFTsBelongingToWallet(context.Background(), network, publicKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaGetNFTsBelongingToWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetNFTsBelongingToWallet`: ListNFTsResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaGetNFTsBelongingToWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**publicKey** | **string** | The public key of the account whose list of owned NFTs you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetNFTsBelongingToWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListNFTsResponse**](ListNFTsResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetTokensBelongingToWallet

> []map[string]interface{} SolanaGetTokensBelongingToWallet(ctx, network, publicKey).IncludeNfts(includeNfts).IncludeZeroBalanceHoldings(includeZeroBalanceHoldings).Execute()

Get address's tokens and respective balances



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
    publicKey := "GKNcUmNacSJo4S2Kq3DuYRYRGw3sNUfJ4tyqd198t6vQ" // string | The public key of the account whose list of owned NFTs you want to get
    includeNfts := false // bool | Whether or not to include NFTs in the response (optional) (default to false)
    includeZeroBalanceHoldings := false // bool | Whether or not to include holdings that have zero balance. This indicates that the wallet held this token or NFT in the past, but no longer holds it. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaGetTokensBelongingToWallet(context.Background(), network, publicKey).IncludeNfts(includeNfts).IncludeZeroBalanceHoldings(includeZeroBalanceHoldings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaGetTokensBelongingToWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetTokensBelongingToWallet`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaGetTokensBelongingToWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**publicKey** | **string** | The public key of the account whose list of owned NFTs you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetTokensBelongingToWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeNfts** | **bool** | Whether or not to include NFTs in the response | [default to false]
 **includeZeroBalanceHoldings** | **bool** | Whether or not to include holdings that have zero balance. This indicates that the wallet held this token or NFT in the past, but no longer holds it. | [default to false]

### Return type

**[]map[string]interface{}**

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetWalletTransactions

> []string SolanaGetWalletTransactions(ctx, network, publicKey).Execute()

Get address's associated transaction signatures



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
    publicKey := "GKNcUmNacSJo4S2Kq3DuYRYRGw3sNUfJ4tyqd198t6vQ" // string | The public key of the account whose list of signatures you want to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaGetWalletTransactions(context.Background(), network, publicKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaGetWalletTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetWalletTransactions`: []string
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaGetWalletTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**publicKey** | **string** | The public key of the account whose list of signatures you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetWalletTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaTransfer

> DoubleTransferResponse SolanaTransfer(ctx).TransferRequest(transferRequest).Execute()

Transfer SOL, a token, or an NFT to another address



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
    transferRequest := *openapiclient.NewTransferRequest("RecipientAddress_example") // TransferRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SolanaWalletApi.SolanaTransfer(context.Background()).TransferRequest(transferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaTransfer`: DoubleTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaWalletApi.SolanaTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferRequest** | [**TransferRequest**](TransferRequest.md) |  | 

### Return type

[**DoubleTransferResponse**](DoubleTransferResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

