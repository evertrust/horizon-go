# RestTriggerTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A message describing the test | 
**RequestBody** | Pointer to **NullableString** | The body from this request | [optional] 
**RequestHeaders** | Pointer to [**[]RESTHeader**](RESTHeader.md) | The headers from this request | [optional] 
**RequestURL** | **string** | The URL requested | 
**ResponseBody** | Pointer to **NullableString** | The body from the response to this request | [optional] 
**ResponseCode** | Pointer to **NullableInt64** | The response code to this request | [optional] 
**ResponseHeaders** | Pointer to [**[]RESTHeader**](RESTHeader.md) | The headers from the response to this request | [optional] 
**Status** | **string** | Status of the test | 
**Type** | **string** | The trigger type that was executed | 

## Methods

### NewRestTriggerTestResponse

`func NewRestTriggerTestResponse(message string, requestURL string, status string, type_ string, ) *RestTriggerTestResponse`

NewRestTriggerTestResponse instantiates a new RestTriggerTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestTriggerTestResponseWithDefaults

`func NewRestTriggerTestResponseWithDefaults() *RestTriggerTestResponse`

NewRestTriggerTestResponseWithDefaults instantiates a new RestTriggerTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RestTriggerTestResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RestTriggerTestResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RestTriggerTestResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRequestBody

`func (o *RestTriggerTestResponse) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *RestTriggerTestResponse) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *RestTriggerTestResponse) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *RestTriggerTestResponse) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### SetRequestBodyNil

`func (o *RestTriggerTestResponse) SetRequestBodyNil(b bool)`

 SetRequestBodyNil sets the value for RequestBody to be an explicit nil

### UnsetRequestBody
`func (o *RestTriggerTestResponse) UnsetRequestBody()`

UnsetRequestBody ensures that no value is present for RequestBody, not even an explicit nil
### GetRequestHeaders

`func (o *RestTriggerTestResponse) GetRequestHeaders() []RESTHeader`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *RestTriggerTestResponse) GetRequestHeadersOk() (*[]RESTHeader, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *RestTriggerTestResponse) SetRequestHeaders(v []RESTHeader)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *RestTriggerTestResponse) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *RestTriggerTestResponse) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *RestTriggerTestResponse) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetRequestURL

`func (o *RestTriggerTestResponse) GetRequestURL() string`

GetRequestURL returns the RequestURL field if non-nil, zero value otherwise.

### GetRequestURLOk

`func (o *RestTriggerTestResponse) GetRequestURLOk() (*string, bool)`

GetRequestURLOk returns a tuple with the RequestURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestURL

`func (o *RestTriggerTestResponse) SetRequestURL(v string)`

SetRequestURL sets RequestURL field to given value.


### GetResponseBody

`func (o *RestTriggerTestResponse) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *RestTriggerTestResponse) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *RestTriggerTestResponse) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *RestTriggerTestResponse) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *RestTriggerTestResponse) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *RestTriggerTestResponse) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseCode

`func (o *RestTriggerTestResponse) GetResponseCode() int64`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *RestTriggerTestResponse) GetResponseCodeOk() (*int64, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *RestTriggerTestResponse) SetResponseCode(v int64)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *RestTriggerTestResponse) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### SetResponseCodeNil

`func (o *RestTriggerTestResponse) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *RestTriggerTestResponse) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
### GetResponseHeaders

`func (o *RestTriggerTestResponse) GetResponseHeaders() []RESTHeader`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *RestTriggerTestResponse) GetResponseHeadersOk() (*[]RESTHeader, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *RestTriggerTestResponse) SetResponseHeaders(v []RESTHeader)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *RestTriggerTestResponse) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *RestTriggerTestResponse) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *RestTriggerTestResponse) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
### GetStatus

`func (o *RestTriggerTestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestTriggerTestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestTriggerTestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *RestTriggerTestResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestTriggerTestResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestTriggerTestResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


