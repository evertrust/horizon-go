# PrincipalInfoSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableString** | The identifier of the principal | [optional] 
**Contact** | Pointer to **NullableString** | The contact e-mail of the principal | [optional] 
**Role** | Pointer to **NullableString** | The role of the principal | [optional] 
**Team** | Pointer to **NullableString** | The team of the principal | [optional] 
**SortedBy** | Pointer to [**[]SortElement**](SortElement.md) | How to sort the results of the search | [optional] 
**PageIndex** | Pointer to **NullableInt64** | Which page result to display | [optional] 
**PageSize** | Pointer to **NullableInt64** | How many results to display per page | [optional] 
**WithCount** | Pointer to **NullableBool** | Whether to include the total number of results in the response | [optional] 

## Methods

### NewPrincipalInfoSearchQuery

`func NewPrincipalInfoSearchQuery() *PrincipalInfoSearchQuery`

NewPrincipalInfoSearchQuery instantiates a new PrincipalInfoSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalInfoSearchQueryWithDefaults

`func NewPrincipalInfoSearchQueryWithDefaults() *PrincipalInfoSearchQuery`

NewPrincipalInfoSearchQueryWithDefaults instantiates a new PrincipalInfoSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PrincipalInfoSearchQuery) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PrincipalInfoSearchQuery) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PrincipalInfoSearchQuery) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PrincipalInfoSearchQuery) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *PrincipalInfoSearchQuery) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *PrincipalInfoSearchQuery) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetContact

`func (o *PrincipalInfoSearchQuery) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PrincipalInfoSearchQuery) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PrincipalInfoSearchQuery) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PrincipalInfoSearchQuery) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PrincipalInfoSearchQuery) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PrincipalInfoSearchQuery) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetRole

`func (o *PrincipalInfoSearchQuery) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PrincipalInfoSearchQuery) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PrincipalInfoSearchQuery) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PrincipalInfoSearchQuery) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PrincipalInfoSearchQuery) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PrincipalInfoSearchQuery) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTeam

`func (o *PrincipalInfoSearchQuery) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PrincipalInfoSearchQuery) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *PrincipalInfoSearchQuery) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *PrincipalInfoSearchQuery) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *PrincipalInfoSearchQuery) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *PrincipalInfoSearchQuery) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetSortedBy

`func (o *PrincipalInfoSearchQuery) GetSortedBy() []SortElement`

GetSortedBy returns the SortedBy field if non-nil, zero value otherwise.

### GetSortedByOk

`func (o *PrincipalInfoSearchQuery) GetSortedByOk() (*[]SortElement, bool)`

GetSortedByOk returns a tuple with the SortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortedBy

`func (o *PrincipalInfoSearchQuery) SetSortedBy(v []SortElement)`

SetSortedBy sets SortedBy field to given value.

### HasSortedBy

`func (o *PrincipalInfoSearchQuery) HasSortedBy() bool`

HasSortedBy returns a boolean if a field has been set.

### SetSortedByNil

`func (o *PrincipalInfoSearchQuery) SetSortedByNil(b bool)`

 SetSortedByNil sets the value for SortedBy to be an explicit nil

### UnsetSortedBy
`func (o *PrincipalInfoSearchQuery) UnsetSortedBy()`

UnsetSortedBy ensures that no value is present for SortedBy, not even an explicit nil
### GetPageIndex

`func (o *PrincipalInfoSearchQuery) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *PrincipalInfoSearchQuery) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *PrincipalInfoSearchQuery) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *PrincipalInfoSearchQuery) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### SetPageIndexNil

`func (o *PrincipalInfoSearchQuery) SetPageIndexNil(b bool)`

 SetPageIndexNil sets the value for PageIndex to be an explicit nil

### UnsetPageIndex
`func (o *PrincipalInfoSearchQuery) UnsetPageIndex()`

UnsetPageIndex ensures that no value is present for PageIndex, not even an explicit nil
### GetPageSize

`func (o *PrincipalInfoSearchQuery) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PrincipalInfoSearchQuery) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PrincipalInfoSearchQuery) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PrincipalInfoSearchQuery) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *PrincipalInfoSearchQuery) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *PrincipalInfoSearchQuery) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetWithCount

`func (o *PrincipalInfoSearchQuery) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *PrincipalInfoSearchQuery) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *PrincipalInfoSearchQuery) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *PrincipalInfoSearchQuery) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.

### SetWithCountNil

`func (o *PrincipalInfoSearchQuery) SetWithCountNil(b bool)`

 SetWithCountNil sets the value for WithCount to be an explicit nil

### UnsetWithCount
`func (o *PrincipalInfoSearchQuery) UnsetWithCount()`

UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


