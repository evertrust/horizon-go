# SecurityTeamUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** |  | [optional] 
**Name** | **string** | The name of the team | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized description of the team | [optional] 
**Contact** | Pointer to **NullableString** | The generic contact e-mail of the Team | [optional] 
**Webhook** | Pointer to [**NullableWebhook**](Webhook.md) | The webhook of the team&#39;s corporate channel (Teams, Slack, Mattermost) | [optional] 
**Manager** | Pointer to **NullableString** | The e-mail address of the team&#39;s manager | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized display name of the team | [optional] 

## Methods

### NewSecurityTeamUpdateRequest

`func NewSecurityTeamUpdateRequest(name string, ) *SecurityTeamUpdateRequest`

NewSecurityTeamUpdateRequest instantiates a new SecurityTeamUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityTeamUpdateRequestWithDefaults

`func NewSecurityTeamUpdateRequestWithDefaults() *SecurityTeamUpdateRequest`

NewSecurityTeamUpdateRequestWithDefaults instantiates a new SecurityTeamUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityTeamUpdateRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityTeamUpdateRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityTeamUpdateRequest) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *SecurityTeamUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SecurityTeamUpdateRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SecurityTeamUpdateRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *SecurityTeamUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityTeamUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityTeamUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SecurityTeamUpdateRequest) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityTeamUpdateRequest) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityTeamUpdateRequest) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityTeamUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityTeamUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityTeamUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContact

`func (o *SecurityTeamUpdateRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SecurityTeamUpdateRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SecurityTeamUpdateRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *SecurityTeamUpdateRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *SecurityTeamUpdateRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *SecurityTeamUpdateRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetWebhook

`func (o *SecurityTeamUpdateRequest) GetWebhook() Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *SecurityTeamUpdateRequest) GetWebhookOk() (*Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *SecurityTeamUpdateRequest) SetWebhook(v Webhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *SecurityTeamUpdateRequest) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *SecurityTeamUpdateRequest) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *SecurityTeamUpdateRequest) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
### GetManager

`func (o *SecurityTeamUpdateRequest) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *SecurityTeamUpdateRequest) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *SecurityTeamUpdateRequest) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *SecurityTeamUpdateRequest) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *SecurityTeamUpdateRequest) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *SecurityTeamUpdateRequest) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetDisplayName

`func (o *SecurityTeamUpdateRequest) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SecurityTeamUpdateRequest) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SecurityTeamUpdateRequest) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SecurityTeamUpdateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SecurityTeamUpdateRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SecurityTeamUpdateRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


