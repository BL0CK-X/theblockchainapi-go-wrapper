# ProjectDeploymentURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the bucket where you will upload your binary  | [optional] 
**DeploymentVersion** | Pointer to **string** | An internal version tracking the version of the binary uploaded  | [optional] 
**Fields** | Pointer to **map[string]interface{}** | Information necessary to sign the upload URL  | [optional] 

## Methods

### NewProjectDeploymentURL

`func NewProjectDeploymentURL() *ProjectDeploymentURL`

NewProjectDeploymentURL instantiates a new ProjectDeploymentURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDeploymentURLWithDefaults

`func NewProjectDeploymentURLWithDefaults() *ProjectDeploymentURL`

NewProjectDeploymentURLWithDefaults instantiates a new ProjectDeploymentURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ProjectDeploymentURL) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectDeploymentURL) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectDeploymentURL) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectDeploymentURL) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDeploymentVersion

`func (o *ProjectDeploymentURL) GetDeploymentVersion() string`

GetDeploymentVersion returns the DeploymentVersion field if non-nil, zero value otherwise.

### GetDeploymentVersionOk

`func (o *ProjectDeploymentURL) GetDeploymentVersionOk() (*string, bool)`

GetDeploymentVersionOk returns a tuple with the DeploymentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersion

`func (o *ProjectDeploymentURL) SetDeploymentVersion(v string)`

SetDeploymentVersion sets DeploymentVersion field to given value.

### HasDeploymentVersion

`func (o *ProjectDeploymentURL) HasDeploymentVersion() bool`

HasDeploymentVersion returns a boolean if a field has been set.

### GetFields

`func (o *ProjectDeploymentURL) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ProjectDeploymentURL) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ProjectDeploymentURL) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *ProjectDeploymentURL) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


