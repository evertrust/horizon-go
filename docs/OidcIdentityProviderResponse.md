# OidcIdentityProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The internal ID of the Identity Provider | 
**ClientCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the client ID  and secret to use to authenticate Horizon against the identity provider | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the identity provider | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the identity provider | [optional] 
**EmailClaim** | **string** | The OpenID information that will be used as the user&#39;s email in Horizon | 
**Enabled** | **bool** | Whether the identity provider can be used to identify against Horizon | 
**EnabledOnUI** | **bool** | Whether the identity provider can be selected on login to the Horizon UI | 
**IdentifierClaim** | **string** | The OpenID information that will be used as the user&#39;s identifier in Horizon | 
**Name** | **string** | The internal name of the identity provider | 
**NameClaim** | **string** | The OpenID information that will be used as the user&#39;s name in Horizon | 
**ProviderMetadataUrl** | **string** | The URL of the identity provider OpenID callback | 
**Proxy** | Pointer to **NullableString** | The name of the proxy to use to reach the identity provider | [optional] 
**Scope** | **string** | The scope where to retrieve the user data from | 
**Timeout** | Pointer to **NullableString** | The timeout value to use when connecting to the identity provider (must be a valid finite duration) | [optional] 
**TrustSystemCAs** | **bool** | Trust AC coming from the system trust store or only trust AC imported in Horizon | [default to true]
**Type** | **string** | The type of Identity provider to register | 

## Methods

### NewOidcIdentityProviderResponse

`func NewOidcIdentityProviderResponse(id string, clientCredentials string, emailClaim string, enabled bool, enabledOnUI bool, identifierClaim string, name string, nameClaim string, providerMetadataUrl string, scope string, trustSystemCAs bool, type_ string, ) *OidcIdentityProviderResponse`

NewOidcIdentityProviderResponse instantiates a new OidcIdentityProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcIdentityProviderResponseWithDefaults

`func NewOidcIdentityProviderResponseWithDefaults() *OidcIdentityProviderResponse`

NewOidcIdentityProviderResponseWithDefaults instantiates a new OidcIdentityProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OidcIdentityProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OidcIdentityProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OidcIdentityProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetClientCredentials

`func (o *OidcIdentityProviderResponse) GetClientCredentials() string`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *OidcIdentityProviderResponse) GetClientCredentialsOk() (*string, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *OidcIdentityProviderResponse) SetClientCredentials(v string)`

SetClientCredentials sets ClientCredentials field to given value.


### GetDescription

`func (o *OidcIdentityProviderResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OidcIdentityProviderResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OidcIdentityProviderResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OidcIdentityProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OidcIdentityProviderResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OidcIdentityProviderResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *OidcIdentityProviderResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OidcIdentityProviderResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OidcIdentityProviderResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OidcIdentityProviderResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *OidcIdentityProviderResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *OidcIdentityProviderResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmailClaim

`func (o *OidcIdentityProviderResponse) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *OidcIdentityProviderResponse) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *OidcIdentityProviderResponse) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.


### GetEnabled

`func (o *OidcIdentityProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OidcIdentityProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OidcIdentityProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledOnUI

`func (o *OidcIdentityProviderResponse) GetEnabledOnUI() bool`

GetEnabledOnUI returns the EnabledOnUI field if non-nil, zero value otherwise.

### GetEnabledOnUIOk

`func (o *OidcIdentityProviderResponse) GetEnabledOnUIOk() (*bool, bool)`

GetEnabledOnUIOk returns a tuple with the EnabledOnUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnUI

`func (o *OidcIdentityProviderResponse) SetEnabledOnUI(v bool)`

SetEnabledOnUI sets EnabledOnUI field to given value.


### GetIdentifierClaim

`func (o *OidcIdentityProviderResponse) GetIdentifierClaim() string`

GetIdentifierClaim returns the IdentifierClaim field if non-nil, zero value otherwise.

### GetIdentifierClaimOk

`func (o *OidcIdentityProviderResponse) GetIdentifierClaimOk() (*string, bool)`

GetIdentifierClaimOk returns a tuple with the IdentifierClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierClaim

`func (o *OidcIdentityProviderResponse) SetIdentifierClaim(v string)`

SetIdentifierClaim sets IdentifierClaim field to given value.


### GetName

`func (o *OidcIdentityProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OidcIdentityProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OidcIdentityProviderResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameClaim

`func (o *OidcIdentityProviderResponse) GetNameClaim() string`

GetNameClaim returns the NameClaim field if non-nil, zero value otherwise.

### GetNameClaimOk

`func (o *OidcIdentityProviderResponse) GetNameClaimOk() (*string, bool)`

GetNameClaimOk returns a tuple with the NameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameClaim

`func (o *OidcIdentityProviderResponse) SetNameClaim(v string)`

SetNameClaim sets NameClaim field to given value.


### GetProviderMetadataUrl

`func (o *OidcIdentityProviderResponse) GetProviderMetadataUrl() string`

GetProviderMetadataUrl returns the ProviderMetadataUrl field if non-nil, zero value otherwise.

### GetProviderMetadataUrlOk

`func (o *OidcIdentityProviderResponse) GetProviderMetadataUrlOk() (*string, bool)`

GetProviderMetadataUrlOk returns a tuple with the ProviderMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadataUrl

`func (o *OidcIdentityProviderResponse) SetProviderMetadataUrl(v string)`

SetProviderMetadataUrl sets ProviderMetadataUrl field to given value.


### GetProxy

`func (o *OidcIdentityProviderResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *OidcIdentityProviderResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *OidcIdentityProviderResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *OidcIdentityProviderResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *OidcIdentityProviderResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *OidcIdentityProviderResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetScope

`func (o *OidcIdentityProviderResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OidcIdentityProviderResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OidcIdentityProviderResponse) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTimeout

`func (o *OidcIdentityProviderResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *OidcIdentityProviderResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *OidcIdentityProviderResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *OidcIdentityProviderResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *OidcIdentityProviderResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *OidcIdentityProviderResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTrustSystemCAs

`func (o *OidcIdentityProviderResponse) GetTrustSystemCAs() bool`

GetTrustSystemCAs returns the TrustSystemCAs field if non-nil, zero value otherwise.

### GetTrustSystemCAsOk

`func (o *OidcIdentityProviderResponse) GetTrustSystemCAsOk() (*bool, bool)`

GetTrustSystemCAsOk returns a tuple with the TrustSystemCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustSystemCAs

`func (o *OidcIdentityProviderResponse) SetTrustSystemCAs(v bool)`

SetTrustSystemCAs sets TrustSystemCAs field to given value.


### GetType

`func (o *OidcIdentityProviderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OidcIdentityProviderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OidcIdentityProviderResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


