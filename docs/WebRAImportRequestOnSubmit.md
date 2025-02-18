# WebRAImportRequestOnSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **NullableString** | The profile name on which to import | [optional] 
**CertificateId** | Pointer to **NullableString** | The id of the certificate to import | [optional] 
**CertificatePem** | Pointer to **NullableString** | The PEM encoded certificate to import | [optional] 
**RequesterComment** | Pointer to **NullableString** | Free-text field editable by the requester to provider more context on the request | [optional] 
**Module** | **string** | The module that will be used to process this request. For a WebRA request, this is always &#x60;webra&#x60; | 
**Workflow** | **string** | What this request will do. For an import request, this is always &#x60;import&#x60; | 
**Template** | Pointer to [**WebRAImportRequestTemplate**](WebRAImportRequestTemplate.md) | The user-data that will be added on certificate import | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]

## Methods

### NewWebRAImportRequestOnSubmit

`func NewWebRAImportRequestOnSubmit(module string, workflow string, ) *WebRAImportRequestOnSubmit`

NewWebRAImportRequestOnSubmit instantiates a new WebRAImportRequestOnSubmit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAImportRequestOnSubmitWithDefaults

`func NewWebRAImportRequestOnSubmitWithDefaults() *WebRAImportRequestOnSubmit`

NewWebRAImportRequestOnSubmitWithDefaults instantiates a new WebRAImportRequestOnSubmit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *WebRAImportRequestOnSubmit) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAImportRequestOnSubmit) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAImportRequestOnSubmit) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *WebRAImportRequestOnSubmit) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *WebRAImportRequestOnSubmit) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *WebRAImportRequestOnSubmit) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetCertificateId

`func (o *WebRAImportRequestOnSubmit) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *WebRAImportRequestOnSubmit) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *WebRAImportRequestOnSubmit) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *WebRAImportRequestOnSubmit) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *WebRAImportRequestOnSubmit) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *WebRAImportRequestOnSubmit) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCertificatePem

`func (o *WebRAImportRequestOnSubmit) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *WebRAImportRequestOnSubmit) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *WebRAImportRequestOnSubmit) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *WebRAImportRequestOnSubmit) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### SetCertificatePemNil

`func (o *WebRAImportRequestOnSubmit) SetCertificatePemNil(b bool)`

 SetCertificatePemNil sets the value for CertificatePem to be an explicit nil

### UnsetCertificatePem
`func (o *WebRAImportRequestOnSubmit) UnsetCertificatePem()`

UnsetCertificatePem ensures that no value is present for CertificatePem, not even an explicit nil
### GetRequesterComment

`func (o *WebRAImportRequestOnSubmit) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *WebRAImportRequestOnSubmit) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *WebRAImportRequestOnSubmit) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *WebRAImportRequestOnSubmit) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *WebRAImportRequestOnSubmit) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *WebRAImportRequestOnSubmit) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetModule

`func (o *WebRAImportRequestOnSubmit) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAImportRequestOnSubmit) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAImportRequestOnSubmit) SetModule(v string)`

SetModule sets Module field to given value.


### GetWorkflow

`func (o *WebRAImportRequestOnSubmit) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAImportRequestOnSubmit) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAImportRequestOnSubmit) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetTemplate

`func (o *WebRAImportRequestOnSubmit) GetTemplate() WebRAImportRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAImportRequestOnSubmit) GetTemplateOk() (*WebRAImportRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAImportRequestOnSubmit) SetTemplate(v WebRAImportRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRAImportRequestOnSubmit) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDryRun

`func (o *WebRAImportRequestOnSubmit) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRAImportRequestOnSubmit) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRAImportRequestOnSubmit) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRAImportRequestOnSubmit) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRAImportRequestOnSubmit) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRAImportRequestOnSubmit) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


