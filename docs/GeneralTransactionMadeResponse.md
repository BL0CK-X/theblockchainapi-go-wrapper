# GeneralTransactionMadeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionBlockchainIdentifier** | Pointer to **string** | The identifier of the transaction | [optional] 
**TransactionConfirmed** | Pointer to **bool** | Whether or not the transaction was confirmed | [optional] 

## Methods

### NewGeneralTransactionMadeResponse

`func NewGeneralTransactionMadeResponse() *GeneralTransactionMadeResponse`

NewGeneralTransactionMadeResponse instantiates a new GeneralTransactionMadeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralTransactionMadeResponseWithDefaults

`func NewGeneralTransactionMadeResponseWithDefaults() *GeneralTransactionMadeResponse`

NewGeneralTransactionMadeResponseWithDefaults instantiates a new GeneralTransactionMadeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionBlockchainIdentifier

`func (o *GeneralTransactionMadeResponse) GetTransactionBlockchainIdentifier() string`

GetTransactionBlockchainIdentifier returns the TransactionBlockchainIdentifier field if non-nil, zero value otherwise.

### GetTransactionBlockchainIdentifierOk

`func (o *GeneralTransactionMadeResponse) GetTransactionBlockchainIdentifierOk() (*string, bool)`

GetTransactionBlockchainIdentifierOk returns a tuple with the TransactionBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionBlockchainIdentifier

`func (o *GeneralTransactionMadeResponse) SetTransactionBlockchainIdentifier(v string)`

SetTransactionBlockchainIdentifier sets TransactionBlockchainIdentifier field to given value.

### HasTransactionBlockchainIdentifier

`func (o *GeneralTransactionMadeResponse) HasTransactionBlockchainIdentifier() bool`

HasTransactionBlockchainIdentifier returns a boolean if a field has been set.

### GetTransactionConfirmed

`func (o *GeneralTransactionMadeResponse) GetTransactionConfirmed() bool`

GetTransactionConfirmed returns the TransactionConfirmed field if non-nil, zero value otherwise.

### GetTransactionConfirmedOk

`func (o *GeneralTransactionMadeResponse) GetTransactionConfirmedOk() (*bool, bool)`

GetTransactionConfirmedOk returns a tuple with the TransactionConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionConfirmed

`func (o *GeneralTransactionMadeResponse) SetTransactionConfirmed(v bool)`

SetTransactionConfirmed sets TransactionConfirmed field to given value.

### HasTransactionConfirmed

`func (o *GeneralTransactionMadeResponse) HasTransactionConfirmed() bool`

HasTransactionConfirmed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


