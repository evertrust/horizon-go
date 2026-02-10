# ExtensionElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**ComputationRule** | Pointer to **NullableString** | A computation rule that will dynamically generate a string value from the request&#39;s context | [optional] 
**Mandatory** | **bool** |  | 
**EditableByRequester** | Pointer to **NullableBool** |  | [optional] 
**EditableByApprover** | Pointer to **NullableBool** |  | [optional] 
**Regex** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExtensionElement

`func NewExtensionElement(type_ string, mandatory bool, ) *ExtensionElement`

NewExtensionElement instantiates a new ExtensionElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionElementWithDefaults

`func NewExtensionElementWithDefaults() *ExtensionElement`

NewExtensionElementWithDefaults instantiates a new ExtensionElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExtensionElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtensionElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtensionElement) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ExtensionElement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExtensionElement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExtensionElement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ExtensionElement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ExtensionElement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ExtensionElement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetComputationRule

`func (o *ExtensionElement) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *ExtensionElement) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *ExtensionElement) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *ExtensionElement) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *ExtensionElement) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *ExtensionElement) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetMandatory

`func (o *ExtensionElement) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *ExtensionElement) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *ExtensionElement) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetEditableByRequester

`func (o *ExtensionElement) GetEditableByRequester() bool`

GetEditableByRequester returns the EditableByRequester field if non-nil, zero value otherwise.

### GetEditableByRequesterOk

`func (o *ExtensionElement) GetEditableByRequesterOk() (*bool, bool)`

GetEditableByRequesterOk returns a tuple with the EditableByRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByRequester

`func (o *ExtensionElement) SetEditableByRequester(v bool)`

SetEditableByRequester sets EditableByRequester field to given value.

### HasEditableByRequester

`func (o *ExtensionElement) HasEditableByRequester() bool`

HasEditableByRequester returns a boolean if a field has been set.

### SetEditableByRequesterNil

`func (o *ExtensionElement) SetEditableByRequesterNil(b bool)`

 SetEditableByRequesterNil sets the value for EditableByRequester to be an explicit nil

### UnsetEditableByRequester
`func (o *ExtensionElement) UnsetEditableByRequester()`

UnsetEditableByRequester ensures that no value is present for EditableByRequester, not even an explicit nil
### GetEditableByApprover

`func (o *ExtensionElement) GetEditableByApprover() bool`

GetEditableByApprover returns the EditableByApprover field if non-nil, zero value otherwise.

### GetEditableByApproverOk

`func (o *ExtensionElement) GetEditableByApproverOk() (*bool, bool)`

GetEditableByApproverOk returns a tuple with the EditableByApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByApprover

`func (o *ExtensionElement) SetEditableByApprover(v bool)`

SetEditableByApprover sets EditableByApprover field to given value.

### HasEditableByApprover

`func (o *ExtensionElement) HasEditableByApprover() bool`

HasEditableByApprover returns a boolean if a field has been set.

### SetEditableByApproverNil

`func (o *ExtensionElement) SetEditableByApproverNil(b bool)`

 SetEditableByApproverNil sets the value for EditableByApprover to be an explicit nil

### UnsetEditableByApprover
`func (o *ExtensionElement) UnsetEditableByApprover()`

UnsetEditableByApprover ensures that no value is present for EditableByApprover, not even an explicit nil
### GetRegex

`func (o *ExtensionElement) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *ExtensionElement) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *ExtensionElement) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *ExtensionElement) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *ExtensionElement) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *ExtensionElement) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


