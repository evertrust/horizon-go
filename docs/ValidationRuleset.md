# ValidationRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | **[]string** | The validation rules for this ruleset | 
**Threshold** | **int64** | Number of rules to validation in order to allow enrollment | 

## Methods

### NewValidationRuleset

`func NewValidationRuleset(rules []string, threshold int64, ) *ValidationRuleset`

NewValidationRuleset instantiates a new ValidationRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationRulesetWithDefaults

`func NewValidationRulesetWithDefaults() *ValidationRuleset`

NewValidationRulesetWithDefaults instantiates a new ValidationRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *ValidationRuleset) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ValidationRuleset) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ValidationRuleset) SetRules(v []string)`

SetRules sets Rules field to given value.


### GetThreshold

`func (o *ValidationRuleset) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ValidationRuleset) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ValidationRuleset) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


