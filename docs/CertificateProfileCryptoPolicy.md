# CertificateProfileCryptoPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Centralized** | Pointer to **NullableBool** | Whether this profile supports centralized enrollment | [optional] [default to false]
**Decentralized** | Pointer to **NullableBool** | Whether this profile supports decentralized enrollment | [optional] [default to false]
**DefaultKeyType** | Pointer to **NullableString** | Default key type used for centralized enrollment | [optional] 
**AuthorizedKeyTypes** | Pointer to **[]string** | List of authorized key types for enrollment | [optional] 
**PreferredEnrollmentMode** | Pointer to **NullableString** | If both centralized and decentralized enrollment are supported, this is the preferred mode | [optional] 
**Escrow** | Pointer to **NullableBool** | Whether this profile will escrow the certificate private keys | [optional] [default to false]
**P12passwordPolicy** | Pointer to **NullableString** | Password policy for the P12 file | [optional] 
**P12passwordMode** | Pointer to **NullableString** | Whether the user will be required to input their PKCS#12 password upon enrollment | [optional] 
**P12storeEncryptionType** | Pointer to **NullableString** | Encryption type for the P12 file | [optional] 
**ShowP12PasswordOnEnroll** | Pointer to **NullableBool** | Whether the PKCS#12 password will be displayed to the user upon enrollment | [optional] 
**ShowP12OnEnroll** | Pointer to **NullableBool** | Whether the PKCS#12 file will be displayed to the user upon enrollment | [optional] 
**ShowP12PasswordOnRecover** | Pointer to **NullableBool** | Whether the PKCS#12 password will be displayed to the user upon recovery | [optional] 
**ShowP12OnRecover** | Pointer to **NullableBool** | Whether the PKCS#12 file will be displayed to the user upon recovery | [optional] 

## Methods

### NewCertificateProfileCryptoPolicy

`func NewCertificateProfileCryptoPolicy() *CertificateProfileCryptoPolicy`

NewCertificateProfileCryptoPolicy instantiates a new CertificateProfileCryptoPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileCryptoPolicyWithDefaults

`func NewCertificateProfileCryptoPolicyWithDefaults() *CertificateProfileCryptoPolicy`

NewCertificateProfileCryptoPolicyWithDefaults instantiates a new CertificateProfileCryptoPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCentralized

`func (o *CertificateProfileCryptoPolicy) GetCentralized() bool`

GetCentralized returns the Centralized field if non-nil, zero value otherwise.

### GetCentralizedOk

`func (o *CertificateProfileCryptoPolicy) GetCentralizedOk() (*bool, bool)`

GetCentralizedOk returns a tuple with the Centralized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentralized

`func (o *CertificateProfileCryptoPolicy) SetCentralized(v bool)`

SetCentralized sets Centralized field to given value.

### HasCentralized

`func (o *CertificateProfileCryptoPolicy) HasCentralized() bool`

HasCentralized returns a boolean if a field has been set.

### SetCentralizedNil

`func (o *CertificateProfileCryptoPolicy) SetCentralizedNil(b bool)`

 SetCentralizedNil sets the value for Centralized to be an explicit nil

### UnsetCentralized
`func (o *CertificateProfileCryptoPolicy) UnsetCentralized()`

UnsetCentralized ensures that no value is present for Centralized, not even an explicit nil
### GetDecentralized

`func (o *CertificateProfileCryptoPolicy) GetDecentralized() bool`

GetDecentralized returns the Decentralized field if non-nil, zero value otherwise.

### GetDecentralizedOk

`func (o *CertificateProfileCryptoPolicy) GetDecentralizedOk() (*bool, bool)`

GetDecentralizedOk returns a tuple with the Decentralized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecentralized

`func (o *CertificateProfileCryptoPolicy) SetDecentralized(v bool)`

SetDecentralized sets Decentralized field to given value.

### HasDecentralized

`func (o *CertificateProfileCryptoPolicy) HasDecentralized() bool`

HasDecentralized returns a boolean if a field has been set.

### SetDecentralizedNil

`func (o *CertificateProfileCryptoPolicy) SetDecentralizedNil(b bool)`

 SetDecentralizedNil sets the value for Decentralized to be an explicit nil

### UnsetDecentralized
`func (o *CertificateProfileCryptoPolicy) UnsetDecentralized()`

UnsetDecentralized ensures that no value is present for Decentralized, not even an explicit nil
### GetDefaultKeyType

`func (o *CertificateProfileCryptoPolicy) GetDefaultKeyType() string`

GetDefaultKeyType returns the DefaultKeyType field if non-nil, zero value otherwise.

### GetDefaultKeyTypeOk

`func (o *CertificateProfileCryptoPolicy) GetDefaultKeyTypeOk() (*string, bool)`

GetDefaultKeyTypeOk returns a tuple with the DefaultKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeyType

`func (o *CertificateProfileCryptoPolicy) SetDefaultKeyType(v string)`

SetDefaultKeyType sets DefaultKeyType field to given value.

### HasDefaultKeyType

`func (o *CertificateProfileCryptoPolicy) HasDefaultKeyType() bool`

HasDefaultKeyType returns a boolean if a field has been set.

### SetDefaultKeyTypeNil

`func (o *CertificateProfileCryptoPolicy) SetDefaultKeyTypeNil(b bool)`

 SetDefaultKeyTypeNil sets the value for DefaultKeyType to be an explicit nil

### UnsetDefaultKeyType
`func (o *CertificateProfileCryptoPolicy) UnsetDefaultKeyType()`

UnsetDefaultKeyType ensures that no value is present for DefaultKeyType, not even an explicit nil
### GetAuthorizedKeyTypes

`func (o *CertificateProfileCryptoPolicy) GetAuthorizedKeyTypes() []string`

GetAuthorizedKeyTypes returns the AuthorizedKeyTypes field if non-nil, zero value otherwise.

### GetAuthorizedKeyTypesOk

`func (o *CertificateProfileCryptoPolicy) GetAuthorizedKeyTypesOk() (*[]string, bool)`

GetAuthorizedKeyTypesOk returns a tuple with the AuthorizedKeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeyTypes

`func (o *CertificateProfileCryptoPolicy) SetAuthorizedKeyTypes(v []string)`

SetAuthorizedKeyTypes sets AuthorizedKeyTypes field to given value.

### HasAuthorizedKeyTypes

`func (o *CertificateProfileCryptoPolicy) HasAuthorizedKeyTypes() bool`

HasAuthorizedKeyTypes returns a boolean if a field has been set.

### SetAuthorizedKeyTypesNil

`func (o *CertificateProfileCryptoPolicy) SetAuthorizedKeyTypesNil(b bool)`

 SetAuthorizedKeyTypesNil sets the value for AuthorizedKeyTypes to be an explicit nil

### UnsetAuthorizedKeyTypes
`func (o *CertificateProfileCryptoPolicy) UnsetAuthorizedKeyTypes()`

UnsetAuthorizedKeyTypes ensures that no value is present for AuthorizedKeyTypes, not even an explicit nil
### GetPreferredEnrollmentMode

`func (o *CertificateProfileCryptoPolicy) GetPreferredEnrollmentMode() string`

GetPreferredEnrollmentMode returns the PreferredEnrollmentMode field if non-nil, zero value otherwise.

### GetPreferredEnrollmentModeOk

`func (o *CertificateProfileCryptoPolicy) GetPreferredEnrollmentModeOk() (*string, bool)`

GetPreferredEnrollmentModeOk returns a tuple with the PreferredEnrollmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredEnrollmentMode

`func (o *CertificateProfileCryptoPolicy) SetPreferredEnrollmentMode(v string)`

SetPreferredEnrollmentMode sets PreferredEnrollmentMode field to given value.

### HasPreferredEnrollmentMode

`func (o *CertificateProfileCryptoPolicy) HasPreferredEnrollmentMode() bool`

HasPreferredEnrollmentMode returns a boolean if a field has been set.

### SetPreferredEnrollmentModeNil

`func (o *CertificateProfileCryptoPolicy) SetPreferredEnrollmentModeNil(b bool)`

 SetPreferredEnrollmentModeNil sets the value for PreferredEnrollmentMode to be an explicit nil

### UnsetPreferredEnrollmentMode
`func (o *CertificateProfileCryptoPolicy) UnsetPreferredEnrollmentMode()`

UnsetPreferredEnrollmentMode ensures that no value is present for PreferredEnrollmentMode, not even an explicit nil
### GetEscrow

`func (o *CertificateProfileCryptoPolicy) GetEscrow() bool`

GetEscrow returns the Escrow field if non-nil, zero value otherwise.

### GetEscrowOk

`func (o *CertificateProfileCryptoPolicy) GetEscrowOk() (*bool, bool)`

GetEscrowOk returns a tuple with the Escrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscrow

`func (o *CertificateProfileCryptoPolicy) SetEscrow(v bool)`

SetEscrow sets Escrow field to given value.

### HasEscrow

`func (o *CertificateProfileCryptoPolicy) HasEscrow() bool`

HasEscrow returns a boolean if a field has been set.

### SetEscrowNil

`func (o *CertificateProfileCryptoPolicy) SetEscrowNil(b bool)`

 SetEscrowNil sets the value for Escrow to be an explicit nil

### UnsetEscrow
`func (o *CertificateProfileCryptoPolicy) UnsetEscrow()`

UnsetEscrow ensures that no value is present for Escrow, not even an explicit nil
### GetP12passwordPolicy

`func (o *CertificateProfileCryptoPolicy) GetP12passwordPolicy() string`

GetP12passwordPolicy returns the P12passwordPolicy field if non-nil, zero value otherwise.

### GetP12passwordPolicyOk

`func (o *CertificateProfileCryptoPolicy) GetP12passwordPolicyOk() (*string, bool)`

GetP12passwordPolicyOk returns a tuple with the P12passwordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP12passwordPolicy

`func (o *CertificateProfileCryptoPolicy) SetP12passwordPolicy(v string)`

SetP12passwordPolicy sets P12passwordPolicy field to given value.

### HasP12passwordPolicy

`func (o *CertificateProfileCryptoPolicy) HasP12passwordPolicy() bool`

HasP12passwordPolicy returns a boolean if a field has been set.

### SetP12passwordPolicyNil

`func (o *CertificateProfileCryptoPolicy) SetP12passwordPolicyNil(b bool)`

 SetP12passwordPolicyNil sets the value for P12passwordPolicy to be an explicit nil

### UnsetP12passwordPolicy
`func (o *CertificateProfileCryptoPolicy) UnsetP12passwordPolicy()`

UnsetP12passwordPolicy ensures that no value is present for P12passwordPolicy, not even an explicit nil
### GetP12passwordMode

`func (o *CertificateProfileCryptoPolicy) GetP12passwordMode() string`

GetP12passwordMode returns the P12passwordMode field if non-nil, zero value otherwise.

### GetP12passwordModeOk

`func (o *CertificateProfileCryptoPolicy) GetP12passwordModeOk() (*string, bool)`

GetP12passwordModeOk returns a tuple with the P12passwordMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP12passwordMode

`func (o *CertificateProfileCryptoPolicy) SetP12passwordMode(v string)`

SetP12passwordMode sets P12passwordMode field to given value.

### HasP12passwordMode

`func (o *CertificateProfileCryptoPolicy) HasP12passwordMode() bool`

HasP12passwordMode returns a boolean if a field has been set.

### SetP12passwordModeNil

`func (o *CertificateProfileCryptoPolicy) SetP12passwordModeNil(b bool)`

 SetP12passwordModeNil sets the value for P12passwordMode to be an explicit nil

### UnsetP12passwordMode
`func (o *CertificateProfileCryptoPolicy) UnsetP12passwordMode()`

UnsetP12passwordMode ensures that no value is present for P12passwordMode, not even an explicit nil
### GetP12storeEncryptionType

`func (o *CertificateProfileCryptoPolicy) GetP12storeEncryptionType() string`

GetP12storeEncryptionType returns the P12storeEncryptionType field if non-nil, zero value otherwise.

### GetP12storeEncryptionTypeOk

`func (o *CertificateProfileCryptoPolicy) GetP12storeEncryptionTypeOk() (*string, bool)`

GetP12storeEncryptionTypeOk returns a tuple with the P12storeEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP12storeEncryptionType

`func (o *CertificateProfileCryptoPolicy) SetP12storeEncryptionType(v string)`

SetP12storeEncryptionType sets P12storeEncryptionType field to given value.

### HasP12storeEncryptionType

`func (o *CertificateProfileCryptoPolicy) HasP12storeEncryptionType() bool`

HasP12storeEncryptionType returns a boolean if a field has been set.

### SetP12storeEncryptionTypeNil

`func (o *CertificateProfileCryptoPolicy) SetP12storeEncryptionTypeNil(b bool)`

 SetP12storeEncryptionTypeNil sets the value for P12storeEncryptionType to be an explicit nil

### UnsetP12storeEncryptionType
`func (o *CertificateProfileCryptoPolicy) UnsetP12storeEncryptionType()`

UnsetP12storeEncryptionType ensures that no value is present for P12storeEncryptionType, not even an explicit nil
### GetShowP12PasswordOnEnroll

`func (o *CertificateProfileCryptoPolicy) GetShowP12PasswordOnEnroll() bool`

GetShowP12PasswordOnEnroll returns the ShowP12PasswordOnEnroll field if non-nil, zero value otherwise.

### GetShowP12PasswordOnEnrollOk

`func (o *CertificateProfileCryptoPolicy) GetShowP12PasswordOnEnrollOk() (*bool, bool)`

GetShowP12PasswordOnEnrollOk returns a tuple with the ShowP12PasswordOnEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowP12PasswordOnEnroll

`func (o *CertificateProfileCryptoPolicy) SetShowP12PasswordOnEnroll(v bool)`

SetShowP12PasswordOnEnroll sets ShowP12PasswordOnEnroll field to given value.

### HasShowP12PasswordOnEnroll

`func (o *CertificateProfileCryptoPolicy) HasShowP12PasswordOnEnroll() bool`

HasShowP12PasswordOnEnroll returns a boolean if a field has been set.

### SetShowP12PasswordOnEnrollNil

`func (o *CertificateProfileCryptoPolicy) SetShowP12PasswordOnEnrollNil(b bool)`

 SetShowP12PasswordOnEnrollNil sets the value for ShowP12PasswordOnEnroll to be an explicit nil

### UnsetShowP12PasswordOnEnroll
`func (o *CertificateProfileCryptoPolicy) UnsetShowP12PasswordOnEnroll()`

UnsetShowP12PasswordOnEnroll ensures that no value is present for ShowP12PasswordOnEnroll, not even an explicit nil
### GetShowP12OnEnroll

`func (o *CertificateProfileCryptoPolicy) GetShowP12OnEnroll() bool`

GetShowP12OnEnroll returns the ShowP12OnEnroll field if non-nil, zero value otherwise.

### GetShowP12OnEnrollOk

`func (o *CertificateProfileCryptoPolicy) GetShowP12OnEnrollOk() (*bool, bool)`

GetShowP12OnEnrollOk returns a tuple with the ShowP12OnEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowP12OnEnroll

`func (o *CertificateProfileCryptoPolicy) SetShowP12OnEnroll(v bool)`

SetShowP12OnEnroll sets ShowP12OnEnroll field to given value.

### HasShowP12OnEnroll

`func (o *CertificateProfileCryptoPolicy) HasShowP12OnEnroll() bool`

HasShowP12OnEnroll returns a boolean if a field has been set.

### SetShowP12OnEnrollNil

`func (o *CertificateProfileCryptoPolicy) SetShowP12OnEnrollNil(b bool)`

 SetShowP12OnEnrollNil sets the value for ShowP12OnEnroll to be an explicit nil

### UnsetShowP12OnEnroll
`func (o *CertificateProfileCryptoPolicy) UnsetShowP12OnEnroll()`

UnsetShowP12OnEnroll ensures that no value is present for ShowP12OnEnroll, not even an explicit nil
### GetShowP12PasswordOnRecover

`func (o *CertificateProfileCryptoPolicy) GetShowP12PasswordOnRecover() bool`

GetShowP12PasswordOnRecover returns the ShowP12PasswordOnRecover field if non-nil, zero value otherwise.

### GetShowP12PasswordOnRecoverOk

`func (o *CertificateProfileCryptoPolicy) GetShowP12PasswordOnRecoverOk() (*bool, bool)`

GetShowP12PasswordOnRecoverOk returns a tuple with the ShowP12PasswordOnRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowP12PasswordOnRecover

`func (o *CertificateProfileCryptoPolicy) SetShowP12PasswordOnRecover(v bool)`

SetShowP12PasswordOnRecover sets ShowP12PasswordOnRecover field to given value.

### HasShowP12PasswordOnRecover

`func (o *CertificateProfileCryptoPolicy) HasShowP12PasswordOnRecover() bool`

HasShowP12PasswordOnRecover returns a boolean if a field has been set.

### SetShowP12PasswordOnRecoverNil

`func (o *CertificateProfileCryptoPolicy) SetShowP12PasswordOnRecoverNil(b bool)`

 SetShowP12PasswordOnRecoverNil sets the value for ShowP12PasswordOnRecover to be an explicit nil

### UnsetShowP12PasswordOnRecover
`func (o *CertificateProfileCryptoPolicy) UnsetShowP12PasswordOnRecover()`

UnsetShowP12PasswordOnRecover ensures that no value is present for ShowP12PasswordOnRecover, not even an explicit nil
### GetShowP12OnRecover

`func (o *CertificateProfileCryptoPolicy) GetShowP12OnRecover() bool`

GetShowP12OnRecover returns the ShowP12OnRecover field if non-nil, zero value otherwise.

### GetShowP12OnRecoverOk

`func (o *CertificateProfileCryptoPolicy) GetShowP12OnRecoverOk() (*bool, bool)`

GetShowP12OnRecoverOk returns a tuple with the ShowP12OnRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowP12OnRecover

`func (o *CertificateProfileCryptoPolicy) SetShowP12OnRecover(v bool)`

SetShowP12OnRecover sets ShowP12OnRecover field to given value.

### HasShowP12OnRecover

`func (o *CertificateProfileCryptoPolicy) HasShowP12OnRecover() bool`

HasShowP12OnRecover returns a boolean if a field has been set.

### SetShowP12OnRecoverNil

`func (o *CertificateProfileCryptoPolicy) SetShowP12OnRecoverNil(b bool)`

 SetShowP12OnRecoverNil sets the value for ShowP12OnRecover to be an explicit nil

### UnsetShowP12OnRecover
`func (o *CertificateProfileCryptoPolicy) UnsetShowP12OnRecover()`

UnsetShowP12OnRecover ensures that no value is present for ShowP12OnRecover, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


