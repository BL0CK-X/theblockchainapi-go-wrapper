# CCProjectCreateRequestCustomerIdToCollect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdType** | **string** | What type of identifier you are collecting, either an \&quot;email\&quot; or \&quot;misc\&quot;ellanous. Miscellanous simply means you are collecting something other than an email.   | 
**Name** | **string** | The name of the customer ID input presented to the user  | 
**Description** | Pointer to **string** | The description / stated purpose of the customer ID input presented to the user  | [optional] [default to ""]
**Required** | Pointer to **bool** | Whether to require the customer ID  | [optional] [default to false]

## Methods

### NewCCProjectCreateRequestCustomerIdToCollect

`func NewCCProjectCreateRequestCustomerIdToCollect(idType string, name string, ) *CCProjectCreateRequestCustomerIdToCollect`

NewCCProjectCreateRequestCustomerIdToCollect instantiates a new CCProjectCreateRequestCustomerIdToCollect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCProjectCreateRequestCustomerIdToCollectWithDefaults

`func NewCCProjectCreateRequestCustomerIdToCollectWithDefaults() *CCProjectCreateRequestCustomerIdToCollect`

NewCCProjectCreateRequestCustomerIdToCollectWithDefaults instantiates a new CCProjectCreateRequestCustomerIdToCollect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdType

`func (o *CCProjectCreateRequestCustomerIdToCollect) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *CCProjectCreateRequestCustomerIdToCollect) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *CCProjectCreateRequestCustomerIdToCollect) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetName

`func (o *CCProjectCreateRequestCustomerIdToCollect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CCProjectCreateRequestCustomerIdToCollect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CCProjectCreateRequestCustomerIdToCollect) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CCProjectCreateRequestCustomerIdToCollect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CCProjectCreateRequestCustomerIdToCollect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CCProjectCreateRequestCustomerIdToCollect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CCProjectCreateRequestCustomerIdToCollect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *CCProjectCreateRequestCustomerIdToCollect) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CCProjectCreateRequestCustomerIdToCollect) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CCProjectCreateRequestCustomerIdToCollect) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CCProjectCreateRequestCustomerIdToCollect) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


