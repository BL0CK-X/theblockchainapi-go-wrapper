# GetAllNFTsResponseMintedNfts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NftMetadata** | Pointer to [**NFT**](NFT.md) |  | [optional] 
**PubKeyHash** | Pointer to **string** | Use this to verify the NFT | [optional] 

## Methods

### NewGetAllNFTsResponseMintedNfts

`func NewGetAllNFTsResponseMintedNfts() *GetAllNFTsResponseMintedNfts`

NewGetAllNFTsResponseMintedNfts instantiates a new GetAllNFTsResponseMintedNfts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllNFTsResponseMintedNftsWithDefaults

`func NewGetAllNFTsResponseMintedNftsWithDefaults() *GetAllNFTsResponseMintedNfts`

NewGetAllNFTsResponseMintedNftsWithDefaults instantiates a new GetAllNFTsResponseMintedNfts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNftMetadata

`func (o *GetAllNFTsResponseMintedNfts) GetNftMetadata() NFT`

GetNftMetadata returns the NftMetadata field if non-nil, zero value otherwise.

### GetNftMetadataOk

`func (o *GetAllNFTsResponseMintedNfts) GetNftMetadataOk() (*NFT, bool)`

GetNftMetadataOk returns a tuple with the NftMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftMetadata

`func (o *GetAllNFTsResponseMintedNfts) SetNftMetadata(v NFT)`

SetNftMetadata sets NftMetadata field to given value.

### HasNftMetadata

`func (o *GetAllNFTsResponseMintedNfts) HasNftMetadata() bool`

HasNftMetadata returns a boolean if a field has been set.

### GetPubKeyHash

`func (o *GetAllNFTsResponseMintedNfts) GetPubKeyHash() string`

GetPubKeyHash returns the PubKeyHash field if non-nil, zero value otherwise.

### GetPubKeyHashOk

`func (o *GetAllNFTsResponseMintedNfts) GetPubKeyHashOk() (*string, bool)`

GetPubKeyHashOk returns a tuple with the PubKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKeyHash

`func (o *GetAllNFTsResponseMintedNfts) SetPubKeyHash(v string)`

SetPubKeyHash sets PubKeyHash field to given value.

### HasPubKeyHash

`func (o *GetAllNFTsResponseMintedNfts) HasPubKeyHash() bool`

HasPubKeyHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


