# StorageConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ArchiveStorage** | Pointer to **string** | Name of a [storage](#tag/system.storage) to use for archive file storage | [optional] 
**MagicLinkReportStorage** | Pointer to **string** | Name of a [storage](#tag/system.storage) to use for magic link reports storage | [optional] 
**Type** | **string** | The type of the configuration entry | 

## Methods

### NewStorageConfigurationResponse

`func NewStorageConfigurationResponse(id string, type_ string, ) *StorageConfigurationResponse`

NewStorageConfigurationResponse instantiates a new StorageConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigurationResponseWithDefaults

`func NewStorageConfigurationResponseWithDefaults() *StorageConfigurationResponse`

NewStorageConfigurationResponseWithDefaults instantiates a new StorageConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageConfigurationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageConfigurationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageConfigurationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetArchiveStorage

`func (o *StorageConfigurationResponse) GetArchiveStorage() string`

GetArchiveStorage returns the ArchiveStorage field if non-nil, zero value otherwise.

### GetArchiveStorageOk

`func (o *StorageConfigurationResponse) GetArchiveStorageOk() (*string, bool)`

GetArchiveStorageOk returns a tuple with the ArchiveStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveStorage

`func (o *StorageConfigurationResponse) SetArchiveStorage(v string)`

SetArchiveStorage sets ArchiveStorage field to given value.

### HasArchiveStorage

`func (o *StorageConfigurationResponse) HasArchiveStorage() bool`

HasArchiveStorage returns a boolean if a field has been set.

### GetMagicLinkReportStorage

`func (o *StorageConfigurationResponse) GetMagicLinkReportStorage() string`

GetMagicLinkReportStorage returns the MagicLinkReportStorage field if non-nil, zero value otherwise.

### GetMagicLinkReportStorageOk

`func (o *StorageConfigurationResponse) GetMagicLinkReportStorageOk() (*string, bool)`

GetMagicLinkReportStorageOk returns a tuple with the MagicLinkReportStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicLinkReportStorage

`func (o *StorageConfigurationResponse) SetMagicLinkReportStorage(v string)`

SetMagicLinkReportStorage sets MagicLinkReportStorage field to given value.

### HasMagicLinkReportStorage

`func (o *StorageConfigurationResponse) HasMagicLinkReportStorage() bool`

HasMagicLinkReportStorage returns a boolean if a field has been set.

### GetType

`func (o *StorageConfigurationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageConfigurationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageConfigurationResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


