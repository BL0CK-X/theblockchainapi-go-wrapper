# GetAllNFTsResponseMinted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NftMetadata** | Pointer to [**NFT**](NFT.md) |  | [optional] 
**PubKeyHash** | Pointer to **string** | Use this to verify the NFT | [optional] 

## Methods

### NewGetAllNFTsResponseMinted

`func NewGetAllNFTsResponseMinted() *GetAllNFTsResponseMinted`

NewGetAllNFTsResponseMinted instantiates a new GetAllNFTsResponseMinted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllNFTsResponseMintedWithDefaults

`func NewGetAllNFTsResponseMintedWithDefaults() *GetAllNFTsResponseMinted`

NewGetAllNFTsResponseMintedWithDefaults instantiates a new GetAllNFTsResponseMinted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNftMetadata

`func (o *GetAllNFTsResponseMinted) GetNftMetadata() NFT`

GetNftMetadata returns the NftMetadata field if non-nil, zero value otherwise.

### GetNftMetadataOk

`func (o *GetAllNFTsResponseMinted) GetNftMetadataOk() (*NFT, bool)`

GetNftMetadataOk returns a tuple with the NftMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftMetadata

`func (o *GetAllNFTsResponseMinted) SetNftMetadata(v NFT)`

SetNftMetadata sets NftMetadata field to given value.

### HasNftMetadata

`func (o *GetAllNFTsResponseMinted) HasNftMetadata() bool`

HasNftMetadata returns a boolean if a field has been set.

### GetPubKeyHash

`func (o *GetAllNFTsResponseMinted) GetPubKeyHash() string`

GetPubKeyHash returns the PubKeyHash field if non-nil, zero value otherwise.

### GetPubKeyHashOk

`func (o *GetAllNFTsResponseMinted) GetPubKeyHashOk() (*string, bool)`

GetPubKeyHashOk returns a tuple with the PubKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKeyHash

`func (o *GetAllNFTsResponseMinted) SetPubKeyHash(v string)`

SetPubKeyHash sets PubKeyHash field to given value.

### HasPubKeyHash

`func (o *GetAllNFTsResponseMinted) HasPubKeyHash() bool`

HasPubKeyHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


