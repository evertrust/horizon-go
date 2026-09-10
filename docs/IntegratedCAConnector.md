# IntegratedCAConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Type** | **string** |  | 

## Methods

### NewIntegratedCAConnector

`func NewIntegratedCAConnector(cryptoType string, name string, type_ string, ) *IntegratedCAConnector`

NewIntegratedCAConnector instantiates a new IntegratedCAConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegratedCAConnectorWithDefaults

`func NewIntegratedCAConnectorWithDefaults() *IntegratedCAConnector`

NewIntegratedCAConnectorWithDefaults instantiates a new IntegratedCAConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsyncParams

`func (o *IntegratedCAConnector) GetAsyncParams() IntegratedCAConnectorAsyncParams`

GetAsyncParams returns the AsyncParams field if non-nil, zero value otherwise.

### GetAsyncParamsOk

`func (o *IntegratedCAConnector) GetAsyncParamsOk() (*IntegratedCAConnectorAsyncParams, bool)`

GetAsyncParamsOk returns a tuple with the AsyncParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncParams

`func (o *IntegratedCAConnector) SetAsyncParams(v IntegratedCAConnectorAsyncParams)`

SetAsyncParams sets AsyncParams field to given value.

### HasAsyncParams

`func (o *IntegratedCAConnector) HasAsyncParams() bool`

HasAsyncParams returns a boolean if a field has been set.

### GetCaCert

`func (o *IntegratedCAConnector) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *IntegratedCAConnector) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *IntegratedCAConnector) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *IntegratedCAConnector) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *IntegratedCAConnector) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *IntegratedCAConnector) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetCaKey

`func (o *IntegratedCAConnector) GetCaKey() SecretString`

GetCaKey returns the CaKey field if non-nil, zero value otherwise.

### GetCaKeyOk

`func (o *IntegratedCAConnector) GetCaKeyOk() (*SecretString, bool)`

GetCaKeyOk returns a tuple with the CaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaKey

`func (o *IntegratedCAConnector) SetCaKey(v SecretString)`

SetCaKey sets CaKey field to given value.

### HasCaKey

`func (o *IntegratedCAConnector) HasCaKey() bool`

HasCaKey returns a boolean if a field has been set.

### SetCaKeyNil

`func (o *IntegratedCAConnector) SetCaKeyNil(b bool)`

 SetCaKeyNil sets the value for CaKey to be an explicit nil

### UnsetCaKey
`func (o *IntegratedCAConnector) UnsetCaKey()`

UnsetCaKey ensures that no value is present for CaKey, not even an explicit nil
### GetCertType

`func (o *IntegratedCAConnector) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *IntegratedCAConnector) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *IntegratedCAConnector) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *IntegratedCAConnector) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### SetCertTypeNil

`func (o *IntegratedCAConnector) SetCertTypeNil(b bool)`

 SetCertTypeNil sets the value for CertType to be an explicit nil

### UnsetCertType
`func (o *IntegratedCAConnector) UnsetCertType()`

UnsetCertType ensures that no value is present for CertType, not even an explicit nil
### GetCheckPop

`func (o *IntegratedCAConnector) GetCheckPop() bool`

GetCheckPop returns the CheckPop field if non-nil, zero value otherwise.

### GetCheckPopOk

`func (o *IntegratedCAConnector) GetCheckPopOk() (*bool, bool)`

GetCheckPopOk returns a tuple with the CheckPop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPop

`func (o *IntegratedCAConnector) SetCheckPop(v bool)`

SetCheckPop sets CheckPop field to given value.

### HasCheckPop

`func (o *IntegratedCAConnector) HasCheckPop() bool`

HasCheckPop returns a boolean if a field has been set.

### SetCheckPopNil

`func (o *IntegratedCAConnector) SetCheckPopNil(b bool)`

 SetCheckPopNil sets the value for CheckPop to be an explicit nil

### UnsetCheckPop
`func (o *IntegratedCAConnector) UnsetCheckPop()`

UnsetCheckPop ensures that no value is present for CheckPop, not even an explicit nil
### GetCrlLifetime

`func (o *IntegratedCAConnector) GetCrlLifetime() string`

GetCrlLifetime returns the CrlLifetime field if non-nil, zero value otherwise.

### GetCrlLifetimeOk

`func (o *IntegratedCAConnector) GetCrlLifetimeOk() (*string, bool)`

GetCrlLifetimeOk returns a tuple with the CrlLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlLifetime

`func (o *IntegratedCAConnector) SetCrlLifetime(v string)`

SetCrlLifetime sets CrlLifetime field to given value.

### HasCrlLifetime

`func (o *IntegratedCAConnector) HasCrlLifetime() bool`

HasCrlLifetime returns a boolean if a field has been set.

### SetCrlLifetimeNil

`func (o *IntegratedCAConnector) SetCrlLifetimeNil(b bool)`

 SetCrlLifetimeNil sets the value for CrlLifetime to be an explicit nil

### UnsetCrlLifetime
`func (o *IntegratedCAConnector) UnsetCrlLifetime()`

UnsetCrlLifetime ensures that no value is present for CrlLifetime, not even an explicit nil
### GetCrlPath

`func (o *IntegratedCAConnector) GetCrlPath() string`

GetCrlPath returns the CrlPath field if non-nil, zero value otherwise.

### GetCrlPathOk

`func (o *IntegratedCAConnector) GetCrlPathOk() (*string, bool)`

GetCrlPathOk returns a tuple with the CrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlPath

`func (o *IntegratedCAConnector) SetCrlPath(v string)`

SetCrlPath sets CrlPath field to given value.

### HasCrlPath

`func (o *IntegratedCAConnector) HasCrlPath() bool`

HasCrlPath returns a boolean if a field has been set.

### SetCrlPathNil

`func (o *IntegratedCAConnector) SetCrlPathNil(b bool)`

 SetCrlPathNil sets the value for CrlPath to be an explicit nil

### UnsetCrlPath
`func (o *IntegratedCAConnector) UnsetCrlPath()`

UnsetCrlPath ensures that no value is present for CrlPath, not even an explicit nil
### GetCrtBackDate

`func (o *IntegratedCAConnector) GetCrtBackDate() string`

GetCrtBackDate returns the CrtBackDate field if non-nil, zero value otherwise.

### GetCrtBackDateOk

`func (o *IntegratedCAConnector) GetCrtBackDateOk() (*string, bool)`

GetCrtBackDateOk returns a tuple with the CrtBackDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtBackDate

`func (o *IntegratedCAConnector) SetCrtBackDate(v string)`

SetCrtBackDate sets CrtBackDate field to given value.

### HasCrtBackDate

`func (o *IntegratedCAConnector) HasCrtBackDate() bool`

HasCrtBackDate returns a boolean if a field has been set.

### SetCrtBackDateNil

`func (o *IntegratedCAConnector) SetCrtBackDateNil(b bool)`

 SetCrtBackDateNil sets the value for CrtBackDate to be an explicit nil

### UnsetCrtBackDate
`func (o *IntegratedCAConnector) UnsetCrtBackDate()`

UnsetCrtBackDate ensures that no value is present for CrtBackDate, not even an explicit nil
### GetCrtLifetime

`func (o *IntegratedCAConnector) GetCrtLifetime() string`

GetCrtLifetime returns the CrtLifetime field if non-nil, zero value otherwise.

### GetCrtLifetimeOk

`func (o *IntegratedCAConnector) GetCrtLifetimeOk() (*string, bool)`

GetCrtLifetimeOk returns a tuple with the CrtLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrtLifetime

`func (o *IntegratedCAConnector) SetCrtLifetime(v string)`

SetCrtLifetime sets CrtLifetime field to given value.

### HasCrtLifetime

`func (o *IntegratedCAConnector) HasCrtLifetime() bool`

HasCrtLifetime returns a boolean if a field has been set.

### SetCrtLifetimeNil

`func (o *IntegratedCAConnector) SetCrtLifetimeNil(b bool)`

 SetCrtLifetimeNil sets the value for CrtLifetime to be an explicit nil

### UnsetCrtLifetime
`func (o *IntegratedCAConnector) UnsetCrtLifetime()`

UnsetCrtLifetime ensures that no value is present for CrtLifetime, not even an explicit nil
### GetCryptoType

`func (o *IntegratedCAConnector) GetCryptoType() string`

GetCryptoType returns the CryptoType field if non-nil, zero value otherwise.

### GetCryptoTypeOk

`func (o *IntegratedCAConnector) GetCryptoTypeOk() (*string, bool)`

GetCryptoTypeOk returns a tuple with the CryptoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoType

`func (o *IntegratedCAConnector) SetCryptoType(v string)`

SetCryptoType sets CryptoType field to given value.


### GetName

`func (o *IntegratedCAConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegratedCAConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegratedCAConnector) SetName(v string)`

SetName sets Name field to given value.


### GetQueue

`func (o *IntegratedCAConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *IntegratedCAConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *IntegratedCAConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *IntegratedCAConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *IntegratedCAConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *IntegratedCAConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetRetryInterval

`func (o *IntegratedCAConnector) GetRetryInterval() string`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *IntegratedCAConnector) GetRetryIntervalOk() (*string, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *IntegratedCAConnector) SetRetryInterval(v string)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *IntegratedCAConnector) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### GetSignAlg

`func (o *IntegratedCAConnector) GetSignAlg() string`

GetSignAlg returns the SignAlg field if non-nil, zero value otherwise.

### GetSignAlgOk

`func (o *IntegratedCAConnector) GetSignAlgOk() (*string, bool)`

GetSignAlgOk returns a tuple with the SignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAlg

`func (o *IntegratedCAConnector) SetSignAlg(v string)`

SetSignAlg sets SignAlg field to given value.

### HasSignAlg

`func (o *IntegratedCAConnector) HasSignAlg() bool`

HasSignAlg returns a boolean if a field has been set.

### SetSignAlgNil

`func (o *IntegratedCAConnector) SetSignAlgNil(b bool)`

 SetSignAlgNil sets the value for SignAlg to be an explicit nil

### UnsetSignAlg
`func (o *IntegratedCAConnector) UnsetSignAlg()`

UnsetSignAlg ensures that no value is present for SignAlg, not even an explicit nil
### GetType

`func (o *IntegratedCAConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegratedCAConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegratedCAConnector) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


