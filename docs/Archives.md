# Archives

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Filename** | **string** |  | 
**ArchiveKeys** | **bool** |  | 
**Filter** | Pointer to **NullableString** | An HCQL filter for the archive | [optional] 
**Before** | **int64** | Date before which all events will be archived | 

## Methods

### NewArchives

`func NewArchives(name string, type_ string, filename string, archiveKeys bool, before int64, ) *Archives`

NewArchives instantiates a new Archives object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivesWithDefaults

`func NewArchivesWithDefaults() *Archives`

NewArchivesWithDefaults instantiates a new Archives object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Archives) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Archives) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Archives) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Archives) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Archives) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Archives) SetType(v string)`

SetType sets Type field to given value.


### GetFilename

`func (o *Archives) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Archives) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Archives) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetArchiveKeys

`func (o *Archives) GetArchiveKeys() bool`

GetArchiveKeys returns the ArchiveKeys field if non-nil, zero value otherwise.

### GetArchiveKeysOk

`func (o *Archives) GetArchiveKeysOk() (*bool, bool)`

GetArchiveKeysOk returns a tuple with the ArchiveKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveKeys

`func (o *Archives) SetArchiveKeys(v bool)`

SetArchiveKeys sets ArchiveKeys field to given value.


### GetFilter

`func (o *Archives) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Archives) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Archives) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Archives) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *Archives) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *Archives) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetBefore

`func (o *Archives) GetBefore() int64`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *Archives) GetBeforeOk() (*int64, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *Archives) SetBefore(v int64)`

SetBefore sets Before field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


