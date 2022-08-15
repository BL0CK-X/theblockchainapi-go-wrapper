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

// TokenMetadataResponse struct for TokenMetadataResponse
type TokenMetadataResponse struct {
	// The symbol of the coin 
	Symbol *string `json:"symbol,omitempty"`
	// The name of the coin 
	Name *string `json:"name,omitempty"`
	Decimals *float32 `json:"decimals,omitempty"`
	// The `mint_authority` (Solana) or `master_minter` (Ethereum) 
	MinterBlockchainIdentifier *string `json:"minter_blockchain_identifier,omitempty"`
	TotalSupply *string `json:"total_supply,omitempty"`
}

// NewTokenMetadataResponse instantiates a new TokenMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenMetadataResponse() *TokenMetadataResponse {
	this := TokenMetadataResponse{}
	return &this
}

// NewTokenMetadataResponseWithDefaults instantiates a new TokenMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenMetadataResponseWithDefaults() *TokenMetadataResponse {
	this := TokenMetadataResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TokenMetadataResponse) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenMetadataResponse) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TokenMetadataResponse) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TokenMetadataResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TokenMetadataResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenMetadataResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TokenMetadataResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TokenMetadataResponse) SetName(v string) {
	o.Name = &v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *TokenMetadataResponse) GetDecimals() float32 {
	if o == nil || o.Decimals == nil {
		var ret float32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenMetadataResponse) GetDecimalsOk() (*float32, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *TokenMetadataResponse) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given float32 and assigns it to the Decimals field.
func (o *TokenMetadataResponse) SetDecimals(v float32) {
	o.Decimals = &v
}

// GetMinterBlockchainIdentifier returns the MinterBlockchainIdentifier field value if set, zero value otherwise.
func (o *TokenMetadataResponse) GetMinterBlockchainIdentifier() string {
	if o == nil || o.MinterBlockchainIdentifier == nil {
		var ret string
		return ret
	}
	return *o.MinterBlockchainIdentifier
}

// GetMinterBlockchainIdentifierOk returns a tuple with the MinterBlockchainIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenMetadataResponse) GetMinterBlockchainIdentifierOk() (*string, bool) {
	if o == nil || o.MinterBlockchainIdentifier == nil {
		return nil, false
	}
	return o.MinterBlockchainIdentifier, true
}

// HasMinterBlockchainIdentifier returns a boolean if a field has been set.
func (o *TokenMetadataResponse) HasMinterBlockchainIdentifier() bool {
	if o != nil && o.MinterBlockchainIdentifier != nil {
		return true
	}

	return false
}

// SetMinterBlockchainIdentifier gets a reference to the given string and assigns it to the MinterBlockchainIdentifier field.
func (o *TokenMetadataResponse) SetMinterBlockchainIdentifier(v string) {
	o.MinterBlockchainIdentifier = &v
}

// GetTotalSupply returns the TotalSupply field value if set, zero value otherwise.
func (o *TokenMetadataResponse) GetTotalSupply() string {
	if o == nil || o.TotalSupply == nil {
		var ret string
		return ret
	}
	return *o.TotalSupply
}

// GetTotalSupplyOk returns a tuple with the TotalSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenMetadataResponse) GetTotalSupplyOk() (*string, bool) {
	if o == nil || o.TotalSupply == nil {
		return nil, false
	}
	return o.TotalSupply, true
}

// HasTotalSupply returns a boolean if a field has been set.
func (o *TokenMetadataResponse) HasTotalSupply() bool {
	if o != nil && o.TotalSupply != nil {
		return true
	}

	return false
}

// SetTotalSupply gets a reference to the given string and assigns it to the TotalSupply field.
func (o *TokenMetadataResponse) SetTotalSupply(v string) {
	o.TotalSupply = &v
}

func (o TokenMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	if o.MinterBlockchainIdentifier != nil {
		toSerialize["minter_blockchain_identifier"] = o.MinterBlockchainIdentifier
	}
	if o.TotalSupply != nil {
		toSerialize["total_supply"] = o.TotalSupply
	}
	return json.Marshal(toSerialize)
}

type NullableTokenMetadataResponse struct {
	value *TokenMetadataResponse
	isSet bool
}

func (v NullableTokenMetadataResponse) Get() *TokenMetadataResponse {
	return v.value
}

func (v *NullableTokenMetadataResponse) Set(val *TokenMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenMetadataResponse(val *TokenMetadataResponse) *NullableTokenMetadataResponse {
	return &NullableTokenMetadataResponse{value: val, isSet: true}
}

func (v NullableTokenMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


