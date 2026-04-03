# ExplainedGradingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apply** | **bool** |  | 
**Condition** | **string** |  | 
**Description** | [**[]LocalizedString**](LocalizedString.md) |  | 
**Eval** | Pointer to **NullableBool** |  | [optional] 
**Obtained** | Pointer to **NullableInt64** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Score** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewExplainedGradingRule

`func NewExplainedGradingRule(apply bool, condition string, description []LocalizedString, ) *ExplainedGradingRule`

NewExplainedGradingRule instantiates a new ExplainedGradingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplainedGradingRuleWithDefaults

`func NewExplainedGradingRuleWithDefaults() *ExplainedGradingRule`

NewExplainedGradingRuleWithDefaults instantiates a new ExplainedGradingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApply

`func (o *ExplainedGradingRule) GetApply() bool`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *ExplainedGradingRule) GetApplyOk() (*bool, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *ExplainedGradingRule) SetApply(v bool)`

SetApply sets Apply field to given value.


### GetCondition

`func (o *ExplainedGradingRule) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ExplainedGradingRule) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ExplainedGradingRule) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetDescription

`func (o *ExplainedGradingRule) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExplainedGradingRule) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExplainedGradingRule) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.


### GetEval

`func (o *ExplainedGradingRule) GetEval() bool`

GetEval returns the Eval field if non-nil, zero value otherwise.

### GetEvalOk

`func (o *ExplainedGradingRule) GetEvalOk() (*bool, bool)`

GetEvalOk returns a tuple with the Eval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEval

`func (o *ExplainedGradingRule) SetEval(v bool)`

SetEval sets Eval field to given value.

### HasEval

`func (o *ExplainedGradingRule) HasEval() bool`

HasEval returns a boolean if a field has been set.

### SetEvalNil

`func (o *ExplainedGradingRule) SetEvalNil(b bool)`

 SetEvalNil sets the value for Eval to be an explicit nil

### UnsetEval
`func (o *ExplainedGradingRule) UnsetEval()`

UnsetEval ensures that no value is present for Eval, not even an explicit nil
### GetObtained

`func (o *ExplainedGradingRule) GetObtained() int64`

GetObtained returns the Obtained field if non-nil, zero value otherwise.

### GetObtainedOk

`func (o *ExplainedGradingRule) GetObtainedOk() (*int64, bool)`

GetObtainedOk returns a tuple with the Obtained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtained

`func (o *ExplainedGradingRule) SetObtained(v int64)`

SetObtained sets Obtained field to given value.

### HasObtained

`func (o *ExplainedGradingRule) HasObtained() bool`

HasObtained returns a boolean if a field has been set.

### SetObtainedNil

`func (o *ExplainedGradingRule) SetObtainedNil(b bool)`

 SetObtainedNil sets the value for Obtained to be an explicit nil

### UnsetObtained
`func (o *ExplainedGradingRule) UnsetObtained()`

UnsetObtained ensures that no value is present for Obtained, not even an explicit nil
### GetScope

`func (o *ExplainedGradingRule) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ExplainedGradingRule) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ExplainedGradingRule) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ExplainedGradingRule) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *ExplainedGradingRule) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *ExplainedGradingRule) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetScore

`func (o *ExplainedGradingRule) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ExplainedGradingRule) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ExplainedGradingRule) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *ExplainedGradingRule) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *ExplainedGradingRule) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *ExplainedGradingRule) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


