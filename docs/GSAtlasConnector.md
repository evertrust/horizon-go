# GSAtlasConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use for technical account on the PKI | 
**HashAlgorithm** | Pointer to **NullableString** |  | [optional] 
**CertificateUsage** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | Pointer to **NullableString** |  | [optional] 
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGSAtlasConnector

`func NewGSAtlasConnector(name string, type_ string, loginCredentials string, authenticationCredentials string, ) *GSAtlasConnector`

NewGSAtlasConnector instantiates a new GSAtlasConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGSAtlasConnectorWithDefaults

`func NewGSAtlasConnectorWithDefaults() *GSAtlasConnector`

NewGSAtlasConnectorWithDefaults instantiates a new GSAtlasConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GSAtlasConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GSAtlasConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GSAtlasConnector) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *GSAtlasConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GSAtlasConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GSAtlasConnector) SetType(v string)`

SetType sets Type field to given value.


### GetLoginCredentials

`func (o *GSAtlasConnector) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *GSAtlasConnector) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *GSAtlasConnector) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetHashAlgorithm

`func (o *GSAtlasConnector) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *GSAtlasConnector) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *GSAtlasConnector) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *GSAtlasConnector) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### SetHashAlgorithmNil

`func (o *GSAtlasConnector) SetHashAlgorithmNil(b bool)`

 SetHashAlgorithmNil sets the value for HashAlgorithm to be an explicit nil

### UnsetHashAlgorithm
`func (o *GSAtlasConnector) UnsetHashAlgorithm()`

UnsetHashAlgorithm ensures that no value is present for HashAlgorithm, not even an explicit nil
### GetCertificateUsage

`func (o *GSAtlasConnector) GetCertificateUsage() string`

GetCertificateUsage returns the CertificateUsage field if non-nil, zero value otherwise.

### GetCertificateUsageOk

`func (o *GSAtlasConnector) GetCertificateUsageOk() (*string, bool)`

GetCertificateUsageOk returns a tuple with the CertificateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUsage

`func (o *GSAtlasConnector) SetCertificateUsage(v string)`

SetCertificateUsage sets CertificateUsage field to given value.

### HasCertificateUsage

`func (o *GSAtlasConnector) HasCertificateUsage() bool`

HasCertificateUsage returns a boolean if a field has been set.

### SetCertificateUsageNil

`func (o *GSAtlasConnector) SetCertificateUsageNil(b bool)`

 SetCertificateUsageNil sets the value for CertificateUsage to be an explicit nil

### UnsetCertificateUsage
`func (o *GSAtlasConnector) UnsetCertificateUsage()`

UnsetCertificateUsage ensures that no value is present for CertificateUsage, not even an explicit nil
### GetRetryInterval

`func (o *GSAtlasConnector) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *GSAtlasConnector) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *GSAtlasConnector) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *GSAtlasConnector) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *GSAtlasConnector) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *GSAtlasConnector) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetAuthenticationCredentials

`func (o *GSAtlasConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *GSAtlasConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *GSAtlasConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetTimeout

`func (o *GSAtlasConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GSAtlasConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GSAtlasConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GSAtlasConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *GSAtlasConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *GSAtlasConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *GSAtlasConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *GSAtlasConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *GSAtlasConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *GSAtlasConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *GSAtlasConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *GSAtlasConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *GSAtlasConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *GSAtlasConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *GSAtlasConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *GSAtlasConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *GSAtlasConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *GSAtlasConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


