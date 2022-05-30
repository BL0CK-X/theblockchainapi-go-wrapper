# \TransactionApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransaction**](TransactionApi.md#GetTransaction) | **Get** /{blockchain}/transaction/{network}/{transaction_blockchain_identifier} | Get the details of a transaction made on a blockchain



## GetTransaction

> GeneralTransaction GetTransaction(ctx, blockchain, network, transactionBlockchainIdentifier).Execute()

Get the details of a transaction made on a blockchain



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
    transactionBlockchainIdentifier := "0x5f36b787daa57bfe8568d69e24eae54ccb00720c6edfc826bd4a7b19c525eef4" // string | The transaction signature of the transaction.  Examples: - Solana: `5wHu1qwD7q5ifaN5nwdcDqNFo53GJqa7nLp2BeeEpcHCusb4GzARz4GjgzsEHMkBMgCJMGa6GSQ1VG96Exv8kt2W` - Ethereum: `0x5f36b787daa57bfe8568d69e24eae54ccb00720c6edfc826bd4a7b19c525eef4`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionApi.GetTransaction(context.Background(), blockchain, network, transactionBlockchainIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: GeneralTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | The blockchain you want to use  | 
**network** | **string** | The network of the blockchain you selected  - Solana: &#x60;devnet&#x60;, &#x60;mainnet-beta&#x60; - Ethereum: &#x60;ropsten&#x60;, &#x60;mainnet&#x60;  Defaults when not provided (not applicable to path parameters): - Solana: &#x60;devnet&#x60; - Ethereum: &#x60;ropsten&#x60; | 
**transactionBlockchainIdentifier** | **string** | The transaction signature of the transaction.  Examples: - Solana: &#x60;5wHu1qwD7q5ifaN5nwdcDqNFo53GJqa7nLp2BeeEpcHCusb4GzARz4GjgzsEHMkBMgCJMGa6GSQ1VG96Exv8kt2W&#x60; - Ethereum: &#x60;0x5f36b787daa57bfe8568d69e24eae54ccb00720c6edfc826bd4a7b19c525eef4&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GeneralTransaction**](GeneralTransaction.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

