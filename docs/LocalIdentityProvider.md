# LocalIdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The internal name of the local identity provider | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the local identity provider | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the local identity provider | [optional] 
**Type** | **string** | The type of identity provider | 
**Enabled** | **bool** | Whether the local identity provider can be used to identify against Horizon | 
**EnabledOnUI** | **bool** | Whether the local identity provider can be selected on login to the Horizon UI | 
**PasswordPolicy** | Pointer to **NullableString** | The password policy to enforce for user passwords on the local identity provider | [optional] 
**EmailTemplate** | Pointer to [**NullableEmailTemplate**](EmailTemplate.md) | The e-mail template to use for password recovery | [optional] 

## Methods

### NewLocalIdentityProvider

`func NewLocalIdentityProvider(name string, type_ string, enabled bool, enabledOnUI bool, ) *LocalIdentityProvider`

NewLocalIdentityProvider instantiates a new LocalIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityProviderWithDefaults

`func NewLocalIdentityProviderWithDefaults() *LocalIdentityProvider`

NewLocalIdentityProviderWithDefaults instantiates a new LocalIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LocalIdentityProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalIdentityProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalIdentityProvider) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *LocalIdentityProvider) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LocalIdentityProvider) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LocalIdentityProvider) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *LocalIdentityProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *LocalIdentityProvider) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *LocalIdentityProvider) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *LocalIdentityProvider) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LocalIdentityProvider) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LocalIdentityProvider) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LocalIdentityProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LocalIdentityProvider) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LocalIdentityProvider) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *LocalIdentityProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocalIdentityProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocalIdentityProvider) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *LocalIdentityProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LocalIdentityProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LocalIdentityProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledOnUI

`func (o *LocalIdentityProvider) GetEnabledOnUI() bool`

GetEnabledOnUI returns the EnabledOnUI field if non-nil, zero value otherwise.

### GetEnabledOnUIOk

`func (o *LocalIdentityProvider) GetEnabledOnUIOk() (*bool, bool)`

GetEnabledOnUIOk returns a tuple with the EnabledOnUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnUI

`func (o *LocalIdentityProvider) SetEnabledOnUI(v bool)`

SetEnabledOnUI sets EnabledOnUI field to given value.


### GetPasswordPolicy

`func (o *LocalIdentityProvider) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *LocalIdentityProvider) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *LocalIdentityProvider) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *LocalIdentityProvider) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *LocalIdentityProvider) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *LocalIdentityProvider) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetEmailTemplate

`func (o *LocalIdentityProvider) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *LocalIdentityProvider) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *LocalIdentityProvider) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *LocalIdentityProvider) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### SetEmailTemplateNil

`func (o *LocalIdentityProvider) SetEmailTemplateNil(b bool)`

 SetEmailTemplateNil sets the value for EmailTemplate to be an explicit nil

### UnsetEmailTemplate
`func (o *LocalIdentityProvider) UnsetEmailTemplate()`

UnsetEmailTemplate ensures that no value is present for EmailTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


