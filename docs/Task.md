# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | **map[string]interface{}** |  | 
**TimeCompleted** | Pointer to **float32** |  | [optional] 
**StatusCode** | **float32** |  | 

## Methods

### NewTask

`func NewTask(response map[string]interface{}, statusCode float32, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *Task) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Task) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Task) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.


### GetTimeCompleted

`func (o *Task) GetTimeCompleted() float32`

GetTimeCompleted returns the TimeCompleted field if non-nil, zero value otherwise.

### GetTimeCompletedOk

`func (o *Task) GetTimeCompletedOk() (*float32, bool)`

GetTimeCompletedOk returns a tuple with the TimeCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeCompleted

`func (o *Task) SetTimeCompleted(v float32)`

SetTimeCompleted sets TimeCompleted field to given value.

### HasTimeCompleted

`func (o *Task) HasTimeCompleted() bool`

HasTimeCompleted returns a boolean if a field has been set.

### GetStatusCode

`func (o *Task) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *Task) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *Task) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


