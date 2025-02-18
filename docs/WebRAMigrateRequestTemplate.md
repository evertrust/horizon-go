# WebRAMigrateRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]RequestLabelElement**](RequestLabelElement.md) |  | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElement**](CertificateOwnerElement.md) |  | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElement**](CertificateTeamElement.md) |  | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadataElement**](CertificateMetadataElement.md) |  | [optional] 
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElement**](CertificateContactEmailElement.md) |  | [optional] 

## Methods

### NewWebRAMigrateRequestTemplate

`func NewWebRAMigrateRequestTemplate() *WebRAMigrateRequestTemplate`

NewWebRAMigrateRequestTemplate instantiates a new WebRAMigrateRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAMigrateRequestTemplateWithDefaults

`func NewWebRAMigrateRequestTemplateWithDefaults() *WebRAMigrateRequestTemplate`

NewWebRAMigrateRequestTemplateWithDefaults instantiates a new WebRAMigrateRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *WebRAMigrateRequestTemplate) GetLabels() []RequestLabelElement`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRAMigrateRequestTemplate) GetLabelsOk() (*[]RequestLabelElement, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRAMigrateRequestTemplate) SetLabels(v []RequestLabelElement)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRAMigrateRequestTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRAMigrateRequestTemplate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRAMigrateRequestTemplate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetOwner

`func (o *WebRAMigrateRequestTemplate) GetOwner() CertificateOwnerElement`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WebRAMigrateRequestTemplate) GetOwnerOk() (*CertificateOwnerElement, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WebRAMigrateRequestTemplate) SetOwner(v CertificateOwnerElement)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WebRAMigrateRequestTemplate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WebRAMigrateRequestTemplate) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WebRAMigrateRequestTemplate) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTeam

`func (o *WebRAMigrateRequestTemplate) GetTeam() CertificateTeamElement`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRAMigrateRequestTemplate) GetTeamOk() (*CertificateTeamElement, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRAMigrateRequestTemplate) SetTeam(v CertificateTeamElement)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRAMigrateRequestTemplate) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRAMigrateRequestTemplate) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRAMigrateRequestTemplate) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetMetadata

`func (o *WebRAMigrateRequestTemplate) GetMetadata() []CertificateMetadataElement`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAMigrateRequestTemplate) GetMetadataOk() (*[]CertificateMetadataElement, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAMigrateRequestTemplate) SetMetadata(v []CertificateMetadataElement)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAMigrateRequestTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRAMigrateRequestTemplate) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRAMigrateRequestTemplate) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetContactEmail

`func (o *WebRAMigrateRequestTemplate) GetContactEmail() CertificateContactEmailElement`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *WebRAMigrateRequestTemplate) GetContactEmailOk() (*CertificateContactEmailElement, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *WebRAMigrateRequestTemplate) SetContactEmail(v CertificateContactEmailElement)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *WebRAMigrateRequestTemplate) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *WebRAMigrateRequestTemplate) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *WebRAMigrateRequestTemplate) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


