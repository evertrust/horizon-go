# CertificateProfileAuthorizationLevels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enroll** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**EnrollApi** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**RequestEnroll** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**ApproveEnroll** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**Revoke** | Pointer to [**AuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**RequestRevoke** | Pointer to [**AuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**ApproveRevoke** | Pointer to [**AuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**Search** | [**AuthorizationLevel**](AuthorizationLevel.md) |  | 
**Update** | [**AuthorizationLevel**](AuthorizationLevel.md) |  | 
**RequestUpdate** | [**AuthorizationLevel**](AuthorizationLevel.md) |  | 
**ApproveUpdate** | [**AuthorizationLevel**](AuthorizationLevel.md) |  | 
**Recover** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**RecoverApi** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**RequestRecover** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**ApproveRecover** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**Migrate** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**RequestMigrate** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**ApproveMigrate** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**Renew** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**RenewApi** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**RequestRenew** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**ApproveRenew** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 
**AuditRequest** | Pointer to [**NullableAuthorizationLevel**](AuthorizationLevel.md) |  | [optional] 

## Methods

### NewCertificateProfileAuthorizationLevels

`func NewCertificateProfileAuthorizationLevels(search AuthorizationLevel, update AuthorizationLevel, requestUpdate AuthorizationLevel, approveUpdate AuthorizationLevel, ) *CertificateProfileAuthorizationLevels`

NewCertificateProfileAuthorizationLevels instantiates a new CertificateProfileAuthorizationLevels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileAuthorizationLevelsWithDefaults

`func NewCertificateProfileAuthorizationLevelsWithDefaults() *CertificateProfileAuthorizationLevels`

NewCertificateProfileAuthorizationLevelsWithDefaults instantiates a new CertificateProfileAuthorizationLevels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnroll

`func (o *CertificateProfileAuthorizationLevels) GetEnroll() AuthorizationLevel`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *CertificateProfileAuthorizationLevels) GetEnrollOk() (*AuthorizationLevel, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *CertificateProfileAuthorizationLevels) SetEnroll(v AuthorizationLevel)`

SetEnroll sets Enroll field to given value.

### HasEnroll

`func (o *CertificateProfileAuthorizationLevels) HasEnroll() bool`

HasEnroll returns a boolean if a field has been set.

### SetEnrollNil

`func (o *CertificateProfileAuthorizationLevels) SetEnrollNil(b bool)`

 SetEnrollNil sets the value for Enroll to be an explicit nil

### UnsetEnroll
`func (o *CertificateProfileAuthorizationLevels) UnsetEnroll()`

UnsetEnroll ensures that no value is present for Enroll, not even an explicit nil
### GetEnrollApi

`func (o *CertificateProfileAuthorizationLevels) GetEnrollApi() AuthorizationLevel`

GetEnrollApi returns the EnrollApi field if non-nil, zero value otherwise.

### GetEnrollApiOk

`func (o *CertificateProfileAuthorizationLevels) GetEnrollApiOk() (*AuthorizationLevel, bool)`

GetEnrollApiOk returns a tuple with the EnrollApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollApi

`func (o *CertificateProfileAuthorizationLevels) SetEnrollApi(v AuthorizationLevel)`

SetEnrollApi sets EnrollApi field to given value.

### HasEnrollApi

`func (o *CertificateProfileAuthorizationLevels) HasEnrollApi() bool`

HasEnrollApi returns a boolean if a field has been set.

### SetEnrollApiNil

`func (o *CertificateProfileAuthorizationLevels) SetEnrollApiNil(b bool)`

 SetEnrollApiNil sets the value for EnrollApi to be an explicit nil

### UnsetEnrollApi
`func (o *CertificateProfileAuthorizationLevels) UnsetEnrollApi()`

UnsetEnrollApi ensures that no value is present for EnrollApi, not even an explicit nil
### GetRequestEnroll

`func (o *CertificateProfileAuthorizationLevels) GetRequestEnroll() AuthorizationLevel`

GetRequestEnroll returns the RequestEnroll field if non-nil, zero value otherwise.

### GetRequestEnrollOk

`func (o *CertificateProfileAuthorizationLevels) GetRequestEnrollOk() (*AuthorizationLevel, bool)`

GetRequestEnrollOk returns a tuple with the RequestEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEnroll

`func (o *CertificateProfileAuthorizationLevels) SetRequestEnroll(v AuthorizationLevel)`

SetRequestEnroll sets RequestEnroll field to given value.

### HasRequestEnroll

`func (o *CertificateProfileAuthorizationLevels) HasRequestEnroll() bool`

HasRequestEnroll returns a boolean if a field has been set.

### SetRequestEnrollNil

`func (o *CertificateProfileAuthorizationLevels) SetRequestEnrollNil(b bool)`

 SetRequestEnrollNil sets the value for RequestEnroll to be an explicit nil

### UnsetRequestEnroll
`func (o *CertificateProfileAuthorizationLevels) UnsetRequestEnroll()`

UnsetRequestEnroll ensures that no value is present for RequestEnroll, not even an explicit nil
### GetApproveEnroll

`func (o *CertificateProfileAuthorizationLevels) GetApproveEnroll() AuthorizationLevel`

GetApproveEnroll returns the ApproveEnroll field if non-nil, zero value otherwise.

### GetApproveEnrollOk

`func (o *CertificateProfileAuthorizationLevels) GetApproveEnrollOk() (*AuthorizationLevel, bool)`

GetApproveEnrollOk returns a tuple with the ApproveEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveEnroll

`func (o *CertificateProfileAuthorizationLevels) SetApproveEnroll(v AuthorizationLevel)`

SetApproveEnroll sets ApproveEnroll field to given value.

### HasApproveEnroll

`func (o *CertificateProfileAuthorizationLevels) HasApproveEnroll() bool`

HasApproveEnroll returns a boolean if a field has been set.

### SetApproveEnrollNil

`func (o *CertificateProfileAuthorizationLevels) SetApproveEnrollNil(b bool)`

 SetApproveEnrollNil sets the value for ApproveEnroll to be an explicit nil

### UnsetApproveEnroll
`func (o *CertificateProfileAuthorizationLevels) UnsetApproveEnroll()`

UnsetApproveEnroll ensures that no value is present for ApproveEnroll, not even an explicit nil
### GetRevoke

`func (o *CertificateProfileAuthorizationLevels) GetRevoke() AuthorizationLevel`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *CertificateProfileAuthorizationLevels) GetRevokeOk() (*AuthorizationLevel, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *CertificateProfileAuthorizationLevels) SetRevoke(v AuthorizationLevel)`

SetRevoke sets Revoke field to given value.

### HasRevoke

`func (o *CertificateProfileAuthorizationLevels) HasRevoke() bool`

HasRevoke returns a boolean if a field has been set.

### GetRequestRevoke

`func (o *CertificateProfileAuthorizationLevels) GetRequestRevoke() AuthorizationLevel`

GetRequestRevoke returns the RequestRevoke field if non-nil, zero value otherwise.

### GetRequestRevokeOk

`func (o *CertificateProfileAuthorizationLevels) GetRequestRevokeOk() (*AuthorizationLevel, bool)`

GetRequestRevokeOk returns a tuple with the RequestRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRevoke

`func (o *CertificateProfileAuthorizationLevels) SetRequestRevoke(v AuthorizationLevel)`

SetRequestRevoke sets RequestRevoke field to given value.

### HasRequestRevoke

`func (o *CertificateProfileAuthorizationLevels) HasRequestRevoke() bool`

HasRequestRevoke returns a boolean if a field has been set.

### GetApproveRevoke

`func (o *CertificateProfileAuthorizationLevels) GetApproveRevoke() AuthorizationLevel`

GetApproveRevoke returns the ApproveRevoke field if non-nil, zero value otherwise.

### GetApproveRevokeOk

`func (o *CertificateProfileAuthorizationLevels) GetApproveRevokeOk() (*AuthorizationLevel, bool)`

GetApproveRevokeOk returns a tuple with the ApproveRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveRevoke

`func (o *CertificateProfileAuthorizationLevels) SetApproveRevoke(v AuthorizationLevel)`

SetApproveRevoke sets ApproveRevoke field to given value.

### HasApproveRevoke

`func (o *CertificateProfileAuthorizationLevels) HasApproveRevoke() bool`

HasApproveRevoke returns a boolean if a field has been set.

### GetSearch

`func (o *CertificateProfileAuthorizationLevels) GetSearch() AuthorizationLevel`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CertificateProfileAuthorizationLevels) GetSearchOk() (*AuthorizationLevel, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CertificateProfileAuthorizationLevels) SetSearch(v AuthorizationLevel)`

SetSearch sets Search field to given value.


### GetUpdate

`func (o *CertificateProfileAuthorizationLevels) GetUpdate() AuthorizationLevel`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *CertificateProfileAuthorizationLevels) GetUpdateOk() (*AuthorizationLevel, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *CertificateProfileAuthorizationLevels) SetUpdate(v AuthorizationLevel)`

SetUpdate sets Update field to given value.


### GetRequestUpdate

`func (o *CertificateProfileAuthorizationLevels) GetRequestUpdate() AuthorizationLevel`

GetRequestUpdate returns the RequestUpdate field if non-nil, zero value otherwise.

### GetRequestUpdateOk

`func (o *CertificateProfileAuthorizationLevels) GetRequestUpdateOk() (*AuthorizationLevel, bool)`

GetRequestUpdateOk returns a tuple with the RequestUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUpdate

`func (o *CertificateProfileAuthorizationLevels) SetRequestUpdate(v AuthorizationLevel)`

SetRequestUpdate sets RequestUpdate field to given value.


### GetApproveUpdate

`func (o *CertificateProfileAuthorizationLevels) GetApproveUpdate() AuthorizationLevel`

GetApproveUpdate returns the ApproveUpdate field if non-nil, zero value otherwise.

### GetApproveUpdateOk

`func (o *CertificateProfileAuthorizationLevels) GetApproveUpdateOk() (*AuthorizationLevel, bool)`

GetApproveUpdateOk returns a tuple with the ApproveUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveUpdate

`func (o *CertificateProfileAuthorizationLevels) SetApproveUpdate(v AuthorizationLevel)`

SetApproveUpdate sets ApproveUpdate field to given value.


### GetRecover

`func (o *CertificateProfileAuthorizationLevels) GetRecover() AuthorizationLevel`

GetRecover returns the Recover field if non-nil, zero value otherwise.

### GetRecoverOk

`func (o *CertificateProfileAuthorizationLevels) GetRecoverOk() (*AuthorizationLevel, bool)`

GetRecoverOk returns a tuple with the Recover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecover

`func (o *CertificateProfileAuthorizationLevels) SetRecover(v AuthorizationLevel)`

SetRecover sets Recover field to given value.

### HasRecover

`func (o *CertificateProfileAuthorizationLevels) HasRecover() bool`

HasRecover returns a boolean if a field has been set.

### SetRecoverNil

`func (o *CertificateProfileAuthorizationLevels) SetRecoverNil(b bool)`

 SetRecoverNil sets the value for Recover to be an explicit nil

### UnsetRecover
`func (o *CertificateProfileAuthorizationLevels) UnsetRecover()`

UnsetRecover ensures that no value is present for Recover, not even an explicit nil
### GetRecoverApi

`func (o *CertificateProfileAuthorizationLevels) GetRecoverApi() AuthorizationLevel`

GetRecoverApi returns the RecoverApi field if non-nil, zero value otherwise.

### GetRecoverApiOk

`func (o *CertificateProfileAuthorizationLevels) GetRecoverApiOk() (*AuthorizationLevel, bool)`

GetRecoverApiOk returns a tuple with the RecoverApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverApi

`func (o *CertificateProfileAuthorizationLevels) SetRecoverApi(v AuthorizationLevel)`

SetRecoverApi sets RecoverApi field to given value.

### HasRecoverApi

`func (o *CertificateProfileAuthorizationLevels) HasRecoverApi() bool`

HasRecoverApi returns a boolean if a field has been set.

### SetRecoverApiNil

`func (o *CertificateProfileAuthorizationLevels) SetRecoverApiNil(b bool)`

 SetRecoverApiNil sets the value for RecoverApi to be an explicit nil

### UnsetRecoverApi
`func (o *CertificateProfileAuthorizationLevels) UnsetRecoverApi()`

UnsetRecoverApi ensures that no value is present for RecoverApi, not even an explicit nil
### GetRequestRecover

`func (o *CertificateProfileAuthorizationLevels) GetRequestRecover() AuthorizationLevel`

GetRequestRecover returns the RequestRecover field if non-nil, zero value otherwise.

### GetRequestRecoverOk

`func (o *CertificateProfileAuthorizationLevels) GetRequestRecoverOk() (*AuthorizationLevel, bool)`

GetRequestRecoverOk returns a tuple with the RequestRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRecover

`func (o *CertificateProfileAuthorizationLevels) SetRequestRecover(v AuthorizationLevel)`

SetRequestRecover sets RequestRecover field to given value.

### HasRequestRecover

`func (o *CertificateProfileAuthorizationLevels) HasRequestRecover() bool`

HasRequestRecover returns a boolean if a field has been set.

### SetRequestRecoverNil

`func (o *CertificateProfileAuthorizationLevels) SetRequestRecoverNil(b bool)`

 SetRequestRecoverNil sets the value for RequestRecover to be an explicit nil

### UnsetRequestRecover
`func (o *CertificateProfileAuthorizationLevels) UnsetRequestRecover()`

UnsetRequestRecover ensures that no value is present for RequestRecover, not even an explicit nil
### GetApproveRecover

`func (o *CertificateProfileAuthorizationLevels) GetApproveRecover() AuthorizationLevel`

GetApproveRecover returns the ApproveRecover field if non-nil, zero value otherwise.

### GetApproveRecoverOk

`func (o *CertificateProfileAuthorizationLevels) GetApproveRecoverOk() (*AuthorizationLevel, bool)`

GetApproveRecoverOk returns a tuple with the ApproveRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveRecover

`func (o *CertificateProfileAuthorizationLevels) SetApproveRecover(v AuthorizationLevel)`

SetApproveRecover sets ApproveRecover field to given value.

### HasApproveRecover

`func (o *CertificateProfileAuthorizationLevels) HasApproveRecover() bool`

HasApproveRecover returns a boolean if a field has been set.

### SetApproveRecoverNil

`func (o *CertificateProfileAuthorizationLevels) SetApproveRecoverNil(b bool)`

 SetApproveRecoverNil sets the value for ApproveRecover to be an explicit nil

### UnsetApproveRecover
`func (o *CertificateProfileAuthorizationLevels) UnsetApproveRecover()`

UnsetApproveRecover ensures that no value is present for ApproveRecover, not even an explicit nil
### GetMigrate

`func (o *CertificateProfileAuthorizationLevels) GetMigrate() AuthorizationLevel`

GetMigrate returns the Migrate field if non-nil, zero value otherwise.

### GetMigrateOk

`func (o *CertificateProfileAuthorizationLevels) GetMigrateOk() (*AuthorizationLevel, bool)`

GetMigrateOk returns a tuple with the Migrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrate

`func (o *CertificateProfileAuthorizationLevels) SetMigrate(v AuthorizationLevel)`

SetMigrate sets Migrate field to given value.

### HasMigrate

`func (o *CertificateProfileAuthorizationLevels) HasMigrate() bool`

HasMigrate returns a boolean if a field has been set.

### SetMigrateNil

`func (o *CertificateProfileAuthorizationLevels) SetMigrateNil(b bool)`

 SetMigrateNil sets the value for Migrate to be an explicit nil

### UnsetMigrate
`func (o *CertificateProfileAuthorizationLevels) UnsetMigrate()`

UnsetMigrate ensures that no value is present for Migrate, not even an explicit nil
### GetRequestMigrate

`func (o *CertificateProfileAuthorizationLevels) GetRequestMigrate() AuthorizationLevel`

GetRequestMigrate returns the RequestMigrate field if non-nil, zero value otherwise.

### GetRequestMigrateOk

`func (o *CertificateProfileAuthorizationLevels) GetRequestMigrateOk() (*AuthorizationLevel, bool)`

GetRequestMigrateOk returns a tuple with the RequestMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMigrate

`func (o *CertificateProfileAuthorizationLevels) SetRequestMigrate(v AuthorizationLevel)`

SetRequestMigrate sets RequestMigrate field to given value.

### HasRequestMigrate

`func (o *CertificateProfileAuthorizationLevels) HasRequestMigrate() bool`

HasRequestMigrate returns a boolean if a field has been set.

### SetRequestMigrateNil

`func (o *CertificateProfileAuthorizationLevels) SetRequestMigrateNil(b bool)`

 SetRequestMigrateNil sets the value for RequestMigrate to be an explicit nil

### UnsetRequestMigrate
`func (o *CertificateProfileAuthorizationLevels) UnsetRequestMigrate()`

UnsetRequestMigrate ensures that no value is present for RequestMigrate, not even an explicit nil
### GetApproveMigrate

`func (o *CertificateProfileAuthorizationLevels) GetApproveMigrate() AuthorizationLevel`

GetApproveMigrate returns the ApproveMigrate field if non-nil, zero value otherwise.

### GetApproveMigrateOk

`func (o *CertificateProfileAuthorizationLevels) GetApproveMigrateOk() (*AuthorizationLevel, bool)`

GetApproveMigrateOk returns a tuple with the ApproveMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveMigrate

`func (o *CertificateProfileAuthorizationLevels) SetApproveMigrate(v AuthorizationLevel)`

SetApproveMigrate sets ApproveMigrate field to given value.

### HasApproveMigrate

`func (o *CertificateProfileAuthorizationLevels) HasApproveMigrate() bool`

HasApproveMigrate returns a boolean if a field has been set.

### SetApproveMigrateNil

`func (o *CertificateProfileAuthorizationLevels) SetApproveMigrateNil(b bool)`

 SetApproveMigrateNil sets the value for ApproveMigrate to be an explicit nil

### UnsetApproveMigrate
`func (o *CertificateProfileAuthorizationLevels) UnsetApproveMigrate()`

UnsetApproveMigrate ensures that no value is present for ApproveMigrate, not even an explicit nil
### GetRenew

`func (o *CertificateProfileAuthorizationLevels) GetRenew() AuthorizationLevel`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *CertificateProfileAuthorizationLevels) GetRenewOk() (*AuthorizationLevel, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *CertificateProfileAuthorizationLevels) SetRenew(v AuthorizationLevel)`

SetRenew sets Renew field to given value.

### HasRenew

`func (o *CertificateProfileAuthorizationLevels) HasRenew() bool`

HasRenew returns a boolean if a field has been set.

### SetRenewNil

`func (o *CertificateProfileAuthorizationLevels) SetRenewNil(b bool)`

 SetRenewNil sets the value for Renew to be an explicit nil

### UnsetRenew
`func (o *CertificateProfileAuthorizationLevels) UnsetRenew()`

UnsetRenew ensures that no value is present for Renew, not even an explicit nil
### GetRenewApi

`func (o *CertificateProfileAuthorizationLevels) GetRenewApi() AuthorizationLevel`

GetRenewApi returns the RenewApi field if non-nil, zero value otherwise.

### GetRenewApiOk

`func (o *CertificateProfileAuthorizationLevels) GetRenewApiOk() (*AuthorizationLevel, bool)`

GetRenewApiOk returns a tuple with the RenewApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewApi

`func (o *CertificateProfileAuthorizationLevels) SetRenewApi(v AuthorizationLevel)`

SetRenewApi sets RenewApi field to given value.

### HasRenewApi

`func (o *CertificateProfileAuthorizationLevels) HasRenewApi() bool`

HasRenewApi returns a boolean if a field has been set.

### SetRenewApiNil

`func (o *CertificateProfileAuthorizationLevels) SetRenewApiNil(b bool)`

 SetRenewApiNil sets the value for RenewApi to be an explicit nil

### UnsetRenewApi
`func (o *CertificateProfileAuthorizationLevels) UnsetRenewApi()`

UnsetRenewApi ensures that no value is present for RenewApi, not even an explicit nil
### GetRequestRenew

`func (o *CertificateProfileAuthorizationLevels) GetRequestRenew() AuthorizationLevel`

GetRequestRenew returns the RequestRenew field if non-nil, zero value otherwise.

### GetRequestRenewOk

`func (o *CertificateProfileAuthorizationLevels) GetRequestRenewOk() (*AuthorizationLevel, bool)`

GetRequestRenewOk returns a tuple with the RequestRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRenew

`func (o *CertificateProfileAuthorizationLevels) SetRequestRenew(v AuthorizationLevel)`

SetRequestRenew sets RequestRenew field to given value.

### HasRequestRenew

`func (o *CertificateProfileAuthorizationLevels) HasRequestRenew() bool`

HasRequestRenew returns a boolean if a field has been set.

### SetRequestRenewNil

`func (o *CertificateProfileAuthorizationLevels) SetRequestRenewNil(b bool)`

 SetRequestRenewNil sets the value for RequestRenew to be an explicit nil

### UnsetRequestRenew
`func (o *CertificateProfileAuthorizationLevels) UnsetRequestRenew()`

UnsetRequestRenew ensures that no value is present for RequestRenew, not even an explicit nil
### GetApproveRenew

`func (o *CertificateProfileAuthorizationLevels) GetApproveRenew() AuthorizationLevel`

GetApproveRenew returns the ApproveRenew field if non-nil, zero value otherwise.

### GetApproveRenewOk

`func (o *CertificateProfileAuthorizationLevels) GetApproveRenewOk() (*AuthorizationLevel, bool)`

GetApproveRenewOk returns a tuple with the ApproveRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveRenew

`func (o *CertificateProfileAuthorizationLevels) SetApproveRenew(v AuthorizationLevel)`

SetApproveRenew sets ApproveRenew field to given value.

### HasApproveRenew

`func (o *CertificateProfileAuthorizationLevels) HasApproveRenew() bool`

HasApproveRenew returns a boolean if a field has been set.

### SetApproveRenewNil

`func (o *CertificateProfileAuthorizationLevels) SetApproveRenewNil(b bool)`

 SetApproveRenewNil sets the value for ApproveRenew to be an explicit nil

### UnsetApproveRenew
`func (o *CertificateProfileAuthorizationLevels) UnsetApproveRenew()`

UnsetApproveRenew ensures that no value is present for ApproveRenew, not even an explicit nil
### GetAuditRequest

`func (o *CertificateProfileAuthorizationLevels) GetAuditRequest() AuthorizationLevel`

GetAuditRequest returns the AuditRequest field if non-nil, zero value otherwise.

### GetAuditRequestOk

`func (o *CertificateProfileAuthorizationLevels) GetAuditRequestOk() (*AuthorizationLevel, bool)`

GetAuditRequestOk returns a tuple with the AuditRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditRequest

`func (o *CertificateProfileAuthorizationLevels) SetAuditRequest(v AuthorizationLevel)`

SetAuditRequest sets AuditRequest field to given value.

### HasAuditRequest

`func (o *CertificateProfileAuthorizationLevels) HasAuditRequest() bool`

HasAuditRequest returns a boolean if a field has been set.

### SetAuditRequestNil

`func (o *CertificateProfileAuthorizationLevels) SetAuditRequestNil(b bool)`

 SetAuditRequestNil sets the value for AuditRequest to be an explicit nil

### UnsetAuditRequest
`func (o *CertificateProfileAuthorizationLevels) UnsetAuditRequest()`

UnsetAuditRequest ensures that no value is present for AuditRequest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


