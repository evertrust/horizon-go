# PasswordCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | These credentials login | 
**Password** | [**SecretString**](SecretString.md) | These credentials password | 
**Type** | **string** |  | 
**Id** | **string** | Object internal ID | 
**Description** | Pointer to **NullableString** | These credentials description | [optional] 
**Expires** | Pointer to **NullableInt64** | The expiration date of these credentials | [optional] 
**Name** | **string** | These credentials identifying name | 
**Targets** | Pointer to **[]string** | On which configuration the credentials are usable | [optional] 
**Triggers** | Pointer to [**CredentialsTriggers**](CredentialsTriggers.md) |  | [optional] 

## Methods

### NewPasswordCredentialsResponse

`func NewPasswordCredentialsResponse(login string, password SecretString, type_ string, id string, name string, ) *PasswordCredentialsResponse`

NewPasswordCredentialsResponse instantiates a new PasswordCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordCredentialsResponseWithDefaults

`func NewPasswordCredentialsResponseWithDefaults() *PasswordCredentialsResponse`

NewPasswordCredentialsResponseWithDefaults instantiates a new PasswordCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *PasswordCredentialsResponse) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PasswordCredentialsResponse) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PasswordCredentialsResponse) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *PasswordCredentialsResponse) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordCredentialsResponse) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordCredentialsResponse) SetPassword(v SecretString)`

SetPassword sets Password field to given value.


### GetType

`func (o *PasswordCredentialsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PasswordCredentialsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PasswordCredentialsResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PasswordCredentialsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PasswordCredentialsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PasswordCredentialsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *PasswordCredentialsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PasswordCredentialsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PasswordCredentialsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PasswordCredentialsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PasswordCredentialsResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PasswordCredentialsResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpires

`func (o *PasswordCredentialsResponse) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PasswordCredentialsResponse) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PasswordCredentialsResponse) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PasswordCredentialsResponse) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *PasswordCredentialsResponse) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *PasswordCredentialsResponse) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetName

`func (o *PasswordCredentialsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordCredentialsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordCredentialsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetTargets

`func (o *PasswordCredentialsResponse) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *PasswordCredentialsResponse) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *PasswordCredentialsResponse) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *PasswordCredentialsResponse) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTriggers

`func (o *PasswordCredentialsResponse) GetTriggers() CredentialsTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *PasswordCredentialsResponse) GetTriggersOk() (*CredentialsTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *PasswordCredentialsResponse) SetTriggers(v CredentialsTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *PasswordCredentialsResponse) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


