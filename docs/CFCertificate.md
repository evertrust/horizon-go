# CFCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | **string** | The certificate&#39;s Distinguished Name | 
**DnElements** | [**[]CFDistinguishedName**](CFDistinguishedName.md) |  | 
**IssuerDn** | **string** | The certificate&#39;s issuer Distinguished Name | 
**Serial** | Pointer to **string** | The certificate&#39;s serial number | [optional] 
**NotBefore** | **int64** | The certificate&#39;s start date in milliseconds since the epoch | 
**NotAfter** | **int64** | The certificate&#39;s expiration date in milliseconds since the epoch | 
**KeyType** | **string** | The certificate&#39;s key type | 
**SigningAlgorithm** | **string** | The certificate&#39;s signing algorithm | 
**Pem** | **string** | The certificate&#39;s PEM-encoded content | 
**SubjectKeyIdentifier** | **string** |  | 
**CertificateThumbprint** | **string** | The certificate&#39;s thumbprint | 
**CertificateSHAOneThumbprint** | **string** | The thumbprint of the certificate using SHAOne algorithm | 
**PublicKeyThumbprint** | **string** | The certificate&#39;s public key thumbprint | 
**KeyUsages** | **[]string** | The certificate key&#39;s usage | 
**IsKeyUsagesCritical** | **bool** | If the key usage of the certificate are critical | 
**ExtendedKeyUsages** | **[]string** | The certificate extended key&#39;s usage | 
**IsExtendedKeyUsagesCritical** | **bool** | If the extended key usage are critical | 
**SelfSigned** | **bool** | Whether the certificate is self-signed | 
**Sans** | Pointer to [**[]SubjectAlternateName**](SubjectAlternateName.md) | The certificate&#39;s SAN | [optional] 
**BasicConstraints** | [**CFCertificateBasicConstraints**](CFCertificateBasicConstraints.md) |  | 
**Extensions** | Pointer to [**[]CertificateExtension**](CertificateExtension.md) | The certificate&#39;s extensions | [optional] 
**Crldps** | Pointer to **[]string** | The certificate&#39;s CRLDP if any | [optional] 
**Aias** | Pointer to [**CFCertificateAias**](CFCertificateAias.md) |  | [optional] 
**Policies** | Pointer to [**[]CFCertificatePoliciesInner**](CFCertificatePoliciesInner.md) |  | [optional] 
**AuthorityKeyIdentifier** | Pointer to **string** | The certificate AKI | [optional] 
**UnsupportedExtensions** | Pointer to [**[]CFCertificateUnsupportedExtensionsInner**](CFCertificateUnsupportedExtensionsInner.md) |  | [optional] 

## Methods

### NewCFCertificate

`func NewCFCertificate(dn string, dnElements []CFDistinguishedName, issuerDn string, notBefore int64, notAfter int64, keyType string, signingAlgorithm string, pem string, subjectKeyIdentifier string, certificateThumbprint string, certificateSHAOneThumbprint string, publicKeyThumbprint string, keyUsages []string, isKeyUsagesCritical bool, extendedKeyUsages []string, isExtendedKeyUsagesCritical bool, selfSigned bool, basicConstraints CFCertificateBasicConstraints, ) *CFCertificate`

NewCFCertificate instantiates a new CFCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCertificateWithDefaults

`func NewCFCertificateWithDefaults() *CFCertificate`

NewCFCertificateWithDefaults instantiates a new CFCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDn

`func (o *CFCertificate) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *CFCertificate) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *CFCertificate) SetDn(v string)`

SetDn sets Dn field to given value.


### GetDnElements

`func (o *CFCertificate) GetDnElements() []CFDistinguishedName`

GetDnElements returns the DnElements field if non-nil, zero value otherwise.

### GetDnElementsOk

`func (o *CFCertificate) GetDnElementsOk() (*[]CFDistinguishedName, bool)`

GetDnElementsOk returns a tuple with the DnElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnElements

`func (o *CFCertificate) SetDnElements(v []CFDistinguishedName)`

SetDnElements sets DnElements field to given value.


### GetIssuerDn

`func (o *CFCertificate) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *CFCertificate) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *CFCertificate) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.


### GetSerial

`func (o *CFCertificate) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CFCertificate) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CFCertificate) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CFCertificate) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNotBefore

`func (o *CFCertificate) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CFCertificate) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CFCertificate) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.


### GetNotAfter

`func (o *CFCertificate) GetNotAfter() int64`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CFCertificate) GetNotAfterOk() (*int64, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CFCertificate) SetNotAfter(v int64)`

SetNotAfter sets NotAfter field to given value.


### GetKeyType

`func (o *CFCertificate) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CFCertificate) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CFCertificate) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetSigningAlgorithm

`func (o *CFCertificate) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CFCertificate) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CFCertificate) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.


### GetPem

`func (o *CFCertificate) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *CFCertificate) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *CFCertificate) SetPem(v string)`

SetPem sets Pem field to given value.


### GetSubjectKeyIdentifier

`func (o *CFCertificate) GetSubjectKeyIdentifier() string`

GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field if non-nil, zero value otherwise.

### GetSubjectKeyIdentifierOk

`func (o *CFCertificate) GetSubjectKeyIdentifierOk() (*string, bool)`

GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyIdentifier

`func (o *CFCertificate) SetSubjectKeyIdentifier(v string)`

SetSubjectKeyIdentifier sets SubjectKeyIdentifier field to given value.


### GetCertificateThumbprint

`func (o *CFCertificate) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *CFCertificate) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *CFCertificate) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.


### GetCertificateSHAOneThumbprint

`func (o *CFCertificate) GetCertificateSHAOneThumbprint() string`

GetCertificateSHAOneThumbprint returns the CertificateSHAOneThumbprint field if non-nil, zero value otherwise.

### GetCertificateSHAOneThumbprintOk

`func (o *CFCertificate) GetCertificateSHAOneThumbprintOk() (*string, bool)`

GetCertificateSHAOneThumbprintOk returns a tuple with the CertificateSHAOneThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSHAOneThumbprint

`func (o *CFCertificate) SetCertificateSHAOneThumbprint(v string)`

SetCertificateSHAOneThumbprint sets CertificateSHAOneThumbprint field to given value.


### GetPublicKeyThumbprint

`func (o *CFCertificate) GetPublicKeyThumbprint() string`

GetPublicKeyThumbprint returns the PublicKeyThumbprint field if non-nil, zero value otherwise.

### GetPublicKeyThumbprintOk

`func (o *CFCertificate) GetPublicKeyThumbprintOk() (*string, bool)`

GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyThumbprint

`func (o *CFCertificate) SetPublicKeyThumbprint(v string)`

SetPublicKeyThumbprint sets PublicKeyThumbprint field to given value.


### GetKeyUsages

`func (o *CFCertificate) GetKeyUsages() []string`

GetKeyUsages returns the KeyUsages field if non-nil, zero value otherwise.

### GetKeyUsagesOk

`func (o *CFCertificate) GetKeyUsagesOk() (*[]string, bool)`

GetKeyUsagesOk returns a tuple with the KeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsages

`func (o *CFCertificate) SetKeyUsages(v []string)`

SetKeyUsages sets KeyUsages field to given value.


### GetIsKeyUsagesCritical

`func (o *CFCertificate) GetIsKeyUsagesCritical() bool`

GetIsKeyUsagesCritical returns the IsKeyUsagesCritical field if non-nil, zero value otherwise.

### GetIsKeyUsagesCriticalOk

`func (o *CFCertificate) GetIsKeyUsagesCriticalOk() (*bool, bool)`

GetIsKeyUsagesCriticalOk returns a tuple with the IsKeyUsagesCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeyUsagesCritical

`func (o *CFCertificate) SetIsKeyUsagesCritical(v bool)`

SetIsKeyUsagesCritical sets IsKeyUsagesCritical field to given value.


### GetExtendedKeyUsages

`func (o *CFCertificate) GetExtendedKeyUsages() []string`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *CFCertificate) GetExtendedKeyUsagesOk() (*[]string, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *CFCertificate) SetExtendedKeyUsages(v []string)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.


### GetIsExtendedKeyUsagesCritical

`func (o *CFCertificate) GetIsExtendedKeyUsagesCritical() bool`

GetIsExtendedKeyUsagesCritical returns the IsExtendedKeyUsagesCritical field if non-nil, zero value otherwise.

### GetIsExtendedKeyUsagesCriticalOk

`func (o *CFCertificate) GetIsExtendedKeyUsagesCriticalOk() (*bool, bool)`

GetIsExtendedKeyUsagesCriticalOk returns a tuple with the IsExtendedKeyUsagesCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExtendedKeyUsagesCritical

`func (o *CFCertificate) SetIsExtendedKeyUsagesCritical(v bool)`

SetIsExtendedKeyUsagesCritical sets IsExtendedKeyUsagesCritical field to given value.


### GetSelfSigned

`func (o *CFCertificate) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *CFCertificate) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *CFCertificate) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.


### GetSans

`func (o *CFCertificate) GetSans() []SubjectAlternateName`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *CFCertificate) GetSansOk() (*[]SubjectAlternateName, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *CFCertificate) SetSans(v []SubjectAlternateName)`

SetSans sets Sans field to given value.

### HasSans

`func (o *CFCertificate) HasSans() bool`

HasSans returns a boolean if a field has been set.

### GetBasicConstraints

`func (o *CFCertificate) GetBasicConstraints() CFCertificateBasicConstraints`

GetBasicConstraints returns the BasicConstraints field if non-nil, zero value otherwise.

### GetBasicConstraintsOk

`func (o *CFCertificate) GetBasicConstraintsOk() (*CFCertificateBasicConstraints, bool)`

GetBasicConstraintsOk returns a tuple with the BasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraints

`func (o *CFCertificate) SetBasicConstraints(v CFCertificateBasicConstraints)`

SetBasicConstraints sets BasicConstraints field to given value.


### GetExtensions

`func (o *CFCertificate) GetExtensions() []CertificateExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CFCertificate) GetExtensionsOk() (*[]CertificateExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CFCertificate) SetExtensions(v []CertificateExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CFCertificate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *CFCertificate) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *CFCertificate) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetCrldps

`func (o *CFCertificate) GetCrldps() []string`

GetCrldps returns the Crldps field if non-nil, zero value otherwise.

### GetCrldpsOk

`func (o *CFCertificate) GetCrldpsOk() (*[]string, bool)`

GetCrldpsOk returns a tuple with the Crldps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrldps

`func (o *CFCertificate) SetCrldps(v []string)`

SetCrldps sets Crldps field to given value.

### HasCrldps

`func (o *CFCertificate) HasCrldps() bool`

HasCrldps returns a boolean if a field has been set.

### GetAias

`func (o *CFCertificate) GetAias() CFCertificateAias`

GetAias returns the Aias field if non-nil, zero value otherwise.

### GetAiasOk

`func (o *CFCertificate) GetAiasOk() (*CFCertificateAias, bool)`

GetAiasOk returns a tuple with the Aias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAias

`func (o *CFCertificate) SetAias(v CFCertificateAias)`

SetAias sets Aias field to given value.

### HasAias

`func (o *CFCertificate) HasAias() bool`

HasAias returns a boolean if a field has been set.

### GetPolicies

`func (o *CFCertificate) GetPolicies() []CFCertificatePoliciesInner`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CFCertificate) GetPoliciesOk() (*[]CFCertificatePoliciesInner, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CFCertificate) SetPolicies(v []CFCertificatePoliciesInner)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CFCertificate) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetAuthorityKeyIdentifier

`func (o *CFCertificate) GetAuthorityKeyIdentifier() string`

GetAuthorityKeyIdentifier returns the AuthorityKeyIdentifier field if non-nil, zero value otherwise.

### GetAuthorityKeyIdentifierOk

`func (o *CFCertificate) GetAuthorityKeyIdentifierOk() (*string, bool)`

GetAuthorityKeyIdentifierOk returns a tuple with the AuthorityKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityKeyIdentifier

`func (o *CFCertificate) SetAuthorityKeyIdentifier(v string)`

SetAuthorityKeyIdentifier sets AuthorityKeyIdentifier field to given value.

### HasAuthorityKeyIdentifier

`func (o *CFCertificate) HasAuthorityKeyIdentifier() bool`

HasAuthorityKeyIdentifier returns a boolean if a field has been set.

### GetUnsupportedExtensions

`func (o *CFCertificate) GetUnsupportedExtensions() []CFCertificateUnsupportedExtensionsInner`

GetUnsupportedExtensions returns the UnsupportedExtensions field if non-nil, zero value otherwise.

### GetUnsupportedExtensionsOk

`func (o *CFCertificate) GetUnsupportedExtensionsOk() (*[]CFCertificateUnsupportedExtensionsInner, bool)`

GetUnsupportedExtensionsOk returns a tuple with the UnsupportedExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedExtensions

`func (o *CFCertificate) SetUnsupportedExtensions(v []CFCertificateUnsupportedExtensionsInner)`

SetUnsupportedExtensions sets UnsupportedExtensions field to given value.

### HasUnsupportedExtensions

`func (o *CFCertificate) HasUnsupportedExtensions() bool`

HasUnsupportedExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


