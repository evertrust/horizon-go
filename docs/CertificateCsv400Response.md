# CertificateCsv400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | The error code of the problem | 
**Message** | **string** | A short, human-readable summary of the problem type | 
**Title** | **string** | A short, human-readable summary of the problem type. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 
**Status** | **int64** | The http status code of the error. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 

## Methods

### NewCertificateCsv400Response

`func NewCertificateCsv400Response(error_ string, message string, title string, status int64, ) *CertificateCsv400Response`

NewCertificateCsv400Response instantiates a new CertificateCsv400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCsv400ResponseWithDefaults

`func NewCertificateCsv400ResponseWithDefaults() *CertificateCsv400Response`

NewCertificateCsv400ResponseWithDefaults instantiates a new CertificateCsv400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CertificateCsv400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CertificateCsv400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CertificateCsv400Response) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *CertificateCsv400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CertificateCsv400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CertificateCsv400Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *CertificateCsv400Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CertificateCsv400Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CertificateCsv400Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *CertificateCsv400Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CertificateCsv400Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CertificateCsv400Response) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CertificateCsv400Response) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *CertificateCsv400Response) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *CertificateCsv400Response) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetStatus

`func (o *CertificateCsv400Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateCsv400Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateCsv400Response) SetStatus(v int64)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


