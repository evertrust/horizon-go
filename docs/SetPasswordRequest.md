# SetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableString** | The local identity identifier. If &#x60;null&#x60;, the password for the currently logged in user will be changed | [optional] 
**Password** | **string** | The new password in clear text | 
**PreviousPassword** | Pointer to **string** | When changing your own password, this value is required and must contain the current password in clear text | [optional] 

## Methods

### NewSetPasswordRequest

`func NewSetPasswordRequest(password string, ) *SetPasswordRequest`

NewSetPasswordRequest instantiates a new SetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPasswordRequestWithDefaults

`func NewSetPasswordRequestWithDefaults() *SetPasswordRequest`

NewSetPasswordRequestWithDefaults instantiates a new SetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *SetPasswordRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SetPasswordRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SetPasswordRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *SetPasswordRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *SetPasswordRequest) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *SetPasswordRequest) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetPassword

`func (o *SetPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SetPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SetPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPreviousPassword

`func (o *SetPasswordRequest) GetPreviousPassword() string`

GetPreviousPassword returns the PreviousPassword field if non-nil, zero value otherwise.

### GetPreviousPasswordOk

`func (o *SetPasswordRequest) GetPreviousPasswordOk() (*string, bool)`

GetPreviousPasswordOk returns a tuple with the PreviousPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPassword

`func (o *SetPasswordRequest) SetPreviousPassword(v string)`

SetPreviousPassword sets PreviousPassword field to given value.

### HasPreviousPassword

`func (o *SetPasswordRequest) HasPreviousPassword() bool`

HasPreviousPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


