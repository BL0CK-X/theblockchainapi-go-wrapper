# EthereumTransactionMadeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionBlockchainIdentifier** | Pointer to **string** | The identifier of the transaction | [optional] 
**TransactionConfirmed** | Pointer to **bool** | Whether or not the transaction was confirmed | [optional] 

## Methods

### NewEthereumTransactionMadeResponse

`func NewEthereumTransactionMadeResponse() *EthereumTransactionMadeResponse`

NewEthereumTransactionMadeResponse instantiates a new EthereumTransactionMadeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthereumTransactionMadeResponseWithDefaults

`func NewEthereumTransactionMadeResponseWithDefaults() *EthereumTransactionMadeResponse`

NewEthereumTransactionMadeResponseWithDefaults instantiates a new EthereumTransactionMadeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionBlockchainIdentifier

`func (o *EthereumTransactionMadeResponse) GetTransactionBlockchainIdentifier() string`

GetTransactionBlockchainIdentifier returns the TransactionBlockchainIdentifier field if non-nil, zero value otherwise.

### GetTransactionBlockchainIdentifierOk

`func (o *EthereumTransactionMadeResponse) GetTransactionBlockchainIdentifierOk() (*string, bool)`

GetTransactionBlockchainIdentifierOk returns a tuple with the TransactionBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionBlockchainIdentifier

`func (o *EthereumTransactionMadeResponse) SetTransactionBlockchainIdentifier(v string)`

SetTransactionBlockchainIdentifier sets TransactionBlockchainIdentifier field to given value.

### HasTransactionBlockchainIdentifier

`func (o *EthereumTransactionMadeResponse) HasTransactionBlockchainIdentifier() bool`

HasTransactionBlockchainIdentifier returns a boolean if a field has been set.

### GetTransactionConfirmed

`func (o *EthereumTransactionMadeResponse) GetTransactionConfirmed() bool`

GetTransactionConfirmed returns the TransactionConfirmed field if non-nil, zero value otherwise.

### GetTransactionConfirmedOk

`func (o *EthereumTransactionMadeResponse) GetTransactionConfirmedOk() (*bool, bool)`

GetTransactionConfirmedOk returns a tuple with the TransactionConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionConfirmed

`func (o *EthereumTransactionMadeResponse) SetTransactionConfirmed(v bool)`

SetTransactionConfirmed sets TransactionConfirmed field to given value.

### HasTransactionConfirmed

`func (o *EthereumTransactionMadeResponse) HasTransactionConfirmed() bool`

HasTransactionConfirmed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


