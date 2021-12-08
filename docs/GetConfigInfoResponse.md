# GetConfigInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to **string** |  | [optional] 
**Creators** | Pointer to [**[]GetConfigInfoResponseCreators**](GetConfigInfoResponseCreators.md) |  | [optional] 
**IsMutable** | Pointer to **bool** |  | [optional] 
**RetainAuthority** | Pointer to **string** | A public key address | [optional] 
**SellerFeeBasisPoints** | Pointer to **float32** | The fee of each sale the creators receive. 100 basis points &#x3D; 1% | [optional] 
**Symbol** | Pointer to **string** | The symbol of the candy machine NFT colletion | [optional] 

## Methods

### NewGetConfigInfoResponse

`func NewGetConfigInfoResponse() *GetConfigInfoResponse`

NewGetConfigInfoResponse instantiates a new GetConfigInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfigInfoResponseWithDefaults

`func NewGetConfigInfoResponseWithDefaults() *GetConfigInfoResponse`

NewGetConfigInfoResponseWithDefaults instantiates a new GetConfigInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *GetConfigInfoResponse) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *GetConfigInfoResponse) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *GetConfigInfoResponse) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *GetConfigInfoResponse) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetCreators

`func (o *GetConfigInfoResponse) GetCreators() []GetConfigInfoResponseCreators`

GetCreators returns the Creators field if non-nil, zero value otherwise.

### GetCreatorsOk

`func (o *GetConfigInfoResponse) GetCreatorsOk() (*[]GetConfigInfoResponseCreators, bool)`

GetCreatorsOk returns a tuple with the Creators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreators

`func (o *GetConfigInfoResponse) SetCreators(v []GetConfigInfoResponseCreators)`

SetCreators sets Creators field to given value.

### HasCreators

`func (o *GetConfigInfoResponse) HasCreators() bool`

HasCreators returns a boolean if a field has been set.

### GetIsMutable

`func (o *GetConfigInfoResponse) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *GetConfigInfoResponse) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *GetConfigInfoResponse) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *GetConfigInfoResponse) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetRetainAuthority

`func (o *GetConfigInfoResponse) GetRetainAuthority() string`

GetRetainAuthority returns the RetainAuthority field if non-nil, zero value otherwise.

### GetRetainAuthorityOk

`func (o *GetConfigInfoResponse) GetRetainAuthorityOk() (*string, bool)`

GetRetainAuthorityOk returns a tuple with the RetainAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAuthority

`func (o *GetConfigInfoResponse) SetRetainAuthority(v string)`

SetRetainAuthority sets RetainAuthority field to given value.

### HasRetainAuthority

`func (o *GetConfigInfoResponse) HasRetainAuthority() bool`

HasRetainAuthority returns a boolean if a field has been set.

### GetSellerFeeBasisPoints

`func (o *GetConfigInfoResponse) GetSellerFeeBasisPoints() float32`

GetSellerFeeBasisPoints returns the SellerFeeBasisPoints field if non-nil, zero value otherwise.

### GetSellerFeeBasisPointsOk

`func (o *GetConfigInfoResponse) GetSellerFeeBasisPointsOk() (*float32, bool)`

GetSellerFeeBasisPointsOk returns a tuple with the SellerFeeBasisPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerFeeBasisPoints

`func (o *GetConfigInfoResponse) SetSellerFeeBasisPoints(v float32)`

SetSellerFeeBasisPoints sets SellerFeeBasisPoints field to given value.

### HasSellerFeeBasisPoints

`func (o *GetConfigInfoResponse) HasSellerFeeBasisPoints() bool`

HasSellerFeeBasisPoints returns a boolean if a field has been set.

### GetSymbol

`func (o *GetConfigInfoResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetConfigInfoResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetConfigInfoResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetConfigInfoResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


