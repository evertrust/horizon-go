# DCVRenewalPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | **string** | A Quartz cron expression defining when domain validation should run | 
**RenewalPeriod** | **NullableString** | Duration before DCV expiry at which renewal should be triggered | 

## Methods

### NewDCVRenewalPolicy

`func NewDCVRenewalPolicy(cron string, renewalPeriod NullableString, ) *DCVRenewalPolicy`

NewDCVRenewalPolicy instantiates a new DCVRenewalPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVRenewalPolicyWithDefaults

`func NewDCVRenewalPolicyWithDefaults() *DCVRenewalPolicy`

NewDCVRenewalPolicyWithDefaults instantiates a new DCVRenewalPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *DCVRenewalPolicy) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *DCVRenewalPolicy) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *DCVRenewalPolicy) SetCron(v string)`

SetCron sets Cron field to given value.


### GetRenewalPeriod

`func (o *DCVRenewalPolicy) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *DCVRenewalPolicy) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *DCVRenewalPolicy) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.


### SetRenewalPeriodNil

`func (o *DCVRenewalPolicy) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *DCVRenewalPolicy) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


