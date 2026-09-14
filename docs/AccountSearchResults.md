# AccountSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**HasMore** | **bool** |  | 
**PageIndex** | **int64** |  | 
**PageSize** | **int64** |  | 
**Results** | [**[]AccountResponse**](AccountResponse.md) |  | 

## Methods

### NewAccountSearchResults

`func NewAccountSearchResults(hasMore bool, pageIndex int64, pageSize int64, results []AccountResponse, ) *AccountSearchResults`

NewAccountSearchResults instantiates a new AccountSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSearchResultsWithDefaults

`func NewAccountSearchResultsWithDefaults() *AccountSearchResults`

NewAccountSearchResultsWithDefaults instantiates a new AccountSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AccountSearchResults) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AccountSearchResults) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AccountSearchResults) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *AccountSearchResults) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHasMore

`func (o *AccountSearchResults) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *AccountSearchResults) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *AccountSearchResults) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetPageIndex

`func (o *AccountSearchResults) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *AccountSearchResults) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *AccountSearchResults) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *AccountSearchResults) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *AccountSearchResults) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *AccountSearchResults) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *AccountSearchResults) GetResults() []AccountResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AccountSearchResults) GetResultsOk() (*[]AccountResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AccountSearchResults) SetResults(v []AccountResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


