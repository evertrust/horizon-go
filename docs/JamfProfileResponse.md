# JamfProfileResponse

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
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**DsFlow** | Pointer to [**[]DataSourceFlowEntry**](DataSourceFlowEntry.md) | Representation of a datasource execution flow | [optional] 
**Enabled** | **bool** |  | 
**EncryptionAlgorithm** | **string** |  | 
**GradingPolicies** | Pointer to **[]string** |  | [optional] 
**MaxCertificatePerHolderPolicy** | Pointer to [**NullableMaxCertificatePerHolderPolicy**](MaxCertificatePerHolderPolicy.md) |  | [optional] 
**Mode** | **string** |  | 
**Module** | **string** |  | 
**Name** | **string** |  | 
**PasswordPolicy** | Pointer to **NullableString** |  | [optional] 
**PkiConnector** | **string** |  | 
**PostPKIOperation** | Pointer to **NullableBool** |  | [optional] 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**RequestsPolicy** | [**RequestsPolicy**](RequestsPolicy.md) |  | 
**ScepRA** | **string** |  | 
**SelfPermissions** | [**CertificateProfileSelfPermissions**](CertificateProfileSelfPermissions.md) |  | 
**ThirdPartyConnector** | **string** |  | 
**ThirdPartyDiscoverySync** | Pointer to **NullableBool** |  | [optional] [default to false]
**Triggers** | Pointer to [**NullableCertificateProfileTriggers**](CertificateProfileTriggers.md) |  | [optional] 

## Methods

### NewJamfProfileResponse

`func NewJamfProfileResponse(id string, authorizationLevels CertificateProfileAuthorizationLevels, caps []string, cryptoPolicy ManagedCertificateProfileCryptoPolicy, enabled bool, encryptionAlgorithm string, mode string, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, scepRA string, selfPermissions CertificateProfileSelfPermissions, thirdPartyConnector string, ) *JamfProfileResponse`

NewJamfProfileResponse instantiates a new JamfProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJamfProfileResponseWithDefaults

`func NewJamfProfileResponseWithDefaults() *JamfProfileResponse`

NewJamfProfileResponseWithDefaults instantiates a new JamfProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JamfProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JamfProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JamfProfileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAuthorizationLevels

`func (o *JamfProfileResponse) GetAuthorizationLevels() CertificateProfileAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *JamfProfileResponse) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *JamfProfileResponse) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetCaps

`func (o *JamfProfileResponse) GetCaps() []string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *JamfProfileResponse) GetCapsOk() (*[]string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *JamfProfileResponse) SetCaps(v []string)`

SetCaps sets Caps field to given value.


### GetCertificateTemplate

`func (o *JamfProfileResponse) GetCertificateTemplate() CertificateTemplate`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *JamfProfileResponse) GetCertificateTemplateOk() (*CertificateTemplate, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *JamfProfileResponse) SetCertificateTemplate(v CertificateTemplate)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *JamfProfileResponse) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *JamfProfileResponse) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *JamfProfileResponse) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetConstraints

`func (o *JamfProfileResponse) GetConstraints() CertificateRequestConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *JamfProfileResponse) GetConstraintsOk() (*CertificateRequestConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *JamfProfileResponse) SetConstraints(v CertificateRequestConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *JamfProfileResponse) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *JamfProfileResponse) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *JamfProfileResponse) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetCryptoPolicy

`func (o *JamfProfileResponse) GetCryptoPolicy() ManagedCertificateProfileCryptoPolicy`

GetCryptoPolicy returns the CryptoPolicy field if non-nil, zero value otherwise.

### GetCryptoPolicyOk

`func (o *JamfProfileResponse) GetCryptoPolicyOk() (*ManagedCertificateProfileCryptoPolicy, bool)`

GetCryptoPolicyOk returns a tuple with the CryptoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPolicy

`func (o *JamfProfileResponse) SetCryptoPolicy(v ManagedCertificateProfileCryptoPolicy)`

SetCryptoPolicy sets CryptoPolicy field to given value.


### GetCsrDataMapping

`func (o *JamfProfileResponse) GetCsrDataMapping() map[string]string`

GetCsrDataMapping returns the CsrDataMapping field if non-nil, zero value otherwise.

### GetCsrDataMappingOk

`func (o *JamfProfileResponse) GetCsrDataMappingOk() (*map[string]string, bool)`

GetCsrDataMappingOk returns a tuple with the CsrDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataMapping

`func (o *JamfProfileResponse) SetCsrDataMapping(v map[string]string)`

SetCsrDataMapping sets CsrDataMapping field to given value.

### HasCsrDataMapping

`func (o *JamfProfileResponse) HasCsrDataMapping() bool`

HasCsrDataMapping returns a boolean if a field has been set.

### SetCsrDataMappingNil

`func (o *JamfProfileResponse) SetCsrDataMappingNil(b bool)`

 SetCsrDataMappingNil sets the value for CsrDataMapping to be an explicit nil

### UnsetCsrDataMapping
`func (o *JamfProfileResponse) UnsetCsrDataMapping()`

UnsetCsrDataMapping ensures that no value is present for CsrDataMapping, not even an explicit nil
### GetDescription

`func (o *JamfProfileResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JamfProfileResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JamfProfileResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JamfProfileResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JamfProfileResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JamfProfileResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDeviceIdField

`func (o *JamfProfileResponse) GetDeviceIdField() string`

GetDeviceIdField returns the DeviceIdField field if non-nil, zero value otherwise.

### GetDeviceIdFieldOk

`func (o *JamfProfileResponse) GetDeviceIdFieldOk() (*string, bool)`

GetDeviceIdFieldOk returns a tuple with the DeviceIdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdField

`func (o *JamfProfileResponse) SetDeviceIdField(v string)`

SetDeviceIdField sets DeviceIdField field to given value.

### HasDeviceIdField

`func (o *JamfProfileResponse) HasDeviceIdField() bool`

HasDeviceIdField returns a boolean if a field has been set.

### SetDeviceIdFieldNil

`func (o *JamfProfileResponse) SetDeviceIdFieldNil(b bool)`

 SetDeviceIdFieldNil sets the value for DeviceIdField to be an explicit nil

### UnsetDeviceIdField
`func (o *JamfProfileResponse) UnsetDeviceIdField()`

UnsetDeviceIdField ensures that no value is present for DeviceIdField, not even an explicit nil
### GetDisplayName

`func (o *JamfProfileResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *JamfProfileResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *JamfProfileResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *JamfProfileResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *JamfProfileResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *JamfProfileResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDsFlow

`func (o *JamfProfileResponse) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *JamfProfileResponse) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *JamfProfileResponse) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *JamfProfileResponse) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### SetDsFlowNil

`func (o *JamfProfileResponse) SetDsFlowNil(b bool)`

 SetDsFlowNil sets the value for DsFlow to be an explicit nil

### UnsetDsFlow
`func (o *JamfProfileResponse) UnsetDsFlow()`

UnsetDsFlow ensures that no value is present for DsFlow, not even an explicit nil
### GetEnabled

`func (o *JamfProfileResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JamfProfileResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JamfProfileResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEncryptionAlgorithm

`func (o *JamfProfileResponse) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *JamfProfileResponse) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *JamfProfileResponse) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.


### GetGradingPolicies

`func (o *JamfProfileResponse) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *JamfProfileResponse) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *JamfProfileResponse) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *JamfProfileResponse) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *JamfProfileResponse) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *JamfProfileResponse) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetMaxCertificatePerHolderPolicy

`func (o *JamfProfileResponse) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy`

GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field if non-nil, zero value otherwise.

### GetMaxCertificatePerHolderPolicyOk

`func (o *JamfProfileResponse) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool)`

GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCertificatePerHolderPolicy

`func (o *JamfProfileResponse) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy)`

SetMaxCertificatePerHolderPolicy sets MaxCertificatePerHolderPolicy field to given value.

### HasMaxCertificatePerHolderPolicy

`func (o *JamfProfileResponse) HasMaxCertificatePerHolderPolicy() bool`

HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.

### SetMaxCertificatePerHolderPolicyNil

`func (o *JamfProfileResponse) SetMaxCertificatePerHolderPolicyNil(b bool)`

 SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil

### UnsetMaxCertificatePerHolderPolicy
`func (o *JamfProfileResponse) UnsetMaxCertificatePerHolderPolicy()`

UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
### GetMode

`func (o *JamfProfileResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *JamfProfileResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *JamfProfileResponse) SetMode(v string)`

SetMode sets Mode field to given value.


### GetModule

`func (o *JamfProfileResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *JamfProfileResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *JamfProfileResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetName

`func (o *JamfProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JamfProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JamfProfileResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordPolicy

`func (o *JamfProfileResponse) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *JamfProfileResponse) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *JamfProfileResponse) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *JamfProfileResponse) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *JamfProfileResponse) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *JamfProfileResponse) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetPkiConnector

`func (o *JamfProfileResponse) GetPkiConnector() string`

GetPkiConnector returns the PkiConnector field if non-nil, zero value otherwise.

### GetPkiConnectorOk

`func (o *JamfProfileResponse) GetPkiConnectorOk() (*string, bool)`

GetPkiConnectorOk returns a tuple with the PkiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnector

`func (o *JamfProfileResponse) SetPkiConnector(v string)`

SetPkiConnector sets PkiConnector field to given value.


### GetPostPKIOperation

`func (o *JamfProfileResponse) GetPostPKIOperation() bool`

GetPostPKIOperation returns the PostPKIOperation field if non-nil, zero value otherwise.

### GetPostPKIOperationOk

`func (o *JamfProfileResponse) GetPostPKIOperationOk() (*bool, bool)`

GetPostPKIOperationOk returns a tuple with the PostPKIOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPKIOperation

`func (o *JamfProfileResponse) SetPostPKIOperation(v bool)`

SetPostPKIOperation sets PostPKIOperation field to given value.

### HasPostPKIOperation

`func (o *JamfProfileResponse) HasPostPKIOperation() bool`

HasPostPKIOperation returns a boolean if a field has been set.

### SetPostPKIOperationNil

`func (o *JamfProfileResponse) SetPostPKIOperationNil(b bool)`

 SetPostPKIOperationNil sets the value for PostPKIOperation to be an explicit nil

### UnsetPostPKIOperation
`func (o *JamfProfileResponse) UnsetPostPKIOperation()`

UnsetPostPKIOperation ensures that no value is present for PostPKIOperation, not even an explicit nil
### GetRenewalPeriod

`func (o *JamfProfileResponse) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *JamfProfileResponse) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *JamfProfileResponse) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *JamfProfileResponse) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *JamfProfileResponse) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *JamfProfileResponse) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetRequestsPolicy

`func (o *JamfProfileResponse) GetRequestsPolicy() RequestsPolicy`

GetRequestsPolicy returns the RequestsPolicy field if non-nil, zero value otherwise.

### GetRequestsPolicyOk

`func (o *JamfProfileResponse) GetRequestsPolicyOk() (*RequestsPolicy, bool)`

GetRequestsPolicyOk returns a tuple with the RequestsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPolicy

`func (o *JamfProfileResponse) SetRequestsPolicy(v RequestsPolicy)`

SetRequestsPolicy sets RequestsPolicy field to given value.


### GetScepRA

`func (o *JamfProfileResponse) GetScepRA() string`

GetScepRA returns the ScepRA field if non-nil, zero value otherwise.

### GetScepRAOk

`func (o *JamfProfileResponse) GetScepRAOk() (*string, bool)`

GetScepRAOk returns a tuple with the ScepRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScepRA

`func (o *JamfProfileResponse) SetScepRA(v string)`

SetScepRA sets ScepRA field to given value.


### GetSelfPermissions

`func (o *JamfProfileResponse) GetSelfPermissions() CertificateProfileSelfPermissions`

GetSelfPermissions returns the SelfPermissions field if non-nil, zero value otherwise.

### GetSelfPermissionsOk

`func (o *JamfProfileResponse) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool)`

GetSelfPermissionsOk returns a tuple with the SelfPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPermissions

`func (o *JamfProfileResponse) SetSelfPermissions(v CertificateProfileSelfPermissions)`

SetSelfPermissions sets SelfPermissions field to given value.


### GetThirdPartyConnector

`func (o *JamfProfileResponse) GetThirdPartyConnector() string`

GetThirdPartyConnector returns the ThirdPartyConnector field if non-nil, zero value otherwise.

### GetThirdPartyConnectorOk

`func (o *JamfProfileResponse) GetThirdPartyConnectorOk() (*string, bool)`

GetThirdPartyConnectorOk returns a tuple with the ThirdPartyConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyConnector

`func (o *JamfProfileResponse) SetThirdPartyConnector(v string)`

SetThirdPartyConnector sets ThirdPartyConnector field to given value.


### GetThirdPartyDiscoverySync

`func (o *JamfProfileResponse) GetThirdPartyDiscoverySync() bool`

GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field if non-nil, zero value otherwise.

### GetThirdPartyDiscoverySyncOk

`func (o *JamfProfileResponse) GetThirdPartyDiscoverySyncOk() (*bool, bool)`

GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyDiscoverySync

`func (o *JamfProfileResponse) SetThirdPartyDiscoverySync(v bool)`

SetThirdPartyDiscoverySync sets ThirdPartyDiscoverySync field to given value.

### HasThirdPartyDiscoverySync

`func (o *JamfProfileResponse) HasThirdPartyDiscoverySync() bool`

HasThirdPartyDiscoverySync returns a boolean if a field has been set.

### SetThirdPartyDiscoverySyncNil

`func (o *JamfProfileResponse) SetThirdPartyDiscoverySyncNil(b bool)`

 SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil

### UnsetThirdPartyDiscoverySync
`func (o *JamfProfileResponse) UnsetThirdPartyDiscoverySync()`

UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
### GetTriggers

`func (o *JamfProfileResponse) GetTriggers() CertificateProfileTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *JamfProfileResponse) GetTriggersOk() (*CertificateProfileTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *JamfProfileResponse) SetTriggers(v CertificateProfileTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *JamfProfileResponse) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *JamfProfileResponse) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *JamfProfileResponse) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


