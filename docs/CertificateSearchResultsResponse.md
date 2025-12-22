# CertificateSearchResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt64** | If &#x60;withCount&#x60; was set to true in the query payload, represents the total number of certificates that were retrieved for that query  | [optional] 
**HasMore** | **bool** | Indicates whether the response represents the last page of results (if set to &#x60;false&#x60;) or not (if set to &#x60;true&#x60;)  | 
**PageIndex** | **int64** | The index of the page that has been retrieved | 
**PageSize** | **int64** | The size of the page that has been retrieved | 
**Results** | [**[]CertificateSearchResult**](CertificateSearchResult.md) | List of certificates that matched the search criteria | 

## Methods

### NewCertificateSearchResultsResponse

`func NewCertificateSearchResultsResponse(hasMore bool, pageIndex int64, pageSize int64, results []CertificateSearchResult, ) *CertificateSearchResultsResponse`

NewCertificateSearchResultsResponse instantiates a new CertificateSearchResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSearchResultsResponseWithDefaults

`func NewCertificateSearchResultsResponseWithDefaults() *CertificateSearchResultsResponse`

NewCertificateSearchResultsResponseWithDefaults instantiates a new CertificateSearchResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CertificateSearchResultsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CertificateSearchResultsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CertificateSearchResultsResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CertificateSearchResultsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *CertificateSearchResultsResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *CertificateSearchResultsResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetHasMore

`func (o *CertificateSearchResultsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *CertificateSearchResultsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *CertificateSearchResultsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetPageIndex

`func (o *CertificateSearchResultsResponse) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *CertificateSearchResultsResponse) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *CertificateSearchResultsResponse) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *CertificateSearchResultsResponse) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CertificateSearchResultsResponse) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CertificateSearchResultsResponse) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *CertificateSearchResultsResponse) GetResults() []CertificateSearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CertificateSearchResultsResponse) GetResultsOk() (*[]CertificateSearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CertificateSearchResultsResponse) SetResults(v []CertificateSearchResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


