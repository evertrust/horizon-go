# CertificateProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** |  | 
**Name** | **string** |  | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) |  | [optional] 
**Enabled** | **bool** |  | 
**Timeout** | **string** |  | 
**Meta** | Pointer to [**NullableDirectoryMeta**](DirectoryMeta.md) |  | [optional] 
**Constraints** | Pointer to [**NullableCertificateRequestConstraints**](CertificateRequestConstraints.md) |  | [optional] 
**AuthorizationMethods** | **[]string** |  | 
**PkiConnector** | **string** |  | 
**Http01Port** | Pointer to **NullableInt64** |  | [optional] 
**TlsAlpn01Port** | Pointer to **NullableInt64** |  | [optional] 
**AuthorizeShortName** | **bool** |  | 
**AuthorizeEmptyContact** | **bool** |  | 
**DefaultContacts** | Pointer to **[]string** |  | [optional] 
**VerifyRetryCount** | **int64** |  | 
**VerifyRetryDelay** | **string** |  | 
**RequireTermsOfService** | **bool** |  | 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**CsrDataMapping** | Pointer to **map[string]string** |  | [optional] 
**MaxCertificatePerHolderPolicy** | Pointer to [**NullableMaxCertificatePerHolderPolicy**](MaxCertificatePerHolderPolicy.md) |  | [optional] 
**MaxDnsName** | Pointer to **NullableInt64** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**AuthorizationLevels** | [**CertificateProfileAuthorizationLevels**](CertificateProfileAuthorizationLevels.md) |  | 
**Triggers** | Pointer to [**NullableCertificateProfileTriggers**](CertificateProfileTriggers.md) |  | [optional] 
**RequestsPolicy** | [**RequestsPolicy**](RequestsPolicy.md) |  | 
**SelfPermissions** | [**CertificateProfileSelfPermissions**](CertificateProfileSelfPermissions.md) |  | 
**CertificateTemplate** | Pointer to [**NullableCertificateTemplate**](CertificateTemplate.md) |  | [optional] 
**CryptoPolicy** | [**CertificateProfileCryptoPolicy**](CertificateProfileCryptoPolicy.md) |  | 
**GradingPolicies** | Pointer to **[]string** |  | [optional] 
**DsFlow** | Pointer to [**[]DataSourceFlowEntry**](DataSourceFlowEntry.md) | Representation of a datasource execution flow | [optional] 
**Ca** | **string** |  | 
**AuthorizationMode** | **string** | The authorization mode to use.  &#x60;authorized&#x60; uses permissions to allow enrollment,  &#x60;auto-validation&#x60; uses the validation ruleset, &#x60;auto-validation-authorized&#x60; uses the validation ruleset, and if enrollment is denied, uses the permissions  | 
**DnWhitelist** | **bool** |  | 
**EnrollAuthorizedCas** | Pointer to **[]string** |  | [optional] 
**RenewalAuthorizedCas** | Pointer to **[]string** |  | [optional] 
**PasswordPolicy** | Pointer to **NullableString** |  | [optional] 
**ValidationRuleset** | Pointer to [**NullableValidationRuleset**](ValidationRuleset.md) |  | [optional] 
**Mode** | **string** |  | 
**ThirdPartyConnector** | **string** |  | 
**ScepRA** | **string** |  | 
**Caps** | **[]string** |  | 
**PostPKIOperation** | Pointer to **NullableBool** |  | [optional] 
**EncryptionAlgorithm** | **string** |  | 
**DeviceIdField** | Pointer to **NullableString** |  | [optional] 
**DeviceIdSeparator** | Pointer to **NullableString** |  | [optional] 
**ExchangeCertificate** | Pointer to **NullableString** |  | [optional] 
**AcmeUrl** | Pointer to **string** |  | [optional] 
**RequireEAB** | **bool** |  | 
**AuthorizedCas** | **[]string** |  | 
**DataFieldIdentifier** | Pointer to **NullableString** | Only when escrow is enabled in the cryptoPolicy, possible values are: &#x60;rfc822name&#x60;, &#x60;othername_upn&#x60;, &#x60;mail&#x60;, &#x60;uid&#x60;, &#x60;cn&#x60; and &#x60;label.&lt;label_name&gt;&#x60;. If a label is used, it should be defined in the certificateTemplate  | [optional] 

## Methods

### NewCertificateProfileUpdateRequest

`func NewCertificateProfileUpdateRequest(module string, name string, enabled bool, timeout string, authorizationMethods []string, pkiConnector string, authorizeShortName bool, authorizeEmptyContact bool, verifyRetryCount int64, verifyRetryDelay string, requireTermsOfService bool, authorizationLevels CertificateProfileAuthorizationLevels, requestsPolicy RequestsPolicy, selfPermissions CertificateProfileSelfPermissions, cryptoPolicy CertificateProfileCryptoPolicy, ca string, authorizationMode string, dnWhitelist bool, mode string, thirdPartyConnector string, scepRA string, caps []string, encryptionAlgorithm string, requireEAB bool, authorizedCas []string, ) *CertificateProfileUpdateRequest`

NewCertificateProfileUpdateRequest instantiates a new CertificateProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileUpdateRequestWithDefaults

`func NewCertificateProfileUpdateRequestWithDefaults() *CertificateProfileUpdateRequest`

NewCertificateProfileUpdateRequestWithDefaults instantiates a new CertificateProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *CertificateProfileUpdateRequest) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CertificateProfileUpdateRequest) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CertificateProfileUpdateRequest) SetModule(v string)`

SetModule sets Module field to given value.


### GetName

`func (o *CertificateProfileUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfileUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfileUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CertificateProfileUpdateRequest) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateProfileUpdateRequest) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateProfileUpdateRequest) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertificateProfileUpdateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CertificateProfileUpdateRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CertificateProfileUpdateRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *CertificateProfileUpdateRequest) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateProfileUpdateRequest) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateProfileUpdateRequest) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateProfileUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateProfileUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateProfileUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *CertificateProfileUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CertificateProfileUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CertificateProfileUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTimeout

`func (o *CertificateProfileUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CertificateProfileUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CertificateProfileUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetMeta

`func (o *CertificateProfileUpdateRequest) GetMeta() DirectoryMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertificateProfileUpdateRequest) GetMetaOk() (*DirectoryMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertificateProfileUpdateRequest) SetMeta(v DirectoryMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CertificateProfileUpdateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CertificateProfileUpdateRequest) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CertificateProfileUpdateRequest) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetConstraints

`func (o *CertificateProfileUpdateRequest) GetConstraints() CertificateRequestConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *CertificateProfileUpdateRequest) GetConstraintsOk() (*CertificateRequestConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *CertificateProfileUpdateRequest) SetConstraints(v CertificateRequestConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *CertificateProfileUpdateRequest) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *CertificateProfileUpdateRequest) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *CertificateProfileUpdateRequest) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetAuthorizationMethods

`func (o *CertificateProfileUpdateRequest) GetAuthorizationMethods() []string`

GetAuthorizationMethods returns the AuthorizationMethods field if non-nil, zero value otherwise.

### GetAuthorizationMethodsOk

`func (o *CertificateProfileUpdateRequest) GetAuthorizationMethodsOk() (*[]string, bool)`

GetAuthorizationMethodsOk returns a tuple with the AuthorizationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMethods

`func (o *CertificateProfileUpdateRequest) SetAuthorizationMethods(v []string)`

SetAuthorizationMethods sets AuthorizationMethods field to given value.


### SetAuthorizationMethodsNil

`func (o *CertificateProfileUpdateRequest) SetAuthorizationMethodsNil(b bool)`

 SetAuthorizationMethodsNil sets the value for AuthorizationMethods to be an explicit nil

### UnsetAuthorizationMethods
`func (o *CertificateProfileUpdateRequest) UnsetAuthorizationMethods()`

UnsetAuthorizationMethods ensures that no value is present for AuthorizationMethods, not even an explicit nil
### GetPkiConnector

`func (o *CertificateProfileUpdateRequest) GetPkiConnector() string`

GetPkiConnector returns the PkiConnector field if non-nil, zero value otherwise.

### GetPkiConnectorOk

`func (o *CertificateProfileUpdateRequest) GetPkiConnectorOk() (*string, bool)`

GetPkiConnectorOk returns a tuple with the PkiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnector

`func (o *CertificateProfileUpdateRequest) SetPkiConnector(v string)`

SetPkiConnector sets PkiConnector field to given value.


### GetHttp01Port

`func (o *CertificateProfileUpdateRequest) GetHttp01Port() int64`

GetHttp01Port returns the Http01Port field if non-nil, zero value otherwise.

### GetHttp01PortOk

`func (o *CertificateProfileUpdateRequest) GetHttp01PortOk() (*int64, bool)`

GetHttp01PortOk returns a tuple with the Http01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp01Port

`func (o *CertificateProfileUpdateRequest) SetHttp01Port(v int64)`

SetHttp01Port sets Http01Port field to given value.

### HasHttp01Port

`func (o *CertificateProfileUpdateRequest) HasHttp01Port() bool`

HasHttp01Port returns a boolean if a field has been set.

### SetHttp01PortNil

`func (o *CertificateProfileUpdateRequest) SetHttp01PortNil(b bool)`

 SetHttp01PortNil sets the value for Http01Port to be an explicit nil

### UnsetHttp01Port
`func (o *CertificateProfileUpdateRequest) UnsetHttp01Port()`

UnsetHttp01Port ensures that no value is present for Http01Port, not even an explicit nil
### GetTlsAlpn01Port

`func (o *CertificateProfileUpdateRequest) GetTlsAlpn01Port() int64`

GetTlsAlpn01Port returns the TlsAlpn01Port field if non-nil, zero value otherwise.

### GetTlsAlpn01PortOk

`func (o *CertificateProfileUpdateRequest) GetTlsAlpn01PortOk() (*int64, bool)`

GetTlsAlpn01PortOk returns a tuple with the TlsAlpn01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAlpn01Port

`func (o *CertificateProfileUpdateRequest) SetTlsAlpn01Port(v int64)`

SetTlsAlpn01Port sets TlsAlpn01Port field to given value.

### HasTlsAlpn01Port

`func (o *CertificateProfileUpdateRequest) HasTlsAlpn01Port() bool`

HasTlsAlpn01Port returns a boolean if a field has been set.

### SetTlsAlpn01PortNil

`func (o *CertificateProfileUpdateRequest) SetTlsAlpn01PortNil(b bool)`

 SetTlsAlpn01PortNil sets the value for TlsAlpn01Port to be an explicit nil

### UnsetTlsAlpn01Port
`func (o *CertificateProfileUpdateRequest) UnsetTlsAlpn01Port()`

UnsetTlsAlpn01Port ensures that no value is present for TlsAlpn01Port, not even an explicit nil
### GetAuthorizeShortName

`func (o *CertificateProfileUpdateRequest) GetAuthorizeShortName() bool`

GetAuthorizeShortName returns the AuthorizeShortName field if non-nil, zero value otherwise.

### GetAuthorizeShortNameOk

`func (o *CertificateProfileUpdateRequest) GetAuthorizeShortNameOk() (*bool, bool)`

GetAuthorizeShortNameOk returns a tuple with the AuthorizeShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeShortName

`func (o *CertificateProfileUpdateRequest) SetAuthorizeShortName(v bool)`

SetAuthorizeShortName sets AuthorizeShortName field to given value.


### GetAuthorizeEmptyContact

`func (o *CertificateProfileUpdateRequest) GetAuthorizeEmptyContact() bool`

GetAuthorizeEmptyContact returns the AuthorizeEmptyContact field if non-nil, zero value otherwise.

### GetAuthorizeEmptyContactOk

`func (o *CertificateProfileUpdateRequest) GetAuthorizeEmptyContactOk() (*bool, bool)`

GetAuthorizeEmptyContactOk returns a tuple with the AuthorizeEmptyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeEmptyContact

`func (o *CertificateProfileUpdateRequest) SetAuthorizeEmptyContact(v bool)`

SetAuthorizeEmptyContact sets AuthorizeEmptyContact field to given value.


### GetDefaultContacts

`func (o *CertificateProfileUpdateRequest) GetDefaultContacts() []string`

GetDefaultContacts returns the DefaultContacts field if non-nil, zero value otherwise.

### GetDefaultContactsOk

`func (o *CertificateProfileUpdateRequest) GetDefaultContactsOk() (*[]string, bool)`

GetDefaultContactsOk returns a tuple with the DefaultContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContacts

`func (o *CertificateProfileUpdateRequest) SetDefaultContacts(v []string)`

SetDefaultContacts sets DefaultContacts field to given value.

### HasDefaultContacts

`func (o *CertificateProfileUpdateRequest) HasDefaultContacts() bool`

HasDefaultContacts returns a boolean if a field has been set.

### SetDefaultContactsNil

`func (o *CertificateProfileUpdateRequest) SetDefaultContactsNil(b bool)`

 SetDefaultContactsNil sets the value for DefaultContacts to be an explicit nil

### UnsetDefaultContacts
`func (o *CertificateProfileUpdateRequest) UnsetDefaultContacts()`

UnsetDefaultContacts ensures that no value is present for DefaultContacts, not even an explicit nil
### GetVerifyRetryCount

`func (o *CertificateProfileUpdateRequest) GetVerifyRetryCount() int64`

GetVerifyRetryCount returns the VerifyRetryCount field if non-nil, zero value otherwise.

### GetVerifyRetryCountOk

`func (o *CertificateProfileUpdateRequest) GetVerifyRetryCountOk() (*int64, bool)`

GetVerifyRetryCountOk returns a tuple with the VerifyRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryCount

`func (o *CertificateProfileUpdateRequest) SetVerifyRetryCount(v int64)`

SetVerifyRetryCount sets VerifyRetryCount field to given value.


### GetVerifyRetryDelay

`func (o *CertificateProfileUpdateRequest) GetVerifyRetryDelay() string`

GetVerifyRetryDelay returns the VerifyRetryDelay field if non-nil, zero value otherwise.

### GetVerifyRetryDelayOk

`func (o *CertificateProfileUpdateRequest) GetVerifyRetryDelayOk() (*string, bool)`

GetVerifyRetryDelayOk returns a tuple with the VerifyRetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryDelay

`func (o *CertificateProfileUpdateRequest) SetVerifyRetryDelay(v string)`

SetVerifyRetryDelay sets VerifyRetryDelay field to given value.


### GetRequireTermsOfService

`func (o *CertificateProfileUpdateRequest) GetRequireTermsOfService() bool`

GetRequireTermsOfService returns the RequireTermsOfService field if non-nil, zero value otherwise.

### GetRequireTermsOfServiceOk

`func (o *CertificateProfileUpdateRequest) GetRequireTermsOfServiceOk() (*bool, bool)`

GetRequireTermsOfServiceOk returns a tuple with the RequireTermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTermsOfService

`func (o *CertificateProfileUpdateRequest) SetRequireTermsOfService(v bool)`

SetRequireTermsOfService sets RequireTermsOfService field to given value.


### GetRenewalPeriod

`func (o *CertificateProfileUpdateRequest) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *CertificateProfileUpdateRequest) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *CertificateProfileUpdateRequest) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *CertificateProfileUpdateRequest) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *CertificateProfileUpdateRequest) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *CertificateProfileUpdateRequest) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetCsrDataMapping

`func (o *CertificateProfileUpdateRequest) GetCsrDataMapping() map[string]string`

GetCsrDataMapping returns the CsrDataMapping field if non-nil, zero value otherwise.

### GetCsrDataMappingOk

`func (o *CertificateProfileUpdateRequest) GetCsrDataMappingOk() (*map[string]string, bool)`

GetCsrDataMappingOk returns a tuple with the CsrDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataMapping

`func (o *CertificateProfileUpdateRequest) SetCsrDataMapping(v map[string]string)`

SetCsrDataMapping sets CsrDataMapping field to given value.

### HasCsrDataMapping

`func (o *CertificateProfileUpdateRequest) HasCsrDataMapping() bool`

HasCsrDataMapping returns a boolean if a field has been set.

### SetCsrDataMappingNil

`func (o *CertificateProfileUpdateRequest) SetCsrDataMappingNil(b bool)`

 SetCsrDataMappingNil sets the value for CsrDataMapping to be an explicit nil

### UnsetCsrDataMapping
`func (o *CertificateProfileUpdateRequest) UnsetCsrDataMapping()`

UnsetCsrDataMapping ensures that no value is present for CsrDataMapping, not even an explicit nil
### GetMaxCertificatePerHolderPolicy

`func (o *CertificateProfileUpdateRequest) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy`

GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field if non-nil, zero value otherwise.

### GetMaxCertificatePerHolderPolicyOk

`func (o *CertificateProfileUpdateRequest) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool)`

GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCertificatePerHolderPolicy

`func (o *CertificateProfileUpdateRequest) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy)`

SetMaxCertificatePerHolderPolicy sets MaxCertificatePerHolderPolicy field to given value.

### HasMaxCertificatePerHolderPolicy

`func (o *CertificateProfileUpdateRequest) HasMaxCertificatePerHolderPolicy() bool`

HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.

### SetMaxCertificatePerHolderPolicyNil

`func (o *CertificateProfileUpdateRequest) SetMaxCertificatePerHolderPolicyNil(b bool)`

 SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil

### UnsetMaxCertificatePerHolderPolicy
`func (o *CertificateProfileUpdateRequest) UnsetMaxCertificatePerHolderPolicy()`

UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
### GetMaxDnsName

`func (o *CertificateProfileUpdateRequest) GetMaxDnsName() int64`

GetMaxDnsName returns the MaxDnsName field if non-nil, zero value otherwise.

### GetMaxDnsNameOk

`func (o *CertificateProfileUpdateRequest) GetMaxDnsNameOk() (*int64, bool)`

GetMaxDnsNameOk returns a tuple with the MaxDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDnsName

`func (o *CertificateProfileUpdateRequest) SetMaxDnsName(v int64)`

SetMaxDnsName sets MaxDnsName field to given value.

### HasMaxDnsName

`func (o *CertificateProfileUpdateRequest) HasMaxDnsName() bool`

HasMaxDnsName returns a boolean if a field has been set.

### SetMaxDnsNameNil

`func (o *CertificateProfileUpdateRequest) SetMaxDnsNameNil(b bool)`

 SetMaxDnsNameNil sets the value for MaxDnsName to be an explicit nil

### UnsetMaxDnsName
`func (o *CertificateProfileUpdateRequest) UnsetMaxDnsName()`

UnsetMaxDnsName ensures that no value is present for MaxDnsName, not even an explicit nil
### GetProxy

`func (o *CertificateProfileUpdateRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CertificateProfileUpdateRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CertificateProfileUpdateRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CertificateProfileUpdateRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CertificateProfileUpdateRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CertificateProfileUpdateRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetAuthorizationLevels

`func (o *CertificateProfileUpdateRequest) GetAuthorizationLevels() CertificateProfileAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *CertificateProfileUpdateRequest) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *CertificateProfileUpdateRequest) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetTriggers

`func (o *CertificateProfileUpdateRequest) GetTriggers() CertificateProfileTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CertificateProfileUpdateRequest) GetTriggersOk() (*CertificateProfileTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CertificateProfileUpdateRequest) SetTriggers(v CertificateProfileTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CertificateProfileUpdateRequest) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *CertificateProfileUpdateRequest) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *CertificateProfileUpdateRequest) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetRequestsPolicy

`func (o *CertificateProfileUpdateRequest) GetRequestsPolicy() RequestsPolicy`

GetRequestsPolicy returns the RequestsPolicy field if non-nil, zero value otherwise.

### GetRequestsPolicyOk

`func (o *CertificateProfileUpdateRequest) GetRequestsPolicyOk() (*RequestsPolicy, bool)`

GetRequestsPolicyOk returns a tuple with the RequestsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPolicy

`func (o *CertificateProfileUpdateRequest) SetRequestsPolicy(v RequestsPolicy)`

SetRequestsPolicy sets RequestsPolicy field to given value.


### GetSelfPermissions

`func (o *CertificateProfileUpdateRequest) GetSelfPermissions() CertificateProfileSelfPermissions`

GetSelfPermissions returns the SelfPermissions field if non-nil, zero value otherwise.

### GetSelfPermissionsOk

`func (o *CertificateProfileUpdateRequest) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool)`

GetSelfPermissionsOk returns a tuple with the SelfPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPermissions

`func (o *CertificateProfileUpdateRequest) SetSelfPermissions(v CertificateProfileSelfPermissions)`

SetSelfPermissions sets SelfPermissions field to given value.


### GetCertificateTemplate

`func (o *CertificateProfileUpdateRequest) GetCertificateTemplate() CertificateTemplate`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *CertificateProfileUpdateRequest) GetCertificateTemplateOk() (*CertificateTemplate, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *CertificateProfileUpdateRequest) SetCertificateTemplate(v CertificateTemplate)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *CertificateProfileUpdateRequest) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *CertificateProfileUpdateRequest) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *CertificateProfileUpdateRequest) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetCryptoPolicy

`func (o *CertificateProfileUpdateRequest) GetCryptoPolicy() CertificateProfileCryptoPolicy`

GetCryptoPolicy returns the CryptoPolicy field if non-nil, zero value otherwise.

### GetCryptoPolicyOk

`func (o *CertificateProfileUpdateRequest) GetCryptoPolicyOk() (*CertificateProfileCryptoPolicy, bool)`

GetCryptoPolicyOk returns a tuple with the CryptoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPolicy

`func (o *CertificateProfileUpdateRequest) SetCryptoPolicy(v CertificateProfileCryptoPolicy)`

SetCryptoPolicy sets CryptoPolicy field to given value.


### GetGradingPolicies

`func (o *CertificateProfileUpdateRequest) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *CertificateProfileUpdateRequest) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *CertificateProfileUpdateRequest) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *CertificateProfileUpdateRequest) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *CertificateProfileUpdateRequest) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *CertificateProfileUpdateRequest) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetDsFlow

`func (o *CertificateProfileUpdateRequest) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *CertificateProfileUpdateRequest) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *CertificateProfileUpdateRequest) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *CertificateProfileUpdateRequest) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### SetDsFlowNil

`func (o *CertificateProfileUpdateRequest) SetDsFlowNil(b bool)`

 SetDsFlowNil sets the value for DsFlow to be an explicit nil

### UnsetDsFlow
`func (o *CertificateProfileUpdateRequest) UnsetDsFlow()`

UnsetDsFlow ensures that no value is present for DsFlow, not even an explicit nil
### GetCa

`func (o *CertificateProfileUpdateRequest) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *CertificateProfileUpdateRequest) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *CertificateProfileUpdateRequest) SetCa(v string)`

SetCa sets Ca field to given value.


### GetAuthorizationMode

`func (o *CertificateProfileUpdateRequest) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *CertificateProfileUpdateRequest) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *CertificateProfileUpdateRequest) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.


### GetDnWhitelist

`func (o *CertificateProfileUpdateRequest) GetDnWhitelist() bool`

GetDnWhitelist returns the DnWhitelist field if non-nil, zero value otherwise.

### GetDnWhitelistOk

`func (o *CertificateProfileUpdateRequest) GetDnWhitelistOk() (*bool, bool)`

GetDnWhitelistOk returns a tuple with the DnWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnWhitelist

`func (o *CertificateProfileUpdateRequest) SetDnWhitelist(v bool)`

SetDnWhitelist sets DnWhitelist field to given value.


### GetEnrollAuthorizedCas

`func (o *CertificateProfileUpdateRequest) GetEnrollAuthorizedCas() []string`

GetEnrollAuthorizedCas returns the EnrollAuthorizedCas field if non-nil, zero value otherwise.

### GetEnrollAuthorizedCasOk

`func (o *CertificateProfileUpdateRequest) GetEnrollAuthorizedCasOk() (*[]string, bool)`

GetEnrollAuthorizedCasOk returns a tuple with the EnrollAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollAuthorizedCas

`func (o *CertificateProfileUpdateRequest) SetEnrollAuthorizedCas(v []string)`

SetEnrollAuthorizedCas sets EnrollAuthorizedCas field to given value.

### HasEnrollAuthorizedCas

`func (o *CertificateProfileUpdateRequest) HasEnrollAuthorizedCas() bool`

HasEnrollAuthorizedCas returns a boolean if a field has been set.

### SetEnrollAuthorizedCasNil

`func (o *CertificateProfileUpdateRequest) SetEnrollAuthorizedCasNil(b bool)`

 SetEnrollAuthorizedCasNil sets the value for EnrollAuthorizedCas to be an explicit nil

### UnsetEnrollAuthorizedCas
`func (o *CertificateProfileUpdateRequest) UnsetEnrollAuthorizedCas()`

UnsetEnrollAuthorizedCas ensures that no value is present for EnrollAuthorizedCas, not even an explicit nil
### GetRenewalAuthorizedCas

`func (o *CertificateProfileUpdateRequest) GetRenewalAuthorizedCas() []string`

GetRenewalAuthorizedCas returns the RenewalAuthorizedCas field if non-nil, zero value otherwise.

### GetRenewalAuthorizedCasOk

`func (o *CertificateProfileUpdateRequest) GetRenewalAuthorizedCasOk() (*[]string, bool)`

GetRenewalAuthorizedCasOk returns a tuple with the RenewalAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalAuthorizedCas

`func (o *CertificateProfileUpdateRequest) SetRenewalAuthorizedCas(v []string)`

SetRenewalAuthorizedCas sets RenewalAuthorizedCas field to given value.

### HasRenewalAuthorizedCas

`func (o *CertificateProfileUpdateRequest) HasRenewalAuthorizedCas() bool`

HasRenewalAuthorizedCas returns a boolean if a field has been set.

### SetRenewalAuthorizedCasNil

`func (o *CertificateProfileUpdateRequest) SetRenewalAuthorizedCasNil(b bool)`

 SetRenewalAuthorizedCasNil sets the value for RenewalAuthorizedCas to be an explicit nil

### UnsetRenewalAuthorizedCas
`func (o *CertificateProfileUpdateRequest) UnsetRenewalAuthorizedCas()`

UnsetRenewalAuthorizedCas ensures that no value is present for RenewalAuthorizedCas, not even an explicit nil
### GetPasswordPolicy

`func (o *CertificateProfileUpdateRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *CertificateProfileUpdateRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *CertificateProfileUpdateRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *CertificateProfileUpdateRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *CertificateProfileUpdateRequest) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *CertificateProfileUpdateRequest) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetValidationRuleset

`func (o *CertificateProfileUpdateRequest) GetValidationRuleset() ValidationRuleset`

GetValidationRuleset returns the ValidationRuleset field if non-nil, zero value otherwise.

### GetValidationRulesetOk

`func (o *CertificateProfileUpdateRequest) GetValidationRulesetOk() (*ValidationRuleset, bool)`

GetValidationRulesetOk returns a tuple with the ValidationRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRuleset

`func (o *CertificateProfileUpdateRequest) SetValidationRuleset(v ValidationRuleset)`

SetValidationRuleset sets ValidationRuleset field to given value.

### HasValidationRuleset

`func (o *CertificateProfileUpdateRequest) HasValidationRuleset() bool`

HasValidationRuleset returns a boolean if a field has been set.

### SetValidationRulesetNil

`func (o *CertificateProfileUpdateRequest) SetValidationRulesetNil(b bool)`

 SetValidationRulesetNil sets the value for ValidationRuleset to be an explicit nil

### UnsetValidationRuleset
`func (o *CertificateProfileUpdateRequest) UnsetValidationRuleset()`

UnsetValidationRuleset ensures that no value is present for ValidationRuleset, not even an explicit nil
### GetMode

`func (o *CertificateProfileUpdateRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CertificateProfileUpdateRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CertificateProfileUpdateRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetThirdPartyConnector

`func (o *CertificateProfileUpdateRequest) GetThirdPartyConnector() string`

GetThirdPartyConnector returns the ThirdPartyConnector field if non-nil, zero value otherwise.

### GetThirdPartyConnectorOk

`func (o *CertificateProfileUpdateRequest) GetThirdPartyConnectorOk() (*string, bool)`

GetThirdPartyConnectorOk returns a tuple with the ThirdPartyConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyConnector

`func (o *CertificateProfileUpdateRequest) SetThirdPartyConnector(v string)`

SetThirdPartyConnector sets ThirdPartyConnector field to given value.


### GetScepRA

`func (o *CertificateProfileUpdateRequest) GetScepRA() string`

GetScepRA returns the ScepRA field if non-nil, zero value otherwise.

### GetScepRAOk

`func (o *CertificateProfileUpdateRequest) GetScepRAOk() (*string, bool)`

GetScepRAOk returns a tuple with the ScepRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScepRA

`func (o *CertificateProfileUpdateRequest) SetScepRA(v string)`

SetScepRA sets ScepRA field to given value.


### GetCaps

`func (o *CertificateProfileUpdateRequest) GetCaps() []string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *CertificateProfileUpdateRequest) GetCapsOk() (*[]string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *CertificateProfileUpdateRequest) SetCaps(v []string)`

SetCaps sets Caps field to given value.


### GetPostPKIOperation

`func (o *CertificateProfileUpdateRequest) GetPostPKIOperation() bool`

GetPostPKIOperation returns the PostPKIOperation field if non-nil, zero value otherwise.

### GetPostPKIOperationOk

`func (o *CertificateProfileUpdateRequest) GetPostPKIOperationOk() (*bool, bool)`

GetPostPKIOperationOk returns a tuple with the PostPKIOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPKIOperation

`func (o *CertificateProfileUpdateRequest) SetPostPKIOperation(v bool)`

SetPostPKIOperation sets PostPKIOperation field to given value.

### HasPostPKIOperation

`func (o *CertificateProfileUpdateRequest) HasPostPKIOperation() bool`

HasPostPKIOperation returns a boolean if a field has been set.

### SetPostPKIOperationNil

`func (o *CertificateProfileUpdateRequest) SetPostPKIOperationNil(b bool)`

 SetPostPKIOperationNil sets the value for PostPKIOperation to be an explicit nil

### UnsetPostPKIOperation
`func (o *CertificateProfileUpdateRequest) UnsetPostPKIOperation()`

UnsetPostPKIOperation ensures that no value is present for PostPKIOperation, not even an explicit nil
### GetEncryptionAlgorithm

`func (o *CertificateProfileUpdateRequest) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *CertificateProfileUpdateRequest) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *CertificateProfileUpdateRequest) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.


### GetDeviceIdField

`func (o *CertificateProfileUpdateRequest) GetDeviceIdField() string`

GetDeviceIdField returns the DeviceIdField field if non-nil, zero value otherwise.

### GetDeviceIdFieldOk

`func (o *CertificateProfileUpdateRequest) GetDeviceIdFieldOk() (*string, bool)`

GetDeviceIdFieldOk returns a tuple with the DeviceIdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdField

`func (o *CertificateProfileUpdateRequest) SetDeviceIdField(v string)`

SetDeviceIdField sets DeviceIdField field to given value.

### HasDeviceIdField

`func (o *CertificateProfileUpdateRequest) HasDeviceIdField() bool`

HasDeviceIdField returns a boolean if a field has been set.

### SetDeviceIdFieldNil

`func (o *CertificateProfileUpdateRequest) SetDeviceIdFieldNil(b bool)`

 SetDeviceIdFieldNil sets the value for DeviceIdField to be an explicit nil

### UnsetDeviceIdField
`func (o *CertificateProfileUpdateRequest) UnsetDeviceIdField()`

UnsetDeviceIdField ensures that no value is present for DeviceIdField, not even an explicit nil
### GetDeviceIdSeparator

`func (o *CertificateProfileUpdateRequest) GetDeviceIdSeparator() string`

GetDeviceIdSeparator returns the DeviceIdSeparator field if non-nil, zero value otherwise.

### GetDeviceIdSeparatorOk

`func (o *CertificateProfileUpdateRequest) GetDeviceIdSeparatorOk() (*string, bool)`

GetDeviceIdSeparatorOk returns a tuple with the DeviceIdSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdSeparator

`func (o *CertificateProfileUpdateRequest) SetDeviceIdSeparator(v string)`

SetDeviceIdSeparator sets DeviceIdSeparator field to given value.

### HasDeviceIdSeparator

`func (o *CertificateProfileUpdateRequest) HasDeviceIdSeparator() bool`

HasDeviceIdSeparator returns a boolean if a field has been set.

### SetDeviceIdSeparatorNil

`func (o *CertificateProfileUpdateRequest) SetDeviceIdSeparatorNil(b bool)`

 SetDeviceIdSeparatorNil sets the value for DeviceIdSeparator to be an explicit nil

### UnsetDeviceIdSeparator
`func (o *CertificateProfileUpdateRequest) UnsetDeviceIdSeparator()`

UnsetDeviceIdSeparator ensures that no value is present for DeviceIdSeparator, not even an explicit nil
### GetExchangeCertificate

`func (o *CertificateProfileUpdateRequest) GetExchangeCertificate() string`

GetExchangeCertificate returns the ExchangeCertificate field if non-nil, zero value otherwise.

### GetExchangeCertificateOk

`func (o *CertificateProfileUpdateRequest) GetExchangeCertificateOk() (*string, bool)`

GetExchangeCertificateOk returns a tuple with the ExchangeCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCertificate

`func (o *CertificateProfileUpdateRequest) SetExchangeCertificate(v string)`

SetExchangeCertificate sets ExchangeCertificate field to given value.

### HasExchangeCertificate

`func (o *CertificateProfileUpdateRequest) HasExchangeCertificate() bool`

HasExchangeCertificate returns a boolean if a field has been set.

### SetExchangeCertificateNil

`func (o *CertificateProfileUpdateRequest) SetExchangeCertificateNil(b bool)`

 SetExchangeCertificateNil sets the value for ExchangeCertificate to be an explicit nil

### UnsetExchangeCertificate
`func (o *CertificateProfileUpdateRequest) UnsetExchangeCertificate()`

UnsetExchangeCertificate ensures that no value is present for ExchangeCertificate, not even an explicit nil
### GetAcmeUrl

`func (o *CertificateProfileUpdateRequest) GetAcmeUrl() string`

GetAcmeUrl returns the AcmeUrl field if non-nil, zero value otherwise.

### GetAcmeUrlOk

`func (o *CertificateProfileUpdateRequest) GetAcmeUrlOk() (*string, bool)`

GetAcmeUrlOk returns a tuple with the AcmeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeUrl

`func (o *CertificateProfileUpdateRequest) SetAcmeUrl(v string)`

SetAcmeUrl sets AcmeUrl field to given value.

### HasAcmeUrl

`func (o *CertificateProfileUpdateRequest) HasAcmeUrl() bool`

HasAcmeUrl returns a boolean if a field has been set.

### GetRequireEAB

`func (o *CertificateProfileUpdateRequest) GetRequireEAB() bool`

GetRequireEAB returns the RequireEAB field if non-nil, zero value otherwise.

### GetRequireEABOk

`func (o *CertificateProfileUpdateRequest) GetRequireEABOk() (*bool, bool)`

GetRequireEABOk returns a tuple with the RequireEAB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEAB

`func (o *CertificateProfileUpdateRequest) SetRequireEAB(v bool)`

SetRequireEAB sets RequireEAB field to given value.


### GetAuthorizedCas

`func (o *CertificateProfileUpdateRequest) GetAuthorizedCas() []string`

GetAuthorizedCas returns the AuthorizedCas field if non-nil, zero value otherwise.

### GetAuthorizedCasOk

`func (o *CertificateProfileUpdateRequest) GetAuthorizedCasOk() (*[]string, bool)`

GetAuthorizedCasOk returns a tuple with the AuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedCas

`func (o *CertificateProfileUpdateRequest) SetAuthorizedCas(v []string)`

SetAuthorizedCas sets AuthorizedCas field to given value.


### SetAuthorizedCasNil

`func (o *CertificateProfileUpdateRequest) SetAuthorizedCasNil(b bool)`

 SetAuthorizedCasNil sets the value for AuthorizedCas to be an explicit nil

### UnsetAuthorizedCas
`func (o *CertificateProfileUpdateRequest) UnsetAuthorizedCas()`

UnsetAuthorizedCas ensures that no value is present for AuthorizedCas, not even an explicit nil
### GetDataFieldIdentifier

`func (o *CertificateProfileUpdateRequest) GetDataFieldIdentifier() string`

GetDataFieldIdentifier returns the DataFieldIdentifier field if non-nil, zero value otherwise.

### GetDataFieldIdentifierOk

`func (o *CertificateProfileUpdateRequest) GetDataFieldIdentifierOk() (*string, bool)`

GetDataFieldIdentifierOk returns a tuple with the DataFieldIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFieldIdentifier

`func (o *CertificateProfileUpdateRequest) SetDataFieldIdentifier(v string)`

SetDataFieldIdentifier sets DataFieldIdentifier field to given value.

### HasDataFieldIdentifier

`func (o *CertificateProfileUpdateRequest) HasDataFieldIdentifier() bool`

HasDataFieldIdentifier returns a boolean if a field has been set.

### SetDataFieldIdentifierNil

`func (o *CertificateProfileUpdateRequest) SetDataFieldIdentifierNil(b bool)`

 SetDataFieldIdentifierNil sets the value for DataFieldIdentifier to be an explicit nil

### UnsetDataFieldIdentifier
`func (o *CertificateProfileUpdateRequest) UnsetDataFieldIdentifier()`

UnsetDataFieldIdentifier ensures that no value is present for DataFieldIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


