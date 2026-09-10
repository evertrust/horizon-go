# AzurednsDCVProvisionerConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
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

### NewAzurednsDCVProvisionerConfigResponse

`func NewAzurednsDCVProvisionerConfigResponse(id string, name string, resourceGroupName string, subscriptionId string, tenantId string, ttl NullableString, type_ string, ) *AzurednsDCVProvisionerConfigResponse`

NewAzurednsDCVProvisionerConfigResponse instantiates a new AzurednsDCVProvisionerConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurednsDCVProvisionerConfigResponseWithDefaults

`func NewAzurednsDCVProvisionerConfigResponseWithDefaults() *AzurednsDCVProvisionerConfigResponse`

NewAzurednsDCVProvisionerConfigResponseWithDefaults instantiates a new AzurednsDCVProvisionerConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzurednsDCVProvisionerConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzurednsDCVProvisionerConfigResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAuthorityHost

`func (o *AzurednsDCVProvisionerConfigResponse) GetAuthorityHost() string`

GetAuthorityHost returns the AuthorityHost field if non-nil, zero value otherwise.

### GetAuthorityHostOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetAuthorityHostOk() (*string, bool)`

GetAuthorityHostOk returns a tuple with the AuthorityHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityHost

`func (o *AzurednsDCVProvisionerConfigResponse) SetAuthorityHost(v string)`

SetAuthorityHost sets AuthorityHost field to given value.

### HasAuthorityHost

`func (o *AzurednsDCVProvisionerConfigResponse) HasAuthorityHost() bool`

HasAuthorityHost returns a boolean if a field has been set.

### SetAuthorityHostNil

`func (o *AzurednsDCVProvisionerConfigResponse) SetAuthorityHostNil(b bool)`

 SetAuthorityHostNil sets the value for AuthorityHost to be an explicit nil

### UnsetAuthorityHost
`func (o *AzurednsDCVProvisionerConfigResponse) UnsetAuthorityHost()`

UnsetAuthorityHost ensures that no value is present for AuthorityHost, not even an explicit nil
### GetCredentials

`func (o *AzurednsDCVProvisionerConfigResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AzurednsDCVProvisionerConfigResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AzurednsDCVProvisionerConfigResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *AzurednsDCVProvisionerConfigResponse) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *AzurednsDCVProvisionerConfigResponse) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetDelegationZone

`func (o *AzurednsDCVProvisionerConfigResponse) GetDelegationZone() string`

GetDelegationZone returns the DelegationZone field if non-nil, zero value otherwise.

### GetDelegationZoneOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetDelegationZoneOk() (*string, bool)`

GetDelegationZoneOk returns a tuple with the DelegationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationZone

`func (o *AzurednsDCVProvisionerConfigResponse) SetDelegationZone(v string)`

SetDelegationZone sets DelegationZone field to given value.

### HasDelegationZone

`func (o *AzurednsDCVProvisionerConfigResponse) HasDelegationZone() bool`

HasDelegationZone returns a boolean if a field has been set.

### SetDelegationZoneNil

`func (o *AzurednsDCVProvisionerConfigResponse) SetDelegationZoneNil(b bool)`

 SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil

### UnsetDelegationZone
`func (o *AzurednsDCVProvisionerConfigResponse) UnsetDelegationZone()`

UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
### GetEndpoint

`func (o *AzurednsDCVProvisionerConfigResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AzurednsDCVProvisionerConfigResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AzurednsDCVProvisionerConfigResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *AzurednsDCVProvisionerConfigResponse) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *AzurednsDCVProvisionerConfigResponse) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetName

`func (o *AzurednsDCVProvisionerConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzurednsDCVProvisionerConfigResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *AzurednsDCVProvisionerConfigResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AzurednsDCVProvisionerConfigResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AzurednsDCVProvisionerConfigResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *AzurednsDCVProvisionerConfigResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *AzurednsDCVProvisionerConfigResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetResourceGroupName

`func (o *AzurednsDCVProvisionerConfigResponse) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzurednsDCVProvisionerConfigResponse) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetSubscriptionId

`func (o *AzurednsDCVProvisionerConfigResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzurednsDCVProvisionerConfigResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *AzurednsDCVProvisionerConfigResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzurednsDCVProvisionerConfigResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTimeout

`func (o *AzurednsDCVProvisionerConfigResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AzurednsDCVProvisionerConfigResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AzurednsDCVProvisionerConfigResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *AzurednsDCVProvisionerConfigResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *AzurednsDCVProvisionerConfigResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetTtl

`func (o *AzurednsDCVProvisionerConfigResponse) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AzurednsDCVProvisionerConfigResponse) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *AzurednsDCVProvisionerConfigResponse) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *AzurednsDCVProvisionerConfigResponse) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *AzurednsDCVProvisionerConfigResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzurednsDCVProvisionerConfigResponse) SetType(v string)`

SetType sets Type field to given value.


### GetZoneIdMappings

`func (o *AzurednsDCVProvisionerConfigResponse) GetZoneIdMappings() []ZoneIdMappings`

GetZoneIdMappings returns the ZoneIdMappings field if non-nil, zero value otherwise.

### GetZoneIdMappingsOk

`func (o *AzurednsDCVProvisionerConfigResponse) GetZoneIdMappingsOk() (*[]ZoneIdMappings, bool)`

GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdMappings

`func (o *AzurednsDCVProvisionerConfigResponse) SetZoneIdMappings(v []ZoneIdMappings)`

SetZoneIdMappings sets ZoneIdMappings field to given value.

### HasZoneIdMappings

`func (o *AzurednsDCVProvisionerConfigResponse) HasZoneIdMappings() bool`

HasZoneIdMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


