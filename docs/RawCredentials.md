# RawCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | [**SecretString**](SecretString.md) | These credentials secret | 
**Type** | **string** |  | 
**Description** | Pointer to **NullableString** | These credentials description | [optional] 
**Expires** | Pointer to **NullableInt64** | The expiration date of these credentials | [optional] 
**Name** | **string** | These credentials identifying name | 
**Targets** | Pointer to **[]string** | On which configuration the credentials are usable | [optional] 
**Triggers** | Pointer to [**CredentialsTriggers**](CredentialsTriggers.md) |  | [optional] 

## Methods

### NewRawCredentials

`func NewRawCredentials(secret SecretString, type_ string, name string, ) *RawCredentials`

NewRawCredentials instantiates a new RawCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawCredentialsWithDefaults

`func NewRawCredentialsWithDefaults() *RawCredentials`

NewRawCredentialsWithDefaults instantiates a new RawCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *RawCredentials) GetSecret() SecretString`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RawCredentials) GetSecretOk() (*SecretString, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RawCredentials) SetSecret(v SecretString)`

SetSecret sets Secret field to given value.


### GetType

`func (o *RawCredentials) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RawCredentials) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RawCredentials) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *RawCredentials) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RawCredentials) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RawCredentials) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RawCredentials) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RawCredentials) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RawCredentials) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpires

`func (o *RawCredentials) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *RawCredentials) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *RawCredentials) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *RawCredentials) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *RawCredentials) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *RawCredentials) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetName

`func (o *RawCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RawCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RawCredentials) SetName(v string)`

SetName sets Name field to given value.


### GetTargets

`func (o *RawCredentials) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *RawCredentials) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *RawCredentials) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *RawCredentials) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTriggers

`func (o *RawCredentials) GetTriggers() CredentialsTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *RawCredentials) GetTriggersOk() (*CredentialsTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *RawCredentials) SetTriggers(v CredentialsTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *RawCredentials) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


