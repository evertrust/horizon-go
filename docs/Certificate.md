# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AutoRenew** | **bool** | &#x60;true&#x60; if auto renewal is enabled on this certificate | 
**Certificate** | **string** | The certificate&#39;s PEM-encoded content | 
**ContactEmail** | Pointer to **NullableString** | The certificate&#39;s contact email. It will be used to send notifications about the certificate&#39;s expiration and revocation | [optional] 
**CrlSynchronized** | Pointer to **NullableBool** | Whether the certificate&#39;s revocation status is synchronized with a CRL | [optional] 
**DiscoveredTrusted** | Pointer to **NullableBool** | If the certificate was discovered and is found to be issued by an existing trusted CA, this field will be set to true. If the certificate was discovered and is not found to be issued by an existing trusted CA, this field will be set to false. If the certificate was not discovered, this field will be null | [optional] 
**DiscoveryData** | Pointer to [**[]HostDiscoveryData**](HostDiscoveryData.md) | A list of metadata containing information on where the certificate was discovered | [optional] 
**DiscoveryInfo** | Pointer to [**[]DiscoveryInfo**](DiscoveryInfo.md) | A list of metadata containing information on how and when the certificate was discovered | [optional] 
**Dn** | **string** | The certificate&#39;s Distinguished Name | 
**Escrowed** | **bool** | Whether the certificate is escrowed | 
**Extensions** | Pointer to [**[]CertificateExtension**](CertificateExtension.md) | The certificate&#39;s extensions | [optional] 
**Grades** | Pointer to [**[]GradingPolicyResult**](GradingPolicyResult.md) | The certificate&#39;s grades for the enabled grading policies | [optional] 
**HolderId** | **string** | The certificate&#39;s holder ID. This is a computed field that is used to count how many similar certificates are in use simultaneously by the same holder | 
**Issuer** | **string** | The certificate&#39;s issuer Distinguished Name | 
**KeyType** | **string** | The certificate&#39;s key type | 
**Labels** | Pointer to [**[]LabelData**](LabelData.md) | The certificate&#39;s labels | [optional] 
**Metadata** | [**[]CertificateMetadata**](CertificateMetadata.md) | The certificate&#39;s technical metadata used internally | 
**Module** | **string** | The certificate&#39;s module | 
**NotAfter** | **int64** | The certificate&#39;s expiration date in milliseconds since the epoch | 
**NotBefore** | **int64** | The certificate&#39;s start date in milliseconds since the epoch | 
**Owner** | Pointer to **NullableString** | The certificate&#39;s owner. This is a reference to a local identity identifier | [optional] 
**Profile** | Pointer to **NullableString** | The certificate&#39;s profile | [optional] 
**PublicKeyThumbprint** | **string** | The certificate&#39;s public key thumbprint | 
**RevocationDate** | Pointer to **NullableInt64** | The certificate&#39;s revocation date in milliseconds since the epoch. This field is only present if the certificate is revoked | [optional] 
**RevocationReason** | Pointer to **NullableString** | The certificate&#39;s revocation reason | [optional] 
**Revoked** | **bool** | Whether the certificate is revoked | 
**SelfSigned** | **bool** | Whether the certificate is self-signed | 
**Serial** | **string** | The certificate&#39;s serial number | 
**SigningAlgorithm** | **string** | The certificate&#39;s signing algorithm | 
**SubjectAlternateNames** | [**[]SubjectAlternateName**](SubjectAlternateName.md) | The certificate&#39;s Subject Alternate Names | 
**Team** | Pointer to **NullableString** | The certificate&#39;s team. This is a reference to a team identifier. It will be used to determine the certificate&#39;s permissions and send notifications | [optional] 
**ThirdPartyData** | Pointer to [**[]ThirdPartyItem**](ThirdPartyItem.md) | The certificate&#39;s information about synchronization with Horizon supported third parties | [optional] 
**Thumbprint** | **string** | The certificate&#39;s thumbprint | 
**TriggerResults** | Pointer to [**[]TriggerResult**](TriggerResult.md) | The result of the execution of triggers on this certificate | [optional] 

## Methods

### NewCertificate

`func NewCertificate(id string, autoRenew bool, certificate string, dn string, escrowed bool, holderId string, issuer string, keyType string, metadata []CertificateMetadata, module string, notAfter int64, notBefore int64, publicKeyThumbprint string, revoked bool, selfSigned bool, serial string, signingAlgorithm string, subjectAlternateNames []SubjectAlternateName, thumbprint string, ) *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Certificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Certificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Certificate) SetId(v string)`

SetId sets Id field to given value.


### GetAutoRenew

`func (o *Certificate) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *Certificate) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *Certificate) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.


### GetCertificate

`func (o *Certificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Certificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Certificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetContactEmail

`func (o *Certificate) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Certificate) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Certificate) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Certificate) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *Certificate) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *Certificate) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetCrlSynchronized

`func (o *Certificate) GetCrlSynchronized() bool`

GetCrlSynchronized returns the CrlSynchronized field if non-nil, zero value otherwise.

### GetCrlSynchronizedOk

`func (o *Certificate) GetCrlSynchronizedOk() (*bool, bool)`

GetCrlSynchronizedOk returns a tuple with the CrlSynchronized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlSynchronized

`func (o *Certificate) SetCrlSynchronized(v bool)`

SetCrlSynchronized sets CrlSynchronized field to given value.

### HasCrlSynchronized

`func (o *Certificate) HasCrlSynchronized() bool`

HasCrlSynchronized returns a boolean if a field has been set.

### SetCrlSynchronizedNil

`func (o *Certificate) SetCrlSynchronizedNil(b bool)`

 SetCrlSynchronizedNil sets the value for CrlSynchronized to be an explicit nil

### UnsetCrlSynchronized
`func (o *Certificate) UnsetCrlSynchronized()`

UnsetCrlSynchronized ensures that no value is present for CrlSynchronized, not even an explicit nil
### GetDiscoveredTrusted

`func (o *Certificate) GetDiscoveredTrusted() bool`

GetDiscoveredTrusted returns the DiscoveredTrusted field if non-nil, zero value otherwise.

### GetDiscoveredTrustedOk

`func (o *Certificate) GetDiscoveredTrustedOk() (*bool, bool)`

GetDiscoveredTrustedOk returns a tuple with the DiscoveredTrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredTrusted

`func (o *Certificate) SetDiscoveredTrusted(v bool)`

SetDiscoveredTrusted sets DiscoveredTrusted field to given value.

### HasDiscoveredTrusted

`func (o *Certificate) HasDiscoveredTrusted() bool`

HasDiscoveredTrusted returns a boolean if a field has been set.

### SetDiscoveredTrustedNil

`func (o *Certificate) SetDiscoveredTrustedNil(b bool)`

 SetDiscoveredTrustedNil sets the value for DiscoveredTrusted to be an explicit nil

### UnsetDiscoveredTrusted
`func (o *Certificate) UnsetDiscoveredTrusted()`

UnsetDiscoveredTrusted ensures that no value is present for DiscoveredTrusted, not even an explicit nil
### GetDiscoveryData

`func (o *Certificate) GetDiscoveryData() []HostDiscoveryData`

GetDiscoveryData returns the DiscoveryData field if non-nil, zero value otherwise.

### GetDiscoveryDataOk

`func (o *Certificate) GetDiscoveryDataOk() (*[]HostDiscoveryData, bool)`

GetDiscoveryDataOk returns a tuple with the DiscoveryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryData

`func (o *Certificate) SetDiscoveryData(v []HostDiscoveryData)`

SetDiscoveryData sets DiscoveryData field to given value.

### HasDiscoveryData

`func (o *Certificate) HasDiscoveryData() bool`

HasDiscoveryData returns a boolean if a field has been set.

### SetDiscoveryDataNil

`func (o *Certificate) SetDiscoveryDataNil(b bool)`

 SetDiscoveryDataNil sets the value for DiscoveryData to be an explicit nil

### UnsetDiscoveryData
`func (o *Certificate) UnsetDiscoveryData()`

UnsetDiscoveryData ensures that no value is present for DiscoveryData, not even an explicit nil
### GetDiscoveryInfo

`func (o *Certificate) GetDiscoveryInfo() []DiscoveryInfo`

GetDiscoveryInfo returns the DiscoveryInfo field if non-nil, zero value otherwise.

### GetDiscoveryInfoOk

`func (o *Certificate) GetDiscoveryInfoOk() (*[]DiscoveryInfo, bool)`

GetDiscoveryInfoOk returns a tuple with the DiscoveryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryInfo

`func (o *Certificate) SetDiscoveryInfo(v []DiscoveryInfo)`

SetDiscoveryInfo sets DiscoveryInfo field to given value.

### HasDiscoveryInfo

`func (o *Certificate) HasDiscoveryInfo() bool`

HasDiscoveryInfo returns a boolean if a field has been set.

### SetDiscoveryInfoNil

`func (o *Certificate) SetDiscoveryInfoNil(b bool)`

 SetDiscoveryInfoNil sets the value for DiscoveryInfo to be an explicit nil

### UnsetDiscoveryInfo
`func (o *Certificate) UnsetDiscoveryInfo()`

UnsetDiscoveryInfo ensures that no value is present for DiscoveryInfo, not even an explicit nil
### GetDn

`func (o *Certificate) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *Certificate) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *Certificate) SetDn(v string)`

SetDn sets Dn field to given value.


### GetEscrowed

`func (o *Certificate) GetEscrowed() bool`

GetEscrowed returns the Escrowed field if non-nil, zero value otherwise.

### GetEscrowedOk

`func (o *Certificate) GetEscrowedOk() (*bool, bool)`

GetEscrowedOk returns a tuple with the Escrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscrowed

`func (o *Certificate) SetEscrowed(v bool)`

SetEscrowed sets Escrowed field to given value.


### GetExtensions

`func (o *Certificate) GetExtensions() []CertificateExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Certificate) GetExtensionsOk() (*[]CertificateExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Certificate) SetExtensions(v []CertificateExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Certificate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGrades

`func (o *Certificate) GetGrades() []GradingPolicyResult`

GetGrades returns the Grades field if non-nil, zero value otherwise.

### GetGradesOk

`func (o *Certificate) GetGradesOk() (*[]GradingPolicyResult, bool)`

GetGradesOk returns a tuple with the Grades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrades

`func (o *Certificate) SetGrades(v []GradingPolicyResult)`

SetGrades sets Grades field to given value.

### HasGrades

`func (o *Certificate) HasGrades() bool`

HasGrades returns a boolean if a field has been set.

### SetGradesNil

`func (o *Certificate) SetGradesNil(b bool)`

 SetGradesNil sets the value for Grades to be an explicit nil

### UnsetGrades
`func (o *Certificate) UnsetGrades()`

UnsetGrades ensures that no value is present for Grades, not even an explicit nil
### GetHolderId

`func (o *Certificate) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *Certificate) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *Certificate) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.


### GetIssuer

`func (o *Certificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Certificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Certificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetKeyType

`func (o *Certificate) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *Certificate) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *Certificate) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetLabels

`func (o *Certificate) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Certificate) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Certificate) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Certificate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *Certificate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Certificate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *Certificate) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Certificate) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Certificate) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.


### GetModule

`func (o *Certificate) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *Certificate) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *Certificate) SetModule(v string)`

SetModule sets Module field to given value.


### GetNotAfter

`func (o *Certificate) GetNotAfter() int64`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *Certificate) GetNotAfterOk() (*int64, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *Certificate) SetNotAfter(v int64)`

SetNotAfter sets NotAfter field to given value.


### GetNotBefore

`func (o *Certificate) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *Certificate) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *Certificate) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.


### GetOwner

`func (o *Certificate) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Certificate) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Certificate) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Certificate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Certificate) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Certificate) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetProfile

`func (o *Certificate) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Certificate) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Certificate) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Certificate) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *Certificate) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *Certificate) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetPublicKeyThumbprint

`func (o *Certificate) GetPublicKeyThumbprint() string`

GetPublicKeyThumbprint returns the PublicKeyThumbprint field if non-nil, zero value otherwise.

### GetPublicKeyThumbprintOk

`func (o *Certificate) GetPublicKeyThumbprintOk() (*string, bool)`

GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyThumbprint

`func (o *Certificate) SetPublicKeyThumbprint(v string)`

SetPublicKeyThumbprint sets PublicKeyThumbprint field to given value.


### GetRevocationDate

`func (o *Certificate) GetRevocationDate() int64`

GetRevocationDate returns the RevocationDate field if non-nil, zero value otherwise.

### GetRevocationDateOk

`func (o *Certificate) GetRevocationDateOk() (*int64, bool)`

GetRevocationDateOk returns a tuple with the RevocationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationDate

`func (o *Certificate) SetRevocationDate(v int64)`

SetRevocationDate sets RevocationDate field to given value.

### HasRevocationDate

`func (o *Certificate) HasRevocationDate() bool`

HasRevocationDate returns a boolean if a field has been set.

### SetRevocationDateNil

`func (o *Certificate) SetRevocationDateNil(b bool)`

 SetRevocationDateNil sets the value for RevocationDate to be an explicit nil

### UnsetRevocationDate
`func (o *Certificate) UnsetRevocationDate()`

UnsetRevocationDate ensures that no value is present for RevocationDate, not even an explicit nil
### GetRevocationReason

`func (o *Certificate) GetRevocationReason() string`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *Certificate) GetRevocationReasonOk() (*string, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *Certificate) SetRevocationReason(v string)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *Certificate) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### SetRevocationReasonNil

`func (o *Certificate) SetRevocationReasonNil(b bool)`

 SetRevocationReasonNil sets the value for RevocationReason to be an explicit nil

### UnsetRevocationReason
`func (o *Certificate) UnsetRevocationReason()`

UnsetRevocationReason ensures that no value is present for RevocationReason, not even an explicit nil
### GetRevoked

`func (o *Certificate) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *Certificate) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *Certificate) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.


### GetSelfSigned

`func (o *Certificate) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *Certificate) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *Certificate) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.


### GetSerial

`func (o *Certificate) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Certificate) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Certificate) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetSigningAlgorithm

`func (o *Certificate) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *Certificate) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *Certificate) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.


### GetSubjectAlternateNames

`func (o *Certificate) GetSubjectAlternateNames() []SubjectAlternateName`

GetSubjectAlternateNames returns the SubjectAlternateNames field if non-nil, zero value otherwise.

### GetSubjectAlternateNamesOk

`func (o *Certificate) GetSubjectAlternateNamesOk() (*[]SubjectAlternateName, bool)`

GetSubjectAlternateNamesOk returns a tuple with the SubjectAlternateNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternateNames

`func (o *Certificate) SetSubjectAlternateNames(v []SubjectAlternateName)`

SetSubjectAlternateNames sets SubjectAlternateNames field to given value.


### GetTeam

`func (o *Certificate) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *Certificate) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *Certificate) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *Certificate) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *Certificate) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *Certificate) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetThirdPartyData

`func (o *Certificate) GetThirdPartyData() []ThirdPartyItem`

GetThirdPartyData returns the ThirdPartyData field if non-nil, zero value otherwise.

### GetThirdPartyDataOk

`func (o *Certificate) GetThirdPartyDataOk() (*[]ThirdPartyItem, bool)`

GetThirdPartyDataOk returns a tuple with the ThirdPartyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyData

`func (o *Certificate) SetThirdPartyData(v []ThirdPartyItem)`

SetThirdPartyData sets ThirdPartyData field to given value.

### HasThirdPartyData

`func (o *Certificate) HasThirdPartyData() bool`

HasThirdPartyData returns a boolean if a field has been set.

### SetThirdPartyDataNil

`func (o *Certificate) SetThirdPartyDataNil(b bool)`

 SetThirdPartyDataNil sets the value for ThirdPartyData to be an explicit nil

### UnsetThirdPartyData
`func (o *Certificate) UnsetThirdPartyData()`

UnsetThirdPartyData ensures that no value is present for ThirdPartyData, not even an explicit nil
### GetThumbprint

`func (o *Certificate) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *Certificate) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *Certificate) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.


### GetTriggerResults

`func (o *Certificate) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *Certificate) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *Certificate) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *Certificate) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *Certificate) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *Certificate) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


