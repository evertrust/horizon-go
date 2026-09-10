# DCVPolicyStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainsStatus** | [**DCVDomainsStatus**](DCVDomainsStatus.md) | Domain validation statuses, with an optional error if the provider could not be reached | 
**Enabled** | **bool** | Whether the DCV policy is enabled | 
**ExecutionTimeout** | **NullableString** | Maximum duration allowed for a single DCV run | 
**ExecutionTimeoutAt** | Pointer to **NullableInt64** | Epoch milliseconds at which the current execution will be forcefully ended, only present when status is running | [optional] 
**Name** | **string** | Unique name of the DCV policy | 
**NextCheckAt** | Pointer to **NullableInt64** | Epoch milliseconds of the next retry check, only present when status is running | [optional] 
**RenewalPeriod** | Pointer to **NullableString** | Duration before expiry at which renewal is triggered; absent if no renewal policy is set | [optional] 
**RetryDelay** | **NullableString** | Delay between retry attempts | 
**Runnable** | **bool** | Whether the policy is enabled and the principal has manage permission | 
**StartedAt** | Pointer to **NullableInt64** | Epoch milliseconds when the current execution started, only present when status is running | [optional] 
**Status** | **string** | Current status of the DCV policy | 

## Methods

### NewDCVPolicyStatusResponse

`func NewDCVPolicyStatusResponse(domainsStatus DCVDomainsStatus, enabled bool, executionTimeout NullableString, name string, retryDelay NullableString, runnable bool, status string, ) *DCVPolicyStatusResponse`

NewDCVPolicyStatusResponse instantiates a new DCVPolicyStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVPolicyStatusResponseWithDefaults

`func NewDCVPolicyStatusResponseWithDefaults() *DCVPolicyStatusResponse`

NewDCVPolicyStatusResponseWithDefaults instantiates a new DCVPolicyStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainsStatus

`func (o *DCVPolicyStatusResponse) GetDomainsStatus() DCVDomainsStatus`

GetDomainsStatus returns the DomainsStatus field if non-nil, zero value otherwise.

### GetDomainsStatusOk

`func (o *DCVPolicyStatusResponse) GetDomainsStatusOk() (*DCVDomainsStatus, bool)`

GetDomainsStatusOk returns a tuple with the DomainsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsStatus

`func (o *DCVPolicyStatusResponse) SetDomainsStatus(v DCVDomainsStatus)`

SetDomainsStatus sets DomainsStatus field to given value.


### GetEnabled

`func (o *DCVPolicyStatusResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DCVPolicyStatusResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DCVPolicyStatusResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExecutionTimeout

`func (o *DCVPolicyStatusResponse) GetExecutionTimeout() string`

GetExecutionTimeout returns the ExecutionTimeout field if non-nil, zero value otherwise.

### GetExecutionTimeoutOk

`func (o *DCVPolicyStatusResponse) GetExecutionTimeoutOk() (*string, bool)`

GetExecutionTimeoutOk returns a tuple with the ExecutionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeout

`func (o *DCVPolicyStatusResponse) SetExecutionTimeout(v string)`

SetExecutionTimeout sets ExecutionTimeout field to given value.


### SetExecutionTimeoutNil

`func (o *DCVPolicyStatusResponse) SetExecutionTimeoutNil(b bool)`

 SetExecutionTimeoutNil sets the value for ExecutionTimeout to be an explicit nil

### UnsetExecutionTimeout
`func (o *DCVPolicyStatusResponse) UnsetExecutionTimeout()`

UnsetExecutionTimeout ensures that no value is present for ExecutionTimeout, not even an explicit nil
### GetExecutionTimeoutAt

`func (o *DCVPolicyStatusResponse) GetExecutionTimeoutAt() int64`

GetExecutionTimeoutAt returns the ExecutionTimeoutAt field if non-nil, zero value otherwise.

### GetExecutionTimeoutAtOk

`func (o *DCVPolicyStatusResponse) GetExecutionTimeoutAtOk() (*int64, bool)`

GetExecutionTimeoutAtOk returns a tuple with the ExecutionTimeoutAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeoutAt

`func (o *DCVPolicyStatusResponse) SetExecutionTimeoutAt(v int64)`

SetExecutionTimeoutAt sets ExecutionTimeoutAt field to given value.

### HasExecutionTimeoutAt

`func (o *DCVPolicyStatusResponse) HasExecutionTimeoutAt() bool`

HasExecutionTimeoutAt returns a boolean if a field has been set.

### SetExecutionTimeoutAtNil

`func (o *DCVPolicyStatusResponse) SetExecutionTimeoutAtNil(b bool)`

 SetExecutionTimeoutAtNil sets the value for ExecutionTimeoutAt to be an explicit nil

### UnsetExecutionTimeoutAt
`func (o *DCVPolicyStatusResponse) UnsetExecutionTimeoutAt()`

UnsetExecutionTimeoutAt ensures that no value is present for ExecutionTimeoutAt, not even an explicit nil
### GetName

`func (o *DCVPolicyStatusResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DCVPolicyStatusResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DCVPolicyStatusResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNextCheckAt

`func (o *DCVPolicyStatusResponse) GetNextCheckAt() int64`

GetNextCheckAt returns the NextCheckAt field if non-nil, zero value otherwise.

### GetNextCheckAtOk

`func (o *DCVPolicyStatusResponse) GetNextCheckAtOk() (*int64, bool)`

GetNextCheckAtOk returns a tuple with the NextCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCheckAt

`func (o *DCVPolicyStatusResponse) SetNextCheckAt(v int64)`

SetNextCheckAt sets NextCheckAt field to given value.

### HasNextCheckAt

`func (o *DCVPolicyStatusResponse) HasNextCheckAt() bool`

HasNextCheckAt returns a boolean if a field has been set.

### SetNextCheckAtNil

`func (o *DCVPolicyStatusResponse) SetNextCheckAtNil(b bool)`

 SetNextCheckAtNil sets the value for NextCheckAt to be an explicit nil

### UnsetNextCheckAt
`func (o *DCVPolicyStatusResponse) UnsetNextCheckAt()`

UnsetNextCheckAt ensures that no value is present for NextCheckAt, not even an explicit nil
### GetRenewalPeriod

`func (o *DCVPolicyStatusResponse) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *DCVPolicyStatusResponse) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *DCVPolicyStatusResponse) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *DCVPolicyStatusResponse) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *DCVPolicyStatusResponse) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *DCVPolicyStatusResponse) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetRetryDelay

`func (o *DCVPolicyStatusResponse) GetRetryDelay() string`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *DCVPolicyStatusResponse) GetRetryDelayOk() (*string, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *DCVPolicyStatusResponse) SetRetryDelay(v string)`

SetRetryDelay sets RetryDelay field to given value.


### SetRetryDelayNil

`func (o *DCVPolicyStatusResponse) SetRetryDelayNil(b bool)`

 SetRetryDelayNil sets the value for RetryDelay to be an explicit nil

### UnsetRetryDelay
`func (o *DCVPolicyStatusResponse) UnsetRetryDelay()`

UnsetRetryDelay ensures that no value is present for RetryDelay, not even an explicit nil
### GetRunnable

`func (o *DCVPolicyStatusResponse) GetRunnable() bool`

GetRunnable returns the Runnable field if non-nil, zero value otherwise.

### GetRunnableOk

`func (o *DCVPolicyStatusResponse) GetRunnableOk() (*bool, bool)`

GetRunnableOk returns a tuple with the Runnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnable

`func (o *DCVPolicyStatusResponse) SetRunnable(v bool)`

SetRunnable sets Runnable field to given value.


### GetStartedAt

`func (o *DCVPolicyStatusResponse) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DCVPolicyStatusResponse) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DCVPolicyStatusResponse) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DCVPolicyStatusResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *DCVPolicyStatusResponse) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *DCVPolicyStatusResponse) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetStatus

`func (o *DCVPolicyStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DCVPolicyStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DCVPolicyStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


