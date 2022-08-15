# DoubleTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionSignature** | Pointer to **string** | The signature of the transaction | [optional] 
**Confirmed** | Pointer to **bool** | Whether or not the transaction was confirmed or simply submitted for processing. The status depends on your input for &#x60;wait_for_confirmation&#x60;. This only is returned when you are submitting a transaction, not when retrieving signatures for a public key, for example. | [optional] 
**B58CompiledTransaction** | Pointer to **string** | A base58 encoded byte array in string representation. Really easy to submit to Phantom. See &lt;a href&#x3D;\&quot;https://github.com/BL0CK-X/blockchain-api/blob/main/examples/tutorials/phantom_tutorials/transfer_solana.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt; for an example on how to submit it to a Phantom wallet for signing. | [optional] 
**CompiledTransaction** | Pointer to **map[string]interface{}** | An array of integers representing the bytes of the transaction | [optional] 

## Methods

### NewDoubleTransferResponse

`func NewDoubleTransferResponse() *DoubleTransferResponse`

NewDoubleTransferResponse instantiates a new DoubleTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoubleTransferResponseWithDefaults

`func NewDoubleTransferResponseWithDefaults() *DoubleTransferResponse`

NewDoubleTransferResponseWithDefaults instantiates a new DoubleTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionSignature

`func (o *DoubleTransferResponse) GetTransactionSignature() string`

GetTransactionSignature returns the TransactionSignature field if non-nil, zero value otherwise.

### GetTransactionSignatureOk

`func (o *DoubleTransferResponse) GetTransactionSignatureOk() (*string, bool)`

GetTransactionSignatureOk returns a tuple with the TransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSignature

`func (o *DoubleTransferResponse) SetTransactionSignature(v string)`

SetTransactionSignature sets TransactionSignature field to given value.

### HasTransactionSignature

`func (o *DoubleTransferResponse) HasTransactionSignature() bool`

HasTransactionSignature returns a boolean if a field has been set.

### GetConfirmed

`func (o *DoubleTransferResponse) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *DoubleTransferResponse) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *DoubleTransferResponse) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *DoubleTransferResponse) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetB58CompiledTransaction

`func (o *DoubleTransferResponse) GetB58CompiledTransaction() string`

GetB58CompiledTransaction returns the B58CompiledTransaction field if non-nil, zero value otherwise.

### GetB58CompiledTransactionOk

`func (o *DoubleTransferResponse) GetB58CompiledTransactionOk() (*string, bool)`

GetB58CompiledTransactionOk returns a tuple with the B58CompiledTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB58CompiledTransaction

`func (o *DoubleTransferResponse) SetB58CompiledTransaction(v string)`

SetB58CompiledTransaction sets B58CompiledTransaction field to given value.

### HasB58CompiledTransaction

`func (o *DoubleTransferResponse) HasB58CompiledTransaction() bool`

HasB58CompiledTransaction returns a boolean if a field has been set.

### GetCompiledTransaction

`func (o *DoubleTransferResponse) GetCompiledTransaction() map[string]interface{}`

GetCompiledTransaction returns the CompiledTransaction field if non-nil, zero value otherwise.

### GetCompiledTransactionOk

`func (o *DoubleTransferResponse) GetCompiledTransactionOk() (*map[string]interface{}, bool)`

GetCompiledTransactionOk returns a tuple with the CompiledTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiledTransaction

`func (o *DoubleTransferResponse) SetCompiledTransaction(v map[string]interface{})`

SetCompiledTransaction sets CompiledTransaction field to given value.

### HasCompiledTransaction

`func (o *DoubleTransferResponse) HasCompiledTransaction() bool`

HasCompiledTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


