# AvalancheCChainPublicAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HexPublicAddress** | Pointer to **string** | The public address in hex form of the wallet. This is commonly used in Ethereum, Binance, and Avalanche C-Chain. (e.g., &#x60;0xB2b2d42C3adA171633E36b427F062f85A642F453&#x60;)  | [optional] 
**HexPublicKey** | Pointer to **string** | To public key in hex form of the wallet. This is hashed to get the hex public address. (e.g., &#x60;0x0f7182c2c2f79aca13847bed68c67662c021df868ee5d20a78df6095e4cd162610c63ec9050989a3755a18255cdd707e50678bfd762db3f0feea647610e974c4&#x60;)  | [optional] 

## Methods

### NewAvalancheCChainPublicAddress

`func NewAvalancheCChainPublicAddress() *AvalancheCChainPublicAddress`

NewAvalancheCChainPublicAddress instantiates a new AvalancheCChainPublicAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvalancheCChainPublicAddressWithDefaults

`func NewAvalancheCChainPublicAddressWithDefaults() *AvalancheCChainPublicAddress`

NewAvalancheCChainPublicAddressWithDefaults instantiates a new AvalancheCChainPublicAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHexPublicAddress

`func (o *AvalancheCChainPublicAddress) GetHexPublicAddress() string`

GetHexPublicAddress returns the HexPublicAddress field if non-nil, zero value otherwise.

### GetHexPublicAddressOk

`func (o *AvalancheCChainPublicAddress) GetHexPublicAddressOk() (*string, bool)`

GetHexPublicAddressOk returns a tuple with the HexPublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPublicAddress

`func (o *AvalancheCChainPublicAddress) SetHexPublicAddress(v string)`

SetHexPublicAddress sets HexPublicAddress field to given value.

### HasHexPublicAddress

`func (o *AvalancheCChainPublicAddress) HasHexPublicAddress() bool`

HasHexPublicAddress returns a boolean if a field has been set.

### GetHexPublicKey

`func (o *AvalancheCChainPublicAddress) GetHexPublicKey() string`

GetHexPublicKey returns the HexPublicKey field if non-nil, zero value otherwise.

### GetHexPublicKeyOk

`func (o *AvalancheCChainPublicAddress) GetHexPublicKeyOk() (*string, bool)`

GetHexPublicKeyOk returns a tuple with the HexPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPublicKey

`func (o *AvalancheCChainPublicAddress) SetHexPublicKey(v string)`

SetHexPublicKey sets HexPublicKey field to given value.

### HasHexPublicKey

`func (o *AvalancheCChainPublicAddress) HasHexPublicKey() bool`

HasHexPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


