# WebRARenewRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**Module** | Pointer to **string** | The module that will be used to process this request. For a WebRA request, this is always &#x60;webra&#x60; | [optional] 
**Template** | Pointer to [**WebRARenewRequestTemplate**](WebRARenewRequestTemplate.md) | The user-data that will be used to generate the certificate | [optional] 
**Workflow** | Pointer to **string** | What this request will do. For a renewal request, this is always &#x60;renew&#x60; | [optional] 

## Methods

### NewWebRARenewRequestBase

`func NewWebRARenewRequestBase() *WebRARenewRequestBase`

NewWebRARenewRequestBase instantiates a new WebRARenewRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARenewRequestBaseWithDefaults

`func NewWebRARenewRequestBaseWithDefaults() *WebRARenewRequestBase`

NewWebRARenewRequestBaseWithDefaults instantiates a new WebRARenewRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *WebRARenewRequestBase) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRARenewRequestBase) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRARenewRequestBase) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRARenewRequestBase) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRARenewRequestBase) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRARenewRequestBase) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetModule

`func (o *WebRARenewRequestBase) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRARenewRequestBase) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRARenewRequestBase) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *WebRARenewRequestBase) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetTemplate

`func (o *WebRARenewRequestBase) GetTemplate() WebRARenewRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRARenewRequestBase) GetTemplateOk() (*WebRARenewRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRARenewRequestBase) SetTemplate(v WebRARenewRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRARenewRequestBase) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWorkflow

`func (o *WebRARenewRequestBase) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARenewRequestBase) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARenewRequestBase) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WebRARenewRequestBase) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


