# GetCandyMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to **string** | The update authority of the candy machine | [optional] 
**Bump** | Pointer to **float32** |  | [optional] 
**CandyMachineId** | Pointer to **string** | The ID of the candy machine  | [optional] 
**ConfigAddress** | Pointer to **string** | The configuration public key address of the candy machine  | [optional] 
**Creators** | Pointer to [**[]GetCandyMetadataResponseCreators**](GetCandyMetadataResponseCreators.md) |  | [optional] 
**GoLiveDate** | Pointer to **float32** | The unix timestamp of the start date of the candy machine  | [optional] 
**IsMutable** | Pointer to **bool** |  | [optional] 
**ItemsAvailable** | Pointer to **float32** | The number of NFTs available for mint from the candy machine  | [optional] 
**ItemsRedeemed** | Pointer to **float32** | The number of NFTs minted already from the candy machine  | [optional] 
**MaxNumberOfLines** | Pointer to **float32** |  | [optional] 
**MaxSupply** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** | The price in Lamports for minting an NFT from the candy machine. 1e9 Lamport  &#x3D; 1 SOL  | [optional] 
**RetainAuthority** | Pointer to **bool** |  | [optional] 
**SellerFeeBasisPoints** | Pointer to **float32** | The royalty the creators receive on each sale after the primary sale (the initial minting) (denominated in basis points (e.g., 75 basis points &#x3D; 0.75%))  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TokenMint** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** | The uuid of the candy machine | [optional] 
**Wallet** | Pointer to **string** | The public key address of the wallet that recevies the proceeds from NFT mints  | [optional] 

## Methods

### NewGetCandyMetadataResponse

`func NewGetCandyMetadataResponse() *GetCandyMetadataResponse`

NewGetCandyMetadataResponse instantiates a new GetCandyMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCandyMetadataResponseWithDefaults

`func NewGetCandyMetadataResponseWithDefaults() *GetCandyMetadataResponse`

NewGetCandyMetadataResponseWithDefaults instantiates a new GetCandyMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *GetCandyMetadataResponse) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *GetCandyMetadataResponse) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *GetCandyMetadataResponse) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *GetCandyMetadataResponse) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetBump

`func (o *GetCandyMetadataResponse) GetBump() float32`

GetBump returns the Bump field if non-nil, zero value otherwise.

### GetBumpOk

`func (o *GetCandyMetadataResponse) GetBumpOk() (*float32, bool)`

GetBumpOk returns a tuple with the Bump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBump

`func (o *GetCandyMetadataResponse) SetBump(v float32)`

SetBump sets Bump field to given value.

### HasBump

`func (o *GetCandyMetadataResponse) HasBump() bool`

HasBump returns a boolean if a field has been set.

### GetCandyMachineId

`func (o *GetCandyMetadataResponse) GetCandyMachineId() string`

GetCandyMachineId returns the CandyMachineId field if non-nil, zero value otherwise.

### GetCandyMachineIdOk

`func (o *GetCandyMetadataResponse) GetCandyMachineIdOk() (*string, bool)`

GetCandyMachineIdOk returns a tuple with the CandyMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyMachineId

`func (o *GetCandyMetadataResponse) SetCandyMachineId(v string)`

SetCandyMachineId sets CandyMachineId field to given value.

### HasCandyMachineId

`func (o *GetCandyMetadataResponse) HasCandyMachineId() bool`

HasCandyMachineId returns a boolean if a field has been set.

### GetConfigAddress

`func (o *GetCandyMetadataResponse) GetConfigAddress() string`

GetConfigAddress returns the ConfigAddress field if non-nil, zero value otherwise.

### GetConfigAddressOk

`func (o *GetCandyMetadataResponse) GetConfigAddressOk() (*string, bool)`

GetConfigAddressOk returns a tuple with the ConfigAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAddress

`func (o *GetCandyMetadataResponse) SetConfigAddress(v string)`

SetConfigAddress sets ConfigAddress field to given value.

### HasConfigAddress

`func (o *GetCandyMetadataResponse) HasConfigAddress() bool`

HasConfigAddress returns a boolean if a field has been set.

### GetCreators

`func (o *GetCandyMetadataResponse) GetCreators() []GetCandyMetadataResponseCreators`

GetCreators returns the Creators field if non-nil, zero value otherwise.

### GetCreatorsOk

`func (o *GetCandyMetadataResponse) GetCreatorsOk() (*[]GetCandyMetadataResponseCreators, bool)`

GetCreatorsOk returns a tuple with the Creators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreators

`func (o *GetCandyMetadataResponse) SetCreators(v []GetCandyMetadataResponseCreators)`

SetCreators sets Creators field to given value.

### HasCreators

`func (o *GetCandyMetadataResponse) HasCreators() bool`

HasCreators returns a boolean if a field has been set.

### GetGoLiveDate

`func (o *GetCandyMetadataResponse) GetGoLiveDate() float32`

GetGoLiveDate returns the GoLiveDate field if non-nil, zero value otherwise.

### GetGoLiveDateOk

`func (o *GetCandyMetadataResponse) GetGoLiveDateOk() (*float32, bool)`

GetGoLiveDateOk returns a tuple with the GoLiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoLiveDate

`func (o *GetCandyMetadataResponse) SetGoLiveDate(v float32)`

SetGoLiveDate sets GoLiveDate field to given value.

### HasGoLiveDate

`func (o *GetCandyMetadataResponse) HasGoLiveDate() bool`

HasGoLiveDate returns a boolean if a field has been set.

### GetIsMutable

`func (o *GetCandyMetadataResponse) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *GetCandyMetadataResponse) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *GetCandyMetadataResponse) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *GetCandyMetadataResponse) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetItemsAvailable

`func (o *GetCandyMetadataResponse) GetItemsAvailable() float32`

GetItemsAvailable returns the ItemsAvailable field if non-nil, zero value otherwise.

### GetItemsAvailableOk

`func (o *GetCandyMetadataResponse) GetItemsAvailableOk() (*float32, bool)`

GetItemsAvailableOk returns a tuple with the ItemsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsAvailable

`func (o *GetCandyMetadataResponse) SetItemsAvailable(v float32)`

SetItemsAvailable sets ItemsAvailable field to given value.

### HasItemsAvailable

`func (o *GetCandyMetadataResponse) HasItemsAvailable() bool`

HasItemsAvailable returns a boolean if a field has been set.

### GetItemsRedeemed

`func (o *GetCandyMetadataResponse) GetItemsRedeemed() float32`

GetItemsRedeemed returns the ItemsRedeemed field if non-nil, zero value otherwise.

### GetItemsRedeemedOk

`func (o *GetCandyMetadataResponse) GetItemsRedeemedOk() (*float32, bool)`

GetItemsRedeemedOk returns a tuple with the ItemsRedeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsRedeemed

`func (o *GetCandyMetadataResponse) SetItemsRedeemed(v float32)`

SetItemsRedeemed sets ItemsRedeemed field to given value.

### HasItemsRedeemed

`func (o *GetCandyMetadataResponse) HasItemsRedeemed() bool`

HasItemsRedeemed returns a boolean if a field has been set.

### GetMaxNumberOfLines

`func (o *GetCandyMetadataResponse) GetMaxNumberOfLines() float32`

GetMaxNumberOfLines returns the MaxNumberOfLines field if non-nil, zero value otherwise.

### GetMaxNumberOfLinesOk

`func (o *GetCandyMetadataResponse) GetMaxNumberOfLinesOk() (*float32, bool)`

GetMaxNumberOfLinesOk returns a tuple with the MaxNumberOfLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfLines

`func (o *GetCandyMetadataResponse) SetMaxNumberOfLines(v float32)`

SetMaxNumberOfLines sets MaxNumberOfLines field to given value.

### HasMaxNumberOfLines

`func (o *GetCandyMetadataResponse) HasMaxNumberOfLines() bool`

HasMaxNumberOfLines returns a boolean if a field has been set.

### GetMaxSupply

`func (o *GetCandyMetadataResponse) GetMaxSupply() float32`

GetMaxSupply returns the MaxSupply field if non-nil, zero value otherwise.

### GetMaxSupplyOk

`func (o *GetCandyMetadataResponse) GetMaxSupplyOk() (*float32, bool)`

GetMaxSupplyOk returns a tuple with the MaxSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupply

`func (o *GetCandyMetadataResponse) SetMaxSupply(v float32)`

SetMaxSupply sets MaxSupply field to given value.

### HasMaxSupply

`func (o *GetCandyMetadataResponse) HasMaxSupply() bool`

HasMaxSupply returns a boolean if a field has been set.

### GetPrice

`func (o *GetCandyMetadataResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetCandyMetadataResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetCandyMetadataResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetCandyMetadataResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRetainAuthority

`func (o *GetCandyMetadataResponse) GetRetainAuthority() bool`

GetRetainAuthority returns the RetainAuthority field if non-nil, zero value otherwise.

### GetRetainAuthorityOk

`func (o *GetCandyMetadataResponse) GetRetainAuthorityOk() (*bool, bool)`

GetRetainAuthorityOk returns a tuple with the RetainAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAuthority

`func (o *GetCandyMetadataResponse) SetRetainAuthority(v bool)`

SetRetainAuthority sets RetainAuthority field to given value.

### HasRetainAuthority

`func (o *GetCandyMetadataResponse) HasRetainAuthority() bool`

HasRetainAuthority returns a boolean if a field has been set.

### GetSellerFeeBasisPoints

`func (o *GetCandyMetadataResponse) GetSellerFeeBasisPoints() float32`

GetSellerFeeBasisPoints returns the SellerFeeBasisPoints field if non-nil, zero value otherwise.

### GetSellerFeeBasisPointsOk

`func (o *GetCandyMetadataResponse) GetSellerFeeBasisPointsOk() (*float32, bool)`

GetSellerFeeBasisPointsOk returns a tuple with the SellerFeeBasisPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerFeeBasisPoints

`func (o *GetCandyMetadataResponse) SetSellerFeeBasisPoints(v float32)`

SetSellerFeeBasisPoints sets SellerFeeBasisPoints field to given value.

### HasSellerFeeBasisPoints

`func (o *GetCandyMetadataResponse) HasSellerFeeBasisPoints() bool`

HasSellerFeeBasisPoints returns a boolean if a field has been set.

### GetSymbol

`func (o *GetCandyMetadataResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetCandyMetadataResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetCandyMetadataResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetCandyMetadataResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTokenMint

`func (o *GetCandyMetadataResponse) GetTokenMint() string`

GetTokenMint returns the TokenMint field if non-nil, zero value otherwise.

### GetTokenMintOk

`func (o *GetCandyMetadataResponse) GetTokenMintOk() (*string, bool)`

GetTokenMintOk returns a tuple with the TokenMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMint

`func (o *GetCandyMetadataResponse) SetTokenMint(v string)`

SetTokenMint sets TokenMint field to given value.

### HasTokenMint

`func (o *GetCandyMetadataResponse) HasTokenMint() bool`

HasTokenMint returns a boolean if a field has been set.

### GetUuid

`func (o *GetCandyMetadataResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetCandyMetadataResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetCandyMetadataResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetCandyMetadataResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWallet

`func (o *GetCandyMetadataResponse) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *GetCandyMetadataResponse) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *GetCandyMetadataResponse) SetWallet(v string)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *GetCandyMetadataResponse) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


