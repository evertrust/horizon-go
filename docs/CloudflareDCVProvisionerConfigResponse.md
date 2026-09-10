# CloudflareDCVProvisionerConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the credentials configuration holding the Cloudflare API key | 
**DelegationZone** | Pointer to **NullableString** | DNS zone used for CNAME delegation when provisioning DCV challenges | [optional] 
**Endpoint** | **string** | Cloudflare API endpoint URL | 
**Name** | **string** | Unique name of the DCV provisioner configuration | 
**Proxy** | Pointer to **NullableString** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | Pointer to **NullableString** | Request timeout | [optional] 
**Ttl** | **NullableString** | TTL for DNS TXT records provisioned during DCV challenges | 
**Type** | **string** | Provisioner type discriminator | 
**ZoneIdMappings** | [**[]ZoneIdMappings**](ZoneIdMappings.md) | A set of Cloudflare DNS zone ID with regex | [default to []]

## Methods

### NewCloudflareDCVProvisionerConfigResponse

`func NewCloudflareDCVProvisionerConfigResponse(id string, credentials string, endpoint string, name string, ttl NullableString, type_ string, zoneIdMappings []ZoneIdMappings, ) *CloudflareDCVProvisionerConfigResponse`

NewCloudflareDCVProvisionerConfigResponse instantiates a new CloudflareDCVProvisionerConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflareDCVProvisionerConfigResponseWithDefaults

`func NewCloudflareDCVProvisionerConfigResponseWithDefaults() *CloudflareDCVProvisionerConfigResponse`

NewCloudflareDCVProvisionerConfigResponseWithDefaults instantiates a new CloudflareDCVProvisionerConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudflareDCVProvisionerConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudflareDCVProvisionerConfigResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *CloudflareDCVProvisionerConfigResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CloudflareDCVProvisionerConfigResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDelegationZone

`func (o *CloudflareDCVProvisionerConfigResponse) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *CloudflareDCVProvisionerConfigResponse) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *CloudflareDCVProvisionerConfigResponse) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *CloudflareDCVProvisionerConfigResponse) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *CloudflareDCVProvisionerConfigResponse) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetEndpoint

`func (o *CloudflareDCVProvisionerConfigResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CloudflareDCVProvisionerConfigResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *CloudflareDCVProvisionerConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudflareDCVProvisionerConfigResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *CloudflareDCVProvisionerConfigResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CloudflareDCVProvisionerConfigResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CloudflareDCVProvisionerConfigResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CloudflareDCVProvisionerConfigResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CloudflareDCVProvisionerConfigResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *CloudflareDCVProvisionerConfigResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CloudflareDCVProvisionerConfigResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CloudflareDCVProvisionerConfigResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *CloudflareDCVProvisionerConfigResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *CloudflareDCVProvisionerConfigResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *CloudflareDCVProvisionerConfigResponse) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CloudflareDCVProvisionerConfigResponse) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *CloudflareDCVProvisionerConfigResponse) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *CloudflareDCVProvisionerConfigResponse) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *CloudflareDCVProvisionerConfigResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudflareDCVProvisionerConfigResponse) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *CloudflareDCVProvisionerConfigResponse) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *CloudflareDCVProvisionerConfigResponse) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *CloudflareDCVProvisionerConfigResponse) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


