# DCVLifecycleEventSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt64** |  | [optional] 
**HasMore** | **bool** |  | 
**PageIndex** | **int64** |  | 
**PageSize** | **int64** |  | 
**Results** | [**[]DCVLifecycleEvent**](DCVLifecycleEvent.md) |  | 

## Methods

### NewDCVLifecycleEventSearchResults

`func NewDCVLifecycleEventSearchResults(hasMore bool, pageIndex int64, pageSize int64, results []DCVLifecycleEvent, ) *DCVLifecycleEventSearchResults`

NewDCVLifecycleEventSearchResults instantiates a new DCVLifecycleEventSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVLifecycleEventSearchResultsWithDefaults

`func NewDCVLifecycleEventSearchResultsWithDefaults() *DCVLifecycleEventSearchResults`

NewDCVLifecycleEventSearchResultsWithDefaults instantiates a new DCVLifecycleEventSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DCVLifecycleEventSearchResults) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DCVLifecycleEventSearchResults) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DCVLifecycleEventSearchResults) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *DCVLifecycleEventSearchResults) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *DCVLifecycleEventSearchResults) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *DCVLifecycleEventSearchResults) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetHasMore

`func (o *DCVLifecycleEventSearchResults) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *DCVLifecycleEventSearchResults) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *DCVLifecycleEventSearchResults) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetPageIndex

`func (o *DCVLifecycleEventSearchResults) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *DCVLifecycleEventSearchResults) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *DCVLifecycleEventSearchResults) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *DCVLifecycleEventSearchResults) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DCVLifecycleEventSearchResults) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DCVLifecycleEventSearchResults) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *DCVLifecycleEventSearchResults) GetResults() []DCVLifecycleEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DCVLifecycleEventSearchResults) GetResultsOk() (*[]DCVLifecycleEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DCVLifecycleEventSearchResults) SetResults(v []DCVLifecycleEvent)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


