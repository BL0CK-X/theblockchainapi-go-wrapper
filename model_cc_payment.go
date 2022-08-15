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

// CCPayment struct for CCPayment
type CCPayment struct {
	// The ID of the project 
	ProjectId *string `json:"project_id,omitempty"`
	// The ID of the associated product 
	ProductId *string `json:"product_id,omitempty"`
	// The ID of the respective plan 
	PlanId *string `json:"plan_id,omitempty"`
	// The unique ID of the payment 
	PaymentId *string `json:"payment_id,omitempty"`
	// The unique identifier of the wallet from which the payment was made. 
	BlockchainIdentifier *string `json:"blockchain_identifier,omitempty"`
	BlockchainPaymentDetails *CCPaymentBlockchainPaymentDetails `json:"blockchain_payment_details,omitempty"`
	// The ID of the customer 
	CustomerId *string `json:"customer_id,omitempty"`
	// The validation code shown to the customer. This is only visible to the customer who paid. They can use this code to redeem their subscription to their product. 
	PaymentValidationCode *string `json:"payment_validation_code,omitempty"`
	// A UNIX time stamp, in seconds, that identifies the end of the period of the subscription 
	PeriodEnd *float32 `json:"period_end,omitempty"`
	// A UNIX time stamp, in seconds, that identifies the start of the period of the subscription 
	PeriodStart *float32 `json:"period_start,omitempty"`
	// The string that uniquely identifies the blockchain transaction 
	TransactionBlockchainIdentifier *string `json:"transaction_blockchain_identifier,omitempty"`
}

// NewCCPayment instantiates a new CCPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCPayment() *CCPayment {
	this := CCPayment{}
	return &this
}

// NewCCPaymentWithDefaults instantiates a new CCPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCPaymentWithDefaults() *CCPayment {
	this := CCPayment{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CCPayment) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CCPayment) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *CCPayment) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *CCPayment) GetProductId() string {
	if o == nil || o.ProductId == nil {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetProductIdOk() (*string, bool) {
	if o == nil || o.ProductId == nil {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *CCPayment) HasProductId() bool {
	if o != nil && o.ProductId != nil {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *CCPayment) SetProductId(v string) {
	o.ProductId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *CCPayment) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *CCPayment) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *CCPayment) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *CCPayment) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *CCPayment) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *CCPayment) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetBlockchainIdentifier returns the BlockchainIdentifier field value if set, zero value otherwise.
func (o *CCPayment) GetBlockchainIdentifier() string {
	if o == nil || o.BlockchainIdentifier == nil {
		var ret string
		return ret
	}
	return *o.BlockchainIdentifier
}

// GetBlockchainIdentifierOk returns a tuple with the BlockchainIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetBlockchainIdentifierOk() (*string, bool) {
	if o == nil || o.BlockchainIdentifier == nil {
		return nil, false
	}
	return o.BlockchainIdentifier, true
}

// HasBlockchainIdentifier returns a boolean if a field has been set.
func (o *CCPayment) HasBlockchainIdentifier() bool {
	if o != nil && o.BlockchainIdentifier != nil {
		return true
	}

	return false
}

// SetBlockchainIdentifier gets a reference to the given string and assigns it to the BlockchainIdentifier field.
func (o *CCPayment) SetBlockchainIdentifier(v string) {
	o.BlockchainIdentifier = &v
}

// GetBlockchainPaymentDetails returns the BlockchainPaymentDetails field value if set, zero value otherwise.
func (o *CCPayment) GetBlockchainPaymentDetails() CCPaymentBlockchainPaymentDetails {
	if o == nil || o.BlockchainPaymentDetails == nil {
		var ret CCPaymentBlockchainPaymentDetails
		return ret
	}
	return *o.BlockchainPaymentDetails
}

// GetBlockchainPaymentDetailsOk returns a tuple with the BlockchainPaymentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetBlockchainPaymentDetailsOk() (*CCPaymentBlockchainPaymentDetails, bool) {
	if o == nil || o.BlockchainPaymentDetails == nil {
		return nil, false
	}
	return o.BlockchainPaymentDetails, true
}

// HasBlockchainPaymentDetails returns a boolean if a field has been set.
func (o *CCPayment) HasBlockchainPaymentDetails() bool {
	if o != nil && o.BlockchainPaymentDetails != nil {
		return true
	}

	return false
}

// SetBlockchainPaymentDetails gets a reference to the given CCPaymentBlockchainPaymentDetails and assigns it to the BlockchainPaymentDetails field.
func (o *CCPayment) SetBlockchainPaymentDetails(v CCPaymentBlockchainPaymentDetails) {
	o.BlockchainPaymentDetails = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CCPayment) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CCPayment) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CCPayment) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetPaymentValidationCode returns the PaymentValidationCode field value if set, zero value otherwise.
func (o *CCPayment) GetPaymentValidationCode() string {
	if o == nil || o.PaymentValidationCode == nil {
		var ret string
		return ret
	}
	return *o.PaymentValidationCode
}

// GetPaymentValidationCodeOk returns a tuple with the PaymentValidationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetPaymentValidationCodeOk() (*string, bool) {
	if o == nil || o.PaymentValidationCode == nil {
		return nil, false
	}
	return o.PaymentValidationCode, true
}

// HasPaymentValidationCode returns a boolean if a field has been set.
func (o *CCPayment) HasPaymentValidationCode() bool {
	if o != nil && o.PaymentValidationCode != nil {
		return true
	}

	return false
}

// SetPaymentValidationCode gets a reference to the given string and assigns it to the PaymentValidationCode field.
func (o *CCPayment) SetPaymentValidationCode(v string) {
	o.PaymentValidationCode = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *CCPayment) GetPeriodEnd() float32 {
	if o == nil || o.PeriodEnd == nil {
		var ret float32
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetPeriodEndOk() (*float32, bool) {
	if o == nil || o.PeriodEnd == nil {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *CCPayment) HasPeriodEnd() bool {
	if o != nil && o.PeriodEnd != nil {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given float32 and assigns it to the PeriodEnd field.
func (o *CCPayment) SetPeriodEnd(v float32) {
	o.PeriodEnd = &v
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *CCPayment) GetPeriodStart() float32 {
	if o == nil || o.PeriodStart == nil {
		var ret float32
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetPeriodStartOk() (*float32, bool) {
	if o == nil || o.PeriodStart == nil {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *CCPayment) HasPeriodStart() bool {
	if o != nil && o.PeriodStart != nil {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given float32 and assigns it to the PeriodStart field.
func (o *CCPayment) SetPeriodStart(v float32) {
	o.PeriodStart = &v
}

// GetTransactionBlockchainIdentifier returns the TransactionBlockchainIdentifier field value if set, zero value otherwise.
func (o *CCPayment) GetTransactionBlockchainIdentifier() string {
	if o == nil || o.TransactionBlockchainIdentifier == nil {
		var ret string
		return ret
	}
	return *o.TransactionBlockchainIdentifier
}

// GetTransactionBlockchainIdentifierOk returns a tuple with the TransactionBlockchainIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCPayment) GetTransactionBlockchainIdentifierOk() (*string, bool) {
	if o == nil || o.TransactionBlockchainIdentifier == nil {
		return nil, false
	}
	return o.TransactionBlockchainIdentifier, true
}

// HasTransactionBlockchainIdentifier returns a boolean if a field has been set.
func (o *CCPayment) HasTransactionBlockchainIdentifier() bool {
	if o != nil && o.TransactionBlockchainIdentifier != nil {
		return true
	}

	return false
}

// SetTransactionBlockchainIdentifier gets a reference to the given string and assigns it to the TransactionBlockchainIdentifier field.
func (o *CCPayment) SetTransactionBlockchainIdentifier(v string) {
	o.TransactionBlockchainIdentifier = &v
}

func (o CCPayment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.ProductId != nil {
		toSerialize["product_id"] = o.ProductId
	}
	if o.PlanId != nil {
		toSerialize["plan_id"] = o.PlanId
	}
	if o.PaymentId != nil {
		toSerialize["payment_id"] = o.PaymentId
	}
	if o.BlockchainIdentifier != nil {
		toSerialize["blockchain_identifier"] = o.BlockchainIdentifier
	}
	if o.BlockchainPaymentDetails != nil {
		toSerialize["blockchain_payment_details"] = o.BlockchainPaymentDetails
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.PaymentValidationCode != nil {
		toSerialize["payment_validation_code"] = o.PaymentValidationCode
	}
	if o.PeriodEnd != nil {
		toSerialize["period_end"] = o.PeriodEnd
	}
	if o.PeriodStart != nil {
		toSerialize["period_start"] = o.PeriodStart
	}
	if o.TransactionBlockchainIdentifier != nil {
		toSerialize["transaction_blockchain_identifier"] = o.TransactionBlockchainIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableCCPayment struct {
	value *CCPayment
	isSet bool
}

func (v NullableCCPayment) Get() *CCPayment {
	return v.value
}

func (v *NullableCCPayment) Set(val *CCPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableCCPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableCCPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCPayment(val *CCPayment) *NullableCCPayment {
	return &NullableCCPayment{value: val, isSet: true}
}

func (v NullableCCPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


