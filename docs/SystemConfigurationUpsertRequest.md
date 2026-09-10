# SystemConfigurationUpsertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggers** | Pointer to [**NullableLicenseTriggers**](LicenseTriggers.md) | Triggers to execute on license events | [optional] 
**Type** | **string** | The type of the configuration entry | 
**Cron** | **string** | Cron defining when to run internal monitor checks | 
**Announcements** | Pointer to [**[]InterfaceCustomizationConfigurationAnnouncementsInner**](InterfaceCustomizationConfigurationAnnouncementsInner.md) | Announcements to display to all users in the Horizon instance | [optional] 
**HeaderEnd** | Pointer to **NullableString** | The HTML color code for the right side of the banner gradient | [optional] 
**HeaderStart** | Pointer to **NullableString** | The HTML color code for the left side of the banner gradient | [optional] 
**Logo** | Pointer to **NullableString** | A logo to display on the product, base64 encoded | [optional] 
**ArchiveStorage** | Pointer to **string** | Name of a [storage](#tag/system.storage) to use for archive file storage | [optional] 
**MagicLinkReportStorage** | Pointer to **string** | Name of a [storage](#tag/system.storage) to use for magic link reports storage | [optional] 

## Methods

### NewSystemConfigurationUpsertRequest

`func NewSystemConfigurationUpsertRequest(type_ string, cron string, ) *SystemConfigurationUpsertRequest`

NewSystemConfigurationUpsertRequest instantiates a new SystemConfigurationUpsertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemConfigurationUpsertRequestWithDefaults

`func NewSystemConfigurationUpsertRequestWithDefaults() *SystemConfigurationUpsertRequest`

NewSystemConfigurationUpsertRequestWithDefaults instantiates a new SystemConfigurationUpsertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggers

`func (o *SystemConfigurationUpsertRequest) GetTriggers() LicenseTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *SystemConfigurationUpsertRequest) GetTriggersOk() (*LicenseTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *SystemConfigurationUpsertRequest) SetTriggers(v LicenseTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *SystemConfigurationUpsertRequest) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *SystemConfigurationUpsertRequest) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *SystemConfigurationUpsertRequest) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetType

`func (o *SystemConfigurationUpsertRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemConfigurationUpsertRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemConfigurationUpsertRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *SystemConfigurationUpsertRequest) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *SystemConfigurationUpsertRequest) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *SystemConfigurationUpsertRequest) SetCron(v string)`

SetCron sets Cron field to given value.


### GetAnnouncements

`func (o *SystemConfigurationUpsertRequest) GetAnnouncements() []InterfaceCustomizationConfigurationAnnouncementsInner`

GetAnnouncements returns the Announcements field if non-nil, zero value otherwise.

### GetAnnouncementsOk

`func (o *SystemConfigurationUpsertRequest) GetAnnouncementsOk() (*[]InterfaceCustomizationConfigurationAnnouncementsInner, bool)`

GetAnnouncementsOk returns a tuple with the Announcements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncements

`func (o *SystemConfigurationUpsertRequest) SetAnnouncements(v []InterfaceCustomizationConfigurationAnnouncementsInner)`

SetAnnouncements sets Announcements field to given value.

### HasAnnouncements

`func (o *SystemConfigurationUpsertRequest) HasAnnouncements() bool`

HasAnnouncements returns a boolean if a field has been set.

### GetHeaderEnd

`func (o *SystemConfigurationUpsertRequest) GetHeaderEnd() string`

GetHeaderEnd returns the HeaderEnd field if non-nil, zero value otherwise.

### GetHeaderEndOk

`func (o *SystemConfigurationUpsertRequest) GetHeaderEndOk() (*string, bool)`

GetHeaderEndOk returns a tuple with the HeaderEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderEnd

`func (o *SystemConfigurationUpsertRequest) SetHeaderEnd(v string)`

SetHeaderEnd sets HeaderEnd field to given value.

### HasHeaderEnd

`func (o *SystemConfigurationUpsertRequest) HasHeaderEnd() bool`

HasHeaderEnd returns a boolean if a field has been set.

### SetHeaderEndNil

`func (o *SystemConfigurationUpsertRequest) SetHeaderEndNil(b bool)`

 SetHeaderEndNil sets the value for HeaderEnd to be an explicit nil

### UnsetHeaderEnd
`func (o *SystemConfigurationUpsertRequest) UnsetHeaderEnd()`

UnsetHeaderEnd ensures that no value is present for HeaderEnd, not even an explicit nil
### GetHeaderStart

`func (o *SystemConfigurationUpsertRequest) GetHeaderStart() string`

GetHeaderStart returns the HeaderStart field if non-nil, zero value otherwise.

### GetHeaderStartOk

`func (o *SystemConfigurationUpsertRequest) GetHeaderStartOk() (*string, bool)`

GetHeaderStartOk returns a tuple with the HeaderStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderStart

`func (o *SystemConfigurationUpsertRequest) SetHeaderStart(v string)`

SetHeaderStart sets HeaderStart field to given value.

### HasHeaderStart

`func (o *SystemConfigurationUpsertRequest) HasHeaderStart() bool`

HasHeaderStart returns a boolean if a field has been set.

### SetHeaderStartNil

`func (o *SystemConfigurationUpsertRequest) SetHeaderStartNil(b bool)`

 SetHeaderStartNil sets the value for HeaderStart to be an explicit nil

### UnsetHeaderStart
`func (o *SystemConfigurationUpsertRequest) UnsetHeaderStart()`

UnsetHeaderStart ensures that no value is present for HeaderStart, not even an explicit nil
### GetLogo

`func (o *SystemConfigurationUpsertRequest) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SystemConfigurationUpsertRequest) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SystemConfigurationUpsertRequest) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *SystemConfigurationUpsertRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *SystemConfigurationUpsertRequest) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *SystemConfigurationUpsertRequest) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetArchiveStorage

`func (o *SystemConfigurationUpsertRequest) GetArchiveStorage() string`

GetArchiveStorage returns the ArchiveStorage field if non-nil, zero value otherwise.

### GetArchiveStorageOk

`func (o *SystemConfigurationUpsertRequest) GetArchiveStorageOk() (*string, bool)`

GetArchiveStorageOk returns a tuple with the ArchiveStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveStorage

`func (o *SystemConfigurationUpsertRequest) SetArchiveStorage(v string)`

SetArchiveStorage sets ArchiveStorage field to given value.

### HasArchiveStorage

`func (o *SystemConfigurationUpsertRequest) HasArchiveStorage() bool`

HasArchiveStorage returns a boolean if a field has been set.

### GetMagicLinkReportStorage

`func (o *SystemConfigurationUpsertRequest) GetMagicLinkReportStorage() string`

GetMagicLinkReportStorage returns the MagicLinkReportStorage field if non-nil, zero value otherwise.

### GetMagicLinkReportStorageOk

`func (o *SystemConfigurationUpsertRequest) GetMagicLinkReportStorageOk() (*string, bool)`

GetMagicLinkReportStorageOk returns a tuple with the MagicLinkReportStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicLinkReportStorage

`func (o *SystemConfigurationUpsertRequest) SetMagicLinkReportStorage(v string)`

SetMagicLinkReportStorage sets MagicLinkReportStorage field to given value.

### HasMagicLinkReportStorage

`func (o *SystemConfigurationUpsertRequest) HasMagicLinkReportStorage() bool`

HasMagicLinkReportStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


