# PrincipalInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Contact** | Pointer to **NullableString** | The contact e-mail of the principal | [optional] 
**CreationDate** | Pointer to **int64** | The creation date of the principal (UNIX Timestamp in milliseconds) | [optional] 
**CustomDashboards** | Pointer to [**[]Dashboard**](Dashboard.md) | The custom dashboards of the principal. This is used by UI only. These values should not be manually set but should be copied on update | [optional] 
**Enabled** | **bool** | If the principal is allowed to login horizon | 
**Identifier** | **string** | The identifier of the principal | 
**LastAuthentication** | Pointer to **NullableInt64** | The last authentication date of the principal (UNIX Timestamp in milliseconds) | [optional] 
**LastModification** | Pointer to **NullableInt64** | The last modification date of the principal (UNIX Timestamp in milliseconds) | [optional] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) | The permissions of the principal | [optional] 
**Preferences** | Pointer to [**NullablePrincipalInfoPreferences**](PrincipalInfoPreferences.md) | The UI preferences of the principal. This is used by UI only. These values should not be manually set but should be copied on update | [optional] 
**Roles** | Pointer to **[]string** | The roles of the principal | [optional] 
**SavedQueries** | Pointer to [**[]PrincipalInfoSavedQuery**](PrincipalInfoSavedQuery.md) | The saved HQL queries of the principal. This is used by UI only. These values should not be manually set but should be copied on update | [optional] 
**Teams** | Pointer to **[]string** | The teams of the principal | [optional] 

## Methods

### NewPrincipalInfoResponse

`func NewPrincipalInfoResponse(id string, enabled bool, identifier string, ) *PrincipalInfoResponse`

NewPrincipalInfoResponse instantiates a new PrincipalInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalInfoResponseWithDefaults

`func NewPrincipalInfoResponseWithDefaults() *PrincipalInfoResponse`

NewPrincipalInfoResponseWithDefaults instantiates a new PrincipalInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrincipalInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetContact

`func (o *PrincipalInfoResponse) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PrincipalInfoResponse) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PrincipalInfoResponse) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PrincipalInfoResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PrincipalInfoResponse) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PrincipalInfoResponse) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetCreationDate

`func (o *PrincipalInfoResponse) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *PrincipalInfoResponse) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *PrincipalInfoResponse) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *PrincipalInfoResponse) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomDashboards

`func (o *PrincipalInfoResponse) GetCustomDashboards() []Dashboard`

GetCustomDashboards returns the CustomDashboards field if non-nil, zero value otherwise.

### GetCustomDashboardsOk

`func (o *PrincipalInfoResponse) GetCustomDashboardsOk() (*[]Dashboard, bool)`

GetCustomDashboardsOk returns a tuple with the CustomDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDashboards

`func (o *PrincipalInfoResponse) SetCustomDashboards(v []Dashboard)`

SetCustomDashboards sets CustomDashboards field to given value.

### HasCustomDashboards

`func (o *PrincipalInfoResponse) HasCustomDashboards() bool`

HasCustomDashboards returns a boolean if a field has been set.

### SetCustomDashboardsNil

`func (o *PrincipalInfoResponse) SetCustomDashboardsNil(b bool)`

 SetCustomDashboardsNil sets the value for CustomDashboards to be an explicit nil

### UnsetCustomDashboards
`func (o *PrincipalInfoResponse) UnsetCustomDashboards()`

UnsetCustomDashboards ensures that no value is present for CustomDashboards, not even an explicit nil
### GetEnabled

`func (o *PrincipalInfoResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PrincipalInfoResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PrincipalInfoResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIdentifier

`func (o *PrincipalInfoResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PrincipalInfoResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PrincipalInfoResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetLastAuthentication

`func (o *PrincipalInfoResponse) GetLastAuthentication() int64`

GetLastAuthentication returns the LastAuthentication field if non-nil, zero value otherwise.

### GetLastAuthenticationOk

`func (o *PrincipalInfoResponse) GetLastAuthenticationOk() (*int64, bool)`

GetLastAuthenticationOk returns a tuple with the LastAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthentication

`func (o *PrincipalInfoResponse) SetLastAuthentication(v int64)`

SetLastAuthentication sets LastAuthentication field to given value.

### HasLastAuthentication

`func (o *PrincipalInfoResponse) HasLastAuthentication() bool`

HasLastAuthentication returns a boolean if a field has been set.

### SetLastAuthenticationNil

`func (o *PrincipalInfoResponse) SetLastAuthenticationNil(b bool)`

 SetLastAuthenticationNil sets the value for LastAuthentication to be an explicit nil

### UnsetLastAuthentication
`func (o *PrincipalInfoResponse) UnsetLastAuthentication()`

UnsetLastAuthentication ensures that no value is present for LastAuthentication, not even an explicit nil
### GetLastModification

`func (o *PrincipalInfoResponse) GetLastModification() int64`

GetLastModification returns the LastModification field if non-nil, zero value otherwise.

### GetLastModificationOk

`func (o *PrincipalInfoResponse) GetLastModificationOk() (*int64, bool)`

GetLastModificationOk returns a tuple with the LastModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModification

`func (o *PrincipalInfoResponse) SetLastModification(v int64)`

SetLastModification sets LastModification field to given value.

### HasLastModification

`func (o *PrincipalInfoResponse) HasLastModification() bool`

HasLastModification returns a boolean if a field has been set.

### SetLastModificationNil

`func (o *PrincipalInfoResponse) SetLastModificationNil(b bool)`

 SetLastModificationNil sets the value for LastModification to be an explicit nil

### UnsetLastModification
`func (o *PrincipalInfoResponse) UnsetLastModification()`

UnsetLastModification ensures that no value is present for LastModification, not even an explicit nil
### GetPermissions

`func (o *PrincipalInfoResponse) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PrincipalInfoResponse) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PrincipalInfoResponse) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PrincipalInfoResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *PrincipalInfoResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *PrincipalInfoResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPreferences

`func (o *PrincipalInfoResponse) GetPreferences() PrincipalInfoPreferences`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *PrincipalInfoResponse) GetPreferencesOk() (*PrincipalInfoPreferences, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *PrincipalInfoResponse) SetPreferences(v PrincipalInfoPreferences)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *PrincipalInfoResponse) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### SetPreferencesNil

`func (o *PrincipalInfoResponse) SetPreferencesNil(b bool)`

 SetPreferencesNil sets the value for Preferences to be an explicit nil

### UnsetPreferences
`func (o *PrincipalInfoResponse) UnsetPreferences()`

UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
### GetRoles

`func (o *PrincipalInfoResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PrincipalInfoResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PrincipalInfoResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PrincipalInfoResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *PrincipalInfoResponse) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *PrincipalInfoResponse) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSavedQueries

`func (o *PrincipalInfoResponse) GetSavedQueries() []PrincipalInfoSavedQuery`

GetSavedQueries returns the SavedQueries field if non-nil, zero value otherwise.

### GetSavedQueriesOk

`func (o *PrincipalInfoResponse) GetSavedQueriesOk() (*[]PrincipalInfoSavedQuery, bool)`

GetSavedQueriesOk returns a tuple with the SavedQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedQueries

`func (o *PrincipalInfoResponse) SetSavedQueries(v []PrincipalInfoSavedQuery)`

SetSavedQueries sets SavedQueries field to given value.

### HasSavedQueries

`func (o *PrincipalInfoResponse) HasSavedQueries() bool`

HasSavedQueries returns a boolean if a field has been set.

### SetSavedQueriesNil

`func (o *PrincipalInfoResponse) SetSavedQueriesNil(b bool)`

 SetSavedQueriesNil sets the value for SavedQueries to be an explicit nil

### UnsetSavedQueries
`func (o *PrincipalInfoResponse) UnsetSavedQueries()`

UnsetSavedQueries ensures that no value is present for SavedQueries, not even an explicit nil
### GetTeams

`func (o *PrincipalInfoResponse) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *PrincipalInfoResponse) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *PrincipalInfoResponse) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *PrincipalInfoResponse) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### SetTeamsNil

`func (o *PrincipalInfoResponse) SetTeamsNil(b bool)`

 SetTeamsNil sets the value for Teams to be an explicit nil

### UnsetTeams
`func (o *PrincipalInfoResponse) UnsetTeams()`

UnsetTeams ensures that no value is present for Teams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


