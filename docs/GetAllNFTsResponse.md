# GetAllNFTsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MintedNfts** | Pointer to [**[]GetAllNFTsResponseMintedNfts**](GetAllNFTsResponseMintedNfts.md) | The minted NFTs. Only filled in for &#x60;v1&#x60; candy machines. Left empty for &#x60;v2&#x60;. | [optional] 
**UnmintedNfts** | Pointer to [**[]GetAllNFTsResponseUnmintedNfts**](GetAllNFTsResponseUnmintedNfts.md) | The unminted NFTs. Only filled in for &#x60;v1&#x60; candy machines. Left empty for &#x60;v2&#x60;. | [optional] 
**AllNfts** | Pointer to [**[]GetAllNFTsResponseUnmintedNfts**](GetAllNFTsResponseUnmintedNfts.md) | The list of all NFTs. Filled in for both &#x60;v1&#x60; and &#x60;v2&#x60; NFTs. | [optional] 
**Accurate** | Pointer to **bool** | Whether or not the division of NFTs among minted and unminted is accurate. If it is not accurate, then it is possible that NFTs have been included in the &#x60;minted&#x60; list that are not actually minted. If it is accurate, then the split of  minted and unminted is correct. &#x60;v1&#x60; candy machines always return a correct minted/unminted split.   | [optional] 

## Methods

### NewGetAllNFTsResponse

`func NewGetAllNFTsResponse() *GetAllNFTsResponse`

NewGetAllNFTsResponse instantiates a new GetAllNFTsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllNFTsResponseWithDefaults

`func NewGetAllNFTsResponseWithDefaults() *GetAllNFTsResponse`

NewGetAllNFTsResponseWithDefaults instantiates a new GetAllNFTsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMintedNfts

`func (o *GetAllNFTsResponse) GetMintedNfts() []GetAllNFTsResponseMintedNfts`

GetMintedNfts returns the MintedNfts field if non-nil, zero value otherwise.

### GetMintedNftsOk

`func (o *GetAllNFTsResponse) GetMintedNftsOk() (*[]GetAllNFTsResponseMintedNfts, bool)`

GetMintedNftsOk returns a tuple with the MintedNfts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintedNfts

`func (o *GetAllNFTsResponse) SetMintedNfts(v []GetAllNFTsResponseMintedNfts)`

SetMintedNfts sets MintedNfts field to given value.

### HasMintedNfts

`func (o *GetAllNFTsResponse) HasMintedNfts() bool`

HasMintedNfts returns a boolean if a field has been set.

### GetUnmintedNfts

`func (o *GetAllNFTsResponse) GetUnmintedNfts() []GetAllNFTsResponseUnmintedNfts`

GetUnmintedNfts returns the UnmintedNfts field if non-nil, zero value otherwise.

### GetUnmintedNftsOk

`func (o *GetAllNFTsResponse) GetUnmintedNftsOk() (*[]GetAllNFTsResponseUnmintedNfts, bool)`

GetUnmintedNftsOk returns a tuple with the UnmintedNfts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmintedNfts

`func (o *GetAllNFTsResponse) SetUnmintedNfts(v []GetAllNFTsResponseUnmintedNfts)`

SetUnmintedNfts sets UnmintedNfts field to given value.

### HasUnmintedNfts

`func (o *GetAllNFTsResponse) HasUnmintedNfts() bool`

HasUnmintedNfts returns a boolean if a field has been set.

### GetAllNfts

`func (o *GetAllNFTsResponse) GetAllNfts() []GetAllNFTsResponseUnmintedNfts`

GetAllNfts returns the AllNfts field if non-nil, zero value otherwise.

### GetAllNftsOk

`func (o *GetAllNFTsResponse) GetAllNftsOk() (*[]GetAllNFTsResponseUnmintedNfts, bool)`

GetAllNftsOk returns a tuple with the AllNfts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllNfts

`func (o *GetAllNFTsResponse) SetAllNfts(v []GetAllNFTsResponseUnmintedNfts)`

SetAllNfts sets AllNfts field to given value.

### HasAllNfts

`func (o *GetAllNFTsResponse) HasAllNfts() bool`

HasAllNfts returns a boolean if a field has been set.

### GetAccurate

`func (o *GetAllNFTsResponse) GetAccurate() bool`

GetAccurate returns the Accurate field if non-nil, zero value otherwise.

### GetAccurateOk

`func (o *GetAllNFTsResponse) GetAccurateOk() (*bool, bool)`

GetAccurateOk returns a tuple with the Accurate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccurate

`func (o *GetAllNFTsResponse) SetAccurate(v bool)`

SetAccurate sets Accurate field to given value.

### HasAccurate

`func (o *GetAllNFTsResponse) HasAccurate() bool`

HasAccurate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


