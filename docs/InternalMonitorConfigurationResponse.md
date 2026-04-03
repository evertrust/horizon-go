# InternalMonitorConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Cron** | **string** | Cron defining when to run internal monitor checks | 
**Type** | **string** | The type of the configuration entry | 

## Methods

### NewInternalMonitorConfigurationResponse

`func NewInternalMonitorConfigurationResponse(id string, cron string, type_ string, ) *InternalMonitorConfigurationResponse`

NewInternalMonitorConfigurationResponse instantiates a new InternalMonitorConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalMonitorConfigurationResponseWithDefaults

`func NewInternalMonitorConfigurationResponseWithDefaults() *InternalMonitorConfigurationResponse`

NewInternalMonitorConfigurationResponseWithDefaults instantiates a new InternalMonitorConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalMonitorConfigurationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalMonitorConfigurationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalMonitorConfigurationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCron

`func (o *InternalMonitorConfigurationResponse) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *InternalMonitorConfigurationResponse) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *InternalMonitorConfigurationResponse) SetCron(v string)`

SetCron sets Cron field to given value.


### GetType

`func (o *InternalMonitorConfigurationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalMonitorConfigurationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalMonitorConfigurationResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


