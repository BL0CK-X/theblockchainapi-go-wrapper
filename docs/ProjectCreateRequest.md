# ProjectCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | Pointer to **string** | The name of the mini-API. This will be shown at the top of the mini-API&#39;s documentation.  | [optional] 
**ProjectDescription** | Pointer to **string** | The description of the mini-API. This will be shown at the top of the mini-API&#39;s documentation, below the title.  | [optional] 
**ContactEmail** | Pointer to **string** | The email where users of your mini-API can contact you. This will be shown at the top of the mini-API&#39;s documentation.  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) | A list of groups. A section contains groups, and groups contain API endpoints.   | [optional] 

## Methods

### NewProjectCreateRequest

`func NewProjectCreateRequest() *ProjectCreateRequest`

NewProjectCreateRequest instantiates a new ProjectCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateRequestWithDefaults

`func NewProjectCreateRequestWithDefaults() *ProjectCreateRequest`

NewProjectCreateRequestWithDefaults instantiates a new ProjectCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *ProjectCreateRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectCreateRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectCreateRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ProjectCreateRequest) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetProjectDescription

`func (o *ProjectCreateRequest) GetProjectDescription() string`

GetProjectDescription returns the ProjectDescription field if non-nil, zero value otherwise.

### GetProjectDescriptionOk

`func (o *ProjectCreateRequest) GetProjectDescriptionOk() (*string, bool)`

GetProjectDescriptionOk returns a tuple with the ProjectDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDescription

`func (o *ProjectCreateRequest) SetProjectDescription(v string)`

SetProjectDescription sets ProjectDescription field to given value.

### HasProjectDescription

`func (o *ProjectCreateRequest) HasProjectDescription() bool`

HasProjectDescription returns a boolean if a field has been set.

### GetContactEmail

`func (o *ProjectCreateRequest) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *ProjectCreateRequest) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *ProjectCreateRequest) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *ProjectCreateRequest) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetGroups

`func (o *ProjectCreateRequest) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProjectCreateRequest) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProjectCreateRequest) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ProjectCreateRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


