# IntunePKCSConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AzureTenant** | Pointer to **string** |  | [optional] 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the App ID and Key to authenticate on Intune PKCS | 
**IntendedPurpose** | Pointer to **NullableString** |  | [optional] 
**KeyName** | **string** |  | 
**MaxStoredCertificatePerHolder** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**ProviderName** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**PubKey** | **string** |  | 
**SearchFilter** | Pointer to **NullableString** |  | [optional] 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewIntunePKCSConnectorResponse

`func NewIntunePKCSConnectorResponse(id string, credentials string, keyName string, name string, pubKey string, throttleDuration string, throttleParallelism int64, type_ string, ) *IntunePKCSConnectorResponse`

NewIntunePKCSConnectorResponse instantiates a new IntunePKCSConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntunePKCSConnectorResponseWithDefaults

`func NewIntunePKCSConnectorResponseWithDefaults() *IntunePKCSConnectorResponse`

NewIntunePKCSConnectorResponseWithDefaults instantiates a new IntunePKCSConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntunePKCSConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntunePKCSConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntunePKCSConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAzureTenant

`func (o *IntunePKCSConnectorResponse) GetAzureTenant() string`

GetAzureTenant returns the AzureTenant field if non-nil, zero value otherwise.

### GetAzureTenantOk

`func (o *IntunePKCSConnectorResponse) GetAzureTenantOk() (*string, bool)`

GetAzureTenantOk returns a tuple with the AzureTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenant

`func (o *IntunePKCSConnectorResponse) SetAzureTenant(v string)`

SetAzureTenant sets AzureTenant field to given value.

### HasAzureTenant

`func (o *IntunePKCSConnectorResponse) HasAzureTenant() bool`

HasAzureTenant returns a boolean if a field has been set.

### GetCredentials

`func (o *IntunePKCSConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *IntunePKCSConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *IntunePKCSConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetIntendedPurpose

`func (o *IntunePKCSConnectorResponse) GetIntendedPurpose() string`

GetIntendedPurpose returns the IntendedPurpose field if non-nil, zero value otherwise.

### GetIntendedPurposeOk

`func (o *IntunePKCSConnectorResponse) GetIntendedPurposeOk() (*string, bool)`

GetIntendedPurposeOk returns a tuple with the IntendedPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedPurpose

`func (o *IntunePKCSConnectorResponse) SetIntendedPurpose(v string)`

SetIntendedPurpose sets IntendedPurpose field to given value.

### HasIntendedPurpose

`func (o *IntunePKCSConnectorResponse) HasIntendedPurpose() bool`

HasIntendedPurpose returns a boolean if a field has been set.

### SetIntendedPurposeNil

`func (o *IntunePKCSConnectorResponse) SetIntendedPurposeNil(b bool)`

 SetIntendedPurposeNil sets the value for IntendedPurpose to be an explicit nil

### UnsetIntendedPurpose
`func (o *IntunePKCSConnectorResponse) UnsetIntendedPurpose()`

UnsetIntendedPurpose ensures that no value is present for IntendedPurpose, not even an explicit nil
### GetKeyName

`func (o *IntunePKCSConnectorResponse) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *IntunePKCSConnectorResponse) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *IntunePKCSConnectorResponse) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetMaxStoredCertificatePerHolder

`func (o *IntunePKCSConnectorResponse) GetMaxStoredCertificatePerHolder() int64`

GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field if non-nil, zero value otherwise.

### GetMaxStoredCertificatePerHolderOk

`func (o *IntunePKCSConnectorResponse) GetMaxStoredCertificatePerHolderOk() (*int64, bool)`

GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredCertificatePerHolder

`func (o *IntunePKCSConnectorResponse) SetMaxStoredCertificatePerHolder(v int64)`

SetMaxStoredCertificatePerHolder sets MaxStoredCertificatePerHolder field to given value.

### HasMaxStoredCertificatePerHolder

`func (o *IntunePKCSConnectorResponse) HasMaxStoredCertificatePerHolder() bool`

HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.

### SetMaxStoredCertificatePerHolderNil

`func (o *IntunePKCSConnectorResponse) SetMaxStoredCertificatePerHolderNil(b bool)`

 SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil

### UnsetMaxStoredCertificatePerHolder
`func (o *IntunePKCSConnectorResponse) UnsetMaxStoredCertificatePerHolder()`

UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
### GetName

`func (o *IntunePKCSConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntunePKCSConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntunePKCSConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProviderName

`func (o *IntunePKCSConnectorResponse) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *IntunePKCSConnectorResponse) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *IntunePKCSConnectorResponse) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *IntunePKCSConnectorResponse) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *IntunePKCSConnectorResponse) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *IntunePKCSConnectorResponse) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetProxy

`func (o *IntunePKCSConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *IntunePKCSConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *IntunePKCSConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *IntunePKCSConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *IntunePKCSConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *IntunePKCSConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetPubKey

`func (o *IntunePKCSConnectorResponse) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *IntunePKCSConnectorResponse) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *IntunePKCSConnectorResponse) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetSearchFilter

`func (o *IntunePKCSConnectorResponse) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *IntunePKCSConnectorResponse) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *IntunePKCSConnectorResponse) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *IntunePKCSConnectorResponse) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.

### SetSearchFilterNil

`func (o *IntunePKCSConnectorResponse) SetSearchFilterNil(b bool)`

 SetSearchFilterNil sets the value for SearchFilter to be an explicit nil

### UnsetSearchFilter
`func (o *IntunePKCSConnectorResponse) UnsetSearchFilter()`

UnsetSearchFilter ensures that no value is present for SearchFilter, not even an explicit nil
### GetThrottleDuration

`func (o *IntunePKCSConnectorResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *IntunePKCSConnectorResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *IntunePKCSConnectorResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *IntunePKCSConnectorResponse) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *IntunePKCSConnectorResponse) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *IntunePKCSConnectorResponse) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *IntunePKCSConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *IntunePKCSConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *IntunePKCSConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *IntunePKCSConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *IntunePKCSConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *IntunePKCSConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *IntunePKCSConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntunePKCSConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntunePKCSConnectorResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


