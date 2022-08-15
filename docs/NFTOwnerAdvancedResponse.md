# NFTOwnerAdvancedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | Pointer to [**NFTOwnerAdvancedResponseContract**](NFTOwnerAdvancedResponseContract.md) |  | [optional] 
**Owner** | Pointer to **string** | The public key of the wallet that has true ownership over the provided NFT. If listed, it is the lister. If placed on a loan, it is the loanee. If staked, it is the staker. If burned, it is the burner. Etc. Returns &#x60;null&#x60; in the edge case that we were unable to find the true owner. If this happens, please report it to us and we will try to adapt for this edge case.  | [optional] 
**Price** | Pointer to **float32** | The price of the NFT, if listed or loaned. If held, staked, or burned, this is null. | [optional] 
**State** | Pointer to **string** | The state of the NFT | [optional] 

## Methods

### NewNFTOwnerAdvancedResponse

`func NewNFTOwnerAdvancedResponse() *NFTOwnerAdvancedResponse`

NewNFTOwnerAdvancedResponse instantiates a new NFTOwnerAdvancedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTOwnerAdvancedResponseWithDefaults

`func NewNFTOwnerAdvancedResponseWithDefaults() *NFTOwnerAdvancedResponse`

NewNFTOwnerAdvancedResponseWithDefaults instantiates a new NFTOwnerAdvancedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *NFTOwnerAdvancedResponse) GetContract() NFTOwnerAdvancedResponseContract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *NFTOwnerAdvancedResponse) GetContractOk() (*NFTOwnerAdvancedResponseContract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *NFTOwnerAdvancedResponse) SetContract(v NFTOwnerAdvancedResponseContract)`

SetContract sets Contract field to given value.

### HasContract

`func (o *NFTOwnerAdvancedResponse) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetOwner

`func (o *NFTOwnerAdvancedResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NFTOwnerAdvancedResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NFTOwnerAdvancedResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NFTOwnerAdvancedResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrice

`func (o *NFTOwnerAdvancedResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NFTOwnerAdvancedResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NFTOwnerAdvancedResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NFTOwnerAdvancedResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetState

`func (o *NFTOwnerAdvancedResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NFTOwnerAdvancedResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NFTOwnerAdvancedResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NFTOwnerAdvancedResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


