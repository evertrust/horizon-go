# WcceTemplateMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | **string** |  | 
**Profile** | **string** |  | 
**EnrollmentMode** | **string** |  | 
**EoboTrustedCas** | Pointer to **[]string** |  | [optional] 
**TemplateVersion** | Pointer to **string** | The version of the Microsoft template. Available from &#x60;2.8.1&#x60; | [optional] [default to "v1"]

## Methods

### NewWcceTemplateMapping

`func NewWcceTemplateMapping(template string, profile string, enrollmentMode string, ) *WcceTemplateMapping`

NewWcceTemplateMapping instantiates a new WcceTemplateMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWcceTemplateMappingWithDefaults

`func NewWcceTemplateMappingWithDefaults() *WcceTemplateMapping`

NewWcceTemplateMappingWithDefaults instantiates a new WcceTemplateMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *WcceTemplateMapping) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WcceTemplateMapping) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WcceTemplateMapping) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetProfile

`func (o *WcceTemplateMapping) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WcceTemplateMapping) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WcceTemplateMapping) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetEnrollmentMode

`func (o *WcceTemplateMapping) GetEnrollmentMode() string`

GetEnrollmentMode returns the EnrollmentMode field if non-nil, zero value otherwise.

### GetEnrollmentModeOk

`func (o *WcceTemplateMapping) GetEnrollmentModeOk() (*string, bool)`

GetEnrollmentModeOk returns a tuple with the EnrollmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMode

`func (o *WcceTemplateMapping) SetEnrollmentMode(v string)`

SetEnrollmentMode sets EnrollmentMode field to given value.


### GetEoboTrustedCas

`func (o *WcceTemplateMapping) GetEoboTrustedCas() []string`

GetEoboTrustedCas returns the EoboTrustedCas field if non-nil, zero value otherwise.

### GetEoboTrustedCasOk

`func (o *WcceTemplateMapping) GetEoboTrustedCasOk() (*[]string, bool)`

GetEoboTrustedCasOk returns a tuple with the EoboTrustedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEoboTrustedCas

`func (o *WcceTemplateMapping) SetEoboTrustedCas(v []string)`

SetEoboTrustedCas sets EoboTrustedCas field to given value.

### HasEoboTrustedCas

`func (o *WcceTemplateMapping) HasEoboTrustedCas() bool`

HasEoboTrustedCas returns a boolean if a field has been set.

### SetEoboTrustedCasNil

`func (o *WcceTemplateMapping) SetEoboTrustedCasNil(b bool)`

 SetEoboTrustedCasNil sets the value for EoboTrustedCas to be an explicit nil

### UnsetEoboTrustedCas
`func (o *WcceTemplateMapping) UnsetEoboTrustedCas()`

UnsetEoboTrustedCas ensures that no value is present for EoboTrustedCas, not even an explicit nil
### GetTemplateVersion

`func (o *WcceTemplateMapping) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *WcceTemplateMapping) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *WcceTemplateMapping) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *WcceTemplateMapping) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


