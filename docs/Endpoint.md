# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | The ID of the project. This is auto-generated upon project creation and cannot currently be changed.  | 
**Version** | **string** | The project version under which the endpoint exists  | 
**Path** | **string** | The path of the endpoint  | 
**OperationId** | **string** | The operation ID of the endpoint.  | 
**ReadableName** | **string** | The name of the endpoint formatted in a readable way (e.g., Get Endpoint Metadata).  | 
**Summary** | Pointer to **string** | The summary of the endpoint to be displayed in the sidebar on the left of the mini-API&#39;s documentation  | [optional] 
**Description** | Pointer to **string** | A description of what the endpoint does. This will be shown in the mini-API documentation.  | [optional] 
**Credits** | **float32** | The price of the endpoint. Valid values are integers from 1 to 100.  | 
**GroupName** | Pointer to **string** | The name of the group of endpoints that the endpoint comes from  | [optional] 
**InputSpecification** | [**[]ParameterSpecification**](ParameterSpecification.md) | A list of dictionaries. Each dictionary describes one parameter for the input, including the name, type, required boolean, and description values of that parameter. | 
**InputExamples** | **map[string]interface{}** | An example of the JSON input that a user might send. Limited to one example currently.  | 
**OutputSpecification** | [**[]ParameterSpecification**](ParameterSpecification.md) | A list of dictionaries. Each dictionary describes one parameter for the input, including the name, type, required boolean, and description values of that parameter. | 
**OutputExamples** | **map[string]interface{}** | An example of the JSON output that a user might send. Limited to one example currently.  | 

## Methods

### NewEndpoint

`func NewEndpoint(projectId string, version string, path string, operationId string, readableName string, credits float32, inputSpecification []ParameterSpecification, inputExamples map[string]interface{}, outputSpecification []ParameterSpecification, outputExamples map[string]interface{}, ) *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *Endpoint) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Endpoint) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Endpoint) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersion

`func (o *Endpoint) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Endpoint) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Endpoint) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPath

`func (o *Endpoint) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Endpoint) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Endpoint) SetPath(v string)`

SetPath sets Path field to given value.


### GetOperationId

`func (o *Endpoint) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *Endpoint) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *Endpoint) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetReadableName

`func (o *Endpoint) GetReadableName() string`

GetReadableName returns the ReadableName field if non-nil, zero value otherwise.

### GetReadableNameOk

`func (o *Endpoint) GetReadableNameOk() (*string, bool)`

GetReadableNameOk returns a tuple with the ReadableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadableName

`func (o *Endpoint) SetReadableName(v string)`

SetReadableName sets ReadableName field to given value.


### GetSummary

`func (o *Endpoint) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Endpoint) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Endpoint) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Endpoint) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *Endpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Endpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Endpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Endpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCredits

`func (o *Endpoint) GetCredits() float32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *Endpoint) GetCreditsOk() (*float32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *Endpoint) SetCredits(v float32)`

SetCredits sets Credits field to given value.


### GetGroupName

`func (o *Endpoint) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *Endpoint) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *Endpoint) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *Endpoint) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetInputSpecification

`func (o *Endpoint) GetInputSpecification() []ParameterSpecification`

GetInputSpecification returns the InputSpecification field if non-nil, zero value otherwise.

### GetInputSpecificationOk

`func (o *Endpoint) GetInputSpecificationOk() (*[]ParameterSpecification, bool)`

GetInputSpecificationOk returns a tuple with the InputSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSpecification

`func (o *Endpoint) SetInputSpecification(v []ParameterSpecification)`

SetInputSpecification sets InputSpecification field to given value.


### GetInputExamples

`func (o *Endpoint) GetInputExamples() map[string]interface{}`

GetInputExamples returns the InputExamples field if non-nil, zero value otherwise.

### GetInputExamplesOk

`func (o *Endpoint) GetInputExamplesOk() (*map[string]interface{}, bool)`

GetInputExamplesOk returns a tuple with the InputExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputExamples

`func (o *Endpoint) SetInputExamples(v map[string]interface{})`

SetInputExamples sets InputExamples field to given value.


### GetOutputSpecification

`func (o *Endpoint) GetOutputSpecification() []ParameterSpecification`

GetOutputSpecification returns the OutputSpecification field if non-nil, zero value otherwise.

### GetOutputSpecificationOk

`func (o *Endpoint) GetOutputSpecificationOk() (*[]ParameterSpecification, bool)`

GetOutputSpecificationOk returns a tuple with the OutputSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSpecification

`func (o *Endpoint) SetOutputSpecification(v []ParameterSpecification)`

SetOutputSpecification sets OutputSpecification field to given value.


### GetOutputExamples

`func (o *Endpoint) GetOutputExamples() map[string]interface{}`

GetOutputExamples returns the OutputExamples field if non-nil, zero value otherwise.

### GetOutputExamplesOk

`func (o *Endpoint) GetOutputExamplesOk() (*map[string]interface{}, bool)`

GetOutputExamplesOk returns a tuple with the OutputExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputExamples

`func (o *Endpoint) SetOutputExamples(v map[string]interface{})`

SetOutputExamples sets OutputExamples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


