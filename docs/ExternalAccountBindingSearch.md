# ExternalAccountBindingSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int64** |  | [optional] 
**PageSize** | Pointer to **int64** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**SortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**WithCount** | Pointer to **bool** |  | [optional] 

## Methods

### NewExternalAccountBindingSearch

`func NewExternalAccountBindingSearch() *ExternalAccountBindingSearch`

NewExternalAccountBindingSearch instantiates a new ExternalAccountBindingSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountBindingSearchWithDefaults

`func NewExternalAccountBindingSearchWithDefaults() *ExternalAccountBindingSearch`

NewExternalAccountBindingSearchWithDefaults instantiates a new ExternalAccountBindingSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *ExternalAccountBindingSearch) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *ExternalAccountBindingSearch) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *ExternalAccountBindingSearch) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *ExternalAccountBindingSearch) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *ExternalAccountBindingSearch) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ExternalAccountBindingSearch) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ExternalAccountBindingSearch) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ExternalAccountBindingSearch) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetQuery

`func (o *ExternalAccountBindingSearch) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ExternalAccountBindingSearch) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ExternalAccountBindingSearch) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ExternalAccountBindingSearch) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSortedBy

`func (o *ExternalAccountBindingSearch) GetSortedBy() []SortElement`

GetSortedBy returns the SortedBy field if non-nil, zero value otherwise.

### GetSortedByOk

`func (o *ExternalAccountBindingSearch) GetSortedByOk() (*[]SortElement, bool)`

GetSortedByOk returns a tuple with the SortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortedBy

`func (o *ExternalAccountBindingSearch) SetSortedBy(v []SortElement)`

SetSortedBy sets SortedBy field to given value.

### HasSortedBy

`func (o *ExternalAccountBindingSearch) HasSortedBy() bool`

HasSortedBy returns a boolean if a field has been set.

### GetWithCount

`func (o *ExternalAccountBindingSearch) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *ExternalAccountBindingSearch) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *ExternalAccountBindingSearch) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *ExternalAccountBindingSearch) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


