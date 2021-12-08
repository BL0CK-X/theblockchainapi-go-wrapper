# \SolanaCandyMachineApi

All URIs are relative to *https://api.theblockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolanaCreateTestCandyMachine**](SolanaCandyMachineApi.md#SolanaCreateTestCandyMachine) | **Post** /solana/nft/candy_machine | Create a test candy machine 
[**SolanaGetAllNFTsFromCandyMachine**](SolanaCandyMachineApi.md#SolanaGetAllNFTsFromCandyMachine) | **Get** /solana/nft/candy_machine/{network}/{candy_machine_id}/nfts | Get the list of all NFTs (minted and unminted) from a Solana Candy Machine 
[**SolanaGetCandyMachineConfigurationDetails**](SolanaCandyMachineApi.md#SolanaGetCandyMachineConfigurationDetails) | **Post** /solana/nft/candy_machine/config/info | Get the details of a Solana Candy Machine configuration 
[**SolanaGetCandyMachineDetails**](SolanaCandyMachineApi.md#SolanaGetCandyMachineDetails) | **Post** /solana/nft/candy_machine/info | Get a Metaplex candy machine&#39;s details 
[**SolanaGetNFTsMintedFromCandyMachine**](SolanaCandyMachineApi.md#SolanaGetNFTsMintedFromCandyMachine) | **Post** /solana/nft/candy_machine/nfts | Get the list of NFTs minted from a Solana Candy Machine 
[**SolanaMintFromCandyMachine**](SolanaCandyMachineApi.md#SolanaMintFromCandyMachine) | **Post** /solana/nft/candy_machine/mint | Mint an NFT from a Metaplex candy machine



## SolanaCreateTestCandyMachine

> CreateTestCandyMachineResponse SolanaCreateTestCandyMachine(ctx).CreateTestCandyMachineRequest(createTestCandyMachineRequest).Execute()

Create a test candy machine 



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
    createTestCandyMachineRequest := *openapiclient.NewCreateTestCandyMachineRequest("SecretRecoveryPhrase_example") // CreateTestCandyMachineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaCandyMachineApi.SolanaCreateTestCandyMachine(context.Background()).CreateTestCandyMachineRequest(createTestCandyMachineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaCandyMachineApi.SolanaCreateTestCandyMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaCreateTestCandyMachine`: CreateTestCandyMachineResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaCandyMachineApi.SolanaCreateTestCandyMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaCreateTestCandyMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTestCandyMachineRequest** | [**CreateTestCandyMachineRequest**](CreateTestCandyMachineRequest.md) |  | 

### Return type

[**CreateTestCandyMachineResponse**](CreateTestCandyMachineResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetAllNFTsFromCandyMachine

> GetAllNFTsResponse SolanaGetAllNFTsFromCandyMachine(ctx, network, candyMachineId).Execute()

Get the list of all NFTs (minted and unminted) from a Solana Candy Machine 



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
    candyMachineId := "FmkrvXRenCGtwBHw3VtBcExp8eTdnau97upaewF4GUEX" // string | The ID of the candy machine

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaCandyMachineApi.SolanaGetAllNFTsFromCandyMachine(context.Background(), network, candyMachineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaCandyMachineApi.SolanaGetAllNFTsFromCandyMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetAllNFTsFromCandyMachine`: GetAllNFTsResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaCandyMachineApi.SolanaGetAllNFTsFromCandyMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**candyMachineId** | **string** | The ID of the candy machine | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetAllNFTsFromCandyMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAllNFTsResponse**](GetAllNFTsResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetCandyMachineConfigurationDetails

> GetConfigInfoResponse SolanaGetCandyMachineConfigurationDetails(ctx).GetConfigInfoRequest(getConfigInfoRequest).Execute()

Get the details of a Solana Candy Machine configuration 



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
    getConfigInfoRequest := *openapiclient.NewGetConfigInfoRequest("ConfigAddress_example") // GetConfigInfoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaCandyMachineApi.SolanaGetCandyMachineConfigurationDetails(context.Background()).GetConfigInfoRequest(getConfigInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaCandyMachineApi.SolanaGetCandyMachineConfigurationDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetCandyMachineConfigurationDetails`: GetConfigInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaCandyMachineApi.SolanaGetCandyMachineConfigurationDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetCandyMachineConfigurationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getConfigInfoRequest** | [**GetConfigInfoRequest**](GetConfigInfoRequest.md) |  | 

### Return type

[**GetConfigInfoResponse**](GetConfigInfoResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetCandyMachineDetails

> GetCandyDetailsResponse SolanaGetCandyMachineDetails(ctx).GetCandyDetailsRequest(getCandyDetailsRequest).Execute()

Get a Metaplex candy machine's details 



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
    getCandyDetailsRequest := *openapiclient.NewGetCandyDetailsRequest("CandyMachineId_example") // GetCandyDetailsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaCandyMachineApi.SolanaGetCandyMachineDetails(context.Background()).GetCandyDetailsRequest(getCandyDetailsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaCandyMachineApi.SolanaGetCandyMachineDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetCandyMachineDetails`: GetCandyDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaCandyMachineApi.SolanaGetCandyMachineDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetCandyMachineDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCandyDetailsRequest** | [**GetCandyDetailsRequest**](GetCandyDetailsRequest.md) |  | 

### Return type

[**GetCandyDetailsResponse**](GetCandyDetailsResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaGetNFTsMintedFromCandyMachine

> []GetMintedNFTsResponse SolanaGetNFTsMintedFromCandyMachine(ctx).GetMintedNFTsRequest(getMintedNFTsRequest).Execute()

Get the list of NFTs minted from a Solana Candy Machine 



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
    getMintedNFTsRequest := *openapiclient.NewGetMintedNFTsRequest("CandyMachineId_example") // GetMintedNFTsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaCandyMachineApi.SolanaGetNFTsMintedFromCandyMachine(context.Background()).GetMintedNFTsRequest(getMintedNFTsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaCandyMachineApi.SolanaGetNFTsMintedFromCandyMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetNFTsMintedFromCandyMachine`: []GetMintedNFTsResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaCandyMachineApi.SolanaGetNFTsMintedFromCandyMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetNFTsMintedFromCandyMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getMintedNFTsRequest** | [**GetMintedNFTsRequest**](GetMintedNFTsRequest.md) |  | 

### Return type

[**[]GetMintedNFTsResponse**](GetMintedNFTsResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaMintFromCandyMachine

> MintNFTResponse SolanaMintFromCandyMachine(ctx).MintNFTRequest(mintNFTRequest).Execute()

Mint an NFT from a Metaplex candy machine



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
    mintNFTRequest := *openapiclient.NewMintNFTRequest("ConfigAddress_example", "SecretRecoveryPhrase_example") // MintNFTRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaCandyMachineApi.SolanaMintFromCandyMachine(context.Background()).MintNFTRequest(mintNFTRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaCandyMachineApi.SolanaMintFromCandyMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaMintFromCandyMachine`: MintNFTResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaCandyMachineApi.SolanaMintFromCandyMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaMintFromCandyMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintNFTRequest** | [**MintNFTRequest**](MintNFTRequest.md) |  | 

### Return type

[**MintNFTResponse**](MintNFTResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

