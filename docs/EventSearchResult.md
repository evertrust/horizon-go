# EventSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to [**NullableEventCode**](EventCode.md) |  | [optional] 
**Details** | Pointer to [**[]EventDetail**](EventDetail.md) |  | [optional] 
**Module** | Pointer to [**NullableEventModule**](EventModule.md) |  | [optional] 
**Node** | Pointer to **NullableString** |  | [optional] 
**RemoveAt** | Pointer to **NullableInt64** |  | [optional] 
**Seal** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewEventSearchResult

`func NewEventSearchResult() *EventSearchResult`

NewEventSearchResult instantiates a new EventSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSearchResultWithDefaults

`func NewEventSearchResultWithDefaults() *EventSearchResult`

NewEventSearchResultWithDefaults instantiates a new EventSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventSearchResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventSearchResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventSearchResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventSearchResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EventSearchResult) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EventSearchResult) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCode

`func (o *EventSearchResult) GetCode() EventCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EventSearchResult) GetCodeOk() (*EventCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EventSearchResult) SetCode(v EventCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *EventSearchResult) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *EventSearchResult) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *EventSearchResult) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDetails

`func (o *EventSearchResult) GetDetails() []EventDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EventSearchResult) GetDetailsOk() (*[]EventDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EventSearchResult) SetDetails(v []EventDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EventSearchResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *EventSearchResult) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *EventSearchResult) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetModule

`func (o *EventSearchResult) GetModule() EventModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *EventSearchResult) GetModuleOk() (*EventModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *EventSearchResult) SetModule(v EventModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *EventSearchResult) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *EventSearchResult) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *EventSearchResult) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetNode

`func (o *EventSearchResult) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *EventSearchResult) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *EventSearchResult) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *EventSearchResult) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *EventSearchResult) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *EventSearchResult) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetRemoveAt

`func (o *EventSearchResult) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *EventSearchResult) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *EventSearchResult) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.

### HasRemoveAt

`func (o *EventSearchResult) HasRemoveAt() bool`

HasRemoveAt returns a boolean if a field has been set.

### SetRemoveAtNil

`func (o *EventSearchResult) SetRemoveAtNil(b bool)`

 SetRemoveAtNil sets the value for RemoveAt to be an explicit nil

### UnsetRemoveAt
`func (o *EventSearchResult) UnsetRemoveAt()`

UnsetRemoveAt ensures that no value is present for RemoveAt, not even an explicit nil
### GetSeal

`func (o *EventSearchResult) GetSeal() string`

GetSeal returns the Seal field if non-nil, zero value otherwise.

### GetSealOk

`func (o *EventSearchResult) GetSealOk() (*string, bool)`

GetSealOk returns a tuple with the Seal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeal

`func (o *EventSearchResult) SetSeal(v string)`

SetSeal sets Seal field to given value.

### HasSeal

`func (o *EventSearchResult) HasSeal() bool`

HasSeal returns a boolean if a field has been set.

### SetSealNil

`func (o *EventSearchResult) SetSealNil(b bool)`

 SetSealNil sets the value for Seal to be an explicit nil

### UnsetSeal
`func (o *EventSearchResult) UnsetSeal()`

UnsetSeal ensures that no value is present for Seal, not even an explicit nil
### GetStatus

`func (o *EventSearchResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventSearchResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventSearchResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventSearchResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EventSearchResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EventSearchResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTimestamp

`func (o *EventSearchResult) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventSearchResult) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventSearchResult) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EventSearchResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *EventSearchResult) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *EventSearchResult) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


