# DcvProviderList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the login/password credentials configuration holding the SCM API client, with the OAuth client id as login and the client secret as password | 
**Endpoint** | **string** | Sectigo Certificate Manager (SCM) API base URL | 
**Name** | **string** | Unique name of the DCV provider configuration | 
**Proxy** | Pointer to **string** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | **NullableString** | Request timeout | 
**Type** | **string** | Provider type discriminator | 
**DefaultEmail** | **string** | Default contact email used when renewing a domain | 
**DefaultPhone** | **string** | Default contact phone number used when renewing a domain | 
**Profile** | **string** | GlobalSign MSSL profile ID the domains belong to | 
**DcvMethod** | **string** | DNS method used to validate a domain. It is a fallback: a domain that already holds a CNAME or TXT validation is re-validated with its own method, and this method applies only to a domain that has never been validated or whose existing validation uses a method that cannot be published over DNS. | 
**OauthTokenEndpoint** | **string** | OAuth token endpoint used to obtain a bearer token for the SCM API. Defaults to Sectigo&#39;s SSO realm. | [default to "https://auth.sso.sectigo.com/auth/realms/apiclients/protocol/openid-connect/token"]
**OrganizationId** | Pointer to **int64** | Restricts the domain listing to this Sectigo organization or department. When unset, every domain of the customer account is listed. | [optional] 

## Methods

### NewDcvProviderList200ResponseInner

`func NewDcvProviderList200ResponseInner(id string, credentials string, endpoint string, name string, timeout NullableString, type_ string, defaultEmail string, defaultPhone string, profile string, dcvMethod string, oauthTokenEndpoint string, ) *DcvProviderList200ResponseInner`

NewDcvProviderList200ResponseInner instantiates a new DcvProviderList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcvProviderList200ResponseInnerWithDefaults

`func NewDcvProviderList200ResponseInnerWithDefaults() *DcvProviderList200ResponseInner`

NewDcvProviderList200ResponseInnerWithDefaults instantiates a new DcvProviderList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DcvProviderList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DcvProviderList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DcvProviderList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *DcvProviderList200ResponseInner) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DcvProviderList200ResponseInner) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DcvProviderList200ResponseInner) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetEndpoint

`func (o *DcvProviderList200ResponseInner) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DcvProviderList200ResponseInner) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DcvProviderList200ResponseInner) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *DcvProviderList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcvProviderList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcvProviderList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *DcvProviderList200ResponseInner) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DcvProviderList200ResponseInner) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DcvProviderList200ResponseInner) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DcvProviderList200ResponseInner) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTimeout

`func (o *DcvProviderList200ResponseInner) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DcvProviderList200ResponseInner) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DcvProviderList200ResponseInner) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *DcvProviderList200ResponseInner) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DcvProviderList200ResponseInner) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *DcvProviderList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcvProviderList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcvProviderList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetDefaultEmail

`func (o *DcvProviderList200ResponseInner) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *DcvProviderList200ResponseInner) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *DcvProviderList200ResponseInner) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.


### GetDefaultPhone

`func (o *DcvProviderList200ResponseInner) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *DcvProviderList200ResponseInner) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *DcvProviderList200ResponseInner) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.


### GetProfile

`func (o *DcvProviderList200ResponseInner) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *DcvProviderList200ResponseInner) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *DcvProviderList200ResponseInner) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetDcvMethod

`func (o *DcvProviderList200ResponseInner) GetDcvMethod() string`

GetDcvMethod returns the DcvMethod field if non-nil, zero value otherwise.

### GetDcvMethodOk

`func (o *DcvProviderList200ResponseInner) GetDcvMethodOk() (*string, bool)`

GetDcvMethodOk returns a tuple with the DcvMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcvMethod

`func (o *DcvProviderList200ResponseInner) SetDcvMethod(v string)`

SetDcvMethod sets DcvMethod field to given value.


### GetOauthTokenEndpoint

`func (o *DcvProviderList200ResponseInner) GetOauthTokenEndpoint() string`

GetOauthTokenEndpoint returns the OauthTokenEndpoint field if non-nil, zero value otherwise.

### GetOauthTokenEndpointOk

`func (o *DcvProviderList200ResponseInner) GetOauthTokenEndpointOk() (*string, bool)`

GetOauthTokenEndpointOk returns a tuple with the OauthTokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenEndpoint

`func (o *DcvProviderList200ResponseInner) SetOauthTokenEndpoint(v string)`

SetOauthTokenEndpoint sets OauthTokenEndpoint field to given value.


### GetOrganizationId

`func (o *DcvProviderList200ResponseInner) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DcvProviderList200ResponseInner) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DcvProviderList200ResponseInner) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DcvProviderList200ResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


