# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**AccountContext**](AccountContext.md) |  | [optional] 
**Value** | Pointer to [**AccountValue**](AccountValue.md) |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Account) GetContext() AccountContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Account) GetContextOk() (*AccountContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Account) SetContext(v AccountContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *Account) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetValue

`func (o *Account) GetValue() AccountValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Account) GetValueOk() (*AccountValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Account) SetValue(v AccountValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *Account) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


