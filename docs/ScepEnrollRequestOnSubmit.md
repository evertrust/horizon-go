# ScepEnrollRequestOnSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | **interface{}** | The SCEP profile name | 
**Dn** | Pointer to **interface{}** | Fill the DN if DN whitelist is enabled. Contains the DN of the challenge | [optional] 
**Password** | Pointer to [**SecretString**](SecretString.md) | The password of the challenge. Must be set if password mode is &#x60;manual&#x60; | [optional] 
**RequesterComment** | Pointer to **NullableString** | Free-text field editable by the requester to provider more context on the request | [optional] 
**Module** | **string** | The module that will be used to process this request. For a SCEP request, this is always &#x60;scep&#x60; | 
**Workflow** | **string** | What this request will do. For an enrollment request, this is always &#x60;enroll&#x60; | 
**Template** | Pointer to [**ScepEnrollRequestTemplate**](ScepEnrollRequestTemplate.md) | The user-data that will be used to generate the challenge | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]

## Methods

### NewScepEnrollRequestOnSubmit

`func NewScepEnrollRequestOnSubmit(profile interface{}, module string, workflow string, ) *ScepEnrollRequestOnSubmit`

NewScepEnrollRequestOnSubmit instantiates a new ScepEnrollRequestOnSubmit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScepEnrollRequestOnSubmitWithDefaults

`func NewScepEnrollRequestOnSubmitWithDefaults() *ScepEnrollRequestOnSubmit`

NewScepEnrollRequestOnSubmitWithDefaults instantiates a new ScepEnrollRequestOnSubmit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *ScepEnrollRequestOnSubmit) GetProfile() interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ScepEnrollRequestOnSubmit) GetProfileOk() (*interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ScepEnrollRequestOnSubmit) SetProfile(v interface{})`

SetProfile sets Profile field to given value.


### SetProfileNil

`func (o *ScepEnrollRequestOnSubmit) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *ScepEnrollRequestOnSubmit) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetDn

`func (o *ScepEnrollRequestOnSubmit) GetDn() interface{}`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ScepEnrollRequestOnSubmit) GetDnOk() (*interface{}, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ScepEnrollRequestOnSubmit) SetDn(v interface{})`

SetDn sets Dn field to given value.

### HasDn

`func (o *ScepEnrollRequestOnSubmit) HasDn() bool`

HasDn returns a boolean if a field has been set.

### SetDnNil

`func (o *ScepEnrollRequestOnSubmit) SetDnNil(b bool)`

 SetDnNil sets the value for Dn to be an explicit nil

### UnsetDn
`func (o *ScepEnrollRequestOnSubmit) UnsetDn()`

UnsetDn ensures that no value is present for Dn, not even an explicit nil
### GetPassword

`func (o *ScepEnrollRequestOnSubmit) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ScepEnrollRequestOnSubmit) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ScepEnrollRequestOnSubmit) SetPassword(v SecretString)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ScepEnrollRequestOnSubmit) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRequesterComment

`func (o *ScepEnrollRequestOnSubmit) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *ScepEnrollRequestOnSubmit) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *ScepEnrollRequestOnSubmit) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *ScepEnrollRequestOnSubmit) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *ScepEnrollRequestOnSubmit) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *ScepEnrollRequestOnSubmit) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetModule

`func (o *ScepEnrollRequestOnSubmit) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ScepEnrollRequestOnSubmit) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ScepEnrollRequestOnSubmit) SetModule(v string)`

SetModule sets Module field to given value.


### GetWorkflow

`func (o *ScepEnrollRequestOnSubmit) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ScepEnrollRequestOnSubmit) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ScepEnrollRequestOnSubmit) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetTemplate

`func (o *ScepEnrollRequestOnSubmit) GetTemplate() ScepEnrollRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ScepEnrollRequestOnSubmit) GetTemplateOk() (*ScepEnrollRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ScepEnrollRequestOnSubmit) SetTemplate(v ScepEnrollRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ScepEnrollRequestOnSubmit) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDryRun

`func (o *ScepEnrollRequestOnSubmit) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ScepEnrollRequestOnSubmit) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ScepEnrollRequestOnSubmit) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ScepEnrollRequestOnSubmit) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *ScepEnrollRequestOnSubmit) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *ScepEnrollRequestOnSubmit) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


