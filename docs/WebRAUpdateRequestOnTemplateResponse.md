# WebRAUpdateRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** | The request workflow | 
**Module** | [**Module**](Module.md) | The module of the certificate to update. | 
**Profile** | **string** | The profile for which to return the template. | 
**Template** | [**WebRAUpdateRequestTemplate**](WebRAUpdateRequestTemplate.md) | The template with the constraint set on the profile | 

## Methods

### NewWebRAUpdateRequestOnTemplateResponse

`func NewWebRAUpdateRequestOnTemplateResponse(workflow string, module Module, profile string, template WebRAUpdateRequestTemplate, ) *WebRAUpdateRequestOnTemplateResponse`

NewWebRAUpdateRequestOnTemplateResponse instantiates a new WebRAUpdateRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAUpdateRequestOnTemplateResponseWithDefaults

`func NewWebRAUpdateRequestOnTemplateResponseWithDefaults() *WebRAUpdateRequestOnTemplateResponse`

NewWebRAUpdateRequestOnTemplateResponseWithDefaults instantiates a new WebRAUpdateRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WebRAUpdateRequestOnTemplateResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAUpdateRequestOnTemplateResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAUpdateRequestOnTemplateResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetModule

`func (o *WebRAUpdateRequestOnTemplateResponse) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAUpdateRequestOnTemplateResponse) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAUpdateRequestOnTemplateResponse) SetModule(v Module)`

SetModule sets Module field to given value.


### GetProfile

`func (o *WebRAUpdateRequestOnTemplateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAUpdateRequestOnTemplateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAUpdateRequestOnTemplateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTemplate

`func (o *WebRAUpdateRequestOnTemplateResponse) GetTemplate() WebRAUpdateRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAUpdateRequestOnTemplateResponse) GetTemplateOk() (*WebRAUpdateRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAUpdateRequestOnTemplateResponse) SetTemplate(v WebRAUpdateRequestTemplate)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


