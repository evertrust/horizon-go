# LocalizedString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lang** | **string** | The ISO 3166-1 (2-letters) code of the language used for the value | 
**Value** | **string** | The localized value | 

## Methods

### NewLocalizedString

`func NewLocalizedString(lang string, value string, ) *LocalizedString`

NewLocalizedString instantiates a new LocalizedString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalizedStringWithDefaults

`func NewLocalizedStringWithDefaults() *LocalizedString`

NewLocalizedStringWithDefaults instantiates a new LocalizedString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLang

`func (o *LocalizedString) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *LocalizedString) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *LocalizedString) SetLang(v string)`

SetLang sets Lang field to given value.


### GetValue

`func (o *LocalizedString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LocalizedString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LocalizedString) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


