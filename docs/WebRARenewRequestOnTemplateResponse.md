# WebRARenewRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** | The module for which to return the template. | 
**Profile** | **string** | The profile on which the renewal occurred | 
**Template** | [**WebRARenewRequestTemplateResponse**](WebRARenewRequestTemplateResponse.md) | The cryptography policy applied during the renewal of a certificate | 
**Workflow** | **string** | The request workflow | 

## Methods

### NewWebRARenewRequestOnTemplateResponse

`func NewWebRARenewRequestOnTemplateResponse(module string, profile string, template WebRARenewRequestTemplateResponse, workflow string, ) *WebRARenewRequestOnTemplateResponse`

NewWebRARenewRequestOnTemplateResponse instantiates a new WebRARenewRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARenewRequestOnTemplateResponseWithDefaults

`func NewWebRARenewRequestOnTemplateResponseWithDefaults() *WebRARenewRequestOnTemplateResponse`

NewWebRARenewRequestOnTemplateResponseWithDefaults instantiates a new WebRARenewRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *WebRARenewRequestOnTemplateResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRARenewRequestOnTemplateResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRARenewRequestOnTemplateResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *WebRARenewRequestOnTemplateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRARenewRequestOnTemplateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRARenewRequestOnTemplateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTemplate

`func (o *WebRARenewRequestOnTemplateResponse) GetTemplate() WebRARenewRequestTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRARenewRequestOnTemplateResponse) GetTemplateOk() (*WebRARenewRequestTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRARenewRequestOnTemplateResponse) SetTemplate(v WebRARenewRequestTemplateResponse)`

SetTemplate sets Template field to given value.


### GetWorkflow

`func (o *WebRARenewRequestOnTemplateResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARenewRequestOnTemplateResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARenewRequestOnTemplateResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


