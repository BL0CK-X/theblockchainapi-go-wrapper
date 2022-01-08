/*
The Blockchain API

# About  Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly.   You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>  # How to Use the API  To use the API, you simply need to create an API key pair.  Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.theblockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.  # Feature Requests  Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@theblockchainapi.com).  # Contact  <figure>     <img          width=\"40px\"         height=\"40px\"          src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"     />     <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption> </figure> <a href=\"mailto:info@theblockchainapi.com\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"         />         <figcaption style=\"textAlign:center;\">Email us: info@theblockchainapi.com</figcaption>     </figure> </a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"         />         <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>     </figure> </a> <a href=\"https://twitter.com/_BlockX_\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"         />          <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>     </figure> </a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"         />         <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>     </figure> </a>  # Security  Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet.   <b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b>   How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.  We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.  Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.  Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.  # Pricing  <b>Each user receives 500 free credits.</b>  You can learn more about our pricing <a href=\"https://dashboard.theblockchainapi.com/billing\" target=\"_blank\">here</a>.   We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.  If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.  # SDKs / API Wrappers  We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.  We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.  `pip install theblockchainapi`  We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-javascript-wrapper\" target=\"_blank\">JavaScript Wrapper</a>.  `npm install theblockchainapi`  We also have auto-generated wrappers for the following languages: - <a href=\"https://github.com/BL0CK-X/theblockchainapi-go-wrapper\" target = \"_blank\">Go</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-java-wrapper\" target = \"_blank\">Java</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-kotlin-wrapper\" target = \"_blank\">Kotlin</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-php-wrapper\" target = \"_blank\">PHP</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-swift-wrapper\" target = \"_blank\">Swift5</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-typescript-wrapper\" target = \"_blank\">TypeScript</a>  If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.  If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

API version: null
Contact: info@theblockchainapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetAllNFTsResponse struct for GetAllNFTsResponse
type GetAllNFTsResponse struct {
	// The minted NFTs. Only filled in for `v1` candy machines. Left empty for `v2`.
	MintedNfts *[]GetAllNFTsResponseMintedNfts `json:"minted_nfts,omitempty"`
	// The unminted NFTs. Only filled in for `v1` candy machines. Left empty for `v2`.
	UnmintedNfts *[]GetAllNFTsResponseUnmintedNfts `json:"unminted_nfts,omitempty"`
	// The list of all NFTs. Filled in for both `v1` and `v2` NFTs.
	AllNfts *[]GetAllNFTsResponseUnmintedNfts `json:"all_nfts,omitempty"`
	// Whether or not the division of NFTs among minted and unminted is accurate. If it is not accurate, then it is possible that NFTs have been included in the `minted` list that are not actually minted. If it is accurate, then the split of  minted and unminted is correct. `v1` candy machines always return a correct minted/unminted split.  
	Accurate *bool `json:"accurate,omitempty"`
}

// NewGetAllNFTsResponse instantiates a new GetAllNFTsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllNFTsResponse() *GetAllNFTsResponse {
	this := GetAllNFTsResponse{}
	return &this
}

// NewGetAllNFTsResponseWithDefaults instantiates a new GetAllNFTsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllNFTsResponseWithDefaults() *GetAllNFTsResponse {
	this := GetAllNFTsResponse{}
	return &this
}

// GetMintedNfts returns the MintedNfts field value if set, zero value otherwise.
func (o *GetAllNFTsResponse) GetMintedNfts() []GetAllNFTsResponseMintedNfts {
	if o == nil || o.MintedNfts == nil {
		var ret []GetAllNFTsResponseMintedNfts
		return ret
	}
	return *o.MintedNfts
}

// GetMintedNftsOk returns a tuple with the MintedNfts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllNFTsResponse) GetMintedNftsOk() (*[]GetAllNFTsResponseMintedNfts, bool) {
	if o == nil || o.MintedNfts == nil {
		return nil, false
	}
	return o.MintedNfts, true
}

// HasMintedNfts returns a boolean if a field has been set.
func (o *GetAllNFTsResponse) HasMintedNfts() bool {
	if o != nil && o.MintedNfts != nil {
		return true
	}

	return false
}

// SetMintedNfts gets a reference to the given []GetAllNFTsResponseMintedNfts and assigns it to the MintedNfts field.
func (o *GetAllNFTsResponse) SetMintedNfts(v []GetAllNFTsResponseMintedNfts) {
	o.MintedNfts = &v
}

// GetUnmintedNfts returns the UnmintedNfts field value if set, zero value otherwise.
func (o *GetAllNFTsResponse) GetUnmintedNfts() []GetAllNFTsResponseUnmintedNfts {
	if o == nil || o.UnmintedNfts == nil {
		var ret []GetAllNFTsResponseUnmintedNfts
		return ret
	}
	return *o.UnmintedNfts
}

// GetUnmintedNftsOk returns a tuple with the UnmintedNfts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllNFTsResponse) GetUnmintedNftsOk() (*[]GetAllNFTsResponseUnmintedNfts, bool) {
	if o == nil || o.UnmintedNfts == nil {
		return nil, false
	}
	return o.UnmintedNfts, true
}

// HasUnmintedNfts returns a boolean if a field has been set.
func (o *GetAllNFTsResponse) HasUnmintedNfts() bool {
	if o != nil && o.UnmintedNfts != nil {
		return true
	}

	return false
}

// SetUnmintedNfts gets a reference to the given []GetAllNFTsResponseUnmintedNfts and assigns it to the UnmintedNfts field.
func (o *GetAllNFTsResponse) SetUnmintedNfts(v []GetAllNFTsResponseUnmintedNfts) {
	o.UnmintedNfts = &v
}

// GetAllNfts returns the AllNfts field value if set, zero value otherwise.
func (o *GetAllNFTsResponse) GetAllNfts() []GetAllNFTsResponseUnmintedNfts {
	if o == nil || o.AllNfts == nil {
		var ret []GetAllNFTsResponseUnmintedNfts
		return ret
	}
	return *o.AllNfts
}

// GetAllNftsOk returns a tuple with the AllNfts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllNFTsResponse) GetAllNftsOk() (*[]GetAllNFTsResponseUnmintedNfts, bool) {
	if o == nil || o.AllNfts == nil {
		return nil, false
	}
	return o.AllNfts, true
}

// HasAllNfts returns a boolean if a field has been set.
func (o *GetAllNFTsResponse) HasAllNfts() bool {
	if o != nil && o.AllNfts != nil {
		return true
	}

	return false
}

// SetAllNfts gets a reference to the given []GetAllNFTsResponseUnmintedNfts and assigns it to the AllNfts field.
func (o *GetAllNFTsResponse) SetAllNfts(v []GetAllNFTsResponseUnmintedNfts) {
	o.AllNfts = &v
}

// GetAccurate returns the Accurate field value if set, zero value otherwise.
func (o *GetAllNFTsResponse) GetAccurate() bool {
	if o == nil || o.Accurate == nil {
		var ret bool
		return ret
	}
	return *o.Accurate
}

// GetAccurateOk returns a tuple with the Accurate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllNFTsResponse) GetAccurateOk() (*bool, bool) {
	if o == nil || o.Accurate == nil {
		return nil, false
	}
	return o.Accurate, true
}

// HasAccurate returns a boolean if a field has been set.
func (o *GetAllNFTsResponse) HasAccurate() bool {
	if o != nil && o.Accurate != nil {
		return true
	}

	return false
}

// SetAccurate gets a reference to the given bool and assigns it to the Accurate field.
func (o *GetAllNFTsResponse) SetAccurate(v bool) {
	o.Accurate = &v
}

func (o GetAllNFTsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MintedNfts != nil {
		toSerialize["minted_nfts"] = o.MintedNfts
	}
	if o.UnmintedNfts != nil {
		toSerialize["unminted_nfts"] = o.UnmintedNfts
	}
	if o.AllNfts != nil {
		toSerialize["all_nfts"] = o.AllNfts
	}
	if o.Accurate != nil {
		toSerialize["accurate"] = o.Accurate
	}
	return json.Marshal(toSerialize)
}

type NullableGetAllNFTsResponse struct {
	value *GetAllNFTsResponse
	isSet bool
}

func (v NullableGetAllNFTsResponse) Get() *GetAllNFTsResponse {
	return v.value
}

func (v *NullableGetAllNFTsResponse) Set(val *GetAllNFTsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllNFTsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllNFTsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllNFTsResponse(val *GetAllNFTsResponse) *NullableGetAllNFTsResponse {
	return &NullableGetAllNFTsResponse{value: val, isSet: true}
}

func (v NullableGetAllNFTsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllNFTsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


