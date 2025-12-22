# TeamSearchDictionaryLocalizedEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the team | [optional] 
**Name** | **string** | The technical name of the team | 

## Methods

### NewTeamSearchDictionaryLocalizedEntry

`func NewTeamSearchDictionaryLocalizedEntry(name string, ) *TeamSearchDictionaryLocalizedEntry`

NewTeamSearchDictionaryLocalizedEntry instantiates a new TeamSearchDictionaryLocalizedEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamSearchDictionaryLocalizedEntryWithDefaults

`func NewTeamSearchDictionaryLocalizedEntryWithDefaults() *TeamSearchDictionaryLocalizedEntry`

NewTeamSearchDictionaryLocalizedEntryWithDefaults instantiates a new TeamSearchDictionaryLocalizedEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TeamSearchDictionaryLocalizedEntry) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TeamSearchDictionaryLocalizedEntry) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TeamSearchDictionaryLocalizedEntry) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TeamSearchDictionaryLocalizedEntry) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TeamSearchDictionaryLocalizedEntry) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TeamSearchDictionaryLocalizedEntry) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetName

`func (o *TeamSearchDictionaryLocalizedEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamSearchDictionaryLocalizedEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamSearchDictionaryLocalizedEntry) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


