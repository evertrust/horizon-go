# SwissSignConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**MpkiCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI.  It should contains the mpkiId as the login and the apiKey as password.  | 
**EndPoint** | **string** | Swiss base endpoint | 
**ProductUuid** | **string** | The product Uuid that need to be retrieved from the swiss sign api&#39;s (&lt;endpoints&gt;/v2/clients) | 

## Methods

### NewSwissSignConnector

`func NewSwissSignConnector(name string, type_ string, mpkiCredentials string, endPoint string, productUuid string, ) *SwissSignConnector`

NewSwissSignConnector instantiates a new SwissSignConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwissSignConnectorWithDefaults

`func NewSwissSignConnectorWithDefaults() *SwissSignConnector`

NewSwissSignConnectorWithDefaults instantiates a new SwissSignConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SwissSignConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwissSignConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwissSignConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *SwissSignConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SwissSignConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SwissSignConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SwissSignConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *SwissSignConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *SwissSignConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *SwissSignConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SwissSignConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SwissSignConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SwissSignConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *SwissSignConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SwissSignConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetQueue

`func (o *SwissSignConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *SwissSignConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *SwissSignConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *SwissSignConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *SwissSignConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *SwissSignConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetType

`func (o *SwissSignConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwissSignConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwissSignConnector) SetType(v string)`

SetType sets Type field to given value.


### GetMpkiCredentials

`func (o *SwissSignConnector) GetMpkiCredentials() string`

GetMpkiCredentials returns the MpkiCredentials field if non-nil, zero value otherwise.

### GetMpkiCredentialsOk

`func (o *SwissSignConnector) GetMpkiCredentialsOk() (*string, bool)`

GetMpkiCredentialsOk returns a tuple with the MpkiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpkiCredentials

`func (o *SwissSignConnector) SetMpkiCredentials(v string)`

SetMpkiCredentials sets MpkiCredentials field to given value.


### GetEndPoint

`func (o *SwissSignConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *SwissSignConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *SwissSignConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetProductUuid

`func (o *SwissSignConnector) GetProductUuid() string`

GetProductUuid returns the ProductUuid field if non-nil, zero value otherwise.

### GetProductUuidOk

`func (o *SwissSignConnector) GetProductUuidOk() (*string, bool)`

GetProductUuidOk returns a tuple with the ProductUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUuid

`func (o *SwissSignConnector) SetProductUuid(v string)`

SetProductUuid sets ProductUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


