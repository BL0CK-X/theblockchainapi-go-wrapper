# BalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** | The public key address of the wallet | 
**Unit** | Pointer to **string** | If you are retrieving the SOL balance, you can select whether to retrieve this in SOL or Lamport units (1 SOL &#x3D; 1e9 Lamports). | [optional] [default to "lamport"]
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]
**MintAddress** | Pointer to **string** | The mint address of the SPL token or NFT. If not provided, the balance returned is in SOL. Make sure to use the correct network.  You can see the mint addresses of popular SPL tokens &lt;a href&#x3D;\&quot;https://raw.githubusercontent.com/solana-labs/token-list/main/src/tokens/solana.tokenlist.json\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt;.  Some example mint addresses of SPL tokens: - USDC: EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v - Mango: MangoCzJ36AjZyKwVj3VnYU4GTonjfVEnJmvvWaxLac - Serum: SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt - Raydium: 4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R - wSOL: So11111111111111111111111111111111111111112 - ATLAS: ATLASXmbPQxBUYbxPsV97usA3fPQYEqzQBUHgiFCUsXx | [optional] [default to "null"]

## Methods

### NewBalanceRequest

`func NewBalanceRequest(publicKey string, ) *BalanceRequest`

NewBalanceRequest instantiates a new BalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceRequestWithDefaults

`func NewBalanceRequestWithDefaults() *BalanceRequest`

NewBalanceRequestWithDefaults instantiates a new BalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *BalanceRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *BalanceRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *BalanceRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetUnit

`func (o *BalanceRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BalanceRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BalanceRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BalanceRequest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetNetwork

`func (o *BalanceRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BalanceRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BalanceRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *BalanceRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetMintAddress

`func (o *BalanceRequest) GetMintAddress() string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *BalanceRequest) GetMintAddressOk() (*string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *BalanceRequest) SetMintAddress(v string)`

SetMintAddress sets MintAddress field to given value.

### HasMintAddress

`func (o *BalanceRequest) HasMintAddress() bool`

HasMintAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


