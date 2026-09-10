# ServiceAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Object internal ID | [optional] 
**Readonly** | **bool** | If true, this object was externally provisioned and cannot be edited | 
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

### NewServiceAccountResponse

`func NewServiceAccountResponse(readonly bool, name string, permissions []Permission, roles []string, trustConfig ServiceAccountTrustConfig, validationRules []string, ) *ServiceAccountResponse`

NewServiceAccountResponse instantiates a new ServiceAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountResponseWithDefaults

`func NewServiceAccountResponseWithDefaults() *ServiceAccountResponse`

NewServiceAccountResponseWithDefaults instantiates a new ServiceAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReadonly

`func (o *ServiceAccountResponse) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *ServiceAccountResponse) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *ServiceAccountResponse) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.


### GetIatFutureRestriction

`func (o *ServiceAccountResponse) GetIatFutureRestriction() string`

GetIatFutureRestriction returns the IatFutureRestriction field if non-nil, zero value otherwise.

### GetIatFutureRestrictionOk

`func (o *ServiceAccountResponse) GetIatFutureRestrictionOk() (*string, bool)`

GetIatFutureRestrictionOk returns a tuple with the IatFutureRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIatFutureRestriction

`func (o *ServiceAccountResponse) SetIatFutureRestriction(v string)`

SetIatFutureRestriction sets IatFutureRestriction field to given value.

### HasIatFutureRestriction

`func (o *ServiceAccountResponse) HasIatFutureRestriction() bool`

HasIatFutureRestriction returns a boolean if a field has been set.

### GetIatPastRestriction

`func (o *ServiceAccountResponse) GetIatPastRestriction() string`

GetIatPastRestriction returns the IatPastRestriction field if non-nil, zero value otherwise.

### GetIatPastRestrictionOk

`func (o *ServiceAccountResponse) GetIatPastRestrictionOk() (*string, bool)`

GetIatPastRestrictionOk returns a tuple with the IatPastRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIatPastRestriction

`func (o *ServiceAccountResponse) SetIatPastRestriction(v string)`

SetIatPastRestriction sets IatPastRestriction field to given value.

### HasIatPastRestriction

`func (o *ServiceAccountResponse) HasIatPastRestriction() bool`

HasIatPastRestriction returns a boolean if a field has been set.

### GetIdentifierMapping

`func (o *ServiceAccountResponse) GetIdentifierMapping() string`

GetIdentifierMapping returns the IdentifierMapping field if non-nil, zero value otherwise.

### GetIdentifierMappingOk

`func (o *ServiceAccountResponse) GetIdentifierMappingOk() (*string, bool)`

GetIdentifierMappingOk returns a tuple with the IdentifierMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierMapping

`func (o *ServiceAccountResponse) SetIdentifierMapping(v string)`

SetIdentifierMapping sets IdentifierMapping field to given value.

### HasIdentifierMapping

`func (o *ServiceAccountResponse) HasIdentifierMapping() bool`

HasIdentifierMapping returns a boolean if a field has been set.

### GetJwtAllowedClockSkew

`func (o *ServiceAccountResponse) GetJwtAllowedClockSkew() string`

GetJwtAllowedClockSkew returns the JwtAllowedClockSkew field if non-nil, zero value otherwise.

### GetJwtAllowedClockSkewOk

`func (o *ServiceAccountResponse) GetJwtAllowedClockSkewOk() (*string, bool)`

GetJwtAllowedClockSkewOk returns a tuple with the JwtAllowedClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAllowedClockSkew

`func (o *ServiceAccountResponse) SetJwtAllowedClockSkew(v string)`

SetJwtAllowedClockSkew sets JwtAllowedClockSkew field to given value.

### HasJwtAllowedClockSkew

`func (o *ServiceAccountResponse) HasJwtAllowedClockSkew() bool`

HasJwtAllowedClockSkew returns a boolean if a field has been set.

### GetName

`func (o *ServiceAccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *ServiceAccountResponse) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ServiceAccountResponse) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ServiceAccountResponse) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.


### GetRoles

`func (o *ServiceAccountResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ServiceAccountResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ServiceAccountResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetTrustConfig

`func (o *ServiceAccountResponse) GetTrustConfig() ServiceAccountTrustConfig`

GetTrustConfig returns the TrustConfig field if non-nil, zero value otherwise.

### GetTrustConfigOk

`func (o *ServiceAccountResponse) GetTrustConfigOk() (*ServiceAccountTrustConfig, bool)`

GetTrustConfigOk returns a tuple with the TrustConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustConfig

`func (o *ServiceAccountResponse) SetTrustConfig(v ServiceAccountTrustConfig)`

SetTrustConfig sets TrustConfig field to given value.


### GetValidationRules

`func (o *ServiceAccountResponse) GetValidationRules() []string`

GetValidationRules returns the ValidationRules field if non-nil, zero value otherwise.

### GetValidationRulesOk

`func (o *ServiceAccountResponse) GetValidationRulesOk() (*[]string, bool)`

GetValidationRulesOk returns a tuple with the ValidationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRules

`func (o *ServiceAccountResponse) SetValidationRules(v []string)`

SetValidationRules sets ValidationRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


