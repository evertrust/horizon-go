# RequestPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approve** | Pointer to **NullableBool** | Whether the principal can approve or deny the request (manage rights) | [optional] 
**Cancel** | Pointer to **NullableBool** | Whether the principal can cancel the request (owner or team of the request) | [optional] 

## Methods

### NewRequestPermissions

`func NewRequestPermissions() *RequestPermissions`

NewRequestPermissions instantiates a new RequestPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPermissionsWithDefaults

`func NewRequestPermissionsWithDefaults() *RequestPermissions`

NewRequestPermissionsWithDefaults instantiates a new RequestPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprove

`func (o *RequestPermissions) GetApprove() bool`

GetApprove returns the Approve field if non-nil, zero value otherwise.

### GetApproveOk

`func (o *RequestPermissions) GetApproveOk() (*bool, bool)`

GetApproveOk returns a tuple with the Approve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprove

`func (o *RequestPermissions) SetApprove(v bool)`

SetApprove sets Approve field to given value.

### HasApprove

`func (o *RequestPermissions) HasApprove() bool`

HasApprove returns a boolean if a field has been set.

### SetApproveNil

`func (o *RequestPermissions) SetApproveNil(b bool)`

 SetApproveNil sets the value for Approve to be an explicit nil

### UnsetApprove
`func (o *RequestPermissions) UnsetApprove()`

UnsetApprove ensures that no value is present for Approve, not even an explicit nil
### GetCancel

`func (o *RequestPermissions) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *RequestPermissions) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *RequestPermissions) SetCancel(v bool)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *RequestPermissions) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### SetCancelNil

`func (o *RequestPermissions) SetCancelNil(b bool)`

 SetCancelNil sets the value for Cancel to be an explicit nil

### UnsetCancel
`func (o *RequestPermissions) UnsetCancel()`

UnsetCancel ensures that no value is present for Cancel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


