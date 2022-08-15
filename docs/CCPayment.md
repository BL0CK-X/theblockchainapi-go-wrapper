# CCPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** | The ID of the project  | [optional] 
**ProductId** | Pointer to **string** | The ID of the associated product  | [optional] 
**PlanId** | Pointer to **string** | The ID of the respective plan  | [optional] 
**PaymentId** | Pointer to **string** | The unique ID of the payment  | [optional] 
**BlockchainIdentifier** | Pointer to **string** | The unique identifier of the wallet from which the payment was made.  | [optional] 
**BlockchainPaymentDetails** | Pointer to [**CCPaymentBlockchainPaymentDetails**](CCPaymentBlockchainPaymentDetails.md) |  | [optional] 
**CustomerId** | Pointer to **string** | The ID of the customer  | [optional] 
**PaymentValidationCode** | Pointer to **string** | The validation code shown to the customer. This is only visible to the customer who paid. They can use this code to redeem their subscription to their product.  | [optional] 
**PeriodEnd** | Pointer to **float32** | A UNIX time stamp, in seconds, that identifies the end of the period of the subscription  | [optional] 
**PeriodStart** | Pointer to **float32** | A UNIX time stamp, in seconds, that identifies the start of the period of the subscription  | [optional] 
**TransactionBlockchainIdentifier** | Pointer to **string** | The string that uniquely identifies the blockchain transaction  | [optional] 

## Methods

### NewCCPayment

`func NewCCPayment() *CCPayment`

NewCCPayment instantiates a new CCPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCPaymentWithDefaults

`func NewCCPaymentWithDefaults() *CCPayment`

NewCCPaymentWithDefaults instantiates a new CCPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CCPayment) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CCPayment) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CCPayment) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CCPayment) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProductId

`func (o *CCPayment) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CCPayment) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CCPayment) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CCPayment) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetPlanId

`func (o *CCPayment) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CCPayment) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CCPayment) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CCPayment) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPaymentId

`func (o *CCPayment) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *CCPayment) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *CCPayment) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *CCPayment) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetBlockchainIdentifier

`func (o *CCPayment) GetBlockchainIdentifier() string`

GetBlockchainIdentifier returns the BlockchainIdentifier field if non-nil, zero value otherwise.

### GetBlockchainIdentifierOk

`func (o *CCPayment) GetBlockchainIdentifierOk() (*string, bool)`

GetBlockchainIdentifierOk returns a tuple with the BlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainIdentifier

`func (o *CCPayment) SetBlockchainIdentifier(v string)`

SetBlockchainIdentifier sets BlockchainIdentifier field to given value.

### HasBlockchainIdentifier

`func (o *CCPayment) HasBlockchainIdentifier() bool`

HasBlockchainIdentifier returns a boolean if a field has been set.

### GetBlockchainPaymentDetails

`func (o *CCPayment) GetBlockchainPaymentDetails() CCPaymentBlockchainPaymentDetails`

GetBlockchainPaymentDetails returns the BlockchainPaymentDetails field if non-nil, zero value otherwise.

### GetBlockchainPaymentDetailsOk

`func (o *CCPayment) GetBlockchainPaymentDetailsOk() (*CCPaymentBlockchainPaymentDetails, bool)`

GetBlockchainPaymentDetailsOk returns a tuple with the BlockchainPaymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainPaymentDetails

`func (o *CCPayment) SetBlockchainPaymentDetails(v CCPaymentBlockchainPaymentDetails)`

SetBlockchainPaymentDetails sets BlockchainPaymentDetails field to given value.

### HasBlockchainPaymentDetails

`func (o *CCPayment) HasBlockchainPaymentDetails() bool`

HasBlockchainPaymentDetails returns a boolean if a field has been set.

### GetCustomerId

`func (o *CCPayment) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CCPayment) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CCPayment) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CCPayment) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetPaymentValidationCode

`func (o *CCPayment) GetPaymentValidationCode() string`

GetPaymentValidationCode returns the PaymentValidationCode field if non-nil, zero value otherwise.

### GetPaymentValidationCodeOk

`func (o *CCPayment) GetPaymentValidationCodeOk() (*string, bool)`

GetPaymentValidationCodeOk returns a tuple with the PaymentValidationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentValidationCode

`func (o *CCPayment) SetPaymentValidationCode(v string)`

SetPaymentValidationCode sets PaymentValidationCode field to given value.

### HasPaymentValidationCode

`func (o *CCPayment) HasPaymentValidationCode() bool`

HasPaymentValidationCode returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *CCPayment) GetPeriodEnd() float32`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *CCPayment) GetPeriodEndOk() (*float32, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *CCPayment) SetPeriodEnd(v float32)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *CCPayment) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *CCPayment) GetPeriodStart() float32`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *CCPayment) GetPeriodStartOk() (*float32, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *CCPayment) SetPeriodStart(v float32)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *CCPayment) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetTransactionBlockchainIdentifier

`func (o *CCPayment) GetTransactionBlockchainIdentifier() string`

GetTransactionBlockchainIdentifier returns the TransactionBlockchainIdentifier field if non-nil, zero value otherwise.

### GetTransactionBlockchainIdentifierOk

`func (o *CCPayment) GetTransactionBlockchainIdentifierOk() (*string, bool)`

GetTransactionBlockchainIdentifierOk returns a tuple with the TransactionBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionBlockchainIdentifier

`func (o *CCPayment) SetTransactionBlockchainIdentifier(v string)`

SetTransactionBlockchainIdentifier sets TransactionBlockchainIdentifier field to given value.

### HasTransactionBlockchainIdentifier

`func (o *CCPayment) HasTransactionBlockchainIdentifier() bool`

HasTransactionBlockchainIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


