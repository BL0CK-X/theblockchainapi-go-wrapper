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

// CCWebhook struct for CCWebhook
type CCWebhook struct {
	// The unique identifier of the webhook 
	WebhookId *string `json:"webhook_id,omitempty"`
	// A signature from the Blockchain API that verifies the webhook is from us. An array of integers. 
	Signature map[string]interface{} `json:"signature,omitempty"`
	// The time we sent the webhook 
	TimeSent *float32 `json:"time_sent,omitempty"`
	Data *CCPayment `json:"data,omitempty"`
}

// NewCCWebhook instantiates a new CCWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCWebhook() *CCWebhook {
	this := CCWebhook{}
	return &this
}

// NewCCWebhookWithDefaults instantiates a new CCWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCWebhookWithDefaults() *CCWebhook {
	this := CCWebhook{}
	return &this
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *CCWebhook) GetWebhookId() string {
	if o == nil || o.WebhookId == nil {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCWebhook) GetWebhookIdOk() (*string, bool) {
	if o == nil || o.WebhookId == nil {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *CCWebhook) HasWebhookId() bool {
	if o != nil && o.WebhookId != nil {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *CCWebhook) SetWebhookId(v string) {
	o.WebhookId = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *CCWebhook) GetSignature() map[string]interface{} {
	if o == nil || o.Signature == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCWebhook) GetSignatureOk() (map[string]interface{}, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *CCWebhook) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given map[string]interface{} and assigns it to the Signature field.
func (o *CCWebhook) SetSignature(v map[string]interface{}) {
	o.Signature = v
}

// GetTimeSent returns the TimeSent field value if set, zero value otherwise.
func (o *CCWebhook) GetTimeSent() float32 {
	if o == nil || o.TimeSent == nil {
		var ret float32
		return ret
	}
	return *o.TimeSent
}

// GetTimeSentOk returns a tuple with the TimeSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCWebhook) GetTimeSentOk() (*float32, bool) {
	if o == nil || o.TimeSent == nil {
		return nil, false
	}
	return o.TimeSent, true
}

// HasTimeSent returns a boolean if a field has been set.
func (o *CCWebhook) HasTimeSent() bool {
	if o != nil && o.TimeSent != nil {
		return true
	}

	return false
}

// SetTimeSent gets a reference to the given float32 and assigns it to the TimeSent field.
func (o *CCWebhook) SetTimeSent(v float32) {
	o.TimeSent = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CCWebhook) GetData() CCPayment {
	if o == nil || o.Data == nil {
		var ret CCPayment
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCWebhook) GetDataOk() (*CCPayment, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CCWebhook) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CCPayment and assigns it to the Data field.
func (o *CCWebhook) SetData(v CCPayment) {
	o.Data = &v
}

func (o CCWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WebhookId != nil {
		toSerialize["webhook_id"] = o.WebhookId
	}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}
	if o.TimeSent != nil {
		toSerialize["time_sent"] = o.TimeSent
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCCWebhook struct {
	value *CCWebhook
	isSet bool
}

func (v NullableCCWebhook) Get() *CCWebhook {
	return v.value
}

func (v *NullableCCWebhook) Set(val *CCWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableCCWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableCCWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCWebhook(val *CCWebhook) *NullableCCWebhook {
	return &NullableCCWebhook{value: val, isSet: true}
}

func (v NullableCCWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


