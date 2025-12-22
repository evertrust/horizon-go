# LicenseConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggers** | Pointer to [**NullableLicenseTriggers**](LicenseTriggers.md) | Triggers to execute on license events | [optional] 
**Type** | **string** | The type of the configuration entry | 

## Methods

### NewLicenseConfiguration

`func NewLicenseConfiguration(type_ string, ) *LicenseConfiguration`

NewLicenseConfiguration instantiates a new LicenseConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseConfigurationWithDefaults

`func NewLicenseConfigurationWithDefaults() *LicenseConfiguration`

NewLicenseConfigurationWithDefaults instantiates a new LicenseConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggers

`func (o *LicenseConfiguration) GetTriggers() LicenseTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *LicenseConfiguration) GetTriggersOk() (*LicenseTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *LicenseConfiguration) SetTriggers(v LicenseTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *LicenseConfiguration) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *LicenseConfiguration) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *LicenseConfiguration) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetType

`func (o *LicenseConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseConfiguration) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


