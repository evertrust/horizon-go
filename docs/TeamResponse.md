# TeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** | The name of the team | 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized description of the team | [optional] 
**Contact** | Pointer to **NullableString** | The generic contact e-mail of the Team | [optional] 
**Webhook** | Pointer to [**NullableWebhook**](Webhook.md) | The webhook of the team&#39;s corporate channel (Teams, Slack, Mattermost) | [optional] 
**Managers** | **[]string** | The identifiers of the team&#39;s managers | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized display name of the team | [optional] 

## Methods

### NewTeamResponse

`func NewTeamResponse(id string, name string, managers []string, ) *TeamResponse`

NewTeamResponse instantiates a new TeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamResponseWithDefaults

`func NewTeamResponseWithDefaults() *TeamResponse`

NewTeamResponseWithDefaults instantiates a new TeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TeamResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TeamResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TeamResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContact

`func (o *TeamResponse) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *TeamResponse) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *TeamResponse) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *TeamResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *TeamResponse) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *TeamResponse) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetWebhook

`func (o *TeamResponse) GetWebhook() Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *TeamResponse) GetWebhookOk() (*Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *TeamResponse) SetWebhook(v Webhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *TeamResponse) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *TeamResponse) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *TeamResponse) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
### GetManagers

`func (o *TeamResponse) GetManagers() []string`

GetManagers returns the Managers field if non-nil, zero value otherwise.

### GetManagersOk

`func (o *TeamResponse) GetManagersOk() (*[]string, bool)`

GetManagersOk returns a tuple with the Managers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagers

`func (o *TeamResponse) SetManagers(v []string)`

SetManagers sets Managers field to given value.


### GetDisplayName

`func (o *TeamResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TeamResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TeamResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TeamResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TeamResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TeamResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


