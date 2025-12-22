# EstEnrollRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**Module** | Pointer to **string** | The module that will be used to process this request. For an EST request, this is always &#x60;est&#x60; | [optional] 
**Template** | Pointer to [**EstEnrollRequestTemplate**](EstEnrollRequestTemplate.md) | The user-data that will be used to generate the certificate | [optional] 
**Workflow** | Pointer to **string** | What this request will do. For an enrollment request, this is always &#x60;enroll&#x60; | [optional] 

## Methods

### NewEstEnrollRequestBase

`func NewEstEnrollRequestBase() *EstEnrollRequestBase`

NewEstEnrollRequestBase instantiates a new EstEnrollRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstEnrollRequestBaseWithDefaults

`func NewEstEnrollRequestBaseWithDefaults() *EstEnrollRequestBase`

NewEstEnrollRequestBaseWithDefaults instantiates a new EstEnrollRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *EstEnrollRequestBase) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *EstEnrollRequestBase) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *EstEnrollRequestBase) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *EstEnrollRequestBase) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *EstEnrollRequestBase) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *EstEnrollRequestBase) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetModule

`func (o *EstEnrollRequestBase) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *EstEnrollRequestBase) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *EstEnrollRequestBase) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *EstEnrollRequestBase) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetTemplate

`func (o *EstEnrollRequestBase) GetTemplate() EstEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EstEnrollRequestBase) GetTemplateOk() (*EstEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EstEnrollRequestBase) SetTemplate(v EstEnrollRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EstEnrollRequestBase) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWorkflow

`func (o *EstEnrollRequestBase) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *EstEnrollRequestBase) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *EstEnrollRequestBase) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *EstEnrollRequestBase) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


