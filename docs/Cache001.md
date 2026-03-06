# Cache001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Message** | **string** |  | 
**Title** | **string** |  | 
**Detail** | Pointer to **NullableString** | A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | [optional] 
**Status** | **int64** | The http status code of the error. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807) | 

## Methods

### NewCache001

`func NewCache001(error_ string, message string, title string, status int64, ) *Cache001`

NewCache001 instantiates a new Cache001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCache001WithDefaults

`func NewCache001WithDefaults() *Cache001`

NewCache001WithDefaults instantiates a new Cache001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Cache001) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Cache001) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Cache001) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *Cache001) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Cache001) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Cache001) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *Cache001) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Cache001) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Cache001) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *Cache001) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Cache001) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Cache001) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Cache001) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *Cache001) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *Cache001) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetStatus

`func (o *Cache001) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cache001) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cache001) SetStatus(v int64)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


