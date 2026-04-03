# RequestLabelElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputationRule** | Pointer to **NullableString** | The computation rule of the label element | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the label element | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the label element | [optional] 
**Editable** | Pointer to **NullableBool** | Whether the label element is editable | [optional] 
**Enum** | Pointer to **[]string** | The enum used to validate the label element | [optional] 
**Label** | **string** | The name of the label | 
**Mandatory** | Pointer to **NullableBool** | Whether the label element is mandatory to submit this request | [optional] 
**Regex** | Pointer to **NullableString** | The regex used to validate the label element | [optional] 
**Suggestions** | Pointer to **[]string** | The suggestions used to recommend the label element values | [optional] 
**Value** | Pointer to **NullableString** | The value of the label element | [optional] 

## Methods

### NewRequestLabelElementResponse

`func NewRequestLabelElementResponse(label string, ) *RequestLabelElementResponse`

NewRequestLabelElementResponse instantiates a new RequestLabelElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestLabelElementResponseWithDefaults

`func NewRequestLabelElementResponseWithDefaults() *RequestLabelElementResponse`

NewRequestLabelElementResponseWithDefaults instantiates a new RequestLabelElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputationRule

`func (o *RequestLabelElementResponse) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *RequestLabelElementResponse) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *RequestLabelElementResponse) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *RequestLabelElementResponse) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *RequestLabelElementResponse) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *RequestLabelElementResponse) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetDescription

`func (o *RequestLabelElementResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestLabelElementResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestLabelElementResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestLabelElementResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RequestLabelElementResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RequestLabelElementResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *RequestLabelElementResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RequestLabelElementResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RequestLabelElementResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RequestLabelElementResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *RequestLabelElementResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *RequestLabelElementResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEditable

`func (o *RequestLabelElementResponse) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *RequestLabelElementResponse) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *RequestLabelElementResponse) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *RequestLabelElementResponse) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *RequestLabelElementResponse) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *RequestLabelElementResponse) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetEnum

`func (o *RequestLabelElementResponse) GetEnum() []string`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *RequestLabelElementResponse) GetEnumOk() (*[]string, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *RequestLabelElementResponse) SetEnum(v []string)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *RequestLabelElementResponse) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### SetEnumNil

`func (o *RequestLabelElementResponse) SetEnumNil(b bool)`

 SetEnumNil sets the value for Enum to be an explicit nil

### UnsetEnum
`func (o *RequestLabelElementResponse) UnsetEnum()`

UnsetEnum ensures that no value is present for Enum, not even an explicit nil
### GetLabel

`func (o *RequestLabelElementResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RequestLabelElementResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RequestLabelElementResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMandatory

`func (o *RequestLabelElementResponse) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *RequestLabelElementResponse) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *RequestLabelElementResponse) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *RequestLabelElementResponse) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### SetMandatoryNil

`func (o *RequestLabelElementResponse) SetMandatoryNil(b bool)`

 SetMandatoryNil sets the value for Mandatory to be an explicit nil

### UnsetMandatory
`func (o *RequestLabelElementResponse) UnsetMandatory()`

UnsetMandatory ensures that no value is present for Mandatory, not even an explicit nil
### GetRegex

`func (o *RequestLabelElementResponse) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *RequestLabelElementResponse) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *RequestLabelElementResponse) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *RequestLabelElementResponse) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *RequestLabelElementResponse) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *RequestLabelElementResponse) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetSuggestions

`func (o *RequestLabelElementResponse) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *RequestLabelElementResponse) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *RequestLabelElementResponse) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *RequestLabelElementResponse) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### SetSuggestionsNil

`func (o *RequestLabelElementResponse) SetSuggestionsNil(b bool)`

 SetSuggestionsNil sets the value for Suggestions to be an explicit nil

### UnsetSuggestions
`func (o *RequestLabelElementResponse) UnsetSuggestions()`

UnsetSuggestions ensures that no value is present for Suggestions, not even an explicit nil
### GetValue

`func (o *RequestLabelElementResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RequestLabelElementResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RequestLabelElementResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RequestLabelElementResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *RequestLabelElementResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *RequestLabelElementResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


