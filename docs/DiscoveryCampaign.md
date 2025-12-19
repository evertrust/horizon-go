# DiscoveryCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewDiscoveryCampaign

`func NewDiscoveryCampaign(authorizationLevels DiscoveryCampaignAuthorizationLevels, enabled bool, eventOnFailure bool, eventOnSuccess bool, eventOnWarning bool, name string, ) *DiscoveryCampaign`

NewDiscoveryCampaign instantiates a new DiscoveryCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryCampaignWithDefaults

`func NewDiscoveryCampaignWithDefaults() *DiscoveryCampaign`

NewDiscoveryCampaignWithDefaults instantiates a new DiscoveryCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationLevels

`func (o *DiscoveryCampaign) GetAuthorizationLevels() DiscoveryCampaignAuthorizationLevels`

GetAuthorizationLevels returns the AuthorizationLevels field if non-nil, zero value otherwise.

### GetAuthorizationLevelsOk

`func (o *DiscoveryCampaign) GetAuthorizationLevelsOk() (*DiscoveryCampaignAuthorizationLevels, bool)`

GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLevels

`func (o *DiscoveryCampaign) SetAuthorizationLevels(v DiscoveryCampaignAuthorizationLevels)`

SetAuthorizationLevels sets AuthorizationLevels field to given value.


### GetDescription

`func (o *DiscoveryCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveryCampaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveryCampaign) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveryCampaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DiscoveryCampaign) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DiscoveryCampaign) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *DiscoveryCampaign) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DiscoveryCampaign) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DiscoveryCampaign) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEventOnFailure

`func (o *DiscoveryCampaign) GetEventOnFailure() bool`

GetEventOnFailure returns the EventOnFailure field if non-nil, zero value otherwise.

### GetEventOnFailureOk

`func (o *DiscoveryCampaign) GetEventOnFailureOk() (*bool, bool)`

GetEventOnFailureOk returns a tuple with the EventOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnFailure

`func (o *DiscoveryCampaign) SetEventOnFailure(v bool)`

SetEventOnFailure sets EventOnFailure field to given value.


### GetEventOnSuccess

`func (o *DiscoveryCampaign) GetEventOnSuccess() bool`

GetEventOnSuccess returns the EventOnSuccess field if non-nil, zero value otherwise.

### GetEventOnSuccessOk

`func (o *DiscoveryCampaign) GetEventOnSuccessOk() (*bool, bool)`

GetEventOnSuccessOk returns a tuple with the EventOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnSuccess

`func (o *DiscoveryCampaign) SetEventOnSuccess(v bool)`

SetEventOnSuccess sets EventOnSuccess field to given value.


### GetEventOnWarning

`func (o *DiscoveryCampaign) GetEventOnWarning() bool`

GetEventOnWarning returns the EventOnWarning field if non-nil, zero value otherwise.

### GetEventOnWarningOk

`func (o *DiscoveryCampaign) GetEventOnWarningOk() (*bool, bool)`

GetEventOnWarningOk returns a tuple with the EventOnWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnWarning

`func (o *DiscoveryCampaign) SetEventOnWarning(v bool)`

SetEventOnWarning sets EventOnWarning field to given value.


### GetGradingPolicies

`func (o *DiscoveryCampaign) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *DiscoveryCampaign) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *DiscoveryCampaign) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *DiscoveryCampaign) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *DiscoveryCampaign) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *DiscoveryCampaign) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetHosts

`func (o *DiscoveryCampaign) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *DiscoveryCampaign) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *DiscoveryCampaign) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *DiscoveryCampaign) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *DiscoveryCampaign) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *DiscoveryCampaign) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetName

`func (o *DiscoveryCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveryCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveryCampaign) SetName(v string)`

SetName sets Name field to given value.


### GetPorts

`func (o *DiscoveryCampaign) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *DiscoveryCampaign) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *DiscoveryCampaign) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *DiscoveryCampaign) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *DiscoveryCampaign) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *DiscoveryCampaign) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


