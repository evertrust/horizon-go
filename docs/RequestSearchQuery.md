# RequestSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **NullableString** | The HRQL query to use for the search, represents the way to filter requests. Filters nothing if not specified | [optional] 
**Fields** | Pointer to **[]string** | The fields to be returned by the search. If this parameter is not specified, everything is returned by default. If this parameter is equal to an empty array, only the &#x60;_id&#x60; field is returned | [optional] 
**SortedBy** | Pointer to [**[]SortElement**](SortElement.md) | The way to sort the search results | [optional] 
**PageIndex** | Pointer to **NullableInt64** | The index of the page to retrieve | [optional] [default to 1]
**PageSize** | Pointer to **NullableInt64** | The maximum number of items to retrieve for one page | [optional] [default to 50]
**WithCount** | Pointer to **NullableBool** | Whether to return the total count of requests matching the HRQL query | [optional] [default to false]
**Scope** | Pointer to **NullableString** | The scope of the search. &#x60;manage&#x60; only searches among requests that the currently logged in user has the rights to manage. &#x60;search&#x60; searches among all visible requests to the logged in user. &#x60;self&#x60; searches among requests that the currently logged in user or its team has issued | [optional] 

## Methods

### NewRequestSearchQuery

`func NewRequestSearchQuery() *RequestSearchQuery`

NewRequestSearchQuery instantiates a new RequestSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSearchQueryWithDefaults

`func NewRequestSearchQueryWithDefaults() *RequestSearchQuery`

NewRequestSearchQueryWithDefaults instantiates a new RequestSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *RequestSearchQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RequestSearchQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RequestSearchQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *RequestSearchQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *RequestSearchQuery) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *RequestSearchQuery) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetFields

`func (o *RequestSearchQuery) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RequestSearchQuery) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RequestSearchQuery) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RequestSearchQuery) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *RequestSearchQuery) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *RequestSearchQuery) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetSortedBy

`func (o *RequestSearchQuery) GetSortedBy() []SortElement`

GetSortedBy returns the SortedBy field if non-nil, zero value otherwise.

### GetSortedByOk

`func (o *RequestSearchQuery) GetSortedByOk() (*[]SortElement, bool)`

GetSortedByOk returns a tuple with the SortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortedBy

`func (o *RequestSearchQuery) SetSortedBy(v []SortElement)`

SetSortedBy sets SortedBy field to given value.

### HasSortedBy

`func (o *RequestSearchQuery) HasSortedBy() bool`

HasSortedBy returns a boolean if a field has been set.

### SetSortedByNil

`func (o *RequestSearchQuery) SetSortedByNil(b bool)`

 SetSortedByNil sets the value for SortedBy to be an explicit nil

### UnsetSortedBy
`func (o *RequestSearchQuery) UnsetSortedBy()`

UnsetSortedBy ensures that no value is present for SortedBy, not even an explicit nil
### GetPageIndex

`func (o *RequestSearchQuery) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *RequestSearchQuery) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *RequestSearchQuery) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *RequestSearchQuery) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### SetPageIndexNil

`func (o *RequestSearchQuery) SetPageIndexNil(b bool)`

 SetPageIndexNil sets the value for PageIndex to be an explicit nil

### UnsetPageIndex
`func (o *RequestSearchQuery) UnsetPageIndex()`

UnsetPageIndex ensures that no value is present for PageIndex, not even an explicit nil
### GetPageSize

`func (o *RequestSearchQuery) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *RequestSearchQuery) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *RequestSearchQuery) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *RequestSearchQuery) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *RequestSearchQuery) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *RequestSearchQuery) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetWithCount

`func (o *RequestSearchQuery) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *RequestSearchQuery) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *RequestSearchQuery) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *RequestSearchQuery) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.

### SetWithCountNil

`func (o *RequestSearchQuery) SetWithCountNil(b bool)`

 SetWithCountNil sets the value for WithCount to be an explicit nil

### UnsetWithCount
`func (o *RequestSearchQuery) UnsetWithCount()`

UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil
### GetScope

`func (o *RequestSearchQuery) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RequestSearchQuery) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RequestSearchQuery) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RequestSearchQuery) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *RequestSearchQuery) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *RequestSearchQuery) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


