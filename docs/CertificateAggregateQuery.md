# CertificateAggregateQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **NullableString** | The HCQL query to use for the search, represents the way to filter certificates. If not specified, it will filter nothing | [optional] 
**GroupBy** | Pointer to **[]string** | The field that the aggregation will take place on | [optional] 
**WithCount** | Pointer to **NullableBool** | If set to &#x60;true&#x60;, the total count of certificates matching the HCQL query will be returned | [optional] 
**SortOrder** | Pointer to **NullableString** |  | [optional] 
**Limit** | Pointer to **NullableInt64** | In case of an aggregate sending a lot of different results, how many must be sent back | [optional] 
**Having** | Pointer to [**NullableHaving**](Having.md) | A condition to apply to the result. Only the aggregates results with more than 5 certificates in them can be kept for example | [optional] 

## Methods

### NewCertificateAggregateQuery

`func NewCertificateAggregateQuery() *CertificateAggregateQuery`

NewCertificateAggregateQuery instantiates a new CertificateAggregateQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAggregateQueryWithDefaults

`func NewCertificateAggregateQueryWithDefaults() *CertificateAggregateQuery`

NewCertificateAggregateQueryWithDefaults instantiates a new CertificateAggregateQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CertificateAggregateQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CertificateAggregateQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CertificateAggregateQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CertificateAggregateQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *CertificateAggregateQuery) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *CertificateAggregateQuery) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetGroupBy

`func (o *CertificateAggregateQuery) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *CertificateAggregateQuery) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *CertificateAggregateQuery) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *CertificateAggregateQuery) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### SetGroupByNil

`func (o *CertificateAggregateQuery) SetGroupByNil(b bool)`

 SetGroupByNil sets the value for GroupBy to be an explicit nil

### UnsetGroupBy
`func (o *CertificateAggregateQuery) UnsetGroupBy()`

UnsetGroupBy ensures that no value is present for GroupBy, not even an explicit nil
### GetWithCount

`func (o *CertificateAggregateQuery) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *CertificateAggregateQuery) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *CertificateAggregateQuery) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *CertificateAggregateQuery) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.

### SetWithCountNil

`func (o *CertificateAggregateQuery) SetWithCountNil(b bool)`

 SetWithCountNil sets the value for WithCount to be an explicit nil

### UnsetWithCount
`func (o *CertificateAggregateQuery) UnsetWithCount()`

UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil
### GetSortOrder

`func (o *CertificateAggregateQuery) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *CertificateAggregateQuery) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *CertificateAggregateQuery) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *CertificateAggregateQuery) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *CertificateAggregateQuery) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *CertificateAggregateQuery) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetLimit

`func (o *CertificateAggregateQuery) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CertificateAggregateQuery) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CertificateAggregateQuery) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CertificateAggregateQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *CertificateAggregateQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *CertificateAggregateQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetHaving

`func (o *CertificateAggregateQuery) GetHaving() Having`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *CertificateAggregateQuery) GetHavingOk() (*Having, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *CertificateAggregateQuery) SetHaving(v Having)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *CertificateAggregateQuery) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### SetHavingNil

`func (o *CertificateAggregateQuery) SetHavingNil(b bool)`

 SetHavingNil sets the value for Having to be an explicit nil

### UnsetHaving
`func (o *CertificateAggregateQuery) UnsetHaving()`

UnsetHaving ensures that no value is present for Having, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


