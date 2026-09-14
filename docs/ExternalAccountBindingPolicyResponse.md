# ExternalAccountBindingPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AllowedProfiles** | **[]string** |  | 
**EmailConstraint** | Pointer to **string** |  | [optional] 
**IdentifierConstraint** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**ValidationMethods** | [**[]AcmeAuthorizationType**](AcmeAuthorizationType.md) |  | 

## Methods

### NewExternalAccountBindingPolicyResponse

`func NewExternalAccountBindingPolicyResponse(id string, allowedProfiles []string, name string, validationMethods []AcmeAuthorizationType, ) *ExternalAccountBindingPolicyResponse`

NewExternalAccountBindingPolicyResponse instantiates a new ExternalAccountBindingPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountBindingPolicyResponseWithDefaults

`func NewExternalAccountBindingPolicyResponseWithDefaults() *ExternalAccountBindingPolicyResponse`

NewExternalAccountBindingPolicyResponseWithDefaults instantiates a new ExternalAccountBindingPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalAccountBindingPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalAccountBindingPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalAccountBindingPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAllowedProfiles

`func (o *ExternalAccountBindingPolicyResponse) GetAllowedProfiles() []string`

GetAllowedProfiles returns the AllowedProfiles field if non-nil, zero value otherwise.

### GetAllowedProfilesOk

`func (o *ExternalAccountBindingPolicyResponse) GetAllowedProfilesOk() (*[]string, bool)`

GetAllowedProfilesOk returns a tuple with the AllowedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedProfiles

`func (o *ExternalAccountBindingPolicyResponse) SetAllowedProfiles(v []string)`

SetAllowedProfiles sets AllowedProfiles field to given value.


### GetEmailConstraint

`func (o *ExternalAccountBindingPolicyResponse) GetEmailConstraint() string`

GetEmailConstraint returns the EmailConstraint field if non-nil, zero value otherwise.

### GetEmailConstraintOk

`func (o *ExternalAccountBindingPolicyResponse) GetEmailConstraintOk() (*string, bool)`

GetEmailConstraintOk returns a tuple with the EmailConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConstraint

`func (o *ExternalAccountBindingPolicyResponse) SetEmailConstraint(v string)`

SetEmailConstraint sets EmailConstraint field to given value.

### HasEmailConstraint

`func (o *ExternalAccountBindingPolicyResponse) HasEmailConstraint() bool`

HasEmailConstraint returns a boolean if a field has been set.

### GetIdentifierConstraint

`func (o *ExternalAccountBindingPolicyResponse) GetIdentifierConstraint() string`

GetIdentifierConstraint returns the IdentifierConstraint field if non-nil, zero value otherwise.

### GetIdentifierConstraintOk

`func (o *ExternalAccountBindingPolicyResponse) GetIdentifierConstraintOk() (*string, bool)`

GetIdentifierConstraintOk returns a tuple with the IdentifierConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierConstraint

`func (o *ExternalAccountBindingPolicyResponse) SetIdentifierConstraint(v string)`

SetIdentifierConstraint sets IdentifierConstraint field to given value.

### HasIdentifierConstraint

`func (o *ExternalAccountBindingPolicyResponse) HasIdentifierConstraint() bool`

HasIdentifierConstraint returns a boolean if a field has been set.

### GetName

`func (o *ExternalAccountBindingPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalAccountBindingPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalAccountBindingPolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetValidationMethods

`func (o *ExternalAccountBindingPolicyResponse) GetValidationMethods() []AcmeAuthorizationType`

GetValidationMethods returns the ValidationMethods field if non-nil, zero value otherwise.

### GetValidationMethodsOk

`func (o *ExternalAccountBindingPolicyResponse) GetValidationMethodsOk() (*[]AcmeAuthorizationType, bool)`

GetValidationMethodsOk returns a tuple with the ValidationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethods

`func (o *ExternalAccountBindingPolicyResponse) SetValidationMethods(v []AcmeAuthorizationType)`

SetValidationMethods sets ValidationMethods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


