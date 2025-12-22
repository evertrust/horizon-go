# RequestableWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Perform** | **bool** | The principal can perform this workflow without validation | 
**Request** | **bool** | The principal has the ability to create a request for this workflow | 
**Workflow** | **string** | A possible workflow on this profile | 

## Methods

### NewRequestableWorkflow

`func NewRequestableWorkflow(perform bool, request bool, workflow string, ) *RequestableWorkflow`

NewRequestableWorkflow instantiates a new RequestableWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestableWorkflowWithDefaults

`func NewRequestableWorkflowWithDefaults() *RequestableWorkflow`

NewRequestableWorkflowWithDefaults instantiates a new RequestableWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerform

`func (o *RequestableWorkflow) GetPerform() bool`

GetPerform returns the Perform field if non-nil, zero value otherwise.

### GetPerformOk

`func (o *RequestableWorkflow) GetPerformOk() (*bool, bool)`

GetPerformOk returns a tuple with the Perform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerform

`func (o *RequestableWorkflow) SetPerform(v bool)`

SetPerform sets Perform field to given value.


### GetRequest

`func (o *RequestableWorkflow) GetRequest() bool`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *RequestableWorkflow) GetRequestOk() (*bool, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *RequestableWorkflow) SetRequest(v bool)`

SetRequest sets Request field to given value.


### GetWorkflow

`func (o *RequestableWorkflow) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RequestableWorkflow) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RequestableWorkflow) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


