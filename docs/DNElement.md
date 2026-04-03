# DNElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputationRule** | Pointer to **NullableString** | A computation rule that will dynamically generate a string value from the request&#39;s context | [optional] 
**EditableByApprover** | Pointer to **NullableBool** |  | [optional] 
**EditableByRequester** | Pointer to **NullableBool** |  | [optional] 
**Mandatory** | **bool** |  | 
**Regex** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDNElement

`func NewDNElement(mandatory bool, type_ string, ) *DNElement`

NewDNElement instantiates a new DNElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNElementWithDefaults

`func NewDNElementWithDefaults() *DNElement`

NewDNElementWithDefaults instantiates a new DNElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputationRule

`func (o *DNElement) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *DNElement) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *DNElement) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *DNElement) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *DNElement) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *DNElement) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetEditableByApprover

`func (o *DNElement) GetEditableByApprover() bool`

GetEditableByApprover returns the EditableByApprover field if non-nil, zero value otherwise.

### GetEditableByApproverOk

`func (o *DNElement) GetEditableByApproverOk() (*bool, bool)`

GetEditableByApproverOk returns a tuple with the EditableByApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByApprover

`func (o *DNElement) SetEditableByApprover(v bool)`

SetEditableByApprover sets EditableByApprover field to given value.

### HasEditableByApprover

`func (o *DNElement) HasEditableByApprover() bool`

HasEditableByApprover returns a boolean if a field has been set.

### SetEditableByApproverNil

`func (o *DNElement) SetEditableByApproverNil(b bool)`

 SetEditableByApproverNil sets the value for EditableByApprover to be an explicit nil

### UnsetEditableByApprover
`func (o *DNElement) UnsetEditableByApprover()`

UnsetEditableByApprover ensures that no value is present for EditableByApprover, not even an explicit nil
### GetEditableByRequester

`func (o *DNElement) GetEditableByRequester() bool`

GetEditableByRequester returns the EditableByRequester field if non-nil, zero value otherwise.

### GetEditableByRequesterOk

`func (o *DNElement) GetEditableByRequesterOk() (*bool, bool)`

GetEditableByRequesterOk returns a tuple with the EditableByRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByRequester

`func (o *DNElement) SetEditableByRequester(v bool)`

SetEditableByRequester sets EditableByRequester field to given value.

### HasEditableByRequester

`func (o *DNElement) HasEditableByRequester() bool`

HasEditableByRequester returns a boolean if a field has been set.

### SetEditableByRequesterNil

`func (o *DNElement) SetEditableByRequesterNil(b bool)`

 SetEditableByRequesterNil sets the value for EditableByRequester to be an explicit nil

### UnsetEditableByRequester
`func (o *DNElement) UnsetEditableByRequester()`

UnsetEditableByRequester ensures that no value is present for EditableByRequester, not even an explicit nil
### GetMandatory

`func (o *DNElement) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *DNElement) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *DNElement) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetRegex

`func (o *DNElement) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *DNElement) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *DNElement) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *DNElement) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *DNElement) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *DNElement) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetType

`func (o *DNElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNElement) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *DNElement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DNElement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DNElement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DNElement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *DNElement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *DNElement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


