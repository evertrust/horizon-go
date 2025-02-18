# CertificateProfileCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Centralized** | **bool** | Centralized enrollment is enabled on this profile | 
**Decentralized** | **bool** | Decentralized enrollment is enabled on this profile | 
**DefaultKeyType** | Pointer to **NullableString** | Key type used when no keyType has been chosen | [optional] 
**AuthorizedKeyTypes** | Pointer to **[]string** | The list of key types that are authorized for enrollment. A null value means all keys are allowed | [optional] 
**PreferredEnrollmentMode** | Pointer to **NullableString** | The enrollment mode that should be prioritized when both are defined | [optional] 
**PasswordPolicy** | Pointer to [**NullablePasswordPolicy**](PasswordPolicy.md) | The selected password policy for this profile. If none is defined and the password mode is &#x60;manual&#x60;, there is no constraint on the password. In &#x60;random&#x60; mode, the &#x60;Horizon-Default&#x60; policy is used | [optional] 
**P12passwordMode** | Pointer to **NullableString** | A &#x60;manual&#x60; password mode means the password for the PKCS#12 must be set in the request. A &#x60;random&#x60; password will be generated on Horizon | [optional] 

## Methods

### NewCertificateProfileCapabilities

`func NewCertificateProfileCapabilities(centralized bool, decentralized bool, ) *CertificateProfileCapabilities`

NewCertificateProfileCapabilities instantiates a new CertificateProfileCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileCapabilitiesWithDefaults

`func NewCertificateProfileCapabilitiesWithDefaults() *CertificateProfileCapabilities`

NewCertificateProfileCapabilitiesWithDefaults instantiates a new CertificateProfileCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCentralized

`func (o *CertificateProfileCapabilities) GetCentralized() bool`

GetCentralized returns the Centralized field if non-nil, zero value otherwise.

### GetCentralizedOk

`func (o *CertificateProfileCapabilities) GetCentralizedOk() (*bool, bool)`

GetCentralizedOk returns a tuple with the Centralized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentralized

`func (o *CertificateProfileCapabilities) SetCentralized(v bool)`

SetCentralized sets Centralized field to given value.


### GetDecentralized

`func (o *CertificateProfileCapabilities) GetDecentralized() bool`

GetDecentralized returns the Decentralized field if non-nil, zero value otherwise.

### GetDecentralizedOk

`func (o *CertificateProfileCapabilities) GetDecentralizedOk() (*bool, bool)`

GetDecentralizedOk returns a tuple with the Decentralized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecentralized

`func (o *CertificateProfileCapabilities) SetDecentralized(v bool)`

SetDecentralized sets Decentralized field to given value.


### GetDefaultKeyType

`func (o *CertificateProfileCapabilities) GetDefaultKeyType() string`

GetDefaultKeyType returns the DefaultKeyType field if non-nil, zero value otherwise.

### GetDefaultKeyTypeOk

`func (o *CertificateProfileCapabilities) GetDefaultKeyTypeOk() (*string, bool)`

GetDefaultKeyTypeOk returns a tuple with the DefaultKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeyType

`func (o *CertificateProfileCapabilities) SetDefaultKeyType(v string)`

SetDefaultKeyType sets DefaultKeyType field to given value.

### HasDefaultKeyType

`func (o *CertificateProfileCapabilities) HasDefaultKeyType() bool`

HasDefaultKeyType returns a boolean if a field has been set.

### SetDefaultKeyTypeNil

`func (o *CertificateProfileCapabilities) SetDefaultKeyTypeNil(b bool)`

 SetDefaultKeyTypeNil sets the value for DefaultKeyType to be an explicit nil

### UnsetDefaultKeyType
`func (o *CertificateProfileCapabilities) UnsetDefaultKeyType()`

UnsetDefaultKeyType ensures that no value is present for DefaultKeyType, not even an explicit nil
### GetAuthorizedKeyTypes

`func (o *CertificateProfileCapabilities) GetAuthorizedKeyTypes() []string`

GetAuthorizedKeyTypes returns the AuthorizedKeyTypes field if non-nil, zero value otherwise.

### GetAuthorizedKeyTypesOk

`func (o *CertificateProfileCapabilities) GetAuthorizedKeyTypesOk() (*[]string, bool)`

GetAuthorizedKeyTypesOk returns a tuple with the AuthorizedKeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeyTypes

`func (o *CertificateProfileCapabilities) SetAuthorizedKeyTypes(v []string)`

SetAuthorizedKeyTypes sets AuthorizedKeyTypes field to given value.

### HasAuthorizedKeyTypes

`func (o *CertificateProfileCapabilities) HasAuthorizedKeyTypes() bool`

HasAuthorizedKeyTypes returns a boolean if a field has been set.

### SetAuthorizedKeyTypesNil

`func (o *CertificateProfileCapabilities) SetAuthorizedKeyTypesNil(b bool)`

 SetAuthorizedKeyTypesNil sets the value for AuthorizedKeyTypes to be an explicit nil

### UnsetAuthorizedKeyTypes
`func (o *CertificateProfileCapabilities) UnsetAuthorizedKeyTypes()`

UnsetAuthorizedKeyTypes ensures that no value is present for AuthorizedKeyTypes, not even an explicit nil
### GetPreferredEnrollmentMode

`func (o *CertificateProfileCapabilities) GetPreferredEnrollmentMode() string`

GetPreferredEnrollmentMode returns the PreferredEnrollmentMode field if non-nil, zero value otherwise.

### GetPreferredEnrollmentModeOk

`func (o *CertificateProfileCapabilities) GetPreferredEnrollmentModeOk() (*string, bool)`

GetPreferredEnrollmentModeOk returns a tuple with the PreferredEnrollmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredEnrollmentMode

`func (o *CertificateProfileCapabilities) SetPreferredEnrollmentMode(v string)`

SetPreferredEnrollmentMode sets PreferredEnrollmentMode field to given value.

### HasPreferredEnrollmentMode

`func (o *CertificateProfileCapabilities) HasPreferredEnrollmentMode() bool`

HasPreferredEnrollmentMode returns a boolean if a field has been set.

### SetPreferredEnrollmentModeNil

`func (o *CertificateProfileCapabilities) SetPreferredEnrollmentModeNil(b bool)`

 SetPreferredEnrollmentModeNil sets the value for PreferredEnrollmentMode to be an explicit nil

### UnsetPreferredEnrollmentMode
`func (o *CertificateProfileCapabilities) UnsetPreferredEnrollmentMode()`

UnsetPreferredEnrollmentMode ensures that no value is present for PreferredEnrollmentMode, not even an explicit nil
### GetPasswordPolicy

`func (o *CertificateProfileCapabilities) GetPasswordPolicy() PasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *CertificateProfileCapabilities) GetPasswordPolicyOk() (*PasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *CertificateProfileCapabilities) SetPasswordPolicy(v PasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *CertificateProfileCapabilities) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *CertificateProfileCapabilities) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *CertificateProfileCapabilities) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetP12passwordMode

`func (o *CertificateProfileCapabilities) GetP12passwordMode() string`

GetP12passwordMode returns the P12passwordMode field if non-nil, zero value otherwise.

### GetP12passwordModeOk

`func (o *CertificateProfileCapabilities) GetP12passwordModeOk() (*string, bool)`

GetP12passwordModeOk returns a tuple with the P12passwordMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP12passwordMode

`func (o *CertificateProfileCapabilities) SetP12passwordMode(v string)`

SetP12passwordMode sets P12passwordMode field to given value.

### HasP12passwordMode

`func (o *CertificateProfileCapabilities) HasP12passwordMode() bool`

HasP12passwordMode returns a boolean if a field has been set.

### SetP12passwordModeNil

`func (o *CertificateProfileCapabilities) SetP12passwordModeNil(b bool)`

 SetP12passwordModeNil sets the value for P12passwordMode to be an explicit nil

### UnsetP12passwordMode
`func (o *CertificateProfileCapabilities) UnsetP12passwordMode()`

UnsetP12passwordMode ensures that no value is present for P12passwordMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


