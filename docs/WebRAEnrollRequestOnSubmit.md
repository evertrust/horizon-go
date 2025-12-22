# WebRAEnrollRequestOnSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**SecretString**](SecretString.md) | The password to decrypt the PKCS12 file. Must be set if password mode is &#x60;manual&#x60; | [optional] 
**Profile** | **string** | The WebRA profile name | 
**RequesterComment** | Pointer to **NullableString** | Free-text field editable by the requester to provider more context on the request | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**Module** | **string** | The module that will be used to process this request. For a WebRA request, this is always &#x60;webra&#x60; | 
**Template** | [**WebRAEnrollRequestTemplate**](WebRAEnrollRequestTemplate.md) | The user-data that will be used to generate the certificate | 
**Workflow** | **string** | What this request will do. For an enrollment request, this is always &#x60;enroll&#x60; | 

## Methods

### NewWebRAEnrollRequestOnSubmit

`func NewWebRAEnrollRequestOnSubmit(profile string, module string, template WebRAEnrollRequestTemplate, workflow string, ) *WebRAEnrollRequestOnSubmit`

NewWebRAEnrollRequestOnSubmit instantiates a new WebRAEnrollRequestOnSubmit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAEnrollRequestOnSubmitWithDefaults

`func NewWebRAEnrollRequestOnSubmitWithDefaults() *WebRAEnrollRequestOnSubmit`

NewWebRAEnrollRequestOnSubmitWithDefaults instantiates a new WebRAEnrollRequestOnSubmit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *WebRAEnrollRequestOnSubmit) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebRAEnrollRequestOnSubmit) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebRAEnrollRequestOnSubmit) SetPassword(v SecretString)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebRAEnrollRequestOnSubmit) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProfile

`func (o *WebRAEnrollRequestOnSubmit) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAEnrollRequestOnSubmit) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAEnrollRequestOnSubmit) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRequesterComment

`func (o *WebRAEnrollRequestOnSubmit) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *WebRAEnrollRequestOnSubmit) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *WebRAEnrollRequestOnSubmit) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *WebRAEnrollRequestOnSubmit) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *WebRAEnrollRequestOnSubmit) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *WebRAEnrollRequestOnSubmit) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetDryRun

`func (o *WebRAEnrollRequestOnSubmit) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAEnrollRequestOnSubmit) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAEnrollRequestOnSubmit) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAEnrollRequestOnSubmit) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAEnrollRequestOnSubmit) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAEnrollRequestOnSubmit) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetModule

`func (o *WebRAEnrollRequestOnSubmit) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAEnrollRequestOnSubmit) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAEnrollRequestOnSubmit) SetModule(v string)`

SetModule sets Module field to given value.


### GetTemplate

`func (o *WebRAEnrollRequestOnSubmit) GetTemplate() WebRAEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAEnrollRequestOnSubmit) GetTemplateOk() (*WebRAEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAEnrollRequestOnSubmit) SetTemplate(v WebRAEnrollRequestTemplate)`

SetTemplate sets Template field to given value.


### GetWorkflow

`func (o *WebRAEnrollRequestOnSubmit) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAEnrollRequestOnSubmit) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAEnrollRequestOnSubmit) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


