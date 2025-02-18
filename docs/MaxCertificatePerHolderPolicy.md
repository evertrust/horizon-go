# MaxCertificatePerHolderPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | **int64** |  | 
**Behavior** | **string** |  | 
**RevocationReason** | Pointer to **NullableString** | One of: &#x60;unspecified&#x60;, &#x60;keycompromise&#x60;, &#x60;cacompromise&#x60;, &#x60;affiliationchange&#x60;, &#x60;superseded&#x60;, &#x60;cessationofoperation&#x60; | [optional] 

## Methods

### NewMaxCertificatePerHolderPolicy

`func NewMaxCertificatePerHolderPolicy(max int64, behavior string, ) *MaxCertificatePerHolderPolicy`

NewMaxCertificatePerHolderPolicy instantiates a new MaxCertificatePerHolderPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxCertificatePerHolderPolicyWithDefaults

`func NewMaxCertificatePerHolderPolicyWithDefaults() *MaxCertificatePerHolderPolicy`

NewMaxCertificatePerHolderPolicyWithDefaults instantiates a new MaxCertificatePerHolderPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *MaxCertificatePerHolderPolicy) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *MaxCertificatePerHolderPolicy) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *MaxCertificatePerHolderPolicy) SetMax(v int64)`

SetMax sets Max field to given value.


### GetBehavior

`func (o *MaxCertificatePerHolderPolicy) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *MaxCertificatePerHolderPolicy) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *MaxCertificatePerHolderPolicy) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.


### GetRevocationReason

`func (o *MaxCertificatePerHolderPolicy) GetRevocationReason() string`

GetRevocationReason returns the RevocationReason field if non-nil, zero value otherwise.

### GetRevocationReasonOk

`func (o *MaxCertificatePerHolderPolicy) GetRevocationReasonOk() (*string, bool)`

GetRevocationReasonOk returns a tuple with the RevocationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationReason

`func (o *MaxCertificatePerHolderPolicy) SetRevocationReason(v string)`

SetRevocationReason sets RevocationReason field to given value.

### HasRevocationReason

`func (o *MaxCertificatePerHolderPolicy) HasRevocationReason() bool`

HasRevocationReason returns a boolean if a field has been set.

### SetRevocationReasonNil

`func (o *MaxCertificatePerHolderPolicy) SetRevocationReasonNil(b bool)`

 SetRevocationReasonNil sets the value for RevocationReason to be an explicit nil

### UnsetRevocationReason
`func (o *MaxCertificatePerHolderPolicy) UnsetRevocationReason()`

UnsetRevocationReason ensures that no value is present for RevocationReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


