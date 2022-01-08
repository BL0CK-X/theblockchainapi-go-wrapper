# MintNFTRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | [**Wallet**](Wallet.md) |  | 
**ConfigAddress** | **string** | The config address of the candy machine. You can retrieve this if you have the candy machine ID using &lt;a href&#x3D;\&quot;#operation/solanaGetCandyMachineDetails\&quot;&gt;this endpoint&lt;/a&gt; and retrieving the &#x60;config_address&#x60; from the response.  A candy machine ID is the same thing as a configuration address for v2 candy machines.  | 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]
**CandyMachineContractVersion** | Pointer to **string** | The candy machine contract of the candy machine from which you&#39;re minting. If you are minting from a &#x60;v1&#x60; candy machine ID, set this to &#x60;v1&#x60;. If you are minting from a &#x60;v2&#x60; candy machine ID, set this to &#x60;v2&#x60;. If you don&#39;t know which the version of your candy machine, check out &lt;a href&#x3D;\&quot;#operation/solanaGetAccountIsCandyMachine\&quot;&gt;this endpoint&lt;/a&gt;.  | [optional] [default to "v1"]

## Methods

### NewMintNFTRequest

`func NewMintNFTRequest(wallet Wallet, configAddress string, ) *MintNFTRequest`

NewMintNFTRequest instantiates a new MintNFTRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintNFTRequestWithDefaults

`func NewMintNFTRequestWithDefaults() *MintNFTRequest`

NewMintNFTRequestWithDefaults instantiates a new MintNFTRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallet

`func (o *MintNFTRequest) GetWallet() Wallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *MintNFTRequest) GetWalletOk() (*Wallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *MintNFTRequest) SetWallet(v Wallet)`

SetWallet sets Wallet field to given value.


### GetConfigAddress

`func (o *MintNFTRequest) GetConfigAddress() string`

GetConfigAddress returns the ConfigAddress field if non-nil, zero value otherwise.

### GetConfigAddressOk

`func (o *MintNFTRequest) GetConfigAddressOk() (*string, bool)`

GetConfigAddressOk returns a tuple with the ConfigAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAddress

`func (o *MintNFTRequest) SetConfigAddress(v string)`

SetConfigAddress sets ConfigAddress field to given value.


### GetNetwork

`func (o *MintNFTRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MintNFTRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MintNFTRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MintNFTRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCandyMachineContractVersion

`func (o *MintNFTRequest) GetCandyMachineContractVersion() string`

GetCandyMachineContractVersion returns the CandyMachineContractVersion field if non-nil, zero value otherwise.

### GetCandyMachineContractVersionOk

`func (o *MintNFTRequest) GetCandyMachineContractVersionOk() (*string, bool)`

GetCandyMachineContractVersionOk returns a tuple with the CandyMachineContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyMachineContractVersion

`func (o *MintNFTRequest) SetCandyMachineContractVersion(v string)`

SetCandyMachineContractVersion sets CandyMachineContractVersion field to given value.

### HasCandyMachineContractVersion

`func (o *MintNFTRequest) HasCandyMachineContractVersion() bool`

HasCandyMachineContractVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


