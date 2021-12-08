# GetConfigInfoResponseCreators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Public key address of a creator | [optional] 
**Share** | Pointer to **float32** | The share of the royalty that the creator receives | [optional] 

## Methods

### NewGetConfigInfoResponseCreators

`func NewGetConfigInfoResponseCreators() *GetConfigInfoResponseCreators`

NewGetConfigInfoResponseCreators instantiates a new GetConfigInfoResponseCreators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfigInfoResponseCreatorsWithDefaults

`func NewGetConfigInfoResponseCreatorsWithDefaults() *GetConfigInfoResponseCreators`

NewGetConfigInfoResponseCreatorsWithDefaults instantiates a new GetConfigInfoResponseCreators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetConfigInfoResponseCreators) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetConfigInfoResponseCreators) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetConfigInfoResponseCreators) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetConfigInfoResponseCreators) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetShare

`func (o *GetConfigInfoResponseCreators) GetShare() float32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *GetConfigInfoResponseCreators) GetShareOk() (*float32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *GetConfigInfoResponseCreators) SetShare(v float32)`

SetShare sets Share field to given value.

### HasShare

`func (o *GetConfigInfoResponseCreators) HasShare() bool`

HasShare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


