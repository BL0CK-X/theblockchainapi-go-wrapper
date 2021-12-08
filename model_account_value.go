/*
The Blockchain API

# About  Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly.   You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>  # How to Use the API  To use the API, you simply need to create an API key pair.  Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.theblockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.  # Feature Requests  Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@theblockchainapi.com).  # Contact  <figure>     <img          width=\"40px\"         height=\"40px\"          src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"     />     <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption> </figure> <a href=\"mailto:info@theblockchainapi.com\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"         />         <figcaption style=\"textAlign:center;\">Email us: info@theblockchainapi.com</figcaption>     </figure> </a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"         />         <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>     </figure> </a> <a href=\"https://twitter.com/_BlockX_\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"         />          <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>     </figure> </a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"         />         <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>     </figure> </a>  # Security  Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet.   <b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b>   How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.  We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.  Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.  Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.  # Pricing  <b>Each user receives 500 free credits every month.</b>  You can learn more about our pricing <a href=\"https://dashboard.theblockchainapi.com/billing\" target=\"_blank\">here</a>.   We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.  If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.  # SDKs / API Wrappers  We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.  We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.  `pip install theblockchainapi`  We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/javascript\" target=\"_blank\">JavaScript Wrapper</a>.  `npm install theblockchainapi`  We also have auto-generated wrappers for the following languages: - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">Go</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/java\" target = \"_blank\">Java</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">Kotlin</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">PHP</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">Swift5</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">TypeScript</a>  If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.  If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

API version: null
Contact: info@theblockchainapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AccountValue struct for AccountValue
type AccountValue struct {
	// More info about the account. What are included depends on the type of account. See examples.
	Data *map[string]interface{} `json:"data,omitempty"`
	// Whether or not this account is marked as executable
	Executable *bool `json:"executable,omitempty"`
	// The owner of the account
	Owner *string `json:"owner,omitempty"`
	RentEpoch *float32 `json:"rent_epoch,omitempty"`
}

// NewAccountValue instantiates a new AccountValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountValue() *AccountValue {
	this := AccountValue{}
	return &this
}

// NewAccountValueWithDefaults instantiates a new AccountValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountValueWithDefaults() *AccountValue {
	this := AccountValue{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AccountValue) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValue) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AccountValue) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *AccountValue) SetData(v map[string]interface{}) {
	o.Data = &v
}

// GetExecutable returns the Executable field value if set, zero value otherwise.
func (o *AccountValue) GetExecutable() bool {
	if o == nil || o.Executable == nil {
		var ret bool
		return ret
	}
	return *o.Executable
}

// GetExecutableOk returns a tuple with the Executable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValue) GetExecutableOk() (*bool, bool) {
	if o == nil || o.Executable == nil {
		return nil, false
	}
	return o.Executable, true
}

// HasExecutable returns a boolean if a field has been set.
func (o *AccountValue) HasExecutable() bool {
	if o != nil && o.Executable != nil {
		return true
	}

	return false
}

// SetExecutable gets a reference to the given bool and assigns it to the Executable field.
func (o *AccountValue) SetExecutable(v bool) {
	o.Executable = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AccountValue) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValue) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *AccountValue) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *AccountValue) SetOwner(v string) {
	o.Owner = &v
}

// GetRentEpoch returns the RentEpoch field value if set, zero value otherwise.
func (o *AccountValue) GetRentEpoch() float32 {
	if o == nil || o.RentEpoch == nil {
		var ret float32
		return ret
	}
	return *o.RentEpoch
}

// GetRentEpochOk returns a tuple with the RentEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountValue) GetRentEpochOk() (*float32, bool) {
	if o == nil || o.RentEpoch == nil {
		return nil, false
	}
	return o.RentEpoch, true
}

// HasRentEpoch returns a boolean if a field has been set.
func (o *AccountValue) HasRentEpoch() bool {
	if o != nil && o.RentEpoch != nil {
		return true
	}

	return false
}

// SetRentEpoch gets a reference to the given float32 and assigns it to the RentEpoch field.
func (o *AccountValue) SetRentEpoch(v float32) {
	o.RentEpoch = &v
}

func (o AccountValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Executable != nil {
		toSerialize["executable"] = o.Executable
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.RentEpoch != nil {
		toSerialize["rent_epoch"] = o.RentEpoch
	}
	return json.Marshal(toSerialize)
}

type NullableAccountValue struct {
	value *AccountValue
	isSet bool
}

func (v NullableAccountValue) Get() *AccountValue {
	return v.value
}

func (v *NullableAccountValue) Set(val *AccountValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountValue(val *AccountValue) *NullableAccountValue {
	return &NullableAccountValue{value: val, isSet: true}
}

func (v NullableAccountValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


