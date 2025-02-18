# CertificatePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enroll** | Pointer to **NullableBool** | Whether the principal is authorized to re-enroll this certificate | [optional] 
**RequestEnroll** | Pointer to **NullableBool** | Whether the principal is authorized to request re-enrollment of this certificate | [optional] 
**Revoke** | **bool** | Whether the principal is authorized to revoke this certificate | 
**RequestRevoke** | **bool** | Whether the principal is authorized to request revocation of this certificate | 
**Update** | **bool** | Whether the principal is authorized to update this certificate | 
**RequestUpdate** | **bool** | Whether the principal is authorized to request update of this certificate | 
**Recover** | Pointer to **NullableBool** | Whether the principal is authorized to recover this certificate | [optional] 
**RequestRecover** | Pointer to **NullableBool** | Whether the principal is authorized to request recovery of this certificate | [optional] 
**Migrate** | **bool** | Whether the principal is authorized to migrate this certificate | 
**RequestMigrate** | **bool** | Whether the principal is authorized to request migration of this certificate | 
**Renew** | Pointer to **NullableBool** | Whether the principal is authorized to renew this certificate | [optional] 
**RequestRenew** | Pointer to **NullableBool** | Whether the principal is authorized to request renewal of this certificate | [optional] 

## Methods

### NewCertificatePermissions

`func NewCertificatePermissions(revoke bool, requestRevoke bool, update bool, requestUpdate bool, migrate bool, requestMigrate bool, ) *CertificatePermissions`

NewCertificatePermissions instantiates a new CertificatePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatePermissionsWithDefaults

`func NewCertificatePermissionsWithDefaults() *CertificatePermissions`

NewCertificatePermissionsWithDefaults instantiates a new CertificatePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnroll

`func (o *CertificatePermissions) GetEnroll() bool`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *CertificatePermissions) GetEnrollOk() (*bool, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *CertificatePermissions) SetEnroll(v bool)`

SetEnroll sets Enroll field to given value.

### HasEnroll

`func (o *CertificatePermissions) HasEnroll() bool`

HasEnroll returns a boolean if a field has been set.

### SetEnrollNil

`func (o *CertificatePermissions) SetEnrollNil(b bool)`

 SetEnrollNil sets the value for Enroll to be an explicit nil

### UnsetEnroll
`func (o *CertificatePermissions) UnsetEnroll()`

UnsetEnroll ensures that no value is present for Enroll, not even an explicit nil
### GetRequestEnroll

`func (o *CertificatePermissions) GetRequestEnroll() bool`

GetRequestEnroll returns the RequestEnroll field if non-nil, zero value otherwise.

### GetRequestEnrollOk

`func (o *CertificatePermissions) GetRequestEnrollOk() (*bool, bool)`

GetRequestEnrollOk returns a tuple with the RequestEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEnroll

`func (o *CertificatePermissions) SetRequestEnroll(v bool)`

SetRequestEnroll sets RequestEnroll field to given value.

### HasRequestEnroll

`func (o *CertificatePermissions) HasRequestEnroll() bool`

HasRequestEnroll returns a boolean if a field has been set.

### SetRequestEnrollNil

`func (o *CertificatePermissions) SetRequestEnrollNil(b bool)`

 SetRequestEnrollNil sets the value for RequestEnroll to be an explicit nil

### UnsetRequestEnroll
`func (o *CertificatePermissions) UnsetRequestEnroll()`

UnsetRequestEnroll ensures that no value is present for RequestEnroll, not even an explicit nil
### GetRevoke

`func (o *CertificatePermissions) GetRevoke() bool`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *CertificatePermissions) GetRevokeOk() (*bool, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *CertificatePermissions) SetRevoke(v bool)`

SetRevoke sets Revoke field to given value.


### GetRequestRevoke

`func (o *CertificatePermissions) GetRequestRevoke() bool`

GetRequestRevoke returns the RequestRevoke field if non-nil, zero value otherwise.

### GetRequestRevokeOk

`func (o *CertificatePermissions) GetRequestRevokeOk() (*bool, bool)`

GetRequestRevokeOk returns a tuple with the RequestRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRevoke

`func (o *CertificatePermissions) SetRequestRevoke(v bool)`

SetRequestRevoke sets RequestRevoke field to given value.


### GetUpdate

`func (o *CertificatePermissions) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *CertificatePermissions) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *CertificatePermissions) SetUpdate(v bool)`

SetUpdate sets Update field to given value.


### GetRequestUpdate

`func (o *CertificatePermissions) GetRequestUpdate() bool`

GetRequestUpdate returns the RequestUpdate field if non-nil, zero value otherwise.

### GetRequestUpdateOk

`func (o *CertificatePermissions) GetRequestUpdateOk() (*bool, bool)`

GetRequestUpdateOk returns a tuple with the RequestUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUpdate

`func (o *CertificatePermissions) SetRequestUpdate(v bool)`

SetRequestUpdate sets RequestUpdate field to given value.


### GetRecover

`func (o *CertificatePermissions) GetRecover() bool`

GetRecover returns the Recover field if non-nil, zero value otherwise.

### GetRecoverOk

`func (o *CertificatePermissions) GetRecoverOk() (*bool, bool)`

GetRecoverOk returns a tuple with the Recover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecover

`func (o *CertificatePermissions) SetRecover(v bool)`

SetRecover sets Recover field to given value.

### HasRecover

`func (o *CertificatePermissions) HasRecover() bool`

HasRecover returns a boolean if a field has been set.

### SetRecoverNil

`func (o *CertificatePermissions) SetRecoverNil(b bool)`

 SetRecoverNil sets the value for Recover to be an explicit nil

### UnsetRecover
`func (o *CertificatePermissions) UnsetRecover()`

UnsetRecover ensures that no value is present for Recover, not even an explicit nil
### GetRequestRecover

`func (o *CertificatePermissions) GetRequestRecover() bool`

GetRequestRecover returns the RequestRecover field if non-nil, zero value otherwise.

### GetRequestRecoverOk

`func (o *CertificatePermissions) GetRequestRecoverOk() (*bool, bool)`

GetRequestRecoverOk returns a tuple with the RequestRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRecover

`func (o *CertificatePermissions) SetRequestRecover(v bool)`

SetRequestRecover sets RequestRecover field to given value.

### HasRequestRecover

`func (o *CertificatePermissions) HasRequestRecover() bool`

HasRequestRecover returns a boolean if a field has been set.

### SetRequestRecoverNil

`func (o *CertificatePermissions) SetRequestRecoverNil(b bool)`

 SetRequestRecoverNil sets the value for RequestRecover to be an explicit nil

### UnsetRequestRecover
`func (o *CertificatePermissions) UnsetRequestRecover()`

UnsetRequestRecover ensures that no value is present for RequestRecover, not even an explicit nil
### GetMigrate

`func (o *CertificatePermissions) GetMigrate() bool`

GetMigrate returns the Migrate field if non-nil, zero value otherwise.

### GetMigrateOk

`func (o *CertificatePermissions) GetMigrateOk() (*bool, bool)`

GetMigrateOk returns a tuple with the Migrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrate

`func (o *CertificatePermissions) SetMigrate(v bool)`

SetMigrate sets Migrate field to given value.


### GetRequestMigrate

`func (o *CertificatePermissions) GetRequestMigrate() bool`

GetRequestMigrate returns the RequestMigrate field if non-nil, zero value otherwise.

### GetRequestMigrateOk

`func (o *CertificatePermissions) GetRequestMigrateOk() (*bool, bool)`

GetRequestMigrateOk returns a tuple with the RequestMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMigrate

`func (o *CertificatePermissions) SetRequestMigrate(v bool)`

SetRequestMigrate sets RequestMigrate field to given value.


### GetRenew

`func (o *CertificatePermissions) GetRenew() bool`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *CertificatePermissions) GetRenewOk() (*bool, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *CertificatePermissions) SetRenew(v bool)`

SetRenew sets Renew field to given value.

### HasRenew

`func (o *CertificatePermissions) HasRenew() bool`

HasRenew returns a boolean if a field has been set.

### SetRenewNil

`func (o *CertificatePermissions) SetRenewNil(b bool)`

 SetRenewNil sets the value for Renew to be an explicit nil

### UnsetRenew
`func (o *CertificatePermissions) UnsetRenew()`

UnsetRenew ensures that no value is present for Renew, not even an explicit nil
### GetRequestRenew

`func (o *CertificatePermissions) GetRequestRenew() bool`

GetRequestRenew returns the RequestRenew field if non-nil, zero value otherwise.

### GetRequestRenewOk

`func (o *CertificatePermissions) GetRequestRenewOk() (*bool, bool)`

GetRequestRenewOk returns a tuple with the RequestRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRenew

`func (o *CertificatePermissions) SetRequestRenew(v bool)`

SetRequestRenew sets RequestRenew field to given value.

### HasRequestRenew

`func (o *CertificatePermissions) HasRequestRenew() bool`

HasRequestRenew returns a boolean if a field has been set.

### SetRequestRenewNil

`func (o *CertificatePermissions) SetRequestRenewNil(b bool)`

 SetRequestRenewNil sets the value for RequestRenew to be an explicit nil

### UnsetRequestRenew
`func (o *CertificatePermissions) UnsetRequestRenew()`

UnsetRequestRenew ensures that no value is present for RequestRenew, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


