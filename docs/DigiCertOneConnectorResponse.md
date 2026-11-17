# DigiCertOneConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ApiCredentials** | **string** | Name of the &#x60;raw&#x60; [credentials](#tag/security.credentials) containing the API key to authenticate on the PKI | 
**BaseUrl** | **string** | The base URL of the DigiCert ONE instance. | 
**Name** | **string** |  | 
**ProfileId** | **string** | Identifier of the DigiCert ONE certificate profile used to issue certificates. | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | **string** | Interval between two polls of a pending certificate request. | 
**Seat** | [**DigiCertOneConnectorSeat**](DigiCertOneConnectorSeat.md) |  | 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Validity** | Pointer to **string** | Overrides the validity of the certificates issued through this connector. When not set, the validity defined by the DigiCert ONE certificate profile is used. The validity is converted to minutes when sent to DigiCert ONE, so any precision below the minute is truncated. | [optional] 

## Methods

### NewDigiCertOneConnectorResponse

`func NewDigiCertOneConnectorResponse(id string, apiCredentials string, baseUrl string, name string, profileId string, retryInterval string, seat DigiCertOneConnectorSeat, type_ string, ) *DigiCertOneConnectorResponse`

NewDigiCertOneConnectorResponse instantiates a new DigiCertOneConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigiCertOneConnectorResponseWithDefaults

`func NewDigiCertOneConnectorResponseWithDefaults() *DigiCertOneConnectorResponse`

NewDigiCertOneConnectorResponseWithDefaults instantiates a new DigiCertOneConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DigiCertOneConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DigiCertOneConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DigiCertOneConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApiCredentials

`func (o *DigiCertOneConnectorResponse) GetApiCredentials() string`

GetApiCredentials returns the ApiCredentials field if non-nil, zero value otherwise.

### GetApiCredentialsOk

`func (o *DigiCertOneConnectorResponse) GetApiCredentialsOk() (*string, bool)`

GetApiCredentialsOk returns a tuple with the ApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentials

`func (o *DigiCertOneConnectorResponse) SetApiCredentials(v string)`

SetApiCredentials sets ApiCredentials field to given value.


### GetBaseUrl

`func (o *DigiCertOneConnectorResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DigiCertOneConnectorResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DigiCertOneConnectorResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetName

`func (o *DigiCertOneConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DigiCertOneConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DigiCertOneConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProfileId

`func (o *DigiCertOneConnectorResponse) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *DigiCertOneConnectorResponse) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *DigiCertOneConnectorResponse) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetProxy

`func (o *DigiCertOneConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *DigiCertOneConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *DigiCertOneConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *DigiCertOneConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *DigiCertOneConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *DigiCertOneConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *DigiCertOneConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *DigiCertOneConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *DigiCertOneConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *DigiCertOneConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *DigiCertOneConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *DigiCertOneConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetRetryInterval

`func (o *DigiCertOneConnectorResponse) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *DigiCertOneConnectorResponse) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *DigiCertOneConnectorResponse) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.


### GetSeat

`func (o *DigiCertOneConnectorResponse) GetSeat() DigiCertOneConnectorSeat`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *DigiCertOneConnectorResponse) GetSeatOk() (*DigiCertOneConnectorSeat, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *DigiCertOneConnectorResponse) SetSeat(v DigiCertOneConnectorSeat)`

SetSeat sets Seat field to given value.


### GetStatus

`func (o *DigiCertOneConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DigiCertOneConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DigiCertOneConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DigiCertOneConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DigiCertOneConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DigiCertOneConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTimeout

`func (o *DigiCertOneConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DigiCertOneConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DigiCertOneConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DigiCertOneConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *DigiCertOneConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DigiCertOneConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *DigiCertOneConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigiCertOneConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigiCertOneConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetValidity

`func (o *DigiCertOneConnectorResponse) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *DigiCertOneConnectorResponse) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *DigiCertOneConnectorResponse) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *DigiCertOneConnectorResponse) HasValidity() bool`

HasValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


