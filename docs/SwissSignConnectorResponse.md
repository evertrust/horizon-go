# SwissSignConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**MpkiCredentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI.  It should contains the mpkiId as the login and the apiKey as password.  | 
**EndPoint** | **string** | Swiss base endpoint | 
**ProductUuid** | **string** | The product Uuid that need to be retrieved from the swiss sign api&#39;s (&lt;endpoints&gt;/v2/clients) | 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 

## Methods

### NewSwissSignConnectorResponse

`func NewSwissSignConnectorResponse(id string, name string, type_ string, mpkiCredentials string, endPoint string, productUuid string, ) *SwissSignConnectorResponse`

NewSwissSignConnectorResponse instantiates a new SwissSignConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwissSignConnectorResponseWithDefaults

`func NewSwissSignConnectorResponseWithDefaults() *SwissSignConnectorResponse`

NewSwissSignConnectorResponseWithDefaults instantiates a new SwissSignConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SwissSignConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SwissSignConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SwissSignConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SwissSignConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwissSignConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwissSignConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *SwissSignConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SwissSignConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SwissSignConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SwissSignConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *SwissSignConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *SwissSignConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *SwissSignConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SwissSignConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SwissSignConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SwissSignConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *SwissSignConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SwissSignConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetQueue

`func (o *SwissSignConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *SwissSignConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *SwissSignConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *SwissSignConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *SwissSignConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *SwissSignConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetType

`func (o *SwissSignConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwissSignConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwissSignConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetMpkiCredentials

`func (o *SwissSignConnectorResponse) GetMpkiCredentials() string`

GetMpkiCredentials returns the MpkiCredentials field if non-nil, zero value otherwise.

### GetMpkiCredentialsOk

`func (o *SwissSignConnectorResponse) GetMpkiCredentialsOk() (*string, bool)`

GetMpkiCredentialsOk returns a tuple with the MpkiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpkiCredentials

`func (o *SwissSignConnectorResponse) SetMpkiCredentials(v string)`

SetMpkiCredentials sets MpkiCredentials field to given value.


### GetEndPoint

`func (o *SwissSignConnectorResponse) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *SwissSignConnectorResponse) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *SwissSignConnectorResponse) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetProductUuid

`func (o *SwissSignConnectorResponse) GetProductUuid() string`

GetProductUuid returns the ProductUuid field if non-nil, zero value otherwise.

### GetProductUuidOk

`func (o *SwissSignConnectorResponse) GetProductUuidOk() (*string, bool)`

GetProductUuidOk returns a tuple with the ProductUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUuid

`func (o *SwissSignConnectorResponse) SetProductUuid(v string)`

SetProductUuid sets ProductUuid field to given value.


### GetStatus

`func (o *SwissSignConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwissSignConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwissSignConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SwissSignConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SwissSignConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SwissSignConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


