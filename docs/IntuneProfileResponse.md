# IntuneProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AuthorizationLevels** | [**CertificateProfileAuthorizationLevels**](CertificateProfileAuthorizationLevels.md) |  | 
**Caps** | **[]string** |  | 
**CertificateTemplate** | Pointer to [**NullableCertificateTemplate**](CertificateTemplate.md) |  | [optional] 
**Constraints** | Pointer to [**NullableCertificateRequestConstraints**](CertificateRequestConstraints.md) |  | [optional] 
**CryptoPolicy** | [**ManagedCertificateProfileCryptoPolicy**](ManagedCertificateProfileCryptoPolicy.md) |  | 
**CsrDataMapping** | Pointer to **map[string]string** |  | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**DeviceIdField** | Pointer to **NullableString** |  | [optional] 
**DeviceIdSeparator** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**DsFlow** | Pointer to [**[]DataSourceFlowEntry**](DataSourceFlowEntry.md) | Representation of a datasource execution flow | [optional] 
**Enabled** | **bool** |  | 
**EncryptionAlgorithm** | **string** |  | 
**GradingPolicies** | Pointer to **[]string** |  | [optional] 
**MaxCertificatePerHolderPolicy** | Pointer to [**NullableMaxCertificatePerHolderPolicy**](MaxCertificatePerHolderPolicy.md) |  | [optional] 
**Mode** | **string** |  | 
**Module** | **string** |  | 
**Name** | **string** |  | 
**PkiConnector** | **string** |  | 
**PostPKIOperation** | Pointer to **NullableBool** |  | [optional] 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**RequestsPolicy** | [**RequestsPolicy**](RequestsPolicy.md) |  | 
**ScepRA** | **string** |  | 
**SelfPermissions** | [**CertificateProfileSelfPermissions**](CertificateProfileSelfPermissions.md) |  | 
**ThirdPartyConnector** | **string** |  | 
**ThirdPartyDiscoverySync** | Pointer to **NullableBool** | Available from &#x60;2.8.2&#x60; | [optional] [default to false]
**Triggers** | Pointer to [**NullableCertificateProfileTriggers**](CertificateProfileTriggers.md) |  | [optional] 

## Methods

### NewIntuneProfileResponse

`func NewIntuneProfileResponse(id string, authorizationLevels CertificateProfileAuthorizationLevels, caps []string, cryptoPolicy ManagedCertificateProfileCryptoPolicy, enabled bool, encryptionAlgorithm string, mode string, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, scepRA string, selfPermissions CertificateProfileSelfPermissions, thirdPartyConnector string, ) *IntuneProfileResponse`

NewIntuneProfileResponse instantiates a new IntuneProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntuneProfileResponseWithDefaults

`func NewIntuneProfileResponseWithDefaults() *IntuneProfileResponse`

NewIntuneProfileResponseWithDefaults instantiates a new IntuneProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntuneProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntuneProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntuneProfileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAuthorizationLevels

`func (o *IntuneProfileResponse) GetAuthorizationLevels() CertificateProfileAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *IntuneProfileResponse) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *IntuneProfileResponse) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetCaps

`func (o *IntuneProfileResponse) GetCaps() []string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *IntuneProfileResponse) GetCapsOk() (*[]string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *IntuneProfileResponse) SetCaps(v []string)`

SetCaps sets Caps field to given value.


### GetCertificateTemplate

`func (o *IntuneProfileResponse) GetCertificateTemplate() CertificateTemplate`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *IntuneProfileResponse) GetCertificateTemplateOk() (*CertificateTemplate, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *IntuneProfileResponse) SetCertificateTemplate(v CertificateTemplate)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *IntuneProfileResponse) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *IntuneProfileResponse) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *IntuneProfileResponse) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetConstraints

`func (o *IntuneProfileResponse) GetConstraints() CertificateRequestConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *IntuneProfileResponse) GetConstraintsOk() (*CertificateRequestConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *IntuneProfileResponse) SetConstraints(v CertificateRequestConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *IntuneProfileResponse) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *IntuneProfileResponse) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *IntuneProfileResponse) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetCryptoPolicy

`func (o *IntuneProfileResponse) GetCryptoPolicy() ManagedCertificateProfileCryptoPolicy`

GetCryptoPolicy returns the CryptoPolicy field if non-nil, zero value otherwise.

### GetCryptoPolicyOk

`func (o *IntuneProfileResponse) GetCryptoPolicyOk() (*ManagedCertificateProfileCryptoPolicy, bool)`

GetCryptoPolicyOk returns a tuple with the CryptoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPolicy

`func (o *IntuneProfileResponse) SetCryptoPolicy(v ManagedCertificateProfileCryptoPolicy)`

SetCryptoPolicy sets CryptoPolicy field to given value.


### GetCsrDataMapping

`func (o *IntuneProfileResponse) GetCsrDataMapping() map[string]string`

GetCsrDataMapping returns the CsrDataMapping field if non-nil, zero value otherwise.

### GetCsrDataMappingOk

`func (o *IntuneProfileResponse) GetCsrDataMappingOk() (*map[string]string, bool)`

GetCsrDataMappingOk returns a tuple with the CsrDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataMapping

`func (o *IntuneProfileResponse) SetCsrDataMapping(v map[string]string)`

SetCsrDataMapping sets CsrDataMapping field to given value.

### HasCsrDataMapping

`func (o *IntuneProfileResponse) HasCsrDataMapping() bool`

HasCsrDataMapping returns a boolean if a field has been set.

### SetCsrDataMappingNil

`func (o *IntuneProfileResponse) SetCsrDataMappingNil(b bool)`

 SetCsrDataMappingNil sets the value for CsrDataMapping to be an explicit nil

### UnsetCsrDataMapping
`func (o *IntuneProfileResponse) UnsetCsrDataMapping()`

UnsetCsrDataMapping ensures that no value is present for CsrDataMapping, not even an explicit nil
### GetDescription

`func (o *IntuneProfileResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntuneProfileResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntuneProfileResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntuneProfileResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IntuneProfileResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IntuneProfileResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDeviceIdField

`func (o *IntuneProfileResponse) GetDeviceIdField() string`

GetDeviceIdField returns the DeviceIdField field if non-nil, zero value otherwise.

### GetDeviceIdFieldOk

`func (o *IntuneProfileResponse) GetDeviceIdFieldOk() (*string, bool)`

GetDeviceIdFieldOk returns a tuple with the DeviceIdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdField

`func (o *IntuneProfileResponse) SetDeviceIdField(v string)`

SetDeviceIdField sets DeviceIdField field to given value.

### HasDeviceIdField

`func (o *IntuneProfileResponse) HasDeviceIdField() bool`

HasDeviceIdField returns a boolean if a field has been set.

### SetDeviceIdFieldNil

`func (o *IntuneProfileResponse) SetDeviceIdFieldNil(b bool)`

 SetDeviceIdFieldNil sets the value for DeviceIdField to be an explicit nil

### UnsetDeviceIdField
`func (o *IntuneProfileResponse) UnsetDeviceIdField()`

UnsetDeviceIdField ensures that no value is present for DeviceIdField, not even an explicit nil
### GetDeviceIdSeparator

`func (o *IntuneProfileResponse) GetDeviceIdSeparator() string`

GetDeviceIdSeparator returns the DeviceIdSeparator field if non-nil, zero value otherwise.

### GetDeviceIdSeparatorOk

`func (o *IntuneProfileResponse) GetDeviceIdSeparatorOk() (*string, bool)`

GetDeviceIdSeparatorOk returns a tuple with the DeviceIdSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdSeparator

`func (o *IntuneProfileResponse) SetDeviceIdSeparator(v string)`

SetDeviceIdSeparator sets DeviceIdSeparator field to given value.

### HasDeviceIdSeparator

`func (o *IntuneProfileResponse) HasDeviceIdSeparator() bool`

HasDeviceIdSeparator returns a boolean if a field has been set.

### SetDeviceIdSeparatorNil

`func (o *IntuneProfileResponse) SetDeviceIdSeparatorNil(b bool)`

 SetDeviceIdSeparatorNil sets the value for DeviceIdSeparator to be an explicit nil

### UnsetDeviceIdSeparator
`func (o *IntuneProfileResponse) UnsetDeviceIdSeparator()`

UnsetDeviceIdSeparator ensures that no value is present for DeviceIdSeparator, not even an explicit nil
### GetDisplayName

`func (o *IntuneProfileResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IntuneProfileResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IntuneProfileResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IntuneProfileResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IntuneProfileResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IntuneProfileResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDsFlow

`func (o *IntuneProfileResponse) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *IntuneProfileResponse) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *IntuneProfileResponse) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *IntuneProfileResponse) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### SetDsFlowNil

`func (o *IntuneProfileResponse) SetDsFlowNil(b bool)`

 SetDsFlowNil sets the value for DsFlow to be an explicit nil

### UnsetDsFlow
`func (o *IntuneProfileResponse) UnsetDsFlow()`

UnsetDsFlow ensures that no value is present for DsFlow, not even an explicit nil
### GetEnabled

`func (o *IntuneProfileResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntuneProfileResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntuneProfileResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEncryptionAlgorithm

`func (o *IntuneProfileResponse) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *IntuneProfileResponse) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *IntuneProfileResponse) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.


### GetGradingPolicies

`func (o *IntuneProfileResponse) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *IntuneProfileResponse) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *IntuneProfileResponse) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *IntuneProfileResponse) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *IntuneProfileResponse) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *IntuneProfileResponse) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetMaxCertificatePerHolderPolicy

`func (o *IntuneProfileResponse) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy`

GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field if non-nil, zero value otherwise.

### GetMaxCertificatePerHolderPolicyOk

`func (o *IntuneProfileResponse) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool)`

GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCertificatePerHolderPolicy

`func (o *IntuneProfileResponse) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy)`

SetMaxCertificatePerHolderPolicy sets MaxCertificatePerHolderPolicy field to given value.

### HasMaxCertificatePerHolderPolicy

`func (o *IntuneProfileResponse) HasMaxCertificatePerHolderPolicy() bool`

HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.

### SetMaxCertificatePerHolderPolicyNil

`func (o *IntuneProfileResponse) SetMaxCertificatePerHolderPolicyNil(b bool)`

 SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil

### UnsetMaxCertificatePerHolderPolicy
`func (o *IntuneProfileResponse) UnsetMaxCertificatePerHolderPolicy()`

UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
### GetMode

`func (o *IntuneProfileResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IntuneProfileResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IntuneProfileResponse) SetMode(v string)`

SetMode sets Mode field to given value.


### GetModule

`func (o *IntuneProfileResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *IntuneProfileResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *IntuneProfileResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetName

`func (o *IntuneProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntuneProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntuneProfileResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPkiConnector

`func (o *IntuneProfileResponse) GetPkiConnector() string`

GetPkiConnector returns the PkiConnector field if non-nil, zero value otherwise.

### GetPkiConnectorOk

`func (o *IntuneProfileResponse) GetPkiConnectorOk() (*string, bool)`

GetPkiConnectorOk returns a tuple with the PkiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnector

`func (o *IntuneProfileResponse) SetPkiConnector(v string)`

SetPkiConnector sets PkiConnector field to given value.


### GetPostPKIOperation

`func (o *IntuneProfileResponse) GetPostPKIOperation() bool`

GetPostPKIOperation returns the PostPKIOperation field if non-nil, zero value otherwise.

### GetPostPKIOperationOk

`func (o *IntuneProfileResponse) GetPostPKIOperationOk() (*bool, bool)`

GetPostPKIOperationOk returns a tuple with the PostPKIOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPKIOperation

`func (o *IntuneProfileResponse) SetPostPKIOperation(v bool)`

SetPostPKIOperation sets PostPKIOperation field to given value.

### HasPostPKIOperation

`func (o *IntuneProfileResponse) HasPostPKIOperation() bool`

HasPostPKIOperation returns a boolean if a field has been set.

### SetPostPKIOperationNil

`func (o *IntuneProfileResponse) SetPostPKIOperationNil(b bool)`

 SetPostPKIOperationNil sets the value for PostPKIOperation to be an explicit nil

### UnsetPostPKIOperation
`func (o *IntuneProfileResponse) UnsetPostPKIOperation()`

UnsetPostPKIOperation ensures that no value is present for PostPKIOperation, not even an explicit nil
### GetRenewalPeriod

`func (o *IntuneProfileResponse) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *IntuneProfileResponse) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *IntuneProfileResponse) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *IntuneProfileResponse) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *IntuneProfileResponse) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *IntuneProfileResponse) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetRequestsPolicy

`func (o *IntuneProfileResponse) GetRequestsPolicy() RequestsPolicy`

GetRequestsPolicy returns the RequestsPolicy field if non-nil, zero value otherwise.

### GetRequestsPolicyOk

`func (o *IntuneProfileResponse) GetRequestsPolicyOk() (*RequestsPolicy, bool)`

GetRequestsPolicyOk returns a tuple with the RequestsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPolicy

`func (o *IntuneProfileResponse) SetRequestsPolicy(v RequestsPolicy)`

SetRequestsPolicy sets RequestsPolicy field to given value.


### GetScepRA

`func (o *IntuneProfileResponse) GetScepRA() string`

GetScepRA returns the ScepRA field if non-nil, zero value otherwise.

### GetScepRAOk

`func (o *IntuneProfileResponse) GetScepRAOk() (*string, bool)`

GetScepRAOk returns a tuple with the ScepRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScepRA

`func (o *IntuneProfileResponse) SetScepRA(v string)`

SetScepRA sets ScepRA field to given value.


### GetSelfPermissions

`func (o *IntuneProfileResponse) GetSelfPermissions() CertificateProfileSelfPermissions`

GetSelfPermissions returns the SelfPermissions field if non-nil, zero value otherwise.

### GetSelfPermissionsOk

`func (o *IntuneProfileResponse) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool)`

GetSelfPermissionsOk returns a tuple with the SelfPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPermissions

`func (o *IntuneProfileResponse) SetSelfPermissions(v CertificateProfileSelfPermissions)`

SetSelfPermissions sets SelfPermissions field to given value.


### GetThirdPartyConnector

`func (o *IntuneProfileResponse) GetThirdPartyConnector() string`

GetThirdPartyConnector returns the ThirdPartyConnector field if non-nil, zero value otherwise.

### GetThirdPartyConnectorOk

`func (o *IntuneProfileResponse) GetThirdPartyConnectorOk() (*string, bool)`

GetThirdPartyConnectorOk returns a tuple with the ThirdPartyConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyConnector

`func (o *IntuneProfileResponse) SetThirdPartyConnector(v string)`

SetThirdPartyConnector sets ThirdPartyConnector field to given value.


### GetThirdPartyDiscoverySync

`func (o *IntuneProfileResponse) GetThirdPartyDiscoverySync() bool`

GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field if non-nil, zero value otherwise.

### GetThirdPartyDiscoverySyncOk

`func (o *IntuneProfileResponse) GetThirdPartyDiscoverySyncOk() (*bool, bool)`

GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyDiscoverySync

`func (o *IntuneProfileResponse) SetThirdPartyDiscoverySync(v bool)`

SetThirdPartyDiscoverySync sets ThirdPartyDiscoverySync field to given value.

### HasThirdPartyDiscoverySync

`func (o *IntuneProfileResponse) HasThirdPartyDiscoverySync() bool`

HasThirdPartyDiscoverySync returns a boolean if a field has been set.

### SetThirdPartyDiscoverySyncNil

`func (o *IntuneProfileResponse) SetThirdPartyDiscoverySyncNil(b bool)`

 SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil

### UnsetThirdPartyDiscoverySync
`func (o *IntuneProfileResponse) UnsetThirdPartyDiscoverySync()`

UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
### GetTriggers

`func (o *IntuneProfileResponse) GetTriggers() CertificateProfileTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *IntuneProfileResponse) GetTriggersOk() (*CertificateProfileTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *IntuneProfileResponse) SetTriggers(v CertificateProfileTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *IntuneProfileResponse) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *IntuneProfileResponse) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *IntuneProfileResponse) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


