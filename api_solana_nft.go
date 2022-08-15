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

// SolanaNFTApiService SolanaNFTApi service
type SolanaNFTApiService service

type ApiSolanaCreateNFTRequest struct {
	ctx context.Context
	ApiService *SolanaNFTApiService
	nFTMintRequest *NFTMintRequest
}

func (r ApiSolanaCreateNFTRequest) NFTMintRequest(nFTMintRequest NFTMintRequest) ApiSolanaCreateNFTRequest {
	r.nFTMintRequest = &nFTMintRequest
	return r
}

func (r ApiSolanaCreateNFTRequest) Execute() (*NFT, *http.Response, error) {
	return r.ApiService.SolanaCreateNFTExecute(r)
}

/*
SolanaCreateNFT Create an NFT on Solana

<a href="https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples/solana-nft/create-an-nft" target="_blank">See examples (Python, JavaScript)</a>.

Create a Metaplex NFT on Solana. 

Read more on this <a href="https://blog.blockchainapi.com/2021/11/16/a-note-on-nfts.html" target="_blank">here</a>.

Note: Please see <a href="https://blog.blockchainapi.com/2022/01/18/how-to-format-off-chain-nft-metadata.html" target="_blank">this article</a> to learn more about what `nft_upload_method` means and how storing the metadata of an NFT works.

If you're using `nft_upload_method = "LINK"`, then to add attributes to the NFT or an image, add them to a JSON file and upload that to Arweave/IPFS/Filecoin. See the JSON format <a href="https://blog.blockchainapi.com/2022/01/18/how-to-format-off-chain-nft-metadata.html">here</a>.

Then supply the link to the JSON file in `nft_url`. 

NOTE: Don't use `nft_metadata`. Values provided here do not do anything at the moment. We are fixing this soon.

`Cost: 5 Credits` (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSolanaCreateNFTRequest
*/
func (a *SolanaNFTApiService) SolanaCreateNFT(ctx context.Context) ApiSolanaCreateNFTRequest {
	return ApiSolanaCreateNFTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NFT
func (a *SolanaNFTApiService) SolanaCreateNFTExecute(r ApiSolanaCreateNFTRequest) (*NFT, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NFT
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SolanaNFTApiService.SolanaCreateNFT")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/solana/nft"

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
	localVarPostBody = r.nFTMintRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v NFTMintErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiSolanaGetNFTRequest struct {
	ctx context.Context
	ApiService *SolanaNFTApiService
	network string
	mintAddress string
}


func (r ApiSolanaGetNFTRequest) Execute() (*NFT, *http.Response, error) {
	return r.ApiService.SolanaGetNFTExecute(r)
}

/*
SolanaGetNFT Get an NFT's metadata

<a href="https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples/solana-nft/get-nft-metadata" target="_blank">See examples (Python, JavaScript)</a>. 
    
Get the metadata of an NFT.

If you're looking for metadata such as attributes and others, you can retrieve them from the link in the URI field of the NFT metadata returned. See the example on the right. The URI is an Arweave URL. That contains the attributes and other information about the NFT. That URL is stored on the Solana blockchain.

`Cost: 0.25 Credit` (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param network The network ID
 @param mintAddress The mint address of the NFT
 @return ApiSolanaGetNFTRequest
*/
func (a *SolanaNFTApiService) SolanaGetNFT(ctx context.Context, network string, mintAddress string) ApiSolanaGetNFTRequest {
	return ApiSolanaGetNFTRequest{
		ApiService: a,
		ctx: ctx,
		network: network,
		mintAddress: mintAddress,
	}
}

// Execute executes the request
//  @return NFT
func (a *SolanaNFTApiService) SolanaGetNFTExecute(r ApiSolanaGetNFTRequest) (*NFT, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NFT
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SolanaNFTApiService.SolanaGetNFT")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/solana/nft/{network}/{mint_address}"
	localVarPath = strings.Replace(localVarPath, "{"+"network"+"}", url.PathEscape(parameterToString(r.network, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mint_address"+"}", url.PathEscape(parameterToString(r.mintAddress, "")), -1)

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

type ApiSolanaGetNFTMintFeeRequest struct {
	ctx context.Context
	ApiService *SolanaNFTApiService
}


func (r ApiSolanaGetNFTMintFeeRequest) Execute() (*NFTMintFee, *http.Response, error) {
	return r.ApiService.SolanaGetNFTMintFeeExecute(r)
}

/*
SolanaGetNFTMintFee Get the NFT mint fee

<a href="https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples/solana-nft/get-nft-mint-fee" target="_blank">See examples (Python, JavaScript)</a>.
    
Get the estimated fee for minting an NFT on the Solana blockchain using the Metaplex protocol.

`Cost: 0 Credit` (Free) (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSolanaGetNFTMintFeeRequest
*/
func (a *SolanaNFTApiService) SolanaGetNFTMintFee(ctx context.Context) ApiSolanaGetNFTMintFeeRequest {
	return ApiSolanaGetNFTMintFeeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NFTMintFee
func (a *SolanaNFTApiService) SolanaGetNFTMintFeeExecute(r ApiSolanaGetNFTMintFeeRequest) (*NFTMintFee, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NFTMintFee
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SolanaNFTApiService.SolanaGetNFTMintFee")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/solana/nft/mint/fee"

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

type ApiSolanaGetNFTOwnerRequest struct {
	ctx context.Context
	ApiService *SolanaNFTApiService
	network string
	mintAddress string
}


func (r ApiSolanaGetNFTOwnerRequest) Execute() (*NFTOwnerResponse, *http.Response, error) {
	return r.ApiService.SolanaGetNFTOwnerExecute(r)
}

/*
SolanaGetNFTOwner Get owner of an NFT

<a href="https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples/solana-nft/get-nft-owner" target="_blank">See examples (Python, JavaScript)</a>. 
    
Get the owner of an NFT. This returns the public key of the wallet that owns the associated token account that owns the NFT.

If you want to get the associated token account that literally owns the NFT, derive the associated token account address from the public key returned and the NFT mint address using <a href="#operation/solanaDeriveAssociatedTokenAccountAddress">this endpoint</a>.

`Cost: 0.25 Credit` (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param network The network ID
 @param mintAddress The mint address of the NFT
 @return ApiSolanaGetNFTOwnerRequest
*/
func (a *SolanaNFTApiService) SolanaGetNFTOwner(ctx context.Context, network string, mintAddress string) ApiSolanaGetNFTOwnerRequest {
	return ApiSolanaGetNFTOwnerRequest{
		ApiService: a,
		ctx: ctx,
		network: network,
		mintAddress: mintAddress,
	}
}

// Execute executes the request
//  @return NFTOwnerResponse
func (a *SolanaNFTApiService) SolanaGetNFTOwnerExecute(r ApiSolanaGetNFTOwnerRequest) (*NFTOwnerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NFTOwnerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SolanaNFTApiService.SolanaGetNFTOwner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/solana/nft/{network}/{mint_address}/owner"
	localVarPath = strings.Replace(localVarPath, "{"+"network"+"}", url.PathEscape(parameterToString(r.network, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mint_address"+"}", url.PathEscape(parameterToString(r.mintAddress, "")), -1)

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

type ApiSolanaGetNFTOwnerAdvancedRequest struct {
	ctx context.Context
	ApiService *SolanaNFTApiService
	network string
	mintAddress string
}


func (r ApiSolanaGetNFTOwnerAdvancedRequest) Execute() (*NFTOwnerAdvancedResponse, *http.Response, error) {
	return r.ApiService.SolanaGetNFTOwnerAdvancedExecute(r)
}

/*
SolanaGetNFTOwnerAdvanced Get owner of an NFT (advanced)

<a href="https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples/solana-nft/get-nft-owner" target="_blank">See examples (Python, JavaScript)</a>. 
    
Get the owner, state, listed price, and listed marketplace (if any) of an NFT. 

Here's are a couple of example responses:
```
{
    'contract': {
        'contract_blockchain_identifier': 'M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K', 
        'contract_id': 'magic-eden-v2', 
        'contract_name': 'Magic Eden v2', 
        'contract_type': 'marketplace'
    }, 
    'owner': '25UJMR3FiMM6noQtPEaCJ6eDU2YQ7myDhikVQXmMuSRW', 
    'price': 50000000, 
    'state': 'listing'
}
```

```
{
    'contract': null, 
    'owner': 'C37PJiJU8WTgoUoFqmB1Maw8hkuENDZoGDQA1pm54Fdd', 
    'price': null, 
    'state': 'holding'
}
```

This function will return whether the NFT is `listed`, `loaned` (Yawww NFT loans), `otc`, `staked`, `burned`, or `held`.

If listed, it will return the contract, the readable name of the contract (e.g., Magic Eden, OpenSea), the contract ID (if any; e.g., open-sea), the owner, and the listed price. From this, you can get the floor of a collection. We currently support Magic Eden (v1, v2), Exchange.Art (auction, singles), CoralCube, Solanart (v1, v2), Yawww Loans, Yawww OTC, OpenSea, Fractal, SolSea, and AlphaArt. 

If loaned, it will return the loan requester as the owner, the loan amount, and the loan contract. We only support the Yawww loaning contract.

If listed on an OTC marketplace, it will return the same information as `listed`. OTC is used to distinguish between marketplaces that respect royalties (OTC) and those that don't (normal ones). The only `OTC` contract we track is Yawwww, at the moment. We do not yet track Solanart v3.

If staked, it will return the owner and the staking contract public key.

If burned, it will return the `burner` as the `owner`.

If held, it will simply return the owner.

If you want to get the literal owner, which may or may not be the same as the owner returned here, call the simplified [get NFT owner function](/#tag/Solana-NFT/operation/solanaGetNFTOwner). For example, Bob might own the NFT, but if it is listed on Magic Eden, then the NFT is held in escrow and "owned" by Magic Eden. The simplified function will return Magic Eden as the owner (the literal owner). This advanced function will tell you the implied owner, which would be Bob.

`Cost: 1.0 Credit` (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param network The network ID
 @param mintAddress The mint address of the NFT
 @return ApiSolanaGetNFTOwnerAdvancedRequest
*/
func (a *SolanaNFTApiService) SolanaGetNFTOwnerAdvanced(ctx context.Context, network string, mintAddress string) ApiSolanaGetNFTOwnerAdvancedRequest {
	return ApiSolanaGetNFTOwnerAdvancedRequest{
		ApiService: a,
		ctx: ctx,
		network: network,
		mintAddress: mintAddress,
	}
}

// Execute executes the request
//  @return NFTOwnerAdvancedResponse
func (a *SolanaNFTApiService) SolanaGetNFTOwnerAdvancedExecute(r ApiSolanaGetNFTOwnerAdvancedRequest) (*NFTOwnerAdvancedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NFTOwnerAdvancedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SolanaNFTApiService.SolanaGetNFTOwnerAdvanced")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/solana/nft/{network}/{mint_address}/owner_advanced"
	localVarPath = strings.Replace(localVarPath, "{"+"network"+"}", url.PathEscape(parameterToString(r.network, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mint_address"+"}", url.PathEscape(parameterToString(r.mintAddress, "")), -1)

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

type ApiSolanaGetNFTsCandyMachineIdRequest struct {
	ctx context.Context
	ApiService *SolanaNFTApiService
	getCandyMachineIDRequest *GetCandyMachineIDRequest
}

func (r ApiSolanaGetNFTsCandyMachineIdRequest) GetCandyMachineIDRequest(getCandyMachineIDRequest GetCandyMachineIDRequest) ApiSolanaGetNFTsCandyMachineIdRequest {
	r.getCandyMachineIDRequest = &getCandyMachineIDRequest
	return r
}

func (r ApiSolanaGetNFTsCandyMachineIdRequest) Execute() (*GetCandyMachineIDResponse, *http.Response, error) {
	return r.ApiService.SolanaGetNFTsCandyMachineIdExecute(r)
}

/*
SolanaGetNFTsCandyMachineId Get the ID of the candy machine of an NFT 

<a href="https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples/solana-nft/get-nft-candy-machine-id" target="_blank">See examples (Python, JavaScript)</a>.

Get the candy machine ID from where the NFT came, if any. NFTs can also be minted without a candy machine.

It's also possible that we return "Not Found" when the NFT actually did come from a version of a candy machine. We check for the most popular versions of candy machine, but it is possible that someone creates their own candy machine version and mints NFTs from it.

`Cost: 1 Credit` (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSolanaGetNFTsCandyMachineIdRequest
*/
func (a *SolanaNFTApiService) SolanaGetNFTsCandyMachineId(ctx context.Context) ApiSolanaGetNFTsCandyMachineIdRequest {
	return ApiSolanaGetNFTsCandyMachineIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCandyMachineIDResponse
func (a *SolanaNFTApiService) SolanaGetNFTsCandyMachineIdExecute(r ApiSolanaGetNFTsCandyMachineIdRequest) (*GetCandyMachineIDResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCandyMachineIDResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SolanaNFTApiService.SolanaGetNFTsCandyMachineId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/solana/nft/candy_machine_id"

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
	localVarPostBody = r.getCandyMachineIDRequest
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

type ApiSolanaSearchNFTsRequest struct {
	ctx context.Context
	ApiService *SolanaNFTApiService
	nFTSearchRequest *NFTSearchRequest
}

func (r ApiSolanaSearchNFTsRequest) NFTSearchRequest(nFTSearchRequest NFTSearchRequest) ApiSolanaSearchNFTsRequest {
	r.nFTSearchRequest = &nFTSearchRequest
	return r
}

func (r ApiSolanaSearchNFTsRequest) Execute() ([]NFTSearchResponse, *http.Response, error) {
	return r.ApiService.SolanaSearchNFTsExecute(r)
}

/*
SolanaSearchNFTs Search NFTs on Solana

<a href="https://github.com/BL0CK-X/the-blockchain-api/tree/main/examples/solana-nft/search-nfts" target="_blank">See examples (Python, JavaScript)</a>.

With this endpoint, you can search for NFTs by their symbol, name of NFTs, uuid, configuration address, and update authority.

The output is a list of NFTs that match your query.

You can also provide multiple search clauses, such as the update authority (`update_authority="G17UmNGnMJ851x3M1JXocgpft1afcYedjPuFpo1ohhCk"`) and symbol begins with "Sol" (`symbol="Sol", symbol_search_method='begins_with'`).

`Cost: 1 Credit` (<a href="#section/Pricing">See Pricing</a>)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSolanaSearchNFTsRequest
*/
func (a *SolanaNFTApiService) SolanaSearchNFTs(ctx context.Context) ApiSolanaSearchNFTsRequest {
	return ApiSolanaSearchNFTsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []NFTSearchResponse
func (a *SolanaNFTApiService) SolanaSearchNFTsExecute(r ApiSolanaSearchNFTsRequest) ([]NFTSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []NFTSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SolanaNFTApiService.SolanaSearchNFTs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/solana/nft/search"

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
	localVarPostBody = r.nFTSearchRequest
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
