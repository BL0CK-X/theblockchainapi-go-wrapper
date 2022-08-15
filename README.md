# Go API client for openapi

# About

Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly. 

You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>

# How to Use the API

To use the API, you simply need to create an API key pair.

Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.blockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.

# Feature Requests

Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@blockchainapi.com).

# Contact

<figure>
    <img 
        width=\"40px\"
        height=\"40px\" 
        src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"
    />
    <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption>
</figure> <a href=\"mailto:info@blockchainapi.com\" target=\"_blank\">
    <figure>
        <img 
            width=\"40px\"
            height=\"40px\" 
            src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"
        />
        <figcaption style=\"textAlign:center;\">Email us: info@blockchainapi.com</figcaption>
    </figure>
</a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">
    <figure>
        <img 
            width=\"40px\"
            height=\"40px\" 
            src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"
        />
        <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>
    </figure>
</a> <a href=\"https://twitter.com/BlockchainAP1\" target=\"_blank\">
    <figure>
        <img 
            width=\"40px\"
            height=\"40px\" 
            src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"
        />
         <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>
    </figure>
</a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">
    <figure>
        <img 
            width=\"40px\"
            height=\"40px\" 
            src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"
        />
        <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>
    </figure>
</a>

# Security

Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet. 

<b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b> 

How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.

We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.

Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.

Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.

# Pricing

<b>Each user receives 50,000 free credits each month.</b> Each endpoint costs a certain amount credits. Scroll below to any endpoint (i.e., function) to see how much each endpoint costs. (Or CMD+F `Cost: 0 Credit`, for example)

You can learn more about our pricing <a href=\"https://dashboard.blockchainapi.com/billing\" target=\"_blank\">here</a>. 

We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.

If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.

# SDKs / API Wrappers

We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.

We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.

`pip install theblockchainapi`

We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-javascript-wrapper\" target=\"_blank\">JavaScript Wrapper</a>.

`npm install theblockchainapi`

We also have auto-generated wrappers for the following languages:
- <a href=\"https://github.com/BL0CK-X/theblockchainapi-go-wrapper\" target = \"_blank\">Go</a>
- <a href=\"https://github.com/BL0CK-X/theblockchainapi-java-wrapper\" target = \"_blank\">Java</a>
- <a href=\"https://github.com/BL0CK-X/theblockchainapi-kotlin-wrapper\" target = \"_blank\">Kotlin</a>
- <a href=\"https://github.com/BL0CK-X/theblockchainapi-php-wrapper\" target = \"_blank\">PHP</a>
- <a href=\"https://github.com/BL0CK-X/theblockchainapi-swift-wrapper\" target = \"_blank\">Swift5</a>
- <a href=\"https://github.com/BL0CK-X/theblockchainapi-typescript-wrapper\" target = \"_blank\">TypeScript</a>
- <a href=\"https://github.com/BL0CK-X/theblockchainapi-csharp-wrapper\" target = \"_blank\">C#</a>
- <a href=\"https://github.com/BL0CK-X/theblockchainapi-dart-wrapper\" target = \"_blank\">Dart</a>

If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: null
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.blockchainapi.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CCPaymentApi* | [**GetCCPayment**](docs/CCPaymentApi.md#getccpayment) | **Get** /checkout/v1/project/{project_id}/payment/{payment_identifier} | Get a payment 
*CCPaymentApi* | [**ListAllCCPayments**](docs/CCPaymentApi.md#listallccpayments) | **Get** /checkout/v1/payments | List all payments  
*CCPaymentApi* | [**ListCCProjectPayments**](docs/CCPaymentApi.md#listccprojectpayments) | **Get** /checkout/v1/project/{project_id}/payments | List a project&#39;s payments 
*CCProjectApi* | [**CreateCCProject**](docs/CCProjectApi.md#createccproject) | **Post** /checkout/v1/project | Create a project 
*CCProjectApi* | [**DeleteCCProject**](docs/CCProjectApi.md#deleteccproject) | **Delete** /checkout/v1/project/{project_id} | Delete a project 
*CCProjectApi* | [**GetCCProject**](docs/CCProjectApi.md#getccproject) | **Get** /checkout/v1/project/{project_id} | Get a project 
*CCProjectApi* | [**ListCCProjects**](docs/CCProjectApi.md#listccprojects) | **Get** /checkout/v1/projects | List projects 
*CCProjectApi* | [**UpdateCCProject**](docs/CCProjectApi.md#updateccproject) | **Put** /checkout/v1/project/{project_id} | Update a project 
*CCWebhookApi* | [**GetCCWebhook**](docs/CCWebhookApi.md#getccwebhook) | **Post** /webhook/v1/project/{project_id}/webhook/{webhook_identifier} | Get a webhook 
*CCWebhookApi* | [**ListCCProjectWebhooks**](docs/CCWebhookApi.md#listccprojectwebhooks) | **Post** /webhook/v1/project/{project_id}/webhooks | List project&#39;s webhooks 
*CCWebhookApi* | [**ValidateCCWebhook**](docs/CCWebhookApi.md#validateccwebhook) | **Post** /webhook/v1/project/{project_id}/webhook/{webhook_identifier}/validate | Validate a webhook 
*NameServiceApi* | [**GetBlockchainIdentifierFromName**](docs/NameServiceApi.md#getblockchainidentifierfromname) | **Post** /{blockchain}/{network}/name_service/name_to_blockchain_identifier | Get the identifier
*NameServiceApi* | [**GetNameForBlockchainIdentifier**](docs/NameServiceApi.md#getnameforblockchainidentifier) | **Post** /{blockchain}/{network}/name_service/blockchain_identifier_to_name | Get the name
*SolanaAccountApi* | [**SolanaGetAccount**](docs/SolanaAccountApi.md#solanagetaccount) | **Get** /solana/account/{network}/{public_key} | Get the details of an account on Solana
*SolanaAccountApi* | [**SolanaGetAccountIsCandyMachine**](docs/SolanaAccountApi.md#solanagetaccountiscandymachine) | **Get** /solana/account/{network}/{public_key}/is_candy_machine | Get if account is candy machine
*SolanaAccountApi* | [**SolanaGetAccountIsNFT**](docs/SolanaAccountApi.md#solanagetaccountisnft) | **Get** /solana/account/{network}/{public_key}/is_nft | Get if account is NFT
*SolanaCandyMachineApi* | [**SolanaGetAllNFTsFromCandyMachine**](docs/SolanaCandyMachineApi.md#solanagetallnftsfromcandymachine) | **Get** /solana/nft/candy_machine/{network}/{candy_machine_id}/nfts | Get CM&#39;s NFTs  
*SolanaCandyMachineApi* | [**SolanaGetCandyMachineMetadata**](docs/SolanaCandyMachineApi.md#solanagetcandymachinemetadata) | **Post** /solana/nft/candy_machine/metadata | Get a CM&#39;s metadata 
*SolanaCandyMachineApi* | [**SolanaListAllCandyMachines**](docs/SolanaCandyMachineApi.md#solanalistallcandymachines) | **Get** /solana/nft/candy_machine/list | List all CMs
*SolanaCandyMachineApi* | [**SolanaSearchCandyMachines**](docs/SolanaCandyMachineApi.md#solanasearchcandymachines) | **Post** /solana/nft/candy_machine/search | Search CMs
*SolanaNFTApi* | [**SolanaCreateNFT**](docs/SolanaNFTApi.md#solanacreatenft) | **Post** /solana/nft | Create an NFT on Solana
*SolanaNFTApi* | [**SolanaGetNFT**](docs/SolanaNFTApi.md#solanagetnft) | **Get** /solana/nft/{network}/{mint_address} | Get an NFT&#39;s metadata
*SolanaNFTApi* | [**SolanaGetNFTMintFee**](docs/SolanaNFTApi.md#solanagetnftmintfee) | **Get** /solana/nft/mint/fee | Get the NFT mint fee
*SolanaNFTApi* | [**SolanaGetNFTOwner**](docs/SolanaNFTApi.md#solanagetnftowner) | **Get** /solana/nft/{network}/{mint_address}/owner | Get owner of an NFT
*SolanaNFTApi* | [**SolanaGetNFTOwnerAdvanced**](docs/SolanaNFTApi.md#solanagetnftowneradvanced) | **Get** /solana/nft/{network}/{mint_address}/owner_advanced | Get owner of an NFT (advanced)
*SolanaNFTApi* | [**SolanaGetNFTsCandyMachineId**](docs/SolanaNFTApi.md#solanagetnftscandymachineid) | **Post** /solana/nft/candy_machine_id | Get the ID of the candy machine of an NFT 
*SolanaNFTApi* | [**SolanaSearchNFTs**](docs/SolanaNFTApi.md#solanasearchnfts) | **Post** /solana/nft/search | Search NFTs on Solana
*SolanaSPLTokenApi* | [**SolanaGetSPLToken**](docs/SolanaSPLTokenApi.md#solanagetspltoken) | **Get** /solana/spl-token/{network}/{public_key} | Get SPL token metadata
*SolanaTransactionApi* | [**SolanaGetTransaction**](docs/SolanaTransactionApi.md#solanagettransaction) | **Get** /solana/transaction/{network}/{tx_signature} | Get the details of a transaction made on Solana
*SolanaWalletApi* | [**SolanaDeriveAssociatedTokenAccountAddress**](docs/SolanaWalletApi.md#solanaderiveassociatedtokenaccountaddress) | **Get** /solana/wallet/{public_key}/associated_token_account/{mint_address} | Derive an associated token account address
*SolanaWalletApi* | [**SolanaDerivePrivateKey**](docs/SolanaWalletApi.md#solanaderiveprivatekey) | **Post** /solana/wallet/private_key | Derive private key
*SolanaWalletApi* | [**SolanaDerivePublicKey**](docs/SolanaWalletApi.md#solanaderivepublickey) | **Post** /solana/wallet/public_key | Derive public key
*SolanaWalletApi* | [**SolanaGeneratePrivateKey**](docs/SolanaWalletApi.md#solanagenerateprivatekey) | **Post** /solana/wallet/generate/private_key | Generate private key
*SolanaWalletApi* | [**SolanaGenerateSecretRecoveryPhrase**](docs/SolanaWalletApi.md#solanageneratesecretrecoveryphrase) | **Post** /solana/wallet/generate/secret_recovery_phrase | Generate secret phrase
*SolanaWalletApi* | [**SolanaGetAirdrop**](docs/SolanaWalletApi.md#solanagetairdrop) | **Post** /solana/wallet/airdrop | Get an airdrop on devnet
*SolanaWalletApi* | [**SolanaGetBalance**](docs/SolanaWalletApi.md#solanagetbalance) | **Post** /solana/wallet/balance | Get wallet&#39;s balance in SOL or any SPL
*SolanaWalletApi* | [**SolanaGetNFTsBelongingToWallet**](docs/SolanaWalletApi.md#solanagetnftsbelongingtowallet) | **Get** /solana/wallet/{network}/{public_key}/nfts | Get address&#39;s NFTs
*SolanaWalletApi* | [**SolanaGetTokensBelongingToWallet**](docs/SolanaWalletApi.md#solanagettokensbelongingtowallet) | **Get** /solana/wallet/{network}/{public_key}/tokens | Get address&#39;s tokens and respective balances
*SolanaWalletApi* | [**SolanaGetWalletTransactions**](docs/SolanaWalletApi.md#solanagetwallettransactions) | **Get** /solana/wallet/{network}/{public_key}/transactions | Get address&#39;s associated transaction signatures
*SolanaWalletApi* | [**SolanaTransfer**](docs/SolanaWalletApi.md#solanatransfer) | **Post** /solana/wallet/transfer | Transfer SOL, a token, or an NFT to another address
*TokenApi* | [**GetTokenMetadata**](docs/TokenApi.md#gettokenmetadata) | **Get** /{blockchain}/{network}/token/{token_blockchain_identifier} | Get a token&#39;s metadata
*TokenApi* | [**ListAllTokens**](docs/TokenApi.md#listalltokens) | **Get** /{blockchain}/{network}/all_tokens | List all tokens
*TransactionApi* | [**GetTransaction**](docs/TransactionApi.md#gettransaction) | **Get** /{blockchain}/transaction/{network}/{transaction_blockchain_identifier} | Get the details of a transaction made on a blockchain
*WalletApi* | [**DerivePrivateKey**](docs/WalletApi.md#deriveprivatekey) | **Post** /{blockchain}/wallet/private_key | Derive private key
*WalletApi* | [**DeriveWalletIdentifier**](docs/WalletApi.md#derivewalletidentifier) | **Post** /{blockchain}/wallet/identifier | Derive wallet identifier
*WalletApi* | [**GeneratePrivateKey**](docs/WalletApi.md#generateprivatekey) | **Post** /{blockchain}/wallet/generate/private_key | Generate private key
*WalletApi* | [**GenerateSeedPhrase**](docs/WalletApi.md#generateseedphrase) | **Post** /{blockchain}/wallet/generate/secret_recovery_phrase | Generate seed phrase
*WalletApi* | [**GetAirdrop**](docs/WalletApi.md#getairdrop) | **Post** /{blockchain}/wallet/airdrop | Get an airdrop
*WalletApi* | [**GetBalance**](docs/WalletApi.md#getbalance) | **Post** /{blockchain}/wallet/balance | Get wallet&#39;s balance of X
*WalletApi* | [**Transfer**](docs/WalletApi.md#transfer) | **Post** /{blockchain}/wallet/transfer | Transfer crypto, a token, or an NFT to another wallet


## Documentation For Models

 - [ATAResponse](docs/ATAResponse.md)
 - [Account](docs/Account.md)
 - [AccountContext](docs/AccountContext.md)
 - [AccountIsCandyMachine](docs/AccountIsCandyMachine.md)
 - [AccountIsNFT](docs/AccountIsNFT.md)
 - [AccountValue](docs/AccountValue.md)
 - [AirdropRequest](docs/AirdropRequest.md)
 - [AvalancheCChainPublicAddress](docs/AvalancheCChainPublicAddress.md)
 - [AvalancheXPChainPublicAddress](docs/AvalancheXPChainPublicAddress.md)
 - [B58PrivateKey](docs/B58PrivateKey.md)
 - [BSCPublicAddress](docs/BSCPublicAddress.md)
 - [BalanceRequest](docs/BalanceRequest.md)
 - [BalanceResponse](docs/BalanceResponse.md)
 - [CCPayment](docs/CCPayment.md)
 - [CCPaymentBlockchainPaymentDetails](docs/CCPaymentBlockchainPaymentDetails.md)
 - [CCProject](docs/CCProject.md)
 - [CCProjectCreateRequest](docs/CCProjectCreateRequest.md)
 - [CCProjectCreateRequestCustomerIdToCollect](docs/CCProjectCreateRequestCustomerIdToCollect.md)
 - [CCProjectCreateRequestPayoutMethod](docs/CCProjectCreateRequestPayoutMethod.md)
 - [CCWebhook](docs/CCWebhook.md)
 - [CCWebhookValidateRequest](docs/CCWebhookValidateRequest.md)
 - [CandyMachineSearchRequest](docs/CandyMachineSearchRequest.md)
 - [DoubleTransferResponse](docs/DoubleTransferResponse.md)
 - [EthereumPublicAddress](docs/EthereumPublicAddress.md)
 - [EthereumTransaction](docs/EthereumTransaction.md)
 - [EthereumTransactionCompiledResponse](docs/EthereumTransactionCompiledResponse.md)
 - [EthereumTransactionMadeResponse](docs/EthereumTransactionMadeResponse.md)
 - [FeePayerWallet](docs/FeePayerWallet.md)
 - [GeneralAirdropRequest](docs/GeneralAirdropRequest.md)
 - [GeneralB58PrivateKey](docs/GeneralB58PrivateKey.md)
 - [GeneralBalanceRequest](docs/GeneralBalanceRequest.md)
 - [GeneralBalanceResponse](docs/GeneralBalanceResponse.md)
 - [GeneralFeePayerWallet](docs/GeneralFeePayerWallet.md)
 - [GeneralGeneratePrivateKeyResponse](docs/GeneralGeneratePrivateKeyResponse.md)
 - [GeneralPrivateKey](docs/GeneralPrivateKey.md)
 - [GeneralSecretPhrase](docs/GeneralSecretPhrase.md)
 - [GeneralSecretRecoveryPhrase](docs/GeneralSecretRecoveryPhrase.md)
 - [GeneralTransaction](docs/GeneralTransaction.md)
 - [GeneralTransactionMadeResponse](docs/GeneralTransactionMadeResponse.md)
 - [GeneralTransactionResponse](docs/GeneralTransactionResponse.md)
 - [GeneralTransferRequest](docs/GeneralTransferRequest.md)
 - [GeneralWallet](docs/GeneralWallet.md)
 - [GeneratePrivateKey](docs/GeneratePrivateKey.md)
 - [GenerateSeedPhraseRequest](docs/GenerateSeedPhraseRequest.md)
 - [GetAllNFTsResponse](docs/GetAllNFTsResponse.md)
 - [GetAllNFTsResponseMintedNfts](docs/GetAllNFTsResponseMintedNfts.md)
 - [GetAllNFTsResponseUnmintedNfts](docs/GetAllNFTsResponseUnmintedNfts.md)
 - [GetCandyMachineIDRequest](docs/GetCandyMachineIDRequest.md)
 - [GetCandyMachineIDResponse](docs/GetCandyMachineIDResponse.md)
 - [GetCandyMetadataErrorResponse](docs/GetCandyMetadataErrorResponse.md)
 - [GetCandyMetadataRequest](docs/GetCandyMetadataRequest.md)
 - [GetCandyMetadataResponse](docs/GetCandyMetadataResponse.md)
 - [GetCandyMetadataResponseCreators](docs/GetCandyMetadataResponseCreators.md)
 - [GetPublicKeyRequest](docs/GetPublicKeyRequest.md)
 - [GetSPLTokenResponse](docs/GetSPLTokenResponse.md)
 - [HexPrivateKey](docs/HexPrivateKey.md)
 - [InputBlockchainIdentifier](docs/InputBlockchainIdentifier.md)
 - [InputName](docs/InputName.md)
 - [ListNFTsResponse](docs/ListNFTsResponse.md)
 - [NFT](docs/NFT.md)
 - [NFTCollection](docs/NFTCollection.md)
 - [NFTData](docs/NFTData.md)
 - [NFTMintErrorResponse](docs/NFTMintErrorResponse.md)
 - [NFTMintFee](docs/NFTMintFee.md)
 - [NFTMintRequest](docs/NFTMintRequest.md)
 - [NFTOwnerAdvancedResponse](docs/NFTOwnerAdvancedResponse.md)
 - [NFTOwnerAdvancedResponseContract](docs/NFTOwnerAdvancedResponseContract.md)
 - [NFTOwnerResponse](docs/NFTOwnerResponse.md)
 - [NFTSearchRequest](docs/NFTSearchRequest.md)
 - [NFTSearchResponse](docs/NFTSearchResponse.md)
 - [NearPublicKey](docs/NearPublicKey.md)
 - [PrivateKey](docs/PrivateKey.md)
 - [PublicKey](docs/PublicKey.md)
 - [SecretPhrase](docs/SecretPhrase.md)
 - [SecretRecoveryPhrase](docs/SecretRecoveryPhrase.md)
 - [SolanaPublicKey](docs/SolanaPublicKey.md)
 - [SolanaTransaction](docs/SolanaTransaction.md)
 - [SolanaTransactionCompiledResponse](docs/SolanaTransactionCompiledResponse.md)
 - [SolanaTransactionMadeResponse](docs/SolanaTransactionMadeResponse.md)
 - [SupplyWalletRequest](docs/SupplyWalletRequest.md)
 - [TokenMetadataResponse](docs/TokenMetadataResponse.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionResult](docs/TransactionResult.md)
 - [TransferRequest](docs/TransferRequest.md)
 - [TransferResponse](docs/TransferResponse.md)
 - [TransferResponseCompiled](docs/TransferResponseCompiled.md)
 - [Wallet](docs/Wallet.md)
 - [WalletIdentifiers](docs/WalletIdentifiers.md)


## Documentation For Authorization



### APIKeyID

- **Type**: API key
- **API key parameter name**: APIKeyID
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: APIKeyID and passed in as the auth context for each request.


### APISecretKey

- **Type**: API key
- **API key parameter name**: APISecretKey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: APISecretKey and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@blockchainapi.com

