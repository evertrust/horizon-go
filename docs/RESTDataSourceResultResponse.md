# RESTDataSourceResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputedHeaders** | Pointer to **[]map[string]interface{}** | Headers that were sent | [optional] 
**ComputedPayload** | Pointer to **NullableString** | Url that was requested | [optional] 
**ComputedUrl** | Pointer to **NullableString** | Url that was requested | [optional] 
**ResponseBody** | Pointer to **NullableString** | Received response body | [optional] 
**ResponseCode** | Pointer to **NullableInt64** | Received response code | [optional] 
**ResponseHeaders** | Pointer to **[]map[string]interface{}** | Headers that were received | [optional] 
**Type** | **string** |  | 
**Dictionary** | [**[]MapEntry**](MapEntry.md) | Data fetched from the datasource | 
**Error** | Pointer to **NullableString** | If &#x60;status&#x60; is &#x60;failure&#x60;, the error message | [optional] 
**Name** | **string** | Name of the executed datasource | 
**Status** | **string** | Status of the execution. &#x60;success&#x60; if the datasource data was fetched correctly, &#x60;failure&#x60; if an error occured and &#x60;ignored&#x60; if inputs were not all filled, resulting in no request being sent | 

## Methods

### NewRESTDataSourceResultResponse

`func NewRESTDataSourceResultResponse(type_ string, dictionary []MapEntry, name string, status string, ) *RESTDataSourceResultResponse`

NewRESTDataSourceResultResponse instantiates a new RESTDataSourceResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRESTDataSourceResultResponseWithDefaults

`func NewRESTDataSourceResultResponseWithDefaults() *RESTDataSourceResultResponse`

NewRESTDataSourceResultResponseWithDefaults instantiates a new RESTDataSourceResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputedHeaders

`func (o *RESTDataSourceResultResponse) GetComputedHeaders() []map[string]interface{}`

GetComputedHeaders returns the ComputedHeaders field if non-nil, zero value otherwise.

### GetComputedHeadersOk

`func (o *RESTDataSourceResultResponse) GetComputedHeadersOk() (*[]map[string]interface{}, bool)`

GetComputedHeadersOk returns a tuple with the ComputedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedHeaders

`func (o *RESTDataSourceResultResponse) SetComputedHeaders(v []map[string]interface{})`

SetComputedHeaders sets ComputedHeaders field to given value.

### HasComputedHeaders

`func (o *RESTDataSourceResultResponse) HasComputedHeaders() bool`

HasComputedHeaders returns a boolean if a field has been set.

### SetComputedHeadersNil

`func (o *RESTDataSourceResultResponse) SetComputedHeadersNil(b bool)`

 SetComputedHeadersNil sets the value for ComputedHeaders to be an explicit nil

### UnsetComputedHeaders
`func (o *RESTDataSourceResultResponse) UnsetComputedHeaders()`

UnsetComputedHeaders ensures that no value is present for ComputedHeaders, not even an explicit nil
### GetComputedPayload

`func (o *RESTDataSourceResultResponse) GetComputedPayload() string`

GetComputedPayload returns the ComputedPayload field if non-nil, zero value otherwise.

### GetComputedPayloadOk

`func (o *RESTDataSourceResultResponse) GetComputedPayloadOk() (*string, bool)`

GetComputedPayloadOk returns a tuple with the ComputedPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedPayload

`func (o *RESTDataSourceResultResponse) SetComputedPayload(v string)`

SetComputedPayload sets ComputedPayload field to given value.

### HasComputedPayload

`func (o *RESTDataSourceResultResponse) HasComputedPayload() bool`

HasComputedPayload returns a boolean if a field has been set.

### SetComputedPayloadNil

`func (o *RESTDataSourceResultResponse) SetComputedPayloadNil(b bool)`

 SetComputedPayloadNil sets the value for ComputedPayload to be an explicit nil

### UnsetComputedPayload
`func (o *RESTDataSourceResultResponse) UnsetComputedPayload()`

UnsetComputedPayload ensures that no value is present for ComputedPayload, not even an explicit nil
### GetComputedUrl

`func (o *RESTDataSourceResultResponse) GetComputedUrl() string`

GetComputedUrl returns the ComputedUrl field if non-nil, zero value otherwise.

### GetComputedUrlOk

`func (o *RESTDataSourceResultResponse) GetComputedUrlOk() (*string, bool)`

GetComputedUrlOk returns a tuple with the ComputedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedUrl

`func (o *RESTDataSourceResultResponse) SetComputedUrl(v string)`

SetComputedUrl sets ComputedUrl field to given value.

### HasComputedUrl

`func (o *RESTDataSourceResultResponse) HasComputedUrl() bool`

HasComputedUrl returns a boolean if a field has been set.

### SetComputedUrlNil

`func (o *RESTDataSourceResultResponse) SetComputedUrlNil(b bool)`

 SetComputedUrlNil sets the value for ComputedUrl to be an explicit nil

### UnsetComputedUrl
`func (o *RESTDataSourceResultResponse) UnsetComputedUrl()`

UnsetComputedUrl ensures that no value is present for ComputedUrl, not even an explicit nil
### GetResponseBody

`func (o *RESTDataSourceResultResponse) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *RESTDataSourceResultResponse) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *RESTDataSourceResultResponse) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *RESTDataSourceResultResponse) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *RESTDataSourceResultResponse) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *RESTDataSourceResultResponse) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseCode

`func (o *RESTDataSourceResultResponse) GetResponseCode() int64`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *RESTDataSourceResultResponse) GetResponseCodeOk() (*int64, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *RESTDataSourceResultResponse) SetResponseCode(v int64)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *RESTDataSourceResultResponse) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### SetResponseCodeNil

`func (o *RESTDataSourceResultResponse) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *RESTDataSourceResultResponse) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
### GetResponseHeaders

`func (o *RESTDataSourceResultResponse) GetResponseHeaders() []map[string]interface{}`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *RESTDataSourceResultResponse) GetResponseHeadersOk() (*[]map[string]interface{}, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *RESTDataSourceResultResponse) SetResponseHeaders(v []map[string]interface{})`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *RESTDataSourceResultResponse) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *RESTDataSourceResultResponse) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *RESTDataSourceResultResponse) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
### GetType

`func (o *RESTDataSourceResultResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RESTDataSourceResultResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RESTDataSourceResultResponse) SetType(v string)`

SetType sets Type field to given value.


### GetDictionary

`func (o *RESTDataSourceResultResponse) GetDictionary() []MapEntry`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *RESTDataSourceResultResponse) GetDictionaryOk() (*[]MapEntry, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *RESTDataSourceResultResponse) SetDictionary(v []MapEntry)`

SetDictionary sets Dictionary field to given value.


### GetError

`func (o *RESTDataSourceResultResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RESTDataSourceResultResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RESTDataSourceResultResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RESTDataSourceResultResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RESTDataSourceResultResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RESTDataSourceResultResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetName

`func (o *RESTDataSourceResultResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RESTDataSourceResultResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RESTDataSourceResultResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *RESTDataSourceResultResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RESTDataSourceResultResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RESTDataSourceResultResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


