# ReportScheduledTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**HqlType** | **string** |  | 
**HqlQuery** | Pointer to **NullableString** |  | [optional] 
**HqlFields** | Pointer to **[]string** |  | [optional] 
**HqlSortedBy** | Pointer to [**[]SortElement**](SortElement.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReportScheduledTask

`func NewReportScheduledTask(type_ string, cron string, enabled bool, name string, recipients []ReportRecipient, from string, title string, isHtml bool, hqlType string, ) *ReportScheduledTask`

NewReportScheduledTask instantiates a new ReportScheduledTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportScheduledTaskWithDefaults

`func NewReportScheduledTaskWithDefaults() *ReportScheduledTask`

NewReportScheduledTaskWithDefaults instantiates a new ReportScheduledTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReportScheduledTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportScheduledTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportScheduledTask) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *ReportScheduledTask) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ReportScheduledTask) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ReportScheduledTask) SetCron(v string)`

SetCron sets Cron field to given value.


### GetHost

`func (o *ReportScheduledTask) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ReportScheduledTask) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ReportScheduledTask) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ReportScheduledTask) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *ReportScheduledTask) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *ReportScheduledTask) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetStatus

`func (o *ReportScheduledTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportScheduledTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportScheduledTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportScheduledTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ReportScheduledTask) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReportScheduledTask) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastExecutionDate

`func (o *ReportScheduledTask) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *ReportScheduledTask) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *ReportScheduledTask) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *ReportScheduledTask) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *ReportScheduledTask) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *ReportScheduledTask) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetLastCompletionDate

`func (o *ReportScheduledTask) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *ReportScheduledTask) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *ReportScheduledTask) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *ReportScheduledTask) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *ReportScheduledTask) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *ReportScheduledTask) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetDetail

`func (o *ReportScheduledTask) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ReportScheduledTask) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ReportScheduledTask) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ReportScheduledTask) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ReportScheduledTask) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ReportScheduledTask) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetExecutionId

`func (o *ReportScheduledTask) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ReportScheduledTask) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ReportScheduledTask) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ReportScheduledTask) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *ReportScheduledTask) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *ReportScheduledTask) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetEnabled

`func (o *ReportScheduledTask) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReportScheduledTask) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReportScheduledTask) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *ReportScheduledTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportScheduledTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportScheduledTask) SetName(v string)`

SetName sets Name field to given value.


### GetFileName

`func (o *ReportScheduledTask) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ReportScheduledTask) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ReportScheduledTask) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ReportScheduledTask) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ReportScheduledTask) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ReportScheduledTask) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetRecipients

`func (o *ReportScheduledTask) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ReportScheduledTask) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ReportScheduledTask) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetFrom

`func (o *ReportScheduledTask) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReportScheduledTask) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReportScheduledTask) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTitle

`func (o *ReportScheduledTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportScheduledTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportScheduledTask) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *ReportScheduledTask) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReportScheduledTask) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReportScheduledTask) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ReportScheduledTask) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ReportScheduledTask) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ReportScheduledTask) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetIsHtml

`func (o *ReportScheduledTask) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *ReportScheduledTask) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *ReportScheduledTask) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetHqlType

`func (o *ReportScheduledTask) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *ReportScheduledTask) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *ReportScheduledTask) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetHqlQuery

`func (o *ReportScheduledTask) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *ReportScheduledTask) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *ReportScheduledTask) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *ReportScheduledTask) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *ReportScheduledTask) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *ReportScheduledTask) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlFields

`func (o *ReportScheduledTask) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *ReportScheduledTask) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *ReportScheduledTask) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *ReportScheduledTask) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *ReportScheduledTask) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *ReportScheduledTask) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlSortedBy

`func (o *ReportScheduledTask) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *ReportScheduledTask) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *ReportScheduledTask) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *ReportScheduledTask) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *ReportScheduledTask) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *ReportScheduledTask) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil
### GetDescription

`func (o *ReportScheduledTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportScheduledTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportScheduledTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportScheduledTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ReportScheduledTask) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ReportScheduledTask) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


