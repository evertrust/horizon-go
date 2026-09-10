# WebhookNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proxy** | Pointer to **NullableString** | Name of a Proxy to use while sending the webhook | [optional] 
**Retries** | Pointer to **interface{}** | Number of retries when the notification fails (non 200 return code) | [optional] 
**Timeout** | Pointer to **string** | Timeout for the webhook request | [optional] 
**Type** | **string** |  | 
**WebhookTemplate** | [**WebhookTemplate**](WebhookTemplate.md) |  | 
**Events** | **[]string** | Event on which the notification runs. This MUST contain only one value. | 
**LicenceUsagePercent** | Pointer to **NullableInt64** | License usage at which the notification needs to run (between 0 and 100). Must be defined on &#x60;on_license_usage&#x60; event and must NOT be defined otherwise. | [optional] 
**Name** | **string** | Name of the notification | 
**RunOnRenewed** | Pointer to **NullableBool** | Must be defined on &#x60;on_expire&#x60; event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed. | [optional] 
**RunPeriod** | Pointer to **NullableString** | Time period at which the notification needs to run. Can only be defined on expiration and pending events. | [optional] 

## Methods

### NewWebhookNotification

`func NewWebhookNotification(type_ string, webhookTemplate WebhookTemplate, events []string, name string, ) *WebhookNotification`

NewWebhookNotification instantiates a new WebhookNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookNotificationWithDefaults

`func NewWebhookNotificationWithDefaults() *WebhookNotification`

NewWebhookNotificationWithDefaults instantiates a new WebhookNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProxy

`func (o *WebhookNotification) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *WebhookNotification) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *WebhookNotification) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *WebhookNotification) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *WebhookNotification) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *WebhookNotification) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRetries

`func (o *WebhookNotification) GetRetries() interface{}`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *WebhookNotification) GetRetriesOk() (*interface{}, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *WebhookNotification) SetRetries(v interface{})`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *WebhookNotification) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *WebhookNotification) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *WebhookNotification) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetTimeout

`func (o *WebhookNotification) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WebhookNotification) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WebhookNotification) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *WebhookNotification) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *WebhookNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookNotification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookNotification) SetType(v string)`

SetType sets Type field to given value.


### GetWebhookTemplate

`func (o *WebhookNotification) GetWebhookTemplate() WebhookTemplate`

GetWebhookTemplate returns the WebhookTemplate field if non-nil, zero value otherwise.

### GetWebhookTemplateOk

`func (o *WebhookNotification) GetWebhookTemplateOk() (*WebhookTemplate, bool)`

GetWebhookTemplateOk returns a tuple with the WebhookTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTemplate

`func (o *WebhookNotification) SetWebhookTemplate(v WebhookTemplate)`

SetWebhookTemplate sets WebhookTemplate field to given value.


### GetEvents

`func (o *WebhookNotification) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookNotification) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookNotification) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetLicenceUsagePercent

`func (o *WebhookNotification) GetLicenceUsagePercent() int64`

GetLicenceUsagePercent returns the LicenceUsagePercent field if non-nil, zero value otherwise.

### GetLicenceUsagePercentOk

`func (o *WebhookNotification) GetLicenceUsagePercentOk() (*int64, bool)`

GetLicenceUsagePercentOk returns a tuple with the LicenceUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenceUsagePercent

`func (o *WebhookNotification) SetLicenceUsagePercent(v int64)`

SetLicenceUsagePercent sets LicenceUsagePercent field to given value.

### HasLicenceUsagePercent

`func (o *WebhookNotification) HasLicenceUsagePercent() bool`

HasLicenceUsagePercent returns a boolean if a field has been set.

### SetLicenceUsagePercentNil

`func (o *WebhookNotification) SetLicenceUsagePercentNil(b bool)`

 SetLicenceUsagePercentNil sets the value for LicenceUsagePercent to be an explicit nil

### UnsetLicenceUsagePercent
`func (o *WebhookNotification) UnsetLicenceUsagePercent()`

UnsetLicenceUsagePercent ensures that no value is present for LicenceUsagePercent, not even an explicit nil
### GetName

`func (o *WebhookNotification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookNotification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookNotification) SetName(v string)`

SetName sets Name field to given value.


### GetRunOnRenewed

`func (o *WebhookNotification) GetRunOnRenewed() bool`

GetRunOnRenewed returns the RunOnRenewed field if non-nil, zero value otherwise.

### GetRunOnRenewedOk

`func (o *WebhookNotification) GetRunOnRenewedOk() (*bool, bool)`

GetRunOnRenewedOk returns a tuple with the RunOnRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnRenewed

`func (o *WebhookNotification) SetRunOnRenewed(v bool)`

SetRunOnRenewed sets RunOnRenewed field to given value.

### HasRunOnRenewed

`func (o *WebhookNotification) HasRunOnRenewed() bool`

HasRunOnRenewed returns a boolean if a field has been set.

### SetRunOnRenewedNil

`func (o *WebhookNotification) SetRunOnRenewedNil(b bool)`

 SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil

### UnsetRunOnRenewed
`func (o *WebhookNotification) UnsetRunOnRenewed()`

UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
### GetRunPeriod

`func (o *WebhookNotification) GetRunPeriod() string`

GetRunPeriod returns the RunPeriod field if non-nil, zero value otherwise.

### GetRunPeriodOk

`func (o *WebhookNotification) GetRunPeriodOk() (*string, bool)`

GetRunPeriodOk returns a tuple with the RunPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPeriod

`func (o *WebhookNotification) SetRunPeriod(v string)`

SetRunPeriod sets RunPeriod field to given value.

### HasRunPeriod

`func (o *WebhookNotification) HasRunPeriod() bool`

HasRunPeriod returns a boolean if a field has been set.

### SetRunPeriodNil

`func (o *WebhookNotification) SetRunPeriodNil(b bool)`

 SetRunPeriodNil sets the value for RunPeriod to be an explicit nil

### UnsetRunPeriod
`func (o *WebhookNotification) UnsetRunPeriod()`

UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


