# EabRenewResponse

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
**MacKey** | **string** | The base64url-encoded MAC key of the External Account Binding. It is only returned in this response and cannot be retrieved afterwards. | 
**MacKeyAlgorithm** | [**ExternalAccountBindingAlgorithm**](ExternalAccountBindingAlgorithm.md) |  | 
**MacKeyId** | **string** | The MAC key identifier of the External Account Binding. It is only returned in this response and cannot be retrieved afterwards. | 
**Name** | **string** |  | 
**NumberOfKeyRegeneration** | **int64** |  | 
**Status** | [**ExternalAccountBindingStatus**](ExternalAccountBindingStatus.md) | The current status of the External Account Binding. See ExternalAccountBindingStatus for impact details. | 
**ValidationMethods** | [**[]AcmeAuthorizationType**](AcmeAuthorizationType.md) |  | 

## Methods

### NewEabRenewResponse

`func NewEabRenewResponse(id string, allowedProfiles []string, createdAt int64, eabPolicy string, macKey string, macKeyAlgorithm ExternalAccountBindingAlgorithm, macKeyId string, name string, numberOfKeyRegeneration int64, status ExternalAccountBindingStatus, validationMethods []AcmeAuthorizationType, ) *EabRenewResponse`

NewEabRenewResponse instantiates a new EabRenewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEabRenewResponseWithDefaults

`func NewEabRenewResponseWithDefaults() *EabRenewResponse`

NewEabRenewResponseWithDefaults instantiates a new EabRenewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EabRenewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EabRenewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EabRenewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAllowedProfiles

`func (o *EabRenewResponse) GetAllowedProfiles() []string`

GetAllowedProfiles returns the AllowedProfiles field if non-nil, zero value otherwise.

### GetAllowedProfilesOk

`func (o *EabRenewResponse) GetAllowedProfilesOk() (*[]string, bool)`

GetAllowedProfilesOk returns a tuple with the AllowedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedProfiles

`func (o *EabRenewResponse) SetAllowedProfiles(v []string)`

SetAllowedProfiles sets AllowedProfiles field to given value.


### GetCompromisedAt

`func (o *EabRenewResponse) GetCompromisedAt() int64`

GetCompromisedAt returns the CompromisedAt field if non-nil, zero value otherwise.

### GetCompromisedAtOk

`func (o *EabRenewResponse) GetCompromisedAtOk() (*int64, bool)`

GetCompromisedAtOk returns a tuple with the CompromisedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedAt

`func (o *EabRenewResponse) SetCompromisedAt(v int64)`

SetCompromisedAt sets CompromisedAt field to given value.

### HasCompromisedAt

`func (o *EabRenewResponse) HasCompromisedAt() bool`

HasCompromisedAt returns a boolean if a field has been set.

### GetCompromissionReason

`func (o *EabRenewResponse) GetCompromissionReason() string`

GetCompromissionReason returns the CompromissionReason field if non-nil, zero value otherwise.

### GetCompromissionReasonOk

`func (o *EabRenewResponse) GetCompromissionReasonOk() (*string, bool)`

GetCompromissionReasonOk returns a tuple with the CompromissionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromissionReason

`func (o *EabRenewResponse) SetCompromissionReason(v string)`

SetCompromissionReason sets CompromissionReason field to given value.

### HasCompromissionReason

`func (o *EabRenewResponse) HasCompromissionReason() bool`

HasCompromissionReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EabRenewResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EabRenewResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EabRenewResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *EabRenewResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EabRenewResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EabRenewResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EabRenewResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEabPolicy

`func (o *EabRenewResponse) GetEabPolicy() string`

GetEabPolicy returns the EabPolicy field if non-nil, zero value otherwise.

### GetEabPolicyOk

`func (o *EabRenewResponse) GetEabPolicyOk() (*string, bool)`

GetEabPolicyOk returns a tuple with the EabPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabPolicy

`func (o *EabRenewResponse) SetEabPolicy(v string)`

SetEabPolicy sets EabPolicy field to given value.


### GetEmailConstraint

`func (o *EabRenewResponse) GetEmailConstraint() string`

GetEmailConstraint returns the EmailConstraint field if non-nil, zero value otherwise.

### GetEmailConstraintOk

`func (o *EabRenewResponse) GetEmailConstraintOk() (*string, bool)`

GetEmailConstraintOk returns a tuple with the EmailConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConstraint

`func (o *EabRenewResponse) SetEmailConstraint(v string)`

SetEmailConstraint sets EmailConstraint field to given value.

### HasEmailConstraint

`func (o *EabRenewResponse) HasEmailConstraint() bool`

HasEmailConstraint returns a boolean if a field has been set.

### GetExpirationDate

`func (o *EabRenewResponse) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *EabRenewResponse) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *EabRenewResponse) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *EabRenewResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIdentifierConstraint

`func (o *EabRenewResponse) GetIdentifierConstraint() string`

GetIdentifierConstraint returns the IdentifierConstraint field if non-nil, zero value otherwise.

### GetIdentifierConstraintOk

`func (o *EabRenewResponse) GetIdentifierConstraintOk() (*string, bool)`

GetIdentifierConstraintOk returns a tuple with the IdentifierConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierConstraint

`func (o *EabRenewResponse) SetIdentifierConstraint(v string)`

SetIdentifierConstraint sets IdentifierConstraint field to given value.

### HasIdentifierConstraint

`func (o *EabRenewResponse) HasIdentifierConstraint() bool`

HasIdentifierConstraint returns a boolean if a field has been set.

### GetMacKey

`func (o *EabRenewResponse) GetMacKey() string`

GetMacKey returns the MacKey field if non-nil, zero value otherwise.

### GetMacKeyOk

`func (o *EabRenewResponse) GetMacKeyOk() (*string, bool)`

GetMacKeyOk returns a tuple with the MacKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacKey

`func (o *EabRenewResponse) SetMacKey(v string)`

SetMacKey sets MacKey field to given value.


### GetMacKeyAlgorithm

`func (o *EabRenewResponse) GetMacKeyAlgorithm() ExternalAccountBindingAlgorithm`

GetMacKeyAlgorithm returns the MacKeyAlgorithm field if non-nil, zero value otherwise.

### GetMacKeyAlgorithmOk

`func (o *EabRenewResponse) GetMacKeyAlgorithmOk() (*ExternalAccountBindingAlgorithm, bool)`

GetMacKeyAlgorithmOk returns a tuple with the MacKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacKeyAlgorithm

`func (o *EabRenewResponse) SetMacKeyAlgorithm(v ExternalAccountBindingAlgorithm)`

SetMacKeyAlgorithm sets MacKeyAlgorithm field to given value.


### GetMacKeyId

`func (o *EabRenewResponse) GetMacKeyId() string`

GetMacKeyId returns the MacKeyId field if non-nil, zero value otherwise.

### GetMacKeyIdOk

`func (o *EabRenewResponse) GetMacKeyIdOk() (*string, bool)`

GetMacKeyIdOk returns a tuple with the MacKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacKeyId

`func (o *EabRenewResponse) SetMacKeyId(v string)`

SetMacKeyId sets MacKeyId field to given value.


### GetName

`func (o *EabRenewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EabRenewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EabRenewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNumberOfKeyRegeneration

`func (o *EabRenewResponse) GetNumberOfKeyRegeneration() int64`

GetNumberOfKeyRegeneration returns the NumberOfKeyRegeneration field if non-nil, zero value otherwise.

### GetNumberOfKeyRegenerationOk

`func (o *EabRenewResponse) GetNumberOfKeyRegenerationOk() (*int64, bool)`

GetNumberOfKeyRegenerationOk returns a tuple with the NumberOfKeyRegeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfKeyRegeneration

`func (o *EabRenewResponse) SetNumberOfKeyRegeneration(v int64)`

SetNumberOfKeyRegeneration sets NumberOfKeyRegeneration field to given value.


### GetStatus

`func (o *EabRenewResponse) GetStatus() ExternalAccountBindingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EabRenewResponse) GetStatusOk() (*ExternalAccountBindingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EabRenewResponse) SetStatus(v ExternalAccountBindingStatus)`

SetStatus sets Status field to given value.


### GetValidationMethods

`func (o *EabRenewResponse) GetValidationMethods() []AcmeAuthorizationType`

GetValidationMethods returns the ValidationMethods field if non-nil, zero value otherwise.

### GetValidationMethodsOk

`func (o *EabRenewResponse) GetValidationMethodsOk() (*[]AcmeAuthorizationType, bool)`

GetValidationMethodsOk returns a tuple with the ValidationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethods

`func (o *EabRenewResponse) SetValidationMethods(v []AcmeAuthorizationType)`

SetValidationMethods sets ValidationMethods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


