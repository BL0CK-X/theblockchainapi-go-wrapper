# NFTMintRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | Pointer to [**Wallet**](Wallet.md) |  | [optional] 
**ReturnCompiledTransaction** | Pointer to **bool** | If &#x60;true&#x60;, the transaction to mint the NFT will not be submitted or signed. It will be returned to you in a raw form that you can then sign with a wallet (e.g., Phantom) or code. No &#x60;wallet&#x60; authentication information is required (thus, you do you have to supply a seed phrase or private key). See a Python example [here](https://github.com/BL0CK-X/blockchain-api/blob/main/third-party-api-examples/me-buy-sell.py). If &#x60;false&#x60; (the default option), then &#x60;wallet&#x60; is required. We sign and submit the transaction for you, which uses your wallet to mint the NFT. No further action is required on your part, and the NFT is minted. Read more on security [here](#section/Security).  | [optional] [default to false]
**Name** | Pointer to **string** | The name of the token. Limited to 32 characters. Stored on the blockchain. | [optional] [default to ""]
**Symbol** | Pointer to **string** | The symbol of the token. Limited to 10 characters. Stored on the blockchain. | [optional] [default to ""]
**Description** | Pointer to **string** | The description of the NFT. Limited to 2000 characters. Not stored on the blockchain.         If you are providing your own &#x60;uri&#x60; (see above), then you do not need to provide this.  If you are not providing your own &#x60;uri&#x60; and you do not provide this, then there wills simply be no description.  Only provide a value for &#x60;description&#x60; if the &#x60;upload_method&#x60; is set to &#x60;S3&#x60; (see the description for &#x60;upload_method&#x60; above). | [optional] [default to ""]
**UploadMethod** | Pointer to **string** | When you choose &#x60;S3&#x60;, all of the &#x60;name&#x60;, &#x60;description&#x60;, &#x60;symbol&#x60;, &#x60;uri_metadata&#x60;, and &#x60;image_url&#x60; are put into a JSON dictionary and uploaded to S3 as a JSON file.  This is uploaded to an AWS S3 bucket we own, and is an option we provide at no charge. The S3 link to this file is stored in the NFT&#39;s account on the blockchain. Learn more  &lt;a href&#x3D;\&quot;https://blockchainapi.com/2022/01/18/how-to-format-off-chain-nft-metadata.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt;.  When you choose &#x60;URI&#x60;, the &#x60;uri&#x60; you provide is stored on the blockchain, and the &#x60;uri_metadata&#x60;, &#x60;description&#x60;, and &#x60;image_url&#x60; are ignored and not stored anywhere. &#x60;S3&#x60; is NOT involved in this case.   An example of a &#x60;uri&#x60; you would provide is an Arweave URL, like this: &#x60;https://arweave.net/_Y8tETU3FbAFZSM1wXNeWPweWwrW9K6oSF1SYi_bH9A&#x60;. | [optional] [default to "S3"]
**Uri** | Pointer to **string** | The &#x60;uri&#x60; you provide is stored on the blockchain, and the &#x60;uri_metadata&#x60;, &#x60;description&#x60;, and &#x60;image_url&#x60; are ignored and not stored anywhere. &#x60;S3&#x60; is NOT involved in this case.   Read more &lt;a href&#x3D;\&quot;https://blockchainapi.com/2022/01/18/how-to-format-off-chain-nft-metadata.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt;.  An example of a &#x60;uri&#x60; you would provide is an Arweave URL, like this: &#x60;https://arweave.net/_Y8tETU3FbAFZSM1wXNeWPweWwrW9K6oSF1SYi_bH9A&#x60;.  Only provide a value for &#x60;uri&#x60; if the &#x60;upload_method&#x60; is set to &#x60;URI&#x60; (see the description for &#x60;upload_method&#x60; above). | [optional] [default to ""]
**ImageUrl** | Pointer to **string** | The URL of the image of the NFT.         If you are providing your own &#x60;uri&#x60; (see above), then you do not need to provide this.  If you are not providing your own &#x60;uri&#x60; and you do not provide this, then there wills simply be no image.  Only provide a value for &#x60;image_url&#x60; if the &#x60;upload_method&#x60; is set to &#x60;S3&#x60; (see the description for &#x60;upload_method&#x60; above). | [optional] [default to ""]
**UriMetadata** | Pointer to **map[string]interface{}** | The off-chain metadata.        If you are providing your own &#x60;uri&#x60; (see above), then you do not need to provide this.  If you are not providing your own &#x60;uri&#x60; and you do not provide this, then there wills simply be no image.  Only provide a value for &#x60;uri_metadata&#x60; if the &#x60;upload_method&#x60; is set to &#x60;S3&#x60; (see the description for &#x60;upload_method&#x60; above).  Learn more about how to format this metadata &lt;a href&#x3D;\&quot;https://blockchainapi.com/2022/01/18/how-to-format-off-chain-nft-metadata.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;here&lt;/a&gt;. | [optional] [default to {}]
**IsMutable** | Pointer to **bool** | Indicates whether or not the NFT created is mutable. If mutable, the NFT can be updated later. Once set to immutable, the NFT is unable to be changed.  | [optional] [default to true]
**IsMasterEdition** | Pointer to **bool** | Whether or not the NFT is a master edition NFT. Saves about 0.001 SOL in transaction costs when set to false.  | [optional] [default to true]
**SellerFeeBasisPoints** | Pointer to **float32** | Valid values from 0 to 10000. Must be an integer.  Represents the number of basis points that the seller receives as a fee upon sale.  E.g., 100 indicates a 1% seller fee. Seller does not receive a fee when \&quot;primary_sale_has_happened\&quot; is set to true.  Will be set to false after first sale has occurred.  | [optional] [default to 0]
**Creators** | Pointer to **[]string** | A JSON encoded string representing an array / list.  The designated creators of the NFT. Length of the creator list must match length of the list of share.  Valid lengths of the list range from 1 to 5. Each item in the list must be a valid public key address.    Only the public key corresponding to the seed phrase provided will be marked as verified. Any other creators supplied will be marked as unverified.  | [optional] [default to ["The Public Key Corresponding to The Seed Phrase, Path, and Passphrase Provided"]]
**Share** | Pointer to **[]int32** | A JSON encoded string representing an array / list.  The share of the royalty that each creator gets. Valid values range from 0 to 100. Sum of the values must equal 100.  Only integer value accepted. Length of the share list must match length of the list of creators.  | [optional] [default to [100]]
**MintToPublicKey** | Pointer to **string** | Assign ownership of the NFT to the public key address given by &#x60;mint_to_public_key&#x60;  | [optional] [default to "The public key of the wallet provided"]
**Network** | Pointer to **string** | This determines which network you choose to run the API calls on. We recommend first testing on the devnet, because minting an NFT costs a little above 0.01 SOL, which is about $1.60 at the time of this writing.  When you run on the mainnet-beta, each successful call will deduct approximately that much. When you run on the devnet, that amount is deducted from a simulated amount, so you are not paying with real SOL. To get SOL on the devnet,   airdrop SOL to this address using the CLI. Keep in mind that you can only do this every so often. If you are rate-limited, consider using a VPN and trying again, or just waiting. To get SOL on the mainnet-beta, you    must transfer real SOL to this account from another wallet (e.g., from another wallet you own, from an exchange, etc.). We hope to make this process easier in the future, and if you have any suggestions, please add them    as an issue on our &lt;a href&#x3D;\&quot;https://github.com/BL0CK-X/the-blockchain-api\&quot; target&#x3D;\&quot;_blank\&quot;&gt;GitHub repository&lt;/a&gt; for the API. To get a fee estimate, make a GET requests to the &lt;a href&#x3D;\&quot;#tag/Solana-NFT/paths/~1solana~1nft~1mint~1fee/get\&quot;&gt;v1/solana/nft/mint/fee endpoint&lt;/a&gt; (details in sidebar).  | [optional] [default to "devnet"]

## Methods

### NewNFTMintRequest

`func NewNFTMintRequest() *NFTMintRequest`

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

### HasWallet

`func (o *NFTMintRequest) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetReturnCompiledTransaction

`func (o *NFTMintRequest) GetReturnCompiledTransaction() bool`

GetReturnCompiledTransaction returns the ReturnCompiledTransaction field if non-nil, zero value otherwise.

### GetReturnCompiledTransactionOk

`func (o *NFTMintRequest) GetReturnCompiledTransactionOk() (*bool, bool)`

GetReturnCompiledTransactionOk returns a tuple with the ReturnCompiledTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCompiledTransaction

`func (o *NFTMintRequest) SetReturnCompiledTransaction(v bool)`

SetReturnCompiledTransaction sets ReturnCompiledTransaction field to given value.

### HasReturnCompiledTransaction

`func (o *NFTMintRequest) HasReturnCompiledTransaction() bool`

HasReturnCompiledTransaction returns a boolean if a field has been set.

### GetName

`func (o *NFTMintRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NFTMintRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NFTMintRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NFTMintRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *NFTMintRequest) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NFTMintRequest) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NFTMintRequest) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *NFTMintRequest) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *NFTMintRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NFTMintRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NFTMintRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NFTMintRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUploadMethod

`func (o *NFTMintRequest) GetUploadMethod() string`

GetUploadMethod returns the UploadMethod field if non-nil, zero value otherwise.

### GetUploadMethodOk

`func (o *NFTMintRequest) GetUploadMethodOk() (*string, bool)`

GetUploadMethodOk returns a tuple with the UploadMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadMethod

`func (o *NFTMintRequest) SetUploadMethod(v string)`

SetUploadMethod sets UploadMethod field to given value.

### HasUploadMethod

`func (o *NFTMintRequest) HasUploadMethod() bool`

HasUploadMethod returns a boolean if a field has been set.

### GetUri

`func (o *NFTMintRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *NFTMintRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *NFTMintRequest) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *NFTMintRequest) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetImageUrl

`func (o *NFTMintRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *NFTMintRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *NFTMintRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *NFTMintRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetUriMetadata

`func (o *NFTMintRequest) GetUriMetadata() map[string]interface{}`

GetUriMetadata returns the UriMetadata field if non-nil, zero value otherwise.

### GetUriMetadataOk

`func (o *NFTMintRequest) GetUriMetadataOk() (*map[string]interface{}, bool)`

GetUriMetadataOk returns a tuple with the UriMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriMetadata

`func (o *NFTMintRequest) SetUriMetadata(v map[string]interface{})`

SetUriMetadata sets UriMetadata field to given value.

### HasUriMetadata

`func (o *NFTMintRequest) HasUriMetadata() bool`

HasUriMetadata returns a boolean if a field has been set.

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


