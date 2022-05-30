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

// Project struct for Project
type Project struct {
	// The ID of the project. This is auto-generated upon project creation and cannot currently be changed. 
	ProjectId *string `json:"project_id,omitempty"`
	// The name of the mini-API. This will be shown at the top of the mini-API's documentation. 
	ProjectName *string `json:"project_name,omitempty"`
	// The description of the mini-API. This will be shown at the top of the mini-API's documentation, below the title. 
	ProjectDescription *string `json:"project_description,omitempty"`
	// The email where users of your mini-API can contact you. This will be shown at the top of the mini-API's documentation. 
	ContactEmail *string `json:"contact_email,omitempty"`
	// The version of the API that the documentation is updated for. You can set the documentation version to any valid version. To see how to format the version string, see the description for `versions`. 
	CurrentDocumentationVersion *string `json:"current_documentation_version,omitempty"`
	// The live versions of the project. An array of strings. We use Python's `version` package to see if it's a valid version and to compare versions (to see which is higher).  Read more about this Python package <a href=\"https://packaging.pypa.io/en/latest/version.html#packaging.version.parse\" target=\"_blank\">here</a>.
	Versions []string `json:"versions,omitempty"`
	// A list of groups. A section contains groups, and groups contain API endpoints.  
	Groups []Group `json:"groups,omitempty"`
	// A list of groups. A section contains groups, and groups contain API endpoints.  
	Endpoints []Endpoint `json:"endpoints,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject() *Project {
	this := Project{}
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Project) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Project) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *Project) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *Project) GetProjectName() string {
	if o == nil || o.ProjectName == nil {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetProjectNameOk() (*string, bool) {
	if o == nil || o.ProjectName == nil {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *Project) HasProjectName() bool {
	if o != nil && o.ProjectName != nil {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *Project) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetProjectDescription returns the ProjectDescription field value if set, zero value otherwise.
func (o *Project) GetProjectDescription() string {
	if o == nil || o.ProjectDescription == nil {
		var ret string
		return ret
	}
	return *o.ProjectDescription
}

// GetProjectDescriptionOk returns a tuple with the ProjectDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetProjectDescriptionOk() (*string, bool) {
	if o == nil || o.ProjectDescription == nil {
		return nil, false
	}
	return o.ProjectDescription, true
}

// HasProjectDescription returns a boolean if a field has been set.
func (o *Project) HasProjectDescription() bool {
	if o != nil && o.ProjectDescription != nil {
		return true
	}

	return false
}

// SetProjectDescription gets a reference to the given string and assigns it to the ProjectDescription field.
func (o *Project) SetProjectDescription(v string) {
	o.ProjectDescription = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *Project) GetContactEmail() string {
	if o == nil || o.ContactEmail == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetContactEmailOk() (*string, bool) {
	if o == nil || o.ContactEmail == nil {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *Project) HasContactEmail() bool {
	if o != nil && o.ContactEmail != nil {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *Project) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetCurrentDocumentationVersion returns the CurrentDocumentationVersion field value if set, zero value otherwise.
func (o *Project) GetCurrentDocumentationVersion() string {
	if o == nil || o.CurrentDocumentationVersion == nil {
		var ret string
		return ret
	}
	return *o.CurrentDocumentationVersion
}

// GetCurrentDocumentationVersionOk returns a tuple with the CurrentDocumentationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCurrentDocumentationVersionOk() (*string, bool) {
	if o == nil || o.CurrentDocumentationVersion == nil {
		return nil, false
	}
	return o.CurrentDocumentationVersion, true
}

// HasCurrentDocumentationVersion returns a boolean if a field has been set.
func (o *Project) HasCurrentDocumentationVersion() bool {
	if o != nil && o.CurrentDocumentationVersion != nil {
		return true
	}

	return false
}

// SetCurrentDocumentationVersion gets a reference to the given string and assigns it to the CurrentDocumentationVersion field.
func (o *Project) SetCurrentDocumentationVersion(v string) {
	o.CurrentDocumentationVersion = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *Project) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetVersionsOk() ([]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *Project) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *Project) SetVersions(v []string) {
	o.Versions = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *Project) GetGroups() []Group {
	if o == nil || o.Groups == nil {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetGroupsOk() ([]Group, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Project) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *Project) SetGroups(v []Group) {
	o.Groups = v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *Project) GetEndpoints() []Endpoint {
	if o == nil || o.Endpoints == nil {
		var ret []Endpoint
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetEndpointsOk() ([]Endpoint, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *Project) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []Endpoint and assigns it to the Endpoints field.
func (o *Project) SetEndpoints(v []Endpoint) {
	o.Endpoints = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.ProjectName != nil {
		toSerialize["project_name"] = o.ProjectName
	}
	if o.ProjectDescription != nil {
		toSerialize["project_description"] = o.ProjectDescription
	}
	if o.ContactEmail != nil {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if o.CurrentDocumentationVersion != nil {
		toSerialize["current_documentation_version"] = o.CurrentDocumentationVersion
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	return json.Marshal(toSerialize)
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


