# GradingPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**Rulesets** | [**[]WeightedGradingRuleset**](WeightedGradingRuleset.md) |  | 

## Methods

### NewGradingPolicyResponse

`func NewGradingPolicyResponse(id string, name string, rulesets []WeightedGradingRuleset, ) *GradingPolicyResponse`

NewGradingPolicyResponse instantiates a new GradingPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGradingPolicyResponseWithDefaults

`func NewGradingPolicyResponseWithDefaults() *GradingPolicyResponse`

NewGradingPolicyResponseWithDefaults instantiates a new GradingPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GradingPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GradingPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GradingPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GradingPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GradingPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GradingPolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GradingPolicyResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GradingPolicyResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GradingPolicyResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GradingPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GradingPolicyResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GradingPolicyResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRulesets

`func (o *GradingPolicyResponse) GetRulesets() []WeightedGradingRuleset`

GetRulesets returns the Rulesets field if non-nil, zero value otherwise.

### GetRulesetsOk

`func (o *GradingPolicyResponse) GetRulesetsOk() (*[]WeightedGradingRuleset, bool)`

GetRulesetsOk returns a tuple with the Rulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesets

`func (o *GradingPolicyResponse) SetRulesets(v []WeightedGradingRuleset)`

SetRulesets sets Rulesets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


