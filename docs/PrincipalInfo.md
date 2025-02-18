# PrincipalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The identifier of the principal | 
**Contact** | Pointer to **NullableString** | The contact e-mail of the principal | [optional] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) | The permissions of the principal | [optional] 
**Roles** | Pointer to **[]string** | The roles of the principal | [optional] 
**Teams** | Pointer to **[]string** | The teams of the principal | [optional] 
**SavedQueries** | Pointer to [**[]PrincipalInfoSavedQuery**](PrincipalInfoSavedQuery.md) | The saved HQL queries of the principal. This is used by UI only. These values should not be manually set but should be copied on update | [optional] 
**CustomDashboards** | Pointer to [**[]Dashboard**](Dashboard.md) | The custom dashboards of the principal. This is used by UI only. These values should not be manually set but should be copied on update | [optional] 
**Preferences** | Pointer to [**NullablePrincipalInfoPreferences**](PrincipalInfoPreferences.md) | The UI preferences of the principal. This is used by UI only. These values should not be manually set but should be copied on update | [optional] 
**Enabled** | **bool** | If the principal is allowed to login horizon | 

## Methods

### NewPrincipalInfo

`func NewPrincipalInfo(identifier string, enabled bool, ) *PrincipalInfo`

NewPrincipalInfo instantiates a new PrincipalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalInfoWithDefaults

`func NewPrincipalInfoWithDefaults() *PrincipalInfo`

NewPrincipalInfoWithDefaults instantiates a new PrincipalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PrincipalInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PrincipalInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PrincipalInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetContact

`func (o *PrincipalInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PrincipalInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PrincipalInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PrincipalInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PrincipalInfo) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PrincipalInfo) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetPermissions

`func (o *PrincipalInfo) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PrincipalInfo) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PrincipalInfo) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PrincipalInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *PrincipalInfo) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *PrincipalInfo) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetRoles

`func (o *PrincipalInfo) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PrincipalInfo) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PrincipalInfo) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PrincipalInfo) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *PrincipalInfo) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *PrincipalInfo) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetTeams

`func (o *PrincipalInfo) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *PrincipalInfo) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *PrincipalInfo) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *PrincipalInfo) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### SetTeamsNil

`func (o *PrincipalInfo) SetTeamsNil(b bool)`

 SetTeamsNil sets the value for Teams to be an explicit nil

### UnsetTeams
`func (o *PrincipalInfo) UnsetTeams()`

UnsetTeams ensures that no value is present for Teams, not even an explicit nil
### GetSavedQueries

`func (o *PrincipalInfo) GetSavedQueries() []PrincipalInfoSavedQuery`

GetSavedQueries returns the SavedQueries field if non-nil, zero value otherwise.

### GetSavedQueriesOk

`func (o *PrincipalInfo) GetSavedQueriesOk() (*[]PrincipalInfoSavedQuery, bool)`

GetSavedQueriesOk returns a tuple with the SavedQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedQueries

`func (o *PrincipalInfo) SetSavedQueries(v []PrincipalInfoSavedQuery)`

SetSavedQueries sets SavedQueries field to given value.

### HasSavedQueries

`func (o *PrincipalInfo) HasSavedQueries() bool`

HasSavedQueries returns a boolean if a field has been set.

### SetSavedQueriesNil

`func (o *PrincipalInfo) SetSavedQueriesNil(b bool)`

 SetSavedQueriesNil sets the value for SavedQueries to be an explicit nil

### UnsetSavedQueries
`func (o *PrincipalInfo) UnsetSavedQueries()`

UnsetSavedQueries ensures that no value is present for SavedQueries, not even an explicit nil
### GetCustomDashboards

`func (o *PrincipalInfo) GetCustomDashboards() []Dashboard`

GetCustomDashboards returns the CustomDashboards field if non-nil, zero value otherwise.

### GetCustomDashboardsOk

`func (o *PrincipalInfo) GetCustomDashboardsOk() (*[]Dashboard, bool)`

GetCustomDashboardsOk returns a tuple with the CustomDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDashboards

`func (o *PrincipalInfo) SetCustomDashboards(v []Dashboard)`

SetCustomDashboards sets CustomDashboards field to given value.

### HasCustomDashboards

`func (o *PrincipalInfo) HasCustomDashboards() bool`

HasCustomDashboards returns a boolean if a field has been set.

### SetCustomDashboardsNil

`func (o *PrincipalInfo) SetCustomDashboardsNil(b bool)`

 SetCustomDashboardsNil sets the value for CustomDashboards to be an explicit nil

### UnsetCustomDashboards
`func (o *PrincipalInfo) UnsetCustomDashboards()`

UnsetCustomDashboards ensures that no value is present for CustomDashboards, not even an explicit nil
### GetPreferences

`func (o *PrincipalInfo) GetPreferences() PrincipalInfoPreferences`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *PrincipalInfo) GetPreferencesOk() (*PrincipalInfoPreferences, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *PrincipalInfo) SetPreferences(v PrincipalInfoPreferences)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *PrincipalInfo) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### SetPreferencesNil

`func (o *PrincipalInfo) SetPreferencesNil(b bool)`

 SetPreferencesNil sets the value for Preferences to be an explicit nil

### UnsetPreferences
`func (o *PrincipalInfo) UnsetPreferences()`

UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
### GetEnabled

`func (o *PrincipalInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PrincipalInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PrincipalInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


