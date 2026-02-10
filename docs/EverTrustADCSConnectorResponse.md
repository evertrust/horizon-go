# EverTrustADCSConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**EndPoint** | **string** |  | 
**CaConfig** | **string** |  | 
**Profile** | **string** |  | 
**Domain** | **string** |  | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use for technical account on the PKI | 
**EnrollmentCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to enroll on the PKI | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 

## Methods

### NewEverTrustADCSConnectorResponse

`func NewEverTrustADCSConnectorResponse(id string, name string, type_ string, endPoint string, caConfig string, profile string, domain string, loginCredentials string, enrollmentCredentials string, ) *EverTrustADCSConnectorResponse`

NewEverTrustADCSConnectorResponse instantiates a new EverTrustADCSConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEverTrustADCSConnectorResponseWithDefaults

`func NewEverTrustADCSConnectorResponseWithDefaults() *EverTrustADCSConnectorResponse`

NewEverTrustADCSConnectorResponseWithDefaults instantiates a new EverTrustADCSConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EverTrustADCSConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EverTrustADCSConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EverTrustADCSConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EverTrustADCSConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EverTrustADCSConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EverTrustADCSConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EverTrustADCSConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EverTrustADCSConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EverTrustADCSConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetEndPoint

`func (o *EverTrustADCSConnectorResponse) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *EverTrustADCSConnectorResponse) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *EverTrustADCSConnectorResponse) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetCaConfig

`func (o *EverTrustADCSConnectorResponse) GetCaConfig() string`

GetCaConfig returns the CaConfig field if non-nil, zero value otherwise.

### GetCaConfigOk

`func (o *EverTrustADCSConnectorResponse) GetCaConfigOk() (*string, bool)`

GetCaConfigOk returns a tuple with the CaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaConfig

`func (o *EverTrustADCSConnectorResponse) SetCaConfig(v string)`

SetCaConfig sets CaConfig field to given value.


### GetProfile

`func (o *EverTrustADCSConnectorResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *EverTrustADCSConnectorResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *EverTrustADCSConnectorResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetDomain

`func (o *EverTrustADCSConnectorResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EverTrustADCSConnectorResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EverTrustADCSConnectorResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetLoginCredentials

`func (o *EverTrustADCSConnectorResponse) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *EverTrustADCSConnectorResponse) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *EverTrustADCSConnectorResponse) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetEnrollmentCredentials

`func (o *EverTrustADCSConnectorResponse) GetEnrollmentCredentials() string`

GetEnrollmentCredentials returns the EnrollmentCredentials field if non-nil, zero value otherwise.

### GetEnrollmentCredentialsOk

`func (o *EverTrustADCSConnectorResponse) GetEnrollmentCredentialsOk() (*string, bool)`

GetEnrollmentCredentialsOk returns a tuple with the EnrollmentCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCredentials

`func (o *EverTrustADCSConnectorResponse) SetEnrollmentCredentials(v string)`

SetEnrollmentCredentials sets EnrollmentCredentials field to given value.


### GetTimeout

`func (o *EverTrustADCSConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EverTrustADCSConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EverTrustADCSConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EverTrustADCSConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *EverTrustADCSConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *EverTrustADCSConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *EverTrustADCSConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *EverTrustADCSConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *EverTrustADCSConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *EverTrustADCSConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *EverTrustADCSConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *EverTrustADCSConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *EverTrustADCSConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *EverTrustADCSConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *EverTrustADCSConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *EverTrustADCSConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *EverTrustADCSConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *EverTrustADCSConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *EverTrustADCSConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EverTrustADCSConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EverTrustADCSConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EverTrustADCSConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EverTrustADCSConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EverTrustADCSConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


