# ExecutionPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | Pointer to [**NullableDateRange**](DateRange.md) |  | [optional] 
**Weeks** | Pointer to **[]int64** |  | [optional] 
**WeekDays** | Pointer to **[]string** |  | [optional] 
**TimeRange** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExecutionPeriod

`func NewExecutionPeriod() *ExecutionPeriod`

NewExecutionPeriod instantiates a new ExecutionPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionPeriodWithDefaults

`func NewExecutionPeriodWithDefaults() *ExecutionPeriod`

NewExecutionPeriodWithDefaults instantiates a new ExecutionPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *ExecutionPeriod) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *ExecutionPeriod) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *ExecutionPeriod) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *ExecutionPeriod) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### SetDateRangeNil

`func (o *ExecutionPeriod) SetDateRangeNil(b bool)`

 SetDateRangeNil sets the value for DateRange to be an explicit nil

### UnsetDateRange
`func (o *ExecutionPeriod) UnsetDateRange()`

UnsetDateRange ensures that no value is present for DateRange, not even an explicit nil
### GetWeeks

`func (o *ExecutionPeriod) GetWeeks() []int64`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *ExecutionPeriod) GetWeeksOk() (*[]int64, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *ExecutionPeriod) SetWeeks(v []int64)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *ExecutionPeriod) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.

### SetWeeksNil

`func (o *ExecutionPeriod) SetWeeksNil(b bool)`

 SetWeeksNil sets the value for Weeks to be an explicit nil

### UnsetWeeks
`func (o *ExecutionPeriod) UnsetWeeks()`

UnsetWeeks ensures that no value is present for Weeks, not even an explicit nil
### GetWeekDays

`func (o *ExecutionPeriod) GetWeekDays() []string`

GetWeekDays returns the WeekDays field if non-nil, zero value otherwise.

### GetWeekDaysOk

`func (o *ExecutionPeriod) GetWeekDaysOk() (*[]string, bool)`

GetWeekDaysOk returns a tuple with the WeekDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDays

`func (o *ExecutionPeriod) SetWeekDays(v []string)`

SetWeekDays sets WeekDays field to given value.

### HasWeekDays

`func (o *ExecutionPeriod) HasWeekDays() bool`

HasWeekDays returns a boolean if a field has been set.

### SetWeekDaysNil

`func (o *ExecutionPeriod) SetWeekDaysNil(b bool)`

 SetWeekDaysNil sets the value for WeekDays to be an explicit nil

### UnsetWeekDays
`func (o *ExecutionPeriod) UnsetWeekDays()`

UnsetWeekDays ensures that no value is present for WeekDays, not even an explicit nil
### GetTimeRange

`func (o *ExecutionPeriod) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *ExecutionPeriod) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *ExecutionPeriod) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *ExecutionPeriod) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### SetTimeRangeNil

`func (o *ExecutionPeriod) SetTimeRangeNil(b bool)`

 SetTimeRangeNil sets the value for TimeRange to be an explicit nil

### UnsetTimeRange
`func (o *ExecutionPeriod) UnsetTimeRange()`

UnsetTimeRange ensures that no value is present for TimeRange, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


