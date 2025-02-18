# LicenseTriggers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnLicenseExpiration** | Pointer to **[]string** | Names of the triggers to execute when the license expires | [optional] 
**OnLicenseUsage** | Pointer to **[]string** | Names of the triggers to execute when the license usage exceeds threshold | [optional] 

## Methods

### NewLicenseTriggers

`func NewLicenseTriggers() *LicenseTriggers`

NewLicenseTriggers instantiates a new LicenseTriggers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseTriggersWithDefaults

`func NewLicenseTriggersWithDefaults() *LicenseTriggers`

NewLicenseTriggersWithDefaults instantiates a new LicenseTriggers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnLicenseExpiration

`func (o *LicenseTriggers) GetOnLicenseExpiration() []string`

GetOnLicenseExpiration returns the OnLicenseExpiration field if non-nil, zero value otherwise.

### GetOnLicenseExpirationOk

`func (o *LicenseTriggers) GetOnLicenseExpirationOk() (*[]string, bool)`

GetOnLicenseExpirationOk returns a tuple with the OnLicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLicenseExpiration

`func (o *LicenseTriggers) SetOnLicenseExpiration(v []string)`

SetOnLicenseExpiration sets OnLicenseExpiration field to given value.

### HasOnLicenseExpiration

`func (o *LicenseTriggers) HasOnLicenseExpiration() bool`

HasOnLicenseExpiration returns a boolean if a field has been set.

### SetOnLicenseExpirationNil

`func (o *LicenseTriggers) SetOnLicenseExpirationNil(b bool)`

 SetOnLicenseExpirationNil sets the value for OnLicenseExpiration to be an explicit nil

### UnsetOnLicenseExpiration
`func (o *LicenseTriggers) UnsetOnLicenseExpiration()`

UnsetOnLicenseExpiration ensures that no value is present for OnLicenseExpiration, not even an explicit nil
### GetOnLicenseUsage

`func (o *LicenseTriggers) GetOnLicenseUsage() []string`

GetOnLicenseUsage returns the OnLicenseUsage field if non-nil, zero value otherwise.

### GetOnLicenseUsageOk

`func (o *LicenseTriggers) GetOnLicenseUsageOk() (*[]string, bool)`

GetOnLicenseUsageOk returns a tuple with the OnLicenseUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLicenseUsage

`func (o *LicenseTriggers) SetOnLicenseUsage(v []string)`

SetOnLicenseUsage sets OnLicenseUsage field to given value.

### HasOnLicenseUsage

`func (o *LicenseTriggers) HasOnLicenseUsage() bool`

HasOnLicenseUsage returns a boolean if a field has been set.

### SetOnLicenseUsageNil

`func (o *LicenseTriggers) SetOnLicenseUsageNil(b bool)`

 SetOnLicenseUsageNil sets the value for OnLicenseUsage to be an explicit nil

### UnsetOnLicenseUsage
`func (o *LicenseTriggers) UnsetOnLicenseUsage()`

UnsetOnLicenseUsage ensures that no value is present for OnLicenseUsage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


