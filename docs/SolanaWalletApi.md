# \SolanaWalletApi

All URIs are relative to *https://api.theblockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolanaDeriveAssociatedTokenAccountAddress**](SolanaWalletApi.md#SolanaDeriveAssociatedTokenAccountAddress) | **Get** /solana/wallet/{public_key}/associated_token_account/{mint_address} | Derive an associated token account address
[**SolanaDerivePublicKey**](SolanaWalletApi.md#SolanaDerivePublicKey) | **Post** /solana/wallet/public_key | Derive public key
[**SolanaGenerateSecretRecoveryPhrase**](SolanaWalletApi.md#SolanaGenerateSecretRecoveryPhrase) | **Post** /solana/wallet/secret_recovery_phrase | Generate secret phrase
[**SolanaGetAirdrop**](SolanaWalletApi.md#SolanaGetAirdrop) | **Post** /solana/wallet/airdrop | Get an airdrop on devnet
[**SolanaGetBalance**](SolanaWalletApi.md#SolanaGetBalance) | **Post** /solana/wallet/balance | Get wallet&#39;s balance in SOL or any SPL
[**SolanaGetNFTsBelongingToWallet**](SolanaWalletApi.md#SolanaGetNFTsBelongingToWallet) | **Get** /solana/wallet/{network}/{public_key}/nfts | Get address&#39;s NFTs
[**SolanaGetTokensBelongingToWallet**](SolanaWalletApi.md#SolanaGetTokensBelongingToWallet) | **Get** /solana/wallet/{network}/{public_key}/tokens | Get address&#39;s tokens and respective balances
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaWalletApi.SolanaDeriveAssociatedTokenAccountAddress(context.Background(), publicKey, mintAddress).Execute()
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
    getPublicKeyRequest := *openapiclient.NewGetPublicKeyRequest("SecretRecoveryPhrase_example") // GetPublicKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaWalletApi.SolanaDerivePublicKey(context.Background()).GetPublicKeyRequest(getPublicKeyRequest).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaWalletApi.SolanaGenerateSecretRecoveryPhrase(context.Background()).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaWalletApi.SolanaGetAirdrop(context.Background()).AirdropRequest(airdropRequest).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaWalletApi.SolanaGetBalance(context.Background()).BalanceRequest(balanceRequest).Execute()
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
    publicKey := "8WRsGBaDcs1X7bHWr4Ad5Nx3bW29BkcmEbyavrLXDC4i" // string | The public key of the account whose list of owned NFTs you want to get

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaWalletApi.SolanaGetNFTsBelongingToWallet(context.Background(), network, publicKey).Execute()
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

> []map[string]interface{} SolanaGetTokensBelongingToWallet(ctx, network, publicKey).ListTokensRequest(listTokensRequest).Execute()

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
    listTokensRequest := *openapiclient.NewListTokensRequest() // ListTokensRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaWalletApi.SolanaGetTokensBelongingToWallet(context.Background(), network, publicKey).ListTokensRequest(listTokensRequest).Execute()
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


 **listTokensRequest** | [**ListTokensRequest**](ListTokensRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaTransfer

> TransferResponse SolanaTransfer(ctx).TransferRequest(transferRequest).Execute()

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
    transferRequest := *openapiclient.NewTransferRequest("RecipientAddress_example", "SecretRecoveryPhrase_example") // TransferRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaWalletApi.SolanaTransfer(context.Background()).TransferRequest(transferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaWalletApi.SolanaTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaTransfer`: TransferResponse
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

[**TransferResponse**](TransferResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

