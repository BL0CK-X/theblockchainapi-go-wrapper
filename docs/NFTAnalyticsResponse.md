# NFTAnalyticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **interface{}** | The start time used in the API request.  | [optional] 
**EndTime** | Pointer to **interface{}** | The end time used in the API request.   | [optional] 
**TransactionHistory** | Pointer to [**NFTAnalyticsResponseTransactionHistory**](NFTAnalyticsResponseTransactionHistory.md) |  | [optional] 
**Floor** | Pointer to **float32** | The minimum active listing price for the collection in the given time period. The listing must have been processed before &#x60;end_time&#x60; and still active (not delisted or purchased) by &#x60;end_time&#x60; in order to affect the floor calculation.  | [optional] 
**Volume** | Pointer to **float32** | The sum of the sale prices for the given time period.  | [optional] 

## Methods

### NewNFTAnalyticsResponse

`func NewNFTAnalyticsResponse() *NFTAnalyticsResponse`

NewNFTAnalyticsResponse instantiates a new NFTAnalyticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTAnalyticsResponseWithDefaults

`func NewNFTAnalyticsResponseWithDefaults() *NFTAnalyticsResponse`

NewNFTAnalyticsResponseWithDefaults instantiates a new NFTAnalyticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *NFTAnalyticsResponse) GetStartTime() interface{}`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *NFTAnalyticsResponse) GetStartTimeOk() (*interface{}, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *NFTAnalyticsResponse) SetStartTime(v interface{})`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *NFTAnalyticsResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *NFTAnalyticsResponse) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *NFTAnalyticsResponse) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *NFTAnalyticsResponse) GetEndTime() interface{}`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *NFTAnalyticsResponse) GetEndTimeOk() (*interface{}, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *NFTAnalyticsResponse) SetEndTime(v interface{})`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *NFTAnalyticsResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *NFTAnalyticsResponse) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *NFTAnalyticsResponse) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetTransactionHistory

`func (o *NFTAnalyticsResponse) GetTransactionHistory() NFTAnalyticsResponseTransactionHistory`

GetTransactionHistory returns the TransactionHistory field if non-nil, zero value otherwise.

### GetTransactionHistoryOk

`func (o *NFTAnalyticsResponse) GetTransactionHistoryOk() (*NFTAnalyticsResponseTransactionHistory, bool)`

GetTransactionHistoryOk returns a tuple with the TransactionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHistory

`func (o *NFTAnalyticsResponse) SetTransactionHistory(v NFTAnalyticsResponseTransactionHistory)`

SetTransactionHistory sets TransactionHistory field to given value.

### HasTransactionHistory

`func (o *NFTAnalyticsResponse) HasTransactionHistory() bool`

HasTransactionHistory returns a boolean if a field has been set.

### GetFloor

`func (o *NFTAnalyticsResponse) GetFloor() float32`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *NFTAnalyticsResponse) GetFloorOk() (*float32, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *NFTAnalyticsResponse) SetFloor(v float32)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *NFTAnalyticsResponse) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### GetVolume

`func (o *NFTAnalyticsResponse) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *NFTAnalyticsResponse) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *NFTAnalyticsResponse) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *NFTAnalyticsResponse) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


