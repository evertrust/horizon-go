# DataSourceOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key of the output | 
**Multi** | Pointer to **NullableBool** | Determine if this attribute is multivalued | [optional] [default to false]
**Selected** | Pointer to **NullableBool** | Determine if the attribute is selected on future fetches | [optional] [default to true]

## Methods

### NewDataSourceOutput

`func NewDataSourceOutput(key string, ) *DataSourceOutput`

NewDataSourceOutput instantiates a new DataSourceOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceOutputWithDefaults

`func NewDataSourceOutputWithDefaults() *DataSourceOutput`

NewDataSourceOutputWithDefaults instantiates a new DataSourceOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DataSourceOutput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DataSourceOutput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DataSourceOutput) SetKey(v string)`

SetKey sets Key field to given value.


### GetMulti

`func (o *DataSourceOutput) GetMulti() bool`

GetMulti returns the Multi field if non-nil, zero value otherwise.

### GetMultiOk

`func (o *DataSourceOutput) GetMultiOk() (*bool, bool)`

GetMultiOk returns a tuple with the Multi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulti

`func (o *DataSourceOutput) SetMulti(v bool)`

SetMulti sets Multi field to given value.

### HasMulti

`func (o *DataSourceOutput) HasMulti() bool`

HasMulti returns a boolean if a field has been set.

### SetMultiNil

`func (o *DataSourceOutput) SetMultiNil(b bool)`

 SetMultiNil sets the value for Multi to be an explicit nil

### UnsetMulti
`func (o *DataSourceOutput) UnsetMulti()`

UnsetMulti ensures that no value is present for Multi, not even an explicit nil
### GetSelected

`func (o *DataSourceOutput) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *DataSourceOutput) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *DataSourceOutput) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *DataSourceOutput) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### SetSelectedNil

`func (o *DataSourceOutput) SetSelectedNil(b bool)`

 SetSelectedNil sets the value for Selected to be an explicit nil

### UnsetSelected
`func (o *DataSourceOutput) UnsetSelected()`

UnsetSelected ensures that no value is present for Selected, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


