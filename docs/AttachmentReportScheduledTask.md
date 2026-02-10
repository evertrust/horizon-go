# AttachmentReportScheduledTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressCsv** | Pointer to **bool** | Should the report be compressed using GZ | [optional] 
**ReportType** | **string** | Send an email with the CSV as an attachment | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Recipients** | [**[]ReportRecipient**](ReportRecipient.md) |  | 
**From** | **string** |  | 
**Title** | **string** |  | 
**Body** | Pointer to **NullableString** |  | [optional] 
**IsHtml** | **bool** |  | 
**HqlType** | **string** |  | 
**HqlQuery** | Pointer to **NullableString** |  | [optional] 
**HqlFields** | Pointer to **[]string** |  | [optional] 
**HqlSortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Cron** | **string** |  | 
**Host** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**LastExecutionDate** | Pointer to **NullableInt64** |  | [optional] 
**LastCompletionDate** | Pointer to **NullableInt64** |  | [optional] 
**Detail** | Pointer to **NullableString** |  | [optional] 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | **bool** |  | 

## Methods

### NewAttachmentReportScheduledTask

`func NewAttachmentReportScheduledTask(reportType string, type_ string, name string, recipients []ReportRecipient, from string, title string, isHtml bool, hqlType string, cron string, enabled bool, ) *AttachmentReportScheduledTask`

NewAttachmentReportScheduledTask instantiates a new AttachmentReportScheduledTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentReportScheduledTaskWithDefaults

`func NewAttachmentReportScheduledTaskWithDefaults() *AttachmentReportScheduledTask`

NewAttachmentReportScheduledTaskWithDefaults instantiates a new AttachmentReportScheduledTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressCsv

`func (o *AttachmentReportScheduledTask) GetCompressCsv() bool`

GetCompressCsv returns the CompressCsv field if non-nil, zero value otherwise.

### GetCompressCsvOk

`func (o *AttachmentReportScheduledTask) GetCompressCsvOk() (*bool, bool)`

GetCompressCsvOk returns a tuple with the CompressCsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressCsv

`func (o *AttachmentReportScheduledTask) SetCompressCsv(v bool)`

SetCompressCsv sets CompressCsv field to given value.

### HasCompressCsv

`func (o *AttachmentReportScheduledTask) HasCompressCsv() bool`

HasCompressCsv returns a boolean if a field has been set.

### GetReportType

`func (o *AttachmentReportScheduledTask) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *AttachmentReportScheduledTask) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *AttachmentReportScheduledTask) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetType

`func (o *AttachmentReportScheduledTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentReportScheduledTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentReportScheduledTask) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AttachmentReportScheduledTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachmentReportScheduledTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachmentReportScheduledTask) SetName(v string)`

SetName sets Name field to given value.


### GetFileName

`func (o *AttachmentReportScheduledTask) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AttachmentReportScheduledTask) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AttachmentReportScheduledTask) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AttachmentReportScheduledTask) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *AttachmentReportScheduledTask) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *AttachmentReportScheduledTask) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetRecipients

`func (o *AttachmentReportScheduledTask) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AttachmentReportScheduledTask) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AttachmentReportScheduledTask) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetFrom

`func (o *AttachmentReportScheduledTask) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AttachmentReportScheduledTask) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AttachmentReportScheduledTask) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTitle

`func (o *AttachmentReportScheduledTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttachmentReportScheduledTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttachmentReportScheduledTask) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *AttachmentReportScheduledTask) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AttachmentReportScheduledTask) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AttachmentReportScheduledTask) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AttachmentReportScheduledTask) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *AttachmentReportScheduledTask) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *AttachmentReportScheduledTask) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetIsHtml

`func (o *AttachmentReportScheduledTask) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *AttachmentReportScheduledTask) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *AttachmentReportScheduledTask) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetHqlType

`func (o *AttachmentReportScheduledTask) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *AttachmentReportScheduledTask) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *AttachmentReportScheduledTask) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetHqlQuery

`func (o *AttachmentReportScheduledTask) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *AttachmentReportScheduledTask) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *AttachmentReportScheduledTask) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *AttachmentReportScheduledTask) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *AttachmentReportScheduledTask) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *AttachmentReportScheduledTask) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlFields

`func (o *AttachmentReportScheduledTask) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *AttachmentReportScheduledTask) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *AttachmentReportScheduledTask) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *AttachmentReportScheduledTask) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *AttachmentReportScheduledTask) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *AttachmentReportScheduledTask) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlSortedBy

`func (o *AttachmentReportScheduledTask) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *AttachmentReportScheduledTask) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *AttachmentReportScheduledTask) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *AttachmentReportScheduledTask) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *AttachmentReportScheduledTask) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *AttachmentReportScheduledTask) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil
### GetDescription

`func (o *AttachmentReportScheduledTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttachmentReportScheduledTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttachmentReportScheduledTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttachmentReportScheduledTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AttachmentReportScheduledTask) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AttachmentReportScheduledTask) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCron

`func (o *AttachmentReportScheduledTask) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *AttachmentReportScheduledTask) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *AttachmentReportScheduledTask) SetCron(v string)`

SetCron sets Cron field to given value.


### GetHost

`func (o *AttachmentReportScheduledTask) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AttachmentReportScheduledTask) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AttachmentReportScheduledTask) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AttachmentReportScheduledTask) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *AttachmentReportScheduledTask) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *AttachmentReportScheduledTask) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetStatus

`func (o *AttachmentReportScheduledTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttachmentReportScheduledTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttachmentReportScheduledTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AttachmentReportScheduledTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AttachmentReportScheduledTask) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AttachmentReportScheduledTask) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastExecutionDate

`func (o *AttachmentReportScheduledTask) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *AttachmentReportScheduledTask) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *AttachmentReportScheduledTask) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *AttachmentReportScheduledTask) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *AttachmentReportScheduledTask) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *AttachmentReportScheduledTask) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetLastCompletionDate

`func (o *AttachmentReportScheduledTask) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *AttachmentReportScheduledTask) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *AttachmentReportScheduledTask) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *AttachmentReportScheduledTask) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *AttachmentReportScheduledTask) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *AttachmentReportScheduledTask) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetDetail

`func (o *AttachmentReportScheduledTask) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AttachmentReportScheduledTask) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AttachmentReportScheduledTask) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AttachmentReportScheduledTask) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *AttachmentReportScheduledTask) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *AttachmentReportScheduledTask) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetExecutionId

`func (o *AttachmentReportScheduledTask) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *AttachmentReportScheduledTask) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *AttachmentReportScheduledTask) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *AttachmentReportScheduledTask) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *AttachmentReportScheduledTask) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *AttachmentReportScheduledTask) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetEnabled

`func (o *AttachmentReportScheduledTask) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AttachmentReportScheduledTask) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AttachmentReportScheduledTask) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


