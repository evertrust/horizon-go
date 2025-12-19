# CertificateLabelSearchDictionaryLocalizedEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized description of the label | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the label | [optional] 
**Name** | **string** | The technical name of the label | 

## Methods

### NewCertificateLabelSearchDictionaryLocalizedEntry

`func NewCertificateLabelSearchDictionaryLocalizedEntry(name string, ) *CertificateLabelSearchDictionaryLocalizedEntry`

NewCertificateLabelSearchDictionaryLocalizedEntry instantiates a new CertificateLabelSearchDictionaryLocalizedEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateLabelSearchDictionaryLocalizedEntryWithDefaults

`func NewCertificateLabelSearchDictionaryLocalizedEntryWithDefaults() *CertificateLabelSearchDictionaryLocalizedEntry`

NewCertificateLabelSearchDictionaryLocalizedEntryWithDefaults instantiates a new CertificateLabelSearchDictionaryLocalizedEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateLabelSearchDictionaryLocalizedEntry) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CertificateLabelSearchDictionaryLocalizedEntry) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetName

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateLabelSearchDictionaryLocalizedEntry) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


