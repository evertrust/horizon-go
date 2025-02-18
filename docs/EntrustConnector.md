# EntrustConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/api.security.credentials) to use for technical account on the PKI | 
**CertType** | **string** |  | 
**RequesterDefaultMail** | **string** |  | 
**RequesterName** | Pointer to **NullableString** |  | [optional] 
**RequesterPhone** | Pointer to **NullableString** |  | [optional] 
**CertLifetime** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableInt64** |  | [optional] 
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/api.security.credentials) to use to authenticate on the PKI | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEntrustConnector

`func NewEntrustConnector(name string, type_ string, loginCredentials string, certType string, requesterDefaultMail string, authenticationCredentials string, ) *EntrustConnector`

NewEntrustConnector instantiates a new EntrustConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrustConnectorWithDefaults

`func NewEntrustConnectorWithDefaults() *EntrustConnector`

NewEntrustConnectorWithDefaults instantiates a new EntrustConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntrustConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntrustConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntrustConnector) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EntrustConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntrustConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntrustConnector) SetType(v string)`

SetType sets Type field to given value.


### GetLoginCredentials

`func (o *EntrustConnector) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *EntrustConnector) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *EntrustConnector) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetCertType

`func (o *EntrustConnector) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *EntrustConnector) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *EntrustConnector) SetCertType(v string)`

SetCertType sets CertType field to given value.


### GetRequesterDefaultMail

`func (o *EntrustConnector) GetRequesterDefaultMail() string`

GetRequesterDefaultMail returns the RequesterDefaultMail field if non-nil, zero value otherwise.

### GetRequesterDefaultMailOk

`func (o *EntrustConnector) GetRequesterDefaultMailOk() (*string, bool)`

GetRequesterDefaultMailOk returns a tuple with the RequesterDefaultMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterDefaultMail

`func (o *EntrustConnector) SetRequesterDefaultMail(v string)`

SetRequesterDefaultMail sets RequesterDefaultMail field to given value.


### GetRequesterName

`func (o *EntrustConnector) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *EntrustConnector) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *EntrustConnector) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *EntrustConnector) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### SetRequesterNameNil

`func (o *EntrustConnector) SetRequesterNameNil(b bool)`

 SetRequesterNameNil sets the value for RequesterName to be an explicit nil

### UnsetRequesterName
`func (o *EntrustConnector) UnsetRequesterName()`

UnsetRequesterName ensures that no value is present for RequesterName, not even an explicit nil
### GetRequesterPhone

`func (o *EntrustConnector) GetRequesterPhone() string`

GetRequesterPhone returns the RequesterPhone field if non-nil, zero value otherwise.

### GetRequesterPhoneOk

`func (o *EntrustConnector) GetRequesterPhoneOk() (*string, bool)`

GetRequesterPhoneOk returns a tuple with the RequesterPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterPhone

`func (o *EntrustConnector) SetRequesterPhone(v string)`

SetRequesterPhone sets RequesterPhone field to given value.

### HasRequesterPhone

`func (o *EntrustConnector) HasRequesterPhone() bool`

HasRequesterPhone returns a boolean if a field has been set.

### SetRequesterPhoneNil

`func (o *EntrustConnector) SetRequesterPhoneNil(b bool)`

 SetRequesterPhoneNil sets the value for RequesterPhone to be an explicit nil

### UnsetRequesterPhone
`func (o *EntrustConnector) UnsetRequesterPhone()`

UnsetRequesterPhone ensures that no value is present for RequesterPhone, not even an explicit nil
### GetCertLifetime

`func (o *EntrustConnector) GetCertLifetime() string`

GetCertLifetime returns the CertLifetime field if non-nil, zero value otherwise.

### GetCertLifetimeOk

`func (o *EntrustConnector) GetCertLifetimeOk() (*string, bool)`

GetCertLifetimeOk returns a tuple with the CertLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertLifetime

`func (o *EntrustConnector) SetCertLifetime(v string)`

SetCertLifetime sets CertLifetime field to given value.

### HasCertLifetime

`func (o *EntrustConnector) HasCertLifetime() bool`

HasCertLifetime returns a boolean if a field has been set.

### SetCertLifetimeNil

`func (o *EntrustConnector) SetCertLifetimeNil(b bool)`

 SetCertLifetimeNil sets the value for CertLifetime to be an explicit nil

### UnsetCertLifetime
`func (o *EntrustConnector) UnsetCertLifetime()`

UnsetCertLifetime ensures that no value is present for CertLifetime, not even an explicit nil
### GetClientId

`func (o *EntrustConnector) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EntrustConnector) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EntrustConnector) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *EntrustConnector) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *EntrustConnector) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *EntrustConnector) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetAuthenticationCredentials

`func (o *EntrustConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *EntrustConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *EntrustConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetTimeout

`func (o *EntrustConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EntrustConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EntrustConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EntrustConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *EntrustConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *EntrustConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *EntrustConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *EntrustConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *EntrustConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *EntrustConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *EntrustConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *EntrustConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *EntrustConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *EntrustConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *EntrustConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *EntrustConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *EntrustConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *EntrustConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


