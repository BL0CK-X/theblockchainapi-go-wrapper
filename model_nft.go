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

// NFT struct for NFT
type NFT struct {
	// Whether or not the NFT mint was confirmed or simply submitted for processing. The status depends on your input for `wait_for_confirmation`. This only is returned when you are minting an NFT, not when searching or retrieving the metadata.
	Confirmed *bool `json:"confirmed,omitempty"`
	Data *NFTData `json:"data,omitempty"`
	IsMutable *bool `json:"is_mutable,omitempty"`
	// The public key address of the NFT 
	Mint *string `json:"mint,omitempty"`
	PrimarySaleHappened *bool `json:"primary_sale_happened,omitempty"`
	// A public key address that is usually that of the person who minted the NFT 
	UpdateAuthority *string `json:"update_authority,omitempty"`
	SellerFeeBasisPoints *float32 `json:"seller_fee_basis_points,omitempty"`
	MintSecretRecoveryPhrase *string `json:"mint_secret_recovery_phrase,omitempty"`
	ExplorerUrl *string `json:"explorer_url,omitempty"`
	// The metadata account of the NFT 
	MetadataAccount *string `json:"metadata_account,omitempty"`
	EditionNonce *float32 `json:"edition_nonce,omitempty"`
	TokenStandard *float32 `json:"token_standard,omitempty"`
	Collection *NFTCollection `json:"collection,omitempty"`
	Uses *float32 `json:"uses,omitempty"`
}

// NewNFT instantiates a new NFT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFT() *NFT {
	this := NFT{}
	return &this
}

// NewNFTWithDefaults instantiates a new NFT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTWithDefaults() *NFT {
	this := NFT{}
	return &this
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *NFT) GetConfirmed() bool {
	if o == nil || o.Confirmed == nil {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetConfirmedOk() (*bool, bool) {
	if o == nil || o.Confirmed == nil {
		return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *NFT) HasConfirmed() bool {
	if o != nil && o.Confirmed != nil {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *NFT) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NFT) GetData() NFTData {
	if o == nil || o.Data == nil {
		var ret NFTData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetDataOk() (*NFTData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NFT) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given NFTData and assigns it to the Data field.
func (o *NFT) SetData(v NFTData) {
	o.Data = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *NFT) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *NFT) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *NFT) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetMint returns the Mint field value if set, zero value otherwise.
func (o *NFT) GetMint() string {
	if o == nil || o.Mint == nil {
		var ret string
		return ret
	}
	return *o.Mint
}

// GetMintOk returns a tuple with the Mint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetMintOk() (*string, bool) {
	if o == nil || o.Mint == nil {
		return nil, false
	}
	return o.Mint, true
}

// HasMint returns a boolean if a field has been set.
func (o *NFT) HasMint() bool {
	if o != nil && o.Mint != nil {
		return true
	}

	return false
}

// SetMint gets a reference to the given string and assigns it to the Mint field.
func (o *NFT) SetMint(v string) {
	o.Mint = &v
}

// GetPrimarySaleHappened returns the PrimarySaleHappened field value if set, zero value otherwise.
func (o *NFT) GetPrimarySaleHappened() bool {
	if o == nil || o.PrimarySaleHappened == nil {
		var ret bool
		return ret
	}
	return *o.PrimarySaleHappened
}

// GetPrimarySaleHappenedOk returns a tuple with the PrimarySaleHappened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetPrimarySaleHappenedOk() (*bool, bool) {
	if o == nil || o.PrimarySaleHappened == nil {
		return nil, false
	}
	return o.PrimarySaleHappened, true
}

// HasPrimarySaleHappened returns a boolean if a field has been set.
func (o *NFT) HasPrimarySaleHappened() bool {
	if o != nil && o.PrimarySaleHappened != nil {
		return true
	}

	return false
}

// SetPrimarySaleHappened gets a reference to the given bool and assigns it to the PrimarySaleHappened field.
func (o *NFT) SetPrimarySaleHappened(v bool) {
	o.PrimarySaleHappened = &v
}

// GetUpdateAuthority returns the UpdateAuthority field value if set, zero value otherwise.
func (o *NFT) GetUpdateAuthority() string {
	if o == nil || o.UpdateAuthority == nil {
		var ret string
		return ret
	}
	return *o.UpdateAuthority
}

// GetUpdateAuthorityOk returns a tuple with the UpdateAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetUpdateAuthorityOk() (*string, bool) {
	if o == nil || o.UpdateAuthority == nil {
		return nil, false
	}
	return o.UpdateAuthority, true
}

// HasUpdateAuthority returns a boolean if a field has been set.
func (o *NFT) HasUpdateAuthority() bool {
	if o != nil && o.UpdateAuthority != nil {
		return true
	}

	return false
}

// SetUpdateAuthority gets a reference to the given string and assigns it to the UpdateAuthority field.
func (o *NFT) SetUpdateAuthority(v string) {
	o.UpdateAuthority = &v
}

// GetSellerFeeBasisPoints returns the SellerFeeBasisPoints field value if set, zero value otherwise.
func (o *NFT) GetSellerFeeBasisPoints() float32 {
	if o == nil || o.SellerFeeBasisPoints == nil {
		var ret float32
		return ret
	}
	return *o.SellerFeeBasisPoints
}

// GetSellerFeeBasisPointsOk returns a tuple with the SellerFeeBasisPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetSellerFeeBasisPointsOk() (*float32, bool) {
	if o == nil || o.SellerFeeBasisPoints == nil {
		return nil, false
	}
	return o.SellerFeeBasisPoints, true
}

// HasSellerFeeBasisPoints returns a boolean if a field has been set.
func (o *NFT) HasSellerFeeBasisPoints() bool {
	if o != nil && o.SellerFeeBasisPoints != nil {
		return true
	}

	return false
}

// SetSellerFeeBasisPoints gets a reference to the given float32 and assigns it to the SellerFeeBasisPoints field.
func (o *NFT) SetSellerFeeBasisPoints(v float32) {
	o.SellerFeeBasisPoints = &v
}

// GetMintSecretRecoveryPhrase returns the MintSecretRecoveryPhrase field value if set, zero value otherwise.
func (o *NFT) GetMintSecretRecoveryPhrase() string {
	if o == nil || o.MintSecretRecoveryPhrase == nil {
		var ret string
		return ret
	}
	return *o.MintSecretRecoveryPhrase
}

// GetMintSecretRecoveryPhraseOk returns a tuple with the MintSecretRecoveryPhrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetMintSecretRecoveryPhraseOk() (*string, bool) {
	if o == nil || o.MintSecretRecoveryPhrase == nil {
		return nil, false
	}
	return o.MintSecretRecoveryPhrase, true
}

// HasMintSecretRecoveryPhrase returns a boolean if a field has been set.
func (o *NFT) HasMintSecretRecoveryPhrase() bool {
	if o != nil && o.MintSecretRecoveryPhrase != nil {
		return true
	}

	return false
}

// SetMintSecretRecoveryPhrase gets a reference to the given string and assigns it to the MintSecretRecoveryPhrase field.
func (o *NFT) SetMintSecretRecoveryPhrase(v string) {
	o.MintSecretRecoveryPhrase = &v
}

// GetExplorerUrl returns the ExplorerUrl field value if set, zero value otherwise.
func (o *NFT) GetExplorerUrl() string {
	if o == nil || o.ExplorerUrl == nil {
		var ret string
		return ret
	}
	return *o.ExplorerUrl
}

// GetExplorerUrlOk returns a tuple with the ExplorerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetExplorerUrlOk() (*string, bool) {
	if o == nil || o.ExplorerUrl == nil {
		return nil, false
	}
	return o.ExplorerUrl, true
}

// HasExplorerUrl returns a boolean if a field has been set.
func (o *NFT) HasExplorerUrl() bool {
	if o != nil && o.ExplorerUrl != nil {
		return true
	}

	return false
}

// SetExplorerUrl gets a reference to the given string and assigns it to the ExplorerUrl field.
func (o *NFT) SetExplorerUrl(v string) {
	o.ExplorerUrl = &v
}

// GetMetadataAccount returns the MetadataAccount field value if set, zero value otherwise.
func (o *NFT) GetMetadataAccount() string {
	if o == nil || o.MetadataAccount == nil {
		var ret string
		return ret
	}
	return *o.MetadataAccount
}

// GetMetadataAccountOk returns a tuple with the MetadataAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetMetadataAccountOk() (*string, bool) {
	if o == nil || o.MetadataAccount == nil {
		return nil, false
	}
	return o.MetadataAccount, true
}

// HasMetadataAccount returns a boolean if a field has been set.
func (o *NFT) HasMetadataAccount() bool {
	if o != nil && o.MetadataAccount != nil {
		return true
	}

	return false
}

// SetMetadataAccount gets a reference to the given string and assigns it to the MetadataAccount field.
func (o *NFT) SetMetadataAccount(v string) {
	o.MetadataAccount = &v
}

// GetEditionNonce returns the EditionNonce field value if set, zero value otherwise.
func (o *NFT) GetEditionNonce() float32 {
	if o == nil || o.EditionNonce == nil {
		var ret float32
		return ret
	}
	return *o.EditionNonce
}

// GetEditionNonceOk returns a tuple with the EditionNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetEditionNonceOk() (*float32, bool) {
	if o == nil || o.EditionNonce == nil {
		return nil, false
	}
	return o.EditionNonce, true
}

// HasEditionNonce returns a boolean if a field has been set.
func (o *NFT) HasEditionNonce() bool {
	if o != nil && o.EditionNonce != nil {
		return true
	}

	return false
}

// SetEditionNonce gets a reference to the given float32 and assigns it to the EditionNonce field.
func (o *NFT) SetEditionNonce(v float32) {
	o.EditionNonce = &v
}

// GetTokenStandard returns the TokenStandard field value if set, zero value otherwise.
func (o *NFT) GetTokenStandard() float32 {
	if o == nil || o.TokenStandard == nil {
		var ret float32
		return ret
	}
	return *o.TokenStandard
}

// GetTokenStandardOk returns a tuple with the TokenStandard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetTokenStandardOk() (*float32, bool) {
	if o == nil || o.TokenStandard == nil {
		return nil, false
	}
	return o.TokenStandard, true
}

// HasTokenStandard returns a boolean if a field has been set.
func (o *NFT) HasTokenStandard() bool {
	if o != nil && o.TokenStandard != nil {
		return true
	}

	return false
}

// SetTokenStandard gets a reference to the given float32 and assigns it to the TokenStandard field.
func (o *NFT) SetTokenStandard(v float32) {
	o.TokenStandard = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *NFT) GetCollection() NFTCollection {
	if o == nil || o.Collection == nil {
		var ret NFTCollection
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetCollectionOk() (*NFTCollection, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *NFT) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given NFTCollection and assigns it to the Collection field.
func (o *NFT) SetCollection(v NFTCollection) {
	o.Collection = &v
}

// GetUses returns the Uses field value if set, zero value otherwise.
func (o *NFT) GetUses() float32 {
	if o == nil || o.Uses == nil {
		var ret float32
		return ret
	}
	return *o.Uses
}

// GetUsesOk returns a tuple with the Uses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetUsesOk() (*float32, bool) {
	if o == nil || o.Uses == nil {
		return nil, false
	}
	return o.Uses, true
}

// HasUses returns a boolean if a field has been set.
func (o *NFT) HasUses() bool {
	if o != nil && o.Uses != nil {
		return true
	}

	return false
}

// SetUses gets a reference to the given float32 and assigns it to the Uses field.
func (o *NFT) SetUses(v float32) {
	o.Uses = &v
}

func (o NFT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Confirmed != nil {
		toSerialize["confirmed"] = o.Confirmed
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.IsMutable != nil {
		toSerialize["is_mutable"] = o.IsMutable
	}
	if o.Mint != nil {
		toSerialize["mint"] = o.Mint
	}
	if o.PrimarySaleHappened != nil {
		toSerialize["primary_sale_happened"] = o.PrimarySaleHappened
	}
	if o.UpdateAuthority != nil {
		toSerialize["update_authority"] = o.UpdateAuthority
	}
	if o.SellerFeeBasisPoints != nil {
		toSerialize["seller_fee_basis_points"] = o.SellerFeeBasisPoints
	}
	if o.MintSecretRecoveryPhrase != nil {
		toSerialize["mint_secret_recovery_phrase"] = o.MintSecretRecoveryPhrase
	}
	if o.ExplorerUrl != nil {
		toSerialize["explorer_url"] = o.ExplorerUrl
	}
	if o.MetadataAccount != nil {
		toSerialize["metadata_account"] = o.MetadataAccount
	}
	if o.EditionNonce != nil {
		toSerialize["edition_nonce"] = o.EditionNonce
	}
	if o.TokenStandard != nil {
		toSerialize["token_standard"] = o.TokenStandard
	}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	if o.Uses != nil {
		toSerialize["uses"] = o.Uses
	}
	return json.Marshal(toSerialize)
}

type NullableNFT struct {
	value *NFT
	isSet bool
}

func (v NullableNFT) Get() *NFT {
	return v.value
}

func (v *NullableNFT) Set(val *NFT) {
	v.value = val
	v.isSet = true
}

func (v NullableNFT) IsSet() bool {
	return v.isSet
}

func (v *NullableNFT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFT(val *NFT) *NullableNFT {
	return &NullableNFT{value: val, isSet: true}
}

func (v NullableNFT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


