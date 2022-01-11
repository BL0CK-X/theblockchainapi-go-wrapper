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

// UploadFileRequest struct for UploadFileRequest
type UploadFileRequest struct {
	UploadMethod *string `json:"upload_method,omitempty"`
}

// NewUploadFileRequest instantiates a new UploadFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadFileRequest() *UploadFileRequest {
	this := UploadFileRequest{}
	var uploadMethod string = "IPFS"
	this.UploadMethod = &uploadMethod
	return &this
}

// NewUploadFileRequestWithDefaults instantiates a new UploadFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadFileRequestWithDefaults() *UploadFileRequest {
	this := UploadFileRequest{}
	var uploadMethod string = "IPFS"
	this.UploadMethod = &uploadMethod
	return &this
}

// GetUploadMethod returns the UploadMethod field value if set, zero value otherwise.
func (o *UploadFileRequest) GetUploadMethod() string {
	if o == nil || o.UploadMethod == nil {
		var ret string
		return ret
	}
	return *o.UploadMethod
}

// GetUploadMethodOk returns a tuple with the UploadMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFileRequest) GetUploadMethodOk() (*string, bool) {
	if o == nil || o.UploadMethod == nil {
		return nil, false
	}
	return o.UploadMethod, true
}

// HasUploadMethod returns a boolean if a field has been set.
func (o *UploadFileRequest) HasUploadMethod() bool {
	if o != nil && o.UploadMethod != nil {
		return true
	}

	return false
}

// SetUploadMethod gets a reference to the given string and assigns it to the UploadMethod field.
func (o *UploadFileRequest) SetUploadMethod(v string) {
	o.UploadMethod = &v
}

func (o UploadFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UploadMethod != nil {
		toSerialize["upload_method"] = o.UploadMethod
	}
	return json.Marshal(toSerialize)
}

type NullableUploadFileRequest struct {
	value *UploadFileRequest
	isSet bool
}

func (v NullableUploadFileRequest) Get() *UploadFileRequest {
	return v.value
}

func (v *NullableUploadFileRequest) Set(val *UploadFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadFileRequest(val *UploadFileRequest) *NullableUploadFileRequest {
	return &NullableUploadFileRequest{value: val, isSet: true}
}

func (v NullableUploadFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


