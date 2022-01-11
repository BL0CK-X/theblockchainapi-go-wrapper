# \FileApi

All URIs are relative to *https://api.blockchainapi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadFile**](FileApi.md#UploadFile) | **Post** /file | Upload a file



## UploadFile

> GetFileResponse UploadFile(ctx).Execute()

Upload a file



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
    resp, r, err := api_client.FileApi.UploadFile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.UploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFile`: GetFileResponse
    fmt.Fprintf(os.Stdout, "Response from `FileApi.UploadFile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


### Return type

[**GetFileResponse**](GetFileResponse.md)

### Authorization

[APIKeyID](../README.md#APIKeyID), [APISecretKey](../README.md#APISecretKey)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

