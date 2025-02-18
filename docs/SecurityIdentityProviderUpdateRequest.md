# SecurityIdentityProviderUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The internal name of the identity provider | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the identity provider | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the identity provider | [optional] 
**Type** | **string** | The type of Identity provider to register | 
**Enabled** | **bool** | Whether the identity provider can be used to identify against Horizon | 
**EnabledOnUI** | **bool** | Whether the identity provider can be selected on login to the Horizon UI | 
**PasswordPolicy** | Pointer to **NullableString** | The password policy to enforce for user passwords on the local identity provider | [optional] 
**EmailTemplate** | Pointer to [**NullableEmailTemplate**](EmailTemplate.md) | The e-mail template to use for password recovery | [optional] 
**Proxy** | Pointer to **NullableString** | The name of the proxy to use to reach the identity provider | [optional] 
**Timeout** | Pointer to **NullableString** | The timeout value to use when connecting to the identity provider (must be a valid finite duration) | [optional] 
**ProviderMetadataUrl** | **string** | The URL of the identity provider OpenID callback | 
**ClientCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/api.security.credentials) containing the client ID  and secret to use to authenticate Horizon against the identity provider | 
**Scope** | **string** | The scope where to retrieve the user data from | 
**TrustSystemCAs** | **bool** | Trust AC coming from the system trust store or only trust AC imported in Horizon | [default to true]
**IdentifierClaim** | **string** | The OpenID information that will be used as the user&#39;s identifier in Horizon | 
**EmailClaim** | **string** | The OpenID information that will be used as the user&#39;s email in Horizon | 
**NameClaim** | **string** | The OpenID information that will be used as the user&#39;s name in Horizon | 

## Methods

### NewSecurityIdentityProviderUpdateRequest

`func NewSecurityIdentityProviderUpdateRequest(name string, type_ string, enabled bool, enabledOnUI bool, providerMetadataUrl string, clientCredentials string, scope string, trustSystemCAs bool, identifierClaim string, emailClaim string, nameClaim string, ) *SecurityIdentityProviderUpdateRequest`

NewSecurityIdentityProviderUpdateRequest instantiates a new SecurityIdentityProviderUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIdentityProviderUpdateRequestWithDefaults

`func NewSecurityIdentityProviderUpdateRequestWithDefaults() *SecurityIdentityProviderUpdateRequest`

NewSecurityIdentityProviderUpdateRequestWithDefaults instantiates a new SecurityIdentityProviderUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityIdentityProviderUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityIdentityProviderUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityIdentityProviderUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *SecurityIdentityProviderUpdateRequest) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SecurityIdentityProviderUpdateRequest) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SecurityIdentityProviderUpdateRequest) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SecurityIdentityProviderUpdateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SecurityIdentityProviderUpdateRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SecurityIdentityProviderUpdateRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *SecurityIdentityProviderUpdateRequest) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityIdentityProviderUpdateRequest) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityIdentityProviderUpdateRequest) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityIdentityProviderUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityIdentityProviderUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityIdentityProviderUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *SecurityIdentityProviderUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityIdentityProviderUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityIdentityProviderUpdateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *SecurityIdentityProviderUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityIdentityProviderUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityIdentityProviderUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledOnUI

`func (o *SecurityIdentityProviderUpdateRequest) GetEnabledOnUI() bool`

GetEnabledOnUI returns the EnabledOnUI field if non-nil, zero value otherwise.

### GetEnabledOnUIOk

`func (o *SecurityIdentityProviderUpdateRequest) GetEnabledOnUIOk() (*bool, bool)`

GetEnabledOnUIOk returns a tuple with the EnabledOnUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnUI

`func (o *SecurityIdentityProviderUpdateRequest) SetEnabledOnUI(v bool)`

SetEnabledOnUI sets EnabledOnUI field to given value.


### GetPasswordPolicy

`func (o *SecurityIdentityProviderUpdateRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *SecurityIdentityProviderUpdateRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *SecurityIdentityProviderUpdateRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *SecurityIdentityProviderUpdateRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *SecurityIdentityProviderUpdateRequest) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *SecurityIdentityProviderUpdateRequest) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetEmailTemplate

`func (o *SecurityIdentityProviderUpdateRequest) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *SecurityIdentityProviderUpdateRequest) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *SecurityIdentityProviderUpdateRequest) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *SecurityIdentityProviderUpdateRequest) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### SetEmailTemplateNil

`func (o *SecurityIdentityProviderUpdateRequest) SetEmailTemplateNil(b bool)`

 SetEmailTemplateNil sets the value for EmailTemplate to be an explicit nil

### UnsetEmailTemplate
`func (o *SecurityIdentityProviderUpdateRequest) UnsetEmailTemplate()`

UnsetEmailTemplate ensures that no value is present for EmailTemplate, not even an explicit nil
### GetProxy

`func (o *SecurityIdentityProviderUpdateRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SecurityIdentityProviderUpdateRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SecurityIdentityProviderUpdateRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SecurityIdentityProviderUpdateRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *SecurityIdentityProviderUpdateRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *SecurityIdentityProviderUpdateRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *SecurityIdentityProviderUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecurityIdentityProviderUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecurityIdentityProviderUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SecurityIdentityProviderUpdateRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *SecurityIdentityProviderUpdateRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SecurityIdentityProviderUpdateRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProviderMetadataUrl

`func (o *SecurityIdentityProviderUpdateRequest) GetProviderMetadataUrl() string`

GetProviderMetadataUrl returns the ProviderMetadataUrl field if non-nil, zero value otherwise.

### GetProviderMetadataUrlOk

`func (o *SecurityIdentityProviderUpdateRequest) GetProviderMetadataUrlOk() (*string, bool)`

GetProviderMetadataUrlOk returns a tuple with the ProviderMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadataUrl

`func (o *SecurityIdentityProviderUpdateRequest) SetProviderMetadataUrl(v string)`

SetProviderMetadataUrl sets ProviderMetadataUrl field to given value.


### GetClientCredentials

`func (o *SecurityIdentityProviderUpdateRequest) GetClientCredentials() string`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *SecurityIdentityProviderUpdateRequest) GetClientCredentialsOk() (*string, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *SecurityIdentityProviderUpdateRequest) SetClientCredentials(v string)`

SetClientCredentials sets ClientCredentials field to given value.


### GetScope

`func (o *SecurityIdentityProviderUpdateRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecurityIdentityProviderUpdateRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecurityIdentityProviderUpdateRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTrustSystemCAs

`func (o *SecurityIdentityProviderUpdateRequest) GetTrustSystemCAs() bool`

GetTrustSystemCAs returns the TrustSystemCAs field if non-nil, zero value otherwise.

### GetTrustSystemCAsOk

`func (o *SecurityIdentityProviderUpdateRequest) GetTrustSystemCAsOk() (*bool, bool)`

GetTrustSystemCAsOk returns a tuple with the TrustSystemCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustSystemCAs

`func (o *SecurityIdentityProviderUpdateRequest) SetTrustSystemCAs(v bool)`

SetTrustSystemCAs sets TrustSystemCAs field to given value.


### GetIdentifierClaim

`func (o *SecurityIdentityProviderUpdateRequest) GetIdentifierClaim() string`

GetIdentifierClaim returns the IdentifierClaim field if non-nil, zero value otherwise.

### GetIdentifierClaimOk

`func (o *SecurityIdentityProviderUpdateRequest) GetIdentifierClaimOk() (*string, bool)`

GetIdentifierClaimOk returns a tuple with the IdentifierClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierClaim

`func (o *SecurityIdentityProviderUpdateRequest) SetIdentifierClaim(v string)`

SetIdentifierClaim sets IdentifierClaim field to given value.


### GetEmailClaim

`func (o *SecurityIdentityProviderUpdateRequest) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *SecurityIdentityProviderUpdateRequest) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *SecurityIdentityProviderUpdateRequest) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.


### GetNameClaim

`func (o *SecurityIdentityProviderUpdateRequest) GetNameClaim() string`

GetNameClaim returns the NameClaim field if non-nil, zero value otherwise.

### GetNameClaimOk

`func (o *SecurityIdentityProviderUpdateRequest) GetNameClaimOk() (*string, bool)`

GetNameClaimOk returns a tuple with the NameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameClaim

`func (o *SecurityIdentityProviderUpdateRequest) SetNameClaim(v string)`

SetNameClaim sets NameClaim field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


