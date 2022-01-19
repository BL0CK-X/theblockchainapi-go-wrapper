# SolSeaListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | [**Wallet**](Wallet.md) |  | 
**NftPrice** | **float32** |  The number of lamports you are expecting to sell the NFT for.  There are 1e9 (1 billion) Lamports in a SOL.    Min price: 1   Max price: 18446744073709551615  | 

## Methods

### NewSolSeaListRequest

`func NewSolSeaListRequest(wallet Wallet, nftPrice float32, ) *SolSeaListRequest`

NewSolSeaListRequest instantiates a new SolSeaListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolSeaListRequestWithDefaults

`func NewSolSeaListRequestWithDefaults() *SolSeaListRequest`

NewSolSeaListRequestWithDefaults instantiates a new SolSeaListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallet

`func (o *SolSeaListRequest) GetWallet() Wallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *SolSeaListRequest) GetWalletOk() (*Wallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *SolSeaListRequest) SetWallet(v Wallet)`

SetWallet sets Wallet field to given value.


### GetNftPrice

`func (o *SolSeaListRequest) GetNftPrice() float32`

GetNftPrice returns the NftPrice field if non-nil, zero value otherwise.

### GetNftPriceOk

`func (o *SolSeaListRequest) GetNftPriceOk() (*float32, bool)`

GetNftPriceOk returns a tuple with the NftPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftPrice

`func (o *SolSeaListRequest) SetNftPrice(v float32)`

SetNftPrice sets NftPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


