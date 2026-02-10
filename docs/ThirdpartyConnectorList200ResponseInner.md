# ThirdpartyConnectorList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**ThrottleDuration** | **string** |  | 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Region** | **string** |  | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing login DN and password. | 
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
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**KeyType** | Pointer to **NullableString** | One of &#x60;rsa-2048&#x60;, &#x60;rsa-3072&#x60;, &#x60;rsa-4096&#x60;, &#x60;rsa-8192&#x60;, &#x60;ec-secp256r1&#x60;, &#x60;ec-secp384r1&#x60;, &#x60;ec-secp521r1&#x60;, &#x60;ed-448&#x60;, &#x60;ed-25519&#x60;, &#x60;mldsa-44&#x60;, &#x60;mldsa-65&#x60;, &#x60;mldsa-87&#x60;, &#x60;slhdsa-sha2-128s&#x60;, &#x60;slhdsa-sha2-128f&#x60;, &#x60;slhdsa-sha2-192s&#x60;, &#x60;slhdsa-sha2-192f&#x60;, &#x60;slhdsa-sha2-256s&#x60;, &#x60;slhdsa-sha2-256f&#x60;, &#x60;slhdsa-sha2-128ssha256&#x60;, &#x60;slhdsa-sha2-128fsha256&#x60;, &#x60;slhdsa-sha2-192ssha512&#x60;, &#x60;slhdsa-sha2-192fsha512&#x60;, &#x60;slhdsa-sha2-256ssha512&#x60;, &#x60;slhdsa-sha2-256fsha512&#x60; or &#x60;&lt;primary key type&gt;+&lt;alternate key type&gt;&#x60; | [optional] 
**Hostname** | **string** |  | 
**LoginProvider** | Pointer to **NullableString** |  | [optional] 
**WithChain** | Pointer to **NullableBool** | Enable the certificate trust chain to be pushed. | [optional] [default to true]
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

### NewThirdpartyConnectorList200ResponseInner

`func NewThirdpartyConnectorList200ResponseInner(id string, type_ string, name string, throttleDuration string, region string, credentials string, throttleParallelism int64, bigIPHostname string, hostname string, tenant string, legacyRevocationMode bool, endpoint string, baseDn string, pubKey string, keyName string, vaultBaseUrl string, project string, location string, userIdentifierAttribute string, certificateAttribute string, ) *ThirdpartyConnectorList200ResponseInner`

NewThirdpartyConnectorList200ResponseInner instantiates a new ThirdpartyConnectorList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdpartyConnectorList200ResponseInnerWithDefaults

`func NewThirdpartyConnectorList200ResponseInnerWithDefaults() *ThirdpartyConnectorList200ResponseInner`

NewThirdpartyConnectorList200ResponseInnerWithDefaults instantiates a new ThirdpartyConnectorList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdpartyConnectorList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdpartyConnectorList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ThirdpartyConnectorList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThirdpartyConnectorList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ThirdpartyConnectorList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThirdpartyConnectorList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetThrottleDuration

`func (o *ThirdpartyConnectorList200ResponseInner) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *ThirdpartyConnectorList200ResponseInner) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetRenewalPeriod

`func (o *ThirdpartyConnectorList200ResponseInner) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *ThirdpartyConnectorList200ResponseInner) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *ThirdpartyConnectorList200ResponseInner) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetTimeout

`func (o *ThirdpartyConnectorList200ResponseInner) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ThirdpartyConnectorList200ResponseInner) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ThirdpartyConnectorList200ResponseInner) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *ThirdpartyConnectorList200ResponseInner) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ThirdpartyConnectorList200ResponseInner) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ThirdpartyConnectorList200ResponseInner) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetRegion

`func (o *ThirdpartyConnectorList200ResponseInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ThirdpartyConnectorList200ResponseInner) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCredentials

`func (o *ThirdpartyConnectorList200ResponseInner) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ThirdpartyConnectorList200ResponseInner) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetResourceGroupName

`func (o *ThirdpartyConnectorList200ResponseInner) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ThirdpartyConnectorList200ResponseInner) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *ThirdpartyConnectorList200ResponseInner) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetRoleArn

`func (o *ThirdpartyConnectorList200ResponseInner) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ThirdpartyConnectorList200ResponseInner) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *ThirdpartyConnectorList200ResponseInner) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetTagKey

`func (o *ThirdpartyConnectorList200ResponseInner) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *ThirdpartyConnectorList200ResponseInner) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.

### HasTagKey

`func (o *ThirdpartyConnectorList200ResponseInner) HasTagKey() bool`

HasTagKey returns a boolean if a field has been set.

### SetTagKeyNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetTagKeyNil(b bool)`

 SetTagKeyNil sets the value for TagKey to be an explicit nil

### UnsetTagKey
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetTagKey()`

UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil
### GetTagValue

`func (o *ThirdpartyConnectorList200ResponseInner) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *ThirdpartyConnectorList200ResponseInner) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.

### HasTagValue

`func (o *ThirdpartyConnectorList200ResponseInner) HasTagValue() bool`

HasTagValue returns a boolean if a field has been set.

### SetTagValueNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetTagValueNil(b bool)`

 SetTagValueNil sets the value for TagValue to be an explicit nil

### UnsetTagValue
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetTagValue()`

UnsetTagValue ensures that no value is present for TagValue, not even an explicit nil
### GetThrottleParallelism

`func (o *ThirdpartyConnectorList200ResponseInner) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *ThirdpartyConnectorList200ResponseInner) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetMaxStoredCertificatePerHolder

`func (o *ThirdpartyConnectorList200ResponseInner) GetMaxStoredCertificatePerHolder() int64`

GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field if non-nil, zero value otherwise.

### GetMaxStoredCertificatePerHolderOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetMaxStoredCertificatePerHolderOk() (*int64, bool)`

GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredCertificatePerHolder

`func (o *ThirdpartyConnectorList200ResponseInner) SetMaxStoredCertificatePerHolder(v int64)`

SetMaxStoredCertificatePerHolder sets MaxStoredCertificatePerHolder field to given value.

### HasMaxStoredCertificatePerHolder

`func (o *ThirdpartyConnectorList200ResponseInner) HasMaxStoredCertificatePerHolder() bool`

HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.

### SetMaxStoredCertificatePerHolderNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetMaxStoredCertificatePerHolderNil(b bool)`

 SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil

### UnsetMaxStoredCertificatePerHolder
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetMaxStoredCertificatePerHolder()`

UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
### GetBigIPHostname

`func (o *ThirdpartyConnectorList200ResponseInner) GetBigIPHostname() string`

GetBigIPHostname returns the BigIPHostname field if non-nil, zero value otherwise.

### GetBigIPHostnameOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetBigIPHostnameOk() (*string, bool)`

GetBigIPHostnameOk returns a tuple with the BigIPHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigIPHostname

`func (o *ThirdpartyConnectorList200ResponseInner) SetBigIPHostname(v string)`

SetBigIPHostname sets BigIPHostname field to given value.


### GetPartition

`func (o *ThirdpartyConnectorList200ResponseInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ThirdpartyConnectorList200ResponseInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *ThirdpartyConnectorList200ResponseInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### SetPartitionNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetPartitionNil(b bool)`

 SetPartitionNil sets the value for Partition to be an explicit nil

### UnsetPartition
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetPartition()`

UnsetPartition ensures that no value is present for Partition, not even an explicit nil
### GetSslParent

`func (o *ThirdpartyConnectorList200ResponseInner) GetSslParent() string`

GetSslParent returns the SslParent field if non-nil, zero value otherwise.

### GetSslParentOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetSslParentOk() (*string, bool)`

GetSslParentOk returns a tuple with the SslParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslParent

`func (o *ThirdpartyConnectorList200ResponseInner) SetSslParent(v string)`

SetSslParent sets SslParent field to given value.

### HasSslParent

`func (o *ThirdpartyConnectorList200ResponseInner) HasSslParent() bool`

HasSslParent returns a boolean if a field has been set.

### SetSslParentNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetSslParentNil(b bool)`

 SetSslParentNil sets the value for SslParent to be an explicit nil

### UnsetSslParent
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetSslParent()`

UnsetSslParent ensures that no value is present for SslParent, not even an explicit nil
### GetPrefix

`func (o *ThirdpartyConnectorList200ResponseInner) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ThirdpartyConnectorList200ResponseInner) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ThirdpartyConnectorList200ResponseInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetCipherGroup

`func (o *ThirdpartyConnectorList200ResponseInner) GetCipherGroup() string`

GetCipherGroup returns the CipherGroup field if non-nil, zero value otherwise.

### GetCipherGroupOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetCipherGroupOk() (*string, bool)`

GetCipherGroupOk returns a tuple with the CipherGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherGroup

`func (o *ThirdpartyConnectorList200ResponseInner) SetCipherGroup(v string)`

SetCipherGroup sets CipherGroup field to given value.

### HasCipherGroup

`func (o *ThirdpartyConnectorList200ResponseInner) HasCipherGroup() bool`

HasCipherGroup returns a boolean if a field has been set.

### SetCipherGroupNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetCipherGroupNil(b bool)`

 SetCipherGroupNil sets the value for CipherGroup to be an explicit nil

### UnsetCipherGroup
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetCipherGroup()`

UnsetCipherGroup ensures that no value is present for CipherGroup, not even an explicit nil
### GetVersion

`func (o *ThirdpartyConnectorList200ResponseInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThirdpartyConnectorList200ResponseInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThirdpartyConnectorList200ResponseInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetTlsInsecure

`func (o *ThirdpartyConnectorList200ResponseInner) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *ThirdpartyConnectorList200ResponseInner) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *ThirdpartyConnectorList200ResponseInner) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetKeyType

`func (o *ThirdpartyConnectorList200ResponseInner) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ThirdpartyConnectorList200ResponseInner) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ThirdpartyConnectorList200ResponseInner) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetHostname

`func (o *ThirdpartyConnectorList200ResponseInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ThirdpartyConnectorList200ResponseInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetLoginProvider

`func (o *ThirdpartyConnectorList200ResponseInner) GetLoginProvider() string`

GetLoginProvider returns the LoginProvider field if non-nil, zero value otherwise.

### GetLoginProviderOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetLoginProviderOk() (*string, bool)`

GetLoginProviderOk returns a tuple with the LoginProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginProvider

`func (o *ThirdpartyConnectorList200ResponseInner) SetLoginProvider(v string)`

SetLoginProvider sets LoginProvider field to given value.

### HasLoginProvider

`func (o *ThirdpartyConnectorList200ResponseInner) HasLoginProvider() bool`

HasLoginProvider returns a boolean if a field has been set.

### SetLoginProviderNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetLoginProviderNil(b bool)`

 SetLoginProviderNil sets the value for LoginProvider to be an explicit nil

### UnsetLoginProvider
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetLoginProvider()`

UnsetLoginProvider ensures that no value is present for LoginProvider, not even an explicit nil
### GetWithChain

`func (o *ThirdpartyConnectorList200ResponseInner) GetWithChain() bool`

GetWithChain returns the WithChain field if non-nil, zero value otherwise.

### GetWithChainOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetWithChainOk() (*bool, bool)`

GetWithChainOk returns a tuple with the WithChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithChain

`func (o *ThirdpartyConnectorList200ResponseInner) SetWithChain(v bool)`

SetWithChain sets WithChain field to given value.

### HasWithChain

`func (o *ThirdpartyConnectorList200ResponseInner) HasWithChain() bool`

HasWithChain returns a boolean if a field has been set.

### SetWithChainNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetWithChainNil(b bool)`

 SetWithChainNil sets the value for WithChain to be an explicit nil

### UnsetWithChain
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetWithChain()`

UnsetWithChain ensures that no value is present for WithChain, not even an explicit nil
### GetTenant

`func (o *ThirdpartyConnectorList200ResponseInner) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ThirdpartyConnectorList200ResponseInner) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetIntuneResourceUrl

`func (o *ThirdpartyConnectorList200ResponseInner) GetIntuneResourceUrl() string`

GetIntuneResourceUrl returns the IntuneResourceUrl field if non-nil, zero value otherwise.

### GetIntuneResourceUrlOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetIntuneResourceUrlOk() (*string, bool)`

GetIntuneResourceUrlOk returns a tuple with the IntuneResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntuneResourceUrl

`func (o *ThirdpartyConnectorList200ResponseInner) SetIntuneResourceUrl(v string)`

SetIntuneResourceUrl sets IntuneResourceUrl field to given value.

### HasIntuneResourceUrl

`func (o *ThirdpartyConnectorList200ResponseInner) HasIntuneResourceUrl() bool`

HasIntuneResourceUrl returns a boolean if a field has been set.

### SetIntuneResourceUrlNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetIntuneResourceUrlNil(b bool)`

 SetIntuneResourceUrlNil sets the value for IntuneResourceUrl to be an explicit nil

### UnsetIntuneResourceUrl
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetIntuneResourceUrl()`

UnsetIntuneResourceUrl ensures that no value is present for IntuneResourceUrl, not even an explicit nil
### GetOsQueryString

`func (o *ThirdpartyConnectorList200ResponseInner) GetOsQueryString() string`

GetOsQueryString returns the OsQueryString field if non-nil, zero value otherwise.

### GetOsQueryStringOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetOsQueryStringOk() (*string, bool)`

GetOsQueryStringOk returns a tuple with the OsQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsQueryString

`func (o *ThirdpartyConnectorList200ResponseInner) SetOsQueryString(v string)`

SetOsQueryString sets OsQueryString field to given value.

### HasOsQueryString

`func (o *ThirdpartyConnectorList200ResponseInner) HasOsQueryString() bool`

HasOsQueryString returns a boolean if a field has been set.

### SetOsQueryStringNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetOsQueryStringNil(b bool)`

 SetOsQueryStringNil sets the value for OsQueryString to be an explicit nil

### UnsetOsQueryString
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetOsQueryString()`

UnsetOsQueryString ensures that no value is present for OsQueryString, not even an explicit nil
### GetLegacyRevocationMode

`func (o *ThirdpartyConnectorList200ResponseInner) GetLegacyRevocationMode() bool`

GetLegacyRevocationMode returns the LegacyRevocationMode field if non-nil, zero value otherwise.

### GetLegacyRevocationModeOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetLegacyRevocationModeOk() (*bool, bool)`

GetLegacyRevocationModeOk returns a tuple with the LegacyRevocationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyRevocationMode

`func (o *ThirdpartyConnectorList200ResponseInner) SetLegacyRevocationMode(v bool)`

SetLegacyRevocationMode sets LegacyRevocationMode field to given value.


### GetEndpoint

`func (o *ThirdpartyConnectorList200ResponseInner) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ThirdpartyConnectorList200ResponseInner) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetPort

`func (o *ThirdpartyConnectorList200ResponseInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ThirdpartyConnectorList200ResponseInner) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ThirdpartyConnectorList200ResponseInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetBaseDn

`func (o *ThirdpartyConnectorList200ResponseInner) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *ThirdpartyConnectorList200ResponseInner) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetFilter

`func (o *ThirdpartyConnectorList200ResponseInner) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ThirdpartyConnectorList200ResponseInner) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ThirdpartyConnectorList200ResponseInner) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetPubKey

`func (o *ThirdpartyConnectorList200ResponseInner) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *ThirdpartyConnectorList200ResponseInner) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetKeyName

`func (o *ThirdpartyConnectorList200ResponseInner) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *ThirdpartyConnectorList200ResponseInner) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetProviderName

`func (o *ThirdpartyConnectorList200ResponseInner) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ThirdpartyConnectorList200ResponseInner) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ThirdpartyConnectorList200ResponseInner) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetIntendedPurpose

`func (o *ThirdpartyConnectorList200ResponseInner) GetIntendedPurpose() string`

GetIntendedPurpose returns the IntendedPurpose field if non-nil, zero value otherwise.

### GetIntendedPurposeOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetIntendedPurposeOk() (*string, bool)`

GetIntendedPurposeOk returns a tuple with the IntendedPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedPurpose

`func (o *ThirdpartyConnectorList200ResponseInner) SetIntendedPurpose(v string)`

SetIntendedPurpose sets IntendedPurpose field to given value.

### HasIntendedPurpose

`func (o *ThirdpartyConnectorList200ResponseInner) HasIntendedPurpose() bool`

HasIntendedPurpose returns a boolean if a field has been set.

### SetIntendedPurposeNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetIntendedPurposeNil(b bool)`

 SetIntendedPurposeNil sets the value for IntendedPurpose to be an explicit nil

### UnsetIntendedPurpose
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetIntendedPurpose()`

UnsetIntendedPurpose ensures that no value is present for IntendedPurpose, not even an explicit nil
### GetSearchFilter

`func (o *ThirdpartyConnectorList200ResponseInner) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *ThirdpartyConnectorList200ResponseInner) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *ThirdpartyConnectorList200ResponseInner) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.

### SetSearchFilterNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetSearchFilterNil(b bool)`

 SetSearchFilterNil sets the value for SearchFilter to be an explicit nil

### UnsetSearchFilter
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetSearchFilter()`

UnsetSearchFilter ensures that no value is present for SearchFilter, not even an explicit nil
### GetVaultBaseUrl

`func (o *ThirdpartyConnectorList200ResponseInner) GetVaultBaseUrl() string`

GetVaultBaseUrl returns the VaultBaseUrl field if non-nil, zero value otherwise.

### GetVaultBaseUrlOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetVaultBaseUrlOk() (*string, bool)`

GetVaultBaseUrlOk returns a tuple with the VaultBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultBaseUrl

`func (o *ThirdpartyConnectorList200ResponseInner) SetVaultBaseUrl(v string)`

SetVaultBaseUrl sets VaultBaseUrl field to given value.


### GetProject

`func (o *ThirdpartyConnectorList200ResponseInner) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ThirdpartyConnectorList200ResponseInner) SetProject(v string)`

SetProject sets Project field to given value.


### GetLocation

`func (o *ThirdpartyConnectorList200ResponseInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ThirdpartyConnectorList200ResponseInner) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCertAttr

`func (o *ThirdpartyConnectorList200ResponseInner) GetCertAttr() string`

GetCertAttr returns the CertAttr field if non-nil, zero value otherwise.

### GetCertAttrOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetCertAttrOk() (*string, bool)`

GetCertAttrOk returns a tuple with the CertAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAttr

`func (o *ThirdpartyConnectorList200ResponseInner) SetCertAttr(v string)`

SetCertAttr sets CertAttr field to given value.

### HasCertAttr

`func (o *ThirdpartyConnectorList200ResponseInner) HasCertAttr() bool`

HasCertAttr returns a boolean if a field has been set.

### SetCertAttrNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetCertAttrNil(b bool)`

 SetCertAttrNil sets the value for CertAttr to be an explicit nil

### UnsetCertAttr
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetCertAttr()`

UnsetCertAttr ensures that no value is present for CertAttr, not even an explicit nil
### GetFollowReferrals

`func (o *ThirdpartyConnectorList200ResponseInner) GetFollowReferrals() bool`

GetFollowReferrals returns the FollowReferrals field if non-nil, zero value otherwise.

### GetFollowReferralsOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetFollowReferralsOk() (*bool, bool)`

GetFollowReferralsOk returns a tuple with the FollowReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowReferrals

`func (o *ThirdpartyConnectorList200ResponseInner) SetFollowReferrals(v bool)`

SetFollowReferrals sets FollowReferrals field to given value.

### HasFollowReferrals

`func (o *ThirdpartyConnectorList200ResponseInner) HasFollowReferrals() bool`

HasFollowReferrals returns a boolean if a field has been set.

### SetFollowReferralsNil

`func (o *ThirdpartyConnectorList200ResponseInner) SetFollowReferralsNil(b bool)`

 SetFollowReferralsNil sets the value for FollowReferrals to be an explicit nil

### UnsetFollowReferrals
`func (o *ThirdpartyConnectorList200ResponseInner) UnsetFollowReferrals()`

UnsetFollowReferrals ensures that no value is present for FollowReferrals, not even an explicit nil
### GetUserIdentifierAttribute

`func (o *ThirdpartyConnectorList200ResponseInner) GetUserIdentifierAttribute() string`

GetUserIdentifierAttribute returns the UserIdentifierAttribute field if non-nil, zero value otherwise.

### GetUserIdentifierAttributeOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetUserIdentifierAttributeOk() (*string, bool)`

GetUserIdentifierAttributeOk returns a tuple with the UserIdentifierAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifierAttribute

`func (o *ThirdpartyConnectorList200ResponseInner) SetUserIdentifierAttribute(v string)`

SetUserIdentifierAttribute sets UserIdentifierAttribute field to given value.


### GetCertificateAttribute

`func (o *ThirdpartyConnectorList200ResponseInner) GetCertificateAttribute() string`

GetCertificateAttribute returns the CertificateAttribute field if non-nil, zero value otherwise.

### GetCertificateAttributeOk

`func (o *ThirdpartyConnectorList200ResponseInner) GetCertificateAttributeOk() (*string, bool)`

GetCertificateAttributeOk returns a tuple with the CertificateAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAttribute

`func (o *ThirdpartyConnectorList200ResponseInner) SetCertificateAttribute(v string)`

SetCertificateAttribute sets CertificateAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


