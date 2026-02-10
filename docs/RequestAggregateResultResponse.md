# RequestAggregateResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt64** | The total number of requests matching the query | [optional] 
**Items** | [**[]RequestAggregateResultResponseItemsInner**](RequestAggregateResultResponseItemsInner.md) | All the groups in this aggregate | 

## Methods

### NewRequestAggregateResultResponse

`func NewRequestAggregateResultResponse(items []RequestAggregateResultResponseItemsInner, ) *RequestAggregateResultResponse`

NewRequestAggregateResultResponse instantiates a new RequestAggregateResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestAggregateResultResponseWithDefaults

`func NewRequestAggregateResultResponseWithDefaults() *RequestAggregateResultResponse`

NewRequestAggregateResultResponseWithDefaults instantiates a new RequestAggregateResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RequestAggregateResultResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RequestAggregateResultResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RequestAggregateResultResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RequestAggregateResultResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *RequestAggregateResultResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *RequestAggregateResultResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetItems

`func (o *RequestAggregateResultResponse) GetItems() []RequestAggregateResultResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RequestAggregateResultResponse) GetItemsOk() (*[]RequestAggregateResultResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RequestAggregateResultResponse) SetItems(v []RequestAggregateResultResponseItemsInner)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


