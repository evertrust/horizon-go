# LinkReportScheduledTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | **string** | Instead of sending an email with the CSV as an attachment, an email containing a link is sent | 
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

## Methods

### NewLinkReportScheduledTask

`func NewLinkReportScheduledTask(reportType string, retentionPeriod string, from string, hqlType string, isHtml bool, recipients []ReportRecipient, title string, type_ string, cron string, enabled bool, name string, ) *LinkReportScheduledTask`

NewLinkReportScheduledTask instantiates a new LinkReportScheduledTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkReportScheduledTaskWithDefaults

`func NewLinkReportScheduledTaskWithDefaults() *LinkReportScheduledTask`

NewLinkReportScheduledTaskWithDefaults instantiates a new LinkReportScheduledTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportType

`func (o *LinkReportScheduledTask) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *LinkReportScheduledTask) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *LinkReportScheduledTask) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetRetentionPeriod

`func (o *LinkReportScheduledTask) GetRetentionPeriod() string`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *LinkReportScheduledTask) GetRetentionPeriodOk() (*string, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *LinkReportScheduledTask) SetRetentionPeriod(v string)`

SetRetentionPeriod sets RetentionPeriod field to given value.


### GetBody

`func (o *LinkReportScheduledTask) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *LinkReportScheduledTask) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *LinkReportScheduledTask) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *LinkReportScheduledTask) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *LinkReportScheduledTask) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *LinkReportScheduledTask) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetFileName

`func (o *LinkReportScheduledTask) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *LinkReportScheduledTask) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *LinkReportScheduledTask) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *LinkReportScheduledTask) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *LinkReportScheduledTask) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *LinkReportScheduledTask) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFrom

`func (o *LinkReportScheduledTask) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LinkReportScheduledTask) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LinkReportScheduledTask) SetFrom(v string)`

SetFrom sets From field to given value.


### GetHqlFields

`func (o *LinkReportScheduledTask) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *LinkReportScheduledTask) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *LinkReportScheduledTask) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *LinkReportScheduledTask) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *LinkReportScheduledTask) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *LinkReportScheduledTask) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlQuery

`func (o *LinkReportScheduledTask) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *LinkReportScheduledTask) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *LinkReportScheduledTask) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *LinkReportScheduledTask) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *LinkReportScheduledTask) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *LinkReportScheduledTask) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlSortedBy

`func (o *LinkReportScheduledTask) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *LinkReportScheduledTask) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *LinkReportScheduledTask) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *LinkReportScheduledTask) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *LinkReportScheduledTask) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *LinkReportScheduledTask) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil
### GetHqlType

`func (o *LinkReportScheduledTask) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *LinkReportScheduledTask) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *LinkReportScheduledTask) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetIsHtml

`func (o *LinkReportScheduledTask) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *LinkReportScheduledTask) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *LinkReportScheduledTask) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetRecipients

`func (o *LinkReportScheduledTask) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *LinkReportScheduledTask) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *LinkReportScheduledTask) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetTitle

`func (o *LinkReportScheduledTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LinkReportScheduledTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LinkReportScheduledTask) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *LinkReportScheduledTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkReportScheduledTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkReportScheduledTask) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *LinkReportScheduledTask) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *LinkReportScheduledTask) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *LinkReportScheduledTask) SetCron(v string)`

SetCron sets Cron field to given value.


### GetDescription

`func (o *LinkReportScheduledTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinkReportScheduledTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinkReportScheduledTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LinkReportScheduledTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetail

`func (o *LinkReportScheduledTask) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *LinkReportScheduledTask) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *LinkReportScheduledTask) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *LinkReportScheduledTask) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *LinkReportScheduledTask) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *LinkReportScheduledTask) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetEnabled

`func (o *LinkReportScheduledTask) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LinkReportScheduledTask) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LinkReportScheduledTask) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExecutionId

`func (o *LinkReportScheduledTask) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *LinkReportScheduledTask) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *LinkReportScheduledTask) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *LinkReportScheduledTask) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *LinkReportScheduledTask) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *LinkReportScheduledTask) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetHost

`func (o *LinkReportScheduledTask) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LinkReportScheduledTask) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LinkReportScheduledTask) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LinkReportScheduledTask) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *LinkReportScheduledTask) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *LinkReportScheduledTask) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLastCompletionDate

`func (o *LinkReportScheduledTask) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *LinkReportScheduledTask) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *LinkReportScheduledTask) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *LinkReportScheduledTask) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *LinkReportScheduledTask) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *LinkReportScheduledTask) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetLastExecutionDate

`func (o *LinkReportScheduledTask) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *LinkReportScheduledTask) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *LinkReportScheduledTask) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *LinkReportScheduledTask) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *LinkReportScheduledTask) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *LinkReportScheduledTask) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetName

`func (o *LinkReportScheduledTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkReportScheduledTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkReportScheduledTask) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *LinkReportScheduledTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkReportScheduledTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkReportScheduledTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LinkReportScheduledTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *LinkReportScheduledTask) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *LinkReportScheduledTask) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


