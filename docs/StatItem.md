# StatItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | The anonymous ID of the user who called your endpoint  | [optional] 
**TimeStamp** | Pointer to **float32** | The time stamp of the ten-minute interval  | [optional] 
**EndpointId** | Pointer to **string** | The full path of the endpoint (e.g., &#x60;project_id&#x60;/&#x60;version&#x60;/&#x60;path&#x60;)  | [optional] 
**ApiCalls** | Pointer to **float32** | The number of API calls by this user to this endpoint for the ten-minute interval  | [optional] 
**Version** | Pointer to **string** | The version of the endpoint | [optional] 
**Path** | Pointer to **string** | The path of the endpoint (not inclusive of the &#x60;version&#x60; or &#x60;project_id&#x60;) | [optional] 

## Methods

### NewStatItem

`func NewStatItem() *StatItem`

NewStatItem instantiates a new StatItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatItemWithDefaults

`func NewStatItemWithDefaults() *StatItem`

NewStatItemWithDefaults instantiates a new StatItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *StatItem) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *StatItem) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *StatItem) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *StatItem) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetTimeStamp

`func (o *StatItem) GetTimeStamp() float32`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *StatItem) GetTimeStampOk() (*float32, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *StatItem) SetTimeStamp(v float32)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *StatItem) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetEndpointId

`func (o *StatItem) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *StatItem) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *StatItem) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *StatItem) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetApiCalls

`func (o *StatItem) GetApiCalls() float32`

GetApiCalls returns the ApiCalls field if non-nil, zero value otherwise.

### GetApiCallsOk

`func (o *StatItem) GetApiCallsOk() (*float32, bool)`

GetApiCallsOk returns a tuple with the ApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCalls

`func (o *StatItem) SetApiCalls(v float32)`

SetApiCalls sets ApiCalls field to given value.

### HasApiCalls

`func (o *StatItem) HasApiCalls() bool`

HasApiCalls returns a boolean if a field has been set.

### GetVersion

`func (o *StatItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StatItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StatItem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StatItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPath

`func (o *StatItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StatItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StatItem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StatItem) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


