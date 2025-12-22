# GradingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | **string** |  | 
**Description** | [**[]LocalizedString**](LocalizedString.md) |  | 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Score** | **int64** |  | 

## Methods

### NewGradingRule

`func NewGradingRule(condition string, description []LocalizedString, score int64, ) *GradingRule`

NewGradingRule instantiates a new GradingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGradingRuleWithDefaults

`func NewGradingRuleWithDefaults() *GradingRule`

NewGradingRuleWithDefaults instantiates a new GradingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *GradingRule) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GradingRule) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GradingRule) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetDescription

`func (o *GradingRule) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GradingRule) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GradingRule) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.


### GetScope

`func (o *GradingRule) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GradingRule) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GradingRule) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GradingRule) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *GradingRule) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *GradingRule) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetScore

`func (o *GradingRule) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GradingRule) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GradingRule) SetScore(v int64)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


