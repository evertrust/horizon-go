# EverTrustADCSConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaConfig** | **string** |  | 
**Domain** | **string** |  | 
**EndPoint** | **string** |  | 
**EnrollmentCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to enroll on the PKI | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use for technical account on the PKI | 
**Name** | **string** |  | 
**Profile** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewEverTrustADCSConnector

`func NewEverTrustADCSConnector(caConfig string, domain string, endPoint string, enrollmentCredentials string, loginCredentials string, name string, profile string, type_ string, ) *EverTrustADCSConnector`

NewEverTrustADCSConnector instantiates a new EverTrustADCSConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEverTrustADCSConnectorWithDefaults

`func NewEverTrustADCSConnectorWithDefaults() *EverTrustADCSConnector`

NewEverTrustADCSConnectorWithDefaults instantiates a new EverTrustADCSConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaConfig

`func (o *EverTrustADCSConnector) GetCaConfig() string`

GetCaConfig returns the CaConfig field if non-nil, zero value otherwise.

### GetCaConfigOk

`func (o *EverTrustADCSConnector) GetCaConfigOk() (*string, bool)`

GetCaConfigOk returns a tuple with the CaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaConfig

`func (o *EverTrustADCSConnector) SetCaConfig(v string)`

SetCaConfig sets CaConfig field to given value.


### GetDomain

`func (o *EverTrustADCSConnector) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EverTrustADCSConnector) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EverTrustADCSConnector) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEndPoint

`func (o *EverTrustADCSConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *EverTrustADCSConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *EverTrustADCSConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetEnrollmentCredentials

`func (o *EverTrustADCSConnector) GetEnrollmentCredentials() string`

GetEnrollmentCredentials returns the EnrollmentCredentials field if non-nil, zero value otherwise.

### GetEnrollmentCredentialsOk

`func (o *EverTrustADCSConnector) GetEnrollmentCredentialsOk() (*string, bool)`

GetEnrollmentCredentialsOk returns a tuple with the EnrollmentCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCredentials

`func (o *EverTrustADCSConnector) SetEnrollmentCredentials(v string)`

SetEnrollmentCredentials sets EnrollmentCredentials field to given value.


### GetLoginCredentials

`func (o *EverTrustADCSConnector) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *EverTrustADCSConnector) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *EverTrustADCSConnector) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetName

`func (o *EverTrustADCSConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EverTrustADCSConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EverTrustADCSConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *EverTrustADCSConnector) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *EverTrustADCSConnector) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *EverTrustADCSConnector) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProxy

`func (o *EverTrustADCSConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *EverTrustADCSConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *EverTrustADCSConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *EverTrustADCSConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *EverTrustADCSConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *EverTrustADCSConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *EverTrustADCSConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *EverTrustADCSConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *EverTrustADCSConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *EverTrustADCSConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *EverTrustADCSConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *EverTrustADCSConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTimeout

`func (o *EverTrustADCSConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EverTrustADCSConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EverTrustADCSConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EverTrustADCSConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *EverTrustADCSConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *EverTrustADCSConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *EverTrustADCSConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EverTrustADCSConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EverTrustADCSConnector) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


