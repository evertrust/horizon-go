# StorageConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveStorage** | Pointer to **string** | Name of a [storage](#tag/system.storage) to use for archive file storage | [optional] 
**MagicLinkReportStorage** | Pointer to **string** | Name of a [storage](#tag/system.storage) to use for magic link reports storage | [optional] 
**Type** | **string** | The type of the configuration entry | 

## Methods

### NewStorageConfiguration

`func NewStorageConfiguration(type_ string, ) *StorageConfiguration`

NewStorageConfiguration instantiates a new StorageConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigurationWithDefaults

`func NewStorageConfigurationWithDefaults() *StorageConfiguration`

NewStorageConfigurationWithDefaults instantiates a new StorageConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveStorage

`func (o *StorageConfiguration) GetArchiveStorage() string`

GetArchiveStorage returns the ArchiveStorage field if non-nil, zero value otherwise.

### GetArchiveStorageOk

`func (o *StorageConfiguration) GetArchiveStorageOk() (*string, bool)`

GetArchiveStorageOk returns a tuple with the ArchiveStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveStorage

`func (o *StorageConfiguration) SetArchiveStorage(v string)`

SetArchiveStorage sets ArchiveStorage field to given value.

### HasArchiveStorage

`func (o *StorageConfiguration) HasArchiveStorage() bool`

HasArchiveStorage returns a boolean if a field has been set.

### GetMagicLinkReportStorage

`func (o *StorageConfiguration) GetMagicLinkReportStorage() string`

GetMagicLinkReportStorage returns the MagicLinkReportStorage field if non-nil, zero value otherwise.

### GetMagicLinkReportStorageOk

`func (o *StorageConfiguration) GetMagicLinkReportStorageOk() (*string, bool)`

GetMagicLinkReportStorageOk returns a tuple with the MagicLinkReportStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicLinkReportStorage

`func (o *StorageConfiguration) SetMagicLinkReportStorage(v string)`

SetMagicLinkReportStorage sets MagicLinkReportStorage field to given value.

### HasMagicLinkReportStorage

`func (o *StorageConfiguration) HasMagicLinkReportStorage() bool`

HasMagicLinkReportStorage returns a boolean if a field has been set.

### GetType

`func (o *StorageConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageConfiguration) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


