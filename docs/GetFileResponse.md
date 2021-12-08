# GetFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** | File size in kilobytes | [optional] 
**UploadTime** | Pointer to **string** | The time when the file was uploaded | [optional] 

## Methods

### NewGetFileResponse

`func NewGetFileResponse() *GetFileResponse`

NewGetFileResponse instantiates a new GetFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFileResponseWithDefaults

`func NewGetFileResponseWithDefaults() *GetFileResponse`

NewGetFileResponseWithDefaults instantiates a new GetFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *GetFileResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *GetFileResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *GetFileResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *GetFileResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSize

`func (o *GetFileResponse) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetFileResponse) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetFileResponse) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetFileResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUploadTime

`func (o *GetFileResponse) GetUploadTime() string`

GetUploadTime returns the UploadTime field if non-nil, zero value otherwise.

### GetUploadTimeOk

`func (o *GetFileResponse) GetUploadTimeOk() (*string, bool)`

GetUploadTimeOk returns a tuple with the UploadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTime

`func (o *GetFileResponse) SetUploadTime(v string)`

SetUploadTime sets UploadTime field to given value.

### HasUploadTime

`func (o *GetFileResponse) HasUploadTime() bool`

HasUploadTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


