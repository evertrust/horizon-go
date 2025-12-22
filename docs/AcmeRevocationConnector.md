# AcmeRevocationConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeDirectoryUrl** | **string** | The directory url of the ACME endpoint | 
**Name** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAcmeRevocationConnector

`func NewAcmeRevocationConnector(acmeDirectoryUrl string, name string, type_ string, ) *AcmeRevocationConnector`

NewAcmeRevocationConnector instantiates a new AcmeRevocationConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeRevocationConnectorWithDefaults

`func NewAcmeRevocationConnectorWithDefaults() *AcmeRevocationConnector`

NewAcmeRevocationConnectorWithDefaults instantiates a new AcmeRevocationConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeDirectoryUrl

`func (o *AcmeRevocationConnector) GetAcmeDirectoryUrl() string`

GetAcmeDirectoryUrl returns the AcmeDirectoryUrl field if non-nil, zero value otherwise.

### GetAcmeDirectoryUrlOk

`func (o *AcmeRevocationConnector) GetAcmeDirectoryUrlOk() (*string, bool)`

GetAcmeDirectoryUrlOk returns a tuple with the AcmeDirectoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectoryUrl

`func (o *AcmeRevocationConnector) SetAcmeDirectoryUrl(v string)`

SetAcmeDirectoryUrl sets AcmeDirectoryUrl field to given value.


### GetName

`func (o *AcmeRevocationConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcmeRevocationConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcmeRevocationConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *AcmeRevocationConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AcmeRevocationConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AcmeRevocationConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AcmeRevocationConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *AcmeRevocationConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *AcmeRevocationConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *AcmeRevocationConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *AcmeRevocationConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *AcmeRevocationConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *AcmeRevocationConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *AcmeRevocationConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *AcmeRevocationConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTimeout

`func (o *AcmeRevocationConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AcmeRevocationConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AcmeRevocationConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AcmeRevocationConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *AcmeRevocationConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *AcmeRevocationConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *AcmeRevocationConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcmeRevocationConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcmeRevocationConnector) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


