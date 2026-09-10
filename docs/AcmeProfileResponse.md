# AcmeProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AuthorizationLevels** | [**CertificateProfileAuthorizationLevels**](CertificateProfileAuthorizationLevels.md) |  | 
**AuthorizationMethods** | Pointer to **[]string** |  | [optional] 
**AuthorizeEmptyContact** | **bool** |  | 
**AuthorizeShortName** | **bool** |  | 
**CertificateTemplate** | Pointer to [**NullableCertificateTemplate**](CertificateTemplate.md) |  | [optional] 
**Constraints** | Pointer to [**NullableCertificateRequestConstraints**](CertificateRequestConstraints.md) |  | [optional] 
**CryptoPolicy** | [**ManagedCertificateProfileCryptoPolicy**](ManagedCertificateProfileCryptoPolicy.md) |  | 
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

## Methods

### NewAcmeProfileResponse

`func NewAcmeProfileResponse(id string, authorizationLevels CertificateProfileAuthorizationLevels, authorizeEmptyContact bool, authorizeShortName bool, cryptoPolicy ManagedCertificateProfileCryptoPolicy, enabled bool, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, requireTermsOfService bool, selfPermissions CertificateProfileSelfPermissions, timeout string, verifyRetryCount int64, verifyRetryDelay string, ) *AcmeProfileResponse`

NewAcmeProfileResponse instantiates a new AcmeProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeProfileResponseWithDefaults

`func NewAcmeProfileResponseWithDefaults() *AcmeProfileResponse`

NewAcmeProfileResponseWithDefaults instantiates a new AcmeProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AcmeProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcmeProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcmeProfileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAuthorizationLevels

`func (o *AcmeProfileResponse) GetAuthorizationLevels() CertificateProfileAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *AcmeProfileResponse) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *AcmeProfileResponse) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetAuthorizationMethods

`func (o *AcmeProfileResponse) GetAuthorizationMethods() []string`

GetAuthorizationMethods returns the AuthorizationMethods field if non-nil, zero value otherwise.

### GetAuthorizationMethodsOk

`func (o *AcmeProfileResponse) GetAuthorizationMethodsOk() (*[]string, bool)`

GetAuthorizationMethodsOk returns a tuple with the AuthorizationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMethods

`func (o *AcmeProfileResponse) SetAuthorizationMethods(v []string)`

SetAuthorizationMethods sets AuthorizationMethods field to given value.

### HasAuthorizationMethods

`func (o *AcmeProfileResponse) HasAuthorizationMethods() bool`

HasAuthorizationMethods returns a boolean if a field has been set.

### SetAuthorizationMethodsNil

`func (o *AcmeProfileResponse) SetAuthorizationMethodsNil(b bool)`

 SetAuthorizationMethodsNil sets the value for AuthorizationMethods to be an explicit nil

### UnsetAuthorizationMethods
`func (o *AcmeProfileResponse) UnsetAuthorizationMethods()`

UnsetAuthorizationMethods ensures that no value is present for AuthorizationMethods, not even an explicit nil
### GetAuthorizeEmptyContact

`func (o *AcmeProfileResponse) GetAuthorizeEmptyContact() bool`

GetAuthorizeEmptyContact returns the AuthorizeEmptyContact field if non-nil, zero value otherwise.

### GetAuthorizeEmptyContactOk

`func (o *AcmeProfileResponse) GetAuthorizeEmptyContactOk() (*bool, bool)`

GetAuthorizeEmptyContactOk returns a tuple with the AuthorizeEmptyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeEmptyContact

`func (o *AcmeProfileResponse) SetAuthorizeEmptyContact(v bool)`

SetAuthorizeEmptyContact sets AuthorizeEmptyContact field to given value.


### GetAuthorizeShortName

`func (o *AcmeProfileResponse) GetAuthorizeShortName() bool`

GetAuthorizeShortName returns the AuthorizeShortName field if non-nil, zero value otherwise.

### GetAuthorizeShortNameOk

`func (o *AcmeProfileResponse) GetAuthorizeShortNameOk() (*bool, bool)`

GetAuthorizeShortNameOk returns a tuple with the AuthorizeShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeShortName

`func (o *AcmeProfileResponse) SetAuthorizeShortName(v bool)`

SetAuthorizeShortName sets AuthorizeShortName field to given value.


### GetCertificateTemplate

`func (o *AcmeProfileResponse) GetCertificateTemplate() CertificateTemplate`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *AcmeProfileResponse) GetCertificateTemplateOk() (*CertificateTemplate, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *AcmeProfileResponse) SetCertificateTemplate(v CertificateTemplate)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *AcmeProfileResponse) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *AcmeProfileResponse) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *AcmeProfileResponse) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetConstraints

`func (o *AcmeProfileResponse) GetConstraints() CertificateRequestConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AcmeProfileResponse) GetConstraintsOk() (*CertificateRequestConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AcmeProfileResponse) SetConstraints(v CertificateRequestConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *AcmeProfileResponse) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *AcmeProfileResponse) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *AcmeProfileResponse) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetCryptoPolicy

`func (o *AcmeProfileResponse) GetCryptoPolicy() ManagedCertificateProfileCryptoPolicy`

GetCryptoPolicy returns the CryptoPolicy field if non-nil, zero value otherwise.

### GetCryptoPolicyOk

`func (o *AcmeProfileResponse) GetCryptoPolicyOk() (*ManagedCertificateProfileCryptoPolicy, bool)`

GetCryptoPolicyOk returns a tuple with the CryptoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPolicy

`func (o *AcmeProfileResponse) SetCryptoPolicy(v ManagedCertificateProfileCryptoPolicy)`

SetCryptoPolicy sets CryptoPolicy field to given value.


### GetCsrDataMapping

`func (o *AcmeProfileResponse) GetCsrDataMapping() map[string]string`

GetCsrDataMapping returns the CsrDataMapping field if non-nil, zero value otherwise.

### GetCsrDataMappingOk

`func (o *AcmeProfileResponse) GetCsrDataMappingOk() (*map[string]string, bool)`

GetCsrDataMappingOk returns a tuple with the CsrDataMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataMapping

`func (o *AcmeProfileResponse) SetCsrDataMapping(v map[string]string)`

SetCsrDataMapping sets CsrDataMapping field to given value.

### HasCsrDataMapping

`func (o *AcmeProfileResponse) HasCsrDataMapping() bool`

HasCsrDataMapping returns a boolean if a field has been set.

### SetCsrDataMappingNil

`func (o *AcmeProfileResponse) SetCsrDataMappingNil(b bool)`

 SetCsrDataMappingNil sets the value for CsrDataMapping to be an explicit nil

### UnsetCsrDataMapping
`func (o *AcmeProfileResponse) UnsetCsrDataMapping()`

UnsetCsrDataMapping ensures that no value is present for CsrDataMapping, not even an explicit nil
### GetDefaultContacts

`func (o *AcmeProfileResponse) GetDefaultContacts() []string`

GetDefaultContacts returns the DefaultContacts field if non-nil, zero value otherwise.

### GetDefaultContactsOk

`func (o *AcmeProfileResponse) GetDefaultContactsOk() (*[]string, bool)`

GetDefaultContactsOk returns a tuple with the DefaultContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContacts

`func (o *AcmeProfileResponse) SetDefaultContacts(v []string)`

SetDefaultContacts sets DefaultContacts field to given value.

### HasDefaultContacts

`func (o *AcmeProfileResponse) HasDefaultContacts() bool`

HasDefaultContacts returns a boolean if a field has been set.

### SetDefaultContactsNil

`func (o *AcmeProfileResponse) SetDefaultContactsNil(b bool)`

 SetDefaultContactsNil sets the value for DefaultContacts to be an explicit nil

### UnsetDefaultContacts
`func (o *AcmeProfileResponse) UnsetDefaultContacts()`

UnsetDefaultContacts ensures that no value is present for DefaultContacts, not even an explicit nil
### GetDescription

`func (o *AcmeProfileResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AcmeProfileResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AcmeProfileResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AcmeProfileResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AcmeProfileResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AcmeProfileResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *AcmeProfileResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AcmeProfileResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AcmeProfileResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AcmeProfileResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AcmeProfileResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AcmeProfileResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDsFlow

`func (o *AcmeProfileResponse) GetDsFlow() []DataSourceFlowEntry`

GetDsFlow returns the DsFlow field if non-nil, zero value otherwise.

### GetDsFlowOk

`func (o *AcmeProfileResponse) GetDsFlowOk() (*[]DataSourceFlowEntry, bool)`

GetDsFlowOk returns a tuple with the DsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFlow

`func (o *AcmeProfileResponse) SetDsFlow(v []DataSourceFlowEntry)`

SetDsFlow sets DsFlow field to given value.

### HasDsFlow

`func (o *AcmeProfileResponse) HasDsFlow() bool`

HasDsFlow returns a boolean if a field has been set.

### SetDsFlowNil

`func (o *AcmeProfileResponse) SetDsFlowNil(b bool)`

 SetDsFlowNil sets the value for DsFlow to be an explicit nil

### UnsetDsFlow
`func (o *AcmeProfileResponse) UnsetDsFlow()`

UnsetDsFlow ensures that no value is present for DsFlow, not even an explicit nil
### GetEnabled

`func (o *AcmeProfileResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AcmeProfileResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AcmeProfileResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGradingPolicies

`func (o *AcmeProfileResponse) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *AcmeProfileResponse) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *AcmeProfileResponse) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *AcmeProfileResponse) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *AcmeProfileResponse) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *AcmeProfileResponse) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetHttp01Port

`func (o *AcmeProfileResponse) GetHttp01Port() int64`

GetHttp01Port returns the Http01Port field if non-nil, zero value otherwise.

### GetHttp01PortOk

`func (o *AcmeProfileResponse) GetHttp01PortOk() (*int64, bool)`

GetHttp01PortOk returns a tuple with the Http01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp01Port

`func (o *AcmeProfileResponse) SetHttp01Port(v int64)`

SetHttp01Port sets Http01Port field to given value.

### HasHttp01Port

`func (o *AcmeProfileResponse) HasHttp01Port() bool`

HasHttp01Port returns a boolean if a field has been set.

### SetHttp01PortNil

`func (o *AcmeProfileResponse) SetHttp01PortNil(b bool)`

 SetHttp01PortNil sets the value for Http01Port to be an explicit nil

### UnsetHttp01Port
`func (o *AcmeProfileResponse) UnsetHttp01Port()`

UnsetHttp01Port ensures that no value is present for Http01Port, not even an explicit nil
### GetMaxCertificatePerHolderPolicy

`func (o *AcmeProfileResponse) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy`

GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field if non-nil, zero value otherwise.

### GetMaxCertificatePerHolderPolicyOk

`func (o *AcmeProfileResponse) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool)`

GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCertificatePerHolderPolicy

`func (o *AcmeProfileResponse) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy)`

SetMaxCertificatePerHolderPolicy sets MaxCertificatePerHolderPolicy field to given value.

### HasMaxCertificatePerHolderPolicy

`func (o *AcmeProfileResponse) HasMaxCertificatePerHolderPolicy() bool`

HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.

### SetMaxCertificatePerHolderPolicyNil

`func (o *AcmeProfileResponse) SetMaxCertificatePerHolderPolicyNil(b bool)`

 SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil

### UnsetMaxCertificatePerHolderPolicy
`func (o *AcmeProfileResponse) UnsetMaxCertificatePerHolderPolicy()`

UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
### GetMaxDnsName

`func (o *AcmeProfileResponse) GetMaxDnsName() int64`

GetMaxDnsName returns the MaxDnsName field if non-nil, zero value otherwise.

### GetMaxDnsNameOk

`func (o *AcmeProfileResponse) GetMaxDnsNameOk() (*int64, bool)`

GetMaxDnsNameOk returns a tuple with the MaxDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDnsName

`func (o *AcmeProfileResponse) SetMaxDnsName(v int64)`

SetMaxDnsName sets MaxDnsName field to given value.

### HasMaxDnsName

`func (o *AcmeProfileResponse) HasMaxDnsName() bool`

HasMaxDnsName returns a boolean if a field has been set.

### SetMaxDnsNameNil

`func (o *AcmeProfileResponse) SetMaxDnsNameNil(b bool)`

 SetMaxDnsNameNil sets the value for MaxDnsName to be an explicit nil

### UnsetMaxDnsName
`func (o *AcmeProfileResponse) UnsetMaxDnsName()`

UnsetMaxDnsName ensures that no value is present for MaxDnsName, not even an explicit nil
### GetMeta

`func (o *AcmeProfileResponse) GetMeta() DirectoryMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AcmeProfileResponse) GetMetaOk() (*DirectoryMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AcmeProfileResponse) SetMeta(v DirectoryMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AcmeProfileResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *AcmeProfileResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *AcmeProfileResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetModule

`func (o *AcmeProfileResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *AcmeProfileResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *AcmeProfileResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetName

`func (o *AcmeProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcmeProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcmeProfileResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPkiConnector

`func (o *AcmeProfileResponse) GetPkiConnector() string`

GetPkiConnector returns the PkiConnector field if non-nil, zero value otherwise.

### GetPkiConnectorOk

`func (o *AcmeProfileResponse) GetPkiConnectorOk() (*string, bool)`

GetPkiConnectorOk returns a tuple with the PkiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnector

`func (o *AcmeProfileResponse) SetPkiConnector(v string)`

SetPkiConnector sets PkiConnector field to given value.


### GetProxy

`func (o *AcmeProfileResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AcmeProfileResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AcmeProfileResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AcmeProfileResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *AcmeProfileResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *AcmeProfileResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRenewalPeriod

`func (o *AcmeProfileResponse) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *AcmeProfileResponse) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *AcmeProfileResponse) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *AcmeProfileResponse) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *AcmeProfileResponse) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *AcmeProfileResponse) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetRequestsPolicy

`func (o *AcmeProfileResponse) GetRequestsPolicy() RequestsPolicy`

GetRequestsPolicy returns the RequestsPolicy field if non-nil, zero value otherwise.

### GetRequestsPolicyOk

`func (o *AcmeProfileResponse) GetRequestsPolicyOk() (*RequestsPolicy, bool)`

GetRequestsPolicyOk returns a tuple with the RequestsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPolicy

`func (o *AcmeProfileResponse) SetRequestsPolicy(v RequestsPolicy)`

SetRequestsPolicy sets RequestsPolicy field to given value.


### GetRequireTermsOfService

`func (o *AcmeProfileResponse) GetRequireTermsOfService() bool`

GetRequireTermsOfService returns the RequireTermsOfService field if non-nil, zero value otherwise.

### GetRequireTermsOfServiceOk

`func (o *AcmeProfileResponse) GetRequireTermsOfServiceOk() (*bool, bool)`

GetRequireTermsOfServiceOk returns a tuple with the RequireTermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTermsOfService

`func (o *AcmeProfileResponse) SetRequireTermsOfService(v bool)`

SetRequireTermsOfService sets RequireTermsOfService field to given value.


### GetSelfPermissions

`func (o *AcmeProfileResponse) GetSelfPermissions() CertificateProfileSelfPermissions`

GetSelfPermissions returns the SelfPermissions field if non-nil, zero value otherwise.

### GetSelfPermissionsOk

`func (o *AcmeProfileResponse) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool)`

GetSelfPermissionsOk returns a tuple with the SelfPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPermissions

`func (o *AcmeProfileResponse) SetSelfPermissions(v CertificateProfileSelfPermissions)`

SetSelfPermissions sets SelfPermissions field to given value.


### GetThirdPartyDiscoverySync

`func (o *AcmeProfileResponse) GetThirdPartyDiscoverySync() bool`

GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field if non-nil, zero value otherwise.

### GetThirdPartyDiscoverySyncOk

`func (o *AcmeProfileResponse) GetThirdPartyDiscoverySyncOk() (*bool, bool)`

GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyDiscoverySync

`func (o *AcmeProfileResponse) SetThirdPartyDiscoverySync(v bool)`

SetThirdPartyDiscoverySync sets ThirdPartyDiscoverySync field to given value.

### HasThirdPartyDiscoverySync

`func (o *AcmeProfileResponse) HasThirdPartyDiscoverySync() bool`

HasThirdPartyDiscoverySync returns a boolean if a field has been set.

### SetThirdPartyDiscoverySyncNil

`func (o *AcmeProfileResponse) SetThirdPartyDiscoverySyncNil(b bool)`

 SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil

### UnsetThirdPartyDiscoverySync
`func (o *AcmeProfileResponse) UnsetThirdPartyDiscoverySync()`

UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
### GetTimeout

`func (o *AcmeProfileResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AcmeProfileResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AcmeProfileResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTlsAlpn01Port

`func (o *AcmeProfileResponse) GetTlsAlpn01Port() int64`

GetTlsAlpn01Port returns the TlsAlpn01Port field if non-nil, zero value otherwise.

### GetTlsAlpn01PortOk

`func (o *AcmeProfileResponse) GetTlsAlpn01PortOk() (*int64, bool)`

GetTlsAlpn01PortOk returns a tuple with the TlsAlpn01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAlpn01Port

`func (o *AcmeProfileResponse) SetTlsAlpn01Port(v int64)`

SetTlsAlpn01Port sets TlsAlpn01Port field to given value.

### HasTlsAlpn01Port

`func (o *AcmeProfileResponse) HasTlsAlpn01Port() bool`

HasTlsAlpn01Port returns a boolean if a field has been set.

### SetTlsAlpn01PortNil

`func (o *AcmeProfileResponse) SetTlsAlpn01PortNil(b bool)`

 SetTlsAlpn01PortNil sets the value for TlsAlpn01Port to be an explicit nil

### UnsetTlsAlpn01Port
`func (o *AcmeProfileResponse) UnsetTlsAlpn01Port()`

UnsetTlsAlpn01Port ensures that no value is present for TlsAlpn01Port, not even an explicit nil
### GetTriggers

`func (o *AcmeProfileResponse) GetTriggers() CertificateProfileTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *AcmeProfileResponse) GetTriggersOk() (*CertificateProfileTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *AcmeProfileResponse) SetTriggers(v CertificateProfileTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *AcmeProfileResponse) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *AcmeProfileResponse) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *AcmeProfileResponse) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetVerifyRetryCount

`func (o *AcmeProfileResponse) GetVerifyRetryCount() int64`

GetVerifyRetryCount returns the VerifyRetryCount field if non-nil, zero value otherwise.

### GetVerifyRetryCountOk

`func (o *AcmeProfileResponse) GetVerifyRetryCountOk() (*int64, bool)`

GetVerifyRetryCountOk returns a tuple with the VerifyRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryCount

`func (o *AcmeProfileResponse) SetVerifyRetryCount(v int64)`

SetVerifyRetryCount sets VerifyRetryCount field to given value.


### GetVerifyRetryDelay

`func (o *AcmeProfileResponse) GetVerifyRetryDelay() string`

GetVerifyRetryDelay returns the VerifyRetryDelay field if non-nil, zero value otherwise.

### GetVerifyRetryDelayOk

`func (o *AcmeProfileResponse) GetVerifyRetryDelayOk() (*string, bool)`

GetVerifyRetryDelayOk returns a tuple with the VerifyRetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyRetryDelay

`func (o *AcmeProfileResponse) SetVerifyRetryDelay(v string)`

SetVerifyRetryDelay sets VerifyRetryDelay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


