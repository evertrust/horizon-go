# SecurityIdentityProviderAdd201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The internal ID of the Identity Provider | 
**Name** | **string** | The internal name of the identity provider | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the identity provider | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the identity provider | [optional] 
**Type** | **string** | The type of Identity provider to register | 
**Enabled** | **bool** | Whether the identity provider can be used to identify against Horizon | 
**EnabledOnUI** | **bool** | Whether the identity provider can be selected on login to the Horizon UI | 
**Proxy** | Pointer to **NullableString** | The name of the proxy to use to reach the identity provider | [optional] 
**Timeout** | Pointer to **NullableString** | The timeout value to use when connecting to the identity provider (must be a valid finite duration) | [optional] 
**ProviderMetadataUrl** | **string** | The URL of the identity provider OpenID callback | 
**ClientCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the client ID  and secret to use to authenticate Horizon against the identity provider | 
**Scope** | **string** | The scope where to retrieve the user data from | 
**TrustSystemCAs** | **bool** | Trust AC coming from the system trust store or only trust AC imported in Horizon | [default to true]
**IdentifierClaim** | **string** | The OpenID information that will be used as the user&#39;s identifier in Horizon | 
**EmailClaim** | **string** | The OpenID information that will be used as the user&#39;s email in Horizon | 
**NameClaim** | **string** | The OpenID information that will be used as the user&#39;s name in Horizon | 

## Methods

### NewSecurityIdentityProviderAdd201Response

`func NewSecurityIdentityProviderAdd201Response(id string, name string, type_ string, enabled bool, enabledOnUI bool, providerMetadataUrl string, clientCredentials string, scope string, trustSystemCAs bool, identifierClaim string, emailClaim string, nameClaim string, ) *SecurityIdentityProviderAdd201Response`

NewSecurityIdentityProviderAdd201Response instantiates a new SecurityIdentityProviderAdd201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIdentityProviderAdd201ResponseWithDefaults

`func NewSecurityIdentityProviderAdd201ResponseWithDefaults() *SecurityIdentityProviderAdd201Response`

NewSecurityIdentityProviderAdd201ResponseWithDefaults instantiates a new SecurityIdentityProviderAdd201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityIdentityProviderAdd201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityIdentityProviderAdd201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityIdentityProviderAdd201Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityIdentityProviderAdd201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityIdentityProviderAdd201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityIdentityProviderAdd201Response) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *SecurityIdentityProviderAdd201Response) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SecurityIdentityProviderAdd201Response) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SecurityIdentityProviderAdd201Response) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SecurityIdentityProviderAdd201Response) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SecurityIdentityProviderAdd201Response) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SecurityIdentityProviderAdd201Response) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *SecurityIdentityProviderAdd201Response) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityIdentityProviderAdd201Response) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityIdentityProviderAdd201Response) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityIdentityProviderAdd201Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityIdentityProviderAdd201Response) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityIdentityProviderAdd201Response) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *SecurityIdentityProviderAdd201Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityIdentityProviderAdd201Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityIdentityProviderAdd201Response) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *SecurityIdentityProviderAdd201Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityIdentityProviderAdd201Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityIdentityProviderAdd201Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledOnUI

`func (o *SecurityIdentityProviderAdd201Response) GetEnabledOnUI() bool`

GetEnabledOnUI returns the EnabledOnUI field if non-nil, zero value otherwise.

### GetEnabledOnUIOk

`func (o *SecurityIdentityProviderAdd201Response) GetEnabledOnUIOk() (*bool, bool)`

GetEnabledOnUIOk returns a tuple with the EnabledOnUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnUI

`func (o *SecurityIdentityProviderAdd201Response) SetEnabledOnUI(v bool)`

SetEnabledOnUI sets EnabledOnUI field to given value.


### GetProxy

`func (o *SecurityIdentityProviderAdd201Response) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SecurityIdentityProviderAdd201Response) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SecurityIdentityProviderAdd201Response) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SecurityIdentityProviderAdd201Response) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *SecurityIdentityProviderAdd201Response) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *SecurityIdentityProviderAdd201Response) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *SecurityIdentityProviderAdd201Response) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecurityIdentityProviderAdd201Response) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecurityIdentityProviderAdd201Response) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SecurityIdentityProviderAdd201Response) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *SecurityIdentityProviderAdd201Response) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SecurityIdentityProviderAdd201Response) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProviderMetadataUrl

`func (o *SecurityIdentityProviderAdd201Response) GetProviderMetadataUrl() string`

GetProviderMetadataUrl returns the ProviderMetadataUrl field if non-nil, zero value otherwise.

### GetProviderMetadataUrlOk

`func (o *SecurityIdentityProviderAdd201Response) GetProviderMetadataUrlOk() (*string, bool)`

GetProviderMetadataUrlOk returns a tuple with the ProviderMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadataUrl

`func (o *SecurityIdentityProviderAdd201Response) SetProviderMetadataUrl(v string)`

SetProviderMetadataUrl sets ProviderMetadataUrl field to given value.


### GetClientCredentials

`func (o *SecurityIdentityProviderAdd201Response) GetClientCredentials() string`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *SecurityIdentityProviderAdd201Response) GetClientCredentialsOk() (*string, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *SecurityIdentityProviderAdd201Response) SetClientCredentials(v string)`

SetClientCredentials sets ClientCredentials field to given value.


### GetScope

`func (o *SecurityIdentityProviderAdd201Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecurityIdentityProviderAdd201Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecurityIdentityProviderAdd201Response) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTrustSystemCAs

`func (o *SecurityIdentityProviderAdd201Response) GetTrustSystemCAs() bool`

GetTrustSystemCAs returns the TrustSystemCAs field if non-nil, zero value otherwise.

### GetTrustSystemCAsOk

`func (o *SecurityIdentityProviderAdd201Response) GetTrustSystemCAsOk() (*bool, bool)`

GetTrustSystemCAsOk returns a tuple with the TrustSystemCAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustSystemCAs

`func (o *SecurityIdentityProviderAdd201Response) SetTrustSystemCAs(v bool)`

SetTrustSystemCAs sets TrustSystemCAs field to given value.


### GetIdentifierClaim

`func (o *SecurityIdentityProviderAdd201Response) GetIdentifierClaim() string`

GetIdentifierClaim returns the IdentifierClaim field if non-nil, zero value otherwise.

### GetIdentifierClaimOk

`func (o *SecurityIdentityProviderAdd201Response) GetIdentifierClaimOk() (*string, bool)`

GetIdentifierClaimOk returns a tuple with the IdentifierClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierClaim

`func (o *SecurityIdentityProviderAdd201Response) SetIdentifierClaim(v string)`

SetIdentifierClaim sets IdentifierClaim field to given value.


### GetEmailClaim

`func (o *SecurityIdentityProviderAdd201Response) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *SecurityIdentityProviderAdd201Response) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *SecurityIdentityProviderAdd201Response) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.


### GetNameClaim

`func (o *SecurityIdentityProviderAdd201Response) GetNameClaim() string`

GetNameClaim returns the NameClaim field if non-nil, zero value otherwise.

### GetNameClaimOk

`func (o *SecurityIdentityProviderAdd201Response) GetNameClaimOk() (*string, bool)`

GetNameClaimOk returns a tuple with the NameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameClaim

`func (o *SecurityIdentityProviderAdd201Response) SetNameClaim(v string)`

SetNameClaim sets NameClaim field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


