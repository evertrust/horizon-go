# WebhookRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Define if the webhook url and type is defined here or will be dynamically taken from the certificate&#39;s team | 
**Webhook** | Pointer to [**Webhook**](Webhook.md) | The definition of the webhook. Mandatory in &#x60;static&#x60; mode. | [optional] 

## Methods

### NewWebhookRecipient

`func NewWebhookRecipient(type_ string, ) *WebhookRecipient`

NewWebhookRecipient instantiates a new WebhookRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRecipientWithDefaults

`func NewWebhookRecipientWithDefaults() *WebhookRecipient`

NewWebhookRecipientWithDefaults instantiates a new WebhookRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookRecipient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookRecipient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookRecipient) SetType(v string)`

SetType sets Type field to given value.


### GetWebhook

`func (o *WebhookRecipient) GetWebhook() Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *WebhookRecipient) GetWebhookOk() (*Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *WebhookRecipient) SetWebhook(v Webhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *WebhookRecipient) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


