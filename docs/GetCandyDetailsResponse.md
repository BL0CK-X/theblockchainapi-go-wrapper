# GetCandyDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The uuid of the candy machine | [optional] 
**GoLiveDate** | Pointer to **float32** | The unix timestamp of the start date of the candy machine  | [optional] 
**Price** | Pointer to **float32** | The price in Lamports for minting an NFT from the candy machine. 1e9 Lamport  &#x3D; 1 SOL  | [optional] 
**ItemsAvailable** | Pointer to **float32** | The number of NFTs available for mint from the candy machine  | [optional] 
**ItemsRedeemed** | Pointer to **float32** | The number of NFTs minted already from the candy machine  | [optional] 
**TokenMint** | Pointer to **string** |  | [optional] 
**ConfigAddress** | Pointer to **string** | The configuration public key address of the candy machine  | [optional] 
**Wallet** | Pointer to **string** | The public key address of the wallet that recevies the proceeds from NFT mints  | [optional] 
**Bump** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetCandyDetailsResponse

`func NewGetCandyDetailsResponse() *GetCandyDetailsResponse`

NewGetCandyDetailsResponse instantiates a new GetCandyDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCandyDetailsResponseWithDefaults

`func NewGetCandyDetailsResponseWithDefaults() *GetCandyDetailsResponse`

NewGetCandyDetailsResponseWithDefaults instantiates a new GetCandyDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *GetCandyDetailsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetCandyDetailsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetCandyDetailsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetCandyDetailsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetGoLiveDate

`func (o *GetCandyDetailsResponse) GetGoLiveDate() float32`

GetGoLiveDate returns the GoLiveDate field if non-nil, zero value otherwise.

### GetGoLiveDateOk

`func (o *GetCandyDetailsResponse) GetGoLiveDateOk() (*float32, bool)`

GetGoLiveDateOk returns a tuple with the GoLiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoLiveDate

`func (o *GetCandyDetailsResponse) SetGoLiveDate(v float32)`

SetGoLiveDate sets GoLiveDate field to given value.

### HasGoLiveDate

`func (o *GetCandyDetailsResponse) HasGoLiveDate() bool`

HasGoLiveDate returns a boolean if a field has been set.

### GetPrice

`func (o *GetCandyDetailsResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetCandyDetailsResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetCandyDetailsResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetCandyDetailsResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetItemsAvailable

`func (o *GetCandyDetailsResponse) GetItemsAvailable() float32`

GetItemsAvailable returns the ItemsAvailable field if non-nil, zero value otherwise.

### GetItemsAvailableOk

`func (o *GetCandyDetailsResponse) GetItemsAvailableOk() (*float32, bool)`

GetItemsAvailableOk returns a tuple with the ItemsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsAvailable

`func (o *GetCandyDetailsResponse) SetItemsAvailable(v float32)`

SetItemsAvailable sets ItemsAvailable field to given value.

### HasItemsAvailable

`func (o *GetCandyDetailsResponse) HasItemsAvailable() bool`

HasItemsAvailable returns a boolean if a field has been set.

### GetItemsRedeemed

`func (o *GetCandyDetailsResponse) GetItemsRedeemed() float32`

GetItemsRedeemed returns the ItemsRedeemed field if non-nil, zero value otherwise.

### GetItemsRedeemedOk

`func (o *GetCandyDetailsResponse) GetItemsRedeemedOk() (*float32, bool)`

GetItemsRedeemedOk returns a tuple with the ItemsRedeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsRedeemed

`func (o *GetCandyDetailsResponse) SetItemsRedeemed(v float32)`

SetItemsRedeemed sets ItemsRedeemed field to given value.

### HasItemsRedeemed

`func (o *GetCandyDetailsResponse) HasItemsRedeemed() bool`

HasItemsRedeemed returns a boolean if a field has been set.

### GetTokenMint

`func (o *GetCandyDetailsResponse) GetTokenMint() string`

GetTokenMint returns the TokenMint field if non-nil, zero value otherwise.

### GetTokenMintOk

`func (o *GetCandyDetailsResponse) GetTokenMintOk() (*string, bool)`

GetTokenMintOk returns a tuple with the TokenMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMint

`func (o *GetCandyDetailsResponse) SetTokenMint(v string)`

SetTokenMint sets TokenMint field to given value.

### HasTokenMint

`func (o *GetCandyDetailsResponse) HasTokenMint() bool`

HasTokenMint returns a boolean if a field has been set.

### GetConfigAddress

`func (o *GetCandyDetailsResponse) GetConfigAddress() string`

GetConfigAddress returns the ConfigAddress field if non-nil, zero value otherwise.

### GetConfigAddressOk

`func (o *GetCandyDetailsResponse) GetConfigAddressOk() (*string, bool)`

GetConfigAddressOk returns a tuple with the ConfigAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAddress

`func (o *GetCandyDetailsResponse) SetConfigAddress(v string)`

SetConfigAddress sets ConfigAddress field to given value.

### HasConfigAddress

`func (o *GetCandyDetailsResponse) HasConfigAddress() bool`

HasConfigAddress returns a boolean if a field has been set.

### GetWallet

`func (o *GetCandyDetailsResponse) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *GetCandyDetailsResponse) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *GetCandyDetailsResponse) SetWallet(v string)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *GetCandyDetailsResponse) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetBump

`func (o *GetCandyDetailsResponse) GetBump() float32`

GetBump returns the Bump field if non-nil, zero value otherwise.

### GetBumpOk

`func (o *GetCandyDetailsResponse) GetBumpOk() (*float32, bool)`

GetBumpOk returns a tuple with the Bump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBump

`func (o *GetCandyDetailsResponse) SetBump(v float32)`

SetBump sets Bump field to given value.

### HasBump

`func (o *GetCandyDetailsResponse) HasBump() bool`

HasBump returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


