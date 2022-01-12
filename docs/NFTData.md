# NFTData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creators** | Pointer to **[]string** | The creators of the NFT  | [optional] 
**Name** | Pointer to **string** | The name of the NFT  | [optional] 
**Share** | Pointer to **[]float32** | The portion of the seller fee that each creator receives. Index 0 corresponds to creator at index 0 in the creator array, and so on. Length of share array must be the same as the length of the creator array.  | [optional] 
**Symbol** | Pointer to **string** | The symbol of the NFT  | [optional] 
**Uri** | Pointer to **string** | The URI of the NFT  | [optional] 
**Verified** | Pointer to **[]float32** | Whether or not the respective creator has signed the minting transaction of the NFT. Index 0 corresponds to creator at index 0 in the creator array, and so on. Length of verified array must be the same as the length of the creator array.  | [optional] 

## Methods

### NewNFTData

`func NewNFTData() *NFTData`

NewNFTData instantiates a new NFTData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTDataWithDefaults

`func NewNFTDataWithDefaults() *NFTData`

NewNFTDataWithDefaults instantiates a new NFTData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreators

`func (o *NFTData) GetCreators() []string`

GetCreators returns the Creators field if non-nil, zero value otherwise.

### GetCreatorsOk

`func (o *NFTData) GetCreatorsOk() (*[]string, bool)`

GetCreatorsOk returns a tuple with the Creators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreators

`func (o *NFTData) SetCreators(v []string)`

SetCreators sets Creators field to given value.

### HasCreators

`func (o *NFTData) HasCreators() bool`

HasCreators returns a boolean if a field has been set.

### GetName

`func (o *NFTData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NFTData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NFTData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NFTData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShare

`func (o *NFTData) GetShare() []float32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *NFTData) GetShareOk() (*[]float32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *NFTData) SetShare(v []float32)`

SetShare sets Share field to given value.

### HasShare

`func (o *NFTData) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSymbol

`func (o *NFTData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NFTData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NFTData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *NFTData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetUri

`func (o *NFTData) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *NFTData) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *NFTData) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *NFTData) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetVerified

`func (o *NFTData) GetVerified() []float32`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *NFTData) GetVerifiedOk() (*[]float32, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *NFTData) SetVerified(v []float32)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *NFTData) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


