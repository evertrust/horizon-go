# ContactEmailPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableString** |  | [optional] 
**ComputationRule** | Pointer to **NullableString** | A computation rule that will dynamically generate a string value from the request&#39;s context | [optional] 
**Mandatory** | **bool** |  | 
**EditableByRequester** | Pointer to **NullableBool** |  | [optional] 
**EditableByApprover** | Pointer to **NullableBool** |  | [optional] 
**Regex** | Pointer to **NullableString** |  | [optional] 
**Whitelist** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 

## Methods

### NewContactEmailPolicy

`func NewContactEmailPolicy(mandatory bool, ) *ContactEmailPolicy`

NewContactEmailPolicy instantiates a new ContactEmailPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactEmailPolicyWithDefaults

`func NewContactEmailPolicyWithDefaults() *ContactEmailPolicy`

NewContactEmailPolicyWithDefaults instantiates a new ContactEmailPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ContactEmailPolicy) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContactEmailPolicy) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContactEmailPolicy) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContactEmailPolicy) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ContactEmailPolicy) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ContactEmailPolicy) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetComputationRule

`func (o *ContactEmailPolicy) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *ContactEmailPolicy) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *ContactEmailPolicy) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *ContactEmailPolicy) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *ContactEmailPolicy) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *ContactEmailPolicy) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetMandatory

`func (o *ContactEmailPolicy) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *ContactEmailPolicy) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *ContactEmailPolicy) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetEditableByRequester

`func (o *ContactEmailPolicy) GetEditableByRequester() bool`

GetEditableByRequester returns the EditableByRequester field if non-nil, zero value otherwise.

### GetEditableByRequesterOk

`func (o *ContactEmailPolicy) GetEditableByRequesterOk() (*bool, bool)`

GetEditableByRequesterOk returns a tuple with the EditableByRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByRequester

`func (o *ContactEmailPolicy) SetEditableByRequester(v bool)`

SetEditableByRequester sets EditableByRequester field to given value.

### HasEditableByRequester

`func (o *ContactEmailPolicy) HasEditableByRequester() bool`

HasEditableByRequester returns a boolean if a field has been set.

### SetEditableByRequesterNil

`func (o *ContactEmailPolicy) SetEditableByRequesterNil(b bool)`

 SetEditableByRequesterNil sets the value for EditableByRequester to be an explicit nil

### UnsetEditableByRequester
`func (o *ContactEmailPolicy) UnsetEditableByRequester()`

UnsetEditableByRequester ensures that no value is present for EditableByRequester, not even an explicit nil
### GetEditableByApprover

`func (o *ContactEmailPolicy) GetEditableByApprover() bool`

GetEditableByApprover returns the EditableByApprover field if non-nil, zero value otherwise.

### GetEditableByApproverOk

`func (o *ContactEmailPolicy) GetEditableByApproverOk() (*bool, bool)`

GetEditableByApproverOk returns a tuple with the EditableByApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByApprover

`func (o *ContactEmailPolicy) SetEditableByApprover(v bool)`

SetEditableByApprover sets EditableByApprover field to given value.

### HasEditableByApprover

`func (o *ContactEmailPolicy) HasEditableByApprover() bool`

HasEditableByApprover returns a boolean if a field has been set.

### SetEditableByApproverNil

`func (o *ContactEmailPolicy) SetEditableByApproverNil(b bool)`

 SetEditableByApproverNil sets the value for EditableByApprover to be an explicit nil

### UnsetEditableByApprover
`func (o *ContactEmailPolicy) UnsetEditableByApprover()`

UnsetEditableByApprover ensures that no value is present for EditableByApprover, not even an explicit nil
### GetRegex

`func (o *ContactEmailPolicy) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *ContactEmailPolicy) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *ContactEmailPolicy) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *ContactEmailPolicy) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *ContactEmailPolicy) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *ContactEmailPolicy) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetWhitelist

`func (o *ContactEmailPolicy) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *ContactEmailPolicy) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *ContactEmailPolicy) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *ContactEmailPolicy) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### SetWhitelistNil

`func (o *ContactEmailPolicy) SetWhitelistNil(b bool)`

 SetWhitelistNil sets the value for Whitelist to be an explicit nil

### UnsetWhitelist
`func (o *ContactEmailPolicy) UnsetWhitelist()`

UnsetWhitelist ensures that no value is present for Whitelist, not even an explicit nil
### GetDescription

`func (o *ContactEmailPolicy) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContactEmailPolicy) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContactEmailPolicy) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContactEmailPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContactEmailPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContactEmailPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


