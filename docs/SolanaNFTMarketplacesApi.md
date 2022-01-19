# \SolanaNFTMarketplacesApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolanaBuyNFT**](SolanaNFTMarketplacesApi.md#SolanaBuyNFT) | **Post** /solana/nft/marketplaces/{exchange}/buy/{network}/{mint_address} | Buy
[**SolanaDelistNFT**](SolanaNFTMarketplacesApi.md#SolanaDelistNFT) | **Post** /solana/nft/marketplaces/{exchange}/delist/{network}/{mint_address} | Delist
[**SolanaGetNFTListing**](SolanaNFTMarketplacesApi.md#SolanaGetNFTListing) | **Get** /solana/nft/marketplaces/listing/{network}/{mint_address} | Get NFT Listing
[**SolanaListNFT**](SolanaNFTMarketplacesApi.md#SolanaListNFT) | **Post** /solana/nft/marketplaces/{exchange}/list/{network}/{mint_address} | List



## SolanaBuyNFT

> BuyResponse SolanaBuyNFT(ctx, network, exchange, mintAddress).BuyRequest(buyRequest).Execute()

Buy



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
    network := "devnet" // string | The network ID
    exchange := "solsea" // string | The NFT exchange to interact with
    mintAddress := "7GA16mQup7ESJbaD6n49VCwVG9kRkyQDzWKhBSLjbYqr" // string | The mint address of the NFT you want to buy
    buyRequest := *openapiclient.NewBuyRequest(openapiclient.Wallet{B58PrivateKey: openapiclient.NewB58PrivateKey("B58PrivateKey_example")}, float32(10000000)) // BuyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaNFTMarketplacesApi.SolanaBuyNFT(context.Background(), network, exchange, mintAddress).BuyRequest(buyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTMarketplacesApi.SolanaBuyNFT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaBuyNFT`: BuyResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTMarketplacesApi.SolanaBuyNFT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID | 
**exchange** | **string** | The NFT exchange to interact with | 
**mintAddress** | **string** | The mint address of the NFT you want to buy | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaBuyNFTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **buyRequest** | [**BuyRequest**](BuyRequest.md) |  | 

### Return type

[**BuyResponse**](BuyResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaDelistNFT

> DelistResponse SolanaDelistNFT(ctx, network, exchange, mintAddress).DelistRequest(delistRequest).Execute()

Delist



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
    network := "devnet" // string | The network ID
    exchange := "solsea" // string | The NFT exchange to interact with
    mintAddress := "7GA16mQup7ESJbaD6n49VCwVG9kRkyQDzWKhBSLjbYqr" // string | The mint address of the NFT you want to buy
    delistRequest := *openapiclient.NewDelistRequest(openapiclient.Wallet{B58PrivateKey: openapiclient.NewB58PrivateKey("B58PrivateKey_example")}) // DelistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaNFTMarketplacesApi.SolanaDelistNFT(context.Background(), network, exchange, mintAddress).DelistRequest(delistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTMarketplacesApi.SolanaDelistNFT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaDelistNFT`: DelistResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTMarketplacesApi.SolanaDelistNFT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID | 
**exchange** | **string** | The NFT exchange to interact with | 
**mintAddress** | **string** | The mint address of the NFT you want to buy | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaDelistNFTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **delistRequest** | [**DelistRequest**](DelistRequest.md) |  | 

### Return type

[**DelistResponse**](DelistResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetNFTListing

> GetNFTListingResponse SolanaGetNFTListing(ctx, network, mintAddress).Execute()

Get NFT Listing



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
    network := "devnet" // string | The network ID
    mintAddress := "7GA16mQup7ESJbaD6n49VCwVG9kRkyQDzWKhBSLjbYqr" // string | The mint address of the NFT you want to buy

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaNFTMarketplacesApi.SolanaGetNFTListing(context.Background(), network, mintAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTMarketplacesApi.SolanaGetNFTListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetNFTListing`: GetNFTListingResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTMarketplacesApi.SolanaGetNFTListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID | 
**mintAddress** | **string** | The mint address of the NFT you want to buy | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetNFTListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNFTListingResponse**](GetNFTListingResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaListNFT

> ListResponse SolanaListNFT(ctx, network, exchange, mintAddress).ListRequest(listRequest).Execute()

List



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
    network := "devnet" // string | The network ID
    exchange := "solsea" // string | The NFT exchange to interact with
    mintAddress := "7GA16mQup7ESJbaD6n49VCwVG9kRkyQDzWKhBSLjbYqr" // string | The mint address of the NFT you want to buy
    listRequest := *openapiclient.NewListRequest(openapiclient.Wallet{B58PrivateKey: openapiclient.NewB58PrivateKey("B58PrivateKey_example")}, float32(10000000)) // ListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaNFTMarketplacesApi.SolanaListNFT(context.Background(), network, exchange, mintAddress).ListRequest(listRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaNFTMarketplacesApi.SolanaListNFT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaListNFT`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaNFTMarketplacesApi.SolanaListNFT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID | 
**exchange** | **string** | The NFT exchange to interact with | 
**mintAddress** | **string** | The mint address of the NFT you want to buy | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaListNFTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **listRequest** | [**ListRequest**](ListRequest.md) |  | 

### Return type

[**ListResponse**](ListResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

