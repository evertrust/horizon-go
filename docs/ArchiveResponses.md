# ArchiveResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Count** | Pointer to **NullableInt64** |  | [optional] 
**CreatedAt** | Pointer to **NullableInt64** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**PurgeAt** | Pointer to **NullableInt64** |  | [optional] 
**Status** | [**ArchiveStatus**](ArchiveStatus.md) |  | 
**ArchiveKeys** | **bool** |  | 
**Filename** | **string** |  | 
**Filter** | Pointer to **NullableString** | An HCQL filter for the archive | [optional] 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Before** | **int64** | Date before which all events will be archived | 

## Methods

### NewArchiveResponses

`func NewArchiveResponses(id string, status ArchiveStatus, archiveKeys bool, filename string, name string, type_ string, before int64, ) *ArchiveResponses`

NewArchiveResponses instantiates a new ArchiveResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveResponsesWithDefaults

`func NewArchiveResponsesWithDefaults() *ArchiveResponses`

NewArchiveResponsesWithDefaults instantiates a new ArchiveResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArchiveResponses) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveResponses) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveResponses) SetId(v string)`

SetId sets Id field to given value.


### GetCount

`func (o *ArchiveResponses) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ArchiveResponses) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ArchiveResponses) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ArchiveResponses) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ArchiveResponses) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ArchiveResponses) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetCreatedAt

`func (o *ArchiveResponses) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArchiveResponses) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArchiveResponses) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArchiveResponses) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ArchiveResponses) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ArchiveResponses) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetError

`func (o *ArchiveResponses) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ArchiveResponses) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ArchiveResponses) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ArchiveResponses) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ArchiveResponses) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ArchiveResponses) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPurgeAt

`func (o *ArchiveResponses) GetPurgeAt() int64`

GetPurgeAt returns the PurgeAt field if non-nil, zero value otherwise.

### GetPurgeAtOk

`func (o *ArchiveResponses) GetPurgeAtOk() (*int64, bool)`

GetPurgeAtOk returns a tuple with the PurgeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeAt

`func (o *ArchiveResponses) SetPurgeAt(v int64)`

SetPurgeAt sets PurgeAt field to given value.

### HasPurgeAt

`func (o *ArchiveResponses) HasPurgeAt() bool`

HasPurgeAt returns a boolean if a field has been set.

### SetPurgeAtNil

`func (o *ArchiveResponses) SetPurgeAtNil(b bool)`

 SetPurgeAtNil sets the value for PurgeAt to be an explicit nil

### UnsetPurgeAt
`func (o *ArchiveResponses) UnsetPurgeAt()`

UnsetPurgeAt ensures that no value is present for PurgeAt, not even an explicit nil
### GetStatus

`func (o *ArchiveResponses) GetStatus() ArchiveStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchiveResponses) GetStatusOk() (*ArchiveStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchiveResponses) SetStatus(v ArchiveStatus)`

SetStatus sets Status field to given value.


### GetArchiveKeys

`func (o *ArchiveResponses) GetArchiveKeys() bool`

GetArchiveKeys returns the ArchiveKeys field if non-nil, zero value otherwise.

### GetArchiveKeysOk

`func (o *ArchiveResponses) GetArchiveKeysOk() (*bool, bool)`

GetArchiveKeysOk returns a tuple with the ArchiveKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveKeys

`func (o *ArchiveResponses) SetArchiveKeys(v bool)`

SetArchiveKeys sets ArchiveKeys field to given value.


### GetFilename

`func (o *ArchiveResponses) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ArchiveResponses) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ArchiveResponses) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetFilter

`func (o *ArchiveResponses) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ArchiveResponses) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ArchiveResponses) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ArchiveResponses) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ArchiveResponses) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ArchiveResponses) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetName

`func (o *ArchiveResponses) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArchiveResponses) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArchiveResponses) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ArchiveResponses) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArchiveResponses) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArchiveResponses) SetType(v string)`

SetType sets Type field to given value.


### GetBefore

`func (o *ArchiveResponses) GetBefore() int64`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ArchiveResponses) GetBeforeOk() (*int64, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ArchiveResponses) SetBefore(v int64)`

SetBefore sets Before field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


