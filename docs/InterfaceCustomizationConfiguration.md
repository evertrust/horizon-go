# InterfaceCustomizationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the configuration entry | 
**Logo** | Pointer to **NullableString** | A logo to display on the product, base64 encoded | [optional] 
**HeaderStart** | Pointer to **NullableString** | The HTML color code for the left side of the banner gradient | [optional] 
**HeaderEnd** | Pointer to **NullableString** | The HTML color code for the right side of the banner gradient | [optional] 

## Methods

### NewInterfaceCustomizationConfiguration

`func NewInterfaceCustomizationConfiguration(type_ string, ) *InterfaceCustomizationConfiguration`

NewInterfaceCustomizationConfiguration instantiates a new InterfaceCustomizationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceCustomizationConfigurationWithDefaults

`func NewInterfaceCustomizationConfigurationWithDefaults() *InterfaceCustomizationConfiguration`

NewInterfaceCustomizationConfigurationWithDefaults instantiates a new InterfaceCustomizationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InterfaceCustomizationConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfaceCustomizationConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfaceCustomizationConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetLogo

`func (o *InterfaceCustomizationConfiguration) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *InterfaceCustomizationConfiguration) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *InterfaceCustomizationConfiguration) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *InterfaceCustomizationConfiguration) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *InterfaceCustomizationConfiguration) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *InterfaceCustomizationConfiguration) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetHeaderStart

`func (o *InterfaceCustomizationConfiguration) GetHeaderStart() string`

GetHeaderStart returns the HeaderStart field if non-nil, zero value otherwise.

### GetHeaderStartOk

`func (o *InterfaceCustomizationConfiguration) GetHeaderStartOk() (*string, bool)`

GetHeaderStartOk returns a tuple with the HeaderStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderStart

`func (o *InterfaceCustomizationConfiguration) SetHeaderStart(v string)`

SetHeaderStart sets HeaderStart field to given value.

### HasHeaderStart

`func (o *InterfaceCustomizationConfiguration) HasHeaderStart() bool`

HasHeaderStart returns a boolean if a field has been set.

### SetHeaderStartNil

`func (o *InterfaceCustomizationConfiguration) SetHeaderStartNil(b bool)`

 SetHeaderStartNil sets the value for HeaderStart to be an explicit nil

### UnsetHeaderStart
`func (o *InterfaceCustomizationConfiguration) UnsetHeaderStart()`

UnsetHeaderStart ensures that no value is present for HeaderStart, not even an explicit nil
### GetHeaderEnd

`func (o *InterfaceCustomizationConfiguration) GetHeaderEnd() string`

GetHeaderEnd returns the HeaderEnd field if non-nil, zero value otherwise.

### GetHeaderEndOk

`func (o *InterfaceCustomizationConfiguration) GetHeaderEndOk() (*string, bool)`

GetHeaderEndOk returns a tuple with the HeaderEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderEnd

`func (o *InterfaceCustomizationConfiguration) SetHeaderEnd(v string)`

SetHeaderEnd sets HeaderEnd field to given value.

### HasHeaderEnd

`func (o *InterfaceCustomizationConfiguration) HasHeaderEnd() bool`

HasHeaderEnd returns a boolean if a field has been set.

### SetHeaderEndNil

`func (o *InterfaceCustomizationConfiguration) SetHeaderEndNil(b bool)`

 SetHeaderEndNil sets the value for HeaderEnd to be an explicit nil

### UnsetHeaderEnd
`func (o *InterfaceCustomizationConfiguration) UnsetHeaderEnd()`

UnsetHeaderEnd ensures that no value is present for HeaderEnd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


