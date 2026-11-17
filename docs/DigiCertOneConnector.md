# DigiCertOneConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiCredentials** | **string** | Name of the &#x60;raw&#x60; [credentials](#tag/security.credentials) containing the API key to authenticate on the PKI | 
**BaseUrl** | **string** | The base URL of the DigiCert ONE instance. | 
**Name** | **string** |  | 
**ProfileId** | **string** | Identifier of the DigiCert ONE certificate profile used to issue certificates. | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | **string** | Interval between two polls of a pending certificate request. | 
**Seat** | [**DigiCertOneConnectorSeat**](DigiCertOneConnectorSeat.md) |  | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Validity** | Pointer to **string** | Overrides the validity of the certificates issued through this connector. When not set, the validity defined by the DigiCert ONE certificate profile is used. The validity is converted to minutes when sent to DigiCert ONE, so any precision below the minute is truncated. | [optional] 

## Methods

### NewDigiCertOneConnector

`func NewDigiCertOneConnector(apiCredentials string, baseUrl string, name string, profileId string, retryInterval string, seat DigiCertOneConnectorSeat, type_ string, ) *DigiCertOneConnector`

NewDigiCertOneConnector instantiates a new DigiCertOneConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigiCertOneConnectorWithDefaults

`func NewDigiCertOneConnectorWithDefaults() *DigiCertOneConnector`

NewDigiCertOneConnectorWithDefaults instantiates a new DigiCertOneConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiCredentials

`func (o *DigiCertOneConnector) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *DigiCertOneConnector) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *DigiCertOneConnector) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetBaseUrl

`func (o *DigiCertOneConnector) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DigiCertOneConnector) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DigiCertOneConnector) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetName

`func (o *DigiCertOneConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DigiCertOneConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DigiCertOneConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProfileId

`func (o *DigiCertOneConnector) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *DigiCertOneConnector) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *DigiCertOneConnector) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetProxy

`func (o *DigiCertOneConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DigiCertOneConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DigiCertOneConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DigiCertOneConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DigiCertOneConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DigiCertOneConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *DigiCertOneConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *DigiCertOneConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *DigiCertOneConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *DigiCertOneConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *DigiCertOneConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *DigiCertOneConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetRetryInterval

`func (o *DigiCertOneConnector) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *DigiCertOneConnector) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *DigiCertOneConnector) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.


### GetSeat

`func (o *DigiCertOneConnector) GetSeat() DigiCertOneConnectorSeat`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *DigiCertOneConnector) GetSeatOk() (*DigiCertOneConnectorSeat, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *DigiCertOneConnector) SetSeat(v DigiCertOneConnectorSeat)`

SetSeat sets Seat field to given value.


### GetTimeout

`func (o *DigiCertOneConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DigiCertOneConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DigiCertOneConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DigiCertOneConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *DigiCertOneConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DigiCertOneConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *DigiCertOneConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigiCertOneConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigiCertOneConnector) SetType(v string)`

SetType sets Type field to given value.


### GetValidity

`func (o *DigiCertOneConnector) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *DigiCertOneConnector) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *DigiCertOneConnector) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *DigiCertOneConnector) HasValidity() bool`

HasValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


