# ExecutionPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**AuthorizedPeriods** | Pointer to [**[]ExecutionPeriod**](ExecutionPeriod.md) |  | [optional] 
**ForbiddenPeriods** | Pointer to [**[]ExecutionPeriod**](ExecutionPeriod.md) |  | [optional] 

## Methods

### NewExecutionPolicyResponse

`func NewExecutionPolicyResponse(id string, name string, ) *ExecutionPolicyResponse`

NewExecutionPolicyResponse instantiates a new ExecutionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionPolicyResponseWithDefaults

`func NewExecutionPolicyResponseWithDefaults() *ExecutionPolicyResponse`

NewExecutionPolicyResponseWithDefaults instantiates a new ExecutionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ExecutionPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecutionPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecutionPolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExecutionPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExecutionPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExecutionPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExecutionPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExecutionPolicyResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExecutionPolicyResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAuthorizedPeriods

`func (o *ExecutionPolicyResponse) GetAuthorizedPeriods() []ExecutionPeriod`

GetAuthorizedPeriods returns the AuthorizedPeriods field if non-nil, zero value otherwise.

### GetAuthorizedPeriodsOk

`func (o *ExecutionPolicyResponse) GetAuthorizedPeriodsOk() (*[]ExecutionPeriod, bool)`

GetAuthorizedPeriodsOk returns a tuple with the AuthorizedPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedPeriods

`func (o *ExecutionPolicyResponse) SetAuthorizedPeriods(v []ExecutionPeriod)`

SetAuthorizedPeriods sets AuthorizedPeriods field to given value.

### HasAuthorizedPeriods

`func (o *ExecutionPolicyResponse) HasAuthorizedPeriods() bool`

HasAuthorizedPeriods returns a boolean if a field has been set.

### SetAuthorizedPeriodsNil

`func (o *ExecutionPolicyResponse) SetAuthorizedPeriodsNil(b bool)`

 SetAuthorizedPeriodsNil sets the value for AuthorizedPeriods to be an explicit nil

### UnsetAuthorizedPeriods
`func (o *ExecutionPolicyResponse) UnsetAuthorizedPeriods()`

UnsetAuthorizedPeriods ensures that no value is present for AuthorizedPeriods, not even an explicit nil
### GetForbiddenPeriods

`func (o *ExecutionPolicyResponse) GetForbiddenPeriods() []ExecutionPeriod`

GetForbiddenPeriods returns the ForbiddenPeriods field if non-nil, zero value otherwise.

### GetForbiddenPeriodsOk

`func (o *ExecutionPolicyResponse) GetForbiddenPeriodsOk() (*[]ExecutionPeriod, bool)`

GetForbiddenPeriodsOk returns a tuple with the ForbiddenPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenPeriods

`func (o *ExecutionPolicyResponse) SetForbiddenPeriods(v []ExecutionPeriod)`

SetForbiddenPeriods sets ForbiddenPeriods field to given value.

### HasForbiddenPeriods

`func (o *ExecutionPolicyResponse) HasForbiddenPeriods() bool`

HasForbiddenPeriods returns a boolean if a field has been set.

### SetForbiddenPeriodsNil

`func (o *ExecutionPolicyResponse) SetForbiddenPeriodsNil(b bool)`

 SetForbiddenPeriodsNil sets the value for ForbiddenPeriods to be an explicit nil

### UnsetForbiddenPeriods
`func (o *ExecutionPolicyResponse) UnsetForbiddenPeriods()`

UnsetForbiddenPeriods ensures that no value is present for ForbiddenPeriods, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


