# SecurityPrincipalInfoUpdate400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The http status code of the error. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Error** | **string** | The error code of the problem | 
**Message** | **string** | A short, human-readable summary of the problem type | 
**Title** | **string** | A short, human-readable summary of the problem type. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 

## Methods

### NewSecurityPrincipalInfoUpdate400Response

`func NewSecurityPrincipalInfoUpdate400Response(status int64, error_ string, message string, title string, ) *SecurityPrincipalInfoUpdate400Response`

NewSecurityPrincipalInfoUpdate400Response instantiates a new SecurityPrincipalInfoUpdate400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPrincipalInfoUpdate400ResponseWithDefaults

`func NewSecurityPrincipalInfoUpdate400ResponseWithDefaults() *SecurityPrincipalInfoUpdate400Response`

NewSecurityPrincipalInfoUpdate400ResponseWithDefaults instantiates a new SecurityPrincipalInfoUpdate400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SecurityPrincipalInfoUpdate400Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityPrincipalInfoUpdate400Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityPrincipalInfoUpdate400Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetError

`func (o *SecurityPrincipalInfoUpdate400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SecurityPrincipalInfoUpdate400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SecurityPrincipalInfoUpdate400Response) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *SecurityPrincipalInfoUpdate400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityPrincipalInfoUpdate400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityPrincipalInfoUpdate400Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *SecurityPrincipalInfoUpdate400Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SecurityPrincipalInfoUpdate400Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SecurityPrincipalInfoUpdate400Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *SecurityPrincipalInfoUpdate400Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SecurityPrincipalInfoUpdate400Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SecurityPrincipalInfoUpdate400Response) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SecurityPrincipalInfoUpdate400Response) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *SecurityPrincipalInfoUpdate400Response) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *SecurityPrincipalInfoUpdate400Response) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


