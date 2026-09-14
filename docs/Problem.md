# Problem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** |  | 
**SubProblems** | Pointer to [**[]Problem**](Problem.md) |  | [optional] 
**Type** | [**ProblemType**](ProblemType.md) |  | 

## Methods

### NewProblem

`func NewProblem(detail string, type_ ProblemType, ) *Problem`

NewProblem instantiates a new Problem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemWithDefaults

`func NewProblemWithDefaults() *Problem`

NewProblemWithDefaults instantiates a new Problem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *Problem) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Problem) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Problem) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSubProblems

`func (o *Problem) GetSubProblems() []Problem`

GetSubProblems returns the SubProblems field if non-nil, zero value otherwise.

### GetSubProblemsOk

`func (o *Problem) GetSubProblemsOk() (*[]Problem, bool)`

GetSubProblemsOk returns a tuple with the SubProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProblems

`func (o *Problem) SetSubProblems(v []Problem)`

SetSubProblems sets SubProblems field to given value.

### HasSubProblems

`func (o *Problem) HasSubProblems() bool`

HasSubProblems returns a boolean if a field has been set.

### GetType

`func (o *Problem) GetType() ProblemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Problem) GetTypeOk() (*ProblemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Problem) SetType(v ProblemType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


