# WebRARecoverRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | [**Module**](Module.md) | The module of the profile. | 
**Profile** | **string** | The profile on which the recovery occurred | 
**Template** | [**WebRARecoverRequestTemplate**](WebRARecoverRequestTemplate.md) | The cryptography policy applied during the recovery of a certificate | 
**Workflow** | **string** | The request workflow | 

## Methods

### NewWebRARecoverRequestOnTemplateResponse

`func NewWebRARecoverRequestOnTemplateResponse(module Module, profile string, template WebRARecoverRequestTemplate, workflow string, ) *WebRARecoverRequestOnTemplateResponse`

NewWebRARecoverRequestOnTemplateResponse instantiates a new WebRARecoverRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARecoverRequestOnTemplateResponseWithDefaults

`func NewWebRARecoverRequestOnTemplateResponseWithDefaults() *WebRARecoverRequestOnTemplateResponse`

NewWebRARecoverRequestOnTemplateResponseWithDefaults instantiates a new WebRARecoverRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetProfile

`func (o *WebRARecoverRequestOnTemplateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRARecoverRequestOnTemplateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRARecoverRequestOnTemplateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


