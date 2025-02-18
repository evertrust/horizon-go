# IntuneConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Tenant** | **string** |  | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/api.security.credentials) containing the App ID and Key to authenticate on Intune | 
**IntuneResourceUrl** | Pointer to **NullableString** |  | [optional] 
**OsQueryString** | Pointer to **NullableString** |  | [optional] 
**LegacyRevocationMode** | **bool** |  | 

## Methods

### NewIntuneConnectorResponse

`func NewIntuneConnectorResponse(id string, type_ string, name string, throttleDuration string, throttleParallelism int64, tenant string, credentials string, legacyRevocationMode bool, ) *IntuneConnectorResponse`

NewIntuneConnectorResponse instantiates a new IntuneConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntuneConnectorResponseWithDefaults

`func NewIntuneConnectorResponseWithDefaults() *IntuneConnectorResponse`

NewIntuneConnectorResponseWithDefaults instantiates a new IntuneConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntuneConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntuneConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntuneConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *IntuneConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntuneConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntuneConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *IntuneConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntuneConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntuneConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetThrottleDuration

`func (o *IntuneConnectorResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *IntuneConnectorResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *IntuneConnectorResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *IntuneConnectorResponse) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *IntuneConnectorResponse) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *IntuneConnectorResponse) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *IntuneConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *IntuneConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *IntuneConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *IntuneConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *IntuneConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *IntuneConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *IntuneConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *IntuneConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *IntuneConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *IntuneConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *IntuneConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *IntuneConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTenant

`func (o *IntuneConnectorResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IntuneConnectorResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IntuneConnectorResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetCredentials

`func (o *IntuneConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *IntuneConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *IntuneConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetIntuneResourceUrl

`func (o *IntuneConnectorResponse) GetIntuneResourceUrl() string`

GetIntuneResourceUrl returns the IntuneResourceUrl field if non-nil, zero value otherwise.

### GetIntuneResourceUrlOk

`func (o *IntuneConnectorResponse) GetIntuneResourceUrlOk() (*string, bool)`

GetIntuneResourceUrlOk returns a tuple with the IntuneResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntuneResourceUrl

`func (o *IntuneConnectorResponse) SetIntuneResourceUrl(v string)`

SetIntuneResourceUrl sets IntuneResourceUrl field to given value.

### HasIntuneResourceUrl

`func (o *IntuneConnectorResponse) HasIntuneResourceUrl() bool`

HasIntuneResourceUrl returns a boolean if a field has been set.

### SetIntuneResourceUrlNil

`func (o *IntuneConnectorResponse) SetIntuneResourceUrlNil(b bool)`

 SetIntuneResourceUrlNil sets the value for IntuneResourceUrl to be an explicit nil

### UnsetIntuneResourceUrl
`func (o *IntuneConnectorResponse) UnsetIntuneResourceUrl()`

UnsetIntuneResourceUrl ensures that no value is present for IntuneResourceUrl, not even an explicit nil
### GetOsQueryString

`func (o *IntuneConnectorResponse) GetOsQueryString() string`

GetOsQueryString returns the OsQueryString field if non-nil, zero value otherwise.

### GetOsQueryStringOk

`func (o *IntuneConnectorResponse) GetOsQueryStringOk() (*string, bool)`

GetOsQueryStringOk returns a tuple with the OsQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsQueryString

`func (o *IntuneConnectorResponse) SetOsQueryString(v string)`

SetOsQueryString sets OsQueryString field to given value.

### HasOsQueryString

`func (o *IntuneConnectorResponse) HasOsQueryString() bool`

HasOsQueryString returns a boolean if a field has been set.

### SetOsQueryStringNil

`func (o *IntuneConnectorResponse) SetOsQueryStringNil(b bool)`

 SetOsQueryStringNil sets the value for OsQueryString to be an explicit nil

### UnsetOsQueryString
`func (o *IntuneConnectorResponse) UnsetOsQueryString()`

UnsetOsQueryString ensures that no value is present for OsQueryString, not even an explicit nil
### GetLegacyRevocationMode

`func (o *IntuneConnectorResponse) GetLegacyRevocationMode() bool`

GetLegacyRevocationMode returns the LegacyRevocationMode field if non-nil, zero value otherwise.

### GetLegacyRevocationModeOk

`func (o *IntuneConnectorResponse) GetLegacyRevocationModeOk() (*bool, bool)`

GetLegacyRevocationModeOk returns a tuple with the LegacyRevocationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyRevocationMode

`func (o *IntuneConnectorResponse) SetLegacyRevocationMode(v bool)`

SetLegacyRevocationMode sets LegacyRevocationMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


