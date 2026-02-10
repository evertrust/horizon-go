# SectigoCMSConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**CustomerUri** | **string** |  | 
**OrganizationId** | **int64** |  | 
**Profile** | **string** |  | 
**RetryInterval** | Pointer to **NullableString** |  | [optional] 
**ValidDays** | Pointer to **NullableString** |  | [optional] 
**EndpointType** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 

## Methods

### NewSectigoCMSConnectorResponse

`func NewSectigoCMSConnectorResponse(id string, name string, type_ string, loginCredentials string, customerUri string, organizationId int64, profile string, ) *SectigoCMSConnectorResponse`

NewSectigoCMSConnectorResponse instantiates a new SectigoCMSConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectigoCMSConnectorResponseWithDefaults

`func NewSectigoCMSConnectorResponseWithDefaults() *SectigoCMSConnectorResponse`

NewSectigoCMSConnectorResponseWithDefaults instantiates a new SectigoCMSConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SectigoCMSConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SectigoCMSConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SectigoCMSConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SectigoCMSConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectigoCMSConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectigoCMSConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SectigoCMSConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SectigoCMSConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SectigoCMSConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetLoginCredentials

`func (o *SectigoCMSConnectorResponse) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *SectigoCMSConnectorResponse) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *SectigoCMSConnectorResponse) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetCustomerUri

`func (o *SectigoCMSConnectorResponse) GetCustomerUri() string`

GetCustomerUri returns the CustomerUri field if non-nil, zero value otherwise.

### GetCustomerUriOk

`func (o *SectigoCMSConnectorResponse) GetCustomerUriOk() (*string, bool)`

GetCustomerUriOk returns a tuple with the CustomerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUri

`func (o *SectigoCMSConnectorResponse) SetCustomerUri(v string)`

SetCustomerUri sets CustomerUri field to given value.


### GetOrganizationId

`func (o *SectigoCMSConnectorResponse) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SectigoCMSConnectorResponse) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SectigoCMSConnectorResponse) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetProfile

`func (o *SectigoCMSConnectorResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SectigoCMSConnectorResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SectigoCMSConnectorResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRetryInterval

`func (o *SectigoCMSConnectorResponse) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *SectigoCMSConnectorResponse) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *SectigoCMSConnectorResponse) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *SectigoCMSConnectorResponse) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *SectigoCMSConnectorResponse) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *SectigoCMSConnectorResponse) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetValidDays

`func (o *SectigoCMSConnectorResponse) GetValidDays() string`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *SectigoCMSConnectorResponse) GetValidDaysOk() (*string, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *SectigoCMSConnectorResponse) SetValidDays(v string)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *SectigoCMSConnectorResponse) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### SetValidDaysNil

`func (o *SectigoCMSConnectorResponse) SetValidDaysNil(b bool)`

 SetValidDaysNil sets the value for ValidDays to be an explicit nil

### UnsetValidDays
`func (o *SectigoCMSConnectorResponse) UnsetValidDays()`

UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
### GetEndpointType

`func (o *SectigoCMSConnectorResponse) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *SectigoCMSConnectorResponse) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *SectigoCMSConnectorResponse) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *SectigoCMSConnectorResponse) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetTimeout

`func (o *SectigoCMSConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SectigoCMSConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SectigoCMSConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SectigoCMSConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *SectigoCMSConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SectigoCMSConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *SectigoCMSConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SectigoCMSConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SectigoCMSConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SectigoCMSConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *SectigoCMSConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *SectigoCMSConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *SectigoCMSConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *SectigoCMSConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *SectigoCMSConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *SectigoCMSConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *SectigoCMSConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *SectigoCMSConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *SectigoCMSConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SectigoCMSConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SectigoCMSConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SectigoCMSConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SectigoCMSConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SectigoCMSConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


