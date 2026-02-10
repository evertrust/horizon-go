# CertificateArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Filename** | **string** |  | 
**ArchiveKeys** | **bool** |  | 
**Filter** | Pointer to **NullableString** | An HCQL filter for the archive | [optional] 

## Methods

### NewCertificateArchive

`func NewCertificateArchive(name string, type_ string, filename string, archiveKeys bool, ) *CertificateArchive`

NewCertificateArchive instantiates a new CertificateArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateArchiveWithDefaults

`func NewCertificateArchiveWithDefaults() *CertificateArchive`

NewCertificateArchiveWithDefaults instantiates a new CertificateArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateArchive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateArchive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateArchive) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CertificateArchive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateArchive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateArchive) SetType(v string)`

SetType sets Type field to given value.


### GetFilename

`func (o *CertificateArchive) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CertificateArchive) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CertificateArchive) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetArchiveKeys

`func (o *CertificateArchive) GetArchiveKeys() bool`

GetArchiveKeys returns the ArchiveKeys field if non-nil, zero value otherwise.

### GetArchiveKeysOk

`func (o *CertificateArchive) GetArchiveKeysOk() (*bool, bool)`

GetArchiveKeysOk returns a tuple with the ArchiveKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveKeys

`func (o *CertificateArchive) SetArchiveKeys(v bool)`

SetArchiveKeys sets ArchiveKeys field to given value.


### GetFilter

`func (o *CertificateArchive) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CertificateArchive) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CertificateArchive) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CertificateArchive) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *CertificateArchive) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *CertificateArchive) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


