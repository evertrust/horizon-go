# MappingTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | The Scim group coming from the scim provider | 
**Team** | **string** | The mapped team on horizon | 

## Methods

### NewMappingTeam

`func NewMappingTeam(group string, team string, ) *MappingTeam`

NewMappingTeam instantiates a new MappingTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingTeamWithDefaults

`func NewMappingTeamWithDefaults() *MappingTeam`

NewMappingTeamWithDefaults instantiates a new MappingTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *MappingTeam) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MappingTeam) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MappingTeam) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetTeam

`func (o *MappingTeam) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MappingTeam) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MappingTeam) SetTeam(v string)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


