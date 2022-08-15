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

// NFTSearchRequest struct for NFTSearchRequest
type NFTSearchRequest struct {
	// The public key of the update authority of the NFT
	UpdateAuthority *string `json:"update_authority,omitempty"`
	// Only `exact_match` supported at this time
	UpdateAuthoritySearchMethod *string `json:"update_authority_search_method,omitempty"`
	// The mint address of the NFT
	MintAddress *string `json:"mint_address,omitempty"`
	// Only `exact_match` supported at this time
	MintAddressSearchMethod *string `json:"mint_address_search_method,omitempty"`
	// The name of the NFT
	Name *string `json:"name,omitempty"`
	NameSearchMethod *string `json:"name_search_method,omitempty"`
	// The NFT's uri
	Uri *string `json:"uri,omitempty"`
	UriSearchMethod *string `json:"uri_search_method,omitempty"`
	// The symbol associated with the candy machine
	Symbol *string `json:"symbol,omitempty"`
	SymbolSearchMethod *string `json:"symbol_search_method,omitempty"`
	Network *string `json:"network,omitempty"`
}

// NewNFTSearchRequest instantiates a new NFTSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFTSearchRequest() *NFTSearchRequest {
	this := NFTSearchRequest{}
	var updateAuthoritySearchMethod string = "exact_match"
	this.UpdateAuthoritySearchMethod = &updateAuthoritySearchMethod
	var mintAddressSearchMethod string = "exact_match"
	this.MintAddressSearchMethod = &mintAddressSearchMethod
	var nameSearchMethod string = "exact_match"
	this.NameSearchMethod = &nameSearchMethod
	var uriSearchMethod string = "exact_match"
	this.UriSearchMethod = &uriSearchMethod
	var symbolSearchMethod string = "exact_match"
	this.SymbolSearchMethod = &symbolSearchMethod
	var network string = "devnet"
	this.Network = &network
	return &this
}

// NewNFTSearchRequestWithDefaults instantiates a new NFTSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTSearchRequestWithDefaults() *NFTSearchRequest {
	this := NFTSearchRequest{}
	var updateAuthoritySearchMethod string = "exact_match"
	this.UpdateAuthoritySearchMethod = &updateAuthoritySearchMethod
	var mintAddressSearchMethod string = "exact_match"
	this.MintAddressSearchMethod = &mintAddressSearchMethod
	var nameSearchMethod string = "exact_match"
	this.NameSearchMethod = &nameSearchMethod
	var uriSearchMethod string = "exact_match"
	this.UriSearchMethod = &uriSearchMethod
	var symbolSearchMethod string = "exact_match"
	this.SymbolSearchMethod = &symbolSearchMethod
	var network string = "devnet"
	this.Network = &network
	return &this
}

// GetUpdateAuthority returns the UpdateAuthority field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetUpdateAuthority() string {
	if o == nil || o.UpdateAuthority == nil {
		var ret string
		return ret
	}
	return *o.UpdateAuthority
}

// GetUpdateAuthorityOk returns a tuple with the UpdateAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetUpdateAuthorityOk() (*string, bool) {
	if o == nil || o.UpdateAuthority == nil {
		return nil, false
	}
	return o.UpdateAuthority, true
}

// HasUpdateAuthority returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasUpdateAuthority() bool {
	if o != nil && o.UpdateAuthority != nil {
		return true
	}

	return false
}

// SetUpdateAuthority gets a reference to the given string and assigns it to the UpdateAuthority field.
func (o *NFTSearchRequest) SetUpdateAuthority(v string) {
	o.UpdateAuthority = &v
}

// GetUpdateAuthoritySearchMethod returns the UpdateAuthoritySearchMethod field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetUpdateAuthoritySearchMethod() string {
	if o == nil || o.UpdateAuthoritySearchMethod == nil {
		var ret string
		return ret
	}
	return *o.UpdateAuthoritySearchMethod
}

// GetUpdateAuthoritySearchMethodOk returns a tuple with the UpdateAuthoritySearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetUpdateAuthoritySearchMethodOk() (*string, bool) {
	if o == nil || o.UpdateAuthoritySearchMethod == nil {
		return nil, false
	}
	return o.UpdateAuthoritySearchMethod, true
}

// HasUpdateAuthoritySearchMethod returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasUpdateAuthoritySearchMethod() bool {
	if o != nil && o.UpdateAuthoritySearchMethod != nil {
		return true
	}

	return false
}

// SetUpdateAuthoritySearchMethod gets a reference to the given string and assigns it to the UpdateAuthoritySearchMethod field.
func (o *NFTSearchRequest) SetUpdateAuthoritySearchMethod(v string) {
	o.UpdateAuthoritySearchMethod = &v
}

// GetMintAddress returns the MintAddress field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetMintAddress() string {
	if o == nil || o.MintAddress == nil {
		var ret string
		return ret
	}
	return *o.MintAddress
}

// GetMintAddressOk returns a tuple with the MintAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetMintAddressOk() (*string, bool) {
	if o == nil || o.MintAddress == nil {
		return nil, false
	}
	return o.MintAddress, true
}

// HasMintAddress returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasMintAddress() bool {
	if o != nil && o.MintAddress != nil {
		return true
	}

	return false
}

// SetMintAddress gets a reference to the given string and assigns it to the MintAddress field.
func (o *NFTSearchRequest) SetMintAddress(v string) {
	o.MintAddress = &v
}

// GetMintAddressSearchMethod returns the MintAddressSearchMethod field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetMintAddressSearchMethod() string {
	if o == nil || o.MintAddressSearchMethod == nil {
		var ret string
		return ret
	}
	return *o.MintAddressSearchMethod
}

// GetMintAddressSearchMethodOk returns a tuple with the MintAddressSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetMintAddressSearchMethodOk() (*string, bool) {
	if o == nil || o.MintAddressSearchMethod == nil {
		return nil, false
	}
	return o.MintAddressSearchMethod, true
}

// HasMintAddressSearchMethod returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasMintAddressSearchMethod() bool {
	if o != nil && o.MintAddressSearchMethod != nil {
		return true
	}

	return false
}

// SetMintAddressSearchMethod gets a reference to the given string and assigns it to the MintAddressSearchMethod field.
func (o *NFTSearchRequest) SetMintAddressSearchMethod(v string) {
	o.MintAddressSearchMethod = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NFTSearchRequest) SetName(v string) {
	o.Name = &v
}

// GetNameSearchMethod returns the NameSearchMethod field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetNameSearchMethod() string {
	if o == nil || o.NameSearchMethod == nil {
		var ret string
		return ret
	}
	return *o.NameSearchMethod
}

// GetNameSearchMethodOk returns a tuple with the NameSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetNameSearchMethodOk() (*string, bool) {
	if o == nil || o.NameSearchMethod == nil {
		return nil, false
	}
	return o.NameSearchMethod, true
}

// HasNameSearchMethod returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasNameSearchMethod() bool {
	if o != nil && o.NameSearchMethod != nil {
		return true
	}

	return false
}

// SetNameSearchMethod gets a reference to the given string and assigns it to the NameSearchMethod field.
func (o *NFTSearchRequest) SetNameSearchMethod(v string) {
	o.NameSearchMethod = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *NFTSearchRequest) SetUri(v string) {
	o.Uri = &v
}

// GetUriSearchMethod returns the UriSearchMethod field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetUriSearchMethod() string {
	if o == nil || o.UriSearchMethod == nil {
		var ret string
		return ret
	}
	return *o.UriSearchMethod
}

// GetUriSearchMethodOk returns a tuple with the UriSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetUriSearchMethodOk() (*string, bool) {
	if o == nil || o.UriSearchMethod == nil {
		return nil, false
	}
	return o.UriSearchMethod, true
}

// HasUriSearchMethod returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasUriSearchMethod() bool {
	if o != nil && o.UriSearchMethod != nil {
		return true
	}

	return false
}

// SetUriSearchMethod gets a reference to the given string and assigns it to the UriSearchMethod field.
func (o *NFTSearchRequest) SetUriSearchMethod(v string) {
	o.UriSearchMethod = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *NFTSearchRequest) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSymbolSearchMethod returns the SymbolSearchMethod field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetSymbolSearchMethod() string {
	if o == nil || o.SymbolSearchMethod == nil {
		var ret string
		return ret
	}
	return *o.SymbolSearchMethod
}

// GetSymbolSearchMethodOk returns a tuple with the SymbolSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetSymbolSearchMethodOk() (*string, bool) {
	if o == nil || o.SymbolSearchMethod == nil {
		return nil, false
	}
	return o.SymbolSearchMethod, true
}

// HasSymbolSearchMethod returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasSymbolSearchMethod() bool {
	if o != nil && o.SymbolSearchMethod != nil {
		return true
	}

	return false
}

// SetSymbolSearchMethod gets a reference to the given string and assigns it to the SymbolSearchMethod field.
func (o *NFTSearchRequest) SetSymbolSearchMethod(v string) {
	o.SymbolSearchMethod = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NFTSearchRequest) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTSearchRequest) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NFTSearchRequest) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *NFTSearchRequest) SetNetwork(v string) {
	o.Network = &v
}

func (o NFTSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UpdateAuthority != nil {
		toSerialize["update_authority"] = o.UpdateAuthority
	}
	if o.UpdateAuthoritySearchMethod != nil {
		toSerialize["update_authority_search_method"] = o.UpdateAuthoritySearchMethod
	}
	if o.MintAddress != nil {
		toSerialize["mint_address"] = o.MintAddress
	}
	if o.MintAddressSearchMethod != nil {
		toSerialize["mint_address_search_method"] = o.MintAddressSearchMethod
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NameSearchMethod != nil {
		toSerialize["name_search_method"] = o.NameSearchMethod
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.UriSearchMethod != nil {
		toSerialize["uri_search_method"] = o.UriSearchMethod
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.SymbolSearchMethod != nil {
		toSerialize["symbol_search_method"] = o.SymbolSearchMethod
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableNFTSearchRequest struct {
	value *NFTSearchRequest
	isSet bool
}

func (v NullableNFTSearchRequest) Get() *NFTSearchRequest {
	return v.value
}

func (v *NullableNFTSearchRequest) Set(val *NFTSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNFTSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNFTSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFTSearchRequest(val *NFTSearchRequest) *NullableNFTSearchRequest {
	return &NullableNFTSearchRequest{value: val, isSet: true}
}

func (v NullableNFTSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFTSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


