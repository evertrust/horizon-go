# SecurityIdentityLocalPasswordSet400ResponseAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **float32** |  | 
**Error** | **string** | The error code of the problem | 
**Message** | **string** | A short, human-readable summary of the problem type | 
**Title** | **string** | A short, human-readable summary of the problem type. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 

## Methods

### NewSecurityIdentityLocalPasswordSet400ResponseAnyOf

`func NewSecurityIdentityLocalPasswordSet400ResponseAnyOf(status float32, error_ string, message string, title string, ) *SecurityIdentityLocalPasswordSet400ResponseAnyOf`

NewSecurityIdentityLocalPasswordSet400ResponseAnyOf instantiates a new SecurityIdentityLocalPasswordSet400ResponseAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIdentityLocalPasswordSet400ResponseAnyOfWithDefaults

`func NewSecurityIdentityLocalPasswordSet400ResponseAnyOfWithDefaults() *SecurityIdentityLocalPasswordSet400ResponseAnyOf`

NewSecurityIdentityLocalPasswordSet400ResponseAnyOfWithDefaults instantiates a new SecurityIdentityLocalPasswordSet400ResponseAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetError

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *SecurityIdentityLocalPasswordSet400ResponseAnyOf) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


