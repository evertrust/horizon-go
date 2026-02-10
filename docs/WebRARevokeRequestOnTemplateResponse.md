# WebRARevokeRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** | The request workflow | 
**Module** | [**Module**](Module.md) | The module on which the certificate will be revoked. | 
**Profile** | **string** | The profile for which to return the template. | 
**Template** | [**WebRARevokeRequestTemplate**](WebRARevokeRequestTemplate.md) | The reason for revoking the certificate | 

## Methods

### NewWebRARevokeRequestOnTemplateResponse

`func NewWebRARevokeRequestOnTemplateResponse(workflow string, module Module, profile string, template WebRARevokeRequestTemplate, ) *WebRARevokeRequestOnTemplateResponse`

NewWebRARevokeRequestOnTemplateResponse instantiates a new WebRARevokeRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARevokeRequestOnTemplateResponseWithDefaults

`func NewWebRARevokeRequestOnTemplateResponseWithDefaults() *WebRARevokeRequestOnTemplateResponse`

NewWebRARevokeRequestOnTemplateResponseWithDefaults instantiates a new WebRARevokeRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WebRARevokeRequestOnTemplateResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARevokeRequestOnTemplateResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARevokeRequestOnTemplateResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetModule

`func (o *WebRARevokeRequestOnTemplateResponse) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRARevokeRequestOnTemplateResponse) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRARevokeRequestOnTemplateResponse) SetModule(v Module)`

SetModule sets Module field to given value.


### GetProfile

`func (o *WebRARevokeRequestOnTemplateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRARevokeRequestOnTemplateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRARevokeRequestOnTemplateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTemplate

`func (o *WebRARevokeRequestOnTemplateResponse) GetTemplate() WebRARevokeRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRARevokeRequestOnTemplateResponse) GetTemplateOk() (*WebRARevokeRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRARevokeRequestOnTemplateResponse) SetTemplate(v WebRARevokeRequestTemplate)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


