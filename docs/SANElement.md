# SANElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ComputationRule** | Pointer to **NullableString** | A computation rule that will dynamically generate a string value from the request&#39;s context | [optional] 
**EditableByRequester** | Pointer to **NullableBool** |  | [optional] 
**EditableByApprover** | Pointer to **NullableBool** |  | [optional] 
**Regex** | Pointer to **NullableString** |  | [optional] 
**Min** | Pointer to **NullableInt64** |  | [optional] 
**Max** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewSANElement

`func NewSANElement(type_ string, ) *SANElement`

NewSANElement instantiates a new SANElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSANElementWithDefaults

`func NewSANElementWithDefaults() *SANElement`

NewSANElementWithDefaults instantiates a new SANElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SANElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SANElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SANElement) SetType(v string)`

SetType sets Type field to given value.


### GetComputationRule

`func (o *SANElement) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *SANElement) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *SANElement) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *SANElement) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *SANElement) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *SANElement) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetEditableByRequester

`func (o *SANElement) GetEditableByRequester() bool`

GetEditableByRequester returns the EditableByRequester field if non-nil, zero value otherwise.

### GetEditableByRequesterOk

`func (o *SANElement) GetEditableByRequesterOk() (*bool, bool)`

GetEditableByRequesterOk returns a tuple with the EditableByRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByRequester

`func (o *SANElement) SetEditableByRequester(v bool)`

SetEditableByRequester sets EditableByRequester field to given value.

### HasEditableByRequester

`func (o *SANElement) HasEditableByRequester() bool`

HasEditableByRequester returns a boolean if a field has been set.

### SetEditableByRequesterNil

`func (o *SANElement) SetEditableByRequesterNil(b bool)`

 SetEditableByRequesterNil sets the value for EditableByRequester to be an explicit nil

### UnsetEditableByRequester
`func (o *SANElement) UnsetEditableByRequester()`

UnsetEditableByRequester ensures that no value is present for EditableByRequester, not even an explicit nil
### GetEditableByApprover

`func (o *SANElement) GetEditableByApprover() bool`

GetEditableByApprover returns the EditableByApprover field if non-nil, zero value otherwise.

### GetEditableByApproverOk

`func (o *SANElement) GetEditableByApproverOk() (*bool, bool)`

GetEditableByApproverOk returns a tuple with the EditableByApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByApprover

`func (o *SANElement) SetEditableByApprover(v bool)`

SetEditableByApprover sets EditableByApprover field to given value.

### HasEditableByApprover

`func (o *SANElement) HasEditableByApprover() bool`

HasEditableByApprover returns a boolean if a field has been set.

### SetEditableByApproverNil

`func (o *SANElement) SetEditableByApproverNil(b bool)`

 SetEditableByApproverNil sets the value for EditableByApprover to be an explicit nil

### UnsetEditableByApprover
`func (o *SANElement) UnsetEditableByApprover()`

UnsetEditableByApprover ensures that no value is present for EditableByApprover, not even an explicit nil
### GetRegex

`func (o *SANElement) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *SANElement) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *SANElement) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *SANElement) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *SANElement) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *SANElement) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetMin

`func (o *SANElement) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *SANElement) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *SANElement) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *SANElement) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *SANElement) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *SANElement) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *SANElement) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *SANElement) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *SANElement) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *SANElement) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *SANElement) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *SANElement) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


