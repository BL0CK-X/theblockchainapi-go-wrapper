# GenerateSeedPhraseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NWords** | Pointer to **float32** | The number of words to generate. Must be one of: 12, 15, 18, 21, or 24. Avalanche&#39;s default is 24. All other chains we support default to 12.  | [optional] 

## Methods

### NewGenerateSeedPhraseRequest

`func NewGenerateSeedPhraseRequest() *GenerateSeedPhraseRequest`

NewGenerateSeedPhraseRequest instantiates a new GenerateSeedPhraseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSeedPhraseRequestWithDefaults

`func NewGenerateSeedPhraseRequestWithDefaults() *GenerateSeedPhraseRequest`

NewGenerateSeedPhraseRequestWithDefaults instantiates a new GenerateSeedPhraseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNWords

`func (o *GenerateSeedPhraseRequest) GetNWords() float32`

GetNWords returns the NWords field if non-nil, zero value otherwise.

### GetNWordsOk

`func (o *GenerateSeedPhraseRequest) GetNWordsOk() (*float32, bool)`

GetNWordsOk returns a tuple with the NWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNWords

`func (o *GenerateSeedPhraseRequest) SetNWords(v float32)`

SetNWords sets NWords field to given value.

### HasNWords

`func (o *GenerateSeedPhraseRequest) HasNWords() bool`

HasNWords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


