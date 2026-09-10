# DcvProvisionerList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **NullableString** | Name of the credentials configuration holding the AWS access key ID and secret | 
**DelegationZone** | Pointer to **NullableString** | DNS zone used for CNAME delegation when provisioning DCV challenges | [optional] 
**Endpoint** | **NullableString** | Route 53 API endpoint URL override (for local testing) | 
**Name** | **string** | Unique name of the DCV provisioner configuration | 
**Proxy** | Pointer to **NullableString** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | Pointer to **NullableString** | Request timeout | [optional] 
**Ttl** | **NullableString** | DNS record cache duration | 
**Type** | **string** | Provisioner type discriminator | 
**ZoneIdMappings** | [**[]ZoneIdMappings**](ZoneIdMappings.md) | A set of Route 53 hosted zone ID with regex | [default to []]
**AuthorityHost** | Pointer to **NullableString** | Azure AD authority host URL override | [optional] 
**ResourceGroupName** | **string** | Azure resource group containing the DNS zones | 
**SubscriptionId** | **string** | Azure subscription ID | 
**TenantId** | **string** | Azure Active Directory tenant ID | 
**DnsName** | **string** | Name of the DNS server on the SOLIDserver to use for provisioning | 
**DnsView** | Pointer to **NullableString** | DNS view to target on the SOLIDserver | [optional] 
**Region** | Pointer to **NullableString** | AWS region override | [optional] 
**RoleArn** | Pointer to **NullableString** | IAM role ARN to assume for cross-account Route 53 access | [optional] 

## Methods

### NewDcvProvisionerList200ResponseInner

`func NewDcvProvisionerList200ResponseInner(id string, credentials NullableString, endpoint NullableString, name string, ttl NullableString, type_ string, zoneIdMappings []ZoneIdMappings, resourceGroupName string, subscriptionId string, tenantId string, dnsName string, ) *DcvProvisionerList200ResponseInner`

NewDcvProvisionerList200ResponseInner instantiates a new DcvProvisionerList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcvProvisionerList200ResponseInnerWithDefaults

`func NewDcvProvisionerList200ResponseInnerWithDefaults() *DcvProvisionerList200ResponseInner`

NewDcvProvisionerList200ResponseInnerWithDefaults instantiates a new DcvProvisionerList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DcvProvisionerList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DcvProvisionerList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DcvProvisionerList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *DcvProvisionerList200ResponseInner) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DcvProvisionerList200ResponseInner) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DcvProvisionerList200ResponseInner) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### SetCredentialsNil

`func (o *DcvProvisionerList200ResponseInner) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *DcvProvisionerList200ResponseInner) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetDelegationZone

`func (o *DcvProvisionerList200ResponseInner) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *DcvProvisionerList200ResponseInner) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *DcvProvisionerList200ResponseInner) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *DcvProvisionerList200ResponseInner) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *DcvProvisionerList200ResponseInner) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *DcvProvisionerList200ResponseInner) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetEndpoint

`func (o *DcvProvisionerList200ResponseInner) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DcvProvisionerList200ResponseInner) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DcvProvisionerList200ResponseInner) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *DcvProvisionerList200ResponseInner) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *DcvProvisionerList200ResponseInner) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetName

`func (o *DcvProvisionerList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcvProvisionerList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcvProvisionerList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *DcvProvisionerList200ResponseInner) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DcvProvisionerList200ResponseInner) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DcvProvisionerList200ResponseInner) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DcvProvisionerList200ResponseInner) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DcvProvisionerList200ResponseInner) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DcvProvisionerList200ResponseInner) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *DcvProvisionerList200ResponseInner) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DcvProvisionerList200ResponseInner) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DcvProvisionerList200ResponseInner) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DcvProvisionerList200ResponseInner) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *DcvProvisionerList200ResponseInner) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DcvProvisionerList200ResponseInner) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *DcvProvisionerList200ResponseInner) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DcvProvisionerList200ResponseInner) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DcvProvisionerList200ResponseInner) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *DcvProvisionerList200ResponseInner) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *DcvProvisionerList200ResponseInner) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *DcvProvisionerList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcvProvisionerList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcvProvisionerList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *DcvProvisionerList200ResponseInner) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *DcvProvisionerList200ResponseInner) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *DcvProvisionerList200ResponseInner) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.


### GetAuthorityHost

`func (o *DcvProvisionerList200ResponseInner) GetAuthorityHost() string`

GetAuthorityHost returns the AuthorityHost field if non-nil, zero value otherwise.

### GetAuthorityHostOk

`func (o *DcvProvisionerList200ResponseInner) GetAuthorityHostOk() (*string, bool)`

GetAuthorityHostOk returns a tuple with the AuthorityHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityHost

`func (o *DcvProvisionerList200ResponseInner) SetAuthorityHost(v string)`

SetAuthorityHost sets AuthorityHost field to given value.

### HasAuthorityHost

`func (o *DcvProvisionerList200ResponseInner) HasAuthorityHost() bool`

HasAuthorityHost returns a boolean if a field has been set.

### SetAuthorityHostNil

`func (o *DcvProvisionerList200ResponseInner) SetAuthorityHostNil(b bool)`

 SetAuthorityHostNil sets the value for AuthorityHost to be an explicit nil

### UnsetAuthorityHost
`func (o *DcvProvisionerList200ResponseInner) UnsetAuthorityHost()`

UnsetAuthorityHost ensures that no value is present for AuthorityHost, not even an explicit nil
### GetResourceGroupName

`func (o *DcvProvisionerList200ResponseInner) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *DcvProvisionerList200ResponseInner) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *DcvProvisionerList200ResponseInner) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetSubscriptionId

`func (o *DcvProvisionerList200ResponseInner) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DcvProvisionerList200ResponseInner) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DcvProvisionerList200ResponseInner) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *DcvProvisionerList200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DcvProvisionerList200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DcvProvisionerList200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetDnsName

`func (o *DcvProvisionerList200ResponseInner) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *DcvProvisionerList200ResponseInner) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *DcvProvisionerList200ResponseInner) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetDnsView

`func (o *DcvProvisionerList200ResponseInner) GetDnsView() string`

GetDnsView returns the DnsView field if non-nil, zero value otherwise.

### GetDnsViewOk

`func (o *DcvProvisionerList200ResponseInner) GetDnsViewOk() (*string, bool)`

GetDnsViewOk returns a tuple with the DnsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsView

`func (o *DcvProvisionerList200ResponseInner) SetDnsView(v string)`

SetDnsView sets DnsView field to given value.

### HasDnsView

`func (o *DcvProvisionerList200ResponseInner) HasDnsView() bool`

HasDnsView returns a boolean if a field has been set.

### SetDnsViewNil

`func (o *DcvProvisionerList200ResponseInner) SetDnsViewNil(b bool)`

 SetDnsViewNil sets the value for DnsView to be an explicit nil

### UnsetDnsView
`func (o *DcvProvisionerList200ResponseInner) UnsetDnsView()`

UnsetDnsView ensures that no value is present for DnsView, not even an explicit nil
### GetRegion

`func (o *DcvProvisionerList200ResponseInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DcvProvisionerList200ResponseInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DcvProvisionerList200ResponseInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DcvProvisionerList200ResponseInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *DcvProvisionerList200ResponseInner) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *DcvProvisionerList200ResponseInner) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRoleArn

`func (o *DcvProvisionerList200ResponseInner) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *DcvProvisionerList200ResponseInner) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *DcvProvisionerList200ResponseInner) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *DcvProvisionerList200ResponseInner) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *DcvProvisionerList200ResponseInner) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *DcvProvisionerList200ResponseInner) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


