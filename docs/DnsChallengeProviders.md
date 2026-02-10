# DnsChallengeProviders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**SetTriggers** | [**[]AcmeRestRequest**](AcmeRestRequest.md) | The triggers that will set the DNS challenge on the provider. | 
**UnsetTriggers** | Pointer to [**[]AcmeRestRequest**](AcmeRestRequest.md) | The triggers that will unset the DNS challenge on the provider. | [optional] 
**Credentials** | **string** | &#x60;raw&#x60; credentials name to use to authenticate on the Nameshield API | 
**EndPoint** | **string** | Nameshield API endpoint | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Timeout** | **string** | Timeout when requesting Nameshield API | 

## Methods

### NewDnsChallengeProviders

`func NewDnsChallengeProviders(type_ string, setTriggers []AcmeRestRequest, credentials string, endPoint string, timeout string, ) *DnsChallengeProviders`

NewDnsChallengeProviders instantiates a new DnsChallengeProviders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsChallengeProvidersWithDefaults

`func NewDnsChallengeProvidersWithDefaults() *DnsChallengeProviders`

NewDnsChallengeProvidersWithDefaults instantiates a new DnsChallengeProviders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DnsChallengeProviders) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsChallengeProviders) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsChallengeProviders) SetType(v string)`

SetType sets Type field to given value.


### GetSetTriggers

`func (o *DnsChallengeProviders) GetSetTriggers() []AcmeRestRequest`

GetSetTriggers returns the SetTriggers field if non-nil, zero value otherwise.

### GetSetTriggersOk

`func (o *DnsChallengeProviders) GetSetTriggersOk() (*[]AcmeRestRequest, bool)`

GetSetTriggersOk returns a tuple with the SetTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTriggers

`func (o *DnsChallengeProviders) SetSetTriggers(v []AcmeRestRequest)`

SetSetTriggers sets SetTriggers field to given value.


### GetUnsetTriggers

`func (o *DnsChallengeProviders) GetUnsetTriggers() []AcmeRestRequest`

GetUnsetTriggers returns the UnsetTriggers field if non-nil, zero value otherwise.

### GetUnsetTriggersOk

`func (o *DnsChallengeProviders) GetUnsetTriggersOk() (*[]AcmeRestRequest, bool)`

GetUnsetTriggersOk returns a tuple with the UnsetTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsetTriggers

`func (o *DnsChallengeProviders) SetUnsetTriggers(v []AcmeRestRequest)`

SetUnsetTriggers sets UnsetTriggers field to given value.

### HasUnsetTriggers

`func (o *DnsChallengeProviders) HasUnsetTriggers() bool`

HasUnsetTriggers returns a boolean if a field has been set.

### GetCredentials

`func (o *DnsChallengeProviders) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DnsChallengeProviders) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DnsChallengeProviders) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetEndPoint

`func (o *DnsChallengeProviders) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *DnsChallengeProviders) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *DnsChallengeProviders) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetProxy

`func (o *DnsChallengeProviders) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DnsChallengeProviders) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DnsChallengeProviders) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DnsChallengeProviders) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DnsChallengeProviders) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DnsChallengeProviders) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *DnsChallengeProviders) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DnsChallengeProviders) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DnsChallengeProviders) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


