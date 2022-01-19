# GetNFTListingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | Pointer to **string** | The ID of the exchange where the NFT is listed  | [optional] 
**ExchangeReadable** | Pointer to **string** | A readable version of the exchange ID  | [optional] 
**ListingTransactionSignature** | Pointer to **string** | The signature of the listing transaction  | [optional] 
**MintAddress** | Pointer to **string** | The mint address of the NFT  | [optional] 
**Price** | Pointer to **float32** | The price of the NFT in Lamports. Represented as an integer. | [optional] 

## Methods

### NewGetNFTListingResponse

`func NewGetNFTListingResponse() *GetNFTListingResponse`

NewGetNFTListingResponse instantiates a new GetNFTListingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNFTListingResponseWithDefaults

`func NewGetNFTListingResponseWithDefaults() *GetNFTListingResponse`

NewGetNFTListingResponseWithDefaults instantiates a new GetNFTListingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *GetNFTListingResponse) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetNFTListingResponse) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetNFTListingResponse) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *GetNFTListingResponse) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeReadable

`func (o *GetNFTListingResponse) GetExchangeReadable() string`

GetExchangeReadable returns the ExchangeReadable field if non-nil, zero value otherwise.

### GetExchangeReadableOk

`func (o *GetNFTListingResponse) GetExchangeReadableOk() (*string, bool)`

GetExchangeReadableOk returns a tuple with the ExchangeReadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeReadable

`func (o *GetNFTListingResponse) SetExchangeReadable(v string)`

SetExchangeReadable sets ExchangeReadable field to given value.

### HasExchangeReadable

`func (o *GetNFTListingResponse) HasExchangeReadable() bool`

HasExchangeReadable returns a boolean if a field has been set.

### GetListingTransactionSignature

`func (o *GetNFTListingResponse) GetListingTransactionSignature() string`

GetListingTransactionSignature returns the ListingTransactionSignature field if non-nil, zero value otherwise.

### GetListingTransactionSignatureOk

`func (o *GetNFTListingResponse) GetListingTransactionSignatureOk() (*string, bool)`

GetListingTransactionSignatureOk returns a tuple with the ListingTransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingTransactionSignature

`func (o *GetNFTListingResponse) SetListingTransactionSignature(v string)`

SetListingTransactionSignature sets ListingTransactionSignature field to given value.

### HasListingTransactionSignature

`func (o *GetNFTListingResponse) HasListingTransactionSignature() bool`

HasListingTransactionSignature returns a boolean if a field has been set.

### GetMintAddress

`func (o *GetNFTListingResponse) GetMintAddress() string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *GetNFTListingResponse) GetMintAddressOk() (*string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *GetNFTListingResponse) SetMintAddress(v string)`

SetMintAddress sets MintAddress field to given value.

### HasMintAddress

`func (o *GetNFTListingResponse) HasMintAddress() bool`

HasMintAddress returns a boolean if a field has been set.

### GetPrice

`func (o *GetNFTListingResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetNFTListingResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetNFTListingResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetNFTListingResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


