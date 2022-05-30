# HexPrivateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HexPrivateKey** | **string** | A hex private key is the typical private key used on Ethereum, for example. It looks like this: &#x60;0x200b9e5baa38b0dc7551645be11b394e9bf2b04532e4af8824bed2b3de2e0dc0&#x60;.  You can use &#x60;hex_private_key&#x60; for chains: &#x60;avalanche&#x60;, &#x60;binance_smart_chain&#x60;, &#x60;ethereum&#x60;. | 

## Methods

### NewHexPrivateKey

`func NewHexPrivateKey(hexPrivateKey string, ) *HexPrivateKey`

NewHexPrivateKey instantiates a new HexPrivateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHexPrivateKeyWithDefaults

`func NewHexPrivateKeyWithDefaults() *HexPrivateKey`

NewHexPrivateKeyWithDefaults instantiates a new HexPrivateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHexPrivateKey

`func (o *HexPrivateKey) GetHexPrivateKey() string`

GetHexPrivateKey returns the HexPrivateKey field if non-nil, zero value otherwise.

### GetHexPrivateKeyOk

`func (o *HexPrivateKey) GetHexPrivateKeyOk() (*string, bool)`

GetHexPrivateKeyOk returns a tuple with the HexPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPrivateKey

`func (o *HexPrivateKey) SetHexPrivateKey(v string)`

SetHexPrivateKey sets HexPrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


