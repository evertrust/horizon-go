# EabUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AllowedProfiles** | **[]string** |  | 
**CompromisedAt** | Pointer to **int64** | The date when the EAB was compromised | [optional] 
**CompromissionReason** | Pointer to **string** | One of: &#x60;unspecified&#x60;, &#x60;keycompromise&#x60;, &#x60;cacompromise&#x60;, &#x60;affiliationchange&#x60;, &#x60;superseded&#x60;, &#x60;cessationofoperation&#x60; | [optional] 
**CreatedAt** | **int64** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EabPolicy** | **string** |  | 
**EmailConstraint** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **int64** |  | [optional] 
**IdentifierConstraint** | Pointer to **string** |  | [optional] 
**MacKeyAlgorithm** | [**ExternalAccountBindingAlgorithm**](ExternalAccountBindingAlgorithm.md) |  | 
**Name** | **string** |  | 
**NumberOfKeyRegeneration** | **int64** |  | 
**Status** | [**ExternalAccountBindingStatus**](ExternalAccountBindingStatus.md) |  | 
**ValidationMethods** | [**[]AcmeAuthorizationType**](AcmeAuthorizationType.md) |  | 

## Methods

### NewEabUpdateResponse

`func NewEabUpdateResponse(id string, allowedProfiles []string, createdAt int64, eabPolicy string, macKeyAlgorithm ExternalAccountBindingAlgorithm, name string, numberOfKeyRegeneration int64, status ExternalAccountBindingStatus, validationMethods []AcmeAuthorizationType, ) *EabUpdateResponse`

NewEabUpdateResponse instantiates a new EabUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEabUpdateResponseWithDefaults

`func NewEabUpdateResponseWithDefaults() *EabUpdateResponse`

NewEabUpdateResponseWithDefaults instantiates a new EabUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EabUpdateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EabUpdateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EabUpdateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAllowedProfiles

`func (o *EabUpdateResponse) GetAllowedProfiles() []string`

GetAllowedProfiles returns the AllowedProfiles field if non-nil, zero value otherwise.

### GetAllowedProfilesOk

`func (o *EabUpdateResponse) GetAllowedProfilesOk() (*[]string, bool)`

GetAllowedProfilesOk returns a tuple with the AllowedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedProfiles

`func (o *EabUpdateResponse) SetAllowedProfiles(v []string)`

SetAllowedProfiles sets AllowedProfiles field to given value.


### GetCompromisedAt

`func (o *EabUpdateResponse) GetCompromisedAt() int64`

GetCompromisedAt returns the CompromisedAt field if non-nil, zero value otherwise.

### GetCompromisedAtOk

`func (o *EabUpdateResponse) GetCompromisedAtOk() (*int64, bool)`

GetCompromisedAtOk returns a tuple with the CompromisedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedAt

`func (o *EabUpdateResponse) SetCompromisedAt(v int64)`

SetCompromisedAt sets CompromisedAt field to given value.

### HasCompromisedAt

`func (o *EabUpdateResponse) HasCompromisedAt() bool`

HasCompromisedAt returns a boolean if a field has been set.

### GetCompromissionReason

`func (o *EabUpdateResponse) GetCompromissionReason() string`

GetCompromissionReason returns the CompromissionReason field if non-nil, zero value otherwise.

### GetCompromissionReasonOk

`func (o *EabUpdateResponse) GetCompromissionReasonOk() (*string, bool)`

GetCompromissionReasonOk returns a tuple with the CompromissionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromissionReason

`func (o *EabUpdateResponse) SetCompromissionReason(v string)`

SetCompromissionReason sets CompromissionReason field to given value.

### HasCompromissionReason

`func (o *EabUpdateResponse) HasCompromissionReason() bool`

HasCompromissionReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EabUpdateResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EabUpdateResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EabUpdateResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *EabUpdateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EabUpdateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EabUpdateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EabUpdateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEabPolicy

`func (o *EabUpdateResponse) GetEabPolicy() string`

GetEabPolicy returns the EabPolicy field if non-nil, zero value otherwise.

### GetEabPolicyOk

`func (o *EabUpdateResponse) GetEabPolicyOk() (*string, bool)`

GetEabPolicyOk returns a tuple with the EabPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabPolicy

`func (o *EabUpdateResponse) SetEabPolicy(v string)`

SetEabPolicy sets EabPolicy field to given value.


### GetEmailConstraint

`func (o *EabUpdateResponse) GetEmailConstraint() string`

GetEmailConstraint returns the EmailConstraint field if non-nil, zero value otherwise.

### GetEmailConstraintOk

`func (o *EabUpdateResponse) GetEmailConstraintOk() (*string, bool)`

GetEmailConstraintOk returns a tuple with the EmailConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConstraint

`func (o *EabUpdateResponse) SetEmailConstraint(v string)`

SetEmailConstraint sets EmailConstraint field to given value.

### HasEmailConstraint

`func (o *EabUpdateResponse) HasEmailConstraint() bool`

HasEmailConstraint returns a boolean if a field has been set.

### GetExpirationDate

`func (o *EabUpdateResponse) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *EabUpdateResponse) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *EabUpdateResponse) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *EabUpdateResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIdentifierConstraint

`func (o *EabUpdateResponse) GetIdentifierConstraint() string`

GetIdentifierConstraint returns the IdentifierConstraint field if non-nil, zero value otherwise.

### GetIdentifierConstraintOk

`func (o *EabUpdateResponse) GetIdentifierConstraintOk() (*string, bool)`

GetIdentifierConstraintOk returns a tuple with the IdentifierConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierConstraint

`func (o *EabUpdateResponse) SetIdentifierConstraint(v string)`

SetIdentifierConstraint sets IdentifierConstraint field to given value.

### HasIdentifierConstraint

`func (o *EabUpdateResponse) HasIdentifierConstraint() bool`

HasIdentifierConstraint returns a boolean if a field has been set.

### GetMacKeyAlgorithm

`func (o *EabUpdateResponse) GetMacKeyAlgorithm() ExternalAccountBindingAlgorithm`

GetMacKeyAlgorithm returns the MacKeyAlgorithm field if non-nil, zero value otherwise.

### GetMacKeyAlgorithmOk

`func (o *EabUpdateResponse) GetMacKeyAlgorithmOk() (*ExternalAccountBindingAlgorithm, bool)`

GetMacKeyAlgorithmOk returns a tuple with the MacKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacKeyAlgorithm

`func (o *EabUpdateResponse) SetMacKeyAlgorithm(v ExternalAccountBindingAlgorithm)`

SetMacKeyAlgorithm sets MacKeyAlgorithm field to given value.


### GetName

`func (o *EabUpdateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EabUpdateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EabUpdateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNumberOfKeyRegeneration

`func (o *EabUpdateResponse) GetNumberOfKeyRegeneration() int64`

GetNumberOfKeyRegeneration returns the NumberOfKeyRegeneration field if non-nil, zero value otherwise.

### GetNumberOfKeyRegenerationOk

`func (o *EabUpdateResponse) GetNumberOfKeyRegenerationOk() (*int64, bool)`

GetNumberOfKeyRegenerationOk returns a tuple with the NumberOfKeyRegeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfKeyRegeneration

`func (o *EabUpdateResponse) SetNumberOfKeyRegeneration(v int64)`

SetNumberOfKeyRegeneration sets NumberOfKeyRegeneration field to given value.


### GetStatus

`func (o *EabUpdateResponse) GetStatus() ExternalAccountBindingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EabUpdateResponse) GetStatusOk() (*ExternalAccountBindingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EabUpdateResponse) SetStatus(v ExternalAccountBindingStatus)`

SetStatus sets Status field to given value.


### GetValidationMethods

`func (o *EabUpdateResponse) GetValidationMethods() []AcmeAuthorizationType`

GetValidationMethods returns the ValidationMethods field if non-nil, zero value otherwise.

### GetValidationMethodsOk

`func (o *EabUpdateResponse) GetValidationMethodsOk() (*[]AcmeAuthorizationType, bool)`

GetValidationMethodsOk returns a tuple with the ValidationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethods

`func (o *EabUpdateResponse) SetValidationMethods(v []AcmeAuthorizationType)`

SetValidationMethods sets ValidationMethods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


