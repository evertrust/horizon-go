# ModuleLicenseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** |  | 
**Items** | **int64** |  | 
**Limit** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewModuleLicenseInfo

`func NewModuleLicenseInfo(module string, items int64, ) *ModuleLicenseInfo`

NewModuleLicenseInfo instantiates a new ModuleLicenseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleLicenseInfoWithDefaults

`func NewModuleLicenseInfoWithDefaults() *ModuleLicenseInfo`

NewModuleLicenseInfoWithDefaults instantiates a new ModuleLicenseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *ModuleLicenseInfo) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ModuleLicenseInfo) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ModuleLicenseInfo) SetModule(v string)`

SetModule sets Module field to given value.


### GetItems

`func (o *ModuleLicenseInfo) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModuleLicenseInfo) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModuleLicenseInfo) SetItems(v int64)`

SetItems sets Items field to given value.


### GetLimit

`func (o *ModuleLicenseInfo) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ModuleLicenseInfo) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ModuleLicenseInfo) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ModuleLicenseInfo) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ModuleLicenseInfo) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ModuleLicenseInfo) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


