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

// EthereumTransactionCompiledResponse struct for EthereumTransactionCompiledResponse
type EthereumTransactionCompiledResponse struct {
	// Indicates the chain that the transaction was compiled for (e.g., ropsten or mainnet). 
	ChainId *string `json:"chain_id,omitempty"`
	// The transaction data 
	Data *string `json:"data,omitempty"`
	// The address expected to sign and submit the transaction 
	From *string `json:"from,omitempty"`
	// The contract. This should match your provided value for `token_blockchain_identifier`. 
	To *string `json:"to,omitempty"`
	Gas *float32 `json:"gas,omitempty"`
	MaxFeePerGas *float32 `json:"max_fee_per_gas,omitempty"`
	MaxPriorityFeePerGas *float32 `json:"max_priority_fee_per_gas,omitempty"`
	Nonce *float32 `json:"nonce,omitempty"`
	Value *float32 `json:"value,omitempty"`
}

// NewEthereumTransactionCompiledResponse instantiates a new EthereumTransactionCompiledResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthereumTransactionCompiledResponse() *EthereumTransactionCompiledResponse {
	this := EthereumTransactionCompiledResponse{}
	return &this
}

// NewEthereumTransactionCompiledResponseWithDefaults instantiates a new EthereumTransactionCompiledResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthereumTransactionCompiledResponseWithDefaults() *EthereumTransactionCompiledResponse {
	this := EthereumTransactionCompiledResponse{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *EthereumTransactionCompiledResponse) GetChainId() string {
	if o == nil || o.ChainId == nil {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionCompiledResponse) GetChainIdOk() (*string, bool) {
	if o == nil || o.ChainId == nil {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *EthereumTransactionCompiledResponse) HasChainId() bool {
	if o != nil && o.ChainId != nil {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *EthereumTransactionCompiledResponse) SetChainId(v string) {
	o.ChainId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EthereumTransactionCompiledResponse) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionCompiledResponse) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EthereumTransactionCompiledResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *EthereumTransactionCompiledResponse) SetData(v string) {
	o.Data = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EthereumTransactionCompiledResponse) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionCompiledResponse) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EthereumTransactionCompiledResponse) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *EthereumTransactionCompiledResponse) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *EthereumTransactionCompiledResponse) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionCompiledResponse) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *EthereumTransactionCompiledResponse) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *EthereumTransactionCompiledResponse) SetTo(v string) {
	o.To = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *EthereumTransactionCompiledResponse) GetGas() float32 {
	if o == nil || o.Gas == nil {
		var ret float32
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionCompiledResponse) GetGasOk() (*float32, bool) {
	if o == nil || o.Gas == nil {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *EthereumTransactionCompiledResponse) HasGas() bool {
	if o != nil && o.Gas != nil {
		return true
	}

	return false
}

// SetGas gets a reference to the given float32 and assigns it to the Gas field.
func (o *EthereumTransactionCompiledResponse) SetGas(v float32) {
	o.Gas = &v
}

// GetMaxFeePerGas returns the MaxFeePerGas field value if set, zero value otherwise.
func (o *EthereumTransactionCompiledResponse) GetMaxFeePerGas() float32 {
	if o == nil || o.MaxFeePerGas == nil {
		var ret float32
		return ret
	}
	return *o.MaxFeePerGas
}

// GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionCompiledResponse) GetMaxFeePerGasOk() (*float32, bool) {
	if o == nil || o.MaxFeePerGas == nil {
		return nil, false
	}
	return o.MaxFeePerGas, true
}

// HasMaxFeePerGas returns a boolean if a field has been set.
func (o *EthereumTransactionCompiledResponse) HasMaxFeePerGas() bool {
	if o != nil && o.MaxFeePerGas != nil {
		return true
	}

	return false
}

// SetMaxFeePerGas gets a reference to the given float32 and assigns it to the MaxFeePerGas field.
func (o *EthereumTransactionCompiledResponse) SetMaxFeePerGas(v float32) {
	o.MaxFeePerGas = &v
}

// GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field value if set, zero value otherwise.
func (o *EthereumTransactionCompiledResponse) GetMaxPriorityFeePerGas() float32 {
	if o == nil || o.MaxPriorityFeePerGas == nil {
		var ret float32
		return ret
	}
	return *o.MaxPriorityFeePerGas
}

// GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionCompiledResponse) GetMaxPriorityFeePerGasOk() (*float32, bool) {
	if o == nil || o.MaxPriorityFeePerGas == nil {
		return nil, false
	}
	return o.MaxPriorityFeePerGas, true
}

// HasMaxPriorityFeePerGas returns a boolean if a field has been set.
func (o *EthereumTransactionCompiledResponse) HasMaxPriorityFeePerGas() bool {
	if o != nil && o.MaxPriorityFeePerGas != nil {
		return true
	}

	return false
}

// SetMaxPriorityFeePerGas gets a reference to the given float32 and assigns it to the MaxPriorityFeePerGas field.
func (o *EthereumTransactionCompiledResponse) SetMaxPriorityFeePerGas(v float32) {
	o.MaxPriorityFeePerGas = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *EthereumTransactionCompiledResponse) GetNonce() float32 {
	if o == nil || o.Nonce == nil {
		var ret float32
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionCompiledResponse) GetNonceOk() (*float32, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *EthereumTransactionCompiledResponse) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given float32 and assigns it to the Nonce field.
func (o *EthereumTransactionCompiledResponse) SetNonce(v float32) {
	o.Nonce = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EthereumTransactionCompiledResponse) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransactionCompiledResponse) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EthereumTransactionCompiledResponse) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *EthereumTransactionCompiledResponse) SetValue(v float32) {
	o.Value = &v
}

func (o EthereumTransactionCompiledResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChainId != nil {
		toSerialize["chain_id"] = o.ChainId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.Gas != nil {
		toSerialize["gas"] = o.Gas
	}
	if o.MaxFeePerGas != nil {
		toSerialize["max_fee_per_gas"] = o.MaxFeePerGas
	}
	if o.MaxPriorityFeePerGas != nil {
		toSerialize["max_priority_fee_per_gas"] = o.MaxPriorityFeePerGas
	}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableEthereumTransactionCompiledResponse struct {
	value *EthereumTransactionCompiledResponse
	isSet bool
}

func (v NullableEthereumTransactionCompiledResponse) Get() *EthereumTransactionCompiledResponse {
	return v.value
}

func (v *NullableEthereumTransactionCompiledResponse) Set(val *EthereumTransactionCompiledResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEthereumTransactionCompiledResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEthereumTransactionCompiledResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthereumTransactionCompiledResponse(val *EthereumTransactionCompiledResponse) *NullableEthereumTransactionCompiledResponse {
	return &NullableEthereumTransactionCompiledResponse{value: val, isSet: true}
}

func (v NullableEthereumTransactionCompiledResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthereumTransactionCompiledResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


