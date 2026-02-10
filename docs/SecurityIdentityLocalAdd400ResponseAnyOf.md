# SecurityIdentityLocalAdd400ResponseAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **float32** |  | 
**Error** | **string** | The error code of the problem | 
**Message** | **string** | A short, human-readable summary of the problem type | 
**Title** | **string** | A short, human-readable summary of the problem type. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 

## Methods

### NewSecurityIdentityLocalAdd400ResponseAnyOf

`func NewSecurityIdentityLocalAdd400ResponseAnyOf(status float32, error_ string, message string, title string, ) *SecurityIdentityLocalAdd400ResponseAnyOf`

NewSecurityIdentityLocalAdd400ResponseAnyOf instantiates a new SecurityIdentityLocalAdd400ResponseAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIdentityLocalAdd400ResponseAnyOfWithDefaults

`func NewSecurityIdentityLocalAdd400ResponseAnyOfWithDefaults() *SecurityIdentityLocalAdd400ResponseAnyOf`

NewSecurityIdentityLocalAdd400ResponseAnyOfWithDefaults instantiates a new SecurityIdentityLocalAdd400ResponseAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetError

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *SecurityIdentityLocalAdd400ResponseAnyOf) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


