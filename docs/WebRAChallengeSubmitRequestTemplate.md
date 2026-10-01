# WebRAChallengeSubmitRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csr** | Pointer to **string** | The certificate signing request to enroll, in decentralized mode. Mutually exclusive with &#x60;keyType&#x60; | [optional] 
**Extensions** | Pointer to [**[]CertificateExtensionElement**](CertificateExtensionElement.md) | List of extension elements that will be used to build the certificate&#39;s extensions. Only accepted if the profile&#39;s certificate template is empty | [optional] 
**KeyType** | Pointer to **string** | The type of key that will be generated, in centralized mode. Mutually exclusive with &#x60;csr&#x60; | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadataElement**](CertificateMetadataElement.md) | List of metadata elements to set on the certificate. Only the &#x60;automation_policy&#x60; metadata may be set here, and only to a policy authorized on the profile | [optional] 
**Sans** | Pointer to [**[]ListSANElement**](ListSANElement.md) | List of SAN elements that will be used to build the certificate&#39;s Subject Alternative Name. Only accepted if the profile&#39;s certificate template is empty | [optional] 
**Subject** | Pointer to [**[]IndexedDNElement**](IndexedDNElement.md) | List of DN elements that will be used to build the certificate&#39;s Distinguished Name. Only accepted if the profile&#39;s certificate template is empty | [optional] 

## Methods

### NewWebRAChallengeSubmitRequestTemplate

`func NewWebRAChallengeSubmitRequestTemplate() *WebRAChallengeSubmitRequestTemplate`

NewWebRAChallengeSubmitRequestTemplate instantiates a new WebRAChallengeSubmitRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAChallengeSubmitRequestTemplateWithDefaults

`func NewWebRAChallengeSubmitRequestTemplateWithDefaults() *WebRAChallengeSubmitRequestTemplate`

NewWebRAChallengeSubmitRequestTemplateWithDefaults instantiates a new WebRAChallengeSubmitRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsr

`func (o *WebRAChallengeSubmitRequestTemplate) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *WebRAChallengeSubmitRequestTemplate) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *WebRAChallengeSubmitRequestTemplate) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *WebRAChallengeSubmitRequestTemplate) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetExtensions

`func (o *WebRAChallengeSubmitRequestTemplate) GetExtensions() []CertificateExtensionElement`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *WebRAChallengeSubmitRequestTemplate) GetExtensionsOk() (*[]CertificateExtensionElement, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *WebRAChallengeSubmitRequestTemplate) SetExtensions(v []CertificateExtensionElement)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *WebRAChallengeSubmitRequestTemplate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetKeyType

`func (o *WebRAChallengeSubmitRequestTemplate) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *WebRAChallengeSubmitRequestTemplate) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *WebRAChallengeSubmitRequestTemplate) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *WebRAChallengeSubmitRequestTemplate) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetMetadata

`func (o *WebRAChallengeSubmitRequestTemplate) GetMetadata() []CertificateMetadataElement`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAChallengeSubmitRequestTemplate) GetMetadataOk() (*[]CertificateMetadataElement, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAChallengeSubmitRequestTemplate) SetMetadata(v []CertificateMetadataElement)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAChallengeSubmitRequestTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSans

`func (o *WebRAChallengeSubmitRequestTemplate) GetSans() []ListSANElement`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *WebRAChallengeSubmitRequestTemplate) GetSansOk() (*[]ListSANElement, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *WebRAChallengeSubmitRequestTemplate) SetSans(v []ListSANElement)`

SetSans sets Sans field to given value.

### HasSans

`func (o *WebRAChallengeSubmitRequestTemplate) HasSans() bool`

HasSans returns a boolean if a field has been set.

### GetSubject

`func (o *WebRAChallengeSubmitRequestTemplate) GetSubject() []IndexedDNElement`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WebRAChallengeSubmitRequestTemplate) GetSubjectOk() (*[]IndexedDNElement, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WebRAChallengeSubmitRequestTemplate) SetSubject(v []IndexedDNElement)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WebRAChallengeSubmitRequestTemplate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


