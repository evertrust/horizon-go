# ScimProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | The description of the Scim profile | [optional] 
**MailType** | Pointer to **NullableString** | The mail type corresponds to the mail coming from the scim provider that must be synchronised in horizon. By default, the mail type is \&quot;work\&quot;. | [optional] [default to "work"]
**Mappings** | Pointer to [**[]ScimProfileMappingsInner**](ScimProfileMappingsInner.md) | The mapping used to synchronize user and group between the scim provider and Horizon. | [optional] 
**Name** | **string** | The name of the Scim profile | 

## Methods

### NewScimProfile

`func NewScimProfile(name string, ) *ScimProfile`

NewScimProfile instantiates a new ScimProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimProfileWithDefaults

`func NewScimProfileWithDefaults() *ScimProfile`

NewScimProfileWithDefaults instantiates a new ScimProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScimProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScimProfile) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScimProfile) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMailType

`func (o *ScimProfile) GetMailType() string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *ScimProfile) GetMailTypeOk() (*string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *ScimProfile) SetMailType(v string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *ScimProfile) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### SetMailTypeNil

`func (o *ScimProfile) SetMailTypeNil(b bool)`

 SetMailTypeNil sets the value for MailType to be an explicit nil

### UnsetMailType
`func (o *ScimProfile) UnsetMailType()`

UnsetMailType ensures that no value is present for MailType, not even an explicit nil
### GetMappings

`func (o *ScimProfile) GetMappings() []ScimProfileMappingsInner`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *ScimProfile) GetMappingsOk() (*[]ScimProfileMappingsInner, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *ScimProfile) SetMappings(v []ScimProfileMappingsInner)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *ScimProfile) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### SetMappingsNil

`func (o *ScimProfile) SetMappingsNil(b bool)`

 SetMappingsNil sets the value for Mappings to be an explicit nil

### UnsetMappings
`func (o *ScimProfile) UnsetMappings()`

UnsetMappings ensures that no value is present for Mappings, not even an explicit nil
### GetName

`func (o *ScimProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimProfile) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


