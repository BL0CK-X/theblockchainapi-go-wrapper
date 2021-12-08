# \SolanaTransactionApi

All URIs are relative to *https://api.theblockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolanaGetTransaction**](SolanaTransactionApi.md#SolanaGetTransaction) | **Get** /solana/transaction/{network}/{tx_signature} | Get the details of a transaction made on Solana



## SolanaGetTransaction

> Transaction SolanaGetTransaction(ctx, network, txSignature).Execute()

Get the details of a transaction made on Solana



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
    txSignature := "5wHu1qwD7q5ifaN5nwdcDqNFo53GJqa7nLp2BeeEpcHCusb4GzARz4GjgzsEHMkBMgCJMGa6GSQ1VG96Exv8kt2W" // string | The transaction signature of the transaction

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SolanaTransactionApi.SolanaGetTransaction(context.Background(), network, txSignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SolanaTransactionApi.SolanaGetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolanaGetTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `SolanaTransactionApi.SolanaGetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**network** | **string** | The network ID (devnet, mainnet-beta) | 
**txSignature** | **string** | The transaction signature of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiSolanaGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Transaction**](Transaction.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

