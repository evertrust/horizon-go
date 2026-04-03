# IndexedDNElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputationRule** | Pointer to **NullableString** | Computation rule input will be evaluated and will override all other inputs | [optional] 
**Editable** | Pointer to **NullableBool** | Whether the field is editable or not for the currently authenticated user | [optional] 
**Element** | **string** | The element type and index. Indexes start at 1! Available elements are: &#x60;cn&#x60;, &#x60;e&#x60;, &#x60;ou&#x60;, &#x60;st&#x60;, &#x60;l&#x60;, &#x60;o&#x60;, &#x60;c&#x60;, &#x60;dc&#x60;, &#x60;uid&#x60;, &#x60;serialNumber&#x60;, &#x60;surname&#x60;, &#x60;givenName&#x60;, &#x60;unstructuredAddress&#x60;, &#x60;unstructuredName&#x60;, &#x60;organizationIdentifier&#x60;, &#x60;uniqueIdentifier&#x60;, &#x60;street&#x60;, &#x60;description&#x60;, &#x60;t&#x60; | 
**Mandatory** | Pointer to **NullableBool** | Whether the field is mandatory or not | [optional] 
**Regex** | Pointer to **NullableString** | A regular expression that will be used to validate the element&#39;s value | [optional] 
**Type** | Pointer to **NullableString** | The formatted element type | [optional] 
**Value** | Pointer to **NullableString** | The element value | [optional] 

## Methods

### NewIndexedDNElementResponse

`func NewIndexedDNElementResponse(element string, ) *IndexedDNElementResponse`

NewIndexedDNElementResponse instantiates a new IndexedDNElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexedDNElementResponseWithDefaults

`func NewIndexedDNElementResponseWithDefaults() *IndexedDNElementResponse`

NewIndexedDNElementResponseWithDefaults instantiates a new IndexedDNElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputationRule

`func (o *IndexedDNElementResponse) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *IndexedDNElementResponse) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *IndexedDNElementResponse) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *IndexedDNElementResponse) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *IndexedDNElementResponse) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *IndexedDNElementResponse) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetEditable

`func (o *IndexedDNElementResponse) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *IndexedDNElementResponse) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *IndexedDNElementResponse) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *IndexedDNElementResponse) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *IndexedDNElementResponse) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *IndexedDNElementResponse) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetElement

`func (o *IndexedDNElementResponse) GetElement() string`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *IndexedDNElementResponse) GetElementOk() (*string, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *IndexedDNElementResponse) SetElement(v string)`

SetElement sets Element field to given value.


### GetMandatory

`func (o *IndexedDNElementResponse) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *IndexedDNElementResponse) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *IndexedDNElementResponse) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *IndexedDNElementResponse) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### SetMandatoryNil

`func (o *IndexedDNElementResponse) SetMandatoryNil(b bool)`

 SetMandatoryNil sets the value for Mandatory to be an explicit nil

### UnsetMandatory
`func (o *IndexedDNElementResponse) UnsetMandatory()`

UnsetMandatory ensures that no value is present for Mandatory, not even an explicit nil
### GetRegex

`func (o *IndexedDNElementResponse) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *IndexedDNElementResponse) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *IndexedDNElementResponse) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *IndexedDNElementResponse) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *IndexedDNElementResponse) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *IndexedDNElementResponse) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetType

`func (o *IndexedDNElementResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IndexedDNElementResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IndexedDNElementResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IndexedDNElementResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *IndexedDNElementResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *IndexedDNElementResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *IndexedDNElementResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IndexedDNElementResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IndexedDNElementResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IndexedDNElementResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *IndexedDNElementResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *IndexedDNElementResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


