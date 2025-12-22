# WebRARecoverRequestOnTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** |  | 
**CertificateId** | Pointer to **NullableString** | Used to pre-fill the template field with the certificate values. | [optional] 
**CertificatePem** | Pointer to **NullableString** | Used to pre-fill the template field with the certificate values. | [optional] 
**Module** | Pointer to **string** | The request module | [optional] 
**Profile** | **string** | The profile for which to return the template. | 

## Methods

### NewWebRARecoverRequestOnTemplate

`func NewWebRARecoverRequestOnTemplate(workflow string, profile string, ) *WebRARecoverRequestOnTemplate`

NewWebRARecoverRequestOnTemplate instantiates a new WebRARecoverRequestOnTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARecoverRequestOnTemplateWithDefaults

`func NewWebRARecoverRequestOnTemplateWithDefaults() *WebRARecoverRequestOnTemplate`

NewWebRARecoverRequestOnTemplateWithDefaults instantiates a new WebRARecoverRequestOnTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WebRARecoverRequestOnTemplate) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRARecoverRequestOnTemplate) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRARecoverRequestOnTemplate) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetCertificateId

`func (o *WebRARecoverRequestOnTemplate) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *WebRARecoverRequestOnTemplate) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *WebRARecoverRequestOnTemplate) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *WebRARecoverRequestOnTemplate) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *WebRARecoverRequestOnTemplate) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *WebRARecoverRequestOnTemplate) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCertificatePem

`func (o *WebRARecoverRequestOnTemplate) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *WebRARecoverRequestOnTemplate) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *WebRARecoverRequestOnTemplate) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *WebRARecoverRequestOnTemplate) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### SetCertificatePemNil

`func (o *WebRARecoverRequestOnTemplate) SetCertificatePemNil(b bool)`

 SetCertificatePemNil sets the value for CertificatePem to be an explicit nil

### UnsetCertificatePem
`func (o *WebRARecoverRequestOnTemplate) UnsetCertificatePem()`

UnsetCertificatePem ensures that no value is present for CertificatePem, not even an explicit nil
### GetModule

`func (o *WebRARecoverRequestOnTemplate) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRARecoverRequestOnTemplate) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRARecoverRequestOnTemplate) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *WebRARecoverRequestOnTemplate) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetProfile

`func (o *WebRARecoverRequestOnTemplate) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRARecoverRequestOnTemplate) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRARecoverRequestOnTemplate) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


