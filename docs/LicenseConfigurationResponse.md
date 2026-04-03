# LicenseConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Triggers** | Pointer to [**NullableLicenseTriggers**](LicenseTriggers.md) | Triggers to execute on license events | [optional] 
**Type** | **string** | The type of the configuration entry | 

## Methods

### NewLicenseConfigurationResponse

`func NewLicenseConfigurationResponse(id string, type_ string, ) *LicenseConfigurationResponse`

NewLicenseConfigurationResponse instantiates a new LicenseConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseConfigurationResponseWithDefaults

`func NewLicenseConfigurationResponseWithDefaults() *LicenseConfigurationResponse`

NewLicenseConfigurationResponseWithDefaults instantiates a new LicenseConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicenseConfigurationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseConfigurationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseConfigurationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTriggers

`func (o *LicenseConfigurationResponse) GetTriggers() LicenseTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *LicenseConfigurationResponse) GetTriggersOk() (*LicenseTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *LicenseConfigurationResponse) SetTriggers(v LicenseTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *LicenseConfigurationResponse) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *LicenseConfigurationResponse) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *LicenseConfigurationResponse) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetType

`func (o *LicenseConfigurationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseConfigurationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseConfigurationResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


