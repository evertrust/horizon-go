# WebRAMigrateRequestTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRenew** | Pointer to [**CertificateAutoRenewElementResponse**](CertificateAutoRenewElementResponse.md) |  | [optional] 
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElementResponse**](CertificateContactEmailElementResponse.md) |  | [optional] 
**Labels** | Pointer to [**[]RequestLabelElementResponse**](RequestLabelElementResponse.md) |  | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadataElementResponse**](CertificateMetadataElementResponse.md) |  | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElementResponse**](CertificateOwnerElementResponse.md) |  | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElementResponse**](CertificateTeamElementResponse.md) |  | [optional] 

## Methods

### NewWebRAMigrateRequestTemplateResponse

`func NewWebRAMigrateRequestTemplateResponse() *WebRAMigrateRequestTemplateResponse`

NewWebRAMigrateRequestTemplateResponse instantiates a new WebRAMigrateRequestTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAMigrateRequestTemplateResponseWithDefaults

`func NewWebRAMigrateRequestTemplateResponseWithDefaults() *WebRAMigrateRequestTemplateResponse`

NewWebRAMigrateRequestTemplateResponseWithDefaults instantiates a new WebRAMigrateRequestTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRenew

`func (o *WebRAMigrateRequestTemplateResponse) GetAutoRenew() CertificateAutoRenewElementResponse`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *WebRAMigrateRequestTemplateResponse) GetAutoRenewOk() (*CertificateAutoRenewElementResponse, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *WebRAMigrateRequestTemplateResponse) SetAutoRenew(v CertificateAutoRenewElementResponse)`

SetAutoRenew sets AutoRenew field to given value.

### HasAutoRenew

`func (o *WebRAMigrateRequestTemplateResponse) HasAutoRenew() bool`

HasAutoRenew returns a boolean if a field has been set.

### GetContactEmail

`func (o *WebRAMigrateRequestTemplateResponse) GetContactEmail() CertificateContactEmailElementResponse`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *WebRAMigrateRequestTemplateResponse) GetContactEmailOk() (*CertificateContactEmailElementResponse, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *WebRAMigrateRequestTemplateResponse) SetContactEmail(v CertificateContactEmailElementResponse)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *WebRAMigrateRequestTemplateResponse) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *WebRAMigrateRequestTemplateResponse) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *WebRAMigrateRequestTemplateResponse) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetLabels

`func (o *WebRAMigrateRequestTemplateResponse) GetLabels() []RequestLabelElementResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRAMigrateRequestTemplateResponse) GetLabelsOk() (*[]RequestLabelElementResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRAMigrateRequestTemplateResponse) SetLabels(v []RequestLabelElementResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRAMigrateRequestTemplateResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRAMigrateRequestTemplateResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRAMigrateRequestTemplateResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *WebRAMigrateRequestTemplateResponse) GetMetadata() []CertificateMetadataElementResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAMigrateRequestTemplateResponse) GetMetadataOk() (*[]CertificateMetadataElementResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAMigrateRequestTemplateResponse) SetMetadata(v []CertificateMetadataElementResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAMigrateRequestTemplateResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRAMigrateRequestTemplateResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRAMigrateRequestTemplateResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOwner

`func (o *WebRAMigrateRequestTemplateResponse) GetOwner() CertificateOwnerElementResponse`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WebRAMigrateRequestTemplateResponse) GetOwnerOk() (*CertificateOwnerElementResponse, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WebRAMigrateRequestTemplateResponse) SetOwner(v CertificateOwnerElementResponse)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WebRAMigrateRequestTemplateResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WebRAMigrateRequestTemplateResponse) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WebRAMigrateRequestTemplateResponse) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTeam

`func (o *WebRAMigrateRequestTemplateResponse) GetTeam() CertificateTeamElementResponse`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRAMigrateRequestTemplateResponse) GetTeamOk() (*CertificateTeamElementResponse, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRAMigrateRequestTemplateResponse) SetTeam(v CertificateTeamElementResponse)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRAMigrateRequestTemplateResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRAMigrateRequestTemplateResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRAMigrateRequestTemplateResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


