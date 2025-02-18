# PrincipalInfoSearchResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]PrincipalInfo**](PrincipalInfo.md) |  | 
**PageIndex** | **int64** | The current page index | 
**PageSize** | **int64** | The current number of returned element per page | 
**Count** | Pointer to **NullableInt64** | The total number of results (if _withCount_ was set to true in the query) | [optional] 
**HasMore** | **bool** | Whether there are more results to display (on another page) | 

## Methods

### NewPrincipalInfoSearchResultsResponse

`func NewPrincipalInfoSearchResultsResponse(results []PrincipalInfo, pageIndex int64, pageSize int64, hasMore bool, ) *PrincipalInfoSearchResultsResponse`

NewPrincipalInfoSearchResultsResponse instantiates a new PrincipalInfoSearchResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalInfoSearchResultsResponseWithDefaults

`func NewPrincipalInfoSearchResultsResponseWithDefaults() *PrincipalInfoSearchResultsResponse`

NewPrincipalInfoSearchResultsResponseWithDefaults instantiates a new PrincipalInfoSearchResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PrincipalInfoSearchResultsResponse) GetResults() []PrincipalInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PrincipalInfoSearchResultsResponse) GetResultsOk() (*[]PrincipalInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PrincipalInfoSearchResultsResponse) SetResults(v []PrincipalInfo)`

SetResults sets Results field to given value.


### GetPageIndex

`func (o *PrincipalInfoSearchResultsResponse) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *PrincipalInfoSearchResultsResponse) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *PrincipalInfoSearchResultsResponse) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *PrincipalInfoSearchResultsResponse) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PrincipalInfoSearchResultsResponse) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PrincipalInfoSearchResultsResponse) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetCount

`func (o *PrincipalInfoSearchResultsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PrincipalInfoSearchResultsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PrincipalInfoSearchResultsResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PrincipalInfoSearchResultsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PrincipalInfoSearchResultsResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PrincipalInfoSearchResultsResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetHasMore

`func (o *PrincipalInfoSearchResultsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PrincipalInfoSearchResultsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PrincipalInfoSearchResultsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


