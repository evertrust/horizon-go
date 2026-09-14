# X509CertificateElements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to [**[]X509ExtensionElementResponse**](X509ExtensionElementResponse.md) | The requested certificate extensions (Microsoft template extensions only) | [optional] 
**Sans** | Pointer to [**[]X509SanElementResponse**](X509SanElementResponse.md) | The requested Subject Alternative Names, grouped by type | [optional] 
**Subject** | Pointer to [**[]X509DnElementResponse**](X509DnElementResponse.md) | The requested subject Distinguished Name elements, in order | [optional] 

## Methods

### NewX509CertificateElements

`func NewX509CertificateElements() *X509CertificateElements`

NewX509CertificateElements instantiates a new X509CertificateElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateElementsWithDefaults

`func NewX509CertificateElementsWithDefaults() *X509CertificateElements`

NewX509CertificateElementsWithDefaults instantiates a new X509CertificateElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *X509CertificateElements) GetExtensions() []X509ExtensionElementResponse`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *X509CertificateElements) GetExtensionsOk() (*[]X509ExtensionElementResponse, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *X509CertificateElements) SetExtensions(v []X509ExtensionElementResponse)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *X509CertificateElements) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetSans

`func (o *X509CertificateElements) GetSans() []X509SanElementResponse`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *X509CertificateElements) GetSansOk() (*[]X509SanElementResponse, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *X509CertificateElements) SetSans(v []X509SanElementResponse)`

SetSans sets Sans field to given value.

### HasSans

`func (o *X509CertificateElements) HasSans() bool`

HasSans returns a boolean if a field has been set.

### GetSubject

`func (o *X509CertificateElements) GetSubject() []X509DnElementResponse`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *X509CertificateElements) GetSubjectOk() (*[]X509DnElementResponse, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *X509CertificateElements) SetSubject(v []X509DnElementResponse)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *X509CertificateElements) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


