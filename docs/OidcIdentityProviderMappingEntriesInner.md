# OidcIdentityProviderMappingEntriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claim** | Pointer to **string** | JWT Claim value that will provide the corresponding teams and roles | [optional] 
**Roles** | Pointer to **[]string** | Name of the roles to map to this claim value | [optional] 
**Teams** | Pointer to **[]string** | Name of the teams to map to this claim value | [optional] 

## Methods

### NewOidcIdentityProviderMappingEntriesInner

`func NewOidcIdentityProviderMappingEntriesInner() *OidcIdentityProviderMappingEntriesInner`

NewOidcIdentityProviderMappingEntriesInner instantiates a new OidcIdentityProviderMappingEntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcIdentityProviderMappingEntriesInnerWithDefaults

`func NewOidcIdentityProviderMappingEntriesInnerWithDefaults() *OidcIdentityProviderMappingEntriesInner`

NewOidcIdentityProviderMappingEntriesInnerWithDefaults instantiates a new OidcIdentityProviderMappingEntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaim

`func (o *OidcIdentityProviderMappingEntriesInner) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *OidcIdentityProviderMappingEntriesInner) GetClaimOk() (*string, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *OidcIdentityProviderMappingEntriesInner) SetClaim(v string)`

SetClaim sets Claim field to given value.

### HasClaim

`func (o *OidcIdentityProviderMappingEntriesInner) HasClaim() bool`

HasClaim returns a boolean if a field has been set.

### GetRoles

`func (o *OidcIdentityProviderMappingEntriesInner) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OidcIdentityProviderMappingEntriesInner) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OidcIdentityProviderMappingEntriesInner) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OidcIdentityProviderMappingEntriesInner) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTeams

`func (o *OidcIdentityProviderMappingEntriesInner) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *OidcIdentityProviderMappingEntriesInner) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *OidcIdentityProviderMappingEntriesInner) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *OidcIdentityProviderMappingEntriesInner) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


