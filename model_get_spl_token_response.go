/*
The Blockchain API

# About  Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly.   You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>  # How to Use the API  To use the API, you simply need to create an API key pair.  Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.blockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.  # Feature Requests  Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@blockchainapi.com).  # Contact  <figure>     <img          width=\"40px\"         height=\"40px\"          src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"     />     <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption> </figure> <a href=\"mailto:info@blockchainapi.com\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"         />         <figcaption style=\"textAlign:center;\">Email us: info@theblockchainapi.com</figcaption>     </figure> </a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"         />         <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>     </figure> </a> <a href=\"https://twitter.com/_BlockX_\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"         />          <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>     </figure> </a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"         />         <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>     </figure> </a>  # Security  Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet.   <b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b>   How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.  We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.  Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.  Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.  # Pricing  <b>Each user receives 500 free credits.</b>  You can learn more about our pricing <a href=\"https://dashboard.blockchainapi.com/billing\" target=\"_blank\">here</a>.   We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.  If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.  # SDKs / API Wrappers  We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.  We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.  `pip install theblockchainapi`  We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-javascript-wrapper\" target=\"_blank\">JavaScript Wrapper</a>.  `npm install theblockchainapi`  We also have auto-generated wrappers for the following languages: - <a href=\"https://github.com/BL0CK-X/theblockchainapi-go-wrapper\" target = \"_blank\">Go</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-java-wrapper\" target = \"_blank\">Java</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-kotlin-wrapper\" target = \"_blank\">Kotlin</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-php-wrapper\" target = \"_blank\">PHP</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-swift-wrapper\" target = \"_blank\">Swift5</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-typescript-wrapper\" target = \"_blank\">TypeScript</a>  If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.  If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

API version: null
Contact: info@blockchainapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetSPLTokenResponse struct for GetSPLTokenResponse
type GetSPLTokenResponse struct {
	// The number of decimals of the token. For example, if the USDC token has 6 decimals, then you need 1 * 10e6 = 1,000,000 tokens to have 1 USDC. The purpose of this is that everything must be stored as  an integer. Thus, if there are $100 USDC in total, there must be $100 * 10e6 tokens in order for the $100 to be divisible into lower denominations than $1. 
	Decimals *float32 `json:"decimals,omitempty"`
	// Public key address
	FreezeAuthority *string `json:"freeze_authority,omitempty"`
	// Public key address
	MintAuthority *string `json:"mint_authority,omitempty"`
	IsInitialized *bool `json:"is_initialized,omitempty"`
	// The supply of the token
	Supply *string `json:"supply,omitempty"`
}

// NewGetSPLTokenResponse instantiates a new GetSPLTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSPLTokenResponse() *GetSPLTokenResponse {
	this := GetSPLTokenResponse{}
	return &this
}

// NewGetSPLTokenResponseWithDefaults instantiates a new GetSPLTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSPLTokenResponseWithDefaults() *GetSPLTokenResponse {
	this := GetSPLTokenResponse{}
	return &this
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *GetSPLTokenResponse) GetDecimals() float32 {
	if o == nil || o.Decimals == nil {
		var ret float32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSPLTokenResponse) GetDecimalsOk() (*float32, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *GetSPLTokenResponse) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given float32 and assigns it to the Decimals field.
func (o *GetSPLTokenResponse) SetDecimals(v float32) {
	o.Decimals = &v
}

// GetFreezeAuthority returns the FreezeAuthority field value if set, zero value otherwise.
func (o *GetSPLTokenResponse) GetFreezeAuthority() string {
	if o == nil || o.FreezeAuthority == nil {
		var ret string
		return ret
	}
	return *o.FreezeAuthority
}

// GetFreezeAuthorityOk returns a tuple with the FreezeAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSPLTokenResponse) GetFreezeAuthorityOk() (*string, bool) {
	if o == nil || o.FreezeAuthority == nil {
		return nil, false
	}
	return o.FreezeAuthority, true
}

// HasFreezeAuthority returns a boolean if a field has been set.
func (o *GetSPLTokenResponse) HasFreezeAuthority() bool {
	if o != nil && o.FreezeAuthority != nil {
		return true
	}

	return false
}

// SetFreezeAuthority gets a reference to the given string and assigns it to the FreezeAuthority field.
func (o *GetSPLTokenResponse) SetFreezeAuthority(v string) {
	o.FreezeAuthority = &v
}

// GetMintAuthority returns the MintAuthority field value if set, zero value otherwise.
func (o *GetSPLTokenResponse) GetMintAuthority() string {
	if o == nil || o.MintAuthority == nil {
		var ret string
		return ret
	}
	return *o.MintAuthority
}

// GetMintAuthorityOk returns a tuple with the MintAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSPLTokenResponse) GetMintAuthorityOk() (*string, bool) {
	if o == nil || o.MintAuthority == nil {
		return nil, false
	}
	return o.MintAuthority, true
}

// HasMintAuthority returns a boolean if a field has been set.
func (o *GetSPLTokenResponse) HasMintAuthority() bool {
	if o != nil && o.MintAuthority != nil {
		return true
	}

	return false
}

// SetMintAuthority gets a reference to the given string and assigns it to the MintAuthority field.
func (o *GetSPLTokenResponse) SetMintAuthority(v string) {
	o.MintAuthority = &v
}

// GetIsInitialized returns the IsInitialized field value if set, zero value otherwise.
func (o *GetSPLTokenResponse) GetIsInitialized() bool {
	if o == nil || o.IsInitialized == nil {
		var ret bool
		return ret
	}
	return *o.IsInitialized
}

// GetIsInitializedOk returns a tuple with the IsInitialized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSPLTokenResponse) GetIsInitializedOk() (*bool, bool) {
	if o == nil || o.IsInitialized == nil {
		return nil, false
	}
	return o.IsInitialized, true
}

// HasIsInitialized returns a boolean if a field has been set.
func (o *GetSPLTokenResponse) HasIsInitialized() bool {
	if o != nil && o.IsInitialized != nil {
		return true
	}

	return false
}

// SetIsInitialized gets a reference to the given bool and assigns it to the IsInitialized field.
func (o *GetSPLTokenResponse) SetIsInitialized(v bool) {
	o.IsInitialized = &v
}

// GetSupply returns the Supply field value if set, zero value otherwise.
func (o *GetSPLTokenResponse) GetSupply() string {
	if o == nil || o.Supply == nil {
		var ret string
		return ret
	}
	return *o.Supply
}

// GetSupplyOk returns a tuple with the Supply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSPLTokenResponse) GetSupplyOk() (*string, bool) {
	if o == nil || o.Supply == nil {
		return nil, false
	}
	return o.Supply, true
}

// HasSupply returns a boolean if a field has been set.
func (o *GetSPLTokenResponse) HasSupply() bool {
	if o != nil && o.Supply != nil {
		return true
	}

	return false
}

// SetSupply gets a reference to the given string and assigns it to the Supply field.
func (o *GetSPLTokenResponse) SetSupply(v string) {
	o.Supply = &v
}

func (o GetSPLTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	if o.FreezeAuthority != nil {
		toSerialize["freeze_authority"] = o.FreezeAuthority
	}
	if o.MintAuthority != nil {
		toSerialize["mint_authority"] = o.MintAuthority
	}
	if o.IsInitialized != nil {
		toSerialize["is_initialized"] = o.IsInitialized
	}
	if o.Supply != nil {
		toSerialize["supply"] = o.Supply
	}
	return json.Marshal(toSerialize)
}

type NullableGetSPLTokenResponse struct {
	value *GetSPLTokenResponse
	isSet bool
}

func (v NullableGetSPLTokenResponse) Get() *GetSPLTokenResponse {
	return v.value
}

func (v *NullableGetSPLTokenResponse) Set(val *GetSPLTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSPLTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSPLTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSPLTokenResponse(val *GetSPLTokenResponse) *NullableGetSPLTokenResponse {
	return &NullableGetSPLTokenResponse{value: val, isSet: true}
}

func (v NullableGetSPLTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSPLTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


