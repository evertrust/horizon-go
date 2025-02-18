# WebRAMigrateRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | Pointer to **string** | The module that will be used to process this request. For a WebRA request, this is always &#x60;webra&#x60; | [optional] 
**Workflow** | Pointer to **string** | What this request will do. For a migration request, this is always &#x60;migrate&#x60; | [optional] 
**Template** | Pointer to [**WebRAMigrateRequestTemplate**](WebRAMigrateRequestTemplate.md) | The user-data that will be used to migrate the certificate | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in a migration | [optional] [default to false]

## Methods

### NewWebRAMigrateRequestBase

`func NewWebRAMigrateRequestBase() *WebRAMigrateRequestBase`

NewWebRAMigrateRequestBase instantiates a new WebRAMigrateRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAMigrateRequestBaseWithDefaults

`func NewWebRAMigrateRequestBaseWithDefaults() *WebRAMigrateRequestBase`

NewWebRAMigrateRequestBaseWithDefaults instantiates a new WebRAMigrateRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *WebRAMigrateRequestBase) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAMigrateRequestBase) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAMigrateRequestBase) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *WebRAMigrateRequestBase) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetWorkflow

`func (o *WebRAMigrateRequestBase) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAMigrateRequestBase) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAMigrateRequestBase) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WebRAMigrateRequestBase) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetTemplate

`func (o *WebRAMigrateRequestBase) GetTemplate() WebRAMigrateRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAMigrateRequestBase) GetTemplateOk() (*WebRAMigrateRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAMigrateRequestBase) SetTemplate(v WebRAMigrateRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRAMigrateRequestBase) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDryRun

`func (o *WebRAMigrateRequestBase) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAMigrateRequestBase) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAMigrateRequestBase) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAMigrateRequestBase) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAMigrateRequestBase) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAMigrateRequestBase) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


