# GetCandyDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CandyMachineId** | **string** | The ID of the candy machine  | 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]

## Methods

### NewGetCandyDetailsRequest

`func NewGetCandyDetailsRequest(candyMachineId string, ) *GetCandyDetailsRequest`

NewGetCandyDetailsRequest instantiates a new GetCandyDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCandyDetailsRequestWithDefaults

`func NewGetCandyDetailsRequestWithDefaults() *GetCandyDetailsRequest`

NewGetCandyDetailsRequestWithDefaults instantiates a new GetCandyDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandyMachineId

`func (o *GetCandyDetailsRequest) GetCandyMachineId() string`

GetCandyMachineId returns the CandyMachineId field if non-nil, zero value otherwise.

### GetCandyMachineIdOk

`func (o *GetCandyDetailsRequest) GetCandyMachineIdOk() (*string, bool)`

GetCandyMachineIdOk returns a tuple with the CandyMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandyMachineId

`func (o *GetCandyDetailsRequest) SetCandyMachineId(v string)`

SetCandyMachineId sets CandyMachineId field to given value.


### GetNetwork

`func (o *GetCandyDetailsRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetCandyDetailsRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetCandyDetailsRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetCandyDetailsRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


