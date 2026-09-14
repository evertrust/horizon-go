# JsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | Algorithm | [optional] 
**E** | Pointer to **string** | Exponent (for RSA) | [optional] 
**K** | Pointer to **string** | Symmetric key | [optional] 
**Kty** | Pointer to **string** | Key type | [optional] 
**N** | Pointer to **string** | Modulus (for RSA) | [optional] 
**Use** | Pointer to **string** | Key usage | [optional] 
**X** | Pointer to **string** | X coordinate (for EC) | [optional] 
**Y** | Pointer to **string** | Y coordinate (for EC) | [optional] 

## Methods

### NewJsonWebKey

`func NewJsonWebKey() *JsonWebKey`

NewJsonWebKey instantiates a new JsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebKeyWithDefaults

`func NewJsonWebKeyWithDefaults() *JsonWebKey`

NewJsonWebKeyWithDefaults instantiates a new JsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *JsonWebKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *JsonWebKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *JsonWebKey) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *JsonWebKey) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetE

`func (o *JsonWebKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *JsonWebKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *JsonWebKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *JsonWebKey) HasE() bool`

HasE returns a boolean if a field has been set.

### GetK

`func (o *JsonWebKey) GetK() string`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *JsonWebKey) GetKOk() (*string, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *JsonWebKey) SetK(v string)`

SetK sets K field to given value.

### HasK

`func (o *JsonWebKey) HasK() bool`

HasK returns a boolean if a field has been set.

### GetKty

`func (o *JsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *JsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *JsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *JsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *JsonWebKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *JsonWebKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *JsonWebKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *JsonWebKey) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUse

`func (o *JsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *JsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *JsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *JsonWebKey) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetX

`func (o *JsonWebKey) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *JsonWebKey) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *JsonWebKey) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *JsonWebKey) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *JsonWebKey) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *JsonWebKey) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *JsonWebKey) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *JsonWebKey) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


