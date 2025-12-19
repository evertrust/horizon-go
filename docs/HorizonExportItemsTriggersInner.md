# HorizonExportItemsTriggersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Connector** | **string** |  | 
**Name** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewHorizonExportItemsTriggersInner

`func NewHorizonExportItemsTriggersInner(id string, connector string, name string, type_ string, ) *HorizonExportItemsTriggersInner`

NewHorizonExportItemsTriggersInner instantiates a new HorizonExportItemsTriggersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHorizonExportItemsTriggersInnerWithDefaults

`func NewHorizonExportItemsTriggersInnerWithDefaults() *HorizonExportItemsTriggersInner`

NewHorizonExportItemsTriggersInnerWithDefaults instantiates a new HorizonExportItemsTriggersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HorizonExportItemsTriggersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HorizonExportItemsTriggersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HorizonExportItemsTriggersInner) SetId(v string)`

SetId sets Id field to given value.


### GetConnector

`func (o *HorizonExportItemsTriggersInner) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *HorizonExportItemsTriggersInner) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *HorizonExportItemsTriggersInner) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetName

`func (o *HorizonExportItemsTriggersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HorizonExportItemsTriggersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HorizonExportItemsTriggersInner) SetName(v string)`

SetName sets Name field to given value.


### GetRetries

`func (o *HorizonExportItemsTriggersInner) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *HorizonExportItemsTriggersInner) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *HorizonExportItemsTriggersInner) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *HorizonExportItemsTriggersInner) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *HorizonExportItemsTriggersInner) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *HorizonExportItemsTriggersInner) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetType

`func (o *HorizonExportItemsTriggersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HorizonExportItemsTriggersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HorizonExportItemsTriggersInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


