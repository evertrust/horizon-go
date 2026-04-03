# IDCAConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AuthenticationCredentials** | **NullableString** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**EndPoint** | **string** |  | 
**Name** | **string** |  | 
**Profile** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewIDCAConnectorResponse

`func NewIDCAConnectorResponse(id string, authenticationCredentials NullableString, endPoint string, name string, profile string, type_ string, ) *IDCAConnectorResponse`

NewIDCAConnectorResponse instantiates a new IDCAConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIDCAConnectorResponseWithDefaults

`func NewIDCAConnectorResponseWithDefaults() *IDCAConnectorResponse`

NewIDCAConnectorResponseWithDefaults instantiates a new IDCAConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IDCAConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IDCAConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IDCAConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAuthenticationCredentials

`func (o *IDCAConnectorResponse) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *IDCAConnectorResponse) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *IDCAConnectorResponse) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### SetAuthenticationCredentialsNil

`func (o *IDCAConnectorResponse) SetAuthenticationCredentialsNil(b bool)`

 SetAuthenticationCredentialsNil sets the value for AuthenticationCredentials to be an explicit nil

### UnsetAuthenticationCredentials
`func (o *IDCAConnectorResponse) UnsetAuthenticationCredentials()`

UnsetAuthenticationCredentials ensures that no value is present for AuthenticationCredentials, not even an explicit nil
### GetEndPoint

`func (o *IDCAConnectorResponse) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *IDCAConnectorResponse) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *IDCAConnectorResponse) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetName

`func (o *IDCAConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IDCAConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IDCAConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *IDCAConnectorResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *IDCAConnectorResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *IDCAConnectorResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProxy

`func (o *IDCAConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *IDCAConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *IDCAConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *IDCAConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *IDCAConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *IDCAConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *IDCAConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *IDCAConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *IDCAConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *IDCAConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *IDCAConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *IDCAConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *IDCAConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IDCAConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IDCAConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IDCAConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *IDCAConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *IDCAConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTimeout

`func (o *IDCAConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *IDCAConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *IDCAConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *IDCAConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *IDCAConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *IDCAConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *IDCAConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IDCAConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IDCAConnectorResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


