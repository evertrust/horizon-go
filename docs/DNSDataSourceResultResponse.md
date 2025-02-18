# DNSDataSourceResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ComputedLookupValues** | Pointer to **[]string** | Lookup values that were requested on the datasource | [optional] 
**Name** | **string** | Name of the executed datasource | 
**Status** | **string** | Status of the execution. &#x60;success&#x60; if the datasource data was fetched correctly, &#x60;failure&#x60; if an error occured and &#x60;ignored&#x60; if inputs were not all filled, resulting in no request being sent | 
**Dictionary** | [**[]MapEntry**](MapEntry.md) | Data fetched from the datasource | 
**Error** | Pointer to **NullableString** | If &#x60;status&#x60; is &#x60;failure&#x60;, the error message | [optional] 

## Methods

### NewDNSDataSourceResultResponse

`func NewDNSDataSourceResultResponse(type_ string, name string, status string, dictionary []MapEntry, ) *DNSDataSourceResultResponse`

NewDNSDataSourceResultResponse instantiates a new DNSDataSourceResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSDataSourceResultResponseWithDefaults

`func NewDNSDataSourceResultResponseWithDefaults() *DNSDataSourceResultResponse`

NewDNSDataSourceResultResponseWithDefaults instantiates a new DNSDataSourceResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSDataSourceResultResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSDataSourceResultResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSDataSourceResultResponse) SetType(v string)`

SetType sets Type field to given value.


### GetComputedLookupValues

`func (o *DNSDataSourceResultResponse) GetComputedLookupValues() []string`

GetComputedLookupValues returns the ComputedLookupValues field if non-nil, zero value otherwise.

### GetComputedLookupValuesOk

`func (o *DNSDataSourceResultResponse) GetComputedLookupValuesOk() (*[]string, bool)`

GetComputedLookupValuesOk returns a tuple with the ComputedLookupValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedLookupValues

`func (o *DNSDataSourceResultResponse) SetComputedLookupValues(v []string)`

SetComputedLookupValues sets ComputedLookupValues field to given value.

### HasComputedLookupValues

`func (o *DNSDataSourceResultResponse) HasComputedLookupValues() bool`

HasComputedLookupValues returns a boolean if a field has been set.

### SetComputedLookupValuesNil

`func (o *DNSDataSourceResultResponse) SetComputedLookupValuesNil(b bool)`

 SetComputedLookupValuesNil sets the value for ComputedLookupValues to be an explicit nil

### UnsetComputedLookupValues
`func (o *DNSDataSourceResultResponse) UnsetComputedLookupValues()`

UnsetComputedLookupValues ensures that no value is present for ComputedLookupValues, not even an explicit nil
### GetName

`func (o *DNSDataSourceResultResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSDataSourceResultResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSDataSourceResultResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DNSDataSourceResultResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DNSDataSourceResultResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DNSDataSourceResultResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDictionary

`func (o *DNSDataSourceResultResponse) GetDictionary() []MapEntry`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *DNSDataSourceResultResponse) GetDictionaryOk() (*[]MapEntry, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *DNSDataSourceResultResponse) SetDictionary(v []MapEntry)`

SetDictionary sets Dictionary field to given value.


### GetError

`func (o *DNSDataSourceResultResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DNSDataSourceResultResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DNSDataSourceResultResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DNSDataSourceResultResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *DNSDataSourceResultResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DNSDataSourceResultResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


