/*
The Blockchain API

# About  Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly.   You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>  # How to Use the API  To use the API, you simply need to create an API key pair.  Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.blockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.  # Feature Requests  Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@blockchainapi.com).  # Contact  <figure>     <img          width=\"40px\"         height=\"40px\"          src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"     />     <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption> </figure> <a href=\"mailto:info@blockchainapi.com\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"         />         <figcaption style=\"textAlign:center;\">Email us: info@blockchainapi.com</figcaption>     </figure> </a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"         />         <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>     </figure> </a> <a href=\"https://twitter.com/BlockchainAP1\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"         />          <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>     </figure> </a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"         />         <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>     </figure> </a>  # Security  Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet.   <b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b>   How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.  We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.  Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.  Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.  # Pricing  <b>Each user receives 50,000 free credits each month.</b> Each endpoint costs a certain amount credits. Scroll below to any endpoint (i.e., function) to see how much each endpoint costs. (Or CMD+F `Cost: 0 Credit`, for example)  You can learn more about our pricing <a href=\"https://dashboard.blockchainapi.com/billing\" target=\"_blank\">here</a>.   We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.  If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.  # SDKs / API Wrappers  We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.  We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.  `pip install theblockchainapi`  We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-javascript-wrapper\" target=\"_blank\">JavaScript Wrapper</a>.  `npm install theblockchainapi`  We also have auto-generated wrappers for the following languages: - <a href=\"https://github.com/BL0CK-X/theblockchainapi-go-wrapper\" target = \"_blank\">Go</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-java-wrapper\" target = \"_blank\">Java</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-kotlin-wrapper\" target = \"_blank\">Kotlin</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-php-wrapper\" target = \"_blank\">PHP</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-swift-wrapper\" target = \"_blank\">Swift5</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-typescript-wrapper\" target = \"_blank\">TypeScript</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-csharp-wrapper\" target = \"_blank\">C#</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-dart-wrapper\" target = \"_blank\">Dart</a>  If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.  If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

API version: null
Contact: info@blockchainapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NFTMintRequest struct for NFTMintRequest
type NFTMintRequest struct {
	// Whether to wait for the NFT mint to be confirmed on the blockchain or simply be processed.  Processed means that our node has picked up the transaction request, but not that it was confirmed by the Solana cluster.   Confirmed means that the cluster voted on your transaction and approved it. To be completely sure that the NFT was minted, you can either set `wait_for_confirmation=True` (call takes 20 seconds with True; about 4 seconds for processed) or you can [get the metadata](/#tag/Solana-NFT/operation/solanaGetNFT) from the mint returned. Once it returns the NFT metadata, then the NFT should have been successfully minted. 
	WaitForConfirmation *bool `json:"wait_for_confirmation,omitempty"`
	Wallet *Wallet `json:"wallet,omitempty"`
	// The name of the token. Limited to 32 characters. Stored on the blockchain.
	Name *string `json:"name,omitempty"`
	// The symbol of the token. Limited to 10 characters. Stored on the blockchain.
	Symbol *string `json:"symbol,omitempty"`
	// The description of the NFT. Limited to 2000 characters. Not stored on the blockchain.         If you are providing your own `uri` (see above), then you do not need to provide this.  If you are not providing your own `uri` and you do not provide this, then there wills simply be no description.  Only provide a value for `description` if the `upload_method` is set to `S3` (see the description for `upload_method` above).
	Description *string `json:"description,omitempty"`
	// When you choose `S3`, all of the `name`, `description`, `symbol`, `uri_metadata`, and `image_url` are put into a JSON dictionary and uploaded to S3 as a JSON file.  This is uploaded to an AWS S3 bucket we own, and is an option we provide at no charge. The S3 link to this file is stored in the NFT's account on the blockchain. Learn more  <a href=\"https://blockchainapi.com/2022/01/18/how-to-format-off-chain-nft-metadata.html\" target=\"_blank\">here</a>.  When you choose `URI`, the `uri` you provide is stored on the blockchain, and the `uri_metadata`, `description`, and `image_url` are ignored and not stored anywhere. `S3` is NOT involved in this case.   An example of a `uri` you would provide is an Arweave URL, like this: `https://arweave.net/_Y8tETU3FbAFZSM1wXNeWPweWwrW9K6oSF1SYi_bH9A`.
	UploadMethod *string `json:"upload_method,omitempty"`
	// The `uri` you provide is stored on the blockchain, and the `uri_metadata`, `description`, and `image_url` are ignored and not stored anywhere. `S3` is NOT involved in this case.   Read more <a href=\"https://blockchainapi.com/2022/01/18/how-to-format-off-chain-nft-metadata.html\" target=\"_blank\">here</a>.  An example of a `uri` you would provide is an Arweave URL, like this: `https://arweave.net/_Y8tETU3FbAFZSM1wXNeWPweWwrW9K6oSF1SYi_bH9A`.  Only provide a value for `uri` if the `upload_method` is set to `URI` (see the description for `upload_method` above).
	Uri *string `json:"uri,omitempty"`
	// The URL of the image of the NFT.         If you are providing your own `uri` (see above), then you do not need to provide this.  If you are not providing your own `uri` and you do not provide this, then there wills simply be no image.  Only provide a value for `image_url` if the `upload_method` is set to `S3` (see the description for `upload_method` above).
	ImageUrl *string `json:"image_url,omitempty"`
	// The off-chain metadata.        If you are providing your own `uri` (see above), then you do not need to provide this.  If you are not providing your own `uri` and you do not provide this, then there wills simply be no image.  Only provide a value for `uri_metadata` if the `upload_method` is set to `S3` (see the description for `upload_method` above).  Learn more about how to format this metadata <a href=\"https://blockchainapi.com/2022/01/18/how-to-format-off-chain-nft-metadata.html\" target=\"_blank\">here</a>.
	UriMetadata map[string]interface{} `json:"uri_metadata,omitempty"`
	// Indicates whether or not the NFT created is mutable. If mutable, the NFT can be updated later. Once set to immutable, the NFT is unable to be changed. 
	IsMutable *bool `json:"is_mutable,omitempty"`
	// Whether or not the NFT is a master edition NFT. Saves about 0.001 SOL in transaction costs when set to false. 
	IsMasterEdition *bool `json:"is_master_edition,omitempty"`
	// Valid values from 0 to 10000. Must be an integer.  Represents the number of basis points that the seller receives as a fee upon sale.  E.g., 100 indicates a 1% seller fee. Seller does not receive a fee when \"primary_sale_has_happened\" is set to true.  Will be set to false after first sale has occurred. 
	SellerFeeBasisPoints *float32 `json:"seller_fee_basis_points,omitempty"`
	// A JSON encoded string representing an array / list.  The designated creators of the NFT. Length of the creator list must match length of the list of share.  Valid lengths of the list range from 1 to 5. Each item in the list must be a valid public key address.    Only the public key corresponding to the seed phrase provided will be marked as verified. Any other creators supplied will be marked as unverified. 
	Creators []string `json:"creators,omitempty"`
	// A JSON encoded string representing an array / list.  The share of the royalty that each creator gets. Valid values range from 0 to 100. Sum of the values must equal 100.  Only integer value accepted. Length of the share list must match length of the list of creators. 
	Share []int32 `json:"share,omitempty"`
	// Assign ownership of the NFT to the public key address given by `mint_to_public_key` 
	MintToPublicKey *string `json:"mint_to_public_key,omitempty"`
	// This determines which network you choose to run the API calls on. We recommend first testing on the devnet, because minting an NFT costs a little above 0.01 SOL, which is about $1.60 at the time of this writing.  When you run on the mainnet-beta, each successful call will deduct approximately that much. When you run on the devnet, that amount is deducted from a simulated amount, so you are not paying with real SOL. To get SOL on the devnet,   airdrop SOL to this address using the CLI. Keep in mind that you can only do this every so often. If you are rate-limited, consider using a VPN and trying again, or just waiting. To get SOL on the mainnet-beta, you    must transfer real SOL to this account from another wallet (e.g., from another wallet you own, from an exchange, etc.). We hope to make this process easier in the future, and if you have any suggestions, please add them    as an issue on our <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">GitHub repository</a> for the API. To get a fee estimate, make a GET requests to the <a href=\"#tag/Solana-NFT/paths/~1solana~1nft~1mint~1fee/get\">v1/solana/nft/mint/fee endpoint</a> (details in sidebar). 
	Network *string `json:"network,omitempty"`
}

// NewNFTMintRequest instantiates a new NFTMintRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFTMintRequest() *NFTMintRequest {
	this := NFTMintRequest{}
	var waitForConfirmation bool = true
	this.WaitForConfirmation = &waitForConfirmation
	var name string = ""
	this.Name = &name
	var symbol string = ""
	this.Symbol = &symbol
	var description string = ""
	this.Description = &description
	var uploadMethod string = "S3"
	this.UploadMethod = &uploadMethod
	var uri string = ""
	this.Uri = &uri
	var imageUrl string = ""
	this.ImageUrl = &imageUrl
	var isMutable bool = true
	this.IsMutable = &isMutable
	var isMasterEdition bool = true
	this.IsMasterEdition = &isMasterEdition
	var sellerFeeBasisPoints float32 = 0
	this.SellerFeeBasisPoints = &sellerFeeBasisPoints
	var mintToPublicKey string = "The public key of the wallet provided"
	this.MintToPublicKey = &mintToPublicKey
	var network string = "devnet"
	this.Network = &network
	return &this
}

// NewNFTMintRequestWithDefaults instantiates a new NFTMintRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTMintRequestWithDefaults() *NFTMintRequest {
	this := NFTMintRequest{}
	var waitForConfirmation bool = true
	this.WaitForConfirmation = &waitForConfirmation
	var name string = ""
	this.Name = &name
	var symbol string = ""
	this.Symbol = &symbol
	var description string = ""
	this.Description = &description
	var uploadMethod string = "S3"
	this.UploadMethod = &uploadMethod
	var uri string = ""
	this.Uri = &uri
	var imageUrl string = ""
	this.ImageUrl = &imageUrl
	var isMutable bool = true
	this.IsMutable = &isMutable
	var isMasterEdition bool = true
	this.IsMasterEdition = &isMasterEdition
	var sellerFeeBasisPoints float32 = 0
	this.SellerFeeBasisPoints = &sellerFeeBasisPoints
	var mintToPublicKey string = "The public key of the wallet provided"
	this.MintToPublicKey = &mintToPublicKey
	var network string = "devnet"
	this.Network = &network
	return &this
}

// GetWaitForConfirmation returns the WaitForConfirmation field value if set, zero value otherwise.
func (o *NFTMintRequest) GetWaitForConfirmation() bool {
	if o == nil || o.WaitForConfirmation == nil {
		var ret bool
		return ret
	}
	return *o.WaitForConfirmation
}

// GetWaitForConfirmationOk returns a tuple with the WaitForConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetWaitForConfirmationOk() (*bool, bool) {
	if o == nil || o.WaitForConfirmation == nil {
		return nil, false
	}
	return o.WaitForConfirmation, true
}

// HasWaitForConfirmation returns a boolean if a field has been set.
func (o *NFTMintRequest) HasWaitForConfirmation() bool {
	if o != nil && o.WaitForConfirmation != nil {
		return true
	}

	return false
}

// SetWaitForConfirmation gets a reference to the given bool and assigns it to the WaitForConfirmation field.
func (o *NFTMintRequest) SetWaitForConfirmation(v bool) {
	o.WaitForConfirmation = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *NFTMintRequest) GetWallet() Wallet {
	if o == nil || o.Wallet == nil {
		var ret Wallet
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetWalletOk() (*Wallet, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *NFTMintRequest) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given Wallet and assigns it to the Wallet field.
func (o *NFTMintRequest) SetWallet(v Wallet) {
	o.Wallet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NFTMintRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NFTMintRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NFTMintRequest) SetName(v string) {
	o.Name = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *NFTMintRequest) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *NFTMintRequest) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *NFTMintRequest) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NFTMintRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NFTMintRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NFTMintRequest) SetDescription(v string) {
	o.Description = &v
}

// GetUploadMethod returns the UploadMethod field value if set, zero value otherwise.
func (o *NFTMintRequest) GetUploadMethod() string {
	if o == nil || o.UploadMethod == nil {
		var ret string
		return ret
	}
	return *o.UploadMethod
}

// GetUploadMethodOk returns a tuple with the UploadMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetUploadMethodOk() (*string, bool) {
	if o == nil || o.UploadMethod == nil {
		return nil, false
	}
	return o.UploadMethod, true
}

// HasUploadMethod returns a boolean if a field has been set.
func (o *NFTMintRequest) HasUploadMethod() bool {
	if o != nil && o.UploadMethod != nil {
		return true
	}

	return false
}

// SetUploadMethod gets a reference to the given string and assigns it to the UploadMethod field.
func (o *NFTMintRequest) SetUploadMethod(v string) {
	o.UploadMethod = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *NFTMintRequest) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *NFTMintRequest) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *NFTMintRequest) SetUri(v string) {
	o.Uri = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *NFTMintRequest) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *NFTMintRequest) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *NFTMintRequest) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetUriMetadata returns the UriMetadata field value if set, zero value otherwise.
func (o *NFTMintRequest) GetUriMetadata() map[string]interface{} {
	if o == nil || o.UriMetadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UriMetadata
}

// GetUriMetadataOk returns a tuple with the UriMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetUriMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.UriMetadata == nil {
		return nil, false
	}
	return o.UriMetadata, true
}

// HasUriMetadata returns a boolean if a field has been set.
func (o *NFTMintRequest) HasUriMetadata() bool {
	if o != nil && o.UriMetadata != nil {
		return true
	}

	return false
}

// SetUriMetadata gets a reference to the given map[string]interface{} and assigns it to the UriMetadata field.
func (o *NFTMintRequest) SetUriMetadata(v map[string]interface{}) {
	o.UriMetadata = v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *NFTMintRequest) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *NFTMintRequest) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *NFTMintRequest) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetIsMasterEdition returns the IsMasterEdition field value if set, zero value otherwise.
func (o *NFTMintRequest) GetIsMasterEdition() bool {
	if o == nil || o.IsMasterEdition == nil {
		var ret bool
		return ret
	}
	return *o.IsMasterEdition
}

// GetIsMasterEditionOk returns a tuple with the IsMasterEdition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetIsMasterEditionOk() (*bool, bool) {
	if o == nil || o.IsMasterEdition == nil {
		return nil, false
	}
	return o.IsMasterEdition, true
}

// HasIsMasterEdition returns a boolean if a field has been set.
func (o *NFTMintRequest) HasIsMasterEdition() bool {
	if o != nil && o.IsMasterEdition != nil {
		return true
	}

	return false
}

// SetIsMasterEdition gets a reference to the given bool and assigns it to the IsMasterEdition field.
func (o *NFTMintRequest) SetIsMasterEdition(v bool) {
	o.IsMasterEdition = &v
}

// GetSellerFeeBasisPoints returns the SellerFeeBasisPoints field value if set, zero value otherwise.
func (o *NFTMintRequest) GetSellerFeeBasisPoints() float32 {
	if o == nil || o.SellerFeeBasisPoints == nil {
		var ret float32
		return ret
	}
	return *o.SellerFeeBasisPoints
}

// GetSellerFeeBasisPointsOk returns a tuple with the SellerFeeBasisPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetSellerFeeBasisPointsOk() (*float32, bool) {
	if o == nil || o.SellerFeeBasisPoints == nil {
		return nil, false
	}
	return o.SellerFeeBasisPoints, true
}

// HasSellerFeeBasisPoints returns a boolean if a field has been set.
func (o *NFTMintRequest) HasSellerFeeBasisPoints() bool {
	if o != nil && o.SellerFeeBasisPoints != nil {
		return true
	}

	return false
}

// SetSellerFeeBasisPoints gets a reference to the given float32 and assigns it to the SellerFeeBasisPoints field.
func (o *NFTMintRequest) SetSellerFeeBasisPoints(v float32) {
	o.SellerFeeBasisPoints = &v
}

// GetCreators returns the Creators field value if set, zero value otherwise.
func (o *NFTMintRequest) GetCreators() []string {
	if o == nil || o.Creators == nil {
		var ret []string
		return ret
	}
	return o.Creators
}

// GetCreatorsOk returns a tuple with the Creators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetCreatorsOk() ([]string, bool) {
	if o == nil || o.Creators == nil {
		return nil, false
	}
	return o.Creators, true
}

// HasCreators returns a boolean if a field has been set.
func (o *NFTMintRequest) HasCreators() bool {
	if o != nil && o.Creators != nil {
		return true
	}

	return false
}

// SetCreators gets a reference to the given []string and assigns it to the Creators field.
func (o *NFTMintRequest) SetCreators(v []string) {
	o.Creators = v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *NFTMintRequest) GetShare() []int32 {
	if o == nil || o.Share == nil {
		var ret []int32
		return ret
	}
	return o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetShareOk() ([]int32, bool) {
	if o == nil || o.Share == nil {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *NFTMintRequest) HasShare() bool {
	if o != nil && o.Share != nil {
		return true
	}

	return false
}

// SetShare gets a reference to the given []int32 and assigns it to the Share field.
func (o *NFTMintRequest) SetShare(v []int32) {
	o.Share = v
}

// GetMintToPublicKey returns the MintToPublicKey field value if set, zero value otherwise.
func (o *NFTMintRequest) GetMintToPublicKey() string {
	if o == nil || o.MintToPublicKey == nil {
		var ret string
		return ret
	}
	return *o.MintToPublicKey
}

// GetMintToPublicKeyOk returns a tuple with the MintToPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetMintToPublicKeyOk() (*string, bool) {
	if o == nil || o.MintToPublicKey == nil {
		return nil, false
	}
	return o.MintToPublicKey, true
}

// HasMintToPublicKey returns a boolean if a field has been set.
func (o *NFTMintRequest) HasMintToPublicKey() bool {
	if o != nil && o.MintToPublicKey != nil {
		return true
	}

	return false
}

// SetMintToPublicKey gets a reference to the given string and assigns it to the MintToPublicKey field.
func (o *NFTMintRequest) SetMintToPublicKey(v string) {
	o.MintToPublicKey = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NFTMintRequest) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NFTMintRequest) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *NFTMintRequest) SetNetwork(v string) {
	o.Network = &v
}

func (o NFTMintRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WaitForConfirmation != nil {
		toSerialize["wait_for_confirmation"] = o.WaitForConfirmation
	}
	if o.Wallet != nil {
		toSerialize["wallet"] = o.Wallet
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.UploadMethod != nil {
		toSerialize["upload_method"] = o.UploadMethod
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.UriMetadata != nil {
		toSerialize["uri_metadata"] = o.UriMetadata
	}
	if o.IsMutable != nil {
		toSerialize["is_mutable"] = o.IsMutable
	}
	if o.IsMasterEdition != nil {
		toSerialize["is_master_edition"] = o.IsMasterEdition
	}
	if o.SellerFeeBasisPoints != nil {
		toSerialize["seller_fee_basis_points"] = o.SellerFeeBasisPoints
	}
	if o.Creators != nil {
		toSerialize["creators"] = o.Creators
	}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	if o.MintToPublicKey != nil {
		toSerialize["mint_to_public_key"] = o.MintToPublicKey
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableNFTMintRequest struct {
	value *NFTMintRequest
	isSet bool
}

func (v NullableNFTMintRequest) Get() *NFTMintRequest {
	return v.value
}

func (v *NullableNFTMintRequest) Set(val *NFTMintRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNFTMintRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNFTMintRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFTMintRequest(val *NFTMintRequest) *NullableNFTMintRequest {
	return &NullableNFTMintRequest{value: val, isSet: true}
}

func (v NullableNFTMintRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFTMintRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


