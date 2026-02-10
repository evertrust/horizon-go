# ResetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Local identity identifier | 
**Uuid** | **string** | The reset UUID received by email by the user after a password reset request | 
**Password** | **string** | The new password to set. It must match the password policy if any has been defined | 

## Methods

### NewResetPasswordRequest

`func NewResetPasswordRequest(identifier string, uuid string, password string, ) *ResetPasswordRequest`

NewResetPasswordRequest instantiates a new ResetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordRequestWithDefaults

`func NewResetPasswordRequestWithDefaults() *ResetPasswordRequest`

NewResetPasswordRequestWithDefaults instantiates a new ResetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ResetPasswordRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ResetPasswordRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ResetPasswordRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetUuid

`func (o *ResetPasswordRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ResetPasswordRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ResetPasswordRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetPassword

`func (o *ResetPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ResetPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ResetPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


