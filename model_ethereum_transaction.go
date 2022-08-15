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

// EthereumTransaction struct for EthereumTransaction
type EthereumTransaction struct {
	AccessList map[string]interface{} `json:"access_list,omitempty"`
	BlockHash *string `json:"block_hash,omitempty"`
	BlockNumber *string `json:"block_number,omitempty"`
	ChainId *string `json:"chain_id,omitempty"`
	From *string `json:"from,omitempty"`
	Gas *string `json:"gas,omitempty"`
	GasPrice *string `json:"gas_price,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Input *string `json:"input,omitempty"`
	MaxFeePerGas *string `json:"max_fee_per_gas,omitempty"`
	MaxPriorityFeePerGas *string `json:"max_priority_fee_per_gas,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	R *string `json:"r,omitempty"`
	S *string `json:"s,omitempty"`
	To *string `json:"to,omitempty"`
	TransactionIndex *string `json:"transaction_index,omitempty"`
	Type *string `json:"type,omitempty"`
	V *string `json:"v,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewEthereumTransaction instantiates a new EthereumTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthereumTransaction() *EthereumTransaction {
	this := EthereumTransaction{}
	return &this
}

// NewEthereumTransactionWithDefaults instantiates a new EthereumTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthereumTransactionWithDefaults() *EthereumTransaction {
	this := EthereumTransaction{}
	return &this
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *EthereumTransaction) GetAccessList() map[string]interface{} {
	if o == nil || o.AccessList == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetAccessListOk() (map[string]interface{}, bool) {
	if o == nil || o.AccessList == nil {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *EthereumTransaction) HasAccessList() bool {
	if o != nil && o.AccessList != nil {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given map[string]interface{} and assigns it to the AccessList field.
func (o *EthereumTransaction) SetAccessList(v map[string]interface{}) {
	o.AccessList = v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *EthereumTransaction) GetBlockHash() string {
	if o == nil || o.BlockHash == nil {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetBlockHashOk() (*string, bool) {
	if o == nil || o.BlockHash == nil {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *EthereumTransaction) HasBlockHash() bool {
	if o != nil && o.BlockHash != nil {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *EthereumTransaction) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise.
func (o *EthereumTransaction) GetBlockNumber() string {
	if o == nil || o.BlockNumber == nil {
		var ret string
		return ret
	}
	return *o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetBlockNumberOk() (*string, bool) {
	if o == nil || o.BlockNumber == nil {
		return nil, false
	}
	return o.BlockNumber, true
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *EthereumTransaction) HasBlockNumber() bool {
	if o != nil && o.BlockNumber != nil {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given string and assigns it to the BlockNumber field.
func (o *EthereumTransaction) SetBlockNumber(v string) {
	o.BlockNumber = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *EthereumTransaction) GetChainId() string {
	if o == nil || o.ChainId == nil {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetChainIdOk() (*string, bool) {
	if o == nil || o.ChainId == nil {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *EthereumTransaction) HasChainId() bool {
	if o != nil && o.ChainId != nil {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *EthereumTransaction) SetChainId(v string) {
	o.ChainId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EthereumTransaction) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EthereumTransaction) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *EthereumTransaction) SetFrom(v string) {
	o.From = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *EthereumTransaction) GetGas() string {
	if o == nil || o.Gas == nil {
		var ret string
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetGasOk() (*string, bool) {
	if o == nil || o.Gas == nil {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *EthereumTransaction) HasGas() bool {
	if o != nil && o.Gas != nil {
		return true
	}

	return false
}

// SetGas gets a reference to the given string and assigns it to the Gas field.
func (o *EthereumTransaction) SetGas(v string) {
	o.Gas = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *EthereumTransaction) GetGasPrice() string {
	if o == nil || o.GasPrice == nil {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetGasPriceOk() (*string, bool) {
	if o == nil || o.GasPrice == nil {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *EthereumTransaction) HasGasPrice() bool {
	if o != nil && o.GasPrice != nil {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *EthereumTransaction) SetGasPrice(v string) {
	o.GasPrice = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *EthereumTransaction) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *EthereumTransaction) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *EthereumTransaction) SetHash(v string) {
	o.Hash = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *EthereumTransaction) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *EthereumTransaction) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *EthereumTransaction) SetInput(v string) {
	o.Input = &v
}

// GetMaxFeePerGas returns the MaxFeePerGas field value if set, zero value otherwise.
func (o *EthereumTransaction) GetMaxFeePerGas() string {
	if o == nil || o.MaxFeePerGas == nil {
		var ret string
		return ret
	}
	return *o.MaxFeePerGas
}

// GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetMaxFeePerGasOk() (*string, bool) {
	if o == nil || o.MaxFeePerGas == nil {
		return nil, false
	}
	return o.MaxFeePerGas, true
}

// HasMaxFeePerGas returns a boolean if a field has been set.
func (o *EthereumTransaction) HasMaxFeePerGas() bool {
	if o != nil && o.MaxFeePerGas != nil {
		return true
	}

	return false
}

// SetMaxFeePerGas gets a reference to the given string and assigns it to the MaxFeePerGas field.
func (o *EthereumTransaction) SetMaxFeePerGas(v string) {
	o.MaxFeePerGas = &v
}

// GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field value if set, zero value otherwise.
func (o *EthereumTransaction) GetMaxPriorityFeePerGas() string {
	if o == nil || o.MaxPriorityFeePerGas == nil {
		var ret string
		return ret
	}
	return *o.MaxPriorityFeePerGas
}

// GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetMaxPriorityFeePerGasOk() (*string, bool) {
	if o == nil || o.MaxPriorityFeePerGas == nil {
		return nil, false
	}
	return o.MaxPriorityFeePerGas, true
}

// HasMaxPriorityFeePerGas returns a boolean if a field has been set.
func (o *EthereumTransaction) HasMaxPriorityFeePerGas() bool {
	if o != nil && o.MaxPriorityFeePerGas != nil {
		return true
	}

	return false
}

// SetMaxPriorityFeePerGas gets a reference to the given string and assigns it to the MaxPriorityFeePerGas field.
func (o *EthereumTransaction) SetMaxPriorityFeePerGas(v string) {
	o.MaxPriorityFeePerGas = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *EthereumTransaction) GetNonce() string {
	if o == nil || o.Nonce == nil {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetNonceOk() (*string, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *EthereumTransaction) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *EthereumTransaction) SetNonce(v string) {
	o.Nonce = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *EthereumTransaction) GetR() string {
	if o == nil || o.R == nil {
		var ret string
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetROk() (*string, bool) {
	if o == nil || o.R == nil {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *EthereumTransaction) HasR() bool {
	if o != nil && o.R != nil {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *EthereumTransaction) SetR(v string) {
	o.R = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *EthereumTransaction) GetS() string {
	if o == nil || o.S == nil {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetSOk() (*string, bool) {
	if o == nil || o.S == nil {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *EthereumTransaction) HasS() bool {
	if o != nil && o.S != nil {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *EthereumTransaction) SetS(v string) {
	o.S = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *EthereumTransaction) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *EthereumTransaction) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *EthereumTransaction) SetTo(v string) {
	o.To = &v
}

// GetTransactionIndex returns the TransactionIndex field value if set, zero value otherwise.
func (o *EthereumTransaction) GetTransactionIndex() string {
	if o == nil || o.TransactionIndex == nil {
		var ret string
		return ret
	}
	return *o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetTransactionIndexOk() (*string, bool) {
	if o == nil || o.TransactionIndex == nil {
		return nil, false
	}
	return o.TransactionIndex, true
}

// HasTransactionIndex returns a boolean if a field has been set.
func (o *EthereumTransaction) HasTransactionIndex() bool {
	if o != nil && o.TransactionIndex != nil {
		return true
	}

	return false
}

// SetTransactionIndex gets a reference to the given string and assigns it to the TransactionIndex field.
func (o *EthereumTransaction) SetTransactionIndex(v string) {
	o.TransactionIndex = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EthereumTransaction) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EthereumTransaction) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EthereumTransaction) SetType(v string) {
	o.Type = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *EthereumTransaction) GetV() string {
	if o == nil || o.V == nil {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetVOk() (*string, bool) {
	if o == nil || o.V == nil {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *EthereumTransaction) HasV() bool {
	if o != nil && o.V != nil {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *EthereumTransaction) SetV(v string) {
	o.V = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EthereumTransaction) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthereumTransaction) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EthereumTransaction) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EthereumTransaction) SetValue(v string) {
	o.Value = &v
}

func (o EthereumTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessList != nil {
		toSerialize["access_list"] = o.AccessList
	}
	if o.BlockHash != nil {
		toSerialize["block_hash"] = o.BlockHash
	}
	if o.BlockNumber != nil {
		toSerialize["block_number"] = o.BlockNumber
	}
	if o.ChainId != nil {
		toSerialize["chain_id"] = o.ChainId
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Gas != nil {
		toSerialize["gas"] = o.Gas
	}
	if o.GasPrice != nil {
		toSerialize["gas_price"] = o.GasPrice
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
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
	if o.R != nil {
		toSerialize["r"] = o.R
	}
	if o.S != nil {
		toSerialize["s"] = o.S
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.TransactionIndex != nil {
		toSerialize["transaction_index"] = o.TransactionIndex
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.V != nil {
		toSerialize["v"] = o.V
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableEthereumTransaction struct {
	value *EthereumTransaction
	isSet bool
}

func (v NullableEthereumTransaction) Get() *EthereumTransaction {
	return v.value
}

func (v *NullableEthereumTransaction) Set(val *EthereumTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableEthereumTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableEthereumTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthereumTransaction(val *EthereumTransaction) *NullableEthereumTransaction {
	return &NullableEthereumTransaction{value: val, isSet: true}
}

func (v NullableEthereumTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthereumTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


