# ScepEnrollRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** |  | 
**Template** | [**ScepEnrollRequestTemplateResponse**](ScepEnrollRequestTemplateResponse.md) | The template with the constraint set on the profile | 
**Workflow** | **string** |  | 
**Profile** | **string** | The profile for which to return the template. | 

## Methods

### NewScepEnrollRequestOnTemplateResponse

`func NewScepEnrollRequestOnTemplateResponse(module string, template ScepEnrollRequestTemplateResponse, workflow string, profile string, ) *ScepEnrollRequestOnTemplateResponse`

NewScepEnrollRequestOnTemplateResponse instantiates a new ScepEnrollRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScepEnrollRequestOnTemplateResponseWithDefaults

`func NewScepEnrollRequestOnTemplateResponseWithDefaults() *ScepEnrollRequestOnTemplateResponse`

NewScepEnrollRequestOnTemplateResponseWithDefaults instantiates a new ScepEnrollRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *ScepEnrollRequestOnTemplateResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ScepEnrollRequestOnTemplateResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ScepEnrollRequestOnTemplateResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetTemplate

`func (o *ScepEnrollRequestOnTemplateResponse) GetTemplate() ScepEnrollRequestTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ScepEnrollRequestOnTemplateResponse) GetTemplateOk() (*ScepEnrollRequestTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ScepEnrollRequestOnTemplateResponse) SetTemplate(v ScepEnrollRequestTemplateResponse)`

SetTemplate sets Template field to given value.


### GetWorkflow

`func (o *ScepEnrollRequestOnTemplateResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ScepEnrollRequestOnTemplateResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ScepEnrollRequestOnTemplateResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetProfile

`func (o *ScepEnrollRequestOnTemplateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ScepEnrollRequestOnTemplateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ScepEnrollRequestOnTemplateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


