# ParameterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the parameter | [optional] 
**Name** | Pointer to **string** | The name of the parameter | [optional] 
**Description** | Pointer to **string** | The name of the parameter | [optional] 
**Required** | Pointer to **bool** | Whether or not this parameter is required | [optional] 

## Methods

### NewParameterSpecification

`func NewParameterSpecification() *ParameterSpecification`

NewParameterSpecification instantiates a new ParameterSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterSpecificationWithDefaults

`func NewParameterSpecificationWithDefaults() *ParameterSpecification`

NewParameterSpecificationWithDefaults instantiates a new ParameterSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParameterSpecification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterSpecification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterSpecification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ParameterSpecification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ParameterSpecification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterSpecification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterSpecification) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParameterSpecification) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ParameterSpecification) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterSpecification) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterSpecification) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterSpecification) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *ParameterSpecification) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ParameterSpecification) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ParameterSpecification) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ParameterSpecification) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


