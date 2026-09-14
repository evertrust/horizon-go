# PanOSFirewallConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the account to authenticate on the firewall | 
**Hostname** | **string** | The hostname or URL of the PAN-OS firewall | 
**JobRetryParameters** | [**RetryParameters**](RetryParameters.md) | Retry policy applied to the asynchronous deployment jobs run by this connector. | 
**Name** | **string** |  | 
**Prefix** | **string** | Certificate name prefix used when deploying certificates | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**Timeout** | **string** |  | 
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**Type** | **string** |  | 
**Vsys** | Pointer to **NullableString** | Virtual system name for multi-VSYS firewalls | [optional] 

## Methods

### NewPanOSFirewallConnectorResponse

`func NewPanOSFirewallConnectorResponse(id string, credentials string, hostname string, jobRetryParameters RetryParameters, name string, prefix string, throttleDuration string, throttleParallelism int64, timeout string, type_ string, ) *PanOSFirewallConnectorResponse`

NewPanOSFirewallConnectorResponse instantiates a new PanOSFirewallConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanOSFirewallConnectorResponseWithDefaults

`func NewPanOSFirewallConnectorResponseWithDefaults() *PanOSFirewallConnectorResponse`

NewPanOSFirewallConnectorResponseWithDefaults instantiates a new PanOSFirewallConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PanOSFirewallConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PanOSFirewallConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PanOSFirewallConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *PanOSFirewallConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PanOSFirewallConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PanOSFirewallConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *PanOSFirewallConnectorResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PanOSFirewallConnectorResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PanOSFirewallConnectorResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetJobRetryParameters

`func (o *PanOSFirewallConnectorResponse) GetJobRetryParameters() RetryParameters`

GetJobRetryParameters returns the JobRetryParameters field if non-nil, zero value otherwise.

### GetJobRetryParametersOk

`func (o *PanOSFirewallConnectorResponse) GetJobRetryParametersOk() (*RetryParameters, bool)`

GetJobRetryParametersOk returns a tuple with the JobRetryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRetryParameters

`func (o *PanOSFirewallConnectorResponse) SetJobRetryParameters(v RetryParameters)`

SetJobRetryParameters sets JobRetryParameters field to given value.


### GetName

`func (o *PanOSFirewallConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PanOSFirewallConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PanOSFirewallConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *PanOSFirewallConnectorResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PanOSFirewallConnectorResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PanOSFirewallConnectorResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetProxy

`func (o *PanOSFirewallConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PanOSFirewallConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PanOSFirewallConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *PanOSFirewallConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *PanOSFirewallConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *PanOSFirewallConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetThrottleDuration

`func (o *PanOSFirewallConnectorResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *PanOSFirewallConnectorResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *PanOSFirewallConnectorResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *PanOSFirewallConnectorResponse) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *PanOSFirewallConnectorResponse) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *PanOSFirewallConnectorResponse) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *PanOSFirewallConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PanOSFirewallConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PanOSFirewallConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTlsInsecure

`func (o *PanOSFirewallConnectorResponse) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *PanOSFirewallConnectorResponse) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *PanOSFirewallConnectorResponse) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *PanOSFirewallConnectorResponse) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *PanOSFirewallConnectorResponse) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *PanOSFirewallConnectorResponse) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetType

`func (o *PanOSFirewallConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PanOSFirewallConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PanOSFirewallConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetVsys

`func (o *PanOSFirewallConnectorResponse) GetVsys() string`

GetVsys returns the Vsys field if non-nil, zero value otherwise.

### GetVsysOk

`func (o *PanOSFirewallConnectorResponse) GetVsysOk() (*string, bool)`

GetVsysOk returns a tuple with the Vsys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsys

`func (o *PanOSFirewallConnectorResponse) SetVsys(v string)`

SetVsys sets Vsys field to given value.

### HasVsys

`func (o *PanOSFirewallConnectorResponse) HasVsys() bool`

HasVsys returns a boolean if a field has been set.

### SetVsysNil

`func (o *PanOSFirewallConnectorResponse) SetVsysNil(b bool)`

 SetVsysNil sets the value for Vsys to be an explicit nil

### UnsetVsys
`func (o *PanOSFirewallConnectorResponse) UnsetVsys()`

UnsetVsys ensures that no value is present for Vsys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


