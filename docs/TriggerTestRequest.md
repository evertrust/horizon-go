# TriggerTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dictionary** | Pointer to [**[]MapEntry**](MapEntry.md) | Dictionary that will be interpreted by the trigger | [optional] 
**Trigger** | [**TriggerTestRequestTrigger**](TriggerTestRequestTrigger.md) |  | 

## Methods

### NewTriggerTestRequest

`func NewTriggerTestRequest(trigger TriggerTestRequestTrigger, ) *TriggerTestRequest`

NewTriggerTestRequest instantiates a new TriggerTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerTestRequestWithDefaults

`func NewTriggerTestRequestWithDefaults() *TriggerTestRequest`

NewTriggerTestRequestWithDefaults instantiates a new TriggerTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDictionary

`func (o *TriggerTestRequest) GetDictionary() []MapEntry`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *TriggerTestRequest) GetDictionaryOk() (*[]MapEntry, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *TriggerTestRequest) SetDictionary(v []MapEntry)`

SetDictionary sets Dictionary field to given value.

### HasDictionary

`func (o *TriggerTestRequest) HasDictionary() bool`

HasDictionary returns a boolean if a field has been set.

### SetDictionaryNil

`func (o *TriggerTestRequest) SetDictionaryNil(b bool)`

 SetDictionaryNil sets the value for Dictionary to be an explicit nil

### UnsetDictionary
`func (o *TriggerTestRequest) UnsetDictionary()`

UnsetDictionary ensures that no value is present for Dictionary, not even an explicit nil
### GetTrigger

`func (o *TriggerTestRequest) GetTrigger() TriggerTestRequestTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *TriggerTestRequest) GetTriggerOk() (*TriggerTestRequestTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *TriggerTestRequest) SetTrigger(v TriggerTestRequestTrigger)`

SetTrigger sets Trigger field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


