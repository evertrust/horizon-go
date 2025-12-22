# WebRAEnrollRequestTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**NullableManagedCertificateProfileCryptoPolicy**](ManagedCertificateProfileCryptoPolicy.md) | Describes how certificates will be enrolled on this profile | [optional] 
**PasswordPolicy** | Pointer to [**NullablePasswordPolicy**](PasswordPolicy.md) | The password policy that will be used to generate the certificate&#39;s PKCS#12 password | [optional] 
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElement**](CertificateContactEmailElement.md) | Information about the certificate&#39;s contact email and how to edit it | [optional] 
**Extensions** | Pointer to [**[]CertificateExtensionElement**](CertificateExtensionElement.md) | Information about the certificate&#39;s extensions and how to edit them | [optional] 
**Labels** | Pointer to [**[]RequestLabelElement**](RequestLabelElement.md) | List of labels used internally to tag and group certificates | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadataElement**](CertificateMetadataElement.md) | The technical metadata for this certificate | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElement**](CertificateOwnerElement.md) | Information about the certificate&#39;s owner and how to edit it | [optional] 
**Sans** | Pointer to [**[]ListSANElement**](ListSANElement.md) | List of SAN elements that will be used to build the certificate&#39;s Subject Alternative Name | [optional] 
**Subject** | Pointer to [**[]IndexedDNElement**](IndexedDNElement.md) | List of DN elements that will be used to build the certificate&#39;s Distinguished Name | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElement**](CertificateTeamElement.md) | Information about the certificate&#39;s team and how to edit it | [optional] 

## Methods

### NewWebRAEnrollRequestTemplateResponse

`func NewWebRAEnrollRequestTemplateResponse() *WebRAEnrollRequestTemplateResponse`

NewWebRAEnrollRequestTemplateResponse instantiates a new WebRAEnrollRequestTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAEnrollRequestTemplateResponseWithDefaults

`func NewWebRAEnrollRequestTemplateResponseWithDefaults() *WebRAEnrollRequestTemplateResponse`

NewWebRAEnrollRequestTemplateResponseWithDefaults instantiates a new WebRAEnrollRequestTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *WebRAEnrollRequestTemplateResponse) GetCapabilities() ManagedCertificateProfileCryptoPolicy`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *WebRAEnrollRequestTemplateResponse) GetCapabilitiesOk() (*ManagedCertificateProfileCryptoPolicy, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *WebRAEnrollRequestTemplateResponse) SetCapabilities(v ManagedCertificateProfileCryptoPolicy)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *WebRAEnrollRequestTemplateResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *WebRAEnrollRequestTemplateResponse) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *WebRAEnrollRequestTemplateResponse) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetPasswordPolicy

`func (o *WebRAEnrollRequestTemplateResponse) GetPasswordPolicy() PasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *WebRAEnrollRequestTemplateResponse) GetPasswordPolicyOk() (*PasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *WebRAEnrollRequestTemplateResponse) SetPasswordPolicy(v PasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *WebRAEnrollRequestTemplateResponse) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *WebRAEnrollRequestTemplateResponse) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *WebRAEnrollRequestTemplateResponse) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetContactEmail

`func (o *WebRAEnrollRequestTemplateResponse) GetContactEmail() CertificateContactEmailElement`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *WebRAEnrollRequestTemplateResponse) GetContactEmailOk() (*CertificateContactEmailElement, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *WebRAEnrollRequestTemplateResponse) SetContactEmail(v CertificateContactEmailElement)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *WebRAEnrollRequestTemplateResponse) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *WebRAEnrollRequestTemplateResponse) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *WebRAEnrollRequestTemplateResponse) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetExtensions

`func (o *WebRAEnrollRequestTemplateResponse) GetExtensions() []CertificateExtensionElement`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *WebRAEnrollRequestTemplateResponse) GetExtensionsOk() (*[]CertificateExtensionElement, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *WebRAEnrollRequestTemplateResponse) SetExtensions(v []CertificateExtensionElement)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *WebRAEnrollRequestTemplateResponse) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *WebRAEnrollRequestTemplateResponse) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *WebRAEnrollRequestTemplateResponse) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetLabels

`func (o *WebRAEnrollRequestTemplateResponse) GetLabels() []RequestLabelElement`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRAEnrollRequestTemplateResponse) GetLabelsOk() (*[]RequestLabelElement, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRAEnrollRequestTemplateResponse) SetLabels(v []RequestLabelElement)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRAEnrollRequestTemplateResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRAEnrollRequestTemplateResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRAEnrollRequestTemplateResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *WebRAEnrollRequestTemplateResponse) GetMetadata() []CertificateMetadataElement`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAEnrollRequestTemplateResponse) GetMetadataOk() (*[]CertificateMetadataElement, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAEnrollRequestTemplateResponse) SetMetadata(v []CertificateMetadataElement)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAEnrollRequestTemplateResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRAEnrollRequestTemplateResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRAEnrollRequestTemplateResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOwner

`func (o *WebRAEnrollRequestTemplateResponse) GetOwner() CertificateOwnerElement`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WebRAEnrollRequestTemplateResponse) GetOwnerOk() (*CertificateOwnerElement, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WebRAEnrollRequestTemplateResponse) SetOwner(v CertificateOwnerElement)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WebRAEnrollRequestTemplateResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WebRAEnrollRequestTemplateResponse) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WebRAEnrollRequestTemplateResponse) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSans

`func (o *WebRAEnrollRequestTemplateResponse) GetSans() []ListSANElement`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *WebRAEnrollRequestTemplateResponse) GetSansOk() (*[]ListSANElement, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *WebRAEnrollRequestTemplateResponse) SetSans(v []ListSANElement)`

SetSans sets Sans field to given value.

### HasSans

`func (o *WebRAEnrollRequestTemplateResponse) HasSans() bool`

HasSans returns a boolean if a field has been set.

### SetSansNil

`func (o *WebRAEnrollRequestTemplateResponse) SetSansNil(b bool)`

 SetSansNil sets the value for Sans to be an explicit nil

### UnsetSans
`func (o *WebRAEnrollRequestTemplateResponse) UnsetSans()`

UnsetSans ensures that no value is present for Sans, not even an explicit nil
### GetSubject

`func (o *WebRAEnrollRequestTemplateResponse) GetSubject() []IndexedDNElement`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WebRAEnrollRequestTemplateResponse) GetSubjectOk() (*[]IndexedDNElement, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WebRAEnrollRequestTemplateResponse) SetSubject(v []IndexedDNElement)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WebRAEnrollRequestTemplateResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *WebRAEnrollRequestTemplateResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *WebRAEnrollRequestTemplateResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTeam

`func (o *WebRAEnrollRequestTemplateResponse) GetTeam() CertificateTeamElement`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRAEnrollRequestTemplateResponse) GetTeamOk() (*CertificateTeamElement, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRAEnrollRequestTemplateResponse) SetTeam(v CertificateTeamElement)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRAEnrollRequestTemplateResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRAEnrollRequestTemplateResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRAEnrollRequestTemplateResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


