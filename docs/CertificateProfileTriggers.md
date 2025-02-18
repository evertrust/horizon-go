# CertificateProfileTriggers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnEnroll** | Pointer to **[]string** |  | [optional] 
**OnSubmitEnroll** | Pointer to **[]string** |  | [optional] 
**OnCancelEnroll** | Pointer to **[]string** |  | [optional] 
**OnApproveEnroll** | Pointer to **[]string** |  | [optional] 
**OnDenyEnroll** | Pointer to **[]string** |  | [optional] 
**OnPendingEnroll** | Pointer to [**[]CertificateProfileAsynchronousTrigger**](CertificateProfileAsynchronousTrigger.md) |  | [optional] 
**OnRevoke** | Pointer to **[]string** |  | [optional] 
**OnSubmitRevoke** | Pointer to **[]string** |  | [optional] 
**OnCancelRevoke** | Pointer to **[]string** |  | [optional] 
**OnApproveRevoke** | Pointer to **[]string** |  | [optional] 
**OnDenyRevoke** | Pointer to **[]string** |  | [optional] 
**OnPendingRevoke** | Pointer to [**[]CertificateProfileAsynchronousTrigger**](CertificateProfileAsynchronousTrigger.md) |  | [optional] 
**OnUpdate** | Pointer to **[]string** |  | [optional] 
**OnSubmitUpdate** | Pointer to **[]string** |  | [optional] 
**OnCancelUpdate** | Pointer to **[]string** |  | [optional] 
**OnApproveUpdate** | Pointer to **[]string** |  | [optional] 
**OnDenyUpdate** | Pointer to **[]string** |  | [optional] 
**OnPendingUpdate** | Pointer to [**[]CertificateProfileAsynchronousTrigger**](CertificateProfileAsynchronousTrigger.md) |  | [optional] 
**OnRecover** | Pointer to **[]string** |  | [optional] 
**OnSubmitRecover** | Pointer to **[]string** |  | [optional] 
**OnCancelRecover** | Pointer to **[]string** |  | [optional] 
**OnApproveRecover** | Pointer to **[]string** |  | [optional] 
**OnDenyRecover** | Pointer to **[]string** |  | [optional] 
**OnPendingRecover** | Pointer to [**[]CertificateProfileAsynchronousTrigger**](CertificateProfileAsynchronousTrigger.md) |  | [optional] 
**OnMigrate** | Pointer to **[]string** |  | [optional] 
**OnSubmitMigrate** | Pointer to **[]string** |  | [optional] 
**OnCancelMigrate** | Pointer to **[]string** |  | [optional] 
**OnApproveMigrate** | Pointer to **[]string** |  | [optional] 
**OnDenyMigrate** | Pointer to **[]string** |  | [optional] 
**OnPendingMigrate** | Pointer to [**[]CertificateProfileAsynchronousTrigger**](CertificateProfileAsynchronousTrigger.md) |  | [optional] 
**OnExpire** | Pointer to [**[]CertificateProfileAsynchronousTrigger**](CertificateProfileAsynchronousTrigger.md) |  | [optional] 
**OnRenew** | Pointer to **[]string** |  | [optional] 
**OnSubmitRenew** | Pointer to **[]string** |  | [optional] 
**OnCancelRenew** | Pointer to **[]string** |  | [optional] 
**OnApproveRenew** | Pointer to **[]string** |  | [optional] 
**OnDenyRenew** | Pointer to **[]string** |  | [optional] 
**OnPendingRenew** | Pointer to [**[]CertificateProfileAsynchronousTrigger**](CertificateProfileAsynchronousTrigger.md) |  | [optional] 

## Methods

### NewCertificateProfileTriggers

`func NewCertificateProfileTriggers() *CertificateProfileTriggers`

NewCertificateProfileTriggers instantiates a new CertificateProfileTriggers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileTriggersWithDefaults

`func NewCertificateProfileTriggersWithDefaults() *CertificateProfileTriggers`

NewCertificateProfileTriggersWithDefaults instantiates a new CertificateProfileTriggers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnEnroll

`func (o *CertificateProfileTriggers) GetOnEnroll() []string`

GetOnEnroll returns the OnEnroll field if non-nil, zero value otherwise.

### GetOnEnrollOk

`func (o *CertificateProfileTriggers) GetOnEnrollOk() (*[]string, bool)`

GetOnEnrollOk returns a tuple with the OnEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnEnroll

`func (o *CertificateProfileTriggers) SetOnEnroll(v []string)`

SetOnEnroll sets OnEnroll field to given value.

### HasOnEnroll

`func (o *CertificateProfileTriggers) HasOnEnroll() bool`

HasOnEnroll returns a boolean if a field has been set.

### SetOnEnrollNil

`func (o *CertificateProfileTriggers) SetOnEnrollNil(b bool)`

 SetOnEnrollNil sets the value for OnEnroll to be an explicit nil

### UnsetOnEnroll
`func (o *CertificateProfileTriggers) UnsetOnEnroll()`

UnsetOnEnroll ensures that no value is present for OnEnroll, not even an explicit nil
### GetOnSubmitEnroll

`func (o *CertificateProfileTriggers) GetOnSubmitEnroll() []string`

GetOnSubmitEnroll returns the OnSubmitEnroll field if non-nil, zero value otherwise.

### GetOnSubmitEnrollOk

`func (o *CertificateProfileTriggers) GetOnSubmitEnrollOk() (*[]string, bool)`

GetOnSubmitEnrollOk returns a tuple with the OnSubmitEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSubmitEnroll

`func (o *CertificateProfileTriggers) SetOnSubmitEnroll(v []string)`

SetOnSubmitEnroll sets OnSubmitEnroll field to given value.

### HasOnSubmitEnroll

`func (o *CertificateProfileTriggers) HasOnSubmitEnroll() bool`

HasOnSubmitEnroll returns a boolean if a field has been set.

### SetOnSubmitEnrollNil

`func (o *CertificateProfileTriggers) SetOnSubmitEnrollNil(b bool)`

 SetOnSubmitEnrollNil sets the value for OnSubmitEnroll to be an explicit nil

### UnsetOnSubmitEnroll
`func (o *CertificateProfileTriggers) UnsetOnSubmitEnroll()`

UnsetOnSubmitEnroll ensures that no value is present for OnSubmitEnroll, not even an explicit nil
### GetOnCancelEnroll

`func (o *CertificateProfileTriggers) GetOnCancelEnroll() []string`

GetOnCancelEnroll returns the OnCancelEnroll field if non-nil, zero value otherwise.

### GetOnCancelEnrollOk

`func (o *CertificateProfileTriggers) GetOnCancelEnrollOk() (*[]string, bool)`

GetOnCancelEnrollOk returns a tuple with the OnCancelEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCancelEnroll

`func (o *CertificateProfileTriggers) SetOnCancelEnroll(v []string)`

SetOnCancelEnroll sets OnCancelEnroll field to given value.

### HasOnCancelEnroll

`func (o *CertificateProfileTriggers) HasOnCancelEnroll() bool`

HasOnCancelEnroll returns a boolean if a field has been set.

### SetOnCancelEnrollNil

`func (o *CertificateProfileTriggers) SetOnCancelEnrollNil(b bool)`

 SetOnCancelEnrollNil sets the value for OnCancelEnroll to be an explicit nil

### UnsetOnCancelEnroll
`func (o *CertificateProfileTriggers) UnsetOnCancelEnroll()`

UnsetOnCancelEnroll ensures that no value is present for OnCancelEnroll, not even an explicit nil
### GetOnApproveEnroll

`func (o *CertificateProfileTriggers) GetOnApproveEnroll() []string`

GetOnApproveEnroll returns the OnApproveEnroll field if non-nil, zero value otherwise.

### GetOnApproveEnrollOk

`func (o *CertificateProfileTriggers) GetOnApproveEnrollOk() (*[]string, bool)`

GetOnApproveEnrollOk returns a tuple with the OnApproveEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnApproveEnroll

`func (o *CertificateProfileTriggers) SetOnApproveEnroll(v []string)`

SetOnApproveEnroll sets OnApproveEnroll field to given value.

### HasOnApproveEnroll

`func (o *CertificateProfileTriggers) HasOnApproveEnroll() bool`

HasOnApproveEnroll returns a boolean if a field has been set.

### SetOnApproveEnrollNil

`func (o *CertificateProfileTriggers) SetOnApproveEnrollNil(b bool)`

 SetOnApproveEnrollNil sets the value for OnApproveEnroll to be an explicit nil

### UnsetOnApproveEnroll
`func (o *CertificateProfileTriggers) UnsetOnApproveEnroll()`

UnsetOnApproveEnroll ensures that no value is present for OnApproveEnroll, not even an explicit nil
### GetOnDenyEnroll

`func (o *CertificateProfileTriggers) GetOnDenyEnroll() []string`

GetOnDenyEnroll returns the OnDenyEnroll field if non-nil, zero value otherwise.

### GetOnDenyEnrollOk

`func (o *CertificateProfileTriggers) GetOnDenyEnrollOk() (*[]string, bool)`

GetOnDenyEnrollOk returns a tuple with the OnDenyEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDenyEnroll

`func (o *CertificateProfileTriggers) SetOnDenyEnroll(v []string)`

SetOnDenyEnroll sets OnDenyEnroll field to given value.

### HasOnDenyEnroll

`func (o *CertificateProfileTriggers) HasOnDenyEnroll() bool`

HasOnDenyEnroll returns a boolean if a field has been set.

### SetOnDenyEnrollNil

`func (o *CertificateProfileTriggers) SetOnDenyEnrollNil(b bool)`

 SetOnDenyEnrollNil sets the value for OnDenyEnroll to be an explicit nil

### UnsetOnDenyEnroll
`func (o *CertificateProfileTriggers) UnsetOnDenyEnroll()`

UnsetOnDenyEnroll ensures that no value is present for OnDenyEnroll, not even an explicit nil
### GetOnPendingEnroll

`func (o *CertificateProfileTriggers) GetOnPendingEnroll() []CertificateProfileAsynchronousTrigger`

GetOnPendingEnroll returns the OnPendingEnroll field if non-nil, zero value otherwise.

### GetOnPendingEnrollOk

`func (o *CertificateProfileTriggers) GetOnPendingEnrollOk() (*[]CertificateProfileAsynchronousTrigger, bool)`

GetOnPendingEnrollOk returns a tuple with the OnPendingEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPendingEnroll

`func (o *CertificateProfileTriggers) SetOnPendingEnroll(v []CertificateProfileAsynchronousTrigger)`

SetOnPendingEnroll sets OnPendingEnroll field to given value.

### HasOnPendingEnroll

`func (o *CertificateProfileTriggers) HasOnPendingEnroll() bool`

HasOnPendingEnroll returns a boolean if a field has been set.

### SetOnPendingEnrollNil

`func (o *CertificateProfileTriggers) SetOnPendingEnrollNil(b bool)`

 SetOnPendingEnrollNil sets the value for OnPendingEnroll to be an explicit nil

### UnsetOnPendingEnroll
`func (o *CertificateProfileTriggers) UnsetOnPendingEnroll()`

UnsetOnPendingEnroll ensures that no value is present for OnPendingEnroll, not even an explicit nil
### GetOnRevoke

`func (o *CertificateProfileTriggers) GetOnRevoke() []string`

GetOnRevoke returns the OnRevoke field if non-nil, zero value otherwise.

### GetOnRevokeOk

`func (o *CertificateProfileTriggers) GetOnRevokeOk() (*[]string, bool)`

GetOnRevokeOk returns a tuple with the OnRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnRevoke

`func (o *CertificateProfileTriggers) SetOnRevoke(v []string)`

SetOnRevoke sets OnRevoke field to given value.

### HasOnRevoke

`func (o *CertificateProfileTriggers) HasOnRevoke() bool`

HasOnRevoke returns a boolean if a field has been set.

### SetOnRevokeNil

`func (o *CertificateProfileTriggers) SetOnRevokeNil(b bool)`

 SetOnRevokeNil sets the value for OnRevoke to be an explicit nil

### UnsetOnRevoke
`func (o *CertificateProfileTriggers) UnsetOnRevoke()`

UnsetOnRevoke ensures that no value is present for OnRevoke, not even an explicit nil
### GetOnSubmitRevoke

`func (o *CertificateProfileTriggers) GetOnSubmitRevoke() []string`

GetOnSubmitRevoke returns the OnSubmitRevoke field if non-nil, zero value otherwise.

### GetOnSubmitRevokeOk

`func (o *CertificateProfileTriggers) GetOnSubmitRevokeOk() (*[]string, bool)`

GetOnSubmitRevokeOk returns a tuple with the OnSubmitRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSubmitRevoke

`func (o *CertificateProfileTriggers) SetOnSubmitRevoke(v []string)`

SetOnSubmitRevoke sets OnSubmitRevoke field to given value.

### HasOnSubmitRevoke

`func (o *CertificateProfileTriggers) HasOnSubmitRevoke() bool`

HasOnSubmitRevoke returns a boolean if a field has been set.

### SetOnSubmitRevokeNil

`func (o *CertificateProfileTriggers) SetOnSubmitRevokeNil(b bool)`

 SetOnSubmitRevokeNil sets the value for OnSubmitRevoke to be an explicit nil

### UnsetOnSubmitRevoke
`func (o *CertificateProfileTriggers) UnsetOnSubmitRevoke()`

UnsetOnSubmitRevoke ensures that no value is present for OnSubmitRevoke, not even an explicit nil
### GetOnCancelRevoke

`func (o *CertificateProfileTriggers) GetOnCancelRevoke() []string`

GetOnCancelRevoke returns the OnCancelRevoke field if non-nil, zero value otherwise.

### GetOnCancelRevokeOk

`func (o *CertificateProfileTriggers) GetOnCancelRevokeOk() (*[]string, bool)`

GetOnCancelRevokeOk returns a tuple with the OnCancelRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCancelRevoke

`func (o *CertificateProfileTriggers) SetOnCancelRevoke(v []string)`

SetOnCancelRevoke sets OnCancelRevoke field to given value.

### HasOnCancelRevoke

`func (o *CertificateProfileTriggers) HasOnCancelRevoke() bool`

HasOnCancelRevoke returns a boolean if a field has been set.

### SetOnCancelRevokeNil

`func (o *CertificateProfileTriggers) SetOnCancelRevokeNil(b bool)`

 SetOnCancelRevokeNil sets the value for OnCancelRevoke to be an explicit nil

### UnsetOnCancelRevoke
`func (o *CertificateProfileTriggers) UnsetOnCancelRevoke()`

UnsetOnCancelRevoke ensures that no value is present for OnCancelRevoke, not even an explicit nil
### GetOnApproveRevoke

`func (o *CertificateProfileTriggers) GetOnApproveRevoke() []string`

GetOnApproveRevoke returns the OnApproveRevoke field if non-nil, zero value otherwise.

### GetOnApproveRevokeOk

`func (o *CertificateProfileTriggers) GetOnApproveRevokeOk() (*[]string, bool)`

GetOnApproveRevokeOk returns a tuple with the OnApproveRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnApproveRevoke

`func (o *CertificateProfileTriggers) SetOnApproveRevoke(v []string)`

SetOnApproveRevoke sets OnApproveRevoke field to given value.

### HasOnApproveRevoke

`func (o *CertificateProfileTriggers) HasOnApproveRevoke() bool`

HasOnApproveRevoke returns a boolean if a field has been set.

### SetOnApproveRevokeNil

`func (o *CertificateProfileTriggers) SetOnApproveRevokeNil(b bool)`

 SetOnApproveRevokeNil sets the value for OnApproveRevoke to be an explicit nil

### UnsetOnApproveRevoke
`func (o *CertificateProfileTriggers) UnsetOnApproveRevoke()`

UnsetOnApproveRevoke ensures that no value is present for OnApproveRevoke, not even an explicit nil
### GetOnDenyRevoke

`func (o *CertificateProfileTriggers) GetOnDenyRevoke() []string`

GetOnDenyRevoke returns the OnDenyRevoke field if non-nil, zero value otherwise.

### GetOnDenyRevokeOk

`func (o *CertificateProfileTriggers) GetOnDenyRevokeOk() (*[]string, bool)`

GetOnDenyRevokeOk returns a tuple with the OnDenyRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDenyRevoke

`func (o *CertificateProfileTriggers) SetOnDenyRevoke(v []string)`

SetOnDenyRevoke sets OnDenyRevoke field to given value.

### HasOnDenyRevoke

`func (o *CertificateProfileTriggers) HasOnDenyRevoke() bool`

HasOnDenyRevoke returns a boolean if a field has been set.

### SetOnDenyRevokeNil

`func (o *CertificateProfileTriggers) SetOnDenyRevokeNil(b bool)`

 SetOnDenyRevokeNil sets the value for OnDenyRevoke to be an explicit nil

### UnsetOnDenyRevoke
`func (o *CertificateProfileTriggers) UnsetOnDenyRevoke()`

UnsetOnDenyRevoke ensures that no value is present for OnDenyRevoke, not even an explicit nil
### GetOnPendingRevoke

`func (o *CertificateProfileTriggers) GetOnPendingRevoke() []CertificateProfileAsynchronousTrigger`

GetOnPendingRevoke returns the OnPendingRevoke field if non-nil, zero value otherwise.

### GetOnPendingRevokeOk

`func (o *CertificateProfileTriggers) GetOnPendingRevokeOk() (*[]CertificateProfileAsynchronousTrigger, bool)`

GetOnPendingRevokeOk returns a tuple with the OnPendingRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPendingRevoke

`func (o *CertificateProfileTriggers) SetOnPendingRevoke(v []CertificateProfileAsynchronousTrigger)`

SetOnPendingRevoke sets OnPendingRevoke field to given value.

### HasOnPendingRevoke

`func (o *CertificateProfileTriggers) HasOnPendingRevoke() bool`

HasOnPendingRevoke returns a boolean if a field has been set.

### SetOnPendingRevokeNil

`func (o *CertificateProfileTriggers) SetOnPendingRevokeNil(b bool)`

 SetOnPendingRevokeNil sets the value for OnPendingRevoke to be an explicit nil

### UnsetOnPendingRevoke
`func (o *CertificateProfileTriggers) UnsetOnPendingRevoke()`

UnsetOnPendingRevoke ensures that no value is present for OnPendingRevoke, not even an explicit nil
### GetOnUpdate

`func (o *CertificateProfileTriggers) GetOnUpdate() []string`

GetOnUpdate returns the OnUpdate field if non-nil, zero value otherwise.

### GetOnUpdateOk

`func (o *CertificateProfileTriggers) GetOnUpdateOk() (*[]string, bool)`

GetOnUpdateOk returns a tuple with the OnUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnUpdate

`func (o *CertificateProfileTriggers) SetOnUpdate(v []string)`

SetOnUpdate sets OnUpdate field to given value.

### HasOnUpdate

`func (o *CertificateProfileTriggers) HasOnUpdate() bool`

HasOnUpdate returns a boolean if a field has been set.

### SetOnUpdateNil

`func (o *CertificateProfileTriggers) SetOnUpdateNil(b bool)`

 SetOnUpdateNil sets the value for OnUpdate to be an explicit nil

### UnsetOnUpdate
`func (o *CertificateProfileTriggers) UnsetOnUpdate()`

UnsetOnUpdate ensures that no value is present for OnUpdate, not even an explicit nil
### GetOnSubmitUpdate

`func (o *CertificateProfileTriggers) GetOnSubmitUpdate() []string`

GetOnSubmitUpdate returns the OnSubmitUpdate field if non-nil, zero value otherwise.

### GetOnSubmitUpdateOk

`func (o *CertificateProfileTriggers) GetOnSubmitUpdateOk() (*[]string, bool)`

GetOnSubmitUpdateOk returns a tuple with the OnSubmitUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSubmitUpdate

`func (o *CertificateProfileTriggers) SetOnSubmitUpdate(v []string)`

SetOnSubmitUpdate sets OnSubmitUpdate field to given value.

### HasOnSubmitUpdate

`func (o *CertificateProfileTriggers) HasOnSubmitUpdate() bool`

HasOnSubmitUpdate returns a boolean if a field has been set.

### SetOnSubmitUpdateNil

`func (o *CertificateProfileTriggers) SetOnSubmitUpdateNil(b bool)`

 SetOnSubmitUpdateNil sets the value for OnSubmitUpdate to be an explicit nil

### UnsetOnSubmitUpdate
`func (o *CertificateProfileTriggers) UnsetOnSubmitUpdate()`

UnsetOnSubmitUpdate ensures that no value is present for OnSubmitUpdate, not even an explicit nil
### GetOnCancelUpdate

`func (o *CertificateProfileTriggers) GetOnCancelUpdate() []string`

GetOnCancelUpdate returns the OnCancelUpdate field if non-nil, zero value otherwise.

### GetOnCancelUpdateOk

`func (o *CertificateProfileTriggers) GetOnCancelUpdateOk() (*[]string, bool)`

GetOnCancelUpdateOk returns a tuple with the OnCancelUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCancelUpdate

`func (o *CertificateProfileTriggers) SetOnCancelUpdate(v []string)`

SetOnCancelUpdate sets OnCancelUpdate field to given value.

### HasOnCancelUpdate

`func (o *CertificateProfileTriggers) HasOnCancelUpdate() bool`

HasOnCancelUpdate returns a boolean if a field has been set.

### SetOnCancelUpdateNil

`func (o *CertificateProfileTriggers) SetOnCancelUpdateNil(b bool)`

 SetOnCancelUpdateNil sets the value for OnCancelUpdate to be an explicit nil

### UnsetOnCancelUpdate
`func (o *CertificateProfileTriggers) UnsetOnCancelUpdate()`

UnsetOnCancelUpdate ensures that no value is present for OnCancelUpdate, not even an explicit nil
### GetOnApproveUpdate

`func (o *CertificateProfileTriggers) GetOnApproveUpdate() []string`

GetOnApproveUpdate returns the OnApproveUpdate field if non-nil, zero value otherwise.

### GetOnApproveUpdateOk

`func (o *CertificateProfileTriggers) GetOnApproveUpdateOk() (*[]string, bool)`

GetOnApproveUpdateOk returns a tuple with the OnApproveUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnApproveUpdate

`func (o *CertificateProfileTriggers) SetOnApproveUpdate(v []string)`

SetOnApproveUpdate sets OnApproveUpdate field to given value.

### HasOnApproveUpdate

`func (o *CertificateProfileTriggers) HasOnApproveUpdate() bool`

HasOnApproveUpdate returns a boolean if a field has been set.

### SetOnApproveUpdateNil

`func (o *CertificateProfileTriggers) SetOnApproveUpdateNil(b bool)`

 SetOnApproveUpdateNil sets the value for OnApproveUpdate to be an explicit nil

### UnsetOnApproveUpdate
`func (o *CertificateProfileTriggers) UnsetOnApproveUpdate()`

UnsetOnApproveUpdate ensures that no value is present for OnApproveUpdate, not even an explicit nil
### GetOnDenyUpdate

`func (o *CertificateProfileTriggers) GetOnDenyUpdate() []string`

GetOnDenyUpdate returns the OnDenyUpdate field if non-nil, zero value otherwise.

### GetOnDenyUpdateOk

`func (o *CertificateProfileTriggers) GetOnDenyUpdateOk() (*[]string, bool)`

GetOnDenyUpdateOk returns a tuple with the OnDenyUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDenyUpdate

`func (o *CertificateProfileTriggers) SetOnDenyUpdate(v []string)`

SetOnDenyUpdate sets OnDenyUpdate field to given value.

### HasOnDenyUpdate

`func (o *CertificateProfileTriggers) HasOnDenyUpdate() bool`

HasOnDenyUpdate returns a boolean if a field has been set.

### SetOnDenyUpdateNil

`func (o *CertificateProfileTriggers) SetOnDenyUpdateNil(b bool)`

 SetOnDenyUpdateNil sets the value for OnDenyUpdate to be an explicit nil

### UnsetOnDenyUpdate
`func (o *CertificateProfileTriggers) UnsetOnDenyUpdate()`

UnsetOnDenyUpdate ensures that no value is present for OnDenyUpdate, not even an explicit nil
### GetOnPendingUpdate

`func (o *CertificateProfileTriggers) GetOnPendingUpdate() []CertificateProfileAsynchronousTrigger`

GetOnPendingUpdate returns the OnPendingUpdate field if non-nil, zero value otherwise.

### GetOnPendingUpdateOk

`func (o *CertificateProfileTriggers) GetOnPendingUpdateOk() (*[]CertificateProfileAsynchronousTrigger, bool)`

GetOnPendingUpdateOk returns a tuple with the OnPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPendingUpdate

`func (o *CertificateProfileTriggers) SetOnPendingUpdate(v []CertificateProfileAsynchronousTrigger)`

SetOnPendingUpdate sets OnPendingUpdate field to given value.

### HasOnPendingUpdate

`func (o *CertificateProfileTriggers) HasOnPendingUpdate() bool`

HasOnPendingUpdate returns a boolean if a field has been set.

### SetOnPendingUpdateNil

`func (o *CertificateProfileTriggers) SetOnPendingUpdateNil(b bool)`

 SetOnPendingUpdateNil sets the value for OnPendingUpdate to be an explicit nil

### UnsetOnPendingUpdate
`func (o *CertificateProfileTriggers) UnsetOnPendingUpdate()`

UnsetOnPendingUpdate ensures that no value is present for OnPendingUpdate, not even an explicit nil
### GetOnRecover

`func (o *CertificateProfileTriggers) GetOnRecover() []string`

GetOnRecover returns the OnRecover field if non-nil, zero value otherwise.

### GetOnRecoverOk

`func (o *CertificateProfileTriggers) GetOnRecoverOk() (*[]string, bool)`

GetOnRecoverOk returns a tuple with the OnRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnRecover

`func (o *CertificateProfileTriggers) SetOnRecover(v []string)`

SetOnRecover sets OnRecover field to given value.

### HasOnRecover

`func (o *CertificateProfileTriggers) HasOnRecover() bool`

HasOnRecover returns a boolean if a field has been set.

### SetOnRecoverNil

`func (o *CertificateProfileTriggers) SetOnRecoverNil(b bool)`

 SetOnRecoverNil sets the value for OnRecover to be an explicit nil

### UnsetOnRecover
`func (o *CertificateProfileTriggers) UnsetOnRecover()`

UnsetOnRecover ensures that no value is present for OnRecover, not even an explicit nil
### GetOnSubmitRecover

`func (o *CertificateProfileTriggers) GetOnSubmitRecover() []string`

GetOnSubmitRecover returns the OnSubmitRecover field if non-nil, zero value otherwise.

### GetOnSubmitRecoverOk

`func (o *CertificateProfileTriggers) GetOnSubmitRecoverOk() (*[]string, bool)`

GetOnSubmitRecoverOk returns a tuple with the OnSubmitRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSubmitRecover

`func (o *CertificateProfileTriggers) SetOnSubmitRecover(v []string)`

SetOnSubmitRecover sets OnSubmitRecover field to given value.

### HasOnSubmitRecover

`func (o *CertificateProfileTriggers) HasOnSubmitRecover() bool`

HasOnSubmitRecover returns a boolean if a field has been set.

### SetOnSubmitRecoverNil

`func (o *CertificateProfileTriggers) SetOnSubmitRecoverNil(b bool)`

 SetOnSubmitRecoverNil sets the value for OnSubmitRecover to be an explicit nil

### UnsetOnSubmitRecover
`func (o *CertificateProfileTriggers) UnsetOnSubmitRecover()`

UnsetOnSubmitRecover ensures that no value is present for OnSubmitRecover, not even an explicit nil
### GetOnCancelRecover

`func (o *CertificateProfileTriggers) GetOnCancelRecover() []string`

GetOnCancelRecover returns the OnCancelRecover field if non-nil, zero value otherwise.

### GetOnCancelRecoverOk

`func (o *CertificateProfileTriggers) GetOnCancelRecoverOk() (*[]string, bool)`

GetOnCancelRecoverOk returns a tuple with the OnCancelRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCancelRecover

`func (o *CertificateProfileTriggers) SetOnCancelRecover(v []string)`

SetOnCancelRecover sets OnCancelRecover field to given value.

### HasOnCancelRecover

`func (o *CertificateProfileTriggers) HasOnCancelRecover() bool`

HasOnCancelRecover returns a boolean if a field has been set.

### SetOnCancelRecoverNil

`func (o *CertificateProfileTriggers) SetOnCancelRecoverNil(b bool)`

 SetOnCancelRecoverNil sets the value for OnCancelRecover to be an explicit nil

### UnsetOnCancelRecover
`func (o *CertificateProfileTriggers) UnsetOnCancelRecover()`

UnsetOnCancelRecover ensures that no value is present for OnCancelRecover, not even an explicit nil
### GetOnApproveRecover

`func (o *CertificateProfileTriggers) GetOnApproveRecover() []string`

GetOnApproveRecover returns the OnApproveRecover field if non-nil, zero value otherwise.

### GetOnApproveRecoverOk

`func (o *CertificateProfileTriggers) GetOnApproveRecoverOk() (*[]string, bool)`

GetOnApproveRecoverOk returns a tuple with the OnApproveRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnApproveRecover

`func (o *CertificateProfileTriggers) SetOnApproveRecover(v []string)`

SetOnApproveRecover sets OnApproveRecover field to given value.

### HasOnApproveRecover

`func (o *CertificateProfileTriggers) HasOnApproveRecover() bool`

HasOnApproveRecover returns a boolean if a field has been set.

### SetOnApproveRecoverNil

`func (o *CertificateProfileTriggers) SetOnApproveRecoverNil(b bool)`

 SetOnApproveRecoverNil sets the value for OnApproveRecover to be an explicit nil

### UnsetOnApproveRecover
`func (o *CertificateProfileTriggers) UnsetOnApproveRecover()`

UnsetOnApproveRecover ensures that no value is present for OnApproveRecover, not even an explicit nil
### GetOnDenyRecover

`func (o *CertificateProfileTriggers) GetOnDenyRecover() []string`

GetOnDenyRecover returns the OnDenyRecover field if non-nil, zero value otherwise.

### GetOnDenyRecoverOk

`func (o *CertificateProfileTriggers) GetOnDenyRecoverOk() (*[]string, bool)`

GetOnDenyRecoverOk returns a tuple with the OnDenyRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDenyRecover

`func (o *CertificateProfileTriggers) SetOnDenyRecover(v []string)`

SetOnDenyRecover sets OnDenyRecover field to given value.

### HasOnDenyRecover

`func (o *CertificateProfileTriggers) HasOnDenyRecover() bool`

HasOnDenyRecover returns a boolean if a field has been set.

### SetOnDenyRecoverNil

`func (o *CertificateProfileTriggers) SetOnDenyRecoverNil(b bool)`

 SetOnDenyRecoverNil sets the value for OnDenyRecover to be an explicit nil

### UnsetOnDenyRecover
`func (o *CertificateProfileTriggers) UnsetOnDenyRecover()`

UnsetOnDenyRecover ensures that no value is present for OnDenyRecover, not even an explicit nil
### GetOnPendingRecover

`func (o *CertificateProfileTriggers) GetOnPendingRecover() []CertificateProfileAsynchronousTrigger`

GetOnPendingRecover returns the OnPendingRecover field if non-nil, zero value otherwise.

### GetOnPendingRecoverOk

`func (o *CertificateProfileTriggers) GetOnPendingRecoverOk() (*[]CertificateProfileAsynchronousTrigger, bool)`

GetOnPendingRecoverOk returns a tuple with the OnPendingRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPendingRecover

`func (o *CertificateProfileTriggers) SetOnPendingRecover(v []CertificateProfileAsynchronousTrigger)`

SetOnPendingRecover sets OnPendingRecover field to given value.

### HasOnPendingRecover

`func (o *CertificateProfileTriggers) HasOnPendingRecover() bool`

HasOnPendingRecover returns a boolean if a field has been set.

### SetOnPendingRecoverNil

`func (o *CertificateProfileTriggers) SetOnPendingRecoverNil(b bool)`

 SetOnPendingRecoverNil sets the value for OnPendingRecover to be an explicit nil

### UnsetOnPendingRecover
`func (o *CertificateProfileTriggers) UnsetOnPendingRecover()`

UnsetOnPendingRecover ensures that no value is present for OnPendingRecover, not even an explicit nil
### GetOnMigrate

`func (o *CertificateProfileTriggers) GetOnMigrate() []string`

GetOnMigrate returns the OnMigrate field if non-nil, zero value otherwise.

### GetOnMigrateOk

`func (o *CertificateProfileTriggers) GetOnMigrateOk() (*[]string, bool)`

GetOnMigrateOk returns a tuple with the OnMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnMigrate

`func (o *CertificateProfileTriggers) SetOnMigrate(v []string)`

SetOnMigrate sets OnMigrate field to given value.

### HasOnMigrate

`func (o *CertificateProfileTriggers) HasOnMigrate() bool`

HasOnMigrate returns a boolean if a field has been set.

### SetOnMigrateNil

`func (o *CertificateProfileTriggers) SetOnMigrateNil(b bool)`

 SetOnMigrateNil sets the value for OnMigrate to be an explicit nil

### UnsetOnMigrate
`func (o *CertificateProfileTriggers) UnsetOnMigrate()`

UnsetOnMigrate ensures that no value is present for OnMigrate, not even an explicit nil
### GetOnSubmitMigrate

`func (o *CertificateProfileTriggers) GetOnSubmitMigrate() []string`

GetOnSubmitMigrate returns the OnSubmitMigrate field if non-nil, zero value otherwise.

### GetOnSubmitMigrateOk

`func (o *CertificateProfileTriggers) GetOnSubmitMigrateOk() (*[]string, bool)`

GetOnSubmitMigrateOk returns a tuple with the OnSubmitMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSubmitMigrate

`func (o *CertificateProfileTriggers) SetOnSubmitMigrate(v []string)`

SetOnSubmitMigrate sets OnSubmitMigrate field to given value.

### HasOnSubmitMigrate

`func (o *CertificateProfileTriggers) HasOnSubmitMigrate() bool`

HasOnSubmitMigrate returns a boolean if a field has been set.

### SetOnSubmitMigrateNil

`func (o *CertificateProfileTriggers) SetOnSubmitMigrateNil(b bool)`

 SetOnSubmitMigrateNil sets the value for OnSubmitMigrate to be an explicit nil

### UnsetOnSubmitMigrate
`func (o *CertificateProfileTriggers) UnsetOnSubmitMigrate()`

UnsetOnSubmitMigrate ensures that no value is present for OnSubmitMigrate, not even an explicit nil
### GetOnCancelMigrate

`func (o *CertificateProfileTriggers) GetOnCancelMigrate() []string`

GetOnCancelMigrate returns the OnCancelMigrate field if non-nil, zero value otherwise.

### GetOnCancelMigrateOk

`func (o *CertificateProfileTriggers) GetOnCancelMigrateOk() (*[]string, bool)`

GetOnCancelMigrateOk returns a tuple with the OnCancelMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCancelMigrate

`func (o *CertificateProfileTriggers) SetOnCancelMigrate(v []string)`

SetOnCancelMigrate sets OnCancelMigrate field to given value.

### HasOnCancelMigrate

`func (o *CertificateProfileTriggers) HasOnCancelMigrate() bool`

HasOnCancelMigrate returns a boolean if a field has been set.

### SetOnCancelMigrateNil

`func (o *CertificateProfileTriggers) SetOnCancelMigrateNil(b bool)`

 SetOnCancelMigrateNil sets the value for OnCancelMigrate to be an explicit nil

### UnsetOnCancelMigrate
`func (o *CertificateProfileTriggers) UnsetOnCancelMigrate()`

UnsetOnCancelMigrate ensures that no value is present for OnCancelMigrate, not even an explicit nil
### GetOnApproveMigrate

`func (o *CertificateProfileTriggers) GetOnApproveMigrate() []string`

GetOnApproveMigrate returns the OnApproveMigrate field if non-nil, zero value otherwise.

### GetOnApproveMigrateOk

`func (o *CertificateProfileTriggers) GetOnApproveMigrateOk() (*[]string, bool)`

GetOnApproveMigrateOk returns a tuple with the OnApproveMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnApproveMigrate

`func (o *CertificateProfileTriggers) SetOnApproveMigrate(v []string)`

SetOnApproveMigrate sets OnApproveMigrate field to given value.

### HasOnApproveMigrate

`func (o *CertificateProfileTriggers) HasOnApproveMigrate() bool`

HasOnApproveMigrate returns a boolean if a field has been set.

### SetOnApproveMigrateNil

`func (o *CertificateProfileTriggers) SetOnApproveMigrateNil(b bool)`

 SetOnApproveMigrateNil sets the value for OnApproveMigrate to be an explicit nil

### UnsetOnApproveMigrate
`func (o *CertificateProfileTriggers) UnsetOnApproveMigrate()`

UnsetOnApproveMigrate ensures that no value is present for OnApproveMigrate, not even an explicit nil
### GetOnDenyMigrate

`func (o *CertificateProfileTriggers) GetOnDenyMigrate() []string`

GetOnDenyMigrate returns the OnDenyMigrate field if non-nil, zero value otherwise.

### GetOnDenyMigrateOk

`func (o *CertificateProfileTriggers) GetOnDenyMigrateOk() (*[]string, bool)`

GetOnDenyMigrateOk returns a tuple with the OnDenyMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDenyMigrate

`func (o *CertificateProfileTriggers) SetOnDenyMigrate(v []string)`

SetOnDenyMigrate sets OnDenyMigrate field to given value.

### HasOnDenyMigrate

`func (o *CertificateProfileTriggers) HasOnDenyMigrate() bool`

HasOnDenyMigrate returns a boolean if a field has been set.

### SetOnDenyMigrateNil

`func (o *CertificateProfileTriggers) SetOnDenyMigrateNil(b bool)`

 SetOnDenyMigrateNil sets the value for OnDenyMigrate to be an explicit nil

### UnsetOnDenyMigrate
`func (o *CertificateProfileTriggers) UnsetOnDenyMigrate()`

UnsetOnDenyMigrate ensures that no value is present for OnDenyMigrate, not even an explicit nil
### GetOnPendingMigrate

`func (o *CertificateProfileTriggers) GetOnPendingMigrate() []CertificateProfileAsynchronousTrigger`

GetOnPendingMigrate returns the OnPendingMigrate field if non-nil, zero value otherwise.

### GetOnPendingMigrateOk

`func (o *CertificateProfileTriggers) GetOnPendingMigrateOk() (*[]CertificateProfileAsynchronousTrigger, bool)`

GetOnPendingMigrateOk returns a tuple with the OnPendingMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPendingMigrate

`func (o *CertificateProfileTriggers) SetOnPendingMigrate(v []CertificateProfileAsynchronousTrigger)`

SetOnPendingMigrate sets OnPendingMigrate field to given value.

### HasOnPendingMigrate

`func (o *CertificateProfileTriggers) HasOnPendingMigrate() bool`

HasOnPendingMigrate returns a boolean if a field has been set.

### SetOnPendingMigrateNil

`func (o *CertificateProfileTriggers) SetOnPendingMigrateNil(b bool)`

 SetOnPendingMigrateNil sets the value for OnPendingMigrate to be an explicit nil

### UnsetOnPendingMigrate
`func (o *CertificateProfileTriggers) UnsetOnPendingMigrate()`

UnsetOnPendingMigrate ensures that no value is present for OnPendingMigrate, not even an explicit nil
### GetOnExpire

`func (o *CertificateProfileTriggers) GetOnExpire() []CertificateProfileAsynchronousTrigger`

GetOnExpire returns the OnExpire field if non-nil, zero value otherwise.

### GetOnExpireOk

`func (o *CertificateProfileTriggers) GetOnExpireOk() (*[]CertificateProfileAsynchronousTrigger, bool)`

GetOnExpireOk returns a tuple with the OnExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnExpire

`func (o *CertificateProfileTriggers) SetOnExpire(v []CertificateProfileAsynchronousTrigger)`

SetOnExpire sets OnExpire field to given value.

### HasOnExpire

`func (o *CertificateProfileTriggers) HasOnExpire() bool`

HasOnExpire returns a boolean if a field has been set.

### SetOnExpireNil

`func (o *CertificateProfileTriggers) SetOnExpireNil(b bool)`

 SetOnExpireNil sets the value for OnExpire to be an explicit nil

### UnsetOnExpire
`func (o *CertificateProfileTriggers) UnsetOnExpire()`

UnsetOnExpire ensures that no value is present for OnExpire, not even an explicit nil
### GetOnRenew

`func (o *CertificateProfileTriggers) GetOnRenew() []string`

GetOnRenew returns the OnRenew field if non-nil, zero value otherwise.

### GetOnRenewOk

`func (o *CertificateProfileTriggers) GetOnRenewOk() (*[]string, bool)`

GetOnRenewOk returns a tuple with the OnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnRenew

`func (o *CertificateProfileTriggers) SetOnRenew(v []string)`

SetOnRenew sets OnRenew field to given value.

### HasOnRenew

`func (o *CertificateProfileTriggers) HasOnRenew() bool`

HasOnRenew returns a boolean if a field has been set.

### SetOnRenewNil

`func (o *CertificateProfileTriggers) SetOnRenewNil(b bool)`

 SetOnRenewNil sets the value for OnRenew to be an explicit nil

### UnsetOnRenew
`func (o *CertificateProfileTriggers) UnsetOnRenew()`

UnsetOnRenew ensures that no value is present for OnRenew, not even an explicit nil
### GetOnSubmitRenew

`func (o *CertificateProfileTriggers) GetOnSubmitRenew() []string`

GetOnSubmitRenew returns the OnSubmitRenew field if non-nil, zero value otherwise.

### GetOnSubmitRenewOk

`func (o *CertificateProfileTriggers) GetOnSubmitRenewOk() (*[]string, bool)`

GetOnSubmitRenewOk returns a tuple with the OnSubmitRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSubmitRenew

`func (o *CertificateProfileTriggers) SetOnSubmitRenew(v []string)`

SetOnSubmitRenew sets OnSubmitRenew field to given value.

### HasOnSubmitRenew

`func (o *CertificateProfileTriggers) HasOnSubmitRenew() bool`

HasOnSubmitRenew returns a boolean if a field has been set.

### SetOnSubmitRenewNil

`func (o *CertificateProfileTriggers) SetOnSubmitRenewNil(b bool)`

 SetOnSubmitRenewNil sets the value for OnSubmitRenew to be an explicit nil

### UnsetOnSubmitRenew
`func (o *CertificateProfileTriggers) UnsetOnSubmitRenew()`

UnsetOnSubmitRenew ensures that no value is present for OnSubmitRenew, not even an explicit nil
### GetOnCancelRenew

`func (o *CertificateProfileTriggers) GetOnCancelRenew() []string`

GetOnCancelRenew returns the OnCancelRenew field if non-nil, zero value otherwise.

### GetOnCancelRenewOk

`func (o *CertificateProfileTriggers) GetOnCancelRenewOk() (*[]string, bool)`

GetOnCancelRenewOk returns a tuple with the OnCancelRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCancelRenew

`func (o *CertificateProfileTriggers) SetOnCancelRenew(v []string)`

SetOnCancelRenew sets OnCancelRenew field to given value.

### HasOnCancelRenew

`func (o *CertificateProfileTriggers) HasOnCancelRenew() bool`

HasOnCancelRenew returns a boolean if a field has been set.

### SetOnCancelRenewNil

`func (o *CertificateProfileTriggers) SetOnCancelRenewNil(b bool)`

 SetOnCancelRenewNil sets the value for OnCancelRenew to be an explicit nil

### UnsetOnCancelRenew
`func (o *CertificateProfileTriggers) UnsetOnCancelRenew()`

UnsetOnCancelRenew ensures that no value is present for OnCancelRenew, not even an explicit nil
### GetOnApproveRenew

`func (o *CertificateProfileTriggers) GetOnApproveRenew() []string`

GetOnApproveRenew returns the OnApproveRenew field if non-nil, zero value otherwise.

### GetOnApproveRenewOk

`func (o *CertificateProfileTriggers) GetOnApproveRenewOk() (*[]string, bool)`

GetOnApproveRenewOk returns a tuple with the OnApproveRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnApproveRenew

`func (o *CertificateProfileTriggers) SetOnApproveRenew(v []string)`

SetOnApproveRenew sets OnApproveRenew field to given value.

### HasOnApproveRenew

`func (o *CertificateProfileTriggers) HasOnApproveRenew() bool`

HasOnApproveRenew returns a boolean if a field has been set.

### SetOnApproveRenewNil

`func (o *CertificateProfileTriggers) SetOnApproveRenewNil(b bool)`

 SetOnApproveRenewNil sets the value for OnApproveRenew to be an explicit nil

### UnsetOnApproveRenew
`func (o *CertificateProfileTriggers) UnsetOnApproveRenew()`

UnsetOnApproveRenew ensures that no value is present for OnApproveRenew, not even an explicit nil
### GetOnDenyRenew

`func (o *CertificateProfileTriggers) GetOnDenyRenew() []string`

GetOnDenyRenew returns the OnDenyRenew field if non-nil, zero value otherwise.

### GetOnDenyRenewOk

`func (o *CertificateProfileTriggers) GetOnDenyRenewOk() (*[]string, bool)`

GetOnDenyRenewOk returns a tuple with the OnDenyRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDenyRenew

`func (o *CertificateProfileTriggers) SetOnDenyRenew(v []string)`

SetOnDenyRenew sets OnDenyRenew field to given value.

### HasOnDenyRenew

`func (o *CertificateProfileTriggers) HasOnDenyRenew() bool`

HasOnDenyRenew returns a boolean if a field has been set.

### SetOnDenyRenewNil

`func (o *CertificateProfileTriggers) SetOnDenyRenewNil(b bool)`

 SetOnDenyRenewNil sets the value for OnDenyRenew to be an explicit nil

### UnsetOnDenyRenew
`func (o *CertificateProfileTriggers) UnsetOnDenyRenew()`

UnsetOnDenyRenew ensures that no value is present for OnDenyRenew, not even an explicit nil
### GetOnPendingRenew

`func (o *CertificateProfileTriggers) GetOnPendingRenew() []CertificateProfileAsynchronousTrigger`

GetOnPendingRenew returns the OnPendingRenew field if non-nil, zero value otherwise.

### GetOnPendingRenewOk

`func (o *CertificateProfileTriggers) GetOnPendingRenewOk() (*[]CertificateProfileAsynchronousTrigger, bool)`

GetOnPendingRenewOk returns a tuple with the OnPendingRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPendingRenew

`func (o *CertificateProfileTriggers) SetOnPendingRenew(v []CertificateProfileAsynchronousTrigger)`

SetOnPendingRenew sets OnPendingRenew field to given value.

### HasOnPendingRenew

`func (o *CertificateProfileTriggers) HasOnPendingRenew() bool`

HasOnPendingRenew returns a boolean if a field has been set.

### SetOnPendingRenewNil

`func (o *CertificateProfileTriggers) SetOnPendingRenewNil(b bool)`

 SetOnPendingRenewNil sets the value for OnPendingRenew to be an explicit nil

### UnsetOnPendingRenew
`func (o *CertificateProfileTriggers) UnsetOnPendingRenew()`

UnsetOnPendingRenew ensures that no value is present for OnPendingRenew, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


