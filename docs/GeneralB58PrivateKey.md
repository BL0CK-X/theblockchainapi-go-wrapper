# GeneralB58PrivateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**B58PrivateKey** | **string** | A private key corresponds to exactly one public key address. A private key can be used to move assets out of the wallet and sign transaction with the corresponding public key.  A base58-encoded private key is a base58-encoded version of the typical private key. It is represented as a string (e.g., &#x60;4waBTVeAVWEAczSdx36uMrR19668ACgQDs7r386vrUes3UCzvXCQ2FPSCVGb1zJrwcULgpNzgABreyQaWSpGBwfx&#x60;).  &lt;a href&#x3D;\&quot;https://phantom.app\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Phantom&lt;/a&gt; is a popular wallet interface on Solana that allows you to export your private key in this format.  You can use &#x60;b58_private_key&#x60; for Chains: &#x60;solana&#x60;, &#x60;near&#x60; | 

## Methods

### NewGeneralB58PrivateKey

`func NewGeneralB58PrivateKey(b58PrivateKey string, ) *GeneralB58PrivateKey`

NewGeneralB58PrivateKey instantiates a new GeneralB58PrivateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralB58PrivateKeyWithDefaults

`func NewGeneralB58PrivateKeyWithDefaults() *GeneralB58PrivateKey`

NewGeneralB58PrivateKeyWithDefaults instantiates a new GeneralB58PrivateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetB58PrivateKey

`func (o *GeneralB58PrivateKey) GetB58PrivateKey() string`

GetB58PrivateKey returns the B58PrivateKey field if non-nil, zero value otherwise.

### GetB58PrivateKeyOk

`func (o *GeneralB58PrivateKey) GetB58PrivateKeyOk() (*string, bool)`

GetB58PrivateKeyOk returns a tuple with the B58PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB58PrivateKey

`func (o *GeneralB58PrivateKey) SetB58PrivateKey(v string)`

SetB58PrivateKey sets B58PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


