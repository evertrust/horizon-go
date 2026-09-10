# ThirdPartyConnectorResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the account to authenticate on Netscaler | 
**Name** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Region** | **string** |  | 
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
**KeyType** | Pointer to **NullableString** | One of &#x60;rsa-2048&#x60;, &#x60;rsa-3072&#x60;, &#x60;rsa-4096&#x60;, &#x60;rsa-8192&#x60;, &#x60;ec-secp256r1&#x60;, &#x60;ec-secp384r1&#x60;, &#x60;ec-secp521r1&#x60;, &#x60;ec-brainpoolp256r1&#x60;, &#x60;ec-brainpoolp384r1&#x60;, &#x60;ec-brainpoolp512r1&#x60;, &#x60;ed-448&#x60;, &#x60;ed-25519&#x60;, &#x60;mldsa-44&#x60;, &#x60;mldsa-65&#x60;, &#x60;mldsa-87&#x60;, &#x60;slhdsa-sha2-128s&#x60;, &#x60;slhdsa-sha2-128f&#x60;, &#x60;slhdsa-sha2-192s&#x60;, &#x60;slhdsa-sha2-192f&#x60;, &#x60;slhdsa-sha2-256s&#x60;, &#x60;slhdsa-sha2-256f&#x60;, &#x60;slhdsa-sha2-128ssha256&#x60;, &#x60;slhdsa-sha2-128fsha256&#x60;, &#x60;slhdsa-sha2-192ssha512&#x60;, &#x60;slhdsa-sha2-192fsha512&#x60;, &#x60;slhdsa-sha2-256ssha512&#x60;, &#x60;slhdsa-sha2-256fsha512&#x60; or &#x60;&lt;primary key type&gt;+&lt;alternate key type&gt;&#x60; | [optional] 
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
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewThirdPartyConnectorResponses

`func NewThirdPartyConnectorResponses(id string, credentials string, name string, region string, throttleDuration string, type_ string, throttleParallelism int64, vaultBaseUrl string, hostname string, bigIPHostname string, location string, project string, legacyRevocationMode bool, keyName string, pubKey string, endpoint string, baseDn string, certificateAttribute string, userIdentifierAttribute string, certificateStorePath string, ) *ThirdPartyConnectorResponses`

NewThirdPartyConnectorResponses instantiates a new ThirdPartyConnectorResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyConnectorResponsesWithDefaults

`func NewThirdPartyConnectorResponsesWithDefaults() *ThirdPartyConnectorResponses`

NewThirdPartyConnectorResponsesWithDefaults instantiates a new ThirdPartyConnectorResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyConnectorResponses) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyConnectorResponses) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyConnectorResponses) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *ThirdPartyConnectorResponses) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ThirdPartyConnectorResponses) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ThirdPartyConnectorResponses) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetName

`func (o *ThirdPartyConnectorResponses) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThirdPartyConnectorResponses) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThirdPartyConnectorResponses) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *ThirdPartyConnectorResponses) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ThirdPartyConnectorResponses) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ThirdPartyConnectorResponses) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ThirdPartyConnectorResponses) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *ThirdPartyConnectorResponses) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *ThirdPartyConnectorResponses) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRegion

`func (o *ThirdPartyConnectorResponses) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ThirdPartyConnectorResponses) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ThirdPartyConnectorResponses) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetResourceGroupName

`func (o *ThirdPartyConnectorResponses) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ThirdPartyConnectorResponses) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ThirdPartyConnectorResponses) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *ThirdPartyConnectorResponses) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *ThirdPartyConnectorResponses) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *ThirdPartyConnectorResponses) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetRoleArn

`func (o *ThirdPartyConnectorResponses) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ThirdPartyConnectorResponses) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ThirdPartyConnectorResponses) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *ThirdPartyConnectorResponses) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *ThirdPartyConnectorResponses) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *ThirdPartyConnectorResponses) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetTagKey

`func (o *ThirdPartyConnectorResponses) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *ThirdPartyConnectorResponses) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *ThirdPartyConnectorResponses) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.

### HasTagKey

`func (o *ThirdPartyConnectorResponses) HasTagKey() bool`

HasTagKey returns a boolean if a field has been set.

### SetTagKeyNil

`func (o *ThirdPartyConnectorResponses) SetTagKeyNil(b bool)`

 SetTagKeyNil sets the value for TagKey to be an explicit nil

### UnsetTagKey
`func (o *ThirdPartyConnectorResponses) UnsetTagKey()`

UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil
### GetTagValue

`func (o *ThirdPartyConnectorResponses) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *ThirdPartyConnectorResponses) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *ThirdPartyConnectorResponses) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.

### HasTagValue

`func (o *ThirdPartyConnectorResponses) HasTagValue() bool`

HasTagValue returns a boolean if a field has been set.

### SetTagValueNil

`func (o *ThirdPartyConnectorResponses) SetTagValueNil(b bool)`

 SetTagValueNil sets the value for TagValue to be an explicit nil

### UnsetTagValue
`func (o *ThirdPartyConnectorResponses) UnsetTagValue()`

UnsetTagValue ensures that no value is present for TagValue, not even an explicit nil
### GetThrottleDuration

`func (o *ThirdPartyConnectorResponses) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *ThirdPartyConnectorResponses) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *ThirdPartyConnectorResponses) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetTimeout

`func (o *ThirdPartyConnectorResponses) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ThirdPartyConnectorResponses) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ThirdPartyConnectorResponses) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ThirdPartyConnectorResponses) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *ThirdPartyConnectorResponses) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *ThirdPartyConnectorResponses) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *ThirdPartyConnectorResponses) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThirdPartyConnectorResponses) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThirdPartyConnectorResponses) SetType(v string)`

SetType sets Type field to given value.


### GetAzureTenant

`func (o *ThirdPartyConnectorResponses) GetAzureTenant() string`

GetAzureTenant returns the AzureTenant field if non-nil, zero value otherwise.

### GetAzureTenantOk

`func (o *ThirdPartyConnectorResponses) GetAzureTenantOk() (*string, bool)`

GetAzureTenantOk returns a tuple with the AzureTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenant

`func (o *ThirdPartyConnectorResponses) SetAzureTenant(v string)`

SetAzureTenant sets AzureTenant field to given value.

### HasAzureTenant

`func (o *ThirdPartyConnectorResponses) HasAzureTenant() bool`

HasAzureTenant returns a boolean if a field has been set.

### GetPrefix

`func (o *ThirdPartyConnectorResponses) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ThirdPartyConnectorResponses) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ThirdPartyConnectorResponses) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ThirdPartyConnectorResponses) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *ThirdPartyConnectorResponses) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *ThirdPartyConnectorResponses) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetThrottleParallelism

`func (o *ThirdPartyConnectorResponses) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *ThirdPartyConnectorResponses) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *ThirdPartyConnectorResponses) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetVaultBaseUrl

`func (o *ThirdPartyConnectorResponses) GetVaultBaseUrl() string`

GetVaultBaseUrl returns the VaultBaseUrl field if non-nil, zero value otherwise.

### GetVaultBaseUrlOk

`func (o *ThirdPartyConnectorResponses) GetVaultBaseUrlOk() (*string, bool)`

GetVaultBaseUrlOk returns a tuple with the VaultBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultBaseUrl

`func (o *ThirdPartyConnectorResponses) SetVaultBaseUrl(v string)`

SetVaultBaseUrl sets VaultBaseUrl field to given value.


### GetHostname

`func (o *ThirdPartyConnectorResponses) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ThirdPartyConnectorResponses) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ThirdPartyConnectorResponses) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetKeyType

`func (o *ThirdPartyConnectorResponses) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ThirdPartyConnectorResponses) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ThirdPartyConnectorResponses) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ThirdPartyConnectorResponses) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *ThirdPartyConnectorResponses) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *ThirdPartyConnectorResponses) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetLoginProvider

`func (o *ThirdPartyConnectorResponses) GetLoginProvider() string`

GetLoginProvider returns the LoginProvider field if non-nil, zero value otherwise.

### GetLoginProviderOk

`func (o *ThirdPartyConnectorResponses) GetLoginProviderOk() (*string, bool)`

GetLoginProviderOk returns a tuple with the LoginProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginProvider

`func (o *ThirdPartyConnectorResponses) SetLoginProvider(v string)`

SetLoginProvider sets LoginProvider field to given value.

### HasLoginProvider

`func (o *ThirdPartyConnectorResponses) HasLoginProvider() bool`

HasLoginProvider returns a boolean if a field has been set.

### SetLoginProviderNil

`func (o *ThirdPartyConnectorResponses) SetLoginProviderNil(b bool)`

 SetLoginProviderNil sets the value for LoginProvider to be an explicit nil

### UnsetLoginProvider
`func (o *ThirdPartyConnectorResponses) UnsetLoginProvider()`

UnsetLoginProvider ensures that no value is present for LoginProvider, not even an explicit nil
### GetTlsInsecure

`func (o *ThirdPartyConnectorResponses) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *ThirdPartyConnectorResponses) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *ThirdPartyConnectorResponses) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *ThirdPartyConnectorResponses) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *ThirdPartyConnectorResponses) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *ThirdPartyConnectorResponses) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetWithChain

`func (o *ThirdPartyConnectorResponses) GetWithChain() bool`

GetWithChain returns the WithChain field if non-nil, zero value otherwise.

### GetWithChainOk

`func (o *ThirdPartyConnectorResponses) GetWithChainOk() (*bool, bool)`

GetWithChainOk returns a tuple with the WithChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithChain

`func (o *ThirdPartyConnectorResponses) SetWithChain(v bool)`

SetWithChain sets WithChain field to given value.

### HasWithChain

`func (o *ThirdPartyConnectorResponses) HasWithChain() bool`

HasWithChain returns a boolean if a field has been set.

### SetWithChainNil

`func (o *ThirdPartyConnectorResponses) SetWithChainNil(b bool)`

 SetWithChainNil sets the value for WithChain to be an explicit nil

### UnsetWithChain
`func (o *ThirdPartyConnectorResponses) UnsetWithChain()`

UnsetWithChain ensures that no value is present for WithChain, not even an explicit nil
### GetBigIPHostname

`func (o *ThirdPartyConnectorResponses) GetBigIPHostname() string`

GetBigIPHostname returns the BigIPHostname field if non-nil, zero value otherwise.

### GetBigIPHostnameOk

`func (o *ThirdPartyConnectorResponses) GetBigIPHostnameOk() (*string, bool)`

GetBigIPHostnameOk returns a tuple with the BigIPHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigIPHostname

`func (o *ThirdPartyConnectorResponses) SetBigIPHostname(v string)`

SetBigIPHostname sets BigIPHostname field to given value.


### GetCipherGroup

`func (o *ThirdPartyConnectorResponses) GetCipherGroup() string`

GetCipherGroup returns the CipherGroup field if non-nil, zero value otherwise.

### GetCipherGroupOk

`func (o *ThirdPartyConnectorResponses) GetCipherGroupOk() (*string, bool)`

GetCipherGroupOk returns a tuple with the CipherGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherGroup

`func (o *ThirdPartyConnectorResponses) SetCipherGroup(v string)`

SetCipherGroup sets CipherGroup field to given value.

### HasCipherGroup

`func (o *ThirdPartyConnectorResponses) HasCipherGroup() bool`

HasCipherGroup returns a boolean if a field has been set.

### SetCipherGroupNil

`func (o *ThirdPartyConnectorResponses) SetCipherGroupNil(b bool)`

 SetCipherGroupNil sets the value for CipherGroup to be an explicit nil

### UnsetCipherGroup
`func (o *ThirdPartyConnectorResponses) UnsetCipherGroup()`

UnsetCipherGroup ensures that no value is present for CipherGroup, not even an explicit nil
### GetMaxStoredCertificatePerHolder

`func (o *ThirdPartyConnectorResponses) GetMaxStoredCertificatePerHolder() int64`

GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field if non-nil, zero value otherwise.

### GetMaxStoredCertificatePerHolderOk

`func (o *ThirdPartyConnectorResponses) GetMaxStoredCertificatePerHolderOk() (*int64, bool)`

GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredCertificatePerHolder

`func (o *ThirdPartyConnectorResponses) SetMaxStoredCertificatePerHolder(v int64)`

SetMaxStoredCertificatePerHolder sets MaxStoredCertificatePerHolder field to given value.

### HasMaxStoredCertificatePerHolder

`func (o *ThirdPartyConnectorResponses) HasMaxStoredCertificatePerHolder() bool`

HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.

### SetMaxStoredCertificatePerHolderNil

`func (o *ThirdPartyConnectorResponses) SetMaxStoredCertificatePerHolderNil(b bool)`

 SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil

### UnsetMaxStoredCertificatePerHolder
`func (o *ThirdPartyConnectorResponses) UnsetMaxStoredCertificatePerHolder()`

UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
### GetPartition

`func (o *ThirdPartyConnectorResponses) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ThirdPartyConnectorResponses) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ThirdPartyConnectorResponses) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *ThirdPartyConnectorResponses) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### SetPartitionNil

`func (o *ThirdPartyConnectorResponses) SetPartitionNil(b bool)`

 SetPartitionNil sets the value for Partition to be an explicit nil

### UnsetPartition
`func (o *ThirdPartyConnectorResponses) UnsetPartition()`

UnsetPartition ensures that no value is present for Partition, not even an explicit nil
### GetSslParent

`func (o *ThirdPartyConnectorResponses) GetSslParent() string`

GetSslParent returns the SslParent field if non-nil, zero value otherwise.

### GetSslParentOk

`func (o *ThirdPartyConnectorResponses) GetSslParentOk() (*string, bool)`

GetSslParentOk returns a tuple with the SslParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslParent

`func (o *ThirdPartyConnectorResponses) SetSslParent(v string)`

SetSslParent sets SslParent field to given value.

### HasSslParent

`func (o *ThirdPartyConnectorResponses) HasSslParent() bool`

HasSslParent returns a boolean if a field has been set.

### SetSslParentNil

`func (o *ThirdPartyConnectorResponses) SetSslParentNil(b bool)`

 SetSslParentNil sets the value for SslParent to be an explicit nil

### UnsetSslParent
`func (o *ThirdPartyConnectorResponses) UnsetSslParent()`

UnsetSslParent ensures that no value is present for SslParent, not even an explicit nil
### GetVersion

`func (o *ThirdPartyConnectorResponses) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThirdPartyConnectorResponses) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThirdPartyConnectorResponses) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThirdPartyConnectorResponses) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ThirdPartyConnectorResponses) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ThirdPartyConnectorResponses) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetLocation

`func (o *ThirdPartyConnectorResponses) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ThirdPartyConnectorResponses) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ThirdPartyConnectorResponses) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetProject

`func (o *ThirdPartyConnectorResponses) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ThirdPartyConnectorResponses) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ThirdPartyConnectorResponses) SetProject(v string)`

SetProject sets Project field to given value.


### GetIntuneResourceUrl

`func (o *ThirdPartyConnectorResponses) GetIntuneResourceUrl() string`

GetIntuneResourceUrl returns the IntuneResourceUrl field if non-nil, zero value otherwise.

### GetIntuneResourceUrlOk

`func (o *ThirdPartyConnectorResponses) GetIntuneResourceUrlOk() (*string, bool)`

GetIntuneResourceUrlOk returns a tuple with the IntuneResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntuneResourceUrl

`func (o *ThirdPartyConnectorResponses) SetIntuneResourceUrl(v string)`

SetIntuneResourceUrl sets IntuneResourceUrl field to given value.

### HasIntuneResourceUrl

`func (o *ThirdPartyConnectorResponses) HasIntuneResourceUrl() bool`

HasIntuneResourceUrl returns a boolean if a field has been set.

### SetIntuneResourceUrlNil

`func (o *ThirdPartyConnectorResponses) SetIntuneResourceUrlNil(b bool)`

 SetIntuneResourceUrlNil sets the value for IntuneResourceUrl to be an explicit nil

### UnsetIntuneResourceUrl
`func (o *ThirdPartyConnectorResponses) UnsetIntuneResourceUrl()`

UnsetIntuneResourceUrl ensures that no value is present for IntuneResourceUrl, not even an explicit nil
### GetLegacyRevocationMode

`func (o *ThirdPartyConnectorResponses) GetLegacyRevocationMode() bool`

GetLegacyRevocationMode returns the LegacyRevocationMode field if non-nil, zero value otherwise.

### GetLegacyRevocationModeOk

`func (o *ThirdPartyConnectorResponses) GetLegacyRevocationModeOk() (*bool, bool)`

GetLegacyRevocationModeOk returns a tuple with the LegacyRevocationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyRevocationMode

`func (o *ThirdPartyConnectorResponses) SetLegacyRevocationMode(v bool)`

SetLegacyRevocationMode sets LegacyRevocationMode field to given value.


### GetOsQueryString

`func (o *ThirdPartyConnectorResponses) GetOsQueryString() string`

GetOsQueryString returns the OsQueryString field if non-nil, zero value otherwise.

### GetOsQueryStringOk

`func (o *ThirdPartyConnectorResponses) GetOsQueryStringOk() (*string, bool)`

GetOsQueryStringOk returns a tuple with the OsQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsQueryString

`func (o *ThirdPartyConnectorResponses) SetOsQueryString(v string)`

SetOsQueryString sets OsQueryString field to given value.

### HasOsQueryString

`func (o *ThirdPartyConnectorResponses) HasOsQueryString() bool`

HasOsQueryString returns a boolean if a field has been set.

### SetOsQueryStringNil

`func (o *ThirdPartyConnectorResponses) SetOsQueryStringNil(b bool)`

 SetOsQueryStringNil sets the value for OsQueryString to be an explicit nil

### UnsetOsQueryString
`func (o *ThirdPartyConnectorResponses) UnsetOsQueryString()`

UnsetOsQueryString ensures that no value is present for OsQueryString, not even an explicit nil
### GetIntendedPurpose

`func (o *ThirdPartyConnectorResponses) GetIntendedPurpose() string`

GetIntendedPurpose returns the IntendedPurpose field if non-nil, zero value otherwise.

### GetIntendedPurposeOk

`func (o *ThirdPartyConnectorResponses) GetIntendedPurposeOk() (*string, bool)`

GetIntendedPurposeOk returns a tuple with the IntendedPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedPurpose

`func (o *ThirdPartyConnectorResponses) SetIntendedPurpose(v string)`

SetIntendedPurpose sets IntendedPurpose field to given value.

### HasIntendedPurpose

`func (o *ThirdPartyConnectorResponses) HasIntendedPurpose() bool`

HasIntendedPurpose returns a boolean if a field has been set.

### SetIntendedPurposeNil

`func (o *ThirdPartyConnectorResponses) SetIntendedPurposeNil(b bool)`

 SetIntendedPurposeNil sets the value for IntendedPurpose to be an explicit nil

### UnsetIntendedPurpose
`func (o *ThirdPartyConnectorResponses) UnsetIntendedPurpose()`

UnsetIntendedPurpose ensures that no value is present for IntendedPurpose, not even an explicit nil
### GetKeyName

`func (o *ThirdPartyConnectorResponses) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *ThirdPartyConnectorResponses) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *ThirdPartyConnectorResponses) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetProviderName

`func (o *ThirdPartyConnectorResponses) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ThirdPartyConnectorResponses) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ThirdPartyConnectorResponses) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ThirdPartyConnectorResponses) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ThirdPartyConnectorResponses) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ThirdPartyConnectorResponses) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetPubKey

`func (o *ThirdPartyConnectorResponses) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *ThirdPartyConnectorResponses) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *ThirdPartyConnectorResponses) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetSearchFilter

`func (o *ThirdPartyConnectorResponses) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *ThirdPartyConnectorResponses) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *ThirdPartyConnectorResponses) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *ThirdPartyConnectorResponses) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.

### SetSearchFilterNil

`func (o *ThirdPartyConnectorResponses) SetSearchFilterNil(b bool)`

 SetSearchFilterNil sets the value for SearchFilter to be an explicit nil

### UnsetSearchFilter
`func (o *ThirdPartyConnectorResponses) UnsetSearchFilter()`

UnsetSearchFilter ensures that no value is present for SearchFilter, not even an explicit nil
### GetEndpoint

`func (o *ThirdPartyConnectorResponses) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ThirdPartyConnectorResponses) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ThirdPartyConnectorResponses) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetBaseDn

`func (o *ThirdPartyConnectorResponses) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *ThirdPartyConnectorResponses) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *ThirdPartyConnectorResponses) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetCertAttr

`func (o *ThirdPartyConnectorResponses) GetCertAttr() string`

GetCertAttr returns the CertAttr field if non-nil, zero value otherwise.

### GetCertAttrOk

`func (o *ThirdPartyConnectorResponses) GetCertAttrOk() (*string, bool)`

GetCertAttrOk returns a tuple with the CertAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAttr

`func (o *ThirdPartyConnectorResponses) SetCertAttr(v string)`

SetCertAttr sets CertAttr field to given value.

### HasCertAttr

`func (o *ThirdPartyConnectorResponses) HasCertAttr() bool`

HasCertAttr returns a boolean if a field has been set.

### SetCertAttrNil

`func (o *ThirdPartyConnectorResponses) SetCertAttrNil(b bool)`

 SetCertAttrNil sets the value for CertAttr to be an explicit nil

### UnsetCertAttr
`func (o *ThirdPartyConnectorResponses) UnsetCertAttr()`

UnsetCertAttr ensures that no value is present for CertAttr, not even an explicit nil
### GetCertificateAttribute

`func (o *ThirdPartyConnectorResponses) GetCertificateAttribute() string`

GetCertificateAttribute returns the CertificateAttribute field if non-nil, zero value otherwise.

### GetCertificateAttributeOk

`func (o *ThirdPartyConnectorResponses) GetCertificateAttributeOk() (*string, bool)`

GetCertificateAttributeOk returns a tuple with the CertificateAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAttribute

`func (o *ThirdPartyConnectorResponses) SetCertificateAttribute(v string)`

SetCertificateAttribute sets CertificateAttribute field to given value.


### GetFilter

`func (o *ThirdPartyConnectorResponses) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ThirdPartyConnectorResponses) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ThirdPartyConnectorResponses) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ThirdPartyConnectorResponses) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ThirdPartyConnectorResponses) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ThirdPartyConnectorResponses) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetFollowReferrals

`func (o *ThirdPartyConnectorResponses) GetFollowReferrals() bool`

GetFollowReferrals returns the FollowReferrals field if non-nil, zero value otherwise.

### GetFollowReferralsOk

`func (o *ThirdPartyConnectorResponses) GetFollowReferralsOk() (*bool, bool)`

GetFollowReferralsOk returns a tuple with the FollowReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowReferrals

`func (o *ThirdPartyConnectorResponses) SetFollowReferrals(v bool)`

SetFollowReferrals sets FollowReferrals field to given value.

### HasFollowReferrals

`func (o *ThirdPartyConnectorResponses) HasFollowReferrals() bool`

HasFollowReferrals returns a boolean if a field has been set.

### SetFollowReferralsNil

`func (o *ThirdPartyConnectorResponses) SetFollowReferralsNil(b bool)`

 SetFollowReferralsNil sets the value for FollowReferrals to be an explicit nil

### UnsetFollowReferrals
`func (o *ThirdPartyConnectorResponses) UnsetFollowReferrals()`

UnsetFollowReferrals ensures that no value is present for FollowReferrals, not even an explicit nil
### GetPort

`func (o *ThirdPartyConnectorResponses) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ThirdPartyConnectorResponses) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ThirdPartyConnectorResponses) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ThirdPartyConnectorResponses) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *ThirdPartyConnectorResponses) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ThirdPartyConnectorResponses) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUserIdentifierAttribute

`func (o *ThirdPartyConnectorResponses) GetUserIdentifierAttribute() string`

GetUserIdentifierAttribute returns the UserIdentifierAttribute field if non-nil, zero value otherwise.

### GetUserIdentifierAttributeOk

`func (o *ThirdPartyConnectorResponses) GetUserIdentifierAttributeOk() (*string, bool)`

GetUserIdentifierAttributeOk returns a tuple with the UserIdentifierAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifierAttribute

`func (o *ThirdPartyConnectorResponses) SetUserIdentifierAttribute(v string)`

SetUserIdentifierAttribute sets UserIdentifierAttribute field to given value.


### GetCertificateStorePath

`func (o *ThirdPartyConnectorResponses) GetCertificateStorePath() string`

GetCertificateStorePath returns the CertificateStorePath field if non-nil, zero value otherwise.

### GetCertificateStorePathOk

`func (o *ThirdPartyConnectorResponses) GetCertificateStorePathOk() (*string, bool)`

GetCertificateStorePathOk returns a tuple with the CertificateStorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStorePath

`func (o *ThirdPartyConnectorResponses) SetCertificateStorePath(v string)`

SetCertificateStorePath sets CertificateStorePath field to given value.


### GetRenewalPeriod

`func (o *ThirdPartyConnectorResponses) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *ThirdPartyConnectorResponses) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *ThirdPartyConnectorResponses) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *ThirdPartyConnectorResponses) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *ThirdPartyConnectorResponses) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *ThirdPartyConnectorResponses) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


