# NFTSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateAuthority** | Pointer to **string** | The public key of the update authority of the NFT | [optional] 
**UpdateAuthoritySearchMethod** | Pointer to **string** | Only &#x60;exact_match&#x60; supported at this time | [optional] [default to "exact_match"]
**MintAddress** | Pointer to **string** | The mint address of the NFT | [optional] 
**MintAddressSearchMethod** | Pointer to **string** | Only &#x60;exact_match&#x60; supported at this time | [optional] [default to "exact_match"]
**Name** | Pointer to **string** | The name of the NFT | [optional] 
**NameSearchMethod** | Pointer to **string** |  | [optional] [default to "exact_match"]
**Uri** | Pointer to **string** | The NFT&#39;s uri | [optional] 
**UriSearchMethod** | Pointer to **string** |  | [optional] [default to "exact_match"]
**Symbol** | Pointer to **string** | The symbol associated with the candy machine | [optional] 
**SymbolSearchMethod** | Pointer to **string** |  | [optional] [default to "exact_match"]
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]

## Methods

### NewNFTSearchRequest

`func NewNFTSearchRequest() *NFTSearchRequest`

NewNFTSearchRequest instantiates a new NFTSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTSearchRequestWithDefaults

`func NewNFTSearchRequestWithDefaults() *NFTSearchRequest`

NewNFTSearchRequestWithDefaults instantiates a new NFTSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateAuthority

`func (o *NFTSearchRequest) GetUpdateAuthority() string`

GetUpdateAuthority returns the UpdateAuthority field if non-nil, zero value otherwise.

### GetUpdateAuthorityOk

`func (o *NFTSearchRequest) GetUpdateAuthorityOk() (*string, bool)`

GetUpdateAuthorityOk returns a tuple with the UpdateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAuthority

`func (o *NFTSearchRequest) SetUpdateAuthority(v string)`

SetUpdateAuthority sets UpdateAuthority field to given value.

### HasUpdateAuthority

`func (o *NFTSearchRequest) HasUpdateAuthority() bool`

HasUpdateAuthority returns a boolean if a field has been set.

### GetUpdateAuthoritySearchMethod

`func (o *NFTSearchRequest) GetUpdateAuthoritySearchMethod() string`

GetUpdateAuthoritySearchMethod returns the UpdateAuthoritySearchMethod field if non-nil, zero value otherwise.

### GetUpdateAuthoritySearchMethodOk

`func (o *NFTSearchRequest) GetUpdateAuthoritySearchMethodOk() (*string, bool)`

GetUpdateAuthoritySearchMethodOk returns a tuple with the UpdateAuthoritySearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAuthoritySearchMethod

`func (o *NFTSearchRequest) SetUpdateAuthoritySearchMethod(v string)`

SetUpdateAuthoritySearchMethod sets UpdateAuthoritySearchMethod field to given value.

### HasUpdateAuthoritySearchMethod

`func (o *NFTSearchRequest) HasUpdateAuthoritySearchMethod() bool`

HasUpdateAuthoritySearchMethod returns a boolean if a field has been set.

### GetMintAddress

`func (o *NFTSearchRequest) GetMintAddress() string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *NFTSearchRequest) GetMintAddressOk() (*string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *NFTSearchRequest) SetMintAddress(v string)`

SetMintAddress sets MintAddress field to given value.

### HasMintAddress

`func (o *NFTSearchRequest) HasMintAddress() bool`

HasMintAddress returns a boolean if a field has been set.

### GetMintAddressSearchMethod

`func (o *NFTSearchRequest) GetMintAddressSearchMethod() string`

GetMintAddressSearchMethod returns the MintAddressSearchMethod field if non-nil, zero value otherwise.

### GetMintAddressSearchMethodOk

`func (o *NFTSearchRequest) GetMintAddressSearchMethodOk() (*string, bool)`

GetMintAddressSearchMethodOk returns a tuple with the MintAddressSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddressSearchMethod

`func (o *NFTSearchRequest) SetMintAddressSearchMethod(v string)`

SetMintAddressSearchMethod sets MintAddressSearchMethod field to given value.

### HasMintAddressSearchMethod

`func (o *NFTSearchRequest) HasMintAddressSearchMethod() bool`

HasMintAddressSearchMethod returns a boolean if a field has been set.

### GetName

`func (o *NFTSearchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NFTSearchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NFTSearchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NFTSearchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameSearchMethod

`func (o *NFTSearchRequest) GetNameSearchMethod() string`

GetNameSearchMethod returns the NameSearchMethod field if non-nil, zero value otherwise.

### GetNameSearchMethodOk

`func (o *NFTSearchRequest) GetNameSearchMethodOk() (*string, bool)`

GetNameSearchMethodOk returns a tuple with the NameSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSearchMethod

`func (o *NFTSearchRequest) SetNameSearchMethod(v string)`

SetNameSearchMethod sets NameSearchMethod field to given value.

### HasNameSearchMethod

`func (o *NFTSearchRequest) HasNameSearchMethod() bool`

HasNameSearchMethod returns a boolean if a field has been set.

### GetUri

`func (o *NFTSearchRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *NFTSearchRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *NFTSearchRequest) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *NFTSearchRequest) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUriSearchMethod

`func (o *NFTSearchRequest) GetUriSearchMethod() string`

GetUriSearchMethod returns the UriSearchMethod field if non-nil, zero value otherwise.

### GetUriSearchMethodOk

`func (o *NFTSearchRequest) GetUriSearchMethodOk() (*string, bool)`

GetUriSearchMethodOk returns a tuple with the UriSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSearchMethod

`func (o *NFTSearchRequest) SetUriSearchMethod(v string)`

SetUriSearchMethod sets UriSearchMethod field to given value.

### HasUriSearchMethod

`func (o *NFTSearchRequest) HasUriSearchMethod() bool`

HasUriSearchMethod returns a boolean if a field has been set.

### GetSymbol

`func (o *NFTSearchRequest) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NFTSearchRequest) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NFTSearchRequest) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *NFTSearchRequest) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSymbolSearchMethod

`func (o *NFTSearchRequest) GetSymbolSearchMethod() string`

GetSymbolSearchMethod returns the SymbolSearchMethod field if non-nil, zero value otherwise.

### GetSymbolSearchMethodOk

`func (o *NFTSearchRequest) GetSymbolSearchMethodOk() (*string, bool)`

GetSymbolSearchMethodOk returns a tuple with the SymbolSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolSearchMethod

`func (o *NFTSearchRequest) SetSymbolSearchMethod(v string)`

SetSymbolSearchMethod sets SymbolSearchMethod field to given value.

### HasSymbolSearchMethod

`func (o *NFTSearchRequest) HasSymbolSearchMethod() bool`

HasSymbolSearchMethod returns a boolean if a field has been set.

### GetNetwork

`func (o *NFTSearchRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NFTSearchRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NFTSearchRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NFTSearchRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


