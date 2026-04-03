# WebRAImportRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** | The module for which to return the template. | 
**Profile** | **string** | The profile for which to return the template. | 
**Template** | [**WebRAImportRequestTemplateResponse**](WebRAImportRequestTemplateResponse.md) | The template with the constraint set on the profile | 
**Workflow** | **string** | The request workflow | 

## Methods

### NewWebRAImportRequestOnTemplateResponse

`func NewWebRAImportRequestOnTemplateResponse(module string, profile string, template WebRAImportRequestTemplateResponse, workflow string, ) *WebRAImportRequestOnTemplateResponse`

NewWebRAImportRequestOnTemplateResponse instantiates a new WebRAImportRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAImportRequestOnTemplateResponseWithDefaults

`func NewWebRAImportRequestOnTemplateResponseWithDefaults() *WebRAImportRequestOnTemplateResponse`

NewWebRAImportRequestOnTemplateResponseWithDefaults instantiates a new WebRAImportRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *WebRAImportRequestOnTemplateResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAImportRequestOnTemplateResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAImportRequestOnTemplateResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *WebRAImportRequestOnTemplateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAImportRequestOnTemplateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAImportRequestOnTemplateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTemplate

`func (o *WebRAImportRequestOnTemplateResponse) GetTemplate() WebRAImportRequestTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAImportRequestOnTemplateResponse) GetTemplateOk() (*WebRAImportRequestTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAImportRequestOnTemplateResponse) SetTemplate(v WebRAImportRequestTemplateResponse)`

SetTemplate sets Template field to given value.


### GetWorkflow

`func (o *WebRAImportRequestOnTemplateResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAImportRequestOnTemplateResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAImportRequestOnTemplateResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


