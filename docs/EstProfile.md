# EstProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** |  | 
**Name** | **string** |  | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**Enabled** | **bool** |  | 
**Ca** | **string** |  | 
**Constraints** | Pointer to [**NullableCertificateRequestConstraints**](CertificateRequestConstraints.md) |  | [optional] 
**PkiConnector** | **string** |  | 
**CsrDataMapping** | Pointer to **map[string]string** |  | [optional] 
**MaxCertificatePerHolderPolicy** | Pointer to [**NullableMaxCertificatePerHolderPolicy**](MaxCertificatePerHolderPolicy.md) |  | [optional] 
**AuthorizationMode** | **string** |  | 
**DnWhitelist** | **bool** |  | 
**EnrollAuthorizedCas** | Pointer to **[]string** |  | [optional] 
**RenewalAuthorizedCas** | Pointer to **[]string** |  | [optional] 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**AuthorizationLevels** | [**CertificateProfileAuthorizationLevels**](CertificateProfileAuthorizationLevels.md) |  | 
**Triggers** | Pointer to [**NullableCertificateProfileTriggers**](CertificateProfileTriggers.md) |  | [optional] 
**RequestsPolicy** | [**RequestsPolicy**](RequestsPolicy.md) |  | 
**PasswordPolicy** | Pointer to **NullableString** |  | [optional] 
**CryptoPolicy** | [**CertificateProfileCryptoPolicy**](CertificateProfileCryptoPolicy.md) |  | 
**SelfPermissions** | [**CertificateProfileSelfPermissions**](CertificateProfileSelfPermissions.md) |  | 
**CertificateTemplate** | Pointer to [**NullableCertificateTemplate**](CertificateTemplate.md) |  | [optional] 
**GradingPolicies** | Pointer to **[]string** |  | [optional] 
**ValidationRuleset** | Pointer to [**NullableValidationRuleset**](ValidationRuleset.md) |  | [optional] 
**DsFlow** | Pointer to [**[]DataSourceFlowEntry**](DataSourceFlowEntry.md) | Representation of a datasource execution flow | [optional] 

## Methods

### NewEstProfile

`func NewEstProfile(module string, name string, enabled bool, ca string, pkiConnector string, authorizationMode string, dnWhitelist bool, authorizationLevels CertificateProfileAuthorizationLevels, requestsPolicy RequestsPolicy, cryptoPolicy CertificateProfileCryptoPolicy, selfPermissions CertificateProfileSelfPermissions, ) *EstProfile`

NewEstProfile instantiates a new EstProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstProfileWithDefaults

`func NewEstProfileWithDefaults() *EstProfile`

NewEstProfileWithDefaults instantiates a new EstProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *EstProfile) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *EstProfile) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *EstProfile) SetModule(v string)`

SetModule sets Module field to given value.


### GetName

`func (o *EstProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EstProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EstProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *EstProfile) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EstProfile) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EstProfile) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EstProfile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *EstProfile) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *EstProfile) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *EstProfile) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EstProfile) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EstProfile) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EstProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EstProfile) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EstProfile) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *EstProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EstProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EstProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCa

`func (o *EstProfile) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *EstProfile) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *EstProfile) SetCa(v string)`

SetCa sets Ca field to given value.


### GetConstraints

`func (o *EstProfile) GetConstraints() CertificateRequestConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *EstProfile) GetConstraintsOk() (*CertificateRequestConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *EstProfile) SetConstraints(v CertificateRequestConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *EstProfile) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *EstProfile) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *EstProfile) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetPkiConnector

`func (o *EstProfile) GetPkiConnector() string`

GetPkiConnector returns the PkiConnector field if non-nil, zero value otherwise.

### GetPkiConnectorOk

`func (o *EstProfile) GetPkiConnectorOk() (*string, bool)`

GetPkiConnectorOk returns a tuple with the PkiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnector

`func (o *EstProfile) SetPkiConnector(v string)`

SetPkiConnector sets PkiConnector field to given value.


### GetCsrDataMapping

`func (o *EstProfile) GetCsrDataMapping() map[string]string`

GetCsrDataMapping returns the CsrDataMapping field if non-nil, zero value otherwise.

### GetCsrDataMappingOk

`func (o *EstProfile) GetCsrDataMappingOk() (*map[string]string, bool)`

GetCsrDataMappingOk returns a tuple with the CsrDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataMapping

`func (o *EstProfile) SetCsrDataMapping(v map[string]string)`

SetCsrDataMapping sets CsrDataMapping field to given value.

### HasCsrDataMapping

`func (o *EstProfile) HasCsrDataMapping() bool`

HasCsrDataMapping returns a boolean if a field has been set.

### SetCsrDataMappingNil

`func (o *EstProfile) SetCsrDataMappingNil(b bool)`

 SetCsrDataMappingNil sets the value for CsrDataMapping to be an explicit nil

### UnsetCsrDataMapping
`func (o *EstProfile) UnsetCsrDataMapping()`

UnsetCsrDataMapping ensures that no value is present for CsrDataMapping, not even an explicit nil
### GetMaxCertificatePerHolderPolicy

`func (o *EstProfile) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy`

GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field if non-nil, zero value otherwise.

### GetMaxCertificatePerHolderPolicyOk

`func (o *EstProfile) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool)`

GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCertificatePerHolderPolicy

`func (o *EstProfile) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy)`

SetMaxCertificatePerHolderPolicy sets MaxCertificatePerHolderPolicy field to given value.

### HasMaxCertificatePerHolderPolicy

`func (o *EstProfile) HasMaxCertificatePerHolderPolicy() bool`

HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.

### SetMaxCertificatePerHolderPolicyNil

`func (o *EstProfile) SetMaxCertificatePerHolderPolicyNil(b bool)`

 SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil

### UnsetMaxCertificatePerHolderPolicy
`func (o *EstProfile) UnsetMaxCertificatePerHolderPolicy()`

UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
### GetAuthorizationMode

`func (o *EstProfile) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *EstProfile) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *EstProfile) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.


### GetDnWhitelist

`func (o *EstProfile) GetDnWhitelist() bool`

GetDnWhitelist returns the DnWhitelist field if non-nil, zero value otherwise.

### GetDnWhitelistOk

`func (o *EstProfile) GetDnWhitelistOk() (*bool, bool)`

GetDnWhitelistOk returns a tuple with the DnWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnWhitelist

`func (o *EstProfile) SetDnWhitelist(v bool)`

SetDnWhitelist sets DnWhitelist field to given value.


### GetEnrollAuthorizedCas

`func (o *EstProfile) GetEnrollAuthorizedCas() []string`

GetEnrollAuthorizedCas returns the EnrollAuthorizedCas field if non-nil, zero value otherwise.

### GetEnrollAuthorizedCasOk

`func (o *EstProfile) GetEnrollAuthorizedCasOk() (*[]string, bool)`

GetEnrollAuthorizedCasOk returns a tuple with the EnrollAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollAuthorizedCas

`func (o *EstProfile) SetEnrollAuthorizedCas(v []string)`

SetEnrollAuthorizedCas sets EnrollAuthorizedCas field to given value.

### HasEnrollAuthorizedCas

`func (o *EstProfile) HasEnrollAuthorizedCas() bool`

HasEnrollAuthorizedCas returns a boolean if a field has been set.

### SetEnrollAuthorizedCasNil

`func (o *EstProfile) SetEnrollAuthorizedCasNil(b bool)`

 SetEnrollAuthorizedCasNil sets the value for EnrollAuthorizedCas to be an explicit nil

### UnsetEnrollAuthorizedCas
`func (o *EstProfile) UnsetEnrollAuthorizedCas()`

UnsetEnrollAuthorizedCas ensures that no value is present for EnrollAuthorizedCas, not even an explicit nil
### GetRenewalAuthorizedCas

`func (o *EstProfile) GetRenewalAuthorizedCas() []string`

GetRenewalAuthorizedCas returns the RenewalAuthorizedCas field if non-nil, zero value otherwise.

### GetRenewalAuthorizedCasOk

`func (o *EstProfile) GetRenewalAuthorizedCasOk() (*[]string, bool)`

GetRenewalAuthorizedCasOk returns a tuple with the RenewalAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalAuthorizedCas

`func (o *EstProfile) SetRenewalAuthorizedCas(v []string)`

SetRenewalAuthorizedCas sets RenewalAuthorizedCas field to given value.

### HasRenewalAuthorizedCas

`func (o *EstProfile) HasRenewalAuthorizedCas() bool`

HasRenewalAuthorizedCas returns a boolean if a field has been set.

### SetRenewalAuthorizedCasNil

`func (o *EstProfile) SetRenewalAuthorizedCasNil(b bool)`

 SetRenewalAuthorizedCasNil sets the value for RenewalAuthorizedCas to be an explicit nil

### UnsetRenewalAuthorizedCas
`func (o *EstProfile) UnsetRenewalAuthorizedCas()`

UnsetRenewalAuthorizedCas ensures that no value is present for RenewalAuthorizedCas, not even an explicit nil
### GetRenewalPeriod

`func (o *EstProfile) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *EstProfile) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *EstProfile) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *EstProfile) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *EstProfile) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *EstProfile) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetAuthorizationLevels

`func (o *EstProfile) GetAuthorizationLevels() CertificateProfileAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *EstProfile) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *EstProfile) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetTriggers

`func (o *EstProfile) GetTriggers() CertificateProfileTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *EstProfile) GetTriggersOk() (*CertificateProfileTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *EstProfile) SetTriggers(v CertificateProfileTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *EstProfile) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *EstProfile) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *EstProfile) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetRequestsPolicy

`func (o *EstProfile) GetRequestsPolicy() RequestsPolicy`

GetRequestsPolicy returns the RequestsPolicy field if non-nil, zero value otherwise.

### GetRequestsPolicyOk

`func (o *EstProfile) GetRequestsPolicyOk() (*RequestsPolicy, bool)`

GetRequestsPolicyOk returns a tuple with the RequestsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPolicy

`func (o *EstProfile) SetRequestsPolicy(v RequestsPolicy)`

SetRequestsPolicy sets RequestsPolicy field to given value.


### GetPasswordPolicy

`func (o *EstProfile) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *EstProfile) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *EstProfile) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *EstProfile) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *EstProfile) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *EstProfile) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetCryptoPolicy

`func (o *EstProfile) GetCryptoPolicy() CertificateProfileCryptoPolicy`

GetCryptoPolicy returns the CryptoPolicy field if non-nil, zero value otherwise.

### GetCryptoPolicyOk

`func (o *EstProfile) GetCryptoPolicyOk() (*CertificateProfileCryptoPolicy, bool)`

GetCryptoPolicyOk returns a tuple with the CryptoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPolicy

`func (o *EstProfile) SetCryptoPolicy(v CertificateProfileCryptoPolicy)`

SetCryptoPolicy sets CryptoPolicy field to given value.


### GetSelfPermissions

`func (o *EstProfile) GetSelfPermissions() CertificateProfileSelfPermissions`

GetSelfPermissions returns the SelfPermissions field if non-nil, zero value otherwise.

### GetSelfPermissionsOk

`func (o *EstProfile) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool)`

GetSelfPermissionsOk returns a tuple with the SelfPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPermissions

`func (o *EstProfile) SetSelfPermissions(v CertificateProfileSelfPermissions)`

SetSelfPermissions sets SelfPermissions field to given value.


### GetCertificateTemplate

`func (o *EstProfile) GetCertificateTemplate() CertificateTemplate`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *EstProfile) GetCertificateTemplateOk() (*CertificateTemplate, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *EstProfile) SetCertificateTemplate(v CertificateTemplate)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *EstProfile) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *EstProfile) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *EstProfile) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetGradingPolicies

`func (o *EstProfile) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *EstProfile) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *EstProfile) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *EstProfile) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *EstProfile) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *EstProfile) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetValidationRuleset

`func (o *EstProfile) GetValidationRuleset() ValidationRuleset`

GetValidationRuleset returns the ValidationRuleset field if non-nil, zero value otherwise.

### GetValidationRulesetOk

`func (o *EstProfile) GetValidationRulesetOk() (*ValidationRuleset, bool)`

GetValidationRulesetOk returns a tuple with the ValidationRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRuleset

`func (o *EstProfile) SetValidationRuleset(v ValidationRuleset)`

SetValidationRuleset sets ValidationRuleset field to given value.

### HasValidationRuleset

`func (o *EstProfile) HasValidationRuleset() bool`

HasValidationRuleset returns a boolean if a field has been set.

### SetValidationRulesetNil

`func (o *EstProfile) SetValidationRulesetNil(b bool)`

 SetValidationRulesetNil sets the value for ValidationRuleset to be an explicit nil

### UnsetValidationRuleset
`func (o *EstProfile) UnsetValidationRuleset()`

UnsetValidationRuleset ensures that no value is present for ValidationRuleset, not even an explicit nil
### GetDsFlow

`func (o *EstProfile) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *EstProfile) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *EstProfile) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *EstProfile) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### SetDsFlowNil

`func (o *EstProfile) SetDsFlowNil(b bool)`

 SetDsFlowNil sets the value for DsFlow to be an explicit nil

### UnsetDsFlow
`func (o *EstProfile) UnsetDsFlow()`

UnsetDsFlow ensures that no value is present for DsFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


