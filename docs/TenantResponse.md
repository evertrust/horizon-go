# TenantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Object internal ID | [optional] 
**DeletableAt** | Pointer to **int64** | Instant after which this tenant will be deletable | [optional] 
**Description** | Pointer to **string** | A simple description for this tenant | [optional] 
**LicenseExpiration** | Pointer to **int64** | Custom license expiration date for this tenant | [optional] 
**LicenseLimit** | **int64** | License limit for this tenant | 
**Name** | **string** | The tenant internal name | 

## Methods

### NewTenantResponse

`func NewTenantResponse(licenseLimit int64, name string, ) *TenantResponse`

NewTenantResponse instantiates a new TenantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantResponseWithDefaults

`func NewTenantResponseWithDefaults() *TenantResponse`

NewTenantResponseWithDefaults instantiates a new TenantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletableAt

`func (o *TenantResponse) GetDeletableAt() int64`

GetDeletableAt returns the DeletableAt field if non-nil, zero value otherwise.

### GetDeletableAtOk

`func (o *TenantResponse) GetDeletableAtOk() (*int64, bool)`

GetDeletableAtOk returns a tuple with the DeletableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletableAt

`func (o *TenantResponse) SetDeletableAt(v int64)`

SetDeletableAt sets DeletableAt field to given value.

### HasDeletableAt

`func (o *TenantResponse) HasDeletableAt() bool`

HasDeletableAt returns a boolean if a field has been set.

### GetDescription

`func (o *TenantResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicenseExpiration

`func (o *TenantResponse) GetLicenseExpiration() int64`

GetLicenseExpiration returns the LicenseExpiration field if non-nil, zero value otherwise.

### GetLicenseExpirationOk

`func (o *TenantResponse) GetLicenseExpirationOk() (*int64, bool)`

GetLicenseExpirationOk returns a tuple with the LicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiration

`func (o *TenantResponse) SetLicenseExpiration(v int64)`

SetLicenseExpiration sets LicenseExpiration field to given value.

### HasLicenseExpiration

`func (o *TenantResponse) HasLicenseExpiration() bool`

HasLicenseExpiration returns a boolean if a field has been set.

### GetLicenseLimit

`func (o *TenantResponse) GetLicenseLimit() int64`

GetLicenseLimit returns the LicenseLimit field if non-nil, zero value otherwise.

### GetLicenseLimitOk

`func (o *TenantResponse) GetLicenseLimitOk() (*int64, bool)`

GetLicenseLimitOk returns a tuple with the LicenseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseLimit

`func (o *TenantResponse) SetLicenseLimit(v int64)`

SetLicenseLimit sets LicenseLimit field to given value.


### GetName

`func (o *TenantResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantResponse) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


