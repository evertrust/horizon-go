# SectigoDCVProviderConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the login/password credentials configuration holding the SCM API client, with the OAuth client id as login and the client secret as password | 
**DcvMethod** | **string** | DNS method used to validate a domain. It is a fallback: a domain that already holds a CNAME or TXT validation is re-validated with its own method, and this method applies only to a domain that has never been validated or whose existing validation uses a method that cannot be published over DNS. | 
**Endpoint** | **string** | Sectigo Certificate Manager (SCM) API base URL | 
**Name** | **string** | Unique name of the DCV provider configuration | 
**OauthTokenEndpoint** | **string** | OAuth token endpoint used to obtain a bearer token for the SCM API. Defaults to Sectigo&#39;s SSO realm. | [default to "https://auth.sso.sectigo.com/auth/realms/apiclients/protocol/openid-connect/token"]
**OrganizationId** | Pointer to **int64** | Restricts the domain listing to this Sectigo organization or department. When unset, every domain of the customer account is listed. | [optional] 
**Proxy** | Pointer to **string** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | **NullableString** | Request timeout | 
**Type** | **string** | Provider type discriminator | 

## Methods

### NewSectigoDCVProviderConfigResponse

`func NewSectigoDCVProviderConfigResponse(id string, credentials string, dcvMethod string, endpoint string, name string, oauthTokenEndpoint string, timeout NullableString, type_ string, ) *SectigoDCVProviderConfigResponse`

NewSectigoDCVProviderConfigResponse instantiates a new SectigoDCVProviderConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectigoDCVProviderConfigResponseWithDefaults

`func NewSectigoDCVProviderConfigResponseWithDefaults() *SectigoDCVProviderConfigResponse`

NewSectigoDCVProviderConfigResponseWithDefaults instantiates a new SectigoDCVProviderConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SectigoDCVProviderConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SectigoDCVProviderConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SectigoDCVProviderConfigResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *SectigoDCVProviderConfigResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SectigoDCVProviderConfigResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SectigoDCVProviderConfigResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDcvMethod

`func (o *SectigoDCVProviderConfigResponse) GetDcvMethod() string`

GetDcvMethod returns the DcvMethod field if non-nil, zero value otherwise.

### GetDcvMethodOk

`func (o *SectigoDCVProviderConfigResponse) GetDcvMethodOk() (*string, bool)`

GetDcvMethodOk returns a tuple with the DcvMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcvMethod

`func (o *SectigoDCVProviderConfigResponse) SetDcvMethod(v string)`

SetDcvMethod sets DcvMethod field to given value.


### GetEndpoint

`func (o *SectigoDCVProviderConfigResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SectigoDCVProviderConfigResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SectigoDCVProviderConfigResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *SectigoDCVProviderConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectigoDCVProviderConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectigoDCVProviderConfigResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOauthTokenEndpoint

`func (o *SectigoDCVProviderConfigResponse) GetOauthTokenEndpoint() string`

GetOauthTokenEndpoint returns the OauthTokenEndpoint field if non-nil, zero value otherwise.

### GetOauthTokenEndpointOk

`func (o *SectigoDCVProviderConfigResponse) GetOauthTokenEndpointOk() (*string, bool)`

GetOauthTokenEndpointOk returns a tuple with the OauthTokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenEndpoint

`func (o *SectigoDCVProviderConfigResponse) SetOauthTokenEndpoint(v string)`

SetOauthTokenEndpoint sets OauthTokenEndpoint field to given value.


### GetOrganizationId

`func (o *SectigoDCVProviderConfigResponse) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SectigoDCVProviderConfigResponse) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SectigoDCVProviderConfigResponse) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SectigoDCVProviderConfigResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProxy

`func (o *SectigoDCVProviderConfigResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SectigoDCVProviderConfigResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SectigoDCVProviderConfigResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SectigoDCVProviderConfigResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTimeout

`func (o *SectigoDCVProviderConfigResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SectigoDCVProviderConfigResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SectigoDCVProviderConfigResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *SectigoDCVProviderConfigResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SectigoDCVProviderConfigResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *SectigoDCVProviderConfigResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SectigoDCVProviderConfigResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SectigoDCVProviderConfigResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


