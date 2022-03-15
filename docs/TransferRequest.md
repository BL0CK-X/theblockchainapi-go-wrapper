# TransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientAddress** | **string** | The public key address of the recipient to whom you want to send a token or NFT | 
**Wallet** | Pointer to [**Wallet**](Wallet.md) |  | [optional] 
**TokenAddress** | Pointer to **string** | If you&#39;re transfering an NFT, supply the &#x60;mint&#x60; (the address of the mint) for the &#x60;token_address&#x60;. If you&#39;re transfering a token, supply the token address found on the explorer (e.g., see &#x60;SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt&#x60; for &lt;a href&#x3D;\&quot;https://explorer.solana.com/address/SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Serum Token&lt;/a&gt;) for the &#x60;token_address&#x60;. If you&#39;re transferring SOL, do not supply a value for &#x60;token_address&#x60;.  | [optional] 
**Network** | Pointer to **string** |  | [optional] [default to "devnet"]
**Amount** | Pointer to **string** | This value must be a string. What you provide here depends on if you are sending an NFT, an SPL token, or SOL.  - NFT: This must be &#39;1&#39;. - SPL Token: This must be an integer in string format. To convert from what you see on a wallet UI (e.g., 1 ATLAS, 1 USDC) to an integer value, you have to multiply that value by 10^&lt;i&gt;x&lt;/i&gt; where &lt;i&gt;x&lt;/i&gt; is the number of decimals. For example, to transfer 0.2 USDC, if USDC uses 6 decimals, then the amount is 0.2 * 10^6 &#x3D; 200000. You can get the number of decimals for a given SPL token &lt;a href&#x3D;\&quot;#operation/solanaGetSPLToken\&quot;&gt;here&lt;/a&gt;. - SOL: Supply this value denominated in SOL in a string format. This does not need to be an integer. For example, if you want to send 0.0005 SOL, then amount &#x3D; \&quot;0.0005\&quot;. | [optional] [default to "1"]
**ReturnCompiledTransaction** | Pointer to **bool** | If &#x60;false&#x60;, we sign and submit the transaction (&#x60;wallet&#x60; is required in this case; do not provide a value for &#x60;sender_public_key&#x60;).  If &#x60;true&#x60;, we compile the transaction (either &#x60;wallet&#x60; or &#x60;sender_public_key&#x60; is required in this case; do not provide both).  | [optional] [default to false]
**SenderPublicKey** | Pointer to **string** | To understand the purpose of &#x60;sender_public_key&#x60; first note the following: There are two ways you can complete a transfer transaction: (1) we complete it for you with your &#x60;wallet&#x60; information or (2) we return the raw instruction data that you can sign and submit yourself (no private keys required). When you provide your &#x60;wallet&#x60; authentication, we are able to determine your wallet&#39;s public key, which is the sender public key of the transaction. When you are not providing your &#x60;wallet&#x60; as authentication, we still need the &#x60;sender_public_key&#x60; to be able to return the compiled transaction. Thus, you provide &#x60;sender_public_key&#x60; only if you are not providing &#x60;wallet&#x60; authentication information and you are returning a compiled transaction. You will receive an error if you provide both &#x60;wallet&#x60; and &#x60;sender_public_key&#x60;. You will receive an error if you provide neither &#x60;wallet&#x60; nor &#x60;sender_public_key&#x60;.  | [optional] [default to "null"]
**FeePayerWallet** | Pointer to [**FeePayerWallet**](FeePayerWallet.md) |  | [optional] 

## Methods

### NewTransferRequest

`func NewTransferRequest(recipientAddress string, ) *TransferRequest`

NewTransferRequest instantiates a new TransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestWithDefaults

`func NewTransferRequestWithDefaults() *TransferRequest`

NewTransferRequestWithDefaults instantiates a new TransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientAddress

`func (o *TransferRequest) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *TransferRequest) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *TransferRequest) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetWallet

`func (o *TransferRequest) GetWallet() Wallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *TransferRequest) GetWalletOk() (*Wallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *TransferRequest) SetWallet(v Wallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *TransferRequest) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetTokenAddress

`func (o *TransferRequest) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TransferRequest) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TransferRequest) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TransferRequest) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetNetwork

`func (o *TransferRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransferRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransferRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *TransferRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetAmount

`func (o *TransferRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransferRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetReturnCompiledTransaction

`func (o *TransferRequest) GetReturnCompiledTransaction() bool`

GetReturnCompiledTransaction returns the ReturnCompiledTransaction field if non-nil, zero value otherwise.

### GetReturnCompiledTransactionOk

`func (o *TransferRequest) GetReturnCompiledTransactionOk() (*bool, bool)`

GetReturnCompiledTransactionOk returns a tuple with the ReturnCompiledTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCompiledTransaction

`func (o *TransferRequest) SetReturnCompiledTransaction(v bool)`

SetReturnCompiledTransaction sets ReturnCompiledTransaction field to given value.

### HasReturnCompiledTransaction

`func (o *TransferRequest) HasReturnCompiledTransaction() bool`

HasReturnCompiledTransaction returns a boolean if a field has been set.

### GetSenderPublicKey

`func (o *TransferRequest) GetSenderPublicKey() string`

GetSenderPublicKey returns the SenderPublicKey field if non-nil, zero value otherwise.

### GetSenderPublicKeyOk

`func (o *TransferRequest) GetSenderPublicKeyOk() (*string, bool)`

GetSenderPublicKeyOk returns a tuple with the SenderPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPublicKey

`func (o *TransferRequest) SetSenderPublicKey(v string)`

SetSenderPublicKey sets SenderPublicKey field to given value.

### HasSenderPublicKey

`func (o *TransferRequest) HasSenderPublicKey() bool`

HasSenderPublicKey returns a boolean if a field has been set.

### GetFeePayerWallet

`func (o *TransferRequest) GetFeePayerWallet() FeePayerWallet`

GetFeePayerWallet returns the FeePayerWallet field if non-nil, zero value otherwise.

### GetFeePayerWalletOk

`func (o *TransferRequest) GetFeePayerWalletOk() (*FeePayerWallet, bool)`

GetFeePayerWalletOk returns a tuple with the FeePayerWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePayerWallet

`func (o *TransferRequest) SetFeePayerWallet(v FeePayerWallet)`

SetFeePayerWallet sets FeePayerWallet field to given value.

### HasFeePayerWallet

`func (o *TransferRequest) HasFeePayerWallet() bool`

HasFeePayerWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


