# ThirdPartyScheduledTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Connector** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DryRun** | **bool** |  | 
**Enroll** | **bool** |  | 
**Module** | **string** |  | 
**Profile** | **string** |  | 
**Renew** | **bool** |  | 
**Revoke** | **bool** |  | 
**Type** | **string** |  | 
**Cron** | **string** |  | 
**Detail** | Pointer to **NullableString** |  | [optional] 
**Enabled** | **bool** |  | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**LastCompletionDate** | Pointer to **NullableInt64** |  | [optional] 
**LastExecutionDate** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewThirdPartyScheduledTaskResponse

`func NewThirdPartyScheduledTaskResponse(id string, connector string, dryRun bool, enroll bool, module string, profile string, renew bool, revoke bool, type_ string, cron string, enabled bool, name string, ) *ThirdPartyScheduledTaskResponse`

NewThirdPartyScheduledTaskResponse instantiates a new ThirdPartyScheduledTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyScheduledTaskResponseWithDefaults

`func NewThirdPartyScheduledTaskResponseWithDefaults() *ThirdPartyScheduledTaskResponse`

NewThirdPartyScheduledTaskResponseWithDefaults instantiates a new ThirdPartyScheduledTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyScheduledTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyScheduledTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyScheduledTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetConnector

`func (o *ThirdPartyScheduledTaskResponse) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ThirdPartyScheduledTaskResponse) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ThirdPartyScheduledTaskResponse) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetDescription

`func (o *ThirdPartyScheduledTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThirdPartyScheduledTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThirdPartyScheduledTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThirdPartyScheduledTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *ThirdPartyScheduledTaskResponse) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ThirdPartyScheduledTaskResponse) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ThirdPartyScheduledTaskResponse) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetEnroll

`func (o *ThirdPartyScheduledTaskResponse) GetEnroll() bool`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *ThirdPartyScheduledTaskResponse) GetEnrollOk() (*bool, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *ThirdPartyScheduledTaskResponse) SetEnroll(v bool)`

SetEnroll sets Enroll field to given value.


### GetModule

`func (o *ThirdPartyScheduledTaskResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ThirdPartyScheduledTaskResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ThirdPartyScheduledTaskResponse) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *ThirdPartyScheduledTaskResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ThirdPartyScheduledTaskResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ThirdPartyScheduledTaskResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRenew

`func (o *ThirdPartyScheduledTaskResponse) GetRenew() bool`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *ThirdPartyScheduledTaskResponse) GetRenewOk() (*bool, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *ThirdPartyScheduledTaskResponse) SetRenew(v bool)`

SetRenew sets Renew field to given value.


### GetRevoke

`func (o *ThirdPartyScheduledTaskResponse) GetRevoke() bool`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *ThirdPartyScheduledTaskResponse) GetRevokeOk() (*bool, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *ThirdPartyScheduledTaskResponse) SetRevoke(v bool)`

SetRevoke sets Revoke field to given value.


### GetType

`func (o *ThirdPartyScheduledTaskResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThirdPartyScheduledTaskResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThirdPartyScheduledTaskResponse) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *ThirdPartyScheduledTaskResponse) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ThirdPartyScheduledTaskResponse) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ThirdPartyScheduledTaskResponse) SetCron(v string)`

SetCron sets Cron field to given value.


### GetDetail

`func (o *ThirdPartyScheduledTaskResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ThirdPartyScheduledTaskResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ThirdPartyScheduledTaskResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ThirdPartyScheduledTaskResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ThirdPartyScheduledTaskResponse) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ThirdPartyScheduledTaskResponse) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetEnabled

`func (o *ThirdPartyScheduledTaskResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyScheduledTaskResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyScheduledTaskResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExecutionId

`func (o *ThirdPartyScheduledTaskResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ThirdPartyScheduledTaskResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ThirdPartyScheduledTaskResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ThirdPartyScheduledTaskResponse) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *ThirdPartyScheduledTaskResponse) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *ThirdPartyScheduledTaskResponse) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetHost

`func (o *ThirdPartyScheduledTaskResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ThirdPartyScheduledTaskResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ThirdPartyScheduledTaskResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ThirdPartyScheduledTaskResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *ThirdPartyScheduledTaskResponse) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *ThirdPartyScheduledTaskResponse) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLastCompletionDate

`func (o *ThirdPartyScheduledTaskResponse) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *ThirdPartyScheduledTaskResponse) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *ThirdPartyScheduledTaskResponse) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *ThirdPartyScheduledTaskResponse) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *ThirdPartyScheduledTaskResponse) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *ThirdPartyScheduledTaskResponse) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetLastExecutionDate

`func (o *ThirdPartyScheduledTaskResponse) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *ThirdPartyScheduledTaskResponse) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *ThirdPartyScheduledTaskResponse) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *ThirdPartyScheduledTaskResponse) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *ThirdPartyScheduledTaskResponse) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *ThirdPartyScheduledTaskResponse) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetName

`func (o *ThirdPartyScheduledTaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThirdPartyScheduledTaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThirdPartyScheduledTaskResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ThirdPartyScheduledTaskResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ThirdPartyScheduledTaskResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ThirdPartyScheduledTaskResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ThirdPartyScheduledTaskResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ThirdPartyScheduledTaskResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ThirdPartyScheduledTaskResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


