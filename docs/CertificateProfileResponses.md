# CertificateProfileResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AuthorizationLevels** | [**CertificateProfileAuthorizationLevels**](CertificateProfileAuthorizationLevels.md) |  | 
**AuthorizationMethods** | **[]string** |  | 
**AuthorizeEmptyContact** | **bool** |  | 
**AuthorizeShortName** | **bool** |  | 
**CertificateTemplate** | Pointer to [**NullableCertificateTemplate**](CertificateTemplate.md) |  | [optional] 
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
**ThirdPartyDiscoverySync** | Pointer to **NullableBool** | Available from &#x60;2.8.2&#x60; | [optional] [default to false]
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
**AcmeUrl** | Pointer to **string** |  | [optional] 
**AuthorizedCas** | **[]string** |  | 
**RequireEAB** | **bool** |  | 
**DataFieldIdentifier** | Pointer to **NullableString** | Only when escrow is enabled in the cryptoPolicy, possible values are: &#x60;rfc822name&#x60;, &#x60;othername_upn&#x60;, &#x60;mail&#x60;, &#x60;uid&#x60;, &#x60;cn&#x60; and &#x60;label.&lt;label_name&gt;&#x60;. If a label is used, it should be defined in the certificateTemplate  | [optional] 
**Version** | **int64** |  | 

## Methods

### NewCertificateProfileResponses

`func NewCertificateProfileResponses(id string, authorizationLevels CertificateProfileAuthorizationLevels, authorizationMethods []string, authorizeEmptyContact bool, authorizeShortName bool, cryptoPolicy MonitoredCertificateProfileCryptoPolicy, enabled bool, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, requireTermsOfService bool, selfPermissions CertificateProfileSelfPermissions, timeout string, verifyRetryCount int64, verifyRetryDelay string, authorizationMode string, ca string, dnWhitelist bool, caps []string, encryptionAlgorithm string, mode string, scepRA string, thirdPartyConnector string, authorizedCas []string, requireEAB bool, version int64, ) *CertificateProfileResponses`

NewCertificateProfileResponses instantiates a new CertificateProfileResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileResponsesWithDefaults

`func NewCertificateProfileResponsesWithDefaults() *CertificateProfileResponses`

NewCertificateProfileResponsesWithDefaults instantiates a new CertificateProfileResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateProfileResponses) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateProfileResponses) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateProfileResponses) SetId(v string)`

SetId sets Id field to given value.


### GetAuthorizationLevels

`func (o *CertificateProfileResponses) GetAuthorizationLevels() CertificateProfileAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *CertificateProfileResponses) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *CertificateProfileResponses) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetAuthorizationMethods

`func (o *CertificateProfileResponses) GetAuthorizationMethods() []string`

GetAuthorizationMethods returns the AuthorizationMethods field if non-nil, zero value otherwise.

### GetAuthorizationMethodsOk

`func (o *CertificateProfileResponses) GetAuthorizationMethodsOk() (*[]string, bool)`

GetAuthorizationMethodsOk returns a tuple with the AuthorizationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMethods

`func (o *CertificateProfileResponses) SetAuthorizationMethods(v []string)`

SetAuthorizationMethods sets AuthorizationMethods field to given value.


### SetAuthorizationMethodsNil

`func (o *CertificateProfileResponses) SetAuthorizationMethodsNil(b bool)`

 SetAuthorizationMethodsNil sets the value for AuthorizationMethods to be an explicit nil

### UnsetAuthorizationMethods
`func (o *CertificateProfileResponses) UnsetAuthorizationMethods()`

UnsetAuthorizationMethods ensures that no value is present for AuthorizationMethods, not even an explicit nil
### GetAuthorizeEmptyContact

`func (o *CertificateProfileResponses) GetAuthorizeEmptyContact() bool`

GetAuthorizeEmptyContact returns the AuthorizeEmptyContact field if non-nil, zero value otherwise.

### GetAuthorizeEmptyContactOk

`func (o *CertificateProfileResponses) GetAuthorizeEmptyContactOk() (*bool, bool)`

GetAuthorizeEmptyContactOk returns a tuple with the AuthorizeEmptyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeEmptyContact

`func (o *CertificateProfileResponses) SetAuthorizeEmptyContact(v bool)`

SetAuthorizeEmptyContact sets AuthorizeEmptyContact field to given value.


### GetAuthorizeShortName

`func (o *CertificateProfileResponses) GetAuthorizeShortName() bool`

GetAuthorizeShortName returns the AuthorizeShortName field if non-nil, zero value otherwise.

### GetAuthorizeShortNameOk

`func (o *CertificateProfileResponses) GetAuthorizeShortNameOk() (*bool, bool)`

GetAuthorizeShortNameOk returns a tuple with the AuthorizeShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeShortName

`func (o *CertificateProfileResponses) SetAuthorizeShortName(v bool)`

SetAuthorizeShortName sets AuthorizeShortName field to given value.


### GetCertificateTemplate

`func (o *CertificateProfileResponses) GetCertificateTemplate() CertificateTemplate`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *CertificateProfileResponses) GetCertificateTemplateOk() (*CertificateTemplate, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *CertificateProfileResponses) SetCertificateTemplate(v CertificateTemplate)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *CertificateProfileResponses) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *CertificateProfileResponses) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *CertificateProfileResponses) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetConstraints

`func (o *CertificateProfileResponses) GetConstraints() CertificateRequestConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *CertificateProfileResponses) GetConstraintsOk() (*CertificateRequestConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *CertificateProfileResponses) SetConstraints(v CertificateRequestConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *CertificateProfileResponses) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *CertificateProfileResponses) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *CertificateProfileResponses) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetCryptoPolicy

`func (o *CertificateProfileResponses) GetCryptoPolicy() MonitoredCertificateProfileCryptoPolicy`

GetCryptoPolicy returns the CryptoPolicy field if non-nil, zero value otherwise.

### GetCryptoPolicyOk

`func (o *CertificateProfileResponses) GetCryptoPolicyOk() (*MonitoredCertificateProfileCryptoPolicy, bool)`

GetCryptoPolicyOk returns a tuple with the CryptoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPolicy

`func (o *CertificateProfileResponses) SetCryptoPolicy(v MonitoredCertificateProfileCryptoPolicy)`

SetCryptoPolicy sets CryptoPolicy field to given value.


### GetCsrDataMapping

`func (o *CertificateProfileResponses) GetCsrDataMapping() map[string]string`

GetCsrDataMapping returns the CsrDataMapping field if non-nil, zero value otherwise.

### GetCsrDataMappingOk

`func (o *CertificateProfileResponses) GetCsrDataMappingOk() (*map[string]string, bool)`

GetCsrDataMappingOk returns a tuple with the CsrDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataMapping

`func (o *CertificateProfileResponses) SetCsrDataMapping(v map[string]string)`

SetCsrDataMapping sets CsrDataMapping field to given value.

### HasCsrDataMapping

`func (o *CertificateProfileResponses) HasCsrDataMapping() bool`

HasCsrDataMapping returns a boolean if a field has been set.

### SetCsrDataMappingNil

`func (o *CertificateProfileResponses) SetCsrDataMappingNil(b bool)`

 SetCsrDataMappingNil sets the value for CsrDataMapping to be an explicit nil

### UnsetCsrDataMapping
`func (o *CertificateProfileResponses) UnsetCsrDataMapping()`

UnsetCsrDataMapping ensures that no value is present for CsrDataMapping, not even an explicit nil
### GetDefaultContacts

`func (o *CertificateProfileResponses) GetDefaultContacts() []string`

GetDefaultContacts returns the DefaultContacts field if non-nil, zero value otherwise.

### GetDefaultContactsOk

`func (o *CertificateProfileResponses) GetDefaultContactsOk() (*[]string, bool)`

GetDefaultContactsOk returns a tuple with the DefaultContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContacts

`func (o *CertificateProfileResponses) SetDefaultContacts(v []string)`

SetDefaultContacts sets DefaultContacts field to given value.

### HasDefaultContacts

`func (o *CertificateProfileResponses) HasDefaultContacts() bool`

HasDefaultContacts returns a boolean if a field has been set.

### SetDefaultContactsNil

`func (o *CertificateProfileResponses) SetDefaultContactsNil(b bool)`

 SetDefaultContactsNil sets the value for DefaultContacts to be an explicit nil

### UnsetDefaultContacts
`func (o *CertificateProfileResponses) UnsetDefaultContacts()`

UnsetDefaultContacts ensures that no value is present for DefaultContacts, not even an explicit nil
### GetDescription

`func (o *CertificateProfileResponses) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateProfileResponses) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateProfileResponses) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateProfileResponses) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateProfileResponses) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateProfileResponses) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *CertificateProfileResponses) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateProfileResponses) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateProfileResponses) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertificateProfileResponses) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CertificateProfileResponses) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CertificateProfileResponses) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDsFlow

`func (o *CertificateProfileResponses) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *CertificateProfileResponses) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *CertificateProfileResponses) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *CertificateProfileResponses) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### SetDsFlowNil

`func (o *CertificateProfileResponses) SetDsFlowNil(b bool)`

 SetDsFlowNil sets the value for DsFlow to be an explicit nil

### UnsetDsFlow
`func (o *CertificateProfileResponses) UnsetDsFlow()`

UnsetDsFlow ensures that no value is present for DsFlow, not even an explicit nil
### GetEnabled

`func (o *CertificateProfileResponses) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CertificateProfileResponses) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CertificateProfileResponses) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGradingPolicies

`func (o *CertificateProfileResponses) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *CertificateProfileResponses) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *CertificateProfileResponses) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *CertificateProfileResponses) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *CertificateProfileResponses) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *CertificateProfileResponses) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetHttp01Port

`func (o *CertificateProfileResponses) GetHttp01Port() int64`

GetHttp01Port returns the Http01Port field if non-nil, zero value otherwise.

### GetHttp01PortOk

`func (o *CertificateProfileResponses) GetHttp01PortOk() (*int64, bool)`

GetHttp01PortOk returns a tuple with the Http01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp01Port

`func (o *CertificateProfileResponses) SetHttp01Port(v int64)`

SetHttp01Port sets Http01Port field to given value.

### HasHttp01Port

`func (o *CertificateProfileResponses) HasHttp01Port() bool`

HasHttp01Port returns a boolean if a field has been set.

### SetHttp01PortNil

`func (o *CertificateProfileResponses) SetHttp01PortNil(b bool)`

 SetHttp01PortNil sets the value for Http01Port to be an explicit nil

### UnsetHttp01Port
`func (o *CertificateProfileResponses) UnsetHttp01Port()`

UnsetHttp01Port ensures that no value is present for Http01Port, not even an explicit nil
### GetMaxCertificatePerHolderPolicy

`func (o *CertificateProfileResponses) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy`

GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field if non-nil, zero value otherwise.

### GetMaxCertificatePerHolderPolicyOk

`func (o *CertificateProfileResponses) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool)`

GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCertificatePerHolderPolicy

`func (o *CertificateProfileResponses) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy)`

SetMaxCertificatePerHolderPolicy sets MaxCertificatePerHolderPolicy field to given value.

### HasMaxCertificatePerHolderPolicy

`func (o *CertificateProfileResponses) HasMaxCertificatePerHolderPolicy() bool`

HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.

### SetMaxCertificatePerHolderPolicyNil

`func (o *CertificateProfileResponses) SetMaxCertificatePerHolderPolicyNil(b bool)`

 SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil

### UnsetMaxCertificatePerHolderPolicy
`func (o *CertificateProfileResponses) UnsetMaxCertificatePerHolderPolicy()`

UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
### GetMaxDnsName

`func (o *CertificateProfileResponses) GetMaxDnsName() int64`

GetMaxDnsName returns the MaxDnsName field if non-nil, zero value otherwise.

### GetMaxDnsNameOk

`func (o *CertificateProfileResponses) GetMaxDnsNameOk() (*int64, bool)`

GetMaxDnsNameOk returns a tuple with the MaxDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDnsName

`func (o *CertificateProfileResponses) SetMaxDnsName(v int64)`

SetMaxDnsName sets MaxDnsName field to given value.

### HasMaxDnsName

`func (o *CertificateProfileResponses) HasMaxDnsName() bool`

HasMaxDnsName returns a boolean if a field has been set.

### SetMaxDnsNameNil

`func (o *CertificateProfileResponses) SetMaxDnsNameNil(b bool)`

 SetMaxDnsNameNil sets the value for MaxDnsName to be an explicit nil

### UnsetMaxDnsName
`func (o *CertificateProfileResponses) UnsetMaxDnsName()`

UnsetMaxDnsName ensures that no value is present for MaxDnsName, not even an explicit nil
### GetMeta

`func (o *CertificateProfileResponses) GetMeta() DirectoryMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertificateProfileResponses) GetMetaOk() (*DirectoryMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertificateProfileResponses) SetMeta(v DirectoryMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CertificateProfileResponses) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CertificateProfileResponses) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CertificateProfileResponses) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetModule

`func (o *CertificateProfileResponses) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CertificateProfileResponses) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CertificateProfileResponses) SetModule(v string)`

SetModule sets Module field to given value.


### GetName

`func (o *CertificateProfileResponses) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfileResponses) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfileResponses) SetName(v string)`

SetName sets Name field to given value.


### GetPkiConnector

`func (o *CertificateProfileResponses) GetPkiConnector() string`

GetPkiConnector returns the PkiConnector field if non-nil, zero value otherwise.

### GetPkiConnectorOk

`func (o *CertificateProfileResponses) GetPkiConnectorOk() (*string, bool)`

GetPkiConnectorOk returns a tuple with the PkiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnector

`func (o *CertificateProfileResponses) SetPkiConnector(v string)`

SetPkiConnector sets PkiConnector field to given value.


### GetProxy

`func (o *CertificateProfileResponses) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CertificateProfileResponses) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CertificateProfileResponses) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CertificateProfileResponses) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CertificateProfileResponses) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CertificateProfileResponses) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRenewalPeriod

`func (o *CertificateProfileResponses) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *CertificateProfileResponses) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *CertificateProfileResponses) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *CertificateProfileResponses) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *CertificateProfileResponses) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *CertificateProfileResponses) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetRequestsPolicy

`func (o *CertificateProfileResponses) GetRequestsPolicy() RequestsPolicy`

GetRequestsPolicy returns the RequestsPolicy field if non-nil, zero value otherwise.

### GetRequestsPolicyOk

`func (o *CertificateProfileResponses) GetRequestsPolicyOk() (*RequestsPolicy, bool)`

GetRequestsPolicyOk returns a tuple with the RequestsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPolicy

`func (o *CertificateProfileResponses) SetRequestsPolicy(v RequestsPolicy)`

SetRequestsPolicy sets RequestsPolicy field to given value.


### GetRequireTermsOfService

`func (o *CertificateProfileResponses) GetRequireTermsOfService() bool`

GetRequireTermsOfService returns the RequireTermsOfService field if non-nil, zero value otherwise.

### GetRequireTermsOfServiceOk

`func (o *CertificateProfileResponses) GetRequireTermsOfServiceOk() (*bool, bool)`

GetRequireTermsOfServiceOk returns a tuple with the RequireTermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTermsOfService

`func (o *CertificateProfileResponses) SetRequireTermsOfService(v bool)`

SetRequireTermsOfService sets RequireTermsOfService field to given value.


### GetSelfPermissions

`func (o *CertificateProfileResponses) GetSelfPermissions() CertificateProfileSelfPermissions`

GetSelfPermissions returns the SelfPermissions field if non-nil, zero value otherwise.

### GetSelfPermissionsOk

`func (o *CertificateProfileResponses) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool)`

GetSelfPermissionsOk returns a tuple with the SelfPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPermissions

`func (o *CertificateProfileResponses) SetSelfPermissions(v CertificateProfileSelfPermissions)`

SetSelfPermissions sets SelfPermissions field to given value.


### GetThirdPartyDiscoverySync

`func (o *CertificateProfileResponses) GetThirdPartyDiscoverySync() bool`

GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field if non-nil, zero value otherwise.

### GetThirdPartyDiscoverySyncOk

`func (o *CertificateProfileResponses) GetThirdPartyDiscoverySyncOk() (*bool, bool)`

GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyDiscoverySync

`func (o *CertificateProfileResponses) SetThirdPartyDiscoverySync(v bool)`

SetThirdPartyDiscoverySync sets ThirdPartyDiscoverySync field to given value.

### HasThirdPartyDiscoverySync

`func (o *CertificateProfileResponses) HasThirdPartyDiscoverySync() bool`

HasThirdPartyDiscoverySync returns a boolean if a field has been set.

### SetThirdPartyDiscoverySyncNil

`func (o *CertificateProfileResponses) SetThirdPartyDiscoverySyncNil(b bool)`

 SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil

### UnsetThirdPartyDiscoverySync
`func (o *CertificateProfileResponses) UnsetThirdPartyDiscoverySync()`

UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
### GetTimeout

`func (o *CertificateProfileResponses) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CertificateProfileResponses) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CertificateProfileResponses) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTlsAlpn01Port

`func (o *CertificateProfileResponses) GetTlsAlpn01Port() int64`

GetTlsAlpn01Port returns the TlsAlpn01Port field if non-nil, zero value otherwise.

### GetTlsAlpn01PortOk

`func (o *CertificateProfileResponses) GetTlsAlpn01PortOk() (*int64, bool)`

GetTlsAlpn01PortOk returns a tuple with the TlsAlpn01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAlpn01Port

`func (o *CertificateProfileResponses) SetTlsAlpn01Port(v int64)`

SetTlsAlpn01Port sets TlsAlpn01Port field to given value.

### HasTlsAlpn01Port

`func (o *CertificateProfileResponses) HasTlsAlpn01Port() bool`

HasTlsAlpn01Port returns a boolean if a field has been set.

### SetTlsAlpn01PortNil

`func (o *CertificateProfileResponses) SetTlsAlpn01PortNil(b bool)`

 SetTlsAlpn01PortNil sets the value for TlsAlpn01Port to be an explicit nil

### UnsetTlsAlpn01Port
`func (o *CertificateProfileResponses) UnsetTlsAlpn01Port()`

UnsetTlsAlpn01Port ensures that no value is present for TlsAlpn01Port, not even an explicit nil
### GetTriggers

`func (o *CertificateProfileResponses) GetTriggers() CertificateProfileTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CertificateProfileResponses) GetTriggersOk() (*CertificateProfileTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CertificateProfileResponses) SetTriggers(v CertificateProfileTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CertificateProfileResponses) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *CertificateProfileResponses) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *CertificateProfileResponses) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetVerifyRetryCount

`func (o *CertificateProfileResponses) GetVerifyRetryCount() int64`

GetVerifyRetryCount returns the VerifyRetryCount field if non-nil, zero value otherwise.

### GetVerifyRetryCountOk

`func (o *CertificateProfileResponses) GetVerifyRetryCountOk() (*int64, bool)`

GetVerifyRetryCountOk returns a tuple with the VerifyRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryCount

`func (o *CertificateProfileResponses) SetVerifyRetryCount(v int64)`

SetVerifyRetryCount sets VerifyRetryCount field to given value.


### GetVerifyRetryDelay

`func (o *CertificateProfileResponses) GetVerifyRetryDelay() string`

GetVerifyRetryDelay returns the VerifyRetryDelay field if non-nil, zero value otherwise.

### GetVerifyRetryDelayOk

`func (o *CertificateProfileResponses) GetVerifyRetryDelayOk() (*string, bool)`

GetVerifyRetryDelayOk returns a tuple with the VerifyRetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryDelay

`func (o *CertificateProfileResponses) SetVerifyRetryDelay(v string)`

SetVerifyRetryDelay sets VerifyRetryDelay field to given value.


### GetAuthorizationMode

`func (o *CertificateProfileResponses) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *CertificateProfileResponses) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *CertificateProfileResponses) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.


### GetCa

`func (o *CertificateProfileResponses) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *CertificateProfileResponses) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *CertificateProfileResponses) SetCa(v string)`

SetCa sets Ca field to given value.


### GetDnWhitelist

`func (o *CertificateProfileResponses) GetDnWhitelist() bool`

GetDnWhitelist returns the DnWhitelist field if non-nil, zero value otherwise.

### GetDnWhitelistOk

`func (o *CertificateProfileResponses) GetDnWhitelistOk() (*bool, bool)`

GetDnWhitelistOk returns a tuple with the DnWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnWhitelist

`func (o *CertificateProfileResponses) SetDnWhitelist(v bool)`

SetDnWhitelist sets DnWhitelist field to given value.


### GetEnrollAuthorizedCas

`func (o *CertificateProfileResponses) GetEnrollAuthorizedCas() []string`

GetEnrollAuthorizedCas returns the EnrollAuthorizedCas field if non-nil, zero value otherwise.

### GetEnrollAuthorizedCasOk

`func (o *CertificateProfileResponses) GetEnrollAuthorizedCasOk() (*[]string, bool)`

GetEnrollAuthorizedCasOk returns a tuple with the EnrollAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollAuthorizedCas

`func (o *CertificateProfileResponses) SetEnrollAuthorizedCas(v []string)`

SetEnrollAuthorizedCas sets EnrollAuthorizedCas field to given value.

### HasEnrollAuthorizedCas

`func (o *CertificateProfileResponses) HasEnrollAuthorizedCas() bool`

HasEnrollAuthorizedCas returns a boolean if a field has been set.

### SetEnrollAuthorizedCasNil

`func (o *CertificateProfileResponses) SetEnrollAuthorizedCasNil(b bool)`

 SetEnrollAuthorizedCasNil sets the value for EnrollAuthorizedCas to be an explicit nil

### UnsetEnrollAuthorizedCas
`func (o *CertificateProfileResponses) UnsetEnrollAuthorizedCas()`

UnsetEnrollAuthorizedCas ensures that no value is present for EnrollAuthorizedCas, not even an explicit nil
### GetPasswordPolicy

`func (o *CertificateProfileResponses) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *CertificateProfileResponses) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *CertificateProfileResponses) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *CertificateProfileResponses) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *CertificateProfileResponses) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *CertificateProfileResponses) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetRenewalAuthorizedCas

`func (o *CertificateProfileResponses) GetRenewalAuthorizedCas() []string`

GetRenewalAuthorizedCas returns the RenewalAuthorizedCas field if non-nil, zero value otherwise.

### GetRenewalAuthorizedCasOk

`func (o *CertificateProfileResponses) GetRenewalAuthorizedCasOk() (*[]string, bool)`

GetRenewalAuthorizedCasOk returns a tuple with the RenewalAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalAuthorizedCas

`func (o *CertificateProfileResponses) SetRenewalAuthorizedCas(v []string)`

SetRenewalAuthorizedCas sets RenewalAuthorizedCas field to given value.

### HasRenewalAuthorizedCas

`func (o *CertificateProfileResponses) HasRenewalAuthorizedCas() bool`

HasRenewalAuthorizedCas returns a boolean if a field has been set.

### SetRenewalAuthorizedCasNil

`func (o *CertificateProfileResponses) SetRenewalAuthorizedCasNil(b bool)`

 SetRenewalAuthorizedCasNil sets the value for RenewalAuthorizedCas to be an explicit nil

### UnsetRenewalAuthorizedCas
`func (o *CertificateProfileResponses) UnsetRenewalAuthorizedCas()`

UnsetRenewalAuthorizedCas ensures that no value is present for RenewalAuthorizedCas, not even an explicit nil
### GetValidationRuleset

`func (o *CertificateProfileResponses) GetValidationRuleset() ValidationRuleset`

GetValidationRuleset returns the ValidationRuleset field if non-nil, zero value otherwise.

### GetValidationRulesetOk

`func (o *CertificateProfileResponses) GetValidationRulesetOk() (*ValidationRuleset, bool)`

GetValidationRulesetOk returns a tuple with the ValidationRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRuleset

`func (o *CertificateProfileResponses) SetValidationRuleset(v ValidationRuleset)`

SetValidationRuleset sets ValidationRuleset field to given value.

### HasValidationRuleset

`func (o *CertificateProfileResponses) HasValidationRuleset() bool`

HasValidationRuleset returns a boolean if a field has been set.

### SetValidationRulesetNil

`func (o *CertificateProfileResponses) SetValidationRulesetNil(b bool)`

 SetValidationRulesetNil sets the value for ValidationRuleset to be an explicit nil

### UnsetValidationRuleset
`func (o *CertificateProfileResponses) UnsetValidationRuleset()`

UnsetValidationRuleset ensures that no value is present for ValidationRuleset, not even an explicit nil
### GetCaps

`func (o *CertificateProfileResponses) GetCaps() []string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *CertificateProfileResponses) GetCapsOk() (*[]string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *CertificateProfileResponses) SetCaps(v []string)`

SetCaps sets Caps field to given value.


### GetDeviceIdField

`func (o *CertificateProfileResponses) GetDeviceIdField() string`

GetDeviceIdField returns the DeviceIdField field if non-nil, zero value otherwise.

### GetDeviceIdFieldOk

`func (o *CertificateProfileResponses) GetDeviceIdFieldOk() (*string, bool)`

GetDeviceIdFieldOk returns a tuple with the DeviceIdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdField

`func (o *CertificateProfileResponses) SetDeviceIdField(v string)`

SetDeviceIdField sets DeviceIdField field to given value.

### HasDeviceIdField

`func (o *CertificateProfileResponses) HasDeviceIdField() bool`

HasDeviceIdField returns a boolean if a field has been set.

### SetDeviceIdFieldNil

`func (o *CertificateProfileResponses) SetDeviceIdFieldNil(b bool)`

 SetDeviceIdFieldNil sets the value for DeviceIdField to be an explicit nil

### UnsetDeviceIdField
`func (o *CertificateProfileResponses) UnsetDeviceIdField()`

UnsetDeviceIdField ensures that no value is present for DeviceIdField, not even an explicit nil
### GetDeviceIdSeparator

`func (o *CertificateProfileResponses) GetDeviceIdSeparator() string`

GetDeviceIdSeparator returns the DeviceIdSeparator field if non-nil, zero value otherwise.

### GetDeviceIdSeparatorOk

`func (o *CertificateProfileResponses) GetDeviceIdSeparatorOk() (*string, bool)`

GetDeviceIdSeparatorOk returns a tuple with the DeviceIdSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdSeparator

`func (o *CertificateProfileResponses) SetDeviceIdSeparator(v string)`

SetDeviceIdSeparator sets DeviceIdSeparator field to given value.

### HasDeviceIdSeparator

`func (o *CertificateProfileResponses) HasDeviceIdSeparator() bool`

HasDeviceIdSeparator returns a boolean if a field has been set.

### SetDeviceIdSeparatorNil

`func (o *CertificateProfileResponses) SetDeviceIdSeparatorNil(b bool)`

 SetDeviceIdSeparatorNil sets the value for DeviceIdSeparator to be an explicit nil

### UnsetDeviceIdSeparator
`func (o *CertificateProfileResponses) UnsetDeviceIdSeparator()`

UnsetDeviceIdSeparator ensures that no value is present for DeviceIdSeparator, not even an explicit nil
### GetEncryptionAlgorithm

`func (o *CertificateProfileResponses) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *CertificateProfileResponses) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *CertificateProfileResponses) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.


### GetMode

`func (o *CertificateProfileResponses) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CertificateProfileResponses) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CertificateProfileResponses) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPostPKIOperation

`func (o *CertificateProfileResponses) GetPostPKIOperation() bool`

GetPostPKIOperation returns the PostPKIOperation field if non-nil, zero value otherwise.

### GetPostPKIOperationOk

`func (o *CertificateProfileResponses) GetPostPKIOperationOk() (*bool, bool)`

GetPostPKIOperationOk returns a tuple with the PostPKIOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPKIOperation

`func (o *CertificateProfileResponses) SetPostPKIOperation(v bool)`

SetPostPKIOperation sets PostPKIOperation field to given value.

### HasPostPKIOperation

`func (o *CertificateProfileResponses) HasPostPKIOperation() bool`

HasPostPKIOperation returns a boolean if a field has been set.

### SetPostPKIOperationNil

`func (o *CertificateProfileResponses) SetPostPKIOperationNil(b bool)`

 SetPostPKIOperationNil sets the value for PostPKIOperation to be an explicit nil

### UnsetPostPKIOperation
`func (o *CertificateProfileResponses) UnsetPostPKIOperation()`

UnsetPostPKIOperation ensures that no value is present for PostPKIOperation, not even an explicit nil
### GetScepRA

`func (o *CertificateProfileResponses) GetScepRA() string`

GetScepRA returns the ScepRA field if non-nil, zero value otherwise.

### GetScepRAOk

`func (o *CertificateProfileResponses) GetScepRAOk() (*string, bool)`

GetScepRAOk returns a tuple with the ScepRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScepRA

`func (o *CertificateProfileResponses) SetScepRA(v string)`

SetScepRA sets ScepRA field to given value.


### GetThirdPartyConnector

`func (o *CertificateProfileResponses) GetThirdPartyConnector() string`

GetThirdPartyConnector returns the ThirdPartyConnector field if non-nil, zero value otherwise.

### GetThirdPartyConnectorOk

`func (o *CertificateProfileResponses) GetThirdPartyConnectorOk() (*string, bool)`

GetThirdPartyConnectorOk returns a tuple with the ThirdPartyConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyConnector

`func (o *CertificateProfileResponses) SetThirdPartyConnector(v string)`

SetThirdPartyConnector sets ThirdPartyConnector field to given value.


### GetExchangeCertificate

`func (o *CertificateProfileResponses) GetExchangeCertificate() string`

GetExchangeCertificate returns the ExchangeCertificate field if non-nil, zero value otherwise.

### GetExchangeCertificateOk

`func (o *CertificateProfileResponses) GetExchangeCertificateOk() (*string, bool)`

GetExchangeCertificateOk returns a tuple with the ExchangeCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCertificate

`func (o *CertificateProfileResponses) SetExchangeCertificate(v string)`

SetExchangeCertificate sets ExchangeCertificate field to given value.

### HasExchangeCertificate

`func (o *CertificateProfileResponses) HasExchangeCertificate() bool`

HasExchangeCertificate returns a boolean if a field has been set.

### SetExchangeCertificateNil

`func (o *CertificateProfileResponses) SetExchangeCertificateNil(b bool)`

 SetExchangeCertificateNil sets the value for ExchangeCertificate to be an explicit nil

### UnsetExchangeCertificate
`func (o *CertificateProfileResponses) UnsetExchangeCertificate()`

UnsetExchangeCertificate ensures that no value is present for ExchangeCertificate, not even an explicit nil
### GetAcmeUrl

`func (o *CertificateProfileResponses) GetAcmeUrl() string`

GetAcmeUrl returns the AcmeUrl field if non-nil, zero value otherwise.

### GetAcmeUrlOk

`func (o *CertificateProfileResponses) GetAcmeUrlOk() (*string, bool)`

GetAcmeUrlOk returns a tuple with the AcmeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeUrl

`func (o *CertificateProfileResponses) SetAcmeUrl(v string)`

SetAcmeUrl sets AcmeUrl field to given value.

### HasAcmeUrl

`func (o *CertificateProfileResponses) HasAcmeUrl() bool`

HasAcmeUrl returns a boolean if a field has been set.

### GetAuthorizedCas

`func (o *CertificateProfileResponses) GetAuthorizedCas() []string`

GetAuthorizedCas returns the AuthorizedCas field if non-nil, zero value otherwise.

### GetAuthorizedCasOk

`func (o *CertificateProfileResponses) GetAuthorizedCasOk() (*[]string, bool)`

GetAuthorizedCasOk returns a tuple with the AuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedCas

`func (o *CertificateProfileResponses) SetAuthorizedCas(v []string)`

SetAuthorizedCas sets AuthorizedCas field to given value.


### SetAuthorizedCasNil

`func (o *CertificateProfileResponses) SetAuthorizedCasNil(b bool)`

 SetAuthorizedCasNil sets the value for AuthorizedCas to be an explicit nil

### UnsetAuthorizedCas
`func (o *CertificateProfileResponses) UnsetAuthorizedCas()`

UnsetAuthorizedCas ensures that no value is present for AuthorizedCas, not even an explicit nil
### GetRequireEAB

`func (o *CertificateProfileResponses) GetRequireEAB() bool`

GetRequireEAB returns the RequireEAB field if non-nil, zero value otherwise.

### GetRequireEABOk

`func (o *CertificateProfileResponses) GetRequireEABOk() (*bool, bool)`

GetRequireEABOk returns a tuple with the RequireEAB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEAB

`func (o *CertificateProfileResponses) SetRequireEAB(v bool)`

SetRequireEAB sets RequireEAB field to given value.


### GetDataFieldIdentifier

`func (o *CertificateProfileResponses) GetDataFieldIdentifier() string`

GetDataFieldIdentifier returns the DataFieldIdentifier field if non-nil, zero value otherwise.

### GetDataFieldIdentifierOk

`func (o *CertificateProfileResponses) GetDataFieldIdentifierOk() (*string, bool)`

GetDataFieldIdentifierOk returns a tuple with the DataFieldIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFieldIdentifier

`func (o *CertificateProfileResponses) SetDataFieldIdentifier(v string)`

SetDataFieldIdentifier sets DataFieldIdentifier field to given value.

### HasDataFieldIdentifier

`func (o *CertificateProfileResponses) HasDataFieldIdentifier() bool`

HasDataFieldIdentifier returns a boolean if a field has been set.

### SetDataFieldIdentifierNil

`func (o *CertificateProfileResponses) SetDataFieldIdentifierNil(b bool)`

 SetDataFieldIdentifierNil sets the value for DataFieldIdentifier to be an explicit nil

### UnsetDataFieldIdentifier
`func (o *CertificateProfileResponses) UnsetDataFieldIdentifier()`

UnsetDataFieldIdentifier ensures that no value is present for DataFieldIdentifier, not even an explicit nil
### GetVersion

`func (o *CertificateProfileResponses) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CertificateProfileResponses) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CertificateProfileResponses) SetVersion(v int64)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


