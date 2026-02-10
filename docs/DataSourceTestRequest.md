# DataSourceTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**[]MapEntry**](MapEntry.md) | Context to interpret dynamic values from | [optional] 
**Ds** | [**DataSourceTestRequestDs**](DataSourceTestRequestDs.md) |  | 

## Methods

### NewDataSourceTestRequest

`func NewDataSourceTestRequest(ds DataSourceTestRequestDs, ) *DataSourceTestRequest`

NewDataSourceTestRequest instantiates a new DataSourceTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceTestRequestWithDefaults

`func NewDataSourceTestRequestWithDefaults() *DataSourceTestRequest`

NewDataSourceTestRequestWithDefaults instantiates a new DataSourceTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *DataSourceTestRequest) GetContext() []MapEntry`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DataSourceTestRequest) GetContextOk() (*[]MapEntry, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DataSourceTestRequest) SetContext(v []MapEntry)`

SetContext sets Context field to given value.

### HasContext

`func (o *DataSourceTestRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *DataSourceTestRequest) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *DataSourceTestRequest) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetDs

`func (o *DataSourceTestRequest) GetDs() DataSourceTestRequestDs`

GetDs returns the Ds field if non-nil, zero value otherwise.

### GetDsOk

`func (o *DataSourceTestRequest) GetDsOk() (*DataSourceTestRequestDs, bool)`

GetDsOk returns a tuple with the Ds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDs

`func (o *DataSourceTestRequest) SetDs(v DataSourceTestRequestDs)`

SetDs sets Ds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


