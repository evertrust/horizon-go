# WebRAUpdateRequestTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElementResponse**](CertificateContactEmailElementResponse.md) | Information about the certificate&#39;s contact email and how to edit it | [optional] 
**Labels** | Pointer to [**[]RequestLabelElementResponse**](RequestLabelElementResponse.md) | Information about the certificate&#39;s labels and how to edit them | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadataElementResponse**](CertificateMetadataElementResponse.md) | Information about the certificate&#39;s metadata and how to edit them | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElementResponse**](CertificateOwnerElementResponse.md) | Information about the certificate&#39;s owner and how to edit it | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElementResponse**](CertificateTeamElementResponse.md) | Information about the certificate&#39;s team and how to edit it | [optional] 

## Methods

### NewWebRAUpdateRequestTemplateResponse

`func NewWebRAUpdateRequestTemplateResponse() *WebRAUpdateRequestTemplateResponse`

NewWebRAUpdateRequestTemplateResponse instantiates a new WebRAUpdateRequestTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAUpdateRequestTemplateResponseWithDefaults

`func NewWebRAUpdateRequestTemplateResponseWithDefaults() *WebRAUpdateRequestTemplateResponse`

NewWebRAUpdateRequestTemplateResponseWithDefaults instantiates a new WebRAUpdateRequestTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *WebRAUpdateRequestTemplateResponse) GetContactEmail() CertificateContactEmailElementResponse`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *WebRAUpdateRequestTemplateResponse) GetContactEmailOk() (*CertificateContactEmailElementResponse, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *WebRAUpdateRequestTemplateResponse) SetContactEmail(v CertificateContactEmailElementResponse)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *WebRAUpdateRequestTemplateResponse) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *WebRAUpdateRequestTemplateResponse) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *WebRAUpdateRequestTemplateResponse) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetLabels

`func (o *WebRAUpdateRequestTemplateResponse) GetLabels() []RequestLabelElementResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRAUpdateRequestTemplateResponse) GetLabelsOk() (*[]RequestLabelElementResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRAUpdateRequestTemplateResponse) SetLabels(v []RequestLabelElementResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRAUpdateRequestTemplateResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRAUpdateRequestTemplateResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRAUpdateRequestTemplateResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *WebRAUpdateRequestTemplateResponse) GetMetadata() []CertificateMetadataElementResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAUpdateRequestTemplateResponse) GetMetadataOk() (*[]CertificateMetadataElementResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAUpdateRequestTemplateResponse) SetMetadata(v []CertificateMetadataElementResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAUpdateRequestTemplateResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRAUpdateRequestTemplateResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRAUpdateRequestTemplateResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOwner

`func (o *WebRAUpdateRequestTemplateResponse) GetOwner() CertificateOwnerElementResponse`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WebRAUpdateRequestTemplateResponse) GetOwnerOk() (*CertificateOwnerElementResponse, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WebRAUpdateRequestTemplateResponse) SetOwner(v CertificateOwnerElementResponse)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WebRAUpdateRequestTemplateResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WebRAUpdateRequestTemplateResponse) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WebRAUpdateRequestTemplateResponse) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTeam

`func (o *WebRAUpdateRequestTemplateResponse) GetTeam() CertificateTeamElementResponse`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRAUpdateRequestTemplateResponse) GetTeamOk() (*CertificateTeamElementResponse, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRAUpdateRequestTemplateResponse) SetTeam(v CertificateTeamElementResponse)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRAUpdateRequestTemplateResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRAUpdateRequestTemplateResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRAUpdateRequestTemplateResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


