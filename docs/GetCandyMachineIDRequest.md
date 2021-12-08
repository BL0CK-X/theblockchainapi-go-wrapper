# GetCandyMachineIDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MintAddress** | **string** | The address of the NFT. NOT the mint authority address | 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]

## Methods

### NewGetCandyMachineIDRequest

`func NewGetCandyMachineIDRequest(mintAddress string, ) *GetCandyMachineIDRequest`

NewGetCandyMachineIDRequest instantiates a new GetCandyMachineIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCandyMachineIDRequestWithDefaults

`func NewGetCandyMachineIDRequestWithDefaults() *GetCandyMachineIDRequest`

NewGetCandyMachineIDRequestWithDefaults instantiates a new GetCandyMachineIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMintAddress

`func (o *GetCandyMachineIDRequest) GetMintAddress() string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *GetCandyMachineIDRequest) GetMintAddressOk() (*string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *GetCandyMachineIDRequest) SetMintAddress(v string)`

SetMintAddress sets MintAddress field to given value.


### GetNetwork

`func (o *GetCandyMachineIDRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetCandyMachineIDRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetCandyMachineIDRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetCandyMachineIDRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


