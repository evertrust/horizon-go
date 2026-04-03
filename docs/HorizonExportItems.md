# HorizonExportItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automations** | Pointer to [**[]AutomationPolicyResponse**](AutomationPolicyResponse.md) |  | [optional] 
**Cas** | Pointer to [**[]CertificateAuthorityExportResponse**](CertificateAuthorityExportResponse.md) |  | [optional] 
**Datasources** | Pointer to [**[]HorizonExportItemsDatasourcesInner**](HorizonExportItemsDatasourcesInner.md) |  | [optional] 
**DiscoveryCampaigns** | Pointer to [**[]DiscoveryCampaignResponse**](DiscoveryCampaignResponse.md) |  | [optional] 
**Executions** | Pointer to [**[]ExecutionPolicyResponse**](ExecutionPolicyResponse.md) |  | [optional] 
**ForestMappings** | Pointer to [**[]WcceForestMappingResponse**](WcceForestMappingResponse.md) |  | [optional] 
**Labels** | Pointer to [**[]LabelResponse**](LabelResponse.md) |  | [optional] 
**Notifications** | Pointer to [**[]HorizonExportItemsNotificationsInner**](HorizonExportItemsNotificationsInner.md) |  | [optional] 
**PasswordPolicies** | Pointer to [**[]PasswordPolicyResponse**](PasswordPolicyResponse.md) |  | [optional] 
**PkiConnectors** | Pointer to [**[]PKIResponses**](PKIResponses.md) |  | [optional] 
**PkiQueues** | Pointer to [**[]PKIQueueResponse**](PKIQueueResponse.md) |  | [optional] 
**Profiles** | Pointer to [**[]CertificateProfileResponses**](CertificateProfileResponses.md) |  | [optional] 
**Proxies** | Pointer to [**[]HttpProxyResponse**](HttpProxyResponse.md) |  | [optional] 
**Reports** | Pointer to [**[]ReportScheduledTaskResponse**](ReportScheduledTaskResponse.md) |  | [optional] 
**Roles** | Pointer to [**[]RoleResponse**](RoleResponse.md) |  | [optional] 
**ScimProfiles** | Pointer to [**[]ScimProfileResponse**](ScimProfileResponse.md) |  | [optional] 
**Storages** | Pointer to [**[]S3StorageBackendConfigResponse**](S3StorageBackendConfigResponse.md) |  | [optional] 
**Teams** | Pointer to [**[]TeamResponse**](TeamResponse.md) |  | [optional] 
**ThirdParties** | Pointer to [**[]ThirdPartyConnectorResponses**](ThirdPartyConnectorResponses.md) |  | [optional] 
**Triggers** | Pointer to [**[]HorizonExportItemsTriggersInner**](HorizonExportItemsTriggersInner.md) |  | [optional] 

## Methods

### NewHorizonExportItems

`func NewHorizonExportItems() *HorizonExportItems`

NewHorizonExportItems instantiates a new HorizonExportItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHorizonExportItemsWithDefaults

`func NewHorizonExportItemsWithDefaults() *HorizonExportItems`

NewHorizonExportItemsWithDefaults instantiates a new HorizonExportItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomations

`func (o *HorizonExportItems) GetAutomations() []AutomationPolicyResponse`

GetAutomations returns the Automations field if non-nil, zero value otherwise.

### GetAutomationsOk

`func (o *HorizonExportItems) GetAutomationsOk() (*[]AutomationPolicyResponse, bool)`

GetAutomationsOk returns a tuple with the Automations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomations

`func (o *HorizonExportItems) SetAutomations(v []AutomationPolicyResponse)`

SetAutomations sets Automations field to given value.

### HasAutomations

`func (o *HorizonExportItems) HasAutomations() bool`

HasAutomations returns a boolean if a field has been set.

### GetCas

`func (o *HorizonExportItems) GetCas() []CertificateAuthorityExportResponse`

GetCas returns the Cas field if non-nil, zero value otherwise.

### GetCasOk

`func (o *HorizonExportItems) GetCasOk() (*[]CertificateAuthorityExportResponse, bool)`

GetCasOk returns a tuple with the Cas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCas

`func (o *HorizonExportItems) SetCas(v []CertificateAuthorityExportResponse)`

SetCas sets Cas field to given value.

### HasCas

`func (o *HorizonExportItems) HasCas() bool`

HasCas returns a boolean if a field has been set.

### GetDatasources

`func (o *HorizonExportItems) GetDatasources() []HorizonExportItemsDatasourcesInner`

GetDatasources returns the Datasources field if non-nil, zero value otherwise.

### GetDatasourcesOk

`func (o *HorizonExportItems) GetDatasourcesOk() (*[]HorizonExportItemsDatasourcesInner, bool)`

GetDatasourcesOk returns a tuple with the Datasources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasources

`func (o *HorizonExportItems) SetDatasources(v []HorizonExportItemsDatasourcesInner)`

SetDatasources sets Datasources field to given value.

### HasDatasources

`func (o *HorizonExportItems) HasDatasources() bool`

HasDatasources returns a boolean if a field has been set.

### GetDiscoveryCampaigns

`func (o *HorizonExportItems) GetDiscoveryCampaigns() []DiscoveryCampaignResponse`

GetDiscoveryCampaigns returns the DiscoveryCampaigns field if non-nil, zero value otherwise.

### GetDiscoveryCampaignsOk

`func (o *HorizonExportItems) GetDiscoveryCampaignsOk() (*[]DiscoveryCampaignResponse, bool)`

GetDiscoveryCampaignsOk returns a tuple with the DiscoveryCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryCampaigns

`func (o *HorizonExportItems) SetDiscoveryCampaigns(v []DiscoveryCampaignResponse)`

SetDiscoveryCampaigns sets DiscoveryCampaigns field to given value.

### HasDiscoveryCampaigns

`func (o *HorizonExportItems) HasDiscoveryCampaigns() bool`

HasDiscoveryCampaigns returns a boolean if a field has been set.

### GetExecutions

`func (o *HorizonExportItems) GetExecutions() []ExecutionPolicyResponse`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *HorizonExportItems) GetExecutionsOk() (*[]ExecutionPolicyResponse, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *HorizonExportItems) SetExecutions(v []ExecutionPolicyResponse)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *HorizonExportItems) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetForestMappings

`func (o *HorizonExportItems) GetForestMappings() []WcceForestMappingResponse`

GetForestMappings returns the ForestMappings field if non-nil, zero value otherwise.

### GetForestMappingsOk

`func (o *HorizonExportItems) GetForestMappingsOk() (*[]WcceForestMappingResponse, bool)`

GetForestMappingsOk returns a tuple with the ForestMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestMappings

`func (o *HorizonExportItems) SetForestMappings(v []WcceForestMappingResponse)`

SetForestMappings sets ForestMappings field to given value.

### HasForestMappings

`func (o *HorizonExportItems) HasForestMappings() bool`

HasForestMappings returns a boolean if a field has been set.

### GetLabels

`func (o *HorizonExportItems) GetLabels() []LabelResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HorizonExportItems) GetLabelsOk() (*[]LabelResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HorizonExportItems) SetLabels(v []LabelResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HorizonExportItems) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNotifications

`func (o *HorizonExportItems) GetNotifications() []HorizonExportItemsNotificationsInner`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *HorizonExportItems) GetNotificationsOk() (*[]HorizonExportItemsNotificationsInner, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *HorizonExportItems) SetNotifications(v []HorizonExportItemsNotificationsInner)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *HorizonExportItems) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetPasswordPolicies

`func (o *HorizonExportItems) GetPasswordPolicies() []PasswordPolicyResponse`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *HorizonExportItems) GetPasswordPoliciesOk() (*[]PasswordPolicyResponse, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicies

`func (o *HorizonExportItems) SetPasswordPolicies(v []PasswordPolicyResponse)`

SetPasswordPolicies sets PasswordPolicies field to given value.

### HasPasswordPolicies

`func (o *HorizonExportItems) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.

### GetPkiConnectors

`func (o *HorizonExportItems) GetPkiConnectors() []PKIResponses`

GetPkiConnectors returns the PkiConnectors field if non-nil, zero value otherwise.

### GetPkiConnectorsOk

`func (o *HorizonExportItems) GetPkiConnectorsOk() (*[]PKIResponses, bool)`

GetPkiConnectorsOk returns a tuple with the PkiConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiConnectors

`func (o *HorizonExportItems) SetPkiConnectors(v []PKIResponses)`

SetPkiConnectors sets PkiConnectors field to given value.

### HasPkiConnectors

`func (o *HorizonExportItems) HasPkiConnectors() bool`

HasPkiConnectors returns a boolean if a field has been set.

### GetPkiQueues

`func (o *HorizonExportItems) GetPkiQueues() []PKIQueueResponse`

GetPkiQueues returns the PkiQueues field if non-nil, zero value otherwise.

### GetPkiQueuesOk

`func (o *HorizonExportItems) GetPkiQueuesOk() (*[]PKIQueueResponse, bool)`

GetPkiQueuesOk returns a tuple with the PkiQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiQueues

`func (o *HorizonExportItems) SetPkiQueues(v []PKIQueueResponse)`

SetPkiQueues sets PkiQueues field to given value.

### HasPkiQueues

`func (o *HorizonExportItems) HasPkiQueues() bool`

HasPkiQueues returns a boolean if a field has been set.

### GetProfiles

`func (o *HorizonExportItems) GetProfiles() []CertificateProfileResponses`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *HorizonExportItems) GetProfilesOk() (*[]CertificateProfileResponses, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *HorizonExportItems) SetProfiles(v []CertificateProfileResponses)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *HorizonExportItems) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetProxies

`func (o *HorizonExportItems) GetProxies() []HttpProxyResponse`

GetProxies returns the Proxies field if non-nil, zero value otherwise.

### GetProxiesOk

`func (o *HorizonExportItems) GetProxiesOk() (*[]HttpProxyResponse, bool)`

GetProxiesOk returns a tuple with the Proxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxies

`func (o *HorizonExportItems) SetProxies(v []HttpProxyResponse)`

SetProxies sets Proxies field to given value.

### HasProxies

`func (o *HorizonExportItems) HasProxies() bool`

HasProxies returns a boolean if a field has been set.

### GetReports

`func (o *HorizonExportItems) GetReports() []ReportScheduledTaskResponse`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *HorizonExportItems) GetReportsOk() (*[]ReportScheduledTaskResponse, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *HorizonExportItems) SetReports(v []ReportScheduledTaskResponse)`

SetReports sets Reports field to given value.

### HasReports

`func (o *HorizonExportItems) HasReports() bool`

HasReports returns a boolean if a field has been set.

### GetRoles

`func (o *HorizonExportItems) GetRoles() []RoleResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *HorizonExportItems) GetRolesOk() (*[]RoleResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *HorizonExportItems) SetRoles(v []RoleResponse)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *HorizonExportItems) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetScimProfiles

`func (o *HorizonExportItems) GetScimProfiles() []ScimProfileResponse`

GetScimProfiles returns the ScimProfiles field if non-nil, zero value otherwise.

### GetScimProfilesOk

`func (o *HorizonExportItems) GetScimProfilesOk() (*[]ScimProfileResponse, bool)`

GetScimProfilesOk returns a tuple with the ScimProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimProfiles

`func (o *HorizonExportItems) SetScimProfiles(v []ScimProfileResponse)`

SetScimProfiles sets ScimProfiles field to given value.

### HasScimProfiles

`func (o *HorizonExportItems) HasScimProfiles() bool`

HasScimProfiles returns a boolean if a field has been set.

### GetStorages

`func (o *HorizonExportItems) GetStorages() []S3StorageBackendConfigResponse`

GetStorages returns the Storages field if non-nil, zero value otherwise.

### GetStoragesOk

`func (o *HorizonExportItems) GetStoragesOk() (*[]S3StorageBackendConfigResponse, bool)`

GetStoragesOk returns a tuple with the Storages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorages

`func (o *HorizonExportItems) SetStorages(v []S3StorageBackendConfigResponse)`

SetStorages sets Storages field to given value.

### HasStorages

`func (o *HorizonExportItems) HasStorages() bool`

HasStorages returns a boolean if a field has been set.

### GetTeams

`func (o *HorizonExportItems) GetTeams() []TeamResponse`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *HorizonExportItems) GetTeamsOk() (*[]TeamResponse, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *HorizonExportItems) SetTeams(v []TeamResponse)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *HorizonExportItems) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetThirdParties

`func (o *HorizonExportItems) GetThirdParties() []ThirdPartyConnectorResponses`

GetThirdParties returns the ThirdParties field if non-nil, zero value otherwise.

### GetThirdPartiesOk

`func (o *HorizonExportItems) GetThirdPartiesOk() (*[]ThirdPartyConnectorResponses, bool)`

GetThirdPartiesOk returns a tuple with the ThirdParties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParties

`func (o *HorizonExportItems) SetThirdParties(v []ThirdPartyConnectorResponses)`

SetThirdParties sets ThirdParties field to given value.

### HasThirdParties

`func (o *HorizonExportItems) HasThirdParties() bool`

HasThirdParties returns a boolean if a field has been set.

### GetTriggers

`func (o *HorizonExportItems) GetTriggers() []HorizonExportItemsTriggersInner`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *HorizonExportItems) GetTriggersOk() (*[]HorizonExportItemsTriggersInner, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *HorizonExportItems) SetTriggers(v []HorizonExportItemsTriggersInner)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *HorizonExportItems) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


