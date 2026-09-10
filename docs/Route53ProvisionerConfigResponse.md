# Route53ProvisionerConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
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
**ZoneIdMappings** | [**[]ZoneIdMappings**](ZoneIdMappings.md) | A set of Route 53 hosted zone ID with regex | [default to []]

## Methods

### NewRoute53ProvisionerConfigResponse

`func NewRoute53ProvisionerConfigResponse(id string, name string, ttl NullableString, type_ string, zoneIdMappings []ZoneIdMappings, ) *Route53ProvisionerConfigResponse`

NewRoute53ProvisionerConfigResponse instantiates a new Route53ProvisionerConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute53ProvisionerConfigResponseWithDefaults

`func NewRoute53ProvisionerConfigResponseWithDefaults() *Route53ProvisionerConfigResponse`

NewRoute53ProvisionerConfigResponseWithDefaults instantiates a new Route53ProvisionerConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Route53ProvisionerConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Route53ProvisionerConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Route53ProvisionerConfigResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *Route53ProvisionerConfigResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Route53ProvisionerConfigResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Route53ProvisionerConfigResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Route53ProvisionerConfigResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *Route53ProvisionerConfigResponse) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *Route53ProvisionerConfigResponse) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetDelegationZone

`func (o *Route53ProvisionerConfigResponse) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *Route53ProvisionerConfigResponse) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *Route53ProvisionerConfigResponse) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *Route53ProvisionerConfigResponse) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *Route53ProvisionerConfigResponse) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *Route53ProvisionerConfigResponse) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetEndpoint

`func (o *Route53ProvisionerConfigResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Route53ProvisionerConfigResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Route53ProvisionerConfigResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Route53ProvisionerConfigResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *Route53ProvisionerConfigResponse) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *Route53ProvisionerConfigResponse) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetName

`func (o *Route53ProvisionerConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route53ProvisionerConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route53ProvisionerConfigResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *Route53ProvisionerConfigResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *Route53ProvisionerConfigResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *Route53ProvisionerConfigResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *Route53ProvisionerConfigResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *Route53ProvisionerConfigResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *Route53ProvisionerConfigResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRegion

`func (o *Route53ProvisionerConfigResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Route53ProvisionerConfigResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Route53ProvisionerConfigResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Route53ProvisionerConfigResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Route53ProvisionerConfigResponse) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Route53ProvisionerConfigResponse) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRoleArn

`func (o *Route53ProvisionerConfigResponse) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *Route53ProvisionerConfigResponse) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *Route53ProvisionerConfigResponse) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *Route53ProvisionerConfigResponse) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *Route53ProvisionerConfigResponse) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *Route53ProvisionerConfigResponse) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetTimeout

`func (o *Route53ProvisionerConfigResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Route53ProvisionerConfigResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Route53ProvisionerConfigResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Route53ProvisionerConfigResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *Route53ProvisionerConfigResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *Route53ProvisionerConfigResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *Route53ProvisionerConfigResponse) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Route53ProvisionerConfigResponse) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Route53ProvisionerConfigResponse) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *Route53ProvisionerConfigResponse) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *Route53ProvisionerConfigResponse) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *Route53ProvisionerConfigResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Route53ProvisionerConfigResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Route53ProvisionerConfigResponse) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *Route53ProvisionerConfigResponse) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *Route53ProvisionerConfigResponse) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *Route53ProvisionerConfigResponse) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


