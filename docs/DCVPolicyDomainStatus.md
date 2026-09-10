# DCVPolicyDomainStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DcvExpiration** | Pointer to **NullableInt64** | Epoch milliseconds of when the DCV validation expires, only present when validated | [optional] 
**DcvMethod** | Pointer to **NullableString** | DCV method used to validate this domain (e.g. dns_txt, dns_cname, or an unsupported method name) | [optional] 
**DcvStatus** | Pointer to **NullableString** | Current DCV validation status for this domain | [optional] 
**Domain** | **string** | The hostname of the domain being validated | 
**ExecutionStatus** | Pointer to **NullableString** | Current execution status of this domain within the active policy run; absent when the policy is not running or if the domain was not chosen | [optional] 
**IsActive** | **bool** | Whether this domain is currently active in the policy | 

## Methods

### NewDCVPolicyDomainStatus

`func NewDCVPolicyDomainStatus(domain string, isActive bool, ) *DCVPolicyDomainStatus`

NewDCVPolicyDomainStatus instantiates a new DCVPolicyDomainStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVPolicyDomainStatusWithDefaults

`func NewDCVPolicyDomainStatusWithDefaults() *DCVPolicyDomainStatus`

NewDCVPolicyDomainStatusWithDefaults instantiates a new DCVPolicyDomainStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDcvExpiration

`func (o *DCVPolicyDomainStatus) GetDcvExpiration() int64`

GetDcvExpiration returns the DcvExpiration field if non-nil, zero value otherwise.

### GetDcvExpirationOk

`func (o *DCVPolicyDomainStatus) GetDcvExpirationOk() (*int64, bool)`

GetDcvExpirationOk returns a tuple with the DcvExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcvExpiration

`func (o *DCVPolicyDomainStatus) SetDcvExpiration(v int64)`

SetDcvExpiration sets DcvExpiration field to given value.

### HasDcvExpiration

`func (o *DCVPolicyDomainStatus) HasDcvExpiration() bool`

HasDcvExpiration returns a boolean if a field has been set.

### SetDcvExpirationNil

`func (o *DCVPolicyDomainStatus) SetDcvExpirationNil(b bool)`

 SetDcvExpirationNil sets the value for DcvExpiration to be an explicit nil

### UnsetDcvExpiration
`func (o *DCVPolicyDomainStatus) UnsetDcvExpiration()`

UnsetDcvExpiration ensures that no value is present for DcvExpiration, not even an explicit nil
### GetDcvMethod

`func (o *DCVPolicyDomainStatus) GetDcvMethod() string`

GetDcvMethod returns the DcvMethod field if non-nil, zero value otherwise.

### GetDcvMethodOk

`func (o *DCVPolicyDomainStatus) GetDcvMethodOk() (*string, bool)`

GetDcvMethodOk returns a tuple with the DcvMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcvMethod

`func (o *DCVPolicyDomainStatus) SetDcvMethod(v string)`

SetDcvMethod sets DcvMethod field to given value.

### HasDcvMethod

`func (o *DCVPolicyDomainStatus) HasDcvMethod() bool`

HasDcvMethod returns a boolean if a field has been set.

### SetDcvMethodNil

`func (o *DCVPolicyDomainStatus) SetDcvMethodNil(b bool)`

 SetDcvMethodNil sets the value for DcvMethod to be an explicit nil

### UnsetDcvMethod
`func (o *DCVPolicyDomainStatus) UnsetDcvMethod()`

UnsetDcvMethod ensures that no value is present for DcvMethod, not even an explicit nil
### GetDcvStatus

`func (o *DCVPolicyDomainStatus) GetDcvStatus() string`

GetDcvStatus returns the DcvStatus field if non-nil, zero value otherwise.

### GetDcvStatusOk

`func (o *DCVPolicyDomainStatus) GetDcvStatusOk() (*string, bool)`

GetDcvStatusOk returns a tuple with the DcvStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcvStatus

`func (o *DCVPolicyDomainStatus) SetDcvStatus(v string)`

SetDcvStatus sets DcvStatus field to given value.

### HasDcvStatus

`func (o *DCVPolicyDomainStatus) HasDcvStatus() bool`

HasDcvStatus returns a boolean if a field has been set.

### SetDcvStatusNil

`func (o *DCVPolicyDomainStatus) SetDcvStatusNil(b bool)`

 SetDcvStatusNil sets the value for DcvStatus to be an explicit nil

### UnsetDcvStatus
`func (o *DCVPolicyDomainStatus) UnsetDcvStatus()`

UnsetDcvStatus ensures that no value is present for DcvStatus, not even an explicit nil
### GetDomain

`func (o *DCVPolicyDomainStatus) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DCVPolicyDomainStatus) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DCVPolicyDomainStatus) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetExecutionStatus

`func (o *DCVPolicyDomainStatus) GetExecutionStatus() string`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *DCVPolicyDomainStatus) GetExecutionStatusOk() (*string, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *DCVPolicyDomainStatus) SetExecutionStatus(v string)`

SetExecutionStatus sets ExecutionStatus field to given value.

### HasExecutionStatus

`func (o *DCVPolicyDomainStatus) HasExecutionStatus() bool`

HasExecutionStatus returns a boolean if a field has been set.

### SetExecutionStatusNil

`func (o *DCVPolicyDomainStatus) SetExecutionStatusNil(b bool)`

 SetExecutionStatusNil sets the value for ExecutionStatus to be an explicit nil

### UnsetExecutionStatus
`func (o *DCVPolicyDomainStatus) UnsetExecutionStatus()`

UnsetExecutionStatus ensures that no value is present for ExecutionStatus, not even an explicit nil
### GetIsActive

`func (o *DCVPolicyDomainStatus) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DCVPolicyDomainStatus) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DCVPolicyDomainStatus) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


