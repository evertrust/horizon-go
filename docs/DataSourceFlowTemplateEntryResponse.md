# DataSourceFlowTemplateEntryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description of the datasource | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | Display name of the datasource | [optional] 
**Inputs** | Pointer to [**[]DataSourceInput**](DataSourceInput.md) | List of inputs to use for this datasource | [optional] 
**Mandatory** | Pointer to **bool** | If true, the flow will stop with an error if this datasource does not return any result | [optional] [default to false]
**Name** | Pointer to **string** | Name of the datasource | [optional] 
**Outputs** | Pointer to [**[]DataSourceOutput**](DataSourceOutput.md) | List of outputs for this datasource | [optional] 
**StopOnSuccess** | Pointer to **bool** | Stop the execution if this datasource&#39;s execution is successful | [optional] [default to false]
**Type** | Pointer to [**DataSourceType**](DataSourceType.md) |  | [optional] 

## Methods

### NewDataSourceFlowTemplateEntryResponse

`func NewDataSourceFlowTemplateEntryResponse() *DataSourceFlowTemplateEntryResponse`

NewDataSourceFlowTemplateEntryResponse instantiates a new DataSourceFlowTemplateEntryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceFlowTemplateEntryResponseWithDefaults

`func NewDataSourceFlowTemplateEntryResponseWithDefaults() *DataSourceFlowTemplateEntryResponse`

NewDataSourceFlowTemplateEntryResponseWithDefaults instantiates a new DataSourceFlowTemplateEntryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DataSourceFlowTemplateEntryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataSourceFlowTemplateEntryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataSourceFlowTemplateEntryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataSourceFlowTemplateEntryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DataSourceFlowTemplateEntryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DataSourceFlowTemplateEntryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *DataSourceFlowTemplateEntryResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DataSourceFlowTemplateEntryResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DataSourceFlowTemplateEntryResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DataSourceFlowTemplateEntryResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DataSourceFlowTemplateEntryResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DataSourceFlowTemplateEntryResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetInputs

`func (o *DataSourceFlowTemplateEntryResponse) GetInputs() []DataSourceInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *DataSourceFlowTemplateEntryResponse) GetInputsOk() (*[]DataSourceInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *DataSourceFlowTemplateEntryResponse) SetInputs(v []DataSourceInput)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *DataSourceFlowTemplateEntryResponse) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *DataSourceFlowTemplateEntryResponse) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *DataSourceFlowTemplateEntryResponse) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetMandatory

`func (o *DataSourceFlowTemplateEntryResponse) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *DataSourceFlowTemplateEntryResponse) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *DataSourceFlowTemplateEntryResponse) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *DataSourceFlowTemplateEntryResponse) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### GetName

`func (o *DataSourceFlowTemplateEntryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceFlowTemplateEntryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceFlowTemplateEntryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataSourceFlowTemplateEntryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *DataSourceFlowTemplateEntryResponse) GetOutputs() []DataSourceOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *DataSourceFlowTemplateEntryResponse) GetOutputsOk() (*[]DataSourceOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *DataSourceFlowTemplateEntryResponse) SetOutputs(v []DataSourceOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *DataSourceFlowTemplateEntryResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *DataSourceFlowTemplateEntryResponse) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *DataSourceFlowTemplateEntryResponse) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetStopOnSuccess

`func (o *DataSourceFlowTemplateEntryResponse) GetStopOnSuccess() bool`

GetStopOnSuccess returns the StopOnSuccess field if non-nil, zero value otherwise.

### GetStopOnSuccessOk

`func (o *DataSourceFlowTemplateEntryResponse) GetStopOnSuccessOk() (*bool, bool)`

GetStopOnSuccessOk returns a tuple with the StopOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnSuccess

`func (o *DataSourceFlowTemplateEntryResponse) SetStopOnSuccess(v bool)`

SetStopOnSuccess sets StopOnSuccess field to given value.

### HasStopOnSuccess

`func (o *DataSourceFlowTemplateEntryResponse) HasStopOnSuccess() bool`

HasStopOnSuccess returns a boolean if a field has been set.

### GetType

`func (o *DataSourceFlowTemplateEntryResponse) GetType() DataSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataSourceFlowTemplateEntryResponse) GetTypeOk() (*DataSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataSourceFlowTemplateEntryResponse) SetType(v DataSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *DataSourceFlowTemplateEntryResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


