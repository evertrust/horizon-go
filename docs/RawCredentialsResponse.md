# RawCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** |  | 
**Secret** | [**SecretString**](SecretString.md) | These credentials secret | 
**Name** | **string** | These credentials identifying name | 
**Description** | Pointer to **NullableString** | These credentials description | [optional] 
**Expires** | Pointer to **NullableInt64** | The expiration date of these credentials | [optional] 
**Triggers** | Pointer to [**CredentialsTriggers**](CredentialsTriggers.md) |  | [optional] 
**Targets** | Pointer to **[]string** | On which configuration the credentials are usable | [optional] 

## Methods

### NewRawCredentialsResponse

`func NewRawCredentialsResponse(id string, type_ string, secret SecretString, name string, ) *RawCredentialsResponse`

NewRawCredentialsResponse instantiates a new RawCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawCredentialsResponseWithDefaults

`func NewRawCredentialsResponseWithDefaults() *RawCredentialsResponse`

NewRawCredentialsResponseWithDefaults instantiates a new RawCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RawCredentialsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RawCredentialsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RawCredentialsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RawCredentialsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RawCredentialsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RawCredentialsResponse) SetType(v string)`

SetType sets Type field to given value.


### GetSecret

`func (o *RawCredentialsResponse) GetSecret() SecretString`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RawCredentialsResponse) GetSecretOk() (*SecretString, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RawCredentialsResponse) SetSecret(v SecretString)`

SetSecret sets Secret field to given value.


### GetName

`func (o *RawCredentialsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RawCredentialsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RawCredentialsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RawCredentialsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RawCredentialsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RawCredentialsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RawCredentialsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RawCredentialsResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RawCredentialsResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpires

`func (o *RawCredentialsResponse) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *RawCredentialsResponse) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *RawCredentialsResponse) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *RawCredentialsResponse) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *RawCredentialsResponse) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *RawCredentialsResponse) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetTriggers

`func (o *RawCredentialsResponse) GetTriggers() CredentialsTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *RawCredentialsResponse) GetTriggersOk() (*CredentialsTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *RawCredentialsResponse) SetTriggers(v CredentialsTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *RawCredentialsResponse) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargets

`func (o *RawCredentialsResponse) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *RawCredentialsResponse) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *RawCredentialsResponse) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *RawCredentialsResponse) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


