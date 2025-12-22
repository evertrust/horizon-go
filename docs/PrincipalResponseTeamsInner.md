# PrincipalResponseTeamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**[]LocalizedStringResponse**](LocalizedStringResponse.md) |  | [optional] 
**DisplayName** | Pointer to [**[]LocalizedStringResponse**](LocalizedStringResponse.md) |  | [optional] 
**ExternallyManaged** | Pointer to **bool** | &#x60;true&#x60; if this team is externally managed (SCIM,...) | [optional] 
**Manager** | Pointer to **bool** | &#x60;true&#x60; if the principal is a manager of this team | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewPrincipalResponseTeamsInner

`func NewPrincipalResponseTeamsInner() *PrincipalResponseTeamsInner`

NewPrincipalResponseTeamsInner instantiates a new PrincipalResponseTeamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalResponseTeamsInnerWithDefaults

`func NewPrincipalResponseTeamsInnerWithDefaults() *PrincipalResponseTeamsInner`

NewPrincipalResponseTeamsInnerWithDefaults instantiates a new PrincipalResponseTeamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PrincipalResponseTeamsInner) GetDescription() []LocalizedStringResponse`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrincipalResponseTeamsInner) GetDescriptionOk() (*[]LocalizedStringResponse, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrincipalResponseTeamsInner) SetDescription(v []LocalizedStringResponse)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrincipalResponseTeamsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *PrincipalResponseTeamsInner) GetDisplayName() []LocalizedStringResponse`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PrincipalResponseTeamsInner) GetDisplayNameOk() (*[]LocalizedStringResponse, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PrincipalResponseTeamsInner) SetDisplayName(v []LocalizedStringResponse)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PrincipalResponseTeamsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternallyManaged

`func (o *PrincipalResponseTeamsInner) GetExternallyManaged() bool`

GetExternallyManaged returns the ExternallyManaged field if non-nil, zero value otherwise.

### GetExternallyManagedOk

`func (o *PrincipalResponseTeamsInner) GetExternallyManagedOk() (*bool, bool)`

GetExternallyManagedOk returns a tuple with the ExternallyManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyManaged

`func (o *PrincipalResponseTeamsInner) SetExternallyManaged(v bool)`

SetExternallyManaged sets ExternallyManaged field to given value.

### HasExternallyManaged

`func (o *PrincipalResponseTeamsInner) HasExternallyManaged() bool`

HasExternallyManaged returns a boolean if a field has been set.

### GetManager

`func (o *PrincipalResponseTeamsInner) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *PrincipalResponseTeamsInner) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *PrincipalResponseTeamsInner) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *PrincipalResponseTeamsInner) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetName

`func (o *PrincipalResponseTeamsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalResponseTeamsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalResponseTeamsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrincipalResponseTeamsInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


