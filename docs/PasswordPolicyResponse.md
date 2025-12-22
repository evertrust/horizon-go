# PasswordPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The internal ID of the password policy | 
**MaxChar** | Pointer to **NullableInt64** | The maximum number of characters of the password | [optional] 
**MinChar** | **int64** | The minimum number of characters of the password | 
**MinDiChar** | Pointer to **NullableInt64** | The minimum number of digits of the password | [optional] 
**MinLoChar** | Pointer to **NullableInt64** | The minimum number of lowercase characters of the password | [optional] 
**MinSpChar** | Pointer to **NullableInt64** | The minimum number of special characters of the password | [optional] 
**MinUpChar** | Pointer to **NullableInt64** | The minimum number of uppercase characters of the password | [optional] 
**Name** | **string** | The name of the password policy | 
**SpChar** | Pointer to **NullableString** | The special characters of the password accepted by the password policy | [optional] 

## Methods

### NewPasswordPolicyResponse

`func NewPasswordPolicyResponse(id string, minChar int64, name string, ) *PasswordPolicyResponse`

NewPasswordPolicyResponse instantiates a new PasswordPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyResponseWithDefaults

`func NewPasswordPolicyResponseWithDefaults() *PasswordPolicyResponse`

NewPasswordPolicyResponseWithDefaults instantiates a new PasswordPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PasswordPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PasswordPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PasswordPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMaxChar

`func (o *PasswordPolicyResponse) GetMaxChar() int64`

GetMaxChar returns the MaxChar field if non-nil, zero value otherwise.

### GetMaxCharOk

`func (o *PasswordPolicyResponse) GetMaxCharOk() (*int64, bool)`

GetMaxCharOk returns a tuple with the MaxChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChar

`func (o *PasswordPolicyResponse) SetMaxChar(v int64)`

SetMaxChar sets MaxChar field to given value.

### HasMaxChar

`func (o *PasswordPolicyResponse) HasMaxChar() bool`

HasMaxChar returns a boolean if a field has been set.

### SetMaxCharNil

`func (o *PasswordPolicyResponse) SetMaxCharNil(b bool)`

 SetMaxCharNil sets the value for MaxChar to be an explicit nil

### UnsetMaxChar
`func (o *PasswordPolicyResponse) UnsetMaxChar()`

UnsetMaxChar ensures that no value is present for MaxChar, not even an explicit nil
### GetMinChar

`func (o *PasswordPolicyResponse) GetMinChar() int64`

GetMinChar returns the MinChar field if non-nil, zero value otherwise.

### GetMinCharOk

`func (o *PasswordPolicyResponse) GetMinCharOk() (*int64, bool)`

GetMinCharOk returns a tuple with the MinChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinChar

`func (o *PasswordPolicyResponse) SetMinChar(v int64)`

SetMinChar sets MinChar field to given value.


### GetMinDiChar

`func (o *PasswordPolicyResponse) GetMinDiChar() int64`

GetMinDiChar returns the MinDiChar field if non-nil, zero value otherwise.

### GetMinDiCharOk

`func (o *PasswordPolicyResponse) GetMinDiCharOk() (*int64, bool)`

GetMinDiCharOk returns a tuple with the MinDiChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiChar

`func (o *PasswordPolicyResponse) SetMinDiChar(v int64)`

SetMinDiChar sets MinDiChar field to given value.

### HasMinDiChar

`func (o *PasswordPolicyResponse) HasMinDiChar() bool`

HasMinDiChar returns a boolean if a field has been set.

### SetMinDiCharNil

`func (o *PasswordPolicyResponse) SetMinDiCharNil(b bool)`

 SetMinDiCharNil sets the value for MinDiChar to be an explicit nil

### UnsetMinDiChar
`func (o *PasswordPolicyResponse) UnsetMinDiChar()`

UnsetMinDiChar ensures that no value is present for MinDiChar, not even an explicit nil
### GetMinLoChar

`func (o *PasswordPolicyResponse) GetMinLoChar() int64`

GetMinLoChar returns the MinLoChar field if non-nil, zero value otherwise.

### GetMinLoCharOk

`func (o *PasswordPolicyResponse) GetMinLoCharOk() (*int64, bool)`

GetMinLoCharOk returns a tuple with the MinLoChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLoChar

`func (o *PasswordPolicyResponse) SetMinLoChar(v int64)`

SetMinLoChar sets MinLoChar field to given value.

### HasMinLoChar

`func (o *PasswordPolicyResponse) HasMinLoChar() bool`

HasMinLoChar returns a boolean if a field has been set.

### SetMinLoCharNil

`func (o *PasswordPolicyResponse) SetMinLoCharNil(b bool)`

 SetMinLoCharNil sets the value for MinLoChar to be an explicit nil

### UnsetMinLoChar
`func (o *PasswordPolicyResponse) UnsetMinLoChar()`

UnsetMinLoChar ensures that no value is present for MinLoChar, not even an explicit nil
### GetMinSpChar

`func (o *PasswordPolicyResponse) GetMinSpChar() int64`

GetMinSpChar returns the MinSpChar field if non-nil, zero value otherwise.

### GetMinSpCharOk

`func (o *PasswordPolicyResponse) GetMinSpCharOk() (*int64, bool)`

GetMinSpCharOk returns a tuple with the MinSpChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSpChar

`func (o *PasswordPolicyResponse) SetMinSpChar(v int64)`

SetMinSpChar sets MinSpChar field to given value.

### HasMinSpChar

`func (o *PasswordPolicyResponse) HasMinSpChar() bool`

HasMinSpChar returns a boolean if a field has been set.

### SetMinSpCharNil

`func (o *PasswordPolicyResponse) SetMinSpCharNil(b bool)`

 SetMinSpCharNil sets the value for MinSpChar to be an explicit nil

### UnsetMinSpChar
`func (o *PasswordPolicyResponse) UnsetMinSpChar()`

UnsetMinSpChar ensures that no value is present for MinSpChar, not even an explicit nil
### GetMinUpChar

`func (o *PasswordPolicyResponse) GetMinUpChar() int64`

GetMinUpChar returns the MinUpChar field if non-nil, zero value otherwise.

### GetMinUpCharOk

`func (o *PasswordPolicyResponse) GetMinUpCharOk() (*int64, bool)`

GetMinUpCharOk returns a tuple with the MinUpChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpChar

`func (o *PasswordPolicyResponse) SetMinUpChar(v int64)`

SetMinUpChar sets MinUpChar field to given value.

### HasMinUpChar

`func (o *PasswordPolicyResponse) HasMinUpChar() bool`

HasMinUpChar returns a boolean if a field has been set.

### SetMinUpCharNil

`func (o *PasswordPolicyResponse) SetMinUpCharNil(b bool)`

 SetMinUpCharNil sets the value for MinUpChar to be an explicit nil

### UnsetMinUpChar
`func (o *PasswordPolicyResponse) UnsetMinUpChar()`

UnsetMinUpChar ensures that no value is present for MinUpChar, not even an explicit nil
### GetName

`func (o *PasswordPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordPolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSpChar

`func (o *PasswordPolicyResponse) GetSpChar() string`

GetSpChar returns the SpChar field if non-nil, zero value otherwise.

### GetSpCharOk

`func (o *PasswordPolicyResponse) GetSpCharOk() (*string, bool)`

GetSpCharOk returns a tuple with the SpChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpChar

`func (o *PasswordPolicyResponse) SetSpChar(v string)`

SetSpChar sets SpChar field to given value.

### HasSpChar

`func (o *PasswordPolicyResponse) HasSpChar() bool`

HasSpChar returns a boolean if a field has been set.

### SetSpCharNil

`func (o *PasswordPolicyResponse) SetSpCharNil(b bool)`

 SetSpCharNil sets the value for SpChar to be an explicit nil

### UnsetSpChar
`func (o *PasswordPolicyResponse) UnsetSpChar()`

UnsetSpChar ensures that no value is present for SpChar, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


