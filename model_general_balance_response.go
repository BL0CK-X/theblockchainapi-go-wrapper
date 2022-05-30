/*
The Blockchain API

# About  Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly.   You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>  # How to Use the API  To use the API, you simply need to create an API key pair.  Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.blockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.  # Feature Requests  Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@blockchainapi.com).  # Contact  <figure>     <img          width=\"40px\"         height=\"40px\"          src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"     />     <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption> </figure> <a href=\"mailto:info@blockchainapi.com\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"         />         <figcaption style=\"textAlign:center;\">Email us: info@blockchainapi.com</figcaption>     </figure> </a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"         />         <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>     </figure> </a> <a href=\"https://twitter.com/_BlockX_\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"         />          <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>     </figure> </a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"         />         <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>     </figure> </a>  # Security  Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet.   <b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b>   How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.  We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.  Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.  Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.  # Pricing  <b>Each user receives 100 free credits. Each user can call endpoints that cost 0 credits up to 50 requests/min before being rate-limited.</b> Thereafter, they can upgrade to have a higher rate limit. The purpose of this is to act like a free trial -- not to function like a freemium model where one can stay on the free tier indefinitely.  You can learn more about our pricing <a href=\"https://dashboard.blockchainapi.com/billing\" target=\"_blank\">here</a>.   We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.  If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.  # SDKs / API Wrappers  We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.  We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.  `pip install theblockchainapi`  We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-javascript-wrapper\" target=\"_blank\">JavaScript Wrapper</a>.  `npm install theblockchainapi`  We also have auto-generated wrappers for the following languages: - <a href=\"https://github.com/BL0CK-X/theblockchainapi-go-wrapper\" target = \"_blank\">Go</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-java-wrapper\" target = \"_blank\">Java</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-kotlin-wrapper\" target = \"_blank\">Kotlin</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-php-wrapper\" target = \"_blank\">PHP</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-swift-wrapper\" target = \"_blank\">Swift5</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-typescript-wrapper\" target = \"_blank\">TypeScript</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-csharp-wrapper\" target = \"_blank\">C#</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-dart-wrapper\" target = \"_blank\">Dart</a>  If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.  If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

API version: null
Contact: info@blockchainapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GeneralBalanceResponse struct for GeneralBalanceResponse
type GeneralBalanceResponse struct {
	// The balance of the token in the wallet 
	Balance float32 `json:"balance"`
	// Not included if retreiving native token (e.g., SOL, ETH, BNB, etc.) balance 
	IntegerBalance *float32 `json:"integer_balance,omitempty"`
	// Not included if retreiving native token (e.g., SOL, ETH, BNB, etc.) balance. 
	Decimals *float32 `json:"decimals,omitempty"`
	// The network of the blockchain you selected  - Solana: `devnet`, `mainnet-beta` - Ethereum: `ropsten`, `mainnet`  Defaults when not provided (not applicable to path parameters): - Solana: `devnet` - Ethereum: `ropsten`
	Network string `json:"network"`
	// Not included if retreiving a token / NFT balance 
	Unit *string `json:"unit,omitempty"`
}

// NewGeneralBalanceResponse instantiates a new GeneralBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralBalanceResponse(balance float32, network string) *GeneralBalanceResponse {
	this := GeneralBalanceResponse{}
	this.Balance = balance
	this.Network = network
	return &this
}

// NewGeneralBalanceResponseWithDefaults instantiates a new GeneralBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralBalanceResponseWithDefaults() *GeneralBalanceResponse {
	this := GeneralBalanceResponse{}
	return &this
}

// GetBalance returns the Balance field value
func (o *GeneralBalanceResponse) GetBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *GeneralBalanceResponse) GetBalanceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *GeneralBalanceResponse) SetBalance(v float32) {
	o.Balance = v
}

// GetIntegerBalance returns the IntegerBalance field value if set, zero value otherwise.
func (o *GeneralBalanceResponse) GetIntegerBalance() float32 {
	if o == nil || o.IntegerBalance == nil {
		var ret float32
		return ret
	}
	return *o.IntegerBalance
}

// GetIntegerBalanceOk returns a tuple with the IntegerBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralBalanceResponse) GetIntegerBalanceOk() (*float32, bool) {
	if o == nil || o.IntegerBalance == nil {
		return nil, false
	}
	return o.IntegerBalance, true
}

// HasIntegerBalance returns a boolean if a field has been set.
func (o *GeneralBalanceResponse) HasIntegerBalance() bool {
	if o != nil && o.IntegerBalance != nil {
		return true
	}

	return false
}

// SetIntegerBalance gets a reference to the given float32 and assigns it to the IntegerBalance field.
func (o *GeneralBalanceResponse) SetIntegerBalance(v float32) {
	o.IntegerBalance = &v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *GeneralBalanceResponse) GetDecimals() float32 {
	if o == nil || o.Decimals == nil {
		var ret float32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralBalanceResponse) GetDecimalsOk() (*float32, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *GeneralBalanceResponse) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given float32 and assigns it to the Decimals field.
func (o *GeneralBalanceResponse) SetDecimals(v float32) {
	o.Decimals = &v
}

// GetNetwork returns the Network field value
func (o *GeneralBalanceResponse) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *GeneralBalanceResponse) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *GeneralBalanceResponse) SetNetwork(v string) {
	o.Network = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *GeneralBalanceResponse) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralBalanceResponse) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *GeneralBalanceResponse) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *GeneralBalanceResponse) SetUnit(v string) {
	o.Unit = &v
}

func (o GeneralBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if o.IntegerBalance != nil {
		toSerialize["integer_balance"] = o.IntegerBalance
	}
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGeneralBalanceResponse struct {
	value *GeneralBalanceResponse
	isSet bool
}

func (v NullableGeneralBalanceResponse) Get() *GeneralBalanceResponse {
	return v.value
}

func (v *NullableGeneralBalanceResponse) Set(val *GeneralBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralBalanceResponse(val *GeneralBalanceResponse) *NullableGeneralBalanceResponse {
	return &NullableGeneralBalanceResponse{value: val, isSet: true}
}

func (v NullableGeneralBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


