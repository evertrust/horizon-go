# DatasourceFlowTemplate200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description of the datasource | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | Display name of the datasource | [optional] 
**Inputs** | Pointer to [**[]DataSourceInput**](DataSourceInput.md) | List of inputs to use for this datasource | [optional] 
**Name** | Pointer to **string** | Name of the datasource | [optional] 
**Outputs** | Pointer to [**[]DataSourceOutput**](DataSourceOutput.md) | List of outputs for this datasource | [optional] 
**StopOnSuccess** | Pointer to **bool** | Stop the execution if this datasource&#39;s execution is successful | [optional] [default to false]
**Type** | Pointer to [**DataSourceType**](DataSourceType.md) |  | [optional] 

## Methods

### NewDatasourceFlowTemplate200ResponseInner

`func NewDatasourceFlowTemplate200ResponseInner() *DatasourceFlowTemplate200ResponseInner`

NewDatasourceFlowTemplate200ResponseInner instantiates a new DatasourceFlowTemplate200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceFlowTemplate200ResponseInnerWithDefaults

`func NewDatasourceFlowTemplate200ResponseInnerWithDefaults() *DatasourceFlowTemplate200ResponseInner`

NewDatasourceFlowTemplate200ResponseInnerWithDefaults instantiates a new DatasourceFlowTemplate200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DatasourceFlowTemplate200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasourceFlowTemplate200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasourceFlowTemplate200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatasourceFlowTemplate200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DatasourceFlowTemplate200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DatasourceFlowTemplate200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *DatasourceFlowTemplate200ResponseInner) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DatasourceFlowTemplate200ResponseInner) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DatasourceFlowTemplate200ResponseInner) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DatasourceFlowTemplate200ResponseInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DatasourceFlowTemplate200ResponseInner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DatasourceFlowTemplate200ResponseInner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetInputs

`func (o *DatasourceFlowTemplate200ResponseInner) GetInputs() []DataSourceInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *DatasourceFlowTemplate200ResponseInner) GetInputsOk() (*[]DataSourceInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *DatasourceFlowTemplate200ResponseInner) SetInputs(v []DataSourceInput)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *DatasourceFlowTemplate200ResponseInner) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *DatasourceFlowTemplate200ResponseInner) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *DatasourceFlowTemplate200ResponseInner) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetName

`func (o *DatasourceFlowTemplate200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasourceFlowTemplate200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasourceFlowTemplate200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatasourceFlowTemplate200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *DatasourceFlowTemplate200ResponseInner) GetOutputs() []DataSourceOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *DatasourceFlowTemplate200ResponseInner) GetOutputsOk() (*[]DataSourceOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *DatasourceFlowTemplate200ResponseInner) SetOutputs(v []DataSourceOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *DatasourceFlowTemplate200ResponseInner) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *DatasourceFlowTemplate200ResponseInner) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *DatasourceFlowTemplate200ResponseInner) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetStopOnSuccess

`func (o *DatasourceFlowTemplate200ResponseInner) GetStopOnSuccess() bool`

GetStopOnSuccess returns the StopOnSuccess field if non-nil, zero value otherwise.

### GetStopOnSuccessOk

`func (o *DatasourceFlowTemplate200ResponseInner) GetStopOnSuccessOk() (*bool, bool)`

GetStopOnSuccessOk returns a tuple with the StopOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnSuccess

`func (o *DatasourceFlowTemplate200ResponseInner) SetStopOnSuccess(v bool)`

SetStopOnSuccess sets StopOnSuccess field to given value.

### HasStopOnSuccess

`func (o *DatasourceFlowTemplate200ResponseInner) HasStopOnSuccess() bool`

HasStopOnSuccess returns a boolean if a field has been set.

### GetType

`func (o *DatasourceFlowTemplate200ResponseInner) GetType() DataSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasourceFlowTemplate200ResponseInner) GetTypeOk() (*DataSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasourceFlowTemplate200ResponseInner) SetType(v DataSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *DatasourceFlowTemplate200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


