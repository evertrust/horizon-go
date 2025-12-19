# WebRARenewRequestTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**NullableManagedCertificateProfileCryptoPolicy**](ManagedCertificateProfileCryptoPolicy.md) | Describes how certificates will be enrolled on this profile | [optional] 
**PasswordPolicy** | Pointer to [**NullablePasswordPolicy**](PasswordPolicy.md) | The password policy that will be used to generate the certificate&#39;s PKCS#12 password | [optional] 

## Methods

### NewWebRARenewRequestTemplateResponse

`func NewWebRARenewRequestTemplateResponse() *WebRARenewRequestTemplateResponse`

NewWebRARenewRequestTemplateResponse instantiates a new WebRARenewRequestTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARenewRequestTemplateResponseWithDefaults

`func NewWebRARenewRequestTemplateResponseWithDefaults() *WebRARenewRequestTemplateResponse`

NewWebRARenewRequestTemplateResponseWithDefaults instantiates a new WebRARenewRequestTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *WebRARenewRequestTemplateResponse) GetCapabilities() ManagedCertificateProfileCryptoPolicy`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *WebRARenewRequestTemplateResponse) GetCapabilitiesOk() (*ManagedCertificateProfileCryptoPolicy, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *WebRARenewRequestTemplateResponse) SetCapabilities(v ManagedCertificateProfileCryptoPolicy)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *WebRARenewRequestTemplateResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *WebRARenewRequestTemplateResponse) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *WebRARenewRequestTemplateResponse) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetPasswordPolicy

`func (o *WebRARenewRequestTemplateResponse) GetPasswordPolicy() PasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *WebRARenewRequestTemplateResponse) GetPasswordPolicyOk() (*PasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *WebRARenewRequestTemplateResponse) SetPasswordPolicy(v PasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *WebRARenewRequestTemplateResponse) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### SetPasswordPolicyNil

`func (o *WebRARenewRequestTemplateResponse) SetPasswordPolicyNil(b bool)`

 SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil

### UnsetPasswordPolicy
`func (o *WebRARenewRequestTemplateResponse) UnsetPasswordPolicy()`

UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


