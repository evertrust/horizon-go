# RequestableCertificateProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the profile | 
**Module** | **string** | The module of the profile | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the profile | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized description of the profile | [optional] 
**Workflows** | [**[]RequestableWorkflow**](RequestableWorkflow.md) | The list of workflows on this profile and the principals permission on them | 
**Capabilities** | [**CertificateProfileCapabilities**](CertificateProfileCapabilities.md) | Crypto information about this profile | 

## Methods

### NewRequestableCertificateProfileResponse

`func NewRequestableCertificateProfileResponse(name string, module string, workflows []RequestableWorkflow, capabilities CertificateProfileCapabilities, ) *RequestableCertificateProfileResponse`

NewRequestableCertificateProfileResponse instantiates a new RequestableCertificateProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestableCertificateProfileResponseWithDefaults

`func NewRequestableCertificateProfileResponseWithDefaults() *RequestableCertificateProfileResponse`

NewRequestableCertificateProfileResponseWithDefaults instantiates a new RequestableCertificateProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestableCertificateProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestableCertificateProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestableCertificateProfileResponse) SetName(v string)`

SetName sets Name field to given value.


### GetModule

`func (o *RequestableCertificateProfileResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RequestableCertificateProfileResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RequestableCertificateProfileResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetDisplayName

`func (o *RequestableCertificateProfileResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RequestableCertificateProfileResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RequestableCertificateProfileResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RequestableCertificateProfileResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *RequestableCertificateProfileResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *RequestableCertificateProfileResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *RequestableCertificateProfileResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestableCertificateProfileResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestableCertificateProfileResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestableCertificateProfileResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RequestableCertificateProfileResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RequestableCertificateProfileResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWorkflows

`func (o *RequestableCertificateProfileResponse) GetWorkflows() []RequestableWorkflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *RequestableCertificateProfileResponse) GetWorkflowsOk() (*[]RequestableWorkflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *RequestableCertificateProfileResponse) SetWorkflows(v []RequestableWorkflow)`

SetWorkflows sets Workflows field to given value.


### GetCapabilities

`func (o *RequestableCertificateProfileResponse) GetCapabilities() CertificateProfileCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *RequestableCertificateProfileResponse) GetCapabilitiesOk() (*CertificateProfileCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *RequestableCertificateProfileResponse) SetCapabilities(v CertificateProfileCapabilities)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


