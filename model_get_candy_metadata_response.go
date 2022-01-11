/*
The Blockchain API

# About  Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly.   You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>  # How to Use the API  To use the API, you simply need to create an API key pair.  Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.blockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.  # Feature Requests  Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@blockchainapi.com).  # Contact  <figure>     <img          width=\"40px\"         height=\"40px\"          src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"     />     <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption> </figure> <a href=\"mailto:info@blockchainapi.com\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"         />         <figcaption style=\"textAlign:center;\">Email us: info@theblockchainapi.com</figcaption>     </figure> </a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"         />         <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>     </figure> </a> <a href=\"https://twitter.com/_BlockX_\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"         />          <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>     </figure> </a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"         />         <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>     </figure> </a>  # Security  Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet.   <b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b>   How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.  We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.  Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.  Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.  # Pricing  <b>Each user receives 500 free credits.</b>  You can learn more about our pricing <a href=\"https://dashboard.blockchainapi.com/billing\" target=\"_blank\">here</a>.   We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.  If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.  # SDKs / API Wrappers  We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.  We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.  `pip install theblockchainapi`  We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-javascript-wrapper\" target=\"_blank\">JavaScript Wrapper</a>.  `npm install theblockchainapi`  We also have auto-generated wrappers for the following languages: - <a href=\"https://github.com/BL0CK-X/theblockchainapi-go-wrapper\" target = \"_blank\">Go</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-java-wrapper\" target = \"_blank\">Java</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-kotlin-wrapper\" target = \"_blank\">Kotlin</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-php-wrapper\" target = \"_blank\">PHP</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-swift-wrapper\" target = \"_blank\">Swift5</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-typescript-wrapper\" target = \"_blank\">TypeScript</a>  If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.  If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

API version: null
Contact: info@blockchainapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetCandyMetadataResponse struct for GetCandyMetadataResponse
type GetCandyMetadataResponse struct {
	// The update authority of the candy machine
	Authority *string `json:"authority,omitempty"`
	Bump *float32 `json:"bump,omitempty"`
	// The ID of the candy machine 
	CandyMachineId *string `json:"candy_machine_id,omitempty"`
	// The configuration public key address of the candy machine 
	ConfigAddress *string `json:"config_address,omitempty"`
	Creators *[]GetCandyMetadataResponseCreators `json:"creators,omitempty"`
	// The unix timestamp of the start date of the candy machine 
	GoLiveDate *float32 `json:"go_live_date,omitempty"`
	IsMutable *bool `json:"is_mutable,omitempty"`
	// The number of NFTs available for mint from the candy machine 
	ItemsAvailable *float32 `json:"items_available,omitempty"`
	// The number of NFTs minted already from the candy machine 
	ItemsRedeemed *float32 `json:"items_redeemed,omitempty"`
	MaxNumberOfLines *float32 `json:"max_number_of_lines,omitempty"`
	MaxSupply *float32 `json:"max_supply,omitempty"`
	// The price in Lamports for minting an NFT from the candy machine. 1e9 Lamport  = 1 SOL 
	Price *float32 `json:"price,omitempty"`
	RetainAuthority *bool `json:"retain_authority,omitempty"`
	// The royalty the creators receive on each sale after the primary sale (the initial minting) (denominated in basis points (e.g., 75 basis points = 0.75%)) 
	SellerFeeBasisPoints *float32 `json:"seller_fee_basis_points,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TokenMint *string `json:"token_mint,omitempty"`
	// The uuid of the candy machine
	Uuid *string `json:"uuid,omitempty"`
	// The public key address of the wallet that recevies the proceeds from NFT mints 
	Wallet *string `json:"wallet,omitempty"`
}

// NewGetCandyMetadataResponse instantiates a new GetCandyMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCandyMetadataResponse() *GetCandyMetadataResponse {
	this := GetCandyMetadataResponse{}
	return &this
}

// NewGetCandyMetadataResponseWithDefaults instantiates a new GetCandyMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCandyMetadataResponseWithDefaults() *GetCandyMetadataResponse {
	this := GetCandyMetadataResponse{}
	return &this
}

// GetAuthority returns the Authority field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetAuthority() string {
	if o == nil || o.Authority == nil {
		var ret string
		return ret
	}
	return *o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetAuthorityOk() (*string, bool) {
	if o == nil || o.Authority == nil {
		return nil, false
	}
	return o.Authority, true
}

// HasAuthority returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasAuthority() bool {
	if o != nil && o.Authority != nil {
		return true
	}

	return false
}

// SetAuthority gets a reference to the given string and assigns it to the Authority field.
func (o *GetCandyMetadataResponse) SetAuthority(v string) {
	o.Authority = &v
}

// GetBump returns the Bump field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetBump() float32 {
	if o == nil || o.Bump == nil {
		var ret float32
		return ret
	}
	return *o.Bump
}

// GetBumpOk returns a tuple with the Bump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetBumpOk() (*float32, bool) {
	if o == nil || o.Bump == nil {
		return nil, false
	}
	return o.Bump, true
}

// HasBump returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasBump() bool {
	if o != nil && o.Bump != nil {
		return true
	}

	return false
}

// SetBump gets a reference to the given float32 and assigns it to the Bump field.
func (o *GetCandyMetadataResponse) SetBump(v float32) {
	o.Bump = &v
}

// GetCandyMachineId returns the CandyMachineId field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetCandyMachineId() string {
	if o == nil || o.CandyMachineId == nil {
		var ret string
		return ret
	}
	return *o.CandyMachineId
}

// GetCandyMachineIdOk returns a tuple with the CandyMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetCandyMachineIdOk() (*string, bool) {
	if o == nil || o.CandyMachineId == nil {
		return nil, false
	}
	return o.CandyMachineId, true
}

// HasCandyMachineId returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasCandyMachineId() bool {
	if o != nil && o.CandyMachineId != nil {
		return true
	}

	return false
}

// SetCandyMachineId gets a reference to the given string and assigns it to the CandyMachineId field.
func (o *GetCandyMetadataResponse) SetCandyMachineId(v string) {
	o.CandyMachineId = &v
}

// GetConfigAddress returns the ConfigAddress field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetConfigAddress() string {
	if o == nil || o.ConfigAddress == nil {
		var ret string
		return ret
	}
	return *o.ConfigAddress
}

// GetConfigAddressOk returns a tuple with the ConfigAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetConfigAddressOk() (*string, bool) {
	if o == nil || o.ConfigAddress == nil {
		return nil, false
	}
	return o.ConfigAddress, true
}

// HasConfigAddress returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasConfigAddress() bool {
	if o != nil && o.ConfigAddress != nil {
		return true
	}

	return false
}

// SetConfigAddress gets a reference to the given string and assigns it to the ConfigAddress field.
func (o *GetCandyMetadataResponse) SetConfigAddress(v string) {
	o.ConfigAddress = &v
}

// GetCreators returns the Creators field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetCreators() []GetCandyMetadataResponseCreators {
	if o == nil || o.Creators == nil {
		var ret []GetCandyMetadataResponseCreators
		return ret
	}
	return *o.Creators
}

// GetCreatorsOk returns a tuple with the Creators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetCreatorsOk() (*[]GetCandyMetadataResponseCreators, bool) {
	if o == nil || o.Creators == nil {
		return nil, false
	}
	return o.Creators, true
}

// HasCreators returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasCreators() bool {
	if o != nil && o.Creators != nil {
		return true
	}

	return false
}

// SetCreators gets a reference to the given []GetCandyMetadataResponseCreators and assigns it to the Creators field.
func (o *GetCandyMetadataResponse) SetCreators(v []GetCandyMetadataResponseCreators) {
	o.Creators = &v
}

// GetGoLiveDate returns the GoLiveDate field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetGoLiveDate() float32 {
	if o == nil || o.GoLiveDate == nil {
		var ret float32
		return ret
	}
	return *o.GoLiveDate
}

// GetGoLiveDateOk returns a tuple with the GoLiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetGoLiveDateOk() (*float32, bool) {
	if o == nil || o.GoLiveDate == nil {
		return nil, false
	}
	return o.GoLiveDate, true
}

// HasGoLiveDate returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasGoLiveDate() bool {
	if o != nil && o.GoLiveDate != nil {
		return true
	}

	return false
}

// SetGoLiveDate gets a reference to the given float32 and assigns it to the GoLiveDate field.
func (o *GetCandyMetadataResponse) SetGoLiveDate(v float32) {
	o.GoLiveDate = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *GetCandyMetadataResponse) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetItemsAvailable returns the ItemsAvailable field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetItemsAvailable() float32 {
	if o == nil || o.ItemsAvailable == nil {
		var ret float32
		return ret
	}
	return *o.ItemsAvailable
}

// GetItemsAvailableOk returns a tuple with the ItemsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetItemsAvailableOk() (*float32, bool) {
	if o == nil || o.ItemsAvailable == nil {
		return nil, false
	}
	return o.ItemsAvailable, true
}

// HasItemsAvailable returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasItemsAvailable() bool {
	if o != nil && o.ItemsAvailable != nil {
		return true
	}

	return false
}

// SetItemsAvailable gets a reference to the given float32 and assigns it to the ItemsAvailable field.
func (o *GetCandyMetadataResponse) SetItemsAvailable(v float32) {
	o.ItemsAvailable = &v
}

// GetItemsRedeemed returns the ItemsRedeemed field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetItemsRedeemed() float32 {
	if o == nil || o.ItemsRedeemed == nil {
		var ret float32
		return ret
	}
	return *o.ItemsRedeemed
}

// GetItemsRedeemedOk returns a tuple with the ItemsRedeemed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetItemsRedeemedOk() (*float32, bool) {
	if o == nil || o.ItemsRedeemed == nil {
		return nil, false
	}
	return o.ItemsRedeemed, true
}

// HasItemsRedeemed returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasItemsRedeemed() bool {
	if o != nil && o.ItemsRedeemed != nil {
		return true
	}

	return false
}

// SetItemsRedeemed gets a reference to the given float32 and assigns it to the ItemsRedeemed field.
func (o *GetCandyMetadataResponse) SetItemsRedeemed(v float32) {
	o.ItemsRedeemed = &v
}

// GetMaxNumberOfLines returns the MaxNumberOfLines field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetMaxNumberOfLines() float32 {
	if o == nil || o.MaxNumberOfLines == nil {
		var ret float32
		return ret
	}
	return *o.MaxNumberOfLines
}

// GetMaxNumberOfLinesOk returns a tuple with the MaxNumberOfLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetMaxNumberOfLinesOk() (*float32, bool) {
	if o == nil || o.MaxNumberOfLines == nil {
		return nil, false
	}
	return o.MaxNumberOfLines, true
}

// HasMaxNumberOfLines returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasMaxNumberOfLines() bool {
	if o != nil && o.MaxNumberOfLines != nil {
		return true
	}

	return false
}

// SetMaxNumberOfLines gets a reference to the given float32 and assigns it to the MaxNumberOfLines field.
func (o *GetCandyMetadataResponse) SetMaxNumberOfLines(v float32) {
	o.MaxNumberOfLines = &v
}

// GetMaxSupply returns the MaxSupply field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetMaxSupply() float32 {
	if o == nil || o.MaxSupply == nil {
		var ret float32
		return ret
	}
	return *o.MaxSupply
}

// GetMaxSupplyOk returns a tuple with the MaxSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetMaxSupplyOk() (*float32, bool) {
	if o == nil || o.MaxSupply == nil {
		return nil, false
	}
	return o.MaxSupply, true
}

// HasMaxSupply returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasMaxSupply() bool {
	if o != nil && o.MaxSupply != nil {
		return true
	}

	return false
}

// SetMaxSupply gets a reference to the given float32 and assigns it to the MaxSupply field.
func (o *GetCandyMetadataResponse) SetMaxSupply(v float32) {
	o.MaxSupply = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *GetCandyMetadataResponse) SetPrice(v float32) {
	o.Price = &v
}

// GetRetainAuthority returns the RetainAuthority field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetRetainAuthority() bool {
	if o == nil || o.RetainAuthority == nil {
		var ret bool
		return ret
	}
	return *o.RetainAuthority
}

// GetRetainAuthorityOk returns a tuple with the RetainAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetRetainAuthorityOk() (*bool, bool) {
	if o == nil || o.RetainAuthority == nil {
		return nil, false
	}
	return o.RetainAuthority, true
}

// HasRetainAuthority returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasRetainAuthority() bool {
	if o != nil && o.RetainAuthority != nil {
		return true
	}

	return false
}

// SetRetainAuthority gets a reference to the given bool and assigns it to the RetainAuthority field.
func (o *GetCandyMetadataResponse) SetRetainAuthority(v bool) {
	o.RetainAuthority = &v
}

// GetSellerFeeBasisPoints returns the SellerFeeBasisPoints field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetSellerFeeBasisPoints() float32 {
	if o == nil || o.SellerFeeBasisPoints == nil {
		var ret float32
		return ret
	}
	return *o.SellerFeeBasisPoints
}

// GetSellerFeeBasisPointsOk returns a tuple with the SellerFeeBasisPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetSellerFeeBasisPointsOk() (*float32, bool) {
	if o == nil || o.SellerFeeBasisPoints == nil {
		return nil, false
	}
	return o.SellerFeeBasisPoints, true
}

// HasSellerFeeBasisPoints returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasSellerFeeBasisPoints() bool {
	if o != nil && o.SellerFeeBasisPoints != nil {
		return true
	}

	return false
}

// SetSellerFeeBasisPoints gets a reference to the given float32 and assigns it to the SellerFeeBasisPoints field.
func (o *GetCandyMetadataResponse) SetSellerFeeBasisPoints(v float32) {
	o.SellerFeeBasisPoints = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetCandyMetadataResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTokenMint returns the TokenMint field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetTokenMint() string {
	if o == nil || o.TokenMint == nil {
		var ret string
		return ret
	}
	return *o.TokenMint
}

// GetTokenMintOk returns a tuple with the TokenMint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetTokenMintOk() (*string, bool) {
	if o == nil || o.TokenMint == nil {
		return nil, false
	}
	return o.TokenMint, true
}

// HasTokenMint returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasTokenMint() bool {
	if o != nil && o.TokenMint != nil {
		return true
	}

	return false
}

// SetTokenMint gets a reference to the given string and assigns it to the TokenMint field.
func (o *GetCandyMetadataResponse) SetTokenMint(v string) {
	o.TokenMint = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GetCandyMetadataResponse) SetUuid(v string) {
	o.Uuid = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *GetCandyMetadataResponse) GetWallet() string {
	if o == nil || o.Wallet == nil {
		var ret string
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandyMetadataResponse) GetWalletOk() (*string, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *GetCandyMetadataResponse) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given string and assigns it to the Wallet field.
func (o *GetCandyMetadataResponse) SetWallet(v string) {
	o.Wallet = &v
}

func (o GetCandyMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authority != nil {
		toSerialize["authority"] = o.Authority
	}
	if o.Bump != nil {
		toSerialize["bump"] = o.Bump
	}
	if o.CandyMachineId != nil {
		toSerialize["candy_machine_id"] = o.CandyMachineId
	}
	if o.ConfigAddress != nil {
		toSerialize["config_address"] = o.ConfigAddress
	}
	if o.Creators != nil {
		toSerialize["creators"] = o.Creators
	}
	if o.GoLiveDate != nil {
		toSerialize["go_live_date"] = o.GoLiveDate
	}
	if o.IsMutable != nil {
		toSerialize["is_mutable"] = o.IsMutable
	}
	if o.ItemsAvailable != nil {
		toSerialize["items_available"] = o.ItemsAvailable
	}
	if o.ItemsRedeemed != nil {
		toSerialize["items_redeemed"] = o.ItemsRedeemed
	}
	if o.MaxNumberOfLines != nil {
		toSerialize["max_number_of_lines"] = o.MaxNumberOfLines
	}
	if o.MaxSupply != nil {
		toSerialize["max_supply"] = o.MaxSupply
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
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
	if o.TokenMint != nil {
		toSerialize["token_mint"] = o.TokenMint
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Wallet != nil {
		toSerialize["wallet"] = o.Wallet
	}
	return json.Marshal(toSerialize)
}

type NullableGetCandyMetadataResponse struct {
	value *GetCandyMetadataResponse
	isSet bool
}

func (v NullableGetCandyMetadataResponse) Get() *GetCandyMetadataResponse {
	return v.value
}

func (v *NullableGetCandyMetadataResponse) Set(val *GetCandyMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCandyMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCandyMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCandyMetadataResponse(val *GetCandyMetadataResponse) *NullableGetCandyMetadataResponse {
	return &NullableGetCandyMetadataResponse{value: val, isSet: true}
}

func (v NullableGetCandyMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCandyMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


