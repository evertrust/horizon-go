# PowerDNSProvisionerConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the credentials configuration holding the PowerDNS API key | 
**DelegationZone** | Pointer to **NullableString** | DNS zone used for CNAME delegation when provisioning DCV challenges | [optional] 
**Endpoint** | **string** | PowerDNS API endpoint URL | 
**Name** | **string** | Unique name of the DCV provisioner configuration | 
**Proxy** | Pointer to **NullableString** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | Pointer to **NullableString** | Request timeout | [optional] 
**Ttl** | **NullableString** | DNS record cache duration | 
**Type** | **string** | Provisioner type discriminator | 
**ZoneIdMappings** | [**[]ZoneIdMappings**](ZoneIdMappings.md) | A set of DNS zone ID with regex | [default to []]

## Methods

### NewPowerDNSProvisionerConfigResponse

`func NewPowerDNSProvisionerConfigResponse(id string, credentials string, endpoint string, name string, ttl NullableString, type_ string, zoneIdMappings []ZoneIdMappings, ) *PowerDNSProvisionerConfigResponse`

NewPowerDNSProvisionerConfigResponse instantiates a new PowerDNSProvisionerConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerDNSProvisionerConfigResponseWithDefaults

`func NewPowerDNSProvisionerConfigResponseWithDefaults() *PowerDNSProvisionerConfigResponse`

NewPowerDNSProvisionerConfigResponseWithDefaults instantiates a new PowerDNSProvisionerConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerDNSProvisionerConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerDNSProvisionerConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerDNSProvisionerConfigResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *PowerDNSProvisionerConfigResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PowerDNSProvisionerConfigResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PowerDNSProvisionerConfigResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDelegationZone

`func (o *PowerDNSProvisionerConfigResponse) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *PowerDNSProvisionerConfigResponse) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *PowerDNSProvisionerConfigResponse) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *PowerDNSProvisionerConfigResponse) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *PowerDNSProvisionerConfigResponse) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *PowerDNSProvisionerConfigResponse) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetEndpoint

`func (o *PowerDNSProvisionerConfigResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PowerDNSProvisionerConfigResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PowerDNSProvisionerConfigResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *PowerDNSProvisionerConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerDNSProvisionerConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerDNSProvisionerConfigResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *PowerDNSProvisionerConfigResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PowerDNSProvisionerConfigResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PowerDNSProvisionerConfigResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *PowerDNSProvisionerConfigResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *PowerDNSProvisionerConfigResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *PowerDNSProvisionerConfigResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *PowerDNSProvisionerConfigResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PowerDNSProvisionerConfigResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PowerDNSProvisionerConfigResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PowerDNSProvisionerConfigResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *PowerDNSProvisionerConfigResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *PowerDNSProvisionerConfigResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *PowerDNSProvisionerConfigResponse) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PowerDNSProvisionerConfigResponse) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PowerDNSProvisionerConfigResponse) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *PowerDNSProvisionerConfigResponse) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *PowerDNSProvisionerConfigResponse) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *PowerDNSProvisionerConfigResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PowerDNSProvisionerConfigResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PowerDNSProvisionerConfigResponse) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *PowerDNSProvisionerConfigResponse) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *PowerDNSProvisionerConfigResponse) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *PowerDNSProvisionerConfigResponse) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


