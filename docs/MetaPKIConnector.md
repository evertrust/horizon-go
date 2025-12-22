# MetaPKIConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationCredentials** | Pointer to **NullableString** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | [optional] 
**EndPoint** | **string** | MetaPKI base endpoint | 
**EndPointIssuingCA** | **string** | Certificate authority of the endpoint | 
**FormPorteurName** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**ProfilCle** | **NullableString** |  | 
**Profile** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**ValidDays** | Pointer to **NullableString** |  | [optional] 
**Workflow** | **NullableString** |  | 

## Methods

### NewMetaPKIConnector

`func NewMetaPKIConnector(endPoint string, endPointIssuingCA string, name string, profilCle NullableString, profile string, type_ string, workflow NullableString, ) *MetaPKIConnector`

NewMetaPKIConnector instantiates a new MetaPKIConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaPKIConnectorWithDefaults

`func NewMetaPKIConnectorWithDefaults() *MetaPKIConnector`

NewMetaPKIConnectorWithDefaults instantiates a new MetaPKIConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationCredentials

`func (o *MetaPKIConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *MetaPKIConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *MetaPKIConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.

### HasAuthenticationCredentials

`func (o *MetaPKIConnector) HasAuthenticationCredentials() bool`

HasAuthenticationCredentials returns a boolean if a field has been set.

### SetAuthenticationCredentialsNil

`func (o *MetaPKIConnector) SetAuthenticationCredentialsNil(b bool)`

 SetAuthenticationCredentialsNil sets the value for AuthenticationCredentials to be an explicit nil

### UnsetAuthenticationCredentials
`func (o *MetaPKIConnector) UnsetAuthenticationCredentials()`

UnsetAuthenticationCredentials ensures that no value is present for AuthenticationCredentials, not even an explicit nil
### GetEndPoint

`func (o *MetaPKIConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *MetaPKIConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *MetaPKIConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetEndPointIssuingCA

`func (o *MetaPKIConnector) GetEndPointIssuingCA() string`

GetEndPointIssuingCA returns the EndPointIssuingCA field if non-nil, zero value otherwise.

### GetEndPointIssuingCAOk

`func (o *MetaPKIConnector) GetEndPointIssuingCAOk() (*string, bool)`

GetEndPointIssuingCAOk returns a tuple with the EndPointIssuingCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointIssuingCA

`func (o *MetaPKIConnector) SetEndPointIssuingCA(v string)`

SetEndPointIssuingCA sets EndPointIssuingCA field to given value.


### GetFormPorteurName

`func (o *MetaPKIConnector) GetFormPorteurName() string`

GetFormPorteurName returns the FormPorteurName field if non-nil, zero value otherwise.

### GetFormPorteurNameOk

`func (o *MetaPKIConnector) GetFormPorteurNameOk() (*string, bool)`

GetFormPorteurNameOk returns a tuple with the FormPorteurName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormPorteurName

`func (o *MetaPKIConnector) SetFormPorteurName(v string)`

SetFormPorteurName sets FormPorteurName field to given value.

### HasFormPorteurName

`func (o *MetaPKIConnector) HasFormPorteurName() bool`

HasFormPorteurName returns a boolean if a field has been set.

### SetFormPorteurNameNil

`func (o *MetaPKIConnector) SetFormPorteurNameNil(b bool)`

 SetFormPorteurNameNil sets the value for FormPorteurName to be an explicit nil

### UnsetFormPorteurName
`func (o *MetaPKIConnector) UnsetFormPorteurName()`

UnsetFormPorteurName ensures that no value is present for FormPorteurName, not even an explicit nil
### GetName

`func (o *MetaPKIConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaPKIConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaPKIConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProfilCle

`func (o *MetaPKIConnector) GetProfilCle() string`

GetProfilCle returns the ProfilCle field if non-nil, zero value otherwise.

### GetProfilCleOk

`func (o *MetaPKIConnector) GetProfilCleOk() (*string, bool)`

GetProfilCleOk returns a tuple with the ProfilCle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilCle

`func (o *MetaPKIConnector) SetProfilCle(v string)`

SetProfilCle sets ProfilCle field to given value.


### SetProfilCleNil

`func (o *MetaPKIConnector) SetProfilCleNil(b bool)`

 SetProfilCleNil sets the value for ProfilCle to be an explicit nil

### UnsetProfilCle
`func (o *MetaPKIConnector) UnsetProfilCle()`

UnsetProfilCle ensures that no value is present for ProfilCle, not even an explicit nil
### GetProfile

`func (o *MetaPKIConnector) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MetaPKIConnector) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MetaPKIConnector) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProxy

`func (o *MetaPKIConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *MetaPKIConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *MetaPKIConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *MetaPKIConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *MetaPKIConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *MetaPKIConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *MetaPKIConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *MetaPKIConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *MetaPKIConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *MetaPKIConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *MetaPKIConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *MetaPKIConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTimeout

`func (o *MetaPKIConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MetaPKIConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MetaPKIConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *MetaPKIConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *MetaPKIConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *MetaPKIConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *MetaPKIConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetaPKIConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetaPKIConnector) SetType(v string)`

SetType sets Type field to given value.


### GetValidDays

`func (o *MetaPKIConnector) GetValidDays() string`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *MetaPKIConnector) GetValidDaysOk() (*string, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *MetaPKIConnector) SetValidDays(v string)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *MetaPKIConnector) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### SetValidDaysNil

`func (o *MetaPKIConnector) SetValidDaysNil(b bool)`

 SetValidDaysNil sets the value for ValidDays to be an explicit nil

### UnsetValidDays
`func (o *MetaPKIConnector) UnsetValidDays()`

UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
### GetWorkflow

`func (o *MetaPKIConnector) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *MetaPKIConnector) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *MetaPKIConnector) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### SetWorkflowNil

`func (o *MetaPKIConnector) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *MetaPKIConnector) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


