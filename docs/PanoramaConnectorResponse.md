# PanoramaConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the account to authenticate on Panorama | 
**Hostname** | **string** | The hostname or URL of the Panorama appliance | 
**JobRetryParameters** | [**RetryParameters**](RetryParameters.md) | Retry policy applied to the asynchronous deployment jobs run by this connector. | 
**Name** | **string** |  | 
**Prefix** | **string** | Certificate name prefix used when deploying certificates | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**SynchronizeDevices** | **bool** | Synchronize (commit) the configuration to the devices managed by the Panorama after pushing the certificate | 
**Template** | Pointer to **NullableString** | Name of the Panorama template to push certificates to | [optional] 
**TemplateStack** | Pointer to **NullableString** | Name of the Panorama template stack to push changes to | [optional] 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**Timeout** | **string** |  | 
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**Type** | **string** |  | 
**Vsys** | Pointer to **NullableString** | Virtual system name when targeting a specific VSYS within a template | [optional] 

## Methods

### NewPanoramaConnectorResponse

`func NewPanoramaConnectorResponse(id string, credentials string, hostname string, jobRetryParameters RetryParameters, name string, prefix string, synchronizeDevices bool, throttleDuration string, throttleParallelism int64, timeout string, type_ string, ) *PanoramaConnectorResponse`

NewPanoramaConnectorResponse instantiates a new PanoramaConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanoramaConnectorResponseWithDefaults

`func NewPanoramaConnectorResponseWithDefaults() *PanoramaConnectorResponse`

NewPanoramaConnectorResponseWithDefaults instantiates a new PanoramaConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PanoramaConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PanoramaConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PanoramaConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *PanoramaConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PanoramaConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PanoramaConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *PanoramaConnectorResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PanoramaConnectorResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PanoramaConnectorResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetJobRetryParameters

`func (o *PanoramaConnectorResponse) GetJobRetryParameters() RetryParameters`

GetJobRetryParameters returns the JobRetryParameters field if non-nil, zero value otherwise.

### GetJobRetryParametersOk

`func (o *PanoramaConnectorResponse) GetJobRetryParametersOk() (*RetryParameters, bool)`

GetJobRetryParametersOk returns a tuple with the JobRetryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRetryParameters

`func (o *PanoramaConnectorResponse) SetJobRetryParameters(v RetryParameters)`

SetJobRetryParameters sets JobRetryParameters field to given value.


### GetName

`func (o *PanoramaConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PanoramaConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PanoramaConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *PanoramaConnectorResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PanoramaConnectorResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PanoramaConnectorResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetProxy

`func (o *PanoramaConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PanoramaConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PanoramaConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *PanoramaConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *PanoramaConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *PanoramaConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetSynchronizeDevices

`func (o *PanoramaConnectorResponse) GetSynchronizeDevices() bool`

GetSynchronizeDevices returns the SynchronizeDevices field if non-nil, zero value otherwise.

### GetSynchronizeDevicesOk

`func (o *PanoramaConnectorResponse) GetSynchronizeDevicesOk() (*bool, bool)`

GetSynchronizeDevicesOk returns a tuple with the SynchronizeDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizeDevices

`func (o *PanoramaConnectorResponse) SetSynchronizeDevices(v bool)`

SetSynchronizeDevices sets SynchronizeDevices field to given value.


### GetTemplate

`func (o *PanoramaConnectorResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PanoramaConnectorResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PanoramaConnectorResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PanoramaConnectorResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *PanoramaConnectorResponse) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *PanoramaConnectorResponse) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetTemplateStack

`func (o *PanoramaConnectorResponse) GetTemplateStack() string`

GetTemplateStack returns the TemplateStack field if non-nil, zero value otherwise.

### GetTemplateStackOk

`func (o *PanoramaConnectorResponse) GetTemplateStackOk() (*string, bool)`

GetTemplateStackOk returns a tuple with the TemplateStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateStack

`func (o *PanoramaConnectorResponse) SetTemplateStack(v string)`

SetTemplateStack sets TemplateStack field to given value.

### HasTemplateStack

`func (o *PanoramaConnectorResponse) HasTemplateStack() bool`

HasTemplateStack returns a boolean if a field has been set.

### SetTemplateStackNil

`func (o *PanoramaConnectorResponse) SetTemplateStackNil(b bool)`

 SetTemplateStackNil sets the value for TemplateStack to be an explicit nil

### UnsetTemplateStack
`func (o *PanoramaConnectorResponse) UnsetTemplateStack()`

UnsetTemplateStack ensures that no value is present for TemplateStack, not even an explicit nil
### GetThrottleDuration

`func (o *PanoramaConnectorResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *PanoramaConnectorResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *PanoramaConnectorResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *PanoramaConnectorResponse) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *PanoramaConnectorResponse) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *PanoramaConnectorResponse) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *PanoramaConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PanoramaConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PanoramaConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTlsInsecure

`func (o *PanoramaConnectorResponse) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *PanoramaConnectorResponse) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *PanoramaConnectorResponse) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *PanoramaConnectorResponse) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *PanoramaConnectorResponse) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *PanoramaConnectorResponse) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetType

`func (o *PanoramaConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PanoramaConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PanoramaConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetVsys

`func (o *PanoramaConnectorResponse) GetVsys() string`

GetVsys returns the Vsys field if non-nil, zero value otherwise.

### GetVsysOk

`func (o *PanoramaConnectorResponse) GetVsysOk() (*string, bool)`

GetVsysOk returns a tuple with the Vsys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsys

`func (o *PanoramaConnectorResponse) SetVsys(v string)`

SetVsys sets Vsys field to given value.

### HasVsys

`func (o *PanoramaConnectorResponse) HasVsys() bool`

HasVsys returns a boolean if a field has been set.

### SetVsysNil

`func (o *PanoramaConnectorResponse) SetVsysNil(b bool)`

 SetVsysNil sets the value for Vsys to be an explicit nil

### UnsetVsys
`func (o *PanoramaConnectorResponse) UnsetVsys()`

UnsetVsys ensures that no value is present for Vsys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


