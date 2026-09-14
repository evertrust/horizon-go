# PanoramaConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) containing the account to authenticate on Panorama | 
**Hostname** | **string** | The hostname or URL of the Panorama appliance | 
**JobRetryParameters** | [**RetryParameters**](RetryParameters.md) | Retry policy applied to the asynchronous deployment jobs run by this connector. | 
**Name** | **string** |  | 
**Prefix** | **string** | Certificate name prefix used when deploying certificates | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**SynchronizeDevices** | Pointer to **NullableBool** | Synchronize (commit) the configuration to the devices managed by the Panorama after pushing the certificate | [optional] [default to false]
**Template** | Pointer to **NullableString** | Name of the Panorama template to push certificates to | [optional] 
**TemplateStack** | Pointer to **NullableString** | Name of the Panorama template stack to push changes to | [optional] 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | **int64** |  | 
**Timeout** | **string** |  | 
**TlsInsecure** | Pointer to **NullableBool** | Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended. | [optional] [default to false]
**Type** | **string** |  | 
**Vsys** | Pointer to **NullableString** | Virtual system name when targeting a specific VSYS within a template. Requires &#39;template&#39; to be set. | [optional] 

## Methods

### NewPanoramaConnector

`func NewPanoramaConnector(credentials string, hostname string, jobRetryParameters RetryParameters, name string, prefix string, throttleDuration string, throttleParallelism int64, timeout string, type_ string, ) *PanoramaConnector`

NewPanoramaConnector instantiates a new PanoramaConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanoramaConnectorWithDefaults

`func NewPanoramaConnectorWithDefaults() *PanoramaConnector`

NewPanoramaConnectorWithDefaults instantiates a new PanoramaConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *PanoramaConnector) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PanoramaConnector) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PanoramaConnector) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetHostname

`func (o *PanoramaConnector) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PanoramaConnector) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PanoramaConnector) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetJobRetryParameters

`func (o *PanoramaConnector) GetJobRetryParameters() RetryParameters`

GetJobRetryParameters returns the JobRetryParameters field if non-nil, zero value otherwise.

### GetJobRetryParametersOk

`func (o *PanoramaConnector) GetJobRetryParametersOk() (*RetryParameters, bool)`

GetJobRetryParametersOk returns a tuple with the JobRetryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRetryParameters

`func (o *PanoramaConnector) SetJobRetryParameters(v RetryParameters)`

SetJobRetryParameters sets JobRetryParameters field to given value.


### GetName

`func (o *PanoramaConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PanoramaConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PanoramaConnector) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *PanoramaConnector) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PanoramaConnector) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PanoramaConnector) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetProxy

`func (o *PanoramaConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PanoramaConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PanoramaConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *PanoramaConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *PanoramaConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *PanoramaConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetSynchronizeDevices

`func (o *PanoramaConnector) GetSynchronizeDevices() bool`

GetSynchronizeDevices returns the SynchronizeDevices field if non-nil, zero value otherwise.

### GetSynchronizeDevicesOk

`func (o *PanoramaConnector) GetSynchronizeDevicesOk() (*bool, bool)`

GetSynchronizeDevicesOk returns a tuple with the SynchronizeDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizeDevices

`func (o *PanoramaConnector) SetSynchronizeDevices(v bool)`

SetSynchronizeDevices sets SynchronizeDevices field to given value.

### HasSynchronizeDevices

`func (o *PanoramaConnector) HasSynchronizeDevices() bool`

HasSynchronizeDevices returns a boolean if a field has been set.

### SetSynchronizeDevicesNil

`func (o *PanoramaConnector) SetSynchronizeDevicesNil(b bool)`

 SetSynchronizeDevicesNil sets the value for SynchronizeDevices to be an explicit nil

### UnsetSynchronizeDevices
`func (o *PanoramaConnector) UnsetSynchronizeDevices()`

UnsetSynchronizeDevices ensures that no value is present for SynchronizeDevices, not even an explicit nil
### GetTemplate

`func (o *PanoramaConnector) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PanoramaConnector) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PanoramaConnector) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PanoramaConnector) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *PanoramaConnector) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *PanoramaConnector) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetTemplateStack

`func (o *PanoramaConnector) GetTemplateStack() string`

GetTemplateStack returns the TemplateStack field if non-nil, zero value otherwise.

### GetTemplateStackOk

`func (o *PanoramaConnector) GetTemplateStackOk() (*string, bool)`

GetTemplateStackOk returns a tuple with the TemplateStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateStack

`func (o *PanoramaConnector) SetTemplateStack(v string)`

SetTemplateStack sets TemplateStack field to given value.

### HasTemplateStack

`func (o *PanoramaConnector) HasTemplateStack() bool`

HasTemplateStack returns a boolean if a field has been set.

### SetTemplateStackNil

`func (o *PanoramaConnector) SetTemplateStackNil(b bool)`

 SetTemplateStackNil sets the value for TemplateStack to be an explicit nil

### UnsetTemplateStack
`func (o *PanoramaConnector) UnsetTemplateStack()`

UnsetTemplateStack ensures that no value is present for TemplateStack, not even an explicit nil
### GetThrottleDuration

`func (o *PanoramaConnector) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *PanoramaConnector) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *PanoramaConnector) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *PanoramaConnector) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *PanoramaConnector) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *PanoramaConnector) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.


### GetTimeout

`func (o *PanoramaConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PanoramaConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PanoramaConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetTlsInsecure

`func (o *PanoramaConnector) GetTlsInsecure() bool`

GetTlsInsecure returns the TlsInsecure field if non-nil, zero value otherwise.

### GetTlsInsecureOk

`func (o *PanoramaConnector) GetTlsInsecureOk() (*bool, bool)`

GetTlsInsecureOk returns a tuple with the TlsInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsInsecure

`func (o *PanoramaConnector) SetTlsInsecure(v bool)`

SetTlsInsecure sets TlsInsecure field to given value.

### HasTlsInsecure

`func (o *PanoramaConnector) HasTlsInsecure() bool`

HasTlsInsecure returns a boolean if a field has been set.

### SetTlsInsecureNil

`func (o *PanoramaConnector) SetTlsInsecureNil(b bool)`

 SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil

### UnsetTlsInsecure
`func (o *PanoramaConnector) UnsetTlsInsecure()`

UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
### GetType

`func (o *PanoramaConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PanoramaConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PanoramaConnector) SetType(v string)`

SetType sets Type field to given value.


### GetVsys

`func (o *PanoramaConnector) GetVsys() string`

GetVsys returns the Vsys field if non-nil, zero value otherwise.

### GetVsysOk

`func (o *PanoramaConnector) GetVsysOk() (*string, bool)`

GetVsysOk returns a tuple with the Vsys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsys

`func (o *PanoramaConnector) SetVsys(v string)`

SetVsys sets Vsys field to given value.

### HasVsys

`func (o *PanoramaConnector) HasVsys() bool`

HasVsys returns a boolean if a field has been set.

### SetVsysNil

`func (o *PanoramaConnector) SetVsysNil(b bool)`

 SetVsysNil sets the value for Vsys to be an explicit nil

### UnsetVsys
`func (o *PanoramaConnector) UnsetVsys()`

UnsetVsys ensures that no value is present for Vsys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


