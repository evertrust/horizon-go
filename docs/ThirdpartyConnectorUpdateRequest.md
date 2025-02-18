# ThirdpartyConnectorUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**ThrottleDuration** | **string** |  | 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Region** | **string** |  | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/api.security.credentials) containing login DN and password. | 
**ResourceGroupName** | Pointer to **NullableString** |  | [optional] 
**RoleArn** | Pointer to **NullableString** |  | [optional] 
**TagKey** | Pointer to **NullableString** |  | [optional] 
**TagValue** | Pointer to **NullableString** |  | [optional] 
**ThrottleParallelism** | **int64** |  | 
**MaxStoredCertificatePerHolder** | Pointer to **NullableInt64** |  | [optional] 
**BigIPHostname** | **string** |  | 
**Partition** | Pointer to **NullableString** |  | [optional] 
**SslParent** | Pointer to **NullableString** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**CipherGroup** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**KeyType** | Pointer to **NullableString** | Any of &#x60;rsa-2048&#x60;, &#x60;rsa-3072&#x60;, &#x60;rsa-4096&#x60;, &#x60;rsa-8192&#x60;, &#x60;ec-secp256r1&#x60;, &#x60;ec-secp384r1&#x60;, &#x60;ec-secp521r1&#x60;, &#x60;ed-448&#x60;, &#x60;ed-25519&#x60;, &#x60;mldsa-44&#x60;, &#x60;mldsa-65&#x60;, &#x60;mldsa-87&#x60; or &#x60;&lt;primary key type&gt;+&lt;alternate key type&gt;&#x60; | [optional] 
**Hostname** | **string** |  | 
**LoginProvider** | Pointer to **NullableString** |  | [optional] 
**Tenant** | **string** |  | 
**IntuneResourceUrl** | Pointer to **NullableString** |  | [optional] 
**OsQueryString** | Pointer to **NullableString** |  | [optional] 
**LegacyRevocationMode** | **bool** |  | 
**Endpoint** | **string** |  | 
**Port** | Pointer to **NullableInt64** |  | [optional] 
**BaseDn** | **string** |  | 
**Filter** | Pointer to **NullableString** |  | [optional] 
**PubKey** | **string** |  | 
**KeyName** | **string** |  | 
**ProviderName** | Pointer to **NullableString** |  | [optional] 
**IntendedPurpose** | Pointer to **NullableString** |  | [optional] 
**SearchFilter** | Pointer to **NullableString** |  | [optional] 
**VaultBaseUrl** | **string** |  | 
**Project** | **string** |  | 
**Location** | **string** |  | 
**CertAttr** | Pointer to **NullableString** |  | [optional] 
**FollowReferrals** | Pointer to **NullableBool** |  | [optional] 
**UserIdentifierAttribute** | **string** |  | 
**CertificateAttribute** | **string** |  | 

## Methods

### NewThirdpartyConnectorUpdateRequest

`func NewThirdpartyConnectorUpdateRequest(type_ string, name string, throttleDuration string, region string, credentials string, throttleParallelism int64, bigIPHostname string, hostname string, tenant string, legacyRevocationMode bool, endpoint string, baseDn string, pubKey string, keyName string, vaultBaseUrl string, project string, location string, userIdentifierAttribute string, certificateAttribute string, ) *ThirdpartyConnectorUpdateRequest`

NewThirdpartyConnectorUpdateRequest instantiates a new ThirdpartyConnectorUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdpartyConnectorUpdateRequestWithDefaults

`func NewThirdpartyConnectorUpdateRequestWithDefaults() *ThirdpartyConnectorUpdateRequest`

NewThirdpartyConnectorUpdateRequestWithDefaults instantiates a new ThirdpartyConnectorUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThirdpartyConnectorUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThirdpartyConnectorUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThirdpartyConnectorUpdateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ThirdpartyConnectorUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThirdpartyConnectorUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThirdpartyConnectorUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetThrottleDuration

`func (o *ThirdpartyConnectorUpdateRequest) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *ThirdpartyConnectorUpdateRequest) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *ThirdpartyConnectorUpdateRequest) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetRenewalPeriod

`func (o *ThirdpartyConnectorUpdateRequest) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *ThirdpartyConnectorUpdateRequest) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *ThirdpartyConnectorUpdateRequest) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *ThirdpartyConnectorUpdateRequest) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *ThirdpartyConnectorUpdateRequest) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *ThirdpartyConnectorUpdateRequest) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetTimeout

`func (o *ThirdpartyConnectorUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ThirdpartyConnectorUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ThirdpartyConnectorUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ThirdpartyConnectorUpdateRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *ThirdpartyConnectorUpdateRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *ThirdpartyConnectorUpdateRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *ThirdpartyConnectorUpdateRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ThirdpartyConnectorUpdateRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ThirdpartyConnectorUpdateRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ThirdpartyConnectorUpdateRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *ThirdpartyConnectorUpdateRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *ThirdpartyConnectorUpdateRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRegion

`func (o *ThirdpartyConnectorUpdateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ThirdpartyConnectorUpdateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ThirdpartyConnectorUpdateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCredentials

`func (o *ThirdpartyConnectorUpdateRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ThirdpartyConnectorUpdateRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ThirdpartyConnectorUpdateRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetResourceGroupName

`func (o *ThirdpartyConnectorUpdateRequest) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ThirdpartyConnectorUpdateRequest) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ThirdpartyConnectorUpdateRequest) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *ThirdpartyConnectorUpdateRequest) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *ThirdpartyConnectorUpdateRequest) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *ThirdpartyConnectorUpdateRequest) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetRoleArn

`func (o *ThirdpartyConnectorUpdateRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ThirdpartyConnectorUpdateRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ThirdpartyConnectorUpdateRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *ThirdpartyConnectorUpdateRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *ThirdpartyConnectorUpdateRequest) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *ThirdpartyConnectorUpdateRequest) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetTagKey

`func (o *ThirdpartyConnectorUpdateRequest) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *ThirdpartyConnectorUpdateRequest) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *ThirdpartyConnectorUpdateRequest) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.

### HasTagKey

`func (o *ThirdpartyConnectorUpdateRequest) HasTagKey() bool`

HasTagKey returns a boolean if a field has been set.

### SetTagKeyNil

`func (o *ThirdpartyConnectorUpdateRequest) SetTagKeyNil(b bool)`

 SetTagKeyNil sets the value for TagKey to be an explicit nil

### UnsetTagKey
`func (o *ThirdpartyConnectorUpdateRequest) UnsetTagKey()`

UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil
### GetTagValue

`func (o *ThirdpartyConnectorUpdateRequest) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *ThirdpartyConnectorUpdateRequest) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *ThirdpartyConnectorUpdateRequest) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.

### HasTagValue

`func (o *ThirdpartyConnectorUpdateRequest) HasTagValue() bool`

HasTagValue returns a boolean if a field has been set.

### SetTagValueNil

`func (o *ThirdpartyConnectorUpdateRequest) SetTagValueNil(b bool)`

 SetTagValueNil sets the value for TagValue to be an explicit nil

### UnsetTagValue
`func (o *ThirdpartyConnectorUpdateRequest) UnsetTagValue()`

UnsetTagValue ensures that no value is present for TagValue, not even an explicit nil
### GetThrottleParallelism

`func (o *ThirdpartyConnectorUpdateRequest) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *ThirdpartyConnectorUpdateRequest) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *ThirdpartyConnectorUpdateRequest) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetMaxStoredCertificatePerHolder

`func (o *ThirdpartyConnectorUpdateRequest) GetMaxStoredCertificatePerHolder() int64`

GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field if non-nil, zero value otherwise.

### GetMaxStoredCertificatePerHolderOk

`func (o *ThirdpartyConnectorUpdateRequest) GetMaxStoredCertificatePerHolderOk() (*int64, bool)`

GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredCertificatePerHolder

`func (o *ThirdpartyConnectorUpdateRequest) SetMaxStoredCertificatePerHolder(v int64)`

SetMaxStoredCertificatePerHolder sets MaxStoredCertificatePerHolder field to given value.

### HasMaxStoredCertificatePerHolder

`func (o *ThirdpartyConnectorUpdateRequest) HasMaxStoredCertificatePerHolder() bool`

HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.

### SetMaxStoredCertificatePerHolderNil

`func (o *ThirdpartyConnectorUpdateRequest) SetMaxStoredCertificatePerHolderNil(b bool)`

 SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil

### UnsetMaxStoredCertificatePerHolder
`func (o *ThirdpartyConnectorUpdateRequest) UnsetMaxStoredCertificatePerHolder()`

UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
### GetBigIPHostname

`func (o *ThirdpartyConnectorUpdateRequest) GetBigIPHostname() string`

GetBigIPHostname returns the BigIPHostname field if non-nil, zero value otherwise.

### GetBigIPHostnameOk

`func (o *ThirdpartyConnectorUpdateRequest) GetBigIPHostnameOk() (*string, bool)`

GetBigIPHostnameOk returns a tuple with the BigIPHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigIPHostname

`func (o *ThirdpartyConnectorUpdateRequest) SetBigIPHostname(v string)`

SetBigIPHostname sets BigIPHostname field to given value.


### GetPartition

`func (o *ThirdpartyConnectorUpdateRequest) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ThirdpartyConnectorUpdateRequest) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ThirdpartyConnectorUpdateRequest) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *ThirdpartyConnectorUpdateRequest) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### SetPartitionNil

`func (o *ThirdpartyConnectorUpdateRequest) SetPartitionNil(b bool)`

 SetPartitionNil sets the value for Partition to be an explicit nil

### UnsetPartition
`func (o *ThirdpartyConnectorUpdateRequest) UnsetPartition()`

UnsetPartition ensures that no value is present for Partition, not even an explicit nil
### GetSslParent

`func (o *ThirdpartyConnectorUpdateRequest) GetSslParent() string`

GetSslParent returns the SslParent field if non-nil, zero value otherwise.

### GetSslParentOk

`func (o *ThirdpartyConnectorUpdateRequest) GetSslParentOk() (*string, bool)`

GetSslParentOk returns a tuple with the SslParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslParent

`func (o *ThirdpartyConnectorUpdateRequest) SetSslParent(v string)`

SetSslParent sets SslParent field to given value.

### HasSslParent

`func (o *ThirdpartyConnectorUpdateRequest) HasSslParent() bool`

HasSslParent returns a boolean if a field has been set.

### SetSslParentNil

`func (o *ThirdpartyConnectorUpdateRequest) SetSslParentNil(b bool)`

 SetSslParentNil sets the value for SslParent to be an explicit nil

### UnsetSslParent
`func (o *ThirdpartyConnectorUpdateRequest) UnsetSslParent()`

UnsetSslParent ensures that no value is present for SslParent, not even an explicit nil
### GetPrefix

`func (o *ThirdpartyConnectorUpdateRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ThirdpartyConnectorUpdateRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ThirdpartyConnectorUpdateRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ThirdpartyConnectorUpdateRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *ThirdpartyConnectorUpdateRequest) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *ThirdpartyConnectorUpdateRequest) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetCipherGroup

`func (o *ThirdpartyConnectorUpdateRequest) GetCipherGroup() string`

GetCipherGroup returns the CipherGroup field if non-nil, zero value otherwise.

### GetCipherGroupOk

`func (o *ThirdpartyConnectorUpdateRequest) GetCipherGroupOk() (*string, bool)`

GetCipherGroupOk returns a tuple with the CipherGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherGroup

`func (o *ThirdpartyConnectorUpdateRequest) SetCipherGroup(v string)`

SetCipherGroup sets CipherGroup field to given value.

### HasCipherGroup

`func (o *ThirdpartyConnectorUpdateRequest) HasCipherGroup() bool`

HasCipherGroup returns a boolean if a field has been set.

### SetCipherGroupNil

`func (o *ThirdpartyConnectorUpdateRequest) SetCipherGroupNil(b bool)`

 SetCipherGroupNil sets the value for CipherGroup to be an explicit nil

### UnsetCipherGroup
`func (o *ThirdpartyConnectorUpdateRequest) UnsetCipherGroup()`

UnsetCipherGroup ensures that no value is present for CipherGroup, not even an explicit nil
### GetVersion

`func (o *ThirdpartyConnectorUpdateRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThirdpartyConnectorUpdateRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThirdpartyConnectorUpdateRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThirdpartyConnectorUpdateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ThirdpartyConnectorUpdateRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ThirdpartyConnectorUpdateRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetKeyType

`func (o *ThirdpartyConnectorUpdateRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ThirdpartyConnectorUpdateRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ThirdpartyConnectorUpdateRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ThirdpartyConnectorUpdateRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *ThirdpartyConnectorUpdateRequest) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *ThirdpartyConnectorUpdateRequest) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetHostname

`func (o *ThirdpartyConnectorUpdateRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ThirdpartyConnectorUpdateRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ThirdpartyConnectorUpdateRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetLoginProvider

`func (o *ThirdpartyConnectorUpdateRequest) GetLoginProvider() string`

GetLoginProvider returns the LoginProvider field if non-nil, zero value otherwise.

### GetLoginProviderOk

`func (o *ThirdpartyConnectorUpdateRequest) GetLoginProviderOk() (*string, bool)`

GetLoginProviderOk returns a tuple with the LoginProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginProvider

`func (o *ThirdpartyConnectorUpdateRequest) SetLoginProvider(v string)`

SetLoginProvider sets LoginProvider field to given value.

### HasLoginProvider

`func (o *ThirdpartyConnectorUpdateRequest) HasLoginProvider() bool`

HasLoginProvider returns a boolean if a field has been set.

### SetLoginProviderNil

`func (o *ThirdpartyConnectorUpdateRequest) SetLoginProviderNil(b bool)`

 SetLoginProviderNil sets the value for LoginProvider to be an explicit nil

### UnsetLoginProvider
`func (o *ThirdpartyConnectorUpdateRequest) UnsetLoginProvider()`

UnsetLoginProvider ensures that no value is present for LoginProvider, not even an explicit nil
### GetTenant

`func (o *ThirdpartyConnectorUpdateRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ThirdpartyConnectorUpdateRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ThirdpartyConnectorUpdateRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetIntuneResourceUrl

`func (o *ThirdpartyConnectorUpdateRequest) GetIntuneResourceUrl() string`

GetIntuneResourceUrl returns the IntuneResourceUrl field if non-nil, zero value otherwise.

### GetIntuneResourceUrlOk

`func (o *ThirdpartyConnectorUpdateRequest) GetIntuneResourceUrlOk() (*string, bool)`

GetIntuneResourceUrlOk returns a tuple with the IntuneResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntuneResourceUrl

`func (o *ThirdpartyConnectorUpdateRequest) SetIntuneResourceUrl(v string)`

SetIntuneResourceUrl sets IntuneResourceUrl field to given value.

### HasIntuneResourceUrl

`func (o *ThirdpartyConnectorUpdateRequest) HasIntuneResourceUrl() bool`

HasIntuneResourceUrl returns a boolean if a field has been set.

### SetIntuneResourceUrlNil

`func (o *ThirdpartyConnectorUpdateRequest) SetIntuneResourceUrlNil(b bool)`

 SetIntuneResourceUrlNil sets the value for IntuneResourceUrl to be an explicit nil

### UnsetIntuneResourceUrl
`func (o *ThirdpartyConnectorUpdateRequest) UnsetIntuneResourceUrl()`

UnsetIntuneResourceUrl ensures that no value is present for IntuneResourceUrl, not even an explicit nil
### GetOsQueryString

`func (o *ThirdpartyConnectorUpdateRequest) GetOsQueryString() string`

GetOsQueryString returns the OsQueryString field if non-nil, zero value otherwise.

### GetOsQueryStringOk

`func (o *ThirdpartyConnectorUpdateRequest) GetOsQueryStringOk() (*string, bool)`

GetOsQueryStringOk returns a tuple with the OsQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsQueryString

`func (o *ThirdpartyConnectorUpdateRequest) SetOsQueryString(v string)`

SetOsQueryString sets OsQueryString field to given value.

### HasOsQueryString

`func (o *ThirdpartyConnectorUpdateRequest) HasOsQueryString() bool`

HasOsQueryString returns a boolean if a field has been set.

### SetOsQueryStringNil

`func (o *ThirdpartyConnectorUpdateRequest) SetOsQueryStringNil(b bool)`

 SetOsQueryStringNil sets the value for OsQueryString to be an explicit nil

### UnsetOsQueryString
`func (o *ThirdpartyConnectorUpdateRequest) UnsetOsQueryString()`

UnsetOsQueryString ensures that no value is present for OsQueryString, not even an explicit nil
### GetLegacyRevocationMode

`func (o *ThirdpartyConnectorUpdateRequest) GetLegacyRevocationMode() bool`

GetLegacyRevocationMode returns the LegacyRevocationMode field if non-nil, zero value otherwise.

### GetLegacyRevocationModeOk

`func (o *ThirdpartyConnectorUpdateRequest) GetLegacyRevocationModeOk() (*bool, bool)`

GetLegacyRevocationModeOk returns a tuple with the LegacyRevocationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyRevocationMode

`func (o *ThirdpartyConnectorUpdateRequest) SetLegacyRevocationMode(v bool)`

SetLegacyRevocationMode sets LegacyRevocationMode field to given value.


### GetEndpoint

`func (o *ThirdpartyConnectorUpdateRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ThirdpartyConnectorUpdateRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ThirdpartyConnectorUpdateRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetPort

`func (o *ThirdpartyConnectorUpdateRequest) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ThirdpartyConnectorUpdateRequest) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ThirdpartyConnectorUpdateRequest) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ThirdpartyConnectorUpdateRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *ThirdpartyConnectorUpdateRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ThirdpartyConnectorUpdateRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetBaseDn

`func (o *ThirdpartyConnectorUpdateRequest) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *ThirdpartyConnectorUpdateRequest) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *ThirdpartyConnectorUpdateRequest) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetFilter

`func (o *ThirdpartyConnectorUpdateRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ThirdpartyConnectorUpdateRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ThirdpartyConnectorUpdateRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ThirdpartyConnectorUpdateRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ThirdpartyConnectorUpdateRequest) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ThirdpartyConnectorUpdateRequest) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetPubKey

`func (o *ThirdpartyConnectorUpdateRequest) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *ThirdpartyConnectorUpdateRequest) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *ThirdpartyConnectorUpdateRequest) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetKeyName

`func (o *ThirdpartyConnectorUpdateRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *ThirdpartyConnectorUpdateRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *ThirdpartyConnectorUpdateRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetProviderName

`func (o *ThirdpartyConnectorUpdateRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ThirdpartyConnectorUpdateRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ThirdpartyConnectorUpdateRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ThirdpartyConnectorUpdateRequest) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ThirdpartyConnectorUpdateRequest) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ThirdpartyConnectorUpdateRequest) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetIntendedPurpose

`func (o *ThirdpartyConnectorUpdateRequest) GetIntendedPurpose() string`

GetIntendedPurpose returns the IntendedPurpose field if non-nil, zero value otherwise.

### GetIntendedPurposeOk

`func (o *ThirdpartyConnectorUpdateRequest) GetIntendedPurposeOk() (*string, bool)`

GetIntendedPurposeOk returns a tuple with the IntendedPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedPurpose

`func (o *ThirdpartyConnectorUpdateRequest) SetIntendedPurpose(v string)`

SetIntendedPurpose sets IntendedPurpose field to given value.

### HasIntendedPurpose

`func (o *ThirdpartyConnectorUpdateRequest) HasIntendedPurpose() bool`

HasIntendedPurpose returns a boolean if a field has been set.

### SetIntendedPurposeNil

`func (o *ThirdpartyConnectorUpdateRequest) SetIntendedPurposeNil(b bool)`

 SetIntendedPurposeNil sets the value for IntendedPurpose to be an explicit nil

### UnsetIntendedPurpose
`func (o *ThirdpartyConnectorUpdateRequest) UnsetIntendedPurpose()`

UnsetIntendedPurpose ensures that no value is present for IntendedPurpose, not even an explicit nil
### GetSearchFilter

`func (o *ThirdpartyConnectorUpdateRequest) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *ThirdpartyConnectorUpdateRequest) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *ThirdpartyConnectorUpdateRequest) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *ThirdpartyConnectorUpdateRequest) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.

### SetSearchFilterNil

`func (o *ThirdpartyConnectorUpdateRequest) SetSearchFilterNil(b bool)`

 SetSearchFilterNil sets the value for SearchFilter to be an explicit nil

### UnsetSearchFilter
`func (o *ThirdpartyConnectorUpdateRequest) UnsetSearchFilter()`

UnsetSearchFilter ensures that no value is present for SearchFilter, not even an explicit nil
### GetVaultBaseUrl

`func (o *ThirdpartyConnectorUpdateRequest) GetVaultBaseUrl() string`

GetVaultBaseUrl returns the VaultBaseUrl field if non-nil, zero value otherwise.

### GetVaultBaseUrlOk

`func (o *ThirdpartyConnectorUpdateRequest) GetVaultBaseUrlOk() (*string, bool)`

GetVaultBaseUrlOk returns a tuple with the VaultBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultBaseUrl

`func (o *ThirdpartyConnectorUpdateRequest) SetVaultBaseUrl(v string)`

SetVaultBaseUrl sets VaultBaseUrl field to given value.


### GetProject

`func (o *ThirdpartyConnectorUpdateRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ThirdpartyConnectorUpdateRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ThirdpartyConnectorUpdateRequest) SetProject(v string)`

SetProject sets Project field to given value.


### GetLocation

`func (o *ThirdpartyConnectorUpdateRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ThirdpartyConnectorUpdateRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ThirdpartyConnectorUpdateRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCertAttr

`func (o *ThirdpartyConnectorUpdateRequest) GetCertAttr() string`

GetCertAttr returns the CertAttr field if non-nil, zero value otherwise.

### GetCertAttrOk

`func (o *ThirdpartyConnectorUpdateRequest) GetCertAttrOk() (*string, bool)`

GetCertAttrOk returns a tuple with the CertAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAttr

`func (o *ThirdpartyConnectorUpdateRequest) SetCertAttr(v string)`

SetCertAttr sets CertAttr field to given value.

### HasCertAttr

`func (o *ThirdpartyConnectorUpdateRequest) HasCertAttr() bool`

HasCertAttr returns a boolean if a field has been set.

### SetCertAttrNil

`func (o *ThirdpartyConnectorUpdateRequest) SetCertAttrNil(b bool)`

 SetCertAttrNil sets the value for CertAttr to be an explicit nil

### UnsetCertAttr
`func (o *ThirdpartyConnectorUpdateRequest) UnsetCertAttr()`

UnsetCertAttr ensures that no value is present for CertAttr, not even an explicit nil
### GetFollowReferrals

`func (o *ThirdpartyConnectorUpdateRequest) GetFollowReferrals() bool`

GetFollowReferrals returns the FollowReferrals field if non-nil, zero value otherwise.

### GetFollowReferralsOk

`func (o *ThirdpartyConnectorUpdateRequest) GetFollowReferralsOk() (*bool, bool)`

GetFollowReferralsOk returns a tuple with the FollowReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowReferrals

`func (o *ThirdpartyConnectorUpdateRequest) SetFollowReferrals(v bool)`

SetFollowReferrals sets FollowReferrals field to given value.

### HasFollowReferrals

`func (o *ThirdpartyConnectorUpdateRequest) HasFollowReferrals() bool`

HasFollowReferrals returns a boolean if a field has been set.

### SetFollowReferralsNil

`func (o *ThirdpartyConnectorUpdateRequest) SetFollowReferralsNil(b bool)`

 SetFollowReferralsNil sets the value for FollowReferrals to be an explicit nil

### UnsetFollowReferrals
`func (o *ThirdpartyConnectorUpdateRequest) UnsetFollowReferrals()`

UnsetFollowReferrals ensures that no value is present for FollowReferrals, not even an explicit nil
### GetUserIdentifierAttribute

`func (o *ThirdpartyConnectorUpdateRequest) GetUserIdentifierAttribute() string`

GetUserIdentifierAttribute returns the UserIdentifierAttribute field if non-nil, zero value otherwise.

### GetUserIdentifierAttributeOk

`func (o *ThirdpartyConnectorUpdateRequest) GetUserIdentifierAttributeOk() (*string, bool)`

GetUserIdentifierAttributeOk returns a tuple with the UserIdentifierAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifierAttribute

`func (o *ThirdpartyConnectorUpdateRequest) SetUserIdentifierAttribute(v string)`

SetUserIdentifierAttribute sets UserIdentifierAttribute field to given value.


### GetCertificateAttribute

`func (o *ThirdpartyConnectorUpdateRequest) GetCertificateAttribute() string`

GetCertificateAttribute returns the CertificateAttribute field if non-nil, zero value otherwise.

### GetCertificateAttributeOk

`func (o *ThirdpartyConnectorUpdateRequest) GetCertificateAttributeOk() (*string, bool)`

GetCertificateAttributeOk returns a tuple with the CertificateAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAttribute

`func (o *ThirdpartyConnectorUpdateRequest) SetCertificateAttribute(v string)`

SetCertificateAttribute sets CertificateAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


