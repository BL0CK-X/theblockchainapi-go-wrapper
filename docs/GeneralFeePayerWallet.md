# GeneralFeePayerWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretRecoveryPhrase** | **string** | The 12, 15, 18, 21, or 24 word phrase that can be used to derive many public key addresses (e.g., bottom army bless castle alter habit dish embody child flame smooth zone).  To derive a wallet identifier (e.g., public key, public address, etc.), you need a secret recovery phrase, a derivation path, and an optional passphrase.   Alternatively, you can derive a wallet identifier with a private key.  Default Lengths: - Avalanche (C): 24 - Avalanche (X, P): 24 - Binance: 12 - Ethereum: 12 - Near: 12 - Solana: 12  You can use &#x60;secret_recovery_phrase&#x60; for Chains: &#x60;avalanche&#x60;, &#x60;binance&#x60;, &#x60;ethereum&#x60;, &#x60;near&#x60;, &#x60;solana&#x60;. | 
**DerivationPath** | Pointer to **string** | Derivation paths are used to derive the wallet identifier from the secret recovery phrase. Only correctly typed paths are accepted.  Defaults: - Avalanche (C): m/44&#39;/60&#39;/0&#39;/0/0 - Avalanche (X, P): m/44&#39;/60&#39;/0&#39;/0/0 - Binance: m/44&#39;/60&#39;/0&#39;/0/0 - Ethereum: m/44&#39;/60&#39;/0&#39;/0/0 - Near: m/44/397/0 - Solana: m/44/501/0/0  *Solana Behavior*  If you provide the empty string \&quot;\&quot; as the value for the derivation path, then we will derive your public key with the same behavior as the default behavior of the Solana CLI.  By default, we use \&quot;m/44/501/0/0\&quot;. This is the path that the Phantom and Sollet wallets use.  You can also arbitrarily increment the default path (\&quot;m/44/501/0/0\&quot;) to generate more wallets (e.g., \&quot;m/44/501/0/1\&quot;, \&quot;m/44/501/0/2\&quot;, ... AND/OR \&quot;m/44/501/1/0\&quot;, \&quot;m/44/501/2/0\&quot;, ...).  Phantom&#39;s Wallet increments the first digit (e.g., \&quot;m/44/501/0/0\&quot;, \&quot;m/44/501/1/0\&quot;, \&quot;m/44/501/2/0\&quot;, ...) to generate more public key addresses.  The SolFlare recommended path is \&quot;m/44/501/0\&quot;.  To learn more about derivation paths, check out &lt;a href&#x3D;\&quot;https://learnmeabitcoin.com/technical/derivation-paths\&quot; target&#x3D;\&quot;_blank\&quot;&gt;this tutorial&lt;/a&gt;. | [optional] 
**Passphrase** | Pointer to **string** | PASSPHRASE !&#x3D; PASSWORD. This is NOT your Phantom password or any other password. It is an optional string you use when creating a wallet. This provides an additional layer of security because a hacker would need both the secret recovery phrase and the passphrase to access the output public key. By default, most wallet UI extensions do not use a passphrase. (You probably did not use a passphrase.) Limited to 500 characters.  | [optional] [default to ""]
**HexPrivateKey** | **string** | A hex private key is the typical private key used on Ethereum, for example. It looks like this: &#x60;0x200b9e5baa38b0dc7551645be11b394e9bf2b04532e4af8824bed2b3de2e0dc0&#x60;.  You can use &#x60;hex_private_key&#x60; for chains: &#x60;avalanche&#x60;, &#x60;binance_smart_chain&#x60;, &#x60;ethereum&#x60;. | 
**PrivateKey** | **map[string]interface{}** | A private key corresponds to exactly one public key address. A private key can be used to move assets out of the wallet and sign transaction with the corresponding public key.  A private key is an array of integers (e.g.,   &#x60;[185,108,153,165,57,193,166,167,58,148,133,121,92,252,242,13,233,246,35,103,185,20,27,56,111,169,12,50,50,36,83,156,173,195,143,75,135,78,204,129,217,231,58,129,69,180,185,86,119,43,200,193,94,112,31,135,68,128,207,26,85,150,68,181]&#x60;).  &lt;a href&#x3D;\&quot;https://solflare.com\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Solflare&lt;/a&gt; is a popular wallet interface on Solana that allows you to export your private key in this format.  You can use &#x60;private_key&#x60; for Chains: &#x60;avalanche&#x60;, &#x60;binance_smart_chain&#x60;, &#x60;ethereum&#x60;, &#x60;solana&#x60;, &#x60;near&#x60;. | 
**B58PrivateKey** | **string** | A private key corresponds to exactly one public key address. A private key can be used to move assets out of the wallet and sign transaction with the corresponding public key.  A base58-encoded private key is a base58-encoded version of the typical private key. It is represented as a string (e.g., &#x60;4waBTVeAVWEAczSdx36uMrR19668ACgQDs7r386vrUes3UCzvXCQ2FPSCVGb1zJrwcULgpNzgABreyQaWSpGBwfx&#x60;).  &lt;a href&#x3D;\&quot;https://phantom.app\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Phantom&lt;/a&gt; is a popular wallet interface on Solana that allows you to export your private key in this format.  You can use &#x60;b58_private_key&#x60; for Chains: &#x60;solana&#x60;, &#x60;near&#x60; | 

## Methods

### NewGeneralFeePayerWallet

`func NewGeneralFeePayerWallet(secretRecoveryPhrase string, hexPrivateKey string, privateKey map[string]interface{}, b58PrivateKey string, ) *GeneralFeePayerWallet`

NewGeneralFeePayerWallet instantiates a new GeneralFeePayerWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralFeePayerWalletWithDefaults

`func NewGeneralFeePayerWalletWithDefaults() *GeneralFeePayerWallet`

NewGeneralFeePayerWalletWithDefaults instantiates a new GeneralFeePayerWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretRecoveryPhrase

`func (o *GeneralFeePayerWallet) GetSecretRecoveryPhrase() string`

GetSecretRecoveryPhrase returns the SecretRecoveryPhrase field if non-nil, zero value otherwise.

### GetSecretRecoveryPhraseOk

`func (o *GeneralFeePayerWallet) GetSecretRecoveryPhraseOk() (*string, bool)`

GetSecretRecoveryPhraseOk returns a tuple with the SecretRecoveryPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRecoveryPhrase

`func (o *GeneralFeePayerWallet) SetSecretRecoveryPhrase(v string)`

SetSecretRecoveryPhrase sets SecretRecoveryPhrase field to given value.


### GetDerivationPath

`func (o *GeneralFeePayerWallet) GetDerivationPath() string`

GetDerivationPath returns the DerivationPath field if non-nil, zero value otherwise.

### GetDerivationPathOk

`func (o *GeneralFeePayerWallet) GetDerivationPathOk() (*string, bool)`

GetDerivationPathOk returns a tuple with the DerivationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationPath

`func (o *GeneralFeePayerWallet) SetDerivationPath(v string)`

SetDerivationPath sets DerivationPath field to given value.

### HasDerivationPath

`func (o *GeneralFeePayerWallet) HasDerivationPath() bool`

HasDerivationPath returns a boolean if a field has been set.

### GetPassphrase

`func (o *GeneralFeePayerWallet) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *GeneralFeePayerWallet) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *GeneralFeePayerWallet) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *GeneralFeePayerWallet) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetHexPrivateKey

`func (o *GeneralFeePayerWallet) GetHexPrivateKey() string`

GetHexPrivateKey returns the HexPrivateKey field if non-nil, zero value otherwise.

### GetHexPrivateKeyOk

`func (o *GeneralFeePayerWallet) GetHexPrivateKeyOk() (*string, bool)`

GetHexPrivateKeyOk returns a tuple with the HexPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPrivateKey

`func (o *GeneralFeePayerWallet) SetHexPrivateKey(v string)`

SetHexPrivateKey sets HexPrivateKey field to given value.


### GetPrivateKey

`func (o *GeneralFeePayerWallet) GetPrivateKey() map[string]interface{}`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GeneralFeePayerWallet) GetPrivateKeyOk() (*map[string]interface{}, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GeneralFeePayerWallet) SetPrivateKey(v map[string]interface{})`

SetPrivateKey sets PrivateKey field to given value.


### GetB58PrivateKey

`func (o *GeneralFeePayerWallet) GetB58PrivateKey() string`

GetB58PrivateKey returns the B58PrivateKey field if non-nil, zero value otherwise.

### GetB58PrivateKeyOk

`func (o *GeneralFeePayerWallet) GetB58PrivateKeyOk() (*string, bool)`

GetB58PrivateKeyOk returns a tuple with the B58PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB58PrivateKey

`func (o *GeneralFeePayerWallet) SetB58PrivateKey(v string)`

SetB58PrivateKey sets B58PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


