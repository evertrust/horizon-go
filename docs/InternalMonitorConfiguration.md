# InternalMonitorConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | **string** | Cron defining when to run internal monitor checks | 
**Type** | **string** | The type of the configuration entry | 

## Methods

### NewInternalMonitorConfiguration

`func NewInternalMonitorConfiguration(cron string, type_ string, ) *InternalMonitorConfiguration`

NewInternalMonitorConfiguration instantiates a new InternalMonitorConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalMonitorConfigurationWithDefaults

`func NewInternalMonitorConfigurationWithDefaults() *InternalMonitorConfiguration`

NewInternalMonitorConfigurationWithDefaults instantiates a new InternalMonitorConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *InternalMonitorConfiguration) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *InternalMonitorConfiguration) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *InternalMonitorConfiguration) SetCron(v string)`

SetCron sets Cron field to given value.


### GetType

`func (o *InternalMonitorConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalMonitorConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalMonitorConfiguration) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


