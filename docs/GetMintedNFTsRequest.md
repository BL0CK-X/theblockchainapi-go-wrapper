# GetMintedNFTsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CandyMachineId** | **string** | The ID of the candy machine | 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]

## Methods

### NewGetMintedNFTsRequest

`func NewGetMintedNFTsRequest(candyMachineId string, ) *GetMintedNFTsRequest`

NewGetMintedNFTsRequest instantiates a new GetMintedNFTsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMintedNFTsRequestWithDefaults

`func NewGetMintedNFTsRequestWithDefaults() *GetMintedNFTsRequest`

NewGetMintedNFTsRequestWithDefaults instantiates a new GetMintedNFTsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandyMachineId

`func (o *GetMintedNFTsRequest) GetCandyMachineId() string`

GetCandyMachineId returns the CandyMachineId field if non-nil, zero value otherwise.

### GetCandyMachineIdOk

`func (o *GetMintedNFTsRequest) GetCandyMachineIdOk() (*string, bool)`

GetCandyMachineIdOk returns a tuple with the CandyMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyMachineId

`func (o *GetMintedNFTsRequest) SetCandyMachineId(v string)`

SetCandyMachineId sets CandyMachineId field to given value.


### GetNetwork

`func (o *GetMintedNFTsRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetMintedNFTsRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetMintedNFTsRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetMintedNFTsRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


