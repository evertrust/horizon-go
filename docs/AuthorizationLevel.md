# AuthorizationLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | **string** | The access level required to perform the action | 
**EnforcedIdentityProviders** | Pointer to [**[]EnforcedIdentityProvider**](EnforcedIdentityProvider.md) | The different identity providers that can be enforced to perform the action | [optional] 

## Methods

### NewAuthorizationLevel

`func NewAuthorizationLevel(accessLevel string, ) *AuthorizationLevel`

NewAuthorizationLevel instantiates a new AuthorizationLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationLevelWithDefaults

`func NewAuthorizationLevelWithDefaults() *AuthorizationLevel`

NewAuthorizationLevelWithDefaults instantiates a new AuthorizationLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *AuthorizationLevel) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *AuthorizationLevel) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *AuthorizationLevel) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.


### GetEnforcedIdentityProviders

`func (o *AuthorizationLevel) GetEnforcedIdentityProviders() []EnforcedIdentityProvider`

GetEnforcedIdentityProviders returns the EnforcedIdentityProviders field if non-nil, zero value otherwise.

### GetEnforcedIdentityProvidersOk

`func (o *AuthorizationLevel) GetEnforcedIdentityProvidersOk() (*[]EnforcedIdentityProvider, bool)`

GetEnforcedIdentityProvidersOk returns a tuple with the EnforcedIdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcedIdentityProviders

`func (o *AuthorizationLevel) SetEnforcedIdentityProviders(v []EnforcedIdentityProvider)`

SetEnforcedIdentityProviders sets EnforcedIdentityProviders field to given value.

### HasEnforcedIdentityProviders

`func (o *AuthorizationLevel) HasEnforcedIdentityProviders() bool`

HasEnforcedIdentityProviders returns a boolean if a field has been set.

### SetEnforcedIdentityProvidersNil

`func (o *AuthorizationLevel) SetEnforcedIdentityProvidersNil(b bool)`

 SetEnforcedIdentityProvidersNil sets the value for EnforcedIdentityProviders to be an explicit nil

### UnsetEnforcedIdentityProviders
`func (o *AuthorizationLevel) UnsetEnforcedIdentityProviders()`

UnsetEnforcedIdentityProviders ensures that no value is present for EnforcedIdentityProviders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


