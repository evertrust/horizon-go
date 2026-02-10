# GCMConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**ThrottleDuration** | **string** |  | 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Project** | **string** |  | 
**Location** | **string** |  | 
**Credentials** | **string** | Name of the &#x60;raw&#x60; [credentials](#tag/security.credentials) containing User Account credentials. | 
**TagKey** | Pointer to **NullableString** |  | [optional] 
**TagValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGCMConnector

`func NewGCMConnector(type_ string, name string, throttleDuration string, project string, location string, credentials string, ) *GCMConnector`

NewGCMConnector instantiates a new GCMConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCMConnectorWithDefaults

`func NewGCMConnectorWithDefaults() *GCMConnector`

NewGCMConnectorWithDefaults instantiates a new GCMConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GCMConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCMConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCMConnector) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *GCMConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GCMConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GCMConnector) SetName(v string)`

SetName sets Name field to given value.


### GetThrottleDuration

`func (o *GCMConnector) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *GCMConnector) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *GCMConnector) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetRenewalPeriod

`func (o *GCMConnector) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *GCMConnector) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *GCMConnector) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *GCMConnector) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *GCMConnector) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *GCMConnector) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetTimeout

`func (o *GCMConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GCMConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GCMConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GCMConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *GCMConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *GCMConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *GCMConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *GCMConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *GCMConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *GCMConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *GCMConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *GCMConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetProject

`func (o *GCMConnector) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GCMConnector) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GCMConnector) SetProject(v string)`

SetProject sets Project field to given value.


### GetLocation

`func (o *GCMConnector) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GCMConnector) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GCMConnector) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCredentials

`func (o *GCMConnector) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GCMConnector) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GCMConnector) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetTagKey

`func (o *GCMConnector) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *GCMConnector) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *GCMConnector) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.

### HasTagKey

`func (o *GCMConnector) HasTagKey() bool`

HasTagKey returns a boolean if a field has been set.

### SetTagKeyNil

`func (o *GCMConnector) SetTagKeyNil(b bool)`

 SetTagKeyNil sets the value for TagKey to be an explicit nil

### UnsetTagKey
`func (o *GCMConnector) UnsetTagKey()`

UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil
### GetTagValue

`func (o *GCMConnector) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *GCMConnector) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *GCMConnector) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.

### HasTagValue

`func (o *GCMConnector) HasTagValue() bool`

HasTagValue returns a boolean if a field has been set.

### SetTagValueNil

`func (o *GCMConnector) SetTagValueNil(b bool)`

 SetTagValueNil sets the value for TagValue to be an explicit nil

### UnsetTagValue
`func (o *GCMConnector) UnsetTagValue()`

UnsetTagValue ensures that no value is present for TagValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


