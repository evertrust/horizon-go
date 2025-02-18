# OwnerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditableByRequester** | **bool** |  | 
**EditableByApprover** | **bool** |  | 
**ComputationRule** | Pointer to **NullableString** | A computation rule that will dynamically generate a string value from the request&#39;s context | [optional] 
**Mandatory** | **bool** |  | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 

## Methods

### NewOwnerPolicy

`func NewOwnerPolicy(editableByRequester bool, editableByApprover bool, mandatory bool, ) *OwnerPolicy`

NewOwnerPolicy instantiates a new OwnerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerPolicyWithDefaults

`func NewOwnerPolicyWithDefaults() *OwnerPolicy`

NewOwnerPolicyWithDefaults instantiates a new OwnerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditableByRequester

`func (o *OwnerPolicy) GetEditableByRequester() bool`

GetEditableByRequester returns the EditableByRequester field if non-nil, zero value otherwise.

### GetEditableByRequesterOk

`func (o *OwnerPolicy) GetEditableByRequesterOk() (*bool, bool)`

GetEditableByRequesterOk returns a tuple with the EditableByRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByRequester

`func (o *OwnerPolicy) SetEditableByRequester(v bool)`

SetEditableByRequester sets EditableByRequester field to given value.


### GetEditableByApprover

`func (o *OwnerPolicy) GetEditableByApprover() bool`

GetEditableByApprover returns the EditableByApprover field if non-nil, zero value otherwise.

### GetEditableByApproverOk

`func (o *OwnerPolicy) GetEditableByApproverOk() (*bool, bool)`

GetEditableByApproverOk returns a tuple with the EditableByApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByApprover

`func (o *OwnerPolicy) SetEditableByApprover(v bool)`

SetEditableByApprover sets EditableByApprover field to given value.


### GetComputationRule

`func (o *OwnerPolicy) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *OwnerPolicy) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *OwnerPolicy) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *OwnerPolicy) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *OwnerPolicy) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *OwnerPolicy) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetMandatory

`func (o *OwnerPolicy) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *OwnerPolicy) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *OwnerPolicy) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetDescription

`func (o *OwnerPolicy) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OwnerPolicy) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OwnerPolicy) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OwnerPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OwnerPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OwnerPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


