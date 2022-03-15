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

// StatItem struct for StatItem
type StatItem struct {
	// The anonymous ID of the user who called your endpoint 
	Uid *string `json:"uid,omitempty"`
	// The time stamp of the ten-minute interval 
	TimeStamp *float32 `json:"time_stamp,omitempty"`
	// The full path of the endpoint (e.g., `project_id`/`version`/`path`) 
	EndpointId *string `json:"endpoint_id,omitempty"`
	// The number of API calls by this user to this endpoint for the ten-minute interval 
	ApiCalls *float32 `json:"api_calls,omitempty"`
	// The version of the endpoint
	Version *string `json:"version,omitempty"`
	// The path of the endpoint (not inclusive of the `version` or `project_id`)
	Path *string `json:"path,omitempty"`
}

// NewStatItem instantiates a new StatItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatItem() *StatItem {
	this := StatItem{}
	return &this
}

// NewStatItemWithDefaults instantiates a new StatItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatItemWithDefaults() *StatItem {
	this := StatItem{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *StatItem) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatItem) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *StatItem) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *StatItem) SetUid(v string) {
	o.Uid = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *StatItem) GetTimeStamp() float32 {
	if o == nil || o.TimeStamp == nil {
		var ret float32
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatItem) GetTimeStampOk() (*float32, bool) {
	if o == nil || o.TimeStamp == nil {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *StatItem) HasTimeStamp() bool {
	if o != nil && o.TimeStamp != nil {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given float32 and assigns it to the TimeStamp field.
func (o *StatItem) SetTimeStamp(v float32) {
	o.TimeStamp = &v
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *StatItem) GetEndpointId() string {
	if o == nil || o.EndpointId == nil {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatItem) GetEndpointIdOk() (*string, bool) {
	if o == nil || o.EndpointId == nil {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *StatItem) HasEndpointId() bool {
	if o != nil && o.EndpointId != nil {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *StatItem) SetEndpointId(v string) {
	o.EndpointId = &v
}

// GetApiCalls returns the ApiCalls field value if set, zero value otherwise.
func (o *StatItem) GetApiCalls() float32 {
	if o == nil || o.ApiCalls == nil {
		var ret float32
		return ret
	}
	return *o.ApiCalls
}

// GetApiCallsOk returns a tuple with the ApiCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatItem) GetApiCallsOk() (*float32, bool) {
	if o == nil || o.ApiCalls == nil {
		return nil, false
	}
	return o.ApiCalls, true
}

// HasApiCalls returns a boolean if a field has been set.
func (o *StatItem) HasApiCalls() bool {
	if o != nil && o.ApiCalls != nil {
		return true
	}

	return false
}

// SetApiCalls gets a reference to the given float32 and assigns it to the ApiCalls field.
func (o *StatItem) SetApiCalls(v float32) {
	o.ApiCalls = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StatItem) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatItem) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StatItem) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StatItem) SetVersion(v string) {
	o.Version = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *StatItem) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatItem) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *StatItem) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *StatItem) SetPath(v string) {
	o.Path = &v
}

func (o StatItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.TimeStamp != nil {
		toSerialize["time_stamp"] = o.TimeStamp
	}
	if o.EndpointId != nil {
		toSerialize["endpoint_id"] = o.EndpointId
	}
	if o.ApiCalls != nil {
		toSerialize["api_calls"] = o.ApiCalls
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableStatItem struct {
	value *StatItem
	isSet bool
}

func (v NullableStatItem) Get() *StatItem {
	return v.value
}

func (v *NullableStatItem) Set(val *StatItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStatItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStatItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatItem(val *StatItem) *NullableStatItem {
	return &NullableStatItem{value: val, isSet: true}
}

func (v NullableStatItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


