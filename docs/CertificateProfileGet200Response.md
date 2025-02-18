# CertificateProfileGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
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

## Methods

### NewCertificateProfileGet200Response

`func NewCertificateProfileGet200Response(id string, module string, name string, enabled bool, timeout string, authorizationMethods []string, pkiConnector string, authorizeShortName bool, authorizeEmptyContact bool, verifyRetryCount int64, verifyRetryDelay string, requireTermsOfService bool, authorizationLevels CertificateProfileAuthorizationLevels, requestsPolicy RequestsPolicy, selfPermissions CertificateProfileSelfPermissions, cryptoPolicy CertificateProfileCryptoPolicy, ca string, authorizationMode string, dnWhitelist bool, mode string, thirdPartyConnector string, scepRA string, caps []string, encryptionAlgorithm string, requireEAB bool, authorizedCas []string, ) *CertificateProfileGet200Response`

NewCertificateProfileGet200Response instantiates a new CertificateProfileGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileGet200ResponseWithDefaults

`func NewCertificateProfileGet200ResponseWithDefaults() *CertificateProfileGet200Response`

NewCertificateProfileGet200ResponseWithDefaults instantiates a new CertificateProfileGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateProfileGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateProfileGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateProfileGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetModule

`func (o *CertificateProfileGet200Response) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CertificateProfileGet200Response) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CertificateProfileGet200Response) SetModule(v string)`

SetModule sets Module field to given value.


### GetName

`func (o *CertificateProfileGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfileGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfileGet200Response) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CertificateProfileGet200Response) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateProfileGet200Response) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateProfileGet200Response) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertificateProfileGet200Response) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CertificateProfileGet200Response) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CertificateProfileGet200Response) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *CertificateProfileGet200Response) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateProfileGet200Response) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateProfileGet200Response) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateProfileGet200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateProfileGet200Response) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateProfileGet200Response) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *CertificateProfileGet200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CertificateProfileGet200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CertificateProfileGet200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTimeout

`func (o *CertificateProfileGet200Response) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CertificateProfileGet200Response) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CertificateProfileGet200Response) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetMeta

`func (o *CertificateProfileGet200Response) GetMeta() DirectoryMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertificateProfileGet200Response) GetMetaOk() (*DirectoryMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertificateProfileGet200Response) SetMeta(v DirectoryMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CertificateProfileGet200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CertificateProfileGet200Response) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CertificateProfileGet200Response) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetConstraints

`func (o *CertificateProfileGet200Response) GetConstraints() CertificateRequestConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *CertificateProfileGet200Response) GetConstraintsOk() (*CertificateRequestConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *CertificateProfileGet200Response) SetConstraints(v CertificateRequestConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *CertificateProfileGet200Response) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *CertificateProfileGet200Response) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *CertificateProfileGet200Response) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetAuthorizationMethods

`func (o *CertificateProfileGet200Response) GetAuthorizationMethods() []string`

GetAuthorizationMethods returns the AuthorizationMethods field if non-nil, zero value otherwise.

### GetAuthorizationMethodsOk

`func (o *CertificateProfileGet200Response) GetAuthorizationMethodsOk() (*[]string, bool)`

GetAuthorizationMethodsOk returns a tuple with the AuthorizationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMethods

`func (o *CertificateProfileGet200Response) SetAuthorizationMethods(v []string)`

SetAuthorizationMethods sets AuthorizationMethods field to given value.


### SetAuthorizationMethodsNil

`func (o *CertificateProfileGet200Response) SetAuthorizationMethodsNil(b bool)`

 SetAuthorizationMethodsNil sets the value for AuthorizationMethods to be an explicit nil

### UnsetAuthorizationMethods
`func (o *CertificateProfileGet200Response) UnsetAuthorizationMethods()`

UnsetAuthorizationMethods ensures that no value is present for AuthorizationMethods, not even an explicit nil
### GetPkiConnector

`func (o *CertificateProfileGet200Response) GetPkiConnector() string`

GetPkiConnector returns the PkiConnector field if non-nil, zero value otherwise.

### GetPkiConnectorOk

`func (o *CertificateProfileGet200Response) GetPkiConnectorOk() (*string, bool)`

GetPkiConnectorOk returns a tuple with the PkiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnector

`func (o *CertificateProfileGet200Response) SetPkiConnector(v string)`

SetPkiConnector sets PkiConnector field to given value.


### GetHttp01Port

`func (o *CertificateProfileGet200Response) GetHttp01Port() int64`

GetHttp01Port returns the Http01Port field if non-nil, zero value otherwise.

### GetHttp01PortOk

`func (o *CertificateProfileGet200Response) GetHttp01PortOk() (*int64, bool)`

GetHttp01PortOk returns a tuple with the Http01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp01Port

`func (o *CertificateProfileGet200Response) SetHttp01Port(v int64)`

SetHttp01Port sets Http01Port field to given value.

### HasHttp01Port

`func (o *CertificateProfileGet200Response) HasHttp01Port() bool`

HasHttp01Port returns a boolean if a field has been set.

### SetHttp01PortNil

`func (o *CertificateProfileGet200Response) SetHttp01PortNil(b bool)`

 SetHttp01PortNil sets the value for Http01Port to be an explicit nil

### UnsetHttp01Port
`func (o *CertificateProfileGet200Response) UnsetHttp01Port()`

UnsetHttp01Port ensures that no value is present for Http01Port, not even an explicit nil
### GetTlsAlpn01Port

`func (o *CertificateProfileGet200Response) GetTlsAlpn01Port() int64`

GetTlsAlpn01Port returns the TlsAlpn01Port field if non-nil, zero value otherwise.

### GetTlsAlpn01PortOk

`func (o *CertificateProfileGet200Response) GetTlsAlpn01PortOk() (*int64, bool)`

GetTlsAlpn01PortOk returns a tuple with the TlsAlpn01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAlpn01Port

`func (o *CertificateProfileGet200Response) SetTlsAlpn01Port(v int64)`

SetTlsAlpn01Port sets TlsAlpn01Port field to given value.

### HasTlsAlpn01Port

`func (o *CertificateProfileGet200Response) HasTlsAlpn01Port() bool`

HasTlsAlpn01Port returns a boolean if a field has been set.

### SetTlsAlpn01PortNil

`func (o *CertificateProfileGet200Response) SetTlsAlpn01PortNil(b bool)`

 SetTlsAlpn01PortNil sets the value for TlsAlpn01Port to be an explicit nil

### UnsetTlsAlpn01Port
`func (o *CertificateProfileGet200Response) UnsetTlsAlpn01Port()`

UnsetTlsAlpn01Port ensures that no value is present for TlsAlpn01Port, not even an explicit nil
### GetAuthorizeShortName

`func (o *CertificateProfileGet200Response) GetAuthorizeShortName() bool`

GetAuthorizeShortName returns the AuthorizeShortName field if non-nil, zero value otherwise.

### GetAuthorizeShortNameOk

`func (o *CertificateProfileGet200Response) GetAuthorizeShortNameOk() (*bool, bool)`

GetAuthorizeShortNameOk returns a tuple with the AuthorizeShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeShortName

`func (o *CertificateProfileGet200Response) SetAuthorizeShortName(v bool)`

SetAuthorizeShortName sets AuthorizeShortName field to given value.


### GetAuthorizeEmptyContact

`func (o *CertificateProfileGet200Response) GetAuthorizeEmptyContact() bool`

GetAuthorizeEmptyContact returns the AuthorizeEmptyContact field if non-nil, zero value otherwise.

### GetAuthorizeEmptyContactOk

`func (o *CertificateProfileGet200Response) GetAuthorizeEmptyContactOk() (*bool, bool)`

GetAuthorizeEmptyContactOk returns a tuple with the AuthorizeEmptyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeEmptyContact

`func (o *CertificateProfileGet200Response) SetAuthorizeEmptyContact(v bool)`

SetAuthorizeEmptyContact sets AuthorizeEmptyContact field to given value.


### GetDefaultContacts

`func (o *CertificateProfileGet200Response) GetDefaultContacts() []string`

GetDefaultContacts returns the DefaultContacts field if non-nil, zero value otherwise.

### GetDefaultContactsOk

`func (o *CertificateProfileGet200Response) GetDefaultContactsOk() (*[]string, bool)`

GetDefaultContactsOk returns a tuple with the DefaultContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContacts

`func (o *CertificateProfileGet200Response) SetDefaultContacts(v []string)`

SetDefaultContacts sets DefaultContacts field to given value.

### HasDefaultContacts

`func (o *CertificateProfileGet200Response) HasDefaultContacts() bool`

HasDefaultContacts returns a boolean if a field has been set.

### SetDefaultContactsNil

`func (o *CertificateProfileGet200Response) SetDefaultContactsNil(b bool)`

 SetDefaultContactsNil sets the value for DefaultContacts to be an explicit nil

### UnsetDefaultContacts
`func (o *CertificateProfileGet200Response) UnsetDefaultContacts()`

UnsetDefaultContacts ensures that no value is present for DefaultContacts, not even an explicit nil
### GetVerifyRetryCount

`func (o *CertificateProfileGet200Response) GetVerifyRetryCount() int64`

GetVerifyRetryCount returns the VerifyRetryCount field if non-nil, zero value otherwise.

### GetVerifyRetryCountOk

`func (o *CertificateProfileGet200Response) GetVerifyRetryCountOk() (*int64, bool)`

GetVerifyRetryCountOk returns a tuple with the VerifyRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryCount

`func (o *CertificateProfileGet200Response) SetVerifyRetryCount(v int64)`

SetVerifyRetryCount sets VerifyRetryCount field to given value.


### GetVerifyRetryDelay

`func (o *CertificateProfileGet200Response) GetVerifyRetryDelay() string`

GetVerifyRetryDelay returns the VerifyRetryDelay field if non-nil, zero value otherwise.

### GetVerifyRetryDelayOk

`func (o *CertificateProfileGet200Response) GetVerifyRetryDelayOk() (*string, bool)`

GetVerifyRetryDelayOk returns a tuple with the VerifyRetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryDelay

`func (o *CertificateProfileGet200Response) SetVerifyRetryDelay(v string)`

SetVerifyRetryDelay sets VerifyRetryDelay field to given value.


### GetRequireTermsOfService

`func (o *CertificateProfileGet200Response) GetRequireTermsOfService() bool`

GetRequireTermsOfService returns the RequireTermsOfService field if non-nil, zero value otherwise.

### GetRequireTermsOfServiceOk

`func (o *CertificateProfileGet200Response) GetRequireTermsOfServiceOk() (*bool, bool)`

GetRequireTermsOfServiceOk returns a tuple with the RequireTermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTermsOfService

`func (o *CertificateProfileGet200Response) SetRequireTermsOfService(v bool)`

SetRequireTermsOfService sets RequireTermsOfService field to given value.


### GetRenewalPeriod

`func (o *CertificateProfileGet200Response) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *CertificateProfileGet200Response) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *CertificateProfileGet200Response) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *CertificateProfileGet200Response) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *CertificateProfileGet200Response) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *CertificateProfileGet200Response) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetCsrDataMapping

`func (o *CertificateProfileGet200Response) GetCsrDataMapping() map[string]string`

GetCsrDataMapping returns the CsrDataMapping field if non-nil, zero value otherwise.

### GetCsrDataMappingOk

`func (o *CertificateProfileGet200Response) GetCsrDataMappingOk() (*map[string]string, bool)`

GetCsrDataMappingOk returns a tuple with the CsrDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataMapping

`func (o *CertificateProfileGet200Response) SetCsrDataMapping(v map[string]string)`

SetCsrDataMapping sets CsrDataMapping field to given value.

### HasCsrDataMapping

`func (o *CertificateProfileGet200Response) HasCsrDataMapping() bool`

HasCsrDataMapping returns a boolean if a field has been set.

### SetCsrDataMappingNil

`func (o *CertificateProfileGet200Response) SetCsrDataMappingNil(b bool)`

 SetCsrDataMappingNil sets the value for CsrDataMapping to be an explicit nil

### UnsetCsrDataMapping
`func (o *CertificateProfileGet200Response) UnsetCsrDataMapping()`

UnsetCsrDataMapping ensures that no value is present for CsrDataMapping, not even an explicit nil
### GetMaxCertificatePerHolderPolicy

`func (o *CertificateProfileGet200Response) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy`

GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field if non-nil, zero value otherwise.

### GetMaxCertificatePerHolderPolicyOk

`func (o *CertificateProfileGet200Response) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool)`

GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCertificatePerHolderPolicy

`func (o *CertificateProfileGet200Response) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy)`

SetMaxCertificatePerHolderPolicy sets MaxCertificatePerHolderPolicy field to given value.

### HasMaxCertificatePerHolderPolicy

`func (o *CertificateProfileGet200Response) HasMaxCertificatePerHolderPolicy() bool`

HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.

### SetMaxCertificatePerHolderPolicyNil

`func (o *CertificateProfileGet200Response) SetMaxCertificatePerHolderPolicyNil(b bool)`

 SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil

### UnsetMaxCertificatePerHolderPolicy
`func (o *CertificateProfileGet200Response) UnsetMaxCertificatePerHolderPolicy()`

UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
### GetMaxDnsName

`func (o *CertificateProfileGet200Response) GetMaxDnsName() int64`

GetMaxDnsName returns the MaxDnsName field if non-nil, zero value otherwise.

### GetMaxDnsNameOk

`func (o *CertificateProfileGet200Response) GetMaxDnsNameOk() (*int64, bool)`

GetMaxDnsNameOk returns a tuple with the MaxDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDnsName

`func (o *CertificateProfileGet200Response) SetMaxDnsName(v int64)`

SetMaxDnsName sets MaxDnsName field to given value.

### HasMaxDnsName

`func (o *CertificateProfileGet200Response) HasMaxDnsName() bool`

HasMaxDnsName returns a boolean if a field has been set.

### SetMaxDnsNameNil

`func (o *CertificateProfileGet200Response) SetMaxDnsNameNil(b bool)`

 SetMaxDnsNameNil sets the value for MaxDnsName to be an explicit nil

### UnsetMaxDnsName
`func (o *CertificateProfileGet200Response) UnsetMaxDnsName()`

UnsetMaxDnsName ensures that no value is present for MaxDnsName, not even an explicit nil
### GetProxy

`func (o *CertificateProfileGet200Response) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CertificateProfileGet200Response) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CertificateProfileGet200Response) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CertificateProfileGet200Response) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CertificateProfileGet200Response) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CertificateProfileGet200Response) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetAuthorizationLevels

`func (o *CertificateProfileGet200Response) GetAuthorizationLevels() CertificateProfileAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *CertificateProfileGet200Response) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *CertificateProfileGet200Response) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetTriggers

`func (o *CertificateProfileGet200Response) GetTriggers() CertificateProfileTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CertificateProfileGet200Response) GetTriggersOk() (*CertificateProfileTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CertificateProfileGet200Response) SetTriggers(v CertificateProfileTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CertificateProfileGet200Response) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *CertificateProfileGet200Response) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *CertificateProfileGet200Response) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetRequestsPolicy

`func (o *CertificateProfileGet200Response) GetRequestsPolicy() RequestsPolicy`

GetRequestsPolicy returns the RequestsPolicy field if non-nil, zero value otherwise.

### GetRequestsPolicyOk

`func (o *CertificateProfileGet200Response) GetRequestsPolicyOk() (*RequestsPolicy, bool)`

GetRequestsPolicyOk returns a tuple with the RequestsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPolicy

`func (o *CertificateProfileGet200Response) SetRequestsPolicy(v RequestsPolicy)`

SetRequestsPolicy sets RequestsPolicy field to given value.


### GetSelfPermissions

`func (o *CertificateProfileGet200Response) GetSelfPermissions() CertificateProfileSelfPermissions`

GetSelfPermissions returns the SelfPermissions field if non-nil, zero value otherwise.

### GetSelfPermissionsOk

`func (o *CertificateProfileGet200Response) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool)`

GetSelfPermissionsOk returns a tuple with the SelfPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPermissions

`func (o *CertificateProfileGet200Response) SetSelfPermissions(v CertificateProfileSelfPermissions)`

SetSelfPermissions sets SelfPermissions field to given value.


### GetCertificateTemplate

`func (o *CertificateProfileGet200Response) GetCertificateTemplate() CertificateTemplate`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *CertificateProfileGet200Response) GetCertificateTemplateOk() (*CertificateTemplate, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *CertificateProfileGet200Response) SetCertificateTemplate(v CertificateTemplate)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *CertificateProfileGet200Response) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *CertificateProfileGet200Response) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *CertificateProfileGet200Response) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetCryptoPolicy

`func (o *CertificateProfileGet200Response) GetCryptoPolicy() CertificateProfileCryptoPolicy`

GetCryptoPolicy returns the CryptoPolicy field if non-nil, zero value otherwise.

### GetCryptoPolicyOk

`func (o *CertificateProfileGet200Response) GetCryptoPolicyOk() (*CertificateProfileCryptoPolicy, bool)`

GetCryptoPolicyOk returns a tuple with the CryptoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPolicy

`func (o *CertificateProfileGet200Response) SetCryptoPolicy(v CertificateProfileCryptoPolicy)`

SetCryptoPolicy sets CryptoPolicy field to given value.


### GetGradingPolicies

`func (o *CertificateProfileGet200Response) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *CertificateProfileGet200Response) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *CertificateProfileGet200Response) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *CertificateProfileGet200Response) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *CertificateProfileGet200Response) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *CertificateProfileGet200Response) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetDsFlow

`func (o *CertificateProfileGet200Response) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *CertificateProfileGet200Response) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *CertificateProfileGet200Response) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *CertificateProfileGet200Response) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### SetDsFlowNil

`func (o *CertificateProfileGet200Response) SetDsFlowNil(b bool)`

 SetDsFlowNil sets the value for DsFlow to be an explicit nil

### UnsetDsFlow
`func (o *CertificateProfileGet200Response) UnsetDsFlow()`

UnsetDsFlow ensures that no value is present for DsFlow, not even an explicit nil
### GetCa

`func (o *CertificateProfileGet200Response) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *CertificateProfileGet200Response) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *CertificateProfileGet200Response) SetCa(v string)`

SetCa sets Ca field to given value.


### GetAuthorizationMode

`func (o *CertificateProfileGet200Response) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *CertificateProfileGet200Response) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *CertificateProfileGet200Response) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.


### GetDnWhitelist

`func (o *CertificateProfileGet200Response) GetDnWhitelist() bool`

GetDnWhitelist returns the DnWhitelist field if non-nil, zero value otherwise.

### GetDnWhitelistOk

`func (o *CertificateProfileGet200Response) GetDnWhitelistOk() (*bool, bool)`

GetDnWhitelistOk returns a tuple with the DnWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnWhitelist

`func (o *CertificateProfileGet200Response) SetDnWhitelist(v bool)`

SetDnWhitelist sets DnWhitelist field to given value.


### GetEnrollAuthorizedCas

`func (o *CertificateProfileGet200Response) GetEnrollAuthorizedCas() []string`

GetEnrollAuthorizedCas returns the EnrollAuthorizedCas field if non-nil, zero value otherwise.

### GetEnrollAuthorizedCasOk

`func (o *CertificateProfileGet200Response) GetEnrollAuthorizedCasOk() (*[]string, bool)`

GetEnrollAuthorizedCasOk returns a tuple with the EnrollAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollAuthorizedCas

`func (o *CertificateProfileGet200Response) SetEnrollAuthorizedCas(v []string)`

SetEnrollAuthorizedCas sets EnrollAuthorizedCas field to given value.

### HasEnrollAuthorizedCas

`func (o *CertificateProfileGet200Response) HasEnrollAuthorizedCas() bool`

HasEnrollAuthorizedCas returns a boolean if a field has been set.

### SetEnrollAuthorizedCasNil

`func (o *CertificateProfileGet200Response) SetEnrollAuthorizedCasNil(b bool)`

 SetEnrollAuthorizedCasNil sets the value for EnrollAuthorizedCas to be an explicit nil

### UnsetEnrollAuthorizedCas
`func (o *CertificateProfileGet200Response) UnsetEnrollAuthorizedCas()`

UnsetEnrollAuthorizedCas ensures that no value is present for EnrollAuthorizedCas, not even an explicit nil
### GetRenewalAuthorizedCas

`func (o *CertificateProfileGet200Response) GetRenewalAuthorizedCas() []string`

GetRenewalAuthorizedCas returns the RenewalAuthorizedCas field if non-nil, zero value otherwise.

### GetRenewalAuthorizedCasOk

`func (o *CertificateProfileGet200Response) GetRenewalAuthorizedCasOk() (*[]string, bool)`

GetRenewalAuthorizedCasOk returns a tuple with the RenewalAuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalAuthorizedCas

`func (o *CertificateProfileGet200Response) SetRenewalAuthorizedCas(v []string)`

SetRenewalAuthorizedCas sets RenewalAuthorizedCas field to given value.

### HasRenewalAuthorizedCas

`func (o *CertificateProfileGet200Response) HasRenewalAuthorizedCas() bool`

HasRenewalAuthorizedCas returns a boolean if a field has been set.

### SetRenewalAuthorizedCasNil

`func (o *CertificateProfileGet200Response) SetRenewalAuthorizedCasNil(b bool)`

 SetRenewalAuthorizedCasNil sets the value for RenewalAuthorizedCas to be an explicit nil

### UnsetRenewalAuthorizedCas
`func (o *CertificateProfileGet200Response) UnsetRenewalAuthorizedCas()`

UnsetRenewalAuthorizedCas ensures that no value is present for RenewalAuthorizedCas, not even an explicit nil
### GetPasswordPolicy

`func (o *CertificateProfileGet200Response) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *CertificateProfileGet200Response) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *CertificateProfileGet200Response) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *CertificateProfileGet200Response) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *CertificateProfileGet200Response) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *CertificateProfileGet200Response) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetValidationRuleset

`func (o *CertificateProfileGet200Response) GetValidationRuleset() ValidationRuleset`

GetValidationRuleset returns the ValidationRuleset field if non-nil, zero value otherwise.

### GetValidationRulesetOk

`func (o *CertificateProfileGet200Response) GetValidationRulesetOk() (*ValidationRuleset, bool)`

GetValidationRulesetOk returns a tuple with the ValidationRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRuleset

`func (o *CertificateProfileGet200Response) SetValidationRuleset(v ValidationRuleset)`

SetValidationRuleset sets ValidationRuleset field to given value.

### HasValidationRuleset

`func (o *CertificateProfileGet200Response) HasValidationRuleset() bool`

HasValidationRuleset returns a boolean if a field has been set.

### SetValidationRulesetNil

`func (o *CertificateProfileGet200Response) SetValidationRulesetNil(b bool)`

 SetValidationRulesetNil sets the value for ValidationRuleset to be an explicit nil

### UnsetValidationRuleset
`func (o *CertificateProfileGet200Response) UnsetValidationRuleset()`

UnsetValidationRuleset ensures that no value is present for ValidationRuleset, not even an explicit nil
### GetMode

`func (o *CertificateProfileGet200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CertificateProfileGet200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CertificateProfileGet200Response) SetMode(v string)`

SetMode sets Mode field to given value.


### GetThirdPartyConnector

`func (o *CertificateProfileGet200Response) GetThirdPartyConnector() string`

GetThirdPartyConnector returns the ThirdPartyConnector field if non-nil, zero value otherwise.

### GetThirdPartyConnectorOk

`func (o *CertificateProfileGet200Response) GetThirdPartyConnectorOk() (*string, bool)`

GetThirdPartyConnectorOk returns a tuple with the ThirdPartyConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyConnector

`func (o *CertificateProfileGet200Response) SetThirdPartyConnector(v string)`

SetThirdPartyConnector sets ThirdPartyConnector field to given value.


### GetScepRA

`func (o *CertificateProfileGet200Response) GetScepRA() string`

GetScepRA returns the ScepRA field if non-nil, zero value otherwise.

### GetScepRAOk

`func (o *CertificateProfileGet200Response) GetScepRAOk() (*string, bool)`

GetScepRAOk returns a tuple with the ScepRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScepRA

`func (o *CertificateProfileGet200Response) SetScepRA(v string)`

SetScepRA sets ScepRA field to given value.


### GetCaps

`func (o *CertificateProfileGet200Response) GetCaps() []string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *CertificateProfileGet200Response) GetCapsOk() (*[]string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *CertificateProfileGet200Response) SetCaps(v []string)`

SetCaps sets Caps field to given value.


### GetPostPKIOperation

`func (o *CertificateProfileGet200Response) GetPostPKIOperation() bool`

GetPostPKIOperation returns the PostPKIOperation field if non-nil, zero value otherwise.

### GetPostPKIOperationOk

`func (o *CertificateProfileGet200Response) GetPostPKIOperationOk() (*bool, bool)`

GetPostPKIOperationOk returns a tuple with the PostPKIOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPKIOperation

`func (o *CertificateProfileGet200Response) SetPostPKIOperation(v bool)`

SetPostPKIOperation sets PostPKIOperation field to given value.

### HasPostPKIOperation

`func (o *CertificateProfileGet200Response) HasPostPKIOperation() bool`

HasPostPKIOperation returns a boolean if a field has been set.

### SetPostPKIOperationNil

`func (o *CertificateProfileGet200Response) SetPostPKIOperationNil(b bool)`

 SetPostPKIOperationNil sets the value for PostPKIOperation to be an explicit nil

### UnsetPostPKIOperation
`func (o *CertificateProfileGet200Response) UnsetPostPKIOperation()`

UnsetPostPKIOperation ensures that no value is present for PostPKIOperation, not even an explicit nil
### GetEncryptionAlgorithm

`func (o *CertificateProfileGet200Response) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *CertificateProfileGet200Response) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *CertificateProfileGet200Response) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.


### GetDeviceIdField

`func (o *CertificateProfileGet200Response) GetDeviceIdField() string`

GetDeviceIdField returns the DeviceIdField field if non-nil, zero value otherwise.

### GetDeviceIdFieldOk

`func (o *CertificateProfileGet200Response) GetDeviceIdFieldOk() (*string, bool)`

GetDeviceIdFieldOk returns a tuple with the DeviceIdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdField

`func (o *CertificateProfileGet200Response) SetDeviceIdField(v string)`

SetDeviceIdField sets DeviceIdField field to given value.

### HasDeviceIdField

`func (o *CertificateProfileGet200Response) HasDeviceIdField() bool`

HasDeviceIdField returns a boolean if a field has been set.

### SetDeviceIdFieldNil

`func (o *CertificateProfileGet200Response) SetDeviceIdFieldNil(b bool)`

 SetDeviceIdFieldNil sets the value for DeviceIdField to be an explicit nil

### UnsetDeviceIdField
`func (o *CertificateProfileGet200Response) UnsetDeviceIdField()`

UnsetDeviceIdField ensures that no value is present for DeviceIdField, not even an explicit nil
### GetDeviceIdSeparator

`func (o *CertificateProfileGet200Response) GetDeviceIdSeparator() string`

GetDeviceIdSeparator returns the DeviceIdSeparator field if non-nil, zero value otherwise.

### GetDeviceIdSeparatorOk

`func (o *CertificateProfileGet200Response) GetDeviceIdSeparatorOk() (*string, bool)`

GetDeviceIdSeparatorOk returns a tuple with the DeviceIdSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdSeparator

`func (o *CertificateProfileGet200Response) SetDeviceIdSeparator(v string)`

SetDeviceIdSeparator sets DeviceIdSeparator field to given value.

### HasDeviceIdSeparator

`func (o *CertificateProfileGet200Response) HasDeviceIdSeparator() bool`

HasDeviceIdSeparator returns a boolean if a field has been set.

### SetDeviceIdSeparatorNil

`func (o *CertificateProfileGet200Response) SetDeviceIdSeparatorNil(b bool)`

 SetDeviceIdSeparatorNil sets the value for DeviceIdSeparator to be an explicit nil

### UnsetDeviceIdSeparator
`func (o *CertificateProfileGet200Response) UnsetDeviceIdSeparator()`

UnsetDeviceIdSeparator ensures that no value is present for DeviceIdSeparator, not even an explicit nil
### GetExchangeCertificate

`func (o *CertificateProfileGet200Response) GetExchangeCertificate() string`

GetExchangeCertificate returns the ExchangeCertificate field if non-nil, zero value otherwise.

### GetExchangeCertificateOk

`func (o *CertificateProfileGet200Response) GetExchangeCertificateOk() (*string, bool)`

GetExchangeCertificateOk returns a tuple with the ExchangeCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCertificate

`func (o *CertificateProfileGet200Response) SetExchangeCertificate(v string)`

SetExchangeCertificate sets ExchangeCertificate field to given value.

### HasExchangeCertificate

`func (o *CertificateProfileGet200Response) HasExchangeCertificate() bool`

HasExchangeCertificate returns a boolean if a field has been set.

### SetExchangeCertificateNil

`func (o *CertificateProfileGet200Response) SetExchangeCertificateNil(b bool)`

 SetExchangeCertificateNil sets the value for ExchangeCertificate to be an explicit nil

### UnsetExchangeCertificate
`func (o *CertificateProfileGet200Response) UnsetExchangeCertificate()`

UnsetExchangeCertificate ensures that no value is present for ExchangeCertificate, not even an explicit nil
### GetAcmeUrl

`func (o *CertificateProfileGet200Response) GetAcmeUrl() string`

GetAcmeUrl returns the AcmeUrl field if non-nil, zero value otherwise.

### GetAcmeUrlOk

`func (o *CertificateProfileGet200Response) GetAcmeUrlOk() (*string, bool)`

GetAcmeUrlOk returns a tuple with the AcmeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeUrl

`func (o *CertificateProfileGet200Response) SetAcmeUrl(v string)`

SetAcmeUrl sets AcmeUrl field to given value.

### HasAcmeUrl

`func (o *CertificateProfileGet200Response) HasAcmeUrl() bool`

HasAcmeUrl returns a boolean if a field has been set.

### GetRequireEAB

`func (o *CertificateProfileGet200Response) GetRequireEAB() bool`

GetRequireEAB returns the RequireEAB field if non-nil, zero value otherwise.

### GetRequireEABOk

`func (o *CertificateProfileGet200Response) GetRequireEABOk() (*bool, bool)`

GetRequireEABOk returns a tuple with the RequireEAB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEAB

`func (o *CertificateProfileGet200Response) SetRequireEAB(v bool)`

SetRequireEAB sets RequireEAB field to given value.


### GetAuthorizedCas

`func (o *CertificateProfileGet200Response) GetAuthorizedCas() []string`

GetAuthorizedCas returns the AuthorizedCas field if non-nil, zero value otherwise.

### GetAuthorizedCasOk

`func (o *CertificateProfileGet200Response) GetAuthorizedCasOk() (*[]string, bool)`

GetAuthorizedCasOk returns a tuple with the AuthorizedCas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedCas

`func (o *CertificateProfileGet200Response) SetAuthorizedCas(v []string)`

SetAuthorizedCas sets AuthorizedCas field to given value.


### SetAuthorizedCasNil

`func (o *CertificateProfileGet200Response) SetAuthorizedCasNil(b bool)`

 SetAuthorizedCasNil sets the value for AuthorizedCas to be an explicit nil

### UnsetAuthorizedCas
`func (o *CertificateProfileGet200Response) UnsetAuthorizedCas()`

UnsetAuthorizedCas ensures that no value is present for AuthorizedCas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


