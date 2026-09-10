# NameshieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ApiCredentials** | **string** | Name of the &#x60;api-key&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**CustomConnectorDataMapping** | Pointer to **map[string]string** | Custom mapping of connector data fields | [optional] 
**CustomerId** | **string** |  | 
**Environment** | **string** | The testing environment will use https://ote-api.nameshield.net endpoint  and the production will use https://api.nameshield.net  | 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**ProductId** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | **string** |  | 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewNameshieldResponse

`func NewNameshieldResponse(id string, apiCredentials string, customerId string, environment string, name string, organizationId string, productId string, retryInterval string, type_ string, ) *NameshieldResponse`

NewNameshieldResponse instantiates a new NameshieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameshieldResponseWithDefaults

`func NewNameshieldResponseWithDefaults() *NameshieldResponse`

NewNameshieldResponseWithDefaults instantiates a new NameshieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NameshieldResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NameshieldResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NameshieldResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApiCredentials

`func (o *NameshieldResponse) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *NameshieldResponse) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *NameshieldResponse) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetCustomConnectorDataMapping

`func (o *NameshieldResponse) GetCustomConnectorDataMapping() map[string]string`

GetCustomConnectorDataMapping returns the CustomConnectorDataMapping field if non-nil, zero value otherwise.

### GetCustomConnectorDataMappingOk

`func (o *NameshieldResponse) GetCustomConnectorDataMappingOk() (*map[string]string, bool)`

GetCustomConnectorDataMappingOk returns a tuple with the CustomConnectorDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConnectorDataMapping

`func (o *NameshieldResponse) SetCustomConnectorDataMapping(v map[string]string)`

SetCustomConnectorDataMapping sets CustomConnectorDataMapping field to given value.

### HasCustomConnectorDataMapping

`func (o *NameshieldResponse) HasCustomConnectorDataMapping() bool`

HasCustomConnectorDataMapping returns a boolean if a field has been set.

### SetCustomConnectorDataMappingNil

`func (o *NameshieldResponse) SetCustomConnectorDataMappingNil(b bool)`

 SetCustomConnectorDataMappingNil sets the value for CustomConnectorDataMapping to be an explicit nil

### UnsetCustomConnectorDataMapping
`func (o *NameshieldResponse) UnsetCustomConnectorDataMapping()`

UnsetCustomConnectorDataMapping ensures that no value is present for CustomConnectorDataMapping, not even an explicit nil
### GetCustomerId

`func (o *NameshieldResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *NameshieldResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *NameshieldResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEnvironment

`func (o *NameshieldResponse) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NameshieldResponse) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NameshieldResponse) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetName

`func (o *NameshieldResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NameshieldResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NameshieldResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *NameshieldResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *NameshieldResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *NameshieldResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProductId

`func (o *NameshieldResponse) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *NameshieldResponse) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *NameshieldResponse) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProxy

`func (o *NameshieldResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *NameshieldResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *NameshieldResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *NameshieldResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *NameshieldResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *NameshieldResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *NameshieldResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *NameshieldResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *NameshieldResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *NameshieldResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *NameshieldResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *NameshieldResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetRetryInterval

`func (o *NameshieldResponse) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *NameshieldResponse) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *NameshieldResponse) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.


### GetStatus

`func (o *NameshieldResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NameshieldResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NameshieldResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NameshieldResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *NameshieldResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *NameshieldResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTimeout

`func (o *NameshieldResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *NameshieldResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *NameshieldResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *NameshieldResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *NameshieldResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *NameshieldResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *NameshieldResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NameshieldResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NameshieldResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


