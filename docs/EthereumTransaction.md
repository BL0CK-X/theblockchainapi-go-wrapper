# EthereumTransaction

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

## Methods

### NewEthereumTransaction

`func NewEthereumTransaction() *EthereumTransaction`

NewEthereumTransaction instantiates a new EthereumTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthereumTransactionWithDefaults

`func NewEthereumTransactionWithDefaults() *EthereumTransaction`

NewEthereumTransactionWithDefaults instantiates a new EthereumTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessList

`func (o *EthereumTransaction) GetAccessList() map[string]interface{}`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *EthereumTransaction) GetAccessListOk() (*map[string]interface{}, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *EthereumTransaction) SetAccessList(v map[string]interface{})`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *EthereumTransaction) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetBlockHash

`func (o *EthereumTransaction) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *EthereumTransaction) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *EthereumTransaction) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *EthereumTransaction) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockNumber

`func (o *EthereumTransaction) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *EthereumTransaction) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *EthereumTransaction) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *EthereumTransaction) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetChainId

`func (o *EthereumTransaction) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *EthereumTransaction) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *EthereumTransaction) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *EthereumTransaction) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetFrom

`func (o *EthereumTransaction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EthereumTransaction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EthereumTransaction) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EthereumTransaction) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGas

`func (o *EthereumTransaction) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *EthereumTransaction) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *EthereumTransaction) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *EthereumTransaction) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *EthereumTransaction) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *EthereumTransaction) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *EthereumTransaction) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *EthereumTransaction) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetHash

`func (o *EthereumTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *EthereumTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *EthereumTransaction) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *EthereumTransaction) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetInput

`func (o *EthereumTransaction) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EthereumTransaction) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EthereumTransaction) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *EthereumTransaction) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *EthereumTransaction) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *EthereumTransaction) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *EthereumTransaction) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *EthereumTransaction) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *EthereumTransaction) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *EthereumTransaction) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *EthereumTransaction) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *EthereumTransaction) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### GetNonce

`func (o *EthereumTransaction) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *EthereumTransaction) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *EthereumTransaction) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *EthereumTransaction) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetR

`func (o *EthereumTransaction) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *EthereumTransaction) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *EthereumTransaction) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *EthereumTransaction) HasR() bool`

HasR returns a boolean if a field has been set.

### GetS

`func (o *EthereumTransaction) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *EthereumTransaction) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *EthereumTransaction) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *EthereumTransaction) HasS() bool`

HasS returns a boolean if a field has been set.

### GetTo

`func (o *EthereumTransaction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EthereumTransaction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EthereumTransaction) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *EthereumTransaction) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTransactionIndex

`func (o *EthereumTransaction) GetTransactionIndex() string`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *EthereumTransaction) GetTransactionIndexOk() (*string, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *EthereumTransaction) SetTransactionIndex(v string)`

SetTransactionIndex sets TransactionIndex field to given value.

### HasTransactionIndex

`func (o *EthereumTransaction) HasTransactionIndex() bool`

HasTransactionIndex returns a boolean if a field has been set.

### GetType

`func (o *EthereumTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EthereumTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EthereumTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EthereumTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetV

`func (o *EthereumTransaction) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *EthereumTransaction) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *EthereumTransaction) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *EthereumTransaction) HasV() bool`

HasV returns a boolean if a field has been set.

### GetValue

`func (o *EthereumTransaction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EthereumTransaction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EthereumTransaction) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EthereumTransaction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


