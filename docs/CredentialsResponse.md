# CredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Object internal ID | [optional] 
**Name** | Pointer to **string** | These credentials identifying name | [optional] 
**Type** | Pointer to **string** | These credentials type | [optional] 
**Description** | Pointer to **NullableString** | These credentials description | [optional] 
**Expires** | Pointer to **NullableInt64** | The expiration date of these credentials | [optional] 
**Triggers** | Pointer to [**CredentialsTriggers**](CredentialsTriggers.md) |  | [optional] 
**Targets** | Pointer to **[]string** | On which configuration the credentials are usable | [optional] 

## Methods

### NewCredentialsResponse

`func NewCredentialsResponse() *CredentialsResponse`

NewCredentialsResponse instantiates a new CredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsResponseWithDefaults

`func NewCredentialsResponseWithDefaults() *CredentialsResponse`

NewCredentialsResponseWithDefaults instantiates a new CredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CredentialsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CredentialsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CredentialsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CredentialsResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *CredentialsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CredentialsResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CredentialsResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpires

`func (o *CredentialsResponse) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CredentialsResponse) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CredentialsResponse) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CredentialsResponse) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *CredentialsResponse) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *CredentialsResponse) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetTriggers

`func (o *CredentialsResponse) GetTriggers() CredentialsTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CredentialsResponse) GetTriggersOk() (*CredentialsTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CredentialsResponse) SetTriggers(v CredentialsTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CredentialsResponse) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargets

`func (o *CredentialsResponse) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CredentialsResponse) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CredentialsResponse) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *CredentialsResponse) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


