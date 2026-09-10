# CertificateProfiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationLevels** | [**CertificateProfileAuthorizationLevels**](CertificateProfileAuthorizationLevels.md) |  | 
**AuthorizationMethods** | **[]string** |  | 
**AuthorizeEmptyContact** | **bool** |  | 
**AuthorizeShortName** | **bool** |  | 
**CertificateTemplate** | [**NullableCertificateTemplate**](CertificateTemplate.md) |  | 
**Constraints** | Pointer to [**NullableCertificateRequestConstraints**](CertificateRequestConstraints.md) |  | [optional] 
**CryptoPolicy** | [**MonitoredCertificateProfileCryptoPolicy**](MonitoredCertificateProfileCryptoPolicy.md) |  | 
**CsrDataMapping** | Pointer to **map[string]string** |  | [optional] 
**DefaultContacts** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**DsFlow** | Pointer to [**[]DataSourceFlowEntry**](DataSourceFlowEntry.md) | Representation of a datasource execution flow | [optional] 
**Enabled** | **bool** |  | 
**GradingPolicies** | Pointer to **[]string** |  | [optional] 
**Http01Port** | Pointer to **NullableInt64** |  | [optional] 
**MaxCertificatePerHolderPolicy** | Pointer to [**NullableMaxCertificatePerHolderPolicy**](MaxCertificatePerHolderPolicy.md) |  | [optional] 
**MaxDnsName** | Pointer to **NullableInt64** |  | [optional] 
**Meta** | Pointer to [**NullableDirectoryMeta**](DirectoryMeta.md) |  | [optional] 
**Module** | **string** |  | 
**Name** | **string** |  | 
**PkiConnector** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**RequestsPolicy** | [**RequestsPolicy**](RequestsPolicy.md) |  | 
**RequireTermsOfService** | **bool** |  | 
**SelfPermissions** | [**CertificateProfileSelfPermissions**](CertificateProfileSelfPermissions.md) |  | 
**ThirdPartyDiscoverySync** | Pointer to **NullableBool** |  | [optional] [default to false]
**Timeout** | **string** |  | 
**TlsAlpn01Port** | Pointer to **NullableInt64** |  | [optional] 
**Triggers** | Pointer to [**NullableCertificateProfileTriggers**](CertificateProfileTriggers.md) |  | [optional] 
**VerifyRetryCount** | **int64** |  | 
**VerifyRetryDelay** | **string** |  | 
**AuthorizationMode** | **string** | The authorization mode to use.  &#x60;authorized&#x60; uses permissions to allow enrollment,  &#x60;auto-validation&#x60; uses the validation ruleset, &#x60;auto-validation-authorized&#x60; uses the validation ruleset, and if enrollment is denied, uses the permissions  | 
**Ca** | **string** |  | 
**DnWhitelist** | **bool** |  | 
**EnrollAuthorizedCas** | Pointer to **[]string** |  | [optional] 
**PasswordPolicy** | Pointer to **NullableString** |  | [optional] 
**RenewalAuthorizedCas** | Pointer to **[]string** |  | [optional] 
**TermsOfService** | Pointer to **string** | Reference to a &#x60;Terms of service&#x60; object. If defined, it will be displayed on the enrollment workflow before starting certificate enrollment | [optional] 
**ValidationRuleset** | Pointer to [**NullableValidationRuleset**](ValidationRuleset.md) |  | [optional] 
**Caps** | **[]string** |  | 
**DeviceIdField** | Pointer to **NullableString** |  | [optional] 
**DeviceIdSeparator** | Pointer to **NullableString** |  | [optional] 
**EncryptionAlgorithm** | **string** |  | 
**Mode** | **string** |  | 
**PostPKIOperation** | Pointer to **NullableBool** |  | [optional] 
**ScepRA** | **string** |  | 
**ThirdPartyConnector** | **string** |  | 
**ExchangeCertificate** | Pointer to **NullableString** |  | [optional] 
**AutoRenewalPolicy** | Pointer to [**AutoRenewalPolicy**](AutoRenewalPolicy.md) |  | [optional] 
**AcmeUrl** | Pointer to **string** |  | [optional] 
**AuthorizedCas** | **[]string** |  | 
**RequireEAB** | **bool** |  | 
**DataFieldIdentifier** | Pointer to **NullableString** | Only when escrow is enabled in the cryptoPolicy, possible values are: &#x60;rfc822name&#x60;, &#x60;othername_upn&#x60;, &#x60;mail&#x60;, &#x60;uid&#x60;, &#x60;cn&#x60; and &#x60;label.&lt;label_name&gt;&#x60;. If a label is used, it should be defined in the certificateTemplate  | [optional] 

## Methods

### NewCertificateProfiles

`func NewCertificateProfiles(authorizationLevels CertificateProfileAuthorizationLevels, authorizationMethods []string, authorizeEmptyContact bool, authorizeShortName bool, certificateTemplate NullableCertificateTemplate, cryptoPolicy MonitoredCertificateProfileCryptoPolicy, enabled bool, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, requireTermsOfService bool, selfPermissions CertificateProfileSelfPermissions, timeout string, verifyRetryCount int64, verifyRetryDelay string, authorizationMode string, ca string, dnWhitelist bool, caps []string, encryptionAlgorithm string, mode string, scepRA string, thirdPartyConnector string, authorizedCas []string, requireEAB bool, ) *CertificateProfiles`

NewCertificateProfiles instantiates a new CertificateProfiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfilesWithDefaults

`func NewCertificateProfilesWithDefaults() *CertificateProfiles`

NewCertificateProfilesWithDefaults instantiates a new CertificateProfiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationLevels

`func (o *CertificateProfiles) GetAuthorizationLevels() CertificateProfileAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *CertificateProfiles) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *CertificateProfiles) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetAuthorizationMethods

`func (o *CertificateProfiles) GetAuthorizationMethods() []string`

GetAuthorizationMethods returns the AuthorizationMethods field if non-nil, zero value otherwise.

### GetAuthorizationMethodsOk

`func (o *CertificateProfiles) GetAuthorizationMethodsOk() (*[]string, bool)`

GetAuthorizationMethodsOk returns a tuple with the AuthorizationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMethods

`func (o *CertificateProfiles) SetAuthorizationMethods(v []string)`

SetAuthorizationMethods sets AuthorizationMethods field to given value.


### SetAuthorizationMethodsNil

`func (o *CertificateProfiles) SetAuthorizationMethodsNil(b bool)`

 SetAuthorizationMethodsNil sets the value for AuthorizationMethods to be an explicit nil

### UnsetAuthorizationMethods
`func (o *CertificateProfiles) UnsetAuthorizationMethods()`

UnsetAuthorizationMethods ensures that no value is present for AuthorizationMethods, not even an explicit nil
### GetAuthorizeEmptyContact

`func (o *CertificateProfiles) GetAuthorizeEmptyContact() bool`

GetAuthorizeEmptyContact returns the AuthorizeEmptyContact field if non-nil, zero value otherwise.

### GetAuthorizeEmptyContactOk

`func (o *CertificateProfiles) GetAuthorizeEmptyContactOk() (*bool, bool)`

GetAuthorizeEmptyContactOk returns a tuple with the AuthorizeEmptyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeEmptyContact

`func (o *CertificateProfiles) SetAuthorizeEmptyContact(v bool)`

SetAuthorizeEmptyContact sets AuthorizeEmptyContact field to given value.


### GetAuthorizeShortName

`func (o *CertificateProfiles) GetAuthorizeShortName() bool`

GetAuthorizeShortName returns the AuthorizeShortName field if non-nil, zero value otherwise.

### GetAuthorizeShortNameOk

`func (o *CertificateProfiles) GetAuthorizeShortNameOk() (*bool, bool)`

GetAuthorizeShortNameOk returns a tuple with the AuthorizeShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeShortName

`func (o *CertificateProfiles) SetAuthorizeShortName(v bool)`

SetAuthorizeShortName sets AuthorizeShortName field to given value.


### GetCertificateTemplate

`func (o *CertificateProfiles) GetCertificateTemplate() CertificateTemplate`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *CertificateProfiles) GetCertificateTemplateOk() (*CertificateTemplate, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *CertificateProfiles) SetCertificateTemplate(v CertificateTemplate)`

SetCertificateTemplate sets CertificateTemplate field to given value.


### SetCertificateTemplateNil

`func (o *CertificateProfiles) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *CertificateProfiles) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetConstraints

`func (o *CertificateProfiles) GetConstraints() CertificateRequestConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *CertificateProfiles) GetConstraintsOk() (*CertificateRequestConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *CertificateProfiles) SetConstraints(v CertificateRequestConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *CertificateProfiles) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *CertificateProfiles) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *CertificateProfiles) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetCryptoPolicy

`func (o *CertificateProfiles) GetCryptoPolicy() MonitoredCertificateProfileCryptoPolicy`

GetCryptoPolicy returns the CryptoPolicy field if non-nil, zero value otherwise.

### GetCryptoPolicyOk

`func (o *CertificateProfiles) GetCryptoPolicyOk() (*MonitoredCertificateProfileCryptoPolicy, bool)`

GetCryptoPolicyOk returns a tuple with the CryptoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPolicy

`func (o *CertificateProfiles) SetCryptoPolicy(v MonitoredCertificateProfileCryptoPolicy)`

SetCryptoPolicy sets CryptoPolicy field to given value.


### GetCsrDataMapping

`func (o *CertificateProfiles) GetCsrDataMapping() map[string]string`

GetCsrDataMapping returns the CsrDataMapping field if non-nil, zero value otherwise.

### GetCsrDataMappingOk

`func (o *CertificateProfiles) GetCsrDataMappingOk() (*map[string]string, bool)`

GetCsrDataMappingOk returns a tuple with the CsrDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataMapping

`func (o *CertificateProfiles) SetCsrDataMapping(v map[string]string)`

SetCsrDataMapping sets CsrDataMapping field to given value.

### HasCsrDataMapping

`func (o *CertificateProfiles) HasCsrDataMapping() bool`

HasCsrDataMapping returns a boolean if a field has been set.

### SetCsrDataMappingNil

`func (o *CertificateProfiles) SetCsrDataMappingNil(b bool)`

 SetCsrDataMappingNil sets the value for CsrDataMapping to be an explicit nil

### UnsetCsrDataMapping
`func (o *CertificateProfiles) UnsetCsrDataMapping()`

UnsetCsrDataMapping ensures that no value is present for CsrDataMapping, not even an explicit nil
### GetDefaultContacts

`func (o *CertificateProfiles) GetDefaultContacts() []string`

GetDefaultContacts returns the DefaultContacts field if non-nil, zero value otherwise.

### GetDefaultContactsOk

`func (o *CertificateProfiles) GetDefaultContactsOk() (*[]string, bool)`

GetDefaultContactsOk returns a tuple with the DefaultContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContacts

`func (o *CertificateProfiles) SetDefaultContacts(v []string)`

SetDefaultContacts sets DefaultContacts field to given value.

### HasDefaultContacts

`func (o *CertificateProfiles) HasDefaultContacts() bool`

HasDefaultContacts returns a boolean if a field has been set.

### SetDefaultContactsNil

`func (o *CertificateProfiles) SetDefaultContactsNil(b bool)`

 SetDefaultContactsNil sets the value for DefaultContacts to be an explicit nil

### UnsetDefaultContacts
`func (o *CertificateProfiles) UnsetDefaultContacts()`

UnsetDefaultContacts ensures that no value is present for DefaultContacts, not even an explicit nil
### GetDescription

`func (o *CertificateProfiles) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateProfiles) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateProfiles) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateProfiles) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateProfiles) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateProfiles) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *CertificateProfiles) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateProfiles) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateProfiles) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertificateProfiles) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CertificateProfiles) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CertificateProfiles) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDsFlow

`func (o *CertificateProfiles) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *CertificateProfiles) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *CertificateProfiles) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *CertificateProfiles) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### SetDsFlowNil

`func (o *CertificateProfiles) SetDsFlowNil(b bool)`

 SetDsFlowNil sets the value for DsFlow to be an explicit nil

### UnsetDsFlow
`func (o *CertificateProfiles) UnsetDsFlow()`

UnsetDsFlow ensures that no value is present for DsFlow, not even an explicit nil
### GetEnabled

`func (o *CertificateProfiles) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CertificateProfiles) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CertificateProfiles) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGradingPolicies

`func (o *CertificateProfiles) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *CertificateProfiles) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *CertificateProfiles) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *CertificateProfiles) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *CertificateProfiles) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *CertificateProfiles) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetHttp01Port

`func (o *CertificateProfiles) GetHttp01Port() int64`

GetHttp01Port returns the Http01Port field if non-nil, zero value otherwise.

### GetHttp01PortOk

`func (o *CertificateProfiles) GetHttp01PortOk() (*int64, bool)`

GetHttp01PortOk returns a tuple with the Http01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp01Port

`func (o *CertificateProfiles) SetHttp01Port(v int64)`

SetHttp01Port sets Http01Port field to given value.

### HasHttp01Port

`func (o *CertificateProfiles) HasHttp01Port() bool`

HasHttp01Port returns a boolean if a field has been set.

### SetHttp01PortNil

`func (o *CertificateProfiles) SetHttp01PortNil(b bool)`

 SetHttp01PortNil sets the value for Http01Port to be an explicit nil

### UnsetHttp01Port
`func (o *CertificateProfiles) UnsetHttp01Port()`

UnsetHttp01Port ensures that no value is present for Http01Port, not even an explicit nil
### GetMaxCertificatePerHolderPolicy

`func (o *CertificateProfiles) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy`

GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field if non-nil, zero value otherwise.

### GetMaxCertificatePerHolderPolicyOk

`func (o *CertificateProfiles) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool)`

GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCertificatePerHolderPolicy

`func (o *CertificateProfiles) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy)`

SetMaxCertificatePerHolderPolicy sets MaxCertificatePerHolderPolicy field to given value.

### HasMaxCertificatePerHolderPolicy

`func (o *CertificateProfiles) HasMaxCertificatePerHolderPolicy() bool`

HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.

### SetMaxCertificatePerHolderPolicyNil

`func (o *CertificateProfiles) SetMaxCertificatePerHolderPolicyNil(b bool)`

 SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil

### UnsetMaxCertificatePerHolderPolicy
`func (o *CertificateProfiles) UnsetMaxCertificatePerHolderPolicy()`

UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
### GetMaxDnsName

`func (o *CertificateProfiles) GetMaxDnsName() int64`

GetMaxDnsName returns the MaxDnsName field if non-nil, zero value otherwise.

### GetMaxDnsNameOk

`func (o *CertificateProfiles) GetMaxDnsNameOk() (*int64, bool)`

GetMaxDnsNameOk returns a tuple with the MaxDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDnsName

`func (o *CertificateProfiles) SetMaxDnsName(v int64)`

SetMaxDnsName sets MaxDnsName field to given value.

### HasMaxDnsName

`func (o *CertificateProfiles) HasMaxDnsName() bool`

HasMaxDnsName returns a boolean if a field has been set.

### SetMaxDnsNameNil

`func (o *CertificateProfiles) SetMaxDnsNameNil(b bool)`

 SetMaxDnsNameNil sets the value for MaxDnsName to be an explicit nil

### UnsetMaxDnsName
`func (o *CertificateProfiles) UnsetMaxDnsName()`

UnsetMaxDnsName ensures that no value is present for MaxDnsName, not even an explicit nil
### GetMeta

`func (o *CertificateProfiles) GetMeta() DirectoryMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertificateProfiles) GetMetaOk() (*DirectoryMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertificateProfiles) SetMeta(v DirectoryMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CertificateProfiles) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CertificateProfiles) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CertificateProfiles) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetModule

`func (o *CertificateProfiles) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CertificateProfiles) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CertificateProfiles) SetModule(v string)`

SetModule sets Module field to given value.


### GetName

`func (o *CertificateProfiles) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfiles) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfiles) SetName(v string)`

SetName sets Name field to given value.


### GetPkiConnector

`func (o *CertificateProfiles) GetPkiConnector() string`

GetPkiConnector returns the PkiConnector field if non-nil, zero value otherwise.

### GetPkiConnectorOk

`func (o *CertificateProfiles) GetPkiConnectorOk() (*string, bool)`

GetPkiConnectorOk returns a tuple with the PkiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnector

`func (o *CertificateProfiles) SetPkiConnector(v string)`

SetPkiConnector sets PkiConnector field to given value.


### GetProxy

`func (o *CertificateProfiles) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CertificateProfiles) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CertificateProfiles) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CertificateProfiles) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CertificateProfiles) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CertificateProfiles) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRenewalPeriod

`func (o *CertificateProfiles) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *CertificateProfiles) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *CertificateProfiles) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *CertificateProfiles) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *CertificateProfiles) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *CertificateProfiles) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetRequestsPolicy

`func (o *CertificateProfiles) GetRequestsPolicy() RequestsPolicy`

GetRequestsPolicy returns the RequestsPolicy field if non-nil, zero value otherwise.

### GetRequestsPolicyOk

`func (o *CertificateProfiles) GetRequestsPolicyOk() (*RequestsPolicy, bool)`

GetRequestsPolicyOk returns a tuple with the RequestsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPolicy

`func (o *CertificateProfiles) SetRequestsPolicy(v RequestsPolicy)`

SetRequestsPolicy sets RequestsPolicy field to given value.


### GetRequireTermsOfService

`func (o *CertificateProfiles) GetRequireTermsOfService() bool`

GetRequireTermsOfService returns the RequireTermsOfService field if non-nil, zero value otherwise.

### GetRequireTermsOfServiceOk

`func (o *CertificateProfiles) GetRequireTermsOfServiceOk() (*bool, bool)`

GetRequireTermsOfServiceOk returns a tuple with the RequireTermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTermsOfService

`func (o *CertificateProfiles) SetRequireTermsOfService(v bool)`

SetRequireTermsOfService sets RequireTermsOfService field to given value.


### GetSelfPermissions

`func (o *CertificateProfiles) GetSelfPermissions() CertificateProfileSelfPermissions`

GetSelfPermissions returns the SelfPermissions field if non-nil, zero value otherwise.

### GetSelfPermissionsOk

`func (o *CertificateProfiles) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool)`

GetSelfPermissionsOk returns a tuple with the SelfPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPermissions

`func (o *CertificateProfiles) SetSelfPermissions(v CertificateProfileSelfPermissions)`

SetSelfPermissions sets SelfPermissions field to given value.


### GetThirdPartyDiscoverySync

`func (o *CertificateProfiles) GetThirdPartyDiscoverySync() bool`

GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field if non-nil, zero value otherwise.

### GetThirdPartyDiscoverySyncOk

`func (o *CertificateProfiles) GetThirdPartyDiscoverySyncOk() (*bool, bool)`

GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyDiscoverySync

`func (o *CertificateProfiles) SetThirdPartyDiscoverySync(v bool)`

SetThirdPartyDiscoverySync sets ThirdPartyDiscoverySync field to given value.

### HasThirdPartyDiscoverySync

`func (o *CertificateProfiles) HasThirdPartyDiscoverySync() bool`

HasThirdPartyDiscoverySync returns a boolean if a field has been set.

### SetThirdPartyDiscoverySyncNil

`func (o *CertificateProfiles) SetThirdPartyDiscoverySyncNil(b bool)`

 SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil

### UnsetThirdPartyDiscoverySync
`func (o *CertificateProfiles) UnsetThirdPartyDiscoverySync()`

UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
### GetTimeout

`func (o *CertificateProfiles) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CertificateProfiles) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CertificateProfiles) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTlsAlpn01Port

`func (o *CertificateProfiles) GetTlsAlpn01Port() int64`

GetTlsAlpn01Port returns the TlsAlpn01Port field if non-nil, zero value otherwise.

### GetTlsAlpn01PortOk

`func (o *CertificateProfiles) GetTlsAlpn01PortOk() (*int64, bool)`

GetTlsAlpn01PortOk returns a tuple with the TlsAlpn01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAlpn01Port

`func (o *CertificateProfiles) SetTlsAlpn01Port(v int64)`

SetTlsAlpn01Port sets TlsAlpn01Port field to given value.

### HasTlsAlpn01Port

`func (o *CertificateProfiles) HasTlsAlpn01Port() bool`

HasTlsAlpn01Port returns a boolean if a field has been set.

### SetTlsAlpn01PortNil

`func (o *CertificateProfiles) SetTlsAlpn01PortNil(b bool)`

 SetTlsAlpn01PortNil sets the value for TlsAlpn01Port to be an explicit nil

### UnsetTlsAlpn01Port
`func (o *CertificateProfiles) UnsetTlsAlpn01Port()`

UnsetTlsAlpn01Port ensures that no value is present for TlsAlpn01Port, not even an explicit nil
### GetTriggers

`func (o *CertificateProfiles) GetTriggers() CertificateProfileTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CertificateProfiles) GetTriggersOk() (*CertificateProfileTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CertificateProfiles) SetTriggers(v CertificateProfileTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CertificateProfiles) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *CertificateProfiles) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *CertificateProfiles) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetVerifyRetryCount

`func (o *CertificateProfiles) GetVerifyRetryCount() int64`

GetVerifyRetryCount returns the VerifyRetryCount field if non-nil, zero value otherwise.

### GetVerifyRetryCountOk

`func (o *CertificateProfiles) GetVerifyRetryCountOk() (*int64, bool)`

GetVerifyRetryCountOk returns a tuple with the VerifyRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryCount

`func (o *CertificateProfiles) SetVerifyRetryCount(v int64)`

SetVerifyRetryCount sets VerifyRetryCount field to given value.


### GetVerifyRetryDelay

`func (o *CertificateProfiles) GetVerifyRetryDelay() string`

GetVerifyRetryDelay returns the VerifyRetryDelay field if non-nil, zero value otherwise.

### GetVerifyRetryDelayOk

`func (o *CertificateProfiles) GetVerifyRetryDelayOk() (*string, bool)`

GetVerifyRetryDelayOk returns a tuple with the VerifyRetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryDelay

`func (o *CertificateProfiles) SetVerifyRetryDelay(v string)`

SetVerifyRetryDelay sets VerifyRetryDelay field to given value.


### GetAuthorizationMode

`func (o *CertificateProfiles) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *CertificateProfiles) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *CertificateProfiles) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.


### GetCa

`func (o *CertificateProfiles) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *CertificateProfiles) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *CertificateProfiles) SetCa(v string)`

SetCa sets Ca field to given value.


### GetDnWhitelist

`func (o *CertificateProfiles) GetDnWhitelist() bool`

GetDnWhitelist returns the DnWhitelist field if non-nil, zero value otherwise.

### GetDnWhitelistOk

`func (o *CertificateProfiles) GetDnWhitelistOk() (*bool, bool)`

GetDnWhitelistOk returns a tuple with the DnWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnWhitelist

`func (o *CertificateProfiles) SetDnWhitelist(v bool)`

SetDnWhitelist sets DnWhitelist field to given value.


### GetEnrollAuthorizedCas

`func (o *CertificateProfiles) GetEnrollAuthorizedCas() []string`

GetEnrollAuthorizedCas returns the EnrollAuthorizedCas field if non-nil, zero value otherwise.

### GetEnrollAuthorizedCasOk

`func (o *CertificateProfiles) GetEnrollAuthorizedCasOk() (*[]string, bool)`

GetEnrollAuthorizedCasOk returns a tuple with the EnrollAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollAuthorizedCas

`func (o *CertificateProfiles) SetEnrollAuthorizedCas(v []string)`

SetEnrollAuthorizedCas sets EnrollAuthorizedCas field to given value.

### HasEnrollAuthorizedCas

`func (o *CertificateProfiles) HasEnrollAuthorizedCas() bool`

HasEnrollAuthorizedCas returns a boolean if a field has been set.

### SetEnrollAuthorizedCasNil

`func (o *CertificateProfiles) SetEnrollAuthorizedCasNil(b bool)`

 SetEnrollAuthorizedCasNil sets the value for EnrollAuthorizedCas to be an explicit nil

### UnsetEnrollAuthorizedCas
`func (o *CertificateProfiles) UnsetEnrollAuthorizedCas()`

UnsetEnrollAuthorizedCas ensures that no value is present for EnrollAuthorizedCas, not even an explicit nil
### GetPasswordPolicy

`func (o *CertificateProfiles) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *CertificateProfiles) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *CertificateProfiles) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *CertificateProfiles) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *CertificateProfiles) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *CertificateProfiles) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetRenewalAuthorizedCas

`func (o *CertificateProfiles) GetRenewalAuthorizedCas() []string`

GetRenewalAuthorizedCas returns the RenewalAuthorizedCas field if non-nil, zero value otherwise.

### GetRenewalAuthorizedCasOk

`func (o *CertificateProfiles) GetRenewalAuthorizedCasOk() (*[]string, bool)`

GetRenewalAuthorizedCasOk returns a tuple with the RenewalAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalAuthorizedCas

`func (o *CertificateProfiles) SetRenewalAuthorizedCas(v []string)`

SetRenewalAuthorizedCas sets RenewalAuthorizedCas field to given value.

### HasRenewalAuthorizedCas

`func (o *CertificateProfiles) HasRenewalAuthorizedCas() bool`

HasRenewalAuthorizedCas returns a boolean if a field has been set.

### SetRenewalAuthorizedCasNil

`func (o *CertificateProfiles) SetRenewalAuthorizedCasNil(b bool)`

 SetRenewalAuthorizedCasNil sets the value for RenewalAuthorizedCas to be an explicit nil

### UnsetRenewalAuthorizedCas
`func (o *CertificateProfiles) UnsetRenewalAuthorizedCas()`

UnsetRenewalAuthorizedCas ensures that no value is present for RenewalAuthorizedCas, not even an explicit nil
### GetTermsOfService

`func (o *CertificateProfiles) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *CertificateProfiles) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *CertificateProfiles) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *CertificateProfiles) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetValidationRuleset

`func (o *CertificateProfiles) GetValidationRuleset() ValidationRuleset`

GetValidationRuleset returns the ValidationRuleset field if non-nil, zero value otherwise.

### GetValidationRulesetOk

`func (o *CertificateProfiles) GetValidationRulesetOk() (*ValidationRuleset, bool)`

GetValidationRulesetOk returns a tuple with the ValidationRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRuleset

`func (o *CertificateProfiles) SetValidationRuleset(v ValidationRuleset)`

SetValidationRuleset sets ValidationRuleset field to given value.

### HasValidationRuleset

`func (o *CertificateProfiles) HasValidationRuleset() bool`

HasValidationRuleset returns a boolean if a field has been set.

### SetValidationRulesetNil

`func (o *CertificateProfiles) SetValidationRulesetNil(b bool)`

 SetValidationRulesetNil sets the value for ValidationRuleset to be an explicit nil

### UnsetValidationRuleset
`func (o *CertificateProfiles) UnsetValidationRuleset()`

UnsetValidationRuleset ensures that no value is present for ValidationRuleset, not even an explicit nil
### GetCaps

`func (o *CertificateProfiles) GetCaps() []string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *CertificateProfiles) GetCapsOk() (*[]string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *CertificateProfiles) SetCaps(v []string)`

SetCaps sets Caps field to given value.


### GetDeviceIdField

`func (o *CertificateProfiles) GetDeviceIdField() string`

GetDeviceIdField returns the DeviceIdField field if non-nil, zero value otherwise.

### GetDeviceIdFieldOk

`func (o *CertificateProfiles) GetDeviceIdFieldOk() (*string, bool)`

GetDeviceIdFieldOk returns a tuple with the DeviceIdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdField

`func (o *CertificateProfiles) SetDeviceIdField(v string)`

SetDeviceIdField sets DeviceIdField field to given value.

### HasDeviceIdField

`func (o *CertificateProfiles) HasDeviceIdField() bool`

HasDeviceIdField returns a boolean if a field has been set.

### SetDeviceIdFieldNil

`func (o *CertificateProfiles) SetDeviceIdFieldNil(b bool)`

 SetDeviceIdFieldNil sets the value for DeviceIdField to be an explicit nil

### UnsetDeviceIdField
`func (o *CertificateProfiles) UnsetDeviceIdField()`

UnsetDeviceIdField ensures that no value is present for DeviceIdField, not even an explicit nil
### GetDeviceIdSeparator

`func (o *CertificateProfiles) GetDeviceIdSeparator() string`

GetDeviceIdSeparator returns the DeviceIdSeparator field if non-nil, zero value otherwise.

### GetDeviceIdSeparatorOk

`func (o *CertificateProfiles) GetDeviceIdSeparatorOk() (*string, bool)`

GetDeviceIdSeparatorOk returns a tuple with the DeviceIdSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdSeparator

`func (o *CertificateProfiles) SetDeviceIdSeparator(v string)`

SetDeviceIdSeparator sets DeviceIdSeparator field to given value.

### HasDeviceIdSeparator

`func (o *CertificateProfiles) HasDeviceIdSeparator() bool`

HasDeviceIdSeparator returns a boolean if a field has been set.

### SetDeviceIdSeparatorNil

`func (o *CertificateProfiles) SetDeviceIdSeparatorNil(b bool)`

 SetDeviceIdSeparatorNil sets the value for DeviceIdSeparator to be an explicit nil

### UnsetDeviceIdSeparator
`func (o *CertificateProfiles) UnsetDeviceIdSeparator()`

UnsetDeviceIdSeparator ensures that no value is present for DeviceIdSeparator, not even an explicit nil
### GetEncryptionAlgorithm

`func (o *CertificateProfiles) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *CertificateProfiles) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *CertificateProfiles) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.


### GetMode

`func (o *CertificateProfiles) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CertificateProfiles) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CertificateProfiles) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPostPKIOperation

`func (o *CertificateProfiles) GetPostPKIOperation() bool`

GetPostPKIOperation returns the PostPKIOperation field if non-nil, zero value otherwise.

### GetPostPKIOperationOk

`func (o *CertificateProfiles) GetPostPKIOperationOk() (*bool, bool)`

GetPostPKIOperationOk returns a tuple with the PostPKIOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPKIOperation

`func (o *CertificateProfiles) SetPostPKIOperation(v bool)`

SetPostPKIOperation sets PostPKIOperation field to given value.

### HasPostPKIOperation

`func (o *CertificateProfiles) HasPostPKIOperation() bool`

HasPostPKIOperation returns a boolean if a field has been set.

### SetPostPKIOperationNil

`func (o *CertificateProfiles) SetPostPKIOperationNil(b bool)`

 SetPostPKIOperationNil sets the value for PostPKIOperation to be an explicit nil

### UnsetPostPKIOperation
`func (o *CertificateProfiles) UnsetPostPKIOperation()`

UnsetPostPKIOperation ensures that no value is present for PostPKIOperation, not even an explicit nil
### GetScepRA

`func (o *CertificateProfiles) GetScepRA() string`

GetScepRA returns the ScepRA field if non-nil, zero value otherwise.

### GetScepRAOk

`func (o *CertificateProfiles) GetScepRAOk() (*string, bool)`

GetScepRAOk returns a tuple with the ScepRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScepRA

`func (o *CertificateProfiles) SetScepRA(v string)`

SetScepRA sets ScepRA field to given value.


### GetThirdPartyConnector

`func (o *CertificateProfiles) GetThirdPartyConnector() string`

GetThirdPartyConnector returns the ThirdPartyConnector field if non-nil, zero value otherwise.

### GetThirdPartyConnectorOk

`func (o *CertificateProfiles) GetThirdPartyConnectorOk() (*string, bool)`

GetThirdPartyConnectorOk returns a tuple with the ThirdPartyConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyConnector

`func (o *CertificateProfiles) SetThirdPartyConnector(v string)`

SetThirdPartyConnector sets ThirdPartyConnector field to given value.


### GetExchangeCertificate

`func (o *CertificateProfiles) GetExchangeCertificate() string`

GetExchangeCertificate returns the ExchangeCertificate field if non-nil, zero value otherwise.

### GetExchangeCertificateOk

`func (o *CertificateProfiles) GetExchangeCertificateOk() (*string, bool)`

GetExchangeCertificateOk returns a tuple with the ExchangeCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCertificate

`func (o *CertificateProfiles) SetExchangeCertificate(v string)`

SetExchangeCertificate sets ExchangeCertificate field to given value.

### HasExchangeCertificate

`func (o *CertificateProfiles) HasExchangeCertificate() bool`

HasExchangeCertificate returns a boolean if a field has been set.

### SetExchangeCertificateNil

`func (o *CertificateProfiles) SetExchangeCertificateNil(b bool)`

 SetExchangeCertificateNil sets the value for ExchangeCertificate to be an explicit nil

### UnsetExchangeCertificate
`func (o *CertificateProfiles) UnsetExchangeCertificate()`

UnsetExchangeCertificate ensures that no value is present for ExchangeCertificate, not even an explicit nil
### GetAutoRenewalPolicy

`func (o *CertificateProfiles) GetAutoRenewalPolicy() AutoRenewalPolicy`

GetAutoRenewalPolicy returns the AutoRenewalPolicy field if non-nil, zero value otherwise.

### GetAutoRenewalPolicyOk

`func (o *CertificateProfiles) GetAutoRenewalPolicyOk() (*AutoRenewalPolicy, bool)`

GetAutoRenewalPolicyOk returns a tuple with the AutoRenewalPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewalPolicy

`func (o *CertificateProfiles) SetAutoRenewalPolicy(v AutoRenewalPolicy)`

SetAutoRenewalPolicy sets AutoRenewalPolicy field to given value.

### HasAutoRenewalPolicy

`func (o *CertificateProfiles) HasAutoRenewalPolicy() bool`

HasAutoRenewalPolicy returns a boolean if a field has been set.

### GetAcmeUrl

`func (o *CertificateProfiles) GetAcmeUrl() string`

GetAcmeUrl returns the AcmeUrl field if non-nil, zero value otherwise.

### GetAcmeUrlOk

`func (o *CertificateProfiles) GetAcmeUrlOk() (*string, bool)`

GetAcmeUrlOk returns a tuple with the AcmeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeUrl

`func (o *CertificateProfiles) SetAcmeUrl(v string)`

SetAcmeUrl sets AcmeUrl field to given value.

### HasAcmeUrl

`func (o *CertificateProfiles) HasAcmeUrl() bool`

HasAcmeUrl returns a boolean if a field has been set.

### GetAuthorizedCas

`func (o *CertificateProfiles) GetAuthorizedCas() []string`

GetAuthorizedCas returns the AuthorizedCas field if non-nil, zero value otherwise.

### GetAuthorizedCasOk

`func (o *CertificateProfiles) GetAuthorizedCasOk() (*[]string, bool)`

GetAuthorizedCasOk returns a tuple with the AuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedCas

`func (o *CertificateProfiles) SetAuthorizedCas(v []string)`

SetAuthorizedCas sets AuthorizedCas field to given value.


### SetAuthorizedCasNil

`func (o *CertificateProfiles) SetAuthorizedCasNil(b bool)`

 SetAuthorizedCasNil sets the value for AuthorizedCas to be an explicit nil

### UnsetAuthorizedCas
`func (o *CertificateProfiles) UnsetAuthorizedCas()`

UnsetAuthorizedCas ensures that no value is present for AuthorizedCas, not even an explicit nil
### GetRequireEAB

`func (o *CertificateProfiles) GetRequireEAB() bool`

GetRequireEAB returns the RequireEAB field if non-nil, zero value otherwise.

### GetRequireEABOk

`func (o *CertificateProfiles) GetRequireEABOk() (*bool, bool)`

GetRequireEABOk returns a tuple with the RequireEAB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEAB

`func (o *CertificateProfiles) SetRequireEAB(v bool)`

SetRequireEAB sets RequireEAB field to given value.


### GetDataFieldIdentifier

`func (o *CertificateProfiles) GetDataFieldIdentifier() string`

GetDataFieldIdentifier returns the DataFieldIdentifier field if non-nil, zero value otherwise.

### GetDataFieldIdentifierOk

`func (o *CertificateProfiles) GetDataFieldIdentifierOk() (*string, bool)`

GetDataFieldIdentifierOk returns a tuple with the DataFieldIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFieldIdentifier

`func (o *CertificateProfiles) SetDataFieldIdentifier(v string)`

SetDataFieldIdentifier sets DataFieldIdentifier field to given value.

### HasDataFieldIdentifier

`func (o *CertificateProfiles) HasDataFieldIdentifier() bool`

HasDataFieldIdentifier returns a boolean if a field has been set.

### SetDataFieldIdentifierNil

`func (o *CertificateProfiles) SetDataFieldIdentifierNil(b bool)`

 SetDataFieldIdentifierNil sets the value for DataFieldIdentifier to be an explicit nil

### UnsetDataFieldIdentifier
`func (o *CertificateProfiles) UnsetDataFieldIdentifier()`

UnsetDataFieldIdentifier ensures that no value is present for DataFieldIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


