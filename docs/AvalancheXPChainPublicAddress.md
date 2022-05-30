# AvalancheXPChainPublicAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BechPublicAddress** | Pointer to **string** | The bech-32 encoded public address of the wallet, commonly used for the X and P chains of Avalanche. (e.g., &#x60;X-avax1rlxm7ygahlzfjgj4s965t0lk0ucm8v48rc9r5r&#x60;)  | [optional] 
**HexPublicAddress** | Pointer to **string** | The public address in hex form of the wallet. This is commonly used in Ethereum, Binance, and Avalanche C-Chain. (e.g., &#x60;0xB2b2d42C3adA171633E36b427F062f85A642F453&#x60;)  | [optional] 

## Methods

### NewAvalancheXPChainPublicAddress

`func NewAvalancheXPChainPublicAddress() *AvalancheXPChainPublicAddress`

NewAvalancheXPChainPublicAddress instantiates a new AvalancheXPChainPublicAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvalancheXPChainPublicAddressWithDefaults

`func NewAvalancheXPChainPublicAddressWithDefaults() *AvalancheXPChainPublicAddress`

NewAvalancheXPChainPublicAddressWithDefaults instantiates a new AvalancheXPChainPublicAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBechPublicAddress

`func (o *AvalancheXPChainPublicAddress) GetBechPublicAddress() string`

GetBechPublicAddress returns the BechPublicAddress field if non-nil, zero value otherwise.

### GetBechPublicAddressOk

`func (o *AvalancheXPChainPublicAddress) GetBechPublicAddressOk() (*string, bool)`

GetBechPublicAddressOk returns a tuple with the BechPublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBechPublicAddress

`func (o *AvalancheXPChainPublicAddress) SetBechPublicAddress(v string)`

SetBechPublicAddress sets BechPublicAddress field to given value.

### HasBechPublicAddress

`func (o *AvalancheXPChainPublicAddress) HasBechPublicAddress() bool`

HasBechPublicAddress returns a boolean if a field has been set.

### GetHexPublicAddress

`func (o *AvalancheXPChainPublicAddress) GetHexPublicAddress() string`

GetHexPublicAddress returns the HexPublicAddress field if non-nil, zero value otherwise.

### GetHexPublicAddressOk

`func (o *AvalancheXPChainPublicAddress) GetHexPublicAddressOk() (*string, bool)`

GetHexPublicAddressOk returns a tuple with the HexPublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPublicAddress

`func (o *AvalancheXPChainPublicAddress) SetHexPublicAddress(v string)`

SetHexPublicAddress sets HexPublicAddress field to given value.

### HasHexPublicAddress

`func (o *AvalancheXPChainPublicAddress) HasHexPublicAddress() bool`

HasHexPublicAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


