# ScepEnrollRequestTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnWhitelist** | Pointer to **NullableBool** | DN whitelist is enabled on this request | [optional] 
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElement**](CertificateContactEmailElement.md) | Information about the certificate&#39;s contact email and how to edit it | [optional] 
**Extensions** | Pointer to [**[]CertificateExtensionElement**](CertificateExtensionElement.md) | Information about the certificate&#39;s extensions and how to edit them | [optional] 
**Labels** | Pointer to [**[]RequestLabelElement**](RequestLabelElement.md) | List of labels used internally to tag and group certificates | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElement**](CertificateOwnerElement.md) | Information about the certificate&#39;s owner and how to edit it | [optional] 
**Sans** | Pointer to [**[]ListSANElement**](ListSANElement.md) | List of SAN elements that will be used to build the certificate&#39;s Subject Alternative Name | [optional] 
**Subject** | Pointer to [**[]IndexedDNElement**](IndexedDNElement.md) | List of DN elements that will be used to build the certificate&#39;s Distinguished Name | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElement**](CertificateTeamElement.md) | Information about the certificate&#39;s team and how to edit it | [optional] 

## Methods

### NewScepEnrollRequestTemplateResponse

`func NewScepEnrollRequestTemplateResponse() *ScepEnrollRequestTemplateResponse`

NewScepEnrollRequestTemplateResponse instantiates a new ScepEnrollRequestTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScepEnrollRequestTemplateResponseWithDefaults

`func NewScepEnrollRequestTemplateResponseWithDefaults() *ScepEnrollRequestTemplateResponse`

NewScepEnrollRequestTemplateResponseWithDefaults instantiates a new ScepEnrollRequestTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnWhitelist

`func (o *ScepEnrollRequestTemplateResponse) GetDnWhitelist() bool`

GetDnWhitelist returns the DnWhitelist field if non-nil, zero value otherwise.

### GetDnWhitelistOk

`func (o *ScepEnrollRequestTemplateResponse) GetDnWhitelistOk() (*bool, bool)`

GetDnWhitelistOk returns a tuple with the DnWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnWhitelist

`func (o *ScepEnrollRequestTemplateResponse) SetDnWhitelist(v bool)`

SetDnWhitelist sets DnWhitelist field to given value.

### HasDnWhitelist

`func (o *ScepEnrollRequestTemplateResponse) HasDnWhitelist() bool`

HasDnWhitelist returns a boolean if a field has been set.

### SetDnWhitelistNil

`func (o *ScepEnrollRequestTemplateResponse) SetDnWhitelistNil(b bool)`

 SetDnWhitelistNil sets the value for DnWhitelist to be an explicit nil

### UnsetDnWhitelist
`func (o *ScepEnrollRequestTemplateResponse) UnsetDnWhitelist()`

UnsetDnWhitelist ensures that no value is present for DnWhitelist, not even an explicit nil
### GetContactEmail

`func (o *ScepEnrollRequestTemplateResponse) GetContactEmail() CertificateContactEmailElement`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *ScepEnrollRequestTemplateResponse) GetContactEmailOk() (*CertificateContactEmailElement, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *ScepEnrollRequestTemplateResponse) SetContactEmail(v CertificateContactEmailElement)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *ScepEnrollRequestTemplateResponse) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *ScepEnrollRequestTemplateResponse) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *ScepEnrollRequestTemplateResponse) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetExtensions

`func (o *ScepEnrollRequestTemplateResponse) GetExtensions() []CertificateExtensionElement`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ScepEnrollRequestTemplateResponse) GetExtensionsOk() (*[]CertificateExtensionElement, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ScepEnrollRequestTemplateResponse) SetExtensions(v []CertificateExtensionElement)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ScepEnrollRequestTemplateResponse) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *ScepEnrollRequestTemplateResponse) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *ScepEnrollRequestTemplateResponse) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetLabels

`func (o *ScepEnrollRequestTemplateResponse) GetLabels() []RequestLabelElement`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ScepEnrollRequestTemplateResponse) GetLabelsOk() (*[]RequestLabelElement, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ScepEnrollRequestTemplateResponse) SetLabels(v []RequestLabelElement)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ScepEnrollRequestTemplateResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ScepEnrollRequestTemplateResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ScepEnrollRequestTemplateResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetOwner

`func (o *ScepEnrollRequestTemplateResponse) GetOwner() CertificateOwnerElement`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ScepEnrollRequestTemplateResponse) GetOwnerOk() (*CertificateOwnerElement, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ScepEnrollRequestTemplateResponse) SetOwner(v CertificateOwnerElement)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ScepEnrollRequestTemplateResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ScepEnrollRequestTemplateResponse) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ScepEnrollRequestTemplateResponse) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSans

`func (o *ScepEnrollRequestTemplateResponse) GetSans() []ListSANElement`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *ScepEnrollRequestTemplateResponse) GetSansOk() (*[]ListSANElement, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *ScepEnrollRequestTemplateResponse) SetSans(v []ListSANElement)`

SetSans sets Sans field to given value.

### HasSans

`func (o *ScepEnrollRequestTemplateResponse) HasSans() bool`

HasSans returns a boolean if a field has been set.

### SetSansNil

`func (o *ScepEnrollRequestTemplateResponse) SetSansNil(b bool)`

 SetSansNil sets the value for Sans to be an explicit nil

### UnsetSans
`func (o *ScepEnrollRequestTemplateResponse) UnsetSans()`

UnsetSans ensures that no value is present for Sans, not even an explicit nil
### GetSubject

`func (o *ScepEnrollRequestTemplateResponse) GetSubject() []IndexedDNElement`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ScepEnrollRequestTemplateResponse) GetSubjectOk() (*[]IndexedDNElement, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ScepEnrollRequestTemplateResponse) SetSubject(v []IndexedDNElement)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ScepEnrollRequestTemplateResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *ScepEnrollRequestTemplateResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *ScepEnrollRequestTemplateResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTeam

`func (o *ScepEnrollRequestTemplateResponse) GetTeam() CertificateTeamElement`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ScepEnrollRequestTemplateResponse) GetTeamOk() (*CertificateTeamElement, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ScepEnrollRequestTemplateResponse) SetTeam(v CertificateTeamElement)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ScepEnrollRequestTemplateResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ScepEnrollRequestTemplateResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ScepEnrollRequestTemplateResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


