# TenantLicenseConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**TenantCertificateLicenseConfiguration**](TenantCertificateLicenseConfiguration.md) |  | [optional] 
**Dcv** | Pointer to [**TenantDCVLicenseConfiguration**](TenantDCVLicenseConfiguration.md) |  | [optional] 
**Expiration** | Pointer to **int64** | Custom license expiration date for this tenant | [optional] 

## Methods

### NewTenantLicenseConfiguration

`func NewTenantLicenseConfiguration() *TenantLicenseConfiguration`

NewTenantLicenseConfiguration instantiates a new TenantLicenseConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantLicenseConfigurationWithDefaults

`func NewTenantLicenseConfigurationWithDefaults() *TenantLicenseConfiguration`

NewTenantLicenseConfigurationWithDefaults instantiates a new TenantLicenseConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *TenantLicenseConfiguration) GetCertificate() TenantCertificateLicenseConfiguration`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TenantLicenseConfiguration) GetCertificateOk() (*TenantCertificateLicenseConfiguration, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TenantLicenseConfiguration) SetCertificate(v TenantCertificateLicenseConfiguration)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *TenantLicenseConfiguration) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetDcv

`func (o *TenantLicenseConfiguration) GetDcv() TenantDCVLicenseConfiguration`

GetDcv returns the Dcv field if non-nil, zero value otherwise.

### GetDcvOk

`func (o *TenantLicenseConfiguration) GetDcvOk() (*TenantDCVLicenseConfiguration, bool)`

GetDcvOk returns a tuple with the Dcv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcv

`func (o *TenantLicenseConfiguration) SetDcv(v TenantDCVLicenseConfiguration)`

SetDcv sets Dcv field to given value.

### HasDcv

`func (o *TenantLicenseConfiguration) HasDcv() bool`

HasDcv returns a boolean if a field has been set.

### GetExpiration

`func (o *TenantLicenseConfiguration) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *TenantLicenseConfiguration) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *TenantLicenseConfiguration) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *TenantLicenseConfiguration) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


