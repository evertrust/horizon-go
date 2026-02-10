# IntunePKCSConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**MaxStoredCertificatePerHolder** | Pointer to **NullableInt64** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Tenant** | **string** |  | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the App ID and Key to authenticate on Intune PKCS | 
**PubKey** | **string** |  | 
**KeyName** | **string** |  | 
**ProviderName** | Pointer to **NullableString** |  | [optional] 
**IntendedPurpose** | Pointer to **NullableString** |  | [optional] 
**SearchFilter** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIntunePKCSConnector

`func NewIntunePKCSConnector(type_ string, name string, throttleDuration string, throttleParallelism int64, tenant string, credentials string, pubKey string, keyName string, ) *IntunePKCSConnector`

NewIntunePKCSConnector instantiates a new IntunePKCSConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntunePKCSConnectorWithDefaults

`func NewIntunePKCSConnectorWithDefaults() *IntunePKCSConnector`

NewIntunePKCSConnectorWithDefaults instantiates a new IntunePKCSConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IntunePKCSConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntunePKCSConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntunePKCSConnector) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *IntunePKCSConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntunePKCSConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntunePKCSConnector) SetName(v string)`

SetName sets Name field to given value.


### GetThrottleDuration

`func (o *IntunePKCSConnector) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *IntunePKCSConnector) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *IntunePKCSConnector) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *IntunePKCSConnector) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *IntunePKCSConnector) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *IntunePKCSConnector) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetRenewalPeriod

`func (o *IntunePKCSConnector) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *IntunePKCSConnector) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *IntunePKCSConnector) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *IntunePKCSConnector) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *IntunePKCSConnector) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *IntunePKCSConnector) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetTimeout

`func (o *IntunePKCSConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *IntunePKCSConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *IntunePKCSConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *IntunePKCSConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *IntunePKCSConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *IntunePKCSConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetMaxStoredCertificatePerHolder

`func (o *IntunePKCSConnector) GetMaxStoredCertificatePerHolder() int64`

GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field if non-nil, zero value otherwise.

### GetMaxStoredCertificatePerHolderOk

`func (o *IntunePKCSConnector) GetMaxStoredCertificatePerHolderOk() (*int64, bool)`

GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredCertificatePerHolder

`func (o *IntunePKCSConnector) SetMaxStoredCertificatePerHolder(v int64)`

SetMaxStoredCertificatePerHolder sets MaxStoredCertificatePerHolder field to given value.

### HasMaxStoredCertificatePerHolder

`func (o *IntunePKCSConnector) HasMaxStoredCertificatePerHolder() bool`

HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.

### SetMaxStoredCertificatePerHolderNil

`func (o *IntunePKCSConnector) SetMaxStoredCertificatePerHolderNil(b bool)`

 SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil

### UnsetMaxStoredCertificatePerHolder
`func (o *IntunePKCSConnector) UnsetMaxStoredCertificatePerHolder()`

UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
### GetProxy

`func (o *IntunePKCSConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *IntunePKCSConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *IntunePKCSConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *IntunePKCSConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *IntunePKCSConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *IntunePKCSConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTenant

`func (o *IntunePKCSConnector) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IntunePKCSConnector) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IntunePKCSConnector) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetCredentials

`func (o *IntunePKCSConnector) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *IntunePKCSConnector) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *IntunePKCSConnector) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetPubKey

`func (o *IntunePKCSConnector) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *IntunePKCSConnector) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *IntunePKCSConnector) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetKeyName

`func (o *IntunePKCSConnector) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *IntunePKCSConnector) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *IntunePKCSConnector) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetProviderName

`func (o *IntunePKCSConnector) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *IntunePKCSConnector) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *IntunePKCSConnector) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *IntunePKCSConnector) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *IntunePKCSConnector) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *IntunePKCSConnector) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetIntendedPurpose

`func (o *IntunePKCSConnector) GetIntendedPurpose() string`

GetIntendedPurpose returns the IntendedPurpose field if non-nil, zero value otherwise.

### GetIntendedPurposeOk

`func (o *IntunePKCSConnector) GetIntendedPurposeOk() (*string, bool)`

GetIntendedPurposeOk returns a tuple with the IntendedPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedPurpose

`func (o *IntunePKCSConnector) SetIntendedPurpose(v string)`

SetIntendedPurpose sets IntendedPurpose field to given value.

### HasIntendedPurpose

`func (o *IntunePKCSConnector) HasIntendedPurpose() bool`

HasIntendedPurpose returns a boolean if a field has been set.

### SetIntendedPurposeNil

`func (o *IntunePKCSConnector) SetIntendedPurposeNil(b bool)`

 SetIntendedPurposeNil sets the value for IntendedPurpose to be an explicit nil

### UnsetIntendedPurpose
`func (o *IntunePKCSConnector) UnsetIntendedPurpose()`

UnsetIntendedPurpose ensures that no value is present for IntendedPurpose, not even an explicit nil
### GetSearchFilter

`func (o *IntunePKCSConnector) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *IntunePKCSConnector) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *IntunePKCSConnector) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *IntunePKCSConnector) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.

### SetSearchFilterNil

`func (o *IntunePKCSConnector) SetSearchFilterNil(b bool)`

 SetSearchFilterNil sets the value for SearchFilter to be an explicit nil

### UnsetSearchFilter
`func (o *IntunePKCSConnector) UnsetSearchFilter()`

UnsetSearchFilter ensures that no value is present for SearchFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


