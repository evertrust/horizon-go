# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The permission string, in the Horizon format : &#x60;&lt;group&gt;:&lt;resource&gt;:&lt;scope&gt;:&lt;action&gt;&#x60;  | 
**Filter** | Pointer to **NullableString** | The filter to apply to the permission in the HPQL format | [optional] 

## Methods

### NewPermission

`func NewPermission(value string, ) *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Permission) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Permission) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Permission) SetValue(v string)`

SetValue sets Value field to given value.


### GetFilter

`func (o *Permission) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Permission) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Permission) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Permission) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *Permission) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *Permission) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


