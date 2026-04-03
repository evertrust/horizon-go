# ThirdPartyConnectors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the account to authenticate on Netscaler | 
**Name** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Region** | **string** |  | 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**ResourceGroupName** | Pointer to **NullableString** |  | [optional] 
**RoleArn** | Pointer to **NullableString** |  | [optional] 
**TagKey** | Pointer to **NullableString** |  | [optional] 
**TagValue** | Pointer to **NullableString** |  | [optional] 
**ThrottleDuration** | **string** |  | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**AzureTenant** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**ThrottleParallelism** | **int64** |  | 
**VaultBaseUrl** | **string** |  | 
**Hostname** | **string** |  | 
**KeyType** | Pointer to **NullableString** | One of &#x60;rsa-2048&#x60;, &#x60;rsa-3072&#x60;, &#x60;rsa-4096&#x60;, &#x60;rsa-8192&#x60;, &#x60;ec-secp256r1&#x60;, &#x60;ec-secp384r1&#x60;, &#x60;ec-secp521r1&#x60;, &#x60;ed-448&#x60;, &#x60;ed-25519&#x60;, &#x60;mldsa-44&#x60;, &#x60;mldsa-65&#x60;, &#x60;mldsa-87&#x60;, &#x60;slhdsa-sha2-128s&#x60;, &#x60;slhdsa-sha2-128f&#x60;, &#x60;slhdsa-sha2-192s&#x60;, &#x60;slhdsa-sha2-192f&#x60;, &#x60;slhdsa-sha2-256s&#x60;, &#x60;slhdsa-sha2-256f&#x60;, &#x60;slhdsa-sha2-128ssha256&#x60;, &#x60;slhdsa-sha2-128fsha256&#x60;, &#x60;slhdsa-sha2-192ssha512&#x60;, &#x60;slhdsa-sha2-192fsha512&#x60;, &#x60;slhdsa-sha2-256ssha512&#x60;, &#x60;slhdsa-sha2-256fsha512&#x60; or &#x60;&lt;primary key type&gt;+&lt;alternate key type&gt;&#x60; | [optional] 
**LoginProvider** | Pointer to **NullableString** |  | [optional] 
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**WithChain** | Pointer to **NullableBool** | Enable the certificate trust chain to be pushed. | [optional] [default to true]
**BigIPHostname** | **string** |  | 
**CipherGroup** | Pointer to **NullableString** |  | [optional] 
**MaxStoredCertificatePerHolder** | Pointer to **NullableInt64** |  | [optional] 
**Partition** | Pointer to **NullableString** |  | [optional] 
**SslParent** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Location** | **string** |  | 
**Project** | **string** |  | 
**IntuneResourceUrl** | Pointer to **NullableString** |  | [optional] 
**LegacyRevocationMode** | **bool** |  | 
**OsQueryString** | Pointer to **NullableString** |  | [optional] 
**IntendedPurpose** | Pointer to **NullableString** |  | [optional] 
**KeyName** | **string** |  | 
**ProviderName** | Pointer to **NullableString** |  | [optional] 
**PubKey** | **string** |  | 
**SearchFilter** | Pointer to **NullableString** |  | [optional] 
**Endpoint** | **string** |  | 
**BaseDn** | **string** |  | 
**CertAttr** | Pointer to **NullableString** |  | [optional] 
**CertificateAttribute** | **string** |  | 
**Filter** | Pointer to **NullableString** |  | [optional] 
**FollowReferrals** | Pointer to **NullableBool** |  | [optional] 
**Port** | Pointer to **NullableInt64** |  | [optional] 
**UserIdentifierAttribute** | **string** |  | 
**CertificateStorePath** | **string** |  | 

## Methods

### NewThirdPartyConnectors

`func NewThirdPartyConnectors(credentials string, name string, region string, throttleDuration string, type_ string, throttleParallelism int64, vaultBaseUrl string, hostname string, bigIPHostname string, location string, project string, legacyRevocationMode bool, keyName string, pubKey string, endpoint string, baseDn string, certificateAttribute string, userIdentifierAttribute string, certificateStorePath string, ) *ThirdPartyConnectors`

NewThirdPartyConnectors instantiates a new ThirdPartyConnectors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyConnectorsWithDefaults

`func NewThirdPartyConnectorsWithDefaults() *ThirdPartyConnectors`

NewThirdPartyConnectorsWithDefaults instantiates a new ThirdPartyConnectors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ThirdPartyConnectors) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ThirdPartyConnectors) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ThirdPartyConnectors) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetName

`func (o *ThirdPartyConnectors) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThirdPartyConnectors) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThirdPartyConnectors) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *ThirdPartyConnectors) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ThirdPartyConnectors) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ThirdPartyConnectors) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ThirdPartyConnectors) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *ThirdPartyConnectors) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *ThirdPartyConnectors) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRegion

`func (o *ThirdPartyConnectors) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ThirdPartyConnectors) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ThirdPartyConnectors) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRenewalPeriod

`func (o *ThirdPartyConnectors) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *ThirdPartyConnectors) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *ThirdPartyConnectors) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *ThirdPartyConnectors) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *ThirdPartyConnectors) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *ThirdPartyConnectors) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetResourceGroupName

`func (o *ThirdPartyConnectors) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ThirdPartyConnectors) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ThirdPartyConnectors) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *ThirdPartyConnectors) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *ThirdPartyConnectors) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *ThirdPartyConnectors) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetRoleArn

`func (o *ThirdPartyConnectors) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ThirdPartyConnectors) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ThirdPartyConnectors) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *ThirdPartyConnectors) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *ThirdPartyConnectors) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *ThirdPartyConnectors) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetTagKey

`func (o *ThirdPartyConnectors) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *ThirdPartyConnectors) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *ThirdPartyConnectors) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.

### HasTagKey

`func (o *ThirdPartyConnectors) HasTagKey() bool`

HasTagKey returns a boolean if a field has been set.

### SetTagKeyNil

`func (o *ThirdPartyConnectors) SetTagKeyNil(b bool)`

 SetTagKeyNil sets the value for TagKey to be an explicit nil

### UnsetTagKey
`func (o *ThirdPartyConnectors) UnsetTagKey()`

UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil
### GetTagValue

`func (o *ThirdPartyConnectors) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *ThirdPartyConnectors) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *ThirdPartyConnectors) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.

### HasTagValue

`func (o *ThirdPartyConnectors) HasTagValue() bool`

HasTagValue returns a boolean if a field has been set.

### SetTagValueNil

`func (o *ThirdPartyConnectors) SetTagValueNil(b bool)`

 SetTagValueNil sets the value for TagValue to be an explicit nil

### UnsetTagValue
`func (o *ThirdPartyConnectors) UnsetTagValue()`

UnsetTagValue ensures that no value is present for TagValue, not even an explicit nil
### GetThrottleDuration

`func (o *ThirdPartyConnectors) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *ThirdPartyConnectors) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *ThirdPartyConnectors) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetTimeout

`func (o *ThirdPartyConnectors) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ThirdPartyConnectors) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ThirdPartyConnectors) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ThirdPartyConnectors) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *ThirdPartyConnectors) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *ThirdPartyConnectors) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *ThirdPartyConnectors) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThirdPartyConnectors) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThirdPartyConnectors) SetType(v string)`

SetType sets Type field to given value.


### GetAzureTenant

`func (o *ThirdPartyConnectors) GetAzureTenant() string`

GetAzureTenant returns the AzureTenant field if non-nil, zero value otherwise.

### GetAzureTenantOk

`func (o *ThirdPartyConnectors) GetAzureTenantOk() (*string, bool)`

GetAzureTenantOk returns a tuple with the AzureTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenant

`func (o *ThirdPartyConnectors) SetAzureTenant(v string)`

SetAzureTenant sets AzureTenant field to given value.

### HasAzureTenant

`func (o *ThirdPartyConnectors) HasAzureTenant() bool`

HasAzureTenant returns a boolean if a field has been set.

### GetPrefix

`func (o *ThirdPartyConnectors) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ThirdPartyConnectors) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ThirdPartyConnectors) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ThirdPartyConnectors) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *ThirdPartyConnectors) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *ThirdPartyConnectors) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetThrottleParallelism

`func (o *ThirdPartyConnectors) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *ThirdPartyConnectors) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *ThirdPartyConnectors) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetVaultBaseUrl

`func (o *ThirdPartyConnectors) GetVaultBaseUrl() string`

GetVaultBaseUrl returns the VaultBaseUrl field if non-nil, zero value otherwise.

### GetVaultBaseUrlOk

`func (o *ThirdPartyConnectors) GetVaultBaseUrlOk() (*string, bool)`

GetVaultBaseUrlOk returns a tuple with the VaultBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultBaseUrl

`func (o *ThirdPartyConnectors) SetVaultBaseUrl(v string)`

SetVaultBaseUrl sets VaultBaseUrl field to given value.


### GetHostname

`func (o *ThirdPartyConnectors) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ThirdPartyConnectors) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ThirdPartyConnectors) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetKeyType

`func (o *ThirdPartyConnectors) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ThirdPartyConnectors) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ThirdPartyConnectors) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ThirdPartyConnectors) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *ThirdPartyConnectors) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *ThirdPartyConnectors) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetLoginProvider

`func (o *ThirdPartyConnectors) GetLoginProvider() string`

GetLoginProvider returns the LoginProvider field if non-nil, zero value otherwise.

### GetLoginProviderOk

`func (o *ThirdPartyConnectors) GetLoginProviderOk() (*string, bool)`

GetLoginProviderOk returns a tuple with the LoginProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginProvider

`func (o *ThirdPartyConnectors) SetLoginProvider(v string)`

SetLoginProvider sets LoginProvider field to given value.

### HasLoginProvider

`func (o *ThirdPartyConnectors) HasLoginProvider() bool`

HasLoginProvider returns a boolean if a field has been set.

### SetLoginProviderNil

`func (o *ThirdPartyConnectors) SetLoginProviderNil(b bool)`

 SetLoginProviderNil sets the value for LoginProvider to be an explicit nil

### UnsetLoginProvider
`func (o *ThirdPartyConnectors) UnsetLoginProvider()`

UnsetLoginProvider ensures that no value is present for LoginProvider, not even an explicit nil
### GetTlsInsecure

`func (o *ThirdPartyConnectors) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *ThirdPartyConnectors) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *ThirdPartyConnectors) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *ThirdPartyConnectors) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *ThirdPartyConnectors) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *ThirdPartyConnectors) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetWithChain

`func (o *ThirdPartyConnectors) GetWithChain() bool`

GetWithChain returns the WithChain field if non-nil, zero value otherwise.

### GetWithChainOk

`func (o *ThirdPartyConnectors) GetWithChainOk() (*bool, bool)`

GetWithChainOk returns a tuple with the WithChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithChain

`func (o *ThirdPartyConnectors) SetWithChain(v bool)`

SetWithChain sets WithChain field to given value.

### HasWithChain

`func (o *ThirdPartyConnectors) HasWithChain() bool`

HasWithChain returns a boolean if a field has been set.

### SetWithChainNil

`func (o *ThirdPartyConnectors) SetWithChainNil(b bool)`

 SetWithChainNil sets the value for WithChain to be an explicit nil

### UnsetWithChain
`func (o *ThirdPartyConnectors) UnsetWithChain()`

UnsetWithChain ensures that no value is present for WithChain, not even an explicit nil
### GetBigIPHostname

`func (o *ThirdPartyConnectors) GetBigIPHostname() string`

GetBigIPHostname returns the BigIPHostname field if non-nil, zero value otherwise.

### GetBigIPHostnameOk

`func (o *ThirdPartyConnectors) GetBigIPHostnameOk() (*string, bool)`

GetBigIPHostnameOk returns a tuple with the BigIPHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigIPHostname

`func (o *ThirdPartyConnectors) SetBigIPHostname(v string)`

SetBigIPHostname sets BigIPHostname field to given value.


### GetCipherGroup

`func (o *ThirdPartyConnectors) GetCipherGroup() string`

GetCipherGroup returns the CipherGroup field if non-nil, zero value otherwise.

### GetCipherGroupOk

`func (o *ThirdPartyConnectors) GetCipherGroupOk() (*string, bool)`

GetCipherGroupOk returns a tuple with the CipherGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherGroup

`func (o *ThirdPartyConnectors) SetCipherGroup(v string)`

SetCipherGroup sets CipherGroup field to given value.

### HasCipherGroup

`func (o *ThirdPartyConnectors) HasCipherGroup() bool`

HasCipherGroup returns a boolean if a field has been set.

### SetCipherGroupNil

`func (o *ThirdPartyConnectors) SetCipherGroupNil(b bool)`

 SetCipherGroupNil sets the value for CipherGroup to be an explicit nil

### UnsetCipherGroup
`func (o *ThirdPartyConnectors) UnsetCipherGroup()`

UnsetCipherGroup ensures that no value is present for CipherGroup, not even an explicit nil
### GetMaxStoredCertificatePerHolder

`func (o *ThirdPartyConnectors) GetMaxStoredCertificatePerHolder() int64`

GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field if non-nil, zero value otherwise.

### GetMaxStoredCertificatePerHolderOk

`func (o *ThirdPartyConnectors) GetMaxStoredCertificatePerHolderOk() (*int64, bool)`

GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredCertificatePerHolder

`func (o *ThirdPartyConnectors) SetMaxStoredCertificatePerHolder(v int64)`

SetMaxStoredCertificatePerHolder sets MaxStoredCertificatePerHolder field to given value.

### HasMaxStoredCertificatePerHolder

`func (o *ThirdPartyConnectors) HasMaxStoredCertificatePerHolder() bool`

HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.

### SetMaxStoredCertificatePerHolderNil

`func (o *ThirdPartyConnectors) SetMaxStoredCertificatePerHolderNil(b bool)`

 SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil

### UnsetMaxStoredCertificatePerHolder
`func (o *ThirdPartyConnectors) UnsetMaxStoredCertificatePerHolder()`

UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
### GetPartition

`func (o *ThirdPartyConnectors) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ThirdPartyConnectors) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ThirdPartyConnectors) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *ThirdPartyConnectors) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### SetPartitionNil

`func (o *ThirdPartyConnectors) SetPartitionNil(b bool)`

 SetPartitionNil sets the value for Partition to be an explicit nil

### UnsetPartition
`func (o *ThirdPartyConnectors) UnsetPartition()`

UnsetPartition ensures that no value is present for Partition, not even an explicit nil
### GetSslParent

`func (o *ThirdPartyConnectors) GetSslParent() string`

GetSslParent returns the SslParent field if non-nil, zero value otherwise.

### GetSslParentOk

`func (o *ThirdPartyConnectors) GetSslParentOk() (*string, bool)`

GetSslParentOk returns a tuple with the SslParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslParent

`func (o *ThirdPartyConnectors) SetSslParent(v string)`

SetSslParent sets SslParent field to given value.

### HasSslParent

`func (o *ThirdPartyConnectors) HasSslParent() bool`

HasSslParent returns a boolean if a field has been set.

### SetSslParentNil

`func (o *ThirdPartyConnectors) SetSslParentNil(b bool)`

 SetSslParentNil sets the value for SslParent to be an explicit nil

### UnsetSslParent
`func (o *ThirdPartyConnectors) UnsetSslParent()`

UnsetSslParent ensures that no value is present for SslParent, not even an explicit nil
### GetVersion

`func (o *ThirdPartyConnectors) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThirdPartyConnectors) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThirdPartyConnectors) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThirdPartyConnectors) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ThirdPartyConnectors) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ThirdPartyConnectors) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetLocation

`func (o *ThirdPartyConnectors) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ThirdPartyConnectors) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ThirdPartyConnectors) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetProject

`func (o *ThirdPartyConnectors) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ThirdPartyConnectors) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ThirdPartyConnectors) SetProject(v string)`

SetProject sets Project field to given value.


### GetIntuneResourceUrl

`func (o *ThirdPartyConnectors) GetIntuneResourceUrl() string`

GetIntuneResourceUrl returns the IntuneResourceUrl field if non-nil, zero value otherwise.

### GetIntuneResourceUrlOk

`func (o *ThirdPartyConnectors) GetIntuneResourceUrlOk() (*string, bool)`

GetIntuneResourceUrlOk returns a tuple with the IntuneResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntuneResourceUrl

`func (o *ThirdPartyConnectors) SetIntuneResourceUrl(v string)`

SetIntuneResourceUrl sets IntuneResourceUrl field to given value.

### HasIntuneResourceUrl

`func (o *ThirdPartyConnectors) HasIntuneResourceUrl() bool`

HasIntuneResourceUrl returns a boolean if a field has been set.

### SetIntuneResourceUrlNil

`func (o *ThirdPartyConnectors) SetIntuneResourceUrlNil(b bool)`

 SetIntuneResourceUrlNil sets the value for IntuneResourceUrl to be an explicit nil

### UnsetIntuneResourceUrl
`func (o *ThirdPartyConnectors) UnsetIntuneResourceUrl()`

UnsetIntuneResourceUrl ensures that no value is present for IntuneResourceUrl, not even an explicit nil
### GetLegacyRevocationMode

`func (o *ThirdPartyConnectors) GetLegacyRevocationMode() bool`

GetLegacyRevocationMode returns the LegacyRevocationMode field if non-nil, zero value otherwise.

### GetLegacyRevocationModeOk

`func (o *ThirdPartyConnectors) GetLegacyRevocationModeOk() (*bool, bool)`

GetLegacyRevocationModeOk returns a tuple with the LegacyRevocationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyRevocationMode

`func (o *ThirdPartyConnectors) SetLegacyRevocationMode(v bool)`

SetLegacyRevocationMode sets LegacyRevocationMode field to given value.


### GetOsQueryString

`func (o *ThirdPartyConnectors) GetOsQueryString() string`

GetOsQueryString returns the OsQueryString field if non-nil, zero value otherwise.

### GetOsQueryStringOk

`func (o *ThirdPartyConnectors) GetOsQueryStringOk() (*string, bool)`

GetOsQueryStringOk returns a tuple with the OsQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsQueryString

`func (o *ThirdPartyConnectors) SetOsQueryString(v string)`

SetOsQueryString sets OsQueryString field to given value.

### HasOsQueryString

`func (o *ThirdPartyConnectors) HasOsQueryString() bool`

HasOsQueryString returns a boolean if a field has been set.

### SetOsQueryStringNil

`func (o *ThirdPartyConnectors) SetOsQueryStringNil(b bool)`

 SetOsQueryStringNil sets the value for OsQueryString to be an explicit nil

### UnsetOsQueryString
`func (o *ThirdPartyConnectors) UnsetOsQueryString()`

UnsetOsQueryString ensures that no value is present for OsQueryString, not even an explicit nil
### GetIntendedPurpose

`func (o *ThirdPartyConnectors) GetIntendedPurpose() string`

GetIntendedPurpose returns the IntendedPurpose field if non-nil, zero value otherwise.

### GetIntendedPurposeOk

`func (o *ThirdPartyConnectors) GetIntendedPurposeOk() (*string, bool)`

GetIntendedPurposeOk returns a tuple with the IntendedPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedPurpose

`func (o *ThirdPartyConnectors) SetIntendedPurpose(v string)`

SetIntendedPurpose sets IntendedPurpose field to given value.

### HasIntendedPurpose

`func (o *ThirdPartyConnectors) HasIntendedPurpose() bool`

HasIntendedPurpose returns a boolean if a field has been set.

### SetIntendedPurposeNil

`func (o *ThirdPartyConnectors) SetIntendedPurposeNil(b bool)`

 SetIntendedPurposeNil sets the value for IntendedPurpose to be an explicit nil

### UnsetIntendedPurpose
`func (o *ThirdPartyConnectors) UnsetIntendedPurpose()`

UnsetIntendedPurpose ensures that no value is present for IntendedPurpose, not even an explicit nil
### GetKeyName

`func (o *ThirdPartyConnectors) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *ThirdPartyConnectors) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *ThirdPartyConnectors) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetProviderName

`func (o *ThirdPartyConnectors) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ThirdPartyConnectors) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ThirdPartyConnectors) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ThirdPartyConnectors) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ThirdPartyConnectors) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ThirdPartyConnectors) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetPubKey

`func (o *ThirdPartyConnectors) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *ThirdPartyConnectors) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *ThirdPartyConnectors) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetSearchFilter

`func (o *ThirdPartyConnectors) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *ThirdPartyConnectors) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *ThirdPartyConnectors) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *ThirdPartyConnectors) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.

### SetSearchFilterNil

`func (o *ThirdPartyConnectors) SetSearchFilterNil(b bool)`

 SetSearchFilterNil sets the value for SearchFilter to be an explicit nil

### UnsetSearchFilter
`func (o *ThirdPartyConnectors) UnsetSearchFilter()`

UnsetSearchFilter ensures that no value is present for SearchFilter, not even an explicit nil
### GetEndpoint

`func (o *ThirdPartyConnectors) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ThirdPartyConnectors) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ThirdPartyConnectors) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetBaseDn

`func (o *ThirdPartyConnectors) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *ThirdPartyConnectors) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *ThirdPartyConnectors) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetCertAttr

`func (o *ThirdPartyConnectors) GetCertAttr() string`

GetCertAttr returns the CertAttr field if non-nil, zero value otherwise.

### GetCertAttrOk

`func (o *ThirdPartyConnectors) GetCertAttrOk() (*string, bool)`

GetCertAttrOk returns a tuple with the CertAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAttr

`func (o *ThirdPartyConnectors) SetCertAttr(v string)`

SetCertAttr sets CertAttr field to given value.

### HasCertAttr

`func (o *ThirdPartyConnectors) HasCertAttr() bool`

HasCertAttr returns a boolean if a field has been set.

### SetCertAttrNil

`func (o *ThirdPartyConnectors) SetCertAttrNil(b bool)`

 SetCertAttrNil sets the value for CertAttr to be an explicit nil

### UnsetCertAttr
`func (o *ThirdPartyConnectors) UnsetCertAttr()`

UnsetCertAttr ensures that no value is present for CertAttr, not even an explicit nil
### GetCertificateAttribute

`func (o *ThirdPartyConnectors) GetCertificateAttribute() string`

GetCertificateAttribute returns the CertificateAttribute field if non-nil, zero value otherwise.

### GetCertificateAttributeOk

`func (o *ThirdPartyConnectors) GetCertificateAttributeOk() (*string, bool)`

GetCertificateAttributeOk returns a tuple with the CertificateAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAttribute

`func (o *ThirdPartyConnectors) SetCertificateAttribute(v string)`

SetCertificateAttribute sets CertificateAttribute field to given value.


### GetFilter

`func (o *ThirdPartyConnectors) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ThirdPartyConnectors) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ThirdPartyConnectors) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ThirdPartyConnectors) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ThirdPartyConnectors) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ThirdPartyConnectors) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetFollowReferrals

`func (o *ThirdPartyConnectors) GetFollowReferrals() bool`

GetFollowReferrals returns the FollowReferrals field if non-nil, zero value otherwise.

### GetFollowReferralsOk

`func (o *ThirdPartyConnectors) GetFollowReferralsOk() (*bool, bool)`

GetFollowReferralsOk returns a tuple with the FollowReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowReferrals

`func (o *ThirdPartyConnectors) SetFollowReferrals(v bool)`

SetFollowReferrals sets FollowReferrals field to given value.

### HasFollowReferrals

`func (o *ThirdPartyConnectors) HasFollowReferrals() bool`

HasFollowReferrals returns a boolean if a field has been set.

### SetFollowReferralsNil

`func (o *ThirdPartyConnectors) SetFollowReferralsNil(b bool)`

 SetFollowReferralsNil sets the value for FollowReferrals to be an explicit nil

### UnsetFollowReferrals
`func (o *ThirdPartyConnectors) UnsetFollowReferrals()`

UnsetFollowReferrals ensures that no value is present for FollowReferrals, not even an explicit nil
### GetPort

`func (o *ThirdPartyConnectors) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ThirdPartyConnectors) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ThirdPartyConnectors) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ThirdPartyConnectors) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *ThirdPartyConnectors) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ThirdPartyConnectors) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUserIdentifierAttribute

`func (o *ThirdPartyConnectors) GetUserIdentifierAttribute() string`

GetUserIdentifierAttribute returns the UserIdentifierAttribute field if non-nil, zero value otherwise.

### GetUserIdentifierAttributeOk

`func (o *ThirdPartyConnectors) GetUserIdentifierAttributeOk() (*string, bool)`

GetUserIdentifierAttributeOk returns a tuple with the UserIdentifierAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifierAttribute

`func (o *ThirdPartyConnectors) SetUserIdentifierAttribute(v string)`

SetUserIdentifierAttribute sets UserIdentifierAttribute field to given value.


### GetCertificateStorePath

`func (o *ThirdPartyConnectors) GetCertificateStorePath() string`

GetCertificateStorePath returns the CertificateStorePath field if non-nil, zero value otherwise.

### GetCertificateStorePathOk

`func (o *ThirdPartyConnectors) GetCertificateStorePathOk() (*string, bool)`

GetCertificateStorePathOk returns a tuple with the CertificateStorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStorePath

`func (o *ThirdPartyConnectors) SetCertificateStorePath(v string)`

SetCertificateStorePath sets CertificateStorePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


