# DatasourceTest200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the datasource executed | 
**ComputedLookupValues** | Pointer to **[]string** | Lookup values that were requested on the datasource | [optional] 
**Name** | **string** | Name of the executed datasource | 
**Status** | **string** | Status of the execution. &#x60;success&#x60; if the datasource data was fetched correctly, &#x60;failure&#x60; if an error occured and &#x60;ignored&#x60; if inputs were not all filled, resulting in no request being sent | 
**Dictionary** | [**[]MapEntry**](MapEntry.md) | Data fetched from the datasource | 
**Error** | Pointer to **NullableString** | If &#x60;status&#x60; is &#x60;failure&#x60;, the error message | [optional] 
**ComputedDN** | Pointer to **NullableString** | DN that was requested on the LDAP server | [optional] 
**ComputedFilter** | Pointer to **NullableString** | Filter that was requested on the LDAP server | [optional] 
**ComputedUrl** | Pointer to **NullableString** | Url that was requested | [optional] 
**ComputedPayload** | Pointer to **NullableString** | Url that was requested | [optional] 
**ComputedHeaders** | Pointer to **[]map[string]interface{}** | Headers that were sent | [optional] 
**ResponseCode** | Pointer to **NullableInt64** | Received response code | [optional] 
**ResponseHeaders** | Pointer to **[]map[string]interface{}** | Headers that were received | [optional] 
**ResponseBody** | Pointer to **NullableString** | Received response body | [optional] 

## Methods

### NewDatasourceTest200Response

`func NewDatasourceTest200Response(type_ string, name string, status string, dictionary []MapEntry, ) *DatasourceTest200Response`

NewDatasourceTest200Response instantiates a new DatasourceTest200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceTest200ResponseWithDefaults

`func NewDatasourceTest200ResponseWithDefaults() *DatasourceTest200Response`

NewDatasourceTest200ResponseWithDefaults instantiates a new DatasourceTest200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DatasourceTest200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasourceTest200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasourceTest200Response) SetType(v string)`

SetType sets Type field to given value.


### GetComputedLookupValues

`func (o *DatasourceTest200Response) GetComputedLookupValues() []string`

GetComputedLookupValues returns the ComputedLookupValues field if non-nil, zero value otherwise.

### GetComputedLookupValuesOk

`func (o *DatasourceTest200Response) GetComputedLookupValuesOk() (*[]string, bool)`

GetComputedLookupValuesOk returns a tuple with the ComputedLookupValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedLookupValues

`func (o *DatasourceTest200Response) SetComputedLookupValues(v []string)`

SetComputedLookupValues sets ComputedLookupValues field to given value.

### HasComputedLookupValues

`func (o *DatasourceTest200Response) HasComputedLookupValues() bool`

HasComputedLookupValues returns a boolean if a field has been set.

### SetComputedLookupValuesNil

`func (o *DatasourceTest200Response) SetComputedLookupValuesNil(b bool)`

 SetComputedLookupValuesNil sets the value for ComputedLookupValues to be an explicit nil

### UnsetComputedLookupValues
`func (o *DatasourceTest200Response) UnsetComputedLookupValues()`

UnsetComputedLookupValues ensures that no value is present for ComputedLookupValues, not even an explicit nil
### GetName

`func (o *DatasourceTest200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasourceTest200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasourceTest200Response) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DatasourceTest200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatasourceTest200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatasourceTest200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDictionary

`func (o *DatasourceTest200Response) GetDictionary() []MapEntry`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *DatasourceTest200Response) GetDictionaryOk() (*[]MapEntry, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *DatasourceTest200Response) SetDictionary(v []MapEntry)`

SetDictionary sets Dictionary field to given value.


### GetError

`func (o *DatasourceTest200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DatasourceTest200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DatasourceTest200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DatasourceTest200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *DatasourceTest200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DatasourceTest200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetComputedDN

`func (o *DatasourceTest200Response) GetComputedDN() string`

GetComputedDN returns the ComputedDN field if non-nil, zero value otherwise.

### GetComputedDNOk

`func (o *DatasourceTest200Response) GetComputedDNOk() (*string, bool)`

GetComputedDNOk returns a tuple with the ComputedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedDN

`func (o *DatasourceTest200Response) SetComputedDN(v string)`

SetComputedDN sets ComputedDN field to given value.

### HasComputedDN

`func (o *DatasourceTest200Response) HasComputedDN() bool`

HasComputedDN returns a boolean if a field has been set.

### SetComputedDNNil

`func (o *DatasourceTest200Response) SetComputedDNNil(b bool)`

 SetComputedDNNil sets the value for ComputedDN to be an explicit nil

### UnsetComputedDN
`func (o *DatasourceTest200Response) UnsetComputedDN()`

UnsetComputedDN ensures that no value is present for ComputedDN, not even an explicit nil
### GetComputedFilter

`func (o *DatasourceTest200Response) GetComputedFilter() string`

GetComputedFilter returns the ComputedFilter field if non-nil, zero value otherwise.

### GetComputedFilterOk

`func (o *DatasourceTest200Response) GetComputedFilterOk() (*string, bool)`

GetComputedFilterOk returns a tuple with the ComputedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFilter

`func (o *DatasourceTest200Response) SetComputedFilter(v string)`

SetComputedFilter sets ComputedFilter field to given value.

### HasComputedFilter

`func (o *DatasourceTest200Response) HasComputedFilter() bool`

HasComputedFilter returns a boolean if a field has been set.

### SetComputedFilterNil

`func (o *DatasourceTest200Response) SetComputedFilterNil(b bool)`

 SetComputedFilterNil sets the value for ComputedFilter to be an explicit nil

### UnsetComputedFilter
`func (o *DatasourceTest200Response) UnsetComputedFilter()`

UnsetComputedFilter ensures that no value is present for ComputedFilter, not even an explicit nil
### GetComputedUrl

`func (o *DatasourceTest200Response) GetComputedUrl() string`

GetComputedUrl returns the ComputedUrl field if non-nil, zero value otherwise.

### GetComputedUrlOk

`func (o *DatasourceTest200Response) GetComputedUrlOk() (*string, bool)`

GetComputedUrlOk returns a tuple with the ComputedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedUrl

`func (o *DatasourceTest200Response) SetComputedUrl(v string)`

SetComputedUrl sets ComputedUrl field to given value.

### HasComputedUrl

`func (o *DatasourceTest200Response) HasComputedUrl() bool`

HasComputedUrl returns a boolean if a field has been set.

### SetComputedUrlNil

`func (o *DatasourceTest200Response) SetComputedUrlNil(b bool)`

 SetComputedUrlNil sets the value for ComputedUrl to be an explicit nil

### UnsetComputedUrl
`func (o *DatasourceTest200Response) UnsetComputedUrl()`

UnsetComputedUrl ensures that no value is present for ComputedUrl, not even an explicit nil
### GetComputedPayload

`func (o *DatasourceTest200Response) GetComputedPayload() string`

GetComputedPayload returns the ComputedPayload field if non-nil, zero value otherwise.

### GetComputedPayloadOk

`func (o *DatasourceTest200Response) GetComputedPayloadOk() (*string, bool)`

GetComputedPayloadOk returns a tuple with the ComputedPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPayload

`func (o *DatasourceTest200Response) SetComputedPayload(v string)`

SetComputedPayload sets ComputedPayload field to given value.

### HasComputedPayload

`func (o *DatasourceTest200Response) HasComputedPayload() bool`

HasComputedPayload returns a boolean if a field has been set.

### SetComputedPayloadNil

`func (o *DatasourceTest200Response) SetComputedPayloadNil(b bool)`

 SetComputedPayloadNil sets the value for ComputedPayload to be an explicit nil

### UnsetComputedPayload
`func (o *DatasourceTest200Response) UnsetComputedPayload()`

UnsetComputedPayload ensures that no value is present for ComputedPayload, not even an explicit nil
### GetComputedHeaders

`func (o *DatasourceTest200Response) GetComputedHeaders() []map[string]interface{}`

GetComputedHeaders returns the ComputedHeaders field if non-nil, zero value otherwise.

### GetComputedHeadersOk

`func (o *DatasourceTest200Response) GetComputedHeadersOk() (*[]map[string]interface{}, bool)`

GetComputedHeadersOk returns a tuple with the ComputedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedHeaders

`func (o *DatasourceTest200Response) SetComputedHeaders(v []map[string]interface{})`

SetComputedHeaders sets ComputedHeaders field to given value.

### HasComputedHeaders

`func (o *DatasourceTest200Response) HasComputedHeaders() bool`

HasComputedHeaders returns a boolean if a field has been set.

### SetComputedHeadersNil

`func (o *DatasourceTest200Response) SetComputedHeadersNil(b bool)`

 SetComputedHeadersNil sets the value for ComputedHeaders to be an explicit nil

### UnsetComputedHeaders
`func (o *DatasourceTest200Response) UnsetComputedHeaders()`

UnsetComputedHeaders ensures that no value is present for ComputedHeaders, not even an explicit nil
### GetResponseCode

`func (o *DatasourceTest200Response) GetResponseCode() int64`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *DatasourceTest200Response) GetResponseCodeOk() (*int64, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *DatasourceTest200Response) SetResponseCode(v int64)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *DatasourceTest200Response) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### SetResponseCodeNil

`func (o *DatasourceTest200Response) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *DatasourceTest200Response) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
### GetResponseHeaders

`func (o *DatasourceTest200Response) GetResponseHeaders() []map[string]interface{}`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *DatasourceTest200Response) GetResponseHeadersOk() (*[]map[string]interface{}, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *DatasourceTest200Response) SetResponseHeaders(v []map[string]interface{})`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *DatasourceTest200Response) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *DatasourceTest200Response) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *DatasourceTest200Response) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
### GetResponseBody

`func (o *DatasourceTest200Response) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *DatasourceTest200Response) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *DatasourceTest200Response) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *DatasourceTest200Response) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *DatasourceTest200Response) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *DatasourceTest200Response) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


