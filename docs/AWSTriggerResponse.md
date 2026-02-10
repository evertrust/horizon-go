# AWSTriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**Connector** | **string** |  | 

## Methods

### NewAWSTriggerResponse

`func NewAWSTriggerResponse(id string, name string, type_ string, connector string, ) *AWSTriggerResponse`

NewAWSTriggerResponse instantiates a new AWSTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSTriggerResponseWithDefaults

`func NewAWSTriggerResponseWithDefaults() *AWSTriggerResponse`

NewAWSTriggerResponseWithDefaults instantiates a new AWSTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AWSTriggerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSTriggerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSTriggerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AWSTriggerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AWSTriggerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AWSTriggerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AWSTriggerResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSTriggerResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSTriggerResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRetries

`func (o *AWSTriggerResponse) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *AWSTriggerResponse) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *AWSTriggerResponse) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *AWSTriggerResponse) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *AWSTriggerResponse) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *AWSTriggerResponse) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetConnector

`func (o *AWSTriggerResponse) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *AWSTriggerResponse) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *AWSTriggerResponse) SetConnector(v string)`

SetConnector sets Connector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


