# ListSANElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputationRule** | Pointer to **NullableString** | Computation rule input will be evaluated and will override all other inputs | [optional] 
**Editable** | Pointer to **NullableBool** | Whether this SAN element is editable by the user | [optional] 
**Max** | Pointer to **NullableInt64** | The maximum number of SAN elements that can be provided | [optional] 
**Min** | Pointer to **NullableInt64** | The minimum number of SAN elements that must be provided | [optional] 
**Regex** | Pointer to **NullableString** | The regex that will be used to validate the SAN value | [optional] 
**Type** | **NullableString** | SAN type | 
**Value** | Pointer to **[]string** | SAN value | [optional] 

## Methods

### NewListSANElementResponse

`func NewListSANElementResponse(type_ NullableString, ) *ListSANElementResponse`

NewListSANElementResponse instantiates a new ListSANElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSANElementResponseWithDefaults

`func NewListSANElementResponseWithDefaults() *ListSANElementResponse`

NewListSANElementResponseWithDefaults instantiates a new ListSANElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputationRule

`func (o *ListSANElementResponse) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *ListSANElementResponse) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *ListSANElementResponse) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *ListSANElementResponse) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *ListSANElementResponse) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *ListSANElementResponse) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetEditable

`func (o *ListSANElementResponse) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListSANElementResponse) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListSANElementResponse) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListSANElementResponse) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *ListSANElementResponse) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *ListSANElementResponse) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetMax

`func (o *ListSANElementResponse) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ListSANElementResponse) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ListSANElementResponse) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ListSANElementResponse) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *ListSANElementResponse) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *ListSANElementResponse) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetMin

`func (o *ListSANElementResponse) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ListSANElementResponse) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ListSANElementResponse) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *ListSANElementResponse) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *ListSANElementResponse) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *ListSANElementResponse) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetRegex

`func (o *ListSANElementResponse) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *ListSANElementResponse) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *ListSANElementResponse) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *ListSANElementResponse) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *ListSANElementResponse) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *ListSANElementResponse) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetType

`func (o *ListSANElementResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListSANElementResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListSANElementResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ListSANElementResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ListSANElementResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *ListSANElementResponse) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListSANElementResponse) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListSANElementResponse) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListSANElementResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ListSANElementResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ListSANElementResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


