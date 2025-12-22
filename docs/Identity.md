# Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **NullableString** | The principal&#39;s certificate (in case of &#x60;X509&#x60; identity provider) | [optional] 
**Email** | Pointer to **NullableString** | The principal&#39;s e-mail | [optional] 
**Identifier** | **string** | The principal&#39;s identifier | 
**IdentityProviderName** | Pointer to **NullableString** | The identity provider&#39;s name this principal is registered on | [optional] 
**IdentityProviderType** | Pointer to **NullableString** | The identity provider&#39;s type this principal is registered on | [optional] 
**Name** | Pointer to **NullableString** | The principal&#39;s name | [optional] 

## Methods

### NewIdentity

`func NewIdentity(identifier string, ) *Identity`

NewIdentity instantiates a new Identity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithDefaults

`func NewIdentityWithDefaults() *Identity`

NewIdentityWithDefaults instantiates a new Identity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *Identity) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Identity) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Identity) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *Identity) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *Identity) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *Identity) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetEmail

`func (o *Identity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Identity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Identity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Identity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Identity) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Identity) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIdentifier

`func (o *Identity) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Identity) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Identity) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetIdentityProviderName

`func (o *Identity) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *Identity) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *Identity) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.

### HasIdentityProviderName

`func (o *Identity) HasIdentityProviderName() bool`

HasIdentityProviderName returns a boolean if a field has been set.

### SetIdentityProviderNameNil

`func (o *Identity) SetIdentityProviderNameNil(b bool)`

 SetIdentityProviderNameNil sets the value for IdentityProviderName to be an explicit nil

### UnsetIdentityProviderName
`func (o *Identity) UnsetIdentityProviderName()`

UnsetIdentityProviderName ensures that no value is present for IdentityProviderName, not even an explicit nil
### GetIdentityProviderType

`func (o *Identity) GetIdentityProviderType() string`

GetIdentityProviderType returns the IdentityProviderType field if non-nil, zero value otherwise.

### GetIdentityProviderTypeOk

`func (o *Identity) GetIdentityProviderTypeOk() (*string, bool)`

GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderType

`func (o *Identity) SetIdentityProviderType(v string)`

SetIdentityProviderType sets IdentityProviderType field to given value.

### HasIdentityProviderType

`func (o *Identity) HasIdentityProviderType() bool`

HasIdentityProviderType returns a boolean if a field has been set.

### SetIdentityProviderTypeNil

`func (o *Identity) SetIdentityProviderTypeNil(b bool)`

 SetIdentityProviderTypeNil sets the value for IdentityProviderType to be an explicit nil

### UnsetIdentityProviderType
`func (o *Identity) UnsetIdentityProviderType()`

UnsetIdentityProviderType ensures that no value is present for IdentityProviderType, not even an explicit nil
### GetName

`func (o *Identity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Identity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Identity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Identity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Identity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Identity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


