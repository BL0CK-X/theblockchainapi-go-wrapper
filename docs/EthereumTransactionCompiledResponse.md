# EthereumTransactionCompiledResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **string** | Indicates the chain that the transaction was compiled for (e.g., ropsten or mainnet).  | [optional] 
**Data** | Pointer to **string** | The transaction data  | [optional] 
**From** | Pointer to **string** | The address expected to sign and submit the transaction  | [optional] 
**To** | Pointer to **string** | The contract. This should match your provided value for &#x60;token_blockchain_identifier&#x60;.  | [optional] 
**Gas** | Pointer to **float32** |  | [optional] 
**MaxFeePerGas** | Pointer to **float32** |  | [optional] 
**MaxPriorityFeePerGas** | Pointer to **float32** |  | [optional] 
**Nonce** | Pointer to **float32** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewEthereumTransactionCompiledResponse

`func NewEthereumTransactionCompiledResponse() *EthereumTransactionCompiledResponse`

NewEthereumTransactionCompiledResponse instantiates a new EthereumTransactionCompiledResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthereumTransactionCompiledResponseWithDefaults

`func NewEthereumTransactionCompiledResponseWithDefaults() *EthereumTransactionCompiledResponse`

NewEthereumTransactionCompiledResponseWithDefaults instantiates a new EthereumTransactionCompiledResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *EthereumTransactionCompiledResponse) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *EthereumTransactionCompiledResponse) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *EthereumTransactionCompiledResponse) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *EthereumTransactionCompiledResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetData

`func (o *EthereumTransactionCompiledResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EthereumTransactionCompiledResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EthereumTransactionCompiledResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *EthereumTransactionCompiledResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFrom

`func (o *EthereumTransactionCompiledResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EthereumTransactionCompiledResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EthereumTransactionCompiledResponse) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EthereumTransactionCompiledResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *EthereumTransactionCompiledResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EthereumTransactionCompiledResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EthereumTransactionCompiledResponse) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *EthereumTransactionCompiledResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetGas

`func (o *EthereumTransactionCompiledResponse) GetGas() float32`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *EthereumTransactionCompiledResponse) GetGasOk() (*float32, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *EthereumTransactionCompiledResponse) SetGas(v float32)`

SetGas sets Gas field to given value.

### HasGas

`func (o *EthereumTransactionCompiledResponse) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *EthereumTransactionCompiledResponse) GetMaxFeePerGas() float32`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *EthereumTransactionCompiledResponse) GetMaxFeePerGasOk() (*float32, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *EthereumTransactionCompiledResponse) SetMaxFeePerGas(v float32)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *EthereumTransactionCompiledResponse) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *EthereumTransactionCompiledResponse) GetMaxPriorityFeePerGas() float32`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *EthereumTransactionCompiledResponse) GetMaxPriorityFeePerGasOk() (*float32, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *EthereumTransactionCompiledResponse) SetMaxPriorityFeePerGas(v float32)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *EthereumTransactionCompiledResponse) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### GetNonce

`func (o *EthereumTransactionCompiledResponse) GetNonce() float32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *EthereumTransactionCompiledResponse) GetNonceOk() (*float32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *EthereumTransactionCompiledResponse) SetNonce(v float32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *EthereumTransactionCompiledResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetValue

`func (o *EthereumTransactionCompiledResponse) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EthereumTransactionCompiledResponse) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EthereumTransactionCompiledResponse) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *EthereumTransactionCompiledResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


