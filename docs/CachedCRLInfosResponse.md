# CachedCRLInfosResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ca** | **string** |  | 
**CRLNumber** | Pointer to **NullableFloat32** |  | [optional] 
**IssuerDn** | Pointer to **NullableString** |  | [optional] 
**ThisUpdate** | Pointer to **NullableInt64** |  | [optional] 
**NextUpdate** | Pointer to **NullableInt64** |  | [optional] 
**LastRefresh** | Pointer to **NullableInt64** |  | [optional] 
**NextRefresh** | Pointer to **NullableInt64** |  | [optional] 
**CrlSize** | **int64** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCachedCRLInfosResponse

`func NewCachedCRLInfosResponse(ca string, crlSize int64, ) *CachedCRLInfosResponse`

NewCachedCRLInfosResponse instantiates a new CachedCRLInfosResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCachedCRLInfosResponseWithDefaults

`func NewCachedCRLInfosResponseWithDefaults() *CachedCRLInfosResponse`

NewCachedCRLInfosResponseWithDefaults instantiates a new CachedCRLInfosResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *CachedCRLInfosResponse) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *CachedCRLInfosResponse) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *CachedCRLInfosResponse) SetCa(v string)`

SetCa sets Ca field to given value.


### GetCRLNumber

`func (o *CachedCRLInfosResponse) GetCRLNumber() float32`

GetCRLNumber returns the CRLNumber field if non-nil, zero value otherwise.

### GetCRLNumberOk

`func (o *CachedCRLInfosResponse) GetCRLNumberOk() (*float32, bool)`

GetCRLNumberOk returns a tuple with the CRLNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCRLNumber

`func (o *CachedCRLInfosResponse) SetCRLNumber(v float32)`

SetCRLNumber sets CRLNumber field to given value.

### HasCRLNumber

`func (o *CachedCRLInfosResponse) HasCRLNumber() bool`

HasCRLNumber returns a boolean if a field has been set.

### SetCRLNumberNil

`func (o *CachedCRLInfosResponse) SetCRLNumberNil(b bool)`

 SetCRLNumberNil sets the value for CRLNumber to be an explicit nil

### UnsetCRLNumber
`func (o *CachedCRLInfosResponse) UnsetCRLNumber()`

UnsetCRLNumber ensures that no value is present for CRLNumber, not even an explicit nil
### GetIssuerDn

`func (o *CachedCRLInfosResponse) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *CachedCRLInfosResponse) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *CachedCRLInfosResponse) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.

### HasIssuerDn

`func (o *CachedCRLInfosResponse) HasIssuerDn() bool`

HasIssuerDn returns a boolean if a field has been set.

### SetIssuerDnNil

`func (o *CachedCRLInfosResponse) SetIssuerDnNil(b bool)`

 SetIssuerDnNil sets the value for IssuerDn to be an explicit nil

### UnsetIssuerDn
`func (o *CachedCRLInfosResponse) UnsetIssuerDn()`

UnsetIssuerDn ensures that no value is present for IssuerDn, not even an explicit nil
### GetThisUpdate

`func (o *CachedCRLInfosResponse) GetThisUpdate() int64`

GetThisUpdate returns the ThisUpdate field if non-nil, zero value otherwise.

### GetThisUpdateOk

`func (o *CachedCRLInfosResponse) GetThisUpdateOk() (*int64, bool)`

GetThisUpdateOk returns a tuple with the ThisUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThisUpdate

`func (o *CachedCRLInfosResponse) SetThisUpdate(v int64)`

SetThisUpdate sets ThisUpdate field to given value.

### HasThisUpdate

`func (o *CachedCRLInfosResponse) HasThisUpdate() bool`

HasThisUpdate returns a boolean if a field has been set.

### SetThisUpdateNil

`func (o *CachedCRLInfosResponse) SetThisUpdateNil(b bool)`

 SetThisUpdateNil sets the value for ThisUpdate to be an explicit nil

### UnsetThisUpdate
`func (o *CachedCRLInfosResponse) UnsetThisUpdate()`

UnsetThisUpdate ensures that no value is present for ThisUpdate, not even an explicit nil
### GetNextUpdate

`func (o *CachedCRLInfosResponse) GetNextUpdate() int64`

GetNextUpdate returns the NextUpdate field if non-nil, zero value otherwise.

### GetNextUpdateOk

`func (o *CachedCRLInfosResponse) GetNextUpdateOk() (*int64, bool)`

GetNextUpdateOk returns a tuple with the NextUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpdate

`func (o *CachedCRLInfosResponse) SetNextUpdate(v int64)`

SetNextUpdate sets NextUpdate field to given value.

### HasNextUpdate

`func (o *CachedCRLInfosResponse) HasNextUpdate() bool`

HasNextUpdate returns a boolean if a field has been set.

### SetNextUpdateNil

`func (o *CachedCRLInfosResponse) SetNextUpdateNil(b bool)`

 SetNextUpdateNil sets the value for NextUpdate to be an explicit nil

### UnsetNextUpdate
`func (o *CachedCRLInfosResponse) UnsetNextUpdate()`

UnsetNextUpdate ensures that no value is present for NextUpdate, not even an explicit nil
### GetLastRefresh

`func (o *CachedCRLInfosResponse) GetLastRefresh() int64`

GetLastRefresh returns the LastRefresh field if non-nil, zero value otherwise.

### GetLastRefreshOk

`func (o *CachedCRLInfosResponse) GetLastRefreshOk() (*int64, bool)`

GetLastRefreshOk returns a tuple with the LastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefresh

`func (o *CachedCRLInfosResponse) SetLastRefresh(v int64)`

SetLastRefresh sets LastRefresh field to given value.

### HasLastRefresh

`func (o *CachedCRLInfosResponse) HasLastRefresh() bool`

HasLastRefresh returns a boolean if a field has been set.

### SetLastRefreshNil

`func (o *CachedCRLInfosResponse) SetLastRefreshNil(b bool)`

 SetLastRefreshNil sets the value for LastRefresh to be an explicit nil

### UnsetLastRefresh
`func (o *CachedCRLInfosResponse) UnsetLastRefresh()`

UnsetLastRefresh ensures that no value is present for LastRefresh, not even an explicit nil
### GetNextRefresh

`func (o *CachedCRLInfosResponse) GetNextRefresh() int64`

GetNextRefresh returns the NextRefresh field if non-nil, zero value otherwise.

### GetNextRefreshOk

`func (o *CachedCRLInfosResponse) GetNextRefreshOk() (*int64, bool)`

GetNextRefreshOk returns a tuple with the NextRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRefresh

`func (o *CachedCRLInfosResponse) SetNextRefresh(v int64)`

SetNextRefresh sets NextRefresh field to given value.

### HasNextRefresh

`func (o *CachedCRLInfosResponse) HasNextRefresh() bool`

HasNextRefresh returns a boolean if a field has been set.

### SetNextRefreshNil

`func (o *CachedCRLInfosResponse) SetNextRefreshNil(b bool)`

 SetNextRefreshNil sets the value for NextRefresh to be an explicit nil

### UnsetNextRefresh
`func (o *CachedCRLInfosResponse) UnsetNextRefresh()`

UnsetNextRefresh ensures that no value is present for NextRefresh, not even an explicit nil
### GetCrlSize

`func (o *CachedCRLInfosResponse) GetCrlSize() int64`

GetCrlSize returns the CrlSize field if non-nil, zero value otherwise.

### GetCrlSizeOk

`func (o *CachedCRLInfosResponse) GetCrlSizeOk() (*int64, bool)`

GetCrlSizeOk returns a tuple with the CrlSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlSize

`func (o *CachedCRLInfosResponse) SetCrlSize(v int64)`

SetCrlSize sets CrlSize field to given value.


### GetError

`func (o *CachedCRLInfosResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CachedCRLInfosResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CachedCRLInfosResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CachedCRLInfosResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *CachedCRLInfosResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *CachedCRLInfosResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


