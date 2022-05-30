# TokenMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | The symbol of the coin  | [optional] 
**Name** | Pointer to **string** | The name of the coin  | [optional] 
**Decimals** | Pointer to **float32** |  | [optional] 
**MinterBlockchainIdentifier** | Pointer to **string** | The &#x60;mint_authority&#x60; (Solana) or &#x60;master_minter&#x60; (Ethereum)  | [optional] 
**TotalSupply** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenMetadataResponse

`func NewTokenMetadataResponse() *TokenMetadataResponse`

NewTokenMetadataResponse instantiates a new TokenMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenMetadataResponseWithDefaults

`func NewTokenMetadataResponseWithDefaults() *TokenMetadataResponse`

NewTokenMetadataResponseWithDefaults instantiates a new TokenMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *TokenMetadataResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenMetadataResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenMetadataResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TokenMetadataResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *TokenMetadataResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenMetadataResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenMetadataResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenMetadataResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDecimals

`func (o *TokenMetadataResponse) GetDecimals() float32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenMetadataResponse) GetDecimalsOk() (*float32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenMetadataResponse) SetDecimals(v float32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TokenMetadataResponse) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetMinterBlockchainIdentifier

`func (o *TokenMetadataResponse) GetMinterBlockchainIdentifier() string`

GetMinterBlockchainIdentifier returns the MinterBlockchainIdentifier field if non-nil, zero value otherwise.

### GetMinterBlockchainIdentifierOk

`func (o *TokenMetadataResponse) GetMinterBlockchainIdentifierOk() (*string, bool)`

GetMinterBlockchainIdentifierOk returns a tuple with the MinterBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinterBlockchainIdentifier

`func (o *TokenMetadataResponse) SetMinterBlockchainIdentifier(v string)`

SetMinterBlockchainIdentifier sets MinterBlockchainIdentifier field to given value.

### HasMinterBlockchainIdentifier

`func (o *TokenMetadataResponse) HasMinterBlockchainIdentifier() bool`

HasMinterBlockchainIdentifier returns a boolean if a field has been set.

### GetTotalSupply

`func (o *TokenMetadataResponse) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *TokenMetadataResponse) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *TokenMetadataResponse) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *TokenMetadataResponse) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


