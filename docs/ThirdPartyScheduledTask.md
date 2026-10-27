# ThirdPartyScheduledTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DryRun** | **bool** |  | 
**Enroll** | **bool** |  | 
**Module** | **string** |  | 
**Profile** | **string** |  | 
**Renew** | Pointer to **bool** |  | [optional] 
**Revoke** | **bool** |  | 
**Type** | **string** |  | 
**Cron** | **string** |  | 
**Enabled** | **bool** |  | 
**Name** | **string** |  | 

## Methods

### NewThirdPartyScheduledTask

`func NewThirdPartyScheduledTask(connector string, dryRun bool, enroll bool, module string, profile string, revoke bool, type_ string, cron string, enabled bool, name string, ) *ThirdPartyScheduledTask`

NewThirdPartyScheduledTask instantiates a new ThirdPartyScheduledTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyScheduledTaskWithDefaults

`func NewThirdPartyScheduledTaskWithDefaults() *ThirdPartyScheduledTask`

NewThirdPartyScheduledTaskWithDefaults instantiates a new ThirdPartyScheduledTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ThirdPartyScheduledTask) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ThirdPartyScheduledTask) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ThirdPartyScheduledTask) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetDescription

`func (o *ThirdPartyScheduledTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyScheduledTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyScheduledTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyScheduledTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ThirdPartyScheduledTask) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ThirdPartyScheduledTask) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDryRun

`func (o *ThirdPartyScheduledTask) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ThirdPartyScheduledTask) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ThirdPartyScheduledTask) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetEnroll

`func (o *ThirdPartyScheduledTask) GetEnroll() bool`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *ThirdPartyScheduledTask) GetEnrollOk() (*bool, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *ThirdPartyScheduledTask) SetEnroll(v bool)`

SetEnroll sets Enroll field to given value.


### GetModule

`func (o *ThirdPartyScheduledTask) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ThirdPartyScheduledTask) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ThirdPartyScheduledTask) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *ThirdPartyScheduledTask) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ThirdPartyScheduledTask) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ThirdPartyScheduledTask) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRenew

`func (o *ThirdPartyScheduledTask) GetRenew() bool`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *ThirdPartyScheduledTask) GetRenewOk() (*bool, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *ThirdPartyScheduledTask) SetRenew(v bool)`

SetRenew sets Renew field to given value.

### HasRenew

`func (o *ThirdPartyScheduledTask) HasRenew() bool`

HasRenew returns a boolean if a field has been set.

### GetRevoke

`func (o *ThirdPartyScheduledTask) GetRevoke() bool`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *ThirdPartyScheduledTask) GetRevokeOk() (*bool, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *ThirdPartyScheduledTask) SetRevoke(v bool)`

SetRevoke sets Revoke field to given value.


### GetType

`func (o *ThirdPartyScheduledTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThirdPartyScheduledTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThirdPartyScheduledTask) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *ThirdPartyScheduledTask) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ThirdPartyScheduledTask) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ThirdPartyScheduledTask) SetCron(v string)`

SetCron sets Cron field to given value.


### GetEnabled

`func (o *ThirdPartyScheduledTask) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyScheduledTask) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyScheduledTask) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *ThirdPartyScheduledTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThirdPartyScheduledTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThirdPartyScheduledTask) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


