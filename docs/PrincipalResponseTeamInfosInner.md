# PrincipalResponseTeamInfosInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**[]LocalizedStringResponse**](LocalizedStringResponse.md) |  | [optional] 
**DisplayName** | Pointer to [**[]LocalizedStringResponse**](LocalizedStringResponse.md) |  | [optional] 
**ExternallyManaged** | Pointer to **bool** | &#x60;true&#x60; if this team is externally managed (SCIM,...) | [optional] 
**Manager** | Pointer to **bool** | &#x60;true&#x60; if the principal is a manager of this team | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewPrincipalResponseTeamInfosInner

`func NewPrincipalResponseTeamInfosInner() *PrincipalResponseTeamInfosInner`

NewPrincipalResponseTeamInfosInner instantiates a new PrincipalResponseTeamInfosInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalResponseTeamInfosInnerWithDefaults

`func NewPrincipalResponseTeamInfosInnerWithDefaults() *PrincipalResponseTeamInfosInner`

NewPrincipalResponseTeamInfosInnerWithDefaults instantiates a new PrincipalResponseTeamInfosInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PrincipalResponseTeamInfosInner) GetDescription() []LocalizedStringResponse`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrincipalResponseTeamInfosInner) GetDescriptionOk() (*[]LocalizedStringResponse, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrincipalResponseTeamInfosInner) SetDescription(v []LocalizedStringResponse)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrincipalResponseTeamInfosInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *PrincipalResponseTeamInfosInner) GetDisplayName() []LocalizedStringResponse`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PrincipalResponseTeamInfosInner) GetDisplayNameOk() (*[]LocalizedStringResponse, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PrincipalResponseTeamInfosInner) SetDisplayName(v []LocalizedStringResponse)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PrincipalResponseTeamInfosInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternallyManaged

`func (o *PrincipalResponseTeamInfosInner) GetExternallyManaged() bool`

GetExternallyManaged returns the ExternallyManaged field if non-nil, zero value otherwise.

### GetExternallyManagedOk

`func (o *PrincipalResponseTeamInfosInner) GetExternallyManagedOk() (*bool, bool)`

GetExternallyManagedOk returns a tuple with the ExternallyManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyManaged

`func (o *PrincipalResponseTeamInfosInner) SetExternallyManaged(v bool)`

SetExternallyManaged sets ExternallyManaged field to given value.

### HasExternallyManaged

`func (o *PrincipalResponseTeamInfosInner) HasExternallyManaged() bool`

HasExternallyManaged returns a boolean if a field has been set.

### GetManager

`func (o *PrincipalResponseTeamInfosInner) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *PrincipalResponseTeamInfosInner) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *PrincipalResponseTeamInfosInner) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *PrincipalResponseTeamInfosInner) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetName

`func (o *PrincipalResponseTeamInfosInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalResponseTeamInfosInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalResponseTeamInfosInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrincipalResponseTeamInfosInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


