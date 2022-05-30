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
	"fmt"
)

// WalletIdentifiers - The wallet indentification information. Click the `>` arrow next to \"wallet\" on the left to see more details. 
type WalletIdentifiers struct {
	AvalancheCChainPublicAddress *AvalancheCChainPublicAddress
	AvalancheXPChainPublicAddress *AvalancheXPChainPublicAddress
	BSCPublicAddress *BSCPublicAddress
	EthereumPublicAddress *EthereumPublicAddress
	NearPublicKey *NearPublicKey
	SolanaPublicKey *SolanaPublicKey
}

// AvalancheCChainPublicAddressAsWalletIdentifiers is a convenience function that returns AvalancheCChainPublicAddress wrapped in WalletIdentifiers
func AvalancheCChainPublicAddressAsWalletIdentifiers(v *AvalancheCChainPublicAddress) WalletIdentifiers {
	return WalletIdentifiers{
		AvalancheCChainPublicAddress: v,
	}
}

// AvalancheXPChainPublicAddressAsWalletIdentifiers is a convenience function that returns AvalancheXPChainPublicAddress wrapped in WalletIdentifiers
func AvalancheXPChainPublicAddressAsWalletIdentifiers(v *AvalancheXPChainPublicAddress) WalletIdentifiers {
	return WalletIdentifiers{
		AvalancheXPChainPublicAddress: v,
	}
}

// BSCPublicAddressAsWalletIdentifiers is a convenience function that returns BSCPublicAddress wrapped in WalletIdentifiers
func BSCPublicAddressAsWalletIdentifiers(v *BSCPublicAddress) WalletIdentifiers {
	return WalletIdentifiers{
		BSCPublicAddress: v,
	}
}

// EthereumPublicAddressAsWalletIdentifiers is a convenience function that returns EthereumPublicAddress wrapped in WalletIdentifiers
func EthereumPublicAddressAsWalletIdentifiers(v *EthereumPublicAddress) WalletIdentifiers {
	return WalletIdentifiers{
		EthereumPublicAddress: v,
	}
}

// NearPublicKeyAsWalletIdentifiers is a convenience function that returns NearPublicKey wrapped in WalletIdentifiers
func NearPublicKeyAsWalletIdentifiers(v *NearPublicKey) WalletIdentifiers {
	return WalletIdentifiers{
		NearPublicKey: v,
	}
}

// SolanaPublicKeyAsWalletIdentifiers is a convenience function that returns SolanaPublicKey wrapped in WalletIdentifiers
func SolanaPublicKeyAsWalletIdentifiers(v *SolanaPublicKey) WalletIdentifiers {
	return WalletIdentifiers{
		SolanaPublicKey: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WalletIdentifiers) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AvalancheCChainPublicAddress
	err = newStrictDecoder(data).Decode(&dst.AvalancheCChainPublicAddress)
	if err == nil {
		jsonAvalancheCChainPublicAddress, _ := json.Marshal(dst.AvalancheCChainPublicAddress)
		if string(jsonAvalancheCChainPublicAddress) == "{}" { // empty struct
			dst.AvalancheCChainPublicAddress = nil
		} else {
			match++
		}
	} else {
		dst.AvalancheCChainPublicAddress = nil
	}

	// try to unmarshal data into AvalancheXPChainPublicAddress
	err = newStrictDecoder(data).Decode(&dst.AvalancheXPChainPublicAddress)
	if err == nil {
		jsonAvalancheXPChainPublicAddress, _ := json.Marshal(dst.AvalancheXPChainPublicAddress)
		if string(jsonAvalancheXPChainPublicAddress) == "{}" { // empty struct
			dst.AvalancheXPChainPublicAddress = nil
		} else {
			match++
		}
	} else {
		dst.AvalancheXPChainPublicAddress = nil
	}

	// try to unmarshal data into BSCPublicAddress
	err = newStrictDecoder(data).Decode(&dst.BSCPublicAddress)
	if err == nil {
		jsonBSCPublicAddress, _ := json.Marshal(dst.BSCPublicAddress)
		if string(jsonBSCPublicAddress) == "{}" { // empty struct
			dst.BSCPublicAddress = nil
		} else {
			match++
		}
	} else {
		dst.BSCPublicAddress = nil
	}

	// try to unmarshal data into EthereumPublicAddress
	err = newStrictDecoder(data).Decode(&dst.EthereumPublicAddress)
	if err == nil {
		jsonEthereumPublicAddress, _ := json.Marshal(dst.EthereumPublicAddress)
		if string(jsonEthereumPublicAddress) == "{}" { // empty struct
			dst.EthereumPublicAddress = nil
		} else {
			match++
		}
	} else {
		dst.EthereumPublicAddress = nil
	}

	// try to unmarshal data into NearPublicKey
	err = newStrictDecoder(data).Decode(&dst.NearPublicKey)
	if err == nil {
		jsonNearPublicKey, _ := json.Marshal(dst.NearPublicKey)
		if string(jsonNearPublicKey) == "{}" { // empty struct
			dst.NearPublicKey = nil
		} else {
			match++
		}
	} else {
		dst.NearPublicKey = nil
	}

	// try to unmarshal data into SolanaPublicKey
	err = newStrictDecoder(data).Decode(&dst.SolanaPublicKey)
	if err == nil {
		jsonSolanaPublicKey, _ := json.Marshal(dst.SolanaPublicKey)
		if string(jsonSolanaPublicKey) == "{}" { // empty struct
			dst.SolanaPublicKey = nil
		} else {
			match++
		}
	} else {
		dst.SolanaPublicKey = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AvalancheCChainPublicAddress = nil
		dst.AvalancheXPChainPublicAddress = nil
		dst.BSCPublicAddress = nil
		dst.EthereumPublicAddress = nil
		dst.NearPublicKey = nil
		dst.SolanaPublicKey = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(WalletIdentifiers)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(WalletIdentifiers)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WalletIdentifiers) MarshalJSON() ([]byte, error) {
	if src.AvalancheCChainPublicAddress != nil {
		return json.Marshal(&src.AvalancheCChainPublicAddress)
	}

	if src.AvalancheXPChainPublicAddress != nil {
		return json.Marshal(&src.AvalancheXPChainPublicAddress)
	}

	if src.BSCPublicAddress != nil {
		return json.Marshal(&src.BSCPublicAddress)
	}

	if src.EthereumPublicAddress != nil {
		return json.Marshal(&src.EthereumPublicAddress)
	}

	if src.NearPublicKey != nil {
		return json.Marshal(&src.NearPublicKey)
	}

	if src.SolanaPublicKey != nil {
		return json.Marshal(&src.SolanaPublicKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WalletIdentifiers) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AvalancheCChainPublicAddress != nil {
		return obj.AvalancheCChainPublicAddress
	}

	if obj.AvalancheXPChainPublicAddress != nil {
		return obj.AvalancheXPChainPublicAddress
	}

	if obj.BSCPublicAddress != nil {
		return obj.BSCPublicAddress
	}

	if obj.EthereumPublicAddress != nil {
		return obj.EthereumPublicAddress
	}

	if obj.NearPublicKey != nil {
		return obj.NearPublicKey
	}

	if obj.SolanaPublicKey != nil {
		return obj.SolanaPublicKey
	}

	// all schemas are nil
	return nil
}

type NullableWalletIdentifiers struct {
	value *WalletIdentifiers
	isSet bool
}

func (v NullableWalletIdentifiers) Get() *WalletIdentifiers {
	return v.value
}

func (v *NullableWalletIdentifiers) Set(val *WalletIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletIdentifiers(val *WalletIdentifiers) *NullableWalletIdentifiers {
	return &NullableWalletIdentifiers{value: val, isSet: true}
}

func (v NullableWalletIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


