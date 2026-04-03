# PKIQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ClusterWide** | **bool** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Size** | **int64** |  | 
**ThrottleDuration** | Pointer to **NullableString** |  | [optional] 
**ThrottleParallelism** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewPKIQueueResponse

`func NewPKIQueueResponse(id string, clusterWide bool, name string, size int64, ) *PKIQueueResponse`

NewPKIQueueResponse instantiates a new PKIQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIQueueResponseWithDefaults

`func NewPKIQueueResponseWithDefaults() *PKIQueueResponse`

NewPKIQueueResponseWithDefaults instantiates a new PKIQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PKIQueueResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PKIQueueResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PKIQueueResponse) SetId(v string)`

SetId sets Id field to given value.


### GetClusterWide

`func (o *PKIQueueResponse) GetClusterWide() bool`

GetClusterWide returns the ClusterWide field if non-nil, zero value otherwise.

### GetClusterWideOk

`func (o *PKIQueueResponse) GetClusterWideOk() (*bool, bool)`

GetClusterWideOk returns a tuple with the ClusterWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterWide

`func (o *PKIQueueResponse) SetClusterWide(v bool)`

SetClusterWide sets ClusterWide field to given value.


### GetDescription

`func (o *PKIQueueResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PKIQueueResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PKIQueueResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PKIQueueResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PKIQueueResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PKIQueueResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *PKIQueueResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PKIQueueResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PKIQueueResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *PKIQueueResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PKIQueueResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PKIQueueResponse) SetSize(v int64)`

SetSize sets Size field to given value.


### GetThrottleDuration

`func (o *PKIQueueResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *PKIQueueResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *PKIQueueResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.

### HasThrottleDuration

`func (o *PKIQueueResponse) HasThrottleDuration() bool`

HasThrottleDuration returns a boolean if a field has been set.

### SetThrottleDurationNil

`func (o *PKIQueueResponse) SetThrottleDurationNil(b bool)`

 SetThrottleDurationNil sets the value for ThrottleDuration to be an explicit nil

### UnsetThrottleDuration
`func (o *PKIQueueResponse) UnsetThrottleDuration()`

UnsetThrottleDuration ensures that no value is present for ThrottleDuration, not even an explicit nil
### GetThrottleParallelism

`func (o *PKIQueueResponse) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *PKIQueueResponse) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *PKIQueueResponse) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.

### HasThrottleParallelism

`func (o *PKIQueueResponse) HasThrottleParallelism() bool`

HasThrottleParallelism returns a boolean if a field has been set.

### SetThrottleParallelismNil

`func (o *PKIQueueResponse) SetThrottleParallelismNil(b bool)`

 SetThrottleParallelismNil sets the value for ThrottleParallelism to be an explicit nil

### UnsetThrottleParallelism
`func (o *PKIQueueResponse) UnsetThrottleParallelism()`

UnsetThrottleParallelism ensures that no value is present for ThrottleParallelism, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


