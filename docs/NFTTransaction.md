# NFTTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTime** | Pointer to **float32** | An epoch time stamp in UTC time that represents the time of the block where the transaction was processed  | [optional] 
**Exchange** | Pointer to **string** | The NFT exchange on which the transaction occurred | [optional] 
**ExchangeReadable** | Pointer to **string** | A readable version of the NFT exchange | [optional] 
**MintAddress** | Pointer to **string** | The mint address of the NFT  | [optional] 
**Operation** | Pointer to **string** | The operation of the transaction | [optional] 
**Seller** | Pointer to **string** | The public key of the wallet that listed the NFT | [optional] 
**Buyer** | Pointer to **string** | The public key of the buyer. This only exists in &#x60;buy&#x60; transactions.  | [optional] 
**TransactionSignature** | Pointer to **string** | The signature of the transaction. You can look up each transaction on explorer.solana.com  | [optional] 

## Methods

### NewNFTTransaction

`func NewNFTTransaction() *NFTTransaction`

NewNFTTransaction instantiates a new NFTTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTTransactionWithDefaults

`func NewNFTTransactionWithDefaults() *NFTTransaction`

NewNFTTransactionWithDefaults instantiates a new NFTTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTime

`func (o *NFTTransaction) GetBlockTime() float32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *NFTTransaction) GetBlockTimeOk() (*float32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *NFTTransaction) SetBlockTime(v float32)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *NFTTransaction) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetExchange

`func (o *NFTTransaction) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *NFTTransaction) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *NFTTransaction) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *NFTTransaction) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeReadable

`func (o *NFTTransaction) GetExchangeReadable() string`

GetExchangeReadable returns the ExchangeReadable field if non-nil, zero value otherwise.

### GetExchangeReadableOk

`func (o *NFTTransaction) GetExchangeReadableOk() (*string, bool)`

GetExchangeReadableOk returns a tuple with the ExchangeReadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeReadable

`func (o *NFTTransaction) SetExchangeReadable(v string)`

SetExchangeReadable sets ExchangeReadable field to given value.

### HasExchangeReadable

`func (o *NFTTransaction) HasExchangeReadable() bool`

HasExchangeReadable returns a boolean if a field has been set.

### GetMintAddress

`func (o *NFTTransaction) GetMintAddress() string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *NFTTransaction) GetMintAddressOk() (*string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *NFTTransaction) SetMintAddress(v string)`

SetMintAddress sets MintAddress field to given value.

### HasMintAddress

`func (o *NFTTransaction) HasMintAddress() bool`

HasMintAddress returns a boolean if a field has been set.

### GetOperation

`func (o *NFTTransaction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *NFTTransaction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *NFTTransaction) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *NFTTransaction) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetSeller

`func (o *NFTTransaction) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *NFTTransaction) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *NFTTransaction) SetSeller(v string)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *NFTTransaction) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetBuyer

`func (o *NFTTransaction) GetBuyer() string`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *NFTTransaction) GetBuyerOk() (*string, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *NFTTransaction) SetBuyer(v string)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *NFTTransaction) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetTransactionSignature

`func (o *NFTTransaction) GetTransactionSignature() string`

GetTransactionSignature returns the TransactionSignature field if non-nil, zero value otherwise.

### GetTransactionSignatureOk

`func (o *NFTTransaction) GetTransactionSignatureOk() (*string, bool)`

GetTransactionSignatureOk returns a tuple with the TransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSignature

`func (o *NFTTransaction) SetTransactionSignature(v string)`

SetTransactionSignature sets TransactionSignature field to given value.

### HasTransactionSignature

`func (o *NFTTransaction) HasTransactionSignature() bool`

HasTransactionSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


