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

// GeneralSecretRecoveryPhrase struct for GeneralSecretRecoveryPhrase
type GeneralSecretRecoveryPhrase struct {
	// The 12, 15, 18, 21, or 24 word phrase that can be used to derive many public key addresses (e.g., bottom army bless castle alter habit dish embody child flame smooth zone).  To derive a wallet identifier (e.g., public key, public address, etc.), you need a secret recovery phrase, a derivation path, and an optional passphrase.   Alternatively, you can derive a wallet identifier with a private key.  Default Lengths: - Avalanche (C): 24 - Avalanche (X, P): 24 - Binance: 12 - Ethereum: 12 - Near: 12 - Solana: 12  You can use `secret_recovery_phrase` for Chains: `avalanche`, `binance`, `ethereum`, `near`, `solana`.
	SecretRecoveryPhrase string `json:"secret_recovery_phrase"`
	// Derivation paths are used to derive the wallet identifier from the secret recovery phrase. Only correctly typed paths are accepted.  Defaults: - Avalanche (C): m/44'/60'/0'/0/0 - Avalanche (X, P): m/44'/60'/0'/0/0 - Binance: m/44'/60'/0'/0/0 - Ethereum: m/44'/60'/0'/0/0 - Near: m/44/397/0 - Solana: m/44/501/0/0  *Solana Behavior*  If you provide the empty string \"\" as the value for the derivation path, then we will derive your public key with the same behavior as the default behavior of the Solana CLI.  By default, we use \"m/44/501/0/0\". This is the path that the Phantom and Sollet wallets use.  You can also arbitrarily increment the default path (\"m/44/501/0/0\") to generate more wallets (e.g., \"m/44/501/0/1\", \"m/44/501/0/2\", ... AND/OR \"m/44/501/1/0\", \"m/44/501/2/0\", ...).  Phantom's Wallet increments the first digit (e.g., \"m/44/501/0/0\", \"m/44/501/1/0\", \"m/44/501/2/0\", ...) to generate more public key addresses.  The SolFlare recommended path is \"m/44/501/0\".  To learn more about derivation paths, check out <a href=\"https://learnmeabitcoin.com/technical/derivation-paths\" target=\"_blank\">this tutorial</a>.
	DerivationPath *string `json:"derivation_path,omitempty"`
	// PASSPHRASE != PASSWORD. This is NOT your Phantom password or any other password. It is an optional string you use when creating a wallet. This provides an additional layer of security because a hacker would need both the secret recovery phrase and the passphrase to access the output public key. By default, most wallet UI extensions do not use a passphrase. (You probably did not use a passphrase.) Limited to 500 characters. 
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewGeneralSecretRecoveryPhrase instantiates a new GeneralSecretRecoveryPhrase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralSecretRecoveryPhrase(secretRecoveryPhrase string) *GeneralSecretRecoveryPhrase {
	this := GeneralSecretRecoveryPhrase{}
	this.SecretRecoveryPhrase = secretRecoveryPhrase
	var passphrase string = ""
	this.Passphrase = &passphrase
	return &this
}

// NewGeneralSecretRecoveryPhraseWithDefaults instantiates a new GeneralSecretRecoveryPhrase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralSecretRecoveryPhraseWithDefaults() *GeneralSecretRecoveryPhrase {
	this := GeneralSecretRecoveryPhrase{}
	var passphrase string = ""
	this.Passphrase = &passphrase
	return &this
}

// GetSecretRecoveryPhrase returns the SecretRecoveryPhrase field value
func (o *GeneralSecretRecoveryPhrase) GetSecretRecoveryPhrase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretRecoveryPhrase
}

// GetSecretRecoveryPhraseOk returns a tuple with the SecretRecoveryPhrase field value
// and a boolean to check if the value has been set.
func (o *GeneralSecretRecoveryPhrase) GetSecretRecoveryPhraseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecretRecoveryPhrase, true
}

// SetSecretRecoveryPhrase sets field value
func (o *GeneralSecretRecoveryPhrase) SetSecretRecoveryPhrase(v string) {
	o.SecretRecoveryPhrase = v
}

// GetDerivationPath returns the DerivationPath field value if set, zero value otherwise.
func (o *GeneralSecretRecoveryPhrase) GetDerivationPath() string {
	if o == nil || o.DerivationPath == nil {
		var ret string
		return ret
	}
	return *o.DerivationPath
}

// GetDerivationPathOk returns a tuple with the DerivationPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSecretRecoveryPhrase) GetDerivationPathOk() (*string, bool) {
	if o == nil || o.DerivationPath == nil {
		return nil, false
	}
	return o.DerivationPath, true
}

// HasDerivationPath returns a boolean if a field has been set.
func (o *GeneralSecretRecoveryPhrase) HasDerivationPath() bool {
	if o != nil && o.DerivationPath != nil {
		return true
	}

	return false
}

// SetDerivationPath gets a reference to the given string and assigns it to the DerivationPath field.
func (o *GeneralSecretRecoveryPhrase) SetDerivationPath(v string) {
	o.DerivationPath = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *GeneralSecretRecoveryPhrase) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralSecretRecoveryPhrase) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *GeneralSecretRecoveryPhrase) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *GeneralSecretRecoveryPhrase) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o GeneralSecretRecoveryPhrase) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableGeneralSecretRecoveryPhrase struct {
	value *GeneralSecretRecoveryPhrase
	isSet bool
}

func (v NullableGeneralSecretRecoveryPhrase) Get() *GeneralSecretRecoveryPhrase {
	return v.value
}

func (v *NullableGeneralSecretRecoveryPhrase) Set(val *GeneralSecretRecoveryPhrase) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralSecretRecoveryPhrase) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralSecretRecoveryPhrase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralSecretRecoveryPhrase(val *GeneralSecretRecoveryPhrase) *NullableGeneralSecretRecoveryPhrase {
	return &NullableGeneralSecretRecoveryPhrase{value: val, isSet: true}
}

func (v NullableGeneralSecretRecoveryPhrase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralSecretRecoveryPhrase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


