# CFCrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Issuer** | **string** |  | 
**ThisUpdate** | **int64** |  | 
**NextUpdate** | **int64** |  | 

## Methods

### NewCFCrl

`func NewCFCrl(issuer string, thisUpdate int64, nextUpdate int64, ) *CFCrl`

NewCFCrl instantiates a new CFCrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCrlWithDefaults

`func NewCFCrlWithDefaults() *CFCrl`

NewCFCrlWithDefaults instantiates a new CFCrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CFCrl) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CFCrl) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CFCrl) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CFCrl) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetVersion

`func (o *CFCrl) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CFCrl) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CFCrl) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CFCrl) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIssuer

`func (o *CFCrl) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CFCrl) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CFCrl) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetThisUpdate

`func (o *CFCrl) GetThisUpdate() int64`

GetThisUpdate returns the ThisUpdate field if non-nil, zero value otherwise.

### GetThisUpdateOk

`func (o *CFCrl) GetThisUpdateOk() (*int64, bool)`

GetThisUpdateOk returns a tuple with the ThisUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThisUpdate

`func (o *CFCrl) SetThisUpdate(v int64)`

SetThisUpdate sets ThisUpdate field to given value.


### GetNextUpdate

`func (o *CFCrl) GetNextUpdate() int64`

GetNextUpdate returns the NextUpdate field if non-nil, zero value otherwise.

### GetNextUpdateOk

`func (o *CFCrl) GetNextUpdateOk() (*int64, bool)`

GetNextUpdateOk returns a tuple with the NextUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpdate

`func (o *CFCrl) SetNextUpdate(v int64)`

SetNextUpdate sets NextUpdate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


