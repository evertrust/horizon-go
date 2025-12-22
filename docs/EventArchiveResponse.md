# EventArchiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Count** | Pointer to **NullableInt64** |  | [optional] 
**CreatedAt** | Pointer to **NullableInt64** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**PurgeAt** | Pointer to **NullableInt64** |  | [optional] 
**Status** | [**ArchiveStatus**](ArchiveStatus.md) |  | 
**Before** | **int64** | Date before which all events will be archived | 
**Filename** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewEventArchiveResponse

`func NewEventArchiveResponse(id string, status ArchiveStatus, before int64, filename string, name string, type_ string, ) *EventArchiveResponse`

NewEventArchiveResponse instantiates a new EventArchiveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventArchiveResponseWithDefaults

`func NewEventArchiveResponseWithDefaults() *EventArchiveResponse`

NewEventArchiveResponseWithDefaults instantiates a new EventArchiveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventArchiveResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventArchiveResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventArchiveResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCount

`func (o *EventArchiveResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventArchiveResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventArchiveResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *EventArchiveResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *EventArchiveResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *EventArchiveResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetCreatedAt

`func (o *EventArchiveResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventArchiveResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventArchiveResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventArchiveResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *EventArchiveResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *EventArchiveResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetError

`func (o *EventArchiveResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EventArchiveResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EventArchiveResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EventArchiveResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *EventArchiveResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *EventArchiveResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPurgeAt

`func (o *EventArchiveResponse) GetPurgeAt() int64`

GetPurgeAt returns the PurgeAt field if non-nil, zero value otherwise.

### GetPurgeAtOk

`func (o *EventArchiveResponse) GetPurgeAtOk() (*int64, bool)`

GetPurgeAtOk returns a tuple with the PurgeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeAt

`func (o *EventArchiveResponse) SetPurgeAt(v int64)`

SetPurgeAt sets PurgeAt field to given value.

### HasPurgeAt

`func (o *EventArchiveResponse) HasPurgeAt() bool`

HasPurgeAt returns a boolean if a field has been set.

### SetPurgeAtNil

`func (o *EventArchiveResponse) SetPurgeAtNil(b bool)`

 SetPurgeAtNil sets the value for PurgeAt to be an explicit nil

### UnsetPurgeAt
`func (o *EventArchiveResponse) UnsetPurgeAt()`

UnsetPurgeAt ensures that no value is present for PurgeAt, not even an explicit nil
### GetStatus

`func (o *EventArchiveResponse) GetStatus() ArchiveStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventArchiveResponse) GetStatusOk() (*ArchiveStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventArchiveResponse) SetStatus(v ArchiveStatus)`

SetStatus sets Status field to given value.


### GetBefore

`func (o *EventArchiveResponse) GetBefore() int64`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *EventArchiveResponse) GetBeforeOk() (*int64, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *EventArchiveResponse) SetBefore(v int64)`

SetBefore sets Before field to given value.


### GetFilename

`func (o *EventArchiveResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *EventArchiveResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *EventArchiveResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetName

`func (o *EventArchiveResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventArchiveResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventArchiveResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EventArchiveResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventArchiveResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventArchiveResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


