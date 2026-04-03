# DiscoveryEventSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **NullableInt64** |  | [optional] 
**PageSize** | Pointer to **NullableInt64** |  | [optional] 
**Query** | Pointer to **NullableString** |  | [optional] 
**SortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**WithCount** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewDiscoveryEventSearchQuery

`func NewDiscoveryEventSearchQuery() *DiscoveryEventSearchQuery`

NewDiscoveryEventSearchQuery instantiates a new DiscoveryEventSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryEventSearchQueryWithDefaults

`func NewDiscoveryEventSearchQueryWithDefaults() *DiscoveryEventSearchQuery`

NewDiscoveryEventSearchQueryWithDefaults instantiates a new DiscoveryEventSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *DiscoveryEventSearchQuery) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *DiscoveryEventSearchQuery) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *DiscoveryEventSearchQuery) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *DiscoveryEventSearchQuery) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### SetPageIndexNil

`func (o *DiscoveryEventSearchQuery) SetPageIndexNil(b bool)`

 SetPageIndexNil sets the value for PageIndex to be an explicit nil

### UnsetPageIndex
`func (o *DiscoveryEventSearchQuery) UnsetPageIndex()`

UnsetPageIndex ensures that no value is present for PageIndex, not even an explicit nil
### GetPageSize

`func (o *DiscoveryEventSearchQuery) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DiscoveryEventSearchQuery) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DiscoveryEventSearchQuery) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DiscoveryEventSearchQuery) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *DiscoveryEventSearchQuery) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *DiscoveryEventSearchQuery) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetQuery

`func (o *DiscoveryEventSearchQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DiscoveryEventSearchQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DiscoveryEventSearchQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DiscoveryEventSearchQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *DiscoveryEventSearchQuery) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *DiscoveryEventSearchQuery) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetSortedBy

`func (o *DiscoveryEventSearchQuery) GetSortedBy() []SortElement`

GetSortedBy returns the SortedBy field if non-nil, zero value otherwise.

### GetSortedByOk

`func (o *DiscoveryEventSearchQuery) GetSortedByOk() (*[]SortElement, bool)`

GetSortedByOk returns a tuple with the SortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortedBy

`func (o *DiscoveryEventSearchQuery) SetSortedBy(v []SortElement)`

SetSortedBy sets SortedBy field to given value.

### HasSortedBy

`func (o *DiscoveryEventSearchQuery) HasSortedBy() bool`

HasSortedBy returns a boolean if a field has been set.

### SetSortedByNil

`func (o *DiscoveryEventSearchQuery) SetSortedByNil(b bool)`

 SetSortedByNil sets the value for SortedBy to be an explicit nil

### UnsetSortedBy
`func (o *DiscoveryEventSearchQuery) UnsetSortedBy()`

UnsetSortedBy ensures that no value is present for SortedBy, not even an explicit nil
### GetWithCount

`func (o *DiscoveryEventSearchQuery) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *DiscoveryEventSearchQuery) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *DiscoveryEventSearchQuery) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *DiscoveryEventSearchQuery) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.

### SetWithCountNil

`func (o *DiscoveryEventSearchQuery) SetWithCountNil(b bool)`

 SetWithCountNil sets the value for WithCount to be an explicit nil

### UnsetWithCount
`func (o *DiscoveryEventSearchQuery) UnsetWithCount()`

UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


