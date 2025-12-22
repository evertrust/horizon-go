# LinkReportScheduledTaskResponse

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

## Methods

### NewLinkReportScheduledTaskResponse

`func NewLinkReportScheduledTaskResponse(id string, reportType string, retentionPeriod string, from string, hqlType string, isHtml bool, name string, recipients []ReportRecipient, title string, type_ string, cron string, enabled bool, ) *LinkReportScheduledTaskResponse`

NewLinkReportScheduledTaskResponse instantiates a new LinkReportScheduledTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkReportScheduledTaskResponseWithDefaults

`func NewLinkReportScheduledTaskResponseWithDefaults() *LinkReportScheduledTaskResponse`

NewLinkReportScheduledTaskResponseWithDefaults instantiates a new LinkReportScheduledTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkReportScheduledTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkReportScheduledTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkReportScheduledTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetReportType

`func (o *LinkReportScheduledTaskResponse) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *LinkReportScheduledTaskResponse) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *LinkReportScheduledTaskResponse) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetRetentionPeriod

`func (o *LinkReportScheduledTaskResponse) GetRetentionPeriod() string`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *LinkReportScheduledTaskResponse) GetRetentionPeriodOk() (*string, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *LinkReportScheduledTaskResponse) SetRetentionPeriod(v string)`

SetRetentionPeriod sets RetentionPeriod field to given value.


### GetBody

`func (o *LinkReportScheduledTaskResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *LinkReportScheduledTaskResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *LinkReportScheduledTaskResponse) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *LinkReportScheduledTaskResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *LinkReportScheduledTaskResponse) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *LinkReportScheduledTaskResponse) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetDescription

`func (o *LinkReportScheduledTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinkReportScheduledTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinkReportScheduledTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LinkReportScheduledTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LinkReportScheduledTaskResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LinkReportScheduledTaskResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFileName

`func (o *LinkReportScheduledTaskResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *LinkReportScheduledTaskResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *LinkReportScheduledTaskResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *LinkReportScheduledTaskResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *LinkReportScheduledTaskResponse) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *LinkReportScheduledTaskResponse) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFrom

`func (o *LinkReportScheduledTaskResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LinkReportScheduledTaskResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LinkReportScheduledTaskResponse) SetFrom(v string)`

SetFrom sets From field to given value.


### GetHqlFields

`func (o *LinkReportScheduledTaskResponse) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *LinkReportScheduledTaskResponse) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *LinkReportScheduledTaskResponse) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *LinkReportScheduledTaskResponse) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *LinkReportScheduledTaskResponse) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *LinkReportScheduledTaskResponse) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlQuery

`func (o *LinkReportScheduledTaskResponse) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *LinkReportScheduledTaskResponse) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *LinkReportScheduledTaskResponse) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *LinkReportScheduledTaskResponse) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *LinkReportScheduledTaskResponse) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *LinkReportScheduledTaskResponse) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlSortedBy

`func (o *LinkReportScheduledTaskResponse) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *LinkReportScheduledTaskResponse) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *LinkReportScheduledTaskResponse) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *LinkReportScheduledTaskResponse) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *LinkReportScheduledTaskResponse) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *LinkReportScheduledTaskResponse) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil
### GetHqlType

`func (o *LinkReportScheduledTaskResponse) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *LinkReportScheduledTaskResponse) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *LinkReportScheduledTaskResponse) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetIsHtml

`func (o *LinkReportScheduledTaskResponse) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *LinkReportScheduledTaskResponse) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *LinkReportScheduledTaskResponse) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetName

`func (o *LinkReportScheduledTaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkReportScheduledTaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkReportScheduledTaskResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRecipients

`func (o *LinkReportScheduledTaskResponse) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *LinkReportScheduledTaskResponse) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *LinkReportScheduledTaskResponse) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetTitle

`func (o *LinkReportScheduledTaskResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LinkReportScheduledTaskResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LinkReportScheduledTaskResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *LinkReportScheduledTaskResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkReportScheduledTaskResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkReportScheduledTaskResponse) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *LinkReportScheduledTaskResponse) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *LinkReportScheduledTaskResponse) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *LinkReportScheduledTaskResponse) SetCron(v string)`

SetCron sets Cron field to given value.


### GetDetail

`func (o *LinkReportScheduledTaskResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *LinkReportScheduledTaskResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *LinkReportScheduledTaskResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *LinkReportScheduledTaskResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *LinkReportScheduledTaskResponse) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *LinkReportScheduledTaskResponse) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetEnabled

`func (o *LinkReportScheduledTaskResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LinkReportScheduledTaskResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LinkReportScheduledTaskResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExecutionId

`func (o *LinkReportScheduledTaskResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *LinkReportScheduledTaskResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *LinkReportScheduledTaskResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *LinkReportScheduledTaskResponse) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *LinkReportScheduledTaskResponse) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *LinkReportScheduledTaskResponse) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetHost

`func (o *LinkReportScheduledTaskResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LinkReportScheduledTaskResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LinkReportScheduledTaskResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LinkReportScheduledTaskResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *LinkReportScheduledTaskResponse) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *LinkReportScheduledTaskResponse) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLastCompletionDate

`func (o *LinkReportScheduledTaskResponse) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *LinkReportScheduledTaskResponse) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *LinkReportScheduledTaskResponse) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *LinkReportScheduledTaskResponse) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *LinkReportScheduledTaskResponse) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *LinkReportScheduledTaskResponse) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetLastExecutionDate

`func (o *LinkReportScheduledTaskResponse) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *LinkReportScheduledTaskResponse) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *LinkReportScheduledTaskResponse) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *LinkReportScheduledTaskResponse) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *LinkReportScheduledTaskResponse) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *LinkReportScheduledTaskResponse) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetStatus

`func (o *LinkReportScheduledTaskResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkReportScheduledTaskResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkReportScheduledTaskResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LinkReportScheduledTaskResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *LinkReportScheduledTaskResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *LinkReportScheduledTaskResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


