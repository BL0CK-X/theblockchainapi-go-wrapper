# CCProjectCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the project. This is shown to your users and should identify your company or organization.  | 
**Description** | Pointer to **string** | The description of your project / company.  | [optional] 
**Webhook** | Pointer to **string** | A URL that identifies where we should make an API request to notify you of a new payment (e.g., api.myproject.com/crypto_payments/webhook). Learn more [here](#tag/CC-Webhook/operation/getCCWebhook).  | [optional] 
**Website** | Pointer to **string** | The website of your project / company.  | [optional] 
**DiscordWebhook** | Pointer to **string** | A Discord webhook. We will send a message to this channel to notify of payment. Learn more [here]().  | [optional] 
**LogoUrl** | Pointer to **string** | A URL of your logo.  | [optional] 
**CustomerIdToCollect** | Pointer to [**CCProjectCreateRequestCustomerIdToCollect**](CCProjectCreateRequestCustomerIdToCollect.md) |  | [optional] 
**SuccessUrl** | Pointer to **string** | Where to redirect customers after payment. If not supplied, customers will be redirected to checkout.blockchainapi.com/me to view their subscriptions.  | [optional] 
**PayoutMethod** | Pointer to [**CCProjectCreateRequestPayoutMethod**](CCProjectCreateRequestPayoutMethod.md) |  | [optional] 

## Methods

### NewCCProjectCreateRequest

`func NewCCProjectCreateRequest(name string, ) *CCProjectCreateRequest`

NewCCProjectCreateRequest instantiates a new CCProjectCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCProjectCreateRequestWithDefaults

`func NewCCProjectCreateRequestWithDefaults() *CCProjectCreateRequest`

NewCCProjectCreateRequestWithDefaults instantiates a new CCProjectCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CCProjectCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CCProjectCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CCProjectCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CCProjectCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CCProjectCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CCProjectCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CCProjectCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWebhook

`func (o *CCProjectCreateRequest) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *CCProjectCreateRequest) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *CCProjectCreateRequest) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *CCProjectCreateRequest) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetWebsite

`func (o *CCProjectCreateRequest) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CCProjectCreateRequest) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CCProjectCreateRequest) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CCProjectCreateRequest) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetDiscordWebhook

`func (o *CCProjectCreateRequest) GetDiscordWebhook() string`

GetDiscordWebhook returns the DiscordWebhook field if non-nil, zero value otherwise.

### GetDiscordWebhookOk

`func (o *CCProjectCreateRequest) GetDiscordWebhookOk() (*string, bool)`

GetDiscordWebhookOk returns a tuple with the DiscordWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordWebhook

`func (o *CCProjectCreateRequest) SetDiscordWebhook(v string)`

SetDiscordWebhook sets DiscordWebhook field to given value.

### HasDiscordWebhook

`func (o *CCProjectCreateRequest) HasDiscordWebhook() bool`

HasDiscordWebhook returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CCProjectCreateRequest) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CCProjectCreateRequest) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CCProjectCreateRequest) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CCProjectCreateRequest) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetCustomerIdToCollect

`func (o *CCProjectCreateRequest) GetCustomerIdToCollect() CCProjectCreateRequestCustomerIdToCollect`

GetCustomerIdToCollect returns the CustomerIdToCollect field if non-nil, zero value otherwise.

### GetCustomerIdToCollectOk

`func (o *CCProjectCreateRequest) GetCustomerIdToCollectOk() (*CCProjectCreateRequestCustomerIdToCollect, bool)`

GetCustomerIdToCollectOk returns a tuple with the CustomerIdToCollect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIdToCollect

`func (o *CCProjectCreateRequest) SetCustomerIdToCollect(v CCProjectCreateRequestCustomerIdToCollect)`

SetCustomerIdToCollect sets CustomerIdToCollect field to given value.

### HasCustomerIdToCollect

`func (o *CCProjectCreateRequest) HasCustomerIdToCollect() bool`

HasCustomerIdToCollect returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *CCProjectCreateRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CCProjectCreateRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CCProjectCreateRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *CCProjectCreateRequest) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetPayoutMethod

`func (o *CCProjectCreateRequest) GetPayoutMethod() CCProjectCreateRequestPayoutMethod`

GetPayoutMethod returns the PayoutMethod field if non-nil, zero value otherwise.

### GetPayoutMethodOk

`func (o *CCProjectCreateRequest) GetPayoutMethodOk() (*CCProjectCreateRequestPayoutMethod, bool)`

GetPayoutMethodOk returns a tuple with the PayoutMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMethod

`func (o *CCProjectCreateRequest) SetPayoutMethod(v CCProjectCreateRequestPayoutMethod)`

SetPayoutMethod sets PayoutMethod field to given value.

### HasPayoutMethod

`func (o *CCProjectCreateRequest) HasPayoutMethod() bool`

HasPayoutMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


