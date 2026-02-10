# F5AS3Trigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**Connector** | **string** |  | 

## Methods

### NewF5AS3Trigger

`func NewF5AS3Trigger(name string, type_ string, connector string, ) *F5AS3Trigger`

NewF5AS3Trigger instantiates a new F5AS3Trigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewF5AS3TriggerWithDefaults

`func NewF5AS3TriggerWithDefaults() *F5AS3Trigger`

NewF5AS3TriggerWithDefaults instantiates a new F5AS3Trigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *F5AS3Trigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *F5AS3Trigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *F5AS3Trigger) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *F5AS3Trigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *F5AS3Trigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *F5AS3Trigger) SetType(v string)`

SetType sets Type field to given value.


### GetRetries

`func (o *F5AS3Trigger) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *F5AS3Trigger) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *F5AS3Trigger) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *F5AS3Trigger) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *F5AS3Trigger) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *F5AS3Trigger) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetConnector

`func (o *F5AS3Trigger) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *F5AS3Trigger) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *F5AS3Trigger) SetConnector(v string)`

SetConnector sets Connector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


