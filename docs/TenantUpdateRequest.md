# TenantUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A simple description for this tenant | [optional] 
**License** | [**TenantLicenseConfiguration**](TenantLicenseConfiguration.md) |  | 
**Name** | **string** | The tenant internal name | 

## Methods

### NewTenantUpdateRequest

`func NewTenantUpdateRequest(license TenantLicenseConfiguration, name string, ) *TenantUpdateRequest`

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

### GetLicense

`func (o *TenantUpdateRequest) GetLicense() TenantLicenseConfiguration`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *TenantUpdateRequest) GetLicenseOk() (*TenantLicenseConfiguration, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *TenantUpdateRequest) SetLicense(v TenantLicenseConfiguration)`

SetLicense sets License field to given value.


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


