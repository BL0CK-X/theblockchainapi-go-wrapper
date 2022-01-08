# GeneratePrivateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | Pointer to **map[string]interface{}** | A private key corresponds to exactly one public key address. A private key can be used to move assets out of the wallet and sign transaction with the corresponding public key.  A private key is an array of integers (e.g.,   &#x60;[185,108,153,165,57,193,166,167,58,148,133,121,92,252,242,13,233,246,35,103,185,20,27,56,111,169,12,50,50,36,83,156,173,195,143,75,135,78,204,129,217,231,58,129,69,180,185,86,119,43,200,193,94,112,31,135,68,128,207,26,85,150,68,181]&#x60;).  &lt;a href&#x3D;\&quot;https://solflare.com\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Solflare&lt;/a&gt; is a popular wallet interface on Solana that allows you to export your private key in this format. | [optional] 
**B58PrivateKey** | Pointer to **string** | A private key corresponds to exactly one public key address. A private key can be used to move assets out of the wallet and sign transaction with the corresponding public key.  A base58-encoded private key is a base58-encoded version of the typical private key. It is represented as a string (e.g., &#x60;4waBTVeAVWEAczSdx36uMrR19668ACgQDs7r386vrUes3UCzvXCQ2FPSCVGb1zJrwcULgpNzgABreyQaWSpGBwfx&#x60;).  &lt;a href&#x3D;\&quot;https://phantom.app\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Phantom&lt;/a&gt; is a popular wallet interface on Solana that allows you to export your private key in this format. | [optional] 

## Methods

### NewGeneratePrivateKey

`func NewGeneratePrivateKey() *GeneratePrivateKey`

NewGeneratePrivateKey instantiates a new GeneratePrivateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratePrivateKeyWithDefaults

`func NewGeneratePrivateKeyWithDefaults() *GeneratePrivateKey`

NewGeneratePrivateKeyWithDefaults instantiates a new GeneratePrivateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *GeneratePrivateKey) GetPrivateKey() map[string]interface{}`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GeneratePrivateKey) GetPrivateKeyOk() (*map[string]interface{}, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GeneratePrivateKey) SetPrivateKey(v map[string]interface{})`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GeneratePrivateKey) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetB58PrivateKey

`func (o *GeneratePrivateKey) GetB58PrivateKey() string`

GetB58PrivateKey returns the B58PrivateKey field if non-nil, zero value otherwise.

### GetB58PrivateKeyOk

`func (o *GeneratePrivateKey) GetB58PrivateKeyOk() (*string, bool)`

GetB58PrivateKeyOk returns a tuple with the B58PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB58PrivateKey

`func (o *GeneratePrivateKey) SetB58PrivateKey(v string)`

SetB58PrivateKey sets B58PrivateKey field to given value.

### HasB58PrivateKey

`func (o *GeneratePrivateKey) HasB58PrivateKey() bool`

HasB58PrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


