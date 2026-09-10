# DigicertDCVProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **string** | Name of the credentials configuration holding the DigiCert API key | 
**Endpoint** | **string** | DigiCert API endpoint URL | 
**Name** | **string** | Unique name of the DCV provider configuration | 
**Proxy** | Pointer to **NullableString** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | Pointer to **NullableString** | Request timeout | [optional] 
**Type** | **string** | Provider type discriminator | 

## Methods

### NewDigicertDCVProviderConfig

`func NewDigicertDCVProviderConfig(credentials string, endpoint string, name string, type_ string, ) *DigicertDCVProviderConfig`

NewDigicertDCVProviderConfig instantiates a new DigicertDCVProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigicertDCVProviderConfigWithDefaults

`func NewDigicertDCVProviderConfigWithDefaults() *DigicertDCVProviderConfig`

NewDigicertDCVProviderConfigWithDefaults instantiates a new DigicertDCVProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *DigicertDCVProviderConfig) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DigicertDCVProviderConfig) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DigicertDCVProviderConfig) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetEndpoint

`func (o *DigicertDCVProviderConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DigicertDCVProviderConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DigicertDCVProviderConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *DigicertDCVProviderConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DigicertDCVProviderConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DigicertDCVProviderConfig) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *DigicertDCVProviderConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DigicertDCVProviderConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DigicertDCVProviderConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DigicertDCVProviderConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DigicertDCVProviderConfig) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DigicertDCVProviderConfig) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *DigicertDCVProviderConfig) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DigicertDCVProviderConfig) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DigicertDCVProviderConfig) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DigicertDCVProviderConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *DigicertDCVProviderConfig) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DigicertDCVProviderConfig) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *DigicertDCVProviderConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigicertDCVProviderConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigicertDCVProviderConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


