# Base

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **[]string** | Event on which the notification runs. This MUST contain only one value. | [optional] 
**LicenceUsagePercent** | Pointer to **NullableInt64** | License usage at which the notification needs to run (between 0 and 100). Must be defined on &#x60;on_license_usage&#x60; event and must NOT be defined otherwise. | [optional] 
**Name** | Pointer to **string** | Name of the notification | [optional] 
**Retries** | Pointer to **NullableInt64** | Number of retries when the notification fails | [optional] 
**RunOnRenewed** | Pointer to **NullableBool** | Must be defined on &#x60;on_expire&#x60; event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed. | [optional] 
**RunPeriod** | Pointer to **NullableString** | Time period at which the notification needs to run. Can only be defined on expiration and pending events. | [optional] 
**Type** | Pointer to **string** | The type of notification | [optional] 

## Methods

### NewBase

`func NewBase() *Base`

NewBase instantiates a new Base object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseWithDefaults

`func NewBaseWithDefaults() *Base`

NewBaseWithDefaults instantiates a new Base object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *Base) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Base) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Base) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Base) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLicenceUsagePercent

`func (o *Base) GetLicenceUsagePercent() int64`

GetLicenceUsagePercent returns the LicenceUsagePercent field if non-nil, zero value otherwise.

### GetLicenceUsagePercentOk

`func (o *Base) GetLicenceUsagePercentOk() (*int64, bool)`

GetLicenceUsagePercentOk returns a tuple with the LicenceUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenceUsagePercent

`func (o *Base) SetLicenceUsagePercent(v int64)`

SetLicenceUsagePercent sets LicenceUsagePercent field to given value.

### HasLicenceUsagePercent

`func (o *Base) HasLicenceUsagePercent() bool`

HasLicenceUsagePercent returns a boolean if a field has been set.

### SetLicenceUsagePercentNil

`func (o *Base) SetLicenceUsagePercentNil(b bool)`

 SetLicenceUsagePercentNil sets the value for LicenceUsagePercent to be an explicit nil

### UnsetLicenceUsagePercent
`func (o *Base) UnsetLicenceUsagePercent()`

UnsetLicenceUsagePercent ensures that no value is present for LicenceUsagePercent, not even an explicit nil
### GetName

`func (o *Base) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Base) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Base) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Base) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetries

`func (o *Base) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *Base) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *Base) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *Base) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *Base) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *Base) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetRunOnRenewed

`func (o *Base) GetRunOnRenewed() bool`

GetRunOnRenewed returns the RunOnRenewed field if non-nil, zero value otherwise.

### GetRunOnRenewedOk

`func (o *Base) GetRunOnRenewedOk() (*bool, bool)`

GetRunOnRenewedOk returns a tuple with the RunOnRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnRenewed

`func (o *Base) SetRunOnRenewed(v bool)`

SetRunOnRenewed sets RunOnRenewed field to given value.

### HasRunOnRenewed

`func (o *Base) HasRunOnRenewed() bool`

HasRunOnRenewed returns a boolean if a field has been set.

### SetRunOnRenewedNil

`func (o *Base) SetRunOnRenewedNil(b bool)`

 SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil

### UnsetRunOnRenewed
`func (o *Base) UnsetRunOnRenewed()`

UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
### GetRunPeriod

`func (o *Base) GetRunPeriod() string`

GetRunPeriod returns the RunPeriod field if non-nil, zero value otherwise.

### GetRunPeriodOk

`func (o *Base) GetRunPeriodOk() (*string, bool)`

GetRunPeriodOk returns a tuple with the RunPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPeriod

`func (o *Base) SetRunPeriod(v string)`

SetRunPeriod sets RunPeriod field to given value.

### HasRunPeriod

`func (o *Base) HasRunPeriod() bool`

HasRunPeriod returns a boolean if a field has been set.

### SetRunPeriodNil

`func (o *Base) SetRunPeriodNil(b bool)`

 SetRunPeriodNil sets the value for RunPeriod to be an explicit nil

### UnsetRunPeriod
`func (o *Base) UnsetRunPeriod()`

UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil
### GetType

`func (o *Base) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Base) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Base) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Base) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


