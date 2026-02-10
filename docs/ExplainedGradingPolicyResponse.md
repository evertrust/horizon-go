# ExplainedGradingPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**Rulesets** | [**[]WeightedGradingRuleset**](WeightedGradingRuleset.md) |  | 
**Certificate** | **string** |  | 
**Score** | Pointer to **NullableFloat32** |  | [optional] 
**Grade** | Pointer to **NullableString** |  | [optional] 
**Explained** | Pointer to [**[]ExplainedGradingRuleset**](ExplainedGradingRuleset.md) |  | [optional] 

## Methods

### NewExplainedGradingPolicyResponse

`func NewExplainedGradingPolicyResponse(name string, rulesets []WeightedGradingRuleset, certificate string, ) *ExplainedGradingPolicyResponse`

NewExplainedGradingPolicyResponse instantiates a new ExplainedGradingPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplainedGradingPolicyResponseWithDefaults

`func NewExplainedGradingPolicyResponseWithDefaults() *ExplainedGradingPolicyResponse`

NewExplainedGradingPolicyResponseWithDefaults instantiates a new ExplainedGradingPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExplainedGradingPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExplainedGradingPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExplainedGradingPolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExplainedGradingPolicyResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExplainedGradingPolicyResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExplainedGradingPolicyResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExplainedGradingPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExplainedGradingPolicyResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExplainedGradingPolicyResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRulesets

`func (o *ExplainedGradingPolicyResponse) GetRulesets() []WeightedGradingRuleset`

GetRulesets returns the Rulesets field if non-nil, zero value otherwise.

### GetRulesetsOk

`func (o *ExplainedGradingPolicyResponse) GetRulesetsOk() (*[]WeightedGradingRuleset, bool)`

GetRulesetsOk returns a tuple with the Rulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesets

`func (o *ExplainedGradingPolicyResponse) SetRulesets(v []WeightedGradingRuleset)`

SetRulesets sets Rulesets field to given value.


### GetCertificate

`func (o *ExplainedGradingPolicyResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ExplainedGradingPolicyResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ExplainedGradingPolicyResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetScore

`func (o *ExplainedGradingPolicyResponse) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ExplainedGradingPolicyResponse) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ExplainedGradingPolicyResponse) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ExplainedGradingPolicyResponse) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *ExplainedGradingPolicyResponse) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *ExplainedGradingPolicyResponse) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetGrade

`func (o *ExplainedGradingPolicyResponse) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *ExplainedGradingPolicyResponse) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *ExplainedGradingPolicyResponse) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *ExplainedGradingPolicyResponse) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### SetGradeNil

`func (o *ExplainedGradingPolicyResponse) SetGradeNil(b bool)`

 SetGradeNil sets the value for Grade to be an explicit nil

### UnsetGrade
`func (o *ExplainedGradingPolicyResponse) UnsetGrade()`

UnsetGrade ensures that no value is present for Grade, not even an explicit nil
### GetExplained

`func (o *ExplainedGradingPolicyResponse) GetExplained() []ExplainedGradingRuleset`

GetExplained returns the Explained field if non-nil, zero value otherwise.

### GetExplainedOk

`func (o *ExplainedGradingPolicyResponse) GetExplainedOk() (*[]ExplainedGradingRuleset, bool)`

GetExplainedOk returns a tuple with the Explained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplained

`func (o *ExplainedGradingPolicyResponse) SetExplained(v []ExplainedGradingRuleset)`

SetExplained sets Explained field to given value.

### HasExplained

`func (o *ExplainedGradingPolicyResponse) HasExplained() bool`

HasExplained returns a boolean if a field has been set.

### SetExplainedNil

`func (o *ExplainedGradingPolicyResponse) SetExplainedNil(b bool)`

 SetExplainedNil sets the value for Explained to be an explicit nil

### UnsetExplained
`func (o *ExplainedGradingPolicyResponse) UnsetExplained()`

UnsetExplained ensures that no value is present for Explained, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


