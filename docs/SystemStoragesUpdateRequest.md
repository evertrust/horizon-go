# SystemStoragesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | Name of the bucket to store items into | 
**ChecksumMode** | Pointer to **string** | S3 Checksum mode | [optional] [default to "when_required"]
**Credentials** | Pointer to **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing AWS Secret Keys. If not defined, environment variables will be used | [optional] 
**Description** | Pointer to **string** | Simple description for this storage | [optional] 
**Endpoint** | Pointer to **string** | Custom endpoint to use for S3 | [optional] 
**ForcePathStyle** | **bool** | If enabled, force S3 path style requests | [default to false]
**Name** | **string** | The name for this storage | 
**PartBufferSize** | **string** |  | [default to "9MB"]
**Proxy** | Pointer to **string** | Reference to the proxy to use for connection to the S3 | [optional] 
**Region** | Pointer to **string** | AWS Region for the S3 storage. If not defined, environment variables will be used | [optional] 
**RoleArn** | Pointer to **string** | AWS Role ARN to impersonate | [optional] 
**Timeout** | **NullableString** | Timeout while connecting to the S3 | 
**Type** | **string** | Type of storage | 

## Methods

### NewSystemStoragesUpdateRequest

`func NewSystemStoragesUpdateRequest(bucket string, forcePathStyle bool, name string, partBufferSize string, timeout NullableString, type_ string, ) *SystemStoragesUpdateRequest`

NewSystemStoragesUpdateRequest instantiates a new SystemStoragesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStoragesUpdateRequestWithDefaults

`func NewSystemStoragesUpdateRequestWithDefaults() *SystemStoragesUpdateRequest`

NewSystemStoragesUpdateRequestWithDefaults instantiates a new SystemStoragesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *SystemStoragesUpdateRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *SystemStoragesUpdateRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *SystemStoragesUpdateRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetChecksumMode

`func (o *SystemStoragesUpdateRequest) GetChecksumMode() string`

GetChecksumMode returns the ChecksumMode field if non-nil, zero value otherwise.

### GetChecksumModeOk

`func (o *SystemStoragesUpdateRequest) GetChecksumModeOk() (*string, bool)`

GetChecksumModeOk returns a tuple with the ChecksumMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMode

`func (o *SystemStoragesUpdateRequest) SetChecksumMode(v string)`

SetChecksumMode sets ChecksumMode field to given value.

### HasChecksumMode

`func (o *SystemStoragesUpdateRequest) HasChecksumMode() bool`

HasChecksumMode returns a boolean if a field has been set.

### GetCredentials

`func (o *SystemStoragesUpdateRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SystemStoragesUpdateRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SystemStoragesUpdateRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SystemStoragesUpdateRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDescription

`func (o *SystemStoragesUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemStoragesUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemStoragesUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemStoragesUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *SystemStoragesUpdateRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SystemStoragesUpdateRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SystemStoragesUpdateRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *SystemStoragesUpdateRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetForcePathStyle

`func (o *SystemStoragesUpdateRequest) GetForcePathStyle() bool`

GetForcePathStyle returns the ForcePathStyle field if non-nil, zero value otherwise.

### GetForcePathStyleOk

`func (o *SystemStoragesUpdateRequest) GetForcePathStyleOk() (*bool, bool)`

GetForcePathStyleOk returns a tuple with the ForcePathStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePathStyle

`func (o *SystemStoragesUpdateRequest) SetForcePathStyle(v bool)`

SetForcePathStyle sets ForcePathStyle field to given value.


### GetName

`func (o *SystemStoragesUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemStoragesUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemStoragesUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPartBufferSize

`func (o *SystemStoragesUpdateRequest) GetPartBufferSize() string`

GetPartBufferSize returns the PartBufferSize field if non-nil, zero value otherwise.

### GetPartBufferSizeOk

`func (o *SystemStoragesUpdateRequest) GetPartBufferSizeOk() (*string, bool)`

GetPartBufferSizeOk returns a tuple with the PartBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartBufferSize

`func (o *SystemStoragesUpdateRequest) SetPartBufferSize(v string)`

SetPartBufferSize sets PartBufferSize field to given value.


### GetProxy

`func (o *SystemStoragesUpdateRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SystemStoragesUpdateRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SystemStoragesUpdateRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SystemStoragesUpdateRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetRegion

`func (o *SystemStoragesUpdateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SystemStoragesUpdateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SystemStoragesUpdateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SystemStoragesUpdateRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleArn

`func (o *SystemStoragesUpdateRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *SystemStoragesUpdateRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *SystemStoragesUpdateRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *SystemStoragesUpdateRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetTimeout

`func (o *SystemStoragesUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SystemStoragesUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SystemStoragesUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *SystemStoragesUpdateRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SystemStoragesUpdateRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *SystemStoragesUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemStoragesUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemStoragesUpdateRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


