# LabelElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputationRule** | Pointer to **NullableString** | The computation rule of the label element | [optional] 
**EditableByApprover** | Pointer to **NullableBool** | Whether the label element is editable by the approver | [optional] 
**EditableByRequester** | Pointer to **NullableBool** | Whether the label element is editable by the requester | [optional] 
**Enum** | Pointer to **[]string** | The whitelist used to validate the label element | [optional] 
**Label** | **string** | The name of the label | 
**Mandatory** | Pointer to **NullableBool** | Whether the label element is mandatory to submit a request | [optional] 
**Regex** | Pointer to **NullableString** | The regex used to validate the label element | [optional] 
**Suggestions** | Pointer to **[]string** | The suggestions used to recommend the label element values | [optional] 
**Value** | Pointer to **NullableString** | The default value of the label element | [optional] 

## Methods

### NewLabelElement

`func NewLabelElement(label string, ) *LabelElement`

NewLabelElement instantiates a new LabelElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelElementWithDefaults

`func NewLabelElementWithDefaults() *LabelElement`

NewLabelElementWithDefaults instantiates a new LabelElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputationRule

`func (o *LabelElement) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *LabelElement) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *LabelElement) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *LabelElement) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *LabelElement) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *LabelElement) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetEditableByApprover

`func (o *LabelElement) GetEditableByApprover() bool`

GetEditableByApprover returns the EditableByApprover field if non-nil, zero value otherwise.

### GetEditableByApproverOk

`func (o *LabelElement) GetEditableByApproverOk() (*bool, bool)`

GetEditableByApproverOk returns a tuple with the EditableByApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByApprover

`func (o *LabelElement) SetEditableByApprover(v bool)`

SetEditableByApprover sets EditableByApprover field to given value.

### HasEditableByApprover

`func (o *LabelElement) HasEditableByApprover() bool`

HasEditableByApprover returns a boolean if a field has been set.

### SetEditableByApproverNil

`func (o *LabelElement) SetEditableByApproverNil(b bool)`

 SetEditableByApproverNil sets the value for EditableByApprover to be an explicit nil

### UnsetEditableByApprover
`func (o *LabelElement) UnsetEditableByApprover()`

UnsetEditableByApprover ensures that no value is present for EditableByApprover, not even an explicit nil
### GetEditableByRequester

`func (o *LabelElement) GetEditableByRequester() bool`

GetEditableByRequester returns the EditableByRequester field if non-nil, zero value otherwise.

### GetEditableByRequesterOk

`func (o *LabelElement) GetEditableByRequesterOk() (*bool, bool)`

GetEditableByRequesterOk returns a tuple with the EditableByRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByRequester

`func (o *LabelElement) SetEditableByRequester(v bool)`

SetEditableByRequester sets EditableByRequester field to given value.

### HasEditableByRequester

`func (o *LabelElement) HasEditableByRequester() bool`

HasEditableByRequester returns a boolean if a field has been set.

### SetEditableByRequesterNil

`func (o *LabelElement) SetEditableByRequesterNil(b bool)`

 SetEditableByRequesterNil sets the value for EditableByRequester to be an explicit nil

### UnsetEditableByRequester
`func (o *LabelElement) UnsetEditableByRequester()`

UnsetEditableByRequester ensures that no value is present for EditableByRequester, not even an explicit nil
### GetEnum

`func (o *LabelElement) GetEnum() []string`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *LabelElement) GetEnumOk() (*[]string, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *LabelElement) SetEnum(v []string)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *LabelElement) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### SetEnumNil

`func (o *LabelElement) SetEnumNil(b bool)`

 SetEnumNil sets the value for Enum to be an explicit nil

### UnsetEnum
`func (o *LabelElement) UnsetEnum()`

UnsetEnum ensures that no value is present for Enum, not even an explicit nil
### GetLabel

`func (o *LabelElement) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LabelElement) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LabelElement) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMandatory

`func (o *LabelElement) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *LabelElement) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *LabelElement) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *LabelElement) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### SetMandatoryNil

`func (o *LabelElement) SetMandatoryNil(b bool)`

 SetMandatoryNil sets the value for Mandatory to be an explicit nil

### UnsetMandatory
`func (o *LabelElement) UnsetMandatory()`

UnsetMandatory ensures that no value is present for Mandatory, not even an explicit nil
### GetRegex

`func (o *LabelElement) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *LabelElement) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *LabelElement) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *LabelElement) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *LabelElement) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *LabelElement) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetSuggestions

`func (o *LabelElement) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *LabelElement) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *LabelElement) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *LabelElement) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### SetSuggestionsNil

`func (o *LabelElement) SetSuggestionsNil(b bool)`

 SetSuggestionsNil sets the value for Suggestions to be an explicit nil

### UnsetSuggestions
`func (o *LabelElement) UnsetSuggestions()`

UnsetSuggestions ensures that no value is present for Suggestions, not even an explicit nil
### GetValue

`func (o *LabelElement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LabelElement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LabelElement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *LabelElement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *LabelElement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *LabelElement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


