# RequestTemplate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** | The request module | 
**Profile** | **string** | The profile for which to return the template. | 
**Template** | [**ScepEnrollRequestTemplateResponse**](ScepEnrollRequestTemplateResponse.md) | The template with the constraint set on the profile | 
**Workflow** | **string** | The workflow for which to return the template. | 

## Methods

### NewRequestTemplate200Response

`func NewRequestTemplate200Response(module string, profile string, template ScepEnrollRequestTemplateResponse, workflow string, ) *RequestTemplate200Response`

NewRequestTemplate200Response instantiates a new RequestTemplate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTemplate200ResponseWithDefaults

`func NewRequestTemplate200ResponseWithDefaults() *RequestTemplate200Response`

NewRequestTemplate200ResponseWithDefaults instantiates a new RequestTemplate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *RequestTemplate200Response) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RequestTemplate200Response) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RequestTemplate200Response) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *RequestTemplate200Response) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *RequestTemplate200Response) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *RequestTemplate200Response) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTemplate

`func (o *RequestTemplate200Response) GetTemplate() ScepEnrollRequestTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *RequestTemplate200Response) GetTemplateOk() (*ScepEnrollRequestTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *RequestTemplate200Response) SetTemplate(v ScepEnrollRequestTemplateResponse)`

SetTemplate sets Template field to given value.


### GetWorkflow

`func (o *RequestTemplate200Response) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RequestTemplate200Response) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RequestTemplate200Response) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


