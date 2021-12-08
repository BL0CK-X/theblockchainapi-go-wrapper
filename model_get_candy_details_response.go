/*
The Blockchain API

# About  Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly.   You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>  # How to Use the API  To use the API, you simply need to create an API key pair.  Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.theblockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.  # Feature Requests  Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@theblockchainapi.com).  # Contact  <figure>     <img          width=\"40px\"         height=\"40px\"          src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"     />     <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption> </figure> <a href=\"mailto:info@theblockchainapi.com\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"         />         <figcaption style=\"textAlign:center;\">Email us: info@theblockchainapi.com</figcaption>     </figure> </a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"         />         <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>     </figure> </a> <a href=\"https://twitter.com/_BlockX_\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"         />          <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>     </figure> </a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"         />         <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>     </figure> </a>  # Security  Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet.   <b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b>   How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.  We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.  Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.  Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.  # Pricing  <b>Each user receives 500 free credits every month.</b>  You can learn more about our pricing <a href=\"https://dashboard.theblockchainapi.com/billing\" target=\"_blank\">here</a>.   We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.  If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.  # SDKs / API Wrappers  We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.  We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.  `pip install theblockchainapi`  We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/javascript\" target=\"_blank\">JavaScript Wrapper</a>.  `npm install theblockchainapi`  We also have auto-generated wrappers for the following languages: - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">Go</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/java\" target = \"_blank\">Java</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">Kotlin</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">PHP</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">Swift5</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/tree/main/go\" target = \"_blank\">TypeScript</a>  If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.  If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

API version: null
Contact: info@theblockchainapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetCandyDetailsResponse struct for GetCandyDetailsResponse
type GetCandyDetailsResponse struct {
	// The uuid of the candy machine
	Uuid *string `json:"uuid,omitempty"`
	// The unix timestamp of the start date of the candy machine 
	GoLiveDate *float32 `json:"go_live_date,omitempty"`
	// The price in Lamports for minting an NFT from the candy machine. 1e9 Lamport  = 1 SOL 
	Price *float32 `json:"price,omitempty"`
	// The number of NFTs available for mint from the candy machine 
	ItemsAvailable *float32 `json:"items_available,omitempty"`
	// The number of NFTs minted already from the candy machine 
	ItemsRedeemed *float32 `json:"items_redeemed,omitempty"`
	TokenMint *string `json:"token_mint,omitempty"`
	// The configuration public key address of the candy machine 
	ConfigAddress *string `json:"config_address,omitempty"`
	// The public key address of the wallet that recevies the proceeds from NFT mints 
	Wallet *string `json:"wallet,omitempty"`
	Bump *float32 `json:"bump,omitempty"`
}

// NewGetCandyDetailsResponse instantiates a new GetCandyDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCandyDetailsResponse() *GetCandyDetailsResponse {
	this := GetCandyDetailsResponse{}
	return &this
}

// NewGetCandyDetailsResponseWithDefaults instantiates a new GetCandyDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCandyDetailsResponseWithDefaults() *GetCandyDetailsResponse {
	this := GetCandyDetailsResponse{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GetCandyDetailsResponse) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyDetailsResponse) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *GetCandyDetailsResponse) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GetCandyDetailsResponse) SetUuid(v string) {
	o.Uuid = &v
}

// GetGoLiveDate returns the GoLiveDate field value if set, zero value otherwise.
func (o *GetCandyDetailsResponse) GetGoLiveDate() float32 {
	if o == nil || o.GoLiveDate == nil {
		var ret float32
		return ret
	}
	return *o.GoLiveDate
}

// GetGoLiveDateOk returns a tuple with the GoLiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyDetailsResponse) GetGoLiveDateOk() (*float32, bool) {
	if o == nil || o.GoLiveDate == nil {
		return nil, false
	}
	return o.GoLiveDate, true
}

// HasGoLiveDate returns a boolean if a field has been set.
func (o *GetCandyDetailsResponse) HasGoLiveDate() bool {
	if o != nil && o.GoLiveDate != nil {
		return true
	}

	return false
}

// SetGoLiveDate gets a reference to the given float32 and assigns it to the GoLiveDate field.
func (o *GetCandyDetailsResponse) SetGoLiveDate(v float32) {
	o.GoLiveDate = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetCandyDetailsResponse) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyDetailsResponse) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetCandyDetailsResponse) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *GetCandyDetailsResponse) SetPrice(v float32) {
	o.Price = &v
}

// GetItemsAvailable returns the ItemsAvailable field value if set, zero value otherwise.
func (o *GetCandyDetailsResponse) GetItemsAvailable() float32 {
	if o == nil || o.ItemsAvailable == nil {
		var ret float32
		return ret
	}
	return *o.ItemsAvailable
}

// GetItemsAvailableOk returns a tuple with the ItemsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyDetailsResponse) GetItemsAvailableOk() (*float32, bool) {
	if o == nil || o.ItemsAvailable == nil {
		return nil, false
	}
	return o.ItemsAvailable, true
}

// HasItemsAvailable returns a boolean if a field has been set.
func (o *GetCandyDetailsResponse) HasItemsAvailable() bool {
	if o != nil && o.ItemsAvailable != nil {
		return true
	}

	return false
}

// SetItemsAvailable gets a reference to the given float32 and assigns it to the ItemsAvailable field.
func (o *GetCandyDetailsResponse) SetItemsAvailable(v float32) {
	o.ItemsAvailable = &v
}

// GetItemsRedeemed returns the ItemsRedeemed field value if set, zero value otherwise.
func (o *GetCandyDetailsResponse) GetItemsRedeemed() float32 {
	if o == nil || o.ItemsRedeemed == nil {
		var ret float32
		return ret
	}
	return *o.ItemsRedeemed
}

// GetItemsRedeemedOk returns a tuple with the ItemsRedeemed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyDetailsResponse) GetItemsRedeemedOk() (*float32, bool) {
	if o == nil || o.ItemsRedeemed == nil {
		return nil, false
	}
	return o.ItemsRedeemed, true
}

// HasItemsRedeemed returns a boolean if a field has been set.
func (o *GetCandyDetailsResponse) HasItemsRedeemed() bool {
	if o != nil && o.ItemsRedeemed != nil {
		return true
	}

	return false
}

// SetItemsRedeemed gets a reference to the given float32 and assigns it to the ItemsRedeemed field.
func (o *GetCandyDetailsResponse) SetItemsRedeemed(v float32) {
	o.ItemsRedeemed = &v
}

// GetTokenMint returns the TokenMint field value if set, zero value otherwise.
func (o *GetCandyDetailsResponse) GetTokenMint() string {
	if o == nil || o.TokenMint == nil {
		var ret string
		return ret
	}
	return *o.TokenMint
}

// GetTokenMintOk returns a tuple with the TokenMint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyDetailsResponse) GetTokenMintOk() (*string, bool) {
	if o == nil || o.TokenMint == nil {
		return nil, false
	}
	return o.TokenMint, true
}

// HasTokenMint returns a boolean if a field has been set.
func (o *GetCandyDetailsResponse) HasTokenMint() bool {
	if o != nil && o.TokenMint != nil {
		return true
	}

	return false
}

// SetTokenMint gets a reference to the given string and assigns it to the TokenMint field.
func (o *GetCandyDetailsResponse) SetTokenMint(v string) {
	o.TokenMint = &v
}

// GetConfigAddress returns the ConfigAddress field value if set, zero value otherwise.
func (o *GetCandyDetailsResponse) GetConfigAddress() string {
	if o == nil || o.ConfigAddress == nil {
		var ret string
		return ret
	}
	return *o.ConfigAddress
}

// GetConfigAddressOk returns a tuple with the ConfigAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyDetailsResponse) GetConfigAddressOk() (*string, bool) {
	if o == nil || o.ConfigAddress == nil {
		return nil, false
	}
	return o.ConfigAddress, true
}

// HasConfigAddress returns a boolean if a field has been set.
func (o *GetCandyDetailsResponse) HasConfigAddress() bool {
	if o != nil && o.ConfigAddress != nil {
		return true
	}

	return false
}

// SetConfigAddress gets a reference to the given string and assigns it to the ConfigAddress field.
func (o *GetCandyDetailsResponse) SetConfigAddress(v string) {
	o.ConfigAddress = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *GetCandyDetailsResponse) GetWallet() string {
	if o == nil || o.Wallet == nil {
		var ret string
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyDetailsResponse) GetWalletOk() (*string, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *GetCandyDetailsResponse) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given string and assigns it to the Wallet field.
func (o *GetCandyDetailsResponse) SetWallet(v string) {
	o.Wallet = &v
}

// GetBump returns the Bump field value if set, zero value otherwise.
func (o *GetCandyDetailsResponse) GetBump() float32 {
	if o == nil || o.Bump == nil {
		var ret float32
		return ret
	}
	return *o.Bump
}

// GetBumpOk returns a tuple with the Bump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyDetailsResponse) GetBumpOk() (*float32, bool) {
	if o == nil || o.Bump == nil {
		return nil, false
	}
	return o.Bump, true
}

// HasBump returns a boolean if a field has been set.
func (o *GetCandyDetailsResponse) HasBump() bool {
	if o != nil && o.Bump != nil {
		return true
	}

	return false
}

// SetBump gets a reference to the given float32 and assigns it to the Bump field.
func (o *GetCandyDetailsResponse) SetBump(v float32) {
	o.Bump = &v
}

func (o GetCandyDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.GoLiveDate != nil {
		toSerialize["go_live_date"] = o.GoLiveDate
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.ItemsAvailable != nil {
		toSerialize["items_available"] = o.ItemsAvailable
	}
	if o.ItemsRedeemed != nil {
		toSerialize["items_redeemed"] = o.ItemsRedeemed
	}
	if o.TokenMint != nil {
		toSerialize["token_mint"] = o.TokenMint
	}
	if o.ConfigAddress != nil {
		toSerialize["config_address"] = o.ConfigAddress
	}
	if o.Wallet != nil {
		toSerialize["wallet"] = o.Wallet
	}
	if o.Bump != nil {
		toSerialize["bump"] = o.Bump
	}
	return json.Marshal(toSerialize)
}

type NullableGetCandyDetailsResponse struct {
	value *GetCandyDetailsResponse
	isSet bool
}

func (v NullableGetCandyDetailsResponse) Get() *GetCandyDetailsResponse {
	return v.value
}

func (v *NullableGetCandyDetailsResponse) Set(val *GetCandyDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCandyDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCandyDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCandyDetailsResponse(val *GetCandyDetailsResponse) *NullableGetCandyDetailsResponse {
	return &NullableGetCandyDetailsResponse{value: val, isSet: true}
}

func (v NullableGetCandyDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCandyDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


