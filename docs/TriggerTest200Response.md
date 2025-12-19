# TriggerTest200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A message describing the test | 
**Status** | **string** | Status of the test | 
**RequestBody** | Pointer to **NullableString** | The body from this request | [optional] 
**RequestHeaders** | Pointer to [**[]RESTHeader**](RESTHeader.md) | The headers from this request | [optional] 
**RequestURL** | **string** | The URL requested | 
**ResponseBody** | Pointer to **NullableString** | The body from the response to this request | [optional] 
**ResponseCode** | Pointer to **NullableInt64** | The response code to this request | [optional] 
**ResponseHeaders** | Pointer to [**[]RESTHeader**](RESTHeader.md) | The headers from the response to this request | [optional] 
**Type** | **string** | The trigger type that was executed | 

## Methods

### NewTriggerTest200Response

`func NewTriggerTest200Response(message string, status string, requestURL string, type_ string, ) *TriggerTest200Response`

NewTriggerTest200Response instantiates a new TriggerTest200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerTest200ResponseWithDefaults

`func NewTriggerTest200ResponseWithDefaults() *TriggerTest200Response`

NewTriggerTest200ResponseWithDefaults instantiates a new TriggerTest200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TriggerTest200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TriggerTest200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TriggerTest200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *TriggerTest200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TriggerTest200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TriggerTest200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRequestBody

`func (o *TriggerTest200Response) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *TriggerTest200Response) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *TriggerTest200Response) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *TriggerTest200Response) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### SetRequestBodyNil

`func (o *TriggerTest200Response) SetRequestBodyNil(b bool)`

 SetRequestBodyNil sets the value for RequestBody to be an explicit nil

### UnsetRequestBody
`func (o *TriggerTest200Response) UnsetRequestBody()`

UnsetRequestBody ensures that no value is present for RequestBody, not even an explicit nil
### GetRequestHeaders

`func (o *TriggerTest200Response) GetRequestHeaders() []RESTHeader`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *TriggerTest200Response) GetRequestHeadersOk() (*[]RESTHeader, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *TriggerTest200Response) SetRequestHeaders(v []RESTHeader)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *TriggerTest200Response) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *TriggerTest200Response) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *TriggerTest200Response) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetRequestURL

`func (o *TriggerTest200Response) GetRequestURL() string`

GetRequestURL returns the RequestURL field if non-nil, zero value otherwise.

### GetRequestURLOk

`func (o *TriggerTest200Response) GetRequestURLOk() (*string, bool)`

GetRequestURLOk returns a tuple with the RequestURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestURL

`func (o *TriggerTest200Response) SetRequestURL(v string)`

SetRequestURL sets RequestURL field to given value.


### GetResponseBody

`func (o *TriggerTest200Response) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *TriggerTest200Response) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *TriggerTest200Response) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *TriggerTest200Response) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *TriggerTest200Response) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *TriggerTest200Response) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseCode

`func (o *TriggerTest200Response) GetResponseCode() int64`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *TriggerTest200Response) GetResponseCodeOk() (*int64, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *TriggerTest200Response) SetResponseCode(v int64)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *TriggerTest200Response) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### SetResponseCodeNil

`func (o *TriggerTest200Response) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *TriggerTest200Response) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
### GetResponseHeaders

`func (o *TriggerTest200Response) GetResponseHeaders() []RESTHeader`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *TriggerTest200Response) GetResponseHeadersOk() (*[]RESTHeader, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *TriggerTest200Response) SetResponseHeaders(v []RESTHeader)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *TriggerTest200Response) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *TriggerTest200Response) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *TriggerTest200Response) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
### GetType

`func (o *TriggerTest200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerTest200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerTest200Response) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


