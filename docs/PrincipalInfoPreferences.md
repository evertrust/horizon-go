# PrincipalInfoPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateFields** | Pointer to **[]string** | The user&#39;s preferred columns on certificate view | [optional] 
**DarkMode** | Pointer to **NullableBool** | Dark Mode is enabled on UI for this user | [optional] [default to false]
**ExpertMode** | Pointer to **bool** | Expert mode is enabled on UI for this user | [optional] [default to false]
**Lang** | Pointer to **NullableString** | The preferred language of the user | [optional] 
**RequestFields** | Pointer to **[]string** | The user&#39;s preferred columns on request view | [optional] 

## Methods

### NewPrincipalInfoPreferences

`func NewPrincipalInfoPreferences() *PrincipalInfoPreferences`

NewPrincipalInfoPreferences instantiates a new PrincipalInfoPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalInfoPreferencesWithDefaults

`func NewPrincipalInfoPreferencesWithDefaults() *PrincipalInfoPreferences`

NewPrincipalInfoPreferencesWithDefaults instantiates a new PrincipalInfoPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateFields

`func (o *PrincipalInfoPreferences) GetCertificateFields() []string`

GetCertificateFields returns the CertificateFields field if non-nil, zero value otherwise.

### GetCertificateFieldsOk

`func (o *PrincipalInfoPreferences) GetCertificateFieldsOk() (*[]string, bool)`

GetCertificateFieldsOk returns a tuple with the CertificateFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFields

`func (o *PrincipalInfoPreferences) SetCertificateFields(v []string)`

SetCertificateFields sets CertificateFields field to given value.

### HasCertificateFields

`func (o *PrincipalInfoPreferences) HasCertificateFields() bool`

HasCertificateFields returns a boolean if a field has been set.

### SetCertificateFieldsNil

`func (o *PrincipalInfoPreferences) SetCertificateFieldsNil(b bool)`

 SetCertificateFieldsNil sets the value for CertificateFields to be an explicit nil

### UnsetCertificateFields
`func (o *PrincipalInfoPreferences) UnsetCertificateFields()`

UnsetCertificateFields ensures that no value is present for CertificateFields, not even an explicit nil
### GetDarkMode

`func (o *PrincipalInfoPreferences) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *PrincipalInfoPreferences) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *PrincipalInfoPreferences) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *PrincipalInfoPreferences) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### SetDarkModeNil

`func (o *PrincipalInfoPreferences) SetDarkModeNil(b bool)`

 SetDarkModeNil sets the value for DarkMode to be an explicit nil

### UnsetDarkMode
`func (o *PrincipalInfoPreferences) UnsetDarkMode()`

UnsetDarkMode ensures that no value is present for DarkMode, not even an explicit nil
### GetExpertMode

`func (o *PrincipalInfoPreferences) GetExpertMode() bool`

GetExpertMode returns the ExpertMode field if non-nil, zero value otherwise.

### GetExpertModeOk

`func (o *PrincipalInfoPreferences) GetExpertModeOk() (*bool, bool)`

GetExpertModeOk returns a tuple with the ExpertMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpertMode

`func (o *PrincipalInfoPreferences) SetExpertMode(v bool)`

SetExpertMode sets ExpertMode field to given value.

### HasExpertMode

`func (o *PrincipalInfoPreferences) HasExpertMode() bool`

HasExpertMode returns a boolean if a field has been set.

### GetLang

`func (o *PrincipalInfoPreferences) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *PrincipalInfoPreferences) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *PrincipalInfoPreferences) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *PrincipalInfoPreferences) HasLang() bool`

HasLang returns a boolean if a field has been set.

### SetLangNil

`func (o *PrincipalInfoPreferences) SetLangNil(b bool)`

 SetLangNil sets the value for Lang to be an explicit nil

### UnsetLang
`func (o *PrincipalInfoPreferences) UnsetLang()`

UnsetLang ensures that no value is present for Lang, not even an explicit nil
### GetRequestFields

`func (o *PrincipalInfoPreferences) GetRequestFields() []string`

GetRequestFields returns the RequestFields field if non-nil, zero value otherwise.

### GetRequestFieldsOk

`func (o *PrincipalInfoPreferences) GetRequestFieldsOk() (*[]string, bool)`

GetRequestFieldsOk returns a tuple with the RequestFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFields

`func (o *PrincipalInfoPreferences) SetRequestFields(v []string)`

SetRequestFields sets RequestFields field to given value.

### HasRequestFields

`func (o *PrincipalInfoPreferences) HasRequestFields() bool`

HasRequestFields returns a boolean if a field has been set.

### SetRequestFieldsNil

`func (o *PrincipalInfoPreferences) SetRequestFieldsNil(b bool)`

 SetRequestFieldsNil sets the value for RequestFields to be an explicit nil

### UnsetRequestFields
`func (o *PrincipalInfoPreferences) UnsetRequestFields()`

UnsetRequestFields ensures that no value is present for RequestFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


