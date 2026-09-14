# OrderSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt64** | Total number of items matching the query | [optional] 
**HasMore** | **bool** | Indicates if there are more results available | [default to false]
**PageIndex** | **int64** | Current page index (0-based) | [default to 0]
**PageSize** | **int64** | Number of items per page | [default to 20]
**Results** | [**[]OrderResponse**](OrderResponse.md) |  | 

## Methods

### NewOrderSearchResults

`func NewOrderSearchResults(hasMore bool, pageIndex int64, pageSize int64, results []OrderResponse, ) *OrderSearchResults`

NewOrderSearchResults instantiates a new OrderSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSearchResultsWithDefaults

`func NewOrderSearchResultsWithDefaults() *OrderSearchResults`

NewOrderSearchResultsWithDefaults instantiates a new OrderSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *OrderSearchResults) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OrderSearchResults) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OrderSearchResults) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *OrderSearchResults) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *OrderSearchResults) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *OrderSearchResults) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetHasMore

`func (o *OrderSearchResults) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OrderSearchResults) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OrderSearchResults) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetPageIndex

`func (o *OrderSearchResults) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *OrderSearchResults) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *OrderSearchResults) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *OrderSearchResults) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OrderSearchResults) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OrderSearchResults) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *OrderSearchResults) GetResults() []OrderResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OrderSearchResults) GetResultsOk() (*[]OrderResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OrderSearchResults) SetResults(v []OrderResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


