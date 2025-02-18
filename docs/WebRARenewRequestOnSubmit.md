# WebRARenewRequestOnSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**NullableSecretString**](SecretString.md) | The password to decrypt the PKCS12 file. Must be set if password mode is &#x60;manual&#x60; | [optional] 
**CertificateId** | Pointer to **NullableString** | The id of the certificate to renew | [optional] 
**CertificatePem** | Pointer to **NullableString** | The PEM encoded certificate to renew | [optional] 
**RequesterComment** | Pointer to **NullableString** | Free-text field editable by the requester to provider more context on the request | [optional] 
**Module** | **string** | The module that will be used to process this request. For a WebRA request, this is always &#x60;webra&#x60; | 
**Workflow** | **string** | What this request will do. For a renewal request, this is always &#x60;renew&#x60; | 
**Template** | Pointer to [**WebRARenewRequestTemplate**](WebRARenewRequestTemplate.md) | The user-data that will be used to generate the certificate | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]

## Methods

### NewWebRARenewRequestOnSubmit

`func NewWebRARenewRequestOnSubmit(module string, workflow string, ) *WebRARenewRequestOnSubmit`

NewWebRARenewRequestOnSubmit instantiates a new WebRARenewRequestOnSubmit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARenewRequestOnSubmitWithDefaults

`func NewWebRARenewRequestOnSubmitWithDefaults() *WebRARenewRequestOnSubmit`

NewWebRARenewRequestOnSubmitWithDefaults instantiates a new WebRARenewRequestOnSubmit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *WebRARenewRequestOnSubmit) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebRARenewRequestOnSubmit) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebRARenewRequestOnSubmit) SetPassword(v SecretString)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebRARenewRequestOnSubmit) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *WebRARenewRequestOnSubmit) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *WebRARenewRequestOnSubmit) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCertificateId

`func (o *WebRARenewRequestOnSubmit) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *WebRARenewRequestOnSubmit) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *WebRARenewRequestOnSubmit) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *WebRARenewRequestOnSubmit) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *WebRARenewRequestOnSubmit) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *WebRARenewRequestOnSubmit) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCertificatePem

`func (o *WebRARenewRequestOnSubmit) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *WebRARenewRequestOnSubmit) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *WebRARenewRequestOnSubmit) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *WebRARenewRequestOnSubmit) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### SetCertificatePemNil

`func (o *WebRARenewRequestOnSubmit) SetCertificatePemNil(b bool)`

 SetCertificatePemNil sets the value for CertificatePem to be an explicit nil

### UnsetCertificatePem
`func (o *WebRARenewRequestOnSubmit) UnsetCertificatePem()`

UnsetCertificatePem ensures that no value is present for CertificatePem, not even an explicit nil
### GetRequesterComment

`func (o *WebRARenewRequestOnSubmit) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *WebRARenewRequestOnSubmit) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *WebRARenewRequestOnSubmit) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *WebRARenewRequestOnSubmit) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *WebRARenewRequestOnSubmit) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *WebRARenewRequestOnSubmit) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetModule

`func (o *WebRARenewRequestOnSubmit) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRARenewRequestOnSubmit) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRARenewRequestOnSubmit) SetModule(v string)`

SetModule sets Module field to given value.


### GetWorkflow

`func (o *WebRARenewRequestOnSubmit) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARenewRequestOnSubmit) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARenewRequestOnSubmit) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetTemplate

`func (o *WebRARenewRequestOnSubmit) GetTemplate() WebRARenewRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRARenewRequestOnSubmit) GetTemplateOk() (*WebRARenewRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRARenewRequestOnSubmit) SetTemplate(v WebRARenewRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRARenewRequestOnSubmit) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDryRun

`func (o *WebRARenewRequestOnSubmit) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRARenewRequestOnSubmit) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRARenewRequestOnSubmit) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRARenewRequestOnSubmit) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRARenewRequestOnSubmit) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRARenewRequestOnSubmit) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


