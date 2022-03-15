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

// NFTAnalyticsResponse struct for NFTAnalyticsResponse
type NFTAnalyticsResponse struct {
	// The start time used in the API request. 
	StartTime interface{} `json:"start_time,omitempty"`
	// The end time used in the API request.  
	EndTime interface{} `json:"end_time,omitempty"`
	TransactionHistory *NFTAnalyticsResponseTransactionHistory `json:"transaction_history,omitempty"`
	// The minimum active listing price for the collection in the given time period. The listing must have been processed before `end_time` and still active (not delisted or purchased) by `end_time` in order to affect the floor calculation. 
	Floor *float32 `json:"floor,omitempty"`
	// The sum of the sale prices for the given time period. 
	Volume *float32 `json:"volume,omitempty"`
}

// NewNFTAnalyticsResponse instantiates a new NFTAnalyticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFTAnalyticsResponse() *NFTAnalyticsResponse {
	this := NFTAnalyticsResponse{}
	return &this
}

// NewNFTAnalyticsResponseWithDefaults instantiates a new NFTAnalyticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTAnalyticsResponseWithDefaults() *NFTAnalyticsResponse {
	this := NFTAnalyticsResponse{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NFTAnalyticsResponse) GetStartTime() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NFTAnalyticsResponse) GetStartTimeOk() (*interface{}, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *NFTAnalyticsResponse) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given interface{} and assigns it to the StartTime field.
func (o *NFTAnalyticsResponse) SetStartTime(v interface{}) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NFTAnalyticsResponse) GetEndTime() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NFTAnalyticsResponse) GetEndTimeOk() (*interface{}, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *NFTAnalyticsResponse) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given interface{} and assigns it to the EndTime field.
func (o *NFTAnalyticsResponse) SetEndTime(v interface{}) {
	o.EndTime = v
}

// GetTransactionHistory returns the TransactionHistory field value if set, zero value otherwise.
func (o *NFTAnalyticsResponse) GetTransactionHistory() NFTAnalyticsResponseTransactionHistory {
	if o == nil || o.TransactionHistory == nil {
		var ret NFTAnalyticsResponseTransactionHistory
		return ret
	}
	return *o.TransactionHistory
}

// GetTransactionHistoryOk returns a tuple with the TransactionHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTAnalyticsResponse) GetTransactionHistoryOk() (*NFTAnalyticsResponseTransactionHistory, bool) {
	if o == nil || o.TransactionHistory == nil {
		return nil, false
	}
	return o.TransactionHistory, true
}

// HasTransactionHistory returns a boolean if a field has been set.
func (o *NFTAnalyticsResponse) HasTransactionHistory() bool {
	if o != nil && o.TransactionHistory != nil {
		return true
	}

	return false
}

// SetTransactionHistory gets a reference to the given NFTAnalyticsResponseTransactionHistory and assigns it to the TransactionHistory field.
func (o *NFTAnalyticsResponse) SetTransactionHistory(v NFTAnalyticsResponseTransactionHistory) {
	o.TransactionHistory = &v
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *NFTAnalyticsResponse) GetFloor() float32 {
	if o == nil || o.Floor == nil {
		var ret float32
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTAnalyticsResponse) GetFloorOk() (*float32, bool) {
	if o == nil || o.Floor == nil {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *NFTAnalyticsResponse) HasFloor() bool {
	if o != nil && o.Floor != nil {
		return true
	}

	return false
}

// SetFloor gets a reference to the given float32 and assigns it to the Floor field.
func (o *NFTAnalyticsResponse) SetFloor(v float32) {
	o.Floor = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *NFTAnalyticsResponse) GetVolume() float32 {
	if o == nil || o.Volume == nil {
		var ret float32
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTAnalyticsResponse) GetVolumeOk() (*float32, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *NFTAnalyticsResponse) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float32 and assigns it to the Volume field.
func (o *NFTAnalyticsResponse) SetVolume(v float32) {
	o.Volume = &v
}

func (o NFTAnalyticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.TransactionHistory != nil {
		toSerialize["transaction_history"] = o.TransactionHistory
	}
	if o.Floor != nil {
		toSerialize["floor"] = o.Floor
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	return json.Marshal(toSerialize)
}

type NullableNFTAnalyticsResponse struct {
	value *NFTAnalyticsResponse
	isSet bool
}

func (v NullableNFTAnalyticsResponse) Get() *NFTAnalyticsResponse {
	return v.value
}

func (v *NullableNFTAnalyticsResponse) Set(val *NFTAnalyticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNFTAnalyticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNFTAnalyticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFTAnalyticsResponse(val *NFTAnalyticsResponse) *NullableNFTAnalyticsResponse {
	return &NullableNFTAnalyticsResponse{value: val, isSet: true}
}

func (v NullableNFTAnalyticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFTAnalyticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


