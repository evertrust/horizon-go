# ListSANElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableString** | SAN type | 
**Value** | Pointer to **[]string** | SAN value | [optional] 

## Methods

### NewListSANElement

`func NewListSANElement(type_ NullableString, ) *ListSANElement`

NewListSANElement instantiates a new ListSANElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSANElementWithDefaults

`func NewListSANElementWithDefaults() *ListSANElement`

NewListSANElementWithDefaults instantiates a new ListSANElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListSANElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListSANElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListSANElement) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ListSANElement) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ListSANElement) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *ListSANElement) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListSANElement) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListSANElement) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListSANElement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ListSANElement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ListSANElement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


