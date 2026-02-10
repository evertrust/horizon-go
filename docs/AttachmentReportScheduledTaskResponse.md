# AttachmentReportScheduledTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**CompressCsv** | Pointer to **bool** | Should the report be compressed using GZ | [optional] 
**ReportType** | **string** |  | 
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

### NewAttachmentReportScheduledTaskResponse

`func NewAttachmentReportScheduledTaskResponse(id string, reportType string, type_ string, name string, recipients []ReportRecipient, from string, title string, isHtml bool, hqlType string, cron string, enabled bool, ) *AttachmentReportScheduledTaskResponse`

NewAttachmentReportScheduledTaskResponse instantiates a new AttachmentReportScheduledTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentReportScheduledTaskResponseWithDefaults

`func NewAttachmentReportScheduledTaskResponseWithDefaults() *AttachmentReportScheduledTaskResponse`

NewAttachmentReportScheduledTaskResponseWithDefaults instantiates a new AttachmentReportScheduledTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentReportScheduledTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentReportScheduledTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentReportScheduledTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCompressCsv

`func (o *AttachmentReportScheduledTaskResponse) GetCompressCsv() bool`

GetCompressCsv returns the CompressCsv field if non-nil, zero value otherwise.

### GetCompressCsvOk

`func (o *AttachmentReportScheduledTaskResponse) GetCompressCsvOk() (*bool, bool)`

GetCompressCsvOk returns a tuple with the CompressCsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressCsv

`func (o *AttachmentReportScheduledTaskResponse) SetCompressCsv(v bool)`

SetCompressCsv sets CompressCsv field to given value.

### HasCompressCsv

`func (o *AttachmentReportScheduledTaskResponse) HasCompressCsv() bool`

HasCompressCsv returns a boolean if a field has been set.

### GetReportType

`func (o *AttachmentReportScheduledTaskResponse) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *AttachmentReportScheduledTaskResponse) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *AttachmentReportScheduledTaskResponse) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetType

`func (o *AttachmentReportScheduledTaskResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentReportScheduledTaskResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentReportScheduledTaskResponse) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AttachmentReportScheduledTaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachmentReportScheduledTaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachmentReportScheduledTaskResponse) SetName(v string)`

SetName sets Name field to given value.


### GetFileName

`func (o *AttachmentReportScheduledTaskResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AttachmentReportScheduledTaskResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AttachmentReportScheduledTaskResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AttachmentReportScheduledTaskResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *AttachmentReportScheduledTaskResponse) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *AttachmentReportScheduledTaskResponse) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetRecipients

`func (o *AttachmentReportScheduledTaskResponse) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AttachmentReportScheduledTaskResponse) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AttachmentReportScheduledTaskResponse) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetFrom

`func (o *AttachmentReportScheduledTaskResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AttachmentReportScheduledTaskResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AttachmentReportScheduledTaskResponse) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTitle

`func (o *AttachmentReportScheduledTaskResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttachmentReportScheduledTaskResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttachmentReportScheduledTaskResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *AttachmentReportScheduledTaskResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AttachmentReportScheduledTaskResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AttachmentReportScheduledTaskResponse) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AttachmentReportScheduledTaskResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *AttachmentReportScheduledTaskResponse) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *AttachmentReportScheduledTaskResponse) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetIsHtml

`func (o *AttachmentReportScheduledTaskResponse) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *AttachmentReportScheduledTaskResponse) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *AttachmentReportScheduledTaskResponse) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetHqlType

`func (o *AttachmentReportScheduledTaskResponse) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *AttachmentReportScheduledTaskResponse) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *AttachmentReportScheduledTaskResponse) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetHqlQuery

`func (o *AttachmentReportScheduledTaskResponse) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *AttachmentReportScheduledTaskResponse) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *AttachmentReportScheduledTaskResponse) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *AttachmentReportScheduledTaskResponse) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *AttachmentReportScheduledTaskResponse) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *AttachmentReportScheduledTaskResponse) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlFields

`func (o *AttachmentReportScheduledTaskResponse) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *AttachmentReportScheduledTaskResponse) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *AttachmentReportScheduledTaskResponse) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *AttachmentReportScheduledTaskResponse) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *AttachmentReportScheduledTaskResponse) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *AttachmentReportScheduledTaskResponse) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlSortedBy

`func (o *AttachmentReportScheduledTaskResponse) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *AttachmentReportScheduledTaskResponse) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *AttachmentReportScheduledTaskResponse) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *AttachmentReportScheduledTaskResponse) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *AttachmentReportScheduledTaskResponse) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *AttachmentReportScheduledTaskResponse) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil
### GetDescription

`func (o *AttachmentReportScheduledTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttachmentReportScheduledTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttachmentReportScheduledTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttachmentReportScheduledTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AttachmentReportScheduledTaskResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AttachmentReportScheduledTaskResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCron

`func (o *AttachmentReportScheduledTaskResponse) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *AttachmentReportScheduledTaskResponse) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *AttachmentReportScheduledTaskResponse) SetCron(v string)`

SetCron sets Cron field to given value.


### GetHost

`func (o *AttachmentReportScheduledTaskResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AttachmentReportScheduledTaskResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AttachmentReportScheduledTaskResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AttachmentReportScheduledTaskResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *AttachmentReportScheduledTaskResponse) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *AttachmentReportScheduledTaskResponse) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetStatus

`func (o *AttachmentReportScheduledTaskResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttachmentReportScheduledTaskResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttachmentReportScheduledTaskResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AttachmentReportScheduledTaskResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AttachmentReportScheduledTaskResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AttachmentReportScheduledTaskResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastExecutionDate

`func (o *AttachmentReportScheduledTaskResponse) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *AttachmentReportScheduledTaskResponse) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *AttachmentReportScheduledTaskResponse) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *AttachmentReportScheduledTaskResponse) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *AttachmentReportScheduledTaskResponse) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *AttachmentReportScheduledTaskResponse) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetLastCompletionDate

`func (o *AttachmentReportScheduledTaskResponse) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *AttachmentReportScheduledTaskResponse) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *AttachmentReportScheduledTaskResponse) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *AttachmentReportScheduledTaskResponse) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *AttachmentReportScheduledTaskResponse) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *AttachmentReportScheduledTaskResponse) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetDetail

`func (o *AttachmentReportScheduledTaskResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AttachmentReportScheduledTaskResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AttachmentReportScheduledTaskResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AttachmentReportScheduledTaskResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *AttachmentReportScheduledTaskResponse) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *AttachmentReportScheduledTaskResponse) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetExecutionId

`func (o *AttachmentReportScheduledTaskResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *AttachmentReportScheduledTaskResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *AttachmentReportScheduledTaskResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *AttachmentReportScheduledTaskResponse) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *AttachmentReportScheduledTaskResponse) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *AttachmentReportScheduledTaskResponse) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetEnabled

`func (o *AttachmentReportScheduledTaskResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AttachmentReportScheduledTaskResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AttachmentReportScheduledTaskResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


