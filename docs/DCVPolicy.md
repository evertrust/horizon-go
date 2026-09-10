# DCVPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the DCV policy is enabled; disabled policies are not scheduled | [optional] [default to true]
**ExecutionTimeout** | **NullableString** | Maximum duration for a single DCV run | 
**Filter** | Pointer to **NullableString** | Optional regex filter applied to domain hostnames | [optional] 
**Name** | **string** | Unique name of the DCV policy | 
**Provider** | **string** | Name of the DCV provider configuration to use | 
**Provisioner** | **string** | Name of the DCV provisioner configuration to use | 
**RenewalPolicy** | Pointer to [**NullableDCVRenewalPolicy**](DCVRenewalPolicy.md) | Renewal scheduling policy; if absent no automatic renewal is triggered | [optional] 
**RetryDelay** | **NullableString** | Delay between retry attempts | 
**Triggers** | Pointer to [**NullableDCVPolicyTriggers**](DCVPolicyTriggers.md) | Optional trigger configuration for DCV lifecycle events | [optional] 

## Methods

### NewDCVPolicy

`func NewDCVPolicy(executionTimeout NullableString, name string, provider string, provisioner string, retryDelay NullableString, ) *DCVPolicy`

NewDCVPolicy instantiates a new DCVPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVPolicyWithDefaults

`func NewDCVPolicyWithDefaults() *DCVPolicy`

NewDCVPolicyWithDefaults instantiates a new DCVPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DCVPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DCVPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DCVPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DCVPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExecutionTimeout

`func (o *DCVPolicy) GetExecutionTimeout() string`

GetExecutionTimeout returns the ExecutionTimeout field if non-nil, zero value otherwise.

### GetExecutionTimeoutOk

`func (o *DCVPolicy) GetExecutionTimeoutOk() (*string, bool)`

GetExecutionTimeoutOk returns a tuple with the ExecutionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeout

`func (o *DCVPolicy) SetExecutionTimeout(v string)`

SetExecutionTimeout sets ExecutionTimeout field to given value.


### SetExecutionTimeoutNil

`func (o *DCVPolicy) SetExecutionTimeoutNil(b bool)`

 SetExecutionTimeoutNil sets the value for ExecutionTimeout to be an explicit nil

### UnsetExecutionTimeout
`func (o *DCVPolicy) UnsetExecutionTimeout()`

UnsetExecutionTimeout ensures that no value is present for ExecutionTimeout, not even an explicit nil
### GetFilter

`func (o *DCVPolicy) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DCVPolicy) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DCVPolicy) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DCVPolicy) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *DCVPolicy) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *DCVPolicy) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetName

`func (o *DCVPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DCVPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DCVPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *DCVPolicy) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DCVPolicy) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DCVPolicy) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProvisioner

`func (o *DCVPolicy) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *DCVPolicy) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *DCVPolicy) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.


### GetRenewalPolicy

`func (o *DCVPolicy) GetRenewalPolicy() DCVRenewalPolicy`

GetRenewalPolicy returns the RenewalPolicy field if non-nil, zero value otherwise.

### GetRenewalPolicyOk

`func (o *DCVPolicy) GetRenewalPolicyOk() (*DCVRenewalPolicy, bool)`

GetRenewalPolicyOk returns a tuple with the RenewalPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPolicy

`func (o *DCVPolicy) SetRenewalPolicy(v DCVRenewalPolicy)`

SetRenewalPolicy sets RenewalPolicy field to given value.

### HasRenewalPolicy

`func (o *DCVPolicy) HasRenewalPolicy() bool`

HasRenewalPolicy returns a boolean if a field has been set.

### SetRenewalPolicyNil

`func (o *DCVPolicy) SetRenewalPolicyNil(b bool)`

 SetRenewalPolicyNil sets the value for RenewalPolicy to be an explicit nil

### UnsetRenewalPolicy
`func (o *DCVPolicy) UnsetRenewalPolicy()`

UnsetRenewalPolicy ensures that no value is present for RenewalPolicy, not even an explicit nil
### GetRetryDelay

`func (o *DCVPolicy) GetRetryDelay() string`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *DCVPolicy) GetRetryDelayOk() (*string, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *DCVPolicy) SetRetryDelay(v string)`

SetRetryDelay sets RetryDelay field to given value.


### SetRetryDelayNil

`func (o *DCVPolicy) SetRetryDelayNil(b bool)`

 SetRetryDelayNil sets the value for RetryDelay to be an explicit nil

### UnsetRetryDelay
`func (o *DCVPolicy) UnsetRetryDelay()`

UnsetRetryDelay ensures that no value is present for RetryDelay, not even an explicit nil
### GetTriggers

`func (o *DCVPolicy) GetTriggers() DCVPolicyTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *DCVPolicy) GetTriggersOk() (*DCVPolicyTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *DCVPolicy) SetTriggers(v DCVPolicyTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *DCVPolicy) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *DCVPolicy) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *DCVPolicy) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


