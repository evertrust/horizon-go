# WebRAImportRequestOnTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** |  | 
**Template** | Pointer to **map[string]interface{}** | Test whether the private key can be imported on the certificate | [optional] 
**CertificateId** | Pointer to **NullableString** | Used to pre-fill the template field with the certificate values. | [optional] 
**CertificatePem** | Pointer to **NullableString** | Used to pre-fill the template field with the certificate values. | [optional] 
**Module** | Pointer to **string** | The request module | [optional] 
**Profile** | **string** | The profile for which to return the template. | 

## Methods

### NewWebRAImportRequestOnTemplate

`func NewWebRAImportRequestOnTemplate(workflow string, profile string, ) *WebRAImportRequestOnTemplate`

NewWebRAImportRequestOnTemplate instantiates a new WebRAImportRequestOnTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAImportRequestOnTemplateWithDefaults

`func NewWebRAImportRequestOnTemplateWithDefaults() *WebRAImportRequestOnTemplate`

NewWebRAImportRequestOnTemplateWithDefaults instantiates a new WebRAImportRequestOnTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WebRAImportRequestOnTemplate) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WebRAImportRequestOnTemplate) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WebRAImportRequestOnTemplate) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetTemplate

`func (o *WebRAImportRequestOnTemplate) GetTemplate() map[string]interface{}`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebRAImportRequestOnTemplate) GetTemplateOk() (*map[string]interface{}, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebRAImportRequestOnTemplate) SetTemplate(v map[string]interface{})`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebRAImportRequestOnTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetCertificateId

`func (o *WebRAImportRequestOnTemplate) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *WebRAImportRequestOnTemplate) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *WebRAImportRequestOnTemplate) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *WebRAImportRequestOnTemplate) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *WebRAImportRequestOnTemplate) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *WebRAImportRequestOnTemplate) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCertificatePem

`func (o *WebRAImportRequestOnTemplate) GetCertificatePem() string`

GetCertificatePem returns the CertificatePem field if non-nil, zero value otherwise.

### GetCertificatePemOk

`func (o *WebRAImportRequestOnTemplate) GetCertificatePemOk() (*string, bool)`

GetCertificatePemOk returns a tuple with the CertificatePem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePem

`func (o *WebRAImportRequestOnTemplate) SetCertificatePem(v string)`

SetCertificatePem sets CertificatePem field to given value.

### HasCertificatePem

`func (o *WebRAImportRequestOnTemplate) HasCertificatePem() bool`

HasCertificatePem returns a boolean if a field has been set.

### SetCertificatePemNil

`func (o *WebRAImportRequestOnTemplate) SetCertificatePemNil(b bool)`

 SetCertificatePemNil sets the value for CertificatePem to be an explicit nil

### UnsetCertificatePem
`func (o *WebRAImportRequestOnTemplate) UnsetCertificatePem()`

UnsetCertificatePem ensures that no value is present for CertificatePem, not even an explicit nil
### GetModule

`func (o *WebRAImportRequestOnTemplate) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebRAImportRequestOnTemplate) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebRAImportRequestOnTemplate) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *WebRAImportRequestOnTemplate) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetProfile

`func (o *WebRAImportRequestOnTemplate) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebRAImportRequestOnTemplate) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebRAImportRequestOnTemplate) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


