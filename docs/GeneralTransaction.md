# GeneralTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessList** | Pointer to **map[string]interface{}** |  | [optional] 
**BlockHash** | Pointer to **string** |  | [optional] 
**BlockNumber** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Gas** | Pointer to **string** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **string** |  | [optional] 
**MaxFeePerGas** | Pointer to **string** |  | [optional] 
**MaxPriorityFeePerGas** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**R** | Pointer to **string** |  | [optional] 
**S** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**TransactionIndex** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**V** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**BlockTime** | Pointer to **float32** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Slot** | Pointer to **float32** |  | [optional] 
**Transaction** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGeneralTransaction

`func NewGeneralTransaction() *GeneralTransaction`

NewGeneralTransaction instantiates a new GeneralTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralTransactionWithDefaults

`func NewGeneralTransactionWithDefaults() *GeneralTransaction`

NewGeneralTransactionWithDefaults instantiates a new GeneralTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessList

`func (o *GeneralTransaction) GetAccessList() map[string]interface{}`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *GeneralTransaction) GetAccessListOk() (*map[string]interface{}, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *GeneralTransaction) SetAccessList(v map[string]interface{})`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *GeneralTransaction) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetBlockHash

`func (o *GeneralTransaction) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *GeneralTransaction) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *GeneralTransaction) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *GeneralTransaction) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockNumber

`func (o *GeneralTransaction) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *GeneralTransaction) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *GeneralTransaction) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *GeneralTransaction) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetChainId

`func (o *GeneralTransaction) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *GeneralTransaction) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *GeneralTransaction) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *GeneralTransaction) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetFrom

`func (o *GeneralTransaction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GeneralTransaction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GeneralTransaction) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GeneralTransaction) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGas

`func (o *GeneralTransaction) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *GeneralTransaction) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *GeneralTransaction) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *GeneralTransaction) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *GeneralTransaction) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GeneralTransaction) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GeneralTransaction) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *GeneralTransaction) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetHash

`func (o *GeneralTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GeneralTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GeneralTransaction) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GeneralTransaction) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetInput

`func (o *GeneralTransaction) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GeneralTransaction) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GeneralTransaction) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *GeneralTransaction) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *GeneralTransaction) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *GeneralTransaction) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *GeneralTransaction) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *GeneralTransaction) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *GeneralTransaction) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *GeneralTransaction) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *GeneralTransaction) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *GeneralTransaction) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### GetNonce

`func (o *GeneralTransaction) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GeneralTransaction) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GeneralTransaction) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *GeneralTransaction) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetR

`func (o *GeneralTransaction) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *GeneralTransaction) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *GeneralTransaction) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *GeneralTransaction) HasR() bool`

HasR returns a boolean if a field has been set.

### GetS

`func (o *GeneralTransaction) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *GeneralTransaction) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *GeneralTransaction) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *GeneralTransaction) HasS() bool`

HasS returns a boolean if a field has been set.

### GetTo

`func (o *GeneralTransaction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GeneralTransaction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GeneralTransaction) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GeneralTransaction) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTransactionIndex

`func (o *GeneralTransaction) GetTransactionIndex() string`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *GeneralTransaction) GetTransactionIndexOk() (*string, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *GeneralTransaction) SetTransactionIndex(v string)`

SetTransactionIndex sets TransactionIndex field to given value.

### HasTransactionIndex

`func (o *GeneralTransaction) HasTransactionIndex() bool`

HasTransactionIndex returns a boolean if a field has been set.

### GetType

`func (o *GeneralTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeneralTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeneralTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GeneralTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetV

`func (o *GeneralTransaction) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *GeneralTransaction) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *GeneralTransaction) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *GeneralTransaction) HasV() bool`

HasV returns a boolean if a field has been set.

### GetValue

`func (o *GeneralTransaction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GeneralTransaction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GeneralTransaction) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GeneralTransaction) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetBlockTime

`func (o *GeneralTransaction) GetBlockTime() float32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *GeneralTransaction) GetBlockTimeOk() (*float32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *GeneralTransaction) SetBlockTime(v float32)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *GeneralTransaction) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetMeta

`func (o *GeneralTransaction) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GeneralTransaction) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GeneralTransaction) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GeneralTransaction) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSlot

`func (o *GeneralTransaction) GetSlot() float32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *GeneralTransaction) GetSlotOk() (*float32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *GeneralTransaction) SetSlot(v float32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *GeneralTransaction) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetTransaction

`func (o *GeneralTransaction) GetTransaction() map[string]interface{}`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *GeneralTransaction) GetTransactionOk() (*map[string]interface{}, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *GeneralTransaction) SetTransaction(v map[string]interface{})`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *GeneralTransaction) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


