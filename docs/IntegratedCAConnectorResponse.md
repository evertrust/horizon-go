# IntegratedCAConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AsyncParams** | Pointer to [**IntegratedCAConnectorAsyncParams**](IntegratedCAConnectorAsyncParams.md) |  | [optional] 
**CaCert** | Pointer to **NullableString** |  | [optional] 
**CaKey** | Pointer to [**NullableSecretString**](SecretString.md) |  | [optional] 
**CertType** | Pointer to **NullableString** |  | [optional] 
**CheckPop** | Pointer to **NullableBool** |  | [optional] 
**CrlLifetime** | Pointer to **NullableString** |  | [optional] 
**CrlPath** | Pointer to **NullableString** |  | [optional] 
**CrtBackDate** | Pointer to **NullableString** |  | [optional] 
**CrtLifetime** | Pointer to **NullableString** |  | [optional] 
**CryptoType** | **string** |  | 
**Name** | **string** |  | 
**Queue** | Pointer to **NullableString** |  | [optional] 
**RetryInterval** | Pointer to **string** | Interval between retry attempts for asynchronous enrollment polling | [optional] 
**SignAlg** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewIntegratedCAConnectorResponse

`func NewIntegratedCAConnectorResponse(id string, cryptoType string, name string, type_ string, ) *IntegratedCAConnectorResponse`

NewIntegratedCAConnectorResponse instantiates a new IntegratedCAConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegratedCAConnectorResponseWithDefaults

`func NewIntegratedCAConnectorResponseWithDefaults() *IntegratedCAConnectorResponse`

NewIntegratedCAConnectorResponseWithDefaults instantiates a new IntegratedCAConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegratedCAConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegratedCAConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegratedCAConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAsyncParams

`func (o *IntegratedCAConnectorResponse) GetAsyncParams() IntegratedCAConnectorAsyncParams`

GetAsyncParams returns the AsyncParams field if non-nil, zero value otherwise.

### GetAsyncParamsOk

`func (o *IntegratedCAConnectorResponse) GetAsyncParamsOk() (*IntegratedCAConnectorAsyncParams, bool)`

GetAsyncParamsOk returns a tuple with the AsyncParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncParams

`func (o *IntegratedCAConnectorResponse) SetAsyncParams(v IntegratedCAConnectorAsyncParams)`

SetAsyncParams sets AsyncParams field to given value.

### HasAsyncParams

`func (o *IntegratedCAConnectorResponse) HasAsyncParams() bool`

HasAsyncParams returns a boolean if a field has been set.

### GetCaCert

`func (o *IntegratedCAConnectorResponse) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *IntegratedCAConnectorResponse) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *IntegratedCAConnectorResponse) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *IntegratedCAConnectorResponse) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *IntegratedCAConnectorResponse) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *IntegratedCAConnectorResponse) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetCaKey

`func (o *IntegratedCAConnectorResponse) GetCaKey() SecretString`

GetCaKey returns the CaKey field if non-nil, zero value otherwise.

### GetCaKeyOk

`func (o *IntegratedCAConnectorResponse) GetCaKeyOk() (*SecretString, bool)`

GetCaKeyOk returns a tuple with the CaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaKey

`func (o *IntegratedCAConnectorResponse) SetCaKey(v SecretString)`

SetCaKey sets CaKey field to given value.

### HasCaKey

`func (o *IntegratedCAConnectorResponse) HasCaKey() bool`

HasCaKey returns a boolean if a field has been set.

### SetCaKeyNil

`func (o *IntegratedCAConnectorResponse) SetCaKeyNil(b bool)`

 SetCaKeyNil sets the value for CaKey to be an explicit nil

### UnsetCaKey
`func (o *IntegratedCAConnectorResponse) UnsetCaKey()`

UnsetCaKey ensures that no value is present for CaKey, not even an explicit nil
### GetCertType

`func (o *IntegratedCAConnectorResponse) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *IntegratedCAConnectorResponse) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *IntegratedCAConnectorResponse) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *IntegratedCAConnectorResponse) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### SetCertTypeNil

`func (o *IntegratedCAConnectorResponse) SetCertTypeNil(b bool)`

 SetCertTypeNil sets the value for CertType to be an explicit nil

### UnsetCertType
`func (o *IntegratedCAConnectorResponse) UnsetCertType()`

UnsetCertType ensures that no value is present for CertType, not even an explicit nil
### GetCheckPop

`func (o *IntegratedCAConnectorResponse) GetCheckPop() bool`

GetCheckPop returns the CheckPop field if non-nil, zero value otherwise.

### GetCheckPopOk

`func (o *IntegratedCAConnectorResponse) GetCheckPopOk() (*bool, bool)`

GetCheckPopOk returns a tuple with the CheckPop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPop

`func (o *IntegratedCAConnectorResponse) SetCheckPop(v bool)`

SetCheckPop sets CheckPop field to given value.

### HasCheckPop

`func (o *IntegratedCAConnectorResponse) HasCheckPop() bool`

HasCheckPop returns a boolean if a field has been set.

### SetCheckPopNil

`func (o *IntegratedCAConnectorResponse) SetCheckPopNil(b bool)`

 SetCheckPopNil sets the value for CheckPop to be an explicit nil

### UnsetCheckPop
`func (o *IntegratedCAConnectorResponse) UnsetCheckPop()`

UnsetCheckPop ensures that no value is present for CheckPop, not even an explicit nil
### GetCrlLifetime

`func (o *IntegratedCAConnectorResponse) GetCrlLifetime() string`

GetCrlLifetime returns the CrlLifetime field if non-nil, zero value otherwise.

### GetCrlLifetimeOk

`func (o *IntegratedCAConnectorResponse) GetCrlLifetimeOk() (*string, bool)`

GetCrlLifetimeOk returns a tuple with the CrlLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlLifetime

`func (o *IntegratedCAConnectorResponse) SetCrlLifetime(v string)`

SetCrlLifetime sets CrlLifetime field to given value.

### HasCrlLifetime

`func (o *IntegratedCAConnectorResponse) HasCrlLifetime() bool`

HasCrlLifetime returns a boolean if a field has been set.

### SetCrlLifetimeNil

`func (o *IntegratedCAConnectorResponse) SetCrlLifetimeNil(b bool)`

 SetCrlLifetimeNil sets the value for CrlLifetime to be an explicit nil

### UnsetCrlLifetime
`func (o *IntegratedCAConnectorResponse) UnsetCrlLifetime()`

UnsetCrlLifetime ensures that no value is present for CrlLifetime, not even an explicit nil
### GetCrlPath

`func (o *IntegratedCAConnectorResponse) GetCrlPath() string`

GetCrlPath returns the CrlPath field if non-nil, zero value otherwise.

### GetCrlPathOk

`func (o *IntegratedCAConnectorResponse) GetCrlPathOk() (*string, bool)`

GetCrlPathOk returns a tuple with the CrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlPath

`func (o *IntegratedCAConnectorResponse) SetCrlPath(v string)`

SetCrlPath sets CrlPath field to given value.

### HasCrlPath

`func (o *IntegratedCAConnectorResponse) HasCrlPath() bool`

HasCrlPath returns a boolean if a field has been set.

### SetCrlPathNil

`func (o *IntegratedCAConnectorResponse) SetCrlPathNil(b bool)`

 SetCrlPathNil sets the value for CrlPath to be an explicit nil

### UnsetCrlPath
`func (o *IntegratedCAConnectorResponse) UnsetCrlPath()`

UnsetCrlPath ensures that no value is present for CrlPath, not even an explicit nil
### GetCrtBackDate

`func (o *IntegratedCAConnectorResponse) GetCrtBackDate() string`

GetCrtBackDate returns the CrtBackDate field if non-nil, zero value otherwise.

### GetCrtBackDateOk

`func (o *IntegratedCAConnectorResponse) GetCrtBackDateOk() (*string, bool)`

GetCrtBackDateOk returns a tuple with the CrtBackDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtBackDate

`func (o *IntegratedCAConnectorResponse) SetCrtBackDate(v string)`

SetCrtBackDate sets CrtBackDate field to given value.

### HasCrtBackDate

`func (o *IntegratedCAConnectorResponse) HasCrtBackDate() bool`

HasCrtBackDate returns a boolean if a field has been set.

### SetCrtBackDateNil

`func (o *IntegratedCAConnectorResponse) SetCrtBackDateNil(b bool)`

 SetCrtBackDateNil sets the value for CrtBackDate to be an explicit nil

### UnsetCrtBackDate
`func (o *IntegratedCAConnectorResponse) UnsetCrtBackDate()`

UnsetCrtBackDate ensures that no value is present for CrtBackDate, not even an explicit nil
### GetCrtLifetime

`func (o *IntegratedCAConnectorResponse) GetCrtLifetime() string`

GetCrtLifetime returns the CrtLifetime field if non-nil, zero value otherwise.

### GetCrtLifetimeOk

`func (o *IntegratedCAConnectorResponse) GetCrtLifetimeOk() (*string, bool)`

GetCrtLifetimeOk returns a tuple with the CrtLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtLifetime

`func (o *IntegratedCAConnectorResponse) SetCrtLifetime(v string)`

SetCrtLifetime sets CrtLifetime field to given value.

### HasCrtLifetime

`func (o *IntegratedCAConnectorResponse) HasCrtLifetime() bool`

HasCrtLifetime returns a boolean if a field has been set.

### SetCrtLifetimeNil

`func (o *IntegratedCAConnectorResponse) SetCrtLifetimeNil(b bool)`

 SetCrtLifetimeNil sets the value for CrtLifetime to be an explicit nil

### UnsetCrtLifetime
`func (o *IntegratedCAConnectorResponse) UnsetCrtLifetime()`

UnsetCrtLifetime ensures that no value is present for CrtLifetime, not even an explicit nil
### GetCryptoType

`func (o *IntegratedCAConnectorResponse) GetCryptoType() string`

GetCryptoType returns the CryptoType field if non-nil, zero value otherwise.

### GetCryptoTypeOk

`func (o *IntegratedCAConnectorResponse) GetCryptoTypeOk() (*string, bool)`

GetCryptoTypeOk returns a tuple with the CryptoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoType

`func (o *IntegratedCAConnectorResponse) SetCryptoType(v string)`

SetCryptoType sets CryptoType field to given value.


### GetName

`func (o *IntegratedCAConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegratedCAConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegratedCAConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetQueue

`func (o *IntegratedCAConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *IntegratedCAConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *IntegratedCAConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *IntegratedCAConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *IntegratedCAConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *IntegratedCAConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetRetryInterval

`func (o *IntegratedCAConnectorResponse) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *IntegratedCAConnectorResponse) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *IntegratedCAConnectorResponse) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *IntegratedCAConnectorResponse) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### GetSignAlg

`func (o *IntegratedCAConnectorResponse) GetSignAlg() string`

GetSignAlg returns the SignAlg field if non-nil, zero value otherwise.

### GetSignAlgOk

`func (o *IntegratedCAConnectorResponse) GetSignAlgOk() (*string, bool)`

GetSignAlgOk returns a tuple with the SignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAlg

`func (o *IntegratedCAConnectorResponse) SetSignAlg(v string)`

SetSignAlg sets SignAlg field to given value.

### HasSignAlg

`func (o *IntegratedCAConnectorResponse) HasSignAlg() bool`

HasSignAlg returns a boolean if a field has been set.

### SetSignAlgNil

`func (o *IntegratedCAConnectorResponse) SetSignAlgNil(b bool)`

 SetSignAlgNil sets the value for SignAlg to be an explicit nil

### UnsetSignAlg
`func (o *IntegratedCAConnectorResponse) UnsetSignAlg()`

UnsetSignAlg ensures that no value is present for SignAlg, not even an explicit nil
### GetStatus

`func (o *IntegratedCAConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegratedCAConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegratedCAConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntegratedCAConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *IntegratedCAConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *IntegratedCAConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *IntegratedCAConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegratedCAConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegratedCAConnectorResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


