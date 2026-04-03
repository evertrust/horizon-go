# AutomationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompliancePolicy** | Pointer to [**NullableCompliancePolicy**](CompliancePolicy.md) |  | [optional] 
**ExecutionPolicy** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Profile** | **string** |  | 
**TrustChains** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAutomationPolicy

`func NewAutomationPolicy(name string, profile string, ) *AutomationPolicy`

NewAutomationPolicy instantiates a new AutomationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationPolicyWithDefaults

`func NewAutomationPolicyWithDefaults() *AutomationPolicy`

NewAutomationPolicyWithDefaults instantiates a new AutomationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompliancePolicy

`func (o *AutomationPolicy) GetCompliancePolicy() CompliancePolicy`

GetCompliancePolicy returns the CompliancePolicy field if non-nil, zero value otherwise.

### GetCompliancePolicyOk

`func (o *AutomationPolicy) GetCompliancePolicyOk() (*CompliancePolicy, bool)`

GetCompliancePolicyOk returns a tuple with the CompliancePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliancePolicy

`func (o *AutomationPolicy) SetCompliancePolicy(v CompliancePolicy)`

SetCompliancePolicy sets CompliancePolicy field to given value.

### HasCompliancePolicy

`func (o *AutomationPolicy) HasCompliancePolicy() bool`

HasCompliancePolicy returns a boolean if a field has been set.

### SetCompliancePolicyNil

`func (o *AutomationPolicy) SetCompliancePolicyNil(b bool)`

 SetCompliancePolicyNil sets the value for CompliancePolicy to be an explicit nil

### UnsetCompliancePolicy
`func (o *AutomationPolicy) UnsetCompliancePolicy()`

UnsetCompliancePolicy ensures that no value is present for CompliancePolicy, not even an explicit nil
### GetExecutionPolicy

`func (o *AutomationPolicy) GetExecutionPolicy() string`

GetExecutionPolicy returns the ExecutionPolicy field if non-nil, zero value otherwise.

### GetExecutionPolicyOk

`func (o *AutomationPolicy) GetExecutionPolicyOk() (*string, bool)`

GetExecutionPolicyOk returns a tuple with the ExecutionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionPolicy

`func (o *AutomationPolicy) SetExecutionPolicy(v string)`

SetExecutionPolicy sets ExecutionPolicy field to given value.

### HasExecutionPolicy

`func (o *AutomationPolicy) HasExecutionPolicy() bool`

HasExecutionPolicy returns a boolean if a field has been set.

### SetExecutionPolicyNil

`func (o *AutomationPolicy) SetExecutionPolicyNil(b bool)`

 SetExecutionPolicyNil sets the value for ExecutionPolicy to be an explicit nil

### UnsetExecutionPolicy
`func (o *AutomationPolicy) UnsetExecutionPolicy()`

UnsetExecutionPolicy ensures that no value is present for ExecutionPolicy, not even an explicit nil
### GetName

`func (o *AutomationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *AutomationPolicy) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AutomationPolicy) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AutomationPolicy) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTrustChains

`func (o *AutomationPolicy) GetTrustChains() []string`

GetTrustChains returns the TrustChains field if non-nil, zero value otherwise.

### GetTrustChainsOk

`func (o *AutomationPolicy) GetTrustChainsOk() (*[]string, bool)`

GetTrustChainsOk returns a tuple with the TrustChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChains

`func (o *AutomationPolicy) SetTrustChains(v []string)`

SetTrustChains sets TrustChains field to given value.

### HasTrustChains

`func (o *AutomationPolicy) HasTrustChains() bool`

HasTrustChains returns a boolean if a field has been set.

### SetTrustChainsNil

`func (o *AutomationPolicy) SetTrustChainsNil(b bool)`

 SetTrustChainsNil sets the value for TrustChains to be an explicit nil

### UnsetTrustChains
`func (o *AutomationPolicy) UnsetTrustChains()`

UnsetTrustChains ensures that no value is present for TrustChains, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


