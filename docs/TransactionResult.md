# TransactionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTime** | Pointer to **float32** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Slot** | Pointer to **float32** |  | [optional] 
**Transaction** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTransactionResult

`func NewTransactionResult() *TransactionResult`

NewTransactionResult instantiates a new TransactionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResultWithDefaults

`func NewTransactionResultWithDefaults() *TransactionResult`

NewTransactionResultWithDefaults instantiates a new TransactionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTime

`func (o *TransactionResult) GetBlockTime() float32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *TransactionResult) GetBlockTimeOk() (*float32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *TransactionResult) SetBlockTime(v float32)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *TransactionResult) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetMeta

`func (o *TransactionResult) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TransactionResult) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TransactionResult) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TransactionResult) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSlot

`func (o *TransactionResult) GetSlot() float32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *TransactionResult) GetSlotOk() (*float32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *TransactionResult) SetSlot(v float32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *TransactionResult) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetTransaction

`func (o *TransactionResult) GetTransaction() map[string]interface{}`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *TransactionResult) GetTransactionOk() (*map[string]interface{}, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *TransactionResult) SetTransaction(v map[string]interface{})`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *TransactionResult) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


