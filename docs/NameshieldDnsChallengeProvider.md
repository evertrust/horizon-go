# NameshieldDnsChallengeProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **string** | &#x60;raw&#x60; credentials name to use to authenticate on the Nameshield API | 
**EndPoint** | **string** | Nameshield API endpoint | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Timeout** | **string** | Timeout when requesting Nameshield API | 
**Type** | **string** |  | 

## Methods

### NewNameshieldDnsChallengeProvider

`func NewNameshieldDnsChallengeProvider(credentials string, endPoint string, timeout string, type_ string, ) *NameshieldDnsChallengeProvider`

NewNameshieldDnsChallengeProvider instantiates a new NameshieldDnsChallengeProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameshieldDnsChallengeProviderWithDefaults

`func NewNameshieldDnsChallengeProviderWithDefaults() *NameshieldDnsChallengeProvider`

NewNameshieldDnsChallengeProviderWithDefaults instantiates a new NameshieldDnsChallengeProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *NameshieldDnsChallengeProvider) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *NameshieldDnsChallengeProvider) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *NameshieldDnsChallengeProvider) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetEndPoint

`func (o *NameshieldDnsChallengeProvider) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *NameshieldDnsChallengeProvider) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *NameshieldDnsChallengeProvider) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetProxy

`func (o *NameshieldDnsChallengeProvider) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *NameshieldDnsChallengeProvider) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *NameshieldDnsChallengeProvider) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *NameshieldDnsChallengeProvider) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *NameshieldDnsChallengeProvider) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *NameshieldDnsChallengeProvider) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *NameshieldDnsChallengeProvider) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *NameshieldDnsChallengeProvider) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *NameshieldDnsChallengeProvider) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetType

`func (o *NameshieldDnsChallengeProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NameshieldDnsChallengeProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NameshieldDnsChallengeProvider) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


