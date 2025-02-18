# CertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**[]CertificateMetadata**](CertificateMetadata.md) | The certificate&#39;s technical metadata used internally | 
**NotAfter** | **int64** | The certificate&#39;s expiration date in milliseconds since the epoch | 
**Thumbprint** | **string** | The certificate&#39;s thumbprint | 
**RevocationDate** | Pointer to **NullableInt64** | The certificate&#39;s revocation date in milliseconds since the epoch. This field is only present if the certificate is revoked | [optional] 
**Certificate** | **string** | The certificate&#39;s PEM-encoded content | 
**Dn** | **string** | The certificate&#39;s Distinguished Name | 
**Grades** | Pointer to [**[]GradingPolicyResult**](GradingPolicyResult.md) | The certificate&#39;s grades for the enabled grading policies | [optional] 
**Revoked** | **bool** | Whether the certificate is revoked | 
**Escrowed** | **bool** | Whether the certificate is escrowed | 
**Issuer** | **string** | The certificate&#39;s issuer Distinguished Name | 
**NotBefore** | **int64** | The certificate&#39;s start date in milliseconds since the epoch | 
**CrlSynchronized** | Pointer to **NullableBool** | Whether the certificate&#39;s revocation status is synchronized with a CRL | [optional] 
**SelfSigned** | **bool** | Whether the certificate is self-signed | 
**DiscoveredTrusted** | Pointer to **NullableBool** | If the certificate was discovered and is found to be issued by an existing trusted CA, this field will be set to true. If the certificate was discovered and is not found to be issued by an existing trusted CA, this field will be set to false. If the certificate was not discovered, this field will be null | [optional] 
**KeyType** | **string** | The certificate&#39;s key type | 
**ThirdPartyData** | Pointer to [**[]ThirdPartyItem**](ThirdPartyItem.md) | The certificate&#39;s information about synchronization with Horizon supported third parties | [optional] 
**Owner** | Pointer to **NullableString** | The certificate&#39;s owner. This is a reference to a local identity identifier | [optional] 
**PublicKeyThumbprint** | **string** | The certificate&#39;s public key thumbprint | 
**ContactEmail** | Pointer to **NullableString** | The certificate&#39;s contact email. It will be used to send notifications about the certificate&#39;s expiration and revocation | [optional] 
**Module** | **string** | The certificate&#39;s module | 
**Profile** | Pointer to **NullableString** | The certificate&#39;s profile | [optional] 
**Team** | Pointer to **NullableString** | The certificate&#39;s team. This is a reference to a team identifier. It will be used to determine the certificate&#39;s permissions and send notifications | [optional] 
**HolderId** | **string** | The certificate&#39;s holder ID. This is a computed field that is used to count how many similar certificates are in use simultaneously by the same holder | 
**Labels** | Pointer to [**[]LabelData**](LabelData.md) | The certificate&#39;s labels | [optional] 
**DiscoveryInfo** | Pointer to [**[]DiscoveryInfo**](DiscoveryInfo.md) | A list of metadata containing information on how and when the certificate was discovered | [optional] 
**SubjectAlternateNames** | [**[]SubjectAlternateName**](SubjectAlternateName.md) | The certificate&#39;s Subject Alternate Names | 
**TriggerResults** | Pointer to [**[]TriggerResult**](TriggerResult.md) | The result of the execution of triggers on this certificate | [optional] 
**Extensions** | Pointer to [**[]CertificateExtension**](CertificateExtension.md) | The certificate&#39;s extensions | [optional] 
**Serial** | **string** | The certificate&#39;s serial number | 
**SigningAlgorithm** | **string** | The certificate&#39;s signing algorithm | 
**DiscoveryData** | Pointer to [**[]HostDiscoveryData**](HostDiscoveryData.md) | A list of metadata containing information on where the certificate was discovered | [optional] 
**Id** | **string** | Object internal ID | 
**RevocationReason** | Pointer to **NullableString** | The certificate&#39;s revocation reason | [optional] 

## Methods

### NewCertificateResponse

`func NewCertificateResponse(metadata []CertificateMetadata, notAfter int64, thumbprint string, certificate string, dn string, revoked bool, escrowed bool, issuer string, notBefore int64, selfSigned bool, keyType string, publicKeyThumbprint string, module string, holderId string, subjectAlternateNames []SubjectAlternateName, serial string, signingAlgorithm string, id string, ) *CertificateResponse`

NewCertificateResponse instantiates a new CertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateResponseWithDefaults

`func NewCertificateResponseWithDefaults() *CertificateResponse`

NewCertificateResponseWithDefaults instantiates a new CertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CertificateResponse) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificateResponse) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificateResponse) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.


### GetNotAfter

`func (o *CertificateResponse) GetNotAfter() int64`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateResponse) GetNotAfterOk() (*int64, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateResponse) SetNotAfter(v int64)`

SetNotAfter sets NotAfter field to given value.


### GetThumbprint

`func (o *CertificateResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CertificateResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CertificateResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.


### GetRevocationDate

`func (o *CertificateResponse) GetRevocationDate() int64`

GetRevocationDate returns the RevocationDate field if non-nil, zero value otherwise.

### GetRevocationDateOk

`func (o *CertificateResponse) GetRevocationDateOk() (*int64, bool)`

GetRevocationDateOk returns a tuple with the RevocationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationDate

`func (o *CertificateResponse) SetRevocationDate(v int64)`

SetRevocationDate sets RevocationDate field to given value.

### HasRevocationDate

`func (o *CertificateResponse) HasRevocationDate() bool`

HasRevocationDate returns a boolean if a field has been set.

### SetRevocationDateNil

`func (o *CertificateResponse) SetRevocationDateNil(b bool)`

 SetRevocationDateNil sets the value for RevocationDate to be an explicit nil

### UnsetRevocationDate
`func (o *CertificateResponse) UnsetRevocationDate()`

UnsetRevocationDate ensures that no value is present for RevocationDate, not even an explicit nil
### GetCertificate

`func (o *CertificateResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetDn

`func (o *CertificateResponse) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *CertificateResponse) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *CertificateResponse) SetDn(v string)`

SetDn sets Dn field to given value.


### GetGrades

`func (o *CertificateResponse) GetGrades() []GradingPolicyResult`

GetGrades returns the Grades field if non-nil, zero value otherwise.

### GetGradesOk

`func (o *CertificateResponse) GetGradesOk() (*[]GradingPolicyResult, bool)`

GetGradesOk returns a tuple with the Grades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrades

`func (o *CertificateResponse) SetGrades(v []GradingPolicyResult)`

SetGrades sets Grades field to given value.

### HasGrades

`func (o *CertificateResponse) HasGrades() bool`

HasGrades returns a boolean if a field has been set.

### SetGradesNil

`func (o *CertificateResponse) SetGradesNil(b bool)`

 SetGradesNil sets the value for Grades to be an explicit nil

### UnsetGrades
`func (o *CertificateResponse) UnsetGrades()`

UnsetGrades ensures that no value is present for Grades, not even an explicit nil
### GetRevoked

`func (o *CertificateResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *CertificateResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *CertificateResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.


### GetEscrowed

`func (o *CertificateResponse) GetEscrowed() bool`

GetEscrowed returns the Escrowed field if non-nil, zero value otherwise.

### GetEscrowedOk

`func (o *CertificateResponse) GetEscrowedOk() (*bool, bool)`

GetEscrowedOk returns a tuple with the Escrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscrowed

`func (o *CertificateResponse) SetEscrowed(v bool)`

SetEscrowed sets Escrowed field to given value.


### GetIssuer

`func (o *CertificateResponse) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateResponse) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateResponse) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetNotBefore

`func (o *CertificateResponse) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateResponse) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateResponse) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.


### GetCrlSynchronized

`func (o *CertificateResponse) GetCrlSynchronized() bool`

GetCrlSynchronized returns the CrlSynchronized field if non-nil, zero value otherwise.

### GetCrlSynchronizedOk

`func (o *CertificateResponse) GetCrlSynchronizedOk() (*bool, bool)`

GetCrlSynchronizedOk returns a tuple with the CrlSynchronized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlSynchronized

`func (o *CertificateResponse) SetCrlSynchronized(v bool)`

SetCrlSynchronized sets CrlSynchronized field to given value.

### HasCrlSynchronized

`func (o *CertificateResponse) HasCrlSynchronized() bool`

HasCrlSynchronized returns a boolean if a field has been set.

### SetCrlSynchronizedNil

`func (o *CertificateResponse) SetCrlSynchronizedNil(b bool)`

 SetCrlSynchronizedNil sets the value for CrlSynchronized to be an explicit nil

### UnsetCrlSynchronized
`func (o *CertificateResponse) UnsetCrlSynchronized()`

UnsetCrlSynchronized ensures that no value is present for CrlSynchronized, not even an explicit nil
### GetSelfSigned

`func (o *CertificateResponse) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *CertificateResponse) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *CertificateResponse) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.


### GetDiscoveredTrusted

`func (o *CertificateResponse) GetDiscoveredTrusted() bool`

GetDiscoveredTrusted returns the DiscoveredTrusted field if non-nil, zero value otherwise.

### GetDiscoveredTrustedOk

`func (o *CertificateResponse) GetDiscoveredTrustedOk() (*bool, bool)`

GetDiscoveredTrustedOk returns a tuple with the DiscoveredTrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredTrusted

`func (o *CertificateResponse) SetDiscoveredTrusted(v bool)`

SetDiscoveredTrusted sets DiscoveredTrusted field to given value.

### HasDiscoveredTrusted

`func (o *CertificateResponse) HasDiscoveredTrusted() bool`

HasDiscoveredTrusted returns a boolean if a field has been set.

### SetDiscoveredTrustedNil

`func (o *CertificateResponse) SetDiscoveredTrustedNil(b bool)`

 SetDiscoveredTrustedNil sets the value for DiscoveredTrusted to be an explicit nil

### UnsetDiscoveredTrusted
`func (o *CertificateResponse) UnsetDiscoveredTrusted()`

UnsetDiscoveredTrusted ensures that no value is present for DiscoveredTrusted, not even an explicit nil
### GetKeyType

`func (o *CertificateResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CertificateResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CertificateResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetThirdPartyData

`func (o *CertificateResponse) GetThirdPartyData() []ThirdPartyItem`

GetThirdPartyData returns the ThirdPartyData field if non-nil, zero value otherwise.

### GetThirdPartyDataOk

`func (o *CertificateResponse) GetThirdPartyDataOk() (*[]ThirdPartyItem, bool)`

GetThirdPartyDataOk returns a tuple with the ThirdPartyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyData

`func (o *CertificateResponse) SetThirdPartyData(v []ThirdPartyItem)`

SetThirdPartyData sets ThirdPartyData field to given value.

### HasThirdPartyData

`func (o *CertificateResponse) HasThirdPartyData() bool`

HasThirdPartyData returns a boolean if a field has been set.

### SetThirdPartyDataNil

`func (o *CertificateResponse) SetThirdPartyDataNil(b bool)`

 SetThirdPartyDataNil sets the value for ThirdPartyData to be an explicit nil

### UnsetThirdPartyData
`func (o *CertificateResponse) UnsetThirdPartyData()`

UnsetThirdPartyData ensures that no value is present for ThirdPartyData, not even an explicit nil
### GetOwner

`func (o *CertificateResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CertificateResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CertificateResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CertificateResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *CertificateResponse) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *CertificateResponse) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetPublicKeyThumbprint

`func (o *CertificateResponse) GetPublicKeyThumbprint() string`

GetPublicKeyThumbprint returns the PublicKeyThumbprint field if non-nil, zero value otherwise.

### GetPublicKeyThumbprintOk

`func (o *CertificateResponse) GetPublicKeyThumbprintOk() (*string, bool)`

GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyThumbprint

`func (o *CertificateResponse) SetPublicKeyThumbprint(v string)`

SetPublicKeyThumbprint sets PublicKeyThumbprint field to given value.


### GetContactEmail

`func (o *CertificateResponse) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *CertificateResponse) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *CertificateResponse) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *CertificateResponse) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *CertificateResponse) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *CertificateResponse) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetModule

`func (o *CertificateResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CertificateResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CertificateResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *CertificateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CertificateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CertificateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CertificateResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *CertificateResponse) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *CertificateResponse) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetTeam

`func (o *CertificateResponse) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *CertificateResponse) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *CertificateResponse) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *CertificateResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *CertificateResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *CertificateResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetHolderId

`func (o *CertificateResponse) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *CertificateResponse) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *CertificateResponse) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.


### GetLabels

`func (o *CertificateResponse) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CertificateResponse) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CertificateResponse) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CertificateResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *CertificateResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *CertificateResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDiscoveryInfo

`func (o *CertificateResponse) GetDiscoveryInfo() []DiscoveryInfo`

GetDiscoveryInfo returns the DiscoveryInfo field if non-nil, zero value otherwise.

### GetDiscoveryInfoOk

`func (o *CertificateResponse) GetDiscoveryInfoOk() (*[]DiscoveryInfo, bool)`

GetDiscoveryInfoOk returns a tuple with the DiscoveryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryInfo

`func (o *CertificateResponse) SetDiscoveryInfo(v []DiscoveryInfo)`

SetDiscoveryInfo sets DiscoveryInfo field to given value.

### HasDiscoveryInfo

`func (o *CertificateResponse) HasDiscoveryInfo() bool`

HasDiscoveryInfo returns a boolean if a field has been set.

### SetDiscoveryInfoNil

`func (o *CertificateResponse) SetDiscoveryInfoNil(b bool)`

 SetDiscoveryInfoNil sets the value for DiscoveryInfo to be an explicit nil

### UnsetDiscoveryInfo
`func (o *CertificateResponse) UnsetDiscoveryInfo()`

UnsetDiscoveryInfo ensures that no value is present for DiscoveryInfo, not even an explicit nil
### GetSubjectAlternateNames

`func (o *CertificateResponse) GetSubjectAlternateNames() []SubjectAlternateName`

GetSubjectAlternateNames returns the SubjectAlternateNames field if non-nil, zero value otherwise.

### GetSubjectAlternateNamesOk

`func (o *CertificateResponse) GetSubjectAlternateNamesOk() (*[]SubjectAlternateName, bool)`

GetSubjectAlternateNamesOk returns a tuple with the SubjectAlternateNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternateNames

`func (o *CertificateResponse) SetSubjectAlternateNames(v []SubjectAlternateName)`

SetSubjectAlternateNames sets SubjectAlternateNames field to given value.


### GetTriggerResults

`func (o *CertificateResponse) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *CertificateResponse) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *CertificateResponse) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *CertificateResponse) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *CertificateResponse) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *CertificateResponse) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil
### GetExtensions

`func (o *CertificateResponse) GetExtensions() []CertificateExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CertificateResponse) GetExtensionsOk() (*[]CertificateExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CertificateResponse) SetExtensions(v []CertificateExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CertificateResponse) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetSerial

`func (o *CertificateResponse) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CertificateResponse) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CertificateResponse) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetSigningAlgorithm

`func (o *CertificateResponse) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CertificateResponse) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CertificateResponse) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.


### GetDiscoveryData

`func (o *CertificateResponse) GetDiscoveryData() []HostDiscoveryData`

GetDiscoveryData returns the DiscoveryData field if non-nil, zero value otherwise.

### GetDiscoveryDataOk

`func (o *CertificateResponse) GetDiscoveryDataOk() (*[]HostDiscoveryData, bool)`

GetDiscoveryDataOk returns a tuple with the DiscoveryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryData

`func (o *CertificateResponse) SetDiscoveryData(v []HostDiscoveryData)`

SetDiscoveryData sets DiscoveryData field to given value.

### HasDiscoveryData

`func (o *CertificateResponse) HasDiscoveryData() bool`

HasDiscoveryData returns a boolean if a field has been set.

### SetDiscoveryDataNil

`func (o *CertificateResponse) SetDiscoveryDataNil(b bool)`

 SetDiscoveryDataNil sets the value for DiscoveryData to be an explicit nil

### UnsetDiscoveryData
`func (o *CertificateResponse) UnsetDiscoveryData()`

UnsetDiscoveryData ensures that no value is present for DiscoveryData, not even an explicit nil
### GetId

`func (o *CertificateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRevocationReason

`func (o *CertificateResponse) GetRevocationReason() string`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *CertificateResponse) GetRevocationReasonOk() (*string, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *CertificateResponse) SetRevocationReason(v string)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *CertificateResponse) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### SetRevocationReasonNil

`func (o *CertificateResponse) SetRevocationReasonNil(b bool)`

 SetRevocationReasonNil sets the value for RevocationReason to be an explicit nil

### UnsetRevocationReason
`func (o *CertificateResponse) UnsetRevocationReason()`

UnsetRevocationReason ensures that no value is present for RevocationReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


