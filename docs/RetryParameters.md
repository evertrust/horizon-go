# RetryParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | **int64** | Maximum number of retry attempts before the job is considered failed. | 
**MaxBackoff** | **string** | Maximum delay between two retries, capping the exponential backoff. | 
**MinBackoff** | **string** | Minimum delay to wait before the first retry. | 
**RandomFactor** | **float64** | Random jitter factor added to each backoff delay to avoid retry storms (e.g. &#x60;0.1&#x60; adds up to 10%). | 

## Methods

### NewRetryParameters

`func NewRetryParameters(attempts int64, maxBackoff string, minBackoff string, randomFactor float64, ) *RetryParameters`

NewRetryParameters instantiates a new RetryParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryParametersWithDefaults

`func NewRetryParametersWithDefaults() *RetryParameters`

NewRetryParametersWithDefaults instantiates a new RetryParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *RetryParameters) GetAttempts() int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *RetryParameters) GetAttemptsOk() (*int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *RetryParameters) SetAttempts(v int64)`

SetAttempts sets Attempts field to given value.


### GetMaxBackoff

`func (o *RetryParameters) GetMaxBackoff() string`

GetMaxBackoff returns the MaxBackoff field if non-nil, zero value otherwise.

### GetMaxBackoffOk

`func (o *RetryParameters) GetMaxBackoffOk() (*string, bool)`

GetMaxBackoffOk returns a tuple with the MaxBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBackoff

`func (o *RetryParameters) SetMaxBackoff(v string)`

SetMaxBackoff sets MaxBackoff field to given value.


### GetMinBackoff

`func (o *RetryParameters) GetMinBackoff() string`

GetMinBackoff returns the MinBackoff field if non-nil, zero value otherwise.

### GetMinBackoffOk

`func (o *RetryParameters) GetMinBackoffOk() (*string, bool)`

GetMinBackoffOk returns a tuple with the MinBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBackoff

`func (o *RetryParameters) SetMinBackoff(v string)`

SetMinBackoff sets MinBackoff field to given value.


### GetRandomFactor

`func (o *RetryParameters) GetRandomFactor() float64`

GetRandomFactor returns the RandomFactor field if non-nil, zero value otherwise.

### GetRandomFactorOk

`func (o *RetryParameters) GetRandomFactorOk() (*float64, bool)`

GetRandomFactorOk returns a tuple with the RandomFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomFactor

`func (o *RetryParameters) SetRandomFactor(v float64)`

SetRandomFactor sets RandomFactor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


