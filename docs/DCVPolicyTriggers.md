# DCVPolicyTriggers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnDcvPolicyEnd** | Pointer to **[]string** | Triggers to fire when a DCV policy run ends | [optional] 
**OnDcvPolicyStart** | Pointer to **[]string** | Triggers to fire when a DCV policy run starts | [optional] 
**OnDcvValidationFailure** | Pointer to **[]string** | Triggers to fire when domain validation fails | [optional] 
**OnDcvValidationRetry** | Pointer to **[]string** | Triggers to fire when domain validation is retried | [optional] 
**OnDcvValidationSuccess** | Pointer to **[]string** | Triggers to fire when domain validation succeeds | [optional] 

## Methods

### NewDCVPolicyTriggers

`func NewDCVPolicyTriggers() *DCVPolicyTriggers`

NewDCVPolicyTriggers instantiates a new DCVPolicyTriggers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVPolicyTriggersWithDefaults

`func NewDCVPolicyTriggersWithDefaults() *DCVPolicyTriggers`

NewDCVPolicyTriggersWithDefaults instantiates a new DCVPolicyTriggers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnDcvPolicyEnd

`func (o *DCVPolicyTriggers) GetOnDcvPolicyEnd() []string`

GetOnDcvPolicyEnd returns the OnDcvPolicyEnd field if non-nil, zero value otherwise.

### GetOnDcvPolicyEndOk

`func (o *DCVPolicyTriggers) GetOnDcvPolicyEndOk() (*[]string, bool)`

GetOnDcvPolicyEndOk returns a tuple with the OnDcvPolicyEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDcvPolicyEnd

`func (o *DCVPolicyTriggers) SetOnDcvPolicyEnd(v []string)`

SetOnDcvPolicyEnd sets OnDcvPolicyEnd field to given value.

### HasOnDcvPolicyEnd

`func (o *DCVPolicyTriggers) HasOnDcvPolicyEnd() bool`

HasOnDcvPolicyEnd returns a boolean if a field has been set.

### SetOnDcvPolicyEndNil

`func (o *DCVPolicyTriggers) SetOnDcvPolicyEndNil(b bool)`

 SetOnDcvPolicyEndNil sets the value for OnDcvPolicyEnd to be an explicit nil

### UnsetOnDcvPolicyEnd
`func (o *DCVPolicyTriggers) UnsetOnDcvPolicyEnd()`

UnsetOnDcvPolicyEnd ensures that no value is present for OnDcvPolicyEnd, not even an explicit nil
### GetOnDcvPolicyStart

`func (o *DCVPolicyTriggers) GetOnDcvPolicyStart() []string`

GetOnDcvPolicyStart returns the OnDcvPolicyStart field if non-nil, zero value otherwise.

### GetOnDcvPolicyStartOk

`func (o *DCVPolicyTriggers) GetOnDcvPolicyStartOk() (*[]string, bool)`

GetOnDcvPolicyStartOk returns a tuple with the OnDcvPolicyStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDcvPolicyStart

`func (o *DCVPolicyTriggers) SetOnDcvPolicyStart(v []string)`

SetOnDcvPolicyStart sets OnDcvPolicyStart field to given value.

### HasOnDcvPolicyStart

`func (o *DCVPolicyTriggers) HasOnDcvPolicyStart() bool`

HasOnDcvPolicyStart returns a boolean if a field has been set.

### SetOnDcvPolicyStartNil

`func (o *DCVPolicyTriggers) SetOnDcvPolicyStartNil(b bool)`

 SetOnDcvPolicyStartNil sets the value for OnDcvPolicyStart to be an explicit nil

### UnsetOnDcvPolicyStart
`func (o *DCVPolicyTriggers) UnsetOnDcvPolicyStart()`

UnsetOnDcvPolicyStart ensures that no value is present for OnDcvPolicyStart, not even an explicit nil
### GetOnDcvValidationFailure

`func (o *DCVPolicyTriggers) GetOnDcvValidationFailure() []string`

GetOnDcvValidationFailure returns the OnDcvValidationFailure field if non-nil, zero value otherwise.

### GetOnDcvValidationFailureOk

`func (o *DCVPolicyTriggers) GetOnDcvValidationFailureOk() (*[]string, bool)`

GetOnDcvValidationFailureOk returns a tuple with the OnDcvValidationFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDcvValidationFailure

`func (o *DCVPolicyTriggers) SetOnDcvValidationFailure(v []string)`

SetOnDcvValidationFailure sets OnDcvValidationFailure field to given value.

### HasOnDcvValidationFailure

`func (o *DCVPolicyTriggers) HasOnDcvValidationFailure() bool`

HasOnDcvValidationFailure returns a boolean if a field has been set.

### SetOnDcvValidationFailureNil

`func (o *DCVPolicyTriggers) SetOnDcvValidationFailureNil(b bool)`

 SetOnDcvValidationFailureNil sets the value for OnDcvValidationFailure to be an explicit nil

### UnsetOnDcvValidationFailure
`func (o *DCVPolicyTriggers) UnsetOnDcvValidationFailure()`

UnsetOnDcvValidationFailure ensures that no value is present for OnDcvValidationFailure, not even an explicit nil
### GetOnDcvValidationRetry

`func (o *DCVPolicyTriggers) GetOnDcvValidationRetry() []string`

GetOnDcvValidationRetry returns the OnDcvValidationRetry field if non-nil, zero value otherwise.

### GetOnDcvValidationRetryOk

`func (o *DCVPolicyTriggers) GetOnDcvValidationRetryOk() (*[]string, bool)`

GetOnDcvValidationRetryOk returns a tuple with the OnDcvValidationRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDcvValidationRetry

`func (o *DCVPolicyTriggers) SetOnDcvValidationRetry(v []string)`

SetOnDcvValidationRetry sets OnDcvValidationRetry field to given value.

### HasOnDcvValidationRetry

`func (o *DCVPolicyTriggers) HasOnDcvValidationRetry() bool`

HasOnDcvValidationRetry returns a boolean if a field has been set.

### SetOnDcvValidationRetryNil

`func (o *DCVPolicyTriggers) SetOnDcvValidationRetryNil(b bool)`

 SetOnDcvValidationRetryNil sets the value for OnDcvValidationRetry to be an explicit nil

### UnsetOnDcvValidationRetry
`func (o *DCVPolicyTriggers) UnsetOnDcvValidationRetry()`

UnsetOnDcvValidationRetry ensures that no value is present for OnDcvValidationRetry, not even an explicit nil
### GetOnDcvValidationSuccess

`func (o *DCVPolicyTriggers) GetOnDcvValidationSuccess() []string`

GetOnDcvValidationSuccess returns the OnDcvValidationSuccess field if non-nil, zero value otherwise.

### GetOnDcvValidationSuccessOk

`func (o *DCVPolicyTriggers) GetOnDcvValidationSuccessOk() (*[]string, bool)`

GetOnDcvValidationSuccessOk returns a tuple with the OnDcvValidationSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDcvValidationSuccess

`func (o *DCVPolicyTriggers) SetOnDcvValidationSuccess(v []string)`

SetOnDcvValidationSuccess sets OnDcvValidationSuccess field to given value.

### HasOnDcvValidationSuccess

`func (o *DCVPolicyTriggers) HasOnDcvValidationSuccess() bool`

HasOnDcvValidationSuccess returns a boolean if a field has been set.

### SetOnDcvValidationSuccessNil

`func (o *DCVPolicyTriggers) SetOnDcvValidationSuccessNil(b bool)`

 SetOnDcvValidationSuccessNil sets the value for OnDcvValidationSuccess to be an explicit nil

### UnsetOnDcvValidationSuccess
`func (o *DCVPolicyTriggers) UnsetOnDcvValidationSuccess()`

UnsetOnDcvValidationSuccess ensures that no value is present for OnDcvValidationSuccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


