# GlobalSignMsslDCVProviderConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the login/password credentials configuration holding the GlobalSign MSSL API credentials | 
**DefaultEmail** | **string** | Default contact email used when renewing a domain | 
**DefaultPhone** | **string** | Default contact phone number used when renewing a domain | 
**Endpoint** | **string** | GlobalSign Managed SSL (MSSL) API endpoint URL | 
**Name** | **string** | Unique name of the DCV provider configuration | 
**Profile** | **string** | GlobalSign MSSL profile ID the domains belong to | 
**Proxy** | Pointer to **string** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | **NullableString** | Request timeout | 
**Type** | **string** | Provider type discriminator | 

## Methods

### NewGlobalSignMsslDCVProviderConfigResponse

`func NewGlobalSignMsslDCVProviderConfigResponse(id string, credentials string, defaultEmail string, defaultPhone string, endpoint string, name string, profile string, timeout NullableString, type_ string, ) *GlobalSignMsslDCVProviderConfigResponse`

NewGlobalSignMsslDCVProviderConfigResponse instantiates a new GlobalSignMsslDCVProviderConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSignMsslDCVProviderConfigResponseWithDefaults

`func NewGlobalSignMsslDCVProviderConfigResponseWithDefaults() *GlobalSignMsslDCVProviderConfigResponse`

NewGlobalSignMsslDCVProviderConfigResponseWithDefaults instantiates a new GlobalSignMsslDCVProviderConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDefaultEmail

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.


### GetDefaultPhone

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.


### GetEndpoint

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProxy

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *GlobalSignMsslDCVProviderConfigResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTimeout

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *GlobalSignMsslDCVProviderConfigResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalSignMsslDCVProviderConfigResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalSignMsslDCVProviderConfigResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


