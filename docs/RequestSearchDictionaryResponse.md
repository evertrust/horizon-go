# RequestSearchDictionaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]CertificateProfileSearchDictionaryLocalizedEntry**](CertificateProfileSearchDictionaryLocalizedEntry.md) | The list of profiles the principal is authorized to search on | [optional] 
**Teams** | Pointer to [**[]TeamSearchDictionaryLocalizedEntry**](TeamSearchDictionaryLocalizedEntry.md) | The list of available teams on this Horizon instance | [optional] 
**Labels** | Pointer to [**[]CertificateLabelSearchDictionaryLocalizedEntry**](CertificateLabelSearchDictionaryLocalizedEntry.md) | The list of labels the principal is authorized to search on | [optional] 
**Metadata** | **[]string** | The list of available metadata in Horizon | 
**Modules** | Pointer to **[]string** | The list of Horizon modules available on this instance | [optional] 

## Methods

### NewRequestSearchDictionaryResponse

`func NewRequestSearchDictionaryResponse(metadata []string, ) *RequestSearchDictionaryResponse`

NewRequestSearchDictionaryResponse instantiates a new RequestSearchDictionaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSearchDictionaryResponseWithDefaults

`func NewRequestSearchDictionaryResponseWithDefaults() *RequestSearchDictionaryResponse`

NewRequestSearchDictionaryResponseWithDefaults instantiates a new RequestSearchDictionaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *RequestSearchDictionaryResponse) GetProfiles() []CertificateProfileSearchDictionaryLocalizedEntry`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *RequestSearchDictionaryResponse) GetProfilesOk() (*[]CertificateProfileSearchDictionaryLocalizedEntry, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *RequestSearchDictionaryResponse) SetProfiles(v []CertificateProfileSearchDictionaryLocalizedEntry)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *RequestSearchDictionaryResponse) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *RequestSearchDictionaryResponse) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *RequestSearchDictionaryResponse) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetTeams

`func (o *RequestSearchDictionaryResponse) GetTeams() []TeamSearchDictionaryLocalizedEntry`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *RequestSearchDictionaryResponse) GetTeamsOk() (*[]TeamSearchDictionaryLocalizedEntry, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *RequestSearchDictionaryResponse) SetTeams(v []TeamSearchDictionaryLocalizedEntry)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *RequestSearchDictionaryResponse) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### SetTeamsNil

`func (o *RequestSearchDictionaryResponse) SetTeamsNil(b bool)`

 SetTeamsNil sets the value for Teams to be an explicit nil

### UnsetTeams
`func (o *RequestSearchDictionaryResponse) UnsetTeams()`

UnsetTeams ensures that no value is present for Teams, not even an explicit nil
### GetLabels

`func (o *RequestSearchDictionaryResponse) GetLabels() []CertificateLabelSearchDictionaryLocalizedEntry`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RequestSearchDictionaryResponse) GetLabelsOk() (*[]CertificateLabelSearchDictionaryLocalizedEntry, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RequestSearchDictionaryResponse) SetLabels(v []CertificateLabelSearchDictionaryLocalizedEntry)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RequestSearchDictionaryResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *RequestSearchDictionaryResponse) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *RequestSearchDictionaryResponse) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *RequestSearchDictionaryResponse) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequestSearchDictionaryResponse) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequestSearchDictionaryResponse) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.


### GetModules

`func (o *RequestSearchDictionaryResponse) GetModules() []string`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *RequestSearchDictionaryResponse) GetModulesOk() (*[]string, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *RequestSearchDictionaryResponse) SetModules(v []string)`

SetModules sets Modules field to given value.

### HasModules

`func (o *RequestSearchDictionaryResponse) HasModules() bool`

HasModules returns a boolean if a field has been set.

### SetModulesNil

`func (o *RequestSearchDictionaryResponse) SetModulesNil(b bool)`

 SetModulesNil sets the value for Modules to be an explicit nil

### UnsetModules
`func (o *RequestSearchDictionaryResponse) UnsetModules()`

UnsetModules ensures that no value is present for Modules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


