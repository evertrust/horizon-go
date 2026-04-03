# WebRAEnrollRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRenew** | Pointer to [**CertificateAutoRenewElement**](CertificateAutoRenewElement.md) |  | [optional] 
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElement**](CertificateContactEmailElement.md) | Information about the certificate&#39;s contact email and how to edit it | [optional] 
**Csr** | Pointer to **NullableString** | If decentralized enrollment is enabled, this field will contain the CSR that will be used to generate the certificate | [optional] 
**Extensions** | Pointer to [**[]CertificateExtensionElement**](CertificateExtensionElement.md) | Information about the certificate&#39;s extensions and how to edit them | [optional] 
**KeyType** | Pointer to **NullableString** | The type of key that will be used to generate the certificate, if in centralized mode | [optional] 
**Labels** | Pointer to [**[]RequestLabelElement**](RequestLabelElement.md) | List of labels used internally to tag and group certificates | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadataElement**](CertificateMetadataElement.md) | The technical metadata for this certificate | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElement**](CertificateOwnerElement.md) | Information about the certificate&#39;s owner and how to edit it | [optional] 
**Sans** | Pointer to [**[]ListSANElement**](ListSANElement.md) | List of SAN elements that will be used to build the certificate&#39;s Subject Alternative Name | [optional] 
**Subject** | Pointer to [**[]IndexedDNElement**](IndexedDNElement.md) | List of DN elements that will be used to build the certificate&#39;s Distinguished Name | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElement**](CertificateTeamElement.md) | Information about the certificate&#39;s team and how to edit it | [optional] 

## Methods

### NewWebRAEnrollRequestTemplate

`func NewWebRAEnrollRequestTemplate() *WebRAEnrollRequestTemplate`

NewWebRAEnrollRequestTemplate instantiates a new WebRAEnrollRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAEnrollRequestTemplateWithDefaults

`func NewWebRAEnrollRequestTemplateWithDefaults() *WebRAEnrollRequestTemplate`

NewWebRAEnrollRequestTemplateWithDefaults instantiates a new WebRAEnrollRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRenew

`func (o *WebRAEnrollRequestTemplate) GetAutoRenew() CertificateAutoRenewElement`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *WebRAEnrollRequestTemplate) GetAutoRenewOk() (*CertificateAutoRenewElement, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *WebRAEnrollRequestTemplate) SetAutoRenew(v CertificateAutoRenewElement)`

SetAutoRenew sets AutoRenew field to given value.

### HasAutoRenew

`func (o *WebRAEnrollRequestTemplate) HasAutoRenew() bool`

HasAutoRenew returns a boolean if a field has been set.

### GetContactEmail

`func (o *WebRAEnrollRequestTemplate) GetContactEmail() CertificateContactEmailElement`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *WebRAEnrollRequestTemplate) GetContactEmailOk() (*CertificateContactEmailElement, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *WebRAEnrollRequestTemplate) SetContactEmail(v CertificateContactEmailElement)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *WebRAEnrollRequestTemplate) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *WebRAEnrollRequestTemplate) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *WebRAEnrollRequestTemplate) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetCsr

`func (o *WebRAEnrollRequestTemplate) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *WebRAEnrollRequestTemplate) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *WebRAEnrollRequestTemplate) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *WebRAEnrollRequestTemplate) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *WebRAEnrollRequestTemplate) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *WebRAEnrollRequestTemplate) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil
### GetExtensions

`func (o *WebRAEnrollRequestTemplate) GetExtensions() []CertificateExtensionElement`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *WebRAEnrollRequestTemplate) GetExtensionsOk() (*[]CertificateExtensionElement, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *WebRAEnrollRequestTemplate) SetExtensions(v []CertificateExtensionElement)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *WebRAEnrollRequestTemplate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *WebRAEnrollRequestTemplate) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *WebRAEnrollRequestTemplate) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetKeyType

`func (o *WebRAEnrollRequestTemplate) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *WebRAEnrollRequestTemplate) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *WebRAEnrollRequestTemplate) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *WebRAEnrollRequestTemplate) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *WebRAEnrollRequestTemplate) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *WebRAEnrollRequestTemplate) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetLabels

`func (o *WebRAEnrollRequestTemplate) GetLabels() []RequestLabelElement`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRAEnrollRequestTemplate) GetLabelsOk() (*[]RequestLabelElement, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRAEnrollRequestTemplate) SetLabels(v []RequestLabelElement)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRAEnrollRequestTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRAEnrollRequestTemplate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRAEnrollRequestTemplate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *WebRAEnrollRequestTemplate) GetMetadata() []CertificateMetadataElement`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAEnrollRequestTemplate) GetMetadataOk() (*[]CertificateMetadataElement, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAEnrollRequestTemplate) SetMetadata(v []CertificateMetadataElement)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAEnrollRequestTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRAEnrollRequestTemplate) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRAEnrollRequestTemplate) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOwner

`func (o *WebRAEnrollRequestTemplate) GetOwner() CertificateOwnerElement`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WebRAEnrollRequestTemplate) GetOwnerOk() (*CertificateOwnerElement, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WebRAEnrollRequestTemplate) SetOwner(v CertificateOwnerElement)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WebRAEnrollRequestTemplate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WebRAEnrollRequestTemplate) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WebRAEnrollRequestTemplate) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSans

`func (o *WebRAEnrollRequestTemplate) GetSans() []ListSANElement`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *WebRAEnrollRequestTemplate) GetSansOk() (*[]ListSANElement, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *WebRAEnrollRequestTemplate) SetSans(v []ListSANElement)`

SetSans sets Sans field to given value.

### HasSans

`func (o *WebRAEnrollRequestTemplate) HasSans() bool`

HasSans returns a boolean if a field has been set.

### SetSansNil

`func (o *WebRAEnrollRequestTemplate) SetSansNil(b bool)`

 SetSansNil sets the value for Sans to be an explicit nil

### UnsetSans
`func (o *WebRAEnrollRequestTemplate) UnsetSans()`

UnsetSans ensures that no value is present for Sans, not even an explicit nil
### GetSubject

`func (o *WebRAEnrollRequestTemplate) GetSubject() []IndexedDNElement`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WebRAEnrollRequestTemplate) GetSubjectOk() (*[]IndexedDNElement, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WebRAEnrollRequestTemplate) SetSubject(v []IndexedDNElement)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WebRAEnrollRequestTemplate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *WebRAEnrollRequestTemplate) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *WebRAEnrollRequestTemplate) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTeam

`func (o *WebRAEnrollRequestTemplate) GetTeam() CertificateTeamElement`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRAEnrollRequestTemplate) GetTeamOk() (*CertificateTeamElement, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRAEnrollRequestTemplate) SetTeam(v CertificateTeamElement)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRAEnrollRequestTemplate) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRAEnrollRequestTemplate) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRAEnrollRequestTemplate) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


