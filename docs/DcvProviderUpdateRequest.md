# DcvProviderUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **string** | Name of the login/password credentials configuration holding the GlobalSign MSSL API credentials | 
**Endpoint** | **string** | GlobalSign Managed SSL (MSSL) API endpoint URL | 
**Name** | **string** | Unique name of the DCV provider configuration | 
**Proxy** | Pointer to **string** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | **NullableString** | Request timeout | 
**Type** | **string** | Provider type discriminator | 
**DefaultEmail** | **string** | Default contact email used when renewing a domain | 
**DefaultPhone** | **string** | Default contact phone number used when renewing a domain | 
**Profile** | **string** | GlobalSign MSSL profile ID the domains belong to | 

## Methods

### NewDcvProviderUpdateRequest

`func NewDcvProviderUpdateRequest(credentials string, endpoint string, name string, timeout NullableString, type_ string, defaultEmail string, defaultPhone string, profile string, ) *DcvProviderUpdateRequest`

NewDcvProviderUpdateRequest instantiates a new DcvProviderUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcvProviderUpdateRequestWithDefaults

`func NewDcvProviderUpdateRequestWithDefaults() *DcvProviderUpdateRequest`

NewDcvProviderUpdateRequestWithDefaults instantiates a new DcvProviderUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *DcvProviderUpdateRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DcvProviderUpdateRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DcvProviderUpdateRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetEndpoint

`func (o *DcvProviderUpdateRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DcvProviderUpdateRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DcvProviderUpdateRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *DcvProviderUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcvProviderUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcvProviderUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *DcvProviderUpdateRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DcvProviderUpdateRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DcvProviderUpdateRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DcvProviderUpdateRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTimeout

`func (o *DcvProviderUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DcvProviderUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DcvProviderUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *DcvProviderUpdateRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DcvProviderUpdateRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *DcvProviderUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcvProviderUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcvProviderUpdateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDefaultEmail

`func (o *DcvProviderUpdateRequest) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *DcvProviderUpdateRequest) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *DcvProviderUpdateRequest) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.


### GetDefaultPhone

`func (o *DcvProviderUpdateRequest) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *DcvProviderUpdateRequest) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *DcvProviderUpdateRequest) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.


### GetProfile

`func (o *DcvProviderUpdateRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *DcvProviderUpdateRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *DcvProviderUpdateRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


