# EventSearchResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]EventSearchResult**](EventSearchResult.md) |  | 
**PageIndex** | **int64** |  | 
**PageSize** | **int64** |  | 
**Count** | Pointer to **NullableInt64** |  | [optional] 
**HasMore** | **bool** |  | 

## Methods

### NewEventSearchResultsResponse

`func NewEventSearchResultsResponse(results []EventSearchResult, pageIndex int64, pageSize int64, hasMore bool, ) *EventSearchResultsResponse`

NewEventSearchResultsResponse instantiates a new EventSearchResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSearchResultsResponseWithDefaults

`func NewEventSearchResultsResponseWithDefaults() *EventSearchResultsResponse`

NewEventSearchResultsResponseWithDefaults instantiates a new EventSearchResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *EventSearchResultsResponse) GetResults() []EventSearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EventSearchResultsResponse) GetResultsOk() (*[]EventSearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EventSearchResultsResponse) SetResults(v []EventSearchResult)`

SetResults sets Results field to given value.


### GetPageIndex

`func (o *EventSearchResultsResponse) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *EventSearchResultsResponse) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *EventSearchResultsResponse) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *EventSearchResultsResponse) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EventSearchResultsResponse) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EventSearchResultsResponse) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetCount

`func (o *EventSearchResultsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventSearchResultsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventSearchResultsResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *EventSearchResultsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *EventSearchResultsResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *EventSearchResultsResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetHasMore

`func (o *EventSearchResultsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *EventSearchResultsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *EventSearchResultsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


