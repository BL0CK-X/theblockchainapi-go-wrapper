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

// NFTOwnerAdvancedResponse struct for NFTOwnerAdvancedResponse
type NFTOwnerAdvancedResponse struct {
	Contract *NFTOwnerAdvancedResponseContract `json:"contract,omitempty"`
	// The public key of the wallet that has true ownership over the provided NFT. If listed, it is the lister. If placed on a loan, it is the loanee. If staked, it is the staker. If burned, it is the burner. Etc. Returns `null` in the edge case that we were unable to find the true owner. If this happens, please report it to us and we will try to adapt for this edge case. 
	Owner *string `json:"owner,omitempty"`
	// The price of the NFT, if listed or loaned. If held, staked, or burned, this is null.
	Price *float32 `json:"price,omitempty"`
	// The state of the NFT
	State *string `json:"state,omitempty"`
}

// NewNFTOwnerAdvancedResponse instantiates a new NFTOwnerAdvancedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFTOwnerAdvancedResponse() *NFTOwnerAdvancedResponse {
	this := NFTOwnerAdvancedResponse{}
	return &this
}

// NewNFTOwnerAdvancedResponseWithDefaults instantiates a new NFTOwnerAdvancedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTOwnerAdvancedResponseWithDefaults() *NFTOwnerAdvancedResponse {
	this := NFTOwnerAdvancedResponse{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *NFTOwnerAdvancedResponse) GetContract() NFTOwnerAdvancedResponseContract {
	if o == nil || o.Contract == nil {
		var ret NFTOwnerAdvancedResponseContract
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTOwnerAdvancedResponse) GetContractOk() (*NFTOwnerAdvancedResponseContract, bool) {
	if o == nil || o.Contract == nil {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *NFTOwnerAdvancedResponse) HasContract() bool {
	if o != nil && o.Contract != nil {
		return true
	}

	return false
}

// SetContract gets a reference to the given NFTOwnerAdvancedResponseContract and assigns it to the Contract field.
func (o *NFTOwnerAdvancedResponse) SetContract(v NFTOwnerAdvancedResponseContract) {
	o.Contract = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *NFTOwnerAdvancedResponse) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTOwnerAdvancedResponse) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *NFTOwnerAdvancedResponse) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *NFTOwnerAdvancedResponse) SetOwner(v string) {
	o.Owner = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *NFTOwnerAdvancedResponse) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTOwnerAdvancedResponse) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *NFTOwnerAdvancedResponse) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *NFTOwnerAdvancedResponse) SetPrice(v float32) {
	o.Price = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NFTOwnerAdvancedResponse) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTOwnerAdvancedResponse) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NFTOwnerAdvancedResponse) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NFTOwnerAdvancedResponse) SetState(v string) {
	o.State = &v
}

func (o NFTOwnerAdvancedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contract != nil {
		toSerialize["contract"] = o.Contract
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableNFTOwnerAdvancedResponse struct {
	value *NFTOwnerAdvancedResponse
	isSet bool
}

func (v NullableNFTOwnerAdvancedResponse) Get() *NFTOwnerAdvancedResponse {
	return v.value
}

func (v *NullableNFTOwnerAdvancedResponse) Set(val *NFTOwnerAdvancedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNFTOwnerAdvancedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNFTOwnerAdvancedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFTOwnerAdvancedResponse(val *NFTOwnerAdvancedResponse) *NullableNFTOwnerAdvancedResponse {
	return &NullableNFTOwnerAdvancedResponse{value: val, isSet: true}
}

func (v NullableNFTOwnerAdvancedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFTOwnerAdvancedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


