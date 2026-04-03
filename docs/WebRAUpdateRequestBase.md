# WebRAUpdateRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**Template** | Pointer to [**WebRAUpdateRequestTemplate**](WebRAUpdateRequestTemplate.md) | The user-data that will be used to update the certificate | [optional] 
**Workflow** | Pointer to **string** | What this request will do. For an update request, this is always &#x60;update&#x60; | [optional] 

## Methods

### NewWebRAUpdateRequestBase

`func NewWebRAUpdateRequestBase() *WebRAUpdateRequestBase`

NewWebRAUpdateRequestBase instantiates a new WebRAUpdateRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAUpdateRequestBaseWithDefaults

`func NewWebRAUpdateRequestBaseWithDefaults() *WebRAUpdateRequestBase`

NewWebRAUpdateRequestBaseWithDefaults instantiates a new WebRAUpdateRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *WebRAUpdateRequestBase) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAUpdateRequestBase) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAUpdateRequestBase) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAUpdateRequestBase) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAUpdateRequestBase) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAUpdateRequestBase) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetTemplate

`func (o *WebRAUpdateRequestBase) GetTemplate() WebRAUpdateRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAUpdateRequestBase) GetTemplateOk() (*WebRAUpdateRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAUpdateRequestBase) SetTemplate(v WebRAUpdateRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRAUpdateRequestBase) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWorkflow

`func (o *WebRAUpdateRequestBase) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAUpdateRequestBase) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAUpdateRequestBase) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WebRAUpdateRequestBase) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


