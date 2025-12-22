# RequestSearchResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The total count of requests matching the HRQL query | [optional] 
**HasMore** | **bool** | Indicates whether the response represents the last page of results (if set to &#x60;false&#x60;) or not (if set to &#x60;true&#x60;) | 
**PageIndex** | **int64** | The index of the results page | 
**PageSize** | **int64** | The maximum number of items on this page | 
**Results** | [**[]RequestSearchResult**](RequestSearchResult.md) | The list of requests matching the HRQL query | 

## Methods

### NewRequestSearchResultsResponse

`func NewRequestSearchResultsResponse(hasMore bool, pageIndex int64, pageSize int64, results []RequestSearchResult, ) *RequestSearchResultsResponse`

NewRequestSearchResultsResponse instantiates a new RequestSearchResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSearchResultsResponseWithDefaults

`func NewRequestSearchResultsResponseWithDefaults() *RequestSearchResultsResponse`

NewRequestSearchResultsResponseWithDefaults instantiates a new RequestSearchResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RequestSearchResultsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RequestSearchResultsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RequestSearchResultsResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RequestSearchResultsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHasMore

`func (o *RequestSearchResultsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *RequestSearchResultsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *RequestSearchResultsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetPageIndex

`func (o *RequestSearchResultsResponse) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *RequestSearchResultsResponse) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *RequestSearchResultsResponse) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *RequestSearchResultsResponse) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *RequestSearchResultsResponse) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *RequestSearchResultsResponse) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *RequestSearchResultsResponse) GetResults() []RequestSearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RequestSearchResultsResponse) GetResultsOk() (*[]RequestSearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RequestSearchResultsResponse) SetResults(v []RequestSearchResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


