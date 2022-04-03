# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** | The platform for which the binary was compiled. Note that Darwin &#x3D; Mac. | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *InlineObject) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *InlineObject) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *InlineObject) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *InlineObject) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


