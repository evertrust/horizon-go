# LocalIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The identifier of the local identity (used by the identity to log in to the web UI) | 
**Email** | Pointer to **NullableString** | The email address of the local identity | [optional] 
**Name** | Pointer to **NullableString** | The display name of the local identity | [optional] 

## Methods

### NewLocalIdentity

`func NewLocalIdentity(identifier string, ) *LocalIdentity`

NewLocalIdentity instantiates a new LocalIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityWithDefaults

`func NewLocalIdentityWithDefaults() *LocalIdentity`

NewLocalIdentityWithDefaults instantiates a new LocalIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *LocalIdentity) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LocalIdentity) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *LocalIdentity) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetEmail

`func (o *LocalIdentity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocalIdentity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocalIdentity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocalIdentity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LocalIdentity) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LocalIdentity) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *LocalIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocalIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LocalIdentity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LocalIdentity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


