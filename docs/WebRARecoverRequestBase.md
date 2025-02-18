# WebRARecoverRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | Pointer to **string** | What this request will do. For a recovery request, this is always &#x60;recover&#x60; | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]

## Methods

### NewWebRARecoverRequestBase

`func NewWebRARecoverRequestBase() *WebRARecoverRequestBase`

NewWebRARecoverRequestBase instantiates a new WebRARecoverRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARecoverRequestBaseWithDefaults

`func NewWebRARecoverRequestBaseWithDefaults() *WebRARecoverRequestBase`

NewWebRARecoverRequestBaseWithDefaults instantiates a new WebRARecoverRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WebRARecoverRequestBase) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARecoverRequestBase) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARecoverRequestBase) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WebRARecoverRequestBase) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetDryRun

`func (o *WebRARecoverRequestBase) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRARecoverRequestBase) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRARecoverRequestBase) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRARecoverRequestBase) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRARecoverRequestBase) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRARecoverRequestBase) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


