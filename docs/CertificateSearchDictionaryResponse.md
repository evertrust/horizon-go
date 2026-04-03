# CertificateSearchDictionaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaigns** | Pointer to **[]string** | The list of discovery campaign the principal is authorized to search on | [optional] 
**GradingPolicies** | Pointer to **[]string** | The list of available grading policies on Horizon | [optional] 
**Labels** | Pointer to [**[]CertificateLabelSearchDictionaryLocalizedEntry**](CertificateLabelSearchDictionaryLocalizedEntry.md) | The list of labels the principal is authorized to search on | [optional] 
**Metadata** | **[]string** | The list of available metadata in Horizon | 
**Modules** | Pointer to **[]string** | The list of Horizon modules available on this instance | [optional] 
**Profiles** | Pointer to [**[]CertificateProfileSearchDictionaryLocalizedEntry**](CertificateProfileSearchDictionaryLocalizedEntry.md) | The list of profiles the principal is authorized to search on | [optional] 
**Teams** | Pointer to [**[]TeamSearchDictionaryLocalizedEntry**](TeamSearchDictionaryLocalizedEntry.md) | The list of available teams on this Horizon instance | [optional] 

## Methods

### NewCertificateSearchDictionaryResponse

`func NewCertificateSearchDictionaryResponse(metadata []string, ) *CertificateSearchDictionaryResponse`

NewCertificateSearchDictionaryResponse instantiates a new CertificateSearchDictionaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSearchDictionaryResponseWithDefaults

`func NewCertificateSearchDictionaryResponseWithDefaults() *CertificateSearchDictionaryResponse`

NewCertificateSearchDictionaryResponseWithDefaults instantiates a new CertificateSearchDictionaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaigns

`func (o *CertificateSearchDictionaryResponse) GetCampaigns() []string`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *CertificateSearchDictionaryResponse) GetCampaignsOk() (*[]string, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *CertificateSearchDictionaryResponse) SetCampaigns(v []string)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *CertificateSearchDictionaryResponse) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### SetCampaignsNil

`func (o *CertificateSearchDictionaryResponse) SetCampaignsNil(b bool)`

 SetCampaignsNil sets the value for Campaigns to be an explicit nil

### UnsetCampaigns
`func (o *CertificateSearchDictionaryResponse) UnsetCampaigns()`

UnsetCampaigns ensures that no value is present for Campaigns, not even an explicit nil
### GetGradingPolicies

`func (o *CertificateSearchDictionaryResponse) GetGradingPolicies() []string`

GetGradingPolicies returns the GradingPolicies field if non-nil, zero value otherwise.

### GetGradingPoliciesOk

`func (o *CertificateSearchDictionaryResponse) GetGradingPoliciesOk() (*[]string, bool)`

GetGradingPoliciesOk returns a tuple with the GradingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingPolicies

`func (o *CertificateSearchDictionaryResponse) SetGradingPolicies(v []string)`

SetGradingPolicies sets GradingPolicies field to given value.

### HasGradingPolicies

`func (o *CertificateSearchDictionaryResponse) HasGradingPolicies() bool`

HasGradingPolicies returns a boolean if a field has been set.

### SetGradingPoliciesNil

`func (o *CertificateSearchDictionaryResponse) SetGradingPoliciesNil(b bool)`

 SetGradingPoliciesNil sets the value for GradingPolicies to be an explicit nil

### UnsetGradingPolicies
`func (o *CertificateSearchDictionaryResponse) UnsetGradingPolicies()`

UnsetGradingPolicies ensures that no value is present for GradingPolicies, not even an explicit nil
### GetLabels

`func (o *CertificateSearchDictionaryResponse) GetLabels() []CertificateLabelSearchDictionaryLocalizedEntry`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CertificateSearchDictionaryResponse) GetLabelsOk() (*[]CertificateLabelSearchDictionaryLocalizedEntry, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CertificateSearchDictionaryResponse) SetLabels(v []CertificateLabelSearchDictionaryLocalizedEntry)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CertificateSearchDictionaryResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *CertificateSearchDictionaryResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *CertificateSearchDictionaryResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *CertificateSearchDictionaryResponse) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificateSearchDictionaryResponse) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificateSearchDictionaryResponse) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.


### GetModules

`func (o *CertificateSearchDictionaryResponse) GetModules() []string`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *CertificateSearchDictionaryResponse) GetModulesOk() (*[]string, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *CertificateSearchDictionaryResponse) SetModules(v []string)`

SetModules sets Modules field to given value.

### HasModules

`func (o *CertificateSearchDictionaryResponse) HasModules() bool`

HasModules returns a boolean if a field has been set.

### SetModulesNil

`func (o *CertificateSearchDictionaryResponse) SetModulesNil(b bool)`

 SetModulesNil sets the value for Modules to be an explicit nil

### UnsetModules
`func (o *CertificateSearchDictionaryResponse) UnsetModules()`

UnsetModules ensures that no value is present for Modules, not even an explicit nil
### GetProfiles

`func (o *CertificateSearchDictionaryResponse) GetProfiles() []CertificateProfileSearchDictionaryLocalizedEntry`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *CertificateSearchDictionaryResponse) GetProfilesOk() (*[]CertificateProfileSearchDictionaryLocalizedEntry, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *CertificateSearchDictionaryResponse) SetProfiles(v []CertificateProfileSearchDictionaryLocalizedEntry)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *CertificateSearchDictionaryResponse) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *CertificateSearchDictionaryResponse) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *CertificateSearchDictionaryResponse) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetTeams

`func (o *CertificateSearchDictionaryResponse) GetTeams() []TeamSearchDictionaryLocalizedEntry`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *CertificateSearchDictionaryResponse) GetTeamsOk() (*[]TeamSearchDictionaryLocalizedEntry, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *CertificateSearchDictionaryResponse) SetTeams(v []TeamSearchDictionaryLocalizedEntry)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *CertificateSearchDictionaryResponse) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### SetTeamsNil

`func (o *CertificateSearchDictionaryResponse) SetTeamsNil(b bool)`

 SetTeamsNil sets the value for Teams to be an explicit nil

### UnsetTeams
`func (o *CertificateSearchDictionaryResponse) UnsetTeams()`

UnsetTeams ensures that no value is present for Teams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


