# NFTCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verified** | Pointer to **float32** | A &#x60;0&#x60; or &#x60;1&#x60; indicating whether or not the associated collection is verified.  | [optional] 
**Key** | Pointer to **string** | Will return &#x60;11111111111111111111111111111111&#x60; if no associated &#x60;collection_id&#x60;. Otherwise it will return the &#x60;collection_id&#x60;, which is a public key representing the verified Metaplex collection. See more [here](https://collections.metaplex.com).  | [optional] 

## Methods

### NewNFTCollection

`func NewNFTCollection() *NFTCollection`

NewNFTCollection instantiates a new NFTCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTCollectionWithDefaults

`func NewNFTCollectionWithDefaults() *NFTCollection`

NewNFTCollectionWithDefaults instantiates a new NFTCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerified

`func (o *NFTCollection) GetVerified() float32`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *NFTCollection) GetVerifiedOk() (*float32, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *NFTCollection) SetVerified(v float32)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *NFTCollection) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetKey

`func (o *NFTCollection) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NFTCollection) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NFTCollection) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *NFTCollection) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


