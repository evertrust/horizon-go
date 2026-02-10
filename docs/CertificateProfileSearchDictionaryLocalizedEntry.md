# CertificateProfileSearchDictionaryLocalizedEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The technical name of the profile | 
**Module** | **string** | The module on which the profile belongs | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the profile | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized description of the profile | [optional] 

## Methods

### NewCertificateProfileSearchDictionaryLocalizedEntry

`func NewCertificateProfileSearchDictionaryLocalizedEntry(name string, module string, ) *CertificateProfileSearchDictionaryLocalizedEntry`

NewCertificateProfileSearchDictionaryLocalizedEntry instantiates a new CertificateProfileSearchDictionaryLocalizedEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileSearchDictionaryLocalizedEntryWithDefaults

`func NewCertificateProfileSearchDictionaryLocalizedEntryWithDefaults() *CertificateProfileSearchDictionaryLocalizedEntry`

NewCertificateProfileSearchDictionaryLocalizedEntryWithDefaults instantiates a new CertificateProfileSearchDictionaryLocalizedEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) SetName(v string)`

SetName sets Name field to given value.


### GetModule

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) SetModule(v string)`

SetModule sets Module field to given value.


### GetDisplayName

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CertificateProfileSearchDictionaryLocalizedEntry) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateProfileSearchDictionaryLocalizedEntry) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateProfileSearchDictionaryLocalizedEntry) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


