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

// TransferRequest struct for TransferRequest
type TransferRequest struct {
	// Whether to wait for the transaction to be confirmed on the blockchain or simply be processed.  Processed means that our node has picked up the transaction request, but not that it was confirmed by the Solana cluster.  Confirmed means that the cluster voted on your transaction and approved it. To be completely sure that the transaction succeeded, you can either set `wait_for_confirmation=True` (call takes 20 seconds with True; about 4 seconds for processed) or you can [get the transaction metadata](/#tag/Solana-Transaction/operation/solanaGetTransaction) using the signature in the response returned. Once it returns the metadata, then the transaction should have succeeded. 
	WaitForConfirmation *bool `json:"wait_for_confirmation,omitempty"`
	// The public key address of the recipient to whom you want to send a token or NFT
	RecipientAddress string `json:"recipient_address"`
	Wallet *Wallet `json:"wallet,omitempty"`
	// If you're transfering an NFT, supply the `mint` (the address of the mint) for the `token_address`. If you're transfering a token, supply the token address found on the explorer (e.g., see `SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt` for <a href=\"https://explorer.solana.com/address/SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt\" target=\"_blank\">Serum Token</a>) for the `token_address`. If you're transferring SOL, do not supply a value for `token_address`. 
	TokenAddress *string `json:"token_address,omitempty"`
	Network *string `json:"network,omitempty"`
	// This value must be a string. What you provide here depends on if you are sending an NFT, an SPL token, or SOL.  - NFT: This must be '1'. - SPL Token: This must be an integer in string format. To convert from what you see on a wallet UI (e.g., 1 ATLAS, 1 USDC) to an integer value, you have to multiply that value by 10^<i>x</i> where <i>x</i> is the number of decimals. For example, to transfer 0.2 USDC, if USDC uses 6 decimals, then the amount is 0.2 * 10^6 = 200000. You can get the number of decimals for a given SPL token <a href=\"#operation/solanaGetSPLToken\">here</a>. - SOL: Supply this value denominated in SOL in a string format. This does not need to be an integer. For example, if you want to send 0.0005 SOL, then amount = \"0.0005\".
	Amount *string `json:"amount,omitempty"`
	// If `false`, we sign and submit the transaction (`wallet` is required in this case; do not provide a value for `sender_public_key`).  If `true`, we compile the transaction (either `wallet` or `sender_public_key` is required in this case; do not provide both). 
	ReturnCompiledTransaction *bool `json:"return_compiled_transaction,omitempty"`
	// To understand the purpose of `sender_public_key` first note the following: There are two ways you can complete a transfer transaction: (1) we complete it for you with your `wallet` information or (2) we return the raw instruction data that you can sign and submit yourself (no private keys required). When you provide your `wallet` authentication, we are able to determine your wallet's public key, which is the sender public key of the transaction. When you are not providing your `wallet` as authentication, we still need the `sender_public_key` to be able to return the compiled transaction. Thus, you provide `sender_public_key` only if you are not providing `wallet` authentication information and you are returning a compiled transaction. You will receive an error if you provide both `wallet` and `sender_public_key`. You will receive an error if you provide neither `wallet` nor `sender_public_key`. 
	SenderPublicKey *string `json:"sender_public_key,omitempty"`
	FeePayerWallet *FeePayerWallet `json:"fee_payer_wallet,omitempty"`
}

// NewTransferRequest instantiates a new TransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRequest(recipientAddress string) *TransferRequest {
	this := TransferRequest{}
	var waitForConfirmation bool = true
	this.WaitForConfirmation = &waitForConfirmation
	this.RecipientAddress = recipientAddress
	var network string = "devnet"
	this.Network = &network
	var amount string = "1"
	this.Amount = &amount
	var returnCompiledTransaction bool = false
	this.ReturnCompiledTransaction = &returnCompiledTransaction
	var senderPublicKey string = "null"
	this.SenderPublicKey = &senderPublicKey
	return &this
}

// NewTransferRequestWithDefaults instantiates a new TransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRequestWithDefaults() *TransferRequest {
	this := TransferRequest{}
	var waitForConfirmation bool = true
	this.WaitForConfirmation = &waitForConfirmation
	var network string = "devnet"
	this.Network = &network
	var amount string = "1"
	this.Amount = &amount
	var returnCompiledTransaction bool = false
	this.ReturnCompiledTransaction = &returnCompiledTransaction
	var senderPublicKey string = "null"
	this.SenderPublicKey = &senderPublicKey
	return &this
}

// GetWaitForConfirmation returns the WaitForConfirmation field value if set, zero value otherwise.
func (o *TransferRequest) GetWaitForConfirmation() bool {
	if o == nil || o.WaitForConfirmation == nil {
		var ret bool
		return ret
	}
	return *o.WaitForConfirmation
}

// GetWaitForConfirmationOk returns a tuple with the WaitForConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetWaitForConfirmationOk() (*bool, bool) {
	if o == nil || o.WaitForConfirmation == nil {
		return nil, false
	}
	return o.WaitForConfirmation, true
}

// HasWaitForConfirmation returns a boolean if a field has been set.
func (o *TransferRequest) HasWaitForConfirmation() bool {
	if o != nil && o.WaitForConfirmation != nil {
		return true
	}

	return false
}

// SetWaitForConfirmation gets a reference to the given bool and assigns it to the WaitForConfirmation field.
func (o *TransferRequest) SetWaitForConfirmation(v bool) {
	o.WaitForConfirmation = &v
}

// GetRecipientAddress returns the RecipientAddress field value
func (o *TransferRequest) GetRecipientAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetRecipientAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientAddress, true
}

// SetRecipientAddress sets field value
func (o *TransferRequest) SetRecipientAddress(v string) {
	o.RecipientAddress = v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *TransferRequest) GetWallet() Wallet {
	if o == nil || o.Wallet == nil {
		var ret Wallet
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetWalletOk() (*Wallet, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *TransferRequest) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given Wallet and assigns it to the Wallet field.
func (o *TransferRequest) SetWallet(v Wallet) {
	o.Wallet = &v
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise.
func (o *TransferRequest) GetTokenAddress() string {
	if o == nil || o.TokenAddress == nil {
		var ret string
		return ret
	}
	return *o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetTokenAddressOk() (*string, bool) {
	if o == nil || o.TokenAddress == nil {
		return nil, false
	}
	return o.TokenAddress, true
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *TransferRequest) HasTokenAddress() bool {
	if o != nil && o.TokenAddress != nil {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given string and assigns it to the TokenAddress field.
func (o *TransferRequest) SetTokenAddress(v string) {
	o.TokenAddress = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *TransferRequest) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *TransferRequest) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *TransferRequest) SetNetwork(v string) {
	o.Network = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransferRequest) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransferRequest) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *TransferRequest) SetAmount(v string) {
	o.Amount = &v
}

// GetReturnCompiledTransaction returns the ReturnCompiledTransaction field value if set, zero value otherwise.
func (o *TransferRequest) GetReturnCompiledTransaction() bool {
	if o == nil || o.ReturnCompiledTransaction == nil {
		var ret bool
		return ret
	}
	return *o.ReturnCompiledTransaction
}

// GetReturnCompiledTransactionOk returns a tuple with the ReturnCompiledTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetReturnCompiledTransactionOk() (*bool, bool) {
	if o == nil || o.ReturnCompiledTransaction == nil {
		return nil, false
	}
	return o.ReturnCompiledTransaction, true
}

// HasReturnCompiledTransaction returns a boolean if a field has been set.
func (o *TransferRequest) HasReturnCompiledTransaction() bool {
	if o != nil && o.ReturnCompiledTransaction != nil {
		return true
	}

	return false
}

// SetReturnCompiledTransaction gets a reference to the given bool and assigns it to the ReturnCompiledTransaction field.
func (o *TransferRequest) SetReturnCompiledTransaction(v bool) {
	o.ReturnCompiledTransaction = &v
}

// GetSenderPublicKey returns the SenderPublicKey field value if set, zero value otherwise.
func (o *TransferRequest) GetSenderPublicKey() string {
	if o == nil || o.SenderPublicKey == nil {
		var ret string
		return ret
	}
	return *o.SenderPublicKey
}

// GetSenderPublicKeyOk returns a tuple with the SenderPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetSenderPublicKeyOk() (*string, bool) {
	if o == nil || o.SenderPublicKey == nil {
		return nil, false
	}
	return o.SenderPublicKey, true
}

// HasSenderPublicKey returns a boolean if a field has been set.
func (o *TransferRequest) HasSenderPublicKey() bool {
	if o != nil && o.SenderPublicKey != nil {
		return true
	}

	return false
}

// SetSenderPublicKey gets a reference to the given string and assigns it to the SenderPublicKey field.
func (o *TransferRequest) SetSenderPublicKey(v string) {
	o.SenderPublicKey = &v
}

// GetFeePayerWallet returns the FeePayerWallet field value if set, zero value otherwise.
func (o *TransferRequest) GetFeePayerWallet() FeePayerWallet {
	if o == nil || o.FeePayerWallet == nil {
		var ret FeePayerWallet
		return ret
	}
	return *o.FeePayerWallet
}

// GetFeePayerWalletOk returns a tuple with the FeePayerWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetFeePayerWalletOk() (*FeePayerWallet, bool) {
	if o == nil || o.FeePayerWallet == nil {
		return nil, false
	}
	return o.FeePayerWallet, true
}

// HasFeePayerWallet returns a boolean if a field has been set.
func (o *TransferRequest) HasFeePayerWallet() bool {
	if o != nil && o.FeePayerWallet != nil {
		return true
	}

	return false
}

// SetFeePayerWallet gets a reference to the given FeePayerWallet and assigns it to the FeePayerWallet field.
func (o *TransferRequest) SetFeePayerWallet(v FeePayerWallet) {
	o.FeePayerWallet = &v
}

func (o TransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WaitForConfirmation != nil {
		toSerialize["wait_for_confirmation"] = o.WaitForConfirmation
	}
	if true {
		toSerialize["recipient_address"] = o.RecipientAddress
	}
	if o.Wallet != nil {
		toSerialize["wallet"] = o.Wallet
	}
	if o.TokenAddress != nil {
		toSerialize["token_address"] = o.TokenAddress
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
	if o.SenderPublicKey != nil {
		toSerialize["sender_public_key"] = o.SenderPublicKey
	}
	if o.FeePayerWallet != nil {
		toSerialize["fee_payer_wallet"] = o.FeePayerWallet
	}
	return json.Marshal(toSerialize)
}

type NullableTransferRequest struct {
	value *TransferRequest
	isSet bool
}

func (v NullableTransferRequest) Get() *TransferRequest {
	return v.value
}

func (v *NullableTransferRequest) Set(val *TransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRequest(val *TransferRequest) *NullableTransferRequest {
	return &NullableTransferRequest{value: val, isSet: true}
}

func (v NullableTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


