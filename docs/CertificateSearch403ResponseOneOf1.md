# CertificateSearch403ResponseOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | The error code of the problem | 
**Message** | **string** | A short, human-readable summary of the problem type | 
**Title** | **string** | A short, human-readable summary of the problem type. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 
**Status** | **float32** |  | 

## Methods

### NewCertificateSearch403ResponseOneOf1

`func NewCertificateSearch403ResponseOneOf1(error_ string, message string, title string, status float32, ) *CertificateSearch403ResponseOneOf1`

NewCertificateSearch403ResponseOneOf1 instantiates a new CertificateSearch403ResponseOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSearch403ResponseOneOf1WithDefaults

`func NewCertificateSearch403ResponseOneOf1WithDefaults() *CertificateSearch403ResponseOneOf1`

NewCertificateSearch403ResponseOneOf1WithDefaults instantiates a new CertificateSearch403ResponseOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CertificateSearch403ResponseOneOf1) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CertificateSearch403ResponseOneOf1) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CertificateSearch403ResponseOneOf1) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *CertificateSearch403ResponseOneOf1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CertificateSearch403ResponseOneOf1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CertificateSearch403ResponseOneOf1) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *CertificateSearch403ResponseOneOf1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CertificateSearch403ResponseOneOf1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CertificateSearch403ResponseOneOf1) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *CertificateSearch403ResponseOneOf1) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CertificateSearch403ResponseOneOf1) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CertificateSearch403ResponseOneOf1) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CertificateSearch403ResponseOneOf1) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *CertificateSearch403ResponseOneOf1) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *CertificateSearch403ResponseOneOf1) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetStatus

`func (o *CertificateSearch403ResponseOneOf1) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateSearch403ResponseOneOf1) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateSearch403ResponseOneOf1) SetStatus(v float32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


