# LDAPDataSourceResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ComputedDN** | Pointer to **NullableString** | DN that was requested on the LDAP server | [optional] 
**ComputedFilter** | Pointer to **NullableString** | Filter that was requested on the LDAP server | [optional] 
**Name** | **string** | Name of the executed datasource | 
**Status** | **string** | Status of the execution. &#x60;success&#x60; if the datasource data was fetched correctly, &#x60;failure&#x60; if an error occured and &#x60;ignored&#x60; if inputs were not all filled, resulting in no request being sent | 
**Dictionary** | [**[]MapEntry**](MapEntry.md) | Data fetched from the datasource | 
**Error** | Pointer to **NullableString** | If &#x60;status&#x60; is &#x60;failure&#x60;, the error message | [optional] 

## Methods

### NewLDAPDataSourceResultResponse

`func NewLDAPDataSourceResultResponse(type_ string, name string, status string, dictionary []MapEntry, ) *LDAPDataSourceResultResponse`

NewLDAPDataSourceResultResponse instantiates a new LDAPDataSourceResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPDataSourceResultResponseWithDefaults

`func NewLDAPDataSourceResultResponseWithDefaults() *LDAPDataSourceResultResponse`

NewLDAPDataSourceResultResponseWithDefaults instantiates a new LDAPDataSourceResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LDAPDataSourceResultResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LDAPDataSourceResultResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LDAPDataSourceResultResponse) SetType(v string)`

SetType sets Type field to given value.


### GetComputedDN

`func (o *LDAPDataSourceResultResponse) GetComputedDN() string`

GetComputedDN returns the ComputedDN field if non-nil, zero value otherwise.

### GetComputedDNOk

`func (o *LDAPDataSourceResultResponse) GetComputedDNOk() (*string, bool)`

GetComputedDNOk returns a tuple with the ComputedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedDN

`func (o *LDAPDataSourceResultResponse) SetComputedDN(v string)`

SetComputedDN sets ComputedDN field to given value.

### HasComputedDN

`func (o *LDAPDataSourceResultResponse) HasComputedDN() bool`

HasComputedDN returns a boolean if a field has been set.

### SetComputedDNNil

`func (o *LDAPDataSourceResultResponse) SetComputedDNNil(b bool)`

 SetComputedDNNil sets the value for ComputedDN to be an explicit nil

### UnsetComputedDN
`func (o *LDAPDataSourceResultResponse) UnsetComputedDN()`

UnsetComputedDN ensures that no value is present for ComputedDN, not even an explicit nil
### GetComputedFilter

`func (o *LDAPDataSourceResultResponse) GetComputedFilter() string`

GetComputedFilter returns the ComputedFilter field if non-nil, zero value otherwise.

### GetComputedFilterOk

`func (o *LDAPDataSourceResultResponse) GetComputedFilterOk() (*string, bool)`

GetComputedFilterOk returns a tuple with the ComputedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFilter

`func (o *LDAPDataSourceResultResponse) SetComputedFilter(v string)`

SetComputedFilter sets ComputedFilter field to given value.

### HasComputedFilter

`func (o *LDAPDataSourceResultResponse) HasComputedFilter() bool`

HasComputedFilter returns a boolean if a field has been set.

### SetComputedFilterNil

`func (o *LDAPDataSourceResultResponse) SetComputedFilterNil(b bool)`

 SetComputedFilterNil sets the value for ComputedFilter to be an explicit nil

### UnsetComputedFilter
`func (o *LDAPDataSourceResultResponse) UnsetComputedFilter()`

UnsetComputedFilter ensures that no value is present for ComputedFilter, not even an explicit nil
### GetName

`func (o *LDAPDataSourceResultResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPDataSourceResultResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPDataSourceResultResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *LDAPDataSourceResultResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LDAPDataSourceResultResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LDAPDataSourceResultResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDictionary

`func (o *LDAPDataSourceResultResponse) GetDictionary() []MapEntry`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *LDAPDataSourceResultResponse) GetDictionaryOk() (*[]MapEntry, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *LDAPDataSourceResultResponse) SetDictionary(v []MapEntry)`

SetDictionary sets Dictionary field to given value.


### GetError

`func (o *LDAPDataSourceResultResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LDAPDataSourceResultResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LDAPDataSourceResultResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *LDAPDataSourceResultResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *LDAPDataSourceResultResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *LDAPDataSourceResultResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


