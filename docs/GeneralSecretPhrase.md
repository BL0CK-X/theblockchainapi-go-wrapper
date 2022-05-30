# GeneralSecretPhrase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretRecoveryPhrase** | Pointer to **string** | The 12, 15, 18, 21, or 24 word phrase that can be used to derive many public key addresses (e.g., bottom army bless castle alter habit dish embody child flame smooth zone).  To derive a wallet identifier (e.g., public key, public address, etc.), you need a secret recovery phrase, a derivation path, and an optional passphrase.   Alternatively, you can derive a wallet identifier with a private key.  Default Lengths: - Avalanche (C): 24 - Avalanche (X, P): 24 - Binance: 12 - Ethereum: 12 - Near: 12 - Solana: 12  You can use &#x60;secret_recovery_phrase&#x60; for Chains: &#x60;avalanche&#x60;, &#x60;binance&#x60;, &#x60;ethereum&#x60;, &#x60;near&#x60;, &#x60;solana&#x60;. | [optional] 

## Methods

### NewGeneralSecretPhrase

`func NewGeneralSecretPhrase() *GeneralSecretPhrase`

NewGeneralSecretPhrase instantiates a new GeneralSecretPhrase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralSecretPhraseWithDefaults

`func NewGeneralSecretPhraseWithDefaults() *GeneralSecretPhrase`

NewGeneralSecretPhraseWithDefaults instantiates a new GeneralSecretPhrase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretRecoveryPhrase

`func (o *GeneralSecretPhrase) GetSecretRecoveryPhrase() string`

GetSecretRecoveryPhrase returns the SecretRecoveryPhrase field if non-nil, zero value otherwise.

### GetSecretRecoveryPhraseOk

`func (o *GeneralSecretPhrase) GetSecretRecoveryPhraseOk() (*string, bool)`

GetSecretRecoveryPhraseOk returns a tuple with the SecretRecoveryPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRecoveryPhrase

`func (o *GeneralSecretPhrase) SetSecretRecoveryPhrase(v string)`

SetSecretRecoveryPhrase sets SecretRecoveryPhrase field to given value.

### HasSecretRecoveryPhrase

`func (o *GeneralSecretPhrase) HasSecretRecoveryPhrase() bool`

HasSecretRecoveryPhrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


