# EJBCARESTConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**CaName** | **string** |  | 
**EeProfile** | Pointer to **NullableString** |  | [optional] 
**EndPoint** | **string** |  | 
**Name** | **string** |  | 
**Profile** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewEJBCARESTConnector

`func NewEJBCARESTConnector(authenticationCredentials string, caName string, endPoint string, name string, profile string, type_ string, ) *EJBCARESTConnector`

NewEJBCARESTConnector instantiates a new EJBCARESTConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEJBCARESTConnectorWithDefaults

`func NewEJBCARESTConnectorWithDefaults() *EJBCARESTConnector`

NewEJBCARESTConnectorWithDefaults instantiates a new EJBCARESTConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationCredentials

`func (o *EJBCARESTConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *EJBCARESTConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *EJBCARESTConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetCaName

`func (o *EJBCARESTConnector) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *EJBCARESTConnector) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *EJBCARESTConnector) SetCaName(v string)`

SetCaName sets CaName field to given value.


### GetEeProfile

`func (o *EJBCARESTConnector) GetEeProfile() string`

GetEeProfile returns the EeProfile field if non-nil, zero value otherwise.

### GetEeProfileOk

`func (o *EJBCARESTConnector) GetEeProfileOk() (*string, bool)`

GetEeProfileOk returns a tuple with the EeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeProfile

`func (o *EJBCARESTConnector) SetEeProfile(v string)`

SetEeProfile sets EeProfile field to given value.

### HasEeProfile

`func (o *EJBCARESTConnector) HasEeProfile() bool`

HasEeProfile returns a boolean if a field has been set.

### SetEeProfileNil

`func (o *EJBCARESTConnector) SetEeProfileNil(b bool)`

 SetEeProfileNil sets the value for EeProfile to be an explicit nil

### UnsetEeProfile
`func (o *EJBCARESTConnector) UnsetEeProfile()`

UnsetEeProfile ensures that no value is present for EeProfile, not even an explicit nil
### GetEndPoint

`func (o *EJBCARESTConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *EJBCARESTConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *EJBCARESTConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetName

`func (o *EJBCARESTConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EJBCARESTConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EJBCARESTConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *EJBCARESTConnector) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *EJBCARESTConnector) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *EJBCARESTConnector) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProxy

`func (o *EJBCARESTConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *EJBCARESTConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *EJBCARESTConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *EJBCARESTConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *EJBCARESTConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *EJBCARESTConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *EJBCARESTConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *EJBCARESTConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *EJBCARESTConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *EJBCARESTConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *EJBCARESTConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *EJBCARESTConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTimeout

`func (o *EJBCARESTConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EJBCARESTConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EJBCARESTConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EJBCARESTConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *EJBCARESTConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *EJBCARESTConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *EJBCARESTConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EJBCARESTConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EJBCARESTConnector) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


