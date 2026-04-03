# WebRAUpdateRequestOnApprove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ApproverComment** | Pointer to **NullableString** | Free-text field editable by the approver to provider more context on the request | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**Template** | Pointer to [**WebRAUpdateRequestTemplate**](WebRAUpdateRequestTemplate.md) | The user-data that will be used to update the certificate | [optional] 
**Workflow** | **string** | What this request will do. For an update request, this is always &#x60;update&#x60; | 

## Methods

### NewWebRAUpdateRequestOnApprove

`func NewWebRAUpdateRequestOnApprove(id string, workflow string, ) *WebRAUpdateRequestOnApprove`

NewWebRAUpdateRequestOnApprove instantiates a new WebRAUpdateRequestOnApprove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAUpdateRequestOnApproveWithDefaults

`func NewWebRAUpdateRequestOnApproveWithDefaults() *WebRAUpdateRequestOnApprove`

NewWebRAUpdateRequestOnApproveWithDefaults instantiates a new WebRAUpdateRequestOnApprove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebRAUpdateRequestOnApprove) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebRAUpdateRequestOnApprove) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebRAUpdateRequestOnApprove) SetId(v string)`

SetId sets Id field to given value.


### GetApproverComment

`func (o *WebRAUpdateRequestOnApprove) GetApproverComment() string`

GetApproverComment returns the ApproverComment field if non-nil, zero value otherwise.

### GetApproverCommentOk

`func (o *WebRAUpdateRequestOnApprove) GetApproverCommentOk() (*string, bool)`

GetApproverCommentOk returns a tuple with the ApproverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverComment

`func (o *WebRAUpdateRequestOnApprove) SetApproverComment(v string)`

SetApproverComment sets ApproverComment field to given value.

### HasApproverComment

`func (o *WebRAUpdateRequestOnApprove) HasApproverComment() bool`

HasApproverComment returns a boolean if a field has been set.

### SetApproverCommentNil

`func (o *WebRAUpdateRequestOnApprove) SetApproverCommentNil(b bool)`

 SetApproverCommentNil sets the value for ApproverComment to be an explicit nil

### UnsetApproverComment
`func (o *WebRAUpdateRequestOnApprove) UnsetApproverComment()`

UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
### GetDryRun

`func (o *WebRAUpdateRequestOnApprove) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAUpdateRequestOnApprove) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAUpdateRequestOnApprove) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAUpdateRequestOnApprove) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAUpdateRequestOnApprove) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAUpdateRequestOnApprove) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetTemplate

`func (o *WebRAUpdateRequestOnApprove) GetTemplate() WebRAUpdateRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAUpdateRequestOnApprove) GetTemplateOk() (*WebRAUpdateRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAUpdateRequestOnApprove) SetTemplate(v WebRAUpdateRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRAUpdateRequestOnApprove) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWorkflow

`func (o *WebRAUpdateRequestOnApprove) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAUpdateRequestOnApprove) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAUpdateRequestOnApprove) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


