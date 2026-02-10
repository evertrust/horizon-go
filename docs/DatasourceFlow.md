# DatasourceFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DsFlow** | Pointer to [**[]DataSourceFlowEntry**](DataSourceFlowEntry.md) | Representation of a datasource execution flow | [optional] 
**Context** | Pointer to [**[]MapEntry**](MapEntry.md) | Input values for the flow | [optional] 

## Methods

### NewDatasourceFlow

`func NewDatasourceFlow() *DatasourceFlow`

NewDatasourceFlow instantiates a new DatasourceFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceFlowWithDefaults

`func NewDatasourceFlowWithDefaults() *DatasourceFlow`

NewDatasourceFlowWithDefaults instantiates a new DatasourceFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDsFlow

`func (o *DatasourceFlow) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *DatasourceFlow) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *DatasourceFlow) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *DatasourceFlow) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### GetContext

`func (o *DatasourceFlow) GetContext() []MapEntry`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DatasourceFlow) GetContextOk() (*[]MapEntry, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DatasourceFlow) SetContext(v []MapEntry)`

SetContext sets Context field to given value.

### HasContext

`func (o *DatasourceFlow) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *DatasourceFlow) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *DatasourceFlow) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


