# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** | The ID of the project. This is auto-generated upon project creation and cannot currently be changed.  | [optional] 
**ProjectName** | Pointer to **string** | The name of the mini-API. This will be shown at the top of the mini-API&#39;s documentation.  | [optional] 
**ProjectDescription** | Pointer to **string** | The description of the mini-API. This will be shown at the top of the mini-API&#39;s documentation, below the title.  | [optional] 
**ContactEmail** | Pointer to **string** | The email where users of your mini-API can contact you. This will be shown at the top of the mini-API&#39;s documentation.  | [optional] 
**CurrentDocumentationVersion** | Pointer to **string** | The version of the API that the documentation is updated for. You can set the documentation version to any valid version. To see how to format the version string, see the description for &#x60;versions&#x60;.  | [optional] 
**Versions** | Pointer to **[]string** | The live versions of the project. An array of strings. We use Python&#39;s &#x60;version&#x60; package to see if it&#39;s a valid version and to compare versions (to see which is higher).  Read more about this Python package &lt;a href&#x3D;\&quot;https://packaging.pypa.io/en/latest/version.html#packaging.version.parse\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt;. | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) | A list of groups. A section contains groups, and groups contain API endpoints.   | [optional] 
**Endpoints** | Pointer to [**[]Endpoint**](Endpoint.md) | A list of groups. A section contains groups, and groups contain API endpoints.   | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *Project) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Project) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Project) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Project) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *Project) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *Project) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *Project) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *Project) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetProjectDescription

`func (o *Project) GetProjectDescription() string`

GetProjectDescription returns the ProjectDescription field if non-nil, zero value otherwise.

### GetProjectDescriptionOk

`func (o *Project) GetProjectDescriptionOk() (*string, bool)`

GetProjectDescriptionOk returns a tuple with the ProjectDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDescription

`func (o *Project) SetProjectDescription(v string)`

SetProjectDescription sets ProjectDescription field to given value.

### HasProjectDescription

`func (o *Project) HasProjectDescription() bool`

HasProjectDescription returns a boolean if a field has been set.

### GetContactEmail

`func (o *Project) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Project) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Project) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Project) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetCurrentDocumentationVersion

`func (o *Project) GetCurrentDocumentationVersion() string`

GetCurrentDocumentationVersion returns the CurrentDocumentationVersion field if non-nil, zero value otherwise.

### GetCurrentDocumentationVersionOk

`func (o *Project) GetCurrentDocumentationVersionOk() (*string, bool)`

GetCurrentDocumentationVersionOk returns a tuple with the CurrentDocumentationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDocumentationVersion

`func (o *Project) SetCurrentDocumentationVersion(v string)`

SetCurrentDocumentationVersion sets CurrentDocumentationVersion field to given value.

### HasCurrentDocumentationVersion

`func (o *Project) HasCurrentDocumentationVersion() bool`

HasCurrentDocumentationVersion returns a boolean if a field has been set.

### GetVersions

`func (o *Project) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Project) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Project) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *Project) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetGroups

`func (o *Project) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Project) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Project) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Project) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetEndpoints

`func (o *Project) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *Project) GetEndpointsOk() (*[]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *Project) SetEndpoints(v []Endpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *Project) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


