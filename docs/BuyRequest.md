# BuyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | [**Wallet**](Wallet.md) |  | 
**SkipChecks** | Pointer to **bool** | Whether or not to skip the provided checks (e.g., Is this NFT not listed? Is this NFT listed for a different price than you set?) and proceed with the transaction.  | [optional] [default to false]
**SellerPublicKey** | Pointer to **string** | The public key of the seller. Only required if providing &#x60;skip_checks&#x60;. Otherwise, don&#39;t provide it.  | [optional] [default to "null"]
**NftPrice** | **float32** | The number of lamports you are expecting to purchase the NFT for. We check the price of the NFT before  purchasing it to ensure that it matches your expectation. There are 1e9 (1 billion) Lamports in a SOL.  | 

## Methods

### NewBuyRequest

`func NewBuyRequest(wallet Wallet, nftPrice float32, ) *BuyRequest`

NewBuyRequest instantiates a new BuyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyRequestWithDefaults

`func NewBuyRequestWithDefaults() *BuyRequest`

NewBuyRequestWithDefaults instantiates a new BuyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallet

`func (o *BuyRequest) GetWallet() Wallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *BuyRequest) GetWalletOk() (*Wallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *BuyRequest) SetWallet(v Wallet)`

SetWallet sets Wallet field to given value.


### GetSkipChecks

`func (o *BuyRequest) GetSkipChecks() bool`

GetSkipChecks returns the SkipChecks field if non-nil, zero value otherwise.

### GetSkipChecksOk

`func (o *BuyRequest) GetSkipChecksOk() (*bool, bool)`

GetSkipChecksOk returns a tuple with the SkipChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipChecks

`func (o *BuyRequest) SetSkipChecks(v bool)`

SetSkipChecks sets SkipChecks field to given value.

### HasSkipChecks

`func (o *BuyRequest) HasSkipChecks() bool`

HasSkipChecks returns a boolean if a field has been set.

### GetSellerPublicKey

`func (o *BuyRequest) GetSellerPublicKey() string`

GetSellerPublicKey returns the SellerPublicKey field if non-nil, zero value otherwise.

### GetSellerPublicKeyOk

`func (o *BuyRequest) GetSellerPublicKeyOk() (*string, bool)`

GetSellerPublicKeyOk returns a tuple with the SellerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerPublicKey

`func (o *BuyRequest) SetSellerPublicKey(v string)`

SetSellerPublicKey sets SellerPublicKey field to given value.

### HasSellerPublicKey

`func (o *BuyRequest) HasSellerPublicKey() bool`

HasSellerPublicKey returns a boolean if a field has been set.

### GetNftPrice

`func (o *BuyRequest) GetNftPrice() float32`

GetNftPrice returns the NftPrice field if non-nil, zero value otherwise.

### GetNftPriceOk

`func (o *BuyRequest) GetNftPriceOk() (*float32, bool)`

GetNftPriceOk returns a tuple with the NftPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftPrice

`func (o *BuyRequest) SetNftPrice(v float32)`

SetNftPrice sets NftPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


