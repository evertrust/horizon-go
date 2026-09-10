# ServiceAccountTrustConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwks** | Pointer to **string** | The JWKS JSON representation | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Proxy** | Pointer to **string** | Proxy to fetch the JWKS with | [optional] 
**Url** | Pointer to **string** | URL to fetch the JWKS from | [optional] 

## Methods

### NewServiceAccountTrustConfig

`func NewServiceAccountTrustConfig() *ServiceAccountTrustConfig`

NewServiceAccountTrustConfig instantiates a new ServiceAccountTrustConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountTrustConfigWithDefaults

`func NewServiceAccountTrustConfigWithDefaults() *ServiceAccountTrustConfig`

NewServiceAccountTrustConfigWithDefaults instantiates a new ServiceAccountTrustConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwks

`func (o *ServiceAccountTrustConfig) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *ServiceAccountTrustConfig) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *ServiceAccountTrustConfig) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *ServiceAccountTrustConfig) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetType

`func (o *ServiceAccountTrustConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceAccountTrustConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceAccountTrustConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceAccountTrustConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProxy

`func (o *ServiceAccountTrustConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ServiceAccountTrustConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ServiceAccountTrustConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ServiceAccountTrustConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetUrl

`func (o *ServiceAccountTrustConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceAccountTrustConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceAccountTrustConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceAccountTrustConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


