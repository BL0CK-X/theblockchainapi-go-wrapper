# GetCandyMetadataResponseCreators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The public key of a creator | [optional] 
**Share** | Pointer to **float32** | The share of the royalty that each creator receives | [optional] 

## Methods

### NewGetCandyMetadataResponseCreators

`func NewGetCandyMetadataResponseCreators() *GetCandyMetadataResponseCreators`

NewGetCandyMetadataResponseCreators instantiates a new GetCandyMetadataResponseCreators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCandyMetadataResponseCreatorsWithDefaults

`func NewGetCandyMetadataResponseCreatorsWithDefaults() *GetCandyMetadataResponseCreators`

NewGetCandyMetadataResponseCreatorsWithDefaults instantiates a new GetCandyMetadataResponseCreators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetCandyMetadataResponseCreators) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetCandyMetadataResponseCreators) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetCandyMetadataResponseCreators) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetCandyMetadataResponseCreators) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetShare

`func (o *GetCandyMetadataResponseCreators) GetShare() float32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *GetCandyMetadataResponseCreators) GetShareOk() (*float32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *GetCandyMetadataResponseCreators) SetShare(v float32)`

SetShare sets Share field to given value.

### HasShare

`func (o *GetCandyMetadataResponseCreators) HasShare() bool`

HasShare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


