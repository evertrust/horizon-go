# AccountSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int64** |  | [optional] 
**PageSize** | Pointer to **int64** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**SortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**WithCount** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountSearch

`func NewAccountSearch() *AccountSearch`

NewAccountSearch instantiates a new AccountSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSearchWithDefaults

`func NewAccountSearchWithDefaults() *AccountSearch`

NewAccountSearchWithDefaults instantiates a new AccountSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *AccountSearch) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *AccountSearch) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *AccountSearch) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *AccountSearch) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *AccountSearch) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *AccountSearch) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *AccountSearch) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *AccountSearch) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetQuery

`func (o *AccountSearch) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AccountSearch) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AccountSearch) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AccountSearch) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSortedBy

`func (o *AccountSearch) GetSortedBy() []SortElement`

GetSortedBy returns the SortedBy field if non-nil, zero value otherwise.

### GetSortedByOk

`func (o *AccountSearch) GetSortedByOk() (*[]SortElement, bool)`

GetSortedByOk returns a tuple with the SortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortedBy

`func (o *AccountSearch) SetSortedBy(v []SortElement)`

SetSortedBy sets SortedBy field to given value.

### HasSortedBy

`func (o *AccountSearch) HasSortedBy() bool`

HasSortedBy returns a boolean if a field has been set.

### GetWithCount

`func (o *AccountSearch) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *AccountSearch) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *AccountSearch) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *AccountSearch) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


