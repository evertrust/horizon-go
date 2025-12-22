# TriggerResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **NullableString** | Contains details on this trigger&#39;s execution | [optional] 
**Event** | **string** | The event that triggered the trigger | 
**LastExecutionDate** | **int64** | The last time this trigger was executed for this certificate and this event | 
**Name** | **string** | The name of the trigger that was executed | 
**NextDelay** | Pointer to **NullableString** | Time that will be waited between the next and the next+1 execution of this trigger | [optional] 
**NextExecutionDate** | Pointer to **NullableInt64** | The next scheduled execution time for this trigger | [optional] 
**Retries** | Pointer to **NullableInt64** | The number of remaining tries before the trigger is abandoned | [optional] 
**Retryable** | **bool** | Is this trigger manually retryable (can be [run](#tag/certificate/operation/certificate.run)) | 
**Status** | **string** | The status of the trigger after its execution | 
**TriggerType** | **string** | The type of the trigger | 

## Methods

### NewTriggerResult

`func NewTriggerResult(event string, lastExecutionDate int64, name string, retryable bool, status string, triggerType string, ) *TriggerResult`

NewTriggerResult instantiates a new TriggerResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerResultWithDefaults

`func NewTriggerResultWithDefaults() *TriggerResult`

NewTriggerResultWithDefaults instantiates a new TriggerResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *TriggerResult) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *TriggerResult) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *TriggerResult) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *TriggerResult) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *TriggerResult) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *TriggerResult) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetEvent

`func (o *TriggerResult) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TriggerResult) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TriggerResult) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetLastExecutionDate

`func (o *TriggerResult) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *TriggerResult) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *TriggerResult) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.


### GetName

`func (o *TriggerResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerResult) SetName(v string)`

SetName sets Name field to given value.


### GetNextDelay

`func (o *TriggerResult) GetNextDelay() string`

GetNextDelay returns the NextDelay field if non-nil, zero value otherwise.

### GetNextDelayOk

`func (o *TriggerResult) GetNextDelayOk() (*string, bool)`

GetNextDelayOk returns a tuple with the NextDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDelay

`func (o *TriggerResult) SetNextDelay(v string)`

SetNextDelay sets NextDelay field to given value.

### HasNextDelay

`func (o *TriggerResult) HasNextDelay() bool`

HasNextDelay returns a boolean if a field has been set.

### SetNextDelayNil

`func (o *TriggerResult) SetNextDelayNil(b bool)`

 SetNextDelayNil sets the value for NextDelay to be an explicit nil

### UnsetNextDelay
`func (o *TriggerResult) UnsetNextDelay()`

UnsetNextDelay ensures that no value is present for NextDelay, not even an explicit nil
### GetNextExecutionDate

`func (o *TriggerResult) GetNextExecutionDate() int64`

GetNextExecutionDate returns the NextExecutionDate field if non-nil, zero value otherwise.

### GetNextExecutionDateOk

`func (o *TriggerResult) GetNextExecutionDateOk() (*int64, bool)`

GetNextExecutionDateOk returns a tuple with the NextExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecutionDate

`func (o *TriggerResult) SetNextExecutionDate(v int64)`

SetNextExecutionDate sets NextExecutionDate field to given value.

### HasNextExecutionDate

`func (o *TriggerResult) HasNextExecutionDate() bool`

HasNextExecutionDate returns a boolean if a field has been set.

### SetNextExecutionDateNil

`func (o *TriggerResult) SetNextExecutionDateNil(b bool)`

 SetNextExecutionDateNil sets the value for NextExecutionDate to be an explicit nil

### UnsetNextExecutionDate
`func (o *TriggerResult) UnsetNextExecutionDate()`

UnsetNextExecutionDate ensures that no value is present for NextExecutionDate, not even an explicit nil
### GetRetries

`func (o *TriggerResult) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *TriggerResult) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *TriggerResult) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *TriggerResult) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *TriggerResult) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *TriggerResult) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetRetryable

`func (o *TriggerResult) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *TriggerResult) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *TriggerResult) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.


### GetStatus

`func (o *TriggerResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TriggerResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TriggerResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTriggerType

`func (o *TriggerResult) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *TriggerResult) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *TriggerResult) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


