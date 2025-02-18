# SchedulerTaskAddRequest

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
**DryRun** | **bool** |  | 
**Module** | **string** |  | 
**Profile** | **string** |  | 
**Connector** | **string** |  | 
**Enroll** | **bool** |  | 
**Revoke** | **bool** |  | 
**Renew** | **bool** |  | 
**Results** | Pointer to [**NullableThirdPartyConnectorSynchronizationResult**](ThirdPartyConnectorSynchronizationResult.md) |  | [optional] 

## Methods

### NewSchedulerTaskAddRequest

`func NewSchedulerTaskAddRequest(type_ string, cron string, enabled bool, name string, recipients []ReportRecipient, from string, title string, isHtml bool, hqlType string, dryRun bool, module string, profile string, connector string, enroll bool, revoke bool, renew bool, ) *SchedulerTaskAddRequest`

NewSchedulerTaskAddRequest instantiates a new SchedulerTaskAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTaskAddRequestWithDefaults

`func NewSchedulerTaskAddRequestWithDefaults() *SchedulerTaskAddRequest`

NewSchedulerTaskAddRequestWithDefaults instantiates a new SchedulerTaskAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchedulerTaskAddRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchedulerTaskAddRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchedulerTaskAddRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *SchedulerTaskAddRequest) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *SchedulerTaskAddRequest) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *SchedulerTaskAddRequest) SetCron(v string)`

SetCron sets Cron field to given value.


### GetHost

`func (o *SchedulerTaskAddRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SchedulerTaskAddRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SchedulerTaskAddRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SchedulerTaskAddRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SchedulerTaskAddRequest) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SchedulerTaskAddRequest) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetStatus

`func (o *SchedulerTaskAddRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchedulerTaskAddRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchedulerTaskAddRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchedulerTaskAddRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SchedulerTaskAddRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SchedulerTaskAddRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastExecutionDate

`func (o *SchedulerTaskAddRequest) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *SchedulerTaskAddRequest) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *SchedulerTaskAddRequest) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *SchedulerTaskAddRequest) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *SchedulerTaskAddRequest) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *SchedulerTaskAddRequest) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetLastCompletionDate

`func (o *SchedulerTaskAddRequest) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *SchedulerTaskAddRequest) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *SchedulerTaskAddRequest) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *SchedulerTaskAddRequest) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *SchedulerTaskAddRequest) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *SchedulerTaskAddRequest) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetDetail

`func (o *SchedulerTaskAddRequest) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SchedulerTaskAddRequest) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SchedulerTaskAddRequest) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SchedulerTaskAddRequest) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *SchedulerTaskAddRequest) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *SchedulerTaskAddRequest) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetExecutionId

`func (o *SchedulerTaskAddRequest) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *SchedulerTaskAddRequest) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *SchedulerTaskAddRequest) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *SchedulerTaskAddRequest) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *SchedulerTaskAddRequest) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *SchedulerTaskAddRequest) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetEnabled

`func (o *SchedulerTaskAddRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SchedulerTaskAddRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SchedulerTaskAddRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *SchedulerTaskAddRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulerTaskAddRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulerTaskAddRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFileName

`func (o *SchedulerTaskAddRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SchedulerTaskAddRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SchedulerTaskAddRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SchedulerTaskAddRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SchedulerTaskAddRequest) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SchedulerTaskAddRequest) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetRecipients

`func (o *SchedulerTaskAddRequest) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *SchedulerTaskAddRequest) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *SchedulerTaskAddRequest) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetFrom

`func (o *SchedulerTaskAddRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SchedulerTaskAddRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SchedulerTaskAddRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTitle

`func (o *SchedulerTaskAddRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SchedulerTaskAddRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SchedulerTaskAddRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *SchedulerTaskAddRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SchedulerTaskAddRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SchedulerTaskAddRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SchedulerTaskAddRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *SchedulerTaskAddRequest) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *SchedulerTaskAddRequest) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetIsHtml

`func (o *SchedulerTaskAddRequest) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *SchedulerTaskAddRequest) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *SchedulerTaskAddRequest) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetHqlType

`func (o *SchedulerTaskAddRequest) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *SchedulerTaskAddRequest) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *SchedulerTaskAddRequest) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetHqlQuery

`func (o *SchedulerTaskAddRequest) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *SchedulerTaskAddRequest) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *SchedulerTaskAddRequest) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *SchedulerTaskAddRequest) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *SchedulerTaskAddRequest) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *SchedulerTaskAddRequest) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlFields

`func (o *SchedulerTaskAddRequest) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *SchedulerTaskAddRequest) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *SchedulerTaskAddRequest) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *SchedulerTaskAddRequest) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *SchedulerTaskAddRequest) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *SchedulerTaskAddRequest) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlSortedBy

`func (o *SchedulerTaskAddRequest) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *SchedulerTaskAddRequest) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *SchedulerTaskAddRequest) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *SchedulerTaskAddRequest) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *SchedulerTaskAddRequest) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *SchedulerTaskAddRequest) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil
### GetDescription

`func (o *SchedulerTaskAddRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchedulerTaskAddRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchedulerTaskAddRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchedulerTaskAddRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SchedulerTaskAddRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SchedulerTaskAddRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDryRun

`func (o *SchedulerTaskAddRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SchedulerTaskAddRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SchedulerTaskAddRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetModule

`func (o *SchedulerTaskAddRequest) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *SchedulerTaskAddRequest) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *SchedulerTaskAddRequest) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *SchedulerTaskAddRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SchedulerTaskAddRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SchedulerTaskAddRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetConnector

`func (o *SchedulerTaskAddRequest) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *SchedulerTaskAddRequest) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *SchedulerTaskAddRequest) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetEnroll

`func (o *SchedulerTaskAddRequest) GetEnroll() bool`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *SchedulerTaskAddRequest) GetEnrollOk() (*bool, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *SchedulerTaskAddRequest) SetEnroll(v bool)`

SetEnroll sets Enroll field to given value.


### GetRevoke

`func (o *SchedulerTaskAddRequest) GetRevoke() bool`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *SchedulerTaskAddRequest) GetRevokeOk() (*bool, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *SchedulerTaskAddRequest) SetRevoke(v bool)`

SetRevoke sets Revoke field to given value.


### GetRenew

`func (o *SchedulerTaskAddRequest) GetRenew() bool`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *SchedulerTaskAddRequest) GetRenewOk() (*bool, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *SchedulerTaskAddRequest) SetRenew(v bool)`

SetRenew sets Renew field to given value.


### GetResults

`func (o *SchedulerTaskAddRequest) GetResults() ThirdPartyConnectorSynchronizationResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SchedulerTaskAddRequest) GetResultsOk() (*ThirdPartyConnectorSynchronizationResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SchedulerTaskAddRequest) SetResults(v ThirdPartyConnectorSynchronizationResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *SchedulerTaskAddRequest) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *SchedulerTaskAddRequest) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *SchedulerTaskAddRequest) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


