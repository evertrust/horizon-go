# RequestTemplate400ResponseOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **float32** |  | 
**Error** | **string** | The error code of the problem | 
**Message** | **string** | A short, human-readable summary of the problem type | 
**Title** | **string** | A short, human-readable summary of the problem type. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 

## Methods

### NewRequestTemplate400ResponseOneOf3

`func NewRequestTemplate400ResponseOneOf3(status float32, error_ string, message string, title string, ) *RequestTemplate400ResponseOneOf3`

NewRequestTemplate400ResponseOneOf3 instantiates a new RequestTemplate400ResponseOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTemplate400ResponseOneOf3WithDefaults

`func NewRequestTemplate400ResponseOneOf3WithDefaults() *RequestTemplate400ResponseOneOf3`

NewRequestTemplate400ResponseOneOf3WithDefaults instantiates a new RequestTemplate400ResponseOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestTemplate400ResponseOneOf3) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestTemplate400ResponseOneOf3) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestTemplate400ResponseOneOf3) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetError

`func (o *RequestTemplate400ResponseOneOf3) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RequestTemplate400ResponseOneOf3) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RequestTemplate400ResponseOneOf3) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *RequestTemplate400ResponseOneOf3) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RequestTemplate400ResponseOneOf3) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RequestTemplate400ResponseOneOf3) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *RequestTemplate400ResponseOneOf3) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestTemplate400ResponseOneOf3) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestTemplate400ResponseOneOf3) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *RequestTemplate400ResponseOneOf3) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RequestTemplate400ResponseOneOf3) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RequestTemplate400ResponseOneOf3) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *RequestTemplate400ResponseOneOf3) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *RequestTemplate400ResponseOneOf3) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *RequestTemplate400ResponseOneOf3) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


