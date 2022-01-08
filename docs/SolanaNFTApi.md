# \SolanaNFTApi

All URIs are relative to *https://api.theblockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolanaCreateNFT**](SolanaNFTApi.md#SolanaCreateNFT) | **Post** /solana/nft | Create an NFT on Solana
[**SolanaGetNFT**](SolanaNFTApi.md#SolanaGetNFT) | **Get** /solana/nft/{network}/{mint_address} | Get an NFT&#39;s metadata
[**SolanaGetNFTMintFee**](SolanaNFTApi.md#SolanaGetNFTMintFee) | **Get** /solana/nft/mint/fee | Get the NFT mint fee
[**SolanaGetNFTOwner**](SolanaNFTApi.md#SolanaGetNFTOwner) | **Get** /solana/nft/{network}/{mint_address}/owner | Get owner of an NFT
[**SolanaGetNFTsCandyMachineId**](SolanaNFTApi.md#SolanaGetNFTsCandyMachineId) | **Post** /solana/nft/candy_machine_id | Get the ID of the candy machine of an NFT 
[**SolanaSearchNFTs**](SolanaNFTApi.md#SolanaSearchNFTs) | **Post** /solana/nft/search | Search NFTs on Solana



## SolanaCreateNFT

> NFT SolanaCreateNFT(ctx).NFTMintRequest(nFTMintRequest).Execute()

Create an NFT on Solana



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
    nFTMintRequest := *openapiclient.NewNFTMintRequest(openapiclient.Wallet{B58PrivateKey: openapiclient.NewB58PrivateKey("B58PrivateKey_example")}) // NFTMintRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaNFTApi.SolanaCreateNFT(context.Background()).NFTMintRequest(nFTMintRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTApi.SolanaCreateNFT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaCreateNFT`: NFT
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTApi.SolanaCreateNFT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaCreateNFTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nFTMintRequest** | [**NFTMintRequest**](NFTMintRequest.md) |  | 

### Return type

[**NFT**](NFT.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetNFT

> NFT SolanaGetNFT(ctx, network, mintAddress).Execute()

Get an NFT's metadata



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
    mintAddress := "EEr5yQpNXf7Bru6Rt5podx56HGW9CEehXqgRGh2wa71w" // string | The mint address of the NFT

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaNFTApi.SolanaGetNFT(context.Background(), network, mintAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTApi.SolanaGetNFT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetNFT`: NFT
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTApi.SolanaGetNFT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**mintAddress** | **string** | The mint address of the NFT | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetNFTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NFT**](NFT.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetNFTMintFee

> NFTMintFee SolanaGetNFTMintFee(ctx).Execute()

Get the NFT mint fee



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
    resp, r, err := api_client.SolanaNFTApi.SolanaGetNFTMintFee(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTApi.SolanaGetNFTMintFee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetNFTMintFee`: NFTMintFee
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTApi.SolanaGetNFTMintFee`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetNFTMintFeeRequest struct via the builder pattern


### Return type

[**NFTMintFee**](NFTMintFee.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetNFTOwner

> NFTOwnerResponse SolanaGetNFTOwner(ctx, network, mintAddress).Execute()

Get owner of an NFT



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
    mintAddress := "4zH3Rwm1QXdfTSUqsYmeUBY4QqQmQEXJVbv4ErSK736Q" // string | The mint address of the NFT

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaNFTApi.SolanaGetNFTOwner(context.Background(), network, mintAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTApi.SolanaGetNFTOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetNFTOwner`: NFTOwnerResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTApi.SolanaGetNFTOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**mintAddress** | **string** | The mint address of the NFT | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetNFTOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NFTOwnerResponse**](NFTOwnerResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetNFTsCandyMachineId

> GetCandyMachineIDResponse SolanaGetNFTsCandyMachineId(ctx).GetCandyMachineIDRequest(getCandyMachineIDRequest).Execute()

Get the ID of the candy machine of an NFT 



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
    getCandyMachineIDRequest := *openapiclient.NewGetCandyMachineIDRequest("MintAddress_example") // GetCandyMachineIDRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaNFTApi.SolanaGetNFTsCandyMachineId(context.Background()).GetCandyMachineIDRequest(getCandyMachineIDRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTApi.SolanaGetNFTsCandyMachineId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetNFTsCandyMachineId`: GetCandyMachineIDResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTApi.SolanaGetNFTsCandyMachineId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetNFTsCandyMachineIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCandyMachineIDRequest** | [**GetCandyMachineIDRequest**](GetCandyMachineIDRequest.md) |  | 

### Return type

[**GetCandyMachineIDResponse**](GetCandyMachineIDResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaSearchNFTs

> []NFTSearchResponse SolanaSearchNFTs(ctx).NFTSearchRequest(nFTSearchRequest).Execute()

Search NFTs on Solana



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
    nFTSearchRequest := *openapiclient.NewNFTSearchRequest() // NFTSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaNFTApi.SolanaSearchNFTs(context.Background()).NFTSearchRequest(nFTSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTApi.SolanaSearchNFTs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaSearchNFTs`: []NFTSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTApi.SolanaSearchNFTs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaSearchNFTsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nFTSearchRequest** | [**NFTSearchRequest**](NFTSearchRequest.md) |  | 

### Return type

[**[]NFTSearchResponse**](NFTSearchResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

