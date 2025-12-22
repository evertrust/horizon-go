# PKIQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterWide** | **bool** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Size** | **int64** |  | 
**ThrottleDuration** | Pointer to **NullableString** |  | [optional] 
**ThrottleParallelism** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewPKIQueue

`func NewPKIQueue(clusterWide bool, name string, size int64, ) *PKIQueue`

NewPKIQueue instantiates a new PKIQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIQueueWithDefaults

`func NewPKIQueueWithDefaults() *PKIQueue`

NewPKIQueueWithDefaults instantiates a new PKIQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterWide

`func (o *PKIQueue) GetClusterWide() bool`

GetClusterWide returns the ClusterWide field if non-nil, zero value otherwise.

### GetClusterWideOk

`func (o *PKIQueue) GetClusterWideOk() (*bool, bool)`

GetClusterWideOk returns a tuple with the ClusterWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterWide

`func (o *PKIQueue) SetClusterWide(v bool)`

SetClusterWide sets ClusterWide field to given value.


### GetDescription

`func (o *PKIQueue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PKIQueue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PKIQueue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PKIQueue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PKIQueue) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PKIQueue) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *PKIQueue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PKIQueue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PKIQueue) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *PKIQueue) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PKIQueue) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PKIQueue) SetSize(v int64)`

SetSize sets Size field to given value.


### GetThrottleDuration

`func (o *PKIQueue) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *PKIQueue) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *PKIQueue) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.

### HasThrottleDuration

`func (o *PKIQueue) HasThrottleDuration() bool`

HasThrottleDuration returns a boolean if a field has been set.

### SetThrottleDurationNil

`func (o *PKIQueue) SetThrottleDurationNil(b bool)`

 SetThrottleDurationNil sets the value for ThrottleDuration to be an explicit nil

### UnsetThrottleDuration
`func (o *PKIQueue) UnsetThrottleDuration()`

UnsetThrottleDuration ensures that no value is present for ThrottleDuration, not even an explicit nil
### GetThrottleParallelism

`func (o *PKIQueue) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *PKIQueue) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *PKIQueue) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.

### HasThrottleParallelism

`func (o *PKIQueue) HasThrottleParallelism() bool`

HasThrottleParallelism returns a boolean if a field has been set.

### SetThrottleParallelismNil

`func (o *PKIQueue) SetThrottleParallelismNil(b bool)`

 SetThrottleParallelismNil sets the value for ThrottleParallelism to be an explicit nil

### UnsetThrottleParallelism
`func (o *PKIQueue) UnsetThrottleParallelism()`

UnsetThrottleParallelism ensures that no value is present for ThrottleParallelism, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


