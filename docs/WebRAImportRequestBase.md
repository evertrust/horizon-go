# WebRAImportRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | Pointer to **string** | The module that will be used to process this request. For a WebRA request, this is always &#x60;webra&#x60; | [optional] 
**Workflow** | Pointer to **string** | What this request will do. For an import request, this is always &#x60;import&#x60; | [optional] 
**Template** | Pointer to [**WebRAImportRequestTemplate**](WebRAImportRequestTemplate.md) | The user-data that will be added on certificate import | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]

## Methods

### NewWebRAImportRequestBase

`func NewWebRAImportRequestBase() *WebRAImportRequestBase`

NewWebRAImportRequestBase instantiates a new WebRAImportRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAImportRequestBaseWithDefaults

`func NewWebRAImportRequestBaseWithDefaults() *WebRAImportRequestBase`

NewWebRAImportRequestBaseWithDefaults instantiates a new WebRAImportRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *WebRAImportRequestBase) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAImportRequestBase) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAImportRequestBase) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *WebRAImportRequestBase) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetWorkflow

`func (o *WebRAImportRequestBase) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAImportRequestBase) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAImportRequestBase) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WebRAImportRequestBase) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetTemplate

`func (o *WebRAImportRequestBase) GetTemplate() WebRAImportRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAImportRequestBase) GetTemplateOk() (*WebRAImportRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAImportRequestBase) SetTemplate(v WebRAImportRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRAImportRequestBase) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDryRun

`func (o *WebRAImportRequestBase) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAImportRequestBase) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAImportRequestBase) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAImportRequestBase) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAImportRequestBase) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAImportRequestBase) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


