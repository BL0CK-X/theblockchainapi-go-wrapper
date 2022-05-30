# GeneralTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientBlockchainIdentifier** | **string** | The blockchain identifier of the recipient to whom you want to send a token or NFT.  On Ethereum, this is the hex public address of the recipient (e.g., &#x60;0x150865ca659204a9a0cd0da00296c6b5db441172&#x60;)  On Solana, this is the public key of the recipient (e.g., &#x60;EW3nXn7X4NTWFKWaJgxKrFNoTSkop1cBUVHA21zrfF6u&#x60;).  | 
**Wallet** | Pointer to [**GeneralWallet**](GeneralWallet.md) |  | [optional] 
**TokenBlockchainIdentifier** | Pointer to **string** | The &#x60;token_blockchain_identifier&#x60; identifies the token you wish to transfer.  - If you&#39;re transferring a native blockchain currency (e.g., SOL, ETH, BNB), then simply do not supply this value. - If you&#39;re transfering an NFT, then supply the token address of the NFT. On Solana, this is the &#x60;mint_address&#x60; or &#x60;mint&#x60; (the address of the mint). - If you&#39;re transfering a token, supply the token address. For Solana, you can find on this on the Solana Explorer (e.g., see &#x60;SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt&#x60; for &lt;a href&#x3D;\&quot;https://explorer.solana.com/address/SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Serum Token&lt;/a&gt;) for the &#x60;token_address&#x60;.  Examples: - Ethereum: &#x60;0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&#x60; - Solana: &#x60;CK1LHEANTu7RFqN3XMzo2AnZhyus2W1vue1njrxLEM1d&#x60; | [optional] 
**Network** | Pointer to **string** | The network of the blockchain you selected  - Solana: &#x60;devnet&#x60;, &#x60;mainnet-beta&#x60; - Ethereum: &#x60;ropsten&#x60;, &#x60;mainnet&#x60;  Defaults when not provided (not applicable to path parameters): - Solana: &#x60;devnet&#x60; - Ethereum: &#x60;ropsten&#x60; | [optional] 
**Amount** | Pointer to **string** | This value must be a string. What you provide here depends on if you are sending an NFT, an SPL token, or SOL.  - Native Currency (e.g., SOL, ETH, BNB): Supply this value denominated in the native currency (e.g., in SOL (but not in Lamports), or in ETH (but not in Wei)) in a string format. This does not need to be an integer. For example, if you want to send 0.0005 SOL, then &#x60;amount &#x3D; \&quot;0.0005\&quot;&#x60;. If you want to send 0.0005 ETH, then &#x60;amount &#x3D; \&quot;0.0005\&quot;&#x60;. - NFT: This must be &#x60;1&#x60;. - Token: This must be an integer in string format. To convert from what you see on a wallet UI (e.g., 1 ATLAS, 1 USDC) to an integer value, you have to multiply that value by 10^&lt;i&gt;x&lt;/i&gt; where &lt;i&gt;x&lt;/i&gt; is the number of decimals. For example, to transfer 0.2 USDC, if USDC uses 6 decimals, then the amount is 0.2 * 10^6 &#x3D; 200000.    - For Solana, you can get the number of decimals for a given SPL token &lt;a href&#x3D;\&quot;#operation/solanaGetSPLToken\&quot;&gt;here&lt;/a&gt;. We are working on analogues of this endpoint for other blockchains. | [optional] [default to "1"]
**ReturnCompiledTransaction** | Pointer to **bool** | - If &#x60;false&#x60;, we sign and submit the transaction (&#x60;wallet&#x60; is required in this case; do not provide a value for &#x60;sender_blockchain_identifier&#x60;).  - If &#x60;true&#x60;, we compile the transaction (either &#x60;wallet&#x60; or &#x60;sender_blockchain_identifier&#x60; is required in this case; do not provide both).    | [optional] [default to false]
**SenderBlockchainIdentifier** | Pointer to **string** | To understand the purpose of &#x60;sender_blockchain_identifier&#x60; first note the following: There are two ways you can complete a transfer transaction:  - (1) we complete it for you with your &#x60;wallet&#x60; information or  - (2) we return the raw instruction data that you can sign and submit yourself (no private keys required).  When you provide your &#x60;wallet&#x60; authentication, we are able to determine your wallet&#39;s blockchain identifier (public key or address) which is the sender public key of the transaction.  When you are not providing your &#x60;wallet&#x60; as authentication, we still need the &#x60;sender_blockchain_identifier&#x60; to be able to return the compiled transaction. Thus, you provide &#x60;sender_blockchain_identifier&#x60; if and only if you are not providing &#x60;wallet&#x60; authentication information **and** you are returning a compiled transaction.  You will receive an error if you provide both &#x60;wallet&#x60; and &#x60;sender_blockchain_identifier&#x60;. You will receive an error if you provide neither &#x60;wallet&#x60; nor &#x60;sender_blockchain_identifier&#x60;. | [optional] [default to "null"]
**FeePayerWallet** | Pointer to [**GeneralFeePayerWallet**](GeneralFeePayerWallet.md) |  | [optional] 

## Methods

### NewGeneralTransferRequest

`func NewGeneralTransferRequest(recipientBlockchainIdentifier string, ) *GeneralTransferRequest`

NewGeneralTransferRequest instantiates a new GeneralTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralTransferRequestWithDefaults

`func NewGeneralTransferRequestWithDefaults() *GeneralTransferRequest`

NewGeneralTransferRequestWithDefaults instantiates a new GeneralTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientBlockchainIdentifier

`func (o *GeneralTransferRequest) GetRecipientBlockchainIdentifier() string`

GetRecipientBlockchainIdentifier returns the RecipientBlockchainIdentifier field if non-nil, zero value otherwise.

### GetRecipientBlockchainIdentifierOk

`func (o *GeneralTransferRequest) GetRecipientBlockchainIdentifierOk() (*string, bool)`

GetRecipientBlockchainIdentifierOk returns a tuple with the RecipientBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientBlockchainIdentifier

`func (o *GeneralTransferRequest) SetRecipientBlockchainIdentifier(v string)`

SetRecipientBlockchainIdentifier sets RecipientBlockchainIdentifier field to given value.


### GetWallet

`func (o *GeneralTransferRequest) GetWallet() GeneralWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *GeneralTransferRequest) GetWalletOk() (*GeneralWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *GeneralTransferRequest) SetWallet(v GeneralWallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *GeneralTransferRequest) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetTokenBlockchainIdentifier

`func (o *GeneralTransferRequest) GetTokenBlockchainIdentifier() string`

GetTokenBlockchainIdentifier returns the TokenBlockchainIdentifier field if non-nil, zero value otherwise.

### GetTokenBlockchainIdentifierOk

`func (o *GeneralTransferRequest) GetTokenBlockchainIdentifierOk() (*string, bool)`

GetTokenBlockchainIdentifierOk returns a tuple with the TokenBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBlockchainIdentifier

`func (o *GeneralTransferRequest) SetTokenBlockchainIdentifier(v string)`

SetTokenBlockchainIdentifier sets TokenBlockchainIdentifier field to given value.

### HasTokenBlockchainIdentifier

`func (o *GeneralTransferRequest) HasTokenBlockchainIdentifier() bool`

HasTokenBlockchainIdentifier returns a boolean if a field has been set.

### GetNetwork

`func (o *GeneralTransferRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GeneralTransferRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GeneralTransferRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GeneralTransferRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetAmount

`func (o *GeneralTransferRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GeneralTransferRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GeneralTransferRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GeneralTransferRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetReturnCompiledTransaction

`func (o *GeneralTransferRequest) GetReturnCompiledTransaction() bool`

GetReturnCompiledTransaction returns the ReturnCompiledTransaction field if non-nil, zero value otherwise.

### GetReturnCompiledTransactionOk

`func (o *GeneralTransferRequest) GetReturnCompiledTransactionOk() (*bool, bool)`

GetReturnCompiledTransactionOk returns a tuple with the ReturnCompiledTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCompiledTransaction

`func (o *GeneralTransferRequest) SetReturnCompiledTransaction(v bool)`

SetReturnCompiledTransaction sets ReturnCompiledTransaction field to given value.

### HasReturnCompiledTransaction

`func (o *GeneralTransferRequest) HasReturnCompiledTransaction() bool`

HasReturnCompiledTransaction returns a boolean if a field has been set.

### GetSenderBlockchainIdentifier

`func (o *GeneralTransferRequest) GetSenderBlockchainIdentifier() string`

GetSenderBlockchainIdentifier returns the SenderBlockchainIdentifier field if non-nil, zero value otherwise.

### GetSenderBlockchainIdentifierOk

`func (o *GeneralTransferRequest) GetSenderBlockchainIdentifierOk() (*string, bool)`

GetSenderBlockchainIdentifierOk returns a tuple with the SenderBlockchainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderBlockchainIdentifier

`func (o *GeneralTransferRequest) SetSenderBlockchainIdentifier(v string)`

SetSenderBlockchainIdentifier sets SenderBlockchainIdentifier field to given value.

### HasSenderBlockchainIdentifier

`func (o *GeneralTransferRequest) HasSenderBlockchainIdentifier() bool`

HasSenderBlockchainIdentifier returns a boolean if a field has been set.

### GetFeePayerWallet

`func (o *GeneralTransferRequest) GetFeePayerWallet() GeneralFeePayerWallet`

GetFeePayerWallet returns the FeePayerWallet field if non-nil, zero value otherwise.

### GetFeePayerWalletOk

`func (o *GeneralTransferRequest) GetFeePayerWalletOk() (*GeneralFeePayerWallet, bool)`

GetFeePayerWalletOk returns a tuple with the FeePayerWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePayerWallet

`func (o *GeneralTransferRequest) SetFeePayerWallet(v GeneralFeePayerWallet)`

SetFeePayerWallet sets FeePayerWallet field to given value.

### HasFeePayerWallet

`func (o *GeneralTransferRequest) HasFeePayerWallet() bool`

HasFeePayerWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


