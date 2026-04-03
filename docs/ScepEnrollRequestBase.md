# ScepEnrollRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**Module** | Pointer to **string** | The module that will be used to process this request. For a SCEP request, this is always &#x60;scep&#x60; | [optional] 
**Template** | Pointer to [**ScepEnrollRequestTemplate**](ScepEnrollRequestTemplate.md) | The user-data that will be used to generate the challenge | [optional] 
**Workflow** | Pointer to **string** | What this request will do. For an enrollment request, this is always &#x60;enroll&#x60; | [optional] 

## Methods

### NewScepEnrollRequestBase

`func NewScepEnrollRequestBase() *ScepEnrollRequestBase`

NewScepEnrollRequestBase instantiates a new ScepEnrollRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScepEnrollRequestBaseWithDefaults

`func NewScepEnrollRequestBaseWithDefaults() *ScepEnrollRequestBase`

NewScepEnrollRequestBaseWithDefaults instantiates a new ScepEnrollRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ScepEnrollRequestBase) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ScepEnrollRequestBase) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ScepEnrollRequestBase) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ScepEnrollRequestBase) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *ScepEnrollRequestBase) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *ScepEnrollRequestBase) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetModule

`func (o *ScepEnrollRequestBase) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ScepEnrollRequestBase) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ScepEnrollRequestBase) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ScepEnrollRequestBase) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetTemplate

`func (o *ScepEnrollRequestBase) GetTemplate() ScepEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ScepEnrollRequestBase) GetTemplateOk() (*ScepEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ScepEnrollRequestBase) SetTemplate(v ScepEnrollRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ScepEnrollRequestBase) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWorkflow

`func (o *ScepEnrollRequestBase) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ScepEnrollRequestBase) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ScepEnrollRequestBase) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ScepEnrollRequestBase) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


