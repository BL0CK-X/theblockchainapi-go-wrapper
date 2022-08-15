# CCPaymentBlockchainPaymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** | The chain on which the payment was made.  | [optional] 
**FeeAmount** | Pointer to **string** | The fee taken by the blockchain API. This is shown in integer form.  | [optional] 
**PayoutAmount** | Pointer to **string** | The amount paid to the project&#39;s payout wallet (the recipient of the payment). This is shown in integer form.  | [optional] 
**TokenBlockchainIdentifier** | Pointer to **string** | The blockchain identifier of the token used to pay. Currently, we only support payments in USDC.  | [optional] 
**TotalPaid** | Pointer to **string** | The total paid by the user when making the payment. This should match the amount set in the respective plan.  | [optional] 
**TransactionSignature** | Pointer to **string** | The transaction signature which you can use to manually or automatically verify the transaction on the blockchain.  | [optional] 

## Methods

### NewCCPaymentBlockchainPaymentDetails

`func NewCCPaymentBlockchainPaymentDetails() *CCPaymentBlockchainPaymentDetails`

NewCCPaymentBlockchainPaymentDetails instantiates a new CCPaymentBlockchainPaymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCPaymentBlockchainPaymentDetailsWithDefaults

`func NewCCPaymentBlockchainPaymentDetailsWithDefaults() *CCPaymentBlockchainPaymentDetails`

NewCCPaymentBlockchainPaymentDetailsWithDefaults instantiates a new CCPaymentBlockchainPaymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *CCPaymentBlockchainPaymentDetails) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CCPaymentBlockchainPaymentDetails) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CCPaymentBlockchainPaymentDetails) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *CCPaymentBlockchainPaymentDetails) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetFeeAmount

`func (o *CCPaymentBlockchainPaymentDetails) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *CCPaymentBlockchainPaymentDetails) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *CCPaymentBlockchainPaymentDetails) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *CCPaymentBlockchainPaymentDetails) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetPayoutAmount

`func (o *CCPaymentBlockchainPaymentDetails) GetPayoutAmount() string`

GetPayoutAmount returns the PayoutAmount field if non-nil, zero value otherwise.

### GetPayoutAmountOk

`func (o *CCPaymentBlockchainPaymentDetails) GetPayoutAmountOk() (*string, bool)`

GetPayoutAmountOk returns a tuple with the PayoutAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutAmount

`func (o *CCPaymentBlockchainPaymentDetails) SetPayoutAmount(v string)`

SetPayoutAmount sets PayoutAmount field to given value.

### HasPayoutAmount

`func (o *CCPaymentBlockchainPaymentDetails) HasPayoutAmount() bool`

HasPayoutAmount returns a boolean if a field has been set.

### GetTokenBlockchainIdentifier

`func (o *CCPaymentBlockchainPaymentDetails) GetTokenBlockchainIdentifier() string`

GetTokenBlockchainIdentifier returns the TokenBlockchainIdentifier field if non-nil, zero value otherwise.

### GetTokenBlockchainIdentifierOk

`func (o *CCPaymentBlockchainPaymentDetails) GetTokenBlockchainIdentifierOk() (*string, bool)`

GetTokenBlockchainIdentifierOk returns a tuple with the TokenBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBlockchainIdentifier

`func (o *CCPaymentBlockchainPaymentDetails) SetTokenBlockchainIdentifier(v string)`

SetTokenBlockchainIdentifier sets TokenBlockchainIdentifier field to given value.

### HasTokenBlockchainIdentifier

`func (o *CCPaymentBlockchainPaymentDetails) HasTokenBlockchainIdentifier() bool`

HasTokenBlockchainIdentifier returns a boolean if a field has been set.

### GetTotalPaid

`func (o *CCPaymentBlockchainPaymentDetails) GetTotalPaid() string`

GetTotalPaid returns the TotalPaid field if non-nil, zero value otherwise.

### GetTotalPaidOk

`func (o *CCPaymentBlockchainPaymentDetails) GetTotalPaidOk() (*string, bool)`

GetTotalPaidOk returns a tuple with the TotalPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaid

`func (o *CCPaymentBlockchainPaymentDetails) SetTotalPaid(v string)`

SetTotalPaid sets TotalPaid field to given value.

### HasTotalPaid

`func (o *CCPaymentBlockchainPaymentDetails) HasTotalPaid() bool`

HasTotalPaid returns a boolean if a field has been set.

### GetTransactionSignature

`func (o *CCPaymentBlockchainPaymentDetails) GetTransactionSignature() string`

GetTransactionSignature returns the TransactionSignature field if non-nil, zero value otherwise.

### GetTransactionSignatureOk

`func (o *CCPaymentBlockchainPaymentDetails) GetTransactionSignatureOk() (*string, bool)`

GetTransactionSignatureOk returns a tuple with the TransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSignature

`func (o *CCPaymentBlockchainPaymentDetails) SetTransactionSignature(v string)`

SetTransactionSignature sets TransactionSignature field to given value.

### HasTransactionSignature

`func (o *CCPaymentBlockchainPaymentDetails) HasTransactionSignature() bool`

HasTransactionSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


