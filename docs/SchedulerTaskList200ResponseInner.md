# SchedulerTaskList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** |  | 
**Cron** | **string** |  | 
**Host** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**LastExecutionDate** | Pointer to **NullableInt64** |  | [optional] 
**LastCompletionDate** | Pointer to **NullableInt64** |  | [optional] 
**Detail** | Pointer to **NullableString** |  | [optional] 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | **bool** |  | 
**Name** | **string** |  | 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Recipients** | [**[]ReportRecipient**](ReportRecipient.md) |  | 
**From** | **string** |  | 
**Title** | **string** |  | 
**Body** | Pointer to **NullableString** |  | [optional] 
**IsHtml** | **bool** |  | 
**CompressCsv** | Pointer to **bool** | Should the report be compressed using GZ. It will divide by two the size of the csv | [optional] 
**HqlType** | **string** |  | 
**HqlQuery** | Pointer to **NullableString** |  | [optional] 
**HqlFields** | Pointer to **[]string** |  | [optional] 
**HqlSortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DryRun** | **bool** |  | 
**Module** | **string** |  | 
**Profile** | **string** |  | 
**Connector** | **string** |  | 
**Enroll** | **bool** |  | 
**Revoke** | **bool** |  | 
**Renew** | **bool** |  | 
**Results** | Pointer to [**NullableThirdPartyConnectorSynchronizationResult**](ThirdPartyConnectorSynchronizationResult.md) |  | [optional] 

## Methods

### NewSchedulerTaskList200ResponseInner

`func NewSchedulerTaskList200ResponseInner(id string, type_ string, cron string, enabled bool, name string, recipients []ReportRecipient, from string, title string, isHtml bool, hqlType string, dryRun bool, module string, profile string, connector string, enroll bool, revoke bool, renew bool, ) *SchedulerTaskList200ResponseInner`

NewSchedulerTaskList200ResponseInner instantiates a new SchedulerTaskList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTaskList200ResponseInnerWithDefaults

`func NewSchedulerTaskList200ResponseInnerWithDefaults() *SchedulerTaskList200ResponseInner`

NewSchedulerTaskList200ResponseInnerWithDefaults instantiates a new SchedulerTaskList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchedulerTaskList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchedulerTaskList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchedulerTaskList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SchedulerTaskList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchedulerTaskList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchedulerTaskList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *SchedulerTaskList200ResponseInner) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *SchedulerTaskList200ResponseInner) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *SchedulerTaskList200ResponseInner) SetCron(v string)`

SetCron sets Cron field to given value.


### GetHost

`func (o *SchedulerTaskList200ResponseInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SchedulerTaskList200ResponseInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SchedulerTaskList200ResponseInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SchedulerTaskList200ResponseInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SchedulerTaskList200ResponseInner) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SchedulerTaskList200ResponseInner) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetStatus

`func (o *SchedulerTaskList200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchedulerTaskList200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchedulerTaskList200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchedulerTaskList200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SchedulerTaskList200ResponseInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SchedulerTaskList200ResponseInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastExecutionDate

`func (o *SchedulerTaskList200ResponseInner) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *SchedulerTaskList200ResponseInner) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *SchedulerTaskList200ResponseInner) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *SchedulerTaskList200ResponseInner) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *SchedulerTaskList200ResponseInner) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *SchedulerTaskList200ResponseInner) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetLastCompletionDate

`func (o *SchedulerTaskList200ResponseInner) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *SchedulerTaskList200ResponseInner) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *SchedulerTaskList200ResponseInner) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *SchedulerTaskList200ResponseInner) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *SchedulerTaskList200ResponseInner) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *SchedulerTaskList200ResponseInner) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetDetail

`func (o *SchedulerTaskList200ResponseInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SchedulerTaskList200ResponseInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SchedulerTaskList200ResponseInner) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SchedulerTaskList200ResponseInner) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *SchedulerTaskList200ResponseInner) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *SchedulerTaskList200ResponseInner) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetExecutionId

`func (o *SchedulerTaskList200ResponseInner) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *SchedulerTaskList200ResponseInner) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *SchedulerTaskList200ResponseInner) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *SchedulerTaskList200ResponseInner) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *SchedulerTaskList200ResponseInner) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *SchedulerTaskList200ResponseInner) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetEnabled

`func (o *SchedulerTaskList200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SchedulerTaskList200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SchedulerTaskList200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *SchedulerTaskList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulerTaskList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulerTaskList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetFileName

`func (o *SchedulerTaskList200ResponseInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SchedulerTaskList200ResponseInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SchedulerTaskList200ResponseInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SchedulerTaskList200ResponseInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SchedulerTaskList200ResponseInner) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SchedulerTaskList200ResponseInner) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetRecipients

`func (o *SchedulerTaskList200ResponseInner) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *SchedulerTaskList200ResponseInner) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *SchedulerTaskList200ResponseInner) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetFrom

`func (o *SchedulerTaskList200ResponseInner) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SchedulerTaskList200ResponseInner) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SchedulerTaskList200ResponseInner) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTitle

`func (o *SchedulerTaskList200ResponseInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SchedulerTaskList200ResponseInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SchedulerTaskList200ResponseInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *SchedulerTaskList200ResponseInner) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SchedulerTaskList200ResponseInner) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SchedulerTaskList200ResponseInner) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SchedulerTaskList200ResponseInner) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *SchedulerTaskList200ResponseInner) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *SchedulerTaskList200ResponseInner) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetIsHtml

`func (o *SchedulerTaskList200ResponseInner) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *SchedulerTaskList200ResponseInner) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *SchedulerTaskList200ResponseInner) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetCompressCsv

`func (o *SchedulerTaskList200ResponseInner) GetCompressCsv() bool`

GetCompressCsv returns the CompressCsv field if non-nil, zero value otherwise.

### GetCompressCsvOk

`func (o *SchedulerTaskList200ResponseInner) GetCompressCsvOk() (*bool, bool)`

GetCompressCsvOk returns a tuple with the CompressCsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressCsv

`func (o *SchedulerTaskList200ResponseInner) SetCompressCsv(v bool)`

SetCompressCsv sets CompressCsv field to given value.

### HasCompressCsv

`func (o *SchedulerTaskList200ResponseInner) HasCompressCsv() bool`

HasCompressCsv returns a boolean if a field has been set.

### GetHqlType

`func (o *SchedulerTaskList200ResponseInner) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *SchedulerTaskList200ResponseInner) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *SchedulerTaskList200ResponseInner) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetHqlQuery

`func (o *SchedulerTaskList200ResponseInner) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *SchedulerTaskList200ResponseInner) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *SchedulerTaskList200ResponseInner) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *SchedulerTaskList200ResponseInner) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *SchedulerTaskList200ResponseInner) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *SchedulerTaskList200ResponseInner) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlFields

`func (o *SchedulerTaskList200ResponseInner) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *SchedulerTaskList200ResponseInner) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *SchedulerTaskList200ResponseInner) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *SchedulerTaskList200ResponseInner) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *SchedulerTaskList200ResponseInner) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *SchedulerTaskList200ResponseInner) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlSortedBy

`func (o *SchedulerTaskList200ResponseInner) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *SchedulerTaskList200ResponseInner) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *SchedulerTaskList200ResponseInner) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *SchedulerTaskList200ResponseInner) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *SchedulerTaskList200ResponseInner) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *SchedulerTaskList200ResponseInner) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil
### GetDescription

`func (o *SchedulerTaskList200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchedulerTaskList200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchedulerTaskList200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchedulerTaskList200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SchedulerTaskList200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SchedulerTaskList200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDryRun

`func (o *SchedulerTaskList200ResponseInner) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SchedulerTaskList200ResponseInner) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SchedulerTaskList200ResponseInner) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetModule

`func (o *SchedulerTaskList200ResponseInner) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *SchedulerTaskList200ResponseInner) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *SchedulerTaskList200ResponseInner) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *SchedulerTaskList200ResponseInner) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SchedulerTaskList200ResponseInner) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SchedulerTaskList200ResponseInner) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetConnector

`func (o *SchedulerTaskList200ResponseInner) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *SchedulerTaskList200ResponseInner) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *SchedulerTaskList200ResponseInner) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetEnroll

`func (o *SchedulerTaskList200ResponseInner) GetEnroll() bool`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *SchedulerTaskList200ResponseInner) GetEnrollOk() (*bool, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *SchedulerTaskList200ResponseInner) SetEnroll(v bool)`

SetEnroll sets Enroll field to given value.


### GetRevoke

`func (o *SchedulerTaskList200ResponseInner) GetRevoke() bool`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *SchedulerTaskList200ResponseInner) GetRevokeOk() (*bool, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *SchedulerTaskList200ResponseInner) SetRevoke(v bool)`

SetRevoke sets Revoke field to given value.


### GetRenew

`func (o *SchedulerTaskList200ResponseInner) GetRenew() bool`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *SchedulerTaskList200ResponseInner) GetRenewOk() (*bool, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *SchedulerTaskList200ResponseInner) SetRenew(v bool)`

SetRenew sets Renew field to given value.


### GetResults

`func (o *SchedulerTaskList200ResponseInner) GetResults() ThirdPartyConnectorSynchronizationResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SchedulerTaskList200ResponseInner) GetResultsOk() (*ThirdPartyConnectorSynchronizationResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SchedulerTaskList200ResponseInner) SetResults(v ThirdPartyConnectorSynchronizationResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *SchedulerTaskList200ResponseInner) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *SchedulerTaskList200ResponseInner) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *SchedulerTaskList200ResponseInner) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


