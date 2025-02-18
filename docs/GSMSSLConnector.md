# GSMSSLConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**EndpointType** | **string** |  | 
**Profile** | **string** |  | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/api.security.credentials) to use for technical account on the PKI | 
**DomainId** | **string** |  | 
**CertificateValidity** | Pointer to **NullableInt64** |  | [optional] 
**DefaultEmail** | Pointer to **NullableString** |  | [optional] 
**DefaultPhone** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGSMSSLConnector

`func NewGSMSSLConnector(name string, type_ string, endpointType string, profile string, loginCredentials string, domainId string, ) *GSMSSLConnector`

NewGSMSSLConnector instantiates a new GSMSSLConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGSMSSLConnectorWithDefaults

`func NewGSMSSLConnectorWithDefaults() *GSMSSLConnector`

NewGSMSSLConnectorWithDefaults instantiates a new GSMSSLConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GSMSSLConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GSMSSLConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GSMSSLConnector) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *GSMSSLConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GSMSSLConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GSMSSLConnector) SetType(v string)`

SetType sets Type field to given value.


### GetEndpointType

`func (o *GSMSSLConnector) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *GSMSSLConnector) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *GSMSSLConnector) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetProfile

`func (o *GSMSSLConnector) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GSMSSLConnector) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GSMSSLConnector) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetLoginCredentials

`func (o *GSMSSLConnector) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *GSMSSLConnector) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *GSMSSLConnector) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetDomainId

`func (o *GSMSSLConnector) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *GSMSSLConnector) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *GSMSSLConnector) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetCertificateValidity

`func (o *GSMSSLConnector) GetCertificateValidity() int64`

GetCertificateValidity returns the CertificateValidity field if non-nil, zero value otherwise.

### GetCertificateValidityOk

`func (o *GSMSSLConnector) GetCertificateValidityOk() (*int64, bool)`

GetCertificateValidityOk returns a tuple with the CertificateValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidity

`func (o *GSMSSLConnector) SetCertificateValidity(v int64)`

SetCertificateValidity sets CertificateValidity field to given value.

### HasCertificateValidity

`func (o *GSMSSLConnector) HasCertificateValidity() bool`

HasCertificateValidity returns a boolean if a field has been set.

### SetCertificateValidityNil

`func (o *GSMSSLConnector) SetCertificateValidityNil(b bool)`

 SetCertificateValidityNil sets the value for CertificateValidity to be an explicit nil

### UnsetCertificateValidity
`func (o *GSMSSLConnector) UnsetCertificateValidity()`

UnsetCertificateValidity ensures that no value is present for CertificateValidity, not even an explicit nil
### GetDefaultEmail

`func (o *GSMSSLConnector) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *GSMSSLConnector) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *GSMSSLConnector) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.

### HasDefaultEmail

`func (o *GSMSSLConnector) HasDefaultEmail() bool`

HasDefaultEmail returns a boolean if a field has been set.

### SetDefaultEmailNil

`func (o *GSMSSLConnector) SetDefaultEmailNil(b bool)`

 SetDefaultEmailNil sets the value for DefaultEmail to be an explicit nil

### UnsetDefaultEmail
`func (o *GSMSSLConnector) UnsetDefaultEmail()`

UnsetDefaultEmail ensures that no value is present for DefaultEmail, not even an explicit nil
### GetDefaultPhone

`func (o *GSMSSLConnector) GetDefaultPhone() string`

GetDefaultPhone returns the DefaultPhone field if non-nil, zero value otherwise.

### GetDefaultPhoneOk

`func (o *GSMSSLConnector) GetDefaultPhoneOk() (*string, bool)`

GetDefaultPhoneOk returns a tuple with the DefaultPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhone

`func (o *GSMSSLConnector) SetDefaultPhone(v string)`

SetDefaultPhone sets DefaultPhone field to given value.

### HasDefaultPhone

`func (o *GSMSSLConnector) HasDefaultPhone() bool`

HasDefaultPhone returns a boolean if a field has been set.

### SetDefaultPhoneNil

`func (o *GSMSSLConnector) SetDefaultPhoneNil(b bool)`

 SetDefaultPhoneNil sets the value for DefaultPhone to be an explicit nil

### UnsetDefaultPhone
`func (o *GSMSSLConnector) UnsetDefaultPhone()`

UnsetDefaultPhone ensures that no value is present for DefaultPhone, not even an explicit nil
### GetRetryInterval

`func (o *GSMSSLConnector) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *GSMSSLConnector) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *GSMSSLConnector) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *GSMSSLConnector) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *GSMSSLConnector) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *GSMSSLConnector) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetTimeout

`func (o *GSMSSLConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GSMSSLConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GSMSSLConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GSMSSLConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *GSMSSLConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *GSMSSLConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *GSMSSLConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *GSMSSLConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *GSMSSLConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *GSMSSLConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *GSMSSLConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *GSMSSLConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *GSMSSLConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *GSMSSLConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *GSMSSLConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *GSMSSLConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *GSMSSLConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *GSMSSLConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


