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

// CCProject You can change any of these parameters after creating the project. None of these parameters are required. 
type CCProject struct {
	// The ID of the project. This is auto-generated upon project creation and cannot currently be changed. 
	ProjectId string `json:"project_id"`
	// The name of the project. This is shown to your users and should identify your company or organization. 
	Name string `json:"name"`
	// The description of your project / company. 
	Description *string `json:"description,omitempty"`
	// A URL that identifies where we should make an API request to notify you of a new payment (e.g., api.myproject.com/crypto_payments/webhook). Learn more [here](#tag/CC-Webhook/operation/getCCWebhook). 
	Webhook *string `json:"webhook,omitempty"`
	// The website of your project / company. 
	Website *string `json:"website,omitempty"`
	// A Discord webhook. We will send a message to this channel to notify of payment. Learn more [here](). 
	DiscordWebhook *string `json:"discord_webhook,omitempty"`
	// A URL of your logo. 
	LogoUrl *string `json:"logo_url,omitempty"`
	CustomerIdToCollect *CCProjectCreateRequestCustomerIdToCollect `json:"customer_id_to_collect,omitempty"`
	// Where to redirect customers after payment. If not supplied, customers will be redirected to checkout.blockchainapi.com/me to view their subscriptions. 
	SuccessUrl *string `json:"success_url,omitempty"`
	PayoutMethod *CCProjectCreateRequestPayoutMethod `json:"payout_method,omitempty"`
}

// NewCCProject instantiates a new CCProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCProject(projectId string, name string) *CCProject {
	this := CCProject{}
	this.ProjectId = projectId
	this.Name = name
	return &this
}

// NewCCProjectWithDefaults instantiates a new CCProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCProjectWithDefaults() *CCProject {
	this := CCProject{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *CCProject) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CCProject) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CCProject) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *CCProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CCProject) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CCProject) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CCProject) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProject) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CCProject) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CCProject) SetDescription(v string) {
	o.Description = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *CCProject) GetWebhook() string {
	if o == nil || o.Webhook == nil {
		var ret string
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProject) GetWebhookOk() (*string, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *CCProject) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given string and assigns it to the Webhook field.
func (o *CCProject) SetWebhook(v string) {
	o.Webhook = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *CCProject) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProject) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *CCProject) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *CCProject) SetWebsite(v string) {
	o.Website = &v
}

// GetDiscordWebhook returns the DiscordWebhook field value if set, zero value otherwise.
func (o *CCProject) GetDiscordWebhook() string {
	if o == nil || o.DiscordWebhook == nil {
		var ret string
		return ret
	}
	return *o.DiscordWebhook
}

// GetDiscordWebhookOk returns a tuple with the DiscordWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProject) GetDiscordWebhookOk() (*string, bool) {
	if o == nil || o.DiscordWebhook == nil {
		return nil, false
	}
	return o.DiscordWebhook, true
}

// HasDiscordWebhook returns a boolean if a field has been set.
func (o *CCProject) HasDiscordWebhook() bool {
	if o != nil && o.DiscordWebhook != nil {
		return true
	}

	return false
}

// SetDiscordWebhook gets a reference to the given string and assigns it to the DiscordWebhook field.
func (o *CCProject) SetDiscordWebhook(v string) {
	o.DiscordWebhook = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *CCProject) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProject) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *CCProject) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *CCProject) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetCustomerIdToCollect returns the CustomerIdToCollect field value if set, zero value otherwise.
func (o *CCProject) GetCustomerIdToCollect() CCProjectCreateRequestCustomerIdToCollect {
	if o == nil || o.CustomerIdToCollect == nil {
		var ret CCProjectCreateRequestCustomerIdToCollect
		return ret
	}
	return *o.CustomerIdToCollect
}

// GetCustomerIdToCollectOk returns a tuple with the CustomerIdToCollect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProject) GetCustomerIdToCollectOk() (*CCProjectCreateRequestCustomerIdToCollect, bool) {
	if o == nil || o.CustomerIdToCollect == nil {
		return nil, false
	}
	return o.CustomerIdToCollect, true
}

// HasCustomerIdToCollect returns a boolean if a field has been set.
func (o *CCProject) HasCustomerIdToCollect() bool {
	if o != nil && o.CustomerIdToCollect != nil {
		return true
	}

	return false
}

// SetCustomerIdToCollect gets a reference to the given CCProjectCreateRequestCustomerIdToCollect and assigns it to the CustomerIdToCollect field.
func (o *CCProject) SetCustomerIdToCollect(v CCProjectCreateRequestCustomerIdToCollect) {
	o.CustomerIdToCollect = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *CCProject) GetSuccessUrl() string {
	if o == nil || o.SuccessUrl == nil {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProject) GetSuccessUrlOk() (*string, bool) {
	if o == nil || o.SuccessUrl == nil {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *CCProject) HasSuccessUrl() bool {
	if o != nil && o.SuccessUrl != nil {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *CCProject) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetPayoutMethod returns the PayoutMethod field value if set, zero value otherwise.
func (o *CCProject) GetPayoutMethod() CCProjectCreateRequestPayoutMethod {
	if o == nil || o.PayoutMethod == nil {
		var ret CCProjectCreateRequestPayoutMethod
		return ret
	}
	return *o.PayoutMethod
}

// GetPayoutMethodOk returns a tuple with the PayoutMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCProject) GetPayoutMethodOk() (*CCProjectCreateRequestPayoutMethod, bool) {
	if o == nil || o.PayoutMethod == nil {
		return nil, false
	}
	return o.PayoutMethod, true
}

// HasPayoutMethod returns a boolean if a field has been set.
func (o *CCProject) HasPayoutMethod() bool {
	if o != nil && o.PayoutMethod != nil {
		return true
	}

	return false
}

// SetPayoutMethod gets a reference to the given CCProjectCreateRequestPayoutMethod and assigns it to the PayoutMethod field.
func (o *CCProject) SetPayoutMethod(v CCProjectCreateRequestPayoutMethod) {
	o.PayoutMethod = &v
}

func (o CCProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project_id"] = o.ProjectId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.DiscordWebhook != nil {
		toSerialize["discord_webhook"] = o.DiscordWebhook
	}
	if o.LogoUrl != nil {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if o.CustomerIdToCollect != nil {
		toSerialize["customer_id_to_collect"] = o.CustomerIdToCollect
	}
	if o.SuccessUrl != nil {
		toSerialize["success_url"] = o.SuccessUrl
	}
	if o.PayoutMethod != nil {
		toSerialize["payout_method"] = o.PayoutMethod
	}
	return json.Marshal(toSerialize)
}

type NullableCCProject struct {
	value *CCProject
	isSet bool
}

func (v NullableCCProject) Get() *CCProject {
	return v.value
}

func (v *NullableCCProject) Set(val *CCProject) {
	v.value = val
	v.isSet = true
}

func (v NullableCCProject) IsSet() bool {
	return v.isSet
}

func (v *NullableCCProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCProject(val *CCProject) *NullableCCProject {
	return &NullableCCProject{value: val, isSet: true}
}

func (v NullableCCProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


