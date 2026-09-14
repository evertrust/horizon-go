# EabStatusUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompromisedAt** | Pointer to **int64** | Unix epoch time in milliseconds. Certificates issued after this date will be revoked. | [optional] 
**CompromissionReason** | Pointer to **NullableString** | One of: &#x60;unspecified&#x60;, &#x60;keycompromise&#x60;, &#x60;cacompromise&#x60;, &#x60;affiliationchange&#x60;, &#x60;superseded&#x60;, &#x60;cessationofoperation&#x60; | [optional] 
**Status** | [**ExternalAccountBindingStatus**](ExternalAccountBindingStatus.md) |  | 

## Methods

### NewEabStatusUpdateRequest

`func NewEabStatusUpdateRequest(status ExternalAccountBindingStatus, ) *EabStatusUpdateRequest`

NewEabStatusUpdateRequest instantiates a new EabStatusUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEabStatusUpdateRequestWithDefaults

`func NewEabStatusUpdateRequestWithDefaults() *EabStatusUpdateRequest`

NewEabStatusUpdateRequestWithDefaults instantiates a new EabStatusUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompromisedAt

`func (o *EabStatusUpdateRequest) GetCompromisedAt() int64`

GetCompromisedAt returns the CompromisedAt field if non-nil, zero value otherwise.

### GetCompromisedAtOk

`func (o *EabStatusUpdateRequest) GetCompromisedAtOk() (*int64, bool)`

GetCompromisedAtOk returns a tuple with the CompromisedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedAt

`func (o *EabStatusUpdateRequest) SetCompromisedAt(v int64)`

SetCompromisedAt sets CompromisedAt field to given value.

### HasCompromisedAt

`func (o *EabStatusUpdateRequest) HasCompromisedAt() bool`

HasCompromisedAt returns a boolean if a field has been set.

### GetCompromissionReason

`func (o *EabStatusUpdateRequest) GetCompromissionReason() string`

GetCompromissionReason returns the CompromissionReason field if non-nil, zero value otherwise.

### GetCompromissionReasonOk

`func (o *EabStatusUpdateRequest) GetCompromissionReasonOk() (*string, bool)`

GetCompromissionReasonOk returns a tuple with the CompromissionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromissionReason

`func (o *EabStatusUpdateRequest) SetCompromissionReason(v string)`

SetCompromissionReason sets CompromissionReason field to given value.

### HasCompromissionReason

`func (o *EabStatusUpdateRequest) HasCompromissionReason() bool`

HasCompromissionReason returns a boolean if a field has been set.

### SetCompromissionReasonNil

`func (o *EabStatusUpdateRequest) SetCompromissionReasonNil(b bool)`

 SetCompromissionReasonNil sets the value for CompromissionReason to be an explicit nil

### UnsetCompromissionReason
`func (o *EabStatusUpdateRequest) UnsetCompromissionReason()`

UnsetCompromissionReason ensures that no value is present for CompromissionReason, not even an explicit nil
### GetStatus

`func (o *EabStatusUpdateRequest) GetStatus() ExternalAccountBindingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EabStatusUpdateRequest) GetStatusOk() (*ExternalAccountBindingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EabStatusUpdateRequest) SetStatus(v ExternalAccountBindingStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


