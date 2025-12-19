# ReportScheduledTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ReportType** | **string** |  | 
**RetentionPeriod** | **string** | Indicates the duration during which the report can be downloaded | 
**Body** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**From** | **string** |  | 
**HqlFields** | Pointer to **[]string** |  | [optional] 
**HqlQuery** | Pointer to **NullableString** |  | [optional] 
**HqlSortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**HqlType** | **string** |  | 
**IsHtml** | **bool** |  | 
**Name** | **string** |  | 
**Recipients** | [**[]ReportRecipient**](ReportRecipient.md) |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Cron** | **string** |  | 
**Detail** | Pointer to **NullableString** |  | [optional] 
**Enabled** | **bool** |  | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**LastCompletionDate** | Pointer to **NullableInt64** |  | [optional] 
**LastExecutionDate** | Pointer to **NullableInt64** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**CompressCsv** | Pointer to **bool** | Should the report be compressed using GZ | [optional] 

## Methods

### NewReportScheduledTaskResponse

`func NewReportScheduledTaskResponse(id string, reportType string, retentionPeriod string, from string, hqlType string, isHtml bool, name string, recipients []ReportRecipient, title string, type_ string, cron string, enabled bool, ) *ReportScheduledTaskResponse`

NewReportScheduledTaskResponse instantiates a new ReportScheduledTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportScheduledTaskResponseWithDefaults

`func NewReportScheduledTaskResponseWithDefaults() *ReportScheduledTaskResponse`

NewReportScheduledTaskResponseWithDefaults instantiates a new ReportScheduledTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportScheduledTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportScheduledTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportScheduledTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetReportType

`func (o *ReportScheduledTaskResponse) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportScheduledTaskResponse) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportScheduledTaskResponse) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetRetentionPeriod

`func (o *ReportScheduledTaskResponse) GetRetentionPeriod() string`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *ReportScheduledTaskResponse) GetRetentionPeriodOk() (*string, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *ReportScheduledTaskResponse) SetRetentionPeriod(v string)`

SetRetentionPeriod sets RetentionPeriod field to given value.


### GetBody

`func (o *ReportScheduledTaskResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReportScheduledTaskResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReportScheduledTaskResponse) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ReportScheduledTaskResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ReportScheduledTaskResponse) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ReportScheduledTaskResponse) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetDescription

`func (o *ReportScheduledTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportScheduledTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportScheduledTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportScheduledTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ReportScheduledTaskResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ReportScheduledTaskResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFileName

`func (o *ReportScheduledTaskResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ReportScheduledTaskResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ReportScheduledTaskResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ReportScheduledTaskResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ReportScheduledTaskResponse) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ReportScheduledTaskResponse) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFrom

`func (o *ReportScheduledTaskResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReportScheduledTaskResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReportScheduledTaskResponse) SetFrom(v string)`

SetFrom sets From field to given value.


### GetHqlFields

`func (o *ReportScheduledTaskResponse) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *ReportScheduledTaskResponse) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *ReportScheduledTaskResponse) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *ReportScheduledTaskResponse) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *ReportScheduledTaskResponse) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *ReportScheduledTaskResponse) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlQuery

`func (o *ReportScheduledTaskResponse) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *ReportScheduledTaskResponse) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *ReportScheduledTaskResponse) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *ReportScheduledTaskResponse) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *ReportScheduledTaskResponse) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *ReportScheduledTaskResponse) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlSortedBy

`func (o *ReportScheduledTaskResponse) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *ReportScheduledTaskResponse) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *ReportScheduledTaskResponse) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *ReportScheduledTaskResponse) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *ReportScheduledTaskResponse) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *ReportScheduledTaskResponse) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil
### GetHqlType

`func (o *ReportScheduledTaskResponse) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *ReportScheduledTaskResponse) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *ReportScheduledTaskResponse) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetIsHtml

`func (o *ReportScheduledTaskResponse) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *ReportScheduledTaskResponse) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *ReportScheduledTaskResponse) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetName

`func (o *ReportScheduledTaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportScheduledTaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportScheduledTaskResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRecipients

`func (o *ReportScheduledTaskResponse) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ReportScheduledTaskResponse) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ReportScheduledTaskResponse) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetTitle

`func (o *ReportScheduledTaskResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportScheduledTaskResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportScheduledTaskResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *ReportScheduledTaskResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportScheduledTaskResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportScheduledTaskResponse) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *ReportScheduledTaskResponse) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ReportScheduledTaskResponse) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ReportScheduledTaskResponse) SetCron(v string)`

SetCron sets Cron field to given value.


### GetDetail

`func (o *ReportScheduledTaskResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ReportScheduledTaskResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ReportScheduledTaskResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ReportScheduledTaskResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ReportScheduledTaskResponse) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ReportScheduledTaskResponse) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetEnabled

`func (o *ReportScheduledTaskResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReportScheduledTaskResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReportScheduledTaskResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExecutionId

`func (o *ReportScheduledTaskResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ReportScheduledTaskResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ReportScheduledTaskResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ReportScheduledTaskResponse) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *ReportScheduledTaskResponse) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *ReportScheduledTaskResponse) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetHost

`func (o *ReportScheduledTaskResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ReportScheduledTaskResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ReportScheduledTaskResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ReportScheduledTaskResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *ReportScheduledTaskResponse) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *ReportScheduledTaskResponse) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLastCompletionDate

`func (o *ReportScheduledTaskResponse) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *ReportScheduledTaskResponse) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *ReportScheduledTaskResponse) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *ReportScheduledTaskResponse) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *ReportScheduledTaskResponse) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *ReportScheduledTaskResponse) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetLastExecutionDate

`func (o *ReportScheduledTaskResponse) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *ReportScheduledTaskResponse) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *ReportScheduledTaskResponse) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *ReportScheduledTaskResponse) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *ReportScheduledTaskResponse) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *ReportScheduledTaskResponse) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetStatus

`func (o *ReportScheduledTaskResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportScheduledTaskResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportScheduledTaskResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportScheduledTaskResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ReportScheduledTaskResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReportScheduledTaskResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCompressCsv

`func (o *ReportScheduledTaskResponse) GetCompressCsv() bool`

GetCompressCsv returns the CompressCsv field if non-nil, zero value otherwise.

### GetCompressCsvOk

`func (o *ReportScheduledTaskResponse) GetCompressCsvOk() (*bool, bool)`

GetCompressCsvOk returns a tuple with the CompressCsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressCsv

`func (o *ReportScheduledTaskResponse) SetCompressCsv(v bool)`

SetCompressCsv sets CompressCsv field to given value.

### HasCompressCsv

`func (o *ReportScheduledTaskResponse) HasCompressCsv() bool`

HasCompressCsv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


