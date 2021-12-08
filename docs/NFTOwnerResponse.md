# NFTOwnerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NftOwner** | **string** | The public key address of the wallet that has ownership over the provided NFT | 

## Methods

### NewNFTOwnerResponse

`func NewNFTOwnerResponse(nftOwner string, ) *NFTOwnerResponse`

NewNFTOwnerResponse instantiates a new NFTOwnerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTOwnerResponseWithDefaults

`func NewNFTOwnerResponseWithDefaults() *NFTOwnerResponse`

NewNFTOwnerResponseWithDefaults instantiates a new NFTOwnerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNftOwner

`func (o *NFTOwnerResponse) GetNftOwner() string`

GetNftOwner returns the NftOwner field if non-nil, zero value otherwise.

### GetNftOwnerOk

`func (o *NFTOwnerResponse) GetNftOwnerOk() (*string, bool)`

GetNftOwnerOk returns a tuple with the NftOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftOwner

`func (o *NFTOwnerResponse) SetNftOwner(v string)`

SetNftOwner sets NftOwner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


