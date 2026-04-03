# AutoRenewalPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | **bool** | The default value for auto renewal status on a new certificate | 
**Editable** | **bool** | If true, the auto renewal status can be updated on new and existing certificates | 

## Methods

### NewAutoRenewalPolicy

`func NewAutoRenewalPolicy(default_ bool, editable bool, ) *AutoRenewalPolicy`

NewAutoRenewalPolicy instantiates a new AutoRenewalPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRenewalPolicyWithDefaults

`func NewAutoRenewalPolicyWithDefaults() *AutoRenewalPolicy`

NewAutoRenewalPolicyWithDefaults instantiates a new AutoRenewalPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *AutoRenewalPolicy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *AutoRenewalPolicy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *AutoRenewalPolicy) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetEditable

`func (o *AutoRenewalPolicy) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AutoRenewalPolicy) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AutoRenewalPolicy) SetEditable(v bool)`

SetEditable sets Editable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


