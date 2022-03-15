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

// Endpoint struct for Endpoint
type Endpoint struct {
	// The ID of the project. This is auto-generated upon project creation and cannot currently be changed. 
	ProjectId string `json:"project_id"`
	// The project version under which the endpoint exists 
	Version string `json:"version"`
	// The path of the endpoint 
	Path string `json:"path"`
	// The operation ID of the endpoint. 
	OperationId string `json:"operation_id"`
	// The name of the endpoint formatted in a readable way (e.g., Get Endpoint Metadata). 
	ReadableName string `json:"readable_name"`
	// The summary of the endpoint to be displayed in the sidebar on the left of the mini-API's documentation 
	Summary *string `json:"summary,omitempty"`
	// A description of what the endpoint does. This will be shown in the mini-API documentation. 
	Description *string `json:"description,omitempty"`
	// The price of the endpoint. Valid values are integers from 1 to 100. 
	Credits float32 `json:"credits"`
	// The name of the group of endpoints that the endpoint comes from 
	GroupName *string `json:"group_name,omitempty"`
	// A list of dictionaries. Each dictionary describes one parameter for the input, including the name, type, required boolean, and description values of that parameter.
	InputSpecification []ParameterSpecification `json:"input_specification"`
	// An example of the JSON input that a user might send. Limited to one example currently. 
	InputExamples map[string]interface{} `json:"input_examples"`
	// A list of dictionaries. Each dictionary describes one parameter for the input, including the name, type, required boolean, and description values of that parameter.
	OutputSpecification []ParameterSpecification `json:"output_specification"`
	// An example of the JSON output that a user might send. Limited to one example currently. 
	OutputExamples map[string]interface{} `json:"output_examples"`
}

// NewEndpoint instantiates a new Endpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpoint(projectId string, version string, path string, operationId string, readableName string, credits float32, inputSpecification []ParameterSpecification, inputExamples map[string]interface{}, outputSpecification []ParameterSpecification, outputExamples map[string]interface{}) *Endpoint {
	this := Endpoint{}
	this.ProjectId = projectId
	this.Version = version
	this.Path = path
	this.OperationId = operationId
	this.ReadableName = readableName
	this.Credits = credits
	this.InputSpecification = inputSpecification
	this.InputExamples = inputExamples
	this.OutputSpecification = outputSpecification
	this.OutputExamples = outputExamples
	return &this
}

// NewEndpointWithDefaults instantiates a new Endpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointWithDefaults() *Endpoint {
	this := Endpoint{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *Endpoint) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Endpoint) SetProjectId(v string) {
	o.ProjectId = v
}

// GetVersion returns the Version field value
func (o *Endpoint) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Endpoint) SetVersion(v string) {
	o.Version = v
}

// GetPath returns the Path field value
func (o *Endpoint) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *Endpoint) SetPath(v string) {
	o.Path = v
}

// GetOperationId returns the OperationId field value
func (o *Endpoint) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetOperationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *Endpoint) SetOperationId(v string) {
	o.OperationId = v
}

// GetReadableName returns the ReadableName field value
func (o *Endpoint) GetReadableName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReadableName
}

// GetReadableNameOk returns a tuple with the ReadableName field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetReadableNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReadableName, true
}

// SetReadableName sets field value
func (o *Endpoint) SetReadableName(v string) {
	o.ReadableName = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Endpoint) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Endpoint) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Endpoint) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *Endpoint) SetSummary(v string) {
	o.Summary = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Endpoint) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Endpoint) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Endpoint) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Endpoint) SetDescription(v string) {
	o.Description = &v
}

// GetCredits returns the Credits field value
func (o *Endpoint) GetCredits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetCreditsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credits, true
}

// SetCredits sets field value
func (o *Endpoint) SetCredits(v float32) {
	o.Credits = v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *Endpoint) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Endpoint) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *Endpoint) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *Endpoint) SetGroupName(v string) {
	o.GroupName = &v
}

// GetInputSpecification returns the InputSpecification field value
func (o *Endpoint) GetInputSpecification() []ParameterSpecification {
	if o == nil {
		var ret []ParameterSpecification
		return ret
	}

	return o.InputSpecification
}

// GetInputSpecificationOk returns a tuple with the InputSpecification field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetInputSpecificationOk() (*[]ParameterSpecification, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InputSpecification, true
}

// SetInputSpecification sets field value
func (o *Endpoint) SetInputSpecification(v []ParameterSpecification) {
	o.InputSpecification = v
}

// GetInputExamples returns the InputExamples field value
func (o *Endpoint) GetInputExamples() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.InputExamples
}

// GetInputExamplesOk returns a tuple with the InputExamples field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetInputExamplesOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InputExamples, true
}

// SetInputExamples sets field value
func (o *Endpoint) SetInputExamples(v map[string]interface{}) {
	o.InputExamples = v
}

// GetOutputSpecification returns the OutputSpecification field value
func (o *Endpoint) GetOutputSpecification() []ParameterSpecification {
	if o == nil {
		var ret []ParameterSpecification
		return ret
	}

	return o.OutputSpecification
}

// GetOutputSpecificationOk returns a tuple with the OutputSpecification field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetOutputSpecificationOk() (*[]ParameterSpecification, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OutputSpecification, true
}

// SetOutputSpecification sets field value
func (o *Endpoint) SetOutputSpecification(v []ParameterSpecification) {
	o.OutputSpecification = v
}

// GetOutputExamples returns the OutputExamples field value
func (o *Endpoint) GetOutputExamples() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.OutputExamples
}

// GetOutputExamplesOk returns a tuple with the OutputExamples field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetOutputExamplesOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OutputExamples, true
}

// SetOutputExamples sets field value
func (o *Endpoint) SetOutputExamples(v map[string]interface{}) {
	o.OutputExamples = v
}

func (o Endpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project_id"] = o.ProjectId
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["operation_id"] = o.OperationId
	}
	if true {
		toSerialize["readable_name"] = o.ReadableName
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["credits"] = o.Credits
	}
	if o.GroupName != nil {
		toSerialize["group_name"] = o.GroupName
	}
	if true {
		toSerialize["input_specification"] = o.InputSpecification
	}
	if true {
		toSerialize["input_examples"] = o.InputExamples
	}
	if true {
		toSerialize["output_specification"] = o.OutputSpecification
	}
	if true {
		toSerialize["output_examples"] = o.OutputExamples
	}
	return json.Marshal(toSerialize)
}

type NullableEndpoint struct {
	value *Endpoint
	isSet bool
}

func (v NullableEndpoint) Get() *Endpoint {
	return v.value
}

func (v *NullableEndpoint) Set(val *Endpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpoint(val *Endpoint) *NullableEndpoint {
	return &NullableEndpoint{value: val, isSet: true}
}

func (v NullableEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


