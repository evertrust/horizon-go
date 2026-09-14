# HavingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | An operator for comparison | 
**Value** | **int64** | An integer for the right hand side of the condition | 

## Methods

### NewHavingResponse

`func NewHavingResponse(operator string, value int64, ) *HavingResponse`

NewHavingResponse instantiates a new HavingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHavingResponseWithDefaults

`func NewHavingResponseWithDefaults() *HavingResponse`

NewHavingResponseWithDefaults instantiates a new HavingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *HavingResponse) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *HavingResponse) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *HavingResponse) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *HavingResponse) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HavingResponse) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HavingResponse) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


