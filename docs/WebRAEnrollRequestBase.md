# WebRAEnrollRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**Module** | Pointer to **string** | The module that will be used to process this request. For a WebRA request, this is always &#x60;webra&#x60; | [optional] 
**Template** | Pointer to [**WebRAEnrollRequestTemplate**](WebRAEnrollRequestTemplate.md) | The user-data that will be used to generate the certificate | [optional] 
**Workflow** | Pointer to **string** | What this request will do. For an enrollment request, this is always &#x60;enroll&#x60; | [optional] 

## Methods

### NewWebRAEnrollRequestBase

`func NewWebRAEnrollRequestBase() *WebRAEnrollRequestBase`

NewWebRAEnrollRequestBase instantiates a new WebRAEnrollRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAEnrollRequestBaseWithDefaults

`func NewWebRAEnrollRequestBaseWithDefaults() *WebRAEnrollRequestBase`

NewWebRAEnrollRequestBaseWithDefaults instantiates a new WebRAEnrollRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *WebRAEnrollRequestBase) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAEnrollRequestBase) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAEnrollRequestBase) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAEnrollRequestBase) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAEnrollRequestBase) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAEnrollRequestBase) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetModule

`func (o *WebRAEnrollRequestBase) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAEnrollRequestBase) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAEnrollRequestBase) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *WebRAEnrollRequestBase) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetTemplate

`func (o *WebRAEnrollRequestBase) GetTemplate() WebRAEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAEnrollRequestBase) GetTemplateOk() (*WebRAEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAEnrollRequestBase) SetTemplate(v WebRAEnrollRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRAEnrollRequestBase) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWorkflow

`func (o *WebRAEnrollRequestBase) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAEnrollRequestBase) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAEnrollRequestBase) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WebRAEnrollRequestBase) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


