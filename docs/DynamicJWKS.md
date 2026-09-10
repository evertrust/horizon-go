# DynamicJWKS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proxy** | Pointer to **string** | Proxy to fetch the JWKS with | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL to fetch the JWKS from | [optional] 

## Methods

### NewDynamicJWKS

`func NewDynamicJWKS() *DynamicJWKS`

NewDynamicJWKS instantiates a new DynamicJWKS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicJWKSWithDefaults

`func NewDynamicJWKSWithDefaults() *DynamicJWKS`

NewDynamicJWKSWithDefaults instantiates a new DynamicJWKS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProxy

`func (o *DynamicJWKS) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DynamicJWKS) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DynamicJWKS) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DynamicJWKS) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetType

`func (o *DynamicJWKS) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DynamicJWKS) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DynamicJWKS) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DynamicJWKS) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *DynamicJWKS) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DynamicJWKS) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DynamicJWKS) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DynamicJWKS) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


