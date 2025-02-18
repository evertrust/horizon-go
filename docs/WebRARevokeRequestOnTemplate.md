# WebRARevokeRequestOnTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** |  | 
**Profile** | Pointer to **string** |  | [optional] 
**CertificateId** | Pointer to **NullableString** | Used to pre-fill the template field with the certificate values. | [optional] 
**CertificatePem** | Pointer to **NullableString** | Used to pre-fill the template field with the certificate values. | [optional] 
**Module** | Pointer to **string** | The request module | [optional] 

## Methods

### NewWebRARevokeRequestOnTemplate

`func NewWebRARevokeRequestOnTemplate(workflow string, ) *WebRARevokeRequestOnTemplate`

NewWebRARevokeRequestOnTemplate instantiates a new WebRARevokeRequestOnTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARevokeRequestOnTemplateWithDefaults

`func NewWebRARevokeRequestOnTemplateWithDefaults() *WebRARevokeRequestOnTemplate`

NewWebRARevokeRequestOnTemplateWithDefaults instantiates a new WebRARevokeRequestOnTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WebRARevokeRequestOnTemplate) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARevokeRequestOnTemplate) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARevokeRequestOnTemplate) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetProfile

`func (o *WebRARevokeRequestOnTemplate) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRARevokeRequestOnTemplate) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRARevokeRequestOnTemplate) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *WebRARevokeRequestOnTemplate) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetCertificateId

`func (o *WebRARevokeRequestOnTemplate) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *WebRARevokeRequestOnTemplate) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *WebRARevokeRequestOnTemplate) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *WebRARevokeRequestOnTemplate) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *WebRARevokeRequestOnTemplate) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *WebRARevokeRequestOnTemplate) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCertificatePem

`func (o *WebRARevokeRequestOnTemplate) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *WebRARevokeRequestOnTemplate) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *WebRARevokeRequestOnTemplate) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *WebRARevokeRequestOnTemplate) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### SetCertificatePemNil

`func (o *WebRARevokeRequestOnTemplate) SetCertificatePemNil(b bool)`

 SetCertificatePemNil sets the value for CertificatePem to be an explicit nil

### UnsetCertificatePem
`func (o *WebRARevokeRequestOnTemplate) UnsetCertificatePem()`

UnsetCertificatePem ensures that no value is present for CertificatePem, not even an explicit nil
### GetModule

`func (o *WebRARevokeRequestOnTemplate) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRARevokeRequestOnTemplate) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRARevokeRequestOnTemplate) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *WebRARevokeRequestOnTemplate) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


