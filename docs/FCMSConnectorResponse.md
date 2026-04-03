# FCMSConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ApiCredentials** | **string** | Name of the &#x60;raw&#x60; [credentials](#tag/security.credentials) containing the API key to authenticate on the PKI | 
**AuthenticationDomainId** | **int64** |  | 
**DefaultOwner** | **string** |  | 
**DeleteOnRevoke** | **bool** |  | 
**EndPoint** | **string** |  | 
**Name** | **string** |  | 
**OwnerGroups** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
**TemplateId** | **int64** |  | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFCMSConnectorResponse

`func NewFCMSConnectorResponse(id string, apiCredentials string, authenticationDomainId int64, defaultOwner string, deleteOnRevoke bool, endPoint string, name string, templateId int64, type_ string, ) *FCMSConnectorResponse`

NewFCMSConnectorResponse instantiates a new FCMSConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFCMSConnectorResponseWithDefaults

`func NewFCMSConnectorResponseWithDefaults() *FCMSConnectorResponse`

NewFCMSConnectorResponseWithDefaults instantiates a new FCMSConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FCMSConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FCMSConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FCMSConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApiCredentials

`func (o *FCMSConnectorResponse) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *FCMSConnectorResponse) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *FCMSConnectorResponse) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetAuthenticationDomainId

`func (o *FCMSConnectorResponse) GetAuthenticationDomainId() int64`

GetAuthenticationDomainId returns the AuthenticationDomainId field if non-nil, zero value otherwise.

### GetAuthenticationDomainIdOk

`func (o *FCMSConnectorResponse) GetAuthenticationDomainIdOk() (*int64, bool)`

GetAuthenticationDomainIdOk returns a tuple with the AuthenticationDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationDomainId

`func (o *FCMSConnectorResponse) SetAuthenticationDomainId(v int64)`

SetAuthenticationDomainId sets AuthenticationDomainId field to given value.


### GetDefaultOwner

`func (o *FCMSConnectorResponse) GetDefaultOwner() string`

GetDefaultOwner returns the DefaultOwner field if non-nil, zero value otherwise.

### GetDefaultOwnerOk

`func (o *FCMSConnectorResponse) GetDefaultOwnerOk() (*string, bool)`

GetDefaultOwnerOk returns a tuple with the DefaultOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOwner

`func (o *FCMSConnectorResponse) SetDefaultOwner(v string)`

SetDefaultOwner sets DefaultOwner field to given value.


### GetDeleteOnRevoke

`func (o *FCMSConnectorResponse) GetDeleteOnRevoke() bool`

GetDeleteOnRevoke returns the DeleteOnRevoke field if non-nil, zero value otherwise.

### GetDeleteOnRevokeOk

`func (o *FCMSConnectorResponse) GetDeleteOnRevokeOk() (*bool, bool)`

GetDeleteOnRevokeOk returns a tuple with the DeleteOnRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnRevoke

`func (o *FCMSConnectorResponse) SetDeleteOnRevoke(v bool)`

SetDeleteOnRevoke sets DeleteOnRevoke field to given value.


### GetEndPoint

`func (o *FCMSConnectorResponse) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *FCMSConnectorResponse) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *FCMSConnectorResponse) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetName

`func (o *FCMSConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FCMSConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FCMSConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerGroups

`func (o *FCMSConnectorResponse) GetOwnerGroups() string`

GetOwnerGroups returns the OwnerGroups field if non-nil, zero value otherwise.

### GetOwnerGroupsOk

`func (o *FCMSConnectorResponse) GetOwnerGroupsOk() (*string, bool)`

GetOwnerGroupsOk returns a tuple with the OwnerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroups

`func (o *FCMSConnectorResponse) SetOwnerGroups(v string)`

SetOwnerGroups sets OwnerGroups field to given value.

### HasOwnerGroups

`func (o *FCMSConnectorResponse) HasOwnerGroups() bool`

HasOwnerGroups returns a boolean if a field has been set.

### SetOwnerGroupsNil

`func (o *FCMSConnectorResponse) SetOwnerGroupsNil(b bool)`

 SetOwnerGroupsNil sets the value for OwnerGroups to be an explicit nil

### UnsetOwnerGroups
`func (o *FCMSConnectorResponse) UnsetOwnerGroups()`

UnsetOwnerGroups ensures that no value is present for OwnerGroups, not even an explicit nil
### GetProxy

`func (o *FCMSConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *FCMSConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *FCMSConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *FCMSConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *FCMSConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *FCMSConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *FCMSConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *FCMSConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *FCMSConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *FCMSConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *FCMSConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *FCMSConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *FCMSConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FCMSConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FCMSConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FCMSConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *FCMSConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FCMSConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTemplateId

`func (o *FCMSConnectorResponse) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *FCMSConnectorResponse) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *FCMSConnectorResponse) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.


### GetTimeout

`func (o *FCMSConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FCMSConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FCMSConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FCMSConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *FCMSConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *FCMSConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *FCMSConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FCMSConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FCMSConnectorResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


