# DataSourceFlowEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ds** | **string** | Name of the datasource to execute for this step | 
**Inputs** | Pointer to [**[]DataSourceInput**](DataSourceInput.md) | List of inputs to use for this datasource | [optional] 
**StopOnSuccess** | Pointer to **bool** | Stop the flow if this datasource is successfully executed | [optional] [default to false]

## Methods

### NewDataSourceFlowEntry

`func NewDataSourceFlowEntry(ds string, ) *DataSourceFlowEntry`

NewDataSourceFlowEntry instantiates a new DataSourceFlowEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceFlowEntryWithDefaults

`func NewDataSourceFlowEntryWithDefaults() *DataSourceFlowEntry`

NewDataSourceFlowEntryWithDefaults instantiates a new DataSourceFlowEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDs

`func (o *DataSourceFlowEntry) GetDs() string`

GetDs returns the Ds field if non-nil, zero value otherwise.

### GetDsOk

`func (o *DataSourceFlowEntry) GetDsOk() (*string, bool)`

GetDsOk returns a tuple with the Ds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDs

`func (o *DataSourceFlowEntry) SetDs(v string)`

SetDs sets Ds field to given value.


### GetInputs

`func (o *DataSourceFlowEntry) GetInputs() []DataSourceInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *DataSourceFlowEntry) GetInputsOk() (*[]DataSourceInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *DataSourceFlowEntry) SetInputs(v []DataSourceInput)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *DataSourceFlowEntry) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *DataSourceFlowEntry) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *DataSourceFlowEntry) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetStopOnSuccess

`func (o *DataSourceFlowEntry) GetStopOnSuccess() bool`

GetStopOnSuccess returns the StopOnSuccess field if non-nil, zero value otherwise.

### GetStopOnSuccessOk

`func (o *DataSourceFlowEntry) GetStopOnSuccessOk() (*bool, bool)`

GetStopOnSuccessOk returns a tuple with the StopOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnSuccess

`func (o *DataSourceFlowEntry) SetStopOnSuccess(v bool)`

SetStopOnSuccess sets StopOnSuccess field to given value.

### HasStopOnSuccess

`func (o *DataSourceFlowEntry) HasStopOnSuccess() bool`

HasStopOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


