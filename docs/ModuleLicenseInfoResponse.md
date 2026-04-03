# ModuleLicenseInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **int64** |  | 
**Limit** | Pointer to **NullableInt64** |  | [optional] 
**Module** | **string** |  | 

## Methods

### NewModuleLicenseInfoResponse

`func NewModuleLicenseInfoResponse(items int64, module string, ) *ModuleLicenseInfoResponse`

NewModuleLicenseInfoResponse instantiates a new ModuleLicenseInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleLicenseInfoResponseWithDefaults

`func NewModuleLicenseInfoResponseWithDefaults() *ModuleLicenseInfoResponse`

NewModuleLicenseInfoResponseWithDefaults instantiates a new ModuleLicenseInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ModuleLicenseInfoResponse) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModuleLicenseInfoResponse) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModuleLicenseInfoResponse) SetItems(v int64)`

SetItems sets Items field to given value.


### GetLimit

`func (o *ModuleLicenseInfoResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ModuleLicenseInfoResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ModuleLicenseInfoResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ModuleLicenseInfoResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ModuleLicenseInfoResponse) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ModuleLicenseInfoResponse) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModule

`func (o *ModuleLicenseInfoResponse) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ModuleLicenseInfoResponse) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ModuleLicenseInfoResponse) SetModule(v string)`

SetModule sets Module field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


