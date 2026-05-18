# SecurityIdentityProviderAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the identity provider | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the identity provider | [optional] 
**EmailTemplate** | Pointer to [**NullableEmailTemplate**](EmailTemplate.md) | The e-mail template to use for password recovery | [optional] 
**Enabled** | **bool** | Whether the identity provider can be used to identify against Horizon | 
**EnabledOnUI** | **bool** | Whether the identity provider can be selected on login to the Horizon UI | 
**Name** | **string** | The internal name of the identity provider | 
**PasswordPolicy** | Pointer to **NullableString** | The password policy to enforce for user passwords on the local identity provider | [optional] 
**Type** | **string** | The type of identity provider | 
**ClientCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the client ID  and secret to use to authenticate Horizon against the identity provider | 
**EmailClaim** | Pointer to **string** | The OpenID information that will be used as the user&#39;s email in Horizon | [optional] [default to "{{email}}"]
**IdentifierClaim** | Pointer to **string** | The OpenID information that will be used as the user&#39;s identifier in Horizon | [optional] [default to "{{email}}"]
**Mapping** | Pointer to [**OidcIdentityProviderMapping**](OidcIdentityProviderMapping.md) |  | [optional] 
**NameClaim** | Pointer to **string** | The OpenID information that will be used as the user&#39;s name in Horizon | [optional] [default to "{{name}}"]
**ProviderMetadataUrl** | **string** | The URL of the identity provider OpenID callback | 
**Proxy** | Pointer to **NullableString** | The name of the proxy to use to reach the identity provider | [optional] 
**Scope** | **string** | The scope where to retrieve the user data from | 
**Timeout** | **string** | The timeout value to use when connecting to the identity provider (must be a valid finite duration) | 
**TrustSystemCAs** | **bool** | Trust AC coming from the system trust store or only trust AC imported in Horizon | [default to true]

## Methods

### NewSecurityIdentityProviderAddRequest

`func NewSecurityIdentityProviderAddRequest(enabled bool, enabledOnUI bool, name string, type_ string, clientCredentials string, providerMetadataUrl string, scope string, timeout string, trustSystemCAs bool, ) *SecurityIdentityProviderAddRequest`

NewSecurityIdentityProviderAddRequest instantiates a new SecurityIdentityProviderAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIdentityProviderAddRequestWithDefaults

`func NewSecurityIdentityProviderAddRequestWithDefaults() *SecurityIdentityProviderAddRequest`

NewSecurityIdentityProviderAddRequestWithDefaults instantiates a new SecurityIdentityProviderAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityIdentityProviderAddRequest) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityIdentityProviderAddRequest) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityIdentityProviderAddRequest) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityIdentityProviderAddRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityIdentityProviderAddRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityIdentityProviderAddRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *SecurityIdentityProviderAddRequest) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SecurityIdentityProviderAddRequest) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SecurityIdentityProviderAddRequest) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SecurityIdentityProviderAddRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SecurityIdentityProviderAddRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SecurityIdentityProviderAddRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmailTemplate

`func (o *SecurityIdentityProviderAddRequest) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *SecurityIdentityProviderAddRequest) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *SecurityIdentityProviderAddRequest) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *SecurityIdentityProviderAddRequest) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### SetEmailTemplateNil

`func (o *SecurityIdentityProviderAddRequest) SetEmailTemplateNil(b bool)`

 SetEmailTemplateNil sets the value for EmailTemplate to be an explicit nil

### UnsetEmailTemplate
`func (o *SecurityIdentityProviderAddRequest) UnsetEmailTemplate()`

UnsetEmailTemplate ensures that no value is present for EmailTemplate, not even an explicit nil
### GetEnabled

`func (o *SecurityIdentityProviderAddRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityIdentityProviderAddRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityIdentityProviderAddRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledOnUI

`func (o *SecurityIdentityProviderAddRequest) GetEnabledOnUI() bool`

GetEnabledOnUI returns the EnabledOnUI field if non-nil, zero value otherwise.

### GetEnabledOnUIOk

`func (o *SecurityIdentityProviderAddRequest) GetEnabledOnUIOk() (*bool, bool)`

GetEnabledOnUIOk returns a tuple with the EnabledOnUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnUI

`func (o *SecurityIdentityProviderAddRequest) SetEnabledOnUI(v bool)`

SetEnabledOnUI sets EnabledOnUI field to given value.


### GetName

`func (o *SecurityIdentityProviderAddRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityIdentityProviderAddRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityIdentityProviderAddRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordPolicy

`func (o *SecurityIdentityProviderAddRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *SecurityIdentityProviderAddRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *SecurityIdentityProviderAddRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *SecurityIdentityProviderAddRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *SecurityIdentityProviderAddRequest) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *SecurityIdentityProviderAddRequest) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetType

`func (o *SecurityIdentityProviderAddRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityIdentityProviderAddRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityIdentityProviderAddRequest) SetType(v string)`

SetType sets Type field to given value.


### GetClientCredentials

`func (o *SecurityIdentityProviderAddRequest) GetClientCredentials() string`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *SecurityIdentityProviderAddRequest) GetClientCredentialsOk() (*string, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *SecurityIdentityProviderAddRequest) SetClientCredentials(v string)`

SetClientCredentials sets ClientCredentials field to given value.


### GetEmailClaim

`func (o *SecurityIdentityProviderAddRequest) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *SecurityIdentityProviderAddRequest) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *SecurityIdentityProviderAddRequest) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.

### HasEmailClaim

`func (o *SecurityIdentityProviderAddRequest) HasEmailClaim() bool`

HasEmailClaim returns a boolean if a field has been set.

### GetIdentifierClaim

`func (o *SecurityIdentityProviderAddRequest) GetIdentifierClaim() string`

GetIdentifierClaim returns the IdentifierClaim field if non-nil, zero value otherwise.

### GetIdentifierClaimOk

`func (o *SecurityIdentityProviderAddRequest) GetIdentifierClaimOk() (*string, bool)`

GetIdentifierClaimOk returns a tuple with the IdentifierClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierClaim

`func (o *SecurityIdentityProviderAddRequest) SetIdentifierClaim(v string)`

SetIdentifierClaim sets IdentifierClaim field to given value.

### HasIdentifierClaim

`func (o *SecurityIdentityProviderAddRequest) HasIdentifierClaim() bool`

HasIdentifierClaim returns a boolean if a field has been set.

### GetMapping

`func (o *SecurityIdentityProviderAddRequest) GetMapping() OidcIdentityProviderMapping`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *SecurityIdentityProviderAddRequest) GetMappingOk() (*OidcIdentityProviderMapping, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *SecurityIdentityProviderAddRequest) SetMapping(v OidcIdentityProviderMapping)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *SecurityIdentityProviderAddRequest) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetNameClaim

`func (o *SecurityIdentityProviderAddRequest) GetNameClaim() string`

GetNameClaim returns the NameClaim field if non-nil, zero value otherwise.

### GetNameClaimOk

`func (o *SecurityIdentityProviderAddRequest) GetNameClaimOk() (*string, bool)`

GetNameClaimOk returns a tuple with the NameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameClaim

`func (o *SecurityIdentityProviderAddRequest) SetNameClaim(v string)`

SetNameClaim sets NameClaim field to given value.

### HasNameClaim

`func (o *SecurityIdentityProviderAddRequest) HasNameClaim() bool`

HasNameClaim returns a boolean if a field has been set.

### GetProviderMetadataUrl

`func (o *SecurityIdentityProviderAddRequest) GetProviderMetadataUrl() string`

GetProviderMetadataUrl returns the ProviderMetadataUrl field if non-nil, zero value otherwise.

### GetProviderMetadataUrlOk

`func (o *SecurityIdentityProviderAddRequest) GetProviderMetadataUrlOk() (*string, bool)`

GetProviderMetadataUrlOk returns a tuple with the ProviderMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadataUrl

`func (o *SecurityIdentityProviderAddRequest) SetProviderMetadataUrl(v string)`

SetProviderMetadataUrl sets ProviderMetadataUrl field to given value.


### GetProxy

`func (o *SecurityIdentityProviderAddRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SecurityIdentityProviderAddRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SecurityIdentityProviderAddRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SecurityIdentityProviderAddRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *SecurityIdentityProviderAddRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *SecurityIdentityProviderAddRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetScope

`func (o *SecurityIdentityProviderAddRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecurityIdentityProviderAddRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecurityIdentityProviderAddRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTimeout

`func (o *SecurityIdentityProviderAddRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecurityIdentityProviderAddRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecurityIdentityProviderAddRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTrustSystemCAs

`func (o *SecurityIdentityProviderAddRequest) GetTrustSystemCAs() bool`

GetTrustSystemCAs returns the TrustSystemCAs field if non-nil, zero value otherwise.

### GetTrustSystemCAsOk

`func (o *SecurityIdentityProviderAddRequest) GetTrustSystemCAsOk() (*bool, bool)`

GetTrustSystemCAsOk returns a tuple with the TrustSystemCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustSystemCAs

`func (o *SecurityIdentityProviderAddRequest) SetTrustSystemCAs(v bool)`

SetTrustSystemCAs sets TrustSystemCAs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


