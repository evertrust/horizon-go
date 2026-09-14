# FortiGateConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**CertificateCredentials** | Pointer to **NullableString** | Optional name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) used for mutual TLS in addition to the API key | [optional] 
**Credentials** | **string** | Name of the &#x60;raw&#x60; [credentials](#tag/security.credentials) holding the FortiGate REST API key | 
**Hostname** | **string** | The hostname or URL of the FortiGate appliance | 
**Name** | **string** |  | 
**Prefix** | **string** | Certificate name prefix used when deploying certificates | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**Timeout** | **string** |  | 
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**Type** | **string** |  | 
**Vdom** | Pointer to **NullableString** | Virtual domain to deploy to; when absent the certificate is imported in the global scope | [optional] 

## Methods

### NewFortiGateConnectorResponse

`func NewFortiGateConnectorResponse(id string, credentials string, hostname string, name string, prefix string, throttleDuration string, throttleParallelism int64, timeout string, type_ string, ) *FortiGateConnectorResponse`

NewFortiGateConnectorResponse instantiates a new FortiGateConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFortiGateConnectorResponseWithDefaults

`func NewFortiGateConnectorResponseWithDefaults() *FortiGateConnectorResponse`

NewFortiGateConnectorResponseWithDefaults instantiates a new FortiGateConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FortiGateConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FortiGateConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FortiGateConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCertificateCredentials

`func (o *FortiGateConnectorResponse) GetCertificateCredentials() string`

GetCertificateCredentials returns the CertificateCredentials field if non-nil, zero value otherwise.

### GetCertificateCredentialsOk

`func (o *FortiGateConnectorResponse) GetCertificateCredentialsOk() (*string, bool)`

GetCertificateCredentialsOk returns a tuple with the CertificateCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCredentials

`func (o *FortiGateConnectorResponse) SetCertificateCredentials(v string)`

SetCertificateCredentials sets CertificateCredentials field to given value.

### HasCertificateCredentials

`func (o *FortiGateConnectorResponse) HasCertificateCredentials() bool`

HasCertificateCredentials returns a boolean if a field has been set.

### SetCertificateCredentialsNil

`func (o *FortiGateConnectorResponse) SetCertificateCredentialsNil(b bool)`

 SetCertificateCredentialsNil sets the value for CertificateCredentials to be an explicit nil

### UnsetCertificateCredentials
`func (o *FortiGateConnectorResponse) UnsetCertificateCredentials()`

UnsetCertificateCredentials ensures that no value is present for CertificateCredentials, not even an explicit nil
### GetCredentials

`func (o *FortiGateConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *FortiGateConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *FortiGateConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *FortiGateConnectorResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FortiGateConnectorResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FortiGateConnectorResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetName

`func (o *FortiGateConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FortiGateConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FortiGateConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *FortiGateConnectorResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *FortiGateConnectorResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *FortiGateConnectorResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetProxy

`func (o *FortiGateConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *FortiGateConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *FortiGateConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *FortiGateConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *FortiGateConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *FortiGateConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetThrottleDuration

`func (o *FortiGateConnectorResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *FortiGateConnectorResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *FortiGateConnectorResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *FortiGateConnectorResponse) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *FortiGateConnectorResponse) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *FortiGateConnectorResponse) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *FortiGateConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FortiGateConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FortiGateConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTlsInsecure

`func (o *FortiGateConnectorResponse) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *FortiGateConnectorResponse) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *FortiGateConnectorResponse) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *FortiGateConnectorResponse) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *FortiGateConnectorResponse) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *FortiGateConnectorResponse) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetType

`func (o *FortiGateConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FortiGateConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FortiGateConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetVdom

`func (o *FortiGateConnectorResponse) GetVdom() string`

GetVdom returns the Vdom field if non-nil, zero value otherwise.

### GetVdomOk

`func (o *FortiGateConnectorResponse) GetVdomOk() (*string, bool)`

GetVdomOk returns a tuple with the Vdom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdom

`func (o *FortiGateConnectorResponse) SetVdom(v string)`

SetVdom sets Vdom field to given value.

### HasVdom

`func (o *FortiGateConnectorResponse) HasVdom() bool`

HasVdom returns a boolean if a field has been set.

### SetVdomNil

`func (o *FortiGateConnectorResponse) SetVdomNil(b bool)`

 SetVdomNil sets the value for Vdom to be an explicit nil

### UnsetVdom
`func (o *FortiGateConnectorResponse) UnsetVdom()`

UnsetVdom ensures that no value is present for Vdom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


