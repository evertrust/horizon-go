# StreamConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**EndPoint** | **string** | Stream&#39;s base endpoint | 
**Template** | **string** | Stream&#39;s certificate template to use for enrollment | 
**Ca** | **string** | Stream&#39;s technical name of the CA on which to enroll | 
**LoginCredentials** | Pointer to **NullableString** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | [optional] 
**AuthenticationCredentials** | Pointer to **NullableString** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStreamConnector

`func NewStreamConnector(name string, type_ string, endPoint string, template string, ca string, ) *StreamConnector`

NewStreamConnector instantiates a new StreamConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamConnectorWithDefaults

`func NewStreamConnectorWithDefaults() *StreamConnector`

NewStreamConnectorWithDefaults instantiates a new StreamConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StreamConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamConnector) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *StreamConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamConnector) SetType(v string)`

SetType sets Type field to given value.


### GetEndPoint

`func (o *StreamConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *StreamConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *StreamConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetTemplate

`func (o *StreamConnector) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *StreamConnector) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *StreamConnector) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetCa

`func (o *StreamConnector) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *StreamConnector) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *StreamConnector) SetCa(v string)`

SetCa sets Ca field to given value.


### GetLoginCredentials

`func (o *StreamConnector) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *StreamConnector) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *StreamConnector) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.

### HasLoginCredentials

`func (o *StreamConnector) HasLoginCredentials() bool`

HasLoginCredentials returns a boolean if a field has been set.

### SetLoginCredentialsNil

`func (o *StreamConnector) SetLoginCredentialsNil(b bool)`

 SetLoginCredentialsNil sets the value for LoginCredentials to be an explicit nil

### UnsetLoginCredentials
`func (o *StreamConnector) UnsetLoginCredentials()`

UnsetLoginCredentials ensures that no value is present for LoginCredentials, not even an explicit nil
### GetAuthenticationCredentials

`func (o *StreamConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *StreamConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *StreamConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *StreamConnector) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### SetAuthenticationCredentialsNil

`func (o *StreamConnector) SetAuthenticationCredentialsNil(b bool)`

 SetAuthenticationCredentialsNil sets the value for AuthenticationCredentials to be an explicit nil

### UnsetAuthenticationCredentials
`func (o *StreamConnector) UnsetAuthenticationCredentials()`

UnsetAuthenticationCredentials ensures that no value is present for AuthenticationCredentials, not even an explicit nil
### GetTimeout

`func (o *StreamConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *StreamConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *StreamConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *StreamConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *StreamConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *StreamConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *StreamConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *StreamConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *StreamConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *StreamConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *StreamConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *StreamConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *StreamConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *StreamConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *StreamConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *StreamConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *StreamConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *StreamConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


