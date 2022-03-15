# EndpointReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | The ID of the project. This is auto-generated upon project creation and cannot currently be changed.  | 
**Version** | **string** | The project version under which the endpoint exists  | 
**Path** | **string** | The path of the endpoint  | 

## Methods

### NewEndpointReference

`func NewEndpointReference(projectId string, version string, path string, ) *EndpointReference`

NewEndpointReference instantiates a new EndpointReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointReferenceWithDefaults

`func NewEndpointReferenceWithDefaults() *EndpointReference`

NewEndpointReferenceWithDefaults instantiates a new EndpointReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *EndpointReference) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EndpointReference) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EndpointReference) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersion

`func (o *EndpointReference) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EndpointReference) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EndpointReference) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPath

`func (o *EndpointReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EndpointReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EndpointReference) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


