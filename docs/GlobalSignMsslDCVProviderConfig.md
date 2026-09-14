# GlobalSignMsslDCVProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewGlobalSignMsslDCVProviderConfig

`func NewGlobalSignMsslDCVProviderConfig(credentials string, defaultEmail string, defaultPhone string, endpoint string, name string, profile string, timeout NullableString, type_ string, ) *GlobalSignMsslDCVProviderConfig`

NewGlobalSignMsslDCVProviderConfig instantiates a new GlobalSignMsslDCVProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSignMsslDCVProviderConfigWithDefaults

`func NewGlobalSignMsslDCVProviderConfigWithDefaults() *GlobalSignMsslDCVProviderConfig`

NewGlobalSignMsslDCVProviderConfigWithDefaults instantiates a new GlobalSignMsslDCVProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GlobalSignMsslDCVProviderConfig) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GlobalSignMsslDCVProviderConfig) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GlobalSignMsslDCVProviderConfig) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDefaultEmail

`func (o *GlobalSignMsslDCVProviderConfig) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *GlobalSignMsslDCVProviderConfig) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *GlobalSignMsslDCVProviderConfig) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.


### GetDefaultPhone

`func (o *GlobalSignMsslDCVProviderConfig) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *GlobalSignMsslDCVProviderConfig) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *GlobalSignMsslDCVProviderConfig) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.


### GetEndpoint

`func (o *GlobalSignMsslDCVProviderConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GlobalSignMsslDCVProviderConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GlobalSignMsslDCVProviderConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *GlobalSignMsslDCVProviderConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalSignMsslDCVProviderConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalSignMsslDCVProviderConfig) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *GlobalSignMsslDCVProviderConfig) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GlobalSignMsslDCVProviderConfig) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GlobalSignMsslDCVProviderConfig) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProxy

`func (o *GlobalSignMsslDCVProviderConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *GlobalSignMsslDCVProviderConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *GlobalSignMsslDCVProviderConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *GlobalSignMsslDCVProviderConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTimeout

`func (o *GlobalSignMsslDCVProviderConfig) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GlobalSignMsslDCVProviderConfig) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GlobalSignMsslDCVProviderConfig) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *GlobalSignMsslDCVProviderConfig) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *GlobalSignMsslDCVProviderConfig) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *GlobalSignMsslDCVProviderConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalSignMsslDCVProviderConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalSignMsslDCVProviderConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


