# PasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the password policy | 
**MinChar** | **int64** | The minimum number of characters of the password | 
**MaxChar** | Pointer to **NullableInt64** | The maximum number of characters of the password | [optional] 
**MinUpChar** | Pointer to **NullableInt64** | The minimum number of uppercase characters of the password | [optional] 
**MinLoChar** | Pointer to **NullableInt64** | The minimum number of lowercase characters of the password | [optional] 
**MinDiChar** | Pointer to **NullableInt64** | The minimum number of digits of the password | [optional] 
**SpChar** | Pointer to **NullableString** | The special characters of the password accepted by the password policy | [optional] 
**MinSpChar** | Pointer to **NullableInt64** | The minimum number of special characters of the password | [optional] 

## Methods

### NewPasswordPolicy

`func NewPasswordPolicy(name string, minChar int64, ) *PasswordPolicy`

NewPasswordPolicy instantiates a new PasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyWithDefaults

`func NewPasswordPolicyWithDefaults() *PasswordPolicy`

NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PasswordPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetMinChar

`func (o *PasswordPolicy) GetMinChar() int64`

GetMinChar returns the MinChar field if non-nil, zero value otherwise.

### GetMinCharOk

`func (o *PasswordPolicy) GetMinCharOk() (*int64, bool)`

GetMinCharOk returns a tuple with the MinChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinChar

`func (o *PasswordPolicy) SetMinChar(v int64)`

SetMinChar sets MinChar field to given value.


### GetMaxChar

`func (o *PasswordPolicy) GetMaxChar() int64`

GetMaxChar returns the MaxChar field if non-nil, zero value otherwise.

### GetMaxCharOk

`func (o *PasswordPolicy) GetMaxCharOk() (*int64, bool)`

GetMaxCharOk returns a tuple with the MaxChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChar

`func (o *PasswordPolicy) SetMaxChar(v int64)`

SetMaxChar sets MaxChar field to given value.

### HasMaxChar

`func (o *PasswordPolicy) HasMaxChar() bool`

HasMaxChar returns a boolean if a field has been set.

### SetMaxCharNil

`func (o *PasswordPolicy) SetMaxCharNil(b bool)`

 SetMaxCharNil sets the value for MaxChar to be an explicit nil

### UnsetMaxChar
`func (o *PasswordPolicy) UnsetMaxChar()`

UnsetMaxChar ensures that no value is present for MaxChar, not even an explicit nil
### GetMinUpChar

`func (o *PasswordPolicy) GetMinUpChar() int64`

GetMinUpChar returns the MinUpChar field if non-nil, zero value otherwise.

### GetMinUpCharOk

`func (o *PasswordPolicy) GetMinUpCharOk() (*int64, bool)`

GetMinUpCharOk returns a tuple with the MinUpChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpChar

`func (o *PasswordPolicy) SetMinUpChar(v int64)`

SetMinUpChar sets MinUpChar field to given value.

### HasMinUpChar

`func (o *PasswordPolicy) HasMinUpChar() bool`

HasMinUpChar returns a boolean if a field has been set.

### SetMinUpCharNil

`func (o *PasswordPolicy) SetMinUpCharNil(b bool)`

 SetMinUpCharNil sets the value for MinUpChar to be an explicit nil

### UnsetMinUpChar
`func (o *PasswordPolicy) UnsetMinUpChar()`

UnsetMinUpChar ensures that no value is present for MinUpChar, not even an explicit nil
### GetMinLoChar

`func (o *PasswordPolicy) GetMinLoChar() int64`

GetMinLoChar returns the MinLoChar field if non-nil, zero value otherwise.

### GetMinLoCharOk

`func (o *PasswordPolicy) GetMinLoCharOk() (*int64, bool)`

GetMinLoCharOk returns a tuple with the MinLoChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLoChar

`func (o *PasswordPolicy) SetMinLoChar(v int64)`

SetMinLoChar sets MinLoChar field to given value.

### HasMinLoChar

`func (o *PasswordPolicy) HasMinLoChar() bool`

HasMinLoChar returns a boolean if a field has been set.

### SetMinLoCharNil

`func (o *PasswordPolicy) SetMinLoCharNil(b bool)`

 SetMinLoCharNil sets the value for MinLoChar to be an explicit nil

### UnsetMinLoChar
`func (o *PasswordPolicy) UnsetMinLoChar()`

UnsetMinLoChar ensures that no value is present for MinLoChar, not even an explicit nil
### GetMinDiChar

`func (o *PasswordPolicy) GetMinDiChar() int64`

GetMinDiChar returns the MinDiChar field if non-nil, zero value otherwise.

### GetMinDiCharOk

`func (o *PasswordPolicy) GetMinDiCharOk() (*int64, bool)`

GetMinDiCharOk returns a tuple with the MinDiChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiChar

`func (o *PasswordPolicy) SetMinDiChar(v int64)`

SetMinDiChar sets MinDiChar field to given value.

### HasMinDiChar

`func (o *PasswordPolicy) HasMinDiChar() bool`

HasMinDiChar returns a boolean if a field has been set.

### SetMinDiCharNil

`func (o *PasswordPolicy) SetMinDiCharNil(b bool)`

 SetMinDiCharNil sets the value for MinDiChar to be an explicit nil

### UnsetMinDiChar
`func (o *PasswordPolicy) UnsetMinDiChar()`

UnsetMinDiChar ensures that no value is present for MinDiChar, not even an explicit nil
### GetSpChar

`func (o *PasswordPolicy) GetSpChar() string`

GetSpChar returns the SpChar field if non-nil, zero value otherwise.

### GetSpCharOk

`func (o *PasswordPolicy) GetSpCharOk() (*string, bool)`

GetSpCharOk returns a tuple with the SpChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpChar

`func (o *PasswordPolicy) SetSpChar(v string)`

SetSpChar sets SpChar field to given value.

### HasSpChar

`func (o *PasswordPolicy) HasSpChar() bool`

HasSpChar returns a boolean if a field has been set.

### SetSpCharNil

`func (o *PasswordPolicy) SetSpCharNil(b bool)`

 SetSpCharNil sets the value for SpChar to be an explicit nil

### UnsetSpChar
`func (o *PasswordPolicy) UnsetSpChar()`

UnsetSpChar ensures that no value is present for SpChar, not even an explicit nil
### GetMinSpChar

`func (o *PasswordPolicy) GetMinSpChar() int64`

GetMinSpChar returns the MinSpChar field if non-nil, zero value otherwise.

### GetMinSpCharOk

`func (o *PasswordPolicy) GetMinSpCharOk() (*int64, bool)`

GetMinSpCharOk returns a tuple with the MinSpChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSpChar

`func (o *PasswordPolicy) SetMinSpChar(v int64)`

SetMinSpChar sets MinSpChar field to given value.

### HasMinSpChar

`func (o *PasswordPolicy) HasMinSpChar() bool`

HasMinSpChar returns a boolean if a field has been set.

### SetMinSpCharNil

`func (o *PasswordPolicy) SetMinSpCharNil(b bool)`

 SetMinSpCharNil sets the value for MinSpChar to be an explicit nil

### UnsetMinSpChar
`func (o *PasswordPolicy) UnsetMinSpChar()`

UnsetMinSpChar ensures that no value is present for MinSpChar, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


