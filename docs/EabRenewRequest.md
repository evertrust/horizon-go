# EabRenewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EabValidityDuration** | Pointer to **NullableString** |  | [optional] 
**MacKeyAlgorithm** | Pointer to [**ExternalAccountBindingAlgorithm**](ExternalAccountBindingAlgorithm.md) |  | [optional] 

## Methods

### NewEabRenewRequest

`func NewEabRenewRequest() *EabRenewRequest`

NewEabRenewRequest instantiates a new EabRenewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEabRenewRequestWithDefaults

`func NewEabRenewRequestWithDefaults() *EabRenewRequest`

NewEabRenewRequestWithDefaults instantiates a new EabRenewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEabValidityDuration

`func (o *EabRenewRequest) GetEabValidityDuration() string`

GetEabValidityDuration returns the EabValidityDuration field if non-nil, zero value otherwise.

### GetEabValidityDurationOk

`func (o *EabRenewRequest) GetEabValidityDurationOk() (*string, bool)`

GetEabValidityDurationOk returns a tuple with the EabValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabValidityDuration

`func (o *EabRenewRequest) SetEabValidityDuration(v string)`

SetEabValidityDuration sets EabValidityDuration field to given value.

### HasEabValidityDuration

`func (o *EabRenewRequest) HasEabValidityDuration() bool`

HasEabValidityDuration returns a boolean if a field has been set.

### SetEabValidityDurationNil

`func (o *EabRenewRequest) SetEabValidityDurationNil(b bool)`

 SetEabValidityDurationNil sets the value for EabValidityDuration to be an explicit nil

### UnsetEabValidityDuration
`func (o *EabRenewRequest) UnsetEabValidityDuration()`

UnsetEabValidityDuration ensures that no value is present for EabValidityDuration, not even an explicit nil
### GetMacKeyAlgorithm

`func (o *EabRenewRequest) GetMacKeyAlgorithm() ExternalAccountBindingAlgorithm`

GetMacKeyAlgorithm returns the MacKeyAlgorithm field if non-nil, zero value otherwise.

### GetMacKeyAlgorithmOk

`func (o *EabRenewRequest) GetMacKeyAlgorithmOk() (*ExternalAccountBindingAlgorithm, bool)`

GetMacKeyAlgorithmOk returns a tuple with the MacKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacKeyAlgorithm

`func (o *EabRenewRequest) SetMacKeyAlgorithm(v ExternalAccountBindingAlgorithm)`

SetMacKeyAlgorithm sets MacKeyAlgorithm field to given value.

### HasMacKeyAlgorithm

`func (o *EabRenewRequest) HasMacKeyAlgorithm() bool`

HasMacKeyAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


