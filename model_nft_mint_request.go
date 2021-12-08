/*
The Blockchain API

# About  Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly.   You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>  # How to Use the API  To use the API, you simply need to create an API key pair.  Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.theblockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.  # Feature Requests  Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@theblockchainapi.com).  # Contact  <figure>     <img          width=\"40px\"         height=\"40px\"          src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"     />     <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption> </figure> <a href=\"mailto:info@theblockchainapi.com\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"         />         <figcaption style=\"textAlign:center;\">Email us: info@theblockchainapi.com</figcaption>     </figure> </a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"         />         <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>     </figure> </a> <a href=\"https://twitter.com/_BlockX_\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"         />          <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>     </figure> </a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"         />         <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>     </figure> </a>  # Security  Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet.   <b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b>   How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.  We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.  Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.  Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.  # Pricing  <b>Each user receives 500 free credits every month.</b>  You can learn more about our pricing <a href=\"https://dashboard.theblockchainapi.com/billing\" target=\"_blank\">here</a>.   We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.  If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.  # SDKs / API Wrappers  We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.  We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.  `pip install theblockchainapi`  We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/javascript\" target=\"_blank\">JavaScript Wrapper</a>.  `npm install theblockchainapi`  We also have auto-generated wrappers for the following languages: - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">Go</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/java\" target = \"_blank\">Java</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">Kotlin</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">PHP</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">Swift5</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">TypeScript</a>  If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.  If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

API version: null
Contact: info@theblockchainapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NFTMintRequest struct for NFTMintRequest
type NFTMintRequest struct {
	// The twelve word phrase that can be used to derive many public key addresses. To derive a public key, you need a secret recovery phrase, a derivation path, and an optional passphrase. See our Security section <a href=\"#section/Security\">here</a>.
	SecretRecoveryPhrase string `json:"secret_recovery_phrase"`
	// Derivation paths are used to derive the public key from the secret recovery phrase. Only certain paths are accepted.  We use \"m/44/501/0/0\" by default, if it is not provided. This is the path that the Phantom and Sollet wallets use. If you provide the empty string \"\" as the value for the derivation path, then we will use the Solana CLI default value. The SolFlare recommended path is \"m/44/501/0\".  You can also arbitrarily increment the default path (\"m/44/501/0/0\") to generate more wallets (e.g., \"m/44/501/0/1\", \"m/44/501/0/2\", ...). This is how Phantom generates more wallets.  To learn more about derivation paths, check out <a href=\"https://learnmeabitcoin.com/technical/derivation-paths\" target=\"_blank\">this tutorial</a>.
	DerivationPath *string `json:"derivation_path,omitempty"`
	// PASSPHRASE != PASSWORD. This is NOT your Phantom password or any other password. It is an optional string you use when creating a wallet. This provides an additional layer of security because a hacker would need both the secret recovery phrase and the passphrase to access the output public key. By default, most wallet UI extensions do not use a passphrase. (You probably did not use a passphrase.) Limited to 500 characters. 
	Passphrase *string `json:"passphrase,omitempty"`
	// The name of the token. Limited to 32 characters. Stored on the blockchain.
	NftName *string `json:"nft_name,omitempty"`
	// The symbol of the token. Limited to 10 characters. Stored on the blockchain.
	NftSymbol *string `json:"nft_symbol,omitempty"`
	// The description of the token. Limited to 2000 characters. Not stored on the blockchain.  This is stored in S3 in a bucket we own, and the link to that file is stored on the blockchain.  If you provide your own link, the link is also stored in that S3 file, which is publicly accessible. However, if you choose the NFT upload method \"LINK\" instead of \"S3\", then we upload the link you  provide for nft_url directly to the blockchain, and S3 is not used at all. Thus, when you provide the \"LINK\" option, the value nft_description is ignored and not used. The Metaplex API does not provide functionality to store any data about your NFT besides a URL. 
	NftDescription *string `json:"nft_description,omitempty"`
	// The URL you provided. Limited to 200 characters. This does not need to be a valid URL. Whether or not this is  stored on the blockchain depends on which upload method you choose. If you choose LINK, then this is stored on the  blockchain directly. If you choose S3, then this link is embedded in a public S3 text file that also contains the metadata, the name,  the symbol, and the description of the NFT. 
	NftUrl *string `json:"nft_url,omitempty"`
	// Any data you provide. Must be a string and properly encoded json. Will be viewable on S3. Limited to 2000 bytes. Not stored on the blockchain.  This is stored in S3 in a bucket we own, and the link to that file is stored on the blockchain.  If you provide your own link, the link is also stored in that S3 file, which is publicly accessible. However, if you choose the NFT upload method \"LINK\" instead of \"S3\", then we upload the link you  provide for nft_url directly to the blockchain, and S3 is not used at all. Thus, when you provide the \"LINK\" option, the value nft_metadata is ignored and not used. The Metaplex API does not provide functionality to store any data about your NFT besides a URL. 
	NftMetadata *string `json:"nft_metadata,omitempty"`
	// When you choose S3, all of the nft_description, nft_name, nft_symbol, nft_metadata, and nft_url are put into a json dictionary and uploaded to S3 as a text file.  This is uploaded to an AWS S3 bucket we own, and is an option we provide at no charge. The S3 link to this file is stored on the NFT on the blockchain.   When you choose LINK, the nft_url you provide is stored on the blockchain, and the nft_metadata and nft_description are ignored and not stored anywhere.  
	NftUploadMethod *string `json:"nft_upload_method,omitempty"`
	// Indicates whether or not the NFT created is mutable. If mutable, the NFT can be updated later. Once set to immutable, the NFT is unable to be changed. 
	IsMutable *bool `json:"is_mutable,omitempty"`
	// Whether or not the NFT is a master edition NFT. Saves about 0.001 SOL in transaction costs when set to false. 
	IsMasterEdition *bool `json:"is_master_edition,omitempty"`
	// Valid values from 0 to 10000. Must be an integer.  Represents the number of basis points that the seller receives as a fee upon sale.  E.g., 100 indicates a 1% seller fee. Seller does not receive a fee when \"primary_sale_has_happened\" is set to true.  Will be set to false after first sale has occurred. 
	SellerFeeBasisPoints *float32 `json:"seller_fee_basis_points,omitempty"`
	// A JSON encoded string representing an array / list.  The designated creators of the NFT. Length of the creator list must match length of the list of share.  Valid lengths of the list range from 1 to 5. Each item in the list must be a valid public key address.    Only the public key corresponding to the seed phrase provided will be marked as verified. Any other creators supplied will be marked as unverified. 
	Creators *[]string `json:"creators,omitempty"`
	// A JSON encoded string representing an array / list.  The share of the royalty that each creator gets. Valid values range from 0 to 100.  Sum of the values must equal 100.  Only integer value accepted. Length of the share list must match length of the list of creators. 
	Share *[]int32 `json:"share,omitempty"`
	// This determines which network you choose to run the API calls on. We recommend first testing on the devnet, because minting an NFT costs a little above 0.01 SOL, which is about $1.60 at the time of this writing.  When you run on the mainnet-beta, each successful call will deduct approximately that much. When you run on the devnet, that amount is deducted from a simulated amount, so you are not paying with real SOL. To get SOL on the devnet,   airdrop SOL to this address using the CLI. Keep in mind that you can only do this every so often. If you are rate-limited, consider using a VPN and trying again, or just waiting. To get SOL on the mainnet-beta, you    must transfer real SOL to this account from another wallet (e.g., from another wallet you own, from an exchange, etc.). We hope to make this process easier in the future, and if you have any suggestions, please add them    as an issue on our <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">GitHub repository</a> for the API. To get a fee estimate, make a GET requests to the <a href=\"#tag/Solana-NFT/paths/~1solana~1nft~1mint~1fee/get\">v1/solana/nft/mint/fee endpoint</a> (details in sidebar). 
	Network *string `json:"network,omitempty"`
}

// NewNFTMintRequest instantiates a new NFTMintRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFTMintRequest(secretRecoveryPhrase string) *NFTMintRequest {
	this := NFTMintRequest{}
	this.SecretRecoveryPhrase = secretRecoveryPhrase
	var derivationPath string = "m/44/501/0/0"
	this.DerivationPath = &derivationPath
	var passphrase string = ""
	this.Passphrase = &passphrase
	var nftName string = ""
	this.NftName = &nftName
	var nftSymbol string = ""
	this.NftSymbol = &nftSymbol
	var nftDescription string = ""
	this.NftDescription = &nftDescription
	var nftUrl string = ""
	this.NftUrl = &nftUrl
	var nftMetadata string = "{}"
	this.NftMetadata = &nftMetadata
	var nftUploadMethod string = "S3"
	this.NftUploadMethod = &nftUploadMethod
	var isMutable bool = true
	this.IsMutable = &isMutable
	var isMasterEdition bool = true
	this.IsMasterEdition = &isMasterEdition
	var sellerFeeBasisPoints float32 = 0
	this.SellerFeeBasisPoints = &sellerFeeBasisPoints
	var network string = "devnet"
	this.Network = &network
	return &this
}

// NewNFTMintRequestWithDefaults instantiates a new NFTMintRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTMintRequestWithDefaults() *NFTMintRequest {
	this := NFTMintRequest{}
	var derivationPath string = "m/44/501/0/0"
	this.DerivationPath = &derivationPath
	var passphrase string = ""
	this.Passphrase = &passphrase
	var nftName string = ""
	this.NftName = &nftName
	var nftSymbol string = ""
	this.NftSymbol = &nftSymbol
	var nftDescription string = ""
	this.NftDescription = &nftDescription
	var nftUrl string = ""
	this.NftUrl = &nftUrl
	var nftMetadata string = "{}"
	this.NftMetadata = &nftMetadata
	var nftUploadMethod string = "S3"
	this.NftUploadMethod = &nftUploadMethod
	var isMutable bool = true
	this.IsMutable = &isMutable
	var isMasterEdition bool = true
	this.IsMasterEdition = &isMasterEdition
	var sellerFeeBasisPoints float32 = 0
	this.SellerFeeBasisPoints = &sellerFeeBasisPoints
	var network string = "devnet"
	this.Network = &network
	return &this
}

// GetSecretRecoveryPhrase returns the SecretRecoveryPhrase field value
func (o *NFTMintRequest) GetSecretRecoveryPhrase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretRecoveryPhrase
}

// GetSecretRecoveryPhraseOk returns a tuple with the SecretRecoveryPhrase field value
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetSecretRecoveryPhraseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecretRecoveryPhrase, true
}

// SetSecretRecoveryPhrase sets field value
func (o *NFTMintRequest) SetSecretRecoveryPhrase(v string) {
	o.SecretRecoveryPhrase = v
}

// GetDerivationPath returns the DerivationPath field value if set, zero value otherwise.
func (o *NFTMintRequest) GetDerivationPath() string {
	if o == nil || o.DerivationPath == nil {
		var ret string
		return ret
	}
	return *o.DerivationPath
}

// GetDerivationPathOk returns a tuple with the DerivationPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetDerivationPathOk() (*string, bool) {
	if o == nil || o.DerivationPath == nil {
		return nil, false
	}
	return o.DerivationPath, true
}

// HasDerivationPath returns a boolean if a field has been set.
func (o *NFTMintRequest) HasDerivationPath() bool {
	if o != nil && o.DerivationPath != nil {
		return true
	}

	return false
}

// SetDerivationPath gets a reference to the given string and assigns it to the DerivationPath field.
func (o *NFTMintRequest) SetDerivationPath(v string) {
	o.DerivationPath = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *NFTMintRequest) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *NFTMintRequest) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *NFTMintRequest) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetNftName returns the NftName field value if set, zero value otherwise.
func (o *NFTMintRequest) GetNftName() string {
	if o == nil || o.NftName == nil {
		var ret string
		return ret
	}
	return *o.NftName
}

// GetNftNameOk returns a tuple with the NftName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetNftNameOk() (*string, bool) {
	if o == nil || o.NftName == nil {
		return nil, false
	}
	return o.NftName, true
}

// HasNftName returns a boolean if a field has been set.
func (o *NFTMintRequest) HasNftName() bool {
	if o != nil && o.NftName != nil {
		return true
	}

	return false
}

// SetNftName gets a reference to the given string and assigns it to the NftName field.
func (o *NFTMintRequest) SetNftName(v string) {
	o.NftName = &v
}

// GetNftSymbol returns the NftSymbol field value if set, zero value otherwise.
func (o *NFTMintRequest) GetNftSymbol() string {
	if o == nil || o.NftSymbol == nil {
		var ret string
		return ret
	}
	return *o.NftSymbol
}

// GetNftSymbolOk returns a tuple with the NftSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetNftSymbolOk() (*string, bool) {
	if o == nil || o.NftSymbol == nil {
		return nil, false
	}
	return o.NftSymbol, true
}

// HasNftSymbol returns a boolean if a field has been set.
func (o *NFTMintRequest) HasNftSymbol() bool {
	if o != nil && o.NftSymbol != nil {
		return true
	}

	return false
}

// SetNftSymbol gets a reference to the given string and assigns it to the NftSymbol field.
func (o *NFTMintRequest) SetNftSymbol(v string) {
	o.NftSymbol = &v
}

// GetNftDescription returns the NftDescription field value if set, zero value otherwise.
func (o *NFTMintRequest) GetNftDescription() string {
	if o == nil || o.NftDescription == nil {
		var ret string
		return ret
	}
	return *o.NftDescription
}

// GetNftDescriptionOk returns a tuple with the NftDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetNftDescriptionOk() (*string, bool) {
	if o == nil || o.NftDescription == nil {
		return nil, false
	}
	return o.NftDescription, true
}

// HasNftDescription returns a boolean if a field has been set.
func (o *NFTMintRequest) HasNftDescription() bool {
	if o != nil && o.NftDescription != nil {
		return true
	}

	return false
}

// SetNftDescription gets a reference to the given string and assigns it to the NftDescription field.
func (o *NFTMintRequest) SetNftDescription(v string) {
	o.NftDescription = &v
}

// GetNftUrl returns the NftUrl field value if set, zero value otherwise.
func (o *NFTMintRequest) GetNftUrl() string {
	if o == nil || o.NftUrl == nil {
		var ret string
		return ret
	}
	return *o.NftUrl
}

// GetNftUrlOk returns a tuple with the NftUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetNftUrlOk() (*string, bool) {
	if o == nil || o.NftUrl == nil {
		return nil, false
	}
	return o.NftUrl, true
}

// HasNftUrl returns a boolean if a field has been set.
func (o *NFTMintRequest) HasNftUrl() bool {
	if o != nil && o.NftUrl != nil {
		return true
	}

	return false
}

// SetNftUrl gets a reference to the given string and assigns it to the NftUrl field.
func (o *NFTMintRequest) SetNftUrl(v string) {
	o.NftUrl = &v
}

// GetNftMetadata returns the NftMetadata field value if set, zero value otherwise.
func (o *NFTMintRequest) GetNftMetadata() string {
	if o == nil || o.NftMetadata == nil {
		var ret string
		return ret
	}
	return *o.NftMetadata
}

// GetNftMetadataOk returns a tuple with the NftMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetNftMetadataOk() (*string, bool) {
	if o == nil || o.NftMetadata == nil {
		return nil, false
	}
	return o.NftMetadata, true
}

// HasNftMetadata returns a boolean if a field has been set.
func (o *NFTMintRequest) HasNftMetadata() bool {
	if o != nil && o.NftMetadata != nil {
		return true
	}

	return false
}

// SetNftMetadata gets a reference to the given string and assigns it to the NftMetadata field.
func (o *NFTMintRequest) SetNftMetadata(v string) {
	o.NftMetadata = &v
}

// GetNftUploadMethod returns the NftUploadMethod field value if set, zero value otherwise.
func (o *NFTMintRequest) GetNftUploadMethod() string {
	if o == nil || o.NftUploadMethod == nil {
		var ret string
		return ret
	}
	return *o.NftUploadMethod
}

// GetNftUploadMethodOk returns a tuple with the NftUploadMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetNftUploadMethodOk() (*string, bool) {
	if o == nil || o.NftUploadMethod == nil {
		return nil, false
	}
	return o.NftUploadMethod, true
}

// HasNftUploadMethod returns a boolean if a field has been set.
func (o *NFTMintRequest) HasNftUploadMethod() bool {
	if o != nil && o.NftUploadMethod != nil {
		return true
	}

	return false
}

// SetNftUploadMethod gets a reference to the given string and assigns it to the NftUploadMethod field.
func (o *NFTMintRequest) SetNftUploadMethod(v string) {
	o.NftUploadMethod = &v
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
	return *o.Creators
}

// GetCreatorsOk returns a tuple with the Creators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetCreatorsOk() (*[]string, bool) {
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
	o.Creators = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *NFTMintRequest) GetShare() []int32 {
	if o == nil || o.Share == nil {
		var ret []int32
		return ret
	}
	return *o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTMintRequest) GetShareOk() (*[]int32, bool) {
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
	o.Share = &v
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
	if true {
		toSerialize["secret_recovery_phrase"] = o.SecretRecoveryPhrase
	}
	if o.DerivationPath != nil {
		toSerialize["derivation_path"] = o.DerivationPath
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	if o.NftName != nil {
		toSerialize["nft_name"] = o.NftName
	}
	if o.NftSymbol != nil {
		toSerialize["nft_symbol"] = o.NftSymbol
	}
	if o.NftDescription != nil {
		toSerialize["nft_description"] = o.NftDescription
	}
	if o.NftUrl != nil {
		toSerialize["nft_url"] = o.NftUrl
	}
	if o.NftMetadata != nil {
		toSerialize["nft_metadata"] = o.NftMetadata
	}
	if o.NftUploadMethod != nil {
		toSerialize["nft_upload_method"] = o.NftUploadMethod
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


