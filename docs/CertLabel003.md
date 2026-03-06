# CertLabel003

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Message** | **string** |  | 
**Title** | **string** |  | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 
**Status** | **int64** | The http status code of the error. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 

## Methods

### NewCertLabel003

`func NewCertLabel003(error_ string, message string, title string, status int64, ) *CertLabel003`

NewCertLabel003 instantiates a new CertLabel003 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertLabel003WithDefaults

`func NewCertLabel003WithDefaults() *CertLabel003`

NewCertLabel003WithDefaults instantiates a new CertLabel003 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CertLabel003) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CertLabel003) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CertLabel003) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *CertLabel003) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CertLabel003) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CertLabel003) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *CertLabel003) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CertLabel003) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CertLabel003) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *CertLabel003) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CertLabel003) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CertLabel003) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CertLabel003) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *CertLabel003) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *CertLabel003) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetStatus

`func (o *CertLabel003) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertLabel003) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertLabel003) SetStatus(v int64)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


