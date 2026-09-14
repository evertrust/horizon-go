# OrderAPISearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int64** | Page index for pagination (0-based) | [optional] [default to 0]
**PageSize** | Pointer to **int64** | Number of items per page | [optional] [default to 20]
**SortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**WithCount** | Pointer to **bool** | Whether to include total count in results | [optional] [default to false]

## Methods

### NewOrderAPISearch

`func NewOrderAPISearch() *OrderAPISearch`

NewOrderAPISearch instantiates a new OrderAPISearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAPISearchWithDefaults

`func NewOrderAPISearchWithDefaults() *OrderAPISearch`

NewOrderAPISearchWithDefaults instantiates a new OrderAPISearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *OrderAPISearch) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *OrderAPISearch) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *OrderAPISearch) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *OrderAPISearch) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *OrderAPISearch) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OrderAPISearch) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OrderAPISearch) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OrderAPISearch) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortedBy

`func (o *OrderAPISearch) GetSortedBy() []SortElement`

GetSortedBy returns the SortedBy field if non-nil, zero value otherwise.

### GetSortedByOk

`func (o *OrderAPISearch) GetSortedByOk() (*[]SortElement, bool)`

GetSortedByOk returns a tuple with the SortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortedBy

`func (o *OrderAPISearch) SetSortedBy(v []SortElement)`

SetSortedBy sets SortedBy field to given value.

### HasSortedBy

`func (o *OrderAPISearch) HasSortedBy() bool`

HasSortedBy returns a boolean if a field has been set.

### GetWithCount

`func (o *OrderAPISearch) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *OrderAPISearch) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *OrderAPISearch) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *OrderAPISearch) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


