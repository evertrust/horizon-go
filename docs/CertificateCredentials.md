# CertificateCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Store** | [**SecretStoreRequest**](SecretStoreRequest.md) | These credentials certificate | 
**Name** | **string** | These credentials identifying name | 
**Description** | Pointer to **NullableString** | These credentials description | [optional] 
**Expires** | Pointer to **NullableInt64** | The expiration date of these credentials | [optional] 
**Triggers** | Pointer to [**CredentialsTriggers**](CredentialsTriggers.md) |  | [optional] 
**Targets** | Pointer to **[]string** | On which configuration the credentials are usable | [optional] 

## Methods

### NewCertificateCredentials

`func NewCertificateCredentials(type_ string, store SecretStoreRequest, name string, ) *CertificateCredentials`

NewCertificateCredentials instantiates a new CertificateCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCredentialsWithDefaults

`func NewCertificateCredentialsWithDefaults() *CertificateCredentials`

NewCertificateCredentialsWithDefaults instantiates a new CertificateCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CertificateCredentials) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateCredentials) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateCredentials) SetType(v string)`

SetType sets Type field to given value.


### GetStore

`func (o *CertificateCredentials) GetStore() SecretStoreRequest`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CertificateCredentials) GetStoreOk() (*SecretStoreRequest, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CertificateCredentials) SetStore(v SecretStoreRequest)`

SetStore sets Store field to given value.


### GetName

`func (o *CertificateCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateCredentials) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertificateCredentials) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateCredentials) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateCredentials) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateCredentials) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateCredentials) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateCredentials) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpires

`func (o *CertificateCredentials) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CertificateCredentials) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CertificateCredentials) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CertificateCredentials) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *CertificateCredentials) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *CertificateCredentials) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetTriggers

`func (o *CertificateCredentials) GetTriggers() CredentialsTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CertificateCredentials) GetTriggersOk() (*CredentialsTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CertificateCredentials) SetTriggers(v CredentialsTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CertificateCredentials) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargets

`func (o *CertificateCredentials) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CertificateCredentials) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CertificateCredentials) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *CertificateCredentials) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


