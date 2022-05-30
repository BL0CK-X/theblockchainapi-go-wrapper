# GeneralGeneratePrivateKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HexPrivateKey** | **string** | A hex private key is the typical private key used on Ethereum, for example. It looks like this: &#x60;0x200b9e5baa38b0dc7551645be11b394e9bf2b04532e4af8824bed2b3de2e0dc0&#x60;.  You can use &#x60;hex_private_key&#x60; for chains: &#x60;avalanche&#x60;, &#x60;binance_smart_chain&#x60;, &#x60;ethereum&#x60;. | 
**PrivateKey** | **map[string]interface{}** | A private key corresponds to exactly one public key address. A private key can be used to move assets out of the wallet and sign transaction with the corresponding public key.  A private key is an array of integers (e.g.,   &#x60;[185,108,153,165,57,193,166,167,58,148,133,121,92,252,242,13,233,246,35,103,185,20,27,56,111,169,12,50,50,36,83,156,173,195,143,75,135,78,204,129,217,231,58,129,69,180,185,86,119,43,200,193,94,112,31,135,68,128,207,26,85,150,68,181]&#x60;).  &lt;a href&#x3D;\&quot;https://solflare.com\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Solflare&lt;/a&gt; is a popular wallet interface on Solana that allows you to export your private key in this format.  You can use &#x60;private_key&#x60; for Chains: &#x60;avalanche&#x60;, &#x60;binance_smart_chain&#x60;, &#x60;ethereum&#x60;, &#x60;solana&#x60;, &#x60;near&#x60;. | 
**B58PrivateKey** | **string** | A private key corresponds to exactly one public key address. A private key can be used to move assets out of the wallet and sign transaction with the corresponding public key.  A base58-encoded private key is a base58-encoded version of the typical private key. It is represented as a string (e.g., &#x60;4waBTVeAVWEAczSdx36uMrR19668ACgQDs7r386vrUes3UCzvXCQ2FPSCVGb1zJrwcULgpNzgABreyQaWSpGBwfx&#x60;).  &lt;a href&#x3D;\&quot;https://phantom.app\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Phantom&lt;/a&gt; is a popular wallet interface on Solana that allows you to export your private key in this format.  You can use &#x60;b58_private_key&#x60; for Chains: &#x60;solana&#x60;, &#x60;near&#x60; | 

## Methods

### NewGeneralGeneratePrivateKeyResponse

`func NewGeneralGeneratePrivateKeyResponse(hexPrivateKey string, privateKey map[string]interface{}, b58PrivateKey string, ) *GeneralGeneratePrivateKeyResponse`

NewGeneralGeneratePrivateKeyResponse instantiates a new GeneralGeneratePrivateKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralGeneratePrivateKeyResponseWithDefaults

`func NewGeneralGeneratePrivateKeyResponseWithDefaults() *GeneralGeneratePrivateKeyResponse`

NewGeneralGeneratePrivateKeyResponseWithDefaults instantiates a new GeneralGeneratePrivateKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHexPrivateKey

`func (o *GeneralGeneratePrivateKeyResponse) GetHexPrivateKey() string`

GetHexPrivateKey returns the HexPrivateKey field if non-nil, zero value otherwise.

### GetHexPrivateKeyOk

`func (o *GeneralGeneratePrivateKeyResponse) GetHexPrivateKeyOk() (*string, bool)`

GetHexPrivateKeyOk returns a tuple with the HexPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPrivateKey

`func (o *GeneralGeneratePrivateKeyResponse) SetHexPrivateKey(v string)`

SetHexPrivateKey sets HexPrivateKey field to given value.


### GetPrivateKey

`func (o *GeneralGeneratePrivateKeyResponse) GetPrivateKey() map[string]interface{}`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GeneralGeneratePrivateKeyResponse) GetPrivateKeyOk() (*map[string]interface{}, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GeneralGeneratePrivateKeyResponse) SetPrivateKey(v map[string]interface{})`

SetPrivateKey sets PrivateKey field to given value.


### GetB58PrivateKey

`func (o *GeneralGeneratePrivateKeyResponse) GetB58PrivateKey() string`

GetB58PrivateKey returns the B58PrivateKey field if non-nil, zero value otherwise.

### GetB58PrivateKeyOk

`func (o *GeneralGeneratePrivateKeyResponse) GetB58PrivateKeyOk() (*string, bool)`

GetB58PrivateKeyOk returns a tuple with the B58PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB58PrivateKey

`func (o *GeneralGeneratePrivateKeyResponse) SetB58PrivateKey(v string)`

SetB58PrivateKey sets B58PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


