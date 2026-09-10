# EJBCARESTConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**CaName** | **string** |  | 
**EeProfile** | Pointer to **NullableString** |  | [optional] 
**EndPoint** | **string** |  | 
**Name** | **string** |  | 
**Profile** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewEJBCARESTConnectorResponse

`func NewEJBCARESTConnectorResponse(id string, authenticationCredentials string, caName string, endPoint string, name string, profile string, type_ string, ) *EJBCARESTConnectorResponse`

NewEJBCARESTConnectorResponse instantiates a new EJBCARESTConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEJBCARESTConnectorResponseWithDefaults

`func NewEJBCARESTConnectorResponseWithDefaults() *EJBCARESTConnectorResponse`

NewEJBCARESTConnectorResponseWithDefaults instantiates a new EJBCARESTConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EJBCARESTConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EJBCARESTConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EJBCARESTConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAuthenticationCredentials

`func (o *EJBCARESTConnectorResponse) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *EJBCARESTConnectorResponse) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *EJBCARESTConnectorResponse) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetCaName

`func (o *EJBCARESTConnectorResponse) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *EJBCARESTConnectorResponse) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *EJBCARESTConnectorResponse) SetCaName(v string)`

SetCaName sets CaName field to given value.


### GetEeProfile

`func (o *EJBCARESTConnectorResponse) GetEeProfile() string`

GetEeProfile returns the EeProfile field if non-nil, zero value otherwise.

### GetEeProfileOk

`func (o *EJBCARESTConnectorResponse) GetEeProfileOk() (*string, bool)`

GetEeProfileOk returns a tuple with the EeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeProfile

`func (o *EJBCARESTConnectorResponse) SetEeProfile(v string)`

SetEeProfile sets EeProfile field to given value.

### HasEeProfile

`func (o *EJBCARESTConnectorResponse) HasEeProfile() bool`

HasEeProfile returns a boolean if a field has been set.

### SetEeProfileNil

`func (o *EJBCARESTConnectorResponse) SetEeProfileNil(b bool)`

 SetEeProfileNil sets the value for EeProfile to be an explicit nil

### UnsetEeProfile
`func (o *EJBCARESTConnectorResponse) UnsetEeProfile()`

UnsetEeProfile ensures that no value is present for EeProfile, not even an explicit nil
### GetEndPoint

`func (o *EJBCARESTConnectorResponse) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *EJBCARESTConnectorResponse) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *EJBCARESTConnectorResponse) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetName

`func (o *EJBCARESTConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EJBCARESTConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EJBCARESTConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *EJBCARESTConnectorResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *EJBCARESTConnectorResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *EJBCARESTConnectorResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProxy

`func (o *EJBCARESTConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *EJBCARESTConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *EJBCARESTConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *EJBCARESTConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *EJBCARESTConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *EJBCARESTConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *EJBCARESTConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *EJBCARESTConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *EJBCARESTConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *EJBCARESTConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *EJBCARESTConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *EJBCARESTConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *EJBCARESTConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EJBCARESTConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EJBCARESTConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EJBCARESTConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EJBCARESTConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EJBCARESTConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTimeout

`func (o *EJBCARESTConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EJBCARESTConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EJBCARESTConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EJBCARESTConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *EJBCARESTConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *EJBCARESTConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *EJBCARESTConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EJBCARESTConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EJBCARESTConnectorResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


