# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **NullableString** | The generic contact e-mail of the Team | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized description of the team | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized display name of the team | [optional] 
**Managers** | Pointer to **[]string** | The identifiers of the team&#39;s managers | [optional] 
**Name** | **string** | The name of the team | 
**Webhook** | Pointer to [**NullableWebhook**](Webhook.md) | The webhook of the team&#39;s corporate channel (Teams, Slack, Mattermost) | [optional] 

## Methods

### NewTeam

`func NewTeam(name string, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *Team) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Team) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Team) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Team) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *Team) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Team) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetDescription

`func (o *Team) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Team) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Team) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Team) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Team) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Team) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *Team) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Team) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Team) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Team) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Team) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Team) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetManagers

`func (o *Team) GetManagers() []string`

GetManagers returns the Managers field if non-nil, zero value otherwise.

### GetManagersOk

`func (o *Team) GetManagersOk() (*[]string, bool)`

GetManagersOk returns a tuple with the Managers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagers

`func (o *Team) SetManagers(v []string)`

SetManagers sets Managers field to given value.

### HasManagers

`func (o *Team) HasManagers() bool`

HasManagers returns a boolean if a field has been set.

### SetManagersNil

`func (o *Team) SetManagersNil(b bool)`

 SetManagersNil sets the value for Managers to be an explicit nil

### UnsetManagers
`func (o *Team) UnsetManagers()`

UnsetManagers ensures that no value is present for Managers, not even an explicit nil
### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.


### GetWebhook

`func (o *Team) GetWebhook() Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *Team) GetWebhookOk() (*Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *Team) SetWebhook(v Webhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *Team) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *Team) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *Team) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


