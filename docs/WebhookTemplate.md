# WebhookTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**WebhookRecipient**](WebhookRecipient.md) | The target of the webhook | 
**Title** | **string** | The title of the webhook notification (special formatting) | 
**Body** | Pointer to **NullableString** | The body of the notification. Can contain dynamic attributes. | [optional] 

## Methods

### NewWebhookTemplate

`func NewWebhookTemplate(to WebhookRecipient, title string, ) *WebhookTemplate`

NewWebhookTemplate instantiates a new WebhookTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookTemplateWithDefaults

`func NewWebhookTemplateWithDefaults() *WebhookTemplate`

NewWebhookTemplateWithDefaults instantiates a new WebhookTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *WebhookTemplate) GetTo() WebhookRecipient`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WebhookTemplate) GetToOk() (*WebhookRecipient, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WebhookTemplate) SetTo(v WebhookRecipient)`

SetTo sets To field to given value.


### GetTitle

`func (o *WebhookTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebhookTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebhookTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *WebhookTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WebhookTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WebhookTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *WebhookTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *WebhookTemplate) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *WebhookTemplate) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


