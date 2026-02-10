# IntunePKCSTriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**Connector** | **string** |  | 

## Methods

### NewIntunePKCSTriggerResponse

`func NewIntunePKCSTriggerResponse(id string, name string, type_ string, connector string, ) *IntunePKCSTriggerResponse`

NewIntunePKCSTriggerResponse instantiates a new IntunePKCSTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntunePKCSTriggerResponseWithDefaults

`func NewIntunePKCSTriggerResponseWithDefaults() *IntunePKCSTriggerResponse`

NewIntunePKCSTriggerResponseWithDefaults instantiates a new IntunePKCSTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntunePKCSTriggerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntunePKCSTriggerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntunePKCSTriggerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IntunePKCSTriggerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntunePKCSTriggerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntunePKCSTriggerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IntunePKCSTriggerResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntunePKCSTriggerResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntunePKCSTriggerResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRetries

`func (o *IntunePKCSTriggerResponse) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *IntunePKCSTriggerResponse) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *IntunePKCSTriggerResponse) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *IntunePKCSTriggerResponse) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *IntunePKCSTriggerResponse) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *IntunePKCSTriggerResponse) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetConnector

`func (o *IntunePKCSTriggerResponse) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *IntunePKCSTriggerResponse) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *IntunePKCSTriggerResponse) SetConnector(v string)`

SetConnector sets Connector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


