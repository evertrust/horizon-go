# OidcIdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the client ID  and secret to use to authenticate Horizon against the identity provider | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the identity provider | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the identity provider | [optional] 
**EmailClaim** | Pointer to **string** | The OpenID information that will be used as the user&#39;s email in Horizon | [optional] [default to "{{email}}"]
**Enabled** | **bool** | Whether the identity provider can be used to identify against Horizon | 
**EnabledOnUI** | **bool** | Whether the identity provider can be selected on login to the Horizon UI | 
**IdentifierClaim** | Pointer to **string** | The OpenID information that will be used as the user&#39;s identifier in Horizon | [optional] [default to "{{email}}"]
**Mapping** | Pointer to [**OidcIdentityProviderMapping**](OidcIdentityProviderMapping.md) |  | [optional] 
**Name** | **string** | The internal name of the identity provider | 
**NameClaim** | Pointer to **string** | The OpenID information that will be used as the user&#39;s name in Horizon | [optional] [default to "{{name}}"]
**ProviderMetadataUrl** | **string** | The URL of the identity provider OpenID callback | 
**Proxy** | Pointer to **NullableString** | The name of the proxy to use to reach the identity provider | [optional] 
**Scope** | **string** | The scope where to retrieve the user data from | 
**Timeout** | **string** | The timeout value to use when connecting to the identity provider (must be a valid finite duration) | 
**TrustSystemCAs** | **bool** | Trust AC coming from the system trust store or only trust AC imported in Horizon | [default to true]
**Type** | **string** | The type of Identity provider to register | 

## Methods

### NewOidcIdentityProvider

`func NewOidcIdentityProvider(clientCredentials string, enabled bool, enabledOnUI bool, name string, providerMetadataUrl string, scope string, timeout string, trustSystemCAs bool, type_ string, ) *OidcIdentityProvider`

NewOidcIdentityProvider instantiates a new OidcIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcIdentityProviderWithDefaults

`func NewOidcIdentityProviderWithDefaults() *OidcIdentityProvider`

NewOidcIdentityProviderWithDefaults instantiates a new OidcIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCredentials

`func (o *OidcIdentityProvider) GetClientCredentials() string`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *OidcIdentityProvider) GetClientCredentialsOk() (*string, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *OidcIdentityProvider) SetClientCredentials(v string)`

SetClientCredentials sets ClientCredentials field to given value.


### GetDescription

`func (o *OidcIdentityProvider) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OidcIdentityProvider) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OidcIdentityProvider) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OidcIdentityProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OidcIdentityProvider) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OidcIdentityProvider) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *OidcIdentityProvider) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OidcIdentityProvider) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OidcIdentityProvider) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OidcIdentityProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *OidcIdentityProvider) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *OidcIdentityProvider) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmailClaim

`func (o *OidcIdentityProvider) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *OidcIdentityProvider) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *OidcIdentityProvider) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.

### HasEmailClaim

`func (o *OidcIdentityProvider) HasEmailClaim() bool`

HasEmailClaim returns a boolean if a field has been set.

### GetEnabled

`func (o *OidcIdentityProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OidcIdentityProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OidcIdentityProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledOnUI

`func (o *OidcIdentityProvider) GetEnabledOnUI() bool`

GetEnabledOnUI returns the EnabledOnUI field if non-nil, zero value otherwise.

### GetEnabledOnUIOk

`func (o *OidcIdentityProvider) GetEnabledOnUIOk() (*bool, bool)`

GetEnabledOnUIOk returns a tuple with the EnabledOnUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnUI

`func (o *OidcIdentityProvider) SetEnabledOnUI(v bool)`

SetEnabledOnUI sets EnabledOnUI field to given value.


### GetIdentifierClaim

`func (o *OidcIdentityProvider) GetIdentifierClaim() string`

GetIdentifierClaim returns the IdentifierClaim field if non-nil, zero value otherwise.

### GetIdentifierClaimOk

`func (o *OidcIdentityProvider) GetIdentifierClaimOk() (*string, bool)`

GetIdentifierClaimOk returns a tuple with the IdentifierClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierClaim

`func (o *OidcIdentityProvider) SetIdentifierClaim(v string)`

SetIdentifierClaim sets IdentifierClaim field to given value.

### HasIdentifierClaim

`func (o *OidcIdentityProvider) HasIdentifierClaim() bool`

HasIdentifierClaim returns a boolean if a field has been set.

### GetMapping

`func (o *OidcIdentityProvider) GetMapping() OidcIdentityProviderMapping`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *OidcIdentityProvider) GetMappingOk() (*OidcIdentityProviderMapping, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *OidcIdentityProvider) SetMapping(v OidcIdentityProviderMapping)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *OidcIdentityProvider) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetName

`func (o *OidcIdentityProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OidcIdentityProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OidcIdentityProvider) SetName(v string)`

SetName sets Name field to given value.


### GetNameClaim

`func (o *OidcIdentityProvider) GetNameClaim() string`

GetNameClaim returns the NameClaim field if non-nil, zero value otherwise.

### GetNameClaimOk

`func (o *OidcIdentityProvider) GetNameClaimOk() (*string, bool)`

GetNameClaimOk returns a tuple with the NameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameClaim

`func (o *OidcIdentityProvider) SetNameClaim(v string)`

SetNameClaim sets NameClaim field to given value.

### HasNameClaim

`func (o *OidcIdentityProvider) HasNameClaim() bool`

HasNameClaim returns a boolean if a field has been set.

### GetProviderMetadataUrl

`func (o *OidcIdentityProvider) GetProviderMetadataUrl() string`

GetProviderMetadataUrl returns the ProviderMetadataUrl field if non-nil, zero value otherwise.

### GetProviderMetadataUrlOk

`func (o *OidcIdentityProvider) GetProviderMetadataUrlOk() (*string, bool)`

GetProviderMetadataUrlOk returns a tuple with the ProviderMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadataUrl

`func (o *OidcIdentityProvider) SetProviderMetadataUrl(v string)`

SetProviderMetadataUrl sets ProviderMetadataUrl field to given value.


### GetProxy

`func (o *OidcIdentityProvider) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *OidcIdentityProvider) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *OidcIdentityProvider) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *OidcIdentityProvider) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *OidcIdentityProvider) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *OidcIdentityProvider) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetScope

`func (o *OidcIdentityProvider) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OidcIdentityProvider) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OidcIdentityProvider) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTimeout

`func (o *OidcIdentityProvider) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *OidcIdentityProvider) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *OidcIdentityProvider) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTrustSystemCAs

`func (o *OidcIdentityProvider) GetTrustSystemCAs() bool`

GetTrustSystemCAs returns the TrustSystemCAs field if non-nil, zero value otherwise.

### GetTrustSystemCAsOk

`func (o *OidcIdentityProvider) GetTrustSystemCAsOk() (*bool, bool)`

GetTrustSystemCAsOk returns a tuple with the TrustSystemCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustSystemCAs

`func (o *OidcIdentityProvider) SetTrustSystemCAs(v bool)`

SetTrustSystemCAs sets TrustSystemCAs field to given value.


### GetType

`func (o *OidcIdentityProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OidcIdentityProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OidcIdentityProvider) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


