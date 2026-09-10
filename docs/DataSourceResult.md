# DataSourceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dictionary** | Pointer to [**[]MapEntry**](MapEntry.md) | Data fetched from the datasource | [optional] 
**Error** | Pointer to **NullableString** | If &#x60;status&#x60; is &#x60;failure&#x60;, the error message | [optional] 
**Name** | Pointer to **string** | Name of the executed datasource | [optional] 
**Status** | Pointer to **string** | Status of the execution. &#x60;success&#x60; if the datasource data was fetched correctly, &#x60;failure&#x60; if an error occured, &#x60;not_found&#x60; if the datasource query returned no results and &#x60;ignored&#x60; if inputs were not all filled, resulting in no request being sent | [optional] 
**Type** | Pointer to **string** | Type of the datasource executed | [optional] 

## Methods

### NewDataSourceResult

`func NewDataSourceResult() *DataSourceResult`

NewDataSourceResult instantiates a new DataSourceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceResultWithDefaults

`func NewDataSourceResultWithDefaults() *DataSourceResult`

NewDataSourceResultWithDefaults instantiates a new DataSourceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDictionary

`func (o *DataSourceResult) GetDictionary() []MapEntry`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *DataSourceResult) GetDictionaryOk() (*[]MapEntry, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *DataSourceResult) SetDictionary(v []MapEntry)`

SetDictionary sets Dictionary field to given value.

### HasDictionary

`func (o *DataSourceResult) HasDictionary() bool`

HasDictionary returns a boolean if a field has been set.

### GetError

`func (o *DataSourceResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DataSourceResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DataSourceResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DataSourceResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *DataSourceResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DataSourceResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetName

`func (o *DataSourceResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataSourceResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DataSourceResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataSourceResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataSourceResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataSourceResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DataSourceResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataSourceResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataSourceResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DataSourceResult) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


