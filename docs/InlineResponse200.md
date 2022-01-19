# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | Pointer to **string** | The ID of the exchange where the NFT is listed  | [optional] 
**ExchangeReadable** | Pointer to **string** | A readable version of the exchange ID  | [optional] 
**ListingTransactionSignature** | Pointer to **string** | The signature of the listing transaction  | [optional] 
**MintAddress** | Pointer to **string** | The mint address of the NFT  | [optional] 
**Price** | Pointer to **float32** | The price of the NFT in Lamports. Represented as an integer. | [optional] 

## Methods

### NewInlineResponse200

`func NewInlineResponse200() *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *InlineResponse200) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InlineResponse200) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InlineResponse200) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InlineResponse200) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeReadable

`func (o *InlineResponse200) GetExchangeReadable() string`

GetExchangeReadable returns the ExchangeReadable field if non-nil, zero value otherwise.

### GetExchangeReadableOk

`func (o *InlineResponse200) GetExchangeReadableOk() (*string, bool)`

GetExchangeReadableOk returns a tuple with the ExchangeReadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeReadable

`func (o *InlineResponse200) SetExchangeReadable(v string)`

SetExchangeReadable sets ExchangeReadable field to given value.

### HasExchangeReadable

`func (o *InlineResponse200) HasExchangeReadable() bool`

HasExchangeReadable returns a boolean if a field has been set.

### GetListingTransactionSignature

`func (o *InlineResponse200) GetListingTransactionSignature() string`

GetListingTransactionSignature returns the ListingTransactionSignature field if non-nil, zero value otherwise.

### GetListingTransactionSignatureOk

`func (o *InlineResponse200) GetListingTransactionSignatureOk() (*string, bool)`

GetListingTransactionSignatureOk returns a tuple with the ListingTransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingTransactionSignature

`func (o *InlineResponse200) SetListingTransactionSignature(v string)`

SetListingTransactionSignature sets ListingTransactionSignature field to given value.

### HasListingTransactionSignature

`func (o *InlineResponse200) HasListingTransactionSignature() bool`

HasListingTransactionSignature returns a boolean if a field has been set.

### GetMintAddress

`func (o *InlineResponse200) GetMintAddress() string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *InlineResponse200) GetMintAddressOk() (*string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *InlineResponse200) SetMintAddress(v string)`

SetMintAddress sets MintAddress field to given value.

### HasMintAddress

`func (o *InlineResponse200) HasMintAddress() bool`

HasMintAddress returns a boolean if a field has been set.

### GetPrice

`func (o *InlineResponse200) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InlineResponse200) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InlineResponse200) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InlineResponse200) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


