# AzureKeyVaultConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Tenant** | **string** |  | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the App ID and Key to authenticate on AKV | 
**VaultBaseUrl** | **string** |  | 
**Prefix** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAzureKeyVaultConnector

`func NewAzureKeyVaultConnector(type_ string, name string, throttleDuration string, throttleParallelism int64, tenant string, credentials string, vaultBaseUrl string, ) *AzureKeyVaultConnector`

NewAzureKeyVaultConnector instantiates a new AzureKeyVaultConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultConnectorWithDefaults

`func NewAzureKeyVaultConnectorWithDefaults() *AzureKeyVaultConnector`

NewAzureKeyVaultConnectorWithDefaults instantiates a new AzureKeyVaultConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AzureKeyVaultConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureKeyVaultConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureKeyVaultConnector) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AzureKeyVaultConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureKeyVaultConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureKeyVaultConnector) SetName(v string)`

SetName sets Name field to given value.


### GetThrottleDuration

`func (o *AzureKeyVaultConnector) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *AzureKeyVaultConnector) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *AzureKeyVaultConnector) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *AzureKeyVaultConnector) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *AzureKeyVaultConnector) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *AzureKeyVaultConnector) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetRenewalPeriod

`func (o *AzureKeyVaultConnector) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *AzureKeyVaultConnector) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *AzureKeyVaultConnector) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *AzureKeyVaultConnector) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *AzureKeyVaultConnector) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *AzureKeyVaultConnector) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetTimeout

`func (o *AzureKeyVaultConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AzureKeyVaultConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AzureKeyVaultConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AzureKeyVaultConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *AzureKeyVaultConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *AzureKeyVaultConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *AzureKeyVaultConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AzureKeyVaultConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AzureKeyVaultConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AzureKeyVaultConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *AzureKeyVaultConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *AzureKeyVaultConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTenant

`func (o *AzureKeyVaultConnector) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AzureKeyVaultConnector) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AzureKeyVaultConnector) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetCredentials

`func (o *AzureKeyVaultConnector) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AzureKeyVaultConnector) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AzureKeyVaultConnector) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetVaultBaseUrl

`func (o *AzureKeyVaultConnector) GetVaultBaseUrl() string`

GetVaultBaseUrl returns the VaultBaseUrl field if non-nil, zero value otherwise.

### GetVaultBaseUrlOk

`func (o *AzureKeyVaultConnector) GetVaultBaseUrlOk() (*string, bool)`

GetVaultBaseUrlOk returns a tuple with the VaultBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultBaseUrl

`func (o *AzureKeyVaultConnector) SetVaultBaseUrl(v string)`

SetVaultBaseUrl sets VaultBaseUrl field to given value.


### GetPrefix

`func (o *AzureKeyVaultConnector) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AzureKeyVaultConnector) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AzureKeyVaultConnector) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *AzureKeyVaultConnector) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *AzureKeyVaultConnector) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *AzureKeyVaultConnector) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


