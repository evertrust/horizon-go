# ScheduledTaskExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**LastCompletionDate** | Pointer to **int64** |  | [optional] 
**LastExecutionDate** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewScheduledTaskExecution

`func NewScheduledTaskExecution() *ScheduledTaskExecution`

NewScheduledTaskExecution instantiates a new ScheduledTaskExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledTaskExecutionWithDefaults

`func NewScheduledTaskExecutionWithDefaults() *ScheduledTaskExecution`

NewScheduledTaskExecutionWithDefaults instantiates a new ScheduledTaskExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ScheduledTaskExecution) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ScheduledTaskExecution) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ScheduledTaskExecution) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ScheduledTaskExecution) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetExecutionId

`func (o *ScheduledTaskExecution) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ScheduledTaskExecution) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ScheduledTaskExecution) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ScheduledTaskExecution) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetHost

`func (o *ScheduledTaskExecution) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ScheduledTaskExecution) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ScheduledTaskExecution) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ScheduledTaskExecution) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetLastCompletionDate

`func (o *ScheduledTaskExecution) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *ScheduledTaskExecution) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *ScheduledTaskExecution) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *ScheduledTaskExecution) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### GetLastExecutionDate

`func (o *ScheduledTaskExecution) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *ScheduledTaskExecution) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *ScheduledTaskExecution) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *ScheduledTaskExecution) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduledTaskExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledTaskExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledTaskExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduledTaskExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


