# WebRARecoverRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordMode** | Pointer to **NullableString** | The password mode of the certificate | [optional] 
**PasswordPolicy** | Pointer to [**NullablePasswordPolicy**](PasswordPolicy.md) | The selected password policy for this profile. If none is defined and the password mode is &#x60;manual&#x60;, there is no constraint on the password. In &#x60;random&#x60; mode, the &#x60;Horizon-Default&#x60; policy is used | [optional] 

## Methods

### NewWebRARecoverRequestTemplate

`func NewWebRARecoverRequestTemplate() *WebRARecoverRequestTemplate`

NewWebRARecoverRequestTemplate instantiates a new WebRARecoverRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARecoverRequestTemplateWithDefaults

`func NewWebRARecoverRequestTemplateWithDefaults() *WebRARecoverRequestTemplate`

NewWebRARecoverRequestTemplateWithDefaults instantiates a new WebRARecoverRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordMode

`func (o *WebRARecoverRequestTemplate) GetPasswordMode() string`

GetPasswordMode returns the PasswordMode field if non-nil, zero value otherwise.

### GetPasswordModeOk

`func (o *WebRARecoverRequestTemplate) GetPasswordModeOk() (*string, bool)`

GetPasswordModeOk returns a tuple with the PasswordMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMode

`func (o *WebRARecoverRequestTemplate) SetPasswordMode(v string)`

SetPasswordMode sets PasswordMode field to given value.

### HasPasswordMode

`func (o *WebRARecoverRequestTemplate) HasPasswordMode() bool`

HasPasswordMode returns a boolean if a field has been set.

### SetPasswordModeNil

`func (o *WebRARecoverRequestTemplate) SetPasswordModeNil(b bool)`

 SetPasswordModeNil sets the value for PasswordMode to be an explicit nil

### UnsetPasswordMode
`func (o *WebRARecoverRequestTemplate) UnsetPasswordMode()`

UnsetPasswordMode ensures that no value is present for PasswordMode, not even an explicit nil
### GetPasswordPolicy

`func (o *WebRARecoverRequestTemplate) GetPasswordPolicy() PasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *WebRARecoverRequestTemplate) GetPasswordPolicyOk() (*PasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *WebRARecoverRequestTemplate) SetPasswordPolicy(v PasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *WebRARecoverRequestTemplate) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *WebRARecoverRequestTemplate) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *WebRARecoverRequestTemplate) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


