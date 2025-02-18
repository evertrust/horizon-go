# TeamPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditableByRequester** | **bool** |  | 
**EditableByApprover** | **bool** |  | 
**Regex** | Pointer to **NullableString** |  | [optional] 
**Whitelist** | Pointer to **[]string** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**ComputationRule** | Pointer to **NullableString** | A computation rule that will dynamically generate a string value from the request&#39;s context | [optional] 
**Mandatory** | **bool** |  | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 

## Methods

### NewTeamPolicy

`func NewTeamPolicy(editableByRequester bool, editableByApprover bool, mandatory bool, ) *TeamPolicy`

NewTeamPolicy instantiates a new TeamPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPolicyWithDefaults

`func NewTeamPolicyWithDefaults() *TeamPolicy`

NewTeamPolicyWithDefaults instantiates a new TeamPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditableByRequester

`func (o *TeamPolicy) GetEditableByRequester() bool`

GetEditableByRequester returns the EditableByRequester field if non-nil, zero value otherwise.

### GetEditableByRequesterOk

`func (o *TeamPolicy) GetEditableByRequesterOk() (*bool, bool)`

GetEditableByRequesterOk returns a tuple with the EditableByRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByRequester

`func (o *TeamPolicy) SetEditableByRequester(v bool)`

SetEditableByRequester sets EditableByRequester field to given value.


### GetEditableByApprover

`func (o *TeamPolicy) GetEditableByApprover() bool`

GetEditableByApprover returns the EditableByApprover field if non-nil, zero value otherwise.

### GetEditableByApproverOk

`func (o *TeamPolicy) GetEditableByApproverOk() (*bool, bool)`

GetEditableByApproverOk returns a tuple with the EditableByApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByApprover

`func (o *TeamPolicy) SetEditableByApprover(v bool)`

SetEditableByApprover sets EditableByApprover field to given value.


### GetRegex

`func (o *TeamPolicy) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *TeamPolicy) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *TeamPolicy) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *TeamPolicy) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *TeamPolicy) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *TeamPolicy) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetWhitelist

`func (o *TeamPolicy) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *TeamPolicy) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *TeamPolicy) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *TeamPolicy) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### SetWhitelistNil

`func (o *TeamPolicy) SetWhitelistNil(b bool)`

 SetWhitelistNil sets the value for Whitelist to be an explicit nil

### UnsetWhitelist
`func (o *TeamPolicy) UnsetWhitelist()`

UnsetWhitelist ensures that no value is present for Whitelist, not even an explicit nil
### GetValue

`func (o *TeamPolicy) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TeamPolicy) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TeamPolicy) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TeamPolicy) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TeamPolicy) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TeamPolicy) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetComputationRule

`func (o *TeamPolicy) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *TeamPolicy) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *TeamPolicy) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *TeamPolicy) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *TeamPolicy) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *TeamPolicy) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetMandatory

`func (o *TeamPolicy) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *TeamPolicy) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *TeamPolicy) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetDescription

`func (o *TeamPolicy) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamPolicy) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamPolicy) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TeamPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TeamPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


