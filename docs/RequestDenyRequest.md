# RequestDenyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the request to deny | 
**Module** | [**Module**](Module.md) |  | 
**Workflow** | [**Workflow**](Workflow.md) |  | 
**ApproverComment** | Pointer to **string** | Free-text field editable by the approver to provider more context on the denial | [optional] 

## Methods

### NewRequestDenyRequest

`func NewRequestDenyRequest(id string, module Module, workflow Workflow, ) *RequestDenyRequest`

NewRequestDenyRequest instantiates a new RequestDenyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDenyRequestWithDefaults

`func NewRequestDenyRequestWithDefaults() *RequestDenyRequest`

NewRequestDenyRequestWithDefaults instantiates a new RequestDenyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestDenyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestDenyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestDenyRequest) SetId(v string)`

SetId sets Id field to given value.


### GetModule

`func (o *RequestDenyRequest) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RequestDenyRequest) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RequestDenyRequest) SetModule(v Module)`

SetModule sets Module field to given value.


### GetWorkflow

`func (o *RequestDenyRequest) GetWorkflow() Workflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RequestDenyRequest) GetWorkflowOk() (*Workflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RequestDenyRequest) SetWorkflow(v Workflow)`

SetWorkflow sets Workflow field to given value.


### GetApproverComment

`func (o *RequestDenyRequest) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *RequestDenyRequest) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *RequestDenyRequest) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *RequestDenyRequest) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


