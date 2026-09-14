# EabUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedProfiles** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EabPolicy** | Pointer to **string** |  | [optional] 
**EmailConstraint** | Pointer to **string** |  | [optional] 
**IdentifierConstraint** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**ValidationMethods** | Pointer to [**[]AcmeAuthorizationType**](AcmeAuthorizationType.md) |  | [optional] 

## Methods

### NewEabUpdateRequest

`func NewEabUpdateRequest(name string, ) *EabUpdateRequest`

NewEabUpdateRequest instantiates a new EabUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEabUpdateRequestWithDefaults

`func NewEabUpdateRequestWithDefaults() *EabUpdateRequest`

NewEabUpdateRequestWithDefaults instantiates a new EabUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedProfiles

`func (o *EabUpdateRequest) GetAllowedProfiles() []string`

GetAllowedProfiles returns the AllowedProfiles field if non-nil, zero value otherwise.

### GetAllowedProfilesOk

`func (o *EabUpdateRequest) GetAllowedProfilesOk() (*[]string, bool)`

GetAllowedProfilesOk returns a tuple with the AllowedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedProfiles

`func (o *EabUpdateRequest) SetAllowedProfiles(v []string)`

SetAllowedProfiles sets AllowedProfiles field to given value.

### HasAllowedProfiles

`func (o *EabUpdateRequest) HasAllowedProfiles() bool`

HasAllowedProfiles returns a boolean if a field has been set.

### GetDescription

`func (o *EabUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EabUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EabUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EabUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEabPolicy

`func (o *EabUpdateRequest) GetEabPolicy() string`

GetEabPolicy returns the EabPolicy field if non-nil, zero value otherwise.

### GetEabPolicyOk

`func (o *EabUpdateRequest) GetEabPolicyOk() (*string, bool)`

GetEabPolicyOk returns a tuple with the EabPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabPolicy

`func (o *EabUpdateRequest) SetEabPolicy(v string)`

SetEabPolicy sets EabPolicy field to given value.

### HasEabPolicy

`func (o *EabUpdateRequest) HasEabPolicy() bool`

HasEabPolicy returns a boolean if a field has been set.

### GetEmailConstraint

`func (o *EabUpdateRequest) GetEmailConstraint() string`

GetEmailConstraint returns the EmailConstraint field if non-nil, zero value otherwise.

### GetEmailConstraintOk

`func (o *EabUpdateRequest) GetEmailConstraintOk() (*string, bool)`

GetEmailConstraintOk returns a tuple with the EmailConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConstraint

`func (o *EabUpdateRequest) SetEmailConstraint(v string)`

SetEmailConstraint sets EmailConstraint field to given value.

### HasEmailConstraint

`func (o *EabUpdateRequest) HasEmailConstraint() bool`

HasEmailConstraint returns a boolean if a field has been set.

### GetIdentifierConstraint

`func (o *EabUpdateRequest) GetIdentifierConstraint() string`

GetIdentifierConstraint returns the IdentifierConstraint field if non-nil, zero value otherwise.

### GetIdentifierConstraintOk

`func (o *EabUpdateRequest) GetIdentifierConstraintOk() (*string, bool)`

GetIdentifierConstraintOk returns a tuple with the IdentifierConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierConstraint

`func (o *EabUpdateRequest) SetIdentifierConstraint(v string)`

SetIdentifierConstraint sets IdentifierConstraint field to given value.

### HasIdentifierConstraint

`func (o *EabUpdateRequest) HasIdentifierConstraint() bool`

HasIdentifierConstraint returns a boolean if a field has been set.

### GetName

`func (o *EabUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EabUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EabUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValidationMethods

`func (o *EabUpdateRequest) GetValidationMethods() []AcmeAuthorizationType`

GetValidationMethods returns the ValidationMethods field if non-nil, zero value otherwise.

### GetValidationMethodsOk

`func (o *EabUpdateRequest) GetValidationMethodsOk() (*[]AcmeAuthorizationType, bool)`

GetValidationMethodsOk returns a tuple with the ValidationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethods

`func (o *EabUpdateRequest) SetValidationMethods(v []AcmeAuthorizationType)`

SetValidationMethods sets ValidationMethods field to given value.

### HasValidationMethods

`func (o *EabUpdateRequest) HasValidationMethods() bool`

HasValidationMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


