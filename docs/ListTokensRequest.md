# ListTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeNfts** | Pointer to **bool** | Whether or not to include NFTs in the response | [optional] [default to false]
**IncludeZeroBalanceHoldings** | Pointer to **bool** | Whether or not to include holdings that have zero balance. This indicates that the wallet held this token or NFT in the past, but no longer holds it. | [optional] [default to false]

## Methods

### NewListTokensRequest

`func NewListTokensRequest() *ListTokensRequest`

NewListTokensRequest instantiates a new ListTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTokensRequestWithDefaults

`func NewListTokensRequestWithDefaults() *ListTokensRequest`

NewListTokensRequestWithDefaults instantiates a new ListTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeNfts

`func (o *ListTokensRequest) GetIncludeNfts() bool`

GetIncludeNfts returns the IncludeNfts field if non-nil, zero value otherwise.

### GetIncludeNftsOk

`func (o *ListTokensRequest) GetIncludeNftsOk() (*bool, bool)`

GetIncludeNftsOk returns a tuple with the IncludeNfts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNfts

`func (o *ListTokensRequest) SetIncludeNfts(v bool)`

SetIncludeNfts sets IncludeNfts field to given value.

### HasIncludeNfts

`func (o *ListTokensRequest) HasIncludeNfts() bool`

HasIncludeNfts returns a boolean if a field has been set.

### GetIncludeZeroBalanceHoldings

`func (o *ListTokensRequest) GetIncludeZeroBalanceHoldings() bool`

GetIncludeZeroBalanceHoldings returns the IncludeZeroBalanceHoldings field if non-nil, zero value otherwise.

### GetIncludeZeroBalanceHoldingsOk

`func (o *ListTokensRequest) GetIncludeZeroBalanceHoldingsOk() (*bool, bool)`

GetIncludeZeroBalanceHoldingsOk returns a tuple with the IncludeZeroBalanceHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZeroBalanceHoldings

`func (o *ListTokensRequest) SetIncludeZeroBalanceHoldings(v bool)`

SetIncludeZeroBalanceHoldings sets IncludeZeroBalanceHoldings field to given value.

### HasIncludeZeroBalanceHoldings

`func (o *ListTokensRequest) HasIncludeZeroBalanceHoldings() bool`

HasIncludeZeroBalanceHoldings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


