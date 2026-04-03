# F5AS3TriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Connector** | **string** |  | 
**Name** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewF5AS3TriggerResponse

`func NewF5AS3TriggerResponse(id string, connector string, name string, type_ string, ) *F5AS3TriggerResponse`

NewF5AS3TriggerResponse instantiates a new F5AS3TriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewF5AS3TriggerResponseWithDefaults

`func NewF5AS3TriggerResponseWithDefaults() *F5AS3TriggerResponse`

NewF5AS3TriggerResponseWithDefaults instantiates a new F5AS3TriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *F5AS3TriggerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *F5AS3TriggerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *F5AS3TriggerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetConnector

`func (o *F5AS3TriggerResponse) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *F5AS3TriggerResponse) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *F5AS3TriggerResponse) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetName

`func (o *F5AS3TriggerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *F5AS3TriggerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *F5AS3TriggerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRetries

`func (o *F5AS3TriggerResponse) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *F5AS3TriggerResponse) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *F5AS3TriggerResponse) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *F5AS3TriggerResponse) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *F5AS3TriggerResponse) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *F5AS3TriggerResponse) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetType

`func (o *F5AS3TriggerResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *F5AS3TriggerResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *F5AS3TriggerResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


