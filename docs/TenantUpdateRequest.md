# TenantUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A simple description for this tenant | [optional] 
**LicenseExpiration** | Pointer to **int64** | Custom license expiration date for this tenant | [optional] 
**LicenseLimit** | **int64** | License limit for this tenant | 
**Name** | **string** | The tenant internal name | 

## Methods

### NewTenantUpdateRequest

`func NewTenantUpdateRequest(licenseLimit int64, name string, ) *TenantUpdateRequest`

NewTenantUpdateRequest instantiates a new TenantUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantUpdateRequestWithDefaults

`func NewTenantUpdateRequestWithDefaults() *TenantUpdateRequest`

NewTenantUpdateRequestWithDefaults instantiates a new TenantUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TenantUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicenseExpiration

`func (o *TenantUpdateRequest) GetLicenseExpiration() int64`

GetLicenseExpiration returns the LicenseExpiration field if non-nil, zero value otherwise.

### GetLicenseExpirationOk

`func (o *TenantUpdateRequest) GetLicenseExpirationOk() (*int64, bool)`

GetLicenseExpirationOk returns a tuple with the LicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiration

`func (o *TenantUpdateRequest) SetLicenseExpiration(v int64)`

SetLicenseExpiration sets LicenseExpiration field to given value.

### HasLicenseExpiration

`func (o *TenantUpdateRequest) HasLicenseExpiration() bool`

HasLicenseExpiration returns a boolean if a field has been set.

### GetLicenseLimit

`func (o *TenantUpdateRequest) GetLicenseLimit() int64`

GetLicenseLimit returns the LicenseLimit field if non-nil, zero value otherwise.

### GetLicenseLimitOk

`func (o *TenantUpdateRequest) GetLicenseLimitOk() (*int64, bool)`

GetLicenseLimitOk returns a tuple with the LicenseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseLimit

`func (o *TenantUpdateRequest) SetLicenseLimit(v int64)`

SetLicenseLimit sets LicenseLimit field to given value.


### GetName

`func (o *TenantUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantUpdateRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


