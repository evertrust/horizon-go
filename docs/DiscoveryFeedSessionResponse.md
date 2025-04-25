# DiscoveryFeedSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Object internal ID | [optional] 
**Campaign** | Pointer to **string** | The name of the discovery campaign the feed session belongs to | [optional] 
**Description** | Pointer to **NullableString** | The description of the discovery feed session | [optional] 
**EventOnSuccess** | Pointer to **NullableBool** | Whether to generate an event on success (defaults to the campaign setting) | [optional] 
**EventOnWarning** | Pointer to **NullableBool** | Whether to generate an event on warning (defaults to the campaign setting) | [optional] 
**EventOnFailure** | Pointer to **NullableBool** | Whether to generate an event on failure (defaults to the campaign setting) | [optional] 
**Hosts** | Pointer to **[]string** | The hosts on which the discovery campaign takes place | [optional] 
**Ports** | Pointer to **[]string** | The ports on which the discovery campaign takes place | [optional] 

## Methods

### NewDiscoveryFeedSessionResponse

`func NewDiscoveryFeedSessionResponse() *DiscoveryFeedSessionResponse`

NewDiscoveryFeedSessionResponse instantiates a new DiscoveryFeedSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryFeedSessionResponseWithDefaults

`func NewDiscoveryFeedSessionResponseWithDefaults() *DiscoveryFeedSessionResponse`

NewDiscoveryFeedSessionResponseWithDefaults instantiates a new DiscoveryFeedSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscoveryFeedSessionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveryFeedSessionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveryFeedSessionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscoveryFeedSessionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCampaign

`func (o *DiscoveryFeedSessionResponse) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *DiscoveryFeedSessionResponse) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *DiscoveryFeedSessionResponse) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *DiscoveryFeedSessionResponse) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetDescription

`func (o *DiscoveryFeedSessionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveryFeedSessionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveryFeedSessionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveryFeedSessionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DiscoveryFeedSessionResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DiscoveryFeedSessionResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEventOnSuccess

`func (o *DiscoveryFeedSessionResponse) GetEventOnSuccess() bool`

GetEventOnSuccess returns the EventOnSuccess field if non-nil, zero value otherwise.

### GetEventOnSuccessOk

`func (o *DiscoveryFeedSessionResponse) GetEventOnSuccessOk() (*bool, bool)`

GetEventOnSuccessOk returns a tuple with the EventOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnSuccess

`func (o *DiscoveryFeedSessionResponse) SetEventOnSuccess(v bool)`

SetEventOnSuccess sets EventOnSuccess field to given value.

### HasEventOnSuccess

`func (o *DiscoveryFeedSessionResponse) HasEventOnSuccess() bool`

HasEventOnSuccess returns a boolean if a field has been set.

### SetEventOnSuccessNil

`func (o *DiscoveryFeedSessionResponse) SetEventOnSuccessNil(b bool)`

 SetEventOnSuccessNil sets the value for EventOnSuccess to be an explicit nil

### UnsetEventOnSuccess
`func (o *DiscoveryFeedSessionResponse) UnsetEventOnSuccess()`

UnsetEventOnSuccess ensures that no value is present for EventOnSuccess, not even an explicit nil
### GetEventOnWarning

`func (o *DiscoveryFeedSessionResponse) GetEventOnWarning() bool`

GetEventOnWarning returns the EventOnWarning field if non-nil, zero value otherwise.

### GetEventOnWarningOk

`func (o *DiscoveryFeedSessionResponse) GetEventOnWarningOk() (*bool, bool)`

GetEventOnWarningOk returns a tuple with the EventOnWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnWarning

`func (o *DiscoveryFeedSessionResponse) SetEventOnWarning(v bool)`

SetEventOnWarning sets EventOnWarning field to given value.

### HasEventOnWarning

`func (o *DiscoveryFeedSessionResponse) HasEventOnWarning() bool`

HasEventOnWarning returns a boolean if a field has been set.

### SetEventOnWarningNil

`func (o *DiscoveryFeedSessionResponse) SetEventOnWarningNil(b bool)`

 SetEventOnWarningNil sets the value for EventOnWarning to be an explicit nil

### UnsetEventOnWarning
`func (o *DiscoveryFeedSessionResponse) UnsetEventOnWarning()`

UnsetEventOnWarning ensures that no value is present for EventOnWarning, not even an explicit nil
### GetEventOnFailure

`func (o *DiscoveryFeedSessionResponse) GetEventOnFailure() bool`

GetEventOnFailure returns the EventOnFailure field if non-nil, zero value otherwise.

### GetEventOnFailureOk

`func (o *DiscoveryFeedSessionResponse) GetEventOnFailureOk() (*bool, bool)`

GetEventOnFailureOk returns a tuple with the EventOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnFailure

`func (o *DiscoveryFeedSessionResponse) SetEventOnFailure(v bool)`

SetEventOnFailure sets EventOnFailure field to given value.

### HasEventOnFailure

`func (o *DiscoveryFeedSessionResponse) HasEventOnFailure() bool`

HasEventOnFailure returns a boolean if a field has been set.

### SetEventOnFailureNil

`func (o *DiscoveryFeedSessionResponse) SetEventOnFailureNil(b bool)`

 SetEventOnFailureNil sets the value for EventOnFailure to be an explicit nil

### UnsetEventOnFailure
`func (o *DiscoveryFeedSessionResponse) UnsetEventOnFailure()`

UnsetEventOnFailure ensures that no value is present for EventOnFailure, not even an explicit nil
### GetHosts

`func (o *DiscoveryFeedSessionResponse) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *DiscoveryFeedSessionResponse) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *DiscoveryFeedSessionResponse) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *DiscoveryFeedSessionResponse) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *DiscoveryFeedSessionResponse) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *DiscoveryFeedSessionResponse) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetPorts

`func (o *DiscoveryFeedSessionResponse) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *DiscoveryFeedSessionResponse) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *DiscoveryFeedSessionResponse) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *DiscoveryFeedSessionResponse) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *DiscoveryFeedSessionResponse) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *DiscoveryFeedSessionResponse) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


