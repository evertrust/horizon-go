# F5ClientConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BigIPHostname** | **string** |  | 
**CipherGroup** | Pointer to **NullableString** |  | [optional] 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the account to authenticate on F5 | 
**LoginProvider** | Pointer to **NullableString** | Name of the F5 BIG-IP authentication provider to use for login (e.g. &#x60;tmos&#x60;). Defaults to the device&#39;s default provider when unset. | [optional] 
**MaxStoredCertificatePerHolder** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**OverrideProfileConfiguration** | Pointer to **NullableBool** | Whether to override the existing SSL profile&#39;s parent profile and cipher group on update. When &#x60;false&#x60;, only the certificate and key are updated. | [optional] [default to true]
**Partition** | Pointer to **NullableString** |  | [optional] [default to "Common"]
**PersistConfiguration** | Pointer to **bool** | When enabled, Horizon saves the F5 running configuration to &#x60;bigip.conf&#x60; after each successful deployment so that pushed changes survive an appliance reboot. Requires an admin-level F5 technical account. | [optional] [default to false]
**Prefix** | Pointer to **NullableString** |  | [optional] [default to "hrz-"]
**Proxy** | Pointer to **NullableString** |  | [optional] 
**SslParent** | Pointer to **NullableString** |  | [optional] [default to "clientssl"]
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**Type** | **string** |  | 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewF5ClientConnector

`func NewF5ClientConnector(bigIPHostname string, credentials string, name string, throttleDuration string, throttleParallelism int64, type_ string, ) *F5ClientConnector`

NewF5ClientConnector instantiates a new F5ClientConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewF5ClientConnectorWithDefaults

`func NewF5ClientConnectorWithDefaults() *F5ClientConnector`

NewF5ClientConnectorWithDefaults instantiates a new F5ClientConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBigIPHostname

`func (o *F5ClientConnector) GetBigIPHostname() string`

GetBigIPHostname returns the BigIPHostname field if non-nil, zero value otherwise.

### GetBigIPHostnameOk

`func (o *F5ClientConnector) GetBigIPHostnameOk() (*string, bool)`

GetBigIPHostnameOk returns a tuple with the BigIPHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigIPHostname

`func (o *F5ClientConnector) SetBigIPHostname(v string)`

SetBigIPHostname sets BigIPHostname field to given value.


### GetCipherGroup

`func (o *F5ClientConnector) GetCipherGroup() string`

GetCipherGroup returns the CipherGroup field if non-nil, zero value otherwise.

### GetCipherGroupOk

`func (o *F5ClientConnector) GetCipherGroupOk() (*string, bool)`

GetCipherGroupOk returns a tuple with the CipherGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherGroup

`func (o *F5ClientConnector) SetCipherGroup(v string)`

SetCipherGroup sets CipherGroup field to given value.

### HasCipherGroup

`func (o *F5ClientConnector) HasCipherGroup() bool`

HasCipherGroup returns a boolean if a field has been set.

### SetCipherGroupNil

`func (o *F5ClientConnector) SetCipherGroupNil(b bool)`

 SetCipherGroupNil sets the value for CipherGroup to be an explicit nil

### UnsetCipherGroup
`func (o *F5ClientConnector) UnsetCipherGroup()`

UnsetCipherGroup ensures that no value is present for CipherGroup, not even an explicit nil
### GetCredentials

`func (o *F5ClientConnector) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *F5ClientConnector) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *F5ClientConnector) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetLoginProvider

`func (o *F5ClientConnector) GetLoginProvider() string`

GetLoginProvider returns the LoginProvider field if non-nil, zero value otherwise.

### GetLoginProviderOk

`func (o *F5ClientConnector) GetLoginProviderOk() (*string, bool)`

GetLoginProviderOk returns a tuple with the LoginProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginProvider

`func (o *F5ClientConnector) SetLoginProvider(v string)`

SetLoginProvider sets LoginProvider field to given value.

### HasLoginProvider

`func (o *F5ClientConnector) HasLoginProvider() bool`

HasLoginProvider returns a boolean if a field has been set.

### SetLoginProviderNil

`func (o *F5ClientConnector) SetLoginProviderNil(b bool)`

 SetLoginProviderNil sets the value for LoginProvider to be an explicit nil

### UnsetLoginProvider
`func (o *F5ClientConnector) UnsetLoginProvider()`

UnsetLoginProvider ensures that no value is present for LoginProvider, not even an explicit nil
### GetMaxStoredCertificatePerHolder

`func (o *F5ClientConnector) GetMaxStoredCertificatePerHolder() int64`

GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field if non-nil, zero value otherwise.

### GetMaxStoredCertificatePerHolderOk

`func (o *F5ClientConnector) GetMaxStoredCertificatePerHolderOk() (*int64, bool)`

GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredCertificatePerHolder

`func (o *F5ClientConnector) SetMaxStoredCertificatePerHolder(v int64)`

SetMaxStoredCertificatePerHolder sets MaxStoredCertificatePerHolder field to given value.

### HasMaxStoredCertificatePerHolder

`func (o *F5ClientConnector) HasMaxStoredCertificatePerHolder() bool`

HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.

### SetMaxStoredCertificatePerHolderNil

`func (o *F5ClientConnector) SetMaxStoredCertificatePerHolderNil(b bool)`

 SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil

### UnsetMaxStoredCertificatePerHolder
`func (o *F5ClientConnector) UnsetMaxStoredCertificatePerHolder()`

UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
### GetName

`func (o *F5ClientConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *F5ClientConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *F5ClientConnector) SetName(v string)`

SetName sets Name field to given value.


### GetOverrideProfileConfiguration

`func (o *F5ClientConnector) GetOverrideProfileConfiguration() bool`

GetOverrideProfileConfiguration returns the OverrideProfileConfiguration field if non-nil, zero value otherwise.

### GetOverrideProfileConfigurationOk

`func (o *F5ClientConnector) GetOverrideProfileConfigurationOk() (*bool, bool)`

GetOverrideProfileConfigurationOk returns a tuple with the OverrideProfileConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideProfileConfiguration

`func (o *F5ClientConnector) SetOverrideProfileConfiguration(v bool)`

SetOverrideProfileConfiguration sets OverrideProfileConfiguration field to given value.

### HasOverrideProfileConfiguration

`func (o *F5ClientConnector) HasOverrideProfileConfiguration() bool`

HasOverrideProfileConfiguration returns a boolean if a field has been set.

### SetOverrideProfileConfigurationNil

`func (o *F5ClientConnector) SetOverrideProfileConfigurationNil(b bool)`

 SetOverrideProfileConfigurationNil sets the value for OverrideProfileConfiguration to be an explicit nil

### UnsetOverrideProfileConfiguration
`func (o *F5ClientConnector) UnsetOverrideProfileConfiguration()`

UnsetOverrideProfileConfiguration ensures that no value is present for OverrideProfileConfiguration, not even an explicit nil
### GetPartition

`func (o *F5ClientConnector) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *F5ClientConnector) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *F5ClientConnector) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *F5ClientConnector) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### SetPartitionNil

`func (o *F5ClientConnector) SetPartitionNil(b bool)`

 SetPartitionNil sets the value for Partition to be an explicit nil

### UnsetPartition
`func (o *F5ClientConnector) UnsetPartition()`

UnsetPartition ensures that no value is present for Partition, not even an explicit nil
### GetPersistConfiguration

`func (o *F5ClientConnector) GetPersistConfiguration() bool`

GetPersistConfiguration returns the PersistConfiguration field if non-nil, zero value otherwise.

### GetPersistConfigurationOk

`func (o *F5ClientConnector) GetPersistConfigurationOk() (*bool, bool)`

GetPersistConfigurationOk returns a tuple with the PersistConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistConfiguration

`func (o *F5ClientConnector) SetPersistConfiguration(v bool)`

SetPersistConfiguration sets PersistConfiguration field to given value.

### HasPersistConfiguration

`func (o *F5ClientConnector) HasPersistConfiguration() bool`

HasPersistConfiguration returns a boolean if a field has been set.

### GetPrefix

`func (o *F5ClientConnector) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *F5ClientConnector) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *F5ClientConnector) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *F5ClientConnector) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *F5ClientConnector) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *F5ClientConnector) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetProxy

`func (o *F5ClientConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *F5ClientConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *F5ClientConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *F5ClientConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *F5ClientConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *F5ClientConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetSslParent

`func (o *F5ClientConnector) GetSslParent() string`

GetSslParent returns the SslParent field if non-nil, zero value otherwise.

### GetSslParentOk

`func (o *F5ClientConnector) GetSslParentOk() (*string, bool)`

GetSslParentOk returns a tuple with the SslParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslParent

`func (o *F5ClientConnector) SetSslParent(v string)`

SetSslParent sets SslParent field to given value.

### HasSslParent

`func (o *F5ClientConnector) HasSslParent() bool`

HasSslParent returns a boolean if a field has been set.

### SetSslParentNil

`func (o *F5ClientConnector) SetSslParentNil(b bool)`

 SetSslParentNil sets the value for SslParent to be an explicit nil

### UnsetSslParent
`func (o *F5ClientConnector) UnsetSslParent()`

UnsetSslParent ensures that no value is present for SslParent, not even an explicit nil
### GetThrottleDuration

`func (o *F5ClientConnector) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *F5ClientConnector) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *F5ClientConnector) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *F5ClientConnector) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *F5ClientConnector) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *F5ClientConnector) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *F5ClientConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *F5ClientConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *F5ClientConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *F5ClientConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *F5ClientConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *F5ClientConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTlsInsecure

`func (o *F5ClientConnector) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *F5ClientConnector) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *F5ClientConnector) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *F5ClientConnector) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *F5ClientConnector) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *F5ClientConnector) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetType

`func (o *F5ClientConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *F5ClientConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *F5ClientConnector) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *F5ClientConnector) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *F5ClientConnector) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *F5ClientConnector) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *F5ClientConnector) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *F5ClientConnector) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *F5ClientConnector) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


