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

// GetCandyMetadataRequest struct for GetCandyMetadataRequest
type GetCandyMetadataRequest struct {
	// The ID of the candy machine. This is the same as `config_address` for `v2` candy machines (supply either). 
	CandyMachineId *string `json:"candy_machine_id,omitempty"`
	// The configuration address of the candy machine. This is the same as `candy_machine_id` for `v2` candy machines (supply either). 
	ConfigAddress *string `json:"config_address,omitempty"`
	// The uuid of the candy machine. This is an alphanumeric string of length six (e.g., HpVdfP), which corresponds to the first six characters of the config_address. 
	Uuid *string `json:"uuid,omitempty"`
	Network *string `json:"network,omitempty"`
	// The candy machine contract of the candy machine for which you are retrieving the metadata. If you are providing `v1` candy machine ID, set this to `v1`. If you are providing `v2` candy machine ID, set this to `v2`. If you don't know which version your candy machine is, check out <a href=\"#operation/solanaGetAccountIsCandyMachine\">this endpoint</a>. 
	CandyMachineContractVersion *string `json:"candy_machine_contract_version,omitempty"`
}

// NewGetCandyMetadataRequest instantiates a new GetCandyMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCandyMetadataRequest() *GetCandyMetadataRequest {
	this := GetCandyMetadataRequest{}
	var network string = "devnet"
	this.Network = &network
	var candyMachineContractVersion string = "v1"
	this.CandyMachineContractVersion = &candyMachineContractVersion
	return &this
}

// NewGetCandyMetadataRequestWithDefaults instantiates a new GetCandyMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCandyMetadataRequestWithDefaults() *GetCandyMetadataRequest {
	this := GetCandyMetadataRequest{}
	var network string = "devnet"
	this.Network = &network
	var candyMachineContractVersion string = "v1"
	this.CandyMachineContractVersion = &candyMachineContractVersion
	return &this
}

// GetCandyMachineId returns the CandyMachineId field value if set, zero value otherwise.
func (o *GetCandyMetadataRequest) GetCandyMachineId() string {
	if o == nil || o.CandyMachineId == nil {
		var ret string
		return ret
	}
	return *o.CandyMachineId
}

// GetCandyMachineIdOk returns a tuple with the CandyMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataRequest) GetCandyMachineIdOk() (*string, bool) {
	if o == nil || o.CandyMachineId == nil {
		return nil, false
	}
	return o.CandyMachineId, true
}

// HasCandyMachineId returns a boolean if a field has been set.
func (o *GetCandyMetadataRequest) HasCandyMachineId() bool {
	if o != nil && o.CandyMachineId != nil {
		return true
	}

	return false
}

// SetCandyMachineId gets a reference to the given string and assigns it to the CandyMachineId field.
func (o *GetCandyMetadataRequest) SetCandyMachineId(v string) {
	o.CandyMachineId = &v
}

// GetConfigAddress returns the ConfigAddress field value if set, zero value otherwise.
func (o *GetCandyMetadataRequest) GetConfigAddress() string {
	if o == nil || o.ConfigAddress == nil {
		var ret string
		return ret
	}
	return *o.ConfigAddress
}

// GetConfigAddressOk returns a tuple with the ConfigAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataRequest) GetConfigAddressOk() (*string, bool) {
	if o == nil || o.ConfigAddress == nil {
		return nil, false
	}
	return o.ConfigAddress, true
}

// HasConfigAddress returns a boolean if a field has been set.
func (o *GetCandyMetadataRequest) HasConfigAddress() bool {
	if o != nil && o.ConfigAddress != nil {
		return true
	}

	return false
}

// SetConfigAddress gets a reference to the given string and assigns it to the ConfigAddress field.
func (o *GetCandyMetadataRequest) SetConfigAddress(v string) {
	o.ConfigAddress = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GetCandyMetadataRequest) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataRequest) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *GetCandyMetadataRequest) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GetCandyMetadataRequest) SetUuid(v string) {
	o.Uuid = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetCandyMetadataRequest) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataRequest) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetCandyMetadataRequest) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *GetCandyMetadataRequest) SetNetwork(v string) {
	o.Network = &v
}

// GetCandyMachineContractVersion returns the CandyMachineContractVersion field value if set, zero value otherwise.
func (o *GetCandyMetadataRequest) GetCandyMachineContractVersion() string {
	if o == nil || o.CandyMachineContractVersion == nil {
		var ret string
		return ret
	}
	return *o.CandyMachineContractVersion
}

// GetCandyMachineContractVersionOk returns a tuple with the CandyMachineContractVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataRequest) GetCandyMachineContractVersionOk() (*string, bool) {
	if o == nil || o.CandyMachineContractVersion == nil {
		return nil, false
	}
	return o.CandyMachineContractVersion, true
}

// HasCandyMachineContractVersion returns a boolean if a field has been set.
func (o *GetCandyMetadataRequest) HasCandyMachineContractVersion() bool {
	if o != nil && o.CandyMachineContractVersion != nil {
		return true
	}

	return false
}

// SetCandyMachineContractVersion gets a reference to the given string and assigns it to the CandyMachineContractVersion field.
func (o *GetCandyMetadataRequest) SetCandyMachineContractVersion(v string) {
	o.CandyMachineContractVersion = &v
}

func (o GetCandyMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CandyMachineId != nil {
		toSerialize["candy_machine_id"] = o.CandyMachineId
	}
	if o.ConfigAddress != nil {
		toSerialize["config_address"] = o.ConfigAddress
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.CandyMachineContractVersion != nil {
		toSerialize["candy_machine_contract_version"] = o.CandyMachineContractVersion
	}
	return json.Marshal(toSerialize)
}

type NullableGetCandyMetadataRequest struct {
	value *GetCandyMetadataRequest
	isSet bool
}

func (v NullableGetCandyMetadataRequest) Get() *GetCandyMetadataRequest {
	return v.value
}

func (v *NullableGetCandyMetadataRequest) Set(val *GetCandyMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCandyMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCandyMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCandyMetadataRequest(val *GetCandyMetadataRequest) *NullableGetCandyMetadataRequest {
	return &NullableGetCandyMetadataRequest{value: val, isSet: true}
}

func (v NullableGetCandyMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCandyMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


