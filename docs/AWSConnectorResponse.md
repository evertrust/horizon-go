# AWSConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**ThrottleDuration** | **string** |  | 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Region** | **string** |  | 
**Credentials** | Pointer to **NullableString** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing Access Key Id and Secret Access Key. If not defined, an account present in environment variables can be used. | [optional] 
**ResourceGroupName** | Pointer to **NullableString** |  | [optional] 
**RoleArn** | Pointer to **NullableString** |  | [optional] 
**TagKey** | Pointer to **NullableString** |  | [optional] 
**TagValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAWSConnectorResponse

`func NewAWSConnectorResponse(id string, type_ string, name string, throttleDuration string, region string, ) *AWSConnectorResponse`

NewAWSConnectorResponse instantiates a new AWSConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSConnectorResponseWithDefaults

`func NewAWSConnectorResponseWithDefaults() *AWSConnectorResponse`

NewAWSConnectorResponseWithDefaults instantiates a new AWSConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AWSConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AWSConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AWSConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AWSConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AWSConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetThrottleDuration

`func (o *AWSConnectorResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *AWSConnectorResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *AWSConnectorResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetRenewalPeriod

`func (o *AWSConnectorResponse) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *AWSConnectorResponse) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *AWSConnectorResponse) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *AWSConnectorResponse) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *AWSConnectorResponse) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *AWSConnectorResponse) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetTimeout

`func (o *AWSConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AWSConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AWSConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AWSConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *AWSConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *AWSConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *AWSConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AWSConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AWSConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AWSConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *AWSConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *AWSConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRegion

`func (o *AWSConnectorResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSConnectorResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSConnectorResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCredentials

`func (o *AWSConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AWSConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AWSConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AWSConnectorResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *AWSConnectorResponse) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *AWSConnectorResponse) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetResourceGroupName

`func (o *AWSConnectorResponse) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AWSConnectorResponse) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AWSConnectorResponse) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AWSConnectorResponse) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *AWSConnectorResponse) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *AWSConnectorResponse) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetRoleArn

`func (o *AWSConnectorResponse) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AWSConnectorResponse) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AWSConnectorResponse) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *AWSConnectorResponse) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *AWSConnectorResponse) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *AWSConnectorResponse) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetTagKey

`func (o *AWSConnectorResponse) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *AWSConnectorResponse) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *AWSConnectorResponse) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.

### HasTagKey

`func (o *AWSConnectorResponse) HasTagKey() bool`

HasTagKey returns a boolean if a field has been set.

### SetTagKeyNil

`func (o *AWSConnectorResponse) SetTagKeyNil(b bool)`

 SetTagKeyNil sets the value for TagKey to be an explicit nil

### UnsetTagKey
`func (o *AWSConnectorResponse) UnsetTagKey()`

UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil
### GetTagValue

`func (o *AWSConnectorResponse) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *AWSConnectorResponse) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *AWSConnectorResponse) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.

### HasTagValue

`func (o *AWSConnectorResponse) HasTagValue() bool`

HasTagValue returns a boolean if a field has been set.

### SetTagValueNil

`func (o *AWSConnectorResponse) SetTagValueNil(b bool)`

 SetTagValueNil sets the value for TagValue to be an explicit nil

### UnsetTagValue
`func (o *AWSConnectorResponse) UnsetTagValue()`

UnsetTagValue ensures that no value is present for TagValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


