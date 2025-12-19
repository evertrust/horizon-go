# PrincipalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDashboards** | Pointer to [**[]Dashboard**](Dashboard.md) | The custom dashboards of the principal | [optional] 
**Identity** | [**Identity**](Identity.md) |  | 
**Permissions** | Pointer to [**[]Permission**](Permission.md) | The permissions of the principal | [optional] 
**Preferences** | Pointer to [**NullablePrincipalInfoPreferences**](PrincipalInfoPreferences.md) | The UI preferences of the principal | [optional] 
**Roles** | Pointer to **[]string** | The roles of the principal | [optional] 
**Teams** | Pointer to [**[]PrincipalResponseTeamsInner**](PrincipalResponseTeamsInner.md) | The teams of the principal | [optional] 

## Methods

### NewPrincipalResponse

`func NewPrincipalResponse(identity Identity, ) *PrincipalResponse`

NewPrincipalResponse instantiates a new PrincipalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalResponseWithDefaults

`func NewPrincipalResponseWithDefaults() *PrincipalResponse`

NewPrincipalResponseWithDefaults instantiates a new PrincipalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDashboards

`func (o *PrincipalResponse) GetCustomDashboards() []Dashboard`

GetCustomDashboards returns the CustomDashboards field if non-nil, zero value otherwise.

### GetCustomDashboardsOk

`func (o *PrincipalResponse) GetCustomDashboardsOk() (*[]Dashboard, bool)`

GetCustomDashboardsOk returns a tuple with the CustomDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDashboards

`func (o *PrincipalResponse) SetCustomDashboards(v []Dashboard)`

SetCustomDashboards sets CustomDashboards field to given value.

### HasCustomDashboards

`func (o *PrincipalResponse) HasCustomDashboards() bool`

HasCustomDashboards returns a boolean if a field has been set.

### SetCustomDashboardsNil

`func (o *PrincipalResponse) SetCustomDashboardsNil(b bool)`

 SetCustomDashboardsNil sets the value for CustomDashboards to be an explicit nil

### UnsetCustomDashboards
`func (o *PrincipalResponse) UnsetCustomDashboards()`

UnsetCustomDashboards ensures that no value is present for CustomDashboards, not even an explicit nil
### GetIdentity

`func (o *PrincipalResponse) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *PrincipalResponse) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *PrincipalResponse) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.


### GetPermissions

`func (o *PrincipalResponse) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PrincipalResponse) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PrincipalResponse) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PrincipalResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *PrincipalResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *PrincipalResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPreferences

`func (o *PrincipalResponse) GetPreferences() PrincipalInfoPreferences`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *PrincipalResponse) GetPreferencesOk() (*PrincipalInfoPreferences, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *PrincipalResponse) SetPreferences(v PrincipalInfoPreferences)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *PrincipalResponse) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### SetPreferencesNil

`func (o *PrincipalResponse) SetPreferencesNil(b bool)`

 SetPreferencesNil sets the value for Preferences to be an explicit nil

### UnsetPreferences
`func (o *PrincipalResponse) UnsetPreferences()`

UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
### GetRoles

`func (o *PrincipalResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PrincipalResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PrincipalResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PrincipalResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *PrincipalResponse) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *PrincipalResponse) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetTeams

`func (o *PrincipalResponse) GetTeams() []PrincipalResponseTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *PrincipalResponse) GetTeamsOk() (*[]PrincipalResponseTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *PrincipalResponse) SetTeams(v []PrincipalResponseTeamsInner)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *PrincipalResponse) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### SetTeamsNil

`func (o *PrincipalResponse) SetTeamsNil(b bool)`

 SetTeamsNil sets the value for Teams to be an explicit nil

### UnsetTeams
`func (o *PrincipalResponse) UnsetTeams()`

UnsetTeams ensures that no value is present for Teams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


