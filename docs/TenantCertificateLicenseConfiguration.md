# TenantCertificateLicenseConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int64** | Certificate license limit for this tenant | 
**Scope** | [**TenantCertificateLicenseModule**](TenantCertificateLicenseModule.md) |  | 

## Methods

### NewTenantCertificateLicenseConfiguration

`func NewTenantCertificateLicenseConfiguration(limit int64, scope TenantCertificateLicenseModule, ) *TenantCertificateLicenseConfiguration`

NewTenantCertificateLicenseConfiguration instantiates a new TenantCertificateLicenseConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCertificateLicenseConfigurationWithDefaults

`func NewTenantCertificateLicenseConfigurationWithDefaults() *TenantCertificateLicenseConfiguration`

NewTenantCertificateLicenseConfigurationWithDefaults instantiates a new TenantCertificateLicenseConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *TenantCertificateLicenseConfiguration) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TenantCertificateLicenseConfiguration) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TenantCertificateLicenseConfiguration) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetScope

`func (o *TenantCertificateLicenseConfiguration) GetScope() TenantCertificateLicenseModule`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TenantCertificateLicenseConfiguration) GetScopeOk() (*TenantCertificateLicenseModule, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TenantCertificateLicenseConfiguration) SetScope(v TenantCertificateLicenseModule)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


