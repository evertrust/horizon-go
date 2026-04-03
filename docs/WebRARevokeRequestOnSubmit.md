# WebRARevokeRequestOnSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | Pointer to **NullableString** | The id of the certificate to revoke | [optional] 
**CertificatePem** | Pointer to **NullableString** | The PEM encoded certificate to revoke | [optional] 
**DryRun** | Pointer to **NullableBool** | If true, the request is validated, but will not result in an enrollment | [optional] [default to false]
**Template** | [**WebRARevokeRequestTemplate**](WebRARevokeRequestTemplate.md) | The user-data that will be used to revoke the certificate | 
**Workflow** | **string** | What this request will do. For a revocation request, this is always &#x60;revoke&#x60; | 

## Methods

### NewWebRARevokeRequestOnSubmit

`func NewWebRARevokeRequestOnSubmit(template WebRARevokeRequestTemplate, workflow string, ) *WebRARevokeRequestOnSubmit`

NewWebRARevokeRequestOnSubmit instantiates a new WebRARevokeRequestOnSubmit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARevokeRequestOnSubmitWithDefaults

`func NewWebRARevokeRequestOnSubmitWithDefaults() *WebRARevokeRequestOnSubmit`

NewWebRARevokeRequestOnSubmitWithDefaults instantiates a new WebRARevokeRequestOnSubmit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *WebRARevokeRequestOnSubmit) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *WebRARevokeRequestOnSubmit) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *WebRARevokeRequestOnSubmit) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *WebRARevokeRequestOnSubmit) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *WebRARevokeRequestOnSubmit) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *WebRARevokeRequestOnSubmit) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCertificatePem

`func (o *WebRARevokeRequestOnSubmit) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *WebRARevokeRequestOnSubmit) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *WebRARevokeRequestOnSubmit) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *WebRARevokeRequestOnSubmit) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### SetCertificatePemNil

`func (o *WebRARevokeRequestOnSubmit) SetCertificatePemNil(b bool)`

 SetCertificatePemNil sets the value for CertificatePem to be an explicit nil

### UnsetCertificatePem
`func (o *WebRARevokeRequestOnSubmit) UnsetCertificatePem()`

UnsetCertificatePem ensures that no value is present for CertificatePem, not even an explicit nil
### GetDryRun

`func (o *WebRARevokeRequestOnSubmit) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *WebRARevokeRequestOnSubmit) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *WebRARevokeRequestOnSubmit) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *WebRARevokeRequestOnSubmit) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRunNil

`func (o *WebRARevokeRequestOnSubmit) SetDryRunNil(b bool)`

 SetDryRunNil sets the value for DryRun to be an explicit nil

### UnsetDryRun
`func (o *WebRARevokeRequestOnSubmit) UnsetDryRun()`

UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
### GetTemplate

`func (o *WebRARevokeRequestOnSubmit) GetTemplate() WebRARevokeRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRARevokeRequestOnSubmit) GetTemplateOk() (*WebRARevokeRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRARevokeRequestOnSubmit) SetTemplate(v WebRARevokeRequestTemplate)`

SetTemplate sets Template field to given value.


### GetWorkflow

`func (o *WebRARevokeRequestOnSubmit) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARevokeRequestOnSubmit) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARevokeRequestOnSubmit) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


