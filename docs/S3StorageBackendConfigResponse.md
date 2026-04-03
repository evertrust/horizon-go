# S3StorageBackendConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
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

### NewS3StorageBackendConfigResponse

`func NewS3StorageBackendConfigResponse(id string, bucket string, forcePathStyle bool, name string, partBufferSize string, timeout NullableString, type_ string, ) *S3StorageBackendConfigResponse`

NewS3StorageBackendConfigResponse instantiates a new S3StorageBackendConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3StorageBackendConfigResponseWithDefaults

`func NewS3StorageBackendConfigResponseWithDefaults() *S3StorageBackendConfigResponse`

NewS3StorageBackendConfigResponseWithDefaults instantiates a new S3StorageBackendConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3StorageBackendConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3StorageBackendConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3StorageBackendConfigResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBucket

`func (o *S3StorageBackendConfigResponse) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3StorageBackendConfigResponse) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3StorageBackendConfigResponse) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetChecksumMode

`func (o *S3StorageBackendConfigResponse) GetChecksumMode() string`

GetChecksumMode returns the ChecksumMode field if non-nil, zero value otherwise.

### GetChecksumModeOk

`func (o *S3StorageBackendConfigResponse) GetChecksumModeOk() (*string, bool)`

GetChecksumModeOk returns a tuple with the ChecksumMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMode

`func (o *S3StorageBackendConfigResponse) SetChecksumMode(v string)`

SetChecksumMode sets ChecksumMode field to given value.

### HasChecksumMode

`func (o *S3StorageBackendConfigResponse) HasChecksumMode() bool`

HasChecksumMode returns a boolean if a field has been set.

### GetCredentials

`func (o *S3StorageBackendConfigResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *S3StorageBackendConfigResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *S3StorageBackendConfigResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *S3StorageBackendConfigResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDescription

`func (o *S3StorageBackendConfigResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *S3StorageBackendConfigResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *S3StorageBackendConfigResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *S3StorageBackendConfigResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *S3StorageBackendConfigResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *S3StorageBackendConfigResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *S3StorageBackendConfigResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *S3StorageBackendConfigResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetForcePathStyle

`func (o *S3StorageBackendConfigResponse) GetForcePathStyle() bool`

GetForcePathStyle returns the ForcePathStyle field if non-nil, zero value otherwise.

### GetForcePathStyleOk

`func (o *S3StorageBackendConfigResponse) GetForcePathStyleOk() (*bool, bool)`

GetForcePathStyleOk returns a tuple with the ForcePathStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePathStyle

`func (o *S3StorageBackendConfigResponse) SetForcePathStyle(v bool)`

SetForcePathStyle sets ForcePathStyle field to given value.


### GetName

`func (o *S3StorageBackendConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3StorageBackendConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3StorageBackendConfigResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPartBufferSize

`func (o *S3StorageBackendConfigResponse) GetPartBufferSize() string`

GetPartBufferSize returns the PartBufferSize field if non-nil, zero value otherwise.

### GetPartBufferSizeOk

`func (o *S3StorageBackendConfigResponse) GetPartBufferSizeOk() (*string, bool)`

GetPartBufferSizeOk returns a tuple with the PartBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartBufferSize

`func (o *S3StorageBackendConfigResponse) SetPartBufferSize(v string)`

SetPartBufferSize sets PartBufferSize field to given value.


### GetProxy

`func (o *S3StorageBackendConfigResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *S3StorageBackendConfigResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *S3StorageBackendConfigResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *S3StorageBackendConfigResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetRegion

`func (o *S3StorageBackendConfigResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3StorageBackendConfigResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3StorageBackendConfigResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *S3StorageBackendConfigResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleArn

`func (o *S3StorageBackendConfigResponse) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *S3StorageBackendConfigResponse) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *S3StorageBackendConfigResponse) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *S3StorageBackendConfigResponse) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetTimeout

`func (o *S3StorageBackendConfigResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *S3StorageBackendConfigResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *S3StorageBackendConfigResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *S3StorageBackendConfigResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *S3StorageBackendConfigResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *S3StorageBackendConfigResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3StorageBackendConfigResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3StorageBackendConfigResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


