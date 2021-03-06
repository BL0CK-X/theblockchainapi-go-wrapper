# ATARequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretRecoveryPhrase** | **string** | The twelve word phrase that can be used to derive many public key addresses. To derive a public key, you need a secret recovery phrase, a derivation path, and an optional passphrase. See our Security section &lt;a href&#x3D;\&quot;#section/Security\&quot;&gt;here&lt;/a&gt;. | 
**DerivationPath** | Pointer to **string** | Derivation paths are used to derive the public key from the secret recovery phrase. Only certain paths are accepted.  We use \&quot;m/44/501/0/0\&quot; by default, if it is not provided. This is the path that the Phantom and Sollet wallets use. If you provide the empty string \&quot;\&quot; as the value for the derivation path, then we will use the Solana CLI default value. The SolFlare recommended path is \&quot;m/44/501/0\&quot;.  You can also arbitrarily increment the default path (\&quot;m/44/501/0/0\&quot;) to generate more wallets (e.g., \&quot;m/44/501/0/1\&quot;, \&quot;m/44/501/0/2\&quot;, ...). This is how Phantom generates more wallets.  To learn more about derivation paths, check out &lt;a href&#x3D;\&quot;https://learnmeabitcoin.com/technical/derivation-paths\&quot; target&#x3D;\&quot;_blank\&quot;&gt;this tutorial&lt;/a&gt;. | [optional] [default to "m/44/501/0/0"]
**Passphrase** | Pointer to **string** | An optional string used to generate an address. This provides an additional layer of security because a hacker would need both the secret recovery phrase and the passphrase to access the output public key. By default, most wallet UI extensions do not use a passphrase. Limited to 500 characters.  | [optional] [default to ""]
**TokenAddress** | **string** | The address of the token to get the ATA for. For an NFT, this is the same as &#x60;mint&#x60; (the mint address). | 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]

## Methods

### NewATARequest

`func NewATARequest(secretRecoveryPhrase string, tokenAddress string, ) *ATARequest`

NewATARequest instantiates a new ATARequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewATARequestWithDefaults

`func NewATARequestWithDefaults() *ATARequest`

NewATARequestWithDefaults instantiates a new ATARequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretRecoveryPhrase

`func (o *ATARequest) GetSecretRecoveryPhrase() string`

GetSecretRecoveryPhrase returns the SecretRecoveryPhrase field if non-nil, zero value otherwise.

### GetSecretRecoveryPhraseOk

`func (o *ATARequest) GetSecretRecoveryPhraseOk() (*string, bool)`

GetSecretRecoveryPhraseOk returns a tuple with the SecretRecoveryPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRecoveryPhrase

`func (o *ATARequest) SetSecretRecoveryPhrase(v string)`

SetSecretRecoveryPhrase sets SecretRecoveryPhrase field to given value.


### GetDerivationPath

`func (o *ATARequest) GetDerivationPath() string`

GetDerivationPath returns the DerivationPath field if non-nil, zero value otherwise.

### GetDerivationPathOk

`func (o *ATARequest) GetDerivationPathOk() (*string, bool)`

GetDerivationPathOk returns a tuple with the DerivationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationPath

`func (o *ATARequest) SetDerivationPath(v string)`

SetDerivationPath sets DerivationPath field to given value.

### HasDerivationPath

`func (o *ATARequest) HasDerivationPath() bool`

HasDerivationPath returns a boolean if a field has been set.

### GetPassphrase

`func (o *ATARequest) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *ATARequest) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *ATARequest) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *ATARequest) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetTokenAddress

`func (o *ATARequest) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *ATARequest) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *ATARequest) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetNetwork

`func (o *ATARequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ATARequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ATARequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ATARequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


