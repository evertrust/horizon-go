# EfficientIpProvisionerConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
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

### NewEfficientIpProvisionerConfigResponse

`func NewEfficientIpProvisionerConfigResponse(id string, credentials string, dnsName string, endpoint string, name string, ttl NullableString, type_ string, ) *EfficientIpProvisionerConfigResponse`

NewEfficientIpProvisionerConfigResponse instantiates a new EfficientIpProvisionerConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEfficientIpProvisionerConfigResponseWithDefaults

`func NewEfficientIpProvisionerConfigResponseWithDefaults() *EfficientIpProvisionerConfigResponse`

NewEfficientIpProvisionerConfigResponseWithDefaults instantiates a new EfficientIpProvisionerConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EfficientIpProvisionerConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EfficientIpProvisionerConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EfficientIpProvisionerConfigResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *EfficientIpProvisionerConfigResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *EfficientIpProvisionerConfigResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *EfficientIpProvisionerConfigResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetDelegationZone

`func (o *EfficientIpProvisionerConfigResponse) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *EfficientIpProvisionerConfigResponse) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *EfficientIpProvisionerConfigResponse) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *EfficientIpProvisionerConfigResponse) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *EfficientIpProvisionerConfigResponse) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *EfficientIpProvisionerConfigResponse) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetDnsName

`func (o *EfficientIpProvisionerConfigResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *EfficientIpProvisionerConfigResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *EfficientIpProvisionerConfigResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetDnsView

`func (o *EfficientIpProvisionerConfigResponse) GetDnsView() string`

GetDnsView returns the DnsView field if non-nil, zero value otherwise.

### GetDnsViewOk

`func (o *EfficientIpProvisionerConfigResponse) GetDnsViewOk() (*string, bool)`

GetDnsViewOk returns a tuple with the DnsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsView

`func (o *EfficientIpProvisionerConfigResponse) SetDnsView(v string)`

SetDnsView sets DnsView field to given value.

### HasDnsView

`func (o *EfficientIpProvisionerConfigResponse) HasDnsView() bool`

HasDnsView returns a boolean if a field has been set.

### SetDnsViewNil

`func (o *EfficientIpProvisionerConfigResponse) SetDnsViewNil(b bool)`

 SetDnsViewNil sets the value for DnsView to be an explicit nil

### UnsetDnsView
`func (o *EfficientIpProvisionerConfigResponse) UnsetDnsView()`

UnsetDnsView ensures that no value is present for DnsView, not even an explicit nil
### GetEndpoint

`func (o *EfficientIpProvisionerConfigResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EfficientIpProvisionerConfigResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EfficientIpProvisionerConfigResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *EfficientIpProvisionerConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EfficientIpProvisionerConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EfficientIpProvisionerConfigResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *EfficientIpProvisionerConfigResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *EfficientIpProvisionerConfigResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *EfficientIpProvisionerConfigResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *EfficientIpProvisionerConfigResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *EfficientIpProvisionerConfigResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *EfficientIpProvisionerConfigResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *EfficientIpProvisionerConfigResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EfficientIpProvisionerConfigResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EfficientIpProvisionerConfigResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EfficientIpProvisionerConfigResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *EfficientIpProvisionerConfigResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *EfficientIpProvisionerConfigResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *EfficientIpProvisionerConfigResponse) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *EfficientIpProvisionerConfigResponse) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *EfficientIpProvisionerConfigResponse) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *EfficientIpProvisionerConfigResponse) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *EfficientIpProvisionerConfigResponse) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *EfficientIpProvisionerConfigResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EfficientIpProvisionerConfigResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EfficientIpProvisionerConfigResponse) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *EfficientIpProvisionerConfigResponse) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *EfficientIpProvisionerConfigResponse) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *EfficientIpProvisionerConfigResponse) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.

### HasZoneIdMappings

`func (o *EfficientIpProvisionerConfigResponse) HasZoneIdMappings() bool`

HasZoneIdMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


