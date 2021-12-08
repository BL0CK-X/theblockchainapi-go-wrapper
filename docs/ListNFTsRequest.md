# ListNFTsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** | The public key of the account whose list of owned NFTs you want to get   | 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]

## Methods

### NewListNFTsRequest

`func NewListNFTsRequest(publicKey string, ) *ListNFTsRequest`

NewListNFTsRequest instantiates a new ListNFTsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNFTsRequestWithDefaults

`func NewListNFTsRequestWithDefaults() *ListNFTsRequest`

NewListNFTsRequestWithDefaults instantiates a new ListNFTsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *ListNFTsRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ListNFTsRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ListNFTsRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetNetwork

`func (o *ListNFTsRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListNFTsRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListNFTsRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ListNFTsRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


