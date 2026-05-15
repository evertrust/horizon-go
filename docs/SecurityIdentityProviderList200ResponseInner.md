# SecurityIdentityProviderList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The internal ID of the Identity Provider | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the identity provider | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the identity provider | [optional] 
**EmailTemplate** | Pointer to [**NullableEmailTemplate**](EmailTemplate.md) | The e-mail template to use for password recovery | [optional] 
**Enabled** | **bool** | Whether the identity provider can be used to identify against Horizon | 
**EnabledOnUI** | **bool** | Whether the identity provider can be selected on login to the Horizon UI | 
**Name** | **string** | The internal name of the identity provider | 
**PasswordPolicy** | Pointer to **NullableString** | The password policy to enforce for user passwords on the local identity provider | [optional] 
**Type** | **string** | The type of identity provider | 
**ClientCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the client ID  and secret to use to authenticate Horizon against the identity provider | 
**EmailClaim** | **string** | The OpenID information that will be used as the user&#39;s email in Horizon | 
**IdentifierClaim** | **string** | The OpenID information that will be used as the user&#39;s identifier in Horizon | 
**Mapping** | Pointer to [**OidcIdentityProviderMapping**](OidcIdentityProviderMapping.md) |  | [optional] 
**NameClaim** | **string** | The OpenID information that will be used as the user&#39;s name in Horizon | 
**ProviderMetadataUrl** | **string** | The URL of the identity provider OpenID callback | 
**Proxy** | Pointer to **NullableString** | The name of the proxy to use to reach the identity provider | [optional] 
**Scope** | **string** | The scope where to retrieve the user data from | 
**Timeout** | Pointer to **NullableString** | The timeout value to use when connecting to the identity provider (must be a valid finite duration) | [optional] 
**TrustSystemCAs** | **bool** | Trust AC coming from the system trust store or only trust AC imported in Horizon | [default to true]

## Methods

### NewSecurityIdentityProviderList200ResponseInner

`func NewSecurityIdentityProviderList200ResponseInner(id string, enabled bool, enabledOnUI bool, name string, type_ string, clientCredentials string, emailClaim string, identifierClaim string, nameClaim string, providerMetadataUrl string, scope string, trustSystemCAs bool, ) *SecurityIdentityProviderList200ResponseInner`

NewSecurityIdentityProviderList200ResponseInner instantiates a new SecurityIdentityProviderList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIdentityProviderList200ResponseInnerWithDefaults

`func NewSecurityIdentityProviderList200ResponseInnerWithDefaults() *SecurityIdentityProviderList200ResponseInner`

NewSecurityIdentityProviderList200ResponseInnerWithDefaults instantiates a new SecurityIdentityProviderList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityIdentityProviderList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityIdentityProviderList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *SecurityIdentityProviderList200ResponseInner) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityIdentityProviderList200ResponseInner) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityIdentityProviderList200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityIdentityProviderList200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityIdentityProviderList200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *SecurityIdentityProviderList200ResponseInner) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SecurityIdentityProviderList200ResponseInner) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SecurityIdentityProviderList200ResponseInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SecurityIdentityProviderList200ResponseInner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SecurityIdentityProviderList200ResponseInner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmailTemplate

`func (o *SecurityIdentityProviderList200ResponseInner) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *SecurityIdentityProviderList200ResponseInner) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *SecurityIdentityProviderList200ResponseInner) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### SetEmailTemplateNil

`func (o *SecurityIdentityProviderList200ResponseInner) SetEmailTemplateNil(b bool)`

 SetEmailTemplateNil sets the value for EmailTemplate to be an explicit nil

### UnsetEmailTemplate
`func (o *SecurityIdentityProviderList200ResponseInner) UnsetEmailTemplate()`

UnsetEmailTemplate ensures that no value is present for EmailTemplate, not even an explicit nil
### GetEnabled

`func (o *SecurityIdentityProviderList200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityIdentityProviderList200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledOnUI

`func (o *SecurityIdentityProviderList200ResponseInner) GetEnabledOnUI() bool`

GetEnabledOnUI returns the EnabledOnUI field if non-nil, zero value otherwise.

### GetEnabledOnUIOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetEnabledOnUIOk() (*bool, bool)`

GetEnabledOnUIOk returns a tuple with the EnabledOnUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnUI

`func (o *SecurityIdentityProviderList200ResponseInner) SetEnabledOnUI(v bool)`

SetEnabledOnUI sets EnabledOnUI field to given value.


### GetName

`func (o *SecurityIdentityProviderList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityIdentityProviderList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordPolicy

`func (o *SecurityIdentityProviderList200ResponseInner) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *SecurityIdentityProviderList200ResponseInner) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *SecurityIdentityProviderList200ResponseInner) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *SecurityIdentityProviderList200ResponseInner) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *SecurityIdentityProviderList200ResponseInner) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetType

`func (o *SecurityIdentityProviderList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityIdentityProviderList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetClientCredentials

`func (o *SecurityIdentityProviderList200ResponseInner) GetClientCredentials() string`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetClientCredentialsOk() (*string, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *SecurityIdentityProviderList200ResponseInner) SetClientCredentials(v string)`

SetClientCredentials sets ClientCredentials field to given value.


### GetEmailClaim

`func (o *SecurityIdentityProviderList200ResponseInner) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *SecurityIdentityProviderList200ResponseInner) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.


### GetIdentifierClaim

`func (o *SecurityIdentityProviderList200ResponseInner) GetIdentifierClaim() string`

GetIdentifierClaim returns the IdentifierClaim field if non-nil, zero value otherwise.

### GetIdentifierClaimOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetIdentifierClaimOk() (*string, bool)`

GetIdentifierClaimOk returns a tuple with the IdentifierClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierClaim

`func (o *SecurityIdentityProviderList200ResponseInner) SetIdentifierClaim(v string)`

SetIdentifierClaim sets IdentifierClaim field to given value.


### GetMapping

`func (o *SecurityIdentityProviderList200ResponseInner) GetMapping() OidcIdentityProviderMapping`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetMappingOk() (*OidcIdentityProviderMapping, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *SecurityIdentityProviderList200ResponseInner) SetMapping(v OidcIdentityProviderMapping)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *SecurityIdentityProviderList200ResponseInner) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetNameClaim

`func (o *SecurityIdentityProviderList200ResponseInner) GetNameClaim() string`

GetNameClaim returns the NameClaim field if non-nil, zero value otherwise.

### GetNameClaimOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetNameClaimOk() (*string, bool)`

GetNameClaimOk returns a tuple with the NameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameClaim

`func (o *SecurityIdentityProviderList200ResponseInner) SetNameClaim(v string)`

SetNameClaim sets NameClaim field to given value.


### GetProviderMetadataUrl

`func (o *SecurityIdentityProviderList200ResponseInner) GetProviderMetadataUrl() string`

GetProviderMetadataUrl returns the ProviderMetadataUrl field if non-nil, zero value otherwise.

### GetProviderMetadataUrlOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetProviderMetadataUrlOk() (*string, bool)`

GetProviderMetadataUrlOk returns a tuple with the ProviderMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadataUrl

`func (o *SecurityIdentityProviderList200ResponseInner) SetProviderMetadataUrl(v string)`

SetProviderMetadataUrl sets ProviderMetadataUrl field to given value.


### GetProxy

`func (o *SecurityIdentityProviderList200ResponseInner) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SecurityIdentityProviderList200ResponseInner) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SecurityIdentityProviderList200ResponseInner) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *SecurityIdentityProviderList200ResponseInner) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *SecurityIdentityProviderList200ResponseInner) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetScope

`func (o *SecurityIdentityProviderList200ResponseInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecurityIdentityProviderList200ResponseInner) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTimeout

`func (o *SecurityIdentityProviderList200ResponseInner) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecurityIdentityProviderList200ResponseInner) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SecurityIdentityProviderList200ResponseInner) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *SecurityIdentityProviderList200ResponseInner) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SecurityIdentityProviderList200ResponseInner) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTrustSystemCAs

`func (o *SecurityIdentityProviderList200ResponseInner) GetTrustSystemCAs() bool`

GetTrustSystemCAs returns the TrustSystemCAs field if non-nil, zero value otherwise.

### GetTrustSystemCAsOk

`func (o *SecurityIdentityProviderList200ResponseInner) GetTrustSystemCAsOk() (*bool, bool)`

GetTrustSystemCAsOk returns a tuple with the TrustSystemCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustSystemCAs

`func (o *SecurityIdentityProviderList200ResponseInner) SetTrustSystemCAs(v bool)`

SetTrustSystemCAs sets TrustSystemCAs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


