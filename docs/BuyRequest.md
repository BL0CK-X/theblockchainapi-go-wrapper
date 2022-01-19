# BuyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | [**Wallet**](Wallet.md) |  | 
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


