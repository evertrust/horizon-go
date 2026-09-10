# DCVLifecycleEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int64** | Current retry attempt count | [optional] 
**Domain** | **string** | The domain | 
**LastError** | Pointer to **NullableString** | Error message from the previous attempt | [optional] 
**Msg** | Pointer to **string** | Error message describing the failure cause | [optional] 
**Policy** | **string** | The DCV policy name | 
**RemoveAt** | **time.Time** | Timestamp when the event will be removed | 
**Status** | **string** |  | 
**Timestamp** | **time.Time** | Timestamp linked to the validation | 

## Methods

### NewDCVLifecycleEvent

`func NewDCVLifecycleEvent(domain string, policy string, removeAt time.Time, status string, timestamp time.Time, ) *DCVLifecycleEvent`

NewDCVLifecycleEvent instantiates a new DCVLifecycleEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVLifecycleEventWithDefaults

`func NewDCVLifecycleEventWithDefaults() *DCVLifecycleEvent`

NewDCVLifecycleEventWithDefaults instantiates a new DCVLifecycleEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *DCVLifecycleEvent) GetAttempt() int64`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *DCVLifecycleEvent) GetAttemptOk() (*int64, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *DCVLifecycleEvent) SetAttempt(v int64)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *DCVLifecycleEvent) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetDomain

`func (o *DCVLifecycleEvent) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DCVLifecycleEvent) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DCVLifecycleEvent) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetLastError

`func (o *DCVLifecycleEvent) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *DCVLifecycleEvent) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *DCVLifecycleEvent) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *DCVLifecycleEvent) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *DCVLifecycleEvent) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *DCVLifecycleEvent) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetMsg

`func (o *DCVLifecycleEvent) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DCVLifecycleEvent) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DCVLifecycleEvent) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *DCVLifecycleEvent) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetPolicy

`func (o *DCVLifecycleEvent) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *DCVLifecycleEvent) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *DCVLifecycleEvent) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetRemoveAt

`func (o *DCVLifecycleEvent) GetRemoveAt() time.Time`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *DCVLifecycleEvent) GetRemoveAtOk() (*time.Time, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *DCVLifecycleEvent) SetRemoveAt(v time.Time)`

SetRemoveAt sets RemoveAt field to given value.


### GetStatus

`func (o *DCVLifecycleEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DCVLifecycleEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DCVLifecycleEvent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *DCVLifecycleEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DCVLifecycleEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DCVLifecycleEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


