# IntunePKCSProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationLevels** | [**CertificateProfileAuthorizationLevels**](CertificateProfileAuthorizationLevels.md) |  | 
**CertificateTemplate** | Pointer to [**NullableCertificateTemplate**](CertificateTemplate.md) |  | [optional] 
**Constraints** | Pointer to [**NullableCertificateRequestConstraints**](CertificateRequestConstraints.md) |  | [optional] 
**CryptoPolicy** | [**ManagedCertificateProfileCryptoPolicy**](ManagedCertificateProfileCryptoPolicy.md) |  | 
**CsrDataMapping** | Pointer to **map[string]string** |  | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**DsFlow** | Pointer to [**[]DataSourceFlowEntry**](DataSourceFlowEntry.md) | Representation of a datasource execution flow | [optional] 
**Enabled** | **bool** |  | 
**GradingPolicies** | Pointer to **[]string** |  | [optional] 
**MaxCertificatePerHolderPolicy** | Pointer to [**NullableMaxCertificatePerHolderPolicy**](MaxCertificatePerHolderPolicy.md) |  | [optional] 
**Module** | **string** |  | 
**Name** | **string** |  | 
**PkiConnector** | **string** |  | 
**RequestsPolicy** | [**RequestsPolicy**](RequestsPolicy.md) |  | 
**SelfPermissions** | [**CertificateProfileSelfPermissions**](CertificateProfileSelfPermissions.md) |  | 
**ThirdPartyConnector** | **string** |  | 
**ThirdPartyDiscoverySync** | Pointer to **NullableBool** |  | [optional] [default to false]
**Triggers** | Pointer to [**NullableCertificateProfileTriggers**](CertificateProfileTriggers.md) |  | [optional] 

## Methods

### NewIntunePKCSProfile

`func NewIntunePKCSProfile(authorizationLevels CertificateProfileAuthorizationLevels, cryptoPolicy ManagedCertificateProfileCryptoPolicy, enabled bool, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, selfPermissions CertificateProfileSelfPermissions, thirdPartyConnector string, ) *IntunePKCSProfile`

NewIntunePKCSProfile instantiates a new IntunePKCSProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntunePKCSProfileWithDefaults

`func NewIntunePKCSProfileWithDefaults() *IntunePKCSProfile`

NewIntunePKCSProfileWithDefaults instantiates a new IntunePKCSProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationLevels

`func (o *IntunePKCSProfile) GetAuthorizationLevels() CertificateProfileAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *IntunePKCSProfile) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *IntunePKCSProfile) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetCertificateTemplate

`func (o *IntunePKCSProfile) GetCertificateTemplate() CertificateTemplate`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *IntunePKCSProfile) GetCertificateTemplateOk() (*CertificateTemplate, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *IntunePKCSProfile) SetCertificateTemplate(v CertificateTemplate)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *IntunePKCSProfile) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *IntunePKCSProfile) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *IntunePKCSProfile) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetConstraints

`func (o *IntunePKCSProfile) GetConstraints() CertificateRequestConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *IntunePKCSProfile) GetConstraintsOk() (*CertificateRequestConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *IntunePKCSProfile) SetConstraints(v CertificateRequestConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *IntunePKCSProfile) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *IntunePKCSProfile) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *IntunePKCSProfile) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetCryptoPolicy

`func (o *IntunePKCSProfile) GetCryptoPolicy() ManagedCertificateProfileCryptoPolicy`

GetCryptoPolicy returns the CryptoPolicy field if non-nil, zero value otherwise.

### GetCryptoPolicyOk

`func (o *IntunePKCSProfile) GetCryptoPolicyOk() (*ManagedCertificateProfileCryptoPolicy, bool)`

GetCryptoPolicyOk returns a tuple with the CryptoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPolicy

`func (o *IntunePKCSProfile) SetCryptoPolicy(v ManagedCertificateProfileCryptoPolicy)`

SetCryptoPolicy sets CryptoPolicy field to given value.


### GetCsrDataMapping

`func (o *IntunePKCSProfile) GetCsrDataMapping() map[string]string`

GetCsrDataMapping returns the CsrDataMapping field if non-nil, zero value otherwise.

### GetCsrDataMappingOk

`func (o *IntunePKCSProfile) GetCsrDataMappingOk() (*map[string]string, bool)`

GetCsrDataMappingOk returns a tuple with the CsrDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataMapping

`func (o *IntunePKCSProfile) SetCsrDataMapping(v map[string]string)`

SetCsrDataMapping sets CsrDataMapping field to given value.

### HasCsrDataMapping

`func (o *IntunePKCSProfile) HasCsrDataMapping() bool`

HasCsrDataMapping returns a boolean if a field has been set.

### SetCsrDataMappingNil

`func (o *IntunePKCSProfile) SetCsrDataMappingNil(b bool)`

 SetCsrDataMappingNil sets the value for CsrDataMapping to be an explicit nil

### UnsetCsrDataMapping
`func (o *IntunePKCSProfile) UnsetCsrDataMapping()`

UnsetCsrDataMapping ensures that no value is present for CsrDataMapping, not even an explicit nil
### GetDescription

`func (o *IntunePKCSProfile) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntunePKCSProfile) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntunePKCSProfile) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntunePKCSProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IntunePKCSProfile) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IntunePKCSProfile) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *IntunePKCSProfile) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IntunePKCSProfile) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IntunePKCSProfile) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IntunePKCSProfile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IntunePKCSProfile) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IntunePKCSProfile) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDsFlow

`func (o *IntunePKCSProfile) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *IntunePKCSProfile) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *IntunePKCSProfile) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *IntunePKCSProfile) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### SetDsFlowNil

`func (o *IntunePKCSProfile) SetDsFlowNil(b bool)`

 SetDsFlowNil sets the value for DsFlow to be an explicit nil

### UnsetDsFlow
`func (o *IntunePKCSProfile) UnsetDsFlow()`

UnsetDsFlow ensures that no value is present for DsFlow, not even an explicit nil
### GetEnabled

`func (o *IntunePKCSProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntunePKCSProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntunePKCSProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGradingPolicies

`func (o *IntunePKCSProfile) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *IntunePKCSProfile) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *IntunePKCSProfile) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *IntunePKCSProfile) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *IntunePKCSProfile) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *IntunePKCSProfile) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetMaxCertificatePerHolderPolicy

`func (o *IntunePKCSProfile) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy`

GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field if non-nil, zero value otherwise.

### GetMaxCertificatePerHolderPolicyOk

`func (o *IntunePKCSProfile) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool)`

GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCertificatePerHolderPolicy

`func (o *IntunePKCSProfile) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy)`

SetMaxCertificatePerHolderPolicy sets MaxCertificatePerHolderPolicy field to given value.

### HasMaxCertificatePerHolderPolicy

`func (o *IntunePKCSProfile) HasMaxCertificatePerHolderPolicy() bool`

HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.

### SetMaxCertificatePerHolderPolicyNil

`func (o *IntunePKCSProfile) SetMaxCertificatePerHolderPolicyNil(b bool)`

 SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil

### UnsetMaxCertificatePerHolderPolicy
`func (o *IntunePKCSProfile) UnsetMaxCertificatePerHolderPolicy()`

UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
### GetModule

`func (o *IntunePKCSProfile) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *IntunePKCSProfile) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *IntunePKCSProfile) SetModule(v string)`

SetModule sets Module field to given value.


### GetName

`func (o *IntunePKCSProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntunePKCSProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntunePKCSProfile) SetName(v string)`

SetName sets Name field to given value.


### GetPkiConnector

`func (o *IntunePKCSProfile) GetPkiConnector() string`

GetPkiConnector returns the PkiConnector field if non-nil, zero value otherwise.

### GetPkiConnectorOk

`func (o *IntunePKCSProfile) GetPkiConnectorOk() (*string, bool)`

GetPkiConnectorOk returns a tuple with the PkiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnector

`func (o *IntunePKCSProfile) SetPkiConnector(v string)`

SetPkiConnector sets PkiConnector field to given value.


### GetRequestsPolicy

`func (o *IntunePKCSProfile) GetRequestsPolicy() RequestsPolicy`

GetRequestsPolicy returns the RequestsPolicy field if non-nil, zero value otherwise.

### GetRequestsPolicyOk

`func (o *IntunePKCSProfile) GetRequestsPolicyOk() (*RequestsPolicy, bool)`

GetRequestsPolicyOk returns a tuple with the RequestsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPolicy

`func (o *IntunePKCSProfile) SetRequestsPolicy(v RequestsPolicy)`

SetRequestsPolicy sets RequestsPolicy field to given value.


### GetSelfPermissions

`func (o *IntunePKCSProfile) GetSelfPermissions() CertificateProfileSelfPermissions`

GetSelfPermissions returns the SelfPermissions field if non-nil, zero value otherwise.

### GetSelfPermissionsOk

`func (o *IntunePKCSProfile) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool)`

GetSelfPermissionsOk returns a tuple with the SelfPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPermissions

`func (o *IntunePKCSProfile) SetSelfPermissions(v CertificateProfileSelfPermissions)`

SetSelfPermissions sets SelfPermissions field to given value.


### GetThirdPartyConnector

`func (o *IntunePKCSProfile) GetThirdPartyConnector() string`

GetThirdPartyConnector returns the ThirdPartyConnector field if non-nil, zero value otherwise.

### GetThirdPartyConnectorOk

`func (o *IntunePKCSProfile) GetThirdPartyConnectorOk() (*string, bool)`

GetThirdPartyConnectorOk returns a tuple with the ThirdPartyConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyConnector

`func (o *IntunePKCSProfile) SetThirdPartyConnector(v string)`

SetThirdPartyConnector sets ThirdPartyConnector field to given value.


### GetThirdPartyDiscoverySync

`func (o *IntunePKCSProfile) GetThirdPartyDiscoverySync() bool`

GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field if non-nil, zero value otherwise.

### GetThirdPartyDiscoverySyncOk

`func (o *IntunePKCSProfile) GetThirdPartyDiscoverySyncOk() (*bool, bool)`

GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyDiscoverySync

`func (o *IntunePKCSProfile) SetThirdPartyDiscoverySync(v bool)`

SetThirdPartyDiscoverySync sets ThirdPartyDiscoverySync field to given value.

### HasThirdPartyDiscoverySync

`func (o *IntunePKCSProfile) HasThirdPartyDiscoverySync() bool`

HasThirdPartyDiscoverySync returns a boolean if a field has been set.

### SetThirdPartyDiscoverySyncNil

`func (o *IntunePKCSProfile) SetThirdPartyDiscoverySyncNil(b bool)`

 SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil

### UnsetThirdPartyDiscoverySync
`func (o *IntunePKCSProfile) UnsetThirdPartyDiscoverySync()`

UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
### GetTriggers

`func (o *IntunePKCSProfile) GetTriggers() CertificateProfileTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *IntunePKCSProfile) GetTriggersOk() (*CertificateProfileTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *IntunePKCSProfile) SetTriggers(v CertificateProfileTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *IntunePKCSProfile) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *IntunePKCSProfile) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *IntunePKCSProfile) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


