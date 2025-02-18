# ExplainedGradingRulesetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Apply** | **bool** |  | 
**Max** | Pointer to **NullableInt64** |  | [optional] 
**Obtained** | Pointer to **NullableInt64** |  | [optional] 
**Score** | Pointer to **NullableFloat32** |  | [optional] 
**Explained** | Pointer to [**[]ExplainedGradingRule**](ExplainedGradingRule.md) |  | [optional] 

## Methods

### NewExplainedGradingRulesetResponse

`func NewExplainedGradingRulesetResponse(name string, apply bool, ) *ExplainedGradingRulesetResponse`

NewExplainedGradingRulesetResponse instantiates a new ExplainedGradingRulesetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplainedGradingRulesetResponseWithDefaults

`func NewExplainedGradingRulesetResponseWithDefaults() *ExplainedGradingRulesetResponse`

NewExplainedGradingRulesetResponseWithDefaults instantiates a new ExplainedGradingRulesetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExplainedGradingRulesetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExplainedGradingRulesetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExplainedGradingRulesetResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExplainedGradingRulesetResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExplainedGradingRulesetResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExplainedGradingRulesetResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExplainedGradingRulesetResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExplainedGradingRulesetResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExplainedGradingRulesetResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCertificate

`func (o *ExplainedGradingRulesetResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ExplainedGradingRulesetResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ExplainedGradingRulesetResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ExplainedGradingRulesetResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ExplainedGradingRulesetResponse) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ExplainedGradingRulesetResponse) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetScope

`func (o *ExplainedGradingRulesetResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ExplainedGradingRulesetResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ExplainedGradingRulesetResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ExplainedGradingRulesetResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *ExplainedGradingRulesetResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *ExplainedGradingRulesetResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetApply

`func (o *ExplainedGradingRulesetResponse) GetApply() bool`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *ExplainedGradingRulesetResponse) GetApplyOk() (*bool, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *ExplainedGradingRulesetResponse) SetApply(v bool)`

SetApply sets Apply field to given value.


### GetMax

`func (o *ExplainedGradingRulesetResponse) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ExplainedGradingRulesetResponse) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ExplainedGradingRulesetResponse) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ExplainedGradingRulesetResponse) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *ExplainedGradingRulesetResponse) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *ExplainedGradingRulesetResponse) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetObtained

`func (o *ExplainedGradingRulesetResponse) GetObtained() int64`

GetObtained returns the Obtained field if non-nil, zero value otherwise.

### GetObtainedOk

`func (o *ExplainedGradingRulesetResponse) GetObtainedOk() (*int64, bool)`

GetObtainedOk returns a tuple with the Obtained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtained

`func (o *ExplainedGradingRulesetResponse) SetObtained(v int64)`

SetObtained sets Obtained field to given value.

### HasObtained

`func (o *ExplainedGradingRulesetResponse) HasObtained() bool`

HasObtained returns a boolean if a field has been set.

### SetObtainedNil

`func (o *ExplainedGradingRulesetResponse) SetObtainedNil(b bool)`

 SetObtainedNil sets the value for Obtained to be an explicit nil

### UnsetObtained
`func (o *ExplainedGradingRulesetResponse) UnsetObtained()`

UnsetObtained ensures that no value is present for Obtained, not even an explicit nil
### GetScore

`func (o *ExplainedGradingRulesetResponse) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ExplainedGradingRulesetResponse) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ExplainedGradingRulesetResponse) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ExplainedGradingRulesetResponse) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *ExplainedGradingRulesetResponse) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *ExplainedGradingRulesetResponse) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetExplained

`func (o *ExplainedGradingRulesetResponse) GetExplained() []ExplainedGradingRule`

GetExplained returns the Explained field if non-nil, zero value otherwise.

### GetExplainedOk

`func (o *ExplainedGradingRulesetResponse) GetExplainedOk() (*[]ExplainedGradingRule, bool)`

GetExplainedOk returns a tuple with the Explained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplained

`func (o *ExplainedGradingRulesetResponse) SetExplained(v []ExplainedGradingRule)`

SetExplained sets Explained field to given value.

### HasExplained

`func (o *ExplainedGradingRulesetResponse) HasExplained() bool`

HasExplained returns a boolean if a field has been set.

### SetExplainedNil

`func (o *ExplainedGradingRulesetResponse) SetExplainedNil(b bool)`

 SetExplainedNil sets the value for Explained to be an explicit nil

### UnsetExplained
`func (o *ExplainedGradingRulesetResponse) UnsetExplained()`

UnsetExplained ensures that no value is present for Explained, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


