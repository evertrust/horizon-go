# WebRARecoverRequestOnSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | Pointer to **NullableString** | The id of the certificate to renew | [optional] 
**CertificatePem** | Pointer to **NullableString** | The PEM encoded certificate to renew | [optional] 
**Password** | Pointer to [**NullableSecretString**](SecretString.md) | The password to decrypt the PKCS12 file. Must be set if password mode is &#x60;manual&#x60; | [optional] 
**RequesterComment** | Pointer to **NullableString** | Free-text field editable by the requester to provider more context on the request | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**Workflow** | **string** | What this request will do. For a recovery request, this is always &#x60;recover&#x60; | 

## Methods

### NewWebRARecoverRequestOnSubmit

`func NewWebRARecoverRequestOnSubmit(workflow string, ) *WebRARecoverRequestOnSubmit`

NewWebRARecoverRequestOnSubmit instantiates a new WebRARecoverRequestOnSubmit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARecoverRequestOnSubmitWithDefaults

`func NewWebRARecoverRequestOnSubmitWithDefaults() *WebRARecoverRequestOnSubmit`

NewWebRARecoverRequestOnSubmitWithDefaults instantiates a new WebRARecoverRequestOnSubmit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *WebRARecoverRequestOnSubmit) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *WebRARecoverRequestOnSubmit) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *WebRARecoverRequestOnSubmit) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *WebRARecoverRequestOnSubmit) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *WebRARecoverRequestOnSubmit) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *WebRARecoverRequestOnSubmit) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCertificatePem

`func (o *WebRARecoverRequestOnSubmit) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *WebRARecoverRequestOnSubmit) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *WebRARecoverRequestOnSubmit) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *WebRARecoverRequestOnSubmit) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### SetCertificatePemNil

`func (o *WebRARecoverRequestOnSubmit) SetCertificatePemNil(b bool)`

 SetCertificatePemNil sets the value for CertificatePem to be an explicit nil

### UnsetCertificatePem
`func (o *WebRARecoverRequestOnSubmit) UnsetCertificatePem()`

UnsetCertificatePem ensures that no value is present for CertificatePem, not even an explicit nil
### GetPassword

`func (o *WebRARecoverRequestOnSubmit) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebRARecoverRequestOnSubmit) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebRARecoverRequestOnSubmit) SetPassword(v SecretString)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebRARecoverRequestOnSubmit) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *WebRARecoverRequestOnSubmit) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *WebRARecoverRequestOnSubmit) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRequesterComment

`func (o *WebRARecoverRequestOnSubmit) GetRequesterComment() string`

GetRequesterComment returns the RequesterComment field if non-nil, zero value otherwise.

### GetRequesterCommentOk

`func (o *WebRARecoverRequestOnSubmit) GetRequesterCommentOk() (*string, bool)`

GetRequesterCommentOk returns a tuple with the RequesterComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterComment

`func (o *WebRARecoverRequestOnSubmit) SetRequesterComment(v string)`

SetRequesterComment sets RequesterComment field to given value.

### HasRequesterComment

`func (o *WebRARecoverRequestOnSubmit) HasRequesterComment() bool`

HasRequesterComment returns a boolean if a field has been set.

### SetRequesterCommentNil

`func (o *WebRARecoverRequestOnSubmit) SetRequesterCommentNil(b bool)`

 SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil

### UnsetRequesterComment
`func (o *WebRARecoverRequestOnSubmit) UnsetRequesterComment()`

UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
### GetDryRun

`func (o *WebRARecoverRequestOnSubmit) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRARecoverRequestOnSubmit) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRARecoverRequestOnSubmit) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRARecoverRequestOnSubmit) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRARecoverRequestOnSubmit) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRARecoverRequestOnSubmit) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetWorkflow

`func (o *WebRARecoverRequestOnSubmit) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARecoverRequestOnSubmit) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARecoverRequestOnSubmit) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


