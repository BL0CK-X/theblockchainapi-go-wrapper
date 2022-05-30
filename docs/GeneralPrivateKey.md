# GeneralPrivateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **map[string]interface{}** | A private key corresponds to exactly one public key address. A private key can be used to move assets out of the wallet and sign transaction with the corresponding public key.  A private key is an array of integers (e.g.,   &#x60;[185,108,153,165,57,193,166,167,58,148,133,121,92,252,242,13,233,246,35,103,185,20,27,56,111,169,12,50,50,36,83,156,173,195,143,75,135,78,204,129,217,231,58,129,69,180,185,86,119,43,200,193,94,112,31,135,68,128,207,26,85,150,68,181]&#x60;).  &lt;a href&#x3D;\&quot;https://solflare.com\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Solflare&lt;/a&gt; is a popular wallet interface on Solana that allows you to export your private key in this format.  You can use &#x60;private_key&#x60; for Chains: &#x60;avalanche&#x60;, &#x60;binance_smart_chain&#x60;, &#x60;ethereum&#x60;, &#x60;solana&#x60;, &#x60;near&#x60;. | 

## Methods

### NewGeneralPrivateKey

`func NewGeneralPrivateKey(privateKey map[string]interface{}, ) *GeneralPrivateKey`

NewGeneralPrivateKey instantiates a new GeneralPrivateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralPrivateKeyWithDefaults

`func NewGeneralPrivateKeyWithDefaults() *GeneralPrivateKey`

NewGeneralPrivateKeyWithDefaults instantiates a new GeneralPrivateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *GeneralPrivateKey) GetPrivateKey() map[string]interface{}`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GeneralPrivateKey) GetPrivateKeyOk() (*map[string]interface{}, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GeneralPrivateKey) SetPrivateKey(v map[string]interface{})`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


