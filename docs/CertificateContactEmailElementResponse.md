# CertificateContactEmailElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputationRule** | Pointer to **NullableString** | Computation rule input will be evaluated and will override all other inputs | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the contact email | [optional] 
**Editable** | Pointer to **NullableBool** | Whether the contact email is editable by the requester | [optional] 
**Mandatory** | Pointer to **NullableBool** | Whether the contact email is mandatory to submit this request | [optional] 
**Regex** | Pointer to **NullableString** | The regular expression to validate the contact email | [optional] 
**Value** | Pointer to **NullableString** | The contact email | [optional] 
**Whitelist** | Pointer to **[]string** | The list of allowed contact emails | [optional] 

## Methods

### NewCertificateContactEmailElementResponse

`func NewCertificateContactEmailElementResponse() *CertificateContactEmailElementResponse`

NewCertificateContactEmailElementResponse instantiates a new CertificateContactEmailElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateContactEmailElementResponseWithDefaults

`func NewCertificateContactEmailElementResponseWithDefaults() *CertificateContactEmailElementResponse`

NewCertificateContactEmailElementResponseWithDefaults instantiates a new CertificateContactEmailElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputationRule

`func (o *CertificateContactEmailElementResponse) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *CertificateContactEmailElementResponse) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *CertificateContactEmailElementResponse) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *CertificateContactEmailElementResponse) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *CertificateContactEmailElementResponse) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *CertificateContactEmailElementResponse) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetDescription

`func (o *CertificateContactEmailElementResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateContactEmailElementResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateContactEmailElementResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateContactEmailElementResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateContactEmailElementResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateContactEmailElementResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEditable

`func (o *CertificateContactEmailElementResponse) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *CertificateContactEmailElementResponse) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *CertificateContactEmailElementResponse) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *CertificateContactEmailElementResponse) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *CertificateContactEmailElementResponse) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *CertificateContactEmailElementResponse) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetMandatory

`func (o *CertificateContactEmailElementResponse) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *CertificateContactEmailElementResponse) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *CertificateContactEmailElementResponse) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *CertificateContactEmailElementResponse) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### SetMandatoryNil

`func (o *CertificateContactEmailElementResponse) SetMandatoryNil(b bool)`

 SetMandatoryNil sets the value for Mandatory to be an explicit nil

### UnsetMandatory
`func (o *CertificateContactEmailElementResponse) UnsetMandatory()`

UnsetMandatory ensures that no value is present for Mandatory, not even an explicit nil
### GetRegex

`func (o *CertificateContactEmailElementResponse) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *CertificateContactEmailElementResponse) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *CertificateContactEmailElementResponse) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *CertificateContactEmailElementResponse) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *CertificateContactEmailElementResponse) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *CertificateContactEmailElementResponse) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetValue

`func (o *CertificateContactEmailElementResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CertificateContactEmailElementResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CertificateContactEmailElementResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CertificateContactEmailElementResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CertificateContactEmailElementResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CertificateContactEmailElementResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetWhitelist

`func (o *CertificateContactEmailElementResponse) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *CertificateContactEmailElementResponse) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *CertificateContactEmailElementResponse) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *CertificateContactEmailElementResponse) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### SetWhitelistNil

`func (o *CertificateContactEmailElementResponse) SetWhitelistNil(b bool)`

 SetWhitelistNil sets the value for Whitelist to be an explicit nil

### UnsetWhitelist
`func (o *CertificateContactEmailElementResponse) UnsetWhitelist()`

UnsetWhitelist ensures that no value is present for Whitelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


