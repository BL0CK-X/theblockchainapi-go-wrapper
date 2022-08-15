# CCWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | Pointer to **string** | The unique identifier of the webhook  | [optional] 
**Signature** | Pointer to **map[string]interface{}** | A signature from the Blockchain API that verifies the webhook is from us. An array of integers.  | [optional] 
**TimeSent** | Pointer to **float32** | The time we sent the webhook  | [optional] 
**Data** | Pointer to [**CCPayment**](CCPayment.md) |  | [optional] 

## Methods

### NewCCWebhook

`func NewCCWebhook() *CCWebhook`

NewCCWebhook instantiates a new CCWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCWebhookWithDefaults

`func NewCCWebhookWithDefaults() *CCWebhook`

NewCCWebhookWithDefaults instantiates a new CCWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *CCWebhook) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *CCWebhook) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *CCWebhook) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *CCWebhook) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetSignature

`func (o *CCWebhook) GetSignature() map[string]interface{}`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CCWebhook) GetSignatureOk() (*map[string]interface{}, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CCWebhook) SetSignature(v map[string]interface{})`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CCWebhook) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTimeSent

`func (o *CCWebhook) GetTimeSent() float32`

GetTimeSent returns the TimeSent field if non-nil, zero value otherwise.

### GetTimeSentOk

`func (o *CCWebhook) GetTimeSentOk() (*float32, bool)`

GetTimeSentOk returns a tuple with the TimeSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSent

`func (o *CCWebhook) SetTimeSent(v float32)`

SetTimeSent sets TimeSent field to given value.

### HasTimeSent

`func (o *CCWebhook) HasTimeSent() bool`

HasTimeSent returns a boolean if a field has been set.

### GetData

`func (o *CCWebhook) GetData() CCPayment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CCWebhook) GetDataOk() (*CCPayment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CCWebhook) SetData(v CCPayment)`

SetData sets Data field to given value.

### HasData

`func (o *CCWebhook) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


