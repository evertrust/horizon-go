# AzurednsDCVProvisionerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorityHost** | Pointer to **NullableString** | Azure AD authority host URL override | [optional] 
**Credentials** | Pointer to **NullableString** | Name of the credentials configuration holding the Azure client ID and secret | [optional] 
**DelegationZone** | Pointer to **NullableString** | DNS zone used for CNAME delegation when provisioning DCV challenges | [optional] 
**Endpoint** | Pointer to **NullableString** | Azure DNS API endpoint URL override | [optional] 
**Name** | **string** | Unique name of the DCV provisioner configuration | 
**Proxy** | Pointer to **NullableString** | Name of the HTTP proxy configuration to use | [optional] 
**ResourceGroupName** | **string** | Azure resource group containing the DNS zones | 
**SubscriptionId** | **string** | Azure subscription ID | 
**TenantId** | **string** | Azure Active Directory tenant ID | 
**Timeout** | Pointer to **NullableString** | Request timeout | [optional] 
**Ttl** | **NullableString** | DNS record cache duration | 
**Type** | **string** | Provisioner type discriminator | 
**ZoneIdMappings** | Pointer to [**[]ZoneIdMappings**](ZoneIdMappings.md) | A set of Azure DNS zone ID with regex | [optional] [default to []]

## Methods

### NewAzurednsDCVProvisionerConfig

`func NewAzurednsDCVProvisionerConfig(name string, resourceGroupName string, subscriptionId string, tenantId string, ttl NullableString, type_ string, ) *AzurednsDCVProvisionerConfig`

NewAzurednsDCVProvisionerConfig instantiates a new AzurednsDCVProvisionerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurednsDCVProvisionerConfigWithDefaults

`func NewAzurednsDCVProvisionerConfigWithDefaults() *AzurednsDCVProvisionerConfig`

NewAzurednsDCVProvisionerConfigWithDefaults instantiates a new AzurednsDCVProvisionerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorityHost

`func (o *AzurednsDCVProvisionerConfig) GetAuthorityHost() string`

GetAuthorityHost returns the AuthorityHost field if non-nil, zero value otherwise.

### GetAuthorityHostOk

`func (o *AzurednsDCVProvisionerConfig) GetAuthorityHostOk() (*string, bool)`

GetAuthorityHostOk returns a tuple with the AuthorityHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityHost

`func (o *AzurednsDCVProvisionerConfig) SetAuthorityHost(v string)`

SetAuthorityHost sets AuthorityHost field to given value.

### HasAuthorityHost

`func (o *AzurednsDCVProvisionerConfig) HasAuthorityHost() bool`

HasAuthorityHost returns a boolean if a field has been set.

### SetAuthorityHostNil

`func (o *AzurednsDCVProvisionerConfig) SetAuthorityHostNil(b bool)`

 SetAuthorityHostNil sets the value for AuthorityHost to be an explicit nil

### UnsetAuthorityHost
`func (o *AzurednsDCVProvisionerConfig) UnsetAuthorityHost()`

UnsetAuthorityHost ensures that no value is present for AuthorityHost, not even an explicit nil
### GetCredentials

`func (o *AzurednsDCVProvisionerConfig) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AzurednsDCVProvisionerConfig) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AzurednsDCVProvisionerConfig) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AzurednsDCVProvisionerConfig) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *AzurednsDCVProvisionerConfig) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *AzurednsDCVProvisionerConfig) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetDelegationZone

`func (o *AzurednsDCVProvisionerConfig) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *AzurednsDCVProvisionerConfig) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *AzurednsDCVProvisionerConfig) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *AzurednsDCVProvisionerConfig) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *AzurednsDCVProvisionerConfig) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *AzurednsDCVProvisionerConfig) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetEndpoint

`func (o *AzurednsDCVProvisionerConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AzurednsDCVProvisionerConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AzurednsDCVProvisionerConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AzurednsDCVProvisionerConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *AzurednsDCVProvisionerConfig) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *AzurednsDCVProvisionerConfig) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetName

`func (o *AzurednsDCVProvisionerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzurednsDCVProvisionerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzurednsDCVProvisionerConfig) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *AzurednsDCVProvisionerConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AzurednsDCVProvisionerConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AzurednsDCVProvisionerConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AzurednsDCVProvisionerConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *AzurednsDCVProvisionerConfig) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *AzurednsDCVProvisionerConfig) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetResourceGroupName

`func (o *AzurednsDCVProvisionerConfig) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzurednsDCVProvisionerConfig) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzurednsDCVProvisionerConfig) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetSubscriptionId

`func (o *AzurednsDCVProvisionerConfig) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzurednsDCVProvisionerConfig) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzurednsDCVProvisionerConfig) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *AzurednsDCVProvisionerConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzurednsDCVProvisionerConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzurednsDCVProvisionerConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTimeout

`func (o *AzurednsDCVProvisionerConfig) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AzurednsDCVProvisionerConfig) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AzurednsDCVProvisionerConfig) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AzurednsDCVProvisionerConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *AzurednsDCVProvisionerConfig) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *AzurednsDCVProvisionerConfig) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *AzurednsDCVProvisionerConfig) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AzurednsDCVProvisionerConfig) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AzurednsDCVProvisionerConfig) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *AzurednsDCVProvisionerConfig) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *AzurednsDCVProvisionerConfig) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *AzurednsDCVProvisionerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzurednsDCVProvisionerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzurednsDCVProvisionerConfig) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *AzurednsDCVProvisionerConfig) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *AzurednsDCVProvisionerConfig) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *AzurednsDCVProvisionerConfig) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.

### HasZoneIdMappings

`func (o *AzurednsDCVProvisionerConfig) HasZoneIdMappings() bool`

HasZoneIdMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


