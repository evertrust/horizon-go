# FortiManagerTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | **string** |  | 
**Name** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFortiManagerTrigger

`func NewFortiManagerTrigger(connector string, name string, type_ string, ) *FortiManagerTrigger`

NewFortiManagerTrigger instantiates a new FortiManagerTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFortiManagerTriggerWithDefaults

`func NewFortiManagerTriggerWithDefaults() *FortiManagerTrigger`

NewFortiManagerTriggerWithDefaults instantiates a new FortiManagerTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *FortiManagerTrigger) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *FortiManagerTrigger) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *FortiManagerTrigger) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetName

`func (o *FortiManagerTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FortiManagerTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FortiManagerTrigger) SetName(v string)`

SetName sets Name field to given value.


### GetRetries

`func (o *FortiManagerTrigger) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *FortiManagerTrigger) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *FortiManagerTrigger) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *FortiManagerTrigger) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *FortiManagerTrigger) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *FortiManagerTrigger) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetType

`func (o *FortiManagerTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FortiManagerTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FortiManagerTrigger) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


