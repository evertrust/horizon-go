# WebRARecoverRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** | The request workflow | 
**Module** | [**Module**](Module.md) | The module of the profile. | 
**Template** | [**WebRARecoverRequestTemplate**](WebRARecoverRequestTemplate.md) | The cryptography policy applied during the recovery of a certificate | 

## Methods

### NewWebRARecoverRequestOnTemplateResponse

`func NewWebRARecoverRequestOnTemplateResponse(workflow string, module Module, template WebRARecoverRequestTemplate, ) *WebRARecoverRequestOnTemplateResponse`

NewWebRARecoverRequestOnTemplateResponse instantiates a new WebRARecoverRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARecoverRequestOnTemplateResponseWithDefaults

`func NewWebRARecoverRequestOnTemplateResponseWithDefaults() *WebRARecoverRequestOnTemplateResponse`

NewWebRARecoverRequestOnTemplateResponseWithDefaults instantiates a new WebRARecoverRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WebRARecoverRequestOnTemplateResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARecoverRequestOnTemplateResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARecoverRequestOnTemplateResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetModule

`func (o *WebRARecoverRequestOnTemplateResponse) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRARecoverRequestOnTemplateResponse) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRARecoverRequestOnTemplateResponse) SetModule(v Module)`

SetModule sets Module field to given value.


### GetTemplate

`func (o *WebRARecoverRequestOnTemplateResponse) GetTemplate() WebRARecoverRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRARecoverRequestOnTemplateResponse) GetTemplateOk() (*WebRARecoverRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRARecoverRequestOnTemplateResponse) SetTemplate(v WebRARecoverRequestTemplate)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


