# CreateTestCandyMachineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | [**Wallet**](Wallet.md) |  | 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]
**CandyMachineContractVersion** | Pointer to **string** | The contract you want to use to create the candy machine. Note: Metaplex disabled the creation of &#x60;v1&#x60; candy machines on their smart contract, and so we no longer support the creation of &#x60;v1&#x60; test candy machines.  | [optional] [default to "v2"]
**IncludeGatekeeper** | Pointer to **bool** | Whether or not to include a gatekeeper for testing purposes. Only applies to v2 candy machines. | [optional] [default to false]

## Methods

### NewCreateTestCandyMachineRequest

`func NewCreateTestCandyMachineRequest(wallet Wallet, ) *CreateTestCandyMachineRequest`

NewCreateTestCandyMachineRequest instantiates a new CreateTestCandyMachineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestCandyMachineRequestWithDefaults

`func NewCreateTestCandyMachineRequestWithDefaults() *CreateTestCandyMachineRequest`

NewCreateTestCandyMachineRequestWithDefaults instantiates a new CreateTestCandyMachineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallet

`func (o *CreateTestCandyMachineRequest) GetWallet() Wallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *CreateTestCandyMachineRequest) GetWalletOk() (*Wallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *CreateTestCandyMachineRequest) SetWallet(v Wallet)`

SetWallet sets Wallet field to given value.


### GetNetwork

`func (o *CreateTestCandyMachineRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateTestCandyMachineRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateTestCandyMachineRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateTestCandyMachineRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCandyMachineContractVersion

`func (o *CreateTestCandyMachineRequest) GetCandyMachineContractVersion() string`

GetCandyMachineContractVersion returns the CandyMachineContractVersion field if non-nil, zero value otherwise.

### GetCandyMachineContractVersionOk

`func (o *CreateTestCandyMachineRequest) GetCandyMachineContractVersionOk() (*string, bool)`

GetCandyMachineContractVersionOk returns a tuple with the CandyMachineContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyMachineContractVersion

`func (o *CreateTestCandyMachineRequest) SetCandyMachineContractVersion(v string)`

SetCandyMachineContractVersion sets CandyMachineContractVersion field to given value.

### HasCandyMachineContractVersion

`func (o *CreateTestCandyMachineRequest) HasCandyMachineContractVersion() bool`

HasCandyMachineContractVersion returns a boolean if a field has been set.

### GetIncludeGatekeeper

`func (o *CreateTestCandyMachineRequest) GetIncludeGatekeeper() bool`

GetIncludeGatekeeper returns the IncludeGatekeeper field if non-nil, zero value otherwise.

### GetIncludeGatekeeperOk

`func (o *CreateTestCandyMachineRequest) GetIncludeGatekeeperOk() (*bool, bool)`

GetIncludeGatekeeperOk returns a tuple with the IncludeGatekeeper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGatekeeper

`func (o *CreateTestCandyMachineRequest) SetIncludeGatekeeper(v bool)`

SetIncludeGatekeeper sets IncludeGatekeeper field to given value.

### HasIncludeGatekeeper

`func (o *CreateTestCandyMachineRequest) HasIncludeGatekeeper() bool`

HasIncludeGatekeeper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


