# BalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **float32** | The balance of the token in the wallet | 
**IntegerBalance** | Pointer to **float32** | Not included if retreiving SOL balance | [optional] 
**Decimals** | Pointer to **float32** | Not included if retreiving SOL balance | [optional] 
**Network** | **string** |  | 
**Unit** | Pointer to **string** | Not included if retreiving an SPL token balance | [optional] 

## Methods

### NewBalanceResponse

`func NewBalanceResponse(balance float32, network string, ) *BalanceResponse`

NewBalanceResponse instantiates a new BalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceResponseWithDefaults

`func NewBalanceResponseWithDefaults() *BalanceResponse`

NewBalanceResponseWithDefaults instantiates a new BalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BalanceResponse) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceResponse) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceResponse) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetIntegerBalance

`func (o *BalanceResponse) GetIntegerBalance() float32`

GetIntegerBalance returns the IntegerBalance field if non-nil, zero value otherwise.

### GetIntegerBalanceOk

`func (o *BalanceResponse) GetIntegerBalanceOk() (*float32, bool)`

GetIntegerBalanceOk returns a tuple with the IntegerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegerBalance

`func (o *BalanceResponse) SetIntegerBalance(v float32)`

SetIntegerBalance sets IntegerBalance field to given value.

### HasIntegerBalance

`func (o *BalanceResponse) HasIntegerBalance() bool`

HasIntegerBalance returns a boolean if a field has been set.

### GetDecimals

`func (o *BalanceResponse) GetDecimals() float32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *BalanceResponse) GetDecimalsOk() (*float32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *BalanceResponse) SetDecimals(v float32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *BalanceResponse) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetNetwork

`func (o *BalanceResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BalanceResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BalanceResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetUnit

`func (o *BalanceResponse) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BalanceResponse) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BalanceResponse) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BalanceResponse) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


