# DiscoveryCampaignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AuthorizationLevels** | [**DiscoveryCampaignAuthorizationLevels**](DiscoveryCampaignAuthorizationLevels.md) | The authorization levels of the discovery campaign | 
**Description** | Pointer to **NullableString** | The description of the discovery campaign | [optional] 
**Enabled** | **bool** | Whether the discovery campaign is enabled, i.e. whether it can be fed | 
**EventOnFailure** | **bool** | Whether to log a Horizon event in case of failure | 
**EventOnSuccess** | **bool** | Whether to log a Horizon event in case of success | 
**EventOnWarning** | **bool** | Whether to log a Horizon event in case of warning | 
**GradingPolicies** | Pointer to **[]string** | The grading policies to apply to grade the discovered certificates on this campaign | [optional] 
**Hosts** | Pointer to **[]string** | The hosts to be scanned by the discovery campaign | [optional] 
**Name** | **string** | The name of the discovery campaign | 
**Ports** | Pointer to **[]string** | The ports to be scanned by the discovery campaign | [optional] [default to ["25","443","8443","689"]]

## Methods

### NewDiscoveryCampaignResponse

`func NewDiscoveryCampaignResponse(id string, authorizationLevels DiscoveryCampaignAuthorizationLevels, enabled bool, eventOnFailure bool, eventOnSuccess bool, eventOnWarning bool, name string, ) *DiscoveryCampaignResponse`

NewDiscoveryCampaignResponse instantiates a new DiscoveryCampaignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryCampaignResponseWithDefaults

`func NewDiscoveryCampaignResponseWithDefaults() *DiscoveryCampaignResponse`

NewDiscoveryCampaignResponseWithDefaults instantiates a new DiscoveryCampaignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscoveryCampaignResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveryCampaignResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveryCampaignResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAuthorizationLevels

`func (o *DiscoveryCampaignResponse) GetAuthorizationLevels() DiscoveryCampaignAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *DiscoveryCampaignResponse) GetAuthorizationLevelsOk() (*DiscoveryCampaignAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *DiscoveryCampaignResponse) SetAuthorizationLevels(v DiscoveryCampaignAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetDescription

`func (o *DiscoveryCampaignResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveryCampaignResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveryCampaignResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveryCampaignResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DiscoveryCampaignResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DiscoveryCampaignResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *DiscoveryCampaignResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DiscoveryCampaignResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DiscoveryCampaignResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEventOnFailure

`func (o *DiscoveryCampaignResponse) GetEventOnFailure() bool`

GetEventOnFailure returns the EventOnFailure field if non-nil, zero value otherwise.

### GetEventOnFailureOk

`func (o *DiscoveryCampaignResponse) GetEventOnFailureOk() (*bool, bool)`

GetEventOnFailureOk returns a tuple with the EventOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnFailure

`func (o *DiscoveryCampaignResponse) SetEventOnFailure(v bool)`

SetEventOnFailure sets EventOnFailure field to given value.


### GetEventOnSuccess

`func (o *DiscoveryCampaignResponse) GetEventOnSuccess() bool`

GetEventOnSuccess returns the EventOnSuccess field if non-nil, zero value otherwise.

### GetEventOnSuccessOk

`func (o *DiscoveryCampaignResponse) GetEventOnSuccessOk() (*bool, bool)`

GetEventOnSuccessOk returns a tuple with the EventOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnSuccess

`func (o *DiscoveryCampaignResponse) SetEventOnSuccess(v bool)`

SetEventOnSuccess sets EventOnSuccess field to given value.


### GetEventOnWarning

`func (o *DiscoveryCampaignResponse) GetEventOnWarning() bool`

GetEventOnWarning returns the EventOnWarning field if non-nil, zero value otherwise.

### GetEventOnWarningOk

`func (o *DiscoveryCampaignResponse) GetEventOnWarningOk() (*bool, bool)`

GetEventOnWarningOk returns a tuple with the EventOnWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnWarning

`func (o *DiscoveryCampaignResponse) SetEventOnWarning(v bool)`

SetEventOnWarning sets EventOnWarning field to given value.


### GetGradingPolicies

`func (o *DiscoveryCampaignResponse) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *DiscoveryCampaignResponse) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *DiscoveryCampaignResponse) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *DiscoveryCampaignResponse) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *DiscoveryCampaignResponse) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *DiscoveryCampaignResponse) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetHosts

`func (o *DiscoveryCampaignResponse) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *DiscoveryCampaignResponse) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *DiscoveryCampaignResponse) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *DiscoveryCampaignResponse) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *DiscoveryCampaignResponse) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *DiscoveryCampaignResponse) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetName

`func (o *DiscoveryCampaignResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveryCampaignResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveryCampaignResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPorts

`func (o *DiscoveryCampaignResponse) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *DiscoveryCampaignResponse) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *DiscoveryCampaignResponse) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *DiscoveryCampaignResponse) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *DiscoveryCampaignResponse) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *DiscoveryCampaignResponse) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


