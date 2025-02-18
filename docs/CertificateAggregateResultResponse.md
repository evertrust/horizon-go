# CertificateAggregateResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CertificateAggregateResultResponseItemsInner**](CertificateAggregateResultResponseItemsInner.md) | All the groups in this aggregate | 
**Count** | Pointer to **NullableInt64** | The total number of certificates matching the query | [optional] 

## Methods

### NewCertificateAggregateResultResponse

`func NewCertificateAggregateResultResponse(items []CertificateAggregateResultResponseItemsInner, ) *CertificateAggregateResultResponse`

NewCertificateAggregateResultResponse instantiates a new CertificateAggregateResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAggregateResultResponseWithDefaults

`func NewCertificateAggregateResultResponseWithDefaults() *CertificateAggregateResultResponse`

NewCertificateAggregateResultResponseWithDefaults instantiates a new CertificateAggregateResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CertificateAggregateResultResponse) GetItems() []CertificateAggregateResultResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CertificateAggregateResultResponse) GetItemsOk() (*[]CertificateAggregateResultResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CertificateAggregateResultResponse) SetItems(v []CertificateAggregateResultResponseItemsInner)`

SetItems sets Items field to given value.


### GetCount

`func (o *CertificateAggregateResultResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CertificateAggregateResultResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CertificateAggregateResultResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CertificateAggregateResultResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *CertificateAggregateResultResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *CertificateAggregateResultResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


