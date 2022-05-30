# GeneralBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **float32** | The balance of the token in the wallet  | 
**IntegerBalance** | Pointer to **float32** | Not included if retreiving native token (e.g., SOL, ETH, BNB, etc.) balance  | [optional] 
**Decimals** | Pointer to **float32** | Not included if retreiving native token (e.g., SOL, ETH, BNB, etc.) balance.  | [optional] 
**Network** | **string** | The network of the blockchain you selected  - Solana: &#x60;devnet&#x60;, &#x60;mainnet-beta&#x60; - Ethereum: &#x60;ropsten&#x60;, &#x60;mainnet&#x60;  Defaults when not provided (not applicable to path parameters): - Solana: &#x60;devnet&#x60; - Ethereum: &#x60;ropsten&#x60; | 
**Unit** | Pointer to **string** | Not included if retreiving a token / NFT balance  | [optional] 

## Methods

### NewGeneralBalanceResponse

`func NewGeneralBalanceResponse(balance float32, network string, ) *GeneralBalanceResponse`

NewGeneralBalanceResponse instantiates a new GeneralBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralBalanceResponseWithDefaults

`func NewGeneralBalanceResponseWithDefaults() *GeneralBalanceResponse`

NewGeneralBalanceResponseWithDefaults instantiates a new GeneralBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *GeneralBalanceResponse) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GeneralBalanceResponse) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GeneralBalanceResponse) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetIntegerBalance

`func (o *GeneralBalanceResponse) GetIntegerBalance() float32`

GetIntegerBalance returns the IntegerBalance field if non-nil, zero value otherwise.

### GetIntegerBalanceOk

`func (o *GeneralBalanceResponse) GetIntegerBalanceOk() (*float32, bool)`

GetIntegerBalanceOk returns a tuple with the IntegerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegerBalance

`func (o *GeneralBalanceResponse) SetIntegerBalance(v float32)`

SetIntegerBalance sets IntegerBalance field to given value.

### HasIntegerBalance

`func (o *GeneralBalanceResponse) HasIntegerBalance() bool`

HasIntegerBalance returns a boolean if a field has been set.

### GetDecimals

`func (o *GeneralBalanceResponse) GetDecimals() float32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *GeneralBalanceResponse) GetDecimalsOk() (*float32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *GeneralBalanceResponse) SetDecimals(v float32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *GeneralBalanceResponse) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetNetwork

`func (o *GeneralBalanceResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GeneralBalanceResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GeneralBalanceResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetUnit

`func (o *GeneralBalanceResponse) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GeneralBalanceResponse) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GeneralBalanceResponse) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *GeneralBalanceResponse) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


