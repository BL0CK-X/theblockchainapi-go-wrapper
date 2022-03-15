# ListNFTsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NftsOwned** | Pointer to **[]string** | A list of mint addresses represented by a string | [optional] 
**NftsMetadata** | Pointer to [**[]NFT**](NFT.md) | A list of the dictionaries, where each dictionary is an NFT&#39;s metadata | [optional] 

## Methods

### NewListNFTsResponse

`func NewListNFTsResponse() *ListNFTsResponse`

NewListNFTsResponse instantiates a new ListNFTsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNFTsResponseWithDefaults

`func NewListNFTsResponseWithDefaults() *ListNFTsResponse`

NewListNFTsResponseWithDefaults instantiates a new ListNFTsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNftsOwned

`func (o *ListNFTsResponse) GetNftsOwned() []string`

GetNftsOwned returns the NftsOwned field if non-nil, zero value otherwise.

### GetNftsOwnedOk

`func (o *ListNFTsResponse) GetNftsOwnedOk() (*[]string, bool)`

GetNftsOwnedOk returns a tuple with the NftsOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftsOwned

`func (o *ListNFTsResponse) SetNftsOwned(v []string)`

SetNftsOwned sets NftsOwned field to given value.

### HasNftsOwned

`func (o *ListNFTsResponse) HasNftsOwned() bool`

HasNftsOwned returns a boolean if a field has been set.

### GetNftsMetadata

`func (o *ListNFTsResponse) GetNftsMetadata() []NFT`

GetNftsMetadata returns the NftsMetadata field if non-nil, zero value otherwise.

### GetNftsMetadataOk

`func (o *ListNFTsResponse) GetNftsMetadataOk() (*[]NFT, bool)`

GetNftsMetadataOk returns a tuple with the NftsMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftsMetadata

`func (o *ListNFTsResponse) SetNftsMetadata(v []NFT)`

SetNftsMetadata sets NftsMetadata field to given value.

### HasNftsMetadata

`func (o *ListNFTsResponse) HasNftsMetadata() bool`

HasNftsMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


