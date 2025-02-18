# EventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Code** | [**EventCode**](EventCode.md) |  | 
**Details** | Pointer to [**[]EventDetail**](EventDetail.md) |  | [optional] 
**Module** | [**EventModule**](EventModule.md) |  | 
**Node** | **string** |  | 
**Timestamp** | **int64** |  | 
**RemoveAt** | Pointer to **NullableInt64** |  | [optional] 
**Seal** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewEventResponse

`func NewEventResponse(code EventCode, module EventModule, node string, timestamp int64, status string, ) *EventResponse`

NewEventResponse instantiates a new EventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseWithDefaults

`func NewEventResponseWithDefaults() *EventResponse`

NewEventResponseWithDefaults instantiates a new EventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EventResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EventResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCode

`func (o *EventResponse) GetCode() EventCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EventResponse) GetCodeOk() (*EventCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EventResponse) SetCode(v EventCode)`

SetCode sets Code field to given value.


### GetDetails

`func (o *EventResponse) GetDetails() []EventDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EventResponse) GetDetailsOk() (*[]EventDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EventResponse) SetDetails(v []EventDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EventResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *EventResponse) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *EventResponse) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetModule

`func (o *EventResponse) GetModule() EventModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *EventResponse) GetModuleOk() (*EventModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *EventResponse) SetModule(v EventModule)`

SetModule sets Module field to given value.


### GetNode

`func (o *EventResponse) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *EventResponse) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *EventResponse) SetNode(v string)`

SetNode sets Node field to given value.


### GetTimestamp

`func (o *EventResponse) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventResponse) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventResponse) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetRemoveAt

`func (o *EventResponse) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *EventResponse) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *EventResponse) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.

### HasRemoveAt

`func (o *EventResponse) HasRemoveAt() bool`

HasRemoveAt returns a boolean if a field has been set.

### SetRemoveAtNil

`func (o *EventResponse) SetRemoveAtNil(b bool)`

 SetRemoveAtNil sets the value for RemoveAt to be an explicit nil

### UnsetRemoveAt
`func (o *EventResponse) UnsetRemoveAt()`

UnsetRemoveAt ensures that no value is present for RemoveAt, not even an explicit nil
### GetSeal

`func (o *EventResponse) GetSeal() string`

GetSeal returns the Seal field if non-nil, zero value otherwise.

### GetSealOk

`func (o *EventResponse) GetSealOk() (*string, bool)`

GetSealOk returns a tuple with the Seal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeal

`func (o *EventResponse) SetSeal(v string)`

SetSeal sets Seal field to given value.

### HasSeal

`func (o *EventResponse) HasSeal() bool`

HasSeal returns a boolean if a field has been set.

### SetSealNil

`func (o *EventResponse) SetSealNil(b bool)`

 SetSealNil sets the value for Seal to be an explicit nil

### UnsetSeal
`func (o *EventResponse) UnsetSeal()`

UnsetSeal ensures that no value is present for Seal, not even an explicit nil
### GetStatus

`func (o *EventResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


