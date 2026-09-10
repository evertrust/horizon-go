# DigicertDCVProviderConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the credentials configuration holding the DigiCert API key | 
**Endpoint** | **string** | DigiCert API endpoint URL | 
**Name** | **string** | Unique name of the DCV provider configuration | 
**Proxy** | Pointer to **NullableString** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | Pointer to **NullableString** | Request timeout | [optional] 
**Type** | **string** | Provider type discriminator | 

## Methods

### NewDigicertDCVProviderConfigResponse

`func NewDigicertDCVProviderConfigResponse(id string, credentials string, endpoint string, name string, type_ string, ) *DigicertDCVProviderConfigResponse`

NewDigicertDCVProviderConfigResponse instantiates a new DigicertDCVProviderConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigicertDCVProviderConfigResponseWithDefaults

`func NewDigicertDCVProviderConfigResponseWithDefaults() *DigicertDCVProviderConfigResponse`

NewDigicertDCVProviderConfigResponseWithDefaults instantiates a new DigicertDCVProviderConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DigicertDCVProviderConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DigicertDCVProviderConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DigicertDCVProviderConfigResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *DigicertDCVProviderConfigResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DigicertDCVProviderConfigResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DigicertDCVProviderConfigResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetEndpoint

`func (o *DigicertDCVProviderConfigResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DigicertDCVProviderConfigResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DigicertDCVProviderConfigResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *DigicertDCVProviderConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DigicertDCVProviderConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DigicertDCVProviderConfigResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *DigicertDCVProviderConfigResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DigicertDCVProviderConfigResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DigicertDCVProviderConfigResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DigicertDCVProviderConfigResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DigicertDCVProviderConfigResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DigicertDCVProviderConfigResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *DigicertDCVProviderConfigResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DigicertDCVProviderConfigResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DigicertDCVProviderConfigResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DigicertDCVProviderConfigResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *DigicertDCVProviderConfigResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DigicertDCVProviderConfigResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *DigicertDCVProviderConfigResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigicertDCVProviderConfigResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigicertDCVProviderConfigResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


