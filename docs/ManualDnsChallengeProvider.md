# ManualDnsChallengeProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetTriggers** | [**[]AcmeRestRequest**](AcmeRestRequest.md) | The triggers that will set the DNS challenge on the provider. | 
**Type** | **string** |  | 
**UnsetTriggers** | Pointer to [**[]AcmeRestRequest**](AcmeRestRequest.md) | The triggers that will unset the DNS challenge on the provider. | [optional] 

## Methods

### NewManualDnsChallengeProvider

`func NewManualDnsChallengeProvider(setTriggers []AcmeRestRequest, type_ string, ) *ManualDnsChallengeProvider`

NewManualDnsChallengeProvider instantiates a new ManualDnsChallengeProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualDnsChallengeProviderWithDefaults

`func NewManualDnsChallengeProviderWithDefaults() *ManualDnsChallengeProvider`

NewManualDnsChallengeProviderWithDefaults instantiates a new ManualDnsChallengeProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetTriggers

`func (o *ManualDnsChallengeProvider) GetSetTriggers() []AcmeRestRequest`

GetSetTriggers returns the SetTriggers field if non-nil, zero value otherwise.

### GetSetTriggersOk

`func (o *ManualDnsChallengeProvider) GetSetTriggersOk() (*[]AcmeRestRequest, bool)`

GetSetTriggersOk returns a tuple with the SetTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTriggers

`func (o *ManualDnsChallengeProvider) SetSetTriggers(v []AcmeRestRequest)`

SetSetTriggers sets SetTriggers field to given value.


### GetType

`func (o *ManualDnsChallengeProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManualDnsChallengeProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManualDnsChallengeProvider) SetType(v string)`

SetType sets Type field to given value.


### GetUnsetTriggers

`func (o *ManualDnsChallengeProvider) GetUnsetTriggers() []AcmeRestRequest`

GetUnsetTriggers returns the UnsetTriggers field if non-nil, zero value otherwise.

### GetUnsetTriggersOk

`func (o *ManualDnsChallengeProvider) GetUnsetTriggersOk() (*[]AcmeRestRequest, bool)`

GetUnsetTriggersOk returns a tuple with the UnsetTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsetTriggers

`func (o *ManualDnsChallengeProvider) SetUnsetTriggers(v []AcmeRestRequest)`

SetUnsetTriggers sets UnsetTriggers field to given value.

### HasUnsetTriggers

`func (o *ManualDnsChallengeProvider) HasUnsetTriggers() bool`

HasUnsetTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


