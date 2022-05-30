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

// GeneralTransferRequest struct for GeneralTransferRequest
type GeneralTransferRequest struct {
	// The blockchain identifier of the recipient to whom you want to send a token or NFT.  On Ethereum, this is the hex public address of the recipient (e.g., `0x150865ca659204a9a0cd0da00296c6b5db441172`)  On Solana, this is the public key of the recipient (e.g., `EW3nXn7X4NTWFKWaJgxKrFNoTSkop1cBUVHA21zrfF6u`). 
	RecipientBlockchainIdentifier string `json:"recipient_blockchain_identifier"`
	Wallet *GeneralWallet `json:"wallet,omitempty"`
	// The `token_blockchain_identifier` identifies the token you wish to transfer.  - If you're transferring a native blockchain currency (e.g., SOL, ETH, BNB), then simply do not supply this value. - If you're transfering an NFT, then supply the token address of the NFT. On Solana, this is the `mint_address` or `mint` (the address of the mint). - If you're transfering a token, supply the token address. For Solana, you can find on this on the Solana Explorer (e.g., see `SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt` for <a href=\"https://explorer.solana.com/address/SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt\" target=\"_blank\">Serum Token</a>) for the `token_address`.  Examples: - Ethereum: `0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48` - Solana: `CK1LHEANTu7RFqN3XMzo2AnZhyus2W1vue1njrxLEM1d`
	TokenBlockchainIdentifier *string `json:"token_blockchain_identifier,omitempty"`
	// The network of the blockchain you selected  - Solana: `devnet`, `mainnet-beta` - Ethereum: `ropsten`, `mainnet`  Defaults when not provided (not applicable to path parameters): - Solana: `devnet` - Ethereum: `ropsten`
	Network *string `json:"network,omitempty"`
	// This value must be a string. What you provide here depends on if you are sending an NFT, an SPL token, or SOL.  - Native Currency (e.g., SOL, ETH, BNB): Supply this value denominated in the native currency (e.g., in SOL (but not in Lamports), or in ETH (but not in Wei)) in a string format. This does not need to be an integer. For example, if you want to send 0.0005 SOL, then `amount = \"0.0005\"`. If you want to send 0.0005 ETH, then `amount = \"0.0005\"`. - NFT: This must be `1`. - Token: This must be an integer in string format. To convert from what you see on a wallet UI (e.g., 1 ATLAS, 1 USDC) to an integer value, you have to multiply that value by 10^<i>x</i> where <i>x</i> is the number of decimals. For example, to transfer 0.2 USDC, if USDC uses 6 decimals, then the amount is 0.2 * 10^6 = 200000.    - For Solana, you can get the number of decimals for a given SPL token <a href=\"#operation/solanaGetSPLToken\">here</a>. We are working on analogues of this endpoint for other blockchains.
	Amount *string `json:"amount,omitempty"`
	// - If `false`, we sign and submit the transaction (`wallet` is required in this case; do not provide a value for `sender_blockchain_identifier`).  - If `true`, we compile the transaction (either `wallet` or `sender_blockchain_identifier` is required in this case; do not provide both).   
	ReturnCompiledTransaction *bool `json:"return_compiled_transaction,omitempty"`
	// To understand the purpose of `sender_blockchain_identifier` first note the following: There are two ways you can complete a transfer transaction:  - (1) we complete it for you with your `wallet` information or  - (2) we return the raw instruction data that you can sign and submit yourself (no private keys required).  When you provide your `wallet` authentication, we are able to determine your wallet's blockchain identifier (public key or address) which is the sender public key of the transaction.  When you are not providing your `wallet` as authentication, we still need the `sender_blockchain_identifier` to be able to return the compiled transaction. Thus, you provide `sender_blockchain_identifier` if and only if you are not providing `wallet` authentication information **and** you are returning a compiled transaction.  You will receive an error if you provide both `wallet` and `sender_blockchain_identifier`. You will receive an error if you provide neither `wallet` nor `sender_blockchain_identifier`.
	SenderBlockchainIdentifier *string `json:"sender_blockchain_identifier,omitempty"`
	FeePayerWallet *GeneralFeePayerWallet `json:"fee_payer_wallet,omitempty"`
}

// NewGeneralTransferRequest instantiates a new GeneralTransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralTransferRequest(recipientBlockchainIdentifier string) *GeneralTransferRequest {
	this := GeneralTransferRequest{}
	this.RecipientBlockchainIdentifier = recipientBlockchainIdentifier
	var amount string = "1"
	this.Amount = &amount
	var returnCompiledTransaction bool = false
	this.ReturnCompiledTransaction = &returnCompiledTransaction
	var senderBlockchainIdentifier string = "null"
	this.SenderBlockchainIdentifier = &senderBlockchainIdentifier
	return &this
}

// NewGeneralTransferRequestWithDefaults instantiates a new GeneralTransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralTransferRequestWithDefaults() *GeneralTransferRequest {
	this := GeneralTransferRequest{}
	var amount string = "1"
	this.Amount = &amount
	var returnCompiledTransaction bool = false
	this.ReturnCompiledTransaction = &returnCompiledTransaction
	var senderBlockchainIdentifier string = "null"
	this.SenderBlockchainIdentifier = &senderBlockchainIdentifier
	return &this
}

// GetRecipientBlockchainIdentifier returns the RecipientBlockchainIdentifier field value
func (o *GeneralTransferRequest) GetRecipientBlockchainIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientBlockchainIdentifier
}

// GetRecipientBlockchainIdentifierOk returns a tuple with the RecipientBlockchainIdentifier field value
// and a boolean to check if the value has been set.
func (o *GeneralTransferRequest) GetRecipientBlockchainIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientBlockchainIdentifier, true
}

// SetRecipientBlockchainIdentifier sets field value
func (o *GeneralTransferRequest) SetRecipientBlockchainIdentifier(v string) {
	o.RecipientBlockchainIdentifier = v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *GeneralTransferRequest) GetWallet() GeneralWallet {
	if o == nil || o.Wallet == nil {
		var ret GeneralWallet
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralTransferRequest) GetWalletOk() (*GeneralWallet, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *GeneralTransferRequest) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given GeneralWallet and assigns it to the Wallet field.
func (o *GeneralTransferRequest) SetWallet(v GeneralWallet) {
	o.Wallet = &v
}

// GetTokenBlockchainIdentifier returns the TokenBlockchainIdentifier field value if set, zero value otherwise.
func (o *GeneralTransferRequest) GetTokenBlockchainIdentifier() string {
	if o == nil || o.TokenBlockchainIdentifier == nil {
		var ret string
		return ret
	}
	return *o.TokenBlockchainIdentifier
}

// GetTokenBlockchainIdentifierOk returns a tuple with the TokenBlockchainIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralTransferRequest) GetTokenBlockchainIdentifierOk() (*string, bool) {
	if o == nil || o.TokenBlockchainIdentifier == nil {
		return nil, false
	}
	return o.TokenBlockchainIdentifier, true
}

// HasTokenBlockchainIdentifier returns a boolean if a field has been set.
func (o *GeneralTransferRequest) HasTokenBlockchainIdentifier() bool {
	if o != nil && o.TokenBlockchainIdentifier != nil {
		return true
	}

	return false
}

// SetTokenBlockchainIdentifier gets a reference to the given string and assigns it to the TokenBlockchainIdentifier field.
func (o *GeneralTransferRequest) SetTokenBlockchainIdentifier(v string) {
	o.TokenBlockchainIdentifier = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GeneralTransferRequest) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralTransferRequest) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GeneralTransferRequest) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *GeneralTransferRequest) SetNetwork(v string) {
	o.Network = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GeneralTransferRequest) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralTransferRequest) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GeneralTransferRequest) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GeneralTransferRequest) SetAmount(v string) {
	o.Amount = &v
}

// GetReturnCompiledTransaction returns the ReturnCompiledTransaction field value if set, zero value otherwise.
func (o *GeneralTransferRequest) GetReturnCompiledTransaction() bool {
	if o == nil || o.ReturnCompiledTransaction == nil {
		var ret bool
		return ret
	}
	return *o.ReturnCompiledTransaction
}

// GetReturnCompiledTransactionOk returns a tuple with the ReturnCompiledTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralTransferRequest) GetReturnCompiledTransactionOk() (*bool, bool) {
	if o == nil || o.ReturnCompiledTransaction == nil {
		return nil, false
	}
	return o.ReturnCompiledTransaction, true
}

// HasReturnCompiledTransaction returns a boolean if a field has been set.
func (o *GeneralTransferRequest) HasReturnCompiledTransaction() bool {
	if o != nil && o.ReturnCompiledTransaction != nil {
		return true
	}

	return false
}

// SetReturnCompiledTransaction gets a reference to the given bool and assigns it to the ReturnCompiledTransaction field.
func (o *GeneralTransferRequest) SetReturnCompiledTransaction(v bool) {
	o.ReturnCompiledTransaction = &v
}

// GetSenderBlockchainIdentifier returns the SenderBlockchainIdentifier field value if set, zero value otherwise.
func (o *GeneralTransferRequest) GetSenderBlockchainIdentifier() string {
	if o == nil || o.SenderBlockchainIdentifier == nil {
		var ret string
		return ret
	}
	return *o.SenderBlockchainIdentifier
}

// GetSenderBlockchainIdentifierOk returns a tuple with the SenderBlockchainIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralTransferRequest) GetSenderBlockchainIdentifierOk() (*string, bool) {
	if o == nil || o.SenderBlockchainIdentifier == nil {
		return nil, false
	}
	return o.SenderBlockchainIdentifier, true
}

// HasSenderBlockchainIdentifier returns a boolean if a field has been set.
func (o *GeneralTransferRequest) HasSenderBlockchainIdentifier() bool {
	if o != nil && o.SenderBlockchainIdentifier != nil {
		return true
	}

	return false
}

// SetSenderBlockchainIdentifier gets a reference to the given string and assigns it to the SenderBlockchainIdentifier field.
func (o *GeneralTransferRequest) SetSenderBlockchainIdentifier(v string) {
	o.SenderBlockchainIdentifier = &v
}

// GetFeePayerWallet returns the FeePayerWallet field value if set, zero value otherwise.
func (o *GeneralTransferRequest) GetFeePayerWallet() GeneralFeePayerWallet {
	if o == nil || o.FeePayerWallet == nil {
		var ret GeneralFeePayerWallet
		return ret
	}
	return *o.FeePayerWallet
}

// GetFeePayerWalletOk returns a tuple with the FeePayerWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralTransferRequest) GetFeePayerWalletOk() (*GeneralFeePayerWallet, bool) {
	if o == nil || o.FeePayerWallet == nil {
		return nil, false
	}
	return o.FeePayerWallet, true
}

// HasFeePayerWallet returns a boolean if a field has been set.
func (o *GeneralTransferRequest) HasFeePayerWallet() bool {
	if o != nil && o.FeePayerWallet != nil {
		return true
	}

	return false
}

// SetFeePayerWallet gets a reference to the given GeneralFeePayerWallet and assigns it to the FeePayerWallet field.
func (o *GeneralTransferRequest) SetFeePayerWallet(v GeneralFeePayerWallet) {
	o.FeePayerWallet = &v
}

func (o GeneralTransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipient_blockchain_identifier"] = o.RecipientBlockchainIdentifier
	}
	if o.Wallet != nil {
		toSerialize["wallet"] = o.Wallet
	}
	if o.TokenBlockchainIdentifier != nil {
		toSerialize["token_blockchain_identifier"] = o.TokenBlockchainIdentifier
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.ReturnCompiledTransaction != nil {
		toSerialize["return_compiled_transaction"] = o.ReturnCompiledTransaction
	}
	if o.SenderBlockchainIdentifier != nil {
		toSerialize["sender_blockchain_identifier"] = o.SenderBlockchainIdentifier
	}
	if o.FeePayerWallet != nil {
		toSerialize["fee_payer_wallet"] = o.FeePayerWallet
	}
	return json.Marshal(toSerialize)
}

type NullableGeneralTransferRequest struct {
	value *GeneralTransferRequest
	isSet bool
}

func (v NullableGeneralTransferRequest) Get() *GeneralTransferRequest {
	return v.value
}

func (v *NullableGeneralTransferRequest) Set(val *GeneralTransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralTransferRequest(val *GeneralTransferRequest) *NullableGeneralTransferRequest {
	return &NullableGeneralTransferRequest{value: val, isSet: true}
}

func (v NullableGeneralTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


