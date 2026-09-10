# StaticJWKS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwks** | Pointer to **string** | The JWKS JSON representation | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewStaticJWKS

`func NewStaticJWKS() *StaticJWKS`

NewStaticJWKS instantiates a new StaticJWKS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticJWKSWithDefaults

`func NewStaticJWKSWithDefaults() *StaticJWKS`

NewStaticJWKSWithDefaults instantiates a new StaticJWKS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwks

`func (o *StaticJWKS) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *StaticJWKS) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *StaticJWKS) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *StaticJWKS) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetType

`func (o *StaticJWKS) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StaticJWKS) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StaticJWKS) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StaticJWKS) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


