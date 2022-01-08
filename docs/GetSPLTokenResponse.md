# GetSPLTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decimals** | Pointer to **float32** | The number of decimals of the token. For example, if the USDC token has 6 decimals, then you need 1 * 10e6 &#x3D; 1,000,000 tokens to have 1 USDC. The purpose of this is that everything must be stored as  an integer. Thus, if there are $100 USDC in total, there must be $100 * 10e6 tokens in order for the $100 to be divisible into lower denominations than $1.  | [optional] 
**FreezeAuthority** | Pointer to **string** | Public key address | [optional] 
**MintAuthority** | Pointer to **string** | Public key address | [optional] 
**IsInitialized** | Pointer to **bool** |  | [optional] 
**Supply** | Pointer to **string** | The supply of the token | [optional] 

## Methods

### NewGetSPLTokenResponse

`func NewGetSPLTokenResponse() *GetSPLTokenResponse`

NewGetSPLTokenResponse instantiates a new GetSPLTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSPLTokenResponseWithDefaults

`func NewGetSPLTokenResponseWithDefaults() *GetSPLTokenResponse`

NewGetSPLTokenResponseWithDefaults instantiates a new GetSPLTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecimals

`func (o *GetSPLTokenResponse) GetDecimals() float32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *GetSPLTokenResponse) GetDecimalsOk() (*float32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *GetSPLTokenResponse) SetDecimals(v float32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *GetSPLTokenResponse) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetFreezeAuthority

`func (o *GetSPLTokenResponse) GetFreezeAuthority() string`

GetFreezeAuthority returns the FreezeAuthority field if non-nil, zero value otherwise.

### GetFreezeAuthorityOk

`func (o *GetSPLTokenResponse) GetFreezeAuthorityOk() (*string, bool)`

GetFreezeAuthorityOk returns a tuple with the FreezeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeAuthority

`func (o *GetSPLTokenResponse) SetFreezeAuthority(v string)`

SetFreezeAuthority sets FreezeAuthority field to given value.

### HasFreezeAuthority

`func (o *GetSPLTokenResponse) HasFreezeAuthority() bool`

HasFreezeAuthority returns a boolean if a field has been set.

### GetMintAuthority

`func (o *GetSPLTokenResponse) GetMintAuthority() string`

GetMintAuthority returns the MintAuthority field if non-nil, zero value otherwise.

### GetMintAuthorityOk

`func (o *GetSPLTokenResponse) GetMintAuthorityOk() (*string, bool)`

GetMintAuthorityOk returns a tuple with the MintAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAuthority

`func (o *GetSPLTokenResponse) SetMintAuthority(v string)`

SetMintAuthority sets MintAuthority field to given value.

### HasMintAuthority

`func (o *GetSPLTokenResponse) HasMintAuthority() bool`

HasMintAuthority returns a boolean if a field has been set.

### GetIsInitialized

`func (o *GetSPLTokenResponse) GetIsInitialized() bool`

GetIsInitialized returns the IsInitialized field if non-nil, zero value otherwise.

### GetIsInitializedOk

`func (o *GetSPLTokenResponse) GetIsInitializedOk() (*bool, bool)`

GetIsInitializedOk returns a tuple with the IsInitialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitialized

`func (o *GetSPLTokenResponse) SetIsInitialized(v bool)`

SetIsInitialized sets IsInitialized field to given value.

### HasIsInitialized

`func (o *GetSPLTokenResponse) HasIsInitialized() bool`

HasIsInitialized returns a boolean if a field has been set.

### GetSupply

`func (o *GetSPLTokenResponse) GetSupply() string`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *GetSPLTokenResponse) GetSupplyOk() (*string, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *GetSPLTokenResponse) SetSupply(v string)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *GetSPLTokenResponse) HasSupply() bool`

HasSupply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


