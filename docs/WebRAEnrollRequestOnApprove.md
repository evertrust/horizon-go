# WebRAEnrollRequestOnApprove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ApproverComment** | Pointer to **NullableString** | Free-text field editable by the approver to provider more context on the request | [optional] 
**Module** | **string** | The module that will be used to process this request. For a WebRA request, this is always &#x60;webra&#x60; | 
**Workflow** | **string** | What this request will do. For an enrollment request, this is always &#x60;enroll&#x60; | 
**Template** | Pointer to [**WebRAEnrollRequestTemplate**](WebRAEnrollRequestTemplate.md) | The user-data that will be used to generate the certificate | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]

## Methods

### NewWebRAEnrollRequestOnApprove

`func NewWebRAEnrollRequestOnApprove(id string, module string, workflow string, ) *WebRAEnrollRequestOnApprove`

NewWebRAEnrollRequestOnApprove instantiates a new WebRAEnrollRequestOnApprove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAEnrollRequestOnApproveWithDefaults

`func NewWebRAEnrollRequestOnApproveWithDefaults() *WebRAEnrollRequestOnApprove`

NewWebRAEnrollRequestOnApproveWithDefaults instantiates a new WebRAEnrollRequestOnApprove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebRAEnrollRequestOnApprove) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebRAEnrollRequestOnApprove) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebRAEnrollRequestOnApprove) SetId(v string)`

SetId sets Id field to given value.


### GetApproverComment

`func (o *WebRAEnrollRequestOnApprove) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *WebRAEnrollRequestOnApprove) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *WebRAEnrollRequestOnApprove) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *WebRAEnrollRequestOnApprove) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *WebRAEnrollRequestOnApprove) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *WebRAEnrollRequestOnApprove) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetModule

`func (o *WebRAEnrollRequestOnApprove) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAEnrollRequestOnApprove) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAEnrollRequestOnApprove) SetModule(v string)`

SetModule sets Module field to given value.


### GetWorkflow

`func (o *WebRAEnrollRequestOnApprove) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAEnrollRequestOnApprove) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAEnrollRequestOnApprove) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetTemplate

`func (o *WebRAEnrollRequestOnApprove) GetTemplate() WebRAEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAEnrollRequestOnApprove) GetTemplateOk() (*WebRAEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAEnrollRequestOnApprove) SetTemplate(v WebRAEnrollRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRAEnrollRequestOnApprove) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDryRun

`func (o *WebRAEnrollRequestOnApprove) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAEnrollRequestOnApprove) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAEnrollRequestOnApprove) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAEnrollRequestOnApprove) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAEnrollRequestOnApprove) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAEnrollRequestOnApprove) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


