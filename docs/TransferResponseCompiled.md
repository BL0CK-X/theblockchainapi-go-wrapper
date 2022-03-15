# TransferResponseCompiled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**B58CompiledTransaction** | Pointer to **string** | A base58 encoded byte array in string representation. Really easy to submit to Phantom. See &lt;a href&#x3D;\&quot;https://github.com/BL0CK-X/blockchain-api/blob/main/examples/tutorials/phantom_tutorials/transfer_solana.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt; for an example on how to submit it to a Phantom wallet for signing. | [optional] 
**CompiledTransaction** | Pointer to **map[string]interface{}** | An array of integers representing the bytes of the transaction | [optional] 

## Methods

### NewTransferResponseCompiled

`func NewTransferResponseCompiled() *TransferResponseCompiled`

NewTransferResponseCompiled instantiates a new TransferResponseCompiled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResponseCompiledWithDefaults

`func NewTransferResponseCompiledWithDefaults() *TransferResponseCompiled`

NewTransferResponseCompiledWithDefaults instantiates a new TransferResponseCompiled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetB58CompiledTransaction

`func (o *TransferResponseCompiled) GetB58CompiledTransaction() string`

GetB58CompiledTransaction returns the B58CompiledTransaction field if non-nil, zero value otherwise.

### GetB58CompiledTransactionOk

`func (o *TransferResponseCompiled) GetB58CompiledTransactionOk() (*string, bool)`

GetB58CompiledTransactionOk returns a tuple with the B58CompiledTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB58CompiledTransaction

`func (o *TransferResponseCompiled) SetB58CompiledTransaction(v string)`

SetB58CompiledTransaction sets B58CompiledTransaction field to given value.

### HasB58CompiledTransaction

`func (o *TransferResponseCompiled) HasB58CompiledTransaction() bool`

HasB58CompiledTransaction returns a boolean if a field has been set.

### GetCompiledTransaction

`func (o *TransferResponseCompiled) GetCompiledTransaction() map[string]interface{}`

GetCompiledTransaction returns the CompiledTransaction field if non-nil, zero value otherwise.

### GetCompiledTransactionOk

`func (o *TransferResponseCompiled) GetCompiledTransactionOk() (*map[string]interface{}, bool)`

GetCompiledTransactionOk returns a tuple with the CompiledTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiledTransaction

`func (o *TransferResponseCompiled) SetCompiledTransaction(v map[string]interface{})`

SetCompiledTransaction sets CompiledTransaction field to given value.

### HasCompiledTransaction

`func (o *TransferResponseCompiled) HasCompiledTransaction() bool`

HasCompiledTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


