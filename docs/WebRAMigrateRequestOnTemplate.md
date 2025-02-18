# WebRAMigrateRequestOnTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** |  | 
**CertificateId** | Pointer to **NullableString** | Used to pre-fill the template field with the certificate values. | [optional] 
**CertificatePem** | Pointer to **NullableString** | Used to pre-fill the template field with the certificate values. | [optional] 
**Module** | Pointer to **string** | The request module | [optional] 
**Profile** | **string** | The profile for which to return the template. | 

## Methods

### NewWebRAMigrateRequestOnTemplate

`func NewWebRAMigrateRequestOnTemplate(workflow string, profile string, ) *WebRAMigrateRequestOnTemplate`

NewWebRAMigrateRequestOnTemplate instantiates a new WebRAMigrateRequestOnTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAMigrateRequestOnTemplateWithDefaults

`func NewWebRAMigrateRequestOnTemplateWithDefaults() *WebRAMigrateRequestOnTemplate`

NewWebRAMigrateRequestOnTemplateWithDefaults instantiates a new WebRAMigrateRequestOnTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WebRAMigrateRequestOnTemplate) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAMigrateRequestOnTemplate) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAMigrateRequestOnTemplate) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetCertificateId

`func (o *WebRAMigrateRequestOnTemplate) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *WebRAMigrateRequestOnTemplate) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *WebRAMigrateRequestOnTemplate) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *WebRAMigrateRequestOnTemplate) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *WebRAMigrateRequestOnTemplate) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *WebRAMigrateRequestOnTemplate) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCertificatePem

`func (o *WebRAMigrateRequestOnTemplate) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *WebRAMigrateRequestOnTemplate) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *WebRAMigrateRequestOnTemplate) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *WebRAMigrateRequestOnTemplate) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### SetCertificatePemNil

`func (o *WebRAMigrateRequestOnTemplate) SetCertificatePemNil(b bool)`

 SetCertificatePemNil sets the value for CertificatePem to be an explicit nil

### UnsetCertificatePem
`func (o *WebRAMigrateRequestOnTemplate) UnsetCertificatePem()`

UnsetCertificatePem ensures that no value is present for CertificatePem, not even an explicit nil
### GetModule

`func (o *WebRAMigrateRequestOnTemplate) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAMigrateRequestOnTemplate) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAMigrateRequestOnTemplate) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *WebRAMigrateRequestOnTemplate) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetProfile

`func (o *WebRAMigrateRequestOnTemplate) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAMigrateRequestOnTemplate) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAMigrateRequestOnTemplate) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


