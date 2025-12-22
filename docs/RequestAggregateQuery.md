# RequestAggregateQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupBy** | Pointer to **[]string** | The field that the aggregation will take place on | [optional] 
**Having** | Pointer to [**NullableHaving**](Having.md) | A condition to apply to the result. Only the aggregates results with more than 5 requests in them can be kept for example | [optional] 
**Limit** | Pointer to **NullableInt64** | In case of an aggregate sending a lot of different results, how many must be sent back | [optional] 
**Query** | Pointer to **NullableString** | The HRQL query to use for the search, represents the way to filter requests. If not specified, it will filter nothing | [optional] 
**Scope** | Pointer to **NullableString** | The scope of the aggregate. &#x60;manage&#x60; only aggregates among requests that the currently logged in user has the rights to manage. &#x60;search&#x60; aggregates among all visible requests to the logged in user. &#x60;self&#x60; aggregates among requests that the currently logged in user or its team has issued | [optional] 
**SortOrder** | Pointer to **NullableString** |  | [optional] 
**WithCount** | Pointer to **NullableBool** | If set to &#x60;true&#x60;, the total count of requests matching the HRQL query will be returned | [optional] 

## Methods

### NewRequestAggregateQuery

`func NewRequestAggregateQuery() *RequestAggregateQuery`

NewRequestAggregateQuery instantiates a new RequestAggregateQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestAggregateQueryWithDefaults

`func NewRequestAggregateQueryWithDefaults() *RequestAggregateQuery`

NewRequestAggregateQueryWithDefaults instantiates a new RequestAggregateQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupBy

`func (o *RequestAggregateQuery) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *RequestAggregateQuery) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *RequestAggregateQuery) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *RequestAggregateQuery) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### SetGroupByNil

`func (o *RequestAggregateQuery) SetGroupByNil(b bool)`

 SetGroupByNil sets the value for GroupBy to be an explicit nil

### UnsetGroupBy
`func (o *RequestAggregateQuery) UnsetGroupBy()`

UnsetGroupBy ensures that no value is present for GroupBy, not even an explicit nil
### GetHaving

`func (o *RequestAggregateQuery) GetHaving() Having`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *RequestAggregateQuery) GetHavingOk() (*Having, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *RequestAggregateQuery) SetHaving(v Having)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *RequestAggregateQuery) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### SetHavingNil

`func (o *RequestAggregateQuery) SetHavingNil(b bool)`

 SetHavingNil sets the value for Having to be an explicit nil

### UnsetHaving
`func (o *RequestAggregateQuery) UnsetHaving()`

UnsetHaving ensures that no value is present for Having, not even an explicit nil
### GetLimit

`func (o *RequestAggregateQuery) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RequestAggregateQuery) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RequestAggregateQuery) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RequestAggregateQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *RequestAggregateQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *RequestAggregateQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetQuery

`func (o *RequestAggregateQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RequestAggregateQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RequestAggregateQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *RequestAggregateQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *RequestAggregateQuery) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *RequestAggregateQuery) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetScope

`func (o *RequestAggregateQuery) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RequestAggregateQuery) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RequestAggregateQuery) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RequestAggregateQuery) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *RequestAggregateQuery) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *RequestAggregateQuery) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetSortOrder

`func (o *RequestAggregateQuery) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *RequestAggregateQuery) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *RequestAggregateQuery) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *RequestAggregateQuery) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *RequestAggregateQuery) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *RequestAggregateQuery) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetWithCount

`func (o *RequestAggregateQuery) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *RequestAggregateQuery) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *RequestAggregateQuery) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *RequestAggregateQuery) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.

### SetWithCountNil

`func (o *RequestAggregateQuery) SetWithCountNil(b bool)`

 SetWithCountNil sets the value for WithCount to be an explicit nil

### UnsetWithCount
`func (o *RequestAggregateQuery) UnsetWithCount()`

UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


