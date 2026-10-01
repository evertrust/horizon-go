# SectigoDCVProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **string** | Name of the login/password credentials configuration holding the SCM API client, with the OAuth client id as login and the client secret as password | 
**DcvMethod** | **string** | DNS method used to validate a domain. It is a fallback: a domain that already holds a CNAME or TXT validation is re-validated with its own method, and this method applies only to a domain that has never been validated or whose existing validation uses a method that cannot be published over DNS. | 
**Endpoint** | **string** | Sectigo Certificate Manager (SCM) API base URL | 
**Name** | **string** | Unique name of the DCV provider configuration | 
**OauthTokenEndpoint** | Pointer to **string** | OAuth token endpoint used to obtain a bearer token for the SCM API. Defaults to Sectigo&#39;s SSO realm. | [optional] [default to "https://auth.sso.sectigo.com/auth/realms/apiclients/protocol/openid-connect/token"]
**OrganizationId** | Pointer to **int64** | Restricts the domain listing to this Sectigo organization or department. When unset, every domain of the customer account is listed. | [optional] 
**Proxy** | Pointer to **string** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | **NullableString** | Request timeout | 
**Type** | **string** | Provider type discriminator | 

## Methods

### NewSectigoDCVProviderConfig

`func NewSectigoDCVProviderConfig(credentials string, dcvMethod string, endpoint string, name string, timeout NullableString, type_ string, ) *SectigoDCVProviderConfig`

NewSectigoDCVProviderConfig instantiates a new SectigoDCVProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectigoDCVProviderConfigWithDefaults

`func NewSectigoDCVProviderConfigWithDefaults() *SectigoDCVProviderConfig`

NewSectigoDCVProviderConfigWithDefaults instantiates a new SectigoDCVProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *SectigoDCVProviderConfig) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SectigoDCVProviderConfig) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SectigoDCVProviderConfig) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDcvMethod

`func (o *SectigoDCVProviderConfig) GetDcvMethod() string`

GetDcvMethod returns the DcvMethod field if non-nil, zero value otherwise.

### GetDcvMethodOk

`func (o *SectigoDCVProviderConfig) GetDcvMethodOk() (*string, bool)`

GetDcvMethodOk returns a tuple with the DcvMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcvMethod

`func (o *SectigoDCVProviderConfig) SetDcvMethod(v string)`

SetDcvMethod sets DcvMethod field to given value.


### GetEndpoint

`func (o *SectigoDCVProviderConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SectigoDCVProviderConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SectigoDCVProviderConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *SectigoDCVProviderConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectigoDCVProviderConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectigoDCVProviderConfig) SetName(v string)`

SetName sets Name field to given value.


### GetOauthTokenEndpoint

`func (o *SectigoDCVProviderConfig) GetOauthTokenEndpoint() string`

GetOauthTokenEndpoint returns the OauthTokenEndpoint field if non-nil, zero value otherwise.

### GetOauthTokenEndpointOk

`func (o *SectigoDCVProviderConfig) GetOauthTokenEndpointOk() (*string, bool)`

GetOauthTokenEndpointOk returns a tuple with the OauthTokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenEndpoint

`func (o *SectigoDCVProviderConfig) SetOauthTokenEndpoint(v string)`

SetOauthTokenEndpoint sets OauthTokenEndpoint field to given value.

### HasOauthTokenEndpoint

`func (o *SectigoDCVProviderConfig) HasOauthTokenEndpoint() bool`

HasOauthTokenEndpoint returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SectigoDCVProviderConfig) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SectigoDCVProviderConfig) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SectigoDCVProviderConfig) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SectigoDCVProviderConfig) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProxy

`func (o *SectigoDCVProviderConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SectigoDCVProviderConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SectigoDCVProviderConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SectigoDCVProviderConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTimeout

`func (o *SectigoDCVProviderConfig) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SectigoDCVProviderConfig) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SectigoDCVProviderConfig) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *SectigoDCVProviderConfig) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SectigoDCVProviderConfig) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *SectigoDCVProviderConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SectigoDCVProviderConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SectigoDCVProviderConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


