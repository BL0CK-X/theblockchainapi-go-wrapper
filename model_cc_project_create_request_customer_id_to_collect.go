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

// CCProjectCreateRequestCustomerIdToCollect The customer ID to collect during checkout. This could be a Twitter handle, a Telegram handle -- anything you use to identify your customers. We will collect it before they pay and link the payment to this ID. If you do not provide this, we will only be able to link the payment through either the customer's payment validation code, or their wallet's public key. If the users are signing in with their wallet in your application, then you can simply check whether they have paid through that. 
type CCProjectCreateRequestCustomerIdToCollect struct {
	// What type of identifier you are collecting, either an \"email\" or \"misc\"ellanous. Miscellanous simply means you are collecting something other than an email.  
	IdType string `json:"id_type"`
	// The name of the customer ID input presented to the user 
	Name string `json:"name"`
	// The description / stated purpose of the customer ID input presented to the user 
	Description *string `json:"description,omitempty"`
	// Whether to require the customer ID 
	Required *bool `json:"required,omitempty"`
}

// NewCCProjectCreateRequestCustomerIdToCollect instantiates a new CCProjectCreateRequestCustomerIdToCollect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCProjectCreateRequestCustomerIdToCollect(idType string, name string) *CCProjectCreateRequestCustomerIdToCollect {
	this := CCProjectCreateRequestCustomerIdToCollect{}
	this.IdType = idType
	this.Name = name
	var description string = ""
	this.Description = &description
	var required bool = false
	this.Required = &required
	return &this
}

// NewCCProjectCreateRequestCustomerIdToCollectWithDefaults instantiates a new CCProjectCreateRequestCustomerIdToCollect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCProjectCreateRequestCustomerIdToCollectWithDefaults() *CCProjectCreateRequestCustomerIdToCollect {
	this := CCProjectCreateRequestCustomerIdToCollect{}
	var description string = ""
	this.Description = &description
	var required bool = false
	this.Required = &required
	return &this
}

// GetIdType returns the IdType field value
func (o *CCProjectCreateRequestCustomerIdToCollect) GetIdType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value
// and a boolean to check if the value has been set.
func (o *CCProjectCreateRequestCustomerIdToCollect) GetIdTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdType, true
}

// SetIdType sets field value
func (o *CCProjectCreateRequestCustomerIdToCollect) SetIdType(v string) {
	o.IdType = v
}

// GetName returns the Name field value
func (o *CCProjectCreateRequestCustomerIdToCollect) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CCProjectCreateRequestCustomerIdToCollect) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CCProjectCreateRequestCustomerIdToCollect) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CCProjectCreateRequestCustomerIdToCollect) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProjectCreateRequestCustomerIdToCollect) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CCProjectCreateRequestCustomerIdToCollect) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CCProjectCreateRequestCustomerIdToCollect) SetDescription(v string) {
	o.Description = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *CCProjectCreateRequestCustomerIdToCollect) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProjectCreateRequestCustomerIdToCollect) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *CCProjectCreateRequestCustomerIdToCollect) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *CCProjectCreateRequestCustomerIdToCollect) SetRequired(v bool) {
	o.Required = &v
}

func (o CCProjectCreateRequestCustomerIdToCollect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id_type"] = o.IdType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	return json.Marshal(toSerialize)
}

type NullableCCProjectCreateRequestCustomerIdToCollect struct {
	value *CCProjectCreateRequestCustomerIdToCollect
	isSet bool
}

func (v NullableCCProjectCreateRequestCustomerIdToCollect) Get() *CCProjectCreateRequestCustomerIdToCollect {
	return v.value
}

func (v *NullableCCProjectCreateRequestCustomerIdToCollect) Set(val *CCProjectCreateRequestCustomerIdToCollect) {
	v.value = val
	v.isSet = true
}

func (v NullableCCProjectCreateRequestCustomerIdToCollect) IsSet() bool {
	return v.isSet
}

func (v *NullableCCProjectCreateRequestCustomerIdToCollect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCProjectCreateRequestCustomerIdToCollect(val *CCProjectCreateRequestCustomerIdToCollect) *NullableCCProjectCreateRequestCustomerIdToCollect {
	return &NullableCCProjectCreateRequestCustomerIdToCollect{value: val, isSet: true}
}

func (v NullableCCProjectCreateRequestCustomerIdToCollect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCProjectCreateRequestCustomerIdToCollect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


