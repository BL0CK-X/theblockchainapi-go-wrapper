# NFTMintRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | [**Wallet**](Wallet.md) |  | 
**NftName** | Pointer to **string** | The name of the token. Limited to 32 characters. Stored on the blockchain. | [optional] [default to ""]
**NftSymbol** | Pointer to **string** | The symbol of the token. Limited to 10 characters. Stored on the blockchain. | [optional] [default to ""]
**NftDescription** | Pointer to **string** | The description of the token. Limited to 2000 characters. Not stored on the blockchain.  This is stored in S3 in a bucket we own, and the link to that file is stored on the blockchain.  If you provide your own link, the link is also stored in that S3 file, which is publicly accessible. However, if you choose the NFT upload method \&quot;LINK\&quot; instead of \&quot;S3\&quot;, then we upload the link you  provide for nft_url directly to the blockchain, and S3 is not used at all. Thus, when you provide the \&quot;LINK\&quot; option, the value nft_description is ignored and not used. The Metaplex API does not provide functionality to store any data about your NFT besides a URL.  | [optional] [default to ""]
**NftUrl** | Pointer to **string** | The URL you provided. Limited to 200 characters. This does not need to be a valid URL. Whether or not this is  stored on the blockchain depends on which upload method you choose. If you choose LINK, then this is stored on the  blockchain directly. If you choose S3, then this link is embedded in a public S3 text file that also contains the metadata, the name,  the symbol, and the description of the NFT.  | [optional] [default to ""]
**NftMetadata** | Pointer to **string** | Any data you provide. Must be a string and properly encoded json. Will be viewable on S3. Limited to 2000 bytes. Not stored on the blockchain.  This is stored in S3 in a bucket we own, and the link to that file is stored on the blockchain.  If you provide your own link, the link is also stored in that S3 file, which is publicly accessible. However, if you choose the NFT upload method \&quot;LINK\&quot; instead of \&quot;S3\&quot;, then we upload the link you  provide for nft_url directly to the blockchain, and S3 is not used at all. Thus, when you provide the \&quot;LINK\&quot; option, the value nft_metadata is ignored and not used. The Metaplex API does not provide functionality to store any data about your NFT besides a URL.  | [optional] [default to "{}"]
**NftUploadMethod** | Pointer to **string** | When you choose S3, all of the nft_description, nft_name, nft_symbol, nft_metadata, and nft_url are put into a json dictionary and uploaded to S3 as a text file.  This is uploaded to an AWS S3 bucket we own, and is an option we provide at no charge. The S3 link to this file is stored on the NFT on the blockchain.   When you choose LINK, the nft_url you provide is stored on the blockchain, and the nft_metadata and nft_description are ignored and not stored anywhere. S3 is then NOT used.  | [optional] [default to "S3"]
**IsMutable** | Pointer to **bool** | Indicates whether or not the NFT created is mutable. If mutable, the NFT can be updated later. Once set to immutable, the NFT is unable to be changed.  | [optional] [default to true]
**IsMasterEdition** | Pointer to **bool** | Whether or not the NFT is a master edition NFT. Saves about 0.001 SOL in transaction costs when set to false.  | [optional] [default to true]
**SellerFeeBasisPoints** | Pointer to **float32** | Valid values from 0 to 10000. Must be an integer.  Represents the number of basis points that the seller receives as a fee upon sale.  E.g., 100 indicates a 1% seller fee. Seller does not receive a fee when \&quot;primary_sale_has_happened\&quot; is set to true.  Will be set to false after first sale has occurred.  | [optional] [default to 0]
**Creators** | Pointer to **[]string** | A JSON encoded string representing an array / list.  The designated creators of the NFT. Length of the creator list must match length of the list of share.  Valid lengths of the list range from 1 to 5. Each item in the list must be a valid public key address.    Only the public key corresponding to the seed phrase provided will be marked as verified. Any other creators supplied will be marked as unverified.  | [optional] [default to ["The Public Key Corresponding to The Seed Phrase, Path, and Passphrase Provided"]]
**Share** | Pointer to **[]int32** | A JSON encoded string representing an array / list.  The share of the royalty that each creator gets. Valid values range from 0 to 100.  Sum of the values must equal 100.  Only integer value accepted. Length of the share list must match length of the list of creators.  | [optional] [default to [100]]
**MintToPublicKey** | Pointer to **string** | Assign ownership of the NFT to the public key address given by &#x60;mint_to_public_key&#x60;  | [optional] [default to "The public key of the wallet provided"]
**Network** | Pointer to **string** | This determines which network you choose to run the API calls on. We recommend first testing on the devnet, because minting an NFT costs a little above 0.01 SOL, which is about $1.60 at the time of this writing.  When you run on the mainnet-beta, each successful call will deduct approximately that much. When you run on the devnet, that amount is deducted from a simulated amount, so you are not paying with real SOL. To get SOL on the devnet,   airdrop SOL to this address using the CLI. Keep in mind that you can only do this every so often. If you are rate-limited, consider using a VPN and trying again, or just waiting. To get SOL on the mainnet-beta, you    must transfer real SOL to this account from another wallet (e.g., from another wallet you own, from an exchange, etc.). We hope to make this process easier in the future, and if you have any suggestions, please add them    as an issue on our &lt;a href&#x3D;\&quot;https://github.com/BL0CK-X/the-blockchain-api\&quot; target&#x3D;\&quot;_blank\&quot;&gt;GitHub repository&lt;/a&gt; for the API. To get a fee estimate, make a GET requests to the &lt;a href&#x3D;\&quot;#tag/Solana-NFT/paths/~1solana~1nft~1mint~1fee/get\&quot;&gt;v1/solana/nft/mint/fee endpoint&lt;/a&gt; (details in sidebar).  | [optional] [default to "devnet"]

## Methods

### NewNFTMintRequest

`func NewNFTMintRequest(wallet Wallet, ) *NFTMintRequest`

NewNFTMintRequest instantiates a new NFTMintRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTMintRequestWithDefaults

`func NewNFTMintRequestWithDefaults() *NFTMintRequest`

NewNFTMintRequestWithDefaults instantiates a new NFTMintRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallet

`func (o *NFTMintRequest) GetWallet() Wallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *NFTMintRequest) GetWalletOk() (*Wallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *NFTMintRequest) SetWallet(v Wallet)`

SetWallet sets Wallet field to given value.


### GetNftName

`func (o *NFTMintRequest) GetNftName() string`

GetNftName returns the NftName field if non-nil, zero value otherwise.

### GetNftNameOk

`func (o *NFTMintRequest) GetNftNameOk() (*string, bool)`

GetNftNameOk returns a tuple with the NftName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftName

`func (o *NFTMintRequest) SetNftName(v string)`

SetNftName sets NftName field to given value.

### HasNftName

`func (o *NFTMintRequest) HasNftName() bool`

HasNftName returns a boolean if a field has been set.

### GetNftSymbol

`func (o *NFTMintRequest) GetNftSymbol() string`

GetNftSymbol returns the NftSymbol field if non-nil, zero value otherwise.

### GetNftSymbolOk

`func (o *NFTMintRequest) GetNftSymbolOk() (*string, bool)`

GetNftSymbolOk returns a tuple with the NftSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftSymbol

`func (o *NFTMintRequest) SetNftSymbol(v string)`

SetNftSymbol sets NftSymbol field to given value.

### HasNftSymbol

`func (o *NFTMintRequest) HasNftSymbol() bool`

HasNftSymbol returns a boolean if a field has been set.

### GetNftDescription

`func (o *NFTMintRequest) GetNftDescription() string`

GetNftDescription returns the NftDescription field if non-nil, zero value otherwise.

### GetNftDescriptionOk

`func (o *NFTMintRequest) GetNftDescriptionOk() (*string, bool)`

GetNftDescriptionOk returns a tuple with the NftDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftDescription

`func (o *NFTMintRequest) SetNftDescription(v string)`

SetNftDescription sets NftDescription field to given value.

### HasNftDescription

`func (o *NFTMintRequest) HasNftDescription() bool`

HasNftDescription returns a boolean if a field has been set.

### GetNftUrl

`func (o *NFTMintRequest) GetNftUrl() string`

GetNftUrl returns the NftUrl field if non-nil, zero value otherwise.

### GetNftUrlOk

`func (o *NFTMintRequest) GetNftUrlOk() (*string, bool)`

GetNftUrlOk returns a tuple with the NftUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftUrl

`func (o *NFTMintRequest) SetNftUrl(v string)`

SetNftUrl sets NftUrl field to given value.

### HasNftUrl

`func (o *NFTMintRequest) HasNftUrl() bool`

HasNftUrl returns a boolean if a field has been set.

### GetNftMetadata

`func (o *NFTMintRequest) GetNftMetadata() string`

GetNftMetadata returns the NftMetadata field if non-nil, zero value otherwise.

### GetNftMetadataOk

`func (o *NFTMintRequest) GetNftMetadataOk() (*string, bool)`

GetNftMetadataOk returns a tuple with the NftMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftMetadata

`func (o *NFTMintRequest) SetNftMetadata(v string)`

SetNftMetadata sets NftMetadata field to given value.

### HasNftMetadata

`func (o *NFTMintRequest) HasNftMetadata() bool`

HasNftMetadata returns a boolean if a field has been set.

### GetNftUploadMethod

`func (o *NFTMintRequest) GetNftUploadMethod() string`

GetNftUploadMethod returns the NftUploadMethod field if non-nil, zero value otherwise.

### GetNftUploadMethodOk

`func (o *NFTMintRequest) GetNftUploadMethodOk() (*string, bool)`

GetNftUploadMethodOk returns a tuple with the NftUploadMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftUploadMethod

`func (o *NFTMintRequest) SetNftUploadMethod(v string)`

SetNftUploadMethod sets NftUploadMethod field to given value.

### HasNftUploadMethod

`func (o *NFTMintRequest) HasNftUploadMethod() bool`

HasNftUploadMethod returns a boolean if a field has been set.

### GetIsMutable

`func (o *NFTMintRequest) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *NFTMintRequest) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *NFTMintRequest) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *NFTMintRequest) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetIsMasterEdition

`func (o *NFTMintRequest) GetIsMasterEdition() bool`

GetIsMasterEdition returns the IsMasterEdition field if non-nil, zero value otherwise.

### GetIsMasterEditionOk

`func (o *NFTMintRequest) GetIsMasterEditionOk() (*bool, bool)`

GetIsMasterEditionOk returns a tuple with the IsMasterEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterEdition

`func (o *NFTMintRequest) SetIsMasterEdition(v bool)`

SetIsMasterEdition sets IsMasterEdition field to given value.

### HasIsMasterEdition

`func (o *NFTMintRequest) HasIsMasterEdition() bool`

HasIsMasterEdition returns a boolean if a field has been set.

### GetSellerFeeBasisPoints

`func (o *NFTMintRequest) GetSellerFeeBasisPoints() float32`

GetSellerFeeBasisPoints returns the SellerFeeBasisPoints field if non-nil, zero value otherwise.

### GetSellerFeeBasisPointsOk

`func (o *NFTMintRequest) GetSellerFeeBasisPointsOk() (*float32, bool)`

GetSellerFeeBasisPointsOk returns a tuple with the SellerFeeBasisPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerFeeBasisPoints

`func (o *NFTMintRequest) SetSellerFeeBasisPoints(v float32)`

SetSellerFeeBasisPoints sets SellerFeeBasisPoints field to given value.

### HasSellerFeeBasisPoints

`func (o *NFTMintRequest) HasSellerFeeBasisPoints() bool`

HasSellerFeeBasisPoints returns a boolean if a field has been set.

### GetCreators

`func (o *NFTMintRequest) GetCreators() []string`

GetCreators returns the Creators field if non-nil, zero value otherwise.

### GetCreatorsOk

`func (o *NFTMintRequest) GetCreatorsOk() (*[]string, bool)`

GetCreatorsOk returns a tuple with the Creators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreators

`func (o *NFTMintRequest) SetCreators(v []string)`

SetCreators sets Creators field to given value.

### HasCreators

`func (o *NFTMintRequest) HasCreators() bool`

HasCreators returns a boolean if a field has been set.

### GetShare

`func (o *NFTMintRequest) GetShare() []int32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *NFTMintRequest) GetShareOk() (*[]int32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *NFTMintRequest) SetShare(v []int32)`

SetShare sets Share field to given value.

### HasShare

`func (o *NFTMintRequest) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetMintToPublicKey

`func (o *NFTMintRequest) GetMintToPublicKey() string`

GetMintToPublicKey returns the MintToPublicKey field if non-nil, zero value otherwise.

### GetMintToPublicKeyOk

`func (o *NFTMintRequest) GetMintToPublicKeyOk() (*string, bool)`

GetMintToPublicKeyOk returns a tuple with the MintToPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintToPublicKey

`func (o *NFTMintRequest) SetMintToPublicKey(v string)`

SetMintToPublicKey sets MintToPublicKey field to given value.

### HasMintToPublicKey

`func (o *NFTMintRequest) HasMintToPublicKey() bool`

HasMintToPublicKey returns a boolean if a field has been set.

### GetNetwork

`func (o *NFTMintRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NFTMintRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NFTMintRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NFTMintRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


