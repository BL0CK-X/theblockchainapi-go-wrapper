# TransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionSignature** | Pointer to **string** | The signature of the transaction | [optional] 
**Confirmed** | Pointer to **bool** | Whether or not the transaction was confirmed or simply submitted for processing. The status depends on your input for &#x60;wait_for_confirmation&#x60;. This only is returned when you are submitting a transaction, not when retrieving signatures for a public key, for example. | [optional] 

## Methods

### NewTransferResponse

`func NewTransferResponse() *TransferResponse`

NewTransferResponse instantiates a new TransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResponseWithDefaults

`func NewTransferResponseWithDefaults() *TransferResponse`

NewTransferResponseWithDefaults instantiates a new TransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionSignature

`func (o *TransferResponse) GetTransactionSignature() string`

GetTransactionSignature returns the TransactionSignature field if non-nil, zero value otherwise.

### GetTransactionSignatureOk

`func (o *TransferResponse) GetTransactionSignatureOk() (*string, bool)`

GetTransactionSignatureOk returns a tuple with the TransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSignature

`func (o *TransferResponse) SetTransactionSignature(v string)`

SetTransactionSignature sets TransactionSignature field to given value.

### HasTransactionSignature

`func (o *TransferResponse) HasTransactionSignature() bool`

HasTransactionSignature returns a boolean if a field has been set.

### GetConfirmed

`func (o *TransferResponse) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *TransferResponse) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *TransferResponse) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *TransferResponse) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


