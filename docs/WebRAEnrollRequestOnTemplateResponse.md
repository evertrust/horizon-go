# WebRAEnrollRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** | The module for which to return the template. | 
**Profile** | **string** | The profile for which to return the template. | 
**Template** | [**WebRAEnrollRequestTemplateResponse**](WebRAEnrollRequestTemplateResponse.md) | The template with the constraint set on the profile | 
**Workflow** | **string** | The request workflow | 

## Methods

### NewWebRAEnrollRequestOnTemplateResponse

`func NewWebRAEnrollRequestOnTemplateResponse(module string, profile string, template WebRAEnrollRequestTemplateResponse, workflow string, ) *WebRAEnrollRequestOnTemplateResponse`

NewWebRAEnrollRequestOnTemplateResponse instantiates a new WebRAEnrollRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAEnrollRequestOnTemplateResponseWithDefaults

`func NewWebRAEnrollRequestOnTemplateResponseWithDefaults() *WebRAEnrollRequestOnTemplateResponse`

NewWebRAEnrollRequestOnTemplateResponseWithDefaults instantiates a new WebRAEnrollRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *WebRAEnrollRequestOnTemplateResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAEnrollRequestOnTemplateResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAEnrollRequestOnTemplateResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *WebRAEnrollRequestOnTemplateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAEnrollRequestOnTemplateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAEnrollRequestOnTemplateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTemplate

`func (o *WebRAEnrollRequestOnTemplateResponse) GetTemplate() WebRAEnrollRequestTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAEnrollRequestOnTemplateResponse) GetTemplateOk() (*WebRAEnrollRequestTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAEnrollRequestOnTemplateResponse) SetTemplate(v WebRAEnrollRequestTemplateResponse)`

SetTemplate sets Template field to given value.


### GetWorkflow

`func (o *WebRAEnrollRequestOnTemplateResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAEnrollRequestOnTemplateResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAEnrollRequestOnTemplateResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


