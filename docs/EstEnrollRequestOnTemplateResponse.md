# EstEnrollRequestOnTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** |  | 
**Module** | **string** |  | 
**Template** | [**EstEnrollRequestTemplateResponse**](EstEnrollRequestTemplateResponse.md) | The template with the constraint set on the profile | 
**Profile** | **string** | The profile for which to return the template. | 

## Methods

### NewEstEnrollRequestOnTemplateResponse

`func NewEstEnrollRequestOnTemplateResponse(workflow string, module string, template EstEnrollRequestTemplateResponse, profile string, ) *EstEnrollRequestOnTemplateResponse`

NewEstEnrollRequestOnTemplateResponse instantiates a new EstEnrollRequestOnTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstEnrollRequestOnTemplateResponseWithDefaults

`func NewEstEnrollRequestOnTemplateResponseWithDefaults() *EstEnrollRequestOnTemplateResponse`

NewEstEnrollRequestOnTemplateResponseWithDefaults instantiates a new EstEnrollRequestOnTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *EstEnrollRequestOnTemplateResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *EstEnrollRequestOnTemplateResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *EstEnrollRequestOnTemplateResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetModule

`func (o *EstEnrollRequestOnTemplateResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *EstEnrollRequestOnTemplateResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *EstEnrollRequestOnTemplateResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetTemplate

`func (o *EstEnrollRequestOnTemplateResponse) GetTemplate() EstEnrollRequestTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EstEnrollRequestOnTemplateResponse) GetTemplateOk() (*EstEnrollRequestTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EstEnrollRequestOnTemplateResponse) SetTemplate(v EstEnrollRequestTemplateResponse)`

SetTemplate sets Template field to given value.


### GetProfile

`func (o *EstEnrollRequestOnTemplateResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *EstEnrollRequestOnTemplateResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *EstEnrollRequestOnTemplateResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


