# \SolanaCandyMachineApi

All URIs are relative to *https://api.theblockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolanaCreateTestCandyMachine**](SolanaCandyMachineApi.md#SolanaCreateTestCandyMachine) | **Post** /solana/nft/candy_machine | Create a test CM
[**SolanaGetAllNFTsFromCandyMachine**](SolanaCandyMachineApi.md#SolanaGetAllNFTsFromCandyMachine) | **Get** /solana/nft/candy_machine/{network}/{candy_machine_id}/nfts | Get CM&#39;s NFTs  
[**SolanaGetCandyMachineMetadata**](SolanaCandyMachineApi.md#SolanaGetCandyMachineMetadata) | **Post** /solana/nft/candy_machine/metadata | Get a CM&#39;s metadata 
[**SolanaListAllCandyMachines**](SolanaCandyMachineApi.md#SolanaListAllCandyMachines) | **Get** /solana/nft/candy_machine/list | List all CMs
[**SolanaMintFromCandyMachine**](SolanaCandyMachineApi.md#SolanaMintFromCandyMachine) | **Post** /solana/nft/candy_machine/mint | Mint from a CM
[**SolanaSearchCandyMachines**](SolanaCandyMachineApi.md#SolanaSearchCandyMachines) | **Post** /solana/nft/candy_machine/search | Search CMs



## SolanaCreateTestCandyMachine

> CreateTestCandyMachineResponse SolanaCreateTestCandyMachine(ctx).CreateTestCandyMachineRequest(createTestCandyMachineRequest).Execute()

Create a test CM



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
    createTestCandyMachineRequest := *openapiclient.NewCreateTestCandyMachineRequest(openapiclient.Wallet{B58PrivateKey: openapiclient.NewB58PrivateKey("B58PrivateKey_example")}) // CreateTestCandyMachineRequest |  (optional)

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

Get CM's NFTs  



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


## SolanaGetCandyMachineMetadata

> GetCandyMetadataResponse SolanaGetCandyMachineMetadata(ctx).GetCandyMetadataRequest(getCandyMetadataRequest).Execute()

Get a CM's metadata 



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
    getCandyMetadataRequest := *openapiclient.NewGetCandyMetadataRequest() // GetCandyMetadataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaCandyMachineApi.SolanaGetCandyMachineMetadata(context.Background()).GetCandyMetadataRequest(getCandyMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaCandyMachineApi.SolanaGetCandyMachineMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetCandyMachineMetadata`: GetCandyMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `SolanaCandyMachineApi.SolanaGetCandyMachineMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetCandyMachineMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCandyMetadataRequest** | [**GetCandyMetadataRequest**](GetCandyMetadataRequest.md) |  | 

### Return type

[**GetCandyMetadataResponse**](GetCandyMetadataResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaListAllCandyMachines

> interface{} SolanaListAllCandyMachines(ctx).Execute()

List all CMs



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
    resp, r, err := api_client.SolanaCandyMachineApi.SolanaListAllCandyMachines(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaCandyMachineApi.SolanaListAllCandyMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaListAllCandyMachines`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SolanaCandyMachineApi.SolanaListAllCandyMachines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaListAllCandyMachinesRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolanaMintFromCandyMachine

> MintNFTResponse SolanaMintFromCandyMachine(ctx).MintNFTRequest(mintNFTRequest).Execute()

Mint from a CM



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
    mintNFTRequest := *openapiclient.NewMintNFTRequest(openapiclient.Wallet{B58PrivateKey: openapiclient.NewB58PrivateKey("B58PrivateKey_example")}, "ConfigAddress_example") // MintNFTRequest |  (optional)

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


## SolanaSearchCandyMachines

> []string SolanaSearchCandyMachines(ctx).CandyMachineSearchRequest(candyMachineSearchRequest).Execute()

Search CMs



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
    candyMachineSearchRequest := *openapiclient.NewCandyMachineSearchRequest() // CandyMachineSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaCandyMachineApi.SolanaSearchCandyMachines(context.Background()).CandyMachineSearchRequest(candyMachineSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaCandyMachineApi.SolanaSearchCandyMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaSearchCandyMachines`: []string
    fmt.Fprintf(os.Stdout, "Response from `SolanaCandyMachineApi.SolanaSearchCandyMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolanaSearchCandyMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **candyMachineSearchRequest** | [**CandyMachineSearchRequest**](CandyMachineSearchRequest.md) |  | 

### Return type

**[]string**

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

