# DCVPolicyLifecycleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the DCV policy is enabled | 
**Filter** | Pointer to **NullableString** | Optional regex filter applied to domain hostnames | [optional] 
**Name** | **string** | Unique name of the DCV policy | 
**Provider** | **string** | Name of the DCV provider configuration used by this policy | 
**Provisioner** | **string** | Name of the DCV provisioner configuration used by this policy | 
**RenewalPolicy** | Pointer to [**NullableDCVRenewalPolicy**](DCVRenewalPolicy.md) | Renewal scheduling policy; absent if no automatic renewal is configured | [optional] 
**Runnable** | **bool** | Whether the policy is enabled and the principal has manage permission | 

## Methods

### NewDCVPolicyLifecycleResponse

`func NewDCVPolicyLifecycleResponse(enabled bool, name string, provider string, provisioner string, runnable bool, ) *DCVPolicyLifecycleResponse`

NewDCVPolicyLifecycleResponse instantiates a new DCVPolicyLifecycleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVPolicyLifecycleResponseWithDefaults

`func NewDCVPolicyLifecycleResponseWithDefaults() *DCVPolicyLifecycleResponse`

NewDCVPolicyLifecycleResponseWithDefaults instantiates a new DCVPolicyLifecycleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DCVPolicyLifecycleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DCVPolicyLifecycleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DCVPolicyLifecycleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFilter

`func (o *DCVPolicyLifecycleResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DCVPolicyLifecycleResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DCVPolicyLifecycleResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DCVPolicyLifecycleResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *DCVPolicyLifecycleResponse) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *DCVPolicyLifecycleResponse) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetName

`func (o *DCVPolicyLifecycleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DCVPolicyLifecycleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DCVPolicyLifecycleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *DCVPolicyLifecycleResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DCVPolicyLifecycleResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DCVPolicyLifecycleResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProvisioner

`func (o *DCVPolicyLifecycleResponse) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *DCVPolicyLifecycleResponse) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *DCVPolicyLifecycleResponse) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.


### GetRenewalPolicy

`func (o *DCVPolicyLifecycleResponse) GetRenewalPolicy() DCVRenewalPolicy`

GetRenewalPolicy returns the RenewalPolicy field if non-nil, zero value otherwise.

### GetRenewalPolicyOk

`func (o *DCVPolicyLifecycleResponse) GetRenewalPolicyOk() (*DCVRenewalPolicy, bool)`

GetRenewalPolicyOk returns a tuple with the RenewalPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPolicy

`func (o *DCVPolicyLifecycleResponse) SetRenewalPolicy(v DCVRenewalPolicy)`

SetRenewalPolicy sets RenewalPolicy field to given value.

### HasRenewalPolicy

`func (o *DCVPolicyLifecycleResponse) HasRenewalPolicy() bool`

HasRenewalPolicy returns a boolean if a field has been set.

### SetRenewalPolicyNil

`func (o *DCVPolicyLifecycleResponse) SetRenewalPolicyNil(b bool)`

 SetRenewalPolicyNil sets the value for RenewalPolicy to be an explicit nil

### UnsetRenewalPolicy
`func (o *DCVPolicyLifecycleResponse) UnsetRenewalPolicy()`

UnsetRenewalPolicy ensures that no value is present for RenewalPolicy, not even an explicit nil
### GetRunnable

`func (o *DCVPolicyLifecycleResponse) GetRunnable() bool`

GetRunnable returns the Runnable field if non-nil, zero value otherwise.

### GetRunnableOk

`func (o *DCVPolicyLifecycleResponse) GetRunnableOk() (*bool, bool)`

GetRunnableOk returns a tuple with the Runnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnable

`func (o *DCVPolicyLifecycleResponse) SetRunnable(v bool)`

SetRunnable sets Runnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


