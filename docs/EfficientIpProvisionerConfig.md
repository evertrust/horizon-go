# EfficientIpProvisionerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **string** | Name of the credentials configuration holding the EfficientIP login and password | 
**DelegationZone** | Pointer to **NullableString** | DNS zone used for CNAME delegation when provisioning DCV challenges | [optional] 
**DnsName** | **string** | Name of the DNS server on the SOLIDserver to use for provisioning | 
**DnsView** | Pointer to **NullableString** | DNS view to target on the SOLIDserver | [optional] 
**Endpoint** | **string** | EfficientIP SOLIDserver API endpoint URL | 
**Name** | **string** | Unique name of the DCV provisioner configuration | 
**Proxy** | Pointer to **NullableString** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | Pointer to **NullableString** | Request timeout | [optional] 
**Ttl** | **NullableString** | DNS record cache duration | 
**Type** | **string** | Provisioner type discriminator | 
**ZoneIdMappings** | Pointer to [**[]ZoneIdMappings**](ZoneIdMappings.md) | A set of DNS zone ID with regex | [optional] [default to []]

## Methods

### NewEfficientIpProvisionerConfig

`func NewEfficientIpProvisionerConfig(credentials string, dnsName string, endpoint string, name string, ttl NullableString, type_ string, ) *EfficientIpProvisionerConfig`

NewEfficientIpProvisionerConfig instantiates a new EfficientIpProvisionerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEfficientIpProvisionerConfigWithDefaults

`func NewEfficientIpProvisionerConfigWithDefaults() *EfficientIpProvisionerConfig`

NewEfficientIpProvisionerConfigWithDefaults instantiates a new EfficientIpProvisionerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *EfficientIpProvisionerConfig) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *EfficientIpProvisionerConfig) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *EfficientIpProvisionerConfig) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDelegationZone

`func (o *EfficientIpProvisionerConfig) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *EfficientIpProvisionerConfig) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *EfficientIpProvisionerConfig) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *EfficientIpProvisionerConfig) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *EfficientIpProvisionerConfig) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *EfficientIpProvisionerConfig) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetDnsName

`func (o *EfficientIpProvisionerConfig) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *EfficientIpProvisionerConfig) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *EfficientIpProvisionerConfig) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetDnsView

`func (o *EfficientIpProvisionerConfig) GetDnsView() string`

GetDnsView returns the DnsView field if non-nil, zero value otherwise.

### GetDnsViewOk

`func (o *EfficientIpProvisionerConfig) GetDnsViewOk() (*string, bool)`

GetDnsViewOk returns a tuple with the DnsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsView

`func (o *EfficientIpProvisionerConfig) SetDnsView(v string)`

SetDnsView sets DnsView field to given value.

### HasDnsView

`func (o *EfficientIpProvisionerConfig) HasDnsView() bool`

HasDnsView returns a boolean if a field has been set.

### SetDnsViewNil

`func (o *EfficientIpProvisionerConfig) SetDnsViewNil(b bool)`

 SetDnsViewNil sets the value for DnsView to be an explicit nil

### UnsetDnsView
`func (o *EfficientIpProvisionerConfig) UnsetDnsView()`

UnsetDnsView ensures that no value is present for DnsView, not even an explicit nil
### GetEndpoint

`func (o *EfficientIpProvisionerConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EfficientIpProvisionerConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EfficientIpProvisionerConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *EfficientIpProvisionerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EfficientIpProvisionerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EfficientIpProvisionerConfig) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *EfficientIpProvisionerConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *EfficientIpProvisionerConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *EfficientIpProvisionerConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *EfficientIpProvisionerConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *EfficientIpProvisionerConfig) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *EfficientIpProvisionerConfig) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *EfficientIpProvisionerConfig) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EfficientIpProvisionerConfig) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EfficientIpProvisionerConfig) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EfficientIpProvisionerConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *EfficientIpProvisionerConfig) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *EfficientIpProvisionerConfig) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *EfficientIpProvisionerConfig) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *EfficientIpProvisionerConfig) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *EfficientIpProvisionerConfig) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *EfficientIpProvisionerConfig) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *EfficientIpProvisionerConfig) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *EfficientIpProvisionerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EfficientIpProvisionerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EfficientIpProvisionerConfig) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *EfficientIpProvisionerConfig) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *EfficientIpProvisionerConfig) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *EfficientIpProvisionerConfig) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.

### HasZoneIdMappings

`func (o *EfficientIpProvisionerConfig) HasZoneIdMappings() bool`

HasZoneIdMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


