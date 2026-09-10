# ServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IatFutureRestriction** | Pointer to **string** | Maximum duration in the future the JWT &#x60;iat&#x60; claim is allowed to be. Must be set together with &#x60;iatPastRestriction&#x60;. | [optional] 
**IatPastRestriction** | Pointer to **string** | Maximum duration in the past the JWT &#x60;iat&#x60; claim is allowed to be. Must be set together with &#x60;iatFutureRestriction&#x60;. | [optional] 
**IdentifierMapping** | Pointer to **string** | Template string used to compute the identifier of the principal authenticated by this service account. | [optional] 
**JwtAllowedClockSkew** | Pointer to **string** | Allowed clock skew when validating JWT time-based claims. | [optional] 
**Name** | **string** | Internal name for the service account | 
**Permissions** | [**[]Permission**](Permission.md) | List of permissions to apply for successfully validated JWTs | 
**Roles** | **[]string** | List of roles to apply for successfully validated JWTs | 
**TrustConfig** | [**ServiceAccountTrustConfig**](ServiceAccountTrustConfig.md) |  | 
**ValidationRules** | **[]string** | List of rules to apply on top of signature verification for the incoming JWT to be trusted | 

## Methods

### NewServiceAccount

`func NewServiceAccount(name string, permissions []Permission, roles []string, trustConfig ServiceAccountTrustConfig, validationRules []string, ) *ServiceAccount`

NewServiceAccount instantiates a new ServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountWithDefaults

`func NewServiceAccountWithDefaults() *ServiceAccount`

NewServiceAccountWithDefaults instantiates a new ServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIatFutureRestriction

`func (o *ServiceAccount) GetIatFutureRestriction() string`

GetIatFutureRestriction returns the IatFutureRestriction field if non-nil, zero value otherwise.

### GetIatFutureRestrictionOk

`func (o *ServiceAccount) GetIatFutureRestrictionOk() (*string, bool)`

GetIatFutureRestrictionOk returns a tuple with the IatFutureRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIatFutureRestriction

`func (o *ServiceAccount) SetIatFutureRestriction(v string)`

SetIatFutureRestriction sets IatFutureRestriction field to given value.

### HasIatFutureRestriction

`func (o *ServiceAccount) HasIatFutureRestriction() bool`

HasIatFutureRestriction returns a boolean if a field has been set.

### GetIatPastRestriction

`func (o *ServiceAccount) GetIatPastRestriction() string`

GetIatPastRestriction returns the IatPastRestriction field if non-nil, zero value otherwise.

### GetIatPastRestrictionOk

`func (o *ServiceAccount) GetIatPastRestrictionOk() (*string, bool)`

GetIatPastRestrictionOk returns a tuple with the IatPastRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIatPastRestriction

`func (o *ServiceAccount) SetIatPastRestriction(v string)`

SetIatPastRestriction sets IatPastRestriction field to given value.

### HasIatPastRestriction

`func (o *ServiceAccount) HasIatPastRestriction() bool`

HasIatPastRestriction returns a boolean if a field has been set.

### GetIdentifierMapping

`func (o *ServiceAccount) GetIdentifierMapping() string`

GetIdentifierMapping returns the IdentifierMapping field if non-nil, zero value otherwise.

### GetIdentifierMappingOk

`func (o *ServiceAccount) GetIdentifierMappingOk() (*string, bool)`

GetIdentifierMappingOk returns a tuple with the IdentifierMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierMapping

`func (o *ServiceAccount) SetIdentifierMapping(v string)`

SetIdentifierMapping sets IdentifierMapping field to given value.

### HasIdentifierMapping

`func (o *ServiceAccount) HasIdentifierMapping() bool`

HasIdentifierMapping returns a boolean if a field has been set.

### GetJwtAllowedClockSkew

`func (o *ServiceAccount) GetJwtAllowedClockSkew() string`

GetJwtAllowedClockSkew returns the JwtAllowedClockSkew field if non-nil, zero value otherwise.

### GetJwtAllowedClockSkewOk

`func (o *ServiceAccount) GetJwtAllowedClockSkewOk() (*string, bool)`

GetJwtAllowedClockSkewOk returns a tuple with the JwtAllowedClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAllowedClockSkew

`func (o *ServiceAccount) SetJwtAllowedClockSkew(v string)`

SetJwtAllowedClockSkew sets JwtAllowedClockSkew field to given value.

### HasJwtAllowedClockSkew

`func (o *ServiceAccount) HasJwtAllowedClockSkew() bool`

HasJwtAllowedClockSkew returns a boolean if a field has been set.

### GetName

`func (o *ServiceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccount) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *ServiceAccount) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ServiceAccount) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ServiceAccount) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.


### GetRoles

`func (o *ServiceAccount) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ServiceAccount) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ServiceAccount) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetTrustConfig

`func (o *ServiceAccount) GetTrustConfig() ServiceAccountTrustConfig`

GetTrustConfig returns the TrustConfig field if non-nil, zero value otherwise.

### GetTrustConfigOk

`func (o *ServiceAccount) GetTrustConfigOk() (*ServiceAccountTrustConfig, bool)`

GetTrustConfigOk returns a tuple with the TrustConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustConfig

`func (o *ServiceAccount) SetTrustConfig(v ServiceAccountTrustConfig)`

SetTrustConfig sets TrustConfig field to given value.


### GetValidationRules

`func (o *ServiceAccount) GetValidationRules() []string`

GetValidationRules returns the ValidationRules field if non-nil, zero value otherwise.

### GetValidationRulesOk

`func (o *ServiceAccount) GetValidationRulesOk() (*[]string, bool)`

GetValidationRulesOk returns a tuple with the ValidationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRules

`func (o *ServiceAccount) SetValidationRules(v []string)`

SetValidationRules sets ValidationRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


