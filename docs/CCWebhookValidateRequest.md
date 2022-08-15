# CCWebhookValidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | **string** | The unique identifier of the webhook.  | 
**WebhookSignature** | **map[string]interface{}** | The signature we inluded in the webhook. The signature is from the Blockchain API and it ensures that the webhook is from us and not a malicioius actor.  | 
**TimeSent** | **float32** | The time the webhook was sent.  | 

## Methods

### NewCCWebhookValidateRequest

`func NewCCWebhookValidateRequest(webhookId string, webhookSignature map[string]interface{}, timeSent float32, ) *CCWebhookValidateRequest`

NewCCWebhookValidateRequest instantiates a new CCWebhookValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCWebhookValidateRequestWithDefaults

`func NewCCWebhookValidateRequestWithDefaults() *CCWebhookValidateRequest`

NewCCWebhookValidateRequestWithDefaults instantiates a new CCWebhookValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *CCWebhookValidateRequest) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *CCWebhookValidateRequest) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *CCWebhookValidateRequest) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetWebhookSignature

`func (o *CCWebhookValidateRequest) GetWebhookSignature() map[string]interface{}`

GetWebhookSignature returns the WebhookSignature field if non-nil, zero value otherwise.

### GetWebhookSignatureOk

`func (o *CCWebhookValidateRequest) GetWebhookSignatureOk() (*map[string]interface{}, bool)`

GetWebhookSignatureOk returns a tuple with the WebhookSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSignature

`func (o *CCWebhookValidateRequest) SetWebhookSignature(v map[string]interface{})`

SetWebhookSignature sets WebhookSignature field to given value.


### GetTimeSent

`func (o *CCWebhookValidateRequest) GetTimeSent() float32`

GetTimeSent returns the TimeSent field if non-nil, zero value otherwise.

### GetTimeSentOk

`func (o *CCWebhookValidateRequest) GetTimeSentOk() (*float32, bool)`

GetTimeSentOk returns a tuple with the TimeSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSent

`func (o *CCWebhookValidateRequest) SetTimeSent(v float32)`

SetTimeSent sets TimeSent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


