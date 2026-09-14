# Authorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenges** | [**[]Challenge**](Challenge.md) |  | 
**Identifier** | [**Identifier**](Identifier.md) |  | 
**Status** | [**AuthorizationStatus**](AuthorizationStatus.md) |  | 
**Wildcard** | **bool** |  | 

## Methods

### NewAuthorization

`func NewAuthorization(challenges []Challenge, identifier Identifier, status AuthorizationStatus, wildcard bool, ) *Authorization`

NewAuthorization instantiates a new Authorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationWithDefaults

`func NewAuthorizationWithDefaults() *Authorization`

NewAuthorizationWithDefaults instantiates a new Authorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenges

`func (o *Authorization) GetChallenges() []Challenge`

GetChallenges returns the Challenges field if non-nil, zero value otherwise.

### GetChallengesOk

`func (o *Authorization) GetChallengesOk() (*[]Challenge, bool)`

GetChallengesOk returns a tuple with the Challenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenges

`func (o *Authorization) SetChallenges(v []Challenge)`

SetChallenges sets Challenges field to given value.


### GetIdentifier

`func (o *Authorization) GetIdentifier() Identifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Authorization) GetIdentifierOk() (*Identifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Authorization) SetIdentifier(v Identifier)`

SetIdentifier sets Identifier field to given value.


### GetStatus

`func (o *Authorization) GetStatus() AuthorizationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Authorization) GetStatusOk() (*AuthorizationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Authorization) SetStatus(v AuthorizationStatus)`

SetStatus sets Status field to given value.


### GetWildcard

`func (o *Authorization) GetWildcard() bool`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *Authorization) GetWildcardOk() (*bool, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *Authorization) SetWildcard(v bool)`

SetWildcard sets Wildcard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


