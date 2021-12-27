# GetCandyMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CandyMachineId** | Pointer to **string** | The ID of the candy machine. This is the same as &#x60;config_address&#x60; for &#x60;v2&#x60; candy machines (supply either).  | [optional] 
**ConfigAddress** | Pointer to **string** | The configuration address of the candy machine. This is the same as &#x60;candy_machine_id&#x60; for &#x60;v2&#x60; candy machines (supply either).  | [optional] 
**Uuid** | Pointer to **string** | The uuid of the candy machine. This is an alphanumeric string of length six (e.g., HpVdfP), which corresponds to the first six characters of the config_address.  | [optional] 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]
**CandyMachineContractVersion** | Pointer to **string** | The candy machine contract of the candy machine for which you are retrieving the metadata. If you are providing &#x60;v1&#x60; candy machine ID, set this to &#x60;v1&#x60;. If you are providing &#x60;v2&#x60; candy machine ID, set this to &#x60;v2&#x60;. If you don&#39;t know which version your candy machine is, check out &lt;a href&#x3D;\&quot;#operation/solanaGetAccountIsCandyMachine\&quot;&gt;this endpoint&lt;/a&gt;.  | [optional] [default to "v1"]

## Methods

### NewGetCandyMetadataRequest

`func NewGetCandyMetadataRequest() *GetCandyMetadataRequest`

NewGetCandyMetadataRequest instantiates a new GetCandyMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCandyMetadataRequestWithDefaults

`func NewGetCandyMetadataRequestWithDefaults() *GetCandyMetadataRequest`

NewGetCandyMetadataRequestWithDefaults instantiates a new GetCandyMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandyMachineId

`func (o *GetCandyMetadataRequest) GetCandyMachineId() string`

GetCandyMachineId returns the CandyMachineId field if non-nil, zero value otherwise.

### GetCandyMachineIdOk

`func (o *GetCandyMetadataRequest) GetCandyMachineIdOk() (*string, bool)`

GetCandyMachineIdOk returns a tuple with the CandyMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyMachineId

`func (o *GetCandyMetadataRequest) SetCandyMachineId(v string)`

SetCandyMachineId sets CandyMachineId field to given value.

### HasCandyMachineId

`func (o *GetCandyMetadataRequest) HasCandyMachineId() bool`

HasCandyMachineId returns a boolean if a field has been set.

### GetConfigAddress

`func (o *GetCandyMetadataRequest) GetConfigAddress() string`

GetConfigAddress returns the ConfigAddress field if non-nil, zero value otherwise.

### GetConfigAddressOk

`func (o *GetCandyMetadataRequest) GetConfigAddressOk() (*string, bool)`

GetConfigAddressOk returns a tuple with the ConfigAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAddress

`func (o *GetCandyMetadataRequest) SetConfigAddress(v string)`

SetConfigAddress sets ConfigAddress field to given value.

### HasConfigAddress

`func (o *GetCandyMetadataRequest) HasConfigAddress() bool`

HasConfigAddress returns a boolean if a field has been set.

### GetUuid

`func (o *GetCandyMetadataRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetCandyMetadataRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetCandyMetadataRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetCandyMetadataRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetNetwork

`func (o *GetCandyMetadataRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetCandyMetadataRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetCandyMetadataRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetCandyMetadataRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCandyMachineContractVersion

`func (o *GetCandyMetadataRequest) GetCandyMachineContractVersion() string`

GetCandyMachineContractVersion returns the CandyMachineContractVersion field if non-nil, zero value otherwise.

### GetCandyMachineContractVersionOk

`func (o *GetCandyMetadataRequest) GetCandyMachineContractVersionOk() (*string, bool)`

GetCandyMachineContractVersionOk returns a tuple with the CandyMachineContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyMachineContractVersion

`func (o *GetCandyMetadataRequest) SetCandyMachineContractVersion(v string)`

SetCandyMachineContractVersion sets CandyMachineContractVersion field to given value.

### HasCandyMachineContractVersion

`func (o *GetCandyMetadataRequest) HasCandyMachineContractVersion() bool`

HasCandyMachineContractVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


