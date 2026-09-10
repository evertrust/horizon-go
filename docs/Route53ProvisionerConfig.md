# Route53ProvisionerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **NullableString** | Name of the credentials configuration holding the AWS access key ID and secret | [optional] 
**DelegationZone** | Pointer to **NullableString** | DNS zone used for CNAME delegation when provisioning DCV challenges | [optional] 
**Endpoint** | Pointer to **NullableString** | Route 53 API endpoint URL override (for local testing) | [optional] 
**Name** | **string** | Unique name of the DCV provisioner configuration | 
**Proxy** | Pointer to **NullableString** | Name of the HTTP proxy configuration to use | [optional] 
**Region** | Pointer to **NullableString** | AWS region override | [optional] 
**RoleArn** | Pointer to **NullableString** | IAM role ARN to assume for cross-account Route 53 access | [optional] 
**Timeout** | Pointer to **NullableString** | Request timeout | [optional] 
**Ttl** | **NullableString** | DNS record cache duration | 
**Type** | **string** | Provisioner type discriminator | 
**ZoneIdMappings** | Pointer to [**[]ZoneIdMappings**](ZoneIdMappings.md) | A set of Route 53 hosted zone ID with regex | [optional] [default to []]

## Methods

### NewRoute53ProvisionerConfig

`func NewRoute53ProvisionerConfig(name string, ttl NullableString, type_ string, ) *Route53ProvisionerConfig`

NewRoute53ProvisionerConfig instantiates a new Route53ProvisionerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute53ProvisionerConfigWithDefaults

`func NewRoute53ProvisionerConfigWithDefaults() *Route53ProvisionerConfig`

NewRoute53ProvisionerConfigWithDefaults instantiates a new Route53ProvisionerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *Route53ProvisionerConfig) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Route53ProvisionerConfig) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Route53ProvisionerConfig) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Route53ProvisionerConfig) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *Route53ProvisionerConfig) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *Route53ProvisionerConfig) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetDelegationZone

`func (o *Route53ProvisionerConfig) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *Route53ProvisionerConfig) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *Route53ProvisionerConfig) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *Route53ProvisionerConfig) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *Route53ProvisionerConfig) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *Route53ProvisionerConfig) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetEndpoint

`func (o *Route53ProvisionerConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Route53ProvisionerConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Route53ProvisionerConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Route53ProvisionerConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *Route53ProvisionerConfig) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *Route53ProvisionerConfig) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetName

`func (o *Route53ProvisionerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route53ProvisionerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route53ProvisionerConfig) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *Route53ProvisionerConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *Route53ProvisionerConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *Route53ProvisionerConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *Route53ProvisionerConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *Route53ProvisionerConfig) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *Route53ProvisionerConfig) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRegion

`func (o *Route53ProvisionerConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Route53ProvisionerConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Route53ProvisionerConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Route53ProvisionerConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Route53ProvisionerConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Route53ProvisionerConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRoleArn

`func (o *Route53ProvisionerConfig) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *Route53ProvisionerConfig) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *Route53ProvisionerConfig) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *Route53ProvisionerConfig) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *Route53ProvisionerConfig) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *Route53ProvisionerConfig) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetTimeout

`func (o *Route53ProvisionerConfig) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Route53ProvisionerConfig) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Route53ProvisionerConfig) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Route53ProvisionerConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *Route53ProvisionerConfig) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *Route53ProvisionerConfig) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *Route53ProvisionerConfig) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Route53ProvisionerConfig) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Route53ProvisionerConfig) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *Route53ProvisionerConfig) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *Route53ProvisionerConfig) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *Route53ProvisionerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Route53ProvisionerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Route53ProvisionerConfig) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *Route53ProvisionerConfig) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *Route53ProvisionerConfig) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *Route53ProvisionerConfig) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.

### HasZoneIdMappings

`func (o *Route53ProvisionerConfig) HasZoneIdMappings() bool`

HasZoneIdMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


