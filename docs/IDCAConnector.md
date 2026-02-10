# IDCAConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**EndPoint** | **string** |  | 
**Profile** | **string** |  | 
**AuthenticationCredentials** | **NullableString** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIDCAConnector

`func NewIDCAConnector(name string, type_ string, endPoint string, profile string, authenticationCredentials NullableString, ) *IDCAConnector`

NewIDCAConnector instantiates a new IDCAConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIDCAConnectorWithDefaults

`func NewIDCAConnectorWithDefaults() *IDCAConnector`

NewIDCAConnectorWithDefaults instantiates a new IDCAConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IDCAConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IDCAConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IDCAConnector) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IDCAConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IDCAConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IDCAConnector) SetType(v string)`

SetType sets Type field to given value.


### GetEndPoint

`func (o *IDCAConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *IDCAConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *IDCAConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetProfile

`func (o *IDCAConnector) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *IDCAConnector) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *IDCAConnector) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetAuthenticationCredentials

`func (o *IDCAConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *IDCAConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *IDCAConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### SetAuthenticationCredentialsNil

`func (o *IDCAConnector) SetAuthenticationCredentialsNil(b bool)`

 SetAuthenticationCredentialsNil sets the value for AuthenticationCredentials to be an explicit nil

### UnsetAuthenticationCredentials
`func (o *IDCAConnector) UnsetAuthenticationCredentials()`

UnsetAuthenticationCredentials ensures that no value is present for AuthenticationCredentials, not even an explicit nil
### GetTimeout

`func (o *IDCAConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *IDCAConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *IDCAConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *IDCAConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *IDCAConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *IDCAConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *IDCAConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *IDCAConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *IDCAConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *IDCAConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *IDCAConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *IDCAConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *IDCAConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *IDCAConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *IDCAConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *IDCAConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *IDCAConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *IDCAConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


