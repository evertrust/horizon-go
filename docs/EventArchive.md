# EventArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | **int64** | Date before which all events will be archived | 
**Filename** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewEventArchive

`func NewEventArchive(before int64, filename string, name string, type_ string, ) *EventArchive`

NewEventArchive instantiates a new EventArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventArchiveWithDefaults

`func NewEventArchiveWithDefaults() *EventArchive`

NewEventArchiveWithDefaults instantiates a new EventArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *EventArchive) GetBefore() int64`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *EventArchive) GetBeforeOk() (*int64, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *EventArchive) SetBefore(v int64)`

SetBefore sets Before field to given value.


### GetFilename

`func (o *EventArchive) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *EventArchive) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *EventArchive) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetName

`func (o *EventArchive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventArchive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventArchive) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EventArchive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventArchive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventArchive) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


