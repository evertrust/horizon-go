# MonitoredCertificateProfileCryptoPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizedKeyTypes** | Pointer to **[]string** | List of authorized key types for enrollment | [optional] 
**Escrow** | Pointer to **NullableBool** | Whether this profile will escrow the certificate private keys | [optional] [default to false]
**P12passwordPolicy** | Pointer to **NullableString** | Password policy for the P12 file | [optional] 
**P12passwordMode** | Pointer to **NullableString** | Whether the user will be required to input their PKCS#12 password upon enrollment | [optional] 
**P12storeEncryptionType** | Pointer to **NullableString** | Encryption type for the P12 file | [optional] 
**ShowP12PasswordOnRecover** | Pointer to **NullableBool** | Whether the PKCS#12 password will be displayed to the user upon recovery | [optional] 
**ShowP12OnRecover** | Pointer to **NullableBool** | Whether the PKCS#12 file will be displayed to the user upon recovery | [optional] 
**KeyAvailability** | Pointer to **NullableString** | Availability of the key in the requests (enroll, recover), as well as time during which a non-escrowed key is available for trigger retries | [optional] 

## Methods

### NewMonitoredCertificateProfileCryptoPolicy

`func NewMonitoredCertificateProfileCryptoPolicy() *MonitoredCertificateProfileCryptoPolicy`

NewMonitoredCertificateProfileCryptoPolicy instantiates a new MonitoredCertificateProfileCryptoPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoredCertificateProfileCryptoPolicyWithDefaults

`func NewMonitoredCertificateProfileCryptoPolicyWithDefaults() *MonitoredCertificateProfileCryptoPolicy`

NewMonitoredCertificateProfileCryptoPolicyWithDefaults instantiates a new MonitoredCertificateProfileCryptoPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedKeyTypes

`func (o *MonitoredCertificateProfileCryptoPolicy) GetAuthorizedKeyTypes() []string`

GetAuthorizedKeyTypes returns the AuthorizedKeyTypes field if non-nil, zero value otherwise.

### GetAuthorizedKeyTypesOk

`func (o *MonitoredCertificateProfileCryptoPolicy) GetAuthorizedKeyTypesOk() (*[]string, bool)`

GetAuthorizedKeyTypesOk returns a tuple with the AuthorizedKeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeyTypes

`func (o *MonitoredCertificateProfileCryptoPolicy) SetAuthorizedKeyTypes(v []string)`

SetAuthorizedKeyTypes sets AuthorizedKeyTypes field to given value.

### HasAuthorizedKeyTypes

`func (o *MonitoredCertificateProfileCryptoPolicy) HasAuthorizedKeyTypes() bool`

HasAuthorizedKeyTypes returns a boolean if a field has been set.

### SetAuthorizedKeyTypesNil

`func (o *MonitoredCertificateProfileCryptoPolicy) SetAuthorizedKeyTypesNil(b bool)`

 SetAuthorizedKeyTypesNil sets the value for AuthorizedKeyTypes to be an explicit nil

### UnsetAuthorizedKeyTypes
`func (o *MonitoredCertificateProfileCryptoPolicy) UnsetAuthorizedKeyTypes()`

UnsetAuthorizedKeyTypes ensures that no value is present for AuthorizedKeyTypes, not even an explicit nil
### GetEscrow

`func (o *MonitoredCertificateProfileCryptoPolicy) GetEscrow() bool`

GetEscrow returns the Escrow field if non-nil, zero value otherwise.

### GetEscrowOk

`func (o *MonitoredCertificateProfileCryptoPolicy) GetEscrowOk() (*bool, bool)`

GetEscrowOk returns a tuple with the Escrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscrow

`func (o *MonitoredCertificateProfileCryptoPolicy) SetEscrow(v bool)`

SetEscrow sets Escrow field to given value.

### HasEscrow

`func (o *MonitoredCertificateProfileCryptoPolicy) HasEscrow() bool`

HasEscrow returns a boolean if a field has been set.

### SetEscrowNil

`func (o *MonitoredCertificateProfileCryptoPolicy) SetEscrowNil(b bool)`

 SetEscrowNil sets the value for Escrow to be an explicit nil

### UnsetEscrow
`func (o *MonitoredCertificateProfileCryptoPolicy) UnsetEscrow()`

UnsetEscrow ensures that no value is present for Escrow, not even an explicit nil
### GetP12passwordPolicy

`func (o *MonitoredCertificateProfileCryptoPolicy) GetP12passwordPolicy() string`

GetP12passwordPolicy returns the P12passwordPolicy field if non-nil, zero value otherwise.

### GetP12passwordPolicyOk

`func (o *MonitoredCertificateProfileCryptoPolicy) GetP12passwordPolicyOk() (*string, bool)`

GetP12passwordPolicyOk returns a tuple with the P12passwordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP12passwordPolicy

`func (o *MonitoredCertificateProfileCryptoPolicy) SetP12passwordPolicy(v string)`

SetP12passwordPolicy sets P12passwordPolicy field to given value.

### HasP12passwordPolicy

`func (o *MonitoredCertificateProfileCryptoPolicy) HasP12passwordPolicy() bool`

HasP12passwordPolicy returns a boolean if a field has been set.

### SetP12passwordPolicyNil

`func (o *MonitoredCertificateProfileCryptoPolicy) SetP12passwordPolicyNil(b bool)`

 SetP12passwordPolicyNil sets the value for P12passwordPolicy to be an explicit nil

### UnsetP12passwordPolicy
`func (o *MonitoredCertificateProfileCryptoPolicy) UnsetP12passwordPolicy()`

UnsetP12passwordPolicy ensures that no value is present for P12passwordPolicy, not even an explicit nil
### GetP12passwordMode

`func (o *MonitoredCertificateProfileCryptoPolicy) GetP12passwordMode() string`

GetP12passwordMode returns the P12passwordMode field if non-nil, zero value otherwise.

### GetP12passwordModeOk

`func (o *MonitoredCertificateProfileCryptoPolicy) GetP12passwordModeOk() (*string, bool)`

GetP12passwordModeOk returns a tuple with the P12passwordMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP12passwordMode

`func (o *MonitoredCertificateProfileCryptoPolicy) SetP12passwordMode(v string)`

SetP12passwordMode sets P12passwordMode field to given value.

### HasP12passwordMode

`func (o *MonitoredCertificateProfileCryptoPolicy) HasP12passwordMode() bool`

HasP12passwordMode returns a boolean if a field has been set.

### SetP12passwordModeNil

`func (o *MonitoredCertificateProfileCryptoPolicy) SetP12passwordModeNil(b bool)`

 SetP12passwordModeNil sets the value for P12passwordMode to be an explicit nil

### UnsetP12passwordMode
`func (o *MonitoredCertificateProfileCryptoPolicy) UnsetP12passwordMode()`

UnsetP12passwordMode ensures that no value is present for P12passwordMode, not even an explicit nil
### GetP12storeEncryptionType

`func (o *MonitoredCertificateProfileCryptoPolicy) GetP12storeEncryptionType() string`

GetP12storeEncryptionType returns the P12storeEncryptionType field if non-nil, zero value otherwise.

### GetP12storeEncryptionTypeOk

`func (o *MonitoredCertificateProfileCryptoPolicy) GetP12storeEncryptionTypeOk() (*string, bool)`

GetP12storeEncryptionTypeOk returns a tuple with the P12storeEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP12storeEncryptionType

`func (o *MonitoredCertificateProfileCryptoPolicy) SetP12storeEncryptionType(v string)`

SetP12storeEncryptionType sets P12storeEncryptionType field to given value.

### HasP12storeEncryptionType

`func (o *MonitoredCertificateProfileCryptoPolicy) HasP12storeEncryptionType() bool`

HasP12storeEncryptionType returns a boolean if a field has been set.

### SetP12storeEncryptionTypeNil

`func (o *MonitoredCertificateProfileCryptoPolicy) SetP12storeEncryptionTypeNil(b bool)`

 SetP12storeEncryptionTypeNil sets the value for P12storeEncryptionType to be an explicit nil

### UnsetP12storeEncryptionType
`func (o *MonitoredCertificateProfileCryptoPolicy) UnsetP12storeEncryptionType()`

UnsetP12storeEncryptionType ensures that no value is present for P12storeEncryptionType, not even an explicit nil
### GetShowP12PasswordOnRecover

`func (o *MonitoredCertificateProfileCryptoPolicy) GetShowP12PasswordOnRecover() bool`

GetShowP12PasswordOnRecover returns the ShowP12PasswordOnRecover field if non-nil, zero value otherwise.

### GetShowP12PasswordOnRecoverOk

`func (o *MonitoredCertificateProfileCryptoPolicy) GetShowP12PasswordOnRecoverOk() (*bool, bool)`

GetShowP12PasswordOnRecoverOk returns a tuple with the ShowP12PasswordOnRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowP12PasswordOnRecover

`func (o *MonitoredCertificateProfileCryptoPolicy) SetShowP12PasswordOnRecover(v bool)`

SetShowP12PasswordOnRecover sets ShowP12PasswordOnRecover field to given value.

### HasShowP12PasswordOnRecover

`func (o *MonitoredCertificateProfileCryptoPolicy) HasShowP12PasswordOnRecover() bool`

HasShowP12PasswordOnRecover returns a boolean if a field has been set.

### SetShowP12PasswordOnRecoverNil

`func (o *MonitoredCertificateProfileCryptoPolicy) SetShowP12PasswordOnRecoverNil(b bool)`

 SetShowP12PasswordOnRecoverNil sets the value for ShowP12PasswordOnRecover to be an explicit nil

### UnsetShowP12PasswordOnRecover
`func (o *MonitoredCertificateProfileCryptoPolicy) UnsetShowP12PasswordOnRecover()`

UnsetShowP12PasswordOnRecover ensures that no value is present for ShowP12PasswordOnRecover, not even an explicit nil
### GetShowP12OnRecover

`func (o *MonitoredCertificateProfileCryptoPolicy) GetShowP12OnRecover() bool`

GetShowP12OnRecover returns the ShowP12OnRecover field if non-nil, zero value otherwise.

### GetShowP12OnRecoverOk

`func (o *MonitoredCertificateProfileCryptoPolicy) GetShowP12OnRecoverOk() (*bool, bool)`

GetShowP12OnRecoverOk returns a tuple with the ShowP12OnRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowP12OnRecover

`func (o *MonitoredCertificateProfileCryptoPolicy) SetShowP12OnRecover(v bool)`

SetShowP12OnRecover sets ShowP12OnRecover field to given value.

### HasShowP12OnRecover

`func (o *MonitoredCertificateProfileCryptoPolicy) HasShowP12OnRecover() bool`

HasShowP12OnRecover returns a boolean if a field has been set.

### SetShowP12OnRecoverNil

`func (o *MonitoredCertificateProfileCryptoPolicy) SetShowP12OnRecoverNil(b bool)`

 SetShowP12OnRecoverNil sets the value for ShowP12OnRecover to be an explicit nil

### UnsetShowP12OnRecover
`func (o *MonitoredCertificateProfileCryptoPolicy) UnsetShowP12OnRecover()`

UnsetShowP12OnRecover ensures that no value is present for ShowP12OnRecover, not even an explicit nil
### GetKeyAvailability

`func (o *MonitoredCertificateProfileCryptoPolicy) GetKeyAvailability() string`

GetKeyAvailability returns the KeyAvailability field if non-nil, zero value otherwise.

### GetKeyAvailabilityOk

`func (o *MonitoredCertificateProfileCryptoPolicy) GetKeyAvailabilityOk() (*string, bool)`

GetKeyAvailabilityOk returns a tuple with the KeyAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAvailability

`func (o *MonitoredCertificateProfileCryptoPolicy) SetKeyAvailability(v string)`

SetKeyAvailability sets KeyAvailability field to given value.

### HasKeyAvailability

`func (o *MonitoredCertificateProfileCryptoPolicy) HasKeyAvailability() bool`

HasKeyAvailability returns a boolean if a field has been set.

### SetKeyAvailabilityNil

`func (o *MonitoredCertificateProfileCryptoPolicy) SetKeyAvailabilityNil(b bool)`

 SetKeyAvailabilityNil sets the value for KeyAvailability to be an explicit nil

### UnsetKeyAvailability
`func (o *MonitoredCertificateProfileCryptoPolicy) UnsetKeyAvailability()`

UnsetKeyAvailability ensures that no value is present for KeyAvailability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


