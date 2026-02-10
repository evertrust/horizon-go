# CertificateArchiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**CreatedAt** | Pointer to **NullableInt64** |  | [optional] 
**PurgeAt** | Pointer to **NullableInt64** |  | [optional] 
**Status** | [**ArchiveStatus**](ArchiveStatus.md) |  | 
**Count** | Pointer to **NullableInt64** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Filename** | **string** |  | 
**ArchiveKeys** | **bool** |  | 
**Filter** | Pointer to **NullableString** | An HCQL filter for the archive | [optional] 

## Methods

### NewCertificateArchiveResponse

`func NewCertificateArchiveResponse(id string, status ArchiveStatus, name string, type_ string, filename string, archiveKeys bool, ) *CertificateArchiveResponse`

NewCertificateArchiveResponse instantiates a new CertificateArchiveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateArchiveResponseWithDefaults

`func NewCertificateArchiveResponseWithDefaults() *CertificateArchiveResponse`

NewCertificateArchiveResponseWithDefaults instantiates a new CertificateArchiveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateArchiveResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateArchiveResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateArchiveResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CertificateArchiveResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CertificateArchiveResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CertificateArchiveResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CertificateArchiveResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CertificateArchiveResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CertificateArchiveResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetPurgeAt

`func (o *CertificateArchiveResponse) GetPurgeAt() int64`

GetPurgeAt returns the PurgeAt field if non-nil, zero value otherwise.

### GetPurgeAtOk

`func (o *CertificateArchiveResponse) GetPurgeAtOk() (*int64, bool)`

GetPurgeAtOk returns a tuple with the PurgeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeAt

`func (o *CertificateArchiveResponse) SetPurgeAt(v int64)`

SetPurgeAt sets PurgeAt field to given value.

### HasPurgeAt

`func (o *CertificateArchiveResponse) HasPurgeAt() bool`

HasPurgeAt returns a boolean if a field has been set.

### SetPurgeAtNil

`func (o *CertificateArchiveResponse) SetPurgeAtNil(b bool)`

 SetPurgeAtNil sets the value for PurgeAt to be an explicit nil

### UnsetPurgeAt
`func (o *CertificateArchiveResponse) UnsetPurgeAt()`

UnsetPurgeAt ensures that no value is present for PurgeAt, not even an explicit nil
### GetStatus

`func (o *CertificateArchiveResponse) GetStatus() ArchiveStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateArchiveResponse) GetStatusOk() (*ArchiveStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateArchiveResponse) SetStatus(v ArchiveStatus)`

SetStatus sets Status field to given value.


### GetCount

`func (o *CertificateArchiveResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CertificateArchiveResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CertificateArchiveResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CertificateArchiveResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *CertificateArchiveResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *CertificateArchiveResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetError

`func (o *CertificateArchiveResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CertificateArchiveResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CertificateArchiveResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CertificateArchiveResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *CertificateArchiveResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *CertificateArchiveResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetName

`func (o *CertificateArchiveResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateArchiveResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateArchiveResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CertificateArchiveResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateArchiveResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateArchiveResponse) SetType(v string)`

SetType sets Type field to given value.


### GetFilename

`func (o *CertificateArchiveResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CertificateArchiveResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CertificateArchiveResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetArchiveKeys

`func (o *CertificateArchiveResponse) GetArchiveKeys() bool`

GetArchiveKeys returns the ArchiveKeys field if non-nil, zero value otherwise.

### GetArchiveKeysOk

`func (o *CertificateArchiveResponse) GetArchiveKeysOk() (*bool, bool)`

GetArchiveKeysOk returns a tuple with the ArchiveKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveKeys

`func (o *CertificateArchiveResponse) SetArchiveKeys(v bool)`

SetArchiveKeys sets ArchiveKeys field to given value.


### GetFilter

`func (o *CertificateArchiveResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CertificateArchiveResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CertificateArchiveResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CertificateArchiveResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *CertificateArchiveResponse) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *CertificateArchiveResponse) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


