# CertificateSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to **[]string** | The fields to be returned by the search. If this parameter is not specified, everything is returned by default. If this parameter is equal to an empty array, only the &#x60;_id&#x60; field is returned.  | [optional] 
**PageIndex** | Pointer to **NullableInt64** | The index of the page to retrieve | [optional] [default to 1]
**PageSize** | Pointer to **NullableInt64** | The maximum number of items to retrieve for one page | [optional] [default to 50]
**Query** | Pointer to **NullableString** | The HCQL query to use for the search, represents the way to filter certificates. If not specified, it will filter nothing | [optional] 
**SortedBy** | Pointer to [**[]SortElement**](SortElement.md) | The way to sort the search results.  | [optional] 
**WithCount** | Pointer to **NullableBool** | If set to &#x60;true&#x60;, the total count of certificates matching the HCQL query will be returned.  | [optional] [default to false]

## Methods

### NewCertificateSearchQuery

`func NewCertificateSearchQuery() *CertificateSearchQuery`

NewCertificateSearchQuery instantiates a new CertificateSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSearchQueryWithDefaults

`func NewCertificateSearchQueryWithDefaults() *CertificateSearchQuery`

NewCertificateSearchQueryWithDefaults instantiates a new CertificateSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *CertificateSearchQuery) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CertificateSearchQuery) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CertificateSearchQuery) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CertificateSearchQuery) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *CertificateSearchQuery) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *CertificateSearchQuery) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetPageIndex

`func (o *CertificateSearchQuery) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *CertificateSearchQuery) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *CertificateSearchQuery) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *CertificateSearchQuery) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### SetPageIndexNil

`func (o *CertificateSearchQuery) SetPageIndexNil(b bool)`

 SetPageIndexNil sets the value for PageIndex to be an explicit nil

### UnsetPageIndex
`func (o *CertificateSearchQuery) UnsetPageIndex()`

UnsetPageIndex ensures that no value is present for PageIndex, not even an explicit nil
### GetPageSize

`func (o *CertificateSearchQuery) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CertificateSearchQuery) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CertificateSearchQuery) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CertificateSearchQuery) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *CertificateSearchQuery) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *CertificateSearchQuery) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetQuery

`func (o *CertificateSearchQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CertificateSearchQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CertificateSearchQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CertificateSearchQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *CertificateSearchQuery) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *CertificateSearchQuery) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetSortedBy

`func (o *CertificateSearchQuery) GetSortedBy() []SortElement`

GetSortedBy returns the SortedBy field if non-nil, zero value otherwise.

### GetSortedByOk

`func (o *CertificateSearchQuery) GetSortedByOk() (*[]SortElement, bool)`

GetSortedByOk returns a tuple with the SortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortedBy

`func (o *CertificateSearchQuery) SetSortedBy(v []SortElement)`

SetSortedBy sets SortedBy field to given value.

### HasSortedBy

`func (o *CertificateSearchQuery) HasSortedBy() bool`

HasSortedBy returns a boolean if a field has been set.

### SetSortedByNil

`func (o *CertificateSearchQuery) SetSortedByNil(b bool)`

 SetSortedByNil sets the value for SortedBy to be an explicit nil

### UnsetSortedBy
`func (o *CertificateSearchQuery) UnsetSortedBy()`

UnsetSortedBy ensures that no value is present for SortedBy, not even an explicit nil
### GetWithCount

`func (o *CertificateSearchQuery) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *CertificateSearchQuery) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *CertificateSearchQuery) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *CertificateSearchQuery) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.

### SetWithCountNil

`func (o *CertificateSearchQuery) SetWithCountNil(b bool)`

 SetWithCountNil sets the value for WithCount to be an explicit nil

### UnsetWithCount
`func (o *CertificateSearchQuery) UnsetWithCount()`

UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


