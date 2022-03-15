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

// NFTTransaction struct for NFTTransaction
type NFTTransaction struct {
	// An epoch time stamp in UTC time that represents the time of the block where the transaction was processed 
	BlockTime *float32 `json:"block_time,omitempty"`
	// The NFT exchange on which the transaction occurred
	Exchange *string `json:"exchange,omitempty"`
	// A readable version of the NFT exchange
	ExchangeReadable *string `json:"exchange_readable,omitempty"`
	// The mint address of the NFT 
	MintAddress *string `json:"mint_address,omitempty"`
	// The operation of the transaction
	Operation *string `json:"operation,omitempty"`
	// The public key of the wallet that listed the NFT
	Seller *string `json:"seller,omitempty"`
	// The public key of the buyer. This only exists in `buy` transactions. 
	Buyer *string `json:"buyer,omitempty"`
	// The signature of the transaction. You can look up each transaction on explorer.solana.com 
	TransactionSignature *string `json:"transaction_signature,omitempty"`
}

// NewNFTTransaction instantiates a new NFTTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFTTransaction() *NFTTransaction {
	this := NFTTransaction{}
	return &this
}

// NewNFTTransactionWithDefaults instantiates a new NFTTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTTransactionWithDefaults() *NFTTransaction {
	this := NFTTransaction{}
	return &this
}

// GetBlockTime returns the BlockTime field value if set, zero value otherwise.
func (o *NFTTransaction) GetBlockTime() float32 {
	if o == nil || o.BlockTime == nil {
		var ret float32
		return ret
	}
	return *o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTTransaction) GetBlockTimeOk() (*float32, bool) {
	if o == nil || o.BlockTime == nil {
		return nil, false
	}
	return o.BlockTime, true
}

// HasBlockTime returns a boolean if a field has been set.
func (o *NFTTransaction) HasBlockTime() bool {
	if o != nil && o.BlockTime != nil {
		return true
	}

	return false
}

// SetBlockTime gets a reference to the given float32 and assigns it to the BlockTime field.
func (o *NFTTransaction) SetBlockTime(v float32) {
	o.BlockTime = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *NFTTransaction) GetExchange() string {
	if o == nil || o.Exchange == nil {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTTransaction) GetExchangeOk() (*string, bool) {
	if o == nil || o.Exchange == nil {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *NFTTransaction) HasExchange() bool {
	if o != nil && o.Exchange != nil {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *NFTTransaction) SetExchange(v string) {
	o.Exchange = &v
}

// GetExchangeReadable returns the ExchangeReadable field value if set, zero value otherwise.
func (o *NFTTransaction) GetExchangeReadable() string {
	if o == nil || o.ExchangeReadable == nil {
		var ret string
		return ret
	}
	return *o.ExchangeReadable
}

// GetExchangeReadableOk returns a tuple with the ExchangeReadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTTransaction) GetExchangeReadableOk() (*string, bool) {
	if o == nil || o.ExchangeReadable == nil {
		return nil, false
	}
	return o.ExchangeReadable, true
}

// HasExchangeReadable returns a boolean if a field has been set.
func (o *NFTTransaction) HasExchangeReadable() bool {
	if o != nil && o.ExchangeReadable != nil {
		return true
	}

	return false
}

// SetExchangeReadable gets a reference to the given string and assigns it to the ExchangeReadable field.
func (o *NFTTransaction) SetExchangeReadable(v string) {
	o.ExchangeReadable = &v
}

// GetMintAddress returns the MintAddress field value if set, zero value otherwise.
func (o *NFTTransaction) GetMintAddress() string {
	if o == nil || o.MintAddress == nil {
		var ret string
		return ret
	}
	return *o.MintAddress
}

// GetMintAddressOk returns a tuple with the MintAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTTransaction) GetMintAddressOk() (*string, bool) {
	if o == nil || o.MintAddress == nil {
		return nil, false
	}
	return o.MintAddress, true
}

// HasMintAddress returns a boolean if a field has been set.
func (o *NFTTransaction) HasMintAddress() bool {
	if o != nil && o.MintAddress != nil {
		return true
	}

	return false
}

// SetMintAddress gets a reference to the given string and assigns it to the MintAddress field.
func (o *NFTTransaction) SetMintAddress(v string) {
	o.MintAddress = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *NFTTransaction) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTTransaction) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *NFTTransaction) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *NFTTransaction) SetOperation(v string) {
	o.Operation = &v
}

// GetSeller returns the Seller field value if set, zero value otherwise.
func (o *NFTTransaction) GetSeller() string {
	if o == nil || o.Seller == nil {
		var ret string
		return ret
	}
	return *o.Seller
}

// GetSellerOk returns a tuple with the Seller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTTransaction) GetSellerOk() (*string, bool) {
	if o == nil || o.Seller == nil {
		return nil, false
	}
	return o.Seller, true
}

// HasSeller returns a boolean if a field has been set.
func (o *NFTTransaction) HasSeller() bool {
	if o != nil && o.Seller != nil {
		return true
	}

	return false
}

// SetSeller gets a reference to the given string and assigns it to the Seller field.
func (o *NFTTransaction) SetSeller(v string) {
	o.Seller = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *NFTTransaction) GetBuyer() string {
	if o == nil || o.Buyer == nil {
		var ret string
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTTransaction) GetBuyerOk() (*string, bool) {
	if o == nil || o.Buyer == nil {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *NFTTransaction) HasBuyer() bool {
	if o != nil && o.Buyer != nil {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given string and assigns it to the Buyer field.
func (o *NFTTransaction) SetBuyer(v string) {
	o.Buyer = &v
}

// GetTransactionSignature returns the TransactionSignature field value if set, zero value otherwise.
func (o *NFTTransaction) GetTransactionSignature() string {
	if o == nil || o.TransactionSignature == nil {
		var ret string
		return ret
	}
	return *o.TransactionSignature
}

// GetTransactionSignatureOk returns a tuple with the TransactionSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFTTransaction) GetTransactionSignatureOk() (*string, bool) {
	if o == nil || o.TransactionSignature == nil {
		return nil, false
	}
	return o.TransactionSignature, true
}

// HasTransactionSignature returns a boolean if a field has been set.
func (o *NFTTransaction) HasTransactionSignature() bool {
	if o != nil && o.TransactionSignature != nil {
		return true
	}

	return false
}

// SetTransactionSignature gets a reference to the given string and assigns it to the TransactionSignature field.
func (o *NFTTransaction) SetTransactionSignature(v string) {
	o.TransactionSignature = &v
}

func (o NFTTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockTime != nil {
		toSerialize["block_time"] = o.BlockTime
	}
	if o.Exchange != nil {
		toSerialize["exchange"] = o.Exchange
	}
	if o.ExchangeReadable != nil {
		toSerialize["exchange_readable"] = o.ExchangeReadable
	}
	if o.MintAddress != nil {
		toSerialize["mint_address"] = o.MintAddress
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Seller != nil {
		toSerialize["seller"] = o.Seller
	}
	if o.Buyer != nil {
		toSerialize["buyer"] = o.Buyer
	}
	if o.TransactionSignature != nil {
		toSerialize["transaction_signature"] = o.TransactionSignature
	}
	return json.Marshal(toSerialize)
}

type NullableNFTTransaction struct {
	value *NFTTransaction
	isSet bool
}

func (v NullableNFTTransaction) Get() *NFTTransaction {
	return v.value
}

func (v *NullableNFTTransaction) Set(val *NFTTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableNFTTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableNFTTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFTTransaction(val *NFTTransaction) *NullableNFTTransaction {
	return &NullableNFTTransaction{value: val, isSet: true}
}

func (v NullableNFTTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFTTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


