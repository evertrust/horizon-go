# SecurityCredentialsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Store** | [**SecretStoreRequest**](SecretStoreRequest.md) | These credentials certificate | 
**Type** | **string** | These credentials type | 
**Description** | Pointer to **NullableString** | These credentials description | [optional] 
**Expires** | Pointer to **NullableInt64** | The expiration date of these credentials | [optional] 
**Name** | **string** | These credentials identifying name | 
**Targets** | Pointer to **[]string** | On which configuration the credentials are usable | [optional] 
**Triggers** | Pointer to [**CredentialsTriggers**](CredentialsTriggers.md) |  | [optional] 
**Login** | **string** | These credentials login | 
**Password** | [**SecretString**](SecretString.md) | These credentials password | 
**Secret** | [**SecretString**](SecretString.md) | These credentials secret | 

## Methods

### NewSecurityCredentialsUpdateRequest

`func NewSecurityCredentialsUpdateRequest(store SecretStoreRequest, type_ string, name string, login string, password SecretString, secret SecretString, ) *SecurityCredentialsUpdateRequest`

NewSecurityCredentialsUpdateRequest instantiates a new SecurityCredentialsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityCredentialsUpdateRequestWithDefaults

`func NewSecurityCredentialsUpdateRequestWithDefaults() *SecurityCredentialsUpdateRequest`

NewSecurityCredentialsUpdateRequestWithDefaults instantiates a new SecurityCredentialsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStore

`func (o *SecurityCredentialsUpdateRequest) GetStore() SecretStoreRequest`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *SecurityCredentialsUpdateRequest) GetStoreOk() (*SecretStoreRequest, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *SecurityCredentialsUpdateRequest) SetStore(v SecretStoreRequest)`

SetStore sets Store field to given value.


### GetType

`func (o *SecurityCredentialsUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityCredentialsUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityCredentialsUpdateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *SecurityCredentialsUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityCredentialsUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityCredentialsUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityCredentialsUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityCredentialsUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityCredentialsUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpires

`func (o *SecurityCredentialsUpdateRequest) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SecurityCredentialsUpdateRequest) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SecurityCredentialsUpdateRequest) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SecurityCredentialsUpdateRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *SecurityCredentialsUpdateRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *SecurityCredentialsUpdateRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetName

`func (o *SecurityCredentialsUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityCredentialsUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityCredentialsUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTargets

`func (o *SecurityCredentialsUpdateRequest) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SecurityCredentialsUpdateRequest) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SecurityCredentialsUpdateRequest) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *SecurityCredentialsUpdateRequest) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTriggers

`func (o *SecurityCredentialsUpdateRequest) GetTriggers() CredentialsTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *SecurityCredentialsUpdateRequest) GetTriggersOk() (*CredentialsTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *SecurityCredentialsUpdateRequest) SetTriggers(v CredentialsTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *SecurityCredentialsUpdateRequest) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetLogin

`func (o *SecurityCredentialsUpdateRequest) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *SecurityCredentialsUpdateRequest) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *SecurityCredentialsUpdateRequest) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *SecurityCredentialsUpdateRequest) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SecurityCredentialsUpdateRequest) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SecurityCredentialsUpdateRequest) SetPassword(v SecretString)`

SetPassword sets Password field to given value.


### GetSecret

`func (o *SecurityCredentialsUpdateRequest) GetSecret() SecretString`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SecurityCredentialsUpdateRequest) GetSecretOk() (*SecretString, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SecurityCredentialsUpdateRequest) SetSecret(v SecretString)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


