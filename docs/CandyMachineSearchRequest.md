# CandyMachineSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateAuthority** | Pointer to **string** | The public key of the update authority of the candy machine | [optional] 
**UpdateAuthoritySearchMethod** | Pointer to **string** | Only &#x60;exact_match&#x60; supported at this time | [optional] [default to "exact_match"]
**ConfigAddress** | Pointer to **string** | The public key of the configuration of the candy machine | [optional] 
**ConfigAddressSearchMethod** | Pointer to **string** | Only &#x60;exact_match&#x60; supported at this time | [optional] [default to "exact_match"]
**Uuid** | Pointer to **string** | The alphanumeric string of length six that corresponds to the candy machine. This is NOT the candy machine ID.  This is the first six letters of the configuration address and is also used to identify the candy machine. An example is &#x60;4zKV6i&#x60;.  | [optional] 
**UuidSearchMethod** | Pointer to **string** | Only &#x60;exact_match&#x60; supported at this time | [optional] [default to "exact_match"]
**Symbol** | Pointer to **string** | The symbol associated with the candy machine | [optional] 
**SymbolSearchMethod** | Pointer to **string** |  | [optional] [default to "exact_match"]
**NftName** | Pointer to **string** | The name of an NFT on the candy machine, minted or unminted. For example, to find The Solana Money Boys candy machine, you already know that each NFT is named \&quot;Solana Money Boy #0\&quot;, \&quot;Solana Money Boy #1\&quot;, and so on. So you could search with  nft_name&#x3D;\&quot;Solana Money Boy #0\&quot;, nft_name_index&#x3D;0, nft_name_search_method&#x3D;&#39;exact_match&#39;, for example, which would return the candy machine ID. This also works with candy machines that are not live but are uploaded.  | [optional] 
**NftNameIndex** | Pointer to **string** | The index of the NFT to check, e.g., the 2nd NFT would have an index of 1. We cannot search all NFTs on a candy machine currently, so you must search an NFT at a particular position, such as the first, second, and so on. In general, nft_name_index&#x3D;0 works for most use cases.  | [optional] [default to "0"]
**NftNameSearchMethod** | Pointer to **string** |  | [optional] [default to "exact_match"]
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]
**CandyMachineContractVersion** | Pointer to **string** | The candy machine contract you want to search.  If you want to search &#x60;v1&#x60; candy machines, set this to &#x60;v1&#x60;. If you want to search &#x60;v2&#x60; candy machines. set this to &#x60;v2&#x60;.  | [optional] [default to "v1"]

## Methods

### NewCandyMachineSearchRequest

`func NewCandyMachineSearchRequest() *CandyMachineSearchRequest`

NewCandyMachineSearchRequest instantiates a new CandyMachineSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCandyMachineSearchRequestWithDefaults

`func NewCandyMachineSearchRequestWithDefaults() *CandyMachineSearchRequest`

NewCandyMachineSearchRequestWithDefaults instantiates a new CandyMachineSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateAuthority

`func (o *CandyMachineSearchRequest) GetUpdateAuthority() string`

GetUpdateAuthority returns the UpdateAuthority field if non-nil, zero value otherwise.

### GetUpdateAuthorityOk

`func (o *CandyMachineSearchRequest) GetUpdateAuthorityOk() (*string, bool)`

GetUpdateAuthorityOk returns a tuple with the UpdateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAuthority

`func (o *CandyMachineSearchRequest) SetUpdateAuthority(v string)`

SetUpdateAuthority sets UpdateAuthority field to given value.

### HasUpdateAuthority

`func (o *CandyMachineSearchRequest) HasUpdateAuthority() bool`

HasUpdateAuthority returns a boolean if a field has been set.

### GetUpdateAuthoritySearchMethod

`func (o *CandyMachineSearchRequest) GetUpdateAuthoritySearchMethod() string`

GetUpdateAuthoritySearchMethod returns the UpdateAuthoritySearchMethod field if non-nil, zero value otherwise.

### GetUpdateAuthoritySearchMethodOk

`func (o *CandyMachineSearchRequest) GetUpdateAuthoritySearchMethodOk() (*string, bool)`

GetUpdateAuthoritySearchMethodOk returns a tuple with the UpdateAuthoritySearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAuthoritySearchMethod

`func (o *CandyMachineSearchRequest) SetUpdateAuthoritySearchMethod(v string)`

SetUpdateAuthoritySearchMethod sets UpdateAuthoritySearchMethod field to given value.

### HasUpdateAuthoritySearchMethod

`func (o *CandyMachineSearchRequest) HasUpdateAuthoritySearchMethod() bool`

HasUpdateAuthoritySearchMethod returns a boolean if a field has been set.

### GetConfigAddress

`func (o *CandyMachineSearchRequest) GetConfigAddress() string`

GetConfigAddress returns the ConfigAddress field if non-nil, zero value otherwise.

### GetConfigAddressOk

`func (o *CandyMachineSearchRequest) GetConfigAddressOk() (*string, bool)`

GetConfigAddressOk returns a tuple with the ConfigAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAddress

`func (o *CandyMachineSearchRequest) SetConfigAddress(v string)`

SetConfigAddress sets ConfigAddress field to given value.

### HasConfigAddress

`func (o *CandyMachineSearchRequest) HasConfigAddress() bool`

HasConfigAddress returns a boolean if a field has been set.

### GetConfigAddressSearchMethod

`func (o *CandyMachineSearchRequest) GetConfigAddressSearchMethod() string`

GetConfigAddressSearchMethod returns the ConfigAddressSearchMethod field if non-nil, zero value otherwise.

### GetConfigAddressSearchMethodOk

`func (o *CandyMachineSearchRequest) GetConfigAddressSearchMethodOk() (*string, bool)`

GetConfigAddressSearchMethodOk returns a tuple with the ConfigAddressSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAddressSearchMethod

`func (o *CandyMachineSearchRequest) SetConfigAddressSearchMethod(v string)`

SetConfigAddressSearchMethod sets ConfigAddressSearchMethod field to given value.

### HasConfigAddressSearchMethod

`func (o *CandyMachineSearchRequest) HasConfigAddressSearchMethod() bool`

HasConfigAddressSearchMethod returns a boolean if a field has been set.

### GetUuid

`func (o *CandyMachineSearchRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CandyMachineSearchRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CandyMachineSearchRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CandyMachineSearchRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUuidSearchMethod

`func (o *CandyMachineSearchRequest) GetUuidSearchMethod() string`

GetUuidSearchMethod returns the UuidSearchMethod field if non-nil, zero value otherwise.

### GetUuidSearchMethodOk

`func (o *CandyMachineSearchRequest) GetUuidSearchMethodOk() (*string, bool)`

GetUuidSearchMethodOk returns a tuple with the UuidSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidSearchMethod

`func (o *CandyMachineSearchRequest) SetUuidSearchMethod(v string)`

SetUuidSearchMethod sets UuidSearchMethod field to given value.

### HasUuidSearchMethod

`func (o *CandyMachineSearchRequest) HasUuidSearchMethod() bool`

HasUuidSearchMethod returns a boolean if a field has been set.

### GetSymbol

`func (o *CandyMachineSearchRequest) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CandyMachineSearchRequest) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CandyMachineSearchRequest) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CandyMachineSearchRequest) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSymbolSearchMethod

`func (o *CandyMachineSearchRequest) GetSymbolSearchMethod() string`

GetSymbolSearchMethod returns the SymbolSearchMethod field if non-nil, zero value otherwise.

### GetSymbolSearchMethodOk

`func (o *CandyMachineSearchRequest) GetSymbolSearchMethodOk() (*string, bool)`

GetSymbolSearchMethodOk returns a tuple with the SymbolSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolSearchMethod

`func (o *CandyMachineSearchRequest) SetSymbolSearchMethod(v string)`

SetSymbolSearchMethod sets SymbolSearchMethod field to given value.

### HasSymbolSearchMethod

`func (o *CandyMachineSearchRequest) HasSymbolSearchMethod() bool`

HasSymbolSearchMethod returns a boolean if a field has been set.

### GetNftName

`func (o *CandyMachineSearchRequest) GetNftName() string`

GetNftName returns the NftName field if non-nil, zero value otherwise.

### GetNftNameOk

`func (o *CandyMachineSearchRequest) GetNftNameOk() (*string, bool)`

GetNftNameOk returns a tuple with the NftName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftName

`func (o *CandyMachineSearchRequest) SetNftName(v string)`

SetNftName sets NftName field to given value.

### HasNftName

`func (o *CandyMachineSearchRequest) HasNftName() bool`

HasNftName returns a boolean if a field has been set.

### GetNftNameIndex

`func (o *CandyMachineSearchRequest) GetNftNameIndex() string`

GetNftNameIndex returns the NftNameIndex field if non-nil, zero value otherwise.

### GetNftNameIndexOk

`func (o *CandyMachineSearchRequest) GetNftNameIndexOk() (*string, bool)`

GetNftNameIndexOk returns a tuple with the NftNameIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftNameIndex

`func (o *CandyMachineSearchRequest) SetNftNameIndex(v string)`

SetNftNameIndex sets NftNameIndex field to given value.

### HasNftNameIndex

`func (o *CandyMachineSearchRequest) HasNftNameIndex() bool`

HasNftNameIndex returns a boolean if a field has been set.

### GetNftNameSearchMethod

`func (o *CandyMachineSearchRequest) GetNftNameSearchMethod() string`

GetNftNameSearchMethod returns the NftNameSearchMethod field if non-nil, zero value otherwise.

### GetNftNameSearchMethodOk

`func (o *CandyMachineSearchRequest) GetNftNameSearchMethodOk() (*string, bool)`

GetNftNameSearchMethodOk returns a tuple with the NftNameSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftNameSearchMethod

`func (o *CandyMachineSearchRequest) SetNftNameSearchMethod(v string)`

SetNftNameSearchMethod sets NftNameSearchMethod field to given value.

### HasNftNameSearchMethod

`func (o *CandyMachineSearchRequest) HasNftNameSearchMethod() bool`

HasNftNameSearchMethod returns a boolean if a field has been set.

### GetNetwork

`func (o *CandyMachineSearchRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CandyMachineSearchRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CandyMachineSearchRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CandyMachineSearchRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCandyMachineContractVersion

`func (o *CandyMachineSearchRequest) GetCandyMachineContractVersion() string`

GetCandyMachineContractVersion returns the CandyMachineContractVersion field if non-nil, zero value otherwise.

### GetCandyMachineContractVersionOk

`func (o *CandyMachineSearchRequest) GetCandyMachineContractVersionOk() (*string, bool)`

GetCandyMachineContractVersionOk returns a tuple with the CandyMachineContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyMachineContractVersion

`func (o *CandyMachineSearchRequest) SetCandyMachineContractVersion(v string)`

SetCandyMachineContractVersion sets CandyMachineContractVersion field to given value.

### HasCandyMachineContractVersion

`func (o *CandyMachineSearchRequest) HasCandyMachineContractVersion() bool`

HasCandyMachineContractVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


