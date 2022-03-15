# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionName** | Pointer to **string** | The name of the section under which this group is listed. A section contains several groups.  | [optional] 
**GroupName** | Pointer to **string** | The name of the group  | [optional] 
**GroupDescription** | Pointer to **string** | The description of the group  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionName

`func (o *Group) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *Group) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *Group) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.

### HasSectionName

`func (o *Group) HasSectionName() bool`

HasSectionName returns a boolean if a field has been set.

### GetGroupName

`func (o *Group) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *Group) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *Group) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *Group) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetGroupDescription

`func (o *Group) GetGroupDescription() string`

GetGroupDescription returns the GroupDescription field if non-nil, zero value otherwise.

### GetGroupDescriptionOk

`func (o *Group) GetGroupDescriptionOk() (*string, bool)`

GetGroupDescriptionOk returns a tuple with the GroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDescription

`func (o *Group) SetGroupDescription(v string)`

SetGroupDescription sets GroupDescription field to given value.

### HasGroupDescription

`func (o *Group) HasGroupDescription() bool`

HasGroupDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


