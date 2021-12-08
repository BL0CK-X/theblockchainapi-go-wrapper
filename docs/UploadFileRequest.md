# UploadFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadMethod** | Pointer to **string** |  | [optional] [default to "IPFS"]

## Methods

### NewUploadFileRequest

`func NewUploadFileRequest() *UploadFileRequest`

NewUploadFileRequest instantiates a new UploadFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFileRequestWithDefaults

`func NewUploadFileRequestWithDefaults() *UploadFileRequest`

NewUploadFileRequestWithDefaults instantiates a new UploadFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadMethod

`func (o *UploadFileRequest) GetUploadMethod() string`

GetUploadMethod returns the UploadMethod field if non-nil, zero value otherwise.

### GetUploadMethodOk

`func (o *UploadFileRequest) GetUploadMethodOk() (*string, bool)`

GetUploadMethodOk returns a tuple with the UploadMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadMethod

`func (o *UploadFileRequest) SetUploadMethod(v string)`

SetUploadMethod sets UploadMethod field to given value.

### HasUploadMethod

`func (o *UploadFileRequest) HasUploadMethod() bool`

HasUploadMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


