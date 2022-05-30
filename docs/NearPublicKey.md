# NearPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HexPublicKey** | Pointer to **string** | The public key in hex form. This is used in the Near blockchain. (e.g., &#x60;0xdf96e3094ab33558dbe14e743aceded3779307b931e12bd6f526fe450db70910&#x60;)  | [optional] 
**HexPublicAddress** | Pointer to **string** | The public address in hex form of the wallet. This is commonly used in Ethereum, Binance, and Avalanche C-Chain. (e.g., &#x60;0xB2b2d42C3adA171633E36b427F062f85A642F453&#x60;)  | [optional] 
**PublicKey** | Pointer to **string** | The public key in base58 form. This is used in the Solana blockchain. (e.g., &#x60;3Gdu3Uf97jw4Kac8xHEm5Hqmob3BrvJhf3We8s9t5Q2c&#x60;)  | [optional] 

## Methods

### NewNearPublicKey

`func NewNearPublicKey() *NearPublicKey`

NewNearPublicKey instantiates a new NearPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNearPublicKeyWithDefaults

`func NewNearPublicKeyWithDefaults() *NearPublicKey`

NewNearPublicKeyWithDefaults instantiates a new NearPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHexPublicKey

`func (o *NearPublicKey) GetHexPublicKey() string`

GetHexPublicKey returns the HexPublicKey field if non-nil, zero value otherwise.

### GetHexPublicKeyOk

`func (o *NearPublicKey) GetHexPublicKeyOk() (*string, bool)`

GetHexPublicKeyOk returns a tuple with the HexPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPublicKey

`func (o *NearPublicKey) SetHexPublicKey(v string)`

SetHexPublicKey sets HexPublicKey field to given value.

### HasHexPublicKey

`func (o *NearPublicKey) HasHexPublicKey() bool`

HasHexPublicKey returns a boolean if a field has been set.

### GetHexPublicAddress

`func (o *NearPublicKey) GetHexPublicAddress() string`

GetHexPublicAddress returns the HexPublicAddress field if non-nil, zero value otherwise.

### GetHexPublicAddressOk

`func (o *NearPublicKey) GetHexPublicAddressOk() (*string, bool)`

GetHexPublicAddressOk returns a tuple with the HexPublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPublicAddress

`func (o *NearPublicKey) SetHexPublicAddress(v string)`

SetHexPublicAddress sets HexPublicAddress field to given value.

### HasHexPublicAddress

`func (o *NearPublicKey) HasHexPublicAddress() bool`

HasHexPublicAddress returns a boolean if a field has been set.

### GetPublicKey

`func (o *NearPublicKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *NearPublicKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *NearPublicKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *NearPublicKey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


