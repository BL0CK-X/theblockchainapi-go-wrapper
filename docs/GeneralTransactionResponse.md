# GeneralTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionBlockchainIdentifier** | Pointer to **string** | The identifier of the transaction | [optional] 
**TransactionConfirmed** | Pointer to **bool** | Whether or not the transaction was confirmed | [optional] 
**B58CompiledTransaction** | Pointer to **string** | A base58 encoded byte array in string representation. Really easy to submit to Phantom. See &lt;a href&#x3D;\&quot;https://github.com/BL0CK-X/blockchain-api/blob/main/examples/tutorials/phantom_tutorials/transfer_solana.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt; for an example on how to submit it to a Phantom wallet for signing. | [optional] 
**CompiledTransaction** | Pointer to **map[string]interface{}** | An array of integers representing the bytes of the transaction | [optional] 
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

### NewGeneralTransactionResponse

`func NewGeneralTransactionResponse() *GeneralTransactionResponse`

NewGeneralTransactionResponse instantiates a new GeneralTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralTransactionResponseWithDefaults

`func NewGeneralTransactionResponseWithDefaults() *GeneralTransactionResponse`

NewGeneralTransactionResponseWithDefaults instantiates a new GeneralTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionBlockchainIdentifier

`func (o *GeneralTransactionResponse) GetTransactionBlockchainIdentifier() string`

GetTransactionBlockchainIdentifier returns the TransactionBlockchainIdentifier field if non-nil, zero value otherwise.

### GetTransactionBlockchainIdentifierOk

`func (o *GeneralTransactionResponse) GetTransactionBlockchainIdentifierOk() (*string, bool)`

GetTransactionBlockchainIdentifierOk returns a tuple with the TransactionBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionBlockchainIdentifier

`func (o *GeneralTransactionResponse) SetTransactionBlockchainIdentifier(v string)`

SetTransactionBlockchainIdentifier sets TransactionBlockchainIdentifier field to given value.

### HasTransactionBlockchainIdentifier

`func (o *GeneralTransactionResponse) HasTransactionBlockchainIdentifier() bool`

HasTransactionBlockchainIdentifier returns a boolean if a field has been set.

### GetTransactionConfirmed

`func (o *GeneralTransactionResponse) GetTransactionConfirmed() bool`

GetTransactionConfirmed returns the TransactionConfirmed field if non-nil, zero value otherwise.

### GetTransactionConfirmedOk

`func (o *GeneralTransactionResponse) GetTransactionConfirmedOk() (*bool, bool)`

GetTransactionConfirmedOk returns a tuple with the TransactionConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionConfirmed

`func (o *GeneralTransactionResponse) SetTransactionConfirmed(v bool)`

SetTransactionConfirmed sets TransactionConfirmed field to given value.

### HasTransactionConfirmed

`func (o *GeneralTransactionResponse) HasTransactionConfirmed() bool`

HasTransactionConfirmed returns a boolean if a field has been set.

### GetB58CompiledTransaction

`func (o *GeneralTransactionResponse) GetB58CompiledTransaction() string`

GetB58CompiledTransaction returns the B58CompiledTransaction field if non-nil, zero value otherwise.

### GetB58CompiledTransactionOk

`func (o *GeneralTransactionResponse) GetB58CompiledTransactionOk() (*string, bool)`

GetB58CompiledTransactionOk returns a tuple with the B58CompiledTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB58CompiledTransaction

`func (o *GeneralTransactionResponse) SetB58CompiledTransaction(v string)`

SetB58CompiledTransaction sets B58CompiledTransaction field to given value.

### HasB58CompiledTransaction

`func (o *GeneralTransactionResponse) HasB58CompiledTransaction() bool`

HasB58CompiledTransaction returns a boolean if a field has been set.

### GetCompiledTransaction

`func (o *GeneralTransactionResponse) GetCompiledTransaction() map[string]interface{}`

GetCompiledTransaction returns the CompiledTransaction field if non-nil, zero value otherwise.

### GetCompiledTransactionOk

`func (o *GeneralTransactionResponse) GetCompiledTransactionOk() (*map[string]interface{}, bool)`

GetCompiledTransactionOk returns a tuple with the CompiledTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiledTransaction

`func (o *GeneralTransactionResponse) SetCompiledTransaction(v map[string]interface{})`

SetCompiledTransaction sets CompiledTransaction field to given value.

### HasCompiledTransaction

`func (o *GeneralTransactionResponse) HasCompiledTransaction() bool`

HasCompiledTransaction returns a boolean if a field has been set.

### GetChainId

`func (o *GeneralTransactionResponse) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *GeneralTransactionResponse) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *GeneralTransactionResponse) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *GeneralTransactionResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetData

`func (o *GeneralTransactionResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GeneralTransactionResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GeneralTransactionResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *GeneralTransactionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFrom

`func (o *GeneralTransactionResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GeneralTransactionResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GeneralTransactionResponse) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GeneralTransactionResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GeneralTransactionResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GeneralTransactionResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GeneralTransactionResponse) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GeneralTransactionResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetGas

`func (o *GeneralTransactionResponse) GetGas() float32`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *GeneralTransactionResponse) GetGasOk() (*float32, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *GeneralTransactionResponse) SetGas(v float32)`

SetGas sets Gas field to given value.

### HasGas

`func (o *GeneralTransactionResponse) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *GeneralTransactionResponse) GetMaxFeePerGas() float32`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *GeneralTransactionResponse) GetMaxFeePerGasOk() (*float32, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *GeneralTransactionResponse) SetMaxFeePerGas(v float32)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *GeneralTransactionResponse) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *GeneralTransactionResponse) GetMaxPriorityFeePerGas() float32`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *GeneralTransactionResponse) GetMaxPriorityFeePerGasOk() (*float32, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *GeneralTransactionResponse) SetMaxPriorityFeePerGas(v float32)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *GeneralTransactionResponse) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### GetNonce

`func (o *GeneralTransactionResponse) GetNonce() float32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GeneralTransactionResponse) GetNonceOk() (*float32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GeneralTransactionResponse) SetNonce(v float32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *GeneralTransactionResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetValue

`func (o *GeneralTransactionResponse) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GeneralTransactionResponse) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GeneralTransactionResponse) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *GeneralTransactionResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


