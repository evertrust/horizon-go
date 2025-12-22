# CFCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aias** | Pointer to [**CFCertificateAias**](CFCertificateAias.md) |  | [optional] 
**AuthorityKeyIdentifier** | Pointer to **string** | The certificate AKI | [optional] 
**BasicConstraints** | [**CFCertificateBasicConstraints**](CFCertificateBasicConstraints.md) |  | 
**CertificateSHAOneThumbprint** | **string** | The thumbprint of the certificate using SHAOne algorithm | 
**CertificateThumbprint** | **string** | The certificate&#39;s thumbprint | 
**Crldps** | Pointer to **[]string** | The certificate&#39;s CRLDP if any | [optional] 
**Dn** | **string** | The certificate&#39;s Distinguished Name | 
**DnElements** | [**[]CFDistinguishedName**](CFDistinguishedName.md) |  | 
**ExtendedKeyUsages** | **[]string** | The certificate extended key&#39;s usage | 
**Extensions** | Pointer to [**[]CertificateExtension**](CertificateExtension.md) | The certificate&#39;s extensions | [optional] 
**IsExtendedKeyUsagesCritical** | **bool** | If the extended key usage are critical | 
**IsKeyUsagesCritical** | **bool** | If the key usage of the certificate are critical | 
**IssuerDn** | **string** | The certificate&#39;s issuer Distinguished Name | 
**KeyType** | **string** | The certificate&#39;s key type | 
**KeyUsages** | **[]string** | The certificate key&#39;s usage | 
**NotAfter** | **int64** | The certificate&#39;s expiration date in milliseconds since the epoch | 
**NotBefore** | **int64** | The certificate&#39;s start date in milliseconds since the epoch | 
**Pem** | **string** | The certificate&#39;s PEM-encoded content | 
**Policies** | Pointer to [**[]CFCertificatePoliciesInner**](CFCertificatePoliciesInner.md) |  | [optional] 
**PublicKeyThumbprint** | **string** | The certificate&#39;s public key thumbprint | 
**Sans** | Pointer to [**[]SubjectAlternateName**](SubjectAlternateName.md) | The certificate&#39;s SAN | [optional] 
**SelfSigned** | **bool** | Whether the certificate is self-signed | 
**Serial** | Pointer to **string** | The certificate&#39;s serial number | [optional] 
**SigningAlgorithm** | **string** | The certificate&#39;s signing algorithm | 
**SubjectKeyIdentifier** | **string** |  | 
**UnsupportedExtensions** | Pointer to [**[]CFCertificateUnsupportedExtensionsInner**](CFCertificateUnsupportedExtensionsInner.md) |  | [optional] 

## Methods

### NewCFCertificateResponse

`func NewCFCertificateResponse(basicConstraints CFCertificateBasicConstraints, certificateSHAOneThumbprint string, certificateThumbprint string, dn string, dnElements []CFDistinguishedName, extendedKeyUsages []string, isExtendedKeyUsagesCritical bool, isKeyUsagesCritical bool, issuerDn string, keyType string, keyUsages []string, notAfter int64, notBefore int64, pem string, publicKeyThumbprint string, selfSigned bool, signingAlgorithm string, subjectKeyIdentifier string, ) *CFCertificateResponse`

NewCFCertificateResponse instantiates a new CFCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCertificateResponseWithDefaults

`func NewCFCertificateResponseWithDefaults() *CFCertificateResponse`

NewCFCertificateResponseWithDefaults instantiates a new CFCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAias

`func (o *CFCertificateResponse) GetAias() CFCertificateAias`

GetAias returns the Aias field if non-nil, zero value otherwise.

### GetAiasOk

`func (o *CFCertificateResponse) GetAiasOk() (*CFCertificateAias, bool)`

GetAiasOk returns a tuple with the Aias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAias

`func (o *CFCertificateResponse) SetAias(v CFCertificateAias)`

SetAias sets Aias field to given value.

### HasAias

`func (o *CFCertificateResponse) HasAias() bool`

HasAias returns a boolean if a field has been set.

### GetAuthorityKeyIdentifier

`func (o *CFCertificateResponse) GetAuthorityKeyIdentifier() string`

GetAuthorityKeyIdentifier returns the AuthorityKeyIdentifier field if non-nil, zero value otherwise.

### GetAuthorityKeyIdentifierOk

`func (o *CFCertificateResponse) GetAuthorityKeyIdentifierOk() (*string, bool)`

GetAuthorityKeyIdentifierOk returns a tuple with the AuthorityKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityKeyIdentifier

`func (o *CFCertificateResponse) SetAuthorityKeyIdentifier(v string)`

SetAuthorityKeyIdentifier sets AuthorityKeyIdentifier field to given value.

### HasAuthorityKeyIdentifier

`func (o *CFCertificateResponse) HasAuthorityKeyIdentifier() bool`

HasAuthorityKeyIdentifier returns a boolean if a field has been set.

### GetBasicConstraints

`func (o *CFCertificateResponse) GetBasicConstraints() CFCertificateBasicConstraints`

GetBasicConstraints returns the BasicConstraints field if non-nil, zero value otherwise.

### GetBasicConstraintsOk

`func (o *CFCertificateResponse) GetBasicConstraintsOk() (*CFCertificateBasicConstraints, bool)`

GetBasicConstraintsOk returns a tuple with the BasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraints

`func (o *CFCertificateResponse) SetBasicConstraints(v CFCertificateBasicConstraints)`

SetBasicConstraints sets BasicConstraints field to given value.


### GetCertificateSHAOneThumbprint

`func (o *CFCertificateResponse) GetCertificateSHAOneThumbprint() string`

GetCertificateSHAOneThumbprint returns the CertificateSHAOneThumbprint field if non-nil, zero value otherwise.

### GetCertificateSHAOneThumbprintOk

`func (o *CFCertificateResponse) GetCertificateSHAOneThumbprintOk() (*string, bool)`

GetCertificateSHAOneThumbprintOk returns a tuple with the CertificateSHAOneThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSHAOneThumbprint

`func (o *CFCertificateResponse) SetCertificateSHAOneThumbprint(v string)`

SetCertificateSHAOneThumbprint sets CertificateSHAOneThumbprint field to given value.


### GetCertificateThumbprint

`func (o *CFCertificateResponse) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *CFCertificateResponse) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *CFCertificateResponse) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.


### GetCrldps

`func (o *CFCertificateResponse) GetCrldps() []string`

GetCrldps returns the Crldps field if non-nil, zero value otherwise.

### GetCrldpsOk

`func (o *CFCertificateResponse) GetCrldpsOk() (*[]string, bool)`

GetCrldpsOk returns a tuple with the Crldps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrldps

`func (o *CFCertificateResponse) SetCrldps(v []string)`

SetCrldps sets Crldps field to given value.

### HasCrldps

`func (o *CFCertificateResponse) HasCrldps() bool`

HasCrldps returns a boolean if a field has been set.

### GetDn

`func (o *CFCertificateResponse) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *CFCertificateResponse) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *CFCertificateResponse) SetDn(v string)`

SetDn sets Dn field to given value.


### GetDnElements

`func (o *CFCertificateResponse) GetDnElements() []CFDistinguishedName`

GetDnElements returns the DnElements field if non-nil, zero value otherwise.

### GetDnElementsOk

`func (o *CFCertificateResponse) GetDnElementsOk() (*[]CFDistinguishedName, bool)`

GetDnElementsOk returns a tuple with the DnElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnElements

`func (o *CFCertificateResponse) SetDnElements(v []CFDistinguishedName)`

SetDnElements sets DnElements field to given value.


### GetExtendedKeyUsages

`func (o *CFCertificateResponse) GetExtendedKeyUsages() []string`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *CFCertificateResponse) GetExtendedKeyUsagesOk() (*[]string, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *CFCertificateResponse) SetExtendedKeyUsages(v []string)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.


### GetExtensions

`func (o *CFCertificateResponse) GetExtensions() []CertificateExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CFCertificateResponse) GetExtensionsOk() (*[]CertificateExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CFCertificateResponse) SetExtensions(v []CertificateExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CFCertificateResponse) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *CFCertificateResponse) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *CFCertificateResponse) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetIsExtendedKeyUsagesCritical

`func (o *CFCertificateResponse) GetIsExtendedKeyUsagesCritical() bool`

GetIsExtendedKeyUsagesCritical returns the IsExtendedKeyUsagesCritical field if non-nil, zero value otherwise.

### GetIsExtendedKeyUsagesCriticalOk

`func (o *CFCertificateResponse) GetIsExtendedKeyUsagesCriticalOk() (*bool, bool)`

GetIsExtendedKeyUsagesCriticalOk returns a tuple with the IsExtendedKeyUsagesCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExtendedKeyUsagesCritical

`func (o *CFCertificateResponse) SetIsExtendedKeyUsagesCritical(v bool)`

SetIsExtendedKeyUsagesCritical sets IsExtendedKeyUsagesCritical field to given value.


### GetIsKeyUsagesCritical

`func (o *CFCertificateResponse) GetIsKeyUsagesCritical() bool`

GetIsKeyUsagesCritical returns the IsKeyUsagesCritical field if non-nil, zero value otherwise.

### GetIsKeyUsagesCriticalOk

`func (o *CFCertificateResponse) GetIsKeyUsagesCriticalOk() (*bool, bool)`

GetIsKeyUsagesCriticalOk returns a tuple with the IsKeyUsagesCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeyUsagesCritical

`func (o *CFCertificateResponse) SetIsKeyUsagesCritical(v bool)`

SetIsKeyUsagesCritical sets IsKeyUsagesCritical field to given value.


### GetIssuerDn

`func (o *CFCertificateResponse) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *CFCertificateResponse) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *CFCertificateResponse) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.


### GetKeyType

`func (o *CFCertificateResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CFCertificateResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CFCertificateResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyUsages

`func (o *CFCertificateResponse) GetKeyUsages() []string`

GetKeyUsages returns the KeyUsages field if non-nil, zero value otherwise.

### GetKeyUsagesOk

`func (o *CFCertificateResponse) GetKeyUsagesOk() (*[]string, bool)`

GetKeyUsagesOk returns a tuple with the KeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsages

`func (o *CFCertificateResponse) SetKeyUsages(v []string)`

SetKeyUsages sets KeyUsages field to given value.


### GetNotAfter

`func (o *CFCertificateResponse) GetNotAfter() int64`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CFCertificateResponse) GetNotAfterOk() (*int64, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CFCertificateResponse) SetNotAfter(v int64)`

SetNotAfter sets NotAfter field to given value.


### GetNotBefore

`func (o *CFCertificateResponse) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CFCertificateResponse) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CFCertificateResponse) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.


### GetPem

`func (o *CFCertificateResponse) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *CFCertificateResponse) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *CFCertificateResponse) SetPem(v string)`

SetPem sets Pem field to given value.


### GetPolicies

`func (o *CFCertificateResponse) GetPolicies() []CFCertificatePoliciesInner`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CFCertificateResponse) GetPoliciesOk() (*[]CFCertificatePoliciesInner, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CFCertificateResponse) SetPolicies(v []CFCertificatePoliciesInner)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CFCertificateResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetPublicKeyThumbprint

`func (o *CFCertificateResponse) GetPublicKeyThumbprint() string`

GetPublicKeyThumbprint returns the PublicKeyThumbprint field if non-nil, zero value otherwise.

### GetPublicKeyThumbprintOk

`func (o *CFCertificateResponse) GetPublicKeyThumbprintOk() (*string, bool)`

GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyThumbprint

`func (o *CFCertificateResponse) SetPublicKeyThumbprint(v string)`

SetPublicKeyThumbprint sets PublicKeyThumbprint field to given value.


### GetSans

`func (o *CFCertificateResponse) GetSans() []SubjectAlternateName`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *CFCertificateResponse) GetSansOk() (*[]SubjectAlternateName, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *CFCertificateResponse) SetSans(v []SubjectAlternateName)`

SetSans sets Sans field to given value.

### HasSans

`func (o *CFCertificateResponse) HasSans() bool`

HasSans returns a boolean if a field has been set.

### GetSelfSigned

`func (o *CFCertificateResponse) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *CFCertificateResponse) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *CFCertificateResponse) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.


### GetSerial

`func (o *CFCertificateResponse) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CFCertificateResponse) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CFCertificateResponse) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CFCertificateResponse) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *CFCertificateResponse) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CFCertificateResponse) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CFCertificateResponse) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.


### GetSubjectKeyIdentifier

`func (o *CFCertificateResponse) GetSubjectKeyIdentifier() string`

GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field if non-nil, zero value otherwise.

### GetSubjectKeyIdentifierOk

`func (o *CFCertificateResponse) GetSubjectKeyIdentifierOk() (*string, bool)`

GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyIdentifier

`func (o *CFCertificateResponse) SetSubjectKeyIdentifier(v string)`

SetSubjectKeyIdentifier sets SubjectKeyIdentifier field to given value.


### GetUnsupportedExtensions

`func (o *CFCertificateResponse) GetUnsupportedExtensions() []CFCertificateUnsupportedExtensionsInner`

GetUnsupportedExtensions returns the UnsupportedExtensions field if non-nil, zero value otherwise.

### GetUnsupportedExtensionsOk

`func (o *CFCertificateResponse) GetUnsupportedExtensionsOk() (*[]CFCertificateUnsupportedExtensionsInner, bool)`

GetUnsupportedExtensionsOk returns a tuple with the UnsupportedExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedExtensions

`func (o *CFCertificateResponse) SetUnsupportedExtensions(v []CFCertificateUnsupportedExtensionsInner)`

SetUnsupportedExtensions sets UnsupportedExtensions field to given value.

### HasUnsupportedExtensions

`func (o *CFCertificateResponse) HasUnsupportedExtensions() bool`

HasUnsupportedExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


