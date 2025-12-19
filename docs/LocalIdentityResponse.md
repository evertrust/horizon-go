# LocalIdentityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal ID | 
**Email** | Pointer to **NullableString** | The email address of the local identity | [optional] 
**Identifier** | **string** | The identifier of the local identity (used by the identity to log in to the web UI) | 
**Name** | Pointer to **NullableString** | The display name of the local identity | [optional] 

## Methods

### NewLocalIdentityResponse

`func NewLocalIdentityResponse(id string, identifier string, ) *LocalIdentityResponse`

NewLocalIdentityResponse instantiates a new LocalIdentityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityResponseWithDefaults

`func NewLocalIdentityResponseWithDefaults() *LocalIdentityResponse`

NewLocalIdentityResponseWithDefaults instantiates a new LocalIdentityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalIdentityResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalIdentityResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalIdentityResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *LocalIdentityResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocalIdentityResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocalIdentityResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocalIdentityResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LocalIdentityResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LocalIdentityResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIdentifier

`func (o *LocalIdentityResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LocalIdentityResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *LocalIdentityResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *LocalIdentityResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalIdentityResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalIdentityResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocalIdentityResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LocalIdentityResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LocalIdentityResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


