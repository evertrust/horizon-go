# CertificateTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to [**[]DNElement**](DNElement.md) |  | [optional] 
**Sans** | Pointer to [**[]SANElement**](SANElement.md) |  | [optional] 
**Extensions** | Pointer to [**[]ExtensionElement**](ExtensionElement.md) |  | [optional] 
**OwnerPolicy** | Pointer to [**NullableOwnerPolicy**](OwnerPolicy.md) |  | [optional] 
**TeamPolicy** | Pointer to [**NullableTeamPolicy**](TeamPolicy.md) |  | [optional] 
**MetadataPolicies** | Pointer to [**[]MetadataPolicy**](MetadataPolicy.md) |  | [optional] 
**Labels** | Pointer to [**[]LabelElement**](LabelElement.md) |  | [optional] 
**ContactEmailPolicy** | Pointer to [**NullableContactEmailPolicy**](ContactEmailPolicy.md) |  | [optional] 

## Methods

### NewCertificateTemplate

`func NewCertificateTemplate() *CertificateTemplate`

NewCertificateTemplate instantiates a new CertificateTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateTemplateWithDefaults

`func NewCertificateTemplateWithDefaults() *CertificateTemplate`

NewCertificateTemplateWithDefaults instantiates a new CertificateTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *CertificateTemplate) GetSubject() []DNElement`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateTemplate) GetSubjectOk() (*[]DNElement, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateTemplate) SetSubject(v []DNElement)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertificateTemplate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *CertificateTemplate) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *CertificateTemplate) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetSans

`func (o *CertificateTemplate) GetSans() []SANElement`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *CertificateTemplate) GetSansOk() (*[]SANElement, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *CertificateTemplate) SetSans(v []SANElement)`

SetSans sets Sans field to given value.

### HasSans

`func (o *CertificateTemplate) HasSans() bool`

HasSans returns a boolean if a field has been set.

### SetSansNil

`func (o *CertificateTemplate) SetSansNil(b bool)`

 SetSansNil sets the value for Sans to be an explicit nil

### UnsetSans
`func (o *CertificateTemplate) UnsetSans()`

UnsetSans ensures that no value is present for Sans, not even an explicit nil
### GetExtensions

`func (o *CertificateTemplate) GetExtensions() []ExtensionElement`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CertificateTemplate) GetExtensionsOk() (*[]ExtensionElement, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CertificateTemplate) SetExtensions(v []ExtensionElement)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CertificateTemplate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *CertificateTemplate) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *CertificateTemplate) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetOwnerPolicy

`func (o *CertificateTemplate) GetOwnerPolicy() OwnerPolicy`

GetOwnerPolicy returns the OwnerPolicy field if non-nil, zero value otherwise.

### GetOwnerPolicyOk

`func (o *CertificateTemplate) GetOwnerPolicyOk() (*OwnerPolicy, bool)`

GetOwnerPolicyOk returns a tuple with the OwnerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPolicy

`func (o *CertificateTemplate) SetOwnerPolicy(v OwnerPolicy)`

SetOwnerPolicy sets OwnerPolicy field to given value.

### HasOwnerPolicy

`func (o *CertificateTemplate) HasOwnerPolicy() bool`

HasOwnerPolicy returns a boolean if a field has been set.

### SetOwnerPolicyNil

`func (o *CertificateTemplate) SetOwnerPolicyNil(b bool)`

 SetOwnerPolicyNil sets the value for OwnerPolicy to be an explicit nil

### UnsetOwnerPolicy
`func (o *CertificateTemplate) UnsetOwnerPolicy()`

UnsetOwnerPolicy ensures that no value is present for OwnerPolicy, not even an explicit nil
### GetTeamPolicy

`func (o *CertificateTemplate) GetTeamPolicy() TeamPolicy`

GetTeamPolicy returns the TeamPolicy field if non-nil, zero value otherwise.

### GetTeamPolicyOk

`func (o *CertificateTemplate) GetTeamPolicyOk() (*TeamPolicy, bool)`

GetTeamPolicyOk returns a tuple with the TeamPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamPolicy

`func (o *CertificateTemplate) SetTeamPolicy(v TeamPolicy)`

SetTeamPolicy sets TeamPolicy field to given value.

### HasTeamPolicy

`func (o *CertificateTemplate) HasTeamPolicy() bool`

HasTeamPolicy returns a boolean if a field has been set.

### SetTeamPolicyNil

`func (o *CertificateTemplate) SetTeamPolicyNil(b bool)`

 SetTeamPolicyNil sets the value for TeamPolicy to be an explicit nil

### UnsetTeamPolicy
`func (o *CertificateTemplate) UnsetTeamPolicy()`

UnsetTeamPolicy ensures that no value is present for TeamPolicy, not even an explicit nil
### GetMetadataPolicies

`func (o *CertificateTemplate) GetMetadataPolicies() []MetadataPolicy`

GetMetadataPolicies returns the MetadataPolicies field if non-nil, zero value otherwise.

### GetMetadataPoliciesOk

`func (o *CertificateTemplate) GetMetadataPoliciesOk() (*[]MetadataPolicy, bool)`

GetMetadataPoliciesOk returns a tuple with the MetadataPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPolicies

`func (o *CertificateTemplate) SetMetadataPolicies(v []MetadataPolicy)`

SetMetadataPolicies sets MetadataPolicies field to given value.

### HasMetadataPolicies

`func (o *CertificateTemplate) HasMetadataPolicies() bool`

HasMetadataPolicies returns a boolean if a field has been set.

### SetMetadataPoliciesNil

`func (o *CertificateTemplate) SetMetadataPoliciesNil(b bool)`

 SetMetadataPoliciesNil sets the value for MetadataPolicies to be an explicit nil

### UnsetMetadataPolicies
`func (o *CertificateTemplate) UnsetMetadataPolicies()`

UnsetMetadataPolicies ensures that no value is present for MetadataPolicies, not even an explicit nil
### GetLabels

`func (o *CertificateTemplate) GetLabels() []LabelElement`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CertificateTemplate) GetLabelsOk() (*[]LabelElement, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CertificateTemplate) SetLabels(v []LabelElement)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CertificateTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *CertificateTemplate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *CertificateTemplate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetContactEmailPolicy

`func (o *CertificateTemplate) GetContactEmailPolicy() ContactEmailPolicy`

GetContactEmailPolicy returns the ContactEmailPolicy field if non-nil, zero value otherwise.

### GetContactEmailPolicyOk

`func (o *CertificateTemplate) GetContactEmailPolicyOk() (*ContactEmailPolicy, bool)`

GetContactEmailPolicyOk returns a tuple with the ContactEmailPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmailPolicy

`func (o *CertificateTemplate) SetContactEmailPolicy(v ContactEmailPolicy)`

SetContactEmailPolicy sets ContactEmailPolicy field to given value.

### HasContactEmailPolicy

`func (o *CertificateTemplate) HasContactEmailPolicy() bool`

HasContactEmailPolicy returns a boolean if a field has been set.

### SetContactEmailPolicyNil

`func (o *CertificateTemplate) SetContactEmailPolicyNil(b bool)`

 SetContactEmailPolicyNil sets the value for ContactEmailPolicy to be an explicit nil

### UnsetContactEmailPolicy
`func (o *CertificateTemplate) UnsetContactEmailPolicy()`

UnsetContactEmailPolicy ensures that no value is present for ContactEmailPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


