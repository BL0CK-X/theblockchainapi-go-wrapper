# GeneralAirdropRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientBlockchainIdentifier** | **string** | The address/public key to which to send the airdrop.  Examples: - Solana: &#x60;GKNcUmNacSJo4S2Kq3DuYRYRGw3sNUfJ4tyqd198t6vQ&#x60; - Ethereum: &#x60;0xa84b9478d203cd25dF722e83C87590f8028f6aAA&#x60; | 

## Methods

### NewGeneralAirdropRequest

`func NewGeneralAirdropRequest(recipientBlockchainIdentifier string, ) *GeneralAirdropRequest`

NewGeneralAirdropRequest instantiates a new GeneralAirdropRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralAirdropRequestWithDefaults

`func NewGeneralAirdropRequestWithDefaults() *GeneralAirdropRequest`

NewGeneralAirdropRequestWithDefaults instantiates a new GeneralAirdropRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientBlockchainIdentifier

`func (o *GeneralAirdropRequest) GetRecipientBlockchainIdentifier() string`

GetRecipientBlockchainIdentifier returns the RecipientBlockchainIdentifier field if non-nil, zero value otherwise.

### GetRecipientBlockchainIdentifierOk

`func (o *GeneralAirdropRequest) GetRecipientBlockchainIdentifierOk() (*string, bool)`

GetRecipientBlockchainIdentifierOk returns a tuple with the RecipientBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientBlockchainIdentifier

`func (o *GeneralAirdropRequest) SetRecipientBlockchainIdentifier(v string)`

SetRecipientBlockchainIdentifier sets RecipientBlockchainIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


