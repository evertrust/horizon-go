# CertificateSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**ContactEmail** | Pointer to **NullableString** |  | [optional] 
**DiscoveredTrusted** | Pointer to **NullableBool** |  | [optional] 
**DiscoveryData** | Pointer to [**[]HostDiscoveryData**](HostDiscoveryData.md) |  | [optional] 
**DiscoveryInfo** | Pointer to [**[]DiscoveryInfo**](DiscoveryInfo.md) |  | [optional] 
**Dn** | Pointer to **NullableString** |  | [optional] 
**Grades** | Pointer to [**[]GradingPolicyResult**](GradingPolicyResult.md) |  | [optional] 
**HolderId** | Pointer to **NullableString** |  | [optional] 
**Issuer** | Pointer to **NullableString** |  | [optional] 
**KeyType** | Pointer to **NullableString** | One of &#x60;rsa-2048&#x60;, &#x60;rsa-3072&#x60;, &#x60;rsa-4096&#x60;, &#x60;rsa-8192&#x60;, &#x60;ec-secp256r1&#x60;, &#x60;ec-secp384r1&#x60;, &#x60;ec-secp521r1&#x60;, &#x60;ec-brainpoolp256r1&#x60;, &#x60;ec-brainpoolp384r1&#x60;, &#x60;ec-brainpoolp512r1&#x60;,  &#x60;ed-448&#x60;, &#x60;ed-25519&#x60;, &#x60;mldsa-44&#x60;, &#x60;mldsa-65&#x60;, &#x60;mldsa-87&#x60;, &#x60;slhdsa-sha2-128s&#x60;, &#x60;slhdsa-sha2-128f&#x60;, &#x60;slhdsa-sha2-192s&#x60;, &#x60;slhdsa-sha2-192f&#x60;, &#x60;slhdsa-sha2-256s&#x60;, &#x60;slhdsa-sha2-256f&#x60;, &#x60;slhdsa-sha2-128ssha256&#x60;, &#x60;slhdsa-sha2-128fsha256&#x60;, &#x60;slhdsa-sha2-192ssha512&#x60;, &#x60;slhdsa-sha2-192fsha512&#x60;, &#x60;slhdsa-sha2-256ssha512&#x60;, &#x60;slhdsa-sha2-256fsha512&#x60; or &#x60;&lt;primary key type&gt;+&lt;alternate key type&gt;&#x60; | [optional] 
**Labels** | Pointer to [**[]LabelData**](LabelData.md) |  | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadata**](CertificateMetadata.md) |  | [optional] 
**Module** | Pointer to **NullableString** |  | [optional] 
**NotAfter** | Pointer to **NullableInt64** |  | [optional] 
**NotBefore** | Pointer to **NullableInt64** |  | [optional] 
**Owner** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to [**NullableCertificatePermissions**](CertificatePermissions.md) |  | [optional] 
**PrivateKey** | Pointer to [**NullableEscrowedPrivateKey**](EscrowedPrivateKey.md) |  | [optional] 
**Profile** | Pointer to **NullableString** |  | [optional] 
**PublicKeyThumbprint** | Pointer to **NullableString** |  | [optional] 
**RevocationDate** | Pointer to **NullableInt64** |  | [optional] 
**RevocationReason** | Pointer to **NullableString** | One of: &#x60;unspecified&#x60;, &#x60;keycompromise&#x60;, &#x60;cacompromise&#x60;, &#x60;affiliationchange&#x60;, &#x60;superseded&#x60;, &#x60;cessationofoperation&#x60; | [optional] 
**SelfSigned** | Pointer to **NullableBool** |  | [optional] 
**Serial** | Pointer to **NullableString** |  | [optional] 
**SigningAlgorithm** | Pointer to **NullableString** |  | [optional] 
**SubjectAlternateNames** | Pointer to [**[]SubjectAlternateName**](SubjectAlternateName.md) |  | [optional] 
**Team** | Pointer to **NullableString** |  | [optional] 
**ThirdPartyData** | Pointer to [**[]ThirdPartyItem**](ThirdPartyItem.md) |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**TriggerResults** | Pointer to [**[]TriggerResult**](TriggerResult.md) |  | [optional] 

## Methods

### NewCertificateSearchResult

`func NewCertificateSearchResult() *CertificateSearchResult`

NewCertificateSearchResult instantiates a new CertificateSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSearchResultWithDefaults

`func NewCertificateSearchResultWithDefaults() *CertificateSearchResult`

NewCertificateSearchResultWithDefaults instantiates a new CertificateSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateSearchResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateSearchResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateSearchResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateSearchResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CertificateSearchResult) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CertificateSearchResult) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCertificate

`func (o *CertificateSearchResult) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateSearchResult) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateSearchResult) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateSearchResult) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *CertificateSearchResult) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *CertificateSearchResult) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetContactEmail

`func (o *CertificateSearchResult) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *CertificateSearchResult) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *CertificateSearchResult) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *CertificateSearchResult) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *CertificateSearchResult) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *CertificateSearchResult) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetDiscoveredTrusted

`func (o *CertificateSearchResult) GetDiscoveredTrusted() bool`

GetDiscoveredTrusted returns the DiscoveredTrusted field if non-nil, zero value otherwise.

### GetDiscoveredTrustedOk

`func (o *CertificateSearchResult) GetDiscoveredTrustedOk() (*bool, bool)`

GetDiscoveredTrustedOk returns a tuple with the DiscoveredTrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredTrusted

`func (o *CertificateSearchResult) SetDiscoveredTrusted(v bool)`

SetDiscoveredTrusted sets DiscoveredTrusted field to given value.

### HasDiscoveredTrusted

`func (o *CertificateSearchResult) HasDiscoveredTrusted() bool`

HasDiscoveredTrusted returns a boolean if a field has been set.

### SetDiscoveredTrustedNil

`func (o *CertificateSearchResult) SetDiscoveredTrustedNil(b bool)`

 SetDiscoveredTrustedNil sets the value for DiscoveredTrusted to be an explicit nil

### UnsetDiscoveredTrusted
`func (o *CertificateSearchResult) UnsetDiscoveredTrusted()`

UnsetDiscoveredTrusted ensures that no value is present for DiscoveredTrusted, not even an explicit nil
### GetDiscoveryData

`func (o *CertificateSearchResult) GetDiscoveryData() []HostDiscoveryData`

GetDiscoveryData returns the DiscoveryData field if non-nil, zero value otherwise.

### GetDiscoveryDataOk

`func (o *CertificateSearchResult) GetDiscoveryDataOk() (*[]HostDiscoveryData, bool)`

GetDiscoveryDataOk returns a tuple with the DiscoveryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryData

`func (o *CertificateSearchResult) SetDiscoveryData(v []HostDiscoveryData)`

SetDiscoveryData sets DiscoveryData field to given value.

### HasDiscoveryData

`func (o *CertificateSearchResult) HasDiscoveryData() bool`

HasDiscoveryData returns a boolean if a field has been set.

### SetDiscoveryDataNil

`func (o *CertificateSearchResult) SetDiscoveryDataNil(b bool)`

 SetDiscoveryDataNil sets the value for DiscoveryData to be an explicit nil

### UnsetDiscoveryData
`func (o *CertificateSearchResult) UnsetDiscoveryData()`

UnsetDiscoveryData ensures that no value is present for DiscoveryData, not even an explicit nil
### GetDiscoveryInfo

`func (o *CertificateSearchResult) GetDiscoveryInfo() []DiscoveryInfo`

GetDiscoveryInfo returns the DiscoveryInfo field if non-nil, zero value otherwise.

### GetDiscoveryInfoOk

`func (o *CertificateSearchResult) GetDiscoveryInfoOk() (*[]DiscoveryInfo, bool)`

GetDiscoveryInfoOk returns a tuple with the DiscoveryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryInfo

`func (o *CertificateSearchResult) SetDiscoveryInfo(v []DiscoveryInfo)`

SetDiscoveryInfo sets DiscoveryInfo field to given value.

### HasDiscoveryInfo

`func (o *CertificateSearchResult) HasDiscoveryInfo() bool`

HasDiscoveryInfo returns a boolean if a field has been set.

### SetDiscoveryInfoNil

`func (o *CertificateSearchResult) SetDiscoveryInfoNil(b bool)`

 SetDiscoveryInfoNil sets the value for DiscoveryInfo to be an explicit nil

### UnsetDiscoveryInfo
`func (o *CertificateSearchResult) UnsetDiscoveryInfo()`

UnsetDiscoveryInfo ensures that no value is present for DiscoveryInfo, not even an explicit nil
### GetDn

`func (o *CertificateSearchResult) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *CertificateSearchResult) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *CertificateSearchResult) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *CertificateSearchResult) HasDn() bool`

HasDn returns a boolean if a field has been set.

### SetDnNil

`func (o *CertificateSearchResult) SetDnNil(b bool)`

 SetDnNil sets the value for Dn to be an explicit nil

### UnsetDn
`func (o *CertificateSearchResult) UnsetDn()`

UnsetDn ensures that no value is present for Dn, not even an explicit nil
### GetGrades

`func (o *CertificateSearchResult) GetGrades() []GradingPolicyResult`

GetGrades returns the Grades field if non-nil, zero value otherwise.

### GetGradesOk

`func (o *CertificateSearchResult) GetGradesOk() (*[]GradingPolicyResult, bool)`

GetGradesOk returns a tuple with the Grades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrades

`func (o *CertificateSearchResult) SetGrades(v []GradingPolicyResult)`

SetGrades sets Grades field to given value.

### HasGrades

`func (o *CertificateSearchResult) HasGrades() bool`

HasGrades returns a boolean if a field has been set.

### SetGradesNil

`func (o *CertificateSearchResult) SetGradesNil(b bool)`

 SetGradesNil sets the value for Grades to be an explicit nil

### UnsetGrades
`func (o *CertificateSearchResult) UnsetGrades()`

UnsetGrades ensures that no value is present for Grades, not even an explicit nil
### GetHolderId

`func (o *CertificateSearchResult) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *CertificateSearchResult) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *CertificateSearchResult) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.

### HasHolderId

`func (o *CertificateSearchResult) HasHolderId() bool`

HasHolderId returns a boolean if a field has been set.

### SetHolderIdNil

`func (o *CertificateSearchResult) SetHolderIdNil(b bool)`

 SetHolderIdNil sets the value for HolderId to be an explicit nil

### UnsetHolderId
`func (o *CertificateSearchResult) UnsetHolderId()`

UnsetHolderId ensures that no value is present for HolderId, not even an explicit nil
### GetIssuer

`func (o *CertificateSearchResult) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateSearchResult) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateSearchResult) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertificateSearchResult) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *CertificateSearchResult) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *CertificateSearchResult) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetKeyType

`func (o *CertificateSearchResult) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CertificateSearchResult) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CertificateSearchResult) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CertificateSearchResult) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *CertificateSearchResult) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *CertificateSearchResult) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetLabels

`func (o *CertificateSearchResult) GetLabels() []LabelData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CertificateSearchResult) GetLabelsOk() (*[]LabelData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CertificateSearchResult) SetLabels(v []LabelData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CertificateSearchResult) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *CertificateSearchResult) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *CertificateSearchResult) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *CertificateSearchResult) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificateSearchResult) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificateSearchResult) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CertificateSearchResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CertificateSearchResult) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CertificateSearchResult) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetModule

`func (o *CertificateSearchResult) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CertificateSearchResult) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CertificateSearchResult) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *CertificateSearchResult) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *CertificateSearchResult) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *CertificateSearchResult) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetNotAfter

`func (o *CertificateSearchResult) GetNotAfter() int64`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateSearchResult) GetNotAfterOk() (*int64, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateSearchResult) SetNotAfter(v int64)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateSearchResult) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### SetNotAfterNil

`func (o *CertificateSearchResult) SetNotAfterNil(b bool)`

 SetNotAfterNil sets the value for NotAfter to be an explicit nil

### UnsetNotAfter
`func (o *CertificateSearchResult) UnsetNotAfter()`

UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil
### GetNotBefore

`func (o *CertificateSearchResult) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateSearchResult) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateSearchResult) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateSearchResult) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### SetNotBeforeNil

`func (o *CertificateSearchResult) SetNotBeforeNil(b bool)`

 SetNotBeforeNil sets the value for NotBefore to be an explicit nil

### UnsetNotBefore
`func (o *CertificateSearchResult) UnsetNotBefore()`

UnsetNotBefore ensures that no value is present for NotBefore, not even an explicit nil
### GetOwner

`func (o *CertificateSearchResult) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CertificateSearchResult) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CertificateSearchResult) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CertificateSearchResult) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *CertificateSearchResult) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *CertificateSearchResult) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetPermissions

`func (o *CertificateSearchResult) GetPermissions() CertificatePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CertificateSearchResult) GetPermissionsOk() (*CertificatePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CertificateSearchResult) SetPermissions(v CertificatePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CertificateSearchResult) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CertificateSearchResult) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CertificateSearchResult) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPrivateKey

`func (o *CertificateSearchResult) GetPrivateKey() EscrowedPrivateKey`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CertificateSearchResult) GetPrivateKeyOk() (*EscrowedPrivateKey, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CertificateSearchResult) SetPrivateKey(v EscrowedPrivateKey)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CertificateSearchResult) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *CertificateSearchResult) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *CertificateSearchResult) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetProfile

`func (o *CertificateSearchResult) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CertificateSearchResult) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CertificateSearchResult) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CertificateSearchResult) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *CertificateSearchResult) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *CertificateSearchResult) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetPublicKeyThumbprint

`func (o *CertificateSearchResult) GetPublicKeyThumbprint() string`

GetPublicKeyThumbprint returns the PublicKeyThumbprint field if non-nil, zero value otherwise.

### GetPublicKeyThumbprintOk

`func (o *CertificateSearchResult) GetPublicKeyThumbprintOk() (*string, bool)`

GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyThumbprint

`func (o *CertificateSearchResult) SetPublicKeyThumbprint(v string)`

SetPublicKeyThumbprint sets PublicKeyThumbprint field to given value.

### HasPublicKeyThumbprint

`func (o *CertificateSearchResult) HasPublicKeyThumbprint() bool`

HasPublicKeyThumbprint returns a boolean if a field has been set.

### SetPublicKeyThumbprintNil

`func (o *CertificateSearchResult) SetPublicKeyThumbprintNil(b bool)`

 SetPublicKeyThumbprintNil sets the value for PublicKeyThumbprint to be an explicit nil

### UnsetPublicKeyThumbprint
`func (o *CertificateSearchResult) UnsetPublicKeyThumbprint()`

UnsetPublicKeyThumbprint ensures that no value is present for PublicKeyThumbprint, not even an explicit nil
### GetRevocationDate

`func (o *CertificateSearchResult) GetRevocationDate() int64`

GetRevocationDate returns the RevocationDate field if non-nil, zero value otherwise.

### GetRevocationDateOk

`func (o *CertificateSearchResult) GetRevocationDateOk() (*int64, bool)`

GetRevocationDateOk returns a tuple with the RevocationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationDate

`func (o *CertificateSearchResult) SetRevocationDate(v int64)`

SetRevocationDate sets RevocationDate field to given value.

### HasRevocationDate

`func (o *CertificateSearchResult) HasRevocationDate() bool`

HasRevocationDate returns a boolean if a field has been set.

### SetRevocationDateNil

`func (o *CertificateSearchResult) SetRevocationDateNil(b bool)`

 SetRevocationDateNil sets the value for RevocationDate to be an explicit nil

### UnsetRevocationDate
`func (o *CertificateSearchResult) UnsetRevocationDate()`

UnsetRevocationDate ensures that no value is present for RevocationDate, not even an explicit nil
### GetRevocationReason

`func (o *CertificateSearchResult) GetRevocationReason() string`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *CertificateSearchResult) GetRevocationReasonOk() (*string, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *CertificateSearchResult) SetRevocationReason(v string)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *CertificateSearchResult) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### SetRevocationReasonNil

`func (o *CertificateSearchResult) SetRevocationReasonNil(b bool)`

 SetRevocationReasonNil sets the value for RevocationReason to be an explicit nil

### UnsetRevocationReason
`func (o *CertificateSearchResult) UnsetRevocationReason()`

UnsetRevocationReason ensures that no value is present for RevocationReason, not even an explicit nil
### GetSelfSigned

`func (o *CertificateSearchResult) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *CertificateSearchResult) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *CertificateSearchResult) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.

### HasSelfSigned

`func (o *CertificateSearchResult) HasSelfSigned() bool`

HasSelfSigned returns a boolean if a field has been set.

### SetSelfSignedNil

`func (o *CertificateSearchResult) SetSelfSignedNil(b bool)`

 SetSelfSignedNil sets the value for SelfSigned to be an explicit nil

### UnsetSelfSigned
`func (o *CertificateSearchResult) UnsetSelfSigned()`

UnsetSelfSigned ensures that no value is present for SelfSigned, not even an explicit nil
### GetSerial

`func (o *CertificateSearchResult) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CertificateSearchResult) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CertificateSearchResult) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CertificateSearchResult) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *CertificateSearchResult) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *CertificateSearchResult) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetSigningAlgorithm

`func (o *CertificateSearchResult) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CertificateSearchResult) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CertificateSearchResult) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *CertificateSearchResult) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### SetSigningAlgorithmNil

`func (o *CertificateSearchResult) SetSigningAlgorithmNil(b bool)`

 SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil

### UnsetSigningAlgorithm
`func (o *CertificateSearchResult) UnsetSigningAlgorithm()`

UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
### GetSubjectAlternateNames

`func (o *CertificateSearchResult) GetSubjectAlternateNames() []SubjectAlternateName`

GetSubjectAlternateNames returns the SubjectAlternateNames field if non-nil, zero value otherwise.

### GetSubjectAlternateNamesOk

`func (o *CertificateSearchResult) GetSubjectAlternateNamesOk() (*[]SubjectAlternateName, bool)`

GetSubjectAlternateNamesOk returns a tuple with the SubjectAlternateNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternateNames

`func (o *CertificateSearchResult) SetSubjectAlternateNames(v []SubjectAlternateName)`

SetSubjectAlternateNames sets SubjectAlternateNames field to given value.

### HasSubjectAlternateNames

`func (o *CertificateSearchResult) HasSubjectAlternateNames() bool`

HasSubjectAlternateNames returns a boolean if a field has been set.

### SetSubjectAlternateNamesNil

`func (o *CertificateSearchResult) SetSubjectAlternateNamesNil(b bool)`

 SetSubjectAlternateNamesNil sets the value for SubjectAlternateNames to be an explicit nil

### UnsetSubjectAlternateNames
`func (o *CertificateSearchResult) UnsetSubjectAlternateNames()`

UnsetSubjectAlternateNames ensures that no value is present for SubjectAlternateNames, not even an explicit nil
### GetTeam

`func (o *CertificateSearchResult) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *CertificateSearchResult) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *CertificateSearchResult) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *CertificateSearchResult) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *CertificateSearchResult) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *CertificateSearchResult) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetThirdPartyData

`func (o *CertificateSearchResult) GetThirdPartyData() []ThirdPartyItem`

GetThirdPartyData returns the ThirdPartyData field if non-nil, zero value otherwise.

### GetThirdPartyDataOk

`func (o *CertificateSearchResult) GetThirdPartyDataOk() (*[]ThirdPartyItem, bool)`

GetThirdPartyDataOk returns a tuple with the ThirdPartyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyData

`func (o *CertificateSearchResult) SetThirdPartyData(v []ThirdPartyItem)`

SetThirdPartyData sets ThirdPartyData field to given value.

### HasThirdPartyData

`func (o *CertificateSearchResult) HasThirdPartyData() bool`

HasThirdPartyData returns a boolean if a field has been set.

### SetThirdPartyDataNil

`func (o *CertificateSearchResult) SetThirdPartyDataNil(b bool)`

 SetThirdPartyDataNil sets the value for ThirdPartyData to be an explicit nil

### UnsetThirdPartyData
`func (o *CertificateSearchResult) UnsetThirdPartyData()`

UnsetThirdPartyData ensures that no value is present for ThirdPartyData, not even an explicit nil
### GetThumbprint

`func (o *CertificateSearchResult) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CertificateSearchResult) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CertificateSearchResult) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CertificateSearchResult) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CertificateSearchResult) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CertificateSearchResult) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetTriggerResults

`func (o *CertificateSearchResult) GetTriggerResults() []TriggerResult`

GetTriggerResults returns the TriggerResults field if non-nil, zero value otherwise.

### GetTriggerResultsOk

`func (o *CertificateSearchResult) GetTriggerResultsOk() (*[]TriggerResult, bool)`

GetTriggerResultsOk returns a tuple with the TriggerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerResults

`func (o *CertificateSearchResult) SetTriggerResults(v []TriggerResult)`

SetTriggerResults sets TriggerResults field to given value.

### HasTriggerResults

`func (o *CertificateSearchResult) HasTriggerResults() bool`

HasTriggerResults returns a boolean if a field has been set.

### SetTriggerResultsNil

`func (o *CertificateSearchResult) SetTriggerResultsNil(b bool)`

 SetTriggerResultsNil sets the value for TriggerResults to be an explicit nil

### UnsetTriggerResults
`func (o *CertificateSearchResult) UnsetTriggerResults()`

UnsetTriggerResults ensures that no value is present for TriggerResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


