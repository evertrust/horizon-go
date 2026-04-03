# TenantCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministratorPassword** | Pointer to **string** | Password to use for the administrator account of this tenant | [optional] 
**Description** | Pointer to **string** | A simple description for this tenant | [optional] 
**LicenseExpiration** | Pointer to **int64** | Custom license expiration date for this tenant | [optional] 
**LicenseLimit** | **int64** | License limit for this tenant | 
**Name** | **string** | The tenant internal name | 

## Methods

### NewTenantCreationRequest

`func NewTenantCreationRequest(licenseLimit int64, name string, ) *TenantCreationRequest`

NewTenantCreationRequest instantiates a new TenantCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCreationRequestWithDefaults

`func NewTenantCreationRequestWithDefaults() *TenantCreationRequest`

NewTenantCreationRequestWithDefaults instantiates a new TenantCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministratorPassword

`func (o *TenantCreationRequest) GetAdministratorPassword() string`

GetAdministratorPassword returns the AdministratorPassword field if non-nil, zero value otherwise.

### GetAdministratorPasswordOk

`func (o *TenantCreationRequest) GetAdministratorPasswordOk() (*string, bool)`

GetAdministratorPasswordOk returns a tuple with the AdministratorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorPassword

`func (o *TenantCreationRequest) SetAdministratorPassword(v string)`

SetAdministratorPassword sets AdministratorPassword field to given value.

### HasAdministratorPassword

`func (o *TenantCreationRequest) HasAdministratorPassword() bool`

HasAdministratorPassword returns a boolean if a field has been set.

### GetDescription

`func (o *TenantCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantCreationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicenseExpiration

`func (o *TenantCreationRequest) GetLicenseExpiration() int64`

GetLicenseExpiration returns the LicenseExpiration field if non-nil, zero value otherwise.

### GetLicenseExpirationOk

`func (o *TenantCreationRequest) GetLicenseExpirationOk() (*int64, bool)`

GetLicenseExpirationOk returns a tuple with the LicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiration

`func (o *TenantCreationRequest) SetLicenseExpiration(v int64)`

SetLicenseExpiration sets LicenseExpiration field to given value.

### HasLicenseExpiration

`func (o *TenantCreationRequest) HasLicenseExpiration() bool`

HasLicenseExpiration returns a boolean if a field has been set.

### GetLicenseLimit

`func (o *TenantCreationRequest) GetLicenseLimit() int64`

GetLicenseLimit returns the LicenseLimit field if non-nil, zero value otherwise.

### GetLicenseLimitOk

`func (o *TenantCreationRequest) GetLicenseLimitOk() (*int64, bool)`

GetLicenseLimitOk returns a tuple with the LicenseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseLimit

`func (o *TenantCreationRequest) SetLicenseLimit(v int64)`

SetLicenseLimit sets LicenseLimit field to given value.


### GetName

`func (o *TenantCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantCreationRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


