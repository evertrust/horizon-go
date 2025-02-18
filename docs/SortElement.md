# SortElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Element** | **string** | The name of the field the query should be sorted by. Must be one of the &#x60;fields&#x60; of the search query | 
**Order** | **string** | The order to use for the sort. &#x60;Asc&#x60; and &#x60;Desc&#x60; sort the values, and &#x60;KeyAsc&#x60; and &#x60;KeyDesc&#x60; sort by the key, in case of a key-value element | 

## Methods

### NewSortElement

`func NewSortElement(element string, order string, ) *SortElement`

NewSortElement instantiates a new SortElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortElementWithDefaults

`func NewSortElementWithDefaults() *SortElement`

NewSortElementWithDefaults instantiates a new SortElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElement

`func (o *SortElement) GetElement() string`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *SortElement) GetElementOk() (*string, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *SortElement) SetElement(v string)`

SetElement sets Element field to given value.


### GetOrder

`func (o *SortElement) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SortElement) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SortElement) SetOrder(v string)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


