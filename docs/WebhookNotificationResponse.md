# WebhookNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Proxy** | Pointer to **NullableString** | Name of a Proxy to use while sending the webhook | [optional] 
**Retries** | Pointer to **interface{}** | Number of retries when the notification fails (non 200 return code) | [optional] 
**Timeout** | Pointer to **string** | Timeout for the webhook request | [optional] 
**Type** | **string** |  | 
**WebhookTemplate** | [**WebhookTemplate**](WebhookTemplate.md) |  | 
**Events** | **[]string** | Event on which the notification runs. This MUST contain only one value. | 
**LicenseUsagePercent** | Pointer to **NullableInt64** | License usage at which the notification needs to run (between 0 and 100). Must be defined on &#x60;on_license_usage&#x60; event and must NOT be defined otherwise. | [optional] 
**Name** | **string** | Name of the notification | 
**RunOnRenewed** | Pointer to **NullableBool** | Must be defined on &#x60;on_expire&#x60; event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed. | [optional] 
**RunPeriod** | Pointer to **NullableString** | Time period at which the notification needs to run. Can only be defined on expiration and pending events. | [optional] 

## Methods

### NewWebhookNotificationResponse

`func NewWebhookNotificationResponse(id string, type_ string, webhookTemplate WebhookTemplate, events []string, name string, ) *WebhookNotificationResponse`

NewWebhookNotificationResponse instantiates a new WebhookNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookNotificationResponseWithDefaults

`func NewWebhookNotificationResponseWithDefaults() *WebhookNotificationResponse`

NewWebhookNotificationResponseWithDefaults instantiates a new WebhookNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookNotificationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookNotificationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookNotificationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetProxy

`func (o *WebhookNotificationResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *WebhookNotificationResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *WebhookNotificationResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *WebhookNotificationResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *WebhookNotificationResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *WebhookNotificationResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRetries

`func (o *WebhookNotificationResponse) GetRetries() interface{}`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *WebhookNotificationResponse) GetRetriesOk() (*interface{}, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *WebhookNotificationResponse) SetRetries(v interface{})`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *WebhookNotificationResponse) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *WebhookNotificationResponse) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *WebhookNotificationResponse) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetTimeout

`func (o *WebhookNotificationResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WebhookNotificationResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WebhookNotificationResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *WebhookNotificationResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *WebhookNotificationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookNotificationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookNotificationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetWebhookTemplate

`func (o *WebhookNotificationResponse) GetWebhookTemplate() WebhookTemplate`

GetWebhookTemplate returns the WebhookTemplate field if non-nil, zero value otherwise.

### GetWebhookTemplateOk

`func (o *WebhookNotificationResponse) GetWebhookTemplateOk() (*WebhookTemplate, bool)`

GetWebhookTemplateOk returns a tuple with the WebhookTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTemplate

`func (o *WebhookNotificationResponse) SetWebhookTemplate(v WebhookTemplate)`

SetWebhookTemplate sets WebhookTemplate field to given value.


### GetEvents

`func (o *WebhookNotificationResponse) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookNotificationResponse) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookNotificationResponse) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetLicenseUsagePercent

`func (o *WebhookNotificationResponse) GetLicenseUsagePercent() int64`

GetLicenseUsagePercent returns the LicenseUsagePercent field if non-nil, zero value otherwise.

### GetLicenseUsagePercentOk

`func (o *WebhookNotificationResponse) GetLicenseUsagePercentOk() (*int64, bool)`

GetLicenseUsagePercentOk returns a tuple with the LicenseUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUsagePercent

`func (o *WebhookNotificationResponse) SetLicenseUsagePercent(v int64)`

SetLicenseUsagePercent sets LicenseUsagePercent field to given value.

### HasLicenseUsagePercent

`func (o *WebhookNotificationResponse) HasLicenseUsagePercent() bool`

HasLicenseUsagePercent returns a boolean if a field has been set.

### SetLicenseUsagePercentNil

`func (o *WebhookNotificationResponse) SetLicenseUsagePercentNil(b bool)`

 SetLicenseUsagePercentNil sets the value for LicenseUsagePercent to be an explicit nil

### UnsetLicenseUsagePercent
`func (o *WebhookNotificationResponse) UnsetLicenseUsagePercent()`

UnsetLicenseUsagePercent ensures that no value is present for LicenseUsagePercent, not even an explicit nil
### GetName

`func (o *WebhookNotificationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookNotificationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookNotificationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRunOnRenewed

`func (o *WebhookNotificationResponse) GetRunOnRenewed() bool`

GetRunOnRenewed returns the RunOnRenewed field if non-nil, zero value otherwise.

### GetRunOnRenewedOk

`func (o *WebhookNotificationResponse) GetRunOnRenewedOk() (*bool, bool)`

GetRunOnRenewedOk returns a tuple with the RunOnRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnRenewed

`func (o *WebhookNotificationResponse) SetRunOnRenewed(v bool)`

SetRunOnRenewed sets RunOnRenewed field to given value.

### HasRunOnRenewed

`func (o *WebhookNotificationResponse) HasRunOnRenewed() bool`

HasRunOnRenewed returns a boolean if a field has been set.

### SetRunOnRenewedNil

`func (o *WebhookNotificationResponse) SetRunOnRenewedNil(b bool)`

 SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil

### UnsetRunOnRenewed
`func (o *WebhookNotificationResponse) UnsetRunOnRenewed()`

UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
### GetRunPeriod

`func (o *WebhookNotificationResponse) GetRunPeriod() string`

GetRunPeriod returns the RunPeriod field if non-nil, zero value otherwise.

### GetRunPeriodOk

`func (o *WebhookNotificationResponse) GetRunPeriodOk() (*string, bool)`

GetRunPeriodOk returns a tuple with the RunPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPeriod

`func (o *WebhookNotificationResponse) SetRunPeriod(v string)`

SetRunPeriod sets RunPeriod field to given value.

### HasRunPeriod

`func (o *WebhookNotificationResponse) HasRunPeriod() bool`

HasRunPeriod returns a boolean if a field has been set.

### SetRunPeriodNil

`func (o *WebhookNotificationResponse) SetRunPeriodNil(b bool)`

 SetRunPeriodNil sets the value for RunPeriod to be an explicit nil

### UnsetRunPeriod
`func (o *WebhookNotificationResponse) UnsetRunPeriod()`

UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


