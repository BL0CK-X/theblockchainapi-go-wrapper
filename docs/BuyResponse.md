# BuyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionSignature** | Pointer to **string** | The signature of the &#x60;buy&#x60; transaction  | [optional] 

## Methods

### NewBuyResponse

`func NewBuyResponse() *BuyResponse`

NewBuyResponse instantiates a new BuyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyResponseWithDefaults

`func NewBuyResponseWithDefaults() *BuyResponse`

NewBuyResponseWithDefaults instantiates a new BuyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionSignature

`func (o *BuyResponse) GetTransactionSignature() string`

GetTransactionSignature returns the TransactionSignature field if non-nil, zero value otherwise.

### GetTransactionSignatureOk

`func (o *BuyResponse) GetTransactionSignatureOk() (*string, bool)`

GetTransactionSignatureOk returns a tuple with the TransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSignature

`func (o *BuyResponse) SetTransactionSignature(v string)`

SetTransactionSignature sets TransactionSignature field to given value.

### HasTransactionSignature

`func (o *BuyResponse) HasTransactionSignature() bool`

HasTransactionSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


