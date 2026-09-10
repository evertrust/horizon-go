# ScheduledTaskResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ReportType** | **string** |  | 
**RetentionPeriod** | **string** | Indicates the duration during which the report can be downloaded | 
**Body** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**From** | **string** |  | 
**HqlFields** | Pointer to **[]string** |  | [optional] 
**HqlQuery** | Pointer to **NullableString** |  | [optional] 
**HqlSortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**HqlType** | **string** |  | 
**IsHtml** | **bool** |  | 
**Recipients** | [**[]ReportRecipient**](ReportRecipient.md) |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Cron** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **NullableString** |  | [optional] 
**Enabled** | **bool** |  | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**LastCompletionDate** | Pointer to **NullableInt64** |  | [optional] 
**LastExecutionDate** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Status** | Pointer to **NullableString** |  | [optional] 
**CompressCsv** | Pointer to **bool** | Should the report be compressed using GZ | [optional] 
**Connector** | **string** |  | 
**DryRun** | **bool** |  | 
**Enroll** | **bool** |  | 
**Module** | **string** |  | 
**Profile** | **string** |  | 
**Renew** | Pointer to **bool** |  | [optional] 
**Revoke** | **bool** |  | 

## Methods

### NewScheduledTaskResponses

`func NewScheduledTaskResponses(id string, reportType string, retentionPeriod string, from string, hqlType string, isHtml bool, recipients []ReportRecipient, title string, type_ string, cron string, enabled bool, name string, connector string, dryRun bool, enroll bool, module string, profile string, revoke bool, ) *ScheduledTaskResponses`

NewScheduledTaskResponses instantiates a new ScheduledTaskResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledTaskResponsesWithDefaults

`func NewScheduledTaskResponsesWithDefaults() *ScheduledTaskResponses`

NewScheduledTaskResponsesWithDefaults instantiates a new ScheduledTaskResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledTaskResponses) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledTaskResponses) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledTaskResponses) SetId(v string)`

SetId sets Id field to given value.


### GetReportType

`func (o *ScheduledTaskResponses) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ScheduledTaskResponses) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ScheduledTaskResponses) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetRetentionPeriod

`func (o *ScheduledTaskResponses) GetRetentionPeriod() string`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *ScheduledTaskResponses) GetRetentionPeriodOk() (*string, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *ScheduledTaskResponses) SetRetentionPeriod(v string)`

SetRetentionPeriod sets RetentionPeriod field to given value.


### GetBody

`func (o *ScheduledTaskResponses) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ScheduledTaskResponses) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ScheduledTaskResponses) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ScheduledTaskResponses) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ScheduledTaskResponses) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ScheduledTaskResponses) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetFileName

`func (o *ScheduledTaskResponses) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ScheduledTaskResponses) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ScheduledTaskResponses) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ScheduledTaskResponses) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ScheduledTaskResponses) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ScheduledTaskResponses) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFrom

`func (o *ScheduledTaskResponses) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ScheduledTaskResponses) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ScheduledTaskResponses) SetFrom(v string)`

SetFrom sets From field to given value.


### GetHqlFields

`func (o *ScheduledTaskResponses) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *ScheduledTaskResponses) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *ScheduledTaskResponses) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *ScheduledTaskResponses) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *ScheduledTaskResponses) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *ScheduledTaskResponses) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlQuery

`func (o *ScheduledTaskResponses) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *ScheduledTaskResponses) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *ScheduledTaskResponses) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *ScheduledTaskResponses) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *ScheduledTaskResponses) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *ScheduledTaskResponses) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlSortedBy

`func (o *ScheduledTaskResponses) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *ScheduledTaskResponses) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *ScheduledTaskResponses) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *ScheduledTaskResponses) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *ScheduledTaskResponses) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *ScheduledTaskResponses) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil
### GetHqlType

`func (o *ScheduledTaskResponses) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *ScheduledTaskResponses) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *ScheduledTaskResponses) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetIsHtml

`func (o *ScheduledTaskResponses) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *ScheduledTaskResponses) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *ScheduledTaskResponses) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetRecipients

`func (o *ScheduledTaskResponses) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ScheduledTaskResponses) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ScheduledTaskResponses) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetTitle

`func (o *ScheduledTaskResponses) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ScheduledTaskResponses) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ScheduledTaskResponses) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *ScheduledTaskResponses) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduledTaskResponses) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduledTaskResponses) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *ScheduledTaskResponses) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ScheduledTaskResponses) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ScheduledTaskResponses) SetCron(v string)`

SetCron sets Cron field to given value.


### GetDescription

`func (o *ScheduledTaskResponses) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduledTaskResponses) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduledTaskResponses) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduledTaskResponses) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetail

`func (o *ScheduledTaskResponses) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ScheduledTaskResponses) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ScheduledTaskResponses) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ScheduledTaskResponses) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ScheduledTaskResponses) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ScheduledTaskResponses) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetEnabled

`func (o *ScheduledTaskResponses) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScheduledTaskResponses) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScheduledTaskResponses) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExecutionId

`func (o *ScheduledTaskResponses) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ScheduledTaskResponses) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ScheduledTaskResponses) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ScheduledTaskResponses) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *ScheduledTaskResponses) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *ScheduledTaskResponses) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetHost

`func (o *ScheduledTaskResponses) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ScheduledTaskResponses) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ScheduledTaskResponses) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ScheduledTaskResponses) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *ScheduledTaskResponses) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *ScheduledTaskResponses) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLastCompletionDate

`func (o *ScheduledTaskResponses) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *ScheduledTaskResponses) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *ScheduledTaskResponses) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *ScheduledTaskResponses) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *ScheduledTaskResponses) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *ScheduledTaskResponses) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetLastExecutionDate

`func (o *ScheduledTaskResponses) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *ScheduledTaskResponses) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *ScheduledTaskResponses) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *ScheduledTaskResponses) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *ScheduledTaskResponses) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *ScheduledTaskResponses) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetName

`func (o *ScheduledTaskResponses) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledTaskResponses) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledTaskResponses) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ScheduledTaskResponses) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledTaskResponses) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledTaskResponses) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduledTaskResponses) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ScheduledTaskResponses) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ScheduledTaskResponses) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCompressCsv

`func (o *ScheduledTaskResponses) GetCompressCsv() bool`

GetCompressCsv returns the CompressCsv field if non-nil, zero value otherwise.

### GetCompressCsvOk

`func (o *ScheduledTaskResponses) GetCompressCsvOk() (*bool, bool)`

GetCompressCsvOk returns a tuple with the CompressCsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressCsv

`func (o *ScheduledTaskResponses) SetCompressCsv(v bool)`

SetCompressCsv sets CompressCsv field to given value.

### HasCompressCsv

`func (o *ScheduledTaskResponses) HasCompressCsv() bool`

HasCompressCsv returns a boolean if a field has been set.

### GetConnector

`func (o *ScheduledTaskResponses) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ScheduledTaskResponses) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ScheduledTaskResponses) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetDryRun

`func (o *ScheduledTaskResponses) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ScheduledTaskResponses) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ScheduledTaskResponses) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetEnroll

`func (o *ScheduledTaskResponses) GetEnroll() bool`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *ScheduledTaskResponses) GetEnrollOk() (*bool, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *ScheduledTaskResponses) SetEnroll(v bool)`

SetEnroll sets Enroll field to given value.


### GetModule

`func (o *ScheduledTaskResponses) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ScheduledTaskResponses) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ScheduledTaskResponses) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *ScheduledTaskResponses) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ScheduledTaskResponses) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ScheduledTaskResponses) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRenew

`func (o *ScheduledTaskResponses) GetRenew() bool`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *ScheduledTaskResponses) GetRenewOk() (*bool, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *ScheduledTaskResponses) SetRenew(v bool)`

SetRenew sets Renew field to given value.

### HasRenew

`func (o *ScheduledTaskResponses) HasRenew() bool`

HasRenew returns a boolean if a field has been set.

### GetRevoke

`func (o *ScheduledTaskResponses) GetRevoke() bool`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *ScheduledTaskResponses) GetRevokeOk() (*bool, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *ScheduledTaskResponses) SetRevoke(v bool)`

SetRevoke sets Revoke field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


