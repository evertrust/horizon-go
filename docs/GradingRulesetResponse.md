# GradingRulesetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Rules** | [**[]GradingRule**](GradingRule.md) |  | 

## Methods

### NewGradingRulesetResponse

`func NewGradingRulesetResponse(id string, name string, rules []GradingRule, ) *GradingRulesetResponse`

NewGradingRulesetResponse instantiates a new GradingRulesetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGradingRulesetResponseWithDefaults

`func NewGradingRulesetResponseWithDefaults() *GradingRulesetResponse`

NewGradingRulesetResponseWithDefaults instantiates a new GradingRulesetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GradingRulesetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GradingRulesetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GradingRulesetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GradingRulesetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GradingRulesetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GradingRulesetResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GradingRulesetResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GradingRulesetResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GradingRulesetResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GradingRulesetResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GradingRulesetResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GradingRulesetResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScope

`func (o *GradingRulesetResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GradingRulesetResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GradingRulesetResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GradingRulesetResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *GradingRulesetResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *GradingRulesetResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetRules

`func (o *GradingRulesetResponse) GetRules() []GradingRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GradingRulesetResponse) GetRulesOk() (*[]GradingRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GradingRulesetResponse) SetRules(v []GradingRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


