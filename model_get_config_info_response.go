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

// GetConfigInfoResponse struct for GetConfigInfoResponse
type GetConfigInfoResponse struct {
	Authority *string `json:"authority,omitempty"`
	Creators *[]GetConfigInfoResponseCreators `json:"creators,omitempty"`
	IsMutable *bool `json:"is_mutable,omitempty"`
	// A public key address
	RetainAuthority *string `json:"retain_authority,omitempty"`
	// The fee of each sale the creators receive. 100 basis points = 1%
	SellerFeeBasisPoints *float32 `json:"seller_fee_basis_points,omitempty"`
	// The symbol of the candy machine NFT colletion
	Symbol *string `json:"symbol,omitempty"`
}

// NewGetConfigInfoResponse instantiates a new GetConfigInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConfigInfoResponse() *GetConfigInfoResponse {
	this := GetConfigInfoResponse{}
	return &this
}

// NewGetConfigInfoResponseWithDefaults instantiates a new GetConfigInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConfigInfoResponseWithDefaults() *GetConfigInfoResponse {
	this := GetConfigInfoResponse{}
	return &this
}

// GetAuthority returns the Authority field value if set, zero value otherwise.
func (o *GetConfigInfoResponse) GetAuthority() string {
	if o == nil || o.Authority == nil {
		var ret string
		return ret
	}
	return *o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfigInfoResponse) GetAuthorityOk() (*string, bool) {
	if o == nil || o.Authority == nil {
		return nil, false
	}
	return o.Authority, true
}

// HasAuthority returns a boolean if a field has been set.
func (o *GetConfigInfoResponse) HasAuthority() bool {
	if o != nil && o.Authority != nil {
		return true
	}

	return false
}

// SetAuthority gets a reference to the given string and assigns it to the Authority field.
func (o *GetConfigInfoResponse) SetAuthority(v string) {
	o.Authority = &v
}

// GetCreators returns the Creators field value if set, zero value otherwise.
func (o *GetConfigInfoResponse) GetCreators() []GetConfigInfoResponseCreators {
	if o == nil || o.Creators == nil {
		var ret []GetConfigInfoResponseCreators
		return ret
	}
	return *o.Creators
}

// GetCreatorsOk returns a tuple with the Creators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfigInfoResponse) GetCreatorsOk() (*[]GetConfigInfoResponseCreators, bool) {
	if o == nil || o.Creators == nil {
		return nil, false
	}
	return o.Creators, true
}

// HasCreators returns a boolean if a field has been set.
func (o *GetConfigInfoResponse) HasCreators() bool {
	if o != nil && o.Creators != nil {
		return true
	}

	return false
}

// SetCreators gets a reference to the given []GetConfigInfoResponseCreators and assigns it to the Creators field.
func (o *GetConfigInfoResponse) SetCreators(v []GetConfigInfoResponseCreators) {
	o.Creators = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *GetConfigInfoResponse) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfigInfoResponse) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *GetConfigInfoResponse) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *GetConfigInfoResponse) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetRetainAuthority returns the RetainAuthority field value if set, zero value otherwise.
func (o *GetConfigInfoResponse) GetRetainAuthority() string {
	if o == nil || o.RetainAuthority == nil {
		var ret string
		return ret
	}
	return *o.RetainAuthority
}

// GetRetainAuthorityOk returns a tuple with the RetainAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfigInfoResponse) GetRetainAuthorityOk() (*string, bool) {
	if o == nil || o.RetainAuthority == nil {
		return nil, false
	}
	return o.RetainAuthority, true
}

// HasRetainAuthority returns a boolean if a field has been set.
func (o *GetConfigInfoResponse) HasRetainAuthority() bool {
	if o != nil && o.RetainAuthority != nil {
		return true
	}

	return false
}

// SetRetainAuthority gets a reference to the given string and assigns it to the RetainAuthority field.
func (o *GetConfigInfoResponse) SetRetainAuthority(v string) {
	o.RetainAuthority = &v
}

// GetSellerFeeBasisPoints returns the SellerFeeBasisPoints field value if set, zero value otherwise.
func (o *GetConfigInfoResponse) GetSellerFeeBasisPoints() float32 {
	if o == nil || o.SellerFeeBasisPoints == nil {
		var ret float32
		return ret
	}
	return *o.SellerFeeBasisPoints
}

// GetSellerFeeBasisPointsOk returns a tuple with the SellerFeeBasisPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfigInfoResponse) GetSellerFeeBasisPointsOk() (*float32, bool) {
	if o == nil || o.SellerFeeBasisPoints == nil {
		return nil, false
	}
	return o.SellerFeeBasisPoints, true
}

// HasSellerFeeBasisPoints returns a boolean if a field has been set.
func (o *GetConfigInfoResponse) HasSellerFeeBasisPoints() bool {
	if o != nil && o.SellerFeeBasisPoints != nil {
		return true
	}

	return false
}

// SetSellerFeeBasisPoints gets a reference to the given float32 and assigns it to the SellerFeeBasisPoints field.
func (o *GetConfigInfoResponse) SetSellerFeeBasisPoints(v float32) {
	o.SellerFeeBasisPoints = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetConfigInfoResponse) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfigInfoResponse) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetConfigInfoResponse) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetConfigInfoResponse) SetSymbol(v string) {
	o.Symbol = &v
}

func (o GetConfigInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authority != nil {
		toSerialize["authority"] = o.Authority
	}
	if o.Creators != nil {
		toSerialize["creators"] = o.Creators
	}
	if o.IsMutable != nil {
		toSerialize["is_mutable"] = o.IsMutable
	}
	if o.RetainAuthority != nil {
		toSerialize["retain_authority"] = o.RetainAuthority
	}
	if o.SellerFeeBasisPoints != nil {
		toSerialize["seller_fee_basis_points"] = o.SellerFeeBasisPoints
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableGetConfigInfoResponse struct {
	value *GetConfigInfoResponse
	isSet bool
}

func (v NullableGetConfigInfoResponse) Get() *GetConfigInfoResponse {
	return v.value
}

func (v *NullableGetConfigInfoResponse) Set(val *GetConfigInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConfigInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConfigInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConfigInfoResponse(val *GetConfigInfoResponse) *NullableGetConfigInfoResponse {
	return &NullableGetConfigInfoResponse{value: val, isSet: true}
}

func (v NullableGetConfigInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConfigInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


