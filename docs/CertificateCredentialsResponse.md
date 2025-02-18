# CertificateCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | **interface{}** | The expiration date of these credentials. Automatically set to the expiration date of the certificate | 
**Type** | **string** |  | 
**Store** | [**SecretStoreResponse**](SecretStoreResponse.md) | These credentials certificate | 
**Name** | **string** | These credentials identifying name | 
**Description** | Pointer to **NullableString** | These credentials description | [optional] 
**Triggers** | Pointer to [**CredentialsTriggers**](CredentialsTriggers.md) |  | [optional] 
**Targets** | Pointer to **[]string** | On which configuration the credentials are usable | [optional] 

## Methods

### NewCertificateCredentialsResponse

`func NewCertificateCredentialsResponse(expires interface{}, type_ string, store SecretStoreResponse, name string, ) *CertificateCredentialsResponse`

NewCertificateCredentialsResponse instantiates a new CertificateCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCredentialsResponseWithDefaults

`func NewCertificateCredentialsResponseWithDefaults() *CertificateCredentialsResponse`

NewCertificateCredentialsResponseWithDefaults instantiates a new CertificateCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *CertificateCredentialsResponse) GetExpires() interface{}`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CertificateCredentialsResponse) GetExpiresOk() (*interface{}, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CertificateCredentialsResponse) SetExpires(v interface{})`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *CertificateCredentialsResponse) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *CertificateCredentialsResponse) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetType

`func (o *CertificateCredentialsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateCredentialsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateCredentialsResponse) SetType(v string)`

SetType sets Type field to given value.


### GetStore

`func (o *CertificateCredentialsResponse) GetStore() SecretStoreResponse`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CertificateCredentialsResponse) GetStoreOk() (*SecretStoreResponse, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CertificateCredentialsResponse) SetStore(v SecretStoreResponse)`

SetStore sets Store field to given value.


### GetName

`func (o *CertificateCredentialsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateCredentialsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateCredentialsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertificateCredentialsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateCredentialsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateCredentialsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateCredentialsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateCredentialsResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateCredentialsResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTriggers

`func (o *CertificateCredentialsResponse) GetTriggers() CredentialsTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CertificateCredentialsResponse) GetTriggersOk() (*CredentialsTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CertificateCredentialsResponse) SetTriggers(v CredentialsTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CertificateCredentialsResponse) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargets

`func (o *CertificateCredentialsResponse) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CertificateCredentialsResponse) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CertificateCredentialsResponse) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *CertificateCredentialsResponse) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


