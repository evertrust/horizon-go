# LocalIdentityProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The internal ID of the Identity Provider | 
**Name** | **string** | The internal name of the local identity provider | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the local identity provider | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the local identity provider | [optional] 
**Type** | **string** | The type of identity provider | 
**Enabled** | **bool** | Whether the local identity provider can be used to identify against Horizon | 
**EnabledOnUI** | **bool** | Whether the local identity provider can be selected on login to the Horizon UI | 
**PasswordPolicy** | Pointer to **NullableString** | The password policy to enforce for user passwords on the local identity provider | [optional] 
**EmailTemplate** | Pointer to [**NullableEmailTemplate**](EmailTemplate.md) | The e-mail template to use for password recovery | [optional] 

## Methods

### NewLocalIdentityProviderResponse

`func NewLocalIdentityProviderResponse(id string, name string, type_ string, enabled bool, enabledOnUI bool, ) *LocalIdentityProviderResponse`

NewLocalIdentityProviderResponse instantiates a new LocalIdentityProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityProviderResponseWithDefaults

`func NewLocalIdentityProviderResponseWithDefaults() *LocalIdentityProviderResponse`

NewLocalIdentityProviderResponseWithDefaults instantiates a new LocalIdentityProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalIdentityProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalIdentityProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalIdentityProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LocalIdentityProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalIdentityProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalIdentityProviderResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *LocalIdentityProviderResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LocalIdentityProviderResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LocalIdentityProviderResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *LocalIdentityProviderResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *LocalIdentityProviderResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *LocalIdentityProviderResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *LocalIdentityProviderResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LocalIdentityProviderResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LocalIdentityProviderResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LocalIdentityProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LocalIdentityProviderResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LocalIdentityProviderResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *LocalIdentityProviderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocalIdentityProviderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocalIdentityProviderResponse) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *LocalIdentityProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LocalIdentityProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LocalIdentityProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledOnUI

`func (o *LocalIdentityProviderResponse) GetEnabledOnUI() bool`

GetEnabledOnUI returns the EnabledOnUI field if non-nil, zero value otherwise.

### GetEnabledOnUIOk

`func (o *LocalIdentityProviderResponse) GetEnabledOnUIOk() (*bool, bool)`

GetEnabledOnUIOk returns a tuple with the EnabledOnUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnUI

`func (o *LocalIdentityProviderResponse) SetEnabledOnUI(v bool)`

SetEnabledOnUI sets EnabledOnUI field to given value.


### GetPasswordPolicy

`func (o *LocalIdentityProviderResponse) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *LocalIdentityProviderResponse) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *LocalIdentityProviderResponse) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *LocalIdentityProviderResponse) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *LocalIdentityProviderResponse) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *LocalIdentityProviderResponse) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
### GetEmailTemplate

`func (o *LocalIdentityProviderResponse) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *LocalIdentityProviderResponse) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *LocalIdentityProviderResponse) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *LocalIdentityProviderResponse) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### SetEmailTemplateNil

`func (o *LocalIdentityProviderResponse) SetEmailTemplateNil(b bool)`

 SetEmailTemplateNil sets the value for EmailTemplate to be an explicit nil

### UnsetEmailTemplate
`func (o *LocalIdentityProviderResponse) UnsetEmailTemplate()`

UnsetEmailTemplate ensures that no value is present for EmailTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


