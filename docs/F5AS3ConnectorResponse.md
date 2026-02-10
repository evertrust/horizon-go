# F5AS3ConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**KeyType** | Pointer to **NullableString** | One of &#x60;rsa-2048&#x60;, &#x60;rsa-3072&#x60;, &#x60;rsa-4096&#x60;, &#x60;rsa-8192&#x60;, &#x60;ec-secp256r1&#x60;, &#x60;ec-secp384r1&#x60;, &#x60;ec-secp521r1&#x60;, &#x60;ed-448&#x60;, &#x60;ed-25519&#x60;, &#x60;mldsa-44&#x60;, &#x60;mldsa-65&#x60;, &#x60;mldsa-87&#x60;, &#x60;slhdsa-sha2-128s&#x60;, &#x60;slhdsa-sha2-128f&#x60;, &#x60;slhdsa-sha2-192s&#x60;, &#x60;slhdsa-sha2-192f&#x60;, &#x60;slhdsa-sha2-256s&#x60;, &#x60;slhdsa-sha2-256f&#x60;, &#x60;slhdsa-sha2-128ssha256&#x60;, &#x60;slhdsa-sha2-128fsha256&#x60;, &#x60;slhdsa-sha2-192ssha512&#x60;, &#x60;slhdsa-sha2-192fsha512&#x60;, &#x60;slhdsa-sha2-256ssha512&#x60;, &#x60;slhdsa-sha2-256fsha512&#x60; or &#x60;&lt;primary key type&gt;+&lt;alternate key type&gt;&#x60; | [optional] 
**RenewalPeriod** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Hostname** | **string** |  | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the account to authenticate on F5 | 
**LoginProvider** | Pointer to **NullableString** |  | [optional] 
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**WithChain** | Pointer to **NullableBool** | Enable the certificate trust chain to be pushed. | [optional] [default to true]

## Methods

### NewF5AS3ConnectorResponse

`func NewF5AS3ConnectorResponse(id string, type_ string, name string, throttleDuration string, throttleParallelism int64, hostname string, credentials string, ) *F5AS3ConnectorResponse`

NewF5AS3ConnectorResponse instantiates a new F5AS3ConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewF5AS3ConnectorResponseWithDefaults

`func NewF5AS3ConnectorResponseWithDefaults() *F5AS3ConnectorResponse`

NewF5AS3ConnectorResponseWithDefaults instantiates a new F5AS3ConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *F5AS3ConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *F5AS3ConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *F5AS3ConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *F5AS3ConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *F5AS3ConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *F5AS3ConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *F5AS3ConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *F5AS3ConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *F5AS3ConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetThrottleDuration

`func (o *F5AS3ConnectorResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *F5AS3ConnectorResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *F5AS3ConnectorResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *F5AS3ConnectorResponse) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *F5AS3ConnectorResponse) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *F5AS3ConnectorResponse) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetKeyType

`func (o *F5AS3ConnectorResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *F5AS3ConnectorResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *F5AS3ConnectorResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *F5AS3ConnectorResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *F5AS3ConnectorResponse) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *F5AS3ConnectorResponse) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetRenewalPeriod

`func (o *F5AS3ConnectorResponse) GetRenewalPeriod() string`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *F5AS3ConnectorResponse) GetRenewalPeriodOk() (*string, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *F5AS3ConnectorResponse) SetRenewalPeriod(v string)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *F5AS3ConnectorResponse) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### SetRenewalPeriodNil

`func (o *F5AS3ConnectorResponse) SetRenewalPeriodNil(b bool)`

 SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil

### UnsetRenewalPeriod
`func (o *F5AS3ConnectorResponse) UnsetRenewalPeriod()`

UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
### GetTimeout

`func (o *F5AS3ConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *F5AS3ConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *F5AS3ConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *F5AS3ConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *F5AS3ConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *F5AS3ConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *F5AS3ConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *F5AS3ConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *F5AS3ConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *F5AS3ConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *F5AS3ConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *F5AS3ConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetHostname

`func (o *F5AS3ConnectorResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *F5AS3ConnectorResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *F5AS3ConnectorResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetCredentials

`func (o *F5AS3ConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *F5AS3ConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *F5AS3ConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetLoginProvider

`func (o *F5AS3ConnectorResponse) GetLoginProvider() string`

GetLoginProvider returns the LoginProvider field if non-nil, zero value otherwise.

### GetLoginProviderOk

`func (o *F5AS3ConnectorResponse) GetLoginProviderOk() (*string, bool)`

GetLoginProviderOk returns a tuple with the LoginProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginProvider

`func (o *F5AS3ConnectorResponse) SetLoginProvider(v string)`

SetLoginProvider sets LoginProvider field to given value.

### HasLoginProvider

`func (o *F5AS3ConnectorResponse) HasLoginProvider() bool`

HasLoginProvider returns a boolean if a field has been set.

### SetLoginProviderNil

`func (o *F5AS3ConnectorResponse) SetLoginProviderNil(b bool)`

 SetLoginProviderNil sets the value for LoginProvider to be an explicit nil

### UnsetLoginProvider
`func (o *F5AS3ConnectorResponse) UnsetLoginProvider()`

UnsetLoginProvider ensures that no value is present for LoginProvider, not even an explicit nil
### GetTlsInsecure

`func (o *F5AS3ConnectorResponse) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *F5AS3ConnectorResponse) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *F5AS3ConnectorResponse) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *F5AS3ConnectorResponse) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *F5AS3ConnectorResponse) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *F5AS3ConnectorResponse) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetWithChain

`func (o *F5AS3ConnectorResponse) GetWithChain() bool`

GetWithChain returns the WithChain field if non-nil, zero value otherwise.

### GetWithChainOk

`func (o *F5AS3ConnectorResponse) GetWithChainOk() (*bool, bool)`

GetWithChainOk returns a tuple with the WithChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithChain

`func (o *F5AS3ConnectorResponse) SetWithChain(v bool)`

SetWithChain sets WithChain field to given value.

### HasWithChain

`func (o *F5AS3ConnectorResponse) HasWithChain() bool`

HasWithChain returns a boolean if a field has been set.

### SetWithChainNil

`func (o *F5AS3ConnectorResponse) SetWithChainNil(b bool)`

 SetWithChainNil sets the value for WithChain to be an explicit nil

### UnsetWithChain
`func (o *F5AS3ConnectorResponse) UnsetWithChain()`

UnsetWithChain ensures that no value is present for WithChain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


