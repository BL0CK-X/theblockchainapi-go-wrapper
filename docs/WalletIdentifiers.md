# WalletIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HexPublicAddress** | Pointer to **string** | The public address in hex form of the wallet. This is commonly used in Ethereum, Binance, and Avalanche C-Chain. (e.g., &#x60;0xB2b2d42C3adA171633E36b427F062f85A642F453&#x60;)  | [optional] 
**HexPublicKey** | Pointer to **string** | The public key in hex form. This is used in the Near blockchain. (e.g., &#x60;0xdf96e3094ab33558dbe14e743aceded3779307b931e12bd6f526fe450db70910&#x60;)  | [optional] 
**BechPublicAddress** | Pointer to **string** | The bech-32 encoded public address of the wallet, commonly used for the X and P chains of Avalanche. (e.g., &#x60;X-avax1rlxm7ygahlzfjgj4s965t0lk0ucm8v48rc9r5r&#x60;)  | [optional] 
**PublicKey** | Pointer to **string** | The public key in base58 form. This is used in the Solana blockchain. (e.g., &#x60;3Gdu3Uf97jw4Kac8xHEm5Hqmob3BrvJhf3We8s9t5Q2c&#x60;)  | [optional] 

## Methods

### NewWalletIdentifiers

`func NewWalletIdentifiers() *WalletIdentifiers`

NewWalletIdentifiers instantiates a new WalletIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletIdentifiersWithDefaults

`func NewWalletIdentifiersWithDefaults() *WalletIdentifiers`

NewWalletIdentifiersWithDefaults instantiates a new WalletIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHexPublicAddress

`func (o *WalletIdentifiers) GetHexPublicAddress() string`

GetHexPublicAddress returns the HexPublicAddress field if non-nil, zero value otherwise.

### GetHexPublicAddressOk

`func (o *WalletIdentifiers) GetHexPublicAddressOk() (*string, bool)`

GetHexPublicAddressOk returns a tuple with the HexPublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPublicAddress

`func (o *WalletIdentifiers) SetHexPublicAddress(v string)`

SetHexPublicAddress sets HexPublicAddress field to given value.

### HasHexPublicAddress

`func (o *WalletIdentifiers) HasHexPublicAddress() bool`

HasHexPublicAddress returns a boolean if a field has been set.

### GetHexPublicKey

`func (o *WalletIdentifiers) GetHexPublicKey() string`

GetHexPublicKey returns the HexPublicKey field if non-nil, zero value otherwise.

### GetHexPublicKeyOk

`func (o *WalletIdentifiers) GetHexPublicKeyOk() (*string, bool)`

GetHexPublicKeyOk returns a tuple with the HexPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPublicKey

`func (o *WalletIdentifiers) SetHexPublicKey(v string)`

SetHexPublicKey sets HexPublicKey field to given value.

### HasHexPublicKey

`func (o *WalletIdentifiers) HasHexPublicKey() bool`

HasHexPublicKey returns a boolean if a field has been set.

### GetBechPublicAddress

`func (o *WalletIdentifiers) GetBechPublicAddress() string`

GetBechPublicAddress returns the BechPublicAddress field if non-nil, zero value otherwise.

### GetBechPublicAddressOk

`func (o *WalletIdentifiers) GetBechPublicAddressOk() (*string, bool)`

GetBechPublicAddressOk returns a tuple with the BechPublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBechPublicAddress

`func (o *WalletIdentifiers) SetBechPublicAddress(v string)`

SetBechPublicAddress sets BechPublicAddress field to given value.

### HasBechPublicAddress

`func (o *WalletIdentifiers) HasBechPublicAddress() bool`

HasBechPublicAddress returns a boolean if a field has been set.

### GetPublicKey

`func (o *WalletIdentifiers) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WalletIdentifiers) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WalletIdentifiers) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WalletIdentifiers) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


