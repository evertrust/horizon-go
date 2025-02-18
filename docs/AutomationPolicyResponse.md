# AutomationPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Object internal ID | [optional] 
**Name** | **string** |  | 
**ExecutionPolicy** | Pointer to **NullableString** |  | [optional] 
**CompliancePolicy** | Pointer to [**NullableCompliancePolicy**](CompliancePolicy.md) |  | [optional] 
**TrustChains** | Pointer to **[]string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 

## Methods

### NewAutomationPolicyResponse

`func NewAutomationPolicyResponse(name string, ) *AutomationPolicyResponse`

NewAutomationPolicyResponse instantiates a new AutomationPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationPolicyResponseWithDefaults

`func NewAutomationPolicyResponseWithDefaults() *AutomationPolicyResponse`

NewAutomationPolicyResponseWithDefaults instantiates a new AutomationPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutomationPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationPolicyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationPolicyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AutomationPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationPolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionPolicy

`func (o *AutomationPolicyResponse) GetExecutionPolicy() string`

GetExecutionPolicy returns the ExecutionPolicy field if non-nil, zero value otherwise.

### GetExecutionPolicyOk

`func (o *AutomationPolicyResponse) GetExecutionPolicyOk() (*string, bool)`

GetExecutionPolicyOk returns a tuple with the ExecutionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionPolicy

`func (o *AutomationPolicyResponse) SetExecutionPolicy(v string)`

SetExecutionPolicy sets ExecutionPolicy field to given value.

### HasExecutionPolicy

`func (o *AutomationPolicyResponse) HasExecutionPolicy() bool`

HasExecutionPolicy returns a boolean if a field has been set.

### SetExecutionPolicyNil

`func (o *AutomationPolicyResponse) SetExecutionPolicyNil(b bool)`

 SetExecutionPolicyNil sets the value for ExecutionPolicy to be an explicit nil

### UnsetExecutionPolicy
`func (o *AutomationPolicyResponse) UnsetExecutionPolicy()`

UnsetExecutionPolicy ensures that no value is present for ExecutionPolicy, not even an explicit nil
### GetCompliancePolicy

`func (o *AutomationPolicyResponse) GetCompliancePolicy() CompliancePolicy`

GetCompliancePolicy returns the CompliancePolicy field if non-nil, zero value otherwise.

### GetCompliancePolicyOk

`func (o *AutomationPolicyResponse) GetCompliancePolicyOk() (*CompliancePolicy, bool)`

GetCompliancePolicyOk returns a tuple with the CompliancePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliancePolicy

`func (o *AutomationPolicyResponse) SetCompliancePolicy(v CompliancePolicy)`

SetCompliancePolicy sets CompliancePolicy field to given value.

### HasCompliancePolicy

`func (o *AutomationPolicyResponse) HasCompliancePolicy() bool`

HasCompliancePolicy returns a boolean if a field has been set.

### SetCompliancePolicyNil

`func (o *AutomationPolicyResponse) SetCompliancePolicyNil(b bool)`

 SetCompliancePolicyNil sets the value for CompliancePolicy to be an explicit nil

### UnsetCompliancePolicy
`func (o *AutomationPolicyResponse) UnsetCompliancePolicy()`

UnsetCompliancePolicy ensures that no value is present for CompliancePolicy, not even an explicit nil
### GetTrustChains

`func (o *AutomationPolicyResponse) GetTrustChains() []string`

GetTrustChains returns the TrustChains field if non-nil, zero value otherwise.

### GetTrustChainsOk

`func (o *AutomationPolicyResponse) GetTrustChainsOk() (*[]string, bool)`

GetTrustChainsOk returns a tuple with the TrustChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChains

`func (o *AutomationPolicyResponse) SetTrustChains(v []string)`

SetTrustChains sets TrustChains field to given value.

### HasTrustChains

`func (o *AutomationPolicyResponse) HasTrustChains() bool`

HasTrustChains returns a boolean if a field has been set.

### SetTrustChainsNil

`func (o *AutomationPolicyResponse) SetTrustChainsNil(b bool)`

 SetTrustChainsNil sets the value for TrustChains to be an explicit nil

### UnsetTrustChains
`func (o *AutomationPolicyResponse) UnsetTrustChains()`

UnsetTrustChains ensures that no value is present for TrustChains, not even an explicit nil
### GetProfile

`func (o *AutomationPolicyResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AutomationPolicyResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AutomationPolicyResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AutomationPolicyResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


