# DcvProvisionerUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **NullableString** | Name of the credentials configuration holding the AWS access key ID and secret | 
**DelegationZone** | Pointer to **NullableString** | DNS zone used for CNAME delegation when provisioning DCV challenges | [optional] 
**Endpoint** | **NullableString** | Route 53 API endpoint URL override (for local testing) | 
**Name** | **string** | Unique name of the DCV provisioner configuration | 
**Proxy** | Pointer to **NullableString** | Name of the HTTP proxy configuration to use | [optional] 
**Timeout** | Pointer to **NullableString** | Request timeout | [optional] 
**Ttl** | **NullableString** | DNS record cache duration | 
**Type** | **string** | Provisioner type discriminator | 
**ZoneIdMappings** | Pointer to [**[]ZoneIdMappings**](ZoneIdMappings.md) | A set of Route 53 hosted zone ID with regex | [optional] [default to []]
**AuthorityHost** | Pointer to **NullableString** | Azure AD authority host URL override | [optional] 
**ResourceGroupName** | **string** | Azure resource group containing the DNS zones | 
**SubscriptionId** | **string** | Azure subscription ID | 
**TenantId** | **string** | Azure Active Directory tenant ID | 
**DnsName** | **string** | Name of the DNS server on the SOLIDserver to use for provisioning | 
**DnsView** | Pointer to **NullableString** | DNS view to target on the SOLIDserver | [optional] 
**Region** | Pointer to **NullableString** | AWS region override | [optional] 
**RoleArn** | Pointer to **NullableString** | IAM role ARN to assume for cross-account Route 53 access | [optional] 

## Methods

### NewDcvProvisionerUpdateRequest

`func NewDcvProvisionerUpdateRequest(credentials NullableString, endpoint NullableString, name string, ttl NullableString, type_ string, resourceGroupName string, subscriptionId string, tenantId string, dnsName string, ) *DcvProvisionerUpdateRequest`

NewDcvProvisionerUpdateRequest instantiates a new DcvProvisionerUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcvProvisionerUpdateRequestWithDefaults

`func NewDcvProvisionerUpdateRequestWithDefaults() *DcvProvisionerUpdateRequest`

NewDcvProvisionerUpdateRequestWithDefaults instantiates a new DcvProvisionerUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *DcvProvisionerUpdateRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DcvProvisionerUpdateRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DcvProvisionerUpdateRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### SetCredentialsNil

`func (o *DcvProvisionerUpdateRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *DcvProvisionerUpdateRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetDelegationZone

`func (o *DcvProvisionerUpdateRequest) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *DcvProvisionerUpdateRequest) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *DcvProvisionerUpdateRequest) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *DcvProvisionerUpdateRequest) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *DcvProvisionerUpdateRequest) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *DcvProvisionerUpdateRequest) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetEndpoint

`func (o *DcvProvisionerUpdateRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DcvProvisionerUpdateRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DcvProvisionerUpdateRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *DcvProvisionerUpdateRequest) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *DcvProvisionerUpdateRequest) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetName

`func (o *DcvProvisionerUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcvProvisionerUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcvProvisionerUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *DcvProvisionerUpdateRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DcvProvisionerUpdateRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DcvProvisionerUpdateRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DcvProvisionerUpdateRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DcvProvisionerUpdateRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DcvProvisionerUpdateRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *DcvProvisionerUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DcvProvisionerUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DcvProvisionerUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DcvProvisionerUpdateRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *DcvProvisionerUpdateRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DcvProvisionerUpdateRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *DcvProvisionerUpdateRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DcvProvisionerUpdateRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DcvProvisionerUpdateRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *DcvProvisionerUpdateRequest) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *DcvProvisionerUpdateRequest) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *DcvProvisionerUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcvProvisionerUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcvProvisionerUpdateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *DcvProvisionerUpdateRequest) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *DcvProvisionerUpdateRequest) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *DcvProvisionerUpdateRequest) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.

### HasZoneIdMappings

`func (o *DcvProvisionerUpdateRequest) HasZoneIdMappings() bool`

HasZoneIdMappings returns a boolean if a field has been set.

### GetAuthorityHost

`func (o *DcvProvisionerUpdateRequest) GetAuthorityHost() string`

GetAuthorityHost returns the AuthorityHost field if non-nil, zero value otherwise.

### GetAuthorityHostOk

`func (o *DcvProvisionerUpdateRequest) GetAuthorityHostOk() (*string, bool)`

GetAuthorityHostOk returns a tuple with the AuthorityHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityHost

`func (o *DcvProvisionerUpdateRequest) SetAuthorityHost(v string)`

SetAuthorityHost sets AuthorityHost field to given value.

### HasAuthorityHost

`func (o *DcvProvisionerUpdateRequest) HasAuthorityHost() bool`

HasAuthorityHost returns a boolean if a field has been set.

### SetAuthorityHostNil

`func (o *DcvProvisionerUpdateRequest) SetAuthorityHostNil(b bool)`

 SetAuthorityHostNil sets the value for AuthorityHost to be an explicit nil

### UnsetAuthorityHost
`func (o *DcvProvisionerUpdateRequest) UnsetAuthorityHost()`

UnsetAuthorityHost ensures that no value is present for AuthorityHost, not even an explicit nil
### GetResourceGroupName

`func (o *DcvProvisionerUpdateRequest) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *DcvProvisionerUpdateRequest) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *DcvProvisionerUpdateRequest) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetSubscriptionId

`func (o *DcvProvisionerUpdateRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DcvProvisionerUpdateRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DcvProvisionerUpdateRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *DcvProvisionerUpdateRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DcvProvisionerUpdateRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DcvProvisionerUpdateRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetDnsName

`func (o *DcvProvisionerUpdateRequest) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *DcvProvisionerUpdateRequest) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *DcvProvisionerUpdateRequest) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetDnsView

`func (o *DcvProvisionerUpdateRequest) GetDnsView() string`

GetDnsView returns the DnsView field if non-nil, zero value otherwise.

### GetDnsViewOk

`func (o *DcvProvisionerUpdateRequest) GetDnsViewOk() (*string, bool)`

GetDnsViewOk returns a tuple with the DnsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsView

`func (o *DcvProvisionerUpdateRequest) SetDnsView(v string)`

SetDnsView sets DnsView field to given value.

### HasDnsView

`func (o *DcvProvisionerUpdateRequest) HasDnsView() bool`

HasDnsView returns a boolean if a field has been set.

### SetDnsViewNil

`func (o *DcvProvisionerUpdateRequest) SetDnsViewNil(b bool)`

 SetDnsViewNil sets the value for DnsView to be an explicit nil

### UnsetDnsView
`func (o *DcvProvisionerUpdateRequest) UnsetDnsView()`

UnsetDnsView ensures that no value is present for DnsView, not even an explicit nil
### GetRegion

`func (o *DcvProvisionerUpdateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DcvProvisionerUpdateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DcvProvisionerUpdateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DcvProvisionerUpdateRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *DcvProvisionerUpdateRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *DcvProvisionerUpdateRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRoleArn

`func (o *DcvProvisionerUpdateRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *DcvProvisionerUpdateRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *DcvProvisionerUpdateRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *DcvProvisionerUpdateRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *DcvProvisionerUpdateRequest) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *DcvProvisionerUpdateRequest) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


