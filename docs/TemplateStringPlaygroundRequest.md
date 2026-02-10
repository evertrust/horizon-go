# TemplateStringPlaygroundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dictionary** | Pointer to **map[string]string** |  | [optional] 
**ComputationRule** | Pointer to **NullableString** | A computation rule that will dynamically generate a string value from the request&#39;s context | [optional] 
**Csr** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTemplateStringPlaygroundRequest

`func NewTemplateStringPlaygroundRequest() *TemplateStringPlaygroundRequest`

NewTemplateStringPlaygroundRequest instantiates a new TemplateStringPlaygroundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateStringPlaygroundRequestWithDefaults

`func NewTemplateStringPlaygroundRequestWithDefaults() *TemplateStringPlaygroundRequest`

NewTemplateStringPlaygroundRequestWithDefaults instantiates a new TemplateStringPlaygroundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDictionary

`func (o *TemplateStringPlaygroundRequest) GetDictionary() map[string]string`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *TemplateStringPlaygroundRequest) GetDictionaryOk() (*map[string]string, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *TemplateStringPlaygroundRequest) SetDictionary(v map[string]string)`

SetDictionary sets Dictionary field to given value.

### HasDictionary

`func (o *TemplateStringPlaygroundRequest) HasDictionary() bool`

HasDictionary returns a boolean if a field has been set.

### SetDictionaryNil

`func (o *TemplateStringPlaygroundRequest) SetDictionaryNil(b bool)`

 SetDictionaryNil sets the value for Dictionary to be an explicit nil

### UnsetDictionary
`func (o *TemplateStringPlaygroundRequest) UnsetDictionary()`

UnsetDictionary ensures that no value is present for Dictionary, not even an explicit nil
### GetComputationRule

`func (o *TemplateStringPlaygroundRequest) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *TemplateStringPlaygroundRequest) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *TemplateStringPlaygroundRequest) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *TemplateStringPlaygroundRequest) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *TemplateStringPlaygroundRequest) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *TemplateStringPlaygroundRequest) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetCsr

`func (o *TemplateStringPlaygroundRequest) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *TemplateStringPlaygroundRequest) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *TemplateStringPlaygroundRequest) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *TemplateStringPlaygroundRequest) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *TemplateStringPlaygroundRequest) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *TemplateStringPlaygroundRequest) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


