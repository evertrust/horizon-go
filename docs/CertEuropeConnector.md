# CertEuropeConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**EndPoint** | **string** |  | 
**LoginCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use for technical account on the PKI | 
**Name** | **string** |  | 
**OfferId** | **string** |  | 
**OrganizationId** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | Pointer to **NullableString** |  | [optional] 
**RevReason** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewCertEuropeConnector

`func NewCertEuropeConnector(authenticationCredentials string, endPoint string, loginCredentials string, name string, offerId string, organizationId string, type_ string, ) *CertEuropeConnector`

NewCertEuropeConnector instantiates a new CertEuropeConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertEuropeConnectorWithDefaults

`func NewCertEuropeConnectorWithDefaults() *CertEuropeConnector`

NewCertEuropeConnectorWithDefaults instantiates a new CertEuropeConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationCredentials

`func (o *CertEuropeConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *CertEuropeConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *CertEuropeConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetEndPoint

`func (o *CertEuropeConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *CertEuropeConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *CertEuropeConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetLoginCredentials

`func (o *CertEuropeConnector) GetLoginCredentials() string`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *CertEuropeConnector) GetLoginCredentialsOk() (*string, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *CertEuropeConnector) SetLoginCredentials(v string)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetName

`func (o *CertEuropeConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertEuropeConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertEuropeConnector) SetName(v string)`

SetName sets Name field to given value.


### GetOfferId

`func (o *CertEuropeConnector) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *CertEuropeConnector) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *CertEuropeConnector) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetOrganizationId

`func (o *CertEuropeConnector) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CertEuropeConnector) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CertEuropeConnector) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProxy

`func (o *CertEuropeConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CertEuropeConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CertEuropeConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CertEuropeConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CertEuropeConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CertEuropeConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *CertEuropeConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CertEuropeConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CertEuropeConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CertEuropeConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *CertEuropeConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *CertEuropeConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetRetryInterval

`func (o *CertEuropeConnector) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *CertEuropeConnector) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *CertEuropeConnector) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *CertEuropeConnector) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### SetRetryIntervalNil

`func (o *CertEuropeConnector) SetRetryIntervalNil(b bool)`

 SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil

### UnsetRetryInterval
`func (o *CertEuropeConnector) UnsetRetryInterval()`

UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
### GetRevReason

`func (o *CertEuropeConnector) GetRevReason() string`

GetRevReason returns the RevReason field if non-nil, zero value otherwise.

### GetRevReasonOk

`func (o *CertEuropeConnector) GetRevReasonOk() (*string, bool)`

GetRevReasonOk returns a tuple with the RevReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevReason

`func (o *CertEuropeConnector) SetRevReason(v string)`

SetRevReason sets RevReason field to given value.

### HasRevReason

`func (o *CertEuropeConnector) HasRevReason() bool`

HasRevReason returns a boolean if a field has been set.

### SetRevReasonNil

`func (o *CertEuropeConnector) SetRevReasonNil(b bool)`

 SetRevReasonNil sets the value for RevReason to be an explicit nil

### UnsetRevReason
`func (o *CertEuropeConnector) UnsetRevReason()`

UnsetRevReason ensures that no value is present for RevReason, not even an explicit nil
### GetTimeout

`func (o *CertEuropeConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CertEuropeConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CertEuropeConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CertEuropeConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *CertEuropeConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *CertEuropeConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *CertEuropeConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertEuropeConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertEuropeConnector) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


