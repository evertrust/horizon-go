# SystemConfigurationList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** | The type of the configuration entry | 
**Triggers** | Pointer to [**NullableLicenseTriggers**](LicenseTriggers.md) | Triggers to execute on license events | [optional] 
**Cron** | **string** | Cron defining when to run internal monitor checks | 
**Logo** | Pointer to **NullableString** | A logo to display on the product, base64 encoded | [optional] 
**HeaderStart** | Pointer to **NullableString** | The HTML color code for the left side of the banner gradient | [optional] 
**HeaderEnd** | Pointer to **NullableString** | The HTML color code for the right side of the banner gradient | [optional] 

## Methods

### NewSystemConfigurationList200ResponseInner

`func NewSystemConfigurationList200ResponseInner(id string, type_ string, cron string, ) *SystemConfigurationList200ResponseInner`

NewSystemConfigurationList200ResponseInner instantiates a new SystemConfigurationList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemConfigurationList200ResponseInnerWithDefaults

`func NewSystemConfigurationList200ResponseInnerWithDefaults() *SystemConfigurationList200ResponseInner`

NewSystemConfigurationList200ResponseInnerWithDefaults instantiates a new SystemConfigurationList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemConfigurationList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemConfigurationList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemConfigurationList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SystemConfigurationList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemConfigurationList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemConfigurationList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetTriggers

`func (o *SystemConfigurationList200ResponseInner) GetTriggers() LicenseTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *SystemConfigurationList200ResponseInner) GetTriggersOk() (*LicenseTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *SystemConfigurationList200ResponseInner) SetTriggers(v LicenseTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *SystemConfigurationList200ResponseInner) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *SystemConfigurationList200ResponseInner) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *SystemConfigurationList200ResponseInner) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetCron

`func (o *SystemConfigurationList200ResponseInner) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *SystemConfigurationList200ResponseInner) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *SystemConfigurationList200ResponseInner) SetCron(v string)`

SetCron sets Cron field to given value.


### GetLogo

`func (o *SystemConfigurationList200ResponseInner) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SystemConfigurationList200ResponseInner) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SystemConfigurationList200ResponseInner) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *SystemConfigurationList200ResponseInner) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *SystemConfigurationList200ResponseInner) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *SystemConfigurationList200ResponseInner) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetHeaderStart

`func (o *SystemConfigurationList200ResponseInner) GetHeaderStart() string`

GetHeaderStart returns the HeaderStart field if non-nil, zero value otherwise.

### GetHeaderStartOk

`func (o *SystemConfigurationList200ResponseInner) GetHeaderStartOk() (*string, bool)`

GetHeaderStartOk returns a tuple with the HeaderStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderStart

`func (o *SystemConfigurationList200ResponseInner) SetHeaderStart(v string)`

SetHeaderStart sets HeaderStart field to given value.

### HasHeaderStart

`func (o *SystemConfigurationList200ResponseInner) HasHeaderStart() bool`

HasHeaderStart returns a boolean if a field has been set.

### SetHeaderStartNil

`func (o *SystemConfigurationList200ResponseInner) SetHeaderStartNil(b bool)`

 SetHeaderStartNil sets the value for HeaderStart to be an explicit nil

### UnsetHeaderStart
`func (o *SystemConfigurationList200ResponseInner) UnsetHeaderStart()`

UnsetHeaderStart ensures that no value is present for HeaderStart, not even an explicit nil
### GetHeaderEnd

`func (o *SystemConfigurationList200ResponseInner) GetHeaderEnd() string`

GetHeaderEnd returns the HeaderEnd field if non-nil, zero value otherwise.

### GetHeaderEndOk

`func (o *SystemConfigurationList200ResponseInner) GetHeaderEndOk() (*string, bool)`

GetHeaderEndOk returns a tuple with the HeaderEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderEnd

`func (o *SystemConfigurationList200ResponseInner) SetHeaderEnd(v string)`

SetHeaderEnd sets HeaderEnd field to given value.

### HasHeaderEnd

`func (o *SystemConfigurationList200ResponseInner) HasHeaderEnd() bool`

HasHeaderEnd returns a boolean if a field has been set.

### SetHeaderEndNil

`func (o *SystemConfigurationList200ResponseInner) SetHeaderEndNil(b bool)`

 SetHeaderEndNil sets the value for HeaderEnd to be an explicit nil

### UnsetHeaderEnd
`func (o *SystemConfigurationList200ResponseInner) UnsetHeaderEnd()`

UnsetHeaderEnd ensures that no value is present for HeaderEnd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


