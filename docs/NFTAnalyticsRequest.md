# NFTAnalyticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MintAddresses** | **[]string** | The list of mint addresses. Each address must be a valid public key. | 
**StartTime** | Pointer to **float32** | Response will not include transactions before &#x60;start_time&#x60;. Set this to &#x60;-1&#x60; in order to get the entire history we have stored.  | [optional] 
**EndTime** | Pointer to **float32** | Response will not include transactions after &#x60;end_time&#x60;   | [optional] 

## Methods

### NewNFTAnalyticsRequest

`func NewNFTAnalyticsRequest(mintAddresses []string, ) *NFTAnalyticsRequest`

NewNFTAnalyticsRequest instantiates a new NFTAnalyticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTAnalyticsRequestWithDefaults

`func NewNFTAnalyticsRequestWithDefaults() *NFTAnalyticsRequest`

NewNFTAnalyticsRequestWithDefaults instantiates a new NFTAnalyticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMintAddresses

`func (o *NFTAnalyticsRequest) GetMintAddresses() []string`

GetMintAddresses returns the MintAddresses field if non-nil, zero value otherwise.

### GetMintAddressesOk

`func (o *NFTAnalyticsRequest) GetMintAddressesOk() (*[]string, bool)`

GetMintAddressesOk returns a tuple with the MintAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddresses

`func (o *NFTAnalyticsRequest) SetMintAddresses(v []string)`

SetMintAddresses sets MintAddresses field to given value.


### GetStartTime

`func (o *NFTAnalyticsRequest) GetStartTime() float32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *NFTAnalyticsRequest) GetStartTimeOk() (*float32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *NFTAnalyticsRequest) SetStartTime(v float32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *NFTAnalyticsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *NFTAnalyticsRequest) GetEndTime() float32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *NFTAnalyticsRequest) GetEndTimeOk() (*float32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *NFTAnalyticsRequest) SetEndTime(v float32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *NFTAnalyticsRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


