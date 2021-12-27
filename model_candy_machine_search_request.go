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

// CandyMachineSearchRequest struct for CandyMachineSearchRequest
type CandyMachineSearchRequest struct {
	// The public key of the update authority of the candy machine
	UpdateAuthority *string `json:"update_authority,omitempty"`
	// Only `exact_match` supported at this time
	UpdateAuthoritySearchMethod *string `json:"update_authority_search_method,omitempty"`
	// The public key of the configuration of the candy machine
	ConfigAddress *string `json:"config_address,omitempty"`
	// Only `exact_match` supported at this time
	ConfigAddressSearchMethod *string `json:"config_address_search_method,omitempty"`
	// The alphanumeric string of length six that corresponds to the candy machine. This is NOT the candy machine ID.  This is the first six letters of the configuration address and is also used to identify the candy machine. An example is `4zKV6i`. 
	Uuid *string `json:"uuid,omitempty"`
	// Only `exact_match` supported at this time
	UuidSearchMethod *string `json:"uuid_search_method,omitempty"`
	// The symbol associated with the candy machine
	Symbol *string `json:"symbol,omitempty"`
	SymbolSearchMethod *string `json:"symbol_search_method,omitempty"`
	// The name of an NFT on the candy machine, minted or unminted. For example, to find The Solana Money Boys candy machine, you already know that each NFT is named \"Solana Money Boy #0\", \"Solana Money Boy #1\", and so on. So you could search with  nft_name=\"Solana Money Boy #0\", nft_name_index=0, nft_name_search_method='exact_match', for example, which would return the candy machine ID. This also works with candy machines that are not live but are uploaded. 
	NftName *string `json:"nft_name,omitempty"`
	// The index of the NFT to check, e.g., the 2nd NFT would have an index of 1. We cannot search all NFTs on a candy machine currently, so you must search an NFT at a particular position, such as the first, second, and so on. In general, nft_name_index=0 works for most use cases. 
	NftNameIndex *string `json:"nft_name_index,omitempty"`
	NftNameSearchMethod *string `json:"nft_name_search_method,omitempty"`
	Network *string `json:"network,omitempty"`
	// The candy machine contract you want to search.  If you want to search `v1` candy machines, set this to `v1`. If you want to search `v2` candy machines. set this to `v2`. 
	CandyMachineContractVersion *string `json:"candy_machine_contract_version,omitempty"`
}

// NewCandyMachineSearchRequest instantiates a new CandyMachineSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCandyMachineSearchRequest() *CandyMachineSearchRequest {
	this := CandyMachineSearchRequest{}
	var updateAuthoritySearchMethod string = "exact_match"
	this.UpdateAuthoritySearchMethod = &updateAuthoritySearchMethod
	var configAddressSearchMethod string = "exact_match"
	this.ConfigAddressSearchMethod = &configAddressSearchMethod
	var uuidSearchMethod string = "exact_match"
	this.UuidSearchMethod = &uuidSearchMethod
	var symbolSearchMethod string = "exact_match"
	this.SymbolSearchMethod = &symbolSearchMethod
	var nftNameIndex string = "0"
	this.NftNameIndex = &nftNameIndex
	var nftNameSearchMethod string = "exact_match"
	this.NftNameSearchMethod = &nftNameSearchMethod
	var network string = "devnet"
	this.Network = &network
	var candyMachineContractVersion string = "v1"
	this.CandyMachineContractVersion = &candyMachineContractVersion
	return &this
}

// NewCandyMachineSearchRequestWithDefaults instantiates a new CandyMachineSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCandyMachineSearchRequestWithDefaults() *CandyMachineSearchRequest {
	this := CandyMachineSearchRequest{}
	var updateAuthoritySearchMethod string = "exact_match"
	this.UpdateAuthoritySearchMethod = &updateAuthoritySearchMethod
	var configAddressSearchMethod string = "exact_match"
	this.ConfigAddressSearchMethod = &configAddressSearchMethod
	var uuidSearchMethod string = "exact_match"
	this.UuidSearchMethod = &uuidSearchMethod
	var symbolSearchMethod string = "exact_match"
	this.SymbolSearchMethod = &symbolSearchMethod
	var nftNameIndex string = "0"
	this.NftNameIndex = &nftNameIndex
	var nftNameSearchMethod string = "exact_match"
	this.NftNameSearchMethod = &nftNameSearchMethod
	var network string = "devnet"
	this.Network = &network
	var candyMachineContractVersion string = "v1"
	this.CandyMachineContractVersion = &candyMachineContractVersion
	return &this
}

// GetUpdateAuthority returns the UpdateAuthority field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetUpdateAuthority() string {
	if o == nil || o.UpdateAuthority == nil {
		var ret string
		return ret
	}
	return *o.UpdateAuthority
}

// GetUpdateAuthorityOk returns a tuple with the UpdateAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetUpdateAuthorityOk() (*string, bool) {
	if o == nil || o.UpdateAuthority == nil {
		return nil, false
	}
	return o.UpdateAuthority, true
}

// HasUpdateAuthority returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasUpdateAuthority() bool {
	if o != nil && o.UpdateAuthority != nil {
		return true
	}

	return false
}

// SetUpdateAuthority gets a reference to the given string and assigns it to the UpdateAuthority field.
func (o *CandyMachineSearchRequest) SetUpdateAuthority(v string) {
	o.UpdateAuthority = &v
}

// GetUpdateAuthoritySearchMethod returns the UpdateAuthoritySearchMethod field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetUpdateAuthoritySearchMethod() string {
	if o == nil || o.UpdateAuthoritySearchMethod == nil {
		var ret string
		return ret
	}
	return *o.UpdateAuthoritySearchMethod
}

// GetUpdateAuthoritySearchMethodOk returns a tuple with the UpdateAuthoritySearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetUpdateAuthoritySearchMethodOk() (*string, bool) {
	if o == nil || o.UpdateAuthoritySearchMethod == nil {
		return nil, false
	}
	return o.UpdateAuthoritySearchMethod, true
}

// HasUpdateAuthoritySearchMethod returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasUpdateAuthoritySearchMethod() bool {
	if o != nil && o.UpdateAuthoritySearchMethod != nil {
		return true
	}

	return false
}

// SetUpdateAuthoritySearchMethod gets a reference to the given string and assigns it to the UpdateAuthoritySearchMethod field.
func (o *CandyMachineSearchRequest) SetUpdateAuthoritySearchMethod(v string) {
	o.UpdateAuthoritySearchMethod = &v
}

// GetConfigAddress returns the ConfigAddress field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetConfigAddress() string {
	if o == nil || o.ConfigAddress == nil {
		var ret string
		return ret
	}
	return *o.ConfigAddress
}

// GetConfigAddressOk returns a tuple with the ConfigAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetConfigAddressOk() (*string, bool) {
	if o == nil || o.ConfigAddress == nil {
		return nil, false
	}
	return o.ConfigAddress, true
}

// HasConfigAddress returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasConfigAddress() bool {
	if o != nil && o.ConfigAddress != nil {
		return true
	}

	return false
}

// SetConfigAddress gets a reference to the given string and assigns it to the ConfigAddress field.
func (o *CandyMachineSearchRequest) SetConfigAddress(v string) {
	o.ConfigAddress = &v
}

// GetConfigAddressSearchMethod returns the ConfigAddressSearchMethod field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetConfigAddressSearchMethod() string {
	if o == nil || o.ConfigAddressSearchMethod == nil {
		var ret string
		return ret
	}
	return *o.ConfigAddressSearchMethod
}

// GetConfigAddressSearchMethodOk returns a tuple with the ConfigAddressSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetConfigAddressSearchMethodOk() (*string, bool) {
	if o == nil || o.ConfigAddressSearchMethod == nil {
		return nil, false
	}
	return o.ConfigAddressSearchMethod, true
}

// HasConfigAddressSearchMethod returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasConfigAddressSearchMethod() bool {
	if o != nil && o.ConfigAddressSearchMethod != nil {
		return true
	}

	return false
}

// SetConfigAddressSearchMethod gets a reference to the given string and assigns it to the ConfigAddressSearchMethod field.
func (o *CandyMachineSearchRequest) SetConfigAddressSearchMethod(v string) {
	o.ConfigAddressSearchMethod = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CandyMachineSearchRequest) SetUuid(v string) {
	o.Uuid = &v
}

// GetUuidSearchMethod returns the UuidSearchMethod field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetUuidSearchMethod() string {
	if o == nil || o.UuidSearchMethod == nil {
		var ret string
		return ret
	}
	return *o.UuidSearchMethod
}

// GetUuidSearchMethodOk returns a tuple with the UuidSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetUuidSearchMethodOk() (*string, bool) {
	if o == nil || o.UuidSearchMethod == nil {
		return nil, false
	}
	return o.UuidSearchMethod, true
}

// HasUuidSearchMethod returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasUuidSearchMethod() bool {
	if o != nil && o.UuidSearchMethod != nil {
		return true
	}

	return false
}

// SetUuidSearchMethod gets a reference to the given string and assigns it to the UuidSearchMethod field.
func (o *CandyMachineSearchRequest) SetUuidSearchMethod(v string) {
	o.UuidSearchMethod = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CandyMachineSearchRequest) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSymbolSearchMethod returns the SymbolSearchMethod field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetSymbolSearchMethod() string {
	if o == nil || o.SymbolSearchMethod == nil {
		var ret string
		return ret
	}
	return *o.SymbolSearchMethod
}

// GetSymbolSearchMethodOk returns a tuple with the SymbolSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetSymbolSearchMethodOk() (*string, bool) {
	if o == nil || o.SymbolSearchMethod == nil {
		return nil, false
	}
	return o.SymbolSearchMethod, true
}

// HasSymbolSearchMethod returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasSymbolSearchMethod() bool {
	if o != nil && o.SymbolSearchMethod != nil {
		return true
	}

	return false
}

// SetSymbolSearchMethod gets a reference to the given string and assigns it to the SymbolSearchMethod field.
func (o *CandyMachineSearchRequest) SetSymbolSearchMethod(v string) {
	o.SymbolSearchMethod = &v
}

// GetNftName returns the NftName field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetNftName() string {
	if o == nil || o.NftName == nil {
		var ret string
		return ret
	}
	return *o.NftName
}

// GetNftNameOk returns a tuple with the NftName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetNftNameOk() (*string, bool) {
	if o == nil || o.NftName == nil {
		return nil, false
	}
	return o.NftName, true
}

// HasNftName returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasNftName() bool {
	if o != nil && o.NftName != nil {
		return true
	}

	return false
}

// SetNftName gets a reference to the given string and assigns it to the NftName field.
func (o *CandyMachineSearchRequest) SetNftName(v string) {
	o.NftName = &v
}

// GetNftNameIndex returns the NftNameIndex field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetNftNameIndex() string {
	if o == nil || o.NftNameIndex == nil {
		var ret string
		return ret
	}
	return *o.NftNameIndex
}

// GetNftNameIndexOk returns a tuple with the NftNameIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetNftNameIndexOk() (*string, bool) {
	if o == nil || o.NftNameIndex == nil {
		return nil, false
	}
	return o.NftNameIndex, true
}

// HasNftNameIndex returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasNftNameIndex() bool {
	if o != nil && o.NftNameIndex != nil {
		return true
	}

	return false
}

// SetNftNameIndex gets a reference to the given string and assigns it to the NftNameIndex field.
func (o *CandyMachineSearchRequest) SetNftNameIndex(v string) {
	o.NftNameIndex = &v
}

// GetNftNameSearchMethod returns the NftNameSearchMethod field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetNftNameSearchMethod() string {
	if o == nil || o.NftNameSearchMethod == nil {
		var ret string
		return ret
	}
	return *o.NftNameSearchMethod
}

// GetNftNameSearchMethodOk returns a tuple with the NftNameSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetNftNameSearchMethodOk() (*string, bool) {
	if o == nil || o.NftNameSearchMethod == nil {
		return nil, false
	}
	return o.NftNameSearchMethod, true
}

// HasNftNameSearchMethod returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasNftNameSearchMethod() bool {
	if o != nil && o.NftNameSearchMethod != nil {
		return true
	}

	return false
}

// SetNftNameSearchMethod gets a reference to the given string and assigns it to the NftNameSearchMethod field.
func (o *CandyMachineSearchRequest) SetNftNameSearchMethod(v string) {
	o.NftNameSearchMethod = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *CandyMachineSearchRequest) SetNetwork(v string) {
	o.Network = &v
}

// GetCandyMachineContractVersion returns the CandyMachineContractVersion field value if set, zero value otherwise.
func (o *CandyMachineSearchRequest) GetCandyMachineContractVersion() string {
	if o == nil || o.CandyMachineContractVersion == nil {
		var ret string
		return ret
	}
	return *o.CandyMachineContractVersion
}

// GetCandyMachineContractVersionOk returns a tuple with the CandyMachineContractVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CandyMachineSearchRequest) GetCandyMachineContractVersionOk() (*string, bool) {
	if o == nil || o.CandyMachineContractVersion == nil {
		return nil, false
	}
	return o.CandyMachineContractVersion, true
}

// HasCandyMachineContractVersion returns a boolean if a field has been set.
func (o *CandyMachineSearchRequest) HasCandyMachineContractVersion() bool {
	if o != nil && o.CandyMachineContractVersion != nil {
		return true
	}

	return false
}

// SetCandyMachineContractVersion gets a reference to the given string and assigns it to the CandyMachineContractVersion field.
func (o *CandyMachineSearchRequest) SetCandyMachineContractVersion(v string) {
	o.CandyMachineContractVersion = &v
}

func (o CandyMachineSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UpdateAuthority != nil {
		toSerialize["update_authority"] = o.UpdateAuthority
	}
	if o.UpdateAuthoritySearchMethod != nil {
		toSerialize["update_authority_search_method"] = o.UpdateAuthoritySearchMethod
	}
	if o.ConfigAddress != nil {
		toSerialize["config_address"] = o.ConfigAddress
	}
	if o.ConfigAddressSearchMethod != nil {
		toSerialize["config_address_search_method"] = o.ConfigAddressSearchMethod
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.UuidSearchMethod != nil {
		toSerialize["uuid_search_method"] = o.UuidSearchMethod
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.SymbolSearchMethod != nil {
		toSerialize["symbol_search_method"] = o.SymbolSearchMethod
	}
	if o.NftName != nil {
		toSerialize["nft_name"] = o.NftName
	}
	if o.NftNameIndex != nil {
		toSerialize["nft_name_index"] = o.NftNameIndex
	}
	if o.NftNameSearchMethod != nil {
		toSerialize["nft_name_search_method"] = o.NftNameSearchMethod
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.CandyMachineContractVersion != nil {
		toSerialize["candy_machine_contract_version"] = o.CandyMachineContractVersion
	}
	return json.Marshal(toSerialize)
}

type NullableCandyMachineSearchRequest struct {
	value *CandyMachineSearchRequest
	isSet bool
}

func (v NullableCandyMachineSearchRequest) Get() *CandyMachineSearchRequest {
	return v.value
}

func (v *NullableCandyMachineSearchRequest) Set(val *CandyMachineSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCandyMachineSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCandyMachineSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCandyMachineSearchRequest(val *CandyMachineSearchRequest) *NullableCandyMachineSearchRequest {
	return &NullableCandyMachineSearchRequest{value: val, isSet: true}
}

func (v NullableCandyMachineSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCandyMachineSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


