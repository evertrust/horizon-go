# REST

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Retries** | Pointer to **interface{}** | Number of retries when the notification fails (depends on &#x60;expectedHttpCodes&#x60;) | [optional] 
**Sequence** | **[]interface{}** | The REST requests to execute, in execution order. Each request enriches the dictionary with its response for the next one | 
**Name** | **string** | Name of the notification | 
**RunPeriod** | Pointer to **NullableString** | Time period at which the notification needs to run. Can only be defined on expiration and pending events. | [optional] 
**LicenseUsagePercent** | Pointer to **NullableInt64** | License usage at which the notification needs to run (between 0 and 100). Must be defined on &#x60;on_license_usage&#x60; event and must NOT be defined otherwise. | [optional] 
**Events** | **[]string** | Event on which the notification runs. This MUST contain only one value. | 
**RunOnRenewed** | Pointer to **NullableBool** | Must be defined on &#x60;on_expire&#x60; event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed. | [optional] 

## Methods

### NewREST

`func NewREST(type_ string, sequence []interface{}, name string, events []string, ) *REST`

NewREST instantiates a new REST object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRESTWithDefaults

`func NewRESTWithDefaults() *REST`

NewRESTWithDefaults instantiates a new REST object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *REST) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *REST) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *REST) SetType(v string)`

SetType sets Type field to given value.


### GetRetries

`func (o *REST) GetRetries() interface{}`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *REST) GetRetriesOk() (*interface{}, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *REST) SetRetries(v interface{})`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *REST) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *REST) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *REST) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetSequence

`func (o *REST) GetSequence() []interface{}`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *REST) GetSequenceOk() (*[]interface{}, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *REST) SetSequence(v []interface{})`

SetSequence sets Sequence field to given value.


### GetName

`func (o *REST) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *REST) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *REST) SetName(v string)`

SetName sets Name field to given value.


### GetRunPeriod

`func (o *REST) GetRunPeriod() string`

GetRunPeriod returns the RunPeriod field if non-nil, zero value otherwise.

### GetRunPeriodOk

`func (o *REST) GetRunPeriodOk() (*string, bool)`

GetRunPeriodOk returns a tuple with the RunPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPeriod

`func (o *REST) SetRunPeriod(v string)`

SetRunPeriod sets RunPeriod field to given value.

### HasRunPeriod

`func (o *REST) HasRunPeriod() bool`

HasRunPeriod returns a boolean if a field has been set.

### SetRunPeriodNil

`func (o *REST) SetRunPeriodNil(b bool)`

 SetRunPeriodNil sets the value for RunPeriod to be an explicit nil

### UnsetRunPeriod
`func (o *REST) UnsetRunPeriod()`

UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil
### GetLicenseUsagePercent

`func (o *REST) GetLicenseUsagePercent() int64`

GetLicenseUsagePercent returns the LicenseUsagePercent field if non-nil, zero value otherwise.

### GetLicenseUsagePercentOk

`func (o *REST) GetLicenseUsagePercentOk() (*int64, bool)`

GetLicenseUsagePercentOk returns a tuple with the LicenseUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUsagePercent

`func (o *REST) SetLicenseUsagePercent(v int64)`

SetLicenseUsagePercent sets LicenseUsagePercent field to given value.

### HasLicenseUsagePercent

`func (o *REST) HasLicenseUsagePercent() bool`

HasLicenseUsagePercent returns a boolean if a field has been set.

### SetLicenseUsagePercentNil

`func (o *REST) SetLicenseUsagePercentNil(b bool)`

 SetLicenseUsagePercentNil sets the value for LicenseUsagePercent to be an explicit nil

### UnsetLicenseUsagePercent
`func (o *REST) UnsetLicenseUsagePercent()`

UnsetLicenseUsagePercent ensures that no value is present for LicenseUsagePercent, not even an explicit nil
### GetEvents

`func (o *REST) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *REST) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *REST) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetRunOnRenewed

`func (o *REST) GetRunOnRenewed() bool`

GetRunOnRenewed returns the RunOnRenewed field if non-nil, zero value otherwise.

### GetRunOnRenewedOk

`func (o *REST) GetRunOnRenewedOk() (*bool, bool)`

GetRunOnRenewedOk returns a tuple with the RunOnRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnRenewed

`func (o *REST) SetRunOnRenewed(v bool)`

SetRunOnRenewed sets RunOnRenewed field to given value.

### HasRunOnRenewed

`func (o *REST) HasRunOnRenewed() bool`

HasRunOnRenewed returns a boolean if a field has been set.

### SetRunOnRenewedNil

`func (o *REST) SetRunOnRenewedNil(b bool)`

 SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil

### UnsetRunOnRenewed
`func (o *REST) UnsetRunOnRenewed()`

UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


