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
	"fmt"
)

// GeneralWallet - The wallet authentication information used to sign and submit the transaction.  Click the `>` arrow next to \"wallet\" on the left to see more details. See our Security section <a href=\"#section/Security\">here</a>. 
type GeneralWallet struct {
	GeneralB58PrivateKey *GeneralB58PrivateKey
	GeneralPrivateKey *GeneralPrivateKey
	GeneralSecretRecoveryPhrase *GeneralSecretRecoveryPhrase
	HexPrivateKey *HexPrivateKey
}

// GeneralB58PrivateKeyAsGeneralWallet is a convenience function that returns GeneralB58PrivateKey wrapped in GeneralWallet
func GeneralB58PrivateKeyAsGeneralWallet(v *GeneralB58PrivateKey) GeneralWallet {
	return GeneralWallet{
		GeneralB58PrivateKey: v,
	}
}

// GeneralPrivateKeyAsGeneralWallet is a convenience function that returns GeneralPrivateKey wrapped in GeneralWallet
func GeneralPrivateKeyAsGeneralWallet(v *GeneralPrivateKey) GeneralWallet {
	return GeneralWallet{
		GeneralPrivateKey: v,
	}
}

// GeneralSecretRecoveryPhraseAsGeneralWallet is a convenience function that returns GeneralSecretRecoveryPhrase wrapped in GeneralWallet
func GeneralSecretRecoveryPhraseAsGeneralWallet(v *GeneralSecretRecoveryPhrase) GeneralWallet {
	return GeneralWallet{
		GeneralSecretRecoveryPhrase: v,
	}
}

// HexPrivateKeyAsGeneralWallet is a convenience function that returns HexPrivateKey wrapped in GeneralWallet
func HexPrivateKeyAsGeneralWallet(v *HexPrivateKey) GeneralWallet {
	return GeneralWallet{
		HexPrivateKey: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GeneralWallet) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GeneralB58PrivateKey
	err = newStrictDecoder(data).Decode(&dst.GeneralB58PrivateKey)
	if err == nil {
		jsonGeneralB58PrivateKey, _ := json.Marshal(dst.GeneralB58PrivateKey)
		if string(jsonGeneralB58PrivateKey) == "{}" { // empty struct
			dst.GeneralB58PrivateKey = nil
		} else {
			match++
		}
	} else {
		dst.GeneralB58PrivateKey = nil
	}

	// try to unmarshal data into GeneralPrivateKey
	err = newStrictDecoder(data).Decode(&dst.GeneralPrivateKey)
	if err == nil {
		jsonGeneralPrivateKey, _ := json.Marshal(dst.GeneralPrivateKey)
		if string(jsonGeneralPrivateKey) == "{}" { // empty struct
			dst.GeneralPrivateKey = nil
		} else {
			match++
		}
	} else {
		dst.GeneralPrivateKey = nil
	}

	// try to unmarshal data into GeneralSecretRecoveryPhrase
	err = newStrictDecoder(data).Decode(&dst.GeneralSecretRecoveryPhrase)
	if err == nil {
		jsonGeneralSecretRecoveryPhrase, _ := json.Marshal(dst.GeneralSecretRecoveryPhrase)
		if string(jsonGeneralSecretRecoveryPhrase) == "{}" { // empty struct
			dst.GeneralSecretRecoveryPhrase = nil
		} else {
			match++
		}
	} else {
		dst.GeneralSecretRecoveryPhrase = nil
	}

	// try to unmarshal data into HexPrivateKey
	err = newStrictDecoder(data).Decode(&dst.HexPrivateKey)
	if err == nil {
		jsonHexPrivateKey, _ := json.Marshal(dst.HexPrivateKey)
		if string(jsonHexPrivateKey) == "{}" { // empty struct
			dst.HexPrivateKey = nil
		} else {
			match++
		}
	} else {
		dst.HexPrivateKey = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GeneralB58PrivateKey = nil
		dst.GeneralPrivateKey = nil
		dst.GeneralSecretRecoveryPhrase = nil
		dst.HexPrivateKey = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GeneralWallet)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GeneralWallet)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GeneralWallet) MarshalJSON() ([]byte, error) {
	if src.GeneralB58PrivateKey != nil {
		return json.Marshal(&src.GeneralB58PrivateKey)
	}

	if src.GeneralPrivateKey != nil {
		return json.Marshal(&src.GeneralPrivateKey)
	}

	if src.GeneralSecretRecoveryPhrase != nil {
		return json.Marshal(&src.GeneralSecretRecoveryPhrase)
	}

	if src.HexPrivateKey != nil {
		return json.Marshal(&src.HexPrivateKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GeneralWallet) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GeneralB58PrivateKey != nil {
		return obj.GeneralB58PrivateKey
	}

	if obj.GeneralPrivateKey != nil {
		return obj.GeneralPrivateKey
	}

	if obj.GeneralSecretRecoveryPhrase != nil {
		return obj.GeneralSecretRecoveryPhrase
	}

	if obj.HexPrivateKey != nil {
		return obj.HexPrivateKey
	}

	// all schemas are nil
	return nil
}

type NullableGeneralWallet struct {
	value *GeneralWallet
	isSet bool
}

func (v NullableGeneralWallet) Get() *GeneralWallet {
	return v.value
}

func (v *NullableGeneralWallet) Set(val *GeneralWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralWallet(val *GeneralWallet) *NullableGeneralWallet {
	return &NullableGeneralWallet{value: val, isSet: true}
}

func (v NullableGeneralWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


