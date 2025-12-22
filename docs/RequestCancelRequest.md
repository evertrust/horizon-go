# RequestCancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the request to cancel | 
**Module** | [**Module**](Module.md) |  | 
**Workflow** | [**Workflow**](Workflow.md) |  | 

## Methods

### NewRequestCancelRequest

`func NewRequestCancelRequest(id string, module Module, workflow Workflow, ) *RequestCancelRequest`

NewRequestCancelRequest instantiates a new RequestCancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCancelRequestWithDefaults

`func NewRequestCancelRequestWithDefaults() *RequestCancelRequest`

NewRequestCancelRequestWithDefaults instantiates a new RequestCancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestCancelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestCancelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestCancelRequest) SetId(v string)`

SetId sets Id field to given value.


### GetModule

`func (o *RequestCancelRequest) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RequestCancelRequest) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RequestCancelRequest) SetModule(v Module)`

SetModule sets Module field to given value.


### GetWorkflow

`func (o *RequestCancelRequest) GetWorkflow() Workflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RequestCancelRequest) GetWorkflowOk() (*Workflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RequestCancelRequest) SetWorkflow(v Workflow)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


