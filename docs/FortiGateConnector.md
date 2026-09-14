# FortiGateConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewFortiGateConnector

`func NewFortiGateConnector(credentials string, hostname string, name string, prefix string, throttleDuration string, throttleParallelism int64, timeout string, type_ string, ) *FortiGateConnector`

NewFortiGateConnector instantiates a new FortiGateConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFortiGateConnectorWithDefaults

`func NewFortiGateConnectorWithDefaults() *FortiGateConnector`

NewFortiGateConnectorWithDefaults instantiates a new FortiGateConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateCredentials

`func (o *FortiGateConnector) GetCertificateCredentials() string`

GetCertificateCredentials returns the CertificateCredentials field if non-nil, zero value otherwise.

### GetCertificateCredentialsOk

`func (o *FortiGateConnector) GetCertificateCredentialsOk() (*string, bool)`

GetCertificateCredentialsOk returns a tuple with the CertificateCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCredentials

`func (o *FortiGateConnector) SetCertificateCredentials(v string)`

SetCertificateCredentials sets CertificateCredentials field to given value.

### HasCertificateCredentials

`func (o *FortiGateConnector) HasCertificateCredentials() bool`

HasCertificateCredentials returns a boolean if a field has been set.

### SetCertificateCredentialsNil

`func (o *FortiGateConnector) SetCertificateCredentialsNil(b bool)`

 SetCertificateCredentialsNil sets the value for CertificateCredentials to be an explicit nil

### UnsetCertificateCredentials
`func (o *FortiGateConnector) UnsetCertificateCredentials()`

UnsetCertificateCredentials ensures that no value is present for CertificateCredentials, not even an explicit nil
### GetCredentials

`func (o *FortiGateConnector) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *FortiGateConnector) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *FortiGateConnector) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *FortiGateConnector) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FortiGateConnector) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FortiGateConnector) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetName

`func (o *FortiGateConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FortiGateConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FortiGateConnector) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *FortiGateConnector) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *FortiGateConnector) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *FortiGateConnector) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetProxy

`func (o *FortiGateConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *FortiGateConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *FortiGateConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *FortiGateConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *FortiGateConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *FortiGateConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetThrottleDuration

`func (o *FortiGateConnector) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *FortiGateConnector) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *FortiGateConnector) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *FortiGateConnector) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *FortiGateConnector) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *FortiGateConnector) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *FortiGateConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FortiGateConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FortiGateConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTlsInsecure

`func (o *FortiGateConnector) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *FortiGateConnector) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *FortiGateConnector) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *FortiGateConnector) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *FortiGateConnector) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *FortiGateConnector) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetType

`func (o *FortiGateConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FortiGateConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FortiGateConnector) SetType(v string)`

SetType sets Type field to given value.


### GetVdom

`func (o *FortiGateConnector) GetVdom() string`

GetVdom returns the Vdom field if non-nil, zero value otherwise.

### GetVdomOk

`func (o *FortiGateConnector) GetVdomOk() (*string, bool)`

GetVdomOk returns a tuple with the Vdom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdom

`func (o *FortiGateConnector) SetVdom(v string)`

SetVdom sets Vdom field to given value.

### HasVdom

`func (o *FortiGateConnector) HasVdom() bool`

HasVdom returns a boolean if a field has been set.

### SetVdomNil

`func (o *FortiGateConnector) SetVdomNil(b bool)`

 SetVdomNil sets the value for Vdom to be an explicit nil

### UnsetVdom
`func (o *FortiGateConnector) UnsetVdom()`

UnsetVdom ensures that no value is present for Vdom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


