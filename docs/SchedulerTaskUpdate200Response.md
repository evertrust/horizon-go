# SchedulerTaskUpdate200Response

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
**DryRun** | **bool** |  | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | **bool** |  | 
**Module** | **string** |  | 
**Profile** | **string** |  | 
**Connector** | **string** |  | 
**Enroll** | **bool** |  | 
**Revoke** | **bool** |  | 
**Renew** | **bool** |  | 
**Results** | Pointer to [**NullableThirdPartyConnectorSynchronizationResult**](ThirdPartyConnectorSynchronizationResult.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
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

## Methods

### NewSchedulerTaskUpdate200Response

`func NewSchedulerTaskUpdate200Response(id string, type_ string, cron string, dryRun bool, enabled bool, module string, profile string, connector string, enroll bool, revoke bool, renew bool, name string, recipients []ReportRecipient, from string, title string, isHtml bool, hqlType string, ) *SchedulerTaskUpdate200Response`

NewSchedulerTaskUpdate200Response instantiates a new SchedulerTaskUpdate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTaskUpdate200ResponseWithDefaults

`func NewSchedulerTaskUpdate200ResponseWithDefaults() *SchedulerTaskUpdate200Response`

NewSchedulerTaskUpdate200ResponseWithDefaults instantiates a new SchedulerTaskUpdate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchedulerTaskUpdate200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchedulerTaskUpdate200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchedulerTaskUpdate200Response) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SchedulerTaskUpdate200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchedulerTaskUpdate200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchedulerTaskUpdate200Response) SetType(v string)`

SetType sets Type field to given value.


### GetCron

`func (o *SchedulerTaskUpdate200Response) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *SchedulerTaskUpdate200Response) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *SchedulerTaskUpdate200Response) SetCron(v string)`

SetCron sets Cron field to given value.


### GetHost

`func (o *SchedulerTaskUpdate200Response) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SchedulerTaskUpdate200Response) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SchedulerTaskUpdate200Response) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SchedulerTaskUpdate200Response) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SchedulerTaskUpdate200Response) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SchedulerTaskUpdate200Response) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetStatus

`func (o *SchedulerTaskUpdate200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchedulerTaskUpdate200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchedulerTaskUpdate200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchedulerTaskUpdate200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SchedulerTaskUpdate200Response) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SchedulerTaskUpdate200Response) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastExecutionDate

`func (o *SchedulerTaskUpdate200Response) GetLastExecutionDate() int64`

GetLastExecutionDate returns the LastExecutionDate field if non-nil, zero value otherwise.

### GetLastExecutionDateOk

`func (o *SchedulerTaskUpdate200Response) GetLastExecutionDateOk() (*int64, bool)`

GetLastExecutionDateOk returns a tuple with the LastExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionDate

`func (o *SchedulerTaskUpdate200Response) SetLastExecutionDate(v int64)`

SetLastExecutionDate sets LastExecutionDate field to given value.

### HasLastExecutionDate

`func (o *SchedulerTaskUpdate200Response) HasLastExecutionDate() bool`

HasLastExecutionDate returns a boolean if a field has been set.

### SetLastExecutionDateNil

`func (o *SchedulerTaskUpdate200Response) SetLastExecutionDateNil(b bool)`

 SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil

### UnsetLastExecutionDate
`func (o *SchedulerTaskUpdate200Response) UnsetLastExecutionDate()`

UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
### GetLastCompletionDate

`func (o *SchedulerTaskUpdate200Response) GetLastCompletionDate() int64`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *SchedulerTaskUpdate200Response) GetLastCompletionDateOk() (*int64, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *SchedulerTaskUpdate200Response) SetLastCompletionDate(v int64)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *SchedulerTaskUpdate200Response) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### SetLastCompletionDateNil

`func (o *SchedulerTaskUpdate200Response) SetLastCompletionDateNil(b bool)`

 SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil

### UnsetLastCompletionDate
`func (o *SchedulerTaskUpdate200Response) UnsetLastCompletionDate()`

UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
### GetDetail

`func (o *SchedulerTaskUpdate200Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SchedulerTaskUpdate200Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SchedulerTaskUpdate200Response) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SchedulerTaskUpdate200Response) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *SchedulerTaskUpdate200Response) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *SchedulerTaskUpdate200Response) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetDryRun

`func (o *SchedulerTaskUpdate200Response) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SchedulerTaskUpdate200Response) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SchedulerTaskUpdate200Response) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetExecutionId

`func (o *SchedulerTaskUpdate200Response) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *SchedulerTaskUpdate200Response) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *SchedulerTaskUpdate200Response) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *SchedulerTaskUpdate200Response) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *SchedulerTaskUpdate200Response) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *SchedulerTaskUpdate200Response) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetEnabled

`func (o *SchedulerTaskUpdate200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SchedulerTaskUpdate200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SchedulerTaskUpdate200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetModule

`func (o *SchedulerTaskUpdate200Response) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *SchedulerTaskUpdate200Response) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *SchedulerTaskUpdate200Response) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *SchedulerTaskUpdate200Response) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SchedulerTaskUpdate200Response) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SchedulerTaskUpdate200Response) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetConnector

`func (o *SchedulerTaskUpdate200Response) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *SchedulerTaskUpdate200Response) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *SchedulerTaskUpdate200Response) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetEnroll

`func (o *SchedulerTaskUpdate200Response) GetEnroll() bool`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *SchedulerTaskUpdate200Response) GetEnrollOk() (*bool, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *SchedulerTaskUpdate200Response) SetEnroll(v bool)`

SetEnroll sets Enroll field to given value.


### GetRevoke

`func (o *SchedulerTaskUpdate200Response) GetRevoke() bool`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *SchedulerTaskUpdate200Response) GetRevokeOk() (*bool, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *SchedulerTaskUpdate200Response) SetRevoke(v bool)`

SetRevoke sets Revoke field to given value.


### GetRenew

`func (o *SchedulerTaskUpdate200Response) GetRenew() bool`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *SchedulerTaskUpdate200Response) GetRenewOk() (*bool, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *SchedulerTaskUpdate200Response) SetRenew(v bool)`

SetRenew sets Renew field to given value.


### GetResults

`func (o *SchedulerTaskUpdate200Response) GetResults() ThirdPartyConnectorSynchronizationResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SchedulerTaskUpdate200Response) GetResultsOk() (*ThirdPartyConnectorSynchronizationResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SchedulerTaskUpdate200Response) SetResults(v ThirdPartyConnectorSynchronizationResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *SchedulerTaskUpdate200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *SchedulerTaskUpdate200Response) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *SchedulerTaskUpdate200Response) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetDescription

`func (o *SchedulerTaskUpdate200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchedulerTaskUpdate200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchedulerTaskUpdate200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchedulerTaskUpdate200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SchedulerTaskUpdate200Response) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SchedulerTaskUpdate200Response) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *SchedulerTaskUpdate200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulerTaskUpdate200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulerTaskUpdate200Response) SetName(v string)`

SetName sets Name field to given value.


### GetFileName

`func (o *SchedulerTaskUpdate200Response) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SchedulerTaskUpdate200Response) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SchedulerTaskUpdate200Response) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SchedulerTaskUpdate200Response) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SchedulerTaskUpdate200Response) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SchedulerTaskUpdate200Response) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetRecipients

`func (o *SchedulerTaskUpdate200Response) GetRecipients() []ReportRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *SchedulerTaskUpdate200Response) GetRecipientsOk() (*[]ReportRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *SchedulerTaskUpdate200Response) SetRecipients(v []ReportRecipient)`

SetRecipients sets Recipients field to given value.


### GetFrom

`func (o *SchedulerTaskUpdate200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SchedulerTaskUpdate200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SchedulerTaskUpdate200Response) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTitle

`func (o *SchedulerTaskUpdate200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SchedulerTaskUpdate200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SchedulerTaskUpdate200Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *SchedulerTaskUpdate200Response) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SchedulerTaskUpdate200Response) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SchedulerTaskUpdate200Response) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SchedulerTaskUpdate200Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *SchedulerTaskUpdate200Response) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *SchedulerTaskUpdate200Response) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetIsHtml

`func (o *SchedulerTaskUpdate200Response) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *SchedulerTaskUpdate200Response) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *SchedulerTaskUpdate200Response) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetHqlType

`func (o *SchedulerTaskUpdate200Response) GetHqlType() string`

GetHqlType returns the HqlType field if non-nil, zero value otherwise.

### GetHqlTypeOk

`func (o *SchedulerTaskUpdate200Response) GetHqlTypeOk() (*string, bool)`

GetHqlTypeOk returns a tuple with the HqlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlType

`func (o *SchedulerTaskUpdate200Response) SetHqlType(v string)`

SetHqlType sets HqlType field to given value.


### GetHqlQuery

`func (o *SchedulerTaskUpdate200Response) GetHqlQuery() string`

GetHqlQuery returns the HqlQuery field if non-nil, zero value otherwise.

### GetHqlQueryOk

`func (o *SchedulerTaskUpdate200Response) GetHqlQueryOk() (*string, bool)`

GetHqlQueryOk returns a tuple with the HqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlQuery

`func (o *SchedulerTaskUpdate200Response) SetHqlQuery(v string)`

SetHqlQuery sets HqlQuery field to given value.

### HasHqlQuery

`func (o *SchedulerTaskUpdate200Response) HasHqlQuery() bool`

HasHqlQuery returns a boolean if a field has been set.

### SetHqlQueryNil

`func (o *SchedulerTaskUpdate200Response) SetHqlQueryNil(b bool)`

 SetHqlQueryNil sets the value for HqlQuery to be an explicit nil

### UnsetHqlQuery
`func (o *SchedulerTaskUpdate200Response) UnsetHqlQuery()`

UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
### GetHqlFields

`func (o *SchedulerTaskUpdate200Response) GetHqlFields() []string`

GetHqlFields returns the HqlFields field if non-nil, zero value otherwise.

### GetHqlFieldsOk

`func (o *SchedulerTaskUpdate200Response) GetHqlFieldsOk() (*[]string, bool)`

GetHqlFieldsOk returns a tuple with the HqlFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlFields

`func (o *SchedulerTaskUpdate200Response) SetHqlFields(v []string)`

SetHqlFields sets HqlFields field to given value.

### HasHqlFields

`func (o *SchedulerTaskUpdate200Response) HasHqlFields() bool`

HasHqlFields returns a boolean if a field has been set.

### SetHqlFieldsNil

`func (o *SchedulerTaskUpdate200Response) SetHqlFieldsNil(b bool)`

 SetHqlFieldsNil sets the value for HqlFields to be an explicit nil

### UnsetHqlFields
`func (o *SchedulerTaskUpdate200Response) UnsetHqlFields()`

UnsetHqlFields ensures that no value is present for HqlFields, not even an explicit nil
### GetHqlSortedBy

`func (o *SchedulerTaskUpdate200Response) GetHqlSortedBy() []SortElement`

GetHqlSortedBy returns the HqlSortedBy field if non-nil, zero value otherwise.

### GetHqlSortedByOk

`func (o *SchedulerTaskUpdate200Response) GetHqlSortedByOk() (*[]SortElement, bool)`

GetHqlSortedByOk returns a tuple with the HqlSortedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHqlSortedBy

`func (o *SchedulerTaskUpdate200Response) SetHqlSortedBy(v []SortElement)`

SetHqlSortedBy sets HqlSortedBy field to given value.

### HasHqlSortedBy

`func (o *SchedulerTaskUpdate200Response) HasHqlSortedBy() bool`

HasHqlSortedBy returns a boolean if a field has been set.

### SetHqlSortedByNil

`func (o *SchedulerTaskUpdate200Response) SetHqlSortedByNil(b bool)`

 SetHqlSortedByNil sets the value for HqlSortedBy to be an explicit nil

### UnsetHqlSortedBy
`func (o *SchedulerTaskUpdate200Response) UnsetHqlSortedBy()`

UnsetHqlSortedBy ensures that no value is present for HqlSortedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


