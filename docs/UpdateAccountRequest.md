# UpdateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompromisedAt** | Pointer to **int64** | Unix epoch time in milliseconds. Certificates issued after this date will be revoked. | [optional] 
**CompromissionReason** | Pointer to **NullableString** | One of: &#x60;unspecified&#x60;, &#x60;keycompromise&#x60;, &#x60;cacompromise&#x60;, &#x60;affiliationchange&#x60;, &#x60;superseded&#x60;, &#x60;cessationofoperation&#x60; | [optional] 
**Status** | [**AccountStatus**](AccountStatus.md) |  | 

## Methods

### NewUpdateAccountRequest

`func NewUpdateAccountRequest(status AccountStatus, ) *UpdateAccountRequest`

NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountRequestWithDefaults

`func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest`

NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompromisedAt

`func (o *UpdateAccountRequest) GetCompromisedAt() int64`

GetCompromisedAt returns the CompromisedAt field if non-nil, zero value otherwise.

### GetCompromisedAtOk

`func (o *UpdateAccountRequest) GetCompromisedAtOk() (*int64, bool)`

GetCompromisedAtOk returns a tuple with the CompromisedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedAt

`func (o *UpdateAccountRequest) SetCompromisedAt(v int64)`

SetCompromisedAt sets CompromisedAt field to given value.

### HasCompromisedAt

`func (o *UpdateAccountRequest) HasCompromisedAt() bool`

HasCompromisedAt returns a boolean if a field has been set.

### GetCompromissionReason

`func (o *UpdateAccountRequest) GetCompromissionReason() string`

GetCompromissionReason returns the CompromissionReason field if non-nil, zero value otherwise.

### GetCompromissionReasonOk

`func (o *UpdateAccountRequest) GetCompromissionReasonOk() (*string, bool)`

GetCompromissionReasonOk returns a tuple with the CompromissionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromissionReason

`func (o *UpdateAccountRequest) SetCompromissionReason(v string)`

SetCompromissionReason sets CompromissionReason field to given value.

### HasCompromissionReason

`func (o *UpdateAccountRequest) HasCompromissionReason() bool`

HasCompromissionReason returns a boolean if a field has been set.

### SetCompromissionReasonNil

`func (o *UpdateAccountRequest) SetCompromissionReasonNil(b bool)`

 SetCompromissionReasonNil sets the value for CompromissionReason to be an explicit nil

### UnsetCompromissionReason
`func (o *UpdateAccountRequest) UnsetCompromissionReason()`

UnsetCompromissionReason ensures that no value is present for CompromissionReason, not even an explicit nil
### GetStatus

`func (o *UpdateAccountRequest) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAccountRequest) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAccountRequest) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


