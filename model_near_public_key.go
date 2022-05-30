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

// NearPublicKey struct for NearPublicKey
type NearPublicKey struct {
	// The public key in hex form. This is used in the Near blockchain. (e.g., `0xdf96e3094ab33558dbe14e743aceded3779307b931e12bd6f526fe450db70910`) 
	HexPublicKey *string `json:"hex_public_key,omitempty"`
	// The public address in hex form of the wallet. This is commonly used in Ethereum, Binance, and Avalanche C-Chain. (e.g., `0xB2b2d42C3adA171633E36b427F062f85A642F453`) 
	HexPublicAddress *string `json:"hex_public_address,omitempty"`
	// The public key in base58 form. This is used in the Solana blockchain. (e.g., `3Gdu3Uf97jw4Kac8xHEm5Hqmob3BrvJhf3We8s9t5Q2c`) 
	PublicKey *string `json:"public_key,omitempty"`
}

// NewNearPublicKey instantiates a new NearPublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNearPublicKey() *NearPublicKey {
	this := NearPublicKey{}
	return &this
}

// NewNearPublicKeyWithDefaults instantiates a new NearPublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNearPublicKeyWithDefaults() *NearPublicKey {
	this := NearPublicKey{}
	return &this
}

// GetHexPublicKey returns the HexPublicKey field value if set, zero value otherwise.
func (o *NearPublicKey) GetHexPublicKey() string {
	if o == nil || o.HexPublicKey == nil {
		var ret string
		return ret
	}
	return *o.HexPublicKey
}

// GetHexPublicKeyOk returns a tuple with the HexPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NearPublicKey) GetHexPublicKeyOk() (*string, bool) {
	if o == nil || o.HexPublicKey == nil {
		return nil, false
	}
	return o.HexPublicKey, true
}

// HasHexPublicKey returns a boolean if a field has been set.
func (o *NearPublicKey) HasHexPublicKey() bool {
	if o != nil && o.HexPublicKey != nil {
		return true
	}

	return false
}

// SetHexPublicKey gets a reference to the given string and assigns it to the HexPublicKey field.
func (o *NearPublicKey) SetHexPublicKey(v string) {
	o.HexPublicKey = &v
}

// GetHexPublicAddress returns the HexPublicAddress field value if set, zero value otherwise.
func (o *NearPublicKey) GetHexPublicAddress() string {
	if o == nil || o.HexPublicAddress == nil {
		var ret string
		return ret
	}
	return *o.HexPublicAddress
}

// GetHexPublicAddressOk returns a tuple with the HexPublicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NearPublicKey) GetHexPublicAddressOk() (*string, bool) {
	if o == nil || o.HexPublicAddress == nil {
		return nil, false
	}
	return o.HexPublicAddress, true
}

// HasHexPublicAddress returns a boolean if a field has been set.
func (o *NearPublicKey) HasHexPublicAddress() bool {
	if o != nil && o.HexPublicAddress != nil {
		return true
	}

	return false
}

// SetHexPublicAddress gets a reference to the given string and assigns it to the HexPublicAddress field.
func (o *NearPublicKey) SetHexPublicAddress(v string) {
	o.HexPublicAddress = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *NearPublicKey) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NearPublicKey) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *NearPublicKey) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *NearPublicKey) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o NearPublicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HexPublicKey != nil {
		toSerialize["hex_public_key"] = o.HexPublicKey
	}
	if o.HexPublicAddress != nil {
		toSerialize["hex_public_address"] = o.HexPublicAddress
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableNearPublicKey struct {
	value *NearPublicKey
	isSet bool
}

func (v NullableNearPublicKey) Get() *NearPublicKey {
	return v.value
}

func (v *NullableNearPublicKey) Set(val *NearPublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableNearPublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableNearPublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNearPublicKey(val *NearPublicKey) *NullableNearPublicKey {
	return &NullableNearPublicKey{value: val, isSet: true}
}

func (v NullableNearPublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNearPublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


