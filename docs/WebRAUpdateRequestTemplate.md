# WebRAUpdateRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElement**](CertificateContactEmailElement.md) | Information about the certificate&#39;s contact email and how to edit it | [optional] 
**Labels** | Pointer to [**[]RequestLabelElement**](RequestLabelElement.md) | Information about the certificate&#39;s labels and how to edit them | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadataElement**](CertificateMetadataElement.md) | Information about the certificate&#39;s metadata and how to edit them | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElement**](CertificateOwnerElement.md) | Information about the certificate&#39;s owner and how to edit it | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElement**](CertificateTeamElement.md) | Information about the certificate&#39;s team and how to edit it | [optional] 

## Methods

### NewWebRAUpdateRequestTemplate

`func NewWebRAUpdateRequestTemplate() *WebRAUpdateRequestTemplate`

NewWebRAUpdateRequestTemplate instantiates a new WebRAUpdateRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAUpdateRequestTemplateWithDefaults

`func NewWebRAUpdateRequestTemplateWithDefaults() *WebRAUpdateRequestTemplate`

NewWebRAUpdateRequestTemplateWithDefaults instantiates a new WebRAUpdateRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *WebRAUpdateRequestTemplate) GetContactEmail() CertificateContactEmailElement`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *WebRAUpdateRequestTemplate) GetContactEmailOk() (*CertificateContactEmailElement, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *WebRAUpdateRequestTemplate) SetContactEmail(v CertificateContactEmailElement)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *WebRAUpdateRequestTemplate) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *WebRAUpdateRequestTemplate) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *WebRAUpdateRequestTemplate) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetLabels

`func (o *WebRAUpdateRequestTemplate) GetLabels() []RequestLabelElement`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRAUpdateRequestTemplate) GetLabelsOk() (*[]RequestLabelElement, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRAUpdateRequestTemplate) SetLabels(v []RequestLabelElement)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRAUpdateRequestTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRAUpdateRequestTemplate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRAUpdateRequestTemplate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *WebRAUpdateRequestTemplate) GetMetadata() []CertificateMetadataElement`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAUpdateRequestTemplate) GetMetadataOk() (*[]CertificateMetadataElement, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAUpdateRequestTemplate) SetMetadata(v []CertificateMetadataElement)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAUpdateRequestTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRAUpdateRequestTemplate) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRAUpdateRequestTemplate) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOwner

`func (o *WebRAUpdateRequestTemplate) GetOwner() CertificateOwnerElement`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WebRAUpdateRequestTemplate) GetOwnerOk() (*CertificateOwnerElement, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WebRAUpdateRequestTemplate) SetOwner(v CertificateOwnerElement)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WebRAUpdateRequestTemplate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WebRAUpdateRequestTemplate) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WebRAUpdateRequestTemplate) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTeam

`func (o *WebRAUpdateRequestTemplate) GetTeam() CertificateTeamElement`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRAUpdateRequestTemplate) GetTeamOk() (*CertificateTeamElement, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRAUpdateRequestTemplate) SetTeam(v CertificateTeamElement)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRAUpdateRequestTemplate) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRAUpdateRequestTemplate) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRAUpdateRequestTemplate) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


