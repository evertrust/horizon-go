# EnabledIdentityProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the identity provider | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The display name of the identity provider | [optional] 
**Name** | **string** | The internal name of the identity provider | 
**Reset** | **bool** | Whether the password reset option is enabled on the identity provider (only for IdentityProvider of type &#x60;Local&#x60;) | 
**Type** | **string** | The type of the identity provider | 

## Methods

### NewEnabledIdentityProviderResponse

`func NewEnabledIdentityProviderResponse(name string, reset bool, type_ string, ) *EnabledIdentityProviderResponse`

NewEnabledIdentityProviderResponse instantiates a new EnabledIdentityProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnabledIdentityProviderResponseWithDefaults

`func NewEnabledIdentityProviderResponseWithDefaults() *EnabledIdentityProviderResponse`

NewEnabledIdentityProviderResponseWithDefaults instantiates a new EnabledIdentityProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EnabledIdentityProviderResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnabledIdentityProviderResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnabledIdentityProviderResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnabledIdentityProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EnabledIdentityProviderResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnabledIdentityProviderResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *EnabledIdentityProviderResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnabledIdentityProviderResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnabledIdentityProviderResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EnabledIdentityProviderResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *EnabledIdentityProviderResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *EnabledIdentityProviderResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetName

`func (o *EnabledIdentityProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnabledIdentityProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnabledIdentityProviderResponse) SetName(v string)`

SetName sets Name field to given value.


### GetReset

`func (o *EnabledIdentityProviderResponse) GetReset() bool`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *EnabledIdentityProviderResponse) GetResetOk() (*bool, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *EnabledIdentityProviderResponse) SetReset(v bool)`

SetReset sets Reset field to given value.


### GetType

`func (o *EnabledIdentityProviderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnabledIdentityProviderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnabledIdentityProviderResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


