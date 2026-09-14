# EabPolicyAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedProfiles** | Pointer to **[]string** |  | [optional] [default to []]
**EmailConstraint** | Pointer to **string** |  | [optional] 
**IdentifierConstraint** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**ValidationMethods** | Pointer to [**[]AcmeAuthorizationType**](AcmeAuthorizationType.md) |  | [optional] [default to []]

## Methods

### NewEabPolicyAddRequest

`func NewEabPolicyAddRequest(name string, ) *EabPolicyAddRequest`

NewEabPolicyAddRequest instantiates a new EabPolicyAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEabPolicyAddRequestWithDefaults

`func NewEabPolicyAddRequestWithDefaults() *EabPolicyAddRequest`

NewEabPolicyAddRequestWithDefaults instantiates a new EabPolicyAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedProfiles

`func (o *EabPolicyAddRequest) GetAllowedProfiles() []string`

GetAllowedProfiles returns the AllowedProfiles field if non-nil, zero value otherwise.

### GetAllowedProfilesOk

`func (o *EabPolicyAddRequest) GetAllowedProfilesOk() (*[]string, bool)`

GetAllowedProfilesOk returns a tuple with the AllowedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedProfiles

`func (o *EabPolicyAddRequest) SetAllowedProfiles(v []string)`

SetAllowedProfiles sets AllowedProfiles field to given value.

### HasAllowedProfiles

`func (o *EabPolicyAddRequest) HasAllowedProfiles() bool`

HasAllowedProfiles returns a boolean if a field has been set.

### GetEmailConstraint

`func (o *EabPolicyAddRequest) GetEmailConstraint() string`

GetEmailConstraint returns the EmailConstraint field if non-nil, zero value otherwise.

### GetEmailConstraintOk

`func (o *EabPolicyAddRequest) GetEmailConstraintOk() (*string, bool)`

GetEmailConstraintOk returns a tuple with the EmailConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConstraint

`func (o *EabPolicyAddRequest) SetEmailConstraint(v string)`

SetEmailConstraint sets EmailConstraint field to given value.

### HasEmailConstraint

`func (o *EabPolicyAddRequest) HasEmailConstraint() bool`

HasEmailConstraint returns a boolean if a field has been set.

### GetIdentifierConstraint

`func (o *EabPolicyAddRequest) GetIdentifierConstraint() string`

GetIdentifierConstraint returns the IdentifierConstraint field if non-nil, zero value otherwise.

### GetIdentifierConstraintOk

`func (o *EabPolicyAddRequest) GetIdentifierConstraintOk() (*string, bool)`

GetIdentifierConstraintOk returns a tuple with the IdentifierConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierConstraint

`func (o *EabPolicyAddRequest) SetIdentifierConstraint(v string)`

SetIdentifierConstraint sets IdentifierConstraint field to given value.

### HasIdentifierConstraint

`func (o *EabPolicyAddRequest) HasIdentifierConstraint() bool`

HasIdentifierConstraint returns a boolean if a field has been set.

### GetName

`func (o *EabPolicyAddRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EabPolicyAddRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EabPolicyAddRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValidationMethods

`func (o *EabPolicyAddRequest) GetValidationMethods() []AcmeAuthorizationType`

GetValidationMethods returns the ValidationMethods field if non-nil, zero value otherwise.

### GetValidationMethodsOk

`func (o *EabPolicyAddRequest) GetValidationMethodsOk() (*[]AcmeAuthorizationType, bool)`

GetValidationMethodsOk returns a tuple with the ValidationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethods

`func (o *EabPolicyAddRequest) SetValidationMethods(v []AcmeAuthorizationType)`

SetValidationMethods sets ValidationMethods field to given value.

### HasValidationMethods

`func (o *EabPolicyAddRequest) HasValidationMethods() bool`

HasValidationMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


