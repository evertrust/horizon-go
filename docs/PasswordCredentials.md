# PasswordCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Login** | **string** | These credentials login | 
**Password** | [**SecretString**](SecretString.md) | These credentials password | 
**Name** | **string** | These credentials identifying name | 
**Description** | Pointer to **NullableString** | These credentials description | [optional] 
**Expires** | Pointer to **NullableInt64** | The expiration date of these credentials | [optional] 
**Triggers** | Pointer to [**CredentialsTriggers**](CredentialsTriggers.md) |  | [optional] 
**Targets** | Pointer to **[]string** | On which configuration the credentials are usable | [optional] 

## Methods

### NewPasswordCredentials

`func NewPasswordCredentials(type_ string, login string, password SecretString, name string, ) *PasswordCredentials`

NewPasswordCredentials instantiates a new PasswordCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordCredentialsWithDefaults

`func NewPasswordCredentialsWithDefaults() *PasswordCredentials`

NewPasswordCredentialsWithDefaults instantiates a new PasswordCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PasswordCredentials) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PasswordCredentials) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PasswordCredentials) SetType(v string)`

SetType sets Type field to given value.


### GetLogin

`func (o *PasswordCredentials) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PasswordCredentials) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PasswordCredentials) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *PasswordCredentials) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordCredentials) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordCredentials) SetPassword(v SecretString)`

SetPassword sets Password field to given value.


### GetName

`func (o *PasswordCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordCredentials) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PasswordCredentials) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PasswordCredentials) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PasswordCredentials) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PasswordCredentials) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PasswordCredentials) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PasswordCredentials) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpires

`func (o *PasswordCredentials) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PasswordCredentials) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PasswordCredentials) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PasswordCredentials) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *PasswordCredentials) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *PasswordCredentials) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetTriggers

`func (o *PasswordCredentials) GetTriggers() CredentialsTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *PasswordCredentials) GetTriggersOk() (*CredentialsTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *PasswordCredentials) SetTriggers(v CredentialsTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *PasswordCredentials) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargets

`func (o *PasswordCredentials) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *PasswordCredentials) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *PasswordCredentials) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *PasswordCredentials) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


