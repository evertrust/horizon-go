# RequestSubmit500ResponseOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | The error code of the problem | 
**Message** | **string** | A short, human-readable summary of the problem type | 
**Title** | **string** | A short, human-readable summary of the problem type. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 
**Status** | **float32** |  | 

## Methods

### NewRequestSubmit500ResponseOneOf4

`func NewRequestSubmit500ResponseOneOf4(error_ string, message string, title string, status float32, ) *RequestSubmit500ResponseOneOf4`

NewRequestSubmit500ResponseOneOf4 instantiates a new RequestSubmit500ResponseOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSubmit500ResponseOneOf4WithDefaults

`func NewRequestSubmit500ResponseOneOf4WithDefaults() *RequestSubmit500ResponseOneOf4`

NewRequestSubmit500ResponseOneOf4WithDefaults instantiates a new RequestSubmit500ResponseOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RequestSubmit500ResponseOneOf4) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RequestSubmit500ResponseOneOf4) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RequestSubmit500ResponseOneOf4) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *RequestSubmit500ResponseOneOf4) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RequestSubmit500ResponseOneOf4) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RequestSubmit500ResponseOneOf4) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *RequestSubmit500ResponseOneOf4) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestSubmit500ResponseOneOf4) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestSubmit500ResponseOneOf4) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *RequestSubmit500ResponseOneOf4) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RequestSubmit500ResponseOneOf4) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RequestSubmit500ResponseOneOf4) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *RequestSubmit500ResponseOneOf4) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *RequestSubmit500ResponseOneOf4) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *RequestSubmit500ResponseOneOf4) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetStatus

`func (o *RequestSubmit500ResponseOneOf4) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestSubmit500ResponseOneOf4) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestSubmit500ResponseOneOf4) SetStatus(v float32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


