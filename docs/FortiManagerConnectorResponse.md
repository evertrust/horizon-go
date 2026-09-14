# FortiManagerConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the account to authenticate on FortiManager | 
**Hostname** | **string** | The hostname or URL of the FortiManager appliance | 
**JobRetryParameters** | [**RetryParameters**](RetryParameters.md) | Retry policy applied to the asynchronous deployment jobs run by this connector. | 
**ManagedDevice** | Pointer to [**NullableFortiManagerConnectorResponseManagedDevice**](FortiManagerConnectorResponseManagedDevice.md) |  | [optional] 
**Name** | **string** |  | 
**Prefix** | **string** | Certificate name prefix used when deploying certificates | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Target** | **string** | Selects what the connector deploys certificates to. &#x60;unit&#x60; targets the FortiManager unit&#39;s own certificate store, &#x60;device&#x60; targets a FortiGate device managed by the FortiManager. When set to &#x60;device&#x60;, &#x60;managedDevice&#x60; is provided; when set to &#x60;unit&#x60;, &#x60;managedDevice&#x60; is omitted. | 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**Timeout** | **string** |  | 
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**Type** | **string** |  | 

## Methods

### NewFortiManagerConnectorResponse

`func NewFortiManagerConnectorResponse(id string, credentials string, hostname string, jobRetryParameters RetryParameters, name string, prefix string, target string, throttleDuration string, throttleParallelism int64, timeout string, type_ string, ) *FortiManagerConnectorResponse`

NewFortiManagerConnectorResponse instantiates a new FortiManagerConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFortiManagerConnectorResponseWithDefaults

`func NewFortiManagerConnectorResponseWithDefaults() *FortiManagerConnectorResponse`

NewFortiManagerConnectorResponseWithDefaults instantiates a new FortiManagerConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FortiManagerConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FortiManagerConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FortiManagerConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *FortiManagerConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *FortiManagerConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *FortiManagerConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *FortiManagerConnectorResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FortiManagerConnectorResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FortiManagerConnectorResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetJobRetryParameters

`func (o *FortiManagerConnectorResponse) GetJobRetryParameters() RetryParameters`

GetJobRetryParameters returns the JobRetryParameters field if non-nil, zero value otherwise.

### GetJobRetryParametersOk

`func (o *FortiManagerConnectorResponse) GetJobRetryParametersOk() (*RetryParameters, bool)`

GetJobRetryParametersOk returns a tuple with the JobRetryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRetryParameters

`func (o *FortiManagerConnectorResponse) SetJobRetryParameters(v RetryParameters)`

SetJobRetryParameters sets JobRetryParameters field to given value.


### GetManagedDevice

`func (o *FortiManagerConnectorResponse) GetManagedDevice() FortiManagerConnectorResponseManagedDevice`

GetManagedDevice returns the ManagedDevice field if non-nil, zero value otherwise.

### GetManagedDeviceOk

`func (o *FortiManagerConnectorResponse) GetManagedDeviceOk() (*FortiManagerConnectorResponseManagedDevice, bool)`

GetManagedDeviceOk returns a tuple with the ManagedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevice

`func (o *FortiManagerConnectorResponse) SetManagedDevice(v FortiManagerConnectorResponseManagedDevice)`

SetManagedDevice sets ManagedDevice field to given value.

### HasManagedDevice

`func (o *FortiManagerConnectorResponse) HasManagedDevice() bool`

HasManagedDevice returns a boolean if a field has been set.

### SetManagedDeviceNil

`func (o *FortiManagerConnectorResponse) SetManagedDeviceNil(b bool)`

 SetManagedDeviceNil sets the value for ManagedDevice to be an explicit nil

### UnsetManagedDevice
`func (o *FortiManagerConnectorResponse) UnsetManagedDevice()`

UnsetManagedDevice ensures that no value is present for ManagedDevice, not even an explicit nil
### GetName

`func (o *FortiManagerConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FortiManagerConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FortiManagerConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *FortiManagerConnectorResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *FortiManagerConnectorResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *FortiManagerConnectorResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetProxy

`func (o *FortiManagerConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *FortiManagerConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *FortiManagerConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *FortiManagerConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *FortiManagerConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *FortiManagerConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTarget

`func (o *FortiManagerConnectorResponse) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FortiManagerConnectorResponse) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FortiManagerConnectorResponse) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetThrottleDuration

`func (o *FortiManagerConnectorResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *FortiManagerConnectorResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *FortiManagerConnectorResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *FortiManagerConnectorResponse) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *FortiManagerConnectorResponse) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *FortiManagerConnectorResponse) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *FortiManagerConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FortiManagerConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FortiManagerConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTlsInsecure

`func (o *FortiManagerConnectorResponse) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *FortiManagerConnectorResponse) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *FortiManagerConnectorResponse) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *FortiManagerConnectorResponse) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *FortiManagerConnectorResponse) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *FortiManagerConnectorResponse) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetType

`func (o *FortiManagerConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FortiManagerConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FortiManagerConnectorResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


