# Nameshield

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiCredentials** | **string** | Name of the &#x60;api-key&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**CustomerId** | **string** |  | 
**Environment** | **string** | The testing environment will use https://ote-api.nameshield.net endpoint  and the production will use https://api.nameshield.net  | 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**ProductId** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewNameshield

`func NewNameshield(apiCredentials string, customerId string, environment string, name string, organizationId string, productId string, type_ string, ) *Nameshield`

NewNameshield instantiates a new Nameshield object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameshieldWithDefaults

`func NewNameshieldWithDefaults() *Nameshield`

NewNameshieldWithDefaults instantiates a new Nameshield object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiCredentials

`func (o *Nameshield) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *Nameshield) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *Nameshield) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetCustomerId

`func (o *Nameshield) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Nameshield) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Nameshield) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEnvironment

`func (o *Nameshield) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Nameshield) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Nameshield) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetName

`func (o *Nameshield) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Nameshield) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Nameshield) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *Nameshield) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Nameshield) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Nameshield) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProductId

`func (o *Nameshield) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Nameshield) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Nameshield) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProxy

`func (o *Nameshield) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *Nameshield) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *Nameshield) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *Nameshield) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *Nameshield) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *Nameshield) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *Nameshield) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *Nameshield) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *Nameshield) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *Nameshield) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *Nameshield) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *Nameshield) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTimeout

`func (o *Nameshield) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Nameshield) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Nameshield) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Nameshield) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *Nameshield) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *Nameshield) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *Nameshield) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Nameshield) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Nameshield) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


