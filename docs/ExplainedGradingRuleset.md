# ExplainedGradingRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apply** | **bool** |  | 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**Explained** | Pointer to [**[]ExplainedGradingRule**](ExplainedGradingRule.md) |  | [optional] 
**Max** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Obtained** | Pointer to **NullableInt64** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Score** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewExplainedGradingRuleset

`func NewExplainedGradingRuleset(apply bool, name string, ) *ExplainedGradingRuleset`

NewExplainedGradingRuleset instantiates a new ExplainedGradingRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplainedGradingRulesetWithDefaults

`func NewExplainedGradingRulesetWithDefaults() *ExplainedGradingRuleset`

NewExplainedGradingRulesetWithDefaults instantiates a new ExplainedGradingRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApply

`func (o *ExplainedGradingRuleset) GetApply() bool`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *ExplainedGradingRuleset) GetApplyOk() (*bool, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *ExplainedGradingRuleset) SetApply(v bool)`

SetApply sets Apply field to given value.


### GetCertificate

`func (o *ExplainedGradingRuleset) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ExplainedGradingRuleset) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ExplainedGradingRuleset) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ExplainedGradingRuleset) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ExplainedGradingRuleset) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ExplainedGradingRuleset) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetDescription

`func (o *ExplainedGradingRuleset) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExplainedGradingRuleset) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExplainedGradingRuleset) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExplainedGradingRuleset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExplainedGradingRuleset) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExplainedGradingRuleset) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExplained

`func (o *ExplainedGradingRuleset) GetExplained() []ExplainedGradingRule`

GetExplained returns the Explained field if non-nil, zero value otherwise.

### GetExplainedOk

`func (o *ExplainedGradingRuleset) GetExplainedOk() (*[]ExplainedGradingRule, bool)`

GetExplainedOk returns a tuple with the Explained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplained

`func (o *ExplainedGradingRuleset) SetExplained(v []ExplainedGradingRule)`

SetExplained sets Explained field to given value.

### HasExplained

`func (o *ExplainedGradingRuleset) HasExplained() bool`

HasExplained returns a boolean if a field has been set.

### SetExplainedNil

`func (o *ExplainedGradingRuleset) SetExplainedNil(b bool)`

 SetExplainedNil sets the value for Explained to be an explicit nil

### UnsetExplained
`func (o *ExplainedGradingRuleset) UnsetExplained()`

UnsetExplained ensures that no value is present for Explained, not even an explicit nil
### GetMax

`func (o *ExplainedGradingRuleset) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ExplainedGradingRuleset) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ExplainedGradingRuleset) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ExplainedGradingRuleset) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *ExplainedGradingRuleset) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *ExplainedGradingRuleset) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetName

`func (o *ExplainedGradingRuleset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExplainedGradingRuleset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExplainedGradingRuleset) SetName(v string)`

SetName sets Name field to given value.


### GetObtained

`func (o *ExplainedGradingRuleset) GetObtained() int64`

GetObtained returns the Obtained field if non-nil, zero value otherwise.

### GetObtainedOk

`func (o *ExplainedGradingRuleset) GetObtainedOk() (*int64, bool)`

GetObtainedOk returns a tuple with the Obtained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtained

`func (o *ExplainedGradingRuleset) SetObtained(v int64)`

SetObtained sets Obtained field to given value.

### HasObtained

`func (o *ExplainedGradingRuleset) HasObtained() bool`

HasObtained returns a boolean if a field has been set.

### SetObtainedNil

`func (o *ExplainedGradingRuleset) SetObtainedNil(b bool)`

 SetObtainedNil sets the value for Obtained to be an explicit nil

### UnsetObtained
`func (o *ExplainedGradingRuleset) UnsetObtained()`

UnsetObtained ensures that no value is present for Obtained, not even an explicit nil
### GetScope

`func (o *ExplainedGradingRuleset) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ExplainedGradingRuleset) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ExplainedGradingRuleset) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ExplainedGradingRuleset) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *ExplainedGradingRuleset) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *ExplainedGradingRuleset) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetScore

`func (o *ExplainedGradingRuleset) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ExplainedGradingRuleset) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ExplainedGradingRuleset) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ExplainedGradingRuleset) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *ExplainedGradingRuleset) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *ExplainedGradingRuleset) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


