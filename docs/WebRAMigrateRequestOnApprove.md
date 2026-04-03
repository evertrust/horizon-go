# WebRAMigrateRequestOnApprove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ApproverComment** | Pointer to **NullableString** | Free-text field editable by the approver to provider more context on the request | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in a migration | [optional] [default to false]
**Module** | **string** | The module that will be used to process this request. For a WebRA request, this is always &#x60;webra&#x60; | 
**Template** | Pointer to [**WebRAMigrateRequestTemplate**](WebRAMigrateRequestTemplate.md) | The user-data that will be used to migrate the certificate | [optional] 
**Workflow** | **string** | What this request will do. For a migration request, this is always &#x60;migrate&#x60; | 

## Methods

### NewWebRAMigrateRequestOnApprove

`func NewWebRAMigrateRequestOnApprove(id string, module string, workflow string, ) *WebRAMigrateRequestOnApprove`

NewWebRAMigrateRequestOnApprove instantiates a new WebRAMigrateRequestOnApprove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAMigrateRequestOnApproveWithDefaults

`func NewWebRAMigrateRequestOnApproveWithDefaults() *WebRAMigrateRequestOnApprove`

NewWebRAMigrateRequestOnApproveWithDefaults instantiates a new WebRAMigrateRequestOnApprove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebRAMigrateRequestOnApprove) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebRAMigrateRequestOnApprove) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebRAMigrateRequestOnApprove) SetId(v string)`

SetId sets Id field to given value.


### GetApproverComment

`func (o *WebRAMigrateRequestOnApprove) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *WebRAMigrateRequestOnApprove) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *WebRAMigrateRequestOnApprove) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *WebRAMigrateRequestOnApprove) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *WebRAMigrateRequestOnApprove) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *WebRAMigrateRequestOnApprove) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetDryRun

`func (o *WebRAMigrateRequestOnApprove) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAMigrateRequestOnApprove) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAMigrateRequestOnApprove) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAMigrateRequestOnApprove) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAMigrateRequestOnApprove) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAMigrateRequestOnApprove) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetModule

`func (o *WebRAMigrateRequestOnApprove) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAMigrateRequestOnApprove) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAMigrateRequestOnApprove) SetModule(v string)`

SetModule sets Module field to given value.


### GetTemplate

`func (o *WebRAMigrateRequestOnApprove) GetTemplate() WebRAMigrateRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAMigrateRequestOnApprove) GetTemplateOk() (*WebRAMigrateRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAMigrateRequestOnApprove) SetTemplate(v WebRAMigrateRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRAMigrateRequestOnApprove) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWorkflow

`func (o *WebRAMigrateRequestOnApprove) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAMigrateRequestOnApprove) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAMigrateRequestOnApprove) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


