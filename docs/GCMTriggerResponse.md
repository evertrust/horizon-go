# GCMTriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**Connector** | **string** |  | 

## Methods

### NewGCMTriggerResponse

`func NewGCMTriggerResponse(id string, name string, type_ string, connector string, ) *GCMTriggerResponse`

NewGCMTriggerResponse instantiates a new GCMTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCMTriggerResponseWithDefaults

`func NewGCMTriggerResponseWithDefaults() *GCMTriggerResponse`

NewGCMTriggerResponseWithDefaults instantiates a new GCMTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GCMTriggerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GCMTriggerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GCMTriggerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GCMTriggerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GCMTriggerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GCMTriggerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *GCMTriggerResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCMTriggerResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCMTriggerResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRetries

`func (o *GCMTriggerResponse) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *GCMTriggerResponse) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *GCMTriggerResponse) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *GCMTriggerResponse) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *GCMTriggerResponse) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *GCMTriggerResponse) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetConnector

`func (o *GCMTriggerResponse) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *GCMTriggerResponse) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *GCMTriggerResponse) SetConnector(v string)`

SetConnector sets Connector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


