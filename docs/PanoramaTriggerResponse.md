# PanoramaTriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Connector** | **string** |  | 
**Events** | [**[]TriggerEvent**](TriggerEvent.md) |  | 
**Name** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewPanoramaTriggerResponse

`func NewPanoramaTriggerResponse(id string, connector string, events []TriggerEvent, name string, type_ string, ) *PanoramaTriggerResponse`

NewPanoramaTriggerResponse instantiates a new PanoramaTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanoramaTriggerResponseWithDefaults

`func NewPanoramaTriggerResponseWithDefaults() *PanoramaTriggerResponse`

NewPanoramaTriggerResponseWithDefaults instantiates a new PanoramaTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PanoramaTriggerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PanoramaTriggerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PanoramaTriggerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetConnector

`func (o *PanoramaTriggerResponse) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *PanoramaTriggerResponse) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *PanoramaTriggerResponse) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetEvents

`func (o *PanoramaTriggerResponse) GetEvents() []TriggerEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PanoramaTriggerResponse) GetEventsOk() (*[]TriggerEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PanoramaTriggerResponse) SetEvents(v []TriggerEvent)`

SetEvents sets Events field to given value.


### GetName

`func (o *PanoramaTriggerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PanoramaTriggerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PanoramaTriggerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRetries

`func (o *PanoramaTriggerResponse) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *PanoramaTriggerResponse) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *PanoramaTriggerResponse) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *PanoramaTriggerResponse) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *PanoramaTriggerResponse) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *PanoramaTriggerResponse) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetType

`func (o *PanoramaTriggerResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PanoramaTriggerResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PanoramaTriggerResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


