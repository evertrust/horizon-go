# EabAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedProfiles** | Pointer to **[]string** |  | [optional] [default to []]
**Description** | Pointer to **string** |  | [optional] 
**EabPolicy** | **string** |  | 
**EabValidityDuration** | Pointer to **NullableString** |  | [optional] 
**EmailConstraint** | Pointer to **string** |  | [optional] 
**IdentifierConstraint** | Pointer to **string** |  | [optional] 
**MacKeyAlgorithm** | [**ExternalAccountBindingAlgorithm**](ExternalAccountBindingAlgorithm.md) |  | 
**Name** | **string** |  | 
**ValidationMethods** | Pointer to [**[]AcmeAuthorizationType**](AcmeAuthorizationType.md) |  | [optional] [default to []]

## Methods

### NewEabAddRequest

`func NewEabAddRequest(eabPolicy string, macKeyAlgorithm ExternalAccountBindingAlgorithm, name string, ) *EabAddRequest`

NewEabAddRequest instantiates a new EabAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEabAddRequestWithDefaults

`func NewEabAddRequestWithDefaults() *EabAddRequest`

NewEabAddRequestWithDefaults instantiates a new EabAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedProfiles

`func (o *EabAddRequest) GetAllowedProfiles() []string`

GetAllowedProfiles returns the AllowedProfiles field if non-nil, zero value otherwise.

### GetAllowedProfilesOk

`func (o *EabAddRequest) GetAllowedProfilesOk() (*[]string, bool)`

GetAllowedProfilesOk returns a tuple with the AllowedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedProfiles

`func (o *EabAddRequest) SetAllowedProfiles(v []string)`

SetAllowedProfiles sets AllowedProfiles field to given value.

### HasAllowedProfiles

`func (o *EabAddRequest) HasAllowedProfiles() bool`

HasAllowedProfiles returns a boolean if a field has been set.

### GetDescription

`func (o *EabAddRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EabAddRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EabAddRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EabAddRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEabPolicy

`func (o *EabAddRequest) GetEabPolicy() string`

GetEabPolicy returns the EabPolicy field if non-nil, zero value otherwise.

### GetEabPolicyOk

`func (o *EabAddRequest) GetEabPolicyOk() (*string, bool)`

GetEabPolicyOk returns a tuple with the EabPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabPolicy

`func (o *EabAddRequest) SetEabPolicy(v string)`

SetEabPolicy sets EabPolicy field to given value.


### GetEabValidityDuration

`func (o *EabAddRequest) GetEabValidityDuration() string`

GetEabValidityDuration returns the EabValidityDuration field if non-nil, zero value otherwise.

### GetEabValidityDurationOk

`func (o *EabAddRequest) GetEabValidityDurationOk() (*string, bool)`

GetEabValidityDurationOk returns a tuple with the EabValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabValidityDuration

`func (o *EabAddRequest) SetEabValidityDuration(v string)`

SetEabValidityDuration sets EabValidityDuration field to given value.

### HasEabValidityDuration

`func (o *EabAddRequest) HasEabValidityDuration() bool`

HasEabValidityDuration returns a boolean if a field has been set.

### SetEabValidityDurationNil

`func (o *EabAddRequest) SetEabValidityDurationNil(b bool)`

 SetEabValidityDurationNil sets the value for EabValidityDuration to be an explicit nil

### UnsetEabValidityDuration
`func (o *EabAddRequest) UnsetEabValidityDuration()`

UnsetEabValidityDuration ensures that no value is present for EabValidityDuration, not even an explicit nil
### GetEmailConstraint

`func (o *EabAddRequest) GetEmailConstraint() string`

GetEmailConstraint returns the EmailConstraint field if non-nil, zero value otherwise.

### GetEmailConstraintOk

`func (o *EabAddRequest) GetEmailConstraintOk() (*string, bool)`

GetEmailConstraintOk returns a tuple with the EmailConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConstraint

`func (o *EabAddRequest) SetEmailConstraint(v string)`

SetEmailConstraint sets EmailConstraint field to given value.

### HasEmailConstraint

`func (o *EabAddRequest) HasEmailConstraint() bool`

HasEmailConstraint returns a boolean if a field has been set.

### GetIdentifierConstraint

`func (o *EabAddRequest) GetIdentifierConstraint() string`

GetIdentifierConstraint returns the IdentifierConstraint field if non-nil, zero value otherwise.

### GetIdentifierConstraintOk

`func (o *EabAddRequest) GetIdentifierConstraintOk() (*string, bool)`

GetIdentifierConstraintOk returns a tuple with the IdentifierConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierConstraint

`func (o *EabAddRequest) SetIdentifierConstraint(v string)`

SetIdentifierConstraint sets IdentifierConstraint field to given value.

### HasIdentifierConstraint

`func (o *EabAddRequest) HasIdentifierConstraint() bool`

HasIdentifierConstraint returns a boolean if a field has been set.

### GetMacKeyAlgorithm

`func (o *EabAddRequest) GetMacKeyAlgorithm() ExternalAccountBindingAlgorithm`

GetMacKeyAlgorithm returns the MacKeyAlgorithm field if non-nil, zero value otherwise.

### GetMacKeyAlgorithmOk

`func (o *EabAddRequest) GetMacKeyAlgorithmOk() (*ExternalAccountBindingAlgorithm, bool)`

GetMacKeyAlgorithmOk returns a tuple with the MacKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacKeyAlgorithm

`func (o *EabAddRequest) SetMacKeyAlgorithm(v ExternalAccountBindingAlgorithm)`

SetMacKeyAlgorithm sets MacKeyAlgorithm field to given value.


### GetName

`func (o *EabAddRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EabAddRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EabAddRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValidationMethods

`func (o *EabAddRequest) GetValidationMethods() []AcmeAuthorizationType`

GetValidationMethods returns the ValidationMethods field if non-nil, zero value otherwise.

### GetValidationMethodsOk

`func (o *EabAddRequest) GetValidationMethodsOk() (*[]AcmeAuthorizationType, bool)`

GetValidationMethodsOk returns a tuple with the ValidationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethods

`func (o *EabAddRequest) SetValidationMethods(v []AcmeAuthorizationType)`

SetValidationMethods sets ValidationMethods field to given value.

### HasValidationMethods

`func (o *EabAddRequest) HasValidationMethods() bool`

HasValidationMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


