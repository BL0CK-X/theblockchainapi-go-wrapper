# AccountValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | More info about the account. What are included depends on the type of account. See examples. | [optional] 
**Executable** | Pointer to **bool** | Whether or not this account is marked as executable | [optional] 
**Owner** | Pointer to **string** | The owner of the account | [optional] 
**RentEpoch** | Pointer to **float32** |  | [optional] 

## Methods

### NewAccountValue

`func NewAccountValue() *AccountValue`

NewAccountValue instantiates a new AccountValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountValueWithDefaults

`func NewAccountValueWithDefaults() *AccountValue`

NewAccountValueWithDefaults instantiates a new AccountValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AccountValue) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountValue) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountValue) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AccountValue) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExecutable

`func (o *AccountValue) GetExecutable() bool`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *AccountValue) GetExecutableOk() (*bool, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *AccountValue) SetExecutable(v bool)`

SetExecutable sets Executable field to given value.

### HasExecutable

`func (o *AccountValue) HasExecutable() bool`

HasExecutable returns a boolean if a field has been set.

### GetOwner

`func (o *AccountValue) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccountValue) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccountValue) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccountValue) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRentEpoch

`func (o *AccountValue) GetRentEpoch() float32`

GetRentEpoch returns the RentEpoch field if non-nil, zero value otherwise.

### GetRentEpochOk

`func (o *AccountValue) GetRentEpochOk() (*float32, bool)`

GetRentEpochOk returns a tuple with the RentEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentEpoch

`func (o *AccountValue) SetRentEpoch(v float32)`

SetRentEpoch sets RentEpoch field to given value.

### HasRentEpoch

`func (o *AccountValue) HasRentEpoch() bool`

HasRentEpoch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


