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

// GeneralBalanceRequest struct for GeneralBalanceRequest
type GeneralBalanceRequest struct {
	// The address / public key of the wallet you're querying.  Examples: - Solana: `GKNcUmNacSJo4S2Kq3DuYRYRGw3sNUfJ4tyqd198t6vQ` - Ethereum: `0xa84b9478d203cd25dF722e83C87590f8028f6aAA`
	BlockchainIdentifier *string `json:"blockchain_identifier,omitempty"`
	// The `unit` parameter is only applicable if you are trying to retrieve the balance of the native token (e.g., SOL, ETH, BNB).   Applicable units: - Solana: `lamport`, `sol` (1 SOL = 1e9 Lamports) - Ethereum: `wei`, `gwei`, `eth`
	Unit *string `json:"unit,omitempty"`
	// The network of the blockchain you selected  - Solana: `devnet`, `mainnet-beta` - Ethereum: `ropsten`, `mainnet`  Defaults when not provided (not applicable to path parameters): - Solana: `devnet` - Ethereum: `ropsten`
	Network *string `json:"network,omitempty"`
	// The `token_blockchain_identifier` identifies the token you wish to transfer.  - If you're transferring a native blockchain currency (e.g., SOL, ETH, BNB), then simply do not supply this value. - If you're transfering an NFT, then supply the token address of the NFT. On Solana, this is the `mint_address` or `mint` (the address of the mint). - If you're transfering a token, supply the token address. For Solana, you can find on this on the Solana Explorer (e.g., see `SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt` for <a href=\"https://explorer.solana.com/address/SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt\" target=\"_blank\">Serum Token</a>) for the `token_address`.  Examples: - Ethereum: `0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48` - Solana: `CK1LHEANTu7RFqN3XMzo2AnZhyus2W1vue1njrxLEM1d`
	TokenBlockchainIdentifier *string `json:"token_blockchain_identifier,omitempty"`
}

// NewGeneralBalanceRequest instantiates a new GeneralBalanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralBalanceRequest() *GeneralBalanceRequest {
	this := GeneralBalanceRequest{}
	var tokenBlockchainIdentifier string = "null"
	this.TokenBlockchainIdentifier = &tokenBlockchainIdentifier
	return &this
}

// NewGeneralBalanceRequestWithDefaults instantiates a new GeneralBalanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralBalanceRequestWithDefaults() *GeneralBalanceRequest {
	this := GeneralBalanceRequest{}
	var tokenBlockchainIdentifier string = "null"
	this.TokenBlockchainIdentifier = &tokenBlockchainIdentifier
	return &this
}

// GetBlockchainIdentifier returns the BlockchainIdentifier field value if set, zero value otherwise.
func (o *GeneralBalanceRequest) GetBlockchainIdentifier() string {
	if o == nil || o.BlockchainIdentifier == nil {
		var ret string
		return ret
	}
	return *o.BlockchainIdentifier
}

// GetBlockchainIdentifierOk returns a tuple with the BlockchainIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralBalanceRequest) GetBlockchainIdentifierOk() (*string, bool) {
	if o == nil || o.BlockchainIdentifier == nil {
		return nil, false
	}
	return o.BlockchainIdentifier, true
}

// HasBlockchainIdentifier returns a boolean if a field has been set.
func (o *GeneralBalanceRequest) HasBlockchainIdentifier() bool {
	if o != nil && o.BlockchainIdentifier != nil {
		return true
	}

	return false
}

// SetBlockchainIdentifier gets a reference to the given string and assigns it to the BlockchainIdentifier field.
func (o *GeneralBalanceRequest) SetBlockchainIdentifier(v string) {
	o.BlockchainIdentifier = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *GeneralBalanceRequest) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralBalanceRequest) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *GeneralBalanceRequest) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *GeneralBalanceRequest) SetUnit(v string) {
	o.Unit = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GeneralBalanceRequest) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralBalanceRequest) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GeneralBalanceRequest) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *GeneralBalanceRequest) SetNetwork(v string) {
	o.Network = &v
}

// GetTokenBlockchainIdentifier returns the TokenBlockchainIdentifier field value if set, zero value otherwise.
func (o *GeneralBalanceRequest) GetTokenBlockchainIdentifier() string {
	if o == nil || o.TokenBlockchainIdentifier == nil {
		var ret string
		return ret
	}
	return *o.TokenBlockchainIdentifier
}

// GetTokenBlockchainIdentifierOk returns a tuple with the TokenBlockchainIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralBalanceRequest) GetTokenBlockchainIdentifierOk() (*string, bool) {
	if o == nil || o.TokenBlockchainIdentifier == nil {
		return nil, false
	}
	return o.TokenBlockchainIdentifier, true
}

// HasTokenBlockchainIdentifier returns a boolean if a field has been set.
func (o *GeneralBalanceRequest) HasTokenBlockchainIdentifier() bool {
	if o != nil && o.TokenBlockchainIdentifier != nil {
		return true
	}

	return false
}

// SetTokenBlockchainIdentifier gets a reference to the given string and assigns it to the TokenBlockchainIdentifier field.
func (o *GeneralBalanceRequest) SetTokenBlockchainIdentifier(v string) {
	o.TokenBlockchainIdentifier = &v
}

func (o GeneralBalanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockchainIdentifier != nil {
		toSerialize["blockchain_identifier"] = o.BlockchainIdentifier
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.TokenBlockchainIdentifier != nil {
		toSerialize["token_blockchain_identifier"] = o.TokenBlockchainIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableGeneralBalanceRequest struct {
	value *GeneralBalanceRequest
	isSet bool
}

func (v NullableGeneralBalanceRequest) Get() *GeneralBalanceRequest {
	return v.value
}

func (v *NullableGeneralBalanceRequest) Set(val *GeneralBalanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralBalanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralBalanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralBalanceRequest(val *GeneralBalanceRequest) *NullableGeneralBalanceRequest {
	return &NullableGeneralBalanceRequest{value: val, isSet: true}
}

func (v NullableGeneralBalanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralBalanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


