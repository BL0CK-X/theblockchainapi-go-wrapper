# GetConfigInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigAddress** | **string** | The config address of the candy machine.  If you don&#39;t know the address of a candy machine, you can use the \&quot;Get Candy Machine Details\&quot; endpoint to get the configuration public key address.  If you don&#39;t know the candy machine ID to call this endpoint, you can find it with the \&quot;Get NFT&#39;s Candy Machine ID\&quot; endpoint.  | 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]

## Methods

### NewGetConfigInfoRequest

`func NewGetConfigInfoRequest(configAddress string, ) *GetConfigInfoRequest`

NewGetConfigInfoRequest instantiates a new GetConfigInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfigInfoRequestWithDefaults

`func NewGetConfigInfoRequestWithDefaults() *GetConfigInfoRequest`

NewGetConfigInfoRequestWithDefaults instantiates a new GetConfigInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigAddress

`func (o *GetConfigInfoRequest) GetConfigAddress() string`

GetConfigAddress returns the ConfigAddress field if non-nil, zero value otherwise.

### GetConfigAddressOk

`func (o *GetConfigInfoRequest) GetConfigAddressOk() (*string, bool)`

GetConfigAddressOk returns a tuple with the ConfigAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAddress

`func (o *GetConfigInfoRequest) SetConfigAddress(v string)`

SetConfigAddress sets ConfigAddress field to given value.


### GetNetwork

`func (o *GetConfigInfoRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetConfigInfoRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetConfigInfoRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetConfigInfoRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


