# EscrowedPrivateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HorizonKey** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**VaultKey** | Pointer to **NullableString** |  | [optional] 
**Transient** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewEscrowedPrivateKey

`func NewEscrowedPrivateKey(horizonKey string, ) *EscrowedPrivateKey`

NewEscrowedPrivateKey instantiates a new EscrowedPrivateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscrowedPrivateKeyWithDefaults

`func NewEscrowedPrivateKeyWithDefaults() *EscrowedPrivateKey`

NewEscrowedPrivateKeyWithDefaults instantiates a new EscrowedPrivateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHorizonKey

`func (o *EscrowedPrivateKey) GetHorizonKey() string`

GetHorizonKey returns the HorizonKey field if non-nil, zero value otherwise.

### GetHorizonKeyOk

`func (o *EscrowedPrivateKey) GetHorizonKeyOk() (*string, bool)`

GetHorizonKeyOk returns a tuple with the HorizonKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizonKey

`func (o *EscrowedPrivateKey) SetHorizonKey(v string)`

SetHorizonKey sets HorizonKey field to given value.


### GetValue

`func (o *EscrowedPrivateKey) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EscrowedPrivateKey) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EscrowedPrivateKey) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EscrowedPrivateKey) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EscrowedPrivateKey) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EscrowedPrivateKey) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetVaultKey

`func (o *EscrowedPrivateKey) GetVaultKey() string`

GetVaultKey returns the VaultKey field if non-nil, zero value otherwise.

### GetVaultKeyOk

`func (o *EscrowedPrivateKey) GetVaultKeyOk() (*string, bool)`

GetVaultKeyOk returns a tuple with the VaultKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultKey

`func (o *EscrowedPrivateKey) SetVaultKey(v string)`

SetVaultKey sets VaultKey field to given value.

### HasVaultKey

`func (o *EscrowedPrivateKey) HasVaultKey() bool`

HasVaultKey returns a boolean if a field has been set.

### SetVaultKeyNil

`func (o *EscrowedPrivateKey) SetVaultKeyNil(b bool)`

 SetVaultKeyNil sets the value for VaultKey to be an explicit nil

### UnsetVaultKey
`func (o *EscrowedPrivateKey) UnsetVaultKey()`

UnsetVaultKey ensures that no value is present for VaultKey, not even an explicit nil
### GetTransient

`func (o *EscrowedPrivateKey) GetTransient() bool`

GetTransient returns the Transient field if non-nil, zero value otherwise.

### GetTransientOk

`func (o *EscrowedPrivateKey) GetTransientOk() (*bool, bool)`

GetTransientOk returns a tuple with the Transient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransient

`func (o *EscrowedPrivateKey) SetTransient(v bool)`

SetTransient sets Transient field to given value.

### HasTransient

`func (o *EscrowedPrivateKey) HasTransient() bool`

HasTransient returns a boolean if a field has been set.

### SetTransientNil

`func (o *EscrowedPrivateKey) SetTransientNil(b bool)`

 SetTransientNil sets the value for Transient to be an explicit nil

### UnsetTransient
`func (o *EscrowedPrivateKey) UnsetTransient()`

UnsetTransient ensures that no value is present for Transient, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


