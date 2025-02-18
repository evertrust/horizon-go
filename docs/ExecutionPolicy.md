# ExecutionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**AuthorizedPeriods** | Pointer to [**[]ExecutionPeriod**](ExecutionPeriod.md) |  | [optional] 
**ForbiddenPeriods** | Pointer to [**[]ExecutionPeriod**](ExecutionPeriod.md) |  | [optional] 

## Methods

### NewExecutionPolicy

`func NewExecutionPolicy(name string, ) *ExecutionPolicy`

NewExecutionPolicy instantiates a new ExecutionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionPolicyWithDefaults

`func NewExecutionPolicyWithDefaults() *ExecutionPolicy`

NewExecutionPolicyWithDefaults instantiates a new ExecutionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExecutionPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecutionPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecutionPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExecutionPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExecutionPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExecutionPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExecutionPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExecutionPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExecutionPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAuthorizedPeriods

`func (o *ExecutionPolicy) GetAuthorizedPeriods() []ExecutionPeriod`

GetAuthorizedPeriods returns the AuthorizedPeriods field if non-nil, zero value otherwise.

### GetAuthorizedPeriodsOk

`func (o *ExecutionPolicy) GetAuthorizedPeriodsOk() (*[]ExecutionPeriod, bool)`

GetAuthorizedPeriodsOk returns a tuple with the AuthorizedPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedPeriods

`func (o *ExecutionPolicy) SetAuthorizedPeriods(v []ExecutionPeriod)`

SetAuthorizedPeriods sets AuthorizedPeriods field to given value.

### HasAuthorizedPeriods

`func (o *ExecutionPolicy) HasAuthorizedPeriods() bool`

HasAuthorizedPeriods returns a boolean if a field has been set.

### SetAuthorizedPeriodsNil

`func (o *ExecutionPolicy) SetAuthorizedPeriodsNil(b bool)`

 SetAuthorizedPeriodsNil sets the value for AuthorizedPeriods to be an explicit nil

### UnsetAuthorizedPeriods
`func (o *ExecutionPolicy) UnsetAuthorizedPeriods()`

UnsetAuthorizedPeriods ensures that no value is present for AuthorizedPeriods, not even an explicit nil
### GetForbiddenPeriods

`func (o *ExecutionPolicy) GetForbiddenPeriods() []ExecutionPeriod`

GetForbiddenPeriods returns the ForbiddenPeriods field if non-nil, zero value otherwise.

### GetForbiddenPeriodsOk

`func (o *ExecutionPolicy) GetForbiddenPeriodsOk() (*[]ExecutionPeriod, bool)`

GetForbiddenPeriodsOk returns a tuple with the ForbiddenPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenPeriods

`func (o *ExecutionPolicy) SetForbiddenPeriods(v []ExecutionPeriod)`

SetForbiddenPeriods sets ForbiddenPeriods field to given value.

### HasForbiddenPeriods

`func (o *ExecutionPolicy) HasForbiddenPeriods() bool`

HasForbiddenPeriods returns a boolean if a field has been set.

### SetForbiddenPeriodsNil

`func (o *ExecutionPolicy) SetForbiddenPeriodsNil(b bool)`

 SetForbiddenPeriodsNil sets the value for ForbiddenPeriods to be an explicit nil

### UnsetForbiddenPeriods
`func (o *ExecutionPolicy) UnsetForbiddenPeriods()`

UnsetForbiddenPeriods ensures that no value is present for ForbiddenPeriods, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


