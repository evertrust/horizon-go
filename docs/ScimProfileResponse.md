# ScimProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** | The name of the Scim profile | 
**Description** | Pointer to **NullableString** | The description of the Scim profile | [optional] 
**MailType** | Pointer to **NullableString** | The mail type corresponds to the mail coming from the scim provider that must be synchronised in horizon. By default, the mail type is \&quot;work\&quot;. | [optional] [default to "work"]
**Mappings** | Pointer to [**[]ScimProfileResponseMappingsInner**](ScimProfileResponseMappingsInner.md) | The mapping used to synchronize user and group between the scim provider and Horizon. | [optional] 

## Methods

### NewScimProfileResponse

`func NewScimProfileResponse(id string, name string, ) *ScimProfileResponse`

NewScimProfileResponse instantiates a new ScimProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimProfileResponseWithDefaults

`func NewScimProfileResponseWithDefaults() *ScimProfileResponse`

NewScimProfileResponseWithDefaults instantiates a new ScimProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScimProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimProfileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ScimProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimProfileResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScimProfileResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimProfileResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimProfileResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimProfileResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScimProfileResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScimProfileResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMailType

`func (o *ScimProfileResponse) GetMailType() string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *ScimProfileResponse) GetMailTypeOk() (*string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *ScimProfileResponse) SetMailType(v string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *ScimProfileResponse) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### SetMailTypeNil

`func (o *ScimProfileResponse) SetMailTypeNil(b bool)`

 SetMailTypeNil sets the value for MailType to be an explicit nil

### UnsetMailType
`func (o *ScimProfileResponse) UnsetMailType()`

UnsetMailType ensures that no value is present for MailType, not even an explicit nil
### GetMappings

`func (o *ScimProfileResponse) GetMappings() []ScimProfileResponseMappingsInner`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *ScimProfileResponse) GetMappingsOk() (*[]ScimProfileResponseMappingsInner, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *ScimProfileResponse) SetMappings(v []ScimProfileResponseMappingsInner)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *ScimProfileResponse) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### SetMappingsNil

`func (o *ScimProfileResponse) SetMappingsNil(b bool)`

 SetMappingsNil sets the value for Mappings to be an explicit nil

### UnsetMappings
`func (o *ScimProfileResponse) UnsetMappings()`

UnsetMappings ensures that no value is present for Mappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


