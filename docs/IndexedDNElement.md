# IndexedDNElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Element** | **string** | The element type and index. Indexes start at 1! Available elements are: &#x60;cn&#x60;, &#x60;e&#x60;, &#x60;ou&#x60;, &#x60;st&#x60;, &#x60;l&#x60;, &#x60;o&#x60;, &#x60;c&#x60;, &#x60;dc&#x60;, &#x60;uid&#x60;, &#x60;serialNumber&#x60;, &#x60;surname&#x60;, &#x60;givenName&#x60;, &#x60;unstructuredAddress&#x60;, &#x60;unstructuredName&#x60;, &#x60;organizationIdentifier&#x60;, &#x60;uniqueIdentifier&#x60;, &#x60;street&#x60;, &#x60;description&#x60;, &#x60;t&#x60; | 
**Value** | Pointer to **NullableString** | The element value | [optional] 

## Methods

### NewIndexedDNElement

`func NewIndexedDNElement(element string, ) *IndexedDNElement`

NewIndexedDNElement instantiates a new IndexedDNElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexedDNElementWithDefaults

`func NewIndexedDNElementWithDefaults() *IndexedDNElement`

NewIndexedDNElementWithDefaults instantiates a new IndexedDNElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElement

`func (o *IndexedDNElement) GetElement() string`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *IndexedDNElement) GetElementOk() (*string, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *IndexedDNElement) SetElement(v string)`

SetElement sets Element field to given value.


### GetValue

`func (o *IndexedDNElement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IndexedDNElement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IndexedDNElement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IndexedDNElement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *IndexedDNElement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *IndexedDNElement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


