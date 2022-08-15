/*
The Blockchain API

# About  Our vision is to make it super easy to interact with the decentralized web. We want you to be able to do this in any coding language and do it easily and quickly.   You're a key part of our vision. We love feature requests! <a href=\"#section/Feature-Requests\">Make one!</a>  # How to Use the API  To use the API, you simply need to create an API key pair.  Doing so takes less than a minute. Simply go to <a target=\"_blank\" href=\"https://dashboard.blockchainapi.com\">the dashboard</a>, create an account, and generate a key pair. You can now use this pair to make API requests. You must create your first pair via the dashboard.  # Feature Requests  Got a feature request? Submit it as an issue on <a target=\"_blank\" href=\"https://github.com/BL0CK-X/the-blockchain-api/issues/new/choose\">our GitHub repo</a> or [email us](mailto:info@blockchainapi.com).  # Contact  <figure>     <img          width=\"40px\"         height=\"40px\"          src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/phone.svg\"     />     <figcaption style=\"textAlign:center;\">Text / Call: +1 (415) 888 4745 </figcaption> </figure> <a href=\"mailto:info@blockchainapi.com\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/email.svg\"         />         <figcaption style=\"textAlign:center;\">Email us: info@blockchainapi.com</figcaption>     </figure> </a> <a href=\"https://discord.gg/d49yzrZRAF\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/discord.svg\"         />         <figcaption style=\"textAlign:center;\">Join our Discord</figcaption>     </figure> </a> <a href=\"https://twitter.com/BlockchainAP1\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/twitter.svg\"         />          <figcaption style=\"textAlign:center;\">Follow us on Twitter</figcaption>     </figure> </a> <a href=\"https://github.com/BL0CK-X/the-blockchain-api\" target=\"_blank\">     <figure>         <img              width=\"40px\"             height=\"40px\"              src=\"https://theblockchainapi-docs.s3.amazonaws.com/icons/github.svg\"         />         <figcaption style=\"textAlign:center;\">Star us on Github</figcaption>     </figure> </a>  # Security  Common dogma is to never give out your seed phrase. We agree. It's a matter of security, and anyone who has your seed phrase can irreversibly empty your wallet.   <b>When using an API endpoint that requires a seed phrase, we highly recommend that users use or create a wallet that they do not use as their primary wallet.</b>   How you make this work depends on what you're doing. If you're minting an NFT for example, we recommend creating a new wallet and then transferring just enough SOL to that wallet to mint the NFT. This is possible on Solana because such transactions cost less than a penny. We will have more tutorials in the future that make it easier for you to be secure when using our API.  We have easy-to-use endpoints for <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1secret_recovery_phrase/post\">creating a new seed phrase</a> and then <a href=\"#tag/Solana-Wallet/paths/~1solana~1wallet~1public_key/post\">deriving a public key</a> to enable you to transfer to that new wallet.  Let's have a constructive dialog about this. Feel free to <a href=\"#section/Contact\">contact us</a>. I made a video discussing this matter <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.  Note: We have had a couple of individuals harrass and threaten us. These individuals did <b>not</b> try our API or speak to anyone who has used it. They simply saw that we require a seed phrase for certain endpoints and figured that the proper response was to attack us. (I explain why we do <a target=\"_blank\" href=\"https://youtu.be/m9unUb8Z9qU\">here</a>.) Such harrassment and threats are not only harmful, but they are also illegal, and we will report offenders. Do not harrass us. Rather, feel free to discuss your concerns with us and we will be more than happy to work with you to come up with a solution.  # Pricing  <b>Each user receives 50,000 free credits each month.</b> Each endpoint costs a certain amount credits. Scroll below to any endpoint (i.e., function) to see how much each endpoint costs. (Or CMD+F `Cost: 0 Credit`, for example)  You can learn more about our pricing <a href=\"https://dashboard.blockchainapi.com/billing\" target=\"_blank\">here</a>.   We frequently do custom plans. If our pricing doesn't work for you, <a href=\"#section/Contact\">contact us</a>.  If you have questions, concerns, feedback, or ideas, <a href=\"#section/Contact\">contact us</a>.  # SDKs / API Wrappers  We have examples using both our <a href=\"https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples\" target=\"_blank\">Python wrapper and our JavaScript wrapper here</a>.  We have built a custom <a href=\"https://github.com/BL0CK-X/the-blockchain-api-python-wrapper\" target=\"_blank\">Python wrapper</a>.  `pip install theblockchainapi`  We also have published a <a href=\"https://github.com/BL0CK-X/theblockchainapi-javascript-wrapper\" target=\"_blank\">JavaScript Wrapper</a>.  `npm install theblockchainapi`  We also have auto-generated wrappers for the following languages: - <a href=\"https://github.com/BL0CK-X/theblockchainapi-go-wrapper\" target = \"_blank\">Go</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-java-wrapper\" target = \"_blank\">Java</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-kotlin-wrapper\" target = \"_blank\">Kotlin</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-php-wrapper\" target = \"_blank\">PHP</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-swift-wrapper\" target = \"_blank\">Swift5</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-typescript-wrapper\" target = \"_blank\">TypeScript</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-csharp-wrapper\" target = \"_blank\">C#</a> - <a href=\"https://github.com/BL0CK-X/theblockchainapi-dart-wrapper\" target = \"_blank\">Dart</a>  If you would like a different language as well, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.  If you run into any bugs with the wrappers, submit an issue <a href=\"https://github.com/BL0CK-X/theblockchainapi-wrappers/issues/new\" target=\"_blank\">here</a>.

API version: null
Contact: info@blockchainapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// WalletApiService WalletApi service
type WalletApiService service

type ApiDerivePrivateKeyRequest struct {
	ctx context.Context
	ApiService *WalletApiService
	blockchain string
	supplyWalletRequest *SupplyWalletRequest
}

func (r ApiDerivePrivateKeyRequest) SupplyWalletRequest(supplyWalletRequest SupplyWalletRequest) ApiDerivePrivateKeyRequest {
	r.supplyWalletRequest = &supplyWalletRequest
	return r
}

func (r ApiDerivePrivateKeyRequest) Execute() (*GeneralGeneratePrivateKeyResponse, *http.Response, error) {
	return r.ApiService.DerivePrivateKeyExecute(r)
}

/*
DerivePrivateKey Derive private key

<a href="https://github.com/BL0CK-X/blockchain-api/tree/main/examples/wallet/derive-private-key" target="_blank">See examples (Python, JavaScript)</a>.

Use this endpoint to securely derive a private key for a wallet. 

For example, if you have a seed phrase and want to derive the corresponding private key, use this endpoint.

`Cost: 0 Credit` (Free) (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain The blockchain you want to use 
 @return ApiDerivePrivateKeyRequest
*/
func (a *WalletApiService) DerivePrivateKey(ctx context.Context, blockchain string) ApiDerivePrivateKeyRequest {
	return ApiDerivePrivateKeyRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
	}
}

// Execute executes the request
//  @return GeneralGeneratePrivateKeyResponse
func (a *WalletApiService) DerivePrivateKeyExecute(r ApiDerivePrivateKeyRequest) (*GeneralGeneratePrivateKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GeneralGeneratePrivateKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletApiService.DerivePrivateKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{blockchain}/wallet/private_key"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.supplyWalletRequest == nil {
		return localVarReturnValue, nil, reportError("supplyWalletRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.supplyWalletRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyID"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APIKeyID"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APISecretKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APISecretKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeriveWalletIdentifierRequest struct {
	ctx context.Context
	ApiService *WalletApiService
	blockchain string
	supplyWalletRequest *SupplyWalletRequest
}

func (r ApiDeriveWalletIdentifierRequest) SupplyWalletRequest(supplyWalletRequest SupplyWalletRequest) ApiDeriveWalletIdentifierRequest {
	r.supplyWalletRequest = &supplyWalletRequest
	return r
}

func (r ApiDeriveWalletIdentifierRequest) Execute() (*WalletIdentifiers, *http.Response, error) {
	return r.ApiService.DeriveWalletIdentifierExecute(r)
}

/*
DeriveWalletIdentifier Derive wallet identifier

<a href="https://github.com/BL0CK-X/blockchain-api/tree/main/examples/wallet/derive-wallet-identifier" target="_blank">See examples (Python, JavaScript)</a>.

Derive the identifier for a wallet.

`Cost: 0 Credit` (Free) (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain The blockchain you want to use 
 @return ApiDeriveWalletIdentifierRequest
*/
func (a *WalletApiService) DeriveWalletIdentifier(ctx context.Context, blockchain string) ApiDeriveWalletIdentifierRequest {
	return ApiDeriveWalletIdentifierRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
	}
}

// Execute executes the request
//  @return WalletIdentifiers
func (a *WalletApiService) DeriveWalletIdentifierExecute(r ApiDeriveWalletIdentifierRequest) (*WalletIdentifiers, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WalletIdentifiers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletApiService.DeriveWalletIdentifier")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{blockchain}/wallet/identifier"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.supplyWalletRequest == nil {
		return localVarReturnValue, nil, reportError("supplyWalletRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.supplyWalletRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyID"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APIKeyID"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APISecretKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APISecretKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGeneratePrivateKeyRequest struct {
	ctx context.Context
	ApiService *WalletApiService
	blockchain string
}


func (r ApiGeneratePrivateKeyRequest) Execute() (*GeneralGeneratePrivateKeyResponse, *http.Response, error) {
	return r.ApiService.GeneratePrivateKeyExecute(r)
}

/*
GeneratePrivateKey Generate private key

<a href="https://github.com/BL0CK-X/blockchain-api/tree/main/examples/wallet/generate-private-key" target="_blank">See examples (Python, JavaScript)</a>.

Use this endpoint to securely and randomly generate a private key for a wallet. 

`Cost: 0 Credit` (Free) (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain The blockchain you want to use 
 @return ApiGeneratePrivateKeyRequest
*/
func (a *WalletApiService) GeneratePrivateKey(ctx context.Context, blockchain string) ApiGeneratePrivateKeyRequest {
	return ApiGeneratePrivateKeyRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
	}
}

// Execute executes the request
//  @return GeneralGeneratePrivateKeyResponse
func (a *WalletApiService) GeneratePrivateKeyExecute(r ApiGeneratePrivateKeyRequest) (*GeneralGeneratePrivateKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GeneralGeneratePrivateKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletApiService.GeneratePrivateKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{blockchain}/wallet/generate/private_key"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyID"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APIKeyID"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APISecretKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APISecretKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGenerateSeedPhraseRequest struct {
	ctx context.Context
	ApiService *WalletApiService
	blockchain string
	generateSeedPhraseRequest *GenerateSeedPhraseRequest
}

func (r ApiGenerateSeedPhraseRequest) GenerateSeedPhraseRequest(generateSeedPhraseRequest GenerateSeedPhraseRequest) ApiGenerateSeedPhraseRequest {
	r.generateSeedPhraseRequest = &generateSeedPhraseRequest
	return r
}

func (r ApiGenerateSeedPhraseRequest) Execute() (*GeneralSecretPhrase, *http.Response, error) {
	return r.ApiService.GenerateSeedPhraseExecute(r)
}

/*
GenerateSeedPhrase Generate seed phrase

<a href="https://github.com/BL0CK-X/blockchain-api/tree/main/examples/wallet/generate-secret-phrase" target="_blank">See examples (Python, JavaScript)</a>.

Use this endpoint to securely and randomly generate a seed phrase for a wallet. 

`Cost: 0 Credit` (Free) (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain The blockchain you want to use 
 @return ApiGenerateSeedPhraseRequest
*/
func (a *WalletApiService) GenerateSeedPhrase(ctx context.Context, blockchain string) ApiGenerateSeedPhraseRequest {
	return ApiGenerateSeedPhraseRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
	}
}

// Execute executes the request
//  @return GeneralSecretPhrase
func (a *WalletApiService) GenerateSeedPhraseExecute(r ApiGenerateSeedPhraseRequest) (*GeneralSecretPhrase, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GeneralSecretPhrase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletApiService.GenerateSeedPhrase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{blockchain}/wallet/generate/secret_recovery_phrase"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.generateSeedPhraseRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyID"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APIKeyID"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APISecretKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APISecretKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAirdropRequest struct {
	ctx context.Context
	ApiService *WalletApiService
	blockchain string
	generalAirdropRequest *GeneralAirdropRequest
}

func (r ApiGetAirdropRequest) GeneralAirdropRequest(generalAirdropRequest GeneralAirdropRequest) ApiGetAirdropRequest {
	r.generalAirdropRequest = &generalAirdropRequest
	return r
}

func (r ApiGetAirdropRequest) Execute() (*GeneralTransactionMadeResponse, *http.Response, error) {
	return r.ApiService.GetAirdropExecute(r)
}

/*
GetAirdrop Get an airdrop

<a href="https://github.com/BL0CK-X/blockchain-api/tree/main/examples/wallet/get-airdrop" target="_blank">See examples (Python, JavaScript)</a>.
    
- Ethereum: Receive an airdrop of 0.001 ETH on Ropsten (not real ETH). 
- Solana: Use this endpoint to get an airdrop of SOL on the devnet (not real SOL). Amount of 0.015, which is the minimum amount of SOL you need to mint an NFT and then transfer it to another wallet.

`Cost: 0 Credit` (Free) (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain The blockchain you want to use 
 @return ApiGetAirdropRequest
*/
func (a *WalletApiService) GetAirdrop(ctx context.Context, blockchain string) ApiGetAirdropRequest {
	return ApiGetAirdropRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
	}
}

// Execute executes the request
//  @return GeneralTransactionMadeResponse
func (a *WalletApiService) GetAirdropExecute(r ApiGetAirdropRequest) (*GeneralTransactionMadeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GeneralTransactionMadeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletApiService.GetAirdrop")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{blockchain}/wallet/airdrop"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.generalAirdropRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyID"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APIKeyID"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APISecretKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APISecretKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetBalanceRequest struct {
	ctx context.Context
	ApiService *WalletApiService
	blockchain string
	generalBalanceRequest *GeneralBalanceRequest
}

func (r ApiGetBalanceRequest) GeneralBalanceRequest(generalBalanceRequest GeneralBalanceRequest) ApiGetBalanceRequest {
	r.generalBalanceRequest = &generalBalanceRequest
	return r
}

func (r ApiGetBalanceRequest) Execute() (*GeneralBalanceResponse, *http.Response, error) {
	return r.ApiService.GetBalanceExecute(r)
}

/*
GetBalance Get wallet's balance of X

<a href="https://github.com/BL0CK-X/blockchain-api/tree/main/examples/wallet/get-wallet-balance" target="_blank">See examples (Python, JavaScript)</a>.
    
See the balance of a wallet in the native blockchain currency (e.g., ETH, SOL) or any token (e.g., ERC-20, NFTs, SPL, etc.).

To get the balance of a specific token, supply the `token_blockchain_identifier` of the token. 

You can also use this endpoint to see whether or not a person owns an NFT. Just supply the `token_blockchain_identifier` of the NFT. A balance of "1" means the person owns the NFT, and a balance of "0" means the person does not own the NFT.

`Cost: 0.25 Credit` (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain The blockchain you want to use 
 @return ApiGetBalanceRequest
*/
func (a *WalletApiService) GetBalance(ctx context.Context, blockchain string) ApiGetBalanceRequest {
	return ApiGetBalanceRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
	}
}

// Execute executes the request
//  @return GeneralBalanceResponse
func (a *WalletApiService) GetBalanceExecute(r ApiGetBalanceRequest) (*GeneralBalanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GeneralBalanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletApiService.GetBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{blockchain}/wallet/balance"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.generalBalanceRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyID"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APIKeyID"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APISecretKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APISecretKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTransferRequest struct {
	ctx context.Context
	ApiService *WalletApiService
	blockchain string
	generalTransferRequest *GeneralTransferRequest
}

func (r ApiTransferRequest) GeneralTransferRequest(generalTransferRequest GeneralTransferRequest) ApiTransferRequest {
	r.generalTransferRequest = &generalTransferRequest
	return r
}

func (r ApiTransferRequest) Execute() (*GeneralTransactionResponse, *http.Response, error) {
	return r.ApiService.TransferExecute(r)
}

/*
Transfer Transfer crypto, a token, or an NFT to another wallet

<a href="https://github.com/BL0CK-X/blockchain-api/tree/main/examples/wallet/transfer-crypto-token-or-nft" target="_blank">See transfer ETH/SOL/crypto example (Python, JavaScript)</a>.

<a href="https://github.com/BL0CK-X/blockchain-api/tree/main/examples/wallet/transfer-crypto-token-or-nft" target="_blank">See transfer NFT/token example (Python, JavaScript)</a>.

This is a powerful function. It might be slightly confusing because there are several optional parameters, so take some time to review it. 

This function can send
- The native currency of the selected blockchain (e.g., SOL, ETH, BNB, etc.)
- A token (e.g., an SPL token, ERC-20 token, etc.)
- An NFT

You can sign and submit the transaction for confirmation; and you can select to simply return the compiled transaction so that you can submit it to the user for signing (e.g., via Phantom or Metamask; no private keys required in this case).

If you're transferring an NFT or a token, supply its respective `token_blockchain_identifier`.

SENDER: Note that the wallet information is used to authorize the sending of the tokens and identifies the source of the tokens. If `return_compiled_transaction = false`, we sign and submit the transaction (`wallet` is required in this case; do not provide a value for `sender_blockchain_identifier`). If `return_compiled_transaction = true`, we compile the transaction (`sender_blockchain_identifier` is required in this case; do not provide `wallet`).

RECIPIENT: `recipient_blockchain_identifier` identifies the receiver. This is entirely separate from the information used for the SENDER above. So, in this API call, there are two wallets involved, but only one (namely, the SENDER) is needed to authorize the transaction (if you want us to sign and submitting it).

FEE_PAYER (Solana only): The fee payer of the transaction defaults to `wallet` (or `sender_blockchain_identifier`). To set a different fee payer, provide a value for `fee_payer_wallet`.

`Cost: 2 Credit` (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain The blockchain you want to use 
 @return ApiTransferRequest
*/
func (a *WalletApiService) Transfer(ctx context.Context, blockchain string) ApiTransferRequest {
	return ApiTransferRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
	}
}

// Execute executes the request
//  @return GeneralTransactionResponse
func (a *WalletApiService) TransferExecute(r ApiTransferRequest) (*GeneralTransactionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GeneralTransactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletApiService.Transfer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{blockchain}/wallet/transfer"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.generalTransferRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyID"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APIKeyID"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APISecretKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["APISecretKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
