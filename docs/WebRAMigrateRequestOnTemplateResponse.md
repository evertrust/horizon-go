# WebRAMigrateRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** | The request workflow | 
**Module** | [**Module**](Module.md) | The module for which to return the template. | 
**Profile** | **string** | The profile for which to return the template. | 
**Template** | [**WebRAMigrateRequestTemplate**](WebRAMigrateRequestTemplate.md) | The template with the constraint set on the profile | 

## Methods

### NewWebRAMigrateRequestOnTemplateResponse

`func NewWebRAMigrateRequestOnTemplateResponse(workflow string, module Module, profile string, template WebRAMigrateRequestTemplate, ) *WebRAMigrateRequestOnTemplateResponse`

NewWebRAMigrateRequestOnTemplateResponse instantiates a new WebRAMigrateRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAMigrateRequestOnTemplateResponseWithDefaults

`func NewWebRAMigrateRequestOnTemplateResponseWithDefaults() *WebRAMigrateRequestOnTemplateResponse`

NewWebRAMigrateRequestOnTemplateResponseWithDefaults instantiates a new WebRAMigrateRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WebRAMigrateRequestOnTemplateResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAMigrateRequestOnTemplateResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAMigrateRequestOnTemplateResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetModule

`func (o *WebRAMigrateRequestOnTemplateResponse) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAMigrateRequestOnTemplateResponse) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAMigrateRequestOnTemplateResponse) SetModule(v Module)`

SetModule sets Module field to given value.


### GetProfile

`func (o *WebRAMigrateRequestOnTemplateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAMigrateRequestOnTemplateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAMigrateRequestOnTemplateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTemplate

`func (o *WebRAMigrateRequestOnTemplateResponse) GetTemplate() WebRAMigrateRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAMigrateRequestOnTemplateResponse) GetTemplateOk() (*WebRAMigrateRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAMigrateRequestOnTemplateResponse) SetTemplate(v WebRAMigrateRequestTemplate)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


