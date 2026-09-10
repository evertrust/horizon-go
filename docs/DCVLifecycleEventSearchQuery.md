# DCVLifecycleEventSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **NullableInt64** |  | [optional] 
**PageSize** | Pointer to **NullableInt64** |  | [optional] 
**SortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**WithCount** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewDCVLifecycleEventSearchQuery

`func NewDCVLifecycleEventSearchQuery() *DCVLifecycleEventSearchQuery`

NewDCVLifecycleEventSearchQuery instantiates a new DCVLifecycleEventSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVLifecycleEventSearchQueryWithDefaults

`func NewDCVLifecycleEventSearchQueryWithDefaults() *DCVLifecycleEventSearchQuery`

NewDCVLifecycleEventSearchQueryWithDefaults instantiates a new DCVLifecycleEventSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *DCVLifecycleEventSearchQuery) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *DCVLifecycleEventSearchQuery) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *DCVLifecycleEventSearchQuery) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *DCVLifecycleEventSearchQuery) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### SetPageIndexNil

`func (o *DCVLifecycleEventSearchQuery) SetPageIndexNil(b bool)`

 SetPageIndexNil sets the value for PageIndex to be an explicit nil

### UnsetPageIndex
`func (o *DCVLifecycleEventSearchQuery) UnsetPageIndex()`

UnsetPageIndex ensures that no value is present for PageIndex, not even an explicit nil
### GetPageSize

`func (o *DCVLifecycleEventSearchQuery) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DCVLifecycleEventSearchQuery) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DCVLifecycleEventSearchQuery) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DCVLifecycleEventSearchQuery) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *DCVLifecycleEventSearchQuery) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *DCVLifecycleEventSearchQuery) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetSortedBy

`func (o *DCVLifecycleEventSearchQuery) GetSortedBy() []SortElement`

GetSortedBy returns the SortedBy field if non-nil, zero value otherwise.

### GetSortedByOk

`func (o *DCVLifecycleEventSearchQuery) GetSortedByOk() (*[]SortElement, bool)`

GetSortedByOk returns a tuple with the SortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortedBy

`func (o *DCVLifecycleEventSearchQuery) SetSortedBy(v []SortElement)`

SetSortedBy sets SortedBy field to given value.

### HasSortedBy

`func (o *DCVLifecycleEventSearchQuery) HasSortedBy() bool`

HasSortedBy returns a boolean if a field has been set.

### SetSortedByNil

`func (o *DCVLifecycleEventSearchQuery) SetSortedByNil(b bool)`

 SetSortedByNil sets the value for SortedBy to be an explicit nil

### UnsetSortedBy
`func (o *DCVLifecycleEventSearchQuery) UnsetSortedBy()`

UnsetSortedBy ensures that no value is present for SortedBy, not even an explicit nil
### GetWithCount

`func (o *DCVLifecycleEventSearchQuery) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *DCVLifecycleEventSearchQuery) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *DCVLifecycleEventSearchQuery) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *DCVLifecycleEventSearchQuery) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.

### SetWithCountNil

`func (o *DCVLifecycleEventSearchQuery) SetWithCountNil(b bool)`

 SetWithCountNil sets the value for WithCount to be an explicit nil

### UnsetWithCount
`func (o *DCVLifecycleEventSearchQuery) UnsetWithCount()`

UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


