# ScepEnrollRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElement**](CertificateContactEmailElement.md) | Information about the certificate&#39;s contact email and how to edit it | [optional] 
**Extensions** | Pointer to [**[]CertificateExtensionElement**](CertificateExtensionElement.md) | Information about the certificate&#39;s extensions and how to edit them | [optional] 
**Labels** | Pointer to [**[]RequestLabelElement**](RequestLabelElement.md) | List of labels used internally to tag and group certificates | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElement**](CertificateOwnerElement.md) | Information about the certificate&#39;s owner and how to edit it | [optional] 
**Sans** | Pointer to [**[]ListSANElement**](ListSANElement.md) | List of SAN elements that will be used to build the certificate&#39;s Subject Alternative Name | [optional] 
**Subject** | Pointer to [**[]IndexedDNElement**](IndexedDNElement.md) | List of DN elements that will be used to build the certificate&#39;s Distinguished Name | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElement**](CertificateTeamElement.md) | Information about the certificate&#39;s team and how to edit it | [optional] 

## Methods

### NewScepEnrollRequestTemplate

`func NewScepEnrollRequestTemplate() *ScepEnrollRequestTemplate`

NewScepEnrollRequestTemplate instantiates a new ScepEnrollRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScepEnrollRequestTemplateWithDefaults

`func NewScepEnrollRequestTemplateWithDefaults() *ScepEnrollRequestTemplate`

NewScepEnrollRequestTemplateWithDefaults instantiates a new ScepEnrollRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *ScepEnrollRequestTemplate) GetContactEmail() CertificateContactEmailElement`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *ScepEnrollRequestTemplate) GetContactEmailOk() (*CertificateContactEmailElement, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *ScepEnrollRequestTemplate) SetContactEmail(v CertificateContactEmailElement)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *ScepEnrollRequestTemplate) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *ScepEnrollRequestTemplate) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *ScepEnrollRequestTemplate) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetExtensions

`func (o *ScepEnrollRequestTemplate) GetExtensions() []CertificateExtensionElement`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ScepEnrollRequestTemplate) GetExtensionsOk() (*[]CertificateExtensionElement, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ScepEnrollRequestTemplate) SetExtensions(v []CertificateExtensionElement)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ScepEnrollRequestTemplate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *ScepEnrollRequestTemplate) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *ScepEnrollRequestTemplate) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetLabels

`func (o *ScepEnrollRequestTemplate) GetLabels() []RequestLabelElement`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ScepEnrollRequestTemplate) GetLabelsOk() (*[]RequestLabelElement, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ScepEnrollRequestTemplate) SetLabels(v []RequestLabelElement)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ScepEnrollRequestTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ScepEnrollRequestTemplate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ScepEnrollRequestTemplate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetOwner

`func (o *ScepEnrollRequestTemplate) GetOwner() CertificateOwnerElement`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ScepEnrollRequestTemplate) GetOwnerOk() (*CertificateOwnerElement, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ScepEnrollRequestTemplate) SetOwner(v CertificateOwnerElement)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ScepEnrollRequestTemplate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ScepEnrollRequestTemplate) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ScepEnrollRequestTemplate) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSans

`func (o *ScepEnrollRequestTemplate) GetSans() []ListSANElement`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *ScepEnrollRequestTemplate) GetSansOk() (*[]ListSANElement, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *ScepEnrollRequestTemplate) SetSans(v []ListSANElement)`

SetSans sets Sans field to given value.

### HasSans

`func (o *ScepEnrollRequestTemplate) HasSans() bool`

HasSans returns a boolean if a field has been set.

### SetSansNil

`func (o *ScepEnrollRequestTemplate) SetSansNil(b bool)`

 SetSansNil sets the value for Sans to be an explicit nil

### UnsetSans
`func (o *ScepEnrollRequestTemplate) UnsetSans()`

UnsetSans ensures that no value is present for Sans, not even an explicit nil
### GetSubject

`func (o *ScepEnrollRequestTemplate) GetSubject() []IndexedDNElement`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ScepEnrollRequestTemplate) GetSubjectOk() (*[]IndexedDNElement, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ScepEnrollRequestTemplate) SetSubject(v []IndexedDNElement)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ScepEnrollRequestTemplate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *ScepEnrollRequestTemplate) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *ScepEnrollRequestTemplate) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTeam

`func (o *ScepEnrollRequestTemplate) GetTeam() CertificateTeamElement`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ScepEnrollRequestTemplate) GetTeamOk() (*CertificateTeamElement, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ScepEnrollRequestTemplate) SetTeam(v CertificateTeamElement)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ScepEnrollRequestTemplate) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ScepEnrollRequestTemplate) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ScepEnrollRequestTemplate) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


