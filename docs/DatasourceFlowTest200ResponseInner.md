# DatasourceFlowTest200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputedLookupValues** | Pointer to **[]string** | Lookup values that were requested on the datasource | [optional] 
**Type** | **string** | Type of the datasource executed | 
**Dictionary** | [**[]MapEntry**](MapEntry.md) | Data fetched from the datasource | 
**Error** | Pointer to **NullableString** | If &#x60;status&#x60; is &#x60;failure&#x60;, the error message | [optional] 
**Name** | **string** | Name of the executed datasource | 
**Status** | **string** | Status of the execution. &#x60;success&#x60; if the datasource data was fetched correctly, &#x60;failure&#x60; if an error occured, &#x60;not_found&#x60; if the datasource query returned no results and &#x60;ignored&#x60; if inputs were not all filled, resulting in no request being sent | 
**ComputedDN** | Pointer to **NullableString** | DN that was requested on the LDAP server | [optional] 
**ComputedFilter** | Pointer to **NullableString** | Filter that was requested on the LDAP server | [optional] 
**ComputedHeaders** | Pointer to **[]map[string]interface{}** | Headers that were sent | [optional] 
**ComputedPayload** | Pointer to **NullableString** | Url that was requested | [optional] 
**ComputedUrl** | Pointer to **NullableString** | Url that was requested | [optional] 
**ResponseBody** | Pointer to **NullableString** | Received response body | [optional] 
**ResponseCode** | Pointer to **NullableInt64** | Received response code | [optional] 
**ResponseHeaders** | Pointer to **[]map[string]interface{}** | Headers that were received | [optional] 

## Methods

### NewDatasourceFlowTest200ResponseInner

`func NewDatasourceFlowTest200ResponseInner(type_ string, dictionary []MapEntry, name string, status string, ) *DatasourceFlowTest200ResponseInner`

NewDatasourceFlowTest200ResponseInner instantiates a new DatasourceFlowTest200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceFlowTest200ResponseInnerWithDefaults

`func NewDatasourceFlowTest200ResponseInnerWithDefaults() *DatasourceFlowTest200ResponseInner`

NewDatasourceFlowTest200ResponseInnerWithDefaults instantiates a new DatasourceFlowTest200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputedLookupValues

`func (o *DatasourceFlowTest200ResponseInner) GetComputedLookupValues() []string`

GetComputedLookupValues returns the ComputedLookupValues field if non-nil, zero value otherwise.

### GetComputedLookupValuesOk

`func (o *DatasourceFlowTest200ResponseInner) GetComputedLookupValuesOk() (*[]string, bool)`

GetComputedLookupValuesOk returns a tuple with the ComputedLookupValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedLookupValues

`func (o *DatasourceFlowTest200ResponseInner) SetComputedLookupValues(v []string)`

SetComputedLookupValues sets ComputedLookupValues field to given value.

### HasComputedLookupValues

`func (o *DatasourceFlowTest200ResponseInner) HasComputedLookupValues() bool`

HasComputedLookupValues returns a boolean if a field has been set.

### SetComputedLookupValuesNil

`func (o *DatasourceFlowTest200ResponseInner) SetComputedLookupValuesNil(b bool)`

 SetComputedLookupValuesNil sets the value for ComputedLookupValues to be an explicit nil

### UnsetComputedLookupValues
`func (o *DatasourceFlowTest200ResponseInner) UnsetComputedLookupValues()`

UnsetComputedLookupValues ensures that no value is present for ComputedLookupValues, not even an explicit nil
### GetType

`func (o *DatasourceFlowTest200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasourceFlowTest200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasourceFlowTest200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetDictionary

`func (o *DatasourceFlowTest200ResponseInner) GetDictionary() []MapEntry`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *DatasourceFlowTest200ResponseInner) GetDictionaryOk() (*[]MapEntry, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *DatasourceFlowTest200ResponseInner) SetDictionary(v []MapEntry)`

SetDictionary sets Dictionary field to given value.


### GetError

`func (o *DatasourceFlowTest200ResponseInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DatasourceFlowTest200ResponseInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DatasourceFlowTest200ResponseInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DatasourceFlowTest200ResponseInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *DatasourceFlowTest200ResponseInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DatasourceFlowTest200ResponseInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetName

`func (o *DatasourceFlowTest200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasourceFlowTest200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasourceFlowTest200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DatasourceFlowTest200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatasourceFlowTest200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatasourceFlowTest200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetComputedDN

`func (o *DatasourceFlowTest200ResponseInner) GetComputedDN() string`

GetComputedDN returns the ComputedDN field if non-nil, zero value otherwise.

### GetComputedDNOk

`func (o *DatasourceFlowTest200ResponseInner) GetComputedDNOk() (*string, bool)`

GetComputedDNOk returns a tuple with the ComputedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedDN

`func (o *DatasourceFlowTest200ResponseInner) SetComputedDN(v string)`

SetComputedDN sets ComputedDN field to given value.

### HasComputedDN

`func (o *DatasourceFlowTest200ResponseInner) HasComputedDN() bool`

HasComputedDN returns a boolean if a field has been set.

### SetComputedDNNil

`func (o *DatasourceFlowTest200ResponseInner) SetComputedDNNil(b bool)`

 SetComputedDNNil sets the value for ComputedDN to be an explicit nil

### UnsetComputedDN
`func (o *DatasourceFlowTest200ResponseInner) UnsetComputedDN()`

UnsetComputedDN ensures that no value is present for ComputedDN, not even an explicit nil
### GetComputedFilter

`func (o *DatasourceFlowTest200ResponseInner) GetComputedFilter() string`

GetComputedFilter returns the ComputedFilter field if non-nil, zero value otherwise.

### GetComputedFilterOk

`func (o *DatasourceFlowTest200ResponseInner) GetComputedFilterOk() (*string, bool)`

GetComputedFilterOk returns a tuple with the ComputedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFilter

`func (o *DatasourceFlowTest200ResponseInner) SetComputedFilter(v string)`

SetComputedFilter sets ComputedFilter field to given value.

### HasComputedFilter

`func (o *DatasourceFlowTest200ResponseInner) HasComputedFilter() bool`

HasComputedFilter returns a boolean if a field has been set.

### SetComputedFilterNil

`func (o *DatasourceFlowTest200ResponseInner) SetComputedFilterNil(b bool)`

 SetComputedFilterNil sets the value for ComputedFilter to be an explicit nil

### UnsetComputedFilter
`func (o *DatasourceFlowTest200ResponseInner) UnsetComputedFilter()`

UnsetComputedFilter ensures that no value is present for ComputedFilter, not even an explicit nil
### GetComputedHeaders

`func (o *DatasourceFlowTest200ResponseInner) GetComputedHeaders() []map[string]interface{}`

GetComputedHeaders returns the ComputedHeaders field if non-nil, zero value otherwise.

### GetComputedHeadersOk

`func (o *DatasourceFlowTest200ResponseInner) GetComputedHeadersOk() (*[]map[string]interface{}, bool)`

GetComputedHeadersOk returns a tuple with the ComputedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedHeaders

`func (o *DatasourceFlowTest200ResponseInner) SetComputedHeaders(v []map[string]interface{})`

SetComputedHeaders sets ComputedHeaders field to given value.

### HasComputedHeaders

`func (o *DatasourceFlowTest200ResponseInner) HasComputedHeaders() bool`

HasComputedHeaders returns a boolean if a field has been set.

### SetComputedHeadersNil

`func (o *DatasourceFlowTest200ResponseInner) SetComputedHeadersNil(b bool)`

 SetComputedHeadersNil sets the value for ComputedHeaders to be an explicit nil

### UnsetComputedHeaders
`func (o *DatasourceFlowTest200ResponseInner) UnsetComputedHeaders()`

UnsetComputedHeaders ensures that no value is present for ComputedHeaders, not even an explicit nil
### GetComputedPayload

`func (o *DatasourceFlowTest200ResponseInner) GetComputedPayload() string`

GetComputedPayload returns the ComputedPayload field if non-nil, zero value otherwise.

### GetComputedPayloadOk

`func (o *DatasourceFlowTest200ResponseInner) GetComputedPayloadOk() (*string, bool)`

GetComputedPayloadOk returns a tuple with the ComputedPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPayload

`func (o *DatasourceFlowTest200ResponseInner) SetComputedPayload(v string)`

SetComputedPayload sets ComputedPayload field to given value.

### HasComputedPayload

`func (o *DatasourceFlowTest200ResponseInner) HasComputedPayload() bool`

HasComputedPayload returns a boolean if a field has been set.

### SetComputedPayloadNil

`func (o *DatasourceFlowTest200ResponseInner) SetComputedPayloadNil(b bool)`

 SetComputedPayloadNil sets the value for ComputedPayload to be an explicit nil

### UnsetComputedPayload
`func (o *DatasourceFlowTest200ResponseInner) UnsetComputedPayload()`

UnsetComputedPayload ensures that no value is present for ComputedPayload, not even an explicit nil
### GetComputedUrl

`func (o *DatasourceFlowTest200ResponseInner) GetComputedUrl() string`

GetComputedUrl returns the ComputedUrl field if non-nil, zero value otherwise.

### GetComputedUrlOk

`func (o *DatasourceFlowTest200ResponseInner) GetComputedUrlOk() (*string, bool)`

GetComputedUrlOk returns a tuple with the ComputedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedUrl

`func (o *DatasourceFlowTest200ResponseInner) SetComputedUrl(v string)`

SetComputedUrl sets ComputedUrl field to given value.

### HasComputedUrl

`func (o *DatasourceFlowTest200ResponseInner) HasComputedUrl() bool`

HasComputedUrl returns a boolean if a field has been set.

### SetComputedUrlNil

`func (o *DatasourceFlowTest200ResponseInner) SetComputedUrlNil(b bool)`

 SetComputedUrlNil sets the value for ComputedUrl to be an explicit nil

### UnsetComputedUrl
`func (o *DatasourceFlowTest200ResponseInner) UnsetComputedUrl()`

UnsetComputedUrl ensures that no value is present for ComputedUrl, not even an explicit nil
### GetResponseBody

`func (o *DatasourceFlowTest200ResponseInner) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *DatasourceFlowTest200ResponseInner) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *DatasourceFlowTest200ResponseInner) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *DatasourceFlowTest200ResponseInner) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *DatasourceFlowTest200ResponseInner) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *DatasourceFlowTest200ResponseInner) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseCode

`func (o *DatasourceFlowTest200ResponseInner) GetResponseCode() int64`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *DatasourceFlowTest200ResponseInner) GetResponseCodeOk() (*int64, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *DatasourceFlowTest200ResponseInner) SetResponseCode(v int64)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *DatasourceFlowTest200ResponseInner) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### SetResponseCodeNil

`func (o *DatasourceFlowTest200ResponseInner) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *DatasourceFlowTest200ResponseInner) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
### GetResponseHeaders

`func (o *DatasourceFlowTest200ResponseInner) GetResponseHeaders() []map[string]interface{}`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *DatasourceFlowTest200ResponseInner) GetResponseHeadersOk() (*[]map[string]interface{}, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *DatasourceFlowTest200ResponseInner) SetResponseHeaders(v []map[string]interface{})`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *DatasourceFlowTest200ResponseInner) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *DatasourceFlowTest200ResponseInner) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *DatasourceFlowTest200ResponseInner) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


