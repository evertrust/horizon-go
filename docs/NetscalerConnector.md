# NetscalerConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateStorePath** | **string** |  | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the account to authenticate on Netscaler | 
**Hostname** | **string** |  | 
**MaxStoredCertificatePerHolder** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**Type** | **string** |  | 

## Methods

### NewNetscalerConnector

`func NewNetscalerConnector(certificateStorePath string, credentials string, hostname string, name string, throttleDuration string, throttleParallelism int64, type_ string, ) *NetscalerConnector`

NewNetscalerConnector instantiates a new NetscalerConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetscalerConnectorWithDefaults

`func NewNetscalerConnectorWithDefaults() *NetscalerConnector`

NewNetscalerConnectorWithDefaults instantiates a new NetscalerConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateStorePath

`func (o *NetscalerConnector) GetCertificateStorePath() string`

GetCertificateStorePath returns the CertificateStorePath field if non-nil, zero value otherwise.

### GetCertificateStorePathOk

`func (o *NetscalerConnector) GetCertificateStorePathOk() (*string, bool)`

GetCertificateStorePathOk returns a tuple with the CertificateStorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStorePath

`func (o *NetscalerConnector) SetCertificateStorePath(v string)`

SetCertificateStorePath sets CertificateStorePath field to given value.


### GetCredentials

`func (o *NetscalerConnector) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *NetscalerConnector) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *NetscalerConnector) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *NetscalerConnector) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NetscalerConnector) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NetscalerConnector) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetMaxStoredCertificatePerHolder

`func (o *NetscalerConnector) GetMaxStoredCertificatePerHolder() int64`

GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field if non-nil, zero value otherwise.

### GetMaxStoredCertificatePerHolderOk

`func (o *NetscalerConnector) GetMaxStoredCertificatePerHolderOk() (*int64, bool)`

GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredCertificatePerHolder

`func (o *NetscalerConnector) SetMaxStoredCertificatePerHolder(v int64)`

SetMaxStoredCertificatePerHolder sets MaxStoredCertificatePerHolder field to given value.

### HasMaxStoredCertificatePerHolder

`func (o *NetscalerConnector) HasMaxStoredCertificatePerHolder() bool`

HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.

### SetMaxStoredCertificatePerHolderNil

`func (o *NetscalerConnector) SetMaxStoredCertificatePerHolderNil(b bool)`

 SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil

### UnsetMaxStoredCertificatePerHolder
`func (o *NetscalerConnector) UnsetMaxStoredCertificatePerHolder()`

UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
### GetName

`func (o *NetscalerConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetscalerConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetscalerConnector) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *NetscalerConnector) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *NetscalerConnector) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *NetscalerConnector) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *NetscalerConnector) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *NetscalerConnector) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *NetscalerConnector) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetProxy

`func (o *NetscalerConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *NetscalerConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *NetscalerConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *NetscalerConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *NetscalerConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *NetscalerConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRenewalPeriod

`func (o *NetscalerConnector) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *NetscalerConnector) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *NetscalerConnector) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *NetscalerConnector) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *NetscalerConnector) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *NetscalerConnector) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetThrottleDuration

`func (o *NetscalerConnector) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *NetscalerConnector) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *NetscalerConnector) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *NetscalerConnector) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *NetscalerConnector) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *NetscalerConnector) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *NetscalerConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *NetscalerConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *NetscalerConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *NetscalerConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *NetscalerConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *NetscalerConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTlsInsecure

`func (o *NetscalerConnector) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *NetscalerConnector) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *NetscalerConnector) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *NetscalerConnector) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *NetscalerConnector) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *NetscalerConnector) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetType

`func (o *NetscalerConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetscalerConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetscalerConnector) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


