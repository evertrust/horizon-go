# SecurityCredentialsList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | **NullableInt64** | The expiration date of these credentials | 
**Type** | **string** | These credentials type | 
**Store** | [**SecretStoreResponse**](SecretStoreResponse.md) | These credentials certificate | 
**Name** | **string** | These credentials identifying name | 
**Description** | Pointer to **NullableString** | These credentials description | [optional] 
**Triggers** | Pointer to [**CredentialsTriggers**](CredentialsTriggers.md) |  | [optional] 
**Targets** | Pointer to **[]string** | On which configuration the credentials are usable | [optional] 
**Login** | **string** | These credentials login | 
**Password** | [**SecretString**](SecretString.md) | These credentials password | 
**Secret** | [**SecretString**](SecretString.md) | These credentials secret | 

## Methods

### NewSecurityCredentialsList200ResponseInner

`func NewSecurityCredentialsList200ResponseInner(expires NullableInt64, type_ string, store SecretStoreResponse, name string, login string, password SecretString, secret SecretString, ) *SecurityCredentialsList200ResponseInner`

NewSecurityCredentialsList200ResponseInner instantiates a new SecurityCredentialsList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityCredentialsList200ResponseInnerWithDefaults

`func NewSecurityCredentialsList200ResponseInnerWithDefaults() *SecurityCredentialsList200ResponseInner`

NewSecurityCredentialsList200ResponseInnerWithDefaults instantiates a new SecurityCredentialsList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *SecurityCredentialsList200ResponseInner) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SecurityCredentialsList200ResponseInner) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SecurityCredentialsList200ResponseInner) SetExpires(v int64)`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *SecurityCredentialsList200ResponseInner) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *SecurityCredentialsList200ResponseInner) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetType

`func (o *SecurityCredentialsList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityCredentialsList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityCredentialsList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetStore

`func (o *SecurityCredentialsList200ResponseInner) GetStore() SecretStoreResponse`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *SecurityCredentialsList200ResponseInner) GetStoreOk() (*SecretStoreResponse, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *SecurityCredentialsList200ResponseInner) SetStore(v SecretStoreResponse)`

SetStore sets Store field to given value.


### GetName

`func (o *SecurityCredentialsList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityCredentialsList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityCredentialsList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SecurityCredentialsList200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityCredentialsList200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityCredentialsList200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityCredentialsList200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityCredentialsList200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityCredentialsList200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTriggers

`func (o *SecurityCredentialsList200ResponseInner) GetTriggers() CredentialsTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *SecurityCredentialsList200ResponseInner) GetTriggersOk() (*CredentialsTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *SecurityCredentialsList200ResponseInner) SetTriggers(v CredentialsTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *SecurityCredentialsList200ResponseInner) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargets

`func (o *SecurityCredentialsList200ResponseInner) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SecurityCredentialsList200ResponseInner) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SecurityCredentialsList200ResponseInner) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *SecurityCredentialsList200ResponseInner) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetLogin

`func (o *SecurityCredentialsList200ResponseInner) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *SecurityCredentialsList200ResponseInner) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *SecurityCredentialsList200ResponseInner) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *SecurityCredentialsList200ResponseInner) GetPassword() SecretString`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SecurityCredentialsList200ResponseInner) GetPasswordOk() (*SecretString, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SecurityCredentialsList200ResponseInner) SetPassword(v SecretString)`

SetPassword sets Password field to given value.


### GetSecret

`func (o *SecurityCredentialsList200ResponseInner) GetSecret() SecretString`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SecurityCredentialsList200ResponseInner) GetSecretOk() (*SecretString, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SecurityCredentialsList200ResponseInner) SetSecret(v SecretString)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


