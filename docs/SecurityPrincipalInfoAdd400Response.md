# SecurityPrincipalInfoAdd400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The http status code of the error. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Error** | **string** | The error code of the problem | 
**Message** | **string** | A short, human-readable summary of the problem type | 
**Title** | **string** | A short, human-readable summary of the problem type. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 

## Methods

### NewSecurityPrincipalInfoAdd400Response

`func NewSecurityPrincipalInfoAdd400Response(status int64, error_ string, message string, title string, ) *SecurityPrincipalInfoAdd400Response`

NewSecurityPrincipalInfoAdd400Response instantiates a new SecurityPrincipalInfoAdd400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPrincipalInfoAdd400ResponseWithDefaults

`func NewSecurityPrincipalInfoAdd400ResponseWithDefaults() *SecurityPrincipalInfoAdd400Response`

NewSecurityPrincipalInfoAdd400ResponseWithDefaults instantiates a new SecurityPrincipalInfoAdd400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SecurityPrincipalInfoAdd400Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityPrincipalInfoAdd400Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityPrincipalInfoAdd400Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetError

`func (o *SecurityPrincipalInfoAdd400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SecurityPrincipalInfoAdd400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SecurityPrincipalInfoAdd400Response) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *SecurityPrincipalInfoAdd400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityPrincipalInfoAdd400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityPrincipalInfoAdd400Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *SecurityPrincipalInfoAdd400Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SecurityPrincipalInfoAdd400Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SecurityPrincipalInfoAdd400Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *SecurityPrincipalInfoAdd400Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SecurityPrincipalInfoAdd400Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SecurityPrincipalInfoAdd400Response) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SecurityPrincipalInfoAdd400Response) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *SecurityPrincipalInfoAdd400Response) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *SecurityPrincipalInfoAdd400Response) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


