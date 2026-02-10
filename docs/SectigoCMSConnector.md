# SectigoCMSConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewSectigoCMSConnector

`func NewSectigoCMSConnector(name string, type_ string, loginCredentials string, customerUri string, organizationId int64, profile string, ) *SectigoCMSConnector`

NewSectigoCMSConnector instantiates a new SectigoCMSConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectigoCMSConnectorWithDefaults

`func NewSectigoCMSConnectorWithDefaults() *SectigoCMSConnector`

NewSectigoCMSConnectorWithDefaults instantiates a new SectigoCMSConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SectigoCMSConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectigoCMSConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectigoCMSConnector) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SectigoCMSConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SectigoCMSConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SectigoCMSConnector) SetType(v string)`

SetType sets Type field to given value.


### GetLoginCredentials

`func (o *SectigoCMSConnector) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *SectigoCMSConnector) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *SectigoCMSConnector) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetCustomerUri

`func (o *SectigoCMSConnector) GetCustomerUri() string`

GetCustomerUri returns the CustomerUri field if non-nil, zero value otherwise.

### GetCustomerUriOk

`func (o *SectigoCMSConnector) GetCustomerUriOk() (*string, bool)`

GetCustomerUriOk returns a tuple with the CustomerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUri

`func (o *SectigoCMSConnector) SetCustomerUri(v string)`

SetCustomerUri sets CustomerUri field to given value.


### GetOrganizationId

`func (o *SectigoCMSConnector) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SectigoCMSConnector) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SectigoCMSConnector) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetProfile

`func (o *SectigoCMSConnector) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SectigoCMSConnector) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SectigoCMSConnector) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRetryInterval

`func (o *SectigoCMSConnector) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *SectigoCMSConnector) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *SectigoCMSConnector) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *SectigoCMSConnector) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *SectigoCMSConnector) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *SectigoCMSConnector) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetValidDays

`func (o *SectigoCMSConnector) GetValidDays() string`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *SectigoCMSConnector) GetValidDaysOk() (*string, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *SectigoCMSConnector) SetValidDays(v string)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *SectigoCMSConnector) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### SetValidDaysNil

`func (o *SectigoCMSConnector) SetValidDaysNil(b bool)`

 SetValidDaysNil sets the value for ValidDays to be an explicit nil

### UnsetValidDays
`func (o *SectigoCMSConnector) UnsetValidDays()`

UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
### GetEndpointType

`func (o *SectigoCMSConnector) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *SectigoCMSConnector) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *SectigoCMSConnector) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *SectigoCMSConnector) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetTimeout

`func (o *SectigoCMSConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SectigoCMSConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SectigoCMSConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SectigoCMSConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *SectigoCMSConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SectigoCMSConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *SectigoCMSConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SectigoCMSConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SectigoCMSConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SectigoCMSConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *SectigoCMSConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *SectigoCMSConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *SectigoCMSConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *SectigoCMSConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *SectigoCMSConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *SectigoCMSConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *SectigoCMSConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *SectigoCMSConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


